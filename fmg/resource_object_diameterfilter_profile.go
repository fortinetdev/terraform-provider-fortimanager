// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Diameter filter profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDiameterFilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDiameterFilterProfileCreate,
		Read:   resourceObjectDiameterFilterProfileRead,
		Update: resourceObjectDiameterFilterProfileUpdate,
		Delete: resourceObjectDiameterFilterProfileDelete,

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
			"cmd_flags_reserve_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command_code_invalid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command_code_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message_length_invalid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"missing_request_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_all_messages": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"protocol_version_invalid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_error_flag_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"track_requests_answers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectDiameterFilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectDiameterFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDiameterFilterProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDiameterFilterProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDiameterFilterProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDiameterFilterProfileRead(d, m)
}

func resourceObjectDiameterFilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectDiameterFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDiameterFilterProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDiameterFilterProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDiameterFilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDiameterFilterProfileRead(d, m)
}

func resourceObjectDiameterFilterProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectDiameterFilterProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDiameterFilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDiameterFilterProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectDiameterFilterProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDiameterFilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDiameterFilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDiameterFilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectDiameterFilterProfileCmdFlagsReserveSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileCommandCodeInvalid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileCommandCodeRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileMessageLengthInvalid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileMissingRequestAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileMonitorAllMessages(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileProtocolVersionInvalid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileRequestErrorFlagSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDiameterFilterProfileTrackRequestsAnswers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDiameterFilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("cmd_flags_reserve_set", flattenObjectDiameterFilterProfileCmdFlagsReserveSet(o["cmd-flags-reserve-set"], d, "cmd_flags_reserve_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["cmd-flags-reserve-set"], "ObjectDiameterFilterProfile-CmdFlagsReserveSet"); ok {
			if err = d.Set("cmd_flags_reserve_set", vv); err != nil {
				return fmt.Errorf("Error reading cmd_flags_reserve_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cmd_flags_reserve_set: %v", err)
		}
	}

	if err = d.Set("command_code_invalid", flattenObjectDiameterFilterProfileCommandCodeInvalid(o["command-code-invalid"], d, "command_code_invalid")); err != nil {
		if vv, ok := fortiAPIPatch(o["command-code-invalid"], "ObjectDiameterFilterProfile-CommandCodeInvalid"); ok {
			if err = d.Set("command_code_invalid", vv); err != nil {
				return fmt.Errorf("Error reading command_code_invalid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command_code_invalid: %v", err)
		}
	}

	if err = d.Set("command_code_range", flattenObjectDiameterFilterProfileCommandCodeRange(o["command-code-range"], d, "command_code_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["command-code-range"], "ObjectDiameterFilterProfile-CommandCodeRange"); ok {
			if err = d.Set("command_code_range", vv); err != nil {
				return fmt.Errorf("Error reading command_code_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command_code_range: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectDiameterFilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectDiameterFilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenObjectDiameterFilterProfileLogPacket(o["log-packet"], d, "log_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-packet"], "ObjectDiameterFilterProfile-LogPacket"); ok {
			if err = d.Set("log_packet", vv); err != nil {
				return fmt.Errorf("Error reading log_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("message_length_invalid", flattenObjectDiameterFilterProfileMessageLengthInvalid(o["message-length-invalid"], d, "message_length_invalid")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-length-invalid"], "ObjectDiameterFilterProfile-MessageLengthInvalid"); ok {
			if err = d.Set("message_length_invalid", vv); err != nil {
				return fmt.Errorf("Error reading message_length_invalid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_length_invalid: %v", err)
		}
	}

	if err = d.Set("missing_request_action", flattenObjectDiameterFilterProfileMissingRequestAction(o["missing-request-action"], d, "missing_request_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["missing-request-action"], "ObjectDiameterFilterProfile-MissingRequestAction"); ok {
			if err = d.Set("missing_request_action", vv); err != nil {
				return fmt.Errorf("Error reading missing_request_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading missing_request_action: %v", err)
		}
	}

	if err = d.Set("monitor_all_messages", flattenObjectDiameterFilterProfileMonitorAllMessages(o["monitor-all-messages"], d, "monitor_all_messages")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-all-messages"], "ObjectDiameterFilterProfile-MonitorAllMessages"); ok {
			if err = d.Set("monitor_all_messages", vv); err != nil {
				return fmt.Errorf("Error reading monitor_all_messages: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_all_messages: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDiameterFilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDiameterFilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol_version_invalid", flattenObjectDiameterFilterProfileProtocolVersionInvalid(o["protocol-version-invalid"], d, "protocol_version_invalid")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol-version-invalid"], "ObjectDiameterFilterProfile-ProtocolVersionInvalid"); ok {
			if err = d.Set("protocol_version_invalid", vv); err != nil {
				return fmt.Errorf("Error reading protocol_version_invalid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol_version_invalid: %v", err)
		}
	}

	if err = d.Set("request_error_flag_set", flattenObjectDiameterFilterProfileRequestErrorFlagSet(o["request-error-flag-set"], d, "request_error_flag_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-error-flag-set"], "ObjectDiameterFilterProfile-RequestErrorFlagSet"); ok {
			if err = d.Set("request_error_flag_set", vv); err != nil {
				return fmt.Errorf("Error reading request_error_flag_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_error_flag_set: %v", err)
		}
	}

	if err = d.Set("track_requests_answers", flattenObjectDiameterFilterProfileTrackRequestsAnswers(o["track-requests-answers"], d, "track_requests_answers")); err != nil {
		if vv, ok := fortiAPIPatch(o["track-requests-answers"], "ObjectDiameterFilterProfile-TrackRequestsAnswers"); ok {
			if err = d.Set("track_requests_answers", vv); err != nil {
				return fmt.Errorf("Error reading track_requests_answers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading track_requests_answers: %v", err)
		}
	}

	return nil
}

func flattenObjectDiameterFilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDiameterFilterProfileCmdFlagsReserveSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileCommandCodeInvalid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileCommandCodeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileMessageLengthInvalid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileMissingRequestAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileMonitorAllMessages(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileProtocolVersionInvalid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileRequestErrorFlagSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDiameterFilterProfileTrackRequestsAnswers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDiameterFilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cmd_flags_reserve_set"); ok || d.HasChange("cmd_flags_reserve_set") {
		t, err := expandObjectDiameterFilterProfileCmdFlagsReserveSet(d, v, "cmd_flags_reserve_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmd-flags-reserve-set"] = t
		}
	}

	if v, ok := d.GetOk("command_code_invalid"); ok || d.HasChange("command_code_invalid") {
		t, err := expandObjectDiameterFilterProfileCommandCodeInvalid(d, v, "command_code_invalid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-code-invalid"] = t
		}
	}

	if v, ok := d.GetOk("command_code_range"); ok || d.HasChange("command_code_range") {
		t, err := expandObjectDiameterFilterProfileCommandCodeRange(d, v, "command_code_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-code-range"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectDiameterFilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok || d.HasChange("log_packet") {
		t, err := expandObjectDiameterFilterProfileLogPacket(d, v, "log_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("message_length_invalid"); ok || d.HasChange("message_length_invalid") {
		t, err := expandObjectDiameterFilterProfileMessageLengthInvalid(d, v, "message_length_invalid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-length-invalid"] = t
		}
	}

	if v, ok := d.GetOk("missing_request_action"); ok || d.HasChange("missing_request_action") {
		t, err := expandObjectDiameterFilterProfileMissingRequestAction(d, v, "missing_request_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["missing-request-action"] = t
		}
	}

	if v, ok := d.GetOk("monitor_all_messages"); ok || d.HasChange("monitor_all_messages") {
		t, err := expandObjectDiameterFilterProfileMonitorAllMessages(d, v, "monitor_all_messages")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-all-messages"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDiameterFilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol_version_invalid"); ok || d.HasChange("protocol_version_invalid") {
		t, err := expandObjectDiameterFilterProfileProtocolVersionInvalid(d, v, "protocol_version_invalid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol-version-invalid"] = t
		}
	}

	if v, ok := d.GetOk("request_error_flag_set"); ok || d.HasChange("request_error_flag_set") {
		t, err := expandObjectDiameterFilterProfileRequestErrorFlagSet(d, v, "request_error_flag_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-error-flag-set"] = t
		}
	}

	if v, ok := d.GetOk("track_requests_answers"); ok || d.HasChange("track_requests_answers") {
		t, err := expandObjectDiameterFilterProfileTrackRequestsAnswers(d, v, "track_requests_answers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["track-requests-answers"] = t
		}
	}

	return &obj, nil
}
