// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Fabric policy related attributes.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePackagesFirewallPolicyFabricPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallPolicyFabricPolicyUpdate,
		Read:   resourcePackagesFirewallPolicyFabricPolicyRead,
		Update: resourcePackagesFirewallPolicyFabricPolicyUpdate,
		Delete: resourcePackagesFirewallPolicyFabricPolicyDelete,

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
			"policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"from": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"to": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePackagesFirewallPolicyFabricPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	policy := d.Get("policy").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["policy"] = policy

	obj, err := getObjectPackagesFirewallPolicyFabricPolicy(d, false)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallPolicyFabricPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdatePackagesFirewallPolicyFabricPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallPolicyFabricPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("PackagesFirewallPolicyFabricPolicy")

	return resourcePackagesFirewallPolicyFabricPolicyRead(d, m)
}

func resourcePackagesFirewallPolicyFabricPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	policy := d.Get("policy").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["policy"] = policy

	obj, err := getObjectPackagesFirewallPolicyFabricPolicy(d, true)

	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallPolicyFabricPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdatePackagesFirewallPolicyFabricPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing PackagesFirewallPolicyFabricPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallPolicyFabricPolicyRead(d *schema.ResourceData, m interface{}) error {
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
	policy := d.Get("policy").(string)
	if pkg_folder_path == "" {
		pkg_folder_path = importOptionChecking(m.(*FortiClient).Cfg, "pkg_folder_path")
		if err = d.Set("pkg_folder_path", pkg_folder_path); err != nil {
			return fmt.Errorf("Error set params pkg_folder_path: %v", err)
		}
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
	if policy == "" {
		policy = importOptionChecking(m.(*FortiClient).Cfg, "policy")
		if policy == "" {
			return fmt.Errorf("Parameter policy is missing")
		}
		if err = d.Set("policy", policy); err != nil {
			return fmt.Errorf("Error set params policy: %v", err)
		}
	}
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["policy"] = policy

	o, err := c.ReadPackagesFirewallPolicyFabricPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading PackagesFirewallPolicyFabricPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallPolicyFabricPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallPolicyFabricPolicy resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallPolicyFabricPolicyFrom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyFabricPolicyTo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectPackagesFirewallPolicyFabricPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("from", flattenPackagesFirewallPolicyFabricPolicyFrom2edl(o["from"], d, "from")); err != nil {
		if vv, ok := fortiAPIPatch(o["from"], "PackagesFirewallPolicyFabricPolicy-From"); ok {
			if err = d.Set("from", vv); err != nil {
				return fmt.Errorf("Error reading from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading from: %v", err)
		}
	}

	if err = d.Set("to", flattenPackagesFirewallPolicyFabricPolicyTo2edl(o["to"], d, "to")); err != nil {
		if vv, ok := fortiAPIPatch(o["to"], "PackagesFirewallPolicyFabricPolicy-To"); ok {
			if err = d.Set("to", vv); err != nil {
				return fmt.Errorf("Error reading to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading to: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallPolicyFabricPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallPolicyFabricPolicyFrom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyFabricPolicyTo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectPackagesFirewallPolicyFabricPolicy(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("from"); ok || d.HasChange("from") {
		t, err := expandPackagesFirewallPolicyFabricPolicyFrom2edl(d, v, "from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["from"] = t
		}
	}

	if v, ok := d.GetOk("to"); ok || d.HasChange("to") {
		t, err := expandPackagesFirewallPolicyFabricPolicyTo2edl(d, v, "to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["to"] = t
		}
	}

	return &obj, nil
}
