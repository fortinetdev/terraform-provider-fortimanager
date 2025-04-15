// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB profile SaaS application.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbProfileSaasApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbProfileSaasApplicationCreate,
		Read:   resourceObjectCasbProfileSaasApplicationRead,
		Update: resourceObjectCasbProfileSaasApplicationUpdate,
		Delete: resourceObjectCasbProfileSaasApplicationDelete,

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
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectCasbProfileSaasApplicationCreate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectCasbProfileSaasApplication(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplication resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCasbProfileSaasApplication(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplication resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileSaasApplicationRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectCasbProfileSaasApplication(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplication resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbProfileSaasApplication(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplication resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileSaasApplicationRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbProfileSaasApplication(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbProfileSaasApplication resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbProfileSaasApplicationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCasbProfileSaasApplication(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplication resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbProfileSaasApplication(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplication resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbProfileSaasApplicationAccessRule2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationAccessRuleAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AccessRule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := i["attribute-filter"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilter2edl(i["attribute-filter"], d, pre_append)
			tmp["attribute_filter"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AccessRule-AttributeFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := i["bypass"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAccessRuleBypass2edl(i["bypass"], d, pre_append)
			tmp["bypass"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AccessRule-Bypass")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAccessRuleName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AccessRule-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilter2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAccessRule-AttributeFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := i["attribute-match"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch2edl(i["attribute-match"], d, pre_append)
			tmp["attribute_match"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAccessRule-AttributeFilter-AttributeMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAccessRule-AttributeFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAccessRuleBypass2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationAccessRuleName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControl2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute2edl(i["attribute"], d, pre_append)
			tmp["attribute"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AdvancedTenantControl-Attribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControlName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-AdvancedTenantControl-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput2edl(i["input"], d, pre_append)
			tmp["input"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAdvancedTenantControl-Attribute-Input")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAdvancedTenantControl-Attribute-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationCustomControl2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilter2edl(i["attribute-filter"], d, pre_append)
			tmp["attribute_filter"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-CustomControl-AttributeFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-CustomControl-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := i["option"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlOption2edl(i["option"], d, pre_append)
			tmp["option"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplication-CustomControl-Option")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilter2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationCustomControl-AttributeFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := i["attribute-match"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch2edl(i["attribute-match"], d, pre_append)
			tmp["attribute_match"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationCustomControl-AttributeFilter-AttributeMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationCustomControl-AttributeFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationCustomControlName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectCasbProfileSaasApplicationCustomControlOption2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbProfileSaasApplicationCustomControlOptionName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationCustomControl-Option-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := i["user-input"]; ok {
			v := flattenObjectCasbProfileSaasApplicationCustomControlOptionUserInput2edl(i["user-input"], d, pre_append)
			tmp["user_input"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationCustomControl-Option-UserInput")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationCustomControlOptionName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationCustomControlOptionUserInput2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationDomainControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationDomainControlDomains2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectCasbProfileSaasApplicationSafeSearch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationSafeSearchControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationTenantControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationTenantControlTenants2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectCasbProfileSaasApplication(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("access_rule", flattenObjectCasbProfileSaasApplicationAccessRule2edl(o["access-rule"], d, "access_rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["access-rule"], "ObjectCasbProfileSaasApplication-AccessRule"); ok {
				if err = d.Set("access_rule", vv); err != nil {
					return fmt.Errorf("Error reading access_rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading access_rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("access_rule"); ok {
			if err = d.Set("access_rule", flattenObjectCasbProfileSaasApplicationAccessRule2edl(o["access-rule"], d, "access_rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["access-rule"], "ObjectCasbProfileSaasApplication-AccessRule"); ok {
					if err = d.Set("access_rule", vv); err != nil {
						return fmt.Errorf("Error reading access_rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading access_rule: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("advanced_tenant_control", flattenObjectCasbProfileSaasApplicationAdvancedTenantControl2edl(o["advanced-tenant-control"], d, "advanced_tenant_control")); err != nil {
			if vv, ok := fortiAPIPatch(o["advanced-tenant-control"], "ObjectCasbProfileSaasApplication-AdvancedTenantControl"); ok {
				if err = d.Set("advanced_tenant_control", vv); err != nil {
					return fmt.Errorf("Error reading advanced_tenant_control: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading advanced_tenant_control: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("advanced_tenant_control"); ok {
			if err = d.Set("advanced_tenant_control", flattenObjectCasbProfileSaasApplicationAdvancedTenantControl2edl(o["advanced-tenant-control"], d, "advanced_tenant_control")); err != nil {
				if vv, ok := fortiAPIPatch(o["advanced-tenant-control"], "ObjectCasbProfileSaasApplication-AdvancedTenantControl"); ok {
					if err = d.Set("advanced_tenant_control", vv); err != nil {
						return fmt.Errorf("Error reading advanced_tenant_control: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading advanced_tenant_control: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("custom_control", flattenObjectCasbProfileSaasApplicationCustomControl2edl(o["custom-control"], d, "custom_control")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-control"], "ObjectCasbProfileSaasApplication-CustomControl"); ok {
				if err = d.Set("custom_control", vv); err != nil {
					return fmt.Errorf("Error reading custom_control: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_control: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_control"); ok {
			if err = d.Set("custom_control", flattenObjectCasbProfileSaasApplicationCustomControl2edl(o["custom-control"], d, "custom_control")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-control"], "ObjectCasbProfileSaasApplication-CustomControl"); ok {
					if err = d.Set("custom_control", vv); err != nil {
						return fmt.Errorf("Error reading custom_control: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_control: %v", err)
				}
			}
		}
	}

	if err = d.Set("domain_control", flattenObjectCasbProfileSaasApplicationDomainControl2edl(o["domain-control"], d, "domain_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-control"], "ObjectCasbProfileSaasApplication-DomainControl"); ok {
			if err = d.Set("domain_control", vv); err != nil {
				return fmt.Errorf("Error reading domain_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_control: %v", err)
		}
	}

	if err = d.Set("domain_control_domains", flattenObjectCasbProfileSaasApplicationDomainControlDomains2edl(o["domain-control-domains"], d, "domain_control_domains")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-control-domains"], "ObjectCasbProfileSaasApplication-DomainControlDomains"); ok {
			if err = d.Set("domain_control_domains", vv); err != nil {
				return fmt.Errorf("Error reading domain_control_domains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_control_domains: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectCasbProfileSaasApplicationLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectCasbProfileSaasApplication-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCasbProfileSaasApplicationName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbProfileSaasApplication-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("safe_search", flattenObjectCasbProfileSaasApplicationSafeSearch2edl(o["safe-search"], d, "safe_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["safe-search"], "ObjectCasbProfileSaasApplication-SafeSearch"); ok {
			if err = d.Set("safe_search", vv); err != nil {
				return fmt.Errorf("Error reading safe_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safe_search: %v", err)
		}
	}

	if err = d.Set("safe_search_control", flattenObjectCasbProfileSaasApplicationSafeSearchControl2edl(o["safe-search-control"], d, "safe_search_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["safe-search-control"], "ObjectCasbProfileSaasApplication-SafeSearchControl"); ok {
			if err = d.Set("safe_search_control", vv); err != nil {
				return fmt.Errorf("Error reading safe_search_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safe_search_control: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectCasbProfileSaasApplicationStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectCasbProfileSaasApplication-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tenant_control", flattenObjectCasbProfileSaasApplicationTenantControl2edl(o["tenant-control"], d, "tenant_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant-control"], "ObjectCasbProfileSaasApplication-TenantControl"); ok {
			if err = d.Set("tenant_control", vv); err != nil {
				return fmt.Errorf("Error reading tenant_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant_control: %v", err)
		}
	}

	if err = d.Set("tenant_control_tenants", flattenObjectCasbProfileSaasApplicationTenantControlTenants2edl(o["tenant-control-tenants"], d, "tenant_control_tenants")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant-control-tenants"], "ObjectCasbProfileSaasApplication-TenantControlTenants"); ok {
			if err = d.Set("tenant_control_tenants", vv); err != nil {
				return fmt.Errorf("Error reading tenant_control_tenants: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant_control_tenants: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbProfileSaasApplicationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbProfileSaasApplicationAccessRule2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectCasbProfileSaasApplicationAccessRuleAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilter2edl(d, i["attribute_filter"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute-filter"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bypass"], _ = expandObjectCasbProfileSaasApplicationAccessRuleBypass2edl(d, i["bypass"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationAccessRuleName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-match"], _ = expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch2edl(d, i["attribute_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId2edl(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleBypass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute2edl(d, i["attribute"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationAdvancedTenantControlName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["input"], _ = expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput2edl(d, i["input"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationCustomControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandObjectCasbProfileSaasApplicationCustomControlAttributeFilter2edl(d, i["attribute_filter"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute-filter"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationCustomControlName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectCasbProfileSaasApplicationCustomControlOption2edl(d, i["option"], pre_append)
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

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-match"], _ = expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch2edl(d, i["attribute_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterId2edl(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectCasbProfileSaasApplicationCustomControlOption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationCustomControlOptionName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-input"], _ = expandObjectCasbProfileSaasApplicationCustomControlOptionUserInput2edl(d, i["user_input"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlOptionName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlOptionUserInput2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationDomainControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationDomainControlDomains2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectCasbProfileSaasApplicationSafeSearch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationSafeSearchControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationTenantControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationTenantControlTenants2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectCasbProfileSaasApplication(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_rule"); ok || d.HasChange("access_rule") {
		t, err := expandObjectCasbProfileSaasApplicationAccessRule2edl(d, v, "access_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-rule"] = t
		}
	}

	if v, ok := d.GetOk("advanced_tenant_control"); ok || d.HasChange("advanced_tenant_control") {
		t, err := expandObjectCasbProfileSaasApplicationAdvancedTenantControl2edl(d, v, "advanced_tenant_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advanced-tenant-control"] = t
		}
	}

	if v, ok := d.GetOk("custom_control"); ok || d.HasChange("custom_control") {
		t, err := expandObjectCasbProfileSaasApplicationCustomControl2edl(d, v, "custom_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-control"] = t
		}
	}

	if v, ok := d.GetOk("domain_control"); ok || d.HasChange("domain_control") {
		t, err := expandObjectCasbProfileSaasApplicationDomainControl2edl(d, v, "domain_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-control"] = t
		}
	}

	if v, ok := d.GetOk("domain_control_domains"); ok || d.HasChange("domain_control_domains") {
		t, err := expandObjectCasbProfileSaasApplicationDomainControlDomains2edl(d, v, "domain_control_domains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-control-domains"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectCasbProfileSaasApplicationLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbProfileSaasApplicationName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("safe_search"); ok || d.HasChange("safe_search") {
		t, err := expandObjectCasbProfileSaasApplicationSafeSearch2edl(d, v, "safe_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search"] = t
		}
	}

	if v, ok := d.GetOk("safe_search_control"); ok || d.HasChange("safe_search_control") {
		t, err := expandObjectCasbProfileSaasApplicationSafeSearchControl2edl(d, v, "safe_search_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search-control"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectCasbProfileSaasApplicationStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tenant_control"); ok || d.HasChange("tenant_control") {
		t, err := expandObjectCasbProfileSaasApplicationTenantControl2edl(d, v, "tenant_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-control"] = t
		}
	}

	if v, ok := d.GetOk("tenant_control_tenants"); ok || d.HasChange("tenant_control_tenants") {
		t, err := expandObjectCasbProfileSaasApplicationTenantControlTenants2edl(d, v, "tenant_control_tenants")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-control-tenants"] = t
		}
	}

	return &obj, nil
}
