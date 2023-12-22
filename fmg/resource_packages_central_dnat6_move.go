// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Packages CentralDnat6Move

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePackagesCentralDnat6Move() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesCentralDnat6MoveUpdate,
		Read:   resourcePackagesCentralDnat6MoveRead,
		Update: resourcePackagesCentralDnat6MoveUpdate,
		Delete: resourcePackagesCentralDnat6MoveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"state_pos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"pkg_folder_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dnat6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePackagesCentralDnat6MoveUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	dnat6 := d.Get("dnat6").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["dnat6"] = dnat6

	target := d.Get("target").(string)
	obj, err := getObjectPackagesCentralDnat6Move(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesCentralDnat6Move resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesCentralDnat6Move(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating PackagesCentralDnat6Move resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("PackagesCentralDnat6Move" + "_" + pkg_folder_path + "_" + pkg + "_" + dnat6 + "_" + target)

	return resourcePackagesCentralDnat6MoveRead(d, m)
}

func resourcePackagesCentralDnat6MoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourcePackagesCentralDnat6MoveRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	sid := d.Get("dnat6").(string)
	did := d.Get("target").(string)
	action := d.Get("option").(string)

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	if pkg_folder_path == "" {
		pkg_folder_path = importOptionChecking(m.(*FortiClient).Cfg, "pkg_folder_path")
	}
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if pkg == "" {
			return fmt.Errorf("Parameter pkg is missing")
		}
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	o, err := c.ReadPackagesCentralDnat6Move(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading PackagesCentralDnat6Move resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		for _, z := range o {
			if v, ok := z.(map[string]interface{}); ok {
				if _, ok := v["name"]; !ok {
					return fmt.Errorf("Error reading PackagesCentralDnat6Move resource: name doesn't exist.")
				}

				vid := fmt.Sprintf("%v", v["name"])

				if vid == sid {
					now_sid = i
				}

				if vid == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading PackagesCentralDnat6Move resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "name(" + sid + ") and target(" + did + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "name(" + sid + ") was deleted"
			} else if now_did == -1 {
				state_pos = "target(" + did + ") was deleted"
			}
		} else {
			bconsistent := true
			if action == "before" {
				if now_sid != now_did-1 {
					bconsistent = false
				}
			}

			if action == "after" {
				if now_sid != now_did+1 {
					bconsistent = false
				}
			}

			if bconsistent == false {
				relative_pos := now_sid - now_did

				if relative_pos > 0 {
					state_pos = "name(" + sid + ") is " + strconv.Itoa(relative_pos) + " behind target(" + did + ")"
				} else {
					state_pos = "name(" + sid + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + did + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenPackagesCentralDnat6MoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesCentralDnat6MoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesCentralDnat6MoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesCentralDnat6Move(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandPackagesCentralDnat6MoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandPackagesCentralDnat6MoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
