// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiGuard Web Filter settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterProfileFtgdWf() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterProfileFtgdWfUpdate,
		Read:   resourceObjectWebfilterProfileFtgdWfRead,
		Update: resourceObjectWebfilterProfileFtgdWfUpdate,
		Delete: resourceObjectWebfilterProfileFtgdWfDelete,

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
			"exempt_quota": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_usr_grp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"category": &schema.Schema{
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
						"override_replacemsg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"warn_duration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"warning_duration_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"warning_prompt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"max_quota_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ovrd": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"quota": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"duration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"override_replacemsg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"unit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"rate_crl_urls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_css_urls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_image_urls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_javascript_urls": &schema.Schema{
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

func resourceObjectWebfilterProfileFtgdWfUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterProfileFtgdWf(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileFtgdWf resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWebfilterProfileFtgdWf(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileFtgdWf resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWebfilterProfileFtgdWf")

	return resourceObjectWebfilterProfileFtgdWfRead(d, m)
}

func resourceObjectWebfilterProfileFtgdWfDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWebfilterProfileFtgdWf(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterProfileFtgdWf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterProfileFtgdWfRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWebfilterProfileFtgdWf(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileFtgdWf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterProfileFtgdWf(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileFtgdWf resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterProfileFtgdWfExemptQuota2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileFtgdWfFilters2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWebfilterProfileFtgdWfFiltersAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := i["auth-usr-grp"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp2edl(i["auth-usr-grp"], d, pre_append)
			tmp["auth_usr_grp"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-AuthUsrGrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersCategory2edl(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersLog2edl(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := i["override-replacemsg"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg2edl(i["override-replacemsg"], d, pre_append)
			tmp["override_replacemsg"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-OverrideReplacemsg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := i["warn-duration"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersWarnDuration2edl(i["warn-duration"], d, pre_append)
			tmp["warn_duration"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-WarnDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := i["warning-duration-type"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersWarningDurationType2edl(i["warning-duration-type"], d, pre_append)
			tmp["warning_duration_type"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-WarningDurationType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := i["warning-prompt"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersWarningPrompt2edl(i["warning-prompt"], d, pre_append)
			tmp["warning_prompt"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-WarningPrompt")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWebfilterProfileFtgdWfFiltersAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWebfilterProfileFtgdWfFiltersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersWarnDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersWarningDurationType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersWarningPrompt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfMaxQuotaTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileFtgdWfOvrd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileFtgdWfQuota2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaCategory2edl(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := i["duration"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaDuration2edl(i["duration"], d, pre_append)
			tmp["duration"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Duration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := i["override-replacemsg"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaOverrideReplacemsg2edl(i["override-replacemsg"], d, pre_append)
			tmp["override_replacemsg"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-OverrideReplacemsg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := i["unit"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaUnit2edl(i["unit"], d, pre_append)
			tmp["unit"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Unit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWebfilterProfileFtgdWfQuotaCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaOverrideReplacemsg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaUnit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfRateCrlUrls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfRateCssUrls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfRateImageUrls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfRateJavascriptUrls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterProfileFtgdWf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("exempt_quota", flattenObjectWebfilterProfileFtgdWfExemptQuota2edl(o["exempt-quota"], d, "exempt_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["exempt-quota"], "ObjectWebfilterProfileFtgdWf-ExemptQuota"); ok {
			if err = d.Set("exempt_quota", vv); err != nil {
				return fmt.Errorf("Error reading exempt_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exempt_quota: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("filters", flattenObjectWebfilterProfileFtgdWfFilters2edl(o["filters"], d, "filters")); err != nil {
			if vv, ok := fortiAPIPatch(o["filters"], "ObjectWebfilterProfileFtgdWf-Filters"); ok {
				if err = d.Set("filters", vv); err != nil {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filters"); ok {
			if err = d.Set("filters", flattenObjectWebfilterProfileFtgdWfFilters2edl(o["filters"], d, "filters")); err != nil {
				if vv, ok := fortiAPIPatch(o["filters"], "ObjectWebfilterProfileFtgdWf-Filters"); ok {
					if err = d.Set("filters", vv); err != nil {
						return fmt.Errorf("Error reading filters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			}
		}
	}

	if err = d.Set("max_quota_timeout", flattenObjectWebfilterProfileFtgdWfMaxQuotaTimeout2edl(o["max-quota-timeout"], d, "max_quota_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-quota-timeout"], "ObjectWebfilterProfileFtgdWf-MaxQuotaTimeout"); ok {
			if err = d.Set("max_quota_timeout", vv); err != nil {
				return fmt.Errorf("Error reading max_quota_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_quota_timeout: %v", err)
		}
	}

	if err = d.Set("options", flattenObjectWebfilterProfileFtgdWfOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ObjectWebfilterProfileFtgdWf-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("ovrd", flattenObjectWebfilterProfileFtgdWfOvrd2edl(o["ovrd"], d, "ovrd")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd"], "ObjectWebfilterProfileFtgdWf-Ovrd"); ok {
			if err = d.Set("ovrd", vv); err != nil {
				return fmt.Errorf("Error reading ovrd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("quota", flattenObjectWebfilterProfileFtgdWfQuota2edl(o["quota"], d, "quota")); err != nil {
			if vv, ok := fortiAPIPatch(o["quota"], "ObjectWebfilterProfileFtgdWf-Quota"); ok {
				if err = d.Set("quota", vv); err != nil {
					return fmt.Errorf("Error reading quota: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading quota: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("quota"); ok {
			if err = d.Set("quota", flattenObjectWebfilterProfileFtgdWfQuota2edl(o["quota"], d, "quota")); err != nil {
				if vv, ok := fortiAPIPatch(o["quota"], "ObjectWebfilterProfileFtgdWf-Quota"); ok {
					if err = d.Set("quota", vv); err != nil {
						return fmt.Errorf("Error reading quota: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading quota: %v", err)
				}
			}
		}
	}

	if err = d.Set("rate_crl_urls", flattenObjectWebfilterProfileFtgdWfRateCrlUrls2edl(o["rate-crl-urls"], d, "rate_crl_urls")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-crl-urls"], "ObjectWebfilterProfileFtgdWf-RateCrlUrls"); ok {
			if err = d.Set("rate_crl_urls", vv); err != nil {
				return fmt.Errorf("Error reading rate_crl_urls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_crl_urls: %v", err)
		}
	}

	if err = d.Set("rate_css_urls", flattenObjectWebfilterProfileFtgdWfRateCssUrls2edl(o["rate-css-urls"], d, "rate_css_urls")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-css-urls"], "ObjectWebfilterProfileFtgdWf-RateCssUrls"); ok {
			if err = d.Set("rate_css_urls", vv); err != nil {
				return fmt.Errorf("Error reading rate_css_urls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_css_urls: %v", err)
		}
	}

	if err = d.Set("rate_image_urls", flattenObjectWebfilterProfileFtgdWfRateImageUrls2edl(o["rate-image-urls"], d, "rate_image_urls")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-image-urls"], "ObjectWebfilterProfileFtgdWf-RateImageUrls"); ok {
			if err = d.Set("rate_image_urls", vv); err != nil {
				return fmt.Errorf("Error reading rate_image_urls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_image_urls: %v", err)
		}
	}

	if err = d.Set("rate_javascript_urls", flattenObjectWebfilterProfileFtgdWfRateJavascriptUrls2edl(o["rate-javascript-urls"], d, "rate_javascript_urls")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-javascript-urls"], "ObjectWebfilterProfileFtgdWf-RateJavascriptUrls"); ok {
			if err = d.Set("rate_javascript_urls", vv); err != nil {
				return fmt.Errorf("Error reading rate_javascript_urls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_javascript_urls: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterProfileFtgdWfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterProfileFtgdWfExemptQuota2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectWebfilterProfileFtgdWfFilters2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectWebfilterProfileFtgdWfFiltersAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-usr-grp"], _ = expandObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp2edl(d, i["auth_usr_grp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandObjectWebfilterProfileFtgdWfFiltersCategory2edl(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWebfilterProfileFtgdWfFiltersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectWebfilterProfileFtgdWfFiltersLog2edl(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-replacemsg"], _ = expandObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg2edl(d, i["override_replacemsg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warn-duration"], _ = expandObjectWebfilterProfileFtgdWfFiltersWarnDuration2edl(d, i["warn_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warning-duration-type"], _ = expandObjectWebfilterProfileFtgdWfFiltersWarningDurationType2edl(d, i["warning_duration_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warning-prompt"], _ = expandObjectWebfilterProfileFtgdWfFiltersWarningPrompt2edl(d, i["warning_prompt"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersWarnDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersWarningDurationType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersWarningPrompt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfMaxQuotaTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileFtgdWfOvrd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectWebfilterProfileFtgdWfQuota2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandObjectWebfilterProfileFtgdWfQuotaCategory2edl(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["duration"], _ = expandObjectWebfilterProfileFtgdWfQuotaDuration2edl(d, i["duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWebfilterProfileFtgdWfQuotaId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-replacemsg"], _ = expandObjectWebfilterProfileFtgdWfQuotaOverrideReplacemsg2edl(d, i["override_replacemsg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectWebfilterProfileFtgdWfQuotaType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["unit"], _ = expandObjectWebfilterProfileFtgdWfQuotaUnit2edl(d, i["unit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectWebfilterProfileFtgdWfQuotaValue2edl(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaOverrideReplacemsg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaUnit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfRateCrlUrls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfRateCssUrls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfRateImageUrls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfRateJavascriptUrls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterProfileFtgdWf(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("exempt_quota"); ok || d.HasChange("exempt_quota") {
		t, err := expandObjectWebfilterProfileFtgdWfExemptQuota2edl(d, v, "exempt_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exempt-quota"] = t
		}
	}

	if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
		t, err := expandObjectWebfilterProfileFtgdWfFilters2edl(d, v, "filters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filters"] = t
		}
	}

	if v, ok := d.GetOk("max_quota_timeout"); ok || d.HasChange("max_quota_timeout") {
		t, err := expandObjectWebfilterProfileFtgdWfMaxQuotaTimeout2edl(d, v, "max_quota_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-quota-timeout"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectWebfilterProfileFtgdWfOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("ovrd"); ok || d.HasChange("ovrd") {
		t, err := expandObjectWebfilterProfileFtgdWfOvrd2edl(d, v, "ovrd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd"] = t
		}
	}

	if v, ok := d.GetOk("quota"); ok || d.HasChange("quota") {
		t, err := expandObjectWebfilterProfileFtgdWfQuota2edl(d, v, "quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quota"] = t
		}
	}

	if v, ok := d.GetOk("rate_crl_urls"); ok || d.HasChange("rate_crl_urls") {
		t, err := expandObjectWebfilterProfileFtgdWfRateCrlUrls2edl(d, v, "rate_crl_urls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-crl-urls"] = t
		}
	}

	if v, ok := d.GetOk("rate_css_urls"); ok || d.HasChange("rate_css_urls") {
		t, err := expandObjectWebfilterProfileFtgdWfRateCssUrls2edl(d, v, "rate_css_urls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-css-urls"] = t
		}
	}

	if v, ok := d.GetOk("rate_image_urls"); ok || d.HasChange("rate_image_urls") {
		t, err := expandObjectWebfilterProfileFtgdWfRateImageUrls2edl(d, v, "rate_image_urls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-image-urls"] = t
		}
	}

	if v, ok := d.GetOk("rate_javascript_urls"); ok || d.HasChange("rate_javascript_urls") {
		t, err := expandObjectWebfilterProfileFtgdWfRateJavascriptUrls2edl(d, v, "rate_javascript_urls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-javascript-urls"] = t
		}
	}

	return &obj, nil
}
