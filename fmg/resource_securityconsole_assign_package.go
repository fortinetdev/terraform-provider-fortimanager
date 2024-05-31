// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Assign or unassign global policy package to ADOM packages.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSecurityconsoleAssignPackage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityconsoleAssignPackageUpdate,
		Read:   resourceSecurityconsoleAssignPackageRead,
		Update: resourceSecurityconsoleAssignPackageUpdate,
		Delete: resourceSecurityconsoleAssignPackageDelete,

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
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"adom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"excluded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceSecurityconsoleAssignPackageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSecurityconsoleAssignPackage(d)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleAssignPackage resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityconsoleAssignPackage(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleAssignPackage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SecurityconsoleAssignPackage")

	return resourceSecurityconsoleAssignPackageRead(d, m)
}

func resourceSecurityconsoleAssignPackageDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSecurityconsoleAssignPackageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenSecurityconsoleAssignPackageFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSecurityconsoleAssignPackagePkg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSecurityconsoleAssignPackageTarget(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adom"
		if _, ok := i["adom"]; ok {
			v := flattenSecurityconsoleAssignPackageTargetAdom(i["adom"], d, pre_append)
			tmp["adom"] = fortiAPISubPartPatch(v, "SecurityconsoleAssignPackage-Target-Adom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "excluded"
		if _, ok := i["excluded"]; ok {
			v := flattenSecurityconsoleAssignPackageTargetExcluded(i["excluded"], d, pre_append)
			tmp["excluded"] = fortiAPISubPartPatch(v, "SecurityconsoleAssignPackage-Target-Excluded")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pkg"
		if _, ok := i["pkg"]; ok {
			v := flattenSecurityconsoleAssignPackageTargetPkg(i["pkg"], d, pre_append)
			tmp["pkg"] = fortiAPISubPartPatch(v, "SecurityconsoleAssignPackage-Target-Pkg")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSecurityconsoleAssignPackageTargetAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSecurityconsoleAssignPackageTargetExcluded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleAssignPackageTargetPkg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectSecurityconsoleAssignPackage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("flags", flattenSecurityconsoleAssignPackageFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "SecurityconsoleAssignPackage-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("pkg", flattenSecurityconsoleAssignPackagePkg(o["pkg"], d, "pkg")); err != nil {
		if vv, ok := fortiAPIPatch(o["pkg"], "SecurityconsoleAssignPackage-Pkg"); ok {
			if err = d.Set("pkg", vv); err != nil {
				return fmt.Errorf("Error reading pkg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pkg: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("target", flattenSecurityconsoleAssignPackageTarget(o["target"], d, "target")); err != nil {
			if vv, ok := fortiAPIPatch(o["target"], "SecurityconsoleAssignPackage-Target"); ok {
				if err = d.Set("target", vv); err != nil {
					return fmt.Errorf("Error reading target: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading target: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("target"); ok {
			if err = d.Set("target", flattenSecurityconsoleAssignPackageTarget(o["target"], d, "target")); err != nil {
				if vv, ok := fortiAPIPatch(o["target"], "SecurityconsoleAssignPackage-Target"); ok {
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

func flattenSecurityconsoleAssignPackageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSecurityconsoleAssignPackageFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSecurityconsoleAssignPackagePkg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSecurityconsoleAssignPackageTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adom"], _ = expandSecurityconsoleAssignPackageTargetAdom(d, i["adom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "excluded"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["excluded"], _ = expandSecurityconsoleAssignPackageTargetExcluded(d, i["excluded"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pkg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pkg"], _ = expandSecurityconsoleAssignPackageTargetPkg(d, i["pkg"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSecurityconsoleAssignPackageTargetAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSecurityconsoleAssignPackageTargetExcluded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleAssignPackageTargetPkg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectSecurityconsoleAssignPackage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandSecurityconsoleAssignPackageFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("pkg"); ok || d.HasChange("pkg") {
		t, err := expandSecurityconsoleAssignPackagePkg(d, v, "pkg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pkg"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandSecurityconsoleAssignPackageTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	return &obj, nil
}
