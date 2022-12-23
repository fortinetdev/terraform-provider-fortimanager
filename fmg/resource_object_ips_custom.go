// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS custom signature.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectIpsCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIpsCustomCreate,
		Read:   resourceObjectIpsCustomRead,
		Update: resourceObjectIpsCustomUpdate,
		Delete: resourceObjectIpsCustomDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": &schema.Schema{
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
			"log_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sig_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectIpsCustomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectIpsCustom(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIpsCustom resource while getting object: %v", err)
	}

	_, err = c.CreateObjectIpsCustom(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectIpsCustom resource: %v", err)
	}

	d.SetId(getStringKey(d, "tag"))

	return resourceObjectIpsCustomRead(d, m)
}

func resourceObjectIpsCustomUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectIpsCustom(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIpsCustom resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectIpsCustom(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIpsCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "tag"))

	return resourceObjectIpsCustomRead(d, m)
}

func resourceObjectIpsCustomDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectIpsCustom(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIpsCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIpsCustomRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectIpsCustom(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIpsCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIpsCustom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIpsCustom resource from API: %v", err)
	}
	return nil
}

func flattenObjectIpsCustomAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsCustomApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsCustomComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsCustomLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsCustomLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsCustomLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsCustomOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsCustomProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsCustomRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsCustomSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsCustomSigName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsCustomSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsCustomStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsCustomTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectIpsCustom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectIpsCustomAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectIpsCustom-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("application", flattenObjectIpsCustomApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "ObjectIpsCustom-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectIpsCustomComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectIpsCustom-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("location", flattenObjectIpsCustomLocation(o["location"], d, "location")); err != nil {
		if vv, ok := fortiAPIPatch(o["location"], "ObjectIpsCustom-Location"); ok {
			if err = d.Set("location", vv); err != nil {
				return fmt.Errorf("Error reading location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectIpsCustomLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectIpsCustom-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenObjectIpsCustomLogPacket(o["log-packet"], d, "log_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-packet"], "ObjectIpsCustom-LogPacket"); ok {
			if err = d.Set("log_packet", vv); err != nil {
				return fmt.Errorf("Error reading log_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("os", flattenObjectIpsCustomOs(o["os"], d, "os")); err != nil {
		if vv, ok := fortiAPIPatch(o["os"], "ObjectIpsCustom-Os"); ok {
			if err = d.Set("os", vv); err != nil {
				return fmt.Errorf("Error reading os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectIpsCustomProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectIpsCustom-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("rule_id", flattenObjectIpsCustomRuleId(o["rule-id"], d, "rule_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["rule-id"], "ObjectIpsCustom-RuleId"); ok {
			if err = d.Set("rule_id", vv); err != nil {
				return fmt.Errorf("Error reading rule_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rule_id: %v", err)
		}
	}

	if err = d.Set("severity", flattenObjectIpsCustomSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "ObjectIpsCustom-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("sig_name", flattenObjectIpsCustomSigName(o["sig-name"], d, "sig_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["sig-name"], "ObjectIpsCustom-SigName"); ok {
			if err = d.Set("sig_name", vv); err != nil {
				return fmt.Errorf("Error reading sig_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sig_name: %v", err)
		}
	}

	if err = d.Set("signature", flattenObjectIpsCustomSignature(o["signature"], d, "signature")); err != nil {
		if vv, ok := fortiAPIPatch(o["signature"], "ObjectIpsCustom-Signature"); ok {
			if err = d.Set("signature", vv); err != nil {
				return fmt.Errorf("Error reading signature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signature: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectIpsCustomStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectIpsCustom-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tag", flattenObjectIpsCustomTag(o["tag"], d, "tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag"], "ObjectIpsCustom-Tag"); ok {
			if err = d.Set("tag", vv); err != nil {
				return fmt.Errorf("Error reading tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	return nil
}

func flattenObjectIpsCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIpsCustomAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsCustomApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsCustomComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsCustomLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsCustomLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsCustomLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsCustomOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsCustomProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsCustomRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsCustomSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsCustomSigName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsCustomSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsCustomStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsCustomTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectIpsCustom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectIpsCustomAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandObjectIpsCustomApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectIpsCustomComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok || d.HasChange("location") {
		t, err := expandObjectIpsCustomLocation(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectIpsCustomLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok || d.HasChange("log_packet") {
		t, err := expandObjectIpsCustomLogPacket(d, v, "log_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok || d.HasChange("os") {
		t, err := expandObjectIpsCustomOs(d, v, "os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectIpsCustomProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("rule_id"); ok || d.HasChange("rule_id") {
		t, err := expandObjectIpsCustomRuleId(d, v, "rule_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule-id"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandObjectIpsCustomSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("sig_name"); ok || d.HasChange("sig_name") {
		t, err := expandObjectIpsCustomSigName(d, v, "sig_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sig-name"] = t
		}
	}

	if v, ok := d.GetOk("signature"); ok || d.HasChange("signature") {
		t, err := expandObjectIpsCustomSignature(d, v, "signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectIpsCustomStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tag"); ok || d.HasChange("tag") {
		t, err := expandObjectIpsCustomTag(d, v, "tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	return &obj, nil
}
