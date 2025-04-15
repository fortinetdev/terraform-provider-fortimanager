// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Packages Pblock

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePackagesPblock() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesPblockCreate,
		Read:   resourcePackagesPblockRead,
		Update: resourcePackagesPblockUpdate,
		Delete: resourcePackagesPblockDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"oid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"packagesettings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"central_nat": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"consolidated_firewall_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fwpolicy_implicit_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fwpolicy6_implicit_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"inspection_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ngfw_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"policy_offload_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_ssh_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePackagesPblockCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesPblock(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesPblock resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreatePackagesPblock(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating PackagesPblock resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesPblockRead(d, m)
}

func resourcePackagesPblockUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesPblock(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesPblock resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdatePackagesPblock(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating PackagesPblock resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesPblockRead(d, m)
}

func resourcePackagesPblockDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeletePackagesPblock(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesPblock resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesPblockRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadPackagesPblock(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading PackagesPblock resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesPblock(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesPblock resource from API: %v", err)
	}
	return nil
}

func flattenPackagesPblockName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPblockOid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPblockPackageSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "central_nat"
	if _, ok := i["central-nat"]; ok {
		result["central_nat"] = flattenPackagesPblockPackageSettingsCentralNat(i["central-nat"], d, pre_append)
	}

	pre_append = pre + ".0." + "consolidated_firewall_mode"
	if _, ok := i["consolidated-firewall-mode"]; ok {
		result["consolidated_firewall_mode"] = flattenPackagesPblockPackageSettingsConsolidatedFirewallMode(i["consolidated-firewall-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwpolicy_implicit_log"
	if _, ok := i["fwpolicy-implicit-log"]; ok {
		result["fwpolicy_implicit_log"] = flattenPackagesPblockPackageSettingsFwpolicyImplicitLog(i["fwpolicy-implicit-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwpolicy6_implicit_log"
	if _, ok := i["fwpolicy6-implicit-log"]; ok {
		result["fwpolicy6_implicit_log"] = flattenPackagesPblockPackageSettingsFwpolicy6ImplicitLog(i["fwpolicy6-implicit-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspection_mode"
	if _, ok := i["inspection-mode"]; ok {
		result["inspection_mode"] = flattenPackagesPblockPackageSettingsInspectionMode(i["inspection-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ngfw_mode"
	if _, ok := i["ngfw-mode"]; ok {
		result["ngfw_mode"] = flattenPackagesPblockPackageSettingsNgfwMode(i["ngfw-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "policy_offload_level"
	if _, ok := i["policy-offload-level"]; ok {
		result["policy_offload_level"] = flattenPackagesPblockPackageSettingsPolicyOffloadLevel(i["policy-offload-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_ssh_profile"
	if _, ok := i["ssl-ssh-profile"]; ok {
		result["ssl_ssh_profile"] = flattenPackagesPblockPackageSettingsSslSshProfile(i["ssl-ssh-profile"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenPackagesPblockPackageSettingsCentralNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPblockPackageSettingsConsolidatedFirewallMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPblockPackageSettingsFwpolicyImplicitLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPblockPackageSettingsFwpolicy6ImplicitLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPblockPackageSettingsInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPblockPackageSettingsNgfwMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPblockPackageSettingsPolicyOffloadLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPblockPackageSettingsSslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenPackagesPblockType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesPblock(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenPackagesPblockName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesPblock-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("oid", flattenPackagesPblockOid(o["oid"], d, "oid")); err != nil {
		if vv, ok := fortiAPIPatch(o["oid"], "PackagesPblock-Oid"); ok {
			if err = d.Set("oid", vv); err != nil {
				return fmt.Errorf("Error reading oid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("packagesettings", flattenPackagesPblockPackageSettings(o["package settings"], d, "packagesettings")); err != nil {
			if vv, ok := fortiAPIPatch(o["package settings"], "PackagesPblock-PackageSettings"); ok {
				if err = d.Set("packagesettings", vv); err != nil {
					return fmt.Errorf("Error reading packagesettings: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading packagesettings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("packagesettings"); ok {
			if err = d.Set("packagesettings", flattenPackagesPblockPackageSettings(o["package settings"], d, "packagesettings")); err != nil {
				if vv, ok := fortiAPIPatch(o["package settings"], "PackagesPblock-PackageSettings"); ok {
					if err = d.Set("packagesettings", vv); err != nil {
						return fmt.Errorf("Error reading packagesettings: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading packagesettings: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenPackagesPblockType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "PackagesPblock-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenPackagesPblockFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesPblockName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPblockOid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPblockPackageSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "central_nat"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["central-nat"], _ = expandPackagesPblockPackageSettingsCentralNat(d, i["central_nat"], pre_append)
	}
	pre_append = pre + ".0." + "consolidated_firewall_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["consolidated-firewall-mode"], _ = expandPackagesPblockPackageSettingsConsolidatedFirewallMode(d, i["consolidated_firewall_mode"], pre_append)
	}
	pre_append = pre + ".0." + "fwpolicy_implicit_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwpolicy-implicit-log"], _ = expandPackagesPblockPackageSettingsFwpolicyImplicitLog(d, i["fwpolicy_implicit_log"], pre_append)
	}
	pre_append = pre + ".0." + "fwpolicy6_implicit_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwpolicy6-implicit-log"], _ = expandPackagesPblockPackageSettingsFwpolicy6ImplicitLog(d, i["fwpolicy6_implicit_log"], pre_append)
	}
	pre_append = pre + ".0." + "inspection_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspection-mode"], _ = expandPackagesPblockPackageSettingsInspectionMode(d, i["inspection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ngfw_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ngfw-mode"], _ = expandPackagesPblockPackageSettingsNgfwMode(d, i["ngfw_mode"], pre_append)
	}
	pre_append = pre + ".0." + "policy_offload_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["policy-offload-level"], _ = expandPackagesPblockPackageSettingsPolicyOffloadLevel(d, i["policy_offload_level"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_ssh_profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssl-ssh-profile"], _ = expandPackagesPblockPackageSettingsSslSshProfile(d, i["ssl_ssh_profile"], pre_append)
	}

	return result, nil
}

func expandPackagesPblockPackageSettingsCentralNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPblockPackageSettingsConsolidatedFirewallMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPblockPackageSettingsFwpolicyImplicitLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPblockPackageSettingsFwpolicy6ImplicitLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPblockPackageSettingsInspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPblockPackageSettingsNgfwMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPblockPackageSettingsPolicyOffloadLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPblockPackageSettingsSslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandPackagesPblockType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesPblock(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesPblockName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("oid"); ok || d.HasChange("oid") {
		t, err := expandPackagesPblockOid(d, v, "oid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oid"] = t
		}
	}

	if v, ok := d.GetOk("packagesettings"); ok || d.HasChange("packagesettings") {
		t, err := expandPackagesPblockPackageSettings(d, v, "packagesettings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["package settings"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandPackagesPblockType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
