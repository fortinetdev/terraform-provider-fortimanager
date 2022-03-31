// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VoIP profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectVoipProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVoipProfileCreate,
		Read:   resourceObjectVoipProfileRead,
		Update: resourceObjectVoipProfileUpdate,
		Delete: resourceObjectVoipProfileDelete,

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
				Computed: true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"sccp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"block_mcast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_call_summary": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_violations": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_calls": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"verify_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"sip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ack_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ack_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_ack": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_bye": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_cancel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_geo_red_options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_info": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_invite": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_long_lines": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_message": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_notify": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_prack": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_publish": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_refer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_register": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_subscribe": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_unknown": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_update": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bye_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bye_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"call_keepalive": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"cancel_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"cancel_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"contact_fixup": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hnt_restrict_source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hosted_nat_traversal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"info_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"info_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invite_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"invite_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ips_rtp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_call_summary": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_violations": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_allow": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_call_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_contact": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_content_length": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_content_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_cseq": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_expires": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_from": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_max_forwards": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_no_proxy_require": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_no_require": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_p_asserted_identity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_rack": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_record_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_rseq": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_a": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_b": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_c": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_i": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_k": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_m": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_o": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_r": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_s": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_t": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_z": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_to": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_via": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_request_line": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_body_length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_dialogs": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_idle_dialogs": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_line_length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"message_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"message_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nat_port_range": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nat_trace": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"no_sdp_fixup": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"notify_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"notify_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"open_contact_pinhole": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"open_record_route_pinhole": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"open_register_pinhole": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"open_via_pinhole": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"options_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prack_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prack_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preserve_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"provisional_invite_expiry_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"publish_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"publish_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"refer_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"refer_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_contact_trace": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"register_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rfc2543_branch": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rtp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_auth_client": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_auth_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_client_renegotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_max_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_min_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_pfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_send_empty_frags": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_server_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"strict_register": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subscribe_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"subscribe_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unknown_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"update_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"update_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceObjectVoipProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVoipProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVoipProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVoipProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVoipProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVoipProfileRead(d, m)
}

func resourceObjectVoipProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVoipProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVoipProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVoipProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVoipProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVoipProfileRead(d, m)
}

func resourceObjectVoipProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectVoipProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVoipProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVoipProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectVoipProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVoipProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVoipProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVoipProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectVoipProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "block_mcast"
	if _, ok := i["block-mcast"]; ok {
		result["block_mcast"] = flattenObjectVoipProfileSccpBlockMcast(i["block-mcast"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_call_summary"
	if _, ok := i["log-call-summary"]; ok {
		result["log_call_summary"] = flattenObjectVoipProfileSccpLogCallSummary(i["log-call-summary"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_violations"
	if _, ok := i["log-violations"]; ok {
		result["log_violations"] = flattenObjectVoipProfileSccpLogViolations(i["log-violations"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_calls"
	if _, ok := i["max-calls"]; ok {
		result["max_calls"] = flattenObjectVoipProfileSccpMaxCalls(i["max-calls"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectVoipProfileSccpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "verify_header"
	if _, ok := i["verify-header"]; ok {
		result["verify_header"] = flattenObjectVoipProfileSccpVerifyHeader(i["verify-header"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectVoipProfileSccpBlockMcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccpLogCallSummary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccpLogViolations(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccpMaxCalls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccpVerifyHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSip(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_rate"
	if _, ok := i["ack-rate"]; ok {
		result["ack_rate"] = flattenObjectVoipProfileSipAckRate(i["ack-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "ack_rate_track"
	if _, ok := i["ack-rate-track"]; ok {
		result["ack_rate_track"] = flattenObjectVoipProfileSipAckRateTrack(i["ack-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_ack"
	if _, ok := i["block-ack"]; ok {
		result["block_ack"] = flattenObjectVoipProfileSipBlockAck(i["block-ack"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_bye"
	if _, ok := i["block-bye"]; ok {
		result["block_bye"] = flattenObjectVoipProfileSipBlockBye(i["block-bye"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_cancel"
	if _, ok := i["block-cancel"]; ok {
		result["block_cancel"] = flattenObjectVoipProfileSipBlockCancel(i["block-cancel"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_geo_red_options"
	if _, ok := i["block-geo-red-options"]; ok {
		result["block_geo_red_options"] = flattenObjectVoipProfileSipBlockGeoRedOptions(i["block-geo-red-options"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_info"
	if _, ok := i["block-info"]; ok {
		result["block_info"] = flattenObjectVoipProfileSipBlockInfo(i["block-info"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_invite"
	if _, ok := i["block-invite"]; ok {
		result["block_invite"] = flattenObjectVoipProfileSipBlockInvite(i["block-invite"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_long_lines"
	if _, ok := i["block-long-lines"]; ok {
		result["block_long_lines"] = flattenObjectVoipProfileSipBlockLongLines(i["block-long-lines"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_message"
	if _, ok := i["block-message"]; ok {
		result["block_message"] = flattenObjectVoipProfileSipBlockMessage(i["block-message"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_notify"
	if _, ok := i["block-notify"]; ok {
		result["block_notify"] = flattenObjectVoipProfileSipBlockNotify(i["block-notify"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_options"
	if _, ok := i["block-options"]; ok {
		result["block_options"] = flattenObjectVoipProfileSipBlockOptions(i["block-options"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_prack"
	if _, ok := i["block-prack"]; ok {
		result["block_prack"] = flattenObjectVoipProfileSipBlockPrack(i["block-prack"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_publish"
	if _, ok := i["block-publish"]; ok {
		result["block_publish"] = flattenObjectVoipProfileSipBlockPublish(i["block-publish"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_refer"
	if _, ok := i["block-refer"]; ok {
		result["block_refer"] = flattenObjectVoipProfileSipBlockRefer(i["block-refer"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_register"
	if _, ok := i["block-register"]; ok {
		result["block_register"] = flattenObjectVoipProfileSipBlockRegister(i["block-register"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_subscribe"
	if _, ok := i["block-subscribe"]; ok {
		result["block_subscribe"] = flattenObjectVoipProfileSipBlockSubscribe(i["block-subscribe"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_unknown"
	if _, ok := i["block-unknown"]; ok {
		result["block_unknown"] = flattenObjectVoipProfileSipBlockUnknown(i["block-unknown"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_update"
	if _, ok := i["block-update"]; ok {
		result["block_update"] = flattenObjectVoipProfileSipBlockUpdate(i["block-update"], d, pre_append)
	}

	pre_append = pre + ".0." + "bye_rate"
	if _, ok := i["bye-rate"]; ok {
		result["bye_rate"] = flattenObjectVoipProfileSipByeRate(i["bye-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "bye_rate_track"
	if _, ok := i["bye-rate-track"]; ok {
		result["bye_rate_track"] = flattenObjectVoipProfileSipByeRateTrack(i["bye-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_keepalive"
	if _, ok := i["call-keepalive"]; ok {
		result["call_keepalive"] = flattenObjectVoipProfileSipCallKeepalive(i["call-keepalive"], d, pre_append)
	}

	pre_append = pre + ".0." + "cancel_rate"
	if _, ok := i["cancel-rate"]; ok {
		result["cancel_rate"] = flattenObjectVoipProfileSipCancelRate(i["cancel-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "cancel_rate_track"
	if _, ok := i["cancel-rate-track"]; ok {
		result["cancel_rate_track"] = flattenObjectVoipProfileSipCancelRateTrack(i["cancel-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "contact_fixup"
	if _, ok := i["contact-fixup"]; ok {
		result["contact_fixup"] = flattenObjectVoipProfileSipContactFixup(i["contact-fixup"], d, pre_append)
	}

	pre_append = pre + ".0." + "hnt_restrict_source_ip"
	if _, ok := i["hnt-restrict-source-ip"]; ok {
		result["hnt_restrict_source_ip"] = flattenObjectVoipProfileSipHntRestrictSourceIp(i["hnt-restrict-source-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "hosted_nat_traversal"
	if _, ok := i["hosted-nat-traversal"]; ok {
		result["hosted_nat_traversal"] = flattenObjectVoipProfileSipHostedNatTraversal(i["hosted-nat-traversal"], d, pre_append)
	}

	pre_append = pre + ".0." + "info_rate"
	if _, ok := i["info-rate"]; ok {
		result["info_rate"] = flattenObjectVoipProfileSipInfoRate(i["info-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "info_rate_track"
	if _, ok := i["info-rate-track"]; ok {
		result["info_rate_track"] = flattenObjectVoipProfileSipInfoRateTrack(i["info-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "invite_rate"
	if _, ok := i["invite-rate"]; ok {
		result["invite_rate"] = flattenObjectVoipProfileSipInviteRate(i["invite-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "invite_rate_track"
	if _, ok := i["invite-rate-track"]; ok {
		result["invite_rate_track"] = flattenObjectVoipProfileSipInviteRateTrack(i["invite-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "ips_rtp"
	if _, ok := i["ips-rtp"]; ok {
		result["ips_rtp"] = flattenObjectVoipProfileSipIpsRtp(i["ips-rtp"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_call_summary"
	if _, ok := i["log-call-summary"]; ok {
		result["log_call_summary"] = flattenObjectVoipProfileSipLogCallSummary(i["log-call-summary"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_violations"
	if _, ok := i["log-violations"]; ok {
		result["log_violations"] = flattenObjectVoipProfileSipLogViolations(i["log-violations"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_allow"
	if _, ok := i["malformed-header-allow"]; ok {
		result["malformed_header_allow"] = flattenObjectVoipProfileSipMalformedHeaderAllow(i["malformed-header-allow"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_call_id"
	if _, ok := i["malformed-header-call-id"]; ok {
		result["malformed_header_call_id"] = flattenObjectVoipProfileSipMalformedHeaderCallId(i["malformed-header-call-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_contact"
	if _, ok := i["malformed-header-contact"]; ok {
		result["malformed_header_contact"] = flattenObjectVoipProfileSipMalformedHeaderContact(i["malformed-header-contact"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_content_length"
	if _, ok := i["malformed-header-content-length"]; ok {
		result["malformed_header_content_length"] = flattenObjectVoipProfileSipMalformedHeaderContentLength(i["malformed-header-content-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_content_type"
	if _, ok := i["malformed-header-content-type"]; ok {
		result["malformed_header_content_type"] = flattenObjectVoipProfileSipMalformedHeaderContentType(i["malformed-header-content-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_cseq"
	if _, ok := i["malformed-header-cseq"]; ok {
		result["malformed_header_cseq"] = flattenObjectVoipProfileSipMalformedHeaderCseq(i["malformed-header-cseq"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_expires"
	if _, ok := i["malformed-header-expires"]; ok {
		result["malformed_header_expires"] = flattenObjectVoipProfileSipMalformedHeaderExpires(i["malformed-header-expires"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_from"
	if _, ok := i["malformed-header-from"]; ok {
		result["malformed_header_from"] = flattenObjectVoipProfileSipMalformedHeaderFrom(i["malformed-header-from"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_max_forwards"
	if _, ok := i["malformed-header-max-forwards"]; ok {
		result["malformed_header_max_forwards"] = flattenObjectVoipProfileSipMalformedHeaderMaxForwards(i["malformed-header-max-forwards"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_no_proxy_require"
	if _, ok := i["malformed-header-no-proxy-require"]; ok {
		result["malformed_header_no_proxy_require"] = flattenObjectVoipProfileSipMalformedHeaderNoProxyRequire(i["malformed-header-no-proxy-require"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_no_require"
	if _, ok := i["malformed-header-no-require"]; ok {
		result["malformed_header_no_require"] = flattenObjectVoipProfileSipMalformedHeaderNoRequire(i["malformed-header-no-require"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_p_asserted_identity"
	if _, ok := i["malformed-header-p-asserted-identity"]; ok {
		result["malformed_header_p_asserted_identity"] = flattenObjectVoipProfileSipMalformedHeaderPAssertedIdentity(i["malformed-header-p-asserted-identity"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_rack"
	if _, ok := i["malformed-header-rack"]; ok {
		result["malformed_header_rack"] = flattenObjectVoipProfileSipMalformedHeaderRack(i["malformed-header-rack"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_record_route"
	if _, ok := i["malformed-header-record-route"]; ok {
		result["malformed_header_record_route"] = flattenObjectVoipProfileSipMalformedHeaderRecordRoute(i["malformed-header-record-route"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_route"
	if _, ok := i["malformed-header-route"]; ok {
		result["malformed_header_route"] = flattenObjectVoipProfileSipMalformedHeaderRoute(i["malformed-header-route"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_rseq"
	if _, ok := i["malformed-header-rseq"]; ok {
		result["malformed_header_rseq"] = flattenObjectVoipProfileSipMalformedHeaderRseq(i["malformed-header-rseq"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_a"
	if _, ok := i["malformed-header-sdp-a"]; ok {
		result["malformed_header_sdp_a"] = flattenObjectVoipProfileSipMalformedHeaderSdpA(i["malformed-header-sdp-a"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_b"
	if _, ok := i["malformed-header-sdp-b"]; ok {
		result["malformed_header_sdp_b"] = flattenObjectVoipProfileSipMalformedHeaderSdpB(i["malformed-header-sdp-b"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_c"
	if _, ok := i["malformed-header-sdp-c"]; ok {
		result["malformed_header_sdp_c"] = flattenObjectVoipProfileSipMalformedHeaderSdpC(i["malformed-header-sdp-c"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_i"
	if _, ok := i["malformed-header-sdp-i"]; ok {
		result["malformed_header_sdp_i"] = flattenObjectVoipProfileSipMalformedHeaderSdpI(i["malformed-header-sdp-i"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_k"
	if _, ok := i["malformed-header-sdp-k"]; ok {
		result["malformed_header_sdp_k"] = flattenObjectVoipProfileSipMalformedHeaderSdpK(i["malformed-header-sdp-k"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_m"
	if _, ok := i["malformed-header-sdp-m"]; ok {
		result["malformed_header_sdp_m"] = flattenObjectVoipProfileSipMalformedHeaderSdpM(i["malformed-header-sdp-m"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_o"
	if _, ok := i["malformed-header-sdp-o"]; ok {
		result["malformed_header_sdp_o"] = flattenObjectVoipProfileSipMalformedHeaderSdpO(i["malformed-header-sdp-o"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_r"
	if _, ok := i["malformed-header-sdp-r"]; ok {
		result["malformed_header_sdp_r"] = flattenObjectVoipProfileSipMalformedHeaderSdpR(i["malformed-header-sdp-r"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_s"
	if _, ok := i["malformed-header-sdp-s"]; ok {
		result["malformed_header_sdp_s"] = flattenObjectVoipProfileSipMalformedHeaderSdpS(i["malformed-header-sdp-s"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_t"
	if _, ok := i["malformed-header-sdp-t"]; ok {
		result["malformed_header_sdp_t"] = flattenObjectVoipProfileSipMalformedHeaderSdpT(i["malformed-header-sdp-t"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_v"
	if _, ok := i["malformed-header-sdp-v"]; ok {
		result["malformed_header_sdp_v"] = flattenObjectVoipProfileSipMalformedHeaderSdpV(i["malformed-header-sdp-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_z"
	if _, ok := i["malformed-header-sdp-z"]; ok {
		result["malformed_header_sdp_z"] = flattenObjectVoipProfileSipMalformedHeaderSdpZ(i["malformed-header-sdp-z"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_to"
	if _, ok := i["malformed-header-to"]; ok {
		result["malformed_header_to"] = flattenObjectVoipProfileSipMalformedHeaderTo(i["malformed-header-to"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_header_via"
	if _, ok := i["malformed-header-via"]; ok {
		result["malformed_header_via"] = flattenObjectVoipProfileSipMalformedHeaderVia(i["malformed-header-via"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed_request_line"
	if _, ok := i["malformed-request-line"]; ok {
		result["malformed_request_line"] = flattenObjectVoipProfileSipMalformedRequestLine(i["malformed-request-line"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_body_length"
	if _, ok := i["max-body-length"]; ok {
		result["max_body_length"] = flattenObjectVoipProfileSipMaxBodyLength(i["max-body-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_dialogs"
	if _, ok := i["max-dialogs"]; ok {
		result["max_dialogs"] = flattenObjectVoipProfileSipMaxDialogs(i["max-dialogs"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_idle_dialogs"
	if _, ok := i["max-idle-dialogs"]; ok {
		result["max_idle_dialogs"] = flattenObjectVoipProfileSipMaxIdleDialogs(i["max-idle-dialogs"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_line_length"
	if _, ok := i["max-line-length"]; ok {
		result["max_line_length"] = flattenObjectVoipProfileSipMaxLineLength(i["max-line-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "message_rate"
	if _, ok := i["message-rate"]; ok {
		result["message_rate"] = flattenObjectVoipProfileSipMessageRate(i["message-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "message_rate_track"
	if _, ok := i["message-rate-track"]; ok {
		result["message_rate_track"] = flattenObjectVoipProfileSipMessageRateTrack(i["message-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "nat_port_range"
	if _, ok := i["nat-port-range"]; ok {
		result["nat_port_range"] = flattenObjectVoipProfileSipNatPortRange(i["nat-port-range"], d, pre_append)
	}

	pre_append = pre + ".0." + "nat_trace"
	if _, ok := i["nat-trace"]; ok {
		result["nat_trace"] = flattenObjectVoipProfileSipNatTrace(i["nat-trace"], d, pre_append)
	}

	pre_append = pre + ".0." + "no_sdp_fixup"
	if _, ok := i["no-sdp-fixup"]; ok {
		result["no_sdp_fixup"] = flattenObjectVoipProfileSipNoSdpFixup(i["no-sdp-fixup"], d, pre_append)
	}

	pre_append = pre + ".0." + "notify_rate"
	if _, ok := i["notify-rate"]; ok {
		result["notify_rate"] = flattenObjectVoipProfileSipNotifyRate(i["notify-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "notify_rate_track"
	if _, ok := i["notify-rate-track"]; ok {
		result["notify_rate_track"] = flattenObjectVoipProfileSipNotifyRateTrack(i["notify-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "open_contact_pinhole"
	if _, ok := i["open-contact-pinhole"]; ok {
		result["open_contact_pinhole"] = flattenObjectVoipProfileSipOpenContactPinhole(i["open-contact-pinhole"], d, pre_append)
	}

	pre_append = pre + ".0." + "open_record_route_pinhole"
	if _, ok := i["open-record-route-pinhole"]; ok {
		result["open_record_route_pinhole"] = flattenObjectVoipProfileSipOpenRecordRoutePinhole(i["open-record-route-pinhole"], d, pre_append)
	}

	pre_append = pre + ".0." + "open_register_pinhole"
	if _, ok := i["open-register-pinhole"]; ok {
		result["open_register_pinhole"] = flattenObjectVoipProfileSipOpenRegisterPinhole(i["open-register-pinhole"], d, pre_append)
	}

	pre_append = pre + ".0." + "open_via_pinhole"
	if _, ok := i["open-via-pinhole"]; ok {
		result["open_via_pinhole"] = flattenObjectVoipProfileSipOpenViaPinhole(i["open-via-pinhole"], d, pre_append)
	}

	pre_append = pre + ".0." + "options_rate"
	if _, ok := i["options-rate"]; ok {
		result["options_rate"] = flattenObjectVoipProfileSipOptionsRate(i["options-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "options_rate_track"
	if _, ok := i["options-rate-track"]; ok {
		result["options_rate_track"] = flattenObjectVoipProfileSipOptionsRateTrack(i["options-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "prack_rate"
	if _, ok := i["prack-rate"]; ok {
		result["prack_rate"] = flattenObjectVoipProfileSipPrackRate(i["prack-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "prack_rate_track"
	if _, ok := i["prack-rate-track"]; ok {
		result["prack_rate_track"] = flattenObjectVoipProfileSipPrackRateTrack(i["prack-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "preserve_override"
	if _, ok := i["preserve-override"]; ok {
		result["preserve_override"] = flattenObjectVoipProfileSipPreserveOverride(i["preserve-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "provisional_invite_expiry_time"
	if _, ok := i["provisional-invite-expiry-time"]; ok {
		result["provisional_invite_expiry_time"] = flattenObjectVoipProfileSipProvisionalInviteExpiryTime(i["provisional-invite-expiry-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "publish_rate"
	if _, ok := i["publish-rate"]; ok {
		result["publish_rate"] = flattenObjectVoipProfileSipPublishRate(i["publish-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "publish_rate_track"
	if _, ok := i["publish-rate-track"]; ok {
		result["publish_rate_track"] = flattenObjectVoipProfileSipPublishRateTrack(i["publish-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "refer_rate"
	if _, ok := i["refer-rate"]; ok {
		result["refer_rate"] = flattenObjectVoipProfileSipReferRate(i["refer-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "refer_rate_track"
	if _, ok := i["refer-rate-track"]; ok {
		result["refer_rate_track"] = flattenObjectVoipProfileSipReferRateTrack(i["refer-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_contact_trace"
	if _, ok := i["register-contact-trace"]; ok {
		result["register_contact_trace"] = flattenObjectVoipProfileSipRegisterContactTrace(i["register-contact-trace"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_rate"
	if _, ok := i["register-rate"]; ok {
		result["register_rate"] = flattenObjectVoipProfileSipRegisterRate(i["register-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_rate_track"
	if _, ok := i["register-rate-track"]; ok {
		result["register_rate_track"] = flattenObjectVoipProfileSipRegisterRateTrack(i["register-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "rfc2543_branch"
	if _, ok := i["rfc2543-branch"]; ok {
		result["rfc2543_branch"] = flattenObjectVoipProfileSipRfc2543Branch(i["rfc2543-branch"], d, pre_append)
	}

	pre_append = pre + ".0." + "rtp"
	if _, ok := i["rtp"]; ok {
		result["rtp"] = flattenObjectVoipProfileSipRtp(i["rtp"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_algorithm"
	if _, ok := i["ssl-algorithm"]; ok {
		result["ssl_algorithm"] = flattenObjectVoipProfileSipSslAlgorithm(i["ssl-algorithm"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_auth_client"
	if _, ok := i["ssl-auth-client"]; ok {
		result["ssl_auth_client"] = flattenObjectVoipProfileSipSslAuthClient(i["ssl-auth-client"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_auth_server"
	if _, ok := i["ssl-auth-server"]; ok {
		result["ssl_auth_server"] = flattenObjectVoipProfileSipSslAuthServer(i["ssl-auth-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_client_certificate"
	if _, ok := i["ssl-client-certificate"]; ok {
		result["ssl_client_certificate"] = flattenObjectVoipProfileSipSslClientCertificate(i["ssl-client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_client_renegotiation"
	if _, ok := i["ssl-client-renegotiation"]; ok {
		result["ssl_client_renegotiation"] = flattenObjectVoipProfileSipSslClientRenegotiation(i["ssl-client-renegotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_max_version"
	if _, ok := i["ssl-max-version"]; ok {
		result["ssl_max_version"] = flattenObjectVoipProfileSipSslMaxVersion(i["ssl-max-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_min_version"
	if _, ok := i["ssl-min-version"]; ok {
		result["ssl_min_version"] = flattenObjectVoipProfileSipSslMinVersion(i["ssl-min-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_mode"
	if _, ok := i["ssl-mode"]; ok {
		result["ssl_mode"] = flattenObjectVoipProfileSipSslMode(i["ssl-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_pfs"
	if _, ok := i["ssl-pfs"]; ok {
		result["ssl_pfs"] = flattenObjectVoipProfileSipSslPfs(i["ssl-pfs"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_send_empty_frags"
	if _, ok := i["ssl-send-empty-frags"]; ok {
		result["ssl_send_empty_frags"] = flattenObjectVoipProfileSipSslSendEmptyFrags(i["ssl-send-empty-frags"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_server_certificate"
	if _, ok := i["ssl-server-certificate"]; ok {
		result["ssl_server_certificate"] = flattenObjectVoipProfileSipSslServerCertificate(i["ssl-server-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectVoipProfileSipStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "strict_register"
	if _, ok := i["strict-register"]; ok {
		result["strict_register"] = flattenObjectVoipProfileSipStrictRegister(i["strict-register"], d, pre_append)
	}

	pre_append = pre + ".0." + "subscribe_rate"
	if _, ok := i["subscribe-rate"]; ok {
		result["subscribe_rate"] = flattenObjectVoipProfileSipSubscribeRate(i["subscribe-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "subscribe_rate_track"
	if _, ok := i["subscribe-rate-track"]; ok {
		result["subscribe_rate_track"] = flattenObjectVoipProfileSipSubscribeRateTrack(i["subscribe-rate-track"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknown_header"
	if _, ok := i["unknown-header"]; ok {
		result["unknown_header"] = flattenObjectVoipProfileSipUnknownHeader(i["unknown-header"], d, pre_append)
	}

	pre_append = pre + ".0." + "update_rate"
	if _, ok := i["update-rate"]; ok {
		result["update_rate"] = flattenObjectVoipProfileSipUpdateRate(i["update-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "update_rate_track"
	if _, ok := i["update-rate-track"]; ok {
		result["update_rate_track"] = flattenObjectVoipProfileSipUpdateRateTrack(i["update-rate-track"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectVoipProfileSipAckRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipAckRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockAck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockBye(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockCancel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockGeoRedOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockInvite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockLongLines(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockNotify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockPrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockPublish(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockRefer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockRegister(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockSubscribe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockUnknown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipByeRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipByeRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipCallKeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipCancelRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipCancelRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipContactFixup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipHntRestrictSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipHostedNatTraversal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipInfoRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipInfoRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipInviteRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipInviteRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipIpsRtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipLogCallSummary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipLogViolations(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderAllow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderCallId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderContact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderContentLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderContentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderCseq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderExpires(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderMaxForwards(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderNoProxyRequire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderNoRequire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderPAssertedIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderRack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderRecordRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderRseq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpA(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpB(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpC(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpI(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpK(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpM(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpO(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpS(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpT(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpZ(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderVia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedRequestLine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMaxBodyLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMaxDialogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMaxIdleDialogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMaxLineLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMessageRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMessageRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipNatPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipNatTrace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipNoSdpFixup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipNotifyRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipNotifyRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOpenContactPinhole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOpenRecordRoutePinhole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOpenRegisterPinhole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOpenViaPinhole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOptionsRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOptionsRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipPrackRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipPrackRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipPreserveOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipProvisionalInviteExpiryTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipPublishRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipPublishRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipReferRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipReferRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipRegisterContactTrace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipRegisterRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipRegisterRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipRfc2543Branch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipRtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslAuthClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslAuthServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslPfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslServerCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipStrictRegister(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSubscribeRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSubscribeRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipUnknownHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipUpdateRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipUpdateRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVoipProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comment", flattenObjectVoipProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectVoipProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenObjectVoipProfileFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "ObjectVoipProfile-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVoipProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVoipProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sccp", flattenObjectVoipProfileSccp(o["sccp"], d, "sccp")); err != nil {
			if vv, ok := fortiAPIPatch(o["sccp"], "ObjectVoipProfile-Sccp"); ok {
				if err = d.Set("sccp", vv); err != nil {
					return fmt.Errorf("Error reading sccp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sccp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sccp"); ok {
			if err = d.Set("sccp", flattenObjectVoipProfileSccp(o["sccp"], d, "sccp")); err != nil {
				if vv, ok := fortiAPIPatch(o["sccp"], "ObjectVoipProfile-Sccp"); ok {
					if err = d.Set("sccp", vv); err != nil {
						return fmt.Errorf("Error reading sccp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sccp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("sip", flattenObjectVoipProfileSip(o["sip"], d, "sip")); err != nil {
			if vv, ok := fortiAPIPatch(o["sip"], "ObjectVoipProfile-Sip"); ok {
				if err = d.Set("sip", vv); err != nil {
					return fmt.Errorf("Error reading sip: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sip"); ok {
			if err = d.Set("sip", flattenObjectVoipProfileSip(o["sip"], d, "sip")); err != nil {
				if vv, ok := fortiAPIPatch(o["sip"], "ObjectVoipProfile-Sip"); ok {
					if err = d.Set("sip", vv); err != nil {
						return fmt.Errorf("Error reading sip: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sip: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectVoipProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVoipProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "block_mcast"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-mcast"], _ = expandObjectVoipProfileSccpBlockMcast(d, i["block_mcast"], pre_append)
	}
	pre_append = pre + ".0." + "log_call_summary"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-call-summary"], _ = expandObjectVoipProfileSccpLogCallSummary(d, i["log_call_summary"], pre_append)
	}
	pre_append = pre + ".0." + "log_violations"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-violations"], _ = expandObjectVoipProfileSccpLogViolations(d, i["log_violations"], pre_append)
	}
	pre_append = pre + ".0." + "max_calls"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-calls"], _ = expandObjectVoipProfileSccpMaxCalls(d, i["max_calls"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectVoipProfileSccpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "verify_header"
	if _, ok := d.GetOk(pre_append); ok {
		result["verify-header"], _ = expandObjectVoipProfileSccpVerifyHeader(d, i["verify_header"], pre_append)
	}

	return result, nil
}

func expandObjectVoipProfileSccpBlockMcast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccpLogCallSummary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccpLogViolations(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccpMaxCalls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccpVerifyHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ack_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["ack-rate"], _ = expandObjectVoipProfileSipAckRate(d, i["ack_rate"], pre_append)
	}
	pre_append = pre + ".0." + "ack_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["ack-rate-track"], _ = expandObjectVoipProfileSipAckRateTrack(d, i["ack_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "block_ack"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-ack"], _ = expandObjectVoipProfileSipBlockAck(d, i["block_ack"], pre_append)
	}
	pre_append = pre + ".0." + "block_bye"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-bye"], _ = expandObjectVoipProfileSipBlockBye(d, i["block_bye"], pre_append)
	}
	pre_append = pre + ".0." + "block_cancel"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-cancel"], _ = expandObjectVoipProfileSipBlockCancel(d, i["block_cancel"], pre_append)
	}
	pre_append = pre + ".0." + "block_geo_red_options"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-geo-red-options"], _ = expandObjectVoipProfileSipBlockGeoRedOptions(d, i["block_geo_red_options"], pre_append)
	}
	pre_append = pre + ".0." + "block_info"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-info"], _ = expandObjectVoipProfileSipBlockInfo(d, i["block_info"], pre_append)
	}
	pre_append = pre + ".0." + "block_invite"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-invite"], _ = expandObjectVoipProfileSipBlockInvite(d, i["block_invite"], pre_append)
	}
	pre_append = pre + ".0." + "block_long_lines"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-long-lines"], _ = expandObjectVoipProfileSipBlockLongLines(d, i["block_long_lines"], pre_append)
	}
	pre_append = pre + ".0." + "block_message"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-message"], _ = expandObjectVoipProfileSipBlockMessage(d, i["block_message"], pre_append)
	}
	pre_append = pre + ".0." + "block_notify"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-notify"], _ = expandObjectVoipProfileSipBlockNotify(d, i["block_notify"], pre_append)
	}
	pre_append = pre + ".0." + "block_options"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-options"], _ = expandObjectVoipProfileSipBlockOptions(d, i["block_options"], pre_append)
	}
	pre_append = pre + ".0." + "block_prack"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-prack"], _ = expandObjectVoipProfileSipBlockPrack(d, i["block_prack"], pre_append)
	}
	pre_append = pre + ".0." + "block_publish"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-publish"], _ = expandObjectVoipProfileSipBlockPublish(d, i["block_publish"], pre_append)
	}
	pre_append = pre + ".0." + "block_refer"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-refer"], _ = expandObjectVoipProfileSipBlockRefer(d, i["block_refer"], pre_append)
	}
	pre_append = pre + ".0." + "block_register"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-register"], _ = expandObjectVoipProfileSipBlockRegister(d, i["block_register"], pre_append)
	}
	pre_append = pre + ".0." + "block_subscribe"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-subscribe"], _ = expandObjectVoipProfileSipBlockSubscribe(d, i["block_subscribe"], pre_append)
	}
	pre_append = pre + ".0." + "block_unknown"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-unknown"], _ = expandObjectVoipProfileSipBlockUnknown(d, i["block_unknown"], pre_append)
	}
	pre_append = pre + ".0." + "block_update"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-update"], _ = expandObjectVoipProfileSipBlockUpdate(d, i["block_update"], pre_append)
	}
	pre_append = pre + ".0." + "bye_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["bye-rate"], _ = expandObjectVoipProfileSipByeRate(d, i["bye_rate"], pre_append)
	}
	pre_append = pre + ".0." + "bye_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["bye-rate-track"], _ = expandObjectVoipProfileSipByeRateTrack(d, i["bye_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "call_keepalive"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-keepalive"], _ = expandObjectVoipProfileSipCallKeepalive(d, i["call_keepalive"], pre_append)
	}
	pre_append = pre + ".0." + "cancel_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["cancel-rate"], _ = expandObjectVoipProfileSipCancelRate(d, i["cancel_rate"], pre_append)
	}
	pre_append = pre + ".0." + "cancel_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["cancel-rate-track"], _ = expandObjectVoipProfileSipCancelRateTrack(d, i["cancel_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "contact_fixup"
	if _, ok := d.GetOk(pre_append); ok {
		result["contact-fixup"], _ = expandObjectVoipProfileSipContactFixup(d, i["contact_fixup"], pre_append)
	}
	pre_append = pre + ".0." + "hnt_restrict_source_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["hnt-restrict-source-ip"], _ = expandObjectVoipProfileSipHntRestrictSourceIp(d, i["hnt_restrict_source_ip"], pre_append)
	}
	pre_append = pre + ".0." + "hosted_nat_traversal"
	if _, ok := d.GetOk(pre_append); ok {
		result["hosted-nat-traversal"], _ = expandObjectVoipProfileSipHostedNatTraversal(d, i["hosted_nat_traversal"], pre_append)
	}
	pre_append = pre + ".0." + "info_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["info-rate"], _ = expandObjectVoipProfileSipInfoRate(d, i["info_rate"], pre_append)
	}
	pre_append = pre + ".0." + "info_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["info-rate-track"], _ = expandObjectVoipProfileSipInfoRateTrack(d, i["info_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "invite_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["invite-rate"], _ = expandObjectVoipProfileSipInviteRate(d, i["invite_rate"], pre_append)
	}
	pre_append = pre + ".0." + "invite_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["invite-rate-track"], _ = expandObjectVoipProfileSipInviteRateTrack(d, i["invite_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "ips_rtp"
	if _, ok := d.GetOk(pre_append); ok {
		result["ips-rtp"], _ = expandObjectVoipProfileSipIpsRtp(d, i["ips_rtp"], pre_append)
	}
	pre_append = pre + ".0." + "log_call_summary"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-call-summary"], _ = expandObjectVoipProfileSipLogCallSummary(d, i["log_call_summary"], pre_append)
	}
	pre_append = pre + ".0." + "log_violations"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-violations"], _ = expandObjectVoipProfileSipLogViolations(d, i["log_violations"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_allow"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-allow"], _ = expandObjectVoipProfileSipMalformedHeaderAllow(d, i["malformed_header_allow"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_call_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-call-id"], _ = expandObjectVoipProfileSipMalformedHeaderCallId(d, i["malformed_header_call_id"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_contact"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-contact"], _ = expandObjectVoipProfileSipMalformedHeaderContact(d, i["malformed_header_contact"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_content_length"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-content-length"], _ = expandObjectVoipProfileSipMalformedHeaderContentLength(d, i["malformed_header_content_length"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_content_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-content-type"], _ = expandObjectVoipProfileSipMalformedHeaderContentType(d, i["malformed_header_content_type"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_cseq"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-cseq"], _ = expandObjectVoipProfileSipMalformedHeaderCseq(d, i["malformed_header_cseq"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_expires"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-expires"], _ = expandObjectVoipProfileSipMalformedHeaderExpires(d, i["malformed_header_expires"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_from"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-from"], _ = expandObjectVoipProfileSipMalformedHeaderFrom(d, i["malformed_header_from"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_max_forwards"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-max-forwards"], _ = expandObjectVoipProfileSipMalformedHeaderMaxForwards(d, i["malformed_header_max_forwards"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_no_proxy_require"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-no-proxy-require"], _ = expandObjectVoipProfileSipMalformedHeaderNoProxyRequire(d, i["malformed_header_no_proxy_require"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_no_require"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-no-require"], _ = expandObjectVoipProfileSipMalformedHeaderNoRequire(d, i["malformed_header_no_require"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_p_asserted_identity"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-p-asserted-identity"], _ = expandObjectVoipProfileSipMalformedHeaderPAssertedIdentity(d, i["malformed_header_p_asserted_identity"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_rack"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-rack"], _ = expandObjectVoipProfileSipMalformedHeaderRack(d, i["malformed_header_rack"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_record_route"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-record-route"], _ = expandObjectVoipProfileSipMalformedHeaderRecordRoute(d, i["malformed_header_record_route"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_route"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-route"], _ = expandObjectVoipProfileSipMalformedHeaderRoute(d, i["malformed_header_route"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_rseq"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-rseq"], _ = expandObjectVoipProfileSipMalformedHeaderRseq(d, i["malformed_header_rseq"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_a"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-a"], _ = expandObjectVoipProfileSipMalformedHeaderSdpA(d, i["malformed_header_sdp_a"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_b"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-b"], _ = expandObjectVoipProfileSipMalformedHeaderSdpB(d, i["malformed_header_sdp_b"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_c"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-c"], _ = expandObjectVoipProfileSipMalformedHeaderSdpC(d, i["malformed_header_sdp_c"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_i"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-i"], _ = expandObjectVoipProfileSipMalformedHeaderSdpI(d, i["malformed_header_sdp_i"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_k"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-k"], _ = expandObjectVoipProfileSipMalformedHeaderSdpK(d, i["malformed_header_sdp_k"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_m"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-m"], _ = expandObjectVoipProfileSipMalformedHeaderSdpM(d, i["malformed_header_sdp_m"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_o"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-o"], _ = expandObjectVoipProfileSipMalformedHeaderSdpO(d, i["malformed_header_sdp_o"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_r"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-r"], _ = expandObjectVoipProfileSipMalformedHeaderSdpR(d, i["malformed_header_sdp_r"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_s"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-s"], _ = expandObjectVoipProfileSipMalformedHeaderSdpS(d, i["malformed_header_sdp_s"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_t"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-t"], _ = expandObjectVoipProfileSipMalformedHeaderSdpT(d, i["malformed_header_sdp_t"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-v"], _ = expandObjectVoipProfileSipMalformedHeaderSdpV(d, i["malformed_header_sdp_v"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_z"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-sdp-z"], _ = expandObjectVoipProfileSipMalformedHeaderSdpZ(d, i["malformed_header_sdp_z"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_to"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-to"], _ = expandObjectVoipProfileSipMalformedHeaderTo(d, i["malformed_header_to"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_header_via"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-header-via"], _ = expandObjectVoipProfileSipMalformedHeaderVia(d, i["malformed_header_via"], pre_append)
	}
	pre_append = pre + ".0." + "malformed_request_line"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed-request-line"], _ = expandObjectVoipProfileSipMalformedRequestLine(d, i["malformed_request_line"], pre_append)
	}
	pre_append = pre + ".0." + "max_body_length"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-body-length"], _ = expandObjectVoipProfileSipMaxBodyLength(d, i["max_body_length"], pre_append)
	}
	pre_append = pre + ".0." + "max_dialogs"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-dialogs"], _ = expandObjectVoipProfileSipMaxDialogs(d, i["max_dialogs"], pre_append)
	}
	pre_append = pre + ".0." + "max_idle_dialogs"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-idle-dialogs"], _ = expandObjectVoipProfileSipMaxIdleDialogs(d, i["max_idle_dialogs"], pre_append)
	}
	pre_append = pre + ".0." + "max_line_length"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-line-length"], _ = expandObjectVoipProfileSipMaxLineLength(d, i["max_line_length"], pre_append)
	}
	pre_append = pre + ".0." + "message_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["message-rate"], _ = expandObjectVoipProfileSipMessageRate(d, i["message_rate"], pre_append)
	}
	pre_append = pre + ".0." + "message_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["message-rate-track"], _ = expandObjectVoipProfileSipMessageRateTrack(d, i["message_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "nat_port_range"
	if _, ok := d.GetOk(pre_append); ok {
		result["nat-port-range"], _ = expandObjectVoipProfileSipNatPortRange(d, i["nat_port_range"], pre_append)
	}
	pre_append = pre + ".0." + "nat_trace"
	if _, ok := d.GetOk(pre_append); ok {
		result["nat-trace"], _ = expandObjectVoipProfileSipNatTrace(d, i["nat_trace"], pre_append)
	}
	pre_append = pre + ".0." + "no_sdp_fixup"
	if _, ok := d.GetOk(pre_append); ok {
		result["no-sdp-fixup"], _ = expandObjectVoipProfileSipNoSdpFixup(d, i["no_sdp_fixup"], pre_append)
	}
	pre_append = pre + ".0." + "notify_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["notify-rate"], _ = expandObjectVoipProfileSipNotifyRate(d, i["notify_rate"], pre_append)
	}
	pre_append = pre + ".0." + "notify_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["notify-rate-track"], _ = expandObjectVoipProfileSipNotifyRateTrack(d, i["notify_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "open_contact_pinhole"
	if _, ok := d.GetOk(pre_append); ok {
		result["open-contact-pinhole"], _ = expandObjectVoipProfileSipOpenContactPinhole(d, i["open_contact_pinhole"], pre_append)
	}
	pre_append = pre + ".0." + "open_record_route_pinhole"
	if _, ok := d.GetOk(pre_append); ok {
		result["open-record-route-pinhole"], _ = expandObjectVoipProfileSipOpenRecordRoutePinhole(d, i["open_record_route_pinhole"], pre_append)
	}
	pre_append = pre + ".0." + "open_register_pinhole"
	if _, ok := d.GetOk(pre_append); ok {
		result["open-register-pinhole"], _ = expandObjectVoipProfileSipOpenRegisterPinhole(d, i["open_register_pinhole"], pre_append)
	}
	pre_append = pre + ".0." + "open_via_pinhole"
	if _, ok := d.GetOk(pre_append); ok {
		result["open-via-pinhole"], _ = expandObjectVoipProfileSipOpenViaPinhole(d, i["open_via_pinhole"], pre_append)
	}
	pre_append = pre + ".0." + "options_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["options-rate"], _ = expandObjectVoipProfileSipOptionsRate(d, i["options_rate"], pre_append)
	}
	pre_append = pre + ".0." + "options_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["options-rate-track"], _ = expandObjectVoipProfileSipOptionsRateTrack(d, i["options_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "prack_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["prack-rate"], _ = expandObjectVoipProfileSipPrackRate(d, i["prack_rate"], pre_append)
	}
	pre_append = pre + ".0." + "prack_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["prack-rate-track"], _ = expandObjectVoipProfileSipPrackRateTrack(d, i["prack_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "preserve_override"
	if _, ok := d.GetOk(pre_append); ok {
		result["preserve-override"], _ = expandObjectVoipProfileSipPreserveOverride(d, i["preserve_override"], pre_append)
	}
	pre_append = pre + ".0." + "provisional_invite_expiry_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["provisional-invite-expiry-time"], _ = expandObjectVoipProfileSipProvisionalInviteExpiryTime(d, i["provisional_invite_expiry_time"], pre_append)
	}
	pre_append = pre + ".0." + "publish_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["publish-rate"], _ = expandObjectVoipProfileSipPublishRate(d, i["publish_rate"], pre_append)
	}
	pre_append = pre + ".0." + "publish_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["publish-rate-track"], _ = expandObjectVoipProfileSipPublishRateTrack(d, i["publish_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "refer_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["refer-rate"], _ = expandObjectVoipProfileSipReferRate(d, i["refer_rate"], pre_append)
	}
	pre_append = pre + ".0." + "refer_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["refer-rate-track"], _ = expandObjectVoipProfileSipReferRateTrack(d, i["refer_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "register_contact_trace"
	if _, ok := d.GetOk(pre_append); ok {
		result["register-contact-trace"], _ = expandObjectVoipProfileSipRegisterContactTrace(d, i["register_contact_trace"], pre_append)
	}
	pre_append = pre + ".0." + "register_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["register-rate"], _ = expandObjectVoipProfileSipRegisterRate(d, i["register_rate"], pre_append)
	}
	pre_append = pre + ".0." + "register_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["register-rate-track"], _ = expandObjectVoipProfileSipRegisterRateTrack(d, i["register_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "rfc2543_branch"
	if _, ok := d.GetOk(pre_append); ok {
		result["rfc2543-branch"], _ = expandObjectVoipProfileSipRfc2543Branch(d, i["rfc2543_branch"], pre_append)
	}
	pre_append = pre + ".0." + "rtp"
	if _, ok := d.GetOk(pre_append); ok {
		result["rtp"], _ = expandObjectVoipProfileSipRtp(d, i["rtp"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_algorithm"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-algorithm"], _ = expandObjectVoipProfileSipSslAlgorithm(d, i["ssl_algorithm"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_auth_client"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-auth-client"], _ = expandObjectVoipProfileSipSslAuthClient(d, i["ssl_auth_client"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_auth_server"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-auth-server"], _ = expandObjectVoipProfileSipSslAuthServer(d, i["ssl_auth_server"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-client-certificate"], _ = expandObjectVoipProfileSipSslClientCertificate(d, i["ssl_client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_client_renegotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-client-renegotiation"], _ = expandObjectVoipProfileSipSslClientRenegotiation(d, i["ssl_client_renegotiation"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_max_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-max-version"], _ = expandObjectVoipProfileSipSslMaxVersion(d, i["ssl_max_version"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_min_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-min-version"], _ = expandObjectVoipProfileSipSslMinVersion(d, i["ssl_min_version"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-mode"], _ = expandObjectVoipProfileSipSslMode(d, i["ssl_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_pfs"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-pfs"], _ = expandObjectVoipProfileSipSslPfs(d, i["ssl_pfs"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_send_empty_frags"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-send-empty-frags"], _ = expandObjectVoipProfileSipSslSendEmptyFrags(d, i["ssl_send_empty_frags"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_server_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-server-certificate"], _ = expandObjectVoipProfileSipSslServerCertificate(d, i["ssl_server_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectVoipProfileSipStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "strict_register"
	if _, ok := d.GetOk(pre_append); ok {
		result["strict-register"], _ = expandObjectVoipProfileSipStrictRegister(d, i["strict_register"], pre_append)
	}
	pre_append = pre + ".0." + "subscribe_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["subscribe-rate"], _ = expandObjectVoipProfileSipSubscribeRate(d, i["subscribe_rate"], pre_append)
	}
	pre_append = pre + ".0." + "subscribe_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["subscribe-rate-track"], _ = expandObjectVoipProfileSipSubscribeRateTrack(d, i["subscribe_rate_track"], pre_append)
	}
	pre_append = pre + ".0." + "unknown_header"
	if _, ok := d.GetOk(pre_append); ok {
		result["unknown-header"], _ = expandObjectVoipProfileSipUnknownHeader(d, i["unknown_header"], pre_append)
	}
	pre_append = pre + ".0." + "update_rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["update-rate"], _ = expandObjectVoipProfileSipUpdateRate(d, i["update_rate"], pre_append)
	}
	pre_append = pre + ".0." + "update_rate_track"
	if _, ok := d.GetOk(pre_append); ok {
		result["update-rate-track"], _ = expandObjectVoipProfileSipUpdateRateTrack(d, i["update_rate_track"], pre_append)
	}

	return result, nil
}

func expandObjectVoipProfileSipAckRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipAckRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockAck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockBye(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockCancel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockGeoRedOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockInvite(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockLongLines(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockNotify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockPrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockPublish(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockRefer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockRegister(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockSubscribe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockUnknown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipByeRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipByeRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipCallKeepalive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipCancelRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipCancelRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipContactFixup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipHntRestrictSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipHostedNatTraversal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipInfoRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipInfoRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipInviteRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipInviteRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipIpsRtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipLogCallSummary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipLogViolations(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderAllow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderCallId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderContact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderContentLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderContentType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderCseq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderExpires(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderMaxForwards(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderNoProxyRequire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderNoRequire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderPAssertedIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderRack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderRecordRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderRseq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpA(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpB(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpC(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpI(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpK(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpM(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpO(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpS(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpT(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpZ(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderVia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedRequestLine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMaxBodyLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMaxDialogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMaxIdleDialogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMaxLineLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMessageRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMessageRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipNatPortRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipNatTrace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipNoSdpFixup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipNotifyRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipNotifyRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOpenContactPinhole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOpenRecordRoutePinhole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOpenRegisterPinhole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOpenViaPinhole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOptionsRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOptionsRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipPrackRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipPrackRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipPreserveOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipProvisionalInviteExpiryTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipPublishRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipPublishRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipReferRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipReferRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipRegisterContactTrace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipRegisterRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipRegisterRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipRfc2543Branch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipRtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslAuthClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslAuthServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslPfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslServerCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipStrictRegister(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSubscribeRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSubscribeRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipUnknownHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipUpdateRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipUpdateRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVoipProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectVoipProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok {
		t, err := expandObjectVoipProfileFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectVoipProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("sccp"); ok {
		t, err := expandObjectVoipProfileSccp(d, v, "sccp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sccp"] = t
		}
	}

	if v, ok := d.GetOk("sip"); ok {
		t, err := expandObjectVoipProfileSip(d, v, "sip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip"] = t
		}
	}

	return &obj, nil
}
