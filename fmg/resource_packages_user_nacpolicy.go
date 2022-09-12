// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NAC policy matching pattern to identify matching NAC devices.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesUserNacPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesUserNacPolicyCreate,
		Read:   resourcePackagesUserNacPolicyRead,
		Update: resourcePackagesUserNacPolicyUpdate,
		Delete: resourcePackagesUserNacPolicyDelete,

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
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ems_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssid_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sw_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePackagesUserNacPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesUserNacPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesUserNacPolicy resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesUserNacPolicy(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesUserNacPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesUserNacPolicyRead(d, m)
}

func resourcePackagesUserNacPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesUserNacPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesUserNacPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesUserNacPolicy(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesUserNacPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesUserNacPolicyRead(d, m)
}

func resourcePackagesUserNacPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	err = c.DeletePackagesUserNacPolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesUserNacPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesUserNacPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	var paralist []string
	paralist = append(paralist, pkg)

	o, err := c.ReadPackagesUserNacPolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesUserNacPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesUserNacPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesUserNacPolicy resource from API: %v", err)
	}
	return nil
}

func flattenPackagesUserNacPolicyCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyEmsTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyHwVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyHwVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicySrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicySsidPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicySwVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesUserNacPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("category", flattenPackagesUserNacPolicyCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "PackagesUserNacPolicy-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("description", flattenPackagesUserNacPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "PackagesUserNacPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("ems_tag", flattenPackagesUserNacPolicyEmsTag(o["ems-tag"], d, "ems_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-tag"], "PackagesUserNacPolicy-EmsTag"); ok {
			if err = d.Set("ems_tag", vv); err != nil {
				return fmt.Errorf("Error reading ems_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_tag: %v", err)
		}
	}

	if err = d.Set("family", flattenPackagesUserNacPolicyFamily(o["family"], d, "family")); err != nil {
		if vv, ok := fortiAPIPatch(o["family"], "PackagesUserNacPolicy-Family"); ok {
			if err = d.Set("family", vv); err != nil {
				return fmt.Errorf("Error reading family: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading family: %v", err)
		}
	}

	if err = d.Set("host", flattenPackagesUserNacPolicyHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "PackagesUserNacPolicy-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("hw_vendor", flattenPackagesUserNacPolicyHwVendor(o["hw-vendor"], d, "hw_vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-vendor"], "PackagesUserNacPolicy-HwVendor"); ok {
			if err = d.Set("hw_vendor", vv); err != nil {
				return fmt.Errorf("Error reading hw_vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_vendor: %v", err)
		}
	}

	if err = d.Set("hw_version", flattenPackagesUserNacPolicyHwVersion(o["hw-version"], d, "hw_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-version"], "PackagesUserNacPolicy-HwVersion"); ok {
			if err = d.Set("hw_version", vv); err != nil {
				return fmt.Errorf("Error reading hw_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_version: %v", err)
		}
	}

	if err = d.Set("mac", flattenPackagesUserNacPolicyMac(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "PackagesUserNacPolicy-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesUserNacPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesUserNacPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("os", flattenPackagesUserNacPolicyOs(o["os"], d, "os")); err != nil {
		if vv, ok := fortiAPIPatch(o["os"], "PackagesUserNacPolicy-Os"); ok {
			if err = d.Set("os", vv); err != nil {
				return fmt.Errorf("Error reading os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("src", flattenPackagesUserNacPolicySrc(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "PackagesUserNacPolicy-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("ssid_policy", flattenPackagesUserNacPolicySsidPolicy(o["ssid-policy"], d, "ssid_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssid-policy"], "PackagesUserNacPolicy-SsidPolicy"); ok {
			if err = d.Set("ssid_policy", vv); err != nil {
				return fmt.Errorf("Error reading ssid_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssid_policy: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesUserNacPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesUserNacPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("sw_version", flattenPackagesUserNacPolicySwVersion(o["sw-version"], d, "sw_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["sw-version"], "PackagesUserNacPolicy-SwVersion"); ok {
			if err = d.Set("sw_version", vv); err != nil {
				return fmt.Errorf("Error reading sw_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sw_version: %v", err)
		}
	}

	if err = d.Set("type", flattenPackagesUserNacPolicyType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "PackagesUserNacPolicy-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user", flattenPackagesUserNacPolicyUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "PackagesUserNacPolicy-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", flattenPackagesUserNacPolicyUserGroup(o["user-group"], d, "user_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-group"], "PackagesUserNacPolicy-UserGroup"); ok {
			if err = d.Set("user_group", vv); err != nil {
				return fmt.Errorf("Error reading user_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	return nil
}

func flattenPackagesUserNacPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesUserNacPolicyCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyEmsTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyHwVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyHwVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicySrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicySsidPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicySwVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesUserNacPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandPackagesUserNacPolicyCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandPackagesUserNacPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("ems_tag"); ok || d.HasChange("ems_tag") {
		t, err := expandPackagesUserNacPolicyEmsTag(d, v, "ems_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-tag"] = t
		}
	}

	if v, ok := d.GetOk("family"); ok || d.HasChange("family") {
		t, err := expandPackagesUserNacPolicyFamily(d, v, "family")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["family"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandPackagesUserNacPolicyHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("hw_vendor"); ok || d.HasChange("hw_vendor") {
		t, err := expandPackagesUserNacPolicyHwVendor(d, v, "hw_vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-vendor"] = t
		}
	}

	if v, ok := d.GetOk("hw_version"); ok || d.HasChange("hw_version") {
		t, err := expandPackagesUserNacPolicyHwVersion(d, v, "hw_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-version"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandPackagesUserNacPolicyMac(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesUserNacPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok || d.HasChange("os") {
		t, err := expandPackagesUserNacPolicyOs(d, v, "os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandPackagesUserNacPolicySrc(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("ssid_policy"); ok || d.HasChange("ssid_policy") {
		t, err := expandPackagesUserNacPolicySsidPolicy(d, v, "ssid_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid-policy"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesUserNacPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("sw_version"); ok || d.HasChange("sw_version") {
		t, err := expandPackagesUserNacPolicySwVersion(d, v, "sw_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-version"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandPackagesUserNacPolicyType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandPackagesUserNacPolicyUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok || d.HasChange("user_group") {
		t, err := expandPackagesUserNacPolicyUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	return &obj, nil
}
