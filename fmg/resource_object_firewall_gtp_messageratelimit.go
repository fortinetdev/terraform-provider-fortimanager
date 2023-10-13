// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Message rate limiting.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallGtpMessageRateLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpMessageRateLimitUpdate,
		Read:   resourceObjectFirewallGtpMessageRateLimitRead,
		Update: resourceObjectFirewallGtpMessageRateLimitUpdate,
		Delete: resourceObjectFirewallGtpMessageRateLimitDelete,

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
			"create_aa_pdp_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"create_aa_pdp_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"create_mbms_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"create_mbms_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"create_pdp_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"create_pdp_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delete_aa_pdp_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delete_aa_pdp_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delete_mbms_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delete_mbms_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delete_pdp_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delete_pdp_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"echo_reponse": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"echo_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"error_indication": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"failure_report_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"failure_report_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fwd_reloc_complete_ack": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fwd_relocation_complete": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fwd_relocation_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fwd_relocation_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fwd_srns_context": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fwd_srns_context_ack": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"g_pdu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"identification_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"identification_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_de_reg_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_de_reg_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_notify_rej_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_notify_rej_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_notify_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_notify_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_reg_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_reg_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_ses_start_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_ses_start_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_ses_stop_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbms_ses_stop_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"note_ms_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"note_ms_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pdu_notify_rej_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pdu_notify_rej_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pdu_notify_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pdu_notify_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ran_info": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"relocation_cancel_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"relocation_cancel_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"send_route_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"send_route_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sgsn_context_ack": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sgsn_context_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sgsn_context_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"support_ext_hdr_notify": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_mbms_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_mbms_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_pdp_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_pdp_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"version_not_support": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallGtpMessageRateLimitUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallGtpMessageRateLimit(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpMessageRateLimit resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtpMessageRateLimit(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpMessageRateLimit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallGtpMessageRateLimit")

	return resourceObjectFirewallGtpMessageRateLimitRead(d, m)
}

func resourceObjectFirewallGtpMessageRateLimitDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallGtpMessageRateLimit(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtpMessageRateLimit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpMessageRateLimitRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallGtpMessageRateLimit(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpMessageRateLimit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtpMessageRateLimit(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpMessageRateLimit resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpMessageRateLimitCreateAaPdpRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitCreateAaPdpResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitCreateMbmsRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitCreateMbmsResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitCreatePdpRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitCreatePdpResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeleteAaPdpRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeleteAaPdpResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeleteMbmsRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeleteMbmsResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeletePdpRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeletePdpResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitEchoReponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitEchoRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitErrorIndication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFailureReportRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFailureReportResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdRelocCompleteAck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdRelocationComplete2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdRelocationRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdRelocationResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdSrnsContext2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdSrnsContextAck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitGPdu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitIdentificationRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitIdentificationResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsDeRegRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsDeRegResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRejRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRejResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsNotifyResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsRegRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsRegResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsSesStartRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsSesStartResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsSesStopRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsSesStopResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitNoteMsRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitNoteMsResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitPduNotifyRejRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitPduNotifyRejResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitPduNotifyRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitPduNotifyResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitRanInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitRelocationCancelRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitRelocationCancelResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSendRouteRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSendRouteResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSgsnContextAck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSgsnContextRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSgsnContextResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSupportExtHdrNotify2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitUpdateMbmsRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitUpdateMbmsResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitUpdatePdpRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitUpdatePdpResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitVersionNotSupport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallGtpMessageRateLimit(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("create_aa_pdp_request", flattenObjectFirewallGtpMessageRateLimitCreateAaPdpRequest2edl(o["create-aa-pdp-request"], d, "create_aa_pdp_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-aa-pdp-request"], "ObjectFirewallGtpMessageRateLimit-CreateAaPdpRequest"); ok {
			if err = d.Set("create_aa_pdp_request", vv); err != nil {
				return fmt.Errorf("Error reading create_aa_pdp_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_aa_pdp_request: %v", err)
		}
	}

	if err = d.Set("create_aa_pdp_response", flattenObjectFirewallGtpMessageRateLimitCreateAaPdpResponse2edl(o["create-aa-pdp-response"], d, "create_aa_pdp_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-aa-pdp-response"], "ObjectFirewallGtpMessageRateLimit-CreateAaPdpResponse"); ok {
			if err = d.Set("create_aa_pdp_response", vv); err != nil {
				return fmt.Errorf("Error reading create_aa_pdp_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_aa_pdp_response: %v", err)
		}
	}

	if err = d.Set("create_mbms_request", flattenObjectFirewallGtpMessageRateLimitCreateMbmsRequest2edl(o["create-mbms-request"], d, "create_mbms_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-mbms-request"], "ObjectFirewallGtpMessageRateLimit-CreateMbmsRequest"); ok {
			if err = d.Set("create_mbms_request", vv); err != nil {
				return fmt.Errorf("Error reading create_mbms_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_mbms_request: %v", err)
		}
	}

	if err = d.Set("create_mbms_response", flattenObjectFirewallGtpMessageRateLimitCreateMbmsResponse2edl(o["create-mbms-response"], d, "create_mbms_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-mbms-response"], "ObjectFirewallGtpMessageRateLimit-CreateMbmsResponse"); ok {
			if err = d.Set("create_mbms_response", vv); err != nil {
				return fmt.Errorf("Error reading create_mbms_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_mbms_response: %v", err)
		}
	}

	if err = d.Set("create_pdp_request", flattenObjectFirewallGtpMessageRateLimitCreatePdpRequest2edl(o["create-pdp-request"], d, "create_pdp_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-pdp-request"], "ObjectFirewallGtpMessageRateLimit-CreatePdpRequest"); ok {
			if err = d.Set("create_pdp_request", vv); err != nil {
				return fmt.Errorf("Error reading create_pdp_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_pdp_request: %v", err)
		}
	}

	if err = d.Set("create_pdp_response", flattenObjectFirewallGtpMessageRateLimitCreatePdpResponse2edl(o["create-pdp-response"], d, "create_pdp_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-pdp-response"], "ObjectFirewallGtpMessageRateLimit-CreatePdpResponse"); ok {
			if err = d.Set("create_pdp_response", vv); err != nil {
				return fmt.Errorf("Error reading create_pdp_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_pdp_response: %v", err)
		}
	}

	if err = d.Set("delete_aa_pdp_request", flattenObjectFirewallGtpMessageRateLimitDeleteAaPdpRequest2edl(o["delete-aa-pdp-request"], d, "delete_aa_pdp_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-aa-pdp-request"], "ObjectFirewallGtpMessageRateLimit-DeleteAaPdpRequest"); ok {
			if err = d.Set("delete_aa_pdp_request", vv); err != nil {
				return fmt.Errorf("Error reading delete_aa_pdp_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_aa_pdp_request: %v", err)
		}
	}

	if err = d.Set("delete_aa_pdp_response", flattenObjectFirewallGtpMessageRateLimitDeleteAaPdpResponse2edl(o["delete-aa-pdp-response"], d, "delete_aa_pdp_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-aa-pdp-response"], "ObjectFirewallGtpMessageRateLimit-DeleteAaPdpResponse"); ok {
			if err = d.Set("delete_aa_pdp_response", vv); err != nil {
				return fmt.Errorf("Error reading delete_aa_pdp_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_aa_pdp_response: %v", err)
		}
	}

	if err = d.Set("delete_mbms_request", flattenObjectFirewallGtpMessageRateLimitDeleteMbmsRequest2edl(o["delete-mbms-request"], d, "delete_mbms_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-mbms-request"], "ObjectFirewallGtpMessageRateLimit-DeleteMbmsRequest"); ok {
			if err = d.Set("delete_mbms_request", vv); err != nil {
				return fmt.Errorf("Error reading delete_mbms_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_mbms_request: %v", err)
		}
	}

	if err = d.Set("delete_mbms_response", flattenObjectFirewallGtpMessageRateLimitDeleteMbmsResponse2edl(o["delete-mbms-response"], d, "delete_mbms_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-mbms-response"], "ObjectFirewallGtpMessageRateLimit-DeleteMbmsResponse"); ok {
			if err = d.Set("delete_mbms_response", vv); err != nil {
				return fmt.Errorf("Error reading delete_mbms_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_mbms_response: %v", err)
		}
	}

	if err = d.Set("delete_pdp_request", flattenObjectFirewallGtpMessageRateLimitDeletePdpRequest2edl(o["delete-pdp-request"], d, "delete_pdp_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-pdp-request"], "ObjectFirewallGtpMessageRateLimit-DeletePdpRequest"); ok {
			if err = d.Set("delete_pdp_request", vv); err != nil {
				return fmt.Errorf("Error reading delete_pdp_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_pdp_request: %v", err)
		}
	}

	if err = d.Set("delete_pdp_response", flattenObjectFirewallGtpMessageRateLimitDeletePdpResponse2edl(o["delete-pdp-response"], d, "delete_pdp_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-pdp-response"], "ObjectFirewallGtpMessageRateLimit-DeletePdpResponse"); ok {
			if err = d.Set("delete_pdp_response", vv); err != nil {
				return fmt.Errorf("Error reading delete_pdp_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_pdp_response: %v", err)
		}
	}

	if err = d.Set("echo_reponse", flattenObjectFirewallGtpMessageRateLimitEchoReponse2edl(o["echo-reponse"], d, "echo_reponse")); err != nil {
		if vv, ok := fortiAPIPatch(o["echo-reponse"], "ObjectFirewallGtpMessageRateLimit-EchoReponse"); ok {
			if err = d.Set("echo_reponse", vv); err != nil {
				return fmt.Errorf("Error reading echo_reponse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading echo_reponse: %v", err)
		}
	}

	if err = d.Set("echo_request", flattenObjectFirewallGtpMessageRateLimitEchoRequest2edl(o["echo-request"], d, "echo_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["echo-request"], "ObjectFirewallGtpMessageRateLimit-EchoRequest"); ok {
			if err = d.Set("echo_request", vv); err != nil {
				return fmt.Errorf("Error reading echo_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading echo_request: %v", err)
		}
	}

	if err = d.Set("error_indication", flattenObjectFirewallGtpMessageRateLimitErrorIndication2edl(o["error-indication"], d, "error_indication")); err != nil {
		if vv, ok := fortiAPIPatch(o["error-indication"], "ObjectFirewallGtpMessageRateLimit-ErrorIndication"); ok {
			if err = d.Set("error_indication", vv); err != nil {
				return fmt.Errorf("Error reading error_indication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading error_indication: %v", err)
		}
	}

	if err = d.Set("failure_report_request", flattenObjectFirewallGtpMessageRateLimitFailureReportRequest2edl(o["failure-report-request"], d, "failure_report_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["failure-report-request"], "ObjectFirewallGtpMessageRateLimit-FailureReportRequest"); ok {
			if err = d.Set("failure_report_request", vv); err != nil {
				return fmt.Errorf("Error reading failure_report_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failure_report_request: %v", err)
		}
	}

	if err = d.Set("failure_report_response", flattenObjectFirewallGtpMessageRateLimitFailureReportResponse2edl(o["failure-report-response"], d, "failure_report_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["failure-report-response"], "ObjectFirewallGtpMessageRateLimit-FailureReportResponse"); ok {
			if err = d.Set("failure_report_response", vv); err != nil {
				return fmt.Errorf("Error reading failure_report_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failure_report_response: %v", err)
		}
	}

	if err = d.Set("fwd_reloc_complete_ack", flattenObjectFirewallGtpMessageRateLimitFwdRelocCompleteAck2edl(o["fwd-reloc-complete-ack"], d, "fwd_reloc_complete_ack")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-reloc-complete-ack"], "ObjectFirewallGtpMessageRateLimit-FwdRelocCompleteAck"); ok {
			if err = d.Set("fwd_reloc_complete_ack", vv); err != nil {
				return fmt.Errorf("Error reading fwd_reloc_complete_ack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_reloc_complete_ack: %v", err)
		}
	}

	if err = d.Set("fwd_relocation_complete", flattenObjectFirewallGtpMessageRateLimitFwdRelocationComplete2edl(o["fwd-relocation-complete"], d, "fwd_relocation_complete")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-relocation-complete"], "ObjectFirewallGtpMessageRateLimit-FwdRelocationComplete"); ok {
			if err = d.Set("fwd_relocation_complete", vv); err != nil {
				return fmt.Errorf("Error reading fwd_relocation_complete: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_relocation_complete: %v", err)
		}
	}

	if err = d.Set("fwd_relocation_request", flattenObjectFirewallGtpMessageRateLimitFwdRelocationRequest2edl(o["fwd-relocation-request"], d, "fwd_relocation_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-relocation-request"], "ObjectFirewallGtpMessageRateLimit-FwdRelocationRequest"); ok {
			if err = d.Set("fwd_relocation_request", vv); err != nil {
				return fmt.Errorf("Error reading fwd_relocation_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_relocation_request: %v", err)
		}
	}

	if err = d.Set("fwd_relocation_response", flattenObjectFirewallGtpMessageRateLimitFwdRelocationResponse2edl(o["fwd-relocation-response"], d, "fwd_relocation_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-relocation-response"], "ObjectFirewallGtpMessageRateLimit-FwdRelocationResponse"); ok {
			if err = d.Set("fwd_relocation_response", vv); err != nil {
				return fmt.Errorf("Error reading fwd_relocation_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_relocation_response: %v", err)
		}
	}

	if err = d.Set("fwd_srns_context", flattenObjectFirewallGtpMessageRateLimitFwdSrnsContext2edl(o["fwd-srns-context"], d, "fwd_srns_context")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-srns-context"], "ObjectFirewallGtpMessageRateLimit-FwdSrnsContext"); ok {
			if err = d.Set("fwd_srns_context", vv); err != nil {
				return fmt.Errorf("Error reading fwd_srns_context: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_srns_context: %v", err)
		}
	}

	if err = d.Set("fwd_srns_context_ack", flattenObjectFirewallGtpMessageRateLimitFwdSrnsContextAck2edl(o["fwd-srns-context-ack"], d, "fwd_srns_context_ack")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-srns-context-ack"], "ObjectFirewallGtpMessageRateLimit-FwdSrnsContextAck"); ok {
			if err = d.Set("fwd_srns_context_ack", vv); err != nil {
				return fmt.Errorf("Error reading fwd_srns_context_ack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_srns_context_ack: %v", err)
		}
	}

	if err = d.Set("g_pdu", flattenObjectFirewallGtpMessageRateLimitGPdu2edl(o["g-pdu"], d, "g_pdu")); err != nil {
		if vv, ok := fortiAPIPatch(o["g-pdu"], "ObjectFirewallGtpMessageRateLimit-GPdu"); ok {
			if err = d.Set("g_pdu", vv); err != nil {
				return fmt.Errorf("Error reading g_pdu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading g_pdu: %v", err)
		}
	}

	if err = d.Set("identification_request", flattenObjectFirewallGtpMessageRateLimitIdentificationRequest2edl(o["identification-request"], d, "identification_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["identification-request"], "ObjectFirewallGtpMessageRateLimit-IdentificationRequest"); ok {
			if err = d.Set("identification_request", vv); err != nil {
				return fmt.Errorf("Error reading identification_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading identification_request: %v", err)
		}
	}

	if err = d.Set("identification_response", flattenObjectFirewallGtpMessageRateLimitIdentificationResponse2edl(o["identification-response"], d, "identification_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["identification-response"], "ObjectFirewallGtpMessageRateLimit-IdentificationResponse"); ok {
			if err = d.Set("identification_response", vv); err != nil {
				return fmt.Errorf("Error reading identification_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading identification_response: %v", err)
		}
	}

	if err = d.Set("mbms_de_reg_request", flattenObjectFirewallGtpMessageRateLimitMbmsDeRegRequest2edl(o["mbms-de-reg-request"], d, "mbms_de_reg_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-de-reg-request"], "ObjectFirewallGtpMessageRateLimit-MbmsDeRegRequest"); ok {
			if err = d.Set("mbms_de_reg_request", vv); err != nil {
				return fmt.Errorf("Error reading mbms_de_reg_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_de_reg_request: %v", err)
		}
	}

	if err = d.Set("mbms_de_reg_response", flattenObjectFirewallGtpMessageRateLimitMbmsDeRegResponse2edl(o["mbms-de-reg-response"], d, "mbms_de_reg_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-de-reg-response"], "ObjectFirewallGtpMessageRateLimit-MbmsDeRegResponse"); ok {
			if err = d.Set("mbms_de_reg_response", vv); err != nil {
				return fmt.Errorf("Error reading mbms_de_reg_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_de_reg_response: %v", err)
		}
	}

	if err = d.Set("mbms_notify_rej_request", flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRejRequest2edl(o["mbms-notify-rej-request"], d, "mbms_notify_rej_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-notify-rej-request"], "ObjectFirewallGtpMessageRateLimit-MbmsNotifyRejRequest"); ok {
			if err = d.Set("mbms_notify_rej_request", vv); err != nil {
				return fmt.Errorf("Error reading mbms_notify_rej_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_notify_rej_request: %v", err)
		}
	}

	if err = d.Set("mbms_notify_rej_response", flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRejResponse2edl(o["mbms-notify-rej-response"], d, "mbms_notify_rej_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-notify-rej-response"], "ObjectFirewallGtpMessageRateLimit-MbmsNotifyRejResponse"); ok {
			if err = d.Set("mbms_notify_rej_response", vv); err != nil {
				return fmt.Errorf("Error reading mbms_notify_rej_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_notify_rej_response: %v", err)
		}
	}

	if err = d.Set("mbms_notify_request", flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRequest2edl(o["mbms-notify-request"], d, "mbms_notify_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-notify-request"], "ObjectFirewallGtpMessageRateLimit-MbmsNotifyRequest"); ok {
			if err = d.Set("mbms_notify_request", vv); err != nil {
				return fmt.Errorf("Error reading mbms_notify_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_notify_request: %v", err)
		}
	}

	if err = d.Set("mbms_notify_response", flattenObjectFirewallGtpMessageRateLimitMbmsNotifyResponse2edl(o["mbms-notify-response"], d, "mbms_notify_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-notify-response"], "ObjectFirewallGtpMessageRateLimit-MbmsNotifyResponse"); ok {
			if err = d.Set("mbms_notify_response", vv); err != nil {
				return fmt.Errorf("Error reading mbms_notify_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_notify_response: %v", err)
		}
	}

	if err = d.Set("mbms_reg_request", flattenObjectFirewallGtpMessageRateLimitMbmsRegRequest2edl(o["mbms-reg-request"], d, "mbms_reg_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-reg-request"], "ObjectFirewallGtpMessageRateLimit-MbmsRegRequest"); ok {
			if err = d.Set("mbms_reg_request", vv); err != nil {
				return fmt.Errorf("Error reading mbms_reg_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_reg_request: %v", err)
		}
	}

	if err = d.Set("mbms_reg_response", flattenObjectFirewallGtpMessageRateLimitMbmsRegResponse2edl(o["mbms-reg-response"], d, "mbms_reg_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-reg-response"], "ObjectFirewallGtpMessageRateLimit-MbmsRegResponse"); ok {
			if err = d.Set("mbms_reg_response", vv); err != nil {
				return fmt.Errorf("Error reading mbms_reg_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_reg_response: %v", err)
		}
	}

	if err = d.Set("mbms_ses_start_request", flattenObjectFirewallGtpMessageRateLimitMbmsSesStartRequest2edl(o["mbms-ses-start-request"], d, "mbms_ses_start_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-ses-start-request"], "ObjectFirewallGtpMessageRateLimit-MbmsSesStartRequest"); ok {
			if err = d.Set("mbms_ses_start_request", vv); err != nil {
				return fmt.Errorf("Error reading mbms_ses_start_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_ses_start_request: %v", err)
		}
	}

	if err = d.Set("mbms_ses_start_response", flattenObjectFirewallGtpMessageRateLimitMbmsSesStartResponse2edl(o["mbms-ses-start-response"], d, "mbms_ses_start_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-ses-start-response"], "ObjectFirewallGtpMessageRateLimit-MbmsSesStartResponse"); ok {
			if err = d.Set("mbms_ses_start_response", vv); err != nil {
				return fmt.Errorf("Error reading mbms_ses_start_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_ses_start_response: %v", err)
		}
	}

	if err = d.Set("mbms_ses_stop_request", flattenObjectFirewallGtpMessageRateLimitMbmsSesStopRequest2edl(o["mbms-ses-stop-request"], d, "mbms_ses_stop_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-ses-stop-request"], "ObjectFirewallGtpMessageRateLimit-MbmsSesStopRequest"); ok {
			if err = d.Set("mbms_ses_stop_request", vv); err != nil {
				return fmt.Errorf("Error reading mbms_ses_stop_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_ses_stop_request: %v", err)
		}
	}

	if err = d.Set("mbms_ses_stop_response", flattenObjectFirewallGtpMessageRateLimitMbmsSesStopResponse2edl(o["mbms-ses-stop-response"], d, "mbms_ses_stop_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-ses-stop-response"], "ObjectFirewallGtpMessageRateLimit-MbmsSesStopResponse"); ok {
			if err = d.Set("mbms_ses_stop_response", vv); err != nil {
				return fmt.Errorf("Error reading mbms_ses_stop_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_ses_stop_response: %v", err)
		}
	}

	if err = d.Set("note_ms_request", flattenObjectFirewallGtpMessageRateLimitNoteMsRequest2edl(o["note-ms-request"], d, "note_ms_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["note-ms-request"], "ObjectFirewallGtpMessageRateLimit-NoteMsRequest"); ok {
			if err = d.Set("note_ms_request", vv); err != nil {
				return fmt.Errorf("Error reading note_ms_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading note_ms_request: %v", err)
		}
	}

	if err = d.Set("note_ms_response", flattenObjectFirewallGtpMessageRateLimitNoteMsResponse2edl(o["note-ms-response"], d, "note_ms_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["note-ms-response"], "ObjectFirewallGtpMessageRateLimit-NoteMsResponse"); ok {
			if err = d.Set("note_ms_response", vv); err != nil {
				return fmt.Errorf("Error reading note_ms_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading note_ms_response: %v", err)
		}
	}

	if err = d.Set("pdu_notify_rej_request", flattenObjectFirewallGtpMessageRateLimitPduNotifyRejRequest2edl(o["pdu-notify-rej-request"], d, "pdu_notify_rej_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdu-notify-rej-request"], "ObjectFirewallGtpMessageRateLimit-PduNotifyRejRequest"); ok {
			if err = d.Set("pdu_notify_rej_request", vv); err != nil {
				return fmt.Errorf("Error reading pdu_notify_rej_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdu_notify_rej_request: %v", err)
		}
	}

	if err = d.Set("pdu_notify_rej_response", flattenObjectFirewallGtpMessageRateLimitPduNotifyRejResponse2edl(o["pdu-notify-rej-response"], d, "pdu_notify_rej_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdu-notify-rej-response"], "ObjectFirewallGtpMessageRateLimit-PduNotifyRejResponse"); ok {
			if err = d.Set("pdu_notify_rej_response", vv); err != nil {
				return fmt.Errorf("Error reading pdu_notify_rej_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdu_notify_rej_response: %v", err)
		}
	}

	if err = d.Set("pdu_notify_request", flattenObjectFirewallGtpMessageRateLimitPduNotifyRequest2edl(o["pdu-notify-request"], d, "pdu_notify_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdu-notify-request"], "ObjectFirewallGtpMessageRateLimit-PduNotifyRequest"); ok {
			if err = d.Set("pdu_notify_request", vv); err != nil {
				return fmt.Errorf("Error reading pdu_notify_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdu_notify_request: %v", err)
		}
	}

	if err = d.Set("pdu_notify_response", flattenObjectFirewallGtpMessageRateLimitPduNotifyResponse2edl(o["pdu-notify-response"], d, "pdu_notify_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdu-notify-response"], "ObjectFirewallGtpMessageRateLimit-PduNotifyResponse"); ok {
			if err = d.Set("pdu_notify_response", vv); err != nil {
				return fmt.Errorf("Error reading pdu_notify_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdu_notify_response: %v", err)
		}
	}

	if err = d.Set("ran_info", flattenObjectFirewallGtpMessageRateLimitRanInfo2edl(o["ran-info"], d, "ran_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["ran-info"], "ObjectFirewallGtpMessageRateLimit-RanInfo"); ok {
			if err = d.Set("ran_info", vv); err != nil {
				return fmt.Errorf("Error reading ran_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ran_info: %v", err)
		}
	}

	if err = d.Set("relocation_cancel_request", flattenObjectFirewallGtpMessageRateLimitRelocationCancelRequest2edl(o["relocation-cancel-request"], d, "relocation_cancel_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["relocation-cancel-request"], "ObjectFirewallGtpMessageRateLimit-RelocationCancelRequest"); ok {
			if err = d.Set("relocation_cancel_request", vv); err != nil {
				return fmt.Errorf("Error reading relocation_cancel_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relocation_cancel_request: %v", err)
		}
	}

	if err = d.Set("relocation_cancel_response", flattenObjectFirewallGtpMessageRateLimitRelocationCancelResponse2edl(o["relocation-cancel-response"], d, "relocation_cancel_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["relocation-cancel-response"], "ObjectFirewallGtpMessageRateLimit-RelocationCancelResponse"); ok {
			if err = d.Set("relocation_cancel_response", vv); err != nil {
				return fmt.Errorf("Error reading relocation_cancel_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relocation_cancel_response: %v", err)
		}
	}

	if err = d.Set("send_route_request", flattenObjectFirewallGtpMessageRateLimitSendRouteRequest2edl(o["send-route-request"], d, "send_route_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-route-request"], "ObjectFirewallGtpMessageRateLimit-SendRouteRequest"); ok {
			if err = d.Set("send_route_request", vv); err != nil {
				return fmt.Errorf("Error reading send_route_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_route_request: %v", err)
		}
	}

	if err = d.Set("send_route_response", flattenObjectFirewallGtpMessageRateLimitSendRouteResponse2edl(o["send-route-response"], d, "send_route_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-route-response"], "ObjectFirewallGtpMessageRateLimit-SendRouteResponse"); ok {
			if err = d.Set("send_route_response", vv); err != nil {
				return fmt.Errorf("Error reading send_route_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_route_response: %v", err)
		}
	}

	if err = d.Set("sgsn_context_ack", flattenObjectFirewallGtpMessageRateLimitSgsnContextAck2edl(o["sgsn-context-ack"], d, "sgsn_context_ack")); err != nil {
		if vv, ok := fortiAPIPatch(o["sgsn-context-ack"], "ObjectFirewallGtpMessageRateLimit-SgsnContextAck"); ok {
			if err = d.Set("sgsn_context_ack", vv); err != nil {
				return fmt.Errorf("Error reading sgsn_context_ack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sgsn_context_ack: %v", err)
		}
	}

	if err = d.Set("sgsn_context_request", flattenObjectFirewallGtpMessageRateLimitSgsnContextRequest2edl(o["sgsn-context-request"], d, "sgsn_context_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["sgsn-context-request"], "ObjectFirewallGtpMessageRateLimit-SgsnContextRequest"); ok {
			if err = d.Set("sgsn_context_request", vv); err != nil {
				return fmt.Errorf("Error reading sgsn_context_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sgsn_context_request: %v", err)
		}
	}

	if err = d.Set("sgsn_context_response", flattenObjectFirewallGtpMessageRateLimitSgsnContextResponse2edl(o["sgsn-context-response"], d, "sgsn_context_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["sgsn-context-response"], "ObjectFirewallGtpMessageRateLimit-SgsnContextResponse"); ok {
			if err = d.Set("sgsn_context_response", vv); err != nil {
				return fmt.Errorf("Error reading sgsn_context_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sgsn_context_response: %v", err)
		}
	}

	if err = d.Set("support_ext_hdr_notify", flattenObjectFirewallGtpMessageRateLimitSupportExtHdrNotify2edl(o["support-ext-hdr-notify"], d, "support_ext_hdr_notify")); err != nil {
		if vv, ok := fortiAPIPatch(o["support-ext-hdr-notify"], "ObjectFirewallGtpMessageRateLimit-SupportExtHdrNotify"); ok {
			if err = d.Set("support_ext_hdr_notify", vv); err != nil {
				return fmt.Errorf("Error reading support_ext_hdr_notify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading support_ext_hdr_notify: %v", err)
		}
	}

	if err = d.Set("update_mbms_request", flattenObjectFirewallGtpMessageRateLimitUpdateMbmsRequest2edl(o["update-mbms-request"], d, "update_mbms_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-mbms-request"], "ObjectFirewallGtpMessageRateLimit-UpdateMbmsRequest"); ok {
			if err = d.Set("update_mbms_request", vv); err != nil {
				return fmt.Errorf("Error reading update_mbms_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_mbms_request: %v", err)
		}
	}

	if err = d.Set("update_mbms_response", flattenObjectFirewallGtpMessageRateLimitUpdateMbmsResponse2edl(o["update-mbms-response"], d, "update_mbms_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-mbms-response"], "ObjectFirewallGtpMessageRateLimit-UpdateMbmsResponse"); ok {
			if err = d.Set("update_mbms_response", vv); err != nil {
				return fmt.Errorf("Error reading update_mbms_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_mbms_response: %v", err)
		}
	}

	if err = d.Set("update_pdp_request", flattenObjectFirewallGtpMessageRateLimitUpdatePdpRequest2edl(o["update-pdp-request"], d, "update_pdp_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-pdp-request"], "ObjectFirewallGtpMessageRateLimit-UpdatePdpRequest"); ok {
			if err = d.Set("update_pdp_request", vv); err != nil {
				return fmt.Errorf("Error reading update_pdp_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_pdp_request: %v", err)
		}
	}

	if err = d.Set("update_pdp_response", flattenObjectFirewallGtpMessageRateLimitUpdatePdpResponse2edl(o["update-pdp-response"], d, "update_pdp_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-pdp-response"], "ObjectFirewallGtpMessageRateLimit-UpdatePdpResponse"); ok {
			if err = d.Set("update_pdp_response", vv); err != nil {
				return fmt.Errorf("Error reading update_pdp_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_pdp_response: %v", err)
		}
	}

	if err = d.Set("version_not_support", flattenObjectFirewallGtpMessageRateLimitVersionNotSupport2edl(o["version-not-support"], d, "version_not_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["version-not-support"], "ObjectFirewallGtpMessageRateLimit-VersionNotSupport"); ok {
			if err = d.Set("version_not_support", vv); err != nil {
				return fmt.Errorf("Error reading version_not_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version_not_support: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpMessageRateLimitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpMessageRateLimitCreateAaPdpRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitCreateAaPdpResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitCreateMbmsRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitCreateMbmsResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitCreatePdpRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitCreatePdpResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeleteAaPdpRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeleteAaPdpResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeleteMbmsRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeleteMbmsResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeletePdpRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeletePdpResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitEchoReponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitEchoRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitErrorIndication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFailureReportRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFailureReportResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdRelocCompleteAck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdRelocationComplete2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdRelocationRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdRelocationResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdSrnsContext2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdSrnsContextAck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitGPdu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitIdentificationRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitIdentificationResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsDeRegRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsDeRegResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsNotifyRejRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsNotifyRejResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsNotifyRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsNotifyResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsRegRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsRegResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsSesStartRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsSesStartResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsSesStopRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsSesStopResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitNoteMsRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitNoteMsResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitPduNotifyRejRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitPduNotifyRejResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitPduNotifyRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitPduNotifyResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitRanInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitRelocationCancelRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitRelocationCancelResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSendRouteRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSendRouteResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSgsnContextAck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSgsnContextRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSgsnContextResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSupportExtHdrNotify2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitUpdateMbmsRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitUpdateMbmsResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitUpdatePdpRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitUpdatePdpResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitVersionNotSupport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallGtpMessageRateLimit(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("create_aa_pdp_request"); ok || d.HasChange("create_aa_pdp_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitCreateAaPdpRequest2edl(d, v, "create_aa_pdp_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-aa-pdp-request"] = t
		}
	}

	if v, ok := d.GetOk("create_aa_pdp_response"); ok || d.HasChange("create_aa_pdp_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitCreateAaPdpResponse2edl(d, v, "create_aa_pdp_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-aa-pdp-response"] = t
		}
	}

	if v, ok := d.GetOk("create_mbms_request"); ok || d.HasChange("create_mbms_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitCreateMbmsRequest2edl(d, v, "create_mbms_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-mbms-request"] = t
		}
	}

	if v, ok := d.GetOk("create_mbms_response"); ok || d.HasChange("create_mbms_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitCreateMbmsResponse2edl(d, v, "create_mbms_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-mbms-response"] = t
		}
	}

	if v, ok := d.GetOk("create_pdp_request"); ok || d.HasChange("create_pdp_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitCreatePdpRequest2edl(d, v, "create_pdp_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-pdp-request"] = t
		}
	}

	if v, ok := d.GetOk("create_pdp_response"); ok || d.HasChange("create_pdp_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitCreatePdpResponse2edl(d, v, "create_pdp_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-pdp-response"] = t
		}
	}

	if v, ok := d.GetOk("delete_aa_pdp_request"); ok || d.HasChange("delete_aa_pdp_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitDeleteAaPdpRequest2edl(d, v, "delete_aa_pdp_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-aa-pdp-request"] = t
		}
	}

	if v, ok := d.GetOk("delete_aa_pdp_response"); ok || d.HasChange("delete_aa_pdp_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitDeleteAaPdpResponse2edl(d, v, "delete_aa_pdp_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-aa-pdp-response"] = t
		}
	}

	if v, ok := d.GetOk("delete_mbms_request"); ok || d.HasChange("delete_mbms_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitDeleteMbmsRequest2edl(d, v, "delete_mbms_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-mbms-request"] = t
		}
	}

	if v, ok := d.GetOk("delete_mbms_response"); ok || d.HasChange("delete_mbms_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitDeleteMbmsResponse2edl(d, v, "delete_mbms_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-mbms-response"] = t
		}
	}

	if v, ok := d.GetOk("delete_pdp_request"); ok || d.HasChange("delete_pdp_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitDeletePdpRequest2edl(d, v, "delete_pdp_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-pdp-request"] = t
		}
	}

	if v, ok := d.GetOk("delete_pdp_response"); ok || d.HasChange("delete_pdp_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitDeletePdpResponse2edl(d, v, "delete_pdp_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-pdp-response"] = t
		}
	}

	if v, ok := d.GetOk("echo_reponse"); ok || d.HasChange("echo_reponse") {
		t, err := expandObjectFirewallGtpMessageRateLimitEchoReponse2edl(d, v, "echo_reponse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["echo-reponse"] = t
		}
	}

	if v, ok := d.GetOk("echo_request"); ok || d.HasChange("echo_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitEchoRequest2edl(d, v, "echo_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["echo-request"] = t
		}
	}

	if v, ok := d.GetOk("error_indication"); ok || d.HasChange("error_indication") {
		t, err := expandObjectFirewallGtpMessageRateLimitErrorIndication2edl(d, v, "error_indication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["error-indication"] = t
		}
	}

	if v, ok := d.GetOk("failure_report_request"); ok || d.HasChange("failure_report_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitFailureReportRequest2edl(d, v, "failure_report_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failure-report-request"] = t
		}
	}

	if v, ok := d.GetOk("failure_report_response"); ok || d.HasChange("failure_report_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitFailureReportResponse2edl(d, v, "failure_report_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failure-report-response"] = t
		}
	}

	if v, ok := d.GetOk("fwd_reloc_complete_ack"); ok || d.HasChange("fwd_reloc_complete_ack") {
		t, err := expandObjectFirewallGtpMessageRateLimitFwdRelocCompleteAck2edl(d, v, "fwd_reloc_complete_ack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-reloc-complete-ack"] = t
		}
	}

	if v, ok := d.GetOk("fwd_relocation_complete"); ok || d.HasChange("fwd_relocation_complete") {
		t, err := expandObjectFirewallGtpMessageRateLimitFwdRelocationComplete2edl(d, v, "fwd_relocation_complete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-relocation-complete"] = t
		}
	}

	if v, ok := d.GetOk("fwd_relocation_request"); ok || d.HasChange("fwd_relocation_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitFwdRelocationRequest2edl(d, v, "fwd_relocation_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-relocation-request"] = t
		}
	}

	if v, ok := d.GetOk("fwd_relocation_response"); ok || d.HasChange("fwd_relocation_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitFwdRelocationResponse2edl(d, v, "fwd_relocation_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-relocation-response"] = t
		}
	}

	if v, ok := d.GetOk("fwd_srns_context"); ok || d.HasChange("fwd_srns_context") {
		t, err := expandObjectFirewallGtpMessageRateLimitFwdSrnsContext2edl(d, v, "fwd_srns_context")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-srns-context"] = t
		}
	}

	if v, ok := d.GetOk("fwd_srns_context_ack"); ok || d.HasChange("fwd_srns_context_ack") {
		t, err := expandObjectFirewallGtpMessageRateLimitFwdSrnsContextAck2edl(d, v, "fwd_srns_context_ack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-srns-context-ack"] = t
		}
	}

	if v, ok := d.GetOk("g_pdu"); ok || d.HasChange("g_pdu") {
		t, err := expandObjectFirewallGtpMessageRateLimitGPdu2edl(d, v, "g_pdu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["g-pdu"] = t
		}
	}

	if v, ok := d.GetOk("identification_request"); ok || d.HasChange("identification_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitIdentificationRequest2edl(d, v, "identification_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identification-request"] = t
		}
	}

	if v, ok := d.GetOk("identification_response"); ok || d.HasChange("identification_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitIdentificationResponse2edl(d, v, "identification_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identification-response"] = t
		}
	}

	if v, ok := d.GetOk("mbms_de_reg_request"); ok || d.HasChange("mbms_de_reg_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsDeRegRequest2edl(d, v, "mbms_de_reg_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-de-reg-request"] = t
		}
	}

	if v, ok := d.GetOk("mbms_de_reg_response"); ok || d.HasChange("mbms_de_reg_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsDeRegResponse2edl(d, v, "mbms_de_reg_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-de-reg-response"] = t
		}
	}

	if v, ok := d.GetOk("mbms_notify_rej_request"); ok || d.HasChange("mbms_notify_rej_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsNotifyRejRequest2edl(d, v, "mbms_notify_rej_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-notify-rej-request"] = t
		}
	}

	if v, ok := d.GetOk("mbms_notify_rej_response"); ok || d.HasChange("mbms_notify_rej_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsNotifyRejResponse2edl(d, v, "mbms_notify_rej_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-notify-rej-response"] = t
		}
	}

	if v, ok := d.GetOk("mbms_notify_request"); ok || d.HasChange("mbms_notify_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsNotifyRequest2edl(d, v, "mbms_notify_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-notify-request"] = t
		}
	}

	if v, ok := d.GetOk("mbms_notify_response"); ok || d.HasChange("mbms_notify_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsNotifyResponse2edl(d, v, "mbms_notify_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-notify-response"] = t
		}
	}

	if v, ok := d.GetOk("mbms_reg_request"); ok || d.HasChange("mbms_reg_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsRegRequest2edl(d, v, "mbms_reg_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-reg-request"] = t
		}
	}

	if v, ok := d.GetOk("mbms_reg_response"); ok || d.HasChange("mbms_reg_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsRegResponse2edl(d, v, "mbms_reg_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-reg-response"] = t
		}
	}

	if v, ok := d.GetOk("mbms_ses_start_request"); ok || d.HasChange("mbms_ses_start_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsSesStartRequest2edl(d, v, "mbms_ses_start_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-ses-start-request"] = t
		}
	}

	if v, ok := d.GetOk("mbms_ses_start_response"); ok || d.HasChange("mbms_ses_start_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsSesStartResponse2edl(d, v, "mbms_ses_start_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-ses-start-response"] = t
		}
	}

	if v, ok := d.GetOk("mbms_ses_stop_request"); ok || d.HasChange("mbms_ses_stop_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsSesStopRequest2edl(d, v, "mbms_ses_stop_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-ses-stop-request"] = t
		}
	}

	if v, ok := d.GetOk("mbms_ses_stop_response"); ok || d.HasChange("mbms_ses_stop_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitMbmsSesStopResponse2edl(d, v, "mbms_ses_stop_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-ses-stop-response"] = t
		}
	}

	if v, ok := d.GetOk("note_ms_request"); ok || d.HasChange("note_ms_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitNoteMsRequest2edl(d, v, "note_ms_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["note-ms-request"] = t
		}
	}

	if v, ok := d.GetOk("note_ms_response"); ok || d.HasChange("note_ms_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitNoteMsResponse2edl(d, v, "note_ms_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["note-ms-response"] = t
		}
	}

	if v, ok := d.GetOk("pdu_notify_rej_request"); ok || d.HasChange("pdu_notify_rej_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitPduNotifyRejRequest2edl(d, v, "pdu_notify_rej_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdu-notify-rej-request"] = t
		}
	}

	if v, ok := d.GetOk("pdu_notify_rej_response"); ok || d.HasChange("pdu_notify_rej_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitPduNotifyRejResponse2edl(d, v, "pdu_notify_rej_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdu-notify-rej-response"] = t
		}
	}

	if v, ok := d.GetOk("pdu_notify_request"); ok || d.HasChange("pdu_notify_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitPduNotifyRequest2edl(d, v, "pdu_notify_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdu-notify-request"] = t
		}
	}

	if v, ok := d.GetOk("pdu_notify_response"); ok || d.HasChange("pdu_notify_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitPduNotifyResponse2edl(d, v, "pdu_notify_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdu-notify-response"] = t
		}
	}

	if v, ok := d.GetOk("ran_info"); ok || d.HasChange("ran_info") {
		t, err := expandObjectFirewallGtpMessageRateLimitRanInfo2edl(d, v, "ran_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ran-info"] = t
		}
	}

	if v, ok := d.GetOk("relocation_cancel_request"); ok || d.HasChange("relocation_cancel_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitRelocationCancelRequest2edl(d, v, "relocation_cancel_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relocation-cancel-request"] = t
		}
	}

	if v, ok := d.GetOk("relocation_cancel_response"); ok || d.HasChange("relocation_cancel_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitRelocationCancelResponse2edl(d, v, "relocation_cancel_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relocation-cancel-response"] = t
		}
	}

	if v, ok := d.GetOk("send_route_request"); ok || d.HasChange("send_route_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitSendRouteRequest2edl(d, v, "send_route_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-route-request"] = t
		}
	}

	if v, ok := d.GetOk("send_route_response"); ok || d.HasChange("send_route_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitSendRouteResponse2edl(d, v, "send_route_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-route-response"] = t
		}
	}

	if v, ok := d.GetOk("sgsn_context_ack"); ok || d.HasChange("sgsn_context_ack") {
		t, err := expandObjectFirewallGtpMessageRateLimitSgsnContextAck2edl(d, v, "sgsn_context_ack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgsn-context-ack"] = t
		}
	}

	if v, ok := d.GetOk("sgsn_context_request"); ok || d.HasChange("sgsn_context_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitSgsnContextRequest2edl(d, v, "sgsn_context_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgsn-context-request"] = t
		}
	}

	if v, ok := d.GetOk("sgsn_context_response"); ok || d.HasChange("sgsn_context_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitSgsnContextResponse2edl(d, v, "sgsn_context_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgsn-context-response"] = t
		}
	}

	if v, ok := d.GetOk("support_ext_hdr_notify"); ok || d.HasChange("support_ext_hdr_notify") {
		t, err := expandObjectFirewallGtpMessageRateLimitSupportExtHdrNotify2edl(d, v, "support_ext_hdr_notify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["support-ext-hdr-notify"] = t
		}
	}

	if v, ok := d.GetOk("update_mbms_request"); ok || d.HasChange("update_mbms_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitUpdateMbmsRequest2edl(d, v, "update_mbms_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-mbms-request"] = t
		}
	}

	if v, ok := d.GetOk("update_mbms_response"); ok || d.HasChange("update_mbms_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitUpdateMbmsResponse2edl(d, v, "update_mbms_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-mbms-response"] = t
		}
	}

	if v, ok := d.GetOk("update_pdp_request"); ok || d.HasChange("update_pdp_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitUpdatePdpRequest2edl(d, v, "update_pdp_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-pdp-request"] = t
		}
	}

	if v, ok := d.GetOk("update_pdp_response"); ok || d.HasChange("update_pdp_response") {
		t, err := expandObjectFirewallGtpMessageRateLimitUpdatePdpResponse2edl(d, v, "update_pdp_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-pdp-response"] = t
		}
	}

	if v, ok := d.GetOk("version_not_support"); ok || d.HasChange("version_not_support") {
		t, err := expandObjectFirewallGtpMessageRateLimitVersionNotSupport2edl(d, v, "version_not_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version-not-support"] = t
		}
	}

	return &obj, nil
}
