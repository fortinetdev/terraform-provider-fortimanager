// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Run script.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDvmdbScriptExecute() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmdbScriptExecuteUpdate,
		Read:   resourceDvmdbScriptExecuteRead,
		Update: resourceDvmdbScriptExecuteUpdate,
		Delete: resourceDvmdbScriptExecuteDelete,

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
			"fmgadom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"package": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pblock": &schema.Schema{
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
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceDvmdbScriptExecuteUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDvmdbScriptExecute(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbScriptExecute resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmdbScriptExecute(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbScriptExecute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DvmdbScriptExecute")

	return resourceDvmdbScriptExecuteRead(d, m)
}

func resourceDvmdbScriptExecuteDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceDvmdbScriptExecuteRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenDvmdbScriptExecuteAdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptExecutePackage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptExecutePblock2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptExecuteScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenDvmdbScriptExecuteScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "DvmdbScriptExecute-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenDvmdbScriptExecuteScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "DvmdbScriptExecute-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDvmdbScriptExecuteScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptExecuteScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptExecuteScript2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmdbScriptExecute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fmgadom", flattenDvmdbScriptExecuteAdom2edl(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "DvmdbScriptExecute-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("package", flattenDvmdbScriptExecutePackage2edl(o["package"], d, "package")); err != nil {
		if vv, ok := fortiAPIPatch(o["package"], "DvmdbScriptExecute-Package"); ok {
			if err = d.Set("package", vv); err != nil {
				return fmt.Errorf("Error reading package: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading package: %v", err)
		}
	}

	if err = d.Set("pblock", flattenDvmdbScriptExecutePblock2edl(o["pblock"], d, "pblock")); err != nil {
		if vv, ok := fortiAPIPatch(o["pblock"], "DvmdbScriptExecute-Pblock"); ok {
			if err = d.Set("pblock", vv); err != nil {
				return fmt.Errorf("Error reading pblock: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pblock: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("scope", flattenDvmdbScriptExecuteScope2edl(o["scope"], d, "scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["scope"], "DvmdbScriptExecute-Scope"); ok {
				if err = d.Set("scope", vv); err != nil {
					return fmt.Errorf("Error reading scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scope"); ok {
			if err = d.Set("scope", flattenDvmdbScriptExecuteScope2edl(o["scope"], d, "scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["scope"], "DvmdbScriptExecute-Scope"); ok {
					if err = d.Set("scope", vv); err != nil {
						return fmt.Errorf("Error reading scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("script", flattenDvmdbScriptExecuteScript2edl(o["script"], d, "script")); err != nil {
		if vv, ok := fortiAPIPatch(o["script"], "DvmdbScriptExecute-Script"); ok {
			if err = d.Set("script", vv); err != nil {
				return fmt.Errorf("Error reading script: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	return nil
}

func flattenDvmdbScriptExecuteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmdbScriptExecuteAdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptExecutePackage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptExecutePblock2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptExecuteScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandDvmdbScriptExecuteScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandDvmdbScriptExecuteScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDvmdbScriptExecuteScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptExecuteScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptExecuteScript2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmdbScriptExecute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandDvmdbScriptExecuteAdom2edl(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("package"); ok || d.HasChange("package") {
		t, err := expandDvmdbScriptExecutePackage2edl(d, v, "package")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["package"] = t
		}
	}

	if v, ok := d.GetOk("pblock"); ok || d.HasChange("pblock") {
		t, err := expandDvmdbScriptExecutePblock2edl(d, v, "pblock")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pblock"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok || d.HasChange("scope") {
		t, err := expandDvmdbScriptExecuteScope2edl(d, v, "scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok || d.HasChange("script") {
		t, err := expandDvmdbScriptExecuteScript2edl(d, v, "script")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	}

	return &obj, nil
}
