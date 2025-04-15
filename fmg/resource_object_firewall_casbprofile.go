// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFirewall CasbProfile

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallCasbProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallCasbProfileCreate,
		Read:   resourceObjectFirewallCasbProfileRead,
		Update: resourceObjectFirewallCasbProfileUpdate,
		Delete: resourceObjectFirewallCasbProfileDelete,

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
			},
			"saas_application": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_rule": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bypass": &schema.Schema{
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
						"custom_control": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"option": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"user_input": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"domain_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domain_control_domains": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
						"safe_search": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"safe_search_control": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"tenant_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tenant_control_tenants": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceObjectFirewallCasbProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallCasbProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallCasbProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallCasbProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallCasbProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallCasbProfileRead(d, m)
}

func resourceObjectFirewallCasbProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallCasbProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallCasbProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallCasbProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallCasbProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallCasbProfileRead(d, m)
}

func resourceObjectFirewallCasbProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallCasbProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallCasbProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallCasbProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallCasbProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallCasbProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallCasbProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallCasbProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallCasbProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_rule"
		if _, ok := i["access-rule"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationAccessRule(i["access-rule"], d, pre_append)
			tmp["access_rule"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfile-SaasApplication-AccessRule")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom_control"
		if _, ok := i["custom-control"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControl(i["custom-control"], d, pre_append)
			tmp["custom_control"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfile-SaasApplication-CustomControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control"
		if _, ok := i["domain-control"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationDomainControl(i["domain-control"], d, pre_append)
			tmp["domain_control"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfile-SaasApplication-DomainControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control_domains"
		if _, ok := i["domain-control-domains"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationDomainControlDomains(i["domain-control-domains"], d, pre_append)
			tmp["domain_control_domains"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfile-SaasApplication-DomainControlDomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfile-SaasApplication-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfile-SaasApplication-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search"
		if _, ok := i["safe-search"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationSafeSearch(i["safe-search"], d, pre_append)
			tmp["safe_search"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfile-SaasApplication-SafeSearch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search_control"
		if _, ok := i["safe-search-control"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationSafeSearchControl(i["safe-search-control"], d, pre_append)
			tmp["safe_search_control"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfile-SaasApplication-SafeSearchControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control"
		if _, ok := i["tenant-control"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationTenantControl(i["tenant-control"], d, pre_append)
			tmp["tenant_control"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfile-SaasApplication-TenantControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control_tenants"
		if _, ok := i["tenant-control-tenants"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationTenantControlTenants(i["tenant-control-tenants"], d, pre_append)
			tmp["tenant_control_tenants"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfile-SaasApplication-TenantControlTenants")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallCasbProfileSaasApplicationAccessRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallCasbProfileSaasApplicationAccessRuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplication-AccessRule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := i["bypass"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationAccessRuleBypass(i["bypass"], d, pre_append)
			tmp["bypass"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplication-AccessRule-Bypass")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationAccessRuleName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplication-AccessRule-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallCasbProfileSaasApplicationAccessRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationAccessRuleBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallCasbProfileSaasApplicationAccessRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControlName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplication-CustomControl-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := i["option"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControlOption(i["option"], d, pre_append)
			tmp["option"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplication-CustomControl-Option")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOption(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplicationCustomControl-Option-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := i["user-input"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput(i["user-input"], d, pre_append)
			tmp["user_input"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplicationCustomControl-Option-UserInput")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallCasbProfileSaasApplicationDomainControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationDomainControlDomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallCasbProfileSaasApplicationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationSafeSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationSafeSearchControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallCasbProfileSaasApplicationTenantControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationTenantControlTenants(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFirewallCasbProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenObjectFirewallCasbProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallCasbProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("saas_application", flattenObjectFirewallCasbProfileSaasApplication(o["saas-application"], d, "saas_application")); err != nil {
			if vv, ok := fortiAPIPatch(o["saas-application"], "ObjectFirewallCasbProfile-SaasApplication"); ok {
				if err = d.Set("saas_application", vv); err != nil {
					return fmt.Errorf("Error reading saas_application: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading saas_application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("saas_application"); ok {
			if err = d.Set("saas_application", flattenObjectFirewallCasbProfileSaasApplication(o["saas-application"], d, "saas_application")); err != nil {
				if vv, ok := fortiAPIPatch(o["saas-application"], "ObjectFirewallCasbProfile-SaasApplication"); ok {
					if err = d.Set("saas_application", vv); err != nil {
						return fmt.Errorf("Error reading saas_application: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading saas_application: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectFirewallCasbProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallCasbProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_rule"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallCasbProfileSaasApplicationAccessRule(d, i["access_rule"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["access-rule"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallCasbProfileSaasApplicationCustomControl(d, i["custom_control"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["custom-control"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain-control"], _ = expandObjectFirewallCasbProfileSaasApplicationDomainControl(d, i["domain_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control_domains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain-control-domains"], _ = expandObjectFirewallCasbProfileSaasApplicationDomainControlDomains(d, i["domain_control_domains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectFirewallCasbProfileSaasApplicationLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFirewallCasbProfileSaasApplicationName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["safe-search"], _ = expandObjectFirewallCasbProfileSaasApplicationSafeSearch(d, i["safe_search"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["safe-search-control"], _ = expandObjectFirewallCasbProfileSaasApplicationSafeSearchControl(d, i["safe_search_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tenant-control"], _ = expandObjectFirewallCasbProfileSaasApplicationTenantControl(d, i["tenant_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control_tenants"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tenant-control-tenants"], _ = expandObjectFirewallCasbProfileSaasApplicationTenantControlTenants(d, i["tenant_control_tenants"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallCasbProfileSaasApplicationAccessRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectFirewallCasbProfileSaasApplicationAccessRuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bypass"], _ = expandObjectFirewallCasbProfileSaasApplicationAccessRuleBypass(d, i["bypass"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFirewallCasbProfileSaasApplicationAccessRuleName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallCasbProfileSaasApplicationAccessRuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationAccessRuleBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallCasbProfileSaasApplicationAccessRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallCasbProfileSaasApplicationCustomControlName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallCasbProfileSaasApplicationCustomControlOption(d, i["option"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["option"] = t
			}
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-input"], _ = expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput(d, i["user_input"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallCasbProfileSaasApplicationDomainControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationDomainControlDomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallCasbProfileSaasApplicationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationSafeSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationSafeSearchControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallCasbProfileSaasApplicationTenantControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationTenantControlTenants(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFirewallCasbProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallCasbProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("saas_application"); ok || d.HasChange("saas_application") {
		t, err := expandObjectFirewallCasbProfileSaasApplication(d, v, "saas_application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saas-application"] = t
		}
	}

	return &obj, nil
}
