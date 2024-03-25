// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallGtpPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpPolicyCreate,
		Read:   resourceObjectFirewallGtpPolicyRead,
		Update: resourceObjectFirewallGtpPolicyUpdate,
		Delete: resourceObjectFirewallGtpPolicyDelete,

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
			"imei": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"imsi": &schema.Schema{
				Type:     schema.TypeString,
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
			"messages": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"msisdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"msisdn_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rai": &schema.Schema{
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
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallGtpPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallGtpPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallGtpPolicy(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpPolicy resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpPolicyRead(d, m)
}

func resourceObjectFirewallGtpPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallGtpPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtpPolicy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpPolicyRead(d, m)
}

func resourceObjectFirewallGtpPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallGtpPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtpPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallGtpPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtpPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpPolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpPolicyAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyApnSelMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallGtpPolicyApnmember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallGtpPolicyId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyImei2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyImsi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyImsiPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyMaxApnRestriction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyMessages2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallGtpPolicyMsisdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyMsisdnPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyRai2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyRatType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallGtpPolicyUli2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallGtpPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectFirewallGtpPolicyAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectFirewallGtpPolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("apn_sel_mode", flattenObjectFirewallGtpPolicyApnSelMode2edl(o["apn-sel-mode"], d, "apn_sel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["apn-sel-mode"], "ObjectFirewallGtpPolicy-ApnSelMode"); ok {
			if err = d.Set("apn_sel_mode", vv); err != nil {
				return fmt.Errorf("Error reading apn_sel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apn_sel_mode: %v", err)
		}
	}

	if err = d.Set("apnmember", flattenObjectFirewallGtpPolicyApnmember2edl(o["apnmember"], d, "apnmember")); err != nil {
		if vv, ok := fortiAPIPatch(o["apnmember"], "ObjectFirewallGtpPolicy-Apnmember"); ok {
			if err = d.Set("apnmember", vv); err != nil {
				return fmt.Errorf("Error reading apnmember: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apnmember: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallGtpPolicyId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallGtpPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("imei", flattenObjectFirewallGtpPolicyImei2edl(o["imei"], d, "imei")); err != nil {
		if vv, ok := fortiAPIPatch(o["imei"], "ObjectFirewallGtpPolicy-Imei"); ok {
			if err = d.Set("imei", vv); err != nil {
				return fmt.Errorf("Error reading imei: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imei: %v", err)
		}
	}

	if err = d.Set("imsi", flattenObjectFirewallGtpPolicyImsi2edl(o["imsi"], d, "imsi")); err != nil {
		if vv, ok := fortiAPIPatch(o["imsi"], "ObjectFirewallGtpPolicy-Imsi"); ok {
			if err = d.Set("imsi", vv); err != nil {
				return fmt.Errorf("Error reading imsi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imsi: %v", err)
		}
	}

	if err = d.Set("imsi_prefix", flattenObjectFirewallGtpPolicyImsiPrefix2edl(o["imsi-prefix"], d, "imsi_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["imsi-prefix"], "ObjectFirewallGtpPolicy-ImsiPrefix"); ok {
			if err = d.Set("imsi_prefix", vv); err != nil {
				return fmt.Errorf("Error reading imsi_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imsi_prefix: %v", err)
		}
	}

	if err = d.Set("max_apn_restriction", flattenObjectFirewallGtpPolicyMaxApnRestriction2edl(o["max-apn-restriction"], d, "max_apn_restriction")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-apn-restriction"], "ObjectFirewallGtpPolicy-MaxApnRestriction"); ok {
			if err = d.Set("max_apn_restriction", vv); err != nil {
				return fmt.Errorf("Error reading max_apn_restriction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_apn_restriction: %v", err)
		}
	}

	if err = d.Set("messages", flattenObjectFirewallGtpPolicyMessages2edl(o["messages"], d, "messages")); err != nil {
		if vv, ok := fortiAPIPatch(o["messages"], "ObjectFirewallGtpPolicy-Messages"); ok {
			if err = d.Set("messages", vv); err != nil {
				return fmt.Errorf("Error reading messages: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading messages: %v", err)
		}
	}

	if err = d.Set("msisdn", flattenObjectFirewallGtpPolicyMsisdn2edl(o["msisdn"], d, "msisdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["msisdn"], "ObjectFirewallGtpPolicy-Msisdn"); ok {
			if err = d.Set("msisdn", vv); err != nil {
				return fmt.Errorf("Error reading msisdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading msisdn: %v", err)
		}
	}

	if err = d.Set("msisdn_prefix", flattenObjectFirewallGtpPolicyMsisdnPrefix2edl(o["msisdn-prefix"], d, "msisdn_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["msisdn-prefix"], "ObjectFirewallGtpPolicy-MsisdnPrefix"); ok {
			if err = d.Set("msisdn_prefix", vv); err != nil {
				return fmt.Errorf("Error reading msisdn_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading msisdn_prefix: %v", err)
		}
	}

	if err = d.Set("rai", flattenObjectFirewallGtpPolicyRai2edl(o["rai"], d, "rai")); err != nil {
		if vv, ok := fortiAPIPatch(o["rai"], "ObjectFirewallGtpPolicy-Rai"); ok {
			if err = d.Set("rai", vv); err != nil {
				return fmt.Errorf("Error reading rai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rai: %v", err)
		}
	}

	if err = d.Set("rat_type", flattenObjectFirewallGtpPolicyRatType2edl(o["rat-type"], d, "rat_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["rat-type"], "ObjectFirewallGtpPolicy-RatType"); ok {
			if err = d.Set("rat_type", vv); err != nil {
				return fmt.Errorf("Error reading rat_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rat_type: %v", err)
		}
	}

	if err = d.Set("uli", flattenObjectFirewallGtpPolicyUli2edl(o["uli"], d, "uli")); err != nil {
		if vv, ok := fortiAPIPatch(o["uli"], "ObjectFirewallGtpPolicy-Uli"); ok {
			if err = d.Set("uli", vv); err != nil {
				return fmt.Errorf("Error reading uli: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uli: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpPolicyAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyApnSelMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyApnmember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallGtpPolicyId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyImei2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyImsi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyImsiPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyMaxApnRestriction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyMessages2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyMsisdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyMsisdnPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyRai2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyRatType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyUli2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallGtpPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectFirewallGtpPolicyAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("apn_sel_mode"); ok || d.HasChange("apn_sel_mode") {
		t, err := expandObjectFirewallGtpPolicyApnSelMode2edl(d, v, "apn_sel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn-sel-mode"] = t
		}
	}

	if v, ok := d.GetOk("apnmember"); ok || d.HasChange("apnmember") {
		t, err := expandObjectFirewallGtpPolicyApnmember2edl(d, v, "apnmember")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apnmember"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallGtpPolicyId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("imei"); ok || d.HasChange("imei") {
		t, err := expandObjectFirewallGtpPolicyImei2edl(d, v, "imei")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imei"] = t
		}
	}

	if v, ok := d.GetOk("imsi"); ok || d.HasChange("imsi") {
		t, err := expandObjectFirewallGtpPolicyImsi2edl(d, v, "imsi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imsi"] = t
		}
	}

	if v, ok := d.GetOk("imsi_prefix"); ok || d.HasChange("imsi_prefix") {
		t, err := expandObjectFirewallGtpPolicyImsiPrefix2edl(d, v, "imsi_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imsi-prefix"] = t
		}
	}

	if v, ok := d.GetOk("max_apn_restriction"); ok || d.HasChange("max_apn_restriction") {
		t, err := expandObjectFirewallGtpPolicyMaxApnRestriction2edl(d, v, "max_apn_restriction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-apn-restriction"] = t
		}
	}

	if v, ok := d.GetOk("messages"); ok || d.HasChange("messages") {
		t, err := expandObjectFirewallGtpPolicyMessages2edl(d, v, "messages")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["messages"] = t
		}
	}

	if v, ok := d.GetOk("msisdn"); ok || d.HasChange("msisdn") {
		t, err := expandObjectFirewallGtpPolicyMsisdn2edl(d, v, "msisdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msisdn"] = t
		}
	}

	if v, ok := d.GetOk("msisdn_prefix"); ok || d.HasChange("msisdn_prefix") {
		t, err := expandObjectFirewallGtpPolicyMsisdnPrefix2edl(d, v, "msisdn_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msisdn-prefix"] = t
		}
	}

	if v, ok := d.GetOk("rai"); ok || d.HasChange("rai") {
		t, err := expandObjectFirewallGtpPolicyRai2edl(d, v, "rai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rai"] = t
		}
	}

	if v, ok := d.GetOk("rat_type"); ok || d.HasChange("rat_type") {
		t, err := expandObjectFirewallGtpPolicyRatType2edl(d, v, "rat_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rat-type"] = t
		}
	}

	if v, ok := d.GetOk("uli"); ok || d.HasChange("uli") {
		t, err := expandObjectFirewallGtpPolicyUli2edl(d, v, "uli")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uli"] = t
		}
	}

	return &obj, nil
}
