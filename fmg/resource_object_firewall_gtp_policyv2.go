// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Apply allow or deny action to each GTPv2-c packet.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallGtpPolicyV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpPolicyV2Create,
		Read:   resourceObjectFirewallGtpPolicyV2Read,
		Update: resourceObjectFirewallGtpPolicyV2Update,
		Delete: resourceObjectFirewallGtpPolicyV2Delete,

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
			"gtp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apn_sel_mode": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"apnmember": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"imsi_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_apn_restriction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mei": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"messages": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"msisdn_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rat_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"uli": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallGtpPolicyV2Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	obj, err := getObjectObjectFirewallGtpPolicyV2(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpPolicyV2 resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallGtpPolicyV2(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpPolicyV2 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpPolicyV2Read(d, m)
}

func resourceObjectFirewallGtpPolicyV2Update(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	obj, err := getObjectObjectFirewallGtpPolicyV2(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpPolicyV2 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtpPolicyV2(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpPolicyV2 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpPolicyV2Read(d, m)
}

func resourceObjectFirewallGtpPolicyV2Delete(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	err = c.DeleteObjectFirewallGtpPolicyV2(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtpPolicyV2 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpPolicyV2Read(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	if gtp == "" {
		gtp = importOptionChecking(m.(*FortiClient).Cfg, "gtp")
		if gtp == "" {
			return fmt.Errorf("Parameter gtp is missing")
		}
		if err = d.Set("gtp", gtp); err != nil {
			return fmt.Errorf("Error set params gtp: %v", err)
		}
	}
	paradict["gtp"] = gtp

	o, err := c.ReadObjectFirewallGtpPolicyV2(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpPolicyV2 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtpPolicyV2(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpPolicyV2 resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpPolicyV2Action2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2ApnSelMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallGtpPolicyV2Apnmember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2ImsiPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2MaxApnRestriction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2Mei2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2Messages2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallGtpPolicyV2MsisdnPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2RatType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallGtpPolicyV2Uli2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFirewallGtpPolicyV2(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectFirewallGtpPolicyV2Action2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectFirewallGtpPolicyV2-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("apn_sel_mode", flattenObjectFirewallGtpPolicyV2ApnSelMode2edl(o["apn-sel-mode"], d, "apn_sel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["apn-sel-mode"], "ObjectFirewallGtpPolicyV2-ApnSelMode"); ok {
			if err = d.Set("apn_sel_mode", vv); err != nil {
				return fmt.Errorf("Error reading apn_sel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apn_sel_mode: %v", err)
		}
	}

	if err = d.Set("apnmember", flattenObjectFirewallGtpPolicyV2Apnmember2edl(o["apnmember"], d, "apnmember")); err != nil {
		if vv, ok := fortiAPIPatch(o["apnmember"], "ObjectFirewallGtpPolicyV2-Apnmember"); ok {
			if err = d.Set("apnmember", vv); err != nil {
				return fmt.Errorf("Error reading apnmember: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apnmember: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallGtpPolicyV2Id2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallGtpPolicyV2-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("imsi_prefix", flattenObjectFirewallGtpPolicyV2ImsiPrefix2edl(o["imsi-prefix"], d, "imsi_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["imsi-prefix"], "ObjectFirewallGtpPolicyV2-ImsiPrefix"); ok {
			if err = d.Set("imsi_prefix", vv); err != nil {
				return fmt.Errorf("Error reading imsi_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imsi_prefix: %v", err)
		}
	}

	if err = d.Set("max_apn_restriction", flattenObjectFirewallGtpPolicyV2MaxApnRestriction2edl(o["max-apn-restriction"], d, "max_apn_restriction")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-apn-restriction"], "ObjectFirewallGtpPolicyV2-MaxApnRestriction"); ok {
			if err = d.Set("max_apn_restriction", vv); err != nil {
				return fmt.Errorf("Error reading max_apn_restriction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_apn_restriction: %v", err)
		}
	}

	if err = d.Set("mei", flattenObjectFirewallGtpPolicyV2Mei2edl(o["mei"], d, "mei")); err != nil {
		if vv, ok := fortiAPIPatch(o["mei"], "ObjectFirewallGtpPolicyV2-Mei"); ok {
			if err = d.Set("mei", vv); err != nil {
				return fmt.Errorf("Error reading mei: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mei: %v", err)
		}
	}

	if err = d.Set("messages", flattenObjectFirewallGtpPolicyV2Messages2edl(o["messages"], d, "messages")); err != nil {
		if vv, ok := fortiAPIPatch(o["messages"], "ObjectFirewallGtpPolicyV2-Messages"); ok {
			if err = d.Set("messages", vv); err != nil {
				return fmt.Errorf("Error reading messages: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading messages: %v", err)
		}
	}

	if err = d.Set("msisdn_prefix", flattenObjectFirewallGtpPolicyV2MsisdnPrefix2edl(o["msisdn-prefix"], d, "msisdn_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["msisdn-prefix"], "ObjectFirewallGtpPolicyV2-MsisdnPrefix"); ok {
			if err = d.Set("msisdn_prefix", vv); err != nil {
				return fmt.Errorf("Error reading msisdn_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading msisdn_prefix: %v", err)
		}
	}

	if err = d.Set("rat_type", flattenObjectFirewallGtpPolicyV2RatType2edl(o["rat-type"], d, "rat_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["rat-type"], "ObjectFirewallGtpPolicyV2-RatType"); ok {
			if err = d.Set("rat_type", vv); err != nil {
				return fmt.Errorf("Error reading rat_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rat_type: %v", err)
		}
	}

	if err = d.Set("uli", flattenObjectFirewallGtpPolicyV2Uli2edl(o["uli"], d, "uli")); err != nil {
		if vv, ok := fortiAPIPatch(o["uli"], "ObjectFirewallGtpPolicyV2-Uli"); ok {
			if err = d.Set("uli", vv); err != nil {
				return fmt.Errorf("Error reading uli: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uli: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpPolicyV2FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpPolicyV2Action2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2ApnSelMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyV2Apnmember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2ImsiPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2MaxApnRestriction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2Mei2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2Messages2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyV2MsisdnPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2RatType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyV2Uli2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFirewallGtpPolicyV2(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectFirewallGtpPolicyV2Action2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("apn_sel_mode"); ok || d.HasChange("apn_sel_mode") {
		t, err := expandObjectFirewallGtpPolicyV2ApnSelMode2edl(d, v, "apn_sel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn-sel-mode"] = t
		}
	}

	if v, ok := d.GetOk("apnmember"); ok || d.HasChange("apnmember") {
		t, err := expandObjectFirewallGtpPolicyV2Apnmember2edl(d, v, "apnmember")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apnmember"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallGtpPolicyV2Id2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("imsi_prefix"); ok || d.HasChange("imsi_prefix") {
		t, err := expandObjectFirewallGtpPolicyV2ImsiPrefix2edl(d, v, "imsi_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imsi-prefix"] = t
		}
	}

	if v, ok := d.GetOk("max_apn_restriction"); ok || d.HasChange("max_apn_restriction") {
		t, err := expandObjectFirewallGtpPolicyV2MaxApnRestriction2edl(d, v, "max_apn_restriction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-apn-restriction"] = t
		}
	}

	if v, ok := d.GetOk("mei"); ok || d.HasChange("mei") {
		t, err := expandObjectFirewallGtpPolicyV2Mei2edl(d, v, "mei")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mei"] = t
		}
	}

	if v, ok := d.GetOk("messages"); ok || d.HasChange("messages") {
		t, err := expandObjectFirewallGtpPolicyV2Messages2edl(d, v, "messages")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["messages"] = t
		}
	}

	if v, ok := d.GetOk("msisdn_prefix"); ok || d.HasChange("msisdn_prefix") {
		t, err := expandObjectFirewallGtpPolicyV2MsisdnPrefix2edl(d, v, "msisdn_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msisdn-prefix"] = t
		}
	}

	if v, ok := d.GetOk("rat_type"); ok || d.HasChange("rat_type") {
		t, err := expandObjectFirewallGtpPolicyV2RatType2edl(d, v, "rat_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rat-type"] = t
		}
	}

	if v, ok := d.GetOk("uli"); ok || d.HasChange("uli") {
		t, err := expandObjectFirewallGtpPolicyV2Uli2edl(d, v, "uli")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uli"] = t
		}
	}

	return &obj, nil
}
