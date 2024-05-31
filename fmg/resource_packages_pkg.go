// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Packages Pkg

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePackagesPkg() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesPkgCreate,
		Read:   resourcePackagesPkgRead,
		Update: resourcePackagesPkgUpdate,
		Delete: resourcePackagesPkgDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"objver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
							Computed: true,
						},
						"consolidated_firewall_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fwpolicy_implicit_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fwpolicy6_implicit_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspection_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ngfw_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policy_offload_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_ssh_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"packagesetting": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"central_nat": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"consolidated_firewall_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fwpolicy_implicit_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fwpolicy6_implicit_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspection_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ngfw_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_ssh_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"scopemember": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"output_folder_path": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"output_pkg_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePackagesPkgCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)

	obj, err := getObjectPackagesPkg(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesPkg resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesPkg(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating PackagesPkg resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesPkgRead(d, m)
}

func resourcePackagesPkgUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)

	obj, err := getObjectPackagesPkg(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesPkg resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesPkg(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating PackagesPkg resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesPkgRead(d, m)
}

func resourcePackagesPkgDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)

	err = c.DeletePackagesPkg(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesPkg resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesPkgRead(d *schema.ResourceData, m interface{}) error {
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
	if pkg_folder_path == "" {
		pkg_folder_path = importOptionChecking(m.(*FortiClient).Cfg, "pkg_folder_path")
	}
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)

	o, err := c.ReadPackagesPkg(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading PackagesPkg resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesPkg(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesPkg resource from API: %v", err)
	}

	pkg_type := d.Get("type").(string)
	if pkg_type == "folder" {
		d.Set("output_folder_path", formatPath(fmt.Sprintf("%s/%s", pkg_folder_path, d.Get("name").(string))))
		d.Set("output_pkg_name", "")
	} else if pkg_type == "pkg" {
		d.Set("output_folder_path", formatPath(pkg_folder_path))
		d.Set("output_pkg_name", d.Get("name").(string))
	} else {
		d.Set("output_folder_path", "")
		d.Set("output_pkg_name", "")
	}

	return nil
}

func flattenPackagesPkgName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgObjVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgOid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "central_nat"
	if _, ok := i["central-nat"]; ok {
		result["central_nat"] = flattenPackagesPkgPackageSettingsCentralNat(i["central-nat"], d, pre_append)
	}

	pre_append = pre + ".0." + "consolidated_firewall_mode"
	if _, ok := i["consolidated-firewall-mode"]; ok {
		result["consolidated_firewall_mode"] = flattenPackagesPkgPackageSettingsConsolidatedFirewallMode(i["consolidated-firewall-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwpolicy_implicit_log"
	if _, ok := i["fwpolicy-implicit-log"]; ok {
		result["fwpolicy_implicit_log"] = flattenPackagesPkgPackageSettingsFwpolicyImplicitLog(i["fwpolicy-implicit-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwpolicy6_implicit_log"
	if _, ok := i["fwpolicy6-implicit-log"]; ok {
		result["fwpolicy6_implicit_log"] = flattenPackagesPkgPackageSettingsFwpolicy6ImplicitLog(i["fwpolicy6-implicit-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspection_mode"
	if _, ok := i["inspection-mode"]; ok {
		result["inspection_mode"] = flattenPackagesPkgPackageSettingsInspectionMode(i["inspection-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ngfw_mode"
	if _, ok := i["ngfw-mode"]; ok {
		result["ngfw_mode"] = flattenPackagesPkgPackageSettingsNgfwMode(i["ngfw-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "policy_offload_level"
	if _, ok := i["policy-offload-level"]; ok {
		result["policy_offload_level"] = flattenPackagesPkgPackageSettingsPolicyOffloadLevel(i["policy-offload-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_ssh_profile"
	if _, ok := i["ssl-ssh-profile"]; ok {
		result["ssl_ssh_profile"] = flattenPackagesPkgPackageSettingsSslSshProfile(i["ssl-ssh-profile"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenPackagesPkgPackageSettingsCentralNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingsConsolidatedFirewallMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingsFwpolicyImplicitLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingsFwpolicy6ImplicitLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingsInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingsNgfwMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingsPolicyOffloadLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingsSslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenPackagesPkgPackageSetting(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "central_nat"
	if _, ok := i["central-nat"]; ok {
		result["central_nat"] = flattenPackagesPkgPackageSettingCentralNat(i["central-nat"], d, pre_append)
	}

	pre_append = pre + ".0." + "consolidated_firewall_mode"
	if _, ok := i["consolidated-firewall-mode"]; ok {
		result["consolidated_firewall_mode"] = flattenPackagesPkgPackageSettingConsolidatedFirewallMode(i["consolidated-firewall-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwpolicy_implicit_log"
	if _, ok := i["fwpolicy-implicit-log"]; ok {
		result["fwpolicy_implicit_log"] = flattenPackagesPkgPackageSettingFwpolicyImplicitLog(i["fwpolicy-implicit-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwpolicy6_implicit_log"
	if _, ok := i["fwpolicy6-implicit-log"]; ok {
		result["fwpolicy6_implicit_log"] = flattenPackagesPkgPackageSettingFwpolicy6ImplicitLog(i["fwpolicy6-implicit-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspection_mode"
	if _, ok := i["inspection-mode"]; ok {
		result["inspection_mode"] = flattenPackagesPkgPackageSettingInspectionMode(i["inspection-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ngfw_mode"
	if _, ok := i["ngfw-mode"]; ok {
		result["ngfw_mode"] = flattenPackagesPkgPackageSettingNgfwMode(i["ngfw-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_ssh_profile"
	if _, ok := i["ssl-ssh-profile"]; ok {
		result["ssl_ssh_profile"] = flattenPackagesPkgPackageSettingSslSshProfile(i["ssl-ssh-profile"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenPackagesPkgPackageSettingCentralNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingConsolidatedFirewallMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingFwpolicyImplicitLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingFwpolicy6ImplicitLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingNgfwMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgPackageSettingSslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgScopeMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenPackagesPkgScopeMemberName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "PackagesPkg-ScopeMember-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenPackagesPkgScopeMemberVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "PackagesPkg-ScopeMember-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenPackagesPkgScopeMemberName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgScopeMemberVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesPkg(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenPackagesPkgName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesPkg-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("objver", flattenPackagesPkgObjVer(o["obj ver"], d, "objver")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj ver"], "PackagesPkg-ObjVer"); ok {
			if err = d.Set("objver", vv); err != nil {
				return fmt.Errorf("Error reading objver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading objver: %v", err)
		}
	}

	if err = d.Set("oid", flattenPackagesPkgOid(o["oid"], d, "oid")); err != nil {
		if vv, ok := fortiAPIPatch(o["oid"], "PackagesPkg-Oid"); ok {
			if err = d.Set("oid", vv); err != nil {
				return fmt.Errorf("Error reading oid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("packagesettings", flattenPackagesPkgPackageSettings(o["package settings"], d, "packagesettings")); err != nil {
			if vv, ok := fortiAPIPatch(o["package settings"], "PackagesPkg-PackageSettings"); ok {
				if err = d.Set("packagesettings", vv); err != nil {
					return fmt.Errorf("Error reading packagesettings: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading packagesettings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("packagesettings"); ok {
			if err = d.Set("packagesettings", flattenPackagesPkgPackageSettings(o["package settings"], d, "packagesettings")); err != nil {
				if vv, ok := fortiAPIPatch(o["package settings"], "PackagesPkg-PackageSettings"); ok {
					if err = d.Set("packagesettings", vv); err != nil {
						return fmt.Errorf("Error reading packagesettings: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading packagesettings: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("packagesetting", flattenPackagesPkgPackageSetting(o["package setting"], d, "packagesetting")); err != nil {
			if vv, ok := fortiAPIPatch(o["package setting"], "PackagesPkg-PackageSetting"); ok {
				if err = d.Set("packagesetting", vv); err != nil {
					return fmt.Errorf("Error reading packagesetting: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading packagesetting: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("packagesetting"); ok {
			if err = d.Set("packagesetting", flattenPackagesPkgPackageSetting(o["package setting"], d, "packagesetting")); err != nil {
				if vv, ok := fortiAPIPatch(o["package setting"], "PackagesPkg-PackageSetting"); ok {
					if err = d.Set("packagesetting", vv); err != nil {
						return fmt.Errorf("Error reading packagesetting: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading packagesetting: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("scopemember", flattenPackagesPkgScopeMember(o["scope member"], d, "scopemember")); err != nil {
			if vv, ok := fortiAPIPatch(o["scope member"], "PackagesPkg-ScopeMember"); ok {
				if err = d.Set("scopemember", vv); err != nil {
					return fmt.Errorf("Error reading scopemember: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scopemember: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scopemember"); ok {
			if err = d.Set("scopemember", flattenPackagesPkgScopeMember(o["scope member"], d, "scopemember")); err != nil {
				if vv, ok := fortiAPIPatch(o["scope member"], "PackagesPkg-ScopeMember"); ok {
					if err = d.Set("scopemember", vv); err != nil {
						return fmt.Errorf("Error reading scopemember: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scopemember: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenPackagesPkgType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "PackagesPkg-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenPackagesPkgFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesPkgName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgObjVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgOid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "central_nat"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["central-nat"], _ = expandPackagesPkgPackageSettingsCentralNat(d, i["central_nat"], pre_append)
	}
	pre_append = pre + ".0." + "consolidated_firewall_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["consolidated-firewall-mode"], _ = expandPackagesPkgPackageSettingsConsolidatedFirewallMode(d, i["consolidated_firewall_mode"], pre_append)
	}
	pre_append = pre + ".0." + "fwpolicy_implicit_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwpolicy-implicit-log"], _ = expandPackagesPkgPackageSettingsFwpolicyImplicitLog(d, i["fwpolicy_implicit_log"], pre_append)
	}
	pre_append = pre + ".0." + "fwpolicy6_implicit_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwpolicy6-implicit-log"], _ = expandPackagesPkgPackageSettingsFwpolicy6ImplicitLog(d, i["fwpolicy6_implicit_log"], pre_append)
	}
	pre_append = pre + ".0." + "inspection_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspection-mode"], _ = expandPackagesPkgPackageSettingsInspectionMode(d, i["inspection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ngfw_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ngfw-mode"], _ = expandPackagesPkgPackageSettingsNgfwMode(d, i["ngfw_mode"], pre_append)
	}
	pre_append = pre + ".0." + "policy_offload_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["policy-offload-level"], _ = expandPackagesPkgPackageSettingsPolicyOffloadLevel(d, i["policy_offload_level"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_ssh_profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssl-ssh-profile"], _ = expandPackagesPkgPackageSettingsSslSshProfile(d, i["ssl_ssh_profile"], pre_append)
	}

	return result, nil
}

func expandPackagesPkgPackageSettingsCentralNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingsConsolidatedFirewallMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingsFwpolicyImplicitLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingsFwpolicy6ImplicitLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingsInspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingsNgfwMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingsPolicyOffloadLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingsSslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandPackagesPkgPackageSetting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "central_nat"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["central-nat"], _ = expandPackagesPkgPackageSettingCentralNat(d, i["central_nat"], pre_append)
	}
	pre_append = pre + ".0." + "consolidated_firewall_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["consolidated-firewall-mode"], _ = expandPackagesPkgPackageSettingConsolidatedFirewallMode(d, i["consolidated_firewall_mode"], pre_append)
	}
	pre_append = pre + ".0." + "fwpolicy_implicit_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwpolicy-implicit-log"], _ = expandPackagesPkgPackageSettingFwpolicyImplicitLog(d, i["fwpolicy_implicit_log"], pre_append)
	}
	pre_append = pre + ".0." + "fwpolicy6_implicit_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwpolicy6-implicit-log"], _ = expandPackagesPkgPackageSettingFwpolicy6ImplicitLog(d, i["fwpolicy6_implicit_log"], pre_append)
	}
	pre_append = pre + ".0." + "inspection_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspection-mode"], _ = expandPackagesPkgPackageSettingInspectionMode(d, i["inspection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ngfw_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ngfw-mode"], _ = expandPackagesPkgPackageSettingNgfwMode(d, i["ngfw_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_ssh_profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ssl-ssh-profile"], _ = expandPackagesPkgPackageSettingSslSshProfile(d, i["ssl_ssh_profile"], pre_append)
	}

	return result, nil
}

func expandPackagesPkgPackageSettingCentralNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingConsolidatedFirewallMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingFwpolicyImplicitLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingFwpolicy6ImplicitLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingInspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingNgfwMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgPackageSettingSslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgScopeMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandPackagesPkgScopeMemberName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandPackagesPkgScopeMemberVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandPackagesPkgScopeMemberName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgScopeMemberVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesPkg(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesPkgName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("objver"); ok || d.HasChange("objver") {
		t, err := expandPackagesPkgObjVer(d, v, "objver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj ver"] = t
		}
	}

	if v, ok := d.GetOk("oid"); ok || d.HasChange("oid") {
		t, err := expandPackagesPkgOid(d, v, "oid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oid"] = t
		}
	}

	if v, ok := d.GetOk("packagesettings"); ok || d.HasChange("packagesettings") {
		t, err := expandPackagesPkgPackageSettings(d, v, "packagesettings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["package settings"] = t
		}
	}

	if v, ok := d.GetOk("packagesetting"); ok || d.HasChange("packagesetting") {
		t, err := expandPackagesPkgPackageSetting(d, v, "packagesetting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["package setting"] = t
		}
	}

	if v, ok := d.GetOk("scopemember"); ok || d.HasChange("scopemember") {
		t, err := expandPackagesPkgScopeMember(d, v, "scopemember")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope member"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandPackagesPkgType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
