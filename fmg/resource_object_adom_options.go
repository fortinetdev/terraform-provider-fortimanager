// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectAdom Options

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectAdomOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectAdomOptionsUpdate,
		Read:   resourceObjectAdomOptionsRead,
		Update: resourceObjectAdomOptionsUpdate,
		Delete: resourceObjectAdomOptionsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"assign_excluded": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"specify_assign_pkg_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectAdomOptionsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectAdomOptions(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAdomOptions resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectAdomOptions(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAdomOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectAdomOptions")

	return resourceObjectAdomOptionsRead(d, m)
}

func resourceObjectAdomOptionsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectAdomOptions(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectAdomOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectAdomOptionsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectAdomOptions(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAdomOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectAdomOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAdomOptions resource from API: %v", err)
	}
	return nil
}

func flattenObjectAdomOptionsAssignExcluded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAdomOptionsSpecifyAssignPkgList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectAdomOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("assign_excluded", flattenObjectAdomOptionsAssignExcluded(o["assign_excluded"], d, "assign_excluded")); err != nil {
		if vv, ok := fortiAPIPatch(o["assign_excluded"], "ObjectAdomOptions-AssignExcluded"); ok {
			if err = d.Set("assign_excluded", vv); err != nil {
				return fmt.Errorf("Error reading assign_excluded: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assign_excluded: %v", err)
		}
	}

	if err = d.Set("specify_assign_pkg_list", flattenObjectAdomOptionsSpecifyAssignPkgList(o["specify_assign_pkg_list"], d, "specify_assign_pkg_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["specify_assign_pkg_list"], "ObjectAdomOptions-SpecifyAssignPkgList"); ok {
			if err = d.Set("specify_assign_pkg_list", vv); err != nil {
				return fmt.Errorf("Error reading specify_assign_pkg_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading specify_assign_pkg_list: %v", err)
		}
	}

	return nil
}

func flattenObjectAdomOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectAdomOptionsAssignExcluded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAdomOptionsSpecifyAssignPkgList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectAdomOptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("assign_excluded"); ok || d.HasChange("assign_excluded") {
		t, err := expandObjectAdomOptionsAssignExcluded(d, v, "assign_excluded")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign_excluded"] = t
		}
	}

	if v, ok := d.GetOk("specify_assign_pkg_list"); ok || d.HasChange("specify_assign_pkg_list") {
		t, err := expandObjectAdomOptionsSpecifyAssignPkgList(d, v, "specify_assign_pkg_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["specify_assign_pkg_list"] = t
		}
	}

	return &obj, nil
}
