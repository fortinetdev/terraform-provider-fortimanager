// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: AntiPhishing profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterProfileAntiphish() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterProfileAntiphishUpdate,
		Read:   resourceObjectWebfilterProfileAntiphishRead,
		Update: resourceObjectWebfilterProfileAntiphishUpdate,
		Delete: resourceObjectWebfilterProfileAntiphishDelete,

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
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_basic_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_uri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_username_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_patterns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inspection_entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiguard_category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ldap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_body_len": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
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

func resourceObjectWebfilterProfileAntiphishUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterProfileAntiphish(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileAntiphish resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWebfilterProfileAntiphish(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileAntiphish resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWebfilterProfileAntiphish")

	return resourceObjectWebfilterProfileAntiphishRead(d, m)
}

func resourceObjectWebfilterProfileAntiphishDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWebfilterProfileAntiphish(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterProfileAntiphish resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterProfileAntiphishRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWebfilterProfileAntiphish(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileAntiphish resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterProfileAntiphish(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileAntiphish resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterProfileAntiphishAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCheckBasicAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCheckUri2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCheckUsernameOnly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCustomPatterns2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWebfilterProfileAntiphishCustomPatternsCategory2edl(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-CustomPatterns-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectWebfilterProfileAntiphishCustomPatternsPattern2edl(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-CustomPatterns-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectWebfilterProfileAntiphishCustomPatternsType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-CustomPatterns-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectWebfilterProfileAntiphishCustomPatternsCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCustomPatternsPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCustomPatternsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishDefaultAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishDomainController2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWebfilterProfileAntiphishInspectionEntries2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWebfilterProfileAntiphishInspectionEntriesAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-InspectionEntries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := i["fortiguard-category"]; ok {
			v := flattenObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory2edl(i["fortiguard-category"], d, pre_append)
			tmp["fortiguard_category"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-InspectionEntries-FortiguardCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWebfilterProfileAntiphishInspectionEntriesName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-InspectionEntries-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectWebfilterProfileAntiphishInspectionEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileAntiphishInspectionEntriesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishLdap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWebfilterProfileAntiphishMaxBodyLen2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterProfileAntiphish(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("authentication", flattenObjectWebfilterProfileAntiphishAuthentication2edl(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "ObjectWebfilterProfileAntiphish-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("check_basic_auth", flattenObjectWebfilterProfileAntiphishCheckBasicAuth2edl(o["check-basic-auth"], d, "check_basic_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-basic-auth"], "ObjectWebfilterProfileAntiphish-CheckBasicAuth"); ok {
			if err = d.Set("check_basic_auth", vv); err != nil {
				return fmt.Errorf("Error reading check_basic_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_basic_auth: %v", err)
		}
	}

	if err = d.Set("check_uri", flattenObjectWebfilterProfileAntiphishCheckUri2edl(o["check-uri"], d, "check_uri")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-uri"], "ObjectWebfilterProfileAntiphish-CheckUri"); ok {
			if err = d.Set("check_uri", vv); err != nil {
				return fmt.Errorf("Error reading check_uri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_uri: %v", err)
		}
	}

	if err = d.Set("check_username_only", flattenObjectWebfilterProfileAntiphishCheckUsernameOnly2edl(o["check-username-only"], d, "check_username_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-username-only"], "ObjectWebfilterProfileAntiphish-CheckUsernameOnly"); ok {
			if err = d.Set("check_username_only", vv); err != nil {
				return fmt.Errorf("Error reading check_username_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_username_only: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_patterns", flattenObjectWebfilterProfileAntiphishCustomPatterns2edl(o["custom-patterns"], d, "custom_patterns")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-patterns"], "ObjectWebfilterProfileAntiphish-CustomPatterns"); ok {
				if err = d.Set("custom_patterns", vv); err != nil {
					return fmt.Errorf("Error reading custom_patterns: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_patterns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_patterns"); ok {
			if err = d.Set("custom_patterns", flattenObjectWebfilterProfileAntiphishCustomPatterns2edl(o["custom-patterns"], d, "custom_patterns")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-patterns"], "ObjectWebfilterProfileAntiphish-CustomPatterns"); ok {
					if err = d.Set("custom_patterns", vv); err != nil {
						return fmt.Errorf("Error reading custom_patterns: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_patterns: %v", err)
				}
			}
		}
	}

	if err = d.Set("default_action", flattenObjectWebfilterProfileAntiphishDefaultAction2edl(o["default-action"], d, "default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-action"], "ObjectWebfilterProfileAntiphish-DefaultAction"); ok {
			if err = d.Set("default_action", vv); err != nil {
				return fmt.Errorf("Error reading default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_action: %v", err)
		}
	}

	if err = d.Set("domain_controller", flattenObjectWebfilterProfileAntiphishDomainController2edl(o["domain-controller"], d, "domain_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-controller"], "ObjectWebfilterProfileAntiphish-DomainController"); ok {
			if err = d.Set("domain_controller", vv); err != nil {
				return fmt.Errorf("Error reading domain_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_controller: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("inspection_entries", flattenObjectWebfilterProfileAntiphishInspectionEntries2edl(o["inspection-entries"], d, "inspection_entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["inspection-entries"], "ObjectWebfilterProfileAntiphish-InspectionEntries"); ok {
				if err = d.Set("inspection_entries", vv); err != nil {
					return fmt.Errorf("Error reading inspection_entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading inspection_entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("inspection_entries"); ok {
			if err = d.Set("inspection_entries", flattenObjectWebfilterProfileAntiphishInspectionEntries2edl(o["inspection-entries"], d, "inspection_entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["inspection-entries"], "ObjectWebfilterProfileAntiphish-InspectionEntries"); ok {
					if err = d.Set("inspection_entries", vv); err != nil {
						return fmt.Errorf("Error reading inspection_entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading inspection_entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("ldap", flattenObjectWebfilterProfileAntiphishLdap2edl(o["ldap"], d, "ldap")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap"], "ObjectWebfilterProfileAntiphish-Ldap"); ok {
			if err = d.Set("ldap", vv); err != nil {
				return fmt.Errorf("Error reading ldap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap: %v", err)
		}
	}

	if err = d.Set("max_body_len", flattenObjectWebfilterProfileAntiphishMaxBodyLen2edl(o["max-body-len"], d, "max_body_len")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-body-len"], "ObjectWebfilterProfileAntiphish-MaxBodyLen"); ok {
			if err = d.Set("max_body_len", vv); err != nil {
				return fmt.Errorf("Error reading max_body_len: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_body_len: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectWebfilterProfileAntiphishStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectWebfilterProfileAntiphish-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterProfileAntiphishFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterProfileAntiphishAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCheckBasicAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCheckUri2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCheckUsernameOnly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCustomPatterns2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["category"], _ = expandObjectWebfilterProfileAntiphishCustomPatternsCategory2edl(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandObjectWebfilterProfileAntiphishCustomPatternsPattern2edl(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectWebfilterProfileAntiphishCustomPatternsType2edl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterProfileAntiphishCustomPatternsCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCustomPatternsPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCustomPatternsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishDefaultAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishDomainController2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWebfilterProfileAntiphishInspectionEntries2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectWebfilterProfileAntiphishInspectionEntriesAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortiguard-category"], _ = expandObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory2edl(d, i["fortiguard_category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWebfilterProfileAntiphishInspectionEntriesName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterProfileAntiphishInspectionEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileAntiphishInspectionEntriesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishLdap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWebfilterProfileAntiphishMaxBodyLen2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterProfileAntiphish(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandObjectWebfilterProfileAntiphishAuthentication2edl(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("check_basic_auth"); ok || d.HasChange("check_basic_auth") {
		t, err := expandObjectWebfilterProfileAntiphishCheckBasicAuth2edl(d, v, "check_basic_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-basic-auth"] = t
		}
	}

	if v, ok := d.GetOk("check_uri"); ok || d.HasChange("check_uri") {
		t, err := expandObjectWebfilterProfileAntiphishCheckUri2edl(d, v, "check_uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-uri"] = t
		}
	}

	if v, ok := d.GetOk("check_username_only"); ok || d.HasChange("check_username_only") {
		t, err := expandObjectWebfilterProfileAntiphishCheckUsernameOnly2edl(d, v, "check_username_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-username-only"] = t
		}
	}

	if v, ok := d.GetOk("custom_patterns"); ok || d.HasChange("custom_patterns") {
		t, err := expandObjectWebfilterProfileAntiphishCustomPatterns2edl(d, v, "custom_patterns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-patterns"] = t
		}
	}

	if v, ok := d.GetOk("default_action"); ok || d.HasChange("default_action") {
		t, err := expandObjectWebfilterProfileAntiphishDefaultAction2edl(d, v, "default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-action"] = t
		}
	}

	if v, ok := d.GetOk("domain_controller"); ok || d.HasChange("domain_controller") {
		t, err := expandObjectWebfilterProfileAntiphishDomainController2edl(d, v, "domain_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-controller"] = t
		}
	}

	if v, ok := d.GetOk("inspection_entries"); ok || d.HasChange("inspection_entries") {
		t, err := expandObjectWebfilterProfileAntiphishInspectionEntries2edl(d, v, "inspection_entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-entries"] = t
		}
	}

	if v, ok := d.GetOk("ldap"); ok || d.HasChange("ldap") {
		t, err := expandObjectWebfilterProfileAntiphishLdap2edl(d, v, "ldap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap"] = t
		}
	}

	if v, ok := d.GetOk("max_body_len"); ok || d.HasChange("max_body_len") {
		t, err := expandObjectWebfilterProfileAntiphishMaxBodyLen2edl(d, v, "max_body_len")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-body-len"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectWebfilterProfileAntiphishStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
