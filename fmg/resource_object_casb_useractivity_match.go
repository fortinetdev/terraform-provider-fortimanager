// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB user activity match rules.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbUserActivityMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbUserActivityMatchCreate,
		Read:   resourceObjectCasbUserActivityMatchRead,
		Update: resourceObjectCasbUserActivityMatchUpdate,
		Delete: resourceObjectCasbUserActivityMatchDelete,

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
			"user_activity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"body_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"case_sensitive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domains": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"jq": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"methods": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"strategy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_extraction": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"body_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"header_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"place": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"jq": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
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

func resourceObjectCasbUserActivityMatchCreate(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	paradict["user_activity"] = user_activity

	obj, err := getObjectObjectCasbUserActivityMatch(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityMatch resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCasbUserActivityMatch(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityMatch resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbUserActivityMatchRead(d, m)
}

func resourceObjectCasbUserActivityMatchUpdate(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	paradict["user_activity"] = user_activity

	obj, err := getObjectObjectCasbUserActivityMatch(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatch resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbUserActivityMatch(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbUserActivityMatchRead(d, m)
}

func resourceObjectCasbUserActivityMatchDelete(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	paradict["user_activity"] = user_activity

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbUserActivityMatch(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbUserActivityMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbUserActivityMatchRead(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	if user_activity == "" {
		user_activity = importOptionChecking(m.(*FortiClient).Cfg, "user_activity")
		if user_activity == "" {
			return fmt.Errorf("Parameter user_activity is missing")
		}
		if err = d.Set("user_activity", user_activity); err != nil {
			return fmt.Errorf("Error set params user_activity: %v", err)
		}
	}
	paradict["user_activity"] = user_activity

	o, err := c.ReadObjectCasbUserActivityMatch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbUserActivityMatch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatch resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbUserActivityMatchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRules2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := i["body-type"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesBodyType2edl(i["body-type"], d, pre_append)
			tmp["body_type"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-BodyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := i["case-sensitive"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesCaseSensitive2edl(i["case-sensitive"], d, pre_append)
			tmp["case_sensitive"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-CaseSensitive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := i["domains"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesDomains2edl(i["domains"], d, pre_append)
			tmp["domains"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-Domains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesHeaderName2edl(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jq"
		if _, ok := i["jq"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesJq2edl(i["jq"], d, pre_append)
			tmp["jq"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-Jq")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := i["match-pattern"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesMatchPattern2edl(i["match-pattern"], d, pre_append)
			tmp["match_pattern"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-MatchPattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := i["match-value"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesMatchValue2edl(i["match-value"], d, pre_append)
			tmp["match_value"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-MatchValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "methods"
		if _, ok := i["methods"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesMethods2edl(i["methods"], d, pre_append)
			tmp["methods"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-Methods")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := i["negate"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesNegate2edl(i["negate"], d, pre_append)
			tmp["negate"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-Negate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectCasbUserActivityMatchRulesType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatch-Rules-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbUserActivityMatchRulesBodyType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesCaseSensitive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesDomains2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbUserActivityMatchRulesHeaderName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesJq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesMatchPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesMatchValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesMethods2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbUserActivityMatchRulesNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchStrategy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtraction2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenObjectCasbUserActivityMatchTenantExtractionFilters2edl(i["filters"], d, pre_append)
	}

	pre_append = pre + ".0." + "jq"
	if _, ok := i["jq"]; ok {
		result["jq"] = flattenObjectCasbUserActivityMatchTenantExtractionJq2edl(i["jq"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectCasbUserActivityMatchTenantExtractionStatus2edl(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "type"
	if _, ok := i["type"]; ok {
		result["type"] = flattenObjectCasbUserActivityMatchTenantExtractionType2edl(i["type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectCasbUserActivityMatchTenantExtractionFilters2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := i["body-type"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantExtractionFiltersBodyType2edl(i["body-type"], d, pre_append)
			tmp["body_type"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantExtraction-Filters-BodyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantExtractionFiltersDirection2edl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantExtraction-Filters-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName2edl(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantExtraction-Filters-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantExtractionFiltersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantExtraction-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := i["place"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantExtractionFiltersPlace2edl(i["place"], d, pre_append)
			tmp["place"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantExtraction-Filters-Place")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersBodyType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersPlace2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionJq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbUserActivityMatch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fosid", flattenObjectCasbUserActivityMatchId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectCasbUserActivityMatch-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rules", flattenObjectCasbUserActivityMatchRules2edl(o["rules"], d, "rules")); err != nil {
			if vv, ok := fortiAPIPatch(o["rules"], "ObjectCasbUserActivityMatch-Rules"); ok {
				if err = d.Set("rules", vv); err != nil {
					return fmt.Errorf("Error reading rules: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rules"); ok {
			if err = d.Set("rules", flattenObjectCasbUserActivityMatchRules2edl(o["rules"], d, "rules")); err != nil {
				if vv, ok := fortiAPIPatch(o["rules"], "ObjectCasbUserActivityMatch-Rules"); ok {
					if err = d.Set("rules", vv); err != nil {
						return fmt.Errorf("Error reading rules: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rules: %v", err)
				}
			}
		}
	}

	if err = d.Set("strategy", flattenObjectCasbUserActivityMatchStrategy2edl(o["strategy"], d, "strategy")); err != nil {
		if vv, ok := fortiAPIPatch(o["strategy"], "ObjectCasbUserActivityMatch-Strategy"); ok {
			if err = d.Set("strategy", vv); err != nil {
				return fmt.Errorf("Error reading strategy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strategy: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tenant_extraction", flattenObjectCasbUserActivityMatchTenantExtraction2edl(o["tenant-extraction"], d, "tenant_extraction")); err != nil {
			if vv, ok := fortiAPIPatch(o["tenant-extraction"], "ObjectCasbUserActivityMatch-TenantExtraction"); ok {
				if err = d.Set("tenant_extraction", vv); err != nil {
					return fmt.Errorf("Error reading tenant_extraction: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tenant_extraction: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tenant_extraction"); ok {
			if err = d.Set("tenant_extraction", flattenObjectCasbUserActivityMatchTenantExtraction2edl(o["tenant-extraction"], d, "tenant_extraction")); err != nil {
				if vv, ok := fortiAPIPatch(o["tenant-extraction"], "ObjectCasbUserActivityMatch-TenantExtraction"); ok {
					if err = d.Set("tenant_extraction", vv); err != nil {
						return fmt.Errorf("Error reading tenant_extraction: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tenant_extraction: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectCasbUserActivityMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbUserActivityMatchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRules2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["body-type"], _ = expandObjectCasbUserActivityMatchRulesBodyType2edl(d, i["body_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitive"], _ = expandObjectCasbUserActivityMatchRulesCaseSensitive2edl(d, i["case_sensitive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domains"], _ = expandObjectCasbUserActivityMatchRulesDomains2edl(d, i["domains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandObjectCasbUserActivityMatchRulesHeaderName2edl(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectCasbUserActivityMatchRulesId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jq"], _ = expandObjectCasbUserActivityMatchRulesJq2edl(d, i["jq"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-pattern"], _ = expandObjectCasbUserActivityMatchRulesMatchPattern2edl(d, i["match_pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-value"], _ = expandObjectCasbUserActivityMatchRulesMatchValue2edl(d, i["match_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "methods"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["methods"], _ = expandObjectCasbUserActivityMatchRulesMethods2edl(d, i["methods"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["negate"], _ = expandObjectCasbUserActivityMatchRulesNegate2edl(d, i["negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectCasbUserActivityMatchRulesType2edl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbUserActivityMatchRulesBodyType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesCaseSensitive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesDomains2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbUserActivityMatchRulesHeaderName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesJq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesMatchPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesMatchValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesMethods2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbUserActivityMatchRulesNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchStrategy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtraction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectCasbUserActivityMatchTenantExtractionFilters2edl(d, i["filters"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["filters"] = t
		}
	}
	pre_append = pre + ".0." + "jq"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["jq"], _ = expandObjectCasbUserActivityMatchTenantExtractionJq2edl(d, i["jq"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectCasbUserActivityMatchTenantExtractionStatus2edl(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["type"], _ = expandObjectCasbUserActivityMatchTenantExtractionType2edl(d, i["type"], pre_append)
	}

	return result, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFilters2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["body-type"], _ = expandObjectCasbUserActivityMatchTenantExtractionFiltersBodyType2edl(d, i["body_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandObjectCasbUserActivityMatchTenantExtractionFiltersDirection2edl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName2edl(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectCasbUserActivityMatchTenantExtractionFiltersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["place"], _ = expandObjectCasbUserActivityMatchTenantExtractionFiltersPlace2edl(d, i["place"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersBodyType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersPlace2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionJq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbUserActivityMatch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectCasbUserActivityMatchId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("rules"); ok || d.HasChange("rules") {
		t, err := expandObjectCasbUserActivityMatchRules2edl(d, v, "rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rules"] = t
		}
	}

	if v, ok := d.GetOk("strategy"); ok || d.HasChange("strategy") {
		t, err := expandObjectCasbUserActivityMatchStrategy2edl(d, v, "strategy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strategy"] = t
		}
	}

	if v, ok := d.GetOk("tenant_extraction"); ok || d.HasChange("tenant_extraction") {
		t, err := expandObjectCasbUserActivityMatchTenantExtraction2edl(d, v, "tenant_extraction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-extraction"] = t
		}
	}

	return &obj, nil
}
