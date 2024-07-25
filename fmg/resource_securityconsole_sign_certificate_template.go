// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Generate and sign certificate on the target device.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSecurityconsoleSignCertificateTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityconsoleSignCertificateTemplateUpdate,
		Read:   resourceSecurityconsoleSignCertificateTemplateRead,
		Update: resourceSecurityconsoleSignCertificateTemplateUpdate,
		Delete: resourceSecurityconsoleSignCertificateTemplateDelete,

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
			"template": &schema.Schema{
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

func resourceSecurityconsoleSignCertificateTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSecurityconsoleSignCertificateTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleSignCertificateTemplate resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityconsoleSignCertificateTemplate(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleSignCertificateTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SecurityconsoleSignCertificateTemplate")

	return resourceSecurityconsoleSignCertificateTemplateRead(d, m)
}

func resourceSecurityconsoleSignCertificateTemplateDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSecurityconsoleSignCertificateTemplateRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenSecurityconsoleSignCertificateTemplateAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleSignCertificateTemplateScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSecurityconsoleSignCertificateTemplateScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SecurityconsoleSignCertificateTemplate-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenSecurityconsoleSignCertificateTemplateScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "SecurityconsoleSignCertificateTemplate-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSecurityconsoleSignCertificateTemplateScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleSignCertificateTemplateScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleSignCertificateTemplateTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectSecurityconsoleSignCertificateTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fmgadom", flattenSecurityconsoleSignCertificateTemplateAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SecurityconsoleSignCertificateTemplate-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("scope", flattenSecurityconsoleSignCertificateTemplateScope(o["scope"], d, "scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["scope"], "SecurityconsoleSignCertificateTemplate-Scope"); ok {
				if err = d.Set("scope", vv); err != nil {
					return fmt.Errorf("Error reading scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scope"); ok {
			if err = d.Set("scope", flattenSecurityconsoleSignCertificateTemplateScope(o["scope"], d, "scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["scope"], "SecurityconsoleSignCertificateTemplate-Scope"); ok {
					if err = d.Set("scope", vv); err != nil {
						return fmt.Errorf("Error reading scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("template", flattenSecurityconsoleSignCertificateTemplateTemplate(o["template"], d, "template")); err != nil {
		if vv, ok := fortiAPIPatch(o["template"], "SecurityconsoleSignCertificateTemplate-Template"); ok {
			if err = d.Set("template", vv); err != nil {
				return fmt.Errorf("Error reading template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template: %v", err)
		}
	}

	return nil
}

func flattenSecurityconsoleSignCertificateTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSecurityconsoleSignCertificateTemplateAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleSignCertificateTemplateScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSecurityconsoleSignCertificateTemplateScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandSecurityconsoleSignCertificateTemplateScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSecurityconsoleSignCertificateTemplateScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleSignCertificateTemplateScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleSignCertificateTemplateTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectSecurityconsoleSignCertificateTemplate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandSecurityconsoleSignCertificateTemplateAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok || d.HasChange("scope") {
		t, err := expandSecurityconsoleSignCertificateTemplateScope(d, v, "scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	if v, ok := d.GetOk("template"); ok || d.HasChange("template") {
		t, err := expandSecurityconsoleSignCertificateTemplateTemplate(d, v, "template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template"] = t
		}
	}

	return &obj, nil
}
