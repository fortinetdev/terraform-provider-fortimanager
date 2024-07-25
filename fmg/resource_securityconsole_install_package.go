// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Copy and install a policy package to devices.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSecurityconsoleInstallPackage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityconsoleInstallPackageUpdate,
		Read:   resourceSecurityconsoleInstallPackageRead,
		Update: resourceSecurityconsoleInstallPackageUpdate,
		Delete: resourceSecurityconsoleInstallPackageDelete,

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
			"adom_rev_comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"adom_rev_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dev_rev_comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"scope": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSecurityconsoleInstallPackageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSecurityconsoleInstallPackage(d)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleInstallPackage resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityconsoleInstallPackage(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleInstallPackage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SecurityconsoleInstallPackage")

	return resourceSecurityconsoleInstallPackageRead(d, m)
}

func resourceSecurityconsoleInstallPackageDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSecurityconsoleInstallPackageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenSecurityconsoleInstallPackageAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleInstallPackageAdomRevComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleInstallPackageAdomRevName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleInstallPackageDevRevComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleInstallPackageFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSecurityconsoleInstallPackagePkg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleInstallPackageScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSecurityconsoleInstallPackageScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SecurityconsoleInstallPackage-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenSecurityconsoleInstallPackageScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "SecurityconsoleInstallPackage-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSecurityconsoleInstallPackageScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleInstallPackageScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSecurityconsoleInstallPackage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fmgadom", flattenSecurityconsoleInstallPackageAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SecurityconsoleInstallPackage-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("adom_rev_comments", flattenSecurityconsoleInstallPackageAdomRevComments(o["adom_rev_comments"], d, "adom_rev_comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom_rev_comments"], "SecurityconsoleInstallPackage-AdomRevComments"); ok {
			if err = d.Set("adom_rev_comments", vv); err != nil {
				return fmt.Errorf("Error reading adom_rev_comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_rev_comments: %v", err)
		}
	}

	if err = d.Set("adom_rev_name", flattenSecurityconsoleInstallPackageAdomRevName(o["adom_rev_name"], d, "adom_rev_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom_rev_name"], "SecurityconsoleInstallPackage-AdomRevName"); ok {
			if err = d.Set("adom_rev_name", vv); err != nil {
				return fmt.Errorf("Error reading adom_rev_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_rev_name: %v", err)
		}
	}

	if err = d.Set("dev_rev_comments", flattenSecurityconsoleInstallPackageDevRevComments(o["dev_rev_comments"], d, "dev_rev_comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev_rev_comments"], "SecurityconsoleInstallPackage-DevRevComments"); ok {
			if err = d.Set("dev_rev_comments", vv); err != nil {
				return fmt.Errorf("Error reading dev_rev_comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_rev_comments: %v", err)
		}
	}

	if err = d.Set("flags", flattenSecurityconsoleInstallPackageFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "SecurityconsoleInstallPackage-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("pkg", flattenSecurityconsoleInstallPackagePkg(o["pkg"], d, "pkg")); err != nil {
		if vv, ok := fortiAPIPatch(o["pkg"], "SecurityconsoleInstallPackage-Pkg"); ok {
			if err = d.Set("pkg", vv); err != nil {
				return fmt.Errorf("Error reading pkg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pkg: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("scope", flattenSecurityconsoleInstallPackageScope(o["scope"], d, "scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["scope"], "SecurityconsoleInstallPackage-Scope"); ok {
				if err = d.Set("scope", vv); err != nil {
					return fmt.Errorf("Error reading scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scope"); ok {
			if err = d.Set("scope", flattenSecurityconsoleInstallPackageScope(o["scope"], d, "scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["scope"], "SecurityconsoleInstallPackage-Scope"); ok {
					if err = d.Set("scope", vv); err != nil {
						return fmt.Errorf("Error reading scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scope: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSecurityconsoleInstallPackageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSecurityconsoleInstallPackageAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleInstallPackageAdomRevComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleInstallPackageAdomRevName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleInstallPackageDevRevComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleInstallPackageFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSecurityconsoleInstallPackagePkg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleInstallPackageScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSecurityconsoleInstallPackageScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandSecurityconsoleInstallPackageScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSecurityconsoleInstallPackageScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleInstallPackageScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityconsoleInstallPackage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandSecurityconsoleInstallPackageAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("adom_rev_comments"); ok || d.HasChange("adom_rev_comments") {
		t, err := expandSecurityconsoleInstallPackageAdomRevComments(d, v, "adom_rev_comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom_rev_comments"] = t
		}
	}

	if v, ok := d.GetOk("adom_rev_name"); ok || d.HasChange("adom_rev_name") {
		t, err := expandSecurityconsoleInstallPackageAdomRevName(d, v, "adom_rev_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom_rev_name"] = t
		}
	}

	if v, ok := d.GetOk("dev_rev_comments"); ok || d.HasChange("dev_rev_comments") {
		t, err := expandSecurityconsoleInstallPackageDevRevComments(d, v, "dev_rev_comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev_rev_comments"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandSecurityconsoleInstallPackageFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("pkg"); ok || d.HasChange("pkg") {
		t, err := expandSecurityconsoleInstallPackagePkg(d, v, "pkg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pkg"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok || d.HasChange("scope") {
		t, err := expandSecurityconsoleInstallPackageScope(d, v, "scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	return &obj, nil
}
