// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WiFi bridge access control list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerAccessControlList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerAccessControlListCreate,
		Read:   resourceObjectWirelessControllerAccessControlListRead,
		Update: resourceObjectWirelessControllerAccessControlListUpdate,
		Delete: resourceObjectWirelessControllerAccessControlListDelete,

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
			"layer3_ipv4_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rule_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"srcport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"layer3_ipv6_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rule_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"srcport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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

func resourceObjectWirelessControllerAccessControlListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerAccessControlList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerAccessControlList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerAccessControlList(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerAccessControlList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerAccessControlListRead(d, m)
}

func resourceObjectWirelessControllerAccessControlListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerAccessControlList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerAccessControlList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerAccessControlList(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerAccessControlList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerAccessControlListRead(d, m)
}

func resourceObjectWirelessControllerAccessControlListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerAccessControlList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerAccessControlList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerAccessControlListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerAccessControlList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerAccessControlList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerAccessControlList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerAccessControlList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerAccessControlListComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4Rules(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv4Rules-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv4Rules-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv4Rules-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
		if _, ok := i["dstport"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstport(i["dstport"], d, pre_append)
			tmp["dstport"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv4Rules-Dstport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv4Rules-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := i["rule-id"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(i["rule-id"], d, pre_append)
			tmp["rule_id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv4Rules-RuleId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv4Rules-Srcaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
		if _, ok := i["srcport"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(i["srcport"], d, pre_append)
			tmp["srcport"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv4Rules-Srcport")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv6Rules(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv6Rules-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv6Rules-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv6Rules-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
		if _, ok := i["dstport"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesDstport(i["dstport"], d, pre_append)
			tmp["dstport"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv6Rules-Dstport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv6Rules-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := i["rule-id"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(i["rule-id"], d, pre_append)
			tmp["rule_id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv6Rules-RuleId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv6Rules-Srcaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
		if _, ok := i["srcport"]; ok {
			v := flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(i["srcport"], d, pre_append)
			tmp["srcport"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerAccessControlList-Layer3Ipv6Rules-Srcport")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesDstport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerAccessControlList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenObjectWirelessControllerAccessControlListComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerAccessControlList-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("layer3_ipv4_rules", flattenObjectWirelessControllerAccessControlListLayer3Ipv4Rules(o["layer3-ipv4-rules"], d, "layer3_ipv4_rules")); err != nil {
			if vv, ok := fortiAPIPatch(o["layer3-ipv4-rules"], "ObjectWirelessControllerAccessControlList-Layer3Ipv4Rules"); ok {
				if err = d.Set("layer3_ipv4_rules", vv); err != nil {
					return fmt.Errorf("Error reading layer3_ipv4_rules: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading layer3_ipv4_rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("layer3_ipv4_rules"); ok {
			if err = d.Set("layer3_ipv4_rules", flattenObjectWirelessControllerAccessControlListLayer3Ipv4Rules(o["layer3-ipv4-rules"], d, "layer3_ipv4_rules")); err != nil {
				if vv, ok := fortiAPIPatch(o["layer3-ipv4-rules"], "ObjectWirelessControllerAccessControlList-Layer3Ipv4Rules"); ok {
					if err = d.Set("layer3_ipv4_rules", vv); err != nil {
						return fmt.Errorf("Error reading layer3_ipv4_rules: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading layer3_ipv4_rules: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("layer3_ipv6_rules", flattenObjectWirelessControllerAccessControlListLayer3Ipv6Rules(o["layer3-ipv6-rules"], d, "layer3_ipv6_rules")); err != nil {
			if vv, ok := fortiAPIPatch(o["layer3-ipv6-rules"], "ObjectWirelessControllerAccessControlList-Layer3Ipv6Rules"); ok {
				if err = d.Set("layer3_ipv6_rules", vv); err != nil {
					return fmt.Errorf("Error reading layer3_ipv6_rules: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading layer3_ipv6_rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("layer3_ipv6_rules"); ok {
			if err = d.Set("layer3_ipv6_rules", flattenObjectWirelessControllerAccessControlListLayer3Ipv6Rules(o["layer3-ipv6-rules"], d, "layer3_ipv6_rules")); err != nil {
				if vv, ok := fortiAPIPatch(o["layer3-ipv6-rules"], "ObjectWirelessControllerAccessControlList-Layer3Ipv6Rules"); ok {
					if err = d.Set("layer3_ipv6_rules", vv); err != nil {
						return fmt.Errorf("Error reading layer3_ipv6_rules: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading layer3_ipv6_rules: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerAccessControlListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerAccessControlList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerAccessControlListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerAccessControlListComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4Rules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstport"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstport(d, i["dstport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rule-id"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(d, i["rule_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(d, i["srcaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcport"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(d, i["srcport"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6Rules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstport"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesDstport(d, i["dstport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rule-id"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(d, i["rule_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(d, i["srcaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcport"], _ = expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(d, i["srcport"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesDstport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerAccessControlList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWirelessControllerAccessControlListComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("layer3_ipv4_rules"); ok || d.HasChange("layer3_ipv4_rules") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv4Rules(d, v, "layer3_ipv4_rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["layer3-ipv4-rules"] = t
		}
	}

	if v, ok := d.GetOk("layer3_ipv6_rules"); ok || d.HasChange("layer3_ipv6_rules") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv6Rules(d, v, "layer3_ipv6_rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["layer3-ipv6-rules"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerAccessControlListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
