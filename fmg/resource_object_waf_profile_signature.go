// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: WAF signatures.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWafProfileSignature() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWafProfileSignatureUpdate,
		Read:   resourceObjectWafProfileSignatureRead,
		Update: resourceObjectWafProfileSignatureUpdate,
		Delete: resourceObjectWafProfileSignatureDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"credit_card_detection_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"custom_signature": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"case_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"target": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"disabled_signature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled_sub_class": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"main_class": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func resourceObjectWafProfileSignatureUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectWafProfileSignature(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileSignature resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWafProfileSignature(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileSignature resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWafProfileSignature")

	return resourceObjectWafProfileSignatureRead(d, m)
}

func resourceObjectWafProfileSignatureDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	err = c.DeleteObjectWafProfileSignature(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWafProfileSignature resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWafProfileSignatureRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectWafProfileSignature(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileSignature resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWafProfileSignature(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileSignature resource from API: %v", err)
	}
	return nil
}

func flattenObjectWafProfileSignatureCreditCardDetectionThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignature2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureCaseSensitivity2edl(i["case-sensitivity"], d, pre_append)
			tmp["case_sensitivity"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-CaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureDirection2edl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureLog2edl(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignaturePattern2edl(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureSeverity2edl(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureStatus2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureTarget2edl(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Target")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectWafProfileSignatureCustomSignatureAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureCaseSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignaturePattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWafProfileSignatureDisabledSignature2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWafProfileSignatureDisabledSubClass2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWafProfileSignatureMainClass2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileSignatureMainClassAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenObjectWafProfileSignatureMainClassId2edl(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileSignatureMainClassLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileSignatureMainClassSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileSignatureMainClassStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileSignatureMainClassAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureMainClassId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2num(v)
}

func flattenObjectWafProfileSignatureMainClassLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureMainClassSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureMainClassStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWafProfileSignature(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("credit_card_detection_threshold", flattenObjectWafProfileSignatureCreditCardDetectionThreshold2edl(o["credit-card-detection-threshold"], d, "credit_card_detection_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["credit-card-detection-threshold"], "ObjectWafProfileSignature-CreditCardDetectionThreshold"); ok {
			if err = d.Set("credit_card_detection_threshold", vv); err != nil {
				return fmt.Errorf("Error reading credit_card_detection_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading credit_card_detection_threshold: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_signature", flattenObjectWafProfileSignatureCustomSignature2edl(o["custom-signature"], d, "custom_signature")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-signature"], "ObjectWafProfileSignature-CustomSignature"); ok {
				if err = d.Set("custom_signature", vv); err != nil {
					return fmt.Errorf("Error reading custom_signature: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_signature: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_signature"); ok {
			if err = d.Set("custom_signature", flattenObjectWafProfileSignatureCustomSignature2edl(o["custom-signature"], d, "custom_signature")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-signature"], "ObjectWafProfileSignature-CustomSignature"); ok {
					if err = d.Set("custom_signature", vv); err != nil {
						return fmt.Errorf("Error reading custom_signature: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_signature: %v", err)
				}
			}
		}
	}

	if err = d.Set("disabled_signature", flattenObjectWafProfileSignatureDisabledSignature2edl(o["disabled-signature"], d, "disabled_signature")); err != nil {
		if vv, ok := fortiAPIPatch(o["disabled-signature"], "ObjectWafProfileSignature-DisabledSignature"); ok {
			if err = d.Set("disabled_signature", vv); err != nil {
				return fmt.Errorf("Error reading disabled_signature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disabled_signature: %v", err)
		}
	}

	if err = d.Set("disabled_sub_class", flattenObjectWafProfileSignatureDisabledSubClass2edl(o["disabled-sub-class"], d, "disabled_sub_class")); err != nil {
		if vv, ok := fortiAPIPatch(o["disabled-sub-class"], "ObjectWafProfileSignature-DisabledSubClass"); ok {
			if err = d.Set("disabled_sub_class", vv); err != nil {
				return fmt.Errorf("Error reading disabled_sub_class: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disabled_sub_class: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("main_class", flattenObjectWafProfileSignatureMainClass2edl(o["main-class"], d, "main_class")); err != nil {
			if vv, ok := fortiAPIPatch(o["main-class"], "ObjectWafProfileSignature-MainClass"); ok {
				if err = d.Set("main_class", vv); err != nil {
					return fmt.Errorf("Error reading main_class: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading main_class: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("main_class"); ok {
			if err = d.Set("main_class", flattenObjectWafProfileSignatureMainClass2edl(o["main-class"], d, "main_class")); err != nil {
				if vv, ok := fortiAPIPatch(o["main-class"], "ObjectWafProfileSignature-MainClass"); ok {
					if err = d.Set("main_class", vv); err != nil {
						return fmt.Errorf("Error reading main_class: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading main_class: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectWafProfileSignatureFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWafProfileSignatureCreditCardDetectionThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignature2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectWafProfileSignatureCustomSignatureAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitivity"], _ = expandObjectWafProfileSignatureCustomSignatureCaseSensitivity2edl(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandObjectWafProfileSignatureCustomSignatureDirection2edl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectWafProfileSignatureCustomSignatureLog2edl(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWafProfileSignatureCustomSignatureName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandObjectWafProfileSignatureCustomSignaturePattern2edl(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandObjectWafProfileSignatureCustomSignatureSeverity2edl(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectWafProfileSignatureCustomSignatureStatus2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandObjectWafProfileSignatureCustomSignatureTarget2edl(d, i["target"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectWafProfileSignatureCustomSignatureAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureCaseSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignaturePattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWafProfileSignatureDisabledSignature2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWafProfileSignatureDisabledSubClass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWafProfileSignatureMainClass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileSignatureMainClassAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandObjectWafProfileSignatureMainClassId2edl(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileSignatureMainClassLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileSignatureMainClassSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileSignatureMainClassStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileSignatureMainClassAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureMainClassId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureMainClassLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureMainClassSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureMainClassStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWafProfileSignature(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("credit_card_detection_threshold"); ok || d.HasChange("credit_card_detection_threshold") {
		t, err := expandObjectWafProfileSignatureCreditCardDetectionThreshold2edl(d, v, "credit_card_detection_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["credit-card-detection-threshold"] = t
		}
	}

	if v, ok := d.GetOk("custom_signature"); ok || d.HasChange("custom_signature") {
		t, err := expandObjectWafProfileSignatureCustomSignature2edl(d, v, "custom_signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-signature"] = t
		}
	}

	if v, ok := d.GetOk("disabled_signature"); ok || d.HasChange("disabled_signature") {
		t, err := expandObjectWafProfileSignatureDisabledSignature2edl(d, v, "disabled_signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disabled-signature"] = t
		}
	}

	if v, ok := d.GetOk("disabled_sub_class"); ok || d.HasChange("disabled_sub_class") {
		t, err := expandObjectWafProfileSignatureDisabledSubClass2edl(d, v, "disabled_sub_class")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disabled-sub-class"] = t
		}
	}

	if v, ok := d.GetOk("main_class"); ok || d.HasChange("main_class") {
		t, err := expandObjectWafProfileSignatureMainClass2edl(d, v, "main_class")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["main-class"] = t
		}
	}

	return &obj, nil
}
