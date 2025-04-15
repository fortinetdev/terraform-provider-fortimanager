// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure CASB profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbProfileCreate,
		Read:   resourceObjectCasbProfileRead,
		Update: resourceObjectCasbProfileUpdate,
		Delete: resourceObjectCasbProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
									"attribute_filter": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"attribute_match": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
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
						"advanced_tenant_control": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"input": &schema.Schema{
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
									"name": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"custom_control": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute_filter": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"attribute_match": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
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
						"status": &schema.Schema{
							Type:     schema.TypeString,
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

func resourceObjectCasbProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCasbProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCasbProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileRead(d, m)
}

func resourceObjectCasbProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCasbProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileRead(d, m)
}

func resourceObjectCasbProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectCasbProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCasbProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationAccessRule(i["access-rule"], d, pre_append)
			tmp["access_rule"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-AccessRule")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advanced_tenant_control"
		if _, ok := i["advanced-tenant-control"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControl(i["advanced-tenant-control"], d, pre_append)
			tmp["advanced_tenant_control"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-AdvancedTenantControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom_control"
		if _, ok := i["custom-control"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControl(i["custom-control"], d, pre_append)
			tmp["custom_control"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-CustomControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control"
		if _, ok := i["domain-control"]; ok {
			v := flattenObjectCasbProfileSaasApplicationDomainControl(i["domain-control"], d, pre_append)
			tmp["domain_control"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-DomainControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control_domains"
		if _, ok := i["domain-control-domains"]; ok {
			v := flattenObjectCasbProfileSaasApplicationDomainControlDomains(i["domain-control-domains"], d, pre_append)
			tmp["domain_control_domains"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-DomainControlDomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectCasbProfileSaasApplicationLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbProfileSaasApplicationName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search"
		if _, ok := i["safe-search"]; ok {
			v := flattenObjectCasbProfileSaasApplicationSafeSearch(i["safe-search"], d, pre_append)
			tmp["safe_search"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-SafeSearch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search_control"
		if _, ok := i["safe-search-control"]; ok {
			v := flattenObjectCasbProfileSaasApplicationSafeSearchControl(i["safe-search-control"], d, pre_append)
			tmp["safe_search_control"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-SafeSearchControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectCasbProfileSaasApplicationStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control"
		if _, ok := i["tenant-control"]; ok {
			v := flattenObjectCasbProfileSaasApplicationTenantControl(i["tenant-control"], d, pre_append)
			tmp["tenant_control"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-TenantControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control_tenants"
		if _, ok := i["tenant-control-tenants"]; ok {
			v := flattenObjectCasbProfileSaasApplicationTenantControlTenants(i["tenant-control-tenants"], d, pre_append)
			tmp["tenant_control_tenants"] = fortiAPISubPartPatch(v, "ObjectCasbProfile-SaasApplication-TenantControlTenants")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationAccessRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationAccessRuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AccessRule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := i["attribute-filter"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(i["attribute-filter"], d, pre_append)
			tmp["attribute_filter"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AccessRule-AttributeFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := i["bypass"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAccessRuleBypass(i["bypass"], d, pre_append)
			tmp["bypass"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AccessRule-Bypass")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAccessRuleName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AccessRule-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAccessRule-AttributeFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := i["attribute-match"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(i["attribute-match"], d, pre_append)
			tmp["attribute_match"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAccessRule-AttributeFilter-AttributeMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAccessRule-AttributeFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAccessRuleBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationAccessRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute"
		if _, ok := i["attribute"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(i["attribute"], d, pre_append)
			tmp["attribute"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AdvancedTenantControl-Attribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControlName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AdvancedTenantControl-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input"
		if _, ok := i["input"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(i["input"], d, pre_append)
			tmp["input"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAdvancedTenantControl-Attribute-Input")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAdvancedTenantControl-Attribute-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationCustomControl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := i["attribute-filter"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilter(i["attribute-filter"], d, pre_append)
			tmp["attribute_filter"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-CustomControl-AttributeFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-CustomControl-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := i["option"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlOption(i["option"], d, pre_append)
			tmp["option"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-CustomControl-Option")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationCustomControl-AttributeFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := i["attribute-match"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(i["attribute-match"], d, pre_append)
			tmp["attribute_match"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationCustomControl-AttributeFilter-AttributeMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationCustomControl-AttributeFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationCustomControlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectCasbProfileSaasApplicationCustomControlOption(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationCustomControlOptionName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationCustomControl-Option-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := i["user-input"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlOptionUserInput(i["user-input"], d, pre_append)
			tmp["user_input"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationCustomControl-Option-UserInput")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationCustomControlOptionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationCustomControlOptionUserInput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationDomainControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationDomainControlDomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectCasbProfileSaasApplicationSafeSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationSafeSearchControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationTenantControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationTenantControlTenants(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectCasbProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenObjectCasbProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectCasbProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCasbProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("saas_application", flattenObjectCasbProfileSaasApplication(o["saas-application"], d, "saas_application")); err != nil {
			if vv, ok := fortiAPIPatch(o["saas-application"], "ObjectCasbProfile-SaasApplication"); ok {
				if err = d.Set("saas_application", vv); err != nil {
					return fmt.Errorf("Error reading saas_application: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading saas_application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("saas_application"); ok {
			if err = d.Set("saas_application", flattenObjectCasbProfileSaasApplication(o["saas-application"], d, "saas_application")); err != nil {
				if vv, ok := fortiAPIPatch(o["saas-application"], "ObjectCasbProfile-SaasApplication"); ok {
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

func flattenObjectCasbProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandObjectCasbProfileSaasApplicationAccessRule(d, i["access_rule"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["access-rule"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advanced_tenant_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectCasbProfileSaasApplicationAdvancedTenantControl(d, i["advanced_tenant_control"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["advanced-tenant-control"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectCasbProfileSaasApplicationCustomControl(d, i["custom_control"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["custom-control"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain-control"], _ = expandObjectCasbProfileSaasApplicationDomainControl(d, i["domain_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control_domains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain-control-domains"], _ = expandObjectCasbProfileSaasApplicationDomainControlDomains(d, i["domain_control_domains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectCasbProfileSaasApplicationLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["safe-search"], _ = expandObjectCasbProfileSaasApplicationSafeSearch(d, i["safe_search"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["safe-search-control"], _ = expandObjectCasbProfileSaasApplicationSafeSearchControl(d, i["safe_search_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectCasbProfileSaasApplicationStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tenant-control"], _ = expandObjectCasbProfileSaasApplicationTenantControl(d, i["tenant_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control_tenants"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tenant-control-tenants"], _ = expandObjectCasbProfileSaasApplicationTenantControlTenants(d, i["tenant_control_tenants"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationAccessRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectCasbProfileSaasApplicationAccessRuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d, i["attribute_filter"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute-filter"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bypass"], _ = expandObjectCasbProfileSaasApplicationAccessRuleBypass(d, i["bypass"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationAccessRuleName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-match"], _ = expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(d, i["attribute_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(d, i["attribute"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationAdvancedTenantControlName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input"], _ = expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(d, i["input"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationCustomControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectCasbProfileSaasApplicationCustomControlAttributeFilter(d, i["attribute_filter"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute-filter"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationCustomControlName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectCasbProfileSaasApplicationCustomControlOption(d, i["option"], pre_append)
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

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-match"], _ = expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(d, i["attribute_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectCasbProfileSaasApplicationCustomControlOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationCustomControlOptionName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-input"], _ = expandObjectCasbProfileSaasApplicationCustomControlOptionUserInput(d, i["user_input"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlOptionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlOptionUserInput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationDomainControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationDomainControlDomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectCasbProfileSaasApplicationSafeSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationSafeSearchControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationTenantControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationTenantControlTenants(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectCasbProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectCasbProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("saas_application"); ok || d.HasChange("saas_application") {
		t, err := expandObjectCasbProfileSaasApplication(d, v, "saas_application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saas-application"] = t
		}
	}

	return &obj, nil
}
