// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Re-install a policy package that had been previously installed.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSecurityconsoleReinstallPackage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityconsoleReinstallPackageUpdate,
		Read:   resourceSecurityconsoleReinstallPackageRead,
		Update: resourceSecurityconsoleReinstallPackageUpdate,
		Delete: resourceSecurityconsoleReinstallPackageDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"fmgadom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scope": &schema.Schema{
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
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSecurityconsoleReinstallPackageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectSecurityconsoleReinstallPackage(d)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleReinstallPackage resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityconsoleReinstallPackage(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleReinstallPackage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SecurityconsoleReinstallPackage")

	return resourceSecurityconsoleReinstallPackageRead(d, m)
}

func resourceSecurityconsoleReinstallPackageDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSecurityconsoleReinstallPackageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenSecurityconsoleReinstallPackageAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleReinstallPackageFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:     "none",
			1:     "cp_all_objs",
			2:     "preview",
			8:     "generate_rev",
			16:    "copy_assigned_pkg",
			128:   "unassign",
			256:   "ifpolicy_only",
			512:   "no_ifpolicy",
			1024:  "objs_only",
			4096:  "auto_lock_ws",
			8192:  "check_pkg_st",
			16384: "copy_only",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenSecurityconsoleReinstallPackageTarget(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pkg"
		if _, ok := i["pkg"]; ok {
			v := flattenSecurityconsoleReinstallPackageTargetPkg(i["pkg"], d, pre_append)
			tmp["pkg"] = fortiAPISubPartPatch(v, "SecurityconsoleReinstallPackage-Target-Pkg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scope"
		if _, ok := i["scope"]; ok {
			v := flattenSecurityconsoleReinstallPackageTargetScope(i["scope"], d, pre_append)
			tmp["scope"] = fortiAPISubPartPatch(v, "SecurityconsoleReinstallPackage-Target-Scope")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSecurityconsoleReinstallPackageTargetPkg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleReinstallPackageTargetScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSecurityconsoleReinstallPackageTargetScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SecurityconsoleReinstallPackageTarget-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenSecurityconsoleReinstallPackageTargetScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "SecurityconsoleReinstallPackageTarget-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSecurityconsoleReinstallPackageTargetScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleReinstallPackageTargetScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSecurityconsoleReinstallPackage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgadom", flattenSecurityconsoleReinstallPackageAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SecurityconsoleReinstallPackage-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("flags", flattenSecurityconsoleReinstallPackageFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "SecurityconsoleReinstallPackage-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("target", flattenSecurityconsoleReinstallPackageTarget(o["target"], d, "target")); err != nil {
			if vv, ok := fortiAPIPatch(o["target"], "SecurityconsoleReinstallPackage-Target"); ok {
				if err = d.Set("target", vv); err != nil {
					return fmt.Errorf("Error reading target: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading target: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("target"); ok {
			if err = d.Set("target", flattenSecurityconsoleReinstallPackageTarget(o["target"], d, "target")); err != nil {
				if vv, ok := fortiAPIPatch(o["target"], "SecurityconsoleReinstallPackage-Target"); ok {
					if err = d.Set("target", vv); err != nil {
						return fmt.Errorf("Error reading target: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading target: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSecurityconsoleReinstallPackageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSecurityconsoleReinstallPackageAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleReinstallPackageFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSecurityconsoleReinstallPackageTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pkg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pkg"], _ = expandSecurityconsoleReinstallPackageTargetPkg(d, i["pkg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scope"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["scope"], _ = expandSecurityconsoleReinstallPackageTargetScope(d, i["scope"], pre_append)
		} else {
			tmp["scope"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSecurityconsoleReinstallPackageTargetPkg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleReinstallPackageTargetScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSecurityconsoleReinstallPackageTargetScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandSecurityconsoleReinstallPackageTargetScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSecurityconsoleReinstallPackageTargetScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleReinstallPackageTargetScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityconsoleReinstallPackage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok {
		t, err := expandSecurityconsoleReinstallPackageAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok {
		t, err := expandSecurityconsoleReinstallPackageFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok {
		t, err := expandSecurityconsoleReinstallPackageTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	return &obj, nil
}
