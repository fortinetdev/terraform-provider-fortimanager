// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SIP.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVoipProfileSip() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVoipProfileSipUpdate,
		Read:   resourceObjectVoipProfileSipRead,
		Update: resourceObjectVoipProfileSipUpdate,
		Delete: resourceObjectVoipProfileSipDelete,

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
			"ack_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"bye_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_id_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"call_keepalive": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cancel_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			"content_type_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"info_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invite_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"max_dialogs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_idle_dialogs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_line_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"message_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"options_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prack_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"publish_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refer_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"ssl_auth_client": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_auth_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_client_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ssl_server_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"update_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVoipProfileSipUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["profile"] = profile

	obj, err := getObjectObjectVoipProfileSip(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVoipProfileSip resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVoipProfileSip(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVoipProfileSip resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectVoipProfileSip")

	return resourceObjectVoipProfileSipRead(d, m)
}

func resourceObjectVoipProfileSipDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["profile"] = profile

	err = c.DeleteObjectVoipProfileSip(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVoipProfileSip resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVoipProfileSipRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectVoipProfileSip(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVoipProfileSip resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVoipProfileSip(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVoipProfileSip resource from API: %v", err)
	}
	return nil
}

func flattenObjectVoipProfileSipAckRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipAckRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockAck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockBye2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockCancel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockGeoRedOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockInvite2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockLongLines2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockMessage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockNotify2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockPrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockPublish2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockRefer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockRegister2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockSubscribe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockUnknown2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipBlockUpdate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipByeRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipByeRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipCallIdRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipCallKeepalive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipCancelRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipCancelRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipContactFixup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipContentTypeRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipHntRestrictSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipHostedNatTraversal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipInfoRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipInfoRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipInviteRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipInviteRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipIpsRtp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipLogCallSummary2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipLogViolations2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderAllow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderCallId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderContact2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderContentLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderContentType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderCseq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderExpires2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderFrom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderMaxForwards2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderNoProxyRequire2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderNoRequire2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderPAssertedIdentity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderRack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderRecordRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderRseq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpA2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpB2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpC2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpI2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpK2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpM2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpO2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpR2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpS2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpT2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderSdpZ2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderTo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedHeaderVia2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMalformedRequestLine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMaxBodyLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMaxDialogs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMaxIdleDialogs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMaxLineLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMessageRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipMessageRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipNatPortRange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipNatTrace2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipNoSdpFixup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipNotifyRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipNotifyRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOpenContactPinhole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOpenRecordRoutePinhole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOpenRegisterPinhole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOpenViaPinhole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOptionsRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipOptionsRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipPrackRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipPrackRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipPreserveOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipProvisionalInviteExpiryTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipPublishRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipPublishRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipReferRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipReferRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipRegisterContactTrace2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipRegisterRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipRegisterRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipRfc2543Branch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipRtp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslAlgorithm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslAuthClient2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslAuthServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslClientRenegotiation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslMaxVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslMinVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslPfs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslSendEmptyFrags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSslServerCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipStrictRegister2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSubscribeRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipSubscribeRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipUnknownHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipUpdateRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSipUpdateRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVoipProfileSip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ack_rate", flattenObjectVoipProfileSipAckRate2edl(o["ack-rate"], d, "ack_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ack-rate"], "ObjectVoipProfileSip-AckRate"); ok {
			if err = d.Set("ack_rate", vv); err != nil {
				return fmt.Errorf("Error reading ack_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ack_rate: %v", err)
		}
	}

	if err = d.Set("ack_rate_track", flattenObjectVoipProfileSipAckRateTrack2edl(o["ack-rate-track"], d, "ack_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["ack-rate-track"], "ObjectVoipProfileSip-AckRateTrack"); ok {
			if err = d.Set("ack_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading ack_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ack_rate_track: %v", err)
		}
	}

	if err = d.Set("block_ack", flattenObjectVoipProfileSipBlockAck2edl(o["block-ack"], d, "block_ack")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-ack"], "ObjectVoipProfileSip-BlockAck"); ok {
			if err = d.Set("block_ack", vv); err != nil {
				return fmt.Errorf("Error reading block_ack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_ack: %v", err)
		}
	}

	if err = d.Set("block_bye", flattenObjectVoipProfileSipBlockBye2edl(o["block-bye"], d, "block_bye")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-bye"], "ObjectVoipProfileSip-BlockBye"); ok {
			if err = d.Set("block_bye", vv); err != nil {
				return fmt.Errorf("Error reading block_bye: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_bye: %v", err)
		}
	}

	if err = d.Set("block_cancel", flattenObjectVoipProfileSipBlockCancel2edl(o["block-cancel"], d, "block_cancel")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-cancel"], "ObjectVoipProfileSip-BlockCancel"); ok {
			if err = d.Set("block_cancel", vv); err != nil {
				return fmt.Errorf("Error reading block_cancel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_cancel: %v", err)
		}
	}

	if err = d.Set("block_geo_red_options", flattenObjectVoipProfileSipBlockGeoRedOptions2edl(o["block-geo-red-options"], d, "block_geo_red_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-geo-red-options"], "ObjectVoipProfileSip-BlockGeoRedOptions"); ok {
			if err = d.Set("block_geo_red_options", vv); err != nil {
				return fmt.Errorf("Error reading block_geo_red_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_geo_red_options: %v", err)
		}
	}

	if err = d.Set("block_info", flattenObjectVoipProfileSipBlockInfo2edl(o["block-info"], d, "block_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-info"], "ObjectVoipProfileSip-BlockInfo"); ok {
			if err = d.Set("block_info", vv); err != nil {
				return fmt.Errorf("Error reading block_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_info: %v", err)
		}
	}

	if err = d.Set("block_invite", flattenObjectVoipProfileSipBlockInvite2edl(o["block-invite"], d, "block_invite")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-invite"], "ObjectVoipProfileSip-BlockInvite"); ok {
			if err = d.Set("block_invite", vv); err != nil {
				return fmt.Errorf("Error reading block_invite: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_invite: %v", err)
		}
	}

	if err = d.Set("block_long_lines", flattenObjectVoipProfileSipBlockLongLines2edl(o["block-long-lines"], d, "block_long_lines")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-long-lines"], "ObjectVoipProfileSip-BlockLongLines"); ok {
			if err = d.Set("block_long_lines", vv); err != nil {
				return fmt.Errorf("Error reading block_long_lines: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_long_lines: %v", err)
		}
	}

	if err = d.Set("block_message", flattenObjectVoipProfileSipBlockMessage2edl(o["block-message"], d, "block_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-message"], "ObjectVoipProfileSip-BlockMessage"); ok {
			if err = d.Set("block_message", vv); err != nil {
				return fmt.Errorf("Error reading block_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_message: %v", err)
		}
	}

	if err = d.Set("block_notify", flattenObjectVoipProfileSipBlockNotify2edl(o["block-notify"], d, "block_notify")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-notify"], "ObjectVoipProfileSip-BlockNotify"); ok {
			if err = d.Set("block_notify", vv); err != nil {
				return fmt.Errorf("Error reading block_notify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_notify: %v", err)
		}
	}

	if err = d.Set("block_options", flattenObjectVoipProfileSipBlockOptions2edl(o["block-options"], d, "block_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-options"], "ObjectVoipProfileSip-BlockOptions"); ok {
			if err = d.Set("block_options", vv); err != nil {
				return fmt.Errorf("Error reading block_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_options: %v", err)
		}
	}

	if err = d.Set("block_prack", flattenObjectVoipProfileSipBlockPrack2edl(o["block-prack"], d, "block_prack")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-prack"], "ObjectVoipProfileSip-BlockPrack"); ok {
			if err = d.Set("block_prack", vv); err != nil {
				return fmt.Errorf("Error reading block_prack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_prack: %v", err)
		}
	}

	if err = d.Set("block_publish", flattenObjectVoipProfileSipBlockPublish2edl(o["block-publish"], d, "block_publish")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-publish"], "ObjectVoipProfileSip-BlockPublish"); ok {
			if err = d.Set("block_publish", vv); err != nil {
				return fmt.Errorf("Error reading block_publish: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_publish: %v", err)
		}
	}

	if err = d.Set("block_refer", flattenObjectVoipProfileSipBlockRefer2edl(o["block-refer"], d, "block_refer")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-refer"], "ObjectVoipProfileSip-BlockRefer"); ok {
			if err = d.Set("block_refer", vv); err != nil {
				return fmt.Errorf("Error reading block_refer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_refer: %v", err)
		}
	}

	if err = d.Set("block_register", flattenObjectVoipProfileSipBlockRegister2edl(o["block-register"], d, "block_register")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-register"], "ObjectVoipProfileSip-BlockRegister"); ok {
			if err = d.Set("block_register", vv); err != nil {
				return fmt.Errorf("Error reading block_register: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_register: %v", err)
		}
	}

	if err = d.Set("block_subscribe", flattenObjectVoipProfileSipBlockSubscribe2edl(o["block-subscribe"], d, "block_subscribe")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-subscribe"], "ObjectVoipProfileSip-BlockSubscribe"); ok {
			if err = d.Set("block_subscribe", vv); err != nil {
				return fmt.Errorf("Error reading block_subscribe: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_subscribe: %v", err)
		}
	}

	if err = d.Set("block_unknown", flattenObjectVoipProfileSipBlockUnknown2edl(o["block-unknown"], d, "block_unknown")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-unknown"], "ObjectVoipProfileSip-BlockUnknown"); ok {
			if err = d.Set("block_unknown", vv); err != nil {
				return fmt.Errorf("Error reading block_unknown: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_unknown: %v", err)
		}
	}

	if err = d.Set("block_update", flattenObjectVoipProfileSipBlockUpdate2edl(o["block-update"], d, "block_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-update"], "ObjectVoipProfileSip-BlockUpdate"); ok {
			if err = d.Set("block_update", vv); err != nil {
				return fmt.Errorf("Error reading block_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_update: %v", err)
		}
	}

	if err = d.Set("bye_rate", flattenObjectVoipProfileSipByeRate2edl(o["bye-rate"], d, "bye_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bye-rate"], "ObjectVoipProfileSip-ByeRate"); ok {
			if err = d.Set("bye_rate", vv); err != nil {
				return fmt.Errorf("Error reading bye_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bye_rate: %v", err)
		}
	}

	if err = d.Set("bye_rate_track", flattenObjectVoipProfileSipByeRateTrack2edl(o["bye-rate-track"], d, "bye_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["bye-rate-track"], "ObjectVoipProfileSip-ByeRateTrack"); ok {
			if err = d.Set("bye_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading bye_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bye_rate_track: %v", err)
		}
	}

	if err = d.Set("call_id_regex", flattenObjectVoipProfileSipCallIdRegex2edl(o["call-id-regex"], d, "call_id_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-id-regex"], "ObjectVoipProfileSip-CallIdRegex"); ok {
			if err = d.Set("call_id_regex", vv); err != nil {
				return fmt.Errorf("Error reading call_id_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_id_regex: %v", err)
		}
	}

	if err = d.Set("call_keepalive", flattenObjectVoipProfileSipCallKeepalive2edl(o["call-keepalive"], d, "call_keepalive")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-keepalive"], "ObjectVoipProfileSip-CallKeepalive"); ok {
			if err = d.Set("call_keepalive", vv); err != nil {
				return fmt.Errorf("Error reading call_keepalive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_keepalive: %v", err)
		}
	}

	if err = d.Set("cancel_rate", flattenObjectVoipProfileSipCancelRate2edl(o["cancel-rate"], d, "cancel_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["cancel-rate"], "ObjectVoipProfileSip-CancelRate"); ok {
			if err = d.Set("cancel_rate", vv); err != nil {
				return fmt.Errorf("Error reading cancel_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cancel_rate: %v", err)
		}
	}

	if err = d.Set("cancel_rate_track", flattenObjectVoipProfileSipCancelRateTrack2edl(o["cancel-rate-track"], d, "cancel_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["cancel-rate-track"], "ObjectVoipProfileSip-CancelRateTrack"); ok {
			if err = d.Set("cancel_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading cancel_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cancel_rate_track: %v", err)
		}
	}

	if err = d.Set("contact_fixup", flattenObjectVoipProfileSipContactFixup2edl(o["contact-fixup"], d, "contact_fixup")); err != nil {
		if vv, ok := fortiAPIPatch(o["contact-fixup"], "ObjectVoipProfileSip-ContactFixup"); ok {
			if err = d.Set("contact_fixup", vv); err != nil {
				return fmt.Errorf("Error reading contact_fixup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading contact_fixup: %v", err)
		}
	}

	if err = d.Set("content_type_regex", flattenObjectVoipProfileSipContentTypeRegex2edl(o["content-type-regex"], d, "content_type_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["content-type-regex"], "ObjectVoipProfileSip-ContentTypeRegex"); ok {
			if err = d.Set("content_type_regex", vv); err != nil {
				return fmt.Errorf("Error reading content_type_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content_type_regex: %v", err)
		}
	}

	if err = d.Set("hnt_restrict_source_ip", flattenObjectVoipProfileSipHntRestrictSourceIp2edl(o["hnt-restrict-source-ip"], d, "hnt_restrict_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["hnt-restrict-source-ip"], "ObjectVoipProfileSip-HntRestrictSourceIp"); ok {
			if err = d.Set("hnt_restrict_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading hnt_restrict_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hnt_restrict_source_ip: %v", err)
		}
	}

	if err = d.Set("hosted_nat_traversal", flattenObjectVoipProfileSipHostedNatTraversal2edl(o["hosted-nat-traversal"], d, "hosted_nat_traversal")); err != nil {
		if vv, ok := fortiAPIPatch(o["hosted-nat-traversal"], "ObjectVoipProfileSip-HostedNatTraversal"); ok {
			if err = d.Set("hosted_nat_traversal", vv); err != nil {
				return fmt.Errorf("Error reading hosted_nat_traversal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hosted_nat_traversal: %v", err)
		}
	}

	if err = d.Set("info_rate", flattenObjectVoipProfileSipInfoRate2edl(o["info-rate"], d, "info_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["info-rate"], "ObjectVoipProfileSip-InfoRate"); ok {
			if err = d.Set("info_rate", vv); err != nil {
				return fmt.Errorf("Error reading info_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading info_rate: %v", err)
		}
	}

	if err = d.Set("info_rate_track", flattenObjectVoipProfileSipInfoRateTrack2edl(o["info-rate-track"], d, "info_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["info-rate-track"], "ObjectVoipProfileSip-InfoRateTrack"); ok {
			if err = d.Set("info_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading info_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading info_rate_track: %v", err)
		}
	}

	if err = d.Set("invite_rate", flattenObjectVoipProfileSipInviteRate2edl(o["invite-rate"], d, "invite_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["invite-rate"], "ObjectVoipProfileSip-InviteRate"); ok {
			if err = d.Set("invite_rate", vv); err != nil {
				return fmt.Errorf("Error reading invite_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invite_rate: %v", err)
		}
	}

	if err = d.Set("invite_rate_track", flattenObjectVoipProfileSipInviteRateTrack2edl(o["invite-rate-track"], d, "invite_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["invite-rate-track"], "ObjectVoipProfileSip-InviteRateTrack"); ok {
			if err = d.Set("invite_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading invite_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invite_rate_track: %v", err)
		}
	}

	if err = d.Set("ips_rtp", flattenObjectVoipProfileSipIpsRtp2edl(o["ips-rtp"], d, "ips_rtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-rtp"], "ObjectVoipProfileSip-IpsRtp"); ok {
			if err = d.Set("ips_rtp", vv); err != nil {
				return fmt.Errorf("Error reading ips_rtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_rtp: %v", err)
		}
	}

	if err = d.Set("log_call_summary", flattenObjectVoipProfileSipLogCallSummary2edl(o["log-call-summary"], d, "log_call_summary")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-call-summary"], "ObjectVoipProfileSip-LogCallSummary"); ok {
			if err = d.Set("log_call_summary", vv); err != nil {
				return fmt.Errorf("Error reading log_call_summary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_call_summary: %v", err)
		}
	}

	if err = d.Set("log_violations", flattenObjectVoipProfileSipLogViolations2edl(o["log-violations"], d, "log_violations")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-violations"], "ObjectVoipProfileSip-LogViolations"); ok {
			if err = d.Set("log_violations", vv); err != nil {
				return fmt.Errorf("Error reading log_violations: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_violations: %v", err)
		}
	}

	if err = d.Set("malformed_header_allow", flattenObjectVoipProfileSipMalformedHeaderAllow2edl(o["malformed-header-allow"], d, "malformed_header_allow")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-allow"], "ObjectVoipProfileSip-MalformedHeaderAllow"); ok {
			if err = d.Set("malformed_header_allow", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_allow: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_allow: %v", err)
		}
	}

	if err = d.Set("malformed_header_call_id", flattenObjectVoipProfileSipMalformedHeaderCallId2edl(o["malformed-header-call-id"], d, "malformed_header_call_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-call-id"], "ObjectVoipProfileSip-MalformedHeaderCallId"); ok {
			if err = d.Set("malformed_header_call_id", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_call_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_call_id: %v", err)
		}
	}

	if err = d.Set("malformed_header_contact", flattenObjectVoipProfileSipMalformedHeaderContact2edl(o["malformed-header-contact"], d, "malformed_header_contact")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-contact"], "ObjectVoipProfileSip-MalformedHeaderContact"); ok {
			if err = d.Set("malformed_header_contact", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_contact: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_contact: %v", err)
		}
	}

	if err = d.Set("malformed_header_content_length", flattenObjectVoipProfileSipMalformedHeaderContentLength2edl(o["malformed-header-content-length"], d, "malformed_header_content_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-content-length"], "ObjectVoipProfileSip-MalformedHeaderContentLength"); ok {
			if err = d.Set("malformed_header_content_length", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_content_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_content_length: %v", err)
		}
	}

	if err = d.Set("malformed_header_content_type", flattenObjectVoipProfileSipMalformedHeaderContentType2edl(o["malformed-header-content-type"], d, "malformed_header_content_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-content-type"], "ObjectVoipProfileSip-MalformedHeaderContentType"); ok {
			if err = d.Set("malformed_header_content_type", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_content_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_content_type: %v", err)
		}
	}

	if err = d.Set("malformed_header_cseq", flattenObjectVoipProfileSipMalformedHeaderCseq2edl(o["malformed-header-cseq"], d, "malformed_header_cseq")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-cseq"], "ObjectVoipProfileSip-MalformedHeaderCseq"); ok {
			if err = d.Set("malformed_header_cseq", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_cseq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_cseq: %v", err)
		}
	}

	if err = d.Set("malformed_header_expires", flattenObjectVoipProfileSipMalformedHeaderExpires2edl(o["malformed-header-expires"], d, "malformed_header_expires")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-expires"], "ObjectVoipProfileSip-MalformedHeaderExpires"); ok {
			if err = d.Set("malformed_header_expires", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_expires: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_expires: %v", err)
		}
	}

	if err = d.Set("malformed_header_from", flattenObjectVoipProfileSipMalformedHeaderFrom2edl(o["malformed-header-from"], d, "malformed_header_from")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-from"], "ObjectVoipProfileSip-MalformedHeaderFrom"); ok {
			if err = d.Set("malformed_header_from", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_from: %v", err)
		}
	}

	if err = d.Set("malformed_header_max_forwards", flattenObjectVoipProfileSipMalformedHeaderMaxForwards2edl(o["malformed-header-max-forwards"], d, "malformed_header_max_forwards")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-max-forwards"], "ObjectVoipProfileSip-MalformedHeaderMaxForwards"); ok {
			if err = d.Set("malformed_header_max_forwards", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_max_forwards: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_max_forwards: %v", err)
		}
	}

	if err = d.Set("malformed_header_no_proxy_require", flattenObjectVoipProfileSipMalformedHeaderNoProxyRequire2edl(o["malformed-header-no-proxy-require"], d, "malformed_header_no_proxy_require")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-no-proxy-require"], "ObjectVoipProfileSip-MalformedHeaderNoProxyRequire"); ok {
			if err = d.Set("malformed_header_no_proxy_require", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_no_proxy_require: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_no_proxy_require: %v", err)
		}
	}

	if err = d.Set("malformed_header_no_require", flattenObjectVoipProfileSipMalformedHeaderNoRequire2edl(o["malformed-header-no-require"], d, "malformed_header_no_require")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-no-require"], "ObjectVoipProfileSip-MalformedHeaderNoRequire"); ok {
			if err = d.Set("malformed_header_no_require", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_no_require: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_no_require: %v", err)
		}
	}

	if err = d.Set("malformed_header_p_asserted_identity", flattenObjectVoipProfileSipMalformedHeaderPAssertedIdentity2edl(o["malformed-header-p-asserted-identity"], d, "malformed_header_p_asserted_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-p-asserted-identity"], "ObjectVoipProfileSip-MalformedHeaderPAssertedIdentity"); ok {
			if err = d.Set("malformed_header_p_asserted_identity", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_p_asserted_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_p_asserted_identity: %v", err)
		}
	}

	if err = d.Set("malformed_header_rack", flattenObjectVoipProfileSipMalformedHeaderRack2edl(o["malformed-header-rack"], d, "malformed_header_rack")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-rack"], "ObjectVoipProfileSip-MalformedHeaderRack"); ok {
			if err = d.Set("malformed_header_rack", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_rack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_rack: %v", err)
		}
	}

	if err = d.Set("malformed_header_record_route", flattenObjectVoipProfileSipMalformedHeaderRecordRoute2edl(o["malformed-header-record-route"], d, "malformed_header_record_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-record-route"], "ObjectVoipProfileSip-MalformedHeaderRecordRoute"); ok {
			if err = d.Set("malformed_header_record_route", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_record_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_record_route: %v", err)
		}
	}

	if err = d.Set("malformed_header_route", flattenObjectVoipProfileSipMalformedHeaderRoute2edl(o["malformed-header-route"], d, "malformed_header_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-route"], "ObjectVoipProfileSip-MalformedHeaderRoute"); ok {
			if err = d.Set("malformed_header_route", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_route: %v", err)
		}
	}

	if err = d.Set("malformed_header_rseq", flattenObjectVoipProfileSipMalformedHeaderRseq2edl(o["malformed-header-rseq"], d, "malformed_header_rseq")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-rseq"], "ObjectVoipProfileSip-MalformedHeaderRseq"); ok {
			if err = d.Set("malformed_header_rseq", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_rseq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_rseq: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_a", flattenObjectVoipProfileSipMalformedHeaderSdpA2edl(o["malformed-header-sdp-a"], d, "malformed_header_sdp_a")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-a"], "ObjectVoipProfileSip-MalformedHeaderSdpA"); ok {
			if err = d.Set("malformed_header_sdp_a", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_a: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_a: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_b", flattenObjectVoipProfileSipMalformedHeaderSdpB2edl(o["malformed-header-sdp-b"], d, "malformed_header_sdp_b")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-b"], "ObjectVoipProfileSip-MalformedHeaderSdpB"); ok {
			if err = d.Set("malformed_header_sdp_b", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_b: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_b: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_c", flattenObjectVoipProfileSipMalformedHeaderSdpC2edl(o["malformed-header-sdp-c"], d, "malformed_header_sdp_c")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-c"], "ObjectVoipProfileSip-MalformedHeaderSdpC"); ok {
			if err = d.Set("malformed_header_sdp_c", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_c: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_c: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_i", flattenObjectVoipProfileSipMalformedHeaderSdpI2edl(o["malformed-header-sdp-i"], d, "malformed_header_sdp_i")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-i"], "ObjectVoipProfileSip-MalformedHeaderSdpI"); ok {
			if err = d.Set("malformed_header_sdp_i", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_i: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_i: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_k", flattenObjectVoipProfileSipMalformedHeaderSdpK2edl(o["malformed-header-sdp-k"], d, "malformed_header_sdp_k")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-k"], "ObjectVoipProfileSip-MalformedHeaderSdpK"); ok {
			if err = d.Set("malformed_header_sdp_k", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_k: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_k: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_m", flattenObjectVoipProfileSipMalformedHeaderSdpM2edl(o["malformed-header-sdp-m"], d, "malformed_header_sdp_m")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-m"], "ObjectVoipProfileSip-MalformedHeaderSdpM"); ok {
			if err = d.Set("malformed_header_sdp_m", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_m: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_m: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_o", flattenObjectVoipProfileSipMalformedHeaderSdpO2edl(o["malformed-header-sdp-o"], d, "malformed_header_sdp_o")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-o"], "ObjectVoipProfileSip-MalformedHeaderSdpO"); ok {
			if err = d.Set("malformed_header_sdp_o", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_o: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_o: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_r", flattenObjectVoipProfileSipMalformedHeaderSdpR2edl(o["malformed-header-sdp-r"], d, "malformed_header_sdp_r")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-r"], "ObjectVoipProfileSip-MalformedHeaderSdpR"); ok {
			if err = d.Set("malformed_header_sdp_r", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_r: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_r: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_s", flattenObjectVoipProfileSipMalformedHeaderSdpS2edl(o["malformed-header-sdp-s"], d, "malformed_header_sdp_s")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-s"], "ObjectVoipProfileSip-MalformedHeaderSdpS"); ok {
			if err = d.Set("malformed_header_sdp_s", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_s: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_s: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_t", flattenObjectVoipProfileSipMalformedHeaderSdpT2edl(o["malformed-header-sdp-t"], d, "malformed_header_sdp_t")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-t"], "ObjectVoipProfileSip-MalformedHeaderSdpT"); ok {
			if err = d.Set("malformed_header_sdp_t", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_t: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_t: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_v", flattenObjectVoipProfileSipMalformedHeaderSdpV2edl(o["malformed-header-sdp-v"], d, "malformed_header_sdp_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-v"], "ObjectVoipProfileSip-MalformedHeaderSdpV"); ok {
			if err = d.Set("malformed_header_sdp_v", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_v: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_z", flattenObjectVoipProfileSipMalformedHeaderSdpZ2edl(o["malformed-header-sdp-z"], d, "malformed_header_sdp_z")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-z"], "ObjectVoipProfileSip-MalformedHeaderSdpZ"); ok {
			if err = d.Set("malformed_header_sdp_z", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_z: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_z: %v", err)
		}
	}

	if err = d.Set("malformed_header_to", flattenObjectVoipProfileSipMalformedHeaderTo2edl(o["malformed-header-to"], d, "malformed_header_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-to"], "ObjectVoipProfileSip-MalformedHeaderTo"); ok {
			if err = d.Set("malformed_header_to", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_to: %v", err)
		}
	}

	if err = d.Set("malformed_header_via", flattenObjectVoipProfileSipMalformedHeaderVia2edl(o["malformed-header-via"], d, "malformed_header_via")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-via"], "ObjectVoipProfileSip-MalformedHeaderVia"); ok {
			if err = d.Set("malformed_header_via", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_via: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_via: %v", err)
		}
	}

	if err = d.Set("malformed_request_line", flattenObjectVoipProfileSipMalformedRequestLine2edl(o["malformed-request-line"], d, "malformed_request_line")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-request-line"], "ObjectVoipProfileSip-MalformedRequestLine"); ok {
			if err = d.Set("malformed_request_line", vv); err != nil {
				return fmt.Errorf("Error reading malformed_request_line: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_request_line: %v", err)
		}
	}

	if err = d.Set("max_body_length", flattenObjectVoipProfileSipMaxBodyLength2edl(o["max-body-length"], d, "max_body_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-body-length"], "ObjectVoipProfileSip-MaxBodyLength"); ok {
			if err = d.Set("max_body_length", vv); err != nil {
				return fmt.Errorf("Error reading max_body_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_body_length: %v", err)
		}
	}

	if err = d.Set("max_dialogs", flattenObjectVoipProfileSipMaxDialogs2edl(o["max-dialogs"], d, "max_dialogs")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-dialogs"], "ObjectVoipProfileSip-MaxDialogs"); ok {
			if err = d.Set("max_dialogs", vv); err != nil {
				return fmt.Errorf("Error reading max_dialogs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_dialogs: %v", err)
		}
	}

	if err = d.Set("max_idle_dialogs", flattenObjectVoipProfileSipMaxIdleDialogs2edl(o["max-idle-dialogs"], d, "max_idle_dialogs")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-idle-dialogs"], "ObjectVoipProfileSip-MaxIdleDialogs"); ok {
			if err = d.Set("max_idle_dialogs", vv); err != nil {
				return fmt.Errorf("Error reading max_idle_dialogs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_idle_dialogs: %v", err)
		}
	}

	if err = d.Set("max_line_length", flattenObjectVoipProfileSipMaxLineLength2edl(o["max-line-length"], d, "max_line_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-line-length"], "ObjectVoipProfileSip-MaxLineLength"); ok {
			if err = d.Set("max_line_length", vv); err != nil {
				return fmt.Errorf("Error reading max_line_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_line_length: %v", err)
		}
	}

	if err = d.Set("message_rate", flattenObjectVoipProfileSipMessageRate2edl(o["message-rate"], d, "message_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-rate"], "ObjectVoipProfileSip-MessageRate"); ok {
			if err = d.Set("message_rate", vv); err != nil {
				return fmt.Errorf("Error reading message_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_rate: %v", err)
		}
	}

	if err = d.Set("message_rate_track", flattenObjectVoipProfileSipMessageRateTrack2edl(o["message-rate-track"], d, "message_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-rate-track"], "ObjectVoipProfileSip-MessageRateTrack"); ok {
			if err = d.Set("message_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading message_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_rate_track: %v", err)
		}
	}

	if err = d.Set("nat_port_range", flattenObjectVoipProfileSipNatPortRange2edl(o["nat-port-range"], d, "nat_port_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-port-range"], "ObjectVoipProfileSip-NatPortRange"); ok {
			if err = d.Set("nat_port_range", vv); err != nil {
				return fmt.Errorf("Error reading nat_port_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_port_range: %v", err)
		}
	}

	if err = d.Set("nat_trace", flattenObjectVoipProfileSipNatTrace2edl(o["nat-trace"], d, "nat_trace")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-trace"], "ObjectVoipProfileSip-NatTrace"); ok {
			if err = d.Set("nat_trace", vv); err != nil {
				return fmt.Errorf("Error reading nat_trace: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_trace: %v", err)
		}
	}

	if err = d.Set("no_sdp_fixup", flattenObjectVoipProfileSipNoSdpFixup2edl(o["no-sdp-fixup"], d, "no_sdp_fixup")); err != nil {
		if vv, ok := fortiAPIPatch(o["no-sdp-fixup"], "ObjectVoipProfileSip-NoSdpFixup"); ok {
			if err = d.Set("no_sdp_fixup", vv); err != nil {
				return fmt.Errorf("Error reading no_sdp_fixup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading no_sdp_fixup: %v", err)
		}
	}

	if err = d.Set("notify_rate", flattenObjectVoipProfileSipNotifyRate2edl(o["notify-rate"], d, "notify_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify-rate"], "ObjectVoipProfileSip-NotifyRate"); ok {
			if err = d.Set("notify_rate", vv); err != nil {
				return fmt.Errorf("Error reading notify_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify_rate: %v", err)
		}
	}

	if err = d.Set("notify_rate_track", flattenObjectVoipProfileSipNotifyRateTrack2edl(o["notify-rate-track"], d, "notify_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify-rate-track"], "ObjectVoipProfileSip-NotifyRateTrack"); ok {
			if err = d.Set("notify_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading notify_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify_rate_track: %v", err)
		}
	}

	if err = d.Set("open_contact_pinhole", flattenObjectVoipProfileSipOpenContactPinhole2edl(o["open-contact-pinhole"], d, "open_contact_pinhole")); err != nil {
		if vv, ok := fortiAPIPatch(o["open-contact-pinhole"], "ObjectVoipProfileSip-OpenContactPinhole"); ok {
			if err = d.Set("open_contact_pinhole", vv); err != nil {
				return fmt.Errorf("Error reading open_contact_pinhole: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading open_contact_pinhole: %v", err)
		}
	}

	if err = d.Set("open_record_route_pinhole", flattenObjectVoipProfileSipOpenRecordRoutePinhole2edl(o["open-record-route-pinhole"], d, "open_record_route_pinhole")); err != nil {
		if vv, ok := fortiAPIPatch(o["open-record-route-pinhole"], "ObjectVoipProfileSip-OpenRecordRoutePinhole"); ok {
			if err = d.Set("open_record_route_pinhole", vv); err != nil {
				return fmt.Errorf("Error reading open_record_route_pinhole: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading open_record_route_pinhole: %v", err)
		}
	}

	if err = d.Set("open_register_pinhole", flattenObjectVoipProfileSipOpenRegisterPinhole2edl(o["open-register-pinhole"], d, "open_register_pinhole")); err != nil {
		if vv, ok := fortiAPIPatch(o["open-register-pinhole"], "ObjectVoipProfileSip-OpenRegisterPinhole"); ok {
			if err = d.Set("open_register_pinhole", vv); err != nil {
				return fmt.Errorf("Error reading open_register_pinhole: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading open_register_pinhole: %v", err)
		}
	}

	if err = d.Set("open_via_pinhole", flattenObjectVoipProfileSipOpenViaPinhole2edl(o["open-via-pinhole"], d, "open_via_pinhole")); err != nil {
		if vv, ok := fortiAPIPatch(o["open-via-pinhole"], "ObjectVoipProfileSip-OpenViaPinhole"); ok {
			if err = d.Set("open_via_pinhole", vv); err != nil {
				return fmt.Errorf("Error reading open_via_pinhole: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading open_via_pinhole: %v", err)
		}
	}

	if err = d.Set("options_rate", flattenObjectVoipProfileSipOptionsRate2edl(o["options-rate"], d, "options_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["options-rate"], "ObjectVoipProfileSip-OptionsRate"); ok {
			if err = d.Set("options_rate", vv); err != nil {
				return fmt.Errorf("Error reading options_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options_rate: %v", err)
		}
	}

	if err = d.Set("options_rate_track", flattenObjectVoipProfileSipOptionsRateTrack2edl(o["options-rate-track"], d, "options_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["options-rate-track"], "ObjectVoipProfileSip-OptionsRateTrack"); ok {
			if err = d.Set("options_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading options_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options_rate_track: %v", err)
		}
	}

	if err = d.Set("prack_rate", flattenObjectVoipProfileSipPrackRate2edl(o["prack-rate"], d, "prack_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["prack-rate"], "ObjectVoipProfileSip-PrackRate"); ok {
			if err = d.Set("prack_rate", vv); err != nil {
				return fmt.Errorf("Error reading prack_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prack_rate: %v", err)
		}
	}

	if err = d.Set("prack_rate_track", flattenObjectVoipProfileSipPrackRateTrack2edl(o["prack-rate-track"], d, "prack_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["prack-rate-track"], "ObjectVoipProfileSip-PrackRateTrack"); ok {
			if err = d.Set("prack_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading prack_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prack_rate_track: %v", err)
		}
	}

	if err = d.Set("preserve_override", flattenObjectVoipProfileSipPreserveOverride2edl(o["preserve-override"], d, "preserve_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["preserve-override"], "ObjectVoipProfileSip-PreserveOverride"); ok {
			if err = d.Set("preserve_override", vv); err != nil {
				return fmt.Errorf("Error reading preserve_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preserve_override: %v", err)
		}
	}

	if err = d.Set("provisional_invite_expiry_time", flattenObjectVoipProfileSipProvisionalInviteExpiryTime2edl(o["provisional-invite-expiry-time"], d, "provisional_invite_expiry_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["provisional-invite-expiry-time"], "ObjectVoipProfileSip-ProvisionalInviteExpiryTime"); ok {
			if err = d.Set("provisional_invite_expiry_time", vv); err != nil {
				return fmt.Errorf("Error reading provisional_invite_expiry_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading provisional_invite_expiry_time: %v", err)
		}
	}

	if err = d.Set("publish_rate", flattenObjectVoipProfileSipPublishRate2edl(o["publish-rate"], d, "publish_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["publish-rate"], "ObjectVoipProfileSip-PublishRate"); ok {
			if err = d.Set("publish_rate", vv); err != nil {
				return fmt.Errorf("Error reading publish_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading publish_rate: %v", err)
		}
	}

	if err = d.Set("publish_rate_track", flattenObjectVoipProfileSipPublishRateTrack2edl(o["publish-rate-track"], d, "publish_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["publish-rate-track"], "ObjectVoipProfileSip-PublishRateTrack"); ok {
			if err = d.Set("publish_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading publish_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading publish_rate_track: %v", err)
		}
	}

	if err = d.Set("refer_rate", flattenObjectVoipProfileSipReferRate2edl(o["refer-rate"], d, "refer_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["refer-rate"], "ObjectVoipProfileSip-ReferRate"); ok {
			if err = d.Set("refer_rate", vv); err != nil {
				return fmt.Errorf("Error reading refer_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading refer_rate: %v", err)
		}
	}

	if err = d.Set("refer_rate_track", flattenObjectVoipProfileSipReferRateTrack2edl(o["refer-rate-track"], d, "refer_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["refer-rate-track"], "ObjectVoipProfileSip-ReferRateTrack"); ok {
			if err = d.Set("refer_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading refer_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading refer_rate_track: %v", err)
		}
	}

	if err = d.Set("register_contact_trace", flattenObjectVoipProfileSipRegisterContactTrace2edl(o["register-contact-trace"], d, "register_contact_trace")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-contact-trace"], "ObjectVoipProfileSip-RegisterContactTrace"); ok {
			if err = d.Set("register_contact_trace", vv); err != nil {
				return fmt.Errorf("Error reading register_contact_trace: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_contact_trace: %v", err)
		}
	}

	if err = d.Set("register_rate", flattenObjectVoipProfileSipRegisterRate2edl(o["register-rate"], d, "register_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-rate"], "ObjectVoipProfileSip-RegisterRate"); ok {
			if err = d.Set("register_rate", vv); err != nil {
				return fmt.Errorf("Error reading register_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_rate: %v", err)
		}
	}

	if err = d.Set("register_rate_track", flattenObjectVoipProfileSipRegisterRateTrack2edl(o["register-rate-track"], d, "register_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-rate-track"], "ObjectVoipProfileSip-RegisterRateTrack"); ok {
			if err = d.Set("register_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading register_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_rate_track: %v", err)
		}
	}

	if err = d.Set("rfc2543_branch", flattenObjectVoipProfileSipRfc2543Branch2edl(o["rfc2543-branch"], d, "rfc2543_branch")); err != nil {
		if vv, ok := fortiAPIPatch(o["rfc2543-branch"], "ObjectVoipProfileSip-Rfc2543Branch"); ok {
			if err = d.Set("rfc2543_branch", vv); err != nil {
				return fmt.Errorf("Error reading rfc2543_branch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rfc2543_branch: %v", err)
		}
	}

	if err = d.Set("rtp", flattenObjectVoipProfileSipRtp2edl(o["rtp"], d, "rtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtp"], "ObjectVoipProfileSip-Rtp"); ok {
			if err = d.Set("rtp", vv); err != nil {
				return fmt.Errorf("Error reading rtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtp: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenObjectVoipProfileSipSslAlgorithm2edl(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "ObjectVoipProfileSip-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_auth_client", flattenObjectVoipProfileSipSslAuthClient2edl(o["ssl-auth-client"], d, "ssl_auth_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-auth-client"], "ObjectVoipProfileSip-SslAuthClient"); ok {
			if err = d.Set("ssl_auth_client", vv); err != nil {
				return fmt.Errorf("Error reading ssl_auth_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_auth_client: %v", err)
		}
	}

	if err = d.Set("ssl_auth_server", flattenObjectVoipProfileSipSslAuthServer2edl(o["ssl-auth-server"], d, "ssl_auth_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-auth-server"], "ObjectVoipProfileSip-SslAuthServer"); ok {
			if err = d.Set("ssl_auth_server", vv); err != nil {
				return fmt.Errorf("Error reading ssl_auth_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_auth_server: %v", err)
		}
	}

	if err = d.Set("ssl_client_certificate", flattenObjectVoipProfileSipSslClientCertificate2edl(o["ssl-client-certificate"], d, "ssl_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-certificate"], "ObjectVoipProfileSip-SslClientCertificate"); ok {
			if err = d.Set("ssl_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_certificate: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenObjectVoipProfileSipSslClientRenegotiation2edl(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-renegotiation"], "ObjectVoipProfileSip-SslClientRenegotiation"); ok {
			if err = d.Set("ssl_client_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenObjectVoipProfileSipSslMaxVersion2edl(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "ObjectVoipProfileSip-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenObjectVoipProfileSipSslMinVersion2edl(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "ObjectVoipProfileSip-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_mode", flattenObjectVoipProfileSipSslMode2edl(o["ssl-mode"], d, "ssl_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mode"], "ObjectVoipProfileSip-SslMode"); ok {
			if err = d.Set("ssl_mode", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if err = d.Set("ssl_pfs", flattenObjectVoipProfileSipSslPfs2edl(o["ssl-pfs"], d, "ssl_pfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-pfs"], "ObjectVoipProfileSip-SslPfs"); ok {
			if err = d.Set("ssl_pfs", vv); err != nil {
				return fmt.Errorf("Error reading ssl_pfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_pfs: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenObjectVoipProfileSipSslSendEmptyFrags2edl(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-send-empty-frags"], "ObjectVoipProfileSip-SslSendEmptyFrags"); ok {
			if err = d.Set("ssl_send_empty_frags", vv); err != nil {
				return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("ssl_server_certificate", flattenObjectVoipProfileSipSslServerCertificate2edl(o["ssl-server-certificate"], d, "ssl_server_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-certificate"], "ObjectVoipProfileSip-SslServerCertificate"); ok {
			if err = d.Set("ssl_server_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_certificate: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectVoipProfileSipStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectVoipProfileSip-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("strict_register", flattenObjectVoipProfileSipStrictRegister2edl(o["strict-register"], d, "strict_register")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-register"], "ObjectVoipProfileSip-StrictRegister"); ok {
			if err = d.Set("strict_register", vv); err != nil {
				return fmt.Errorf("Error reading strict_register: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_register: %v", err)
		}
	}

	if err = d.Set("subscribe_rate", flattenObjectVoipProfileSipSubscribeRate2edl(o["subscribe-rate"], d, "subscribe_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["subscribe-rate"], "ObjectVoipProfileSip-SubscribeRate"); ok {
			if err = d.Set("subscribe_rate", vv); err != nil {
				return fmt.Errorf("Error reading subscribe_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subscribe_rate: %v", err)
		}
	}

	if err = d.Set("subscribe_rate_track", flattenObjectVoipProfileSipSubscribeRateTrack2edl(o["subscribe-rate-track"], d, "subscribe_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["subscribe-rate-track"], "ObjectVoipProfileSip-SubscribeRateTrack"); ok {
			if err = d.Set("subscribe_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading subscribe_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subscribe_rate_track: %v", err)
		}
	}

	if err = d.Set("unknown_header", flattenObjectVoipProfileSipUnknownHeader2edl(o["unknown-header"], d, "unknown_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-header"], "ObjectVoipProfileSip-UnknownHeader"); ok {
			if err = d.Set("unknown_header", vv); err != nil {
				return fmt.Errorf("Error reading unknown_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_header: %v", err)
		}
	}

	if err = d.Set("update_rate", flattenObjectVoipProfileSipUpdateRate2edl(o["update-rate"], d, "update_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-rate"], "ObjectVoipProfileSip-UpdateRate"); ok {
			if err = d.Set("update_rate", vv); err != nil {
				return fmt.Errorf("Error reading update_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_rate: %v", err)
		}
	}

	if err = d.Set("update_rate_track", flattenObjectVoipProfileSipUpdateRateTrack2edl(o["update-rate-track"], d, "update_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-rate-track"], "ObjectVoipProfileSip-UpdateRateTrack"); ok {
			if err = d.Set("update_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading update_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_rate_track: %v", err)
		}
	}

	return nil
}

func flattenObjectVoipProfileSipFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVoipProfileSipAckRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipAckRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockAck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockBye2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockCancel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockGeoRedOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockInvite2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockLongLines2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockMessage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockNotify2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockPrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockPublish2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockRefer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockRegister2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockSubscribe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockUnknown2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipBlockUpdate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipByeRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipByeRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipCallIdRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipCallKeepalive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipCancelRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipCancelRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipContactFixup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipContentTypeRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipHntRestrictSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipHostedNatTraversal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipInfoRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipInfoRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipInviteRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipInviteRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipIpsRtp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipLogCallSummary2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipLogViolations2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderAllow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderCallId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderContact2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderContentLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderContentType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderCseq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderExpires2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderFrom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderMaxForwards2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderNoProxyRequire2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderNoRequire2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderPAssertedIdentity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderRack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderRecordRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderRseq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpA2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpB2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpC2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpI2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpK2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpM2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpO2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpR2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpS2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpT2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderSdpZ2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderTo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedHeaderVia2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMalformedRequestLine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMaxBodyLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMaxDialogs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMaxIdleDialogs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMaxLineLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMessageRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipMessageRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipNatPortRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipNatTrace2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipNoSdpFixup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipNotifyRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipNotifyRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOpenContactPinhole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOpenRecordRoutePinhole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOpenRegisterPinhole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOpenViaPinhole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOptionsRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipOptionsRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipPrackRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipPrackRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipPreserveOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipProvisionalInviteExpiryTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipPublishRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipPublishRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipReferRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipReferRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipRegisterContactTrace2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipRegisterRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipRegisterRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipRfc2543Branch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipRtp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslAlgorithm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslAuthClient2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslAuthServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslClientRenegotiation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslMaxVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslMinVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslPfs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslSendEmptyFrags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSslServerCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipStrictRegister2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSubscribeRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipSubscribeRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipUnknownHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipUpdateRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSipUpdateRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVoipProfileSip(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ack_rate"); ok || d.HasChange("ack_rate") {
		t, err := expandObjectVoipProfileSipAckRate2edl(d, v, "ack_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ack-rate"] = t
		}
	}

	if v, ok := d.GetOk("ack_rate_track"); ok || d.HasChange("ack_rate_track") {
		t, err := expandObjectVoipProfileSipAckRateTrack2edl(d, v, "ack_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ack-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("block_ack"); ok || d.HasChange("block_ack") {
		t, err := expandObjectVoipProfileSipBlockAck2edl(d, v, "block_ack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-ack"] = t
		}
	}

	if v, ok := d.GetOk("block_bye"); ok || d.HasChange("block_bye") {
		t, err := expandObjectVoipProfileSipBlockBye2edl(d, v, "block_bye")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-bye"] = t
		}
	}

	if v, ok := d.GetOk("block_cancel"); ok || d.HasChange("block_cancel") {
		t, err := expandObjectVoipProfileSipBlockCancel2edl(d, v, "block_cancel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-cancel"] = t
		}
	}

	if v, ok := d.GetOk("block_geo_red_options"); ok || d.HasChange("block_geo_red_options") {
		t, err := expandObjectVoipProfileSipBlockGeoRedOptions2edl(d, v, "block_geo_red_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-geo-red-options"] = t
		}
	}

	if v, ok := d.GetOk("block_info"); ok || d.HasChange("block_info") {
		t, err := expandObjectVoipProfileSipBlockInfo2edl(d, v, "block_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-info"] = t
		}
	}

	if v, ok := d.GetOk("block_invite"); ok || d.HasChange("block_invite") {
		t, err := expandObjectVoipProfileSipBlockInvite2edl(d, v, "block_invite")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-invite"] = t
		}
	}

	if v, ok := d.GetOk("block_long_lines"); ok || d.HasChange("block_long_lines") {
		t, err := expandObjectVoipProfileSipBlockLongLines2edl(d, v, "block_long_lines")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-long-lines"] = t
		}
	}

	if v, ok := d.GetOk("block_message"); ok || d.HasChange("block_message") {
		t, err := expandObjectVoipProfileSipBlockMessage2edl(d, v, "block_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-message"] = t
		}
	}

	if v, ok := d.GetOk("block_notify"); ok || d.HasChange("block_notify") {
		t, err := expandObjectVoipProfileSipBlockNotify2edl(d, v, "block_notify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-notify"] = t
		}
	}

	if v, ok := d.GetOk("block_options"); ok || d.HasChange("block_options") {
		t, err := expandObjectVoipProfileSipBlockOptions2edl(d, v, "block_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-options"] = t
		}
	}

	if v, ok := d.GetOk("block_prack"); ok || d.HasChange("block_prack") {
		t, err := expandObjectVoipProfileSipBlockPrack2edl(d, v, "block_prack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-prack"] = t
		}
	}

	if v, ok := d.GetOk("block_publish"); ok || d.HasChange("block_publish") {
		t, err := expandObjectVoipProfileSipBlockPublish2edl(d, v, "block_publish")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-publish"] = t
		}
	}

	if v, ok := d.GetOk("block_refer"); ok || d.HasChange("block_refer") {
		t, err := expandObjectVoipProfileSipBlockRefer2edl(d, v, "block_refer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-refer"] = t
		}
	}

	if v, ok := d.GetOk("block_register"); ok || d.HasChange("block_register") {
		t, err := expandObjectVoipProfileSipBlockRegister2edl(d, v, "block_register")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-register"] = t
		}
	}

	if v, ok := d.GetOk("block_subscribe"); ok || d.HasChange("block_subscribe") {
		t, err := expandObjectVoipProfileSipBlockSubscribe2edl(d, v, "block_subscribe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-subscribe"] = t
		}
	}

	if v, ok := d.GetOk("block_unknown"); ok || d.HasChange("block_unknown") {
		t, err := expandObjectVoipProfileSipBlockUnknown2edl(d, v, "block_unknown")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-unknown"] = t
		}
	}

	if v, ok := d.GetOk("block_update"); ok || d.HasChange("block_update") {
		t, err := expandObjectVoipProfileSipBlockUpdate2edl(d, v, "block_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-update"] = t
		}
	}

	if v, ok := d.GetOk("bye_rate"); ok || d.HasChange("bye_rate") {
		t, err := expandObjectVoipProfileSipByeRate2edl(d, v, "bye_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bye-rate"] = t
		}
	}

	if v, ok := d.GetOk("bye_rate_track"); ok || d.HasChange("bye_rate_track") {
		t, err := expandObjectVoipProfileSipByeRateTrack2edl(d, v, "bye_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bye-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("call_id_regex"); ok || d.HasChange("call_id_regex") {
		t, err := expandObjectVoipProfileSipCallIdRegex2edl(d, v, "call_id_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-id-regex"] = t
		}
	}

	if v, ok := d.GetOk("call_keepalive"); ok || d.HasChange("call_keepalive") {
		t, err := expandObjectVoipProfileSipCallKeepalive2edl(d, v, "call_keepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-keepalive"] = t
		}
	}

	if v, ok := d.GetOk("cancel_rate"); ok || d.HasChange("cancel_rate") {
		t, err := expandObjectVoipProfileSipCancelRate2edl(d, v, "cancel_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cancel-rate"] = t
		}
	}

	if v, ok := d.GetOk("cancel_rate_track"); ok || d.HasChange("cancel_rate_track") {
		t, err := expandObjectVoipProfileSipCancelRateTrack2edl(d, v, "cancel_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cancel-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("contact_fixup"); ok || d.HasChange("contact_fixup") {
		t, err := expandObjectVoipProfileSipContactFixup2edl(d, v, "contact_fixup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact-fixup"] = t
		}
	}

	if v, ok := d.GetOk("content_type_regex"); ok || d.HasChange("content_type_regex") {
		t, err := expandObjectVoipProfileSipContentTypeRegex2edl(d, v, "content_type_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-type-regex"] = t
		}
	}

	if v, ok := d.GetOk("hnt_restrict_source_ip"); ok || d.HasChange("hnt_restrict_source_ip") {
		t, err := expandObjectVoipProfileSipHntRestrictSourceIp2edl(d, v, "hnt_restrict_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hnt-restrict-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("hosted_nat_traversal"); ok || d.HasChange("hosted_nat_traversal") {
		t, err := expandObjectVoipProfileSipHostedNatTraversal2edl(d, v, "hosted_nat_traversal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosted-nat-traversal"] = t
		}
	}

	if v, ok := d.GetOk("info_rate"); ok || d.HasChange("info_rate") {
		t, err := expandObjectVoipProfileSipInfoRate2edl(d, v, "info_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["info-rate"] = t
		}
	}

	if v, ok := d.GetOk("info_rate_track"); ok || d.HasChange("info_rate_track") {
		t, err := expandObjectVoipProfileSipInfoRateTrack2edl(d, v, "info_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["info-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("invite_rate"); ok || d.HasChange("invite_rate") {
		t, err := expandObjectVoipProfileSipInviteRate2edl(d, v, "invite_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invite-rate"] = t
		}
	}

	if v, ok := d.GetOk("invite_rate_track"); ok || d.HasChange("invite_rate_track") {
		t, err := expandObjectVoipProfileSipInviteRateTrack2edl(d, v, "invite_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invite-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("ips_rtp"); ok || d.HasChange("ips_rtp") {
		t, err := expandObjectVoipProfileSipIpsRtp2edl(d, v, "ips_rtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-rtp"] = t
		}
	}

	if v, ok := d.GetOk("log_call_summary"); ok || d.HasChange("log_call_summary") {
		t, err := expandObjectVoipProfileSipLogCallSummary2edl(d, v, "log_call_summary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-call-summary"] = t
		}
	}

	if v, ok := d.GetOk("log_violations"); ok || d.HasChange("log_violations") {
		t, err := expandObjectVoipProfileSipLogViolations2edl(d, v, "log_violations")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-violations"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_allow"); ok || d.HasChange("malformed_header_allow") {
		t, err := expandObjectVoipProfileSipMalformedHeaderAllow2edl(d, v, "malformed_header_allow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-allow"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_call_id"); ok || d.HasChange("malformed_header_call_id") {
		t, err := expandObjectVoipProfileSipMalformedHeaderCallId2edl(d, v, "malformed_header_call_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-call-id"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_contact"); ok || d.HasChange("malformed_header_contact") {
		t, err := expandObjectVoipProfileSipMalformedHeaderContact2edl(d, v, "malformed_header_contact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-contact"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_content_length"); ok || d.HasChange("malformed_header_content_length") {
		t, err := expandObjectVoipProfileSipMalformedHeaderContentLength2edl(d, v, "malformed_header_content_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-content-length"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_content_type"); ok || d.HasChange("malformed_header_content_type") {
		t, err := expandObjectVoipProfileSipMalformedHeaderContentType2edl(d, v, "malformed_header_content_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-content-type"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_cseq"); ok || d.HasChange("malformed_header_cseq") {
		t, err := expandObjectVoipProfileSipMalformedHeaderCseq2edl(d, v, "malformed_header_cseq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-cseq"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_expires"); ok || d.HasChange("malformed_header_expires") {
		t, err := expandObjectVoipProfileSipMalformedHeaderExpires2edl(d, v, "malformed_header_expires")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-expires"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_from"); ok || d.HasChange("malformed_header_from") {
		t, err := expandObjectVoipProfileSipMalformedHeaderFrom2edl(d, v, "malformed_header_from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-from"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_max_forwards"); ok || d.HasChange("malformed_header_max_forwards") {
		t, err := expandObjectVoipProfileSipMalformedHeaderMaxForwards2edl(d, v, "malformed_header_max_forwards")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-max-forwards"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_no_proxy_require"); ok || d.HasChange("malformed_header_no_proxy_require") {
		t, err := expandObjectVoipProfileSipMalformedHeaderNoProxyRequire2edl(d, v, "malformed_header_no_proxy_require")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-no-proxy-require"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_no_require"); ok || d.HasChange("malformed_header_no_require") {
		t, err := expandObjectVoipProfileSipMalformedHeaderNoRequire2edl(d, v, "malformed_header_no_require")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-no-require"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_p_asserted_identity"); ok || d.HasChange("malformed_header_p_asserted_identity") {
		t, err := expandObjectVoipProfileSipMalformedHeaderPAssertedIdentity2edl(d, v, "malformed_header_p_asserted_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-p-asserted-identity"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_rack"); ok || d.HasChange("malformed_header_rack") {
		t, err := expandObjectVoipProfileSipMalformedHeaderRack2edl(d, v, "malformed_header_rack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-rack"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_record_route"); ok || d.HasChange("malformed_header_record_route") {
		t, err := expandObjectVoipProfileSipMalformedHeaderRecordRoute2edl(d, v, "malformed_header_record_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-record-route"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_route"); ok || d.HasChange("malformed_header_route") {
		t, err := expandObjectVoipProfileSipMalformedHeaderRoute2edl(d, v, "malformed_header_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-route"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_rseq"); ok || d.HasChange("malformed_header_rseq") {
		t, err := expandObjectVoipProfileSipMalformedHeaderRseq2edl(d, v, "malformed_header_rseq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-rseq"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_a"); ok || d.HasChange("malformed_header_sdp_a") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpA2edl(d, v, "malformed_header_sdp_a")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-a"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_b"); ok || d.HasChange("malformed_header_sdp_b") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpB2edl(d, v, "malformed_header_sdp_b")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-b"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_c"); ok || d.HasChange("malformed_header_sdp_c") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpC2edl(d, v, "malformed_header_sdp_c")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-c"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_i"); ok || d.HasChange("malformed_header_sdp_i") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpI2edl(d, v, "malformed_header_sdp_i")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-i"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_k"); ok || d.HasChange("malformed_header_sdp_k") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpK2edl(d, v, "malformed_header_sdp_k")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-k"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_m"); ok || d.HasChange("malformed_header_sdp_m") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpM2edl(d, v, "malformed_header_sdp_m")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-m"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_o"); ok || d.HasChange("malformed_header_sdp_o") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpO2edl(d, v, "malformed_header_sdp_o")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-o"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_r"); ok || d.HasChange("malformed_header_sdp_r") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpR2edl(d, v, "malformed_header_sdp_r")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-r"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_s"); ok || d.HasChange("malformed_header_sdp_s") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpS2edl(d, v, "malformed_header_sdp_s")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-s"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_t"); ok || d.HasChange("malformed_header_sdp_t") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpT2edl(d, v, "malformed_header_sdp_t")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-t"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_v"); ok || d.HasChange("malformed_header_sdp_v") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpV2edl(d, v, "malformed_header_sdp_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-v"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_z"); ok || d.HasChange("malformed_header_sdp_z") {
		t, err := expandObjectVoipProfileSipMalformedHeaderSdpZ2edl(d, v, "malformed_header_sdp_z")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-z"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_to"); ok || d.HasChange("malformed_header_to") {
		t, err := expandObjectVoipProfileSipMalformedHeaderTo2edl(d, v, "malformed_header_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-to"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_via"); ok || d.HasChange("malformed_header_via") {
		t, err := expandObjectVoipProfileSipMalformedHeaderVia2edl(d, v, "malformed_header_via")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-via"] = t
		}
	}

	if v, ok := d.GetOk("malformed_request_line"); ok || d.HasChange("malformed_request_line") {
		t, err := expandObjectVoipProfileSipMalformedRequestLine2edl(d, v, "malformed_request_line")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-request-line"] = t
		}
	}

	if v, ok := d.GetOk("max_body_length"); ok || d.HasChange("max_body_length") {
		t, err := expandObjectVoipProfileSipMaxBodyLength2edl(d, v, "max_body_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-body-length"] = t
		}
	}

	if v, ok := d.GetOk("max_dialogs"); ok || d.HasChange("max_dialogs") {
		t, err := expandObjectVoipProfileSipMaxDialogs2edl(d, v, "max_dialogs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-dialogs"] = t
		}
	}

	if v, ok := d.GetOk("max_idle_dialogs"); ok || d.HasChange("max_idle_dialogs") {
		t, err := expandObjectVoipProfileSipMaxIdleDialogs2edl(d, v, "max_idle_dialogs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-idle-dialogs"] = t
		}
	}

	if v, ok := d.GetOk("max_line_length"); ok || d.HasChange("max_line_length") {
		t, err := expandObjectVoipProfileSipMaxLineLength2edl(d, v, "max_line_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-line-length"] = t
		}
	}

	if v, ok := d.GetOk("message_rate"); ok || d.HasChange("message_rate") {
		t, err := expandObjectVoipProfileSipMessageRate2edl(d, v, "message_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-rate"] = t
		}
	}

	if v, ok := d.GetOk("message_rate_track"); ok || d.HasChange("message_rate_track") {
		t, err := expandObjectVoipProfileSipMessageRateTrack2edl(d, v, "message_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("nat_port_range"); ok || d.HasChange("nat_port_range") {
		t, err := expandObjectVoipProfileSipNatPortRange2edl(d, v, "nat_port_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-port-range"] = t
		}
	}

	if v, ok := d.GetOk("nat_trace"); ok || d.HasChange("nat_trace") {
		t, err := expandObjectVoipProfileSipNatTrace2edl(d, v, "nat_trace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-trace"] = t
		}
	}

	if v, ok := d.GetOk("no_sdp_fixup"); ok || d.HasChange("no_sdp_fixup") {
		t, err := expandObjectVoipProfileSipNoSdpFixup2edl(d, v, "no_sdp_fixup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["no-sdp-fixup"] = t
		}
	}

	if v, ok := d.GetOk("notify_rate"); ok || d.HasChange("notify_rate") {
		t, err := expandObjectVoipProfileSipNotifyRate2edl(d, v, "notify_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-rate"] = t
		}
	}

	if v, ok := d.GetOk("notify_rate_track"); ok || d.HasChange("notify_rate_track") {
		t, err := expandObjectVoipProfileSipNotifyRateTrack2edl(d, v, "notify_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("open_contact_pinhole"); ok || d.HasChange("open_contact_pinhole") {
		t, err := expandObjectVoipProfileSipOpenContactPinhole2edl(d, v, "open_contact_pinhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-contact-pinhole"] = t
		}
	}

	if v, ok := d.GetOk("open_record_route_pinhole"); ok || d.HasChange("open_record_route_pinhole") {
		t, err := expandObjectVoipProfileSipOpenRecordRoutePinhole2edl(d, v, "open_record_route_pinhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-record-route-pinhole"] = t
		}
	}

	if v, ok := d.GetOk("open_register_pinhole"); ok || d.HasChange("open_register_pinhole") {
		t, err := expandObjectVoipProfileSipOpenRegisterPinhole2edl(d, v, "open_register_pinhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-register-pinhole"] = t
		}
	}

	if v, ok := d.GetOk("open_via_pinhole"); ok || d.HasChange("open_via_pinhole") {
		t, err := expandObjectVoipProfileSipOpenViaPinhole2edl(d, v, "open_via_pinhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-via-pinhole"] = t
		}
	}

	if v, ok := d.GetOk("options_rate"); ok || d.HasChange("options_rate") {
		t, err := expandObjectVoipProfileSipOptionsRate2edl(d, v, "options_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options-rate"] = t
		}
	}

	if v, ok := d.GetOk("options_rate_track"); ok || d.HasChange("options_rate_track") {
		t, err := expandObjectVoipProfileSipOptionsRateTrack2edl(d, v, "options_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("prack_rate"); ok || d.HasChange("prack_rate") {
		t, err := expandObjectVoipProfileSipPrackRate2edl(d, v, "prack_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prack-rate"] = t
		}
	}

	if v, ok := d.GetOk("prack_rate_track"); ok || d.HasChange("prack_rate_track") {
		t, err := expandObjectVoipProfileSipPrackRateTrack2edl(d, v, "prack_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prack-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("preserve_override"); ok || d.HasChange("preserve_override") {
		t, err := expandObjectVoipProfileSipPreserveOverride2edl(d, v, "preserve_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preserve-override"] = t
		}
	}

	if v, ok := d.GetOk("provisional_invite_expiry_time"); ok || d.HasChange("provisional_invite_expiry_time") {
		t, err := expandObjectVoipProfileSipProvisionalInviteExpiryTime2edl(d, v, "provisional_invite_expiry_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["provisional-invite-expiry-time"] = t
		}
	}

	if v, ok := d.GetOk("publish_rate"); ok || d.HasChange("publish_rate") {
		t, err := expandObjectVoipProfileSipPublishRate2edl(d, v, "publish_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["publish-rate"] = t
		}
	}

	if v, ok := d.GetOk("publish_rate_track"); ok || d.HasChange("publish_rate_track") {
		t, err := expandObjectVoipProfileSipPublishRateTrack2edl(d, v, "publish_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["publish-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("refer_rate"); ok || d.HasChange("refer_rate") {
		t, err := expandObjectVoipProfileSipReferRate2edl(d, v, "refer_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["refer-rate"] = t
		}
	}

	if v, ok := d.GetOk("refer_rate_track"); ok || d.HasChange("refer_rate_track") {
		t, err := expandObjectVoipProfileSipReferRateTrack2edl(d, v, "refer_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["refer-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("register_contact_trace"); ok || d.HasChange("register_contact_trace") {
		t, err := expandObjectVoipProfileSipRegisterContactTrace2edl(d, v, "register_contact_trace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-contact-trace"] = t
		}
	}

	if v, ok := d.GetOk("register_rate"); ok || d.HasChange("register_rate") {
		t, err := expandObjectVoipProfileSipRegisterRate2edl(d, v, "register_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-rate"] = t
		}
	}

	if v, ok := d.GetOk("register_rate_track"); ok || d.HasChange("register_rate_track") {
		t, err := expandObjectVoipProfileSipRegisterRateTrack2edl(d, v, "register_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("rfc2543_branch"); ok || d.HasChange("rfc2543_branch") {
		t, err := expandObjectVoipProfileSipRfc2543Branch2edl(d, v, "rfc2543_branch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rfc2543-branch"] = t
		}
	}

	if v, ok := d.GetOk("rtp"); ok || d.HasChange("rtp") {
		t, err := expandObjectVoipProfileSipRtp2edl(d, v, "rtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtp"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandObjectVoipProfileSipSslAlgorithm2edl(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_auth_client"); ok || d.HasChange("ssl_auth_client") {
		t, err := expandObjectVoipProfileSipSslAuthClient2edl(d, v, "ssl_auth_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-auth-client"] = t
		}
	}

	if v, ok := d.GetOk("ssl_auth_server"); ok || d.HasChange("ssl_auth_server") {
		t, err := expandObjectVoipProfileSipSslAuthServer2edl(d, v, "ssl_auth_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-auth-server"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_certificate"); ok || d.HasChange("ssl_client_certificate") {
		t, err := expandObjectVoipProfileSipSslClientCertificate2edl(d, v, "ssl_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok || d.HasChange("ssl_client_renegotiation") {
		t, err := expandObjectVoipProfileSipSslClientRenegotiation2edl(d, v, "ssl_client_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandObjectVoipProfileSipSslMaxVersion2edl(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandObjectVoipProfileSipSslMinVersion2edl(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok || d.HasChange("ssl_mode") {
		t, err := expandObjectVoipProfileSipSslMode2edl(d, v, "ssl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_pfs"); ok || d.HasChange("ssl_pfs") {
		t, err := expandObjectVoipProfileSipSslPfs2edl(d, v, "ssl_pfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-pfs"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok || d.HasChange("ssl_send_empty_frags") {
		t, err := expandObjectVoipProfileSipSslSendEmptyFrags2edl(d, v, "ssl_send_empty_frags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_certificate"); ok || d.HasChange("ssl_server_certificate") {
		t, err := expandObjectVoipProfileSipSslServerCertificate2edl(d, v, "ssl_server_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-certificate"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectVoipProfileSipStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("strict_register"); ok || d.HasChange("strict_register") {
		t, err := expandObjectVoipProfileSipStrictRegister2edl(d, v, "strict_register")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-register"] = t
		}
	}

	if v, ok := d.GetOk("subscribe_rate"); ok || d.HasChange("subscribe_rate") {
		t, err := expandObjectVoipProfileSipSubscribeRate2edl(d, v, "subscribe_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subscribe-rate"] = t
		}
	}

	if v, ok := d.GetOk("subscribe_rate_track"); ok || d.HasChange("subscribe_rate_track") {
		t, err := expandObjectVoipProfileSipSubscribeRateTrack2edl(d, v, "subscribe_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subscribe-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("unknown_header"); ok || d.HasChange("unknown_header") {
		t, err := expandObjectVoipProfileSipUnknownHeader2edl(d, v, "unknown_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-header"] = t
		}
	}

	if v, ok := d.GetOk("update_rate"); ok || d.HasChange("update_rate") {
		t, err := expandObjectVoipProfileSipUpdateRate2edl(d, v, "update_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-rate"] = t
		}
	}

	if v, ok := d.GetOk("update_rate_track"); ok || d.HasChange("update_rate_track") {
		t, err := expandObjectVoipProfileSipUpdateRateTrack2edl(d, v, "update_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-rate-track"] = t
		}
	}

	return &obj, nil
}
