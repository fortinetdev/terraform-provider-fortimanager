// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 IP pools.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallIppool6() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallIppool6Create,
		Read:   resourceObjectFirewallIppool6Read,
		Update: resourceObjectFirewallIppool6Update,
		Delete: resourceObjectFirewallIppool6Delete,

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
			"add_nat46_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_scope": &schema.Schema{
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
						"add_nat46_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"endip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"internal_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nat46": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"startip": &schema.Schema{
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
			"endip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"startip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
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

func resourceObjectFirewallIppool6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallIppool6(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallIppool6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallIppool6(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallIppool6 resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallIppool6Read(d, m)
}

func resourceObjectFirewallIppool6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallIppool6(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallIppool6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallIppool6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallIppool6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallIppool6Read(d, m)
}

func resourceObjectFirewallIppool6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallIppool6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallIppool6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallIppool6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallIppool6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallIppool6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallIppool6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallIppool6 resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallIppool6AddNat46Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6DynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectFirewallIppool6DynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat46_route"
		if _, ok := i["add-nat46-route"]; ok {
			v := flattenObjectFirewallIppool6DynamicMappingAddNat46Route(i["add-nat46-route"], d, pre_append)
			tmp["add_nat46_route"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6-DynamicMapping-AddNat46Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			v := flattenObjectFirewallIppool6DynamicMappingComments(i["comments"], d, pre_append)
			tmp["comments"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6-DynamicMapping-Comments")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endip"
		if _, ok := i["endip"]; ok {
			v := flattenObjectFirewallIppool6DynamicMappingEndip(i["endip"], d, pre_append)
			tmp["endip"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6-DynamicMapping-Endip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_prefix"
		if _, ok := i["external-prefix"]; ok {
			v := flattenObjectFirewallIppool6DynamicMappingExternalPrefix(i["external-prefix"], d, pre_append)
			tmp["external_prefix"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6-DynamicMapping-ExternalPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internal_prefix"
		if _, ok := i["internal-prefix"]; ok {
			v := flattenObjectFirewallIppool6DynamicMappingInternalPrefix(i["internal-prefix"], d, pre_append)
			tmp["internal_prefix"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6-DynamicMapping-InternalPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat46"
		if _, ok := i["nat46"]; ok {
			v := flattenObjectFirewallIppool6DynamicMappingNat46(i["nat46"], d, pre_append)
			tmp["nat46"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6-DynamicMapping-Nat46")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startip"
		if _, ok := i["startip"]; ok {
			v := flattenObjectFirewallIppool6DynamicMappingStartip(i["startip"], d, pre_append)
			tmp["startip"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6-DynamicMapping-Startip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallIppool6DynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6-DynamicMapping-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallIppool6DynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallIppool6DynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6DynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFirewallIppool6DynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool6DynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallIppool6DynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6DynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6DynamicMappingAddNat46Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6DynamicMappingComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6DynamicMappingEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6DynamicMappingExternalPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6DynamicMappingInternalPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6DynamicMappingNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6DynamicMappingStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6DynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6Endip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6ExternalPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6InternalPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6Nat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6Startip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppool6Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallIppool6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("add_nat46_route", flattenObjectFirewallIppool6AddNat46Route(o["add-nat46-route"], d, "add_nat46_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-nat46-route"], "ObjectFirewallIppool6-AddNat46Route"); ok {
			if err = d.Set("add_nat46_route", vv); err != nil {
				return fmt.Errorf("Error reading add_nat46_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_nat46_route: %v", err)
		}
	}

	if err = d.Set("comments", flattenObjectFirewallIppool6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectFirewallIppool6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectFirewallIppool6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallIppool6-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectFirewallIppool6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallIppool6-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("endip", flattenObjectFirewallIppool6Endip(o["endip"], d, "endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["endip"], "ObjectFirewallIppool6-Endip"); ok {
			if err = d.Set("endip", vv); err != nil {
				return fmt.Errorf("Error reading endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endip: %v", err)
		}
	}

	if err = d.Set("external_prefix", flattenObjectFirewallIppool6ExternalPrefix(o["external-prefix"], d, "external_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-prefix"], "ObjectFirewallIppool6-ExternalPrefix"); ok {
			if err = d.Set("external_prefix", vv); err != nil {
				return fmt.Errorf("Error reading external_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_prefix: %v", err)
		}
	}

	if err = d.Set("internal_prefix", flattenObjectFirewallIppool6InternalPrefix(o["internal-prefix"], d, "internal_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-prefix"], "ObjectFirewallIppool6-InternalPrefix"); ok {
			if err = d.Set("internal_prefix", vv); err != nil {
				return fmt.Errorf("Error reading internal_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_prefix: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallIppool6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallIppool6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat46", flattenObjectFirewallIppool6Nat46(o["nat46"], d, "nat46")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46"], "ObjectFirewallIppool6-Nat46"); ok {
			if err = d.Set("nat46", vv); err != nil {
				return fmt.Errorf("Error reading nat46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("startip", flattenObjectFirewallIppool6Startip(o["startip"], d, "startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["startip"], "ObjectFirewallIppool6-Startip"); ok {
			if err = d.Set("startip", vv); err != nil {
				return fmt.Errorf("Error reading startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading startip: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallIppool6Type(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallIppool6-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallIppool6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallIppool6AddNat46Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6DynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallIppool6DynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat46_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-nat46-route"], _ = expandObjectFirewallIppool6DynamicMappingAddNat46Route(d, i["add_nat46_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comments"], _ = expandObjectFirewallIppool6DynamicMappingComments(d, i["comments"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["endip"], _ = expandObjectFirewallIppool6DynamicMappingEndip(d, i["endip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-prefix"], _ = expandObjectFirewallIppool6DynamicMappingExternalPrefix(d, i["external_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internal_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internal-prefix"], _ = expandObjectFirewallIppool6DynamicMappingInternalPrefix(d, i["internal_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat46"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat46"], _ = expandObjectFirewallIppool6DynamicMappingNat46(d, i["nat46"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["startip"], _ = expandObjectFirewallIppool6DynamicMappingStartip(d, i["startip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallIppool6DynamicMappingType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallIppool6DynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallIppool6DynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectFirewallIppool6DynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallIppool6DynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6DynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6DynamicMappingAddNat46Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6DynamicMappingComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6DynamicMappingEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6DynamicMappingExternalPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6DynamicMappingInternalPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6DynamicMappingNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6DynamicMappingStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6DynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6Endip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6ExternalPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6InternalPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6Nat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6Startip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppool6Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallIppool6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_nat46_route"); ok || d.HasChange("add_nat46_route") {
		t, err := expandObjectFirewallIppool6AddNat46Route(d, v, "add_nat46_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat46-route"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandObjectFirewallIppool6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectFirewallIppool6DynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("endip"); ok || d.HasChange("endip") {
		t, err := expandObjectFirewallIppool6Endip(d, v, "endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("external_prefix"); ok || d.HasChange("external_prefix") {
		t, err := expandObjectFirewallIppool6ExternalPrefix(d, v, "external_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-prefix"] = t
		}
	}

	if v, ok := d.GetOk("internal_prefix"); ok || d.HasChange("internal_prefix") {
		t, err := expandObjectFirewallIppool6InternalPrefix(d, v, "internal_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-prefix"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallIppool6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok || d.HasChange("nat46") {
		t, err := expandObjectFirewallIppool6Nat46(d, v, "nat46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok || d.HasChange("startip") {
		t, err := expandObjectFirewallIppool6Startip(d, v, "startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallIppool6Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
