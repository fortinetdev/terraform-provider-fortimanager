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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
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
			"packagesetting": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
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
						"ssl_ssh_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"subobj": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourcePackagesPkgCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectPackagesPkg(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesPkg resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesPkg(obj, adomv, nil)

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

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectPackagesPkg(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesPkg resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesPkg(obj, adomv, mkey, nil)
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

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeletePackagesPkg(adomv, mkey, nil)
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

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadPackagesPkg(adomv, mkey, nil)
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

		result = append(result, tmp)

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

func flattenPackagesPkgSubobj(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenPackagesPkgSubobjName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "PackagesPkg-Subobj-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenPackagesPkgSubobjName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesPkgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesPkg(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

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

	if isImportTable() {
		if err = d.Set("subobj", flattenPackagesPkgSubobj(o["subobj"], d, "subobj")); err != nil {
			if vv, ok := fortiAPIPatch(o["subobj"], "PackagesPkg-Subobj"); ok {
				if err = d.Set("subobj", vv); err != nil {
					return fmt.Errorf("Error reading subobj: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading subobj: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("subobj"); ok {
			if err = d.Set("subobj", flattenPackagesPkgSubobj(o["subobj"], d, "subobj")); err != nil {
				if vv, ok := fortiAPIPatch(o["subobj"], "PackagesPkg-Subobj"); ok {
					if err = d.Set("subobj", vv); err != nil {
						return fmt.Errorf("Error reading subobj: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading subobj: %v", err)
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

func expandPackagesPkgPackageSetting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "central_nat"
	if _, ok := d.GetOk(pre_append); ok {
		result["central-nat"], _ = expandPackagesPkgPackageSettingCentralNat(d, i["central_nat"], pre_append)
	}
	pre_append = pre + ".0." + "consolidated_firewall_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["consolidated-firewall-mode"], _ = expandPackagesPkgPackageSettingConsolidatedFirewallMode(d, i["consolidated_firewall_mode"], pre_append)
	}
	pre_append = pre + ".0." + "fwpolicy_implicit_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwpolicy-implicit-log"], _ = expandPackagesPkgPackageSettingFwpolicyImplicitLog(d, i["fwpolicy_implicit_log"], pre_append)
	}
	pre_append = pre + ".0." + "fwpolicy6_implicit_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwpolicy6-implicit-log"], _ = expandPackagesPkgPackageSettingFwpolicy6ImplicitLog(d, i["fwpolicy6_implicit_log"], pre_append)
	}
	pre_append = pre + ".0." + "inspection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspection-mode"], _ = expandPackagesPkgPackageSettingInspectionMode(d, i["inspection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ngfw_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["ngfw-mode"], _ = expandPackagesPkgPackageSettingNgfwMode(d, i["ngfw_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_ssh_profile"
	if _, ok := d.GetOk(pre_append); ok {
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
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandPackagesPkgScopeMemberName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandPackagesPkgScopeMemberVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

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

func expandPackagesPkgSubobj(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandPackagesPkgSubobjName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandPackagesPkgSubobjName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesPkgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesPkg(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandPackagesPkgName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("objver"); ok {
		t, err := expandPackagesPkgObjVer(d, v, "objver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj ver"] = t
		}
	}

	if v, ok := d.GetOk("oid"); ok {
		t, err := expandPackagesPkgOid(d, v, "oid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oid"] = t
		}
	}

	if v, ok := d.GetOk("packagesetting"); ok {
		t, err := expandPackagesPkgPackageSetting(d, v, "packagesetting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["package setting"] = t
		}
	}

	if v, ok := d.GetOk("scopemember"); ok {
		t, err := expandPackagesPkgScopeMember(d, v, "scopemember")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope member"] = t
		}
	}

	if v, ok := d.GetOk("subobj"); ok {
		t, err := expandPackagesPkgSubobj(d, v, "subobj")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subobj"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandPackagesPkgType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
