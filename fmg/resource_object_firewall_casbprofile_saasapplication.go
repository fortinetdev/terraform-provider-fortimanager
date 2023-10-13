// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFirewall CasbProfileSaasApplication

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallCasbProfileSaasApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallCasbProfileSaasApplicationCreate,
		Read:   resourceObjectFirewallCasbProfileSaasApplicationRead,
		Update: resourceObjectFirewallCasbProfileSaasApplicationUpdate,
		Delete: resourceObjectFirewallCasbProfileSaasApplicationDelete,

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
			"casb_profile": &schema.Schema{
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

func resourceObjectFirewallCasbProfileSaasApplicationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	casb_profile := d.Get("casb_profile").(string)
	paradict["casb_profile"] = casb_profile

	obj, err := getObjectObjectFirewallCasbProfileSaasApplication(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallCasbProfileSaasApplication resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallCasbProfileSaasApplication(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallCasbProfileSaasApplication resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallCasbProfileSaasApplicationRead(d, m)
}

func resourceObjectFirewallCasbProfileSaasApplicationUpdate(d *schema.ResourceData, m interface{}) error {
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

	casb_profile := d.Get("casb_profile").(string)
	paradict["casb_profile"] = casb_profile

	obj, err := getObjectObjectFirewallCasbProfileSaasApplication(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallCasbProfileSaasApplication resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallCasbProfileSaasApplication(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallCasbProfileSaasApplication resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallCasbProfileSaasApplicationRead(d, m)
}

func resourceObjectFirewallCasbProfileSaasApplicationDelete(d *schema.ResourceData, m interface{}) error {
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

	casb_profile := d.Get("casb_profile").(string)
	paradict["casb_profile"] = casb_profile

	err = c.DeleteObjectFirewallCasbProfileSaasApplication(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallCasbProfileSaasApplication resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallCasbProfileSaasApplicationRead(d *schema.ResourceData, m interface{}) error {
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

	casb_profile := d.Get("casb_profile").(string)
	if casb_profile == "" {
		casb_profile = importOptionChecking(m.(*FortiClient).Cfg, "casb_profile")
		if casb_profile == "" {
			return fmt.Errorf("Parameter casb_profile is missing")
		}
		if err = d.Set("casb_profile", casb_profile); err != nil {
			return fmt.Errorf("Error set params casb_profile: %v", err)
		}
	}
	paradict["casb_profile"] = casb_profile

	o, err := c.ReadObjectFirewallCasbProfileSaasApplication(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallCasbProfileSaasApplication resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallCasbProfileSaasApplication(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallCasbProfileSaasApplication resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallCasbProfileSaasApplicationAccessRule2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallCasbProfileSaasApplicationAccessRuleAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplication-AccessRule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := i["bypass"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationAccessRuleBypass2edl(i["bypass"], d, pre_append)
			tmp["bypass"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplication-AccessRule-Bypass")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationAccessRuleName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplication-AccessRule-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallCasbProfileSaasApplicationAccessRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationAccessRuleBypass2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallCasbProfileSaasApplicationAccessRuleName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControl2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControlName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplication-CustomControl-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := i["option"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControlOption2edl(i["option"], d, pre_append)
			tmp["option"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplication-CustomControl-Option")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOption2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplicationCustomControl-Option-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := i["user-input"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput2edl(i["user-input"], d, pre_append)
			tmp["user_input"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplicationCustomControl-Option-UserInput")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallCasbProfileSaasApplicationDomainControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationDomainControlDomains2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallCasbProfileSaasApplicationLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationSafeSearch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationSafeSearchControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallCasbProfileSaasApplicationTenantControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationTenantControlTenants2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFirewallCasbProfileSaasApplication(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("access_rule", flattenObjectFirewallCasbProfileSaasApplicationAccessRule2edl(o["access-rule"], d, "access_rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["access-rule"], "ObjectFirewallCasbProfileSaasApplication-AccessRule"); ok {
				if err = d.Set("access_rule", vv); err != nil {
					return fmt.Errorf("Error reading access_rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading access_rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("access_rule"); ok {
			if err = d.Set("access_rule", flattenObjectFirewallCasbProfileSaasApplicationAccessRule2edl(o["access-rule"], d, "access_rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["access-rule"], "ObjectFirewallCasbProfileSaasApplication-AccessRule"); ok {
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
		if err = d.Set("custom_control", flattenObjectFirewallCasbProfileSaasApplicationCustomControl2edl(o["custom-control"], d, "custom_control")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-control"], "ObjectFirewallCasbProfileSaasApplication-CustomControl"); ok {
				if err = d.Set("custom_control", vv); err != nil {
					return fmt.Errorf("Error reading custom_control: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_control: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_control"); ok {
			if err = d.Set("custom_control", flattenObjectFirewallCasbProfileSaasApplicationCustomControl2edl(o["custom-control"], d, "custom_control")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-control"], "ObjectFirewallCasbProfileSaasApplication-CustomControl"); ok {
					if err = d.Set("custom_control", vv); err != nil {
						return fmt.Errorf("Error reading custom_control: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_control: %v", err)
				}
			}
		}
	}

	if err = d.Set("domain_control", flattenObjectFirewallCasbProfileSaasApplicationDomainControl2edl(o["domain-control"], d, "domain_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-control"], "ObjectFirewallCasbProfileSaasApplication-DomainControl"); ok {
			if err = d.Set("domain_control", vv); err != nil {
				return fmt.Errorf("Error reading domain_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_control: %v", err)
		}
	}

	if err = d.Set("domain_control_domains", flattenObjectFirewallCasbProfileSaasApplicationDomainControlDomains2edl(o["domain-control-domains"], d, "domain_control_domains")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-control-domains"], "ObjectFirewallCasbProfileSaasApplication-DomainControlDomains"); ok {
			if err = d.Set("domain_control_domains", vv); err != nil {
				return fmt.Errorf("Error reading domain_control_domains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_control_domains: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectFirewallCasbProfileSaasApplicationLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectFirewallCasbProfileSaasApplication-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallCasbProfileSaasApplicationName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallCasbProfileSaasApplication-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("safe_search", flattenObjectFirewallCasbProfileSaasApplicationSafeSearch2edl(o["safe-search"], d, "safe_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["safe-search"], "ObjectFirewallCasbProfileSaasApplication-SafeSearch"); ok {
			if err = d.Set("safe_search", vv); err != nil {
				return fmt.Errorf("Error reading safe_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safe_search: %v", err)
		}
	}

	if err = d.Set("safe_search_control", flattenObjectFirewallCasbProfileSaasApplicationSafeSearchControl2edl(o["safe-search-control"], d, "safe_search_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["safe-search-control"], "ObjectFirewallCasbProfileSaasApplication-SafeSearchControl"); ok {
			if err = d.Set("safe_search_control", vv); err != nil {
				return fmt.Errorf("Error reading safe_search_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safe_search_control: %v", err)
		}
	}

	if err = d.Set("tenant_control", flattenObjectFirewallCasbProfileSaasApplicationTenantControl2edl(o["tenant-control"], d, "tenant_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant-control"], "ObjectFirewallCasbProfileSaasApplication-TenantControl"); ok {
			if err = d.Set("tenant_control", vv); err != nil {
				return fmt.Errorf("Error reading tenant_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant_control: %v", err)
		}
	}

	if err = d.Set("tenant_control_tenants", flattenObjectFirewallCasbProfileSaasApplicationTenantControlTenants2edl(o["tenant-control-tenants"], d, "tenant_control_tenants")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant-control-tenants"], "ObjectFirewallCasbProfileSaasApplication-TenantControlTenants"); ok {
			if err = d.Set("tenant_control_tenants", vv); err != nil {
				return fmt.Errorf("Error reading tenant_control_tenants: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant_control_tenants: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallCasbProfileSaasApplicationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallCasbProfileSaasApplicationAccessRule2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectFirewallCasbProfileSaasApplicationAccessRuleAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bypass"], _ = expandObjectFirewallCasbProfileSaasApplicationAccessRuleBypass2edl(d, i["bypass"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFirewallCasbProfileSaasApplicationAccessRuleName2edl(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallCasbProfileSaasApplicationAccessRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationAccessRuleBypass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallCasbProfileSaasApplicationAccessRuleName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallCasbProfileSaasApplicationCustomControlName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallCasbProfileSaasApplicationCustomControlOption2edl(d, i["option"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["option"] = t
			}
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-input"], _ = expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput2edl(d, i["user_input"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallCasbProfileSaasApplicationDomainControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationDomainControlDomains2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallCasbProfileSaasApplicationLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationSafeSearch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationSafeSearchControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallCasbProfileSaasApplicationTenantControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationTenantControlTenants2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFirewallCasbProfileSaasApplication(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_rule"); ok || d.HasChange("access_rule") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationAccessRule2edl(d, v, "access_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-rule"] = t
		}
	}

	if v, ok := d.GetOk("custom_control"); ok || d.HasChange("custom_control") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationCustomControl2edl(d, v, "custom_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-control"] = t
		}
	}

	if v, ok := d.GetOk("domain_control"); ok || d.HasChange("domain_control") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationDomainControl2edl(d, v, "domain_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-control"] = t
		}
	}

	if v, ok := d.GetOk("domain_control_domains"); ok || d.HasChange("domain_control_domains") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationDomainControlDomains2edl(d, v, "domain_control_domains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-control-domains"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("safe_search"); ok || d.HasChange("safe_search") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationSafeSearch2edl(d, v, "safe_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search"] = t
		}
	}

	if v, ok := d.GetOk("safe_search_control"); ok || d.HasChange("safe_search_control") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationSafeSearchControl2edl(d, v, "safe_search_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search-control"] = t
		}
	}

	if v, ok := d.GetOk("tenant_control"); ok || d.HasChange("tenant_control") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationTenantControl2edl(d, v, "tenant_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-control"] = t
		}
	}

	if v, ok := d.GetOk("tenant_control_tenants"); ok || d.HasChange("tenant_control_tenants") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationTenantControlTenants2edl(d, v, "tenant_control_tenants")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-control-tenants"] = t
		}
	}

	return &obj, nil
}
