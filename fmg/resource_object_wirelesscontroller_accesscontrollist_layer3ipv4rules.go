// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: AP ACL layer3 ipv4 rule list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerAccessControlListLayer3Ipv4Rules() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerAccessControlListLayer3Ipv4RulesCreate,
		Read:   resourceObjectWirelessControllerAccessControlListLayer3Ipv4RulesRead,
		Update: resourceObjectWirelessControllerAccessControlListLayer3Ipv4RulesUpdate,
		Delete: resourceObjectWirelessControllerAccessControlListLayer3Ipv4RulesDelete,

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
			"access_control_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
				Computed: true,
			},
			"dstport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rule_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerAccessControlListLayer3Ipv4RulesCreate(d *schema.ResourceData, m interface{}) error {
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

	access_control_list := d.Get("access_control_list").(string)
	paradict["access_control_list"] = access_control_list

	obj, err := getObjectObjectWirelessControllerAccessControlListLayer3Ipv4Rules(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerAccessControlListLayer3Ipv4Rules resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerAccessControlListLayer3Ipv4Rules(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerAccessControlListLayer3Ipv4Rules resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "rule_id")))

	return resourceObjectWirelessControllerAccessControlListLayer3Ipv4RulesRead(d, m)
}

func resourceObjectWirelessControllerAccessControlListLayer3Ipv4RulesUpdate(d *schema.ResourceData, m interface{}) error {
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

	access_control_list := d.Get("access_control_list").(string)
	paradict["access_control_list"] = access_control_list

	obj, err := getObjectObjectWirelessControllerAccessControlListLayer3Ipv4Rules(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerAccessControlListLayer3Ipv4Rules resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerAccessControlListLayer3Ipv4Rules(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerAccessControlListLayer3Ipv4Rules resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "rule_id")))

	return resourceObjectWirelessControllerAccessControlListLayer3Ipv4RulesRead(d, m)
}

func resourceObjectWirelessControllerAccessControlListLayer3Ipv4RulesDelete(d *schema.ResourceData, m interface{}) error {
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

	access_control_list := d.Get("access_control_list").(string)
	paradict["access_control_list"] = access_control_list

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerAccessControlListLayer3Ipv4Rules(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerAccessControlListLayer3Ipv4Rules resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerAccessControlListLayer3Ipv4RulesRead(d *schema.ResourceData, m interface{}) error {
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

	access_control_list := d.Get("access_control_list").(string)
	if access_control_list == "" {
		access_control_list = importOptionChecking(m.(*FortiClient).Cfg, "access_control_list")
		if access_control_list == "" {
			return fmt.Errorf("Parameter access_control_list is missing")
		}
		if err = d.Set("access_control_list", access_control_list); err != nil {
			return fmt.Errorf("Error set params access_control_list: %v", err)
		}
	}
	paradict["access_control_list"] = access_control_list

	o, err := c.ReadObjectWirelessControllerAccessControlListLayer3Ipv4Rules(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerAccessControlListLayer3Ipv4Rules resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerAccessControlListLayer3Ipv4Rules(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerAccessControlListLayer3Ipv4Rules resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerAccessControlListLayer3Ipv4Rules(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectWirelessControllerAccessControlListLayer3Ipv4Rules-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerAccessControlListLayer3Ipv4Rules-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr2edl(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "ObjectWirelessControllerAccessControlListLayer3Ipv4Rules-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstport", flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstport2edl(o["dstport"], d, "dstport")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstport"], "ObjectWirelessControllerAccessControlListLayer3Ipv4Rules-Dstport"); ok {
			if err = d.Set("dstport", vv); err != nil {
				return fmt.Errorf("Error reading dstport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstport: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectWirelessControllerAccessControlListLayer3Ipv4Rules-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("rule_id", flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesRuleId2edl(o["rule-id"], d, "rule_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["rule-id"], "ObjectWirelessControllerAccessControlListLayer3Ipv4Rules-RuleId"); ok {
			if err = d.Set("rule_id", vv); err != nil {
				return fmt.Errorf("Error reading rule_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rule_id: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr2edl(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "ObjectWirelessControllerAccessControlListLayer3Ipv4Rules-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcport", flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcport2edl(o["srcport"], d, "srcport")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcport"], "ObjectWirelessControllerAccessControlListLayer3Ipv4Rules-Srcport"); ok {
			if err = d.Set("srcport", vv); err != nil {
				return fmt.Errorf("Error reading srcport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcport: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv4RulesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerAccessControlListLayer3Ipv4Rules(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr2edl(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstport"); ok || d.HasChange("dstport") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesDstport2edl(d, v, "dstport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstport"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("rule_id"); ok || d.HasChange("rule_id") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesRuleId2edl(d, v, "rule_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule-id"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr2edl(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcport"); ok || d.HasChange("srcport") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv4RulesSrcport2edl(d, v, "srcport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcport"] = t
		}
	}

	return &obj, nil
}
