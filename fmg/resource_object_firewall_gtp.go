// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure GTP.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallGtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpCreate,
		Read:   resourceObjectFirewallGtpRead,
		Update: resourceObjectFirewallGtpUpdate,
		Delete: resourceObjectFirewallGtpDelete,

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
			"addr_notify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apn": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"apnmember": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"selection_mode": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"apn_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorized_ggsns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorized_ggsns6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorized_sgsns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorized_sgsns6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"context_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"control_plane_message_rate_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"default_apn_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_imsi_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_ip_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_noip_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_policy_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"denied_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"echo_request_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"extension_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forwarded_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_tunnel_limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtp_in_gtp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtpu_denied_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtpu_forwarded_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtpu_log_freq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"half_close_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"half_open_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"handover_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"handover_group6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ie_remove_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"remove_ies": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sgsn_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sgsn_addr6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ie_remover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ie_validation": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apn_restriction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"charging_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"charging_gateway_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_user_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gsn_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"imei": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"imsi": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mm_context": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ms_tzone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ms_validated": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msisdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nsapi": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pdp_context": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"qos_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rat_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"reordering_required": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"selection_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uli": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ie_white_list_v0v1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ie_white_list_v2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"imsi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"apnmember": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mcc_mnc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msisdn_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"selection_mode": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"imsi_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_notify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalid_reserved_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalid_sgsns_to_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalid_sgsns6_to_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dstaddr6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcaddr6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"log_freq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_gtpu_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_imsi_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_msisdn_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_message_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"message_filter_v0v1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message_filter_v2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message_rate_limit": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"create_aa_pdp_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"create_aa_pdp_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"create_mbms_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"create_mbms_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"create_pdp_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"create_pdp_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"delete_aa_pdp_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"delete_aa_pdp_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"delete_mbms_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"delete_mbms_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"delete_pdp_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"delete_pdp_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"echo_reponse": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"echo_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"error_indication": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"failure_report_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"failure_report_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fwd_reloc_complete_ack": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fwd_relocation_complete": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fwd_relocation_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fwd_relocation_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fwd_srns_context": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fwd_srns_context_ack": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"g_pdu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"identification_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"identification_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_de_reg_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_de_reg_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_notify_rej_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_notify_rej_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_notify_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_notify_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_reg_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_reg_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_ses_start_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_ses_start_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_ses_stop_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mbms_ses_stop_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"note_ms_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"note_ms_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pdu_notify_rej_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pdu_notify_rej_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pdu_notify_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pdu_notify_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ran_info": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"relocation_cancel_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"relocation_cancel_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"send_route_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"send_route_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sgsn_context_ack": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sgsn_context_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sgsn_context_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"support_ext_hdr_notify": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"update_mbms_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"update_mbms_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"update_pdp_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"update_pdp_response": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"version_not_support": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"message_rate_limit_v0": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"create_pdp_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"delete_pdp_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"echo_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"message_rate_limit_v1": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"create_pdp_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"delete_pdp_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"echo_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"message_rate_limit_v2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"create_session_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"delete_session_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"echo_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"min_message_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"miss_must_ie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_mode": &schema.Schema{
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
			"noip_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"noip_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"start": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"out_of_state_ie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"out_of_state_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_apn_shaper": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rate_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"imei": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"imsi_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_apn_restriction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"rai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
					},
				},
			},
			"policy_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_v2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"imsi_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_apn_restriction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mei": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
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
				},
			},
			"port_notify": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rate_limit_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_limited_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_sampling_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"remove_if_echo_expires": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_if_recovery_differ": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserved_ie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_delete_when_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_delete_when_timeout_v2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spoof_src_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state_invalid_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sub_second_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sub_second_sampling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_count_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tunnel_limit_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"unknown_version_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_plane_message_rate_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"warning_threshold": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceObjectFirewallGtpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallGtp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtp resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallGtp(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtp resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallGtpRead(d, m)
}

func resourceObjectFirewallGtpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallGtp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtp resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtp(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallGtpRead(d, m)
}

func resourceObjectFirewallGtpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallGtp(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallGtp(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtp resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpAddrNotify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpApn(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallGtpApnAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Apn-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apnmember"
		if _, ok := i["apnmember"]; ok {
			v := flattenObjectFirewallGtpApnApnmember(i["apnmember"], d, pre_append)
			tmp["apnmember"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Apn-Apnmember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallGtpApnId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Apn-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "selection_mode"
		if _, ok := i["selection-mode"]; ok {
			v := flattenObjectFirewallGtpApnSelectionMode(i["selection-mode"], d, pre_append)
			tmp["selection_mode"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Apn-SelectionMode")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallGtpApnAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpApnApnmember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpApnId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpApnSelectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "ms",
			2: "net",
			4: "vrf",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpApnFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpAuthorizedGgsns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpAuthorizedGgsns6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpAuthorizedSgsns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpAuthorizedSgsns6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpContextId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpControlPlaneMessageRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpDefaultApnAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpDefaultImsiAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpDefaultIpAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpDefaultNoipAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpDefaultPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpDeniedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpEchoRequestInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpExtensionLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpForwardedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpGlobalTunnelLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpGtpInGtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpGtpuDeniedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpGtpuForwardedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpGtpuLogFreq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpHalfCloseTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpHalfOpenTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpHandoverGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpHandoverGroup6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeRemovePolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallGtpIeRemovePolicyId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-IeRemovePolicy-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_ies"
		if _, ok := i["remove-ies"]; ok {
			v := flattenObjectFirewallGtpIeRemovePolicyRemoveIes(i["remove-ies"], d, pre_append)
			tmp["remove_ies"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-IeRemovePolicy-RemoveIes")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sgsn_addr"
		if _, ok := i["sgsn-addr"]; ok {
			v := flattenObjectFirewallGtpIeRemovePolicySgsnAddr(i["sgsn-addr"], d, pre_append)
			tmp["sgsn_addr"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-IeRemovePolicy-SgsnAddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sgsn_addr6"
		if _, ok := i["sgsn-addr6"]; ok {
			v := flattenObjectFirewallGtpIeRemovePolicySgsnAddr6(i["sgsn-addr6"], d, pre_append)
			tmp["sgsn_addr6"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-IeRemovePolicy-SgsnAddr6")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallGtpIeRemovePolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeRemovePolicyRemoveIes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "apn-restriction",
			2:  "rat-type",
			4:  "rai",
			8:  "uli",
			16: "imei",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeRemovePolicySgsnAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeRemovePolicySgsnAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeRemover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidation(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "apn_restriction"
	if _, ok := i["apn-restriction"]; ok {
		result["apn_restriction"] = flattenObjectFirewallGtpIeValidationApnRestriction(i["apn-restriction"], d, pre_append)
	}

	pre_append = pre + ".0." + "charging_ID"
	if _, ok := i["charging-ID"]; ok {
		result["charging_ID"] = flattenObjectFirewallGtpIeValidationChargingId(i["charging-ID"], d, pre_append)
	}

	pre_append = pre + ".0." + "charging_gateway_addr"
	if _, ok := i["charging-gateway-addr"]; ok {
		result["charging_gateway_addr"] = flattenObjectFirewallGtpIeValidationChargingGatewayAddr(i["charging-gateway-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "end_user_addr"
	if _, ok := i["end-user-addr"]; ok {
		result["end_user_addr"] = flattenObjectFirewallGtpIeValidationEndUserAddr(i["end-user-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "gsn_addr"
	if _, ok := i["gsn-addr"]; ok {
		result["gsn_addr"] = flattenObjectFirewallGtpIeValidationGsnAddr(i["gsn-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "imei"
	if _, ok := i["imei"]; ok {
		result["imei"] = flattenObjectFirewallGtpIeValidationImei(i["imei"], d, pre_append)
	}

	pre_append = pre + ".0." + "imsi"
	if _, ok := i["imsi"]; ok {
		result["imsi"] = flattenObjectFirewallGtpIeValidationImsi(i["imsi"], d, pre_append)
	}

	pre_append = pre + ".0." + "mm_context"
	if _, ok := i["mm-context"]; ok {
		result["mm_context"] = flattenObjectFirewallGtpIeValidationMmContext(i["mm-context"], d, pre_append)
	}

	pre_append = pre + ".0." + "ms_tzone"
	if _, ok := i["ms-tzone"]; ok {
		result["ms_tzone"] = flattenObjectFirewallGtpIeValidationMsTzone(i["ms-tzone"], d, pre_append)
	}

	pre_append = pre + ".0." + "ms_validated"
	if _, ok := i["ms-validated"]; ok {
		result["ms_validated"] = flattenObjectFirewallGtpIeValidationMsValidated(i["ms-validated"], d, pre_append)
	}

	pre_append = pre + ".0." + "msisdn"
	if _, ok := i["msisdn"]; ok {
		result["msisdn"] = flattenObjectFirewallGtpIeValidationMsisdn(i["msisdn"], d, pre_append)
	}

	pre_append = pre + ".0." + "nsapi"
	if _, ok := i["nsapi"]; ok {
		result["nsapi"] = flattenObjectFirewallGtpIeValidationNsapi(i["nsapi"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdp_context"
	if _, ok := i["pdp-context"]; ok {
		result["pdp_context"] = flattenObjectFirewallGtpIeValidationPdpContext(i["pdp-context"], d, pre_append)
	}

	pre_append = pre + ".0." + "qos_profile"
	if _, ok := i["qos-profile"]; ok {
		result["qos_profile"] = flattenObjectFirewallGtpIeValidationQosProfile(i["qos-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "rai"
	if _, ok := i["rai"]; ok {
		result["rai"] = flattenObjectFirewallGtpIeValidationRai(i["rai"], d, pre_append)
	}

	pre_append = pre + ".0." + "rat_type"
	if _, ok := i["rat-type"]; ok {
		result["rat_type"] = flattenObjectFirewallGtpIeValidationRatType(i["rat-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "reordering_required"
	if _, ok := i["reordering-required"]; ok {
		result["reordering_required"] = flattenObjectFirewallGtpIeValidationReorderingRequired(i["reordering-required"], d, pre_append)
	}

	pre_append = pre + ".0." + "selection_mode"
	if _, ok := i["selection-mode"]; ok {
		result["selection_mode"] = flattenObjectFirewallGtpIeValidationSelectionMode(i["selection-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "uli"
	if _, ok := i["uli"]; ok {
		result["uli"] = flattenObjectFirewallGtpIeValidationUli(i["uli"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallGtpIeValidationApnRestriction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationChargingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationChargingGatewayAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationEndUserAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationGsnAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationImei(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationImsi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationMmContext(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationMsTzone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationMsValidated(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationMsisdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationNsapi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationPdpContext(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationQosProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationRai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationRatType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationReorderingRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationSelectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeValidationUli(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIeWhiteListV0V1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeWhiteListV2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpImsi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallGtpImsiAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Imsi-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apnmember"
		if _, ok := i["apnmember"]; ok {
			v := flattenObjectFirewallGtpImsiApnmember(i["apnmember"], d, pre_append)
			tmp["apnmember"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Imsi-Apnmember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallGtpImsiId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Imsi-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcc_mnc"
		if _, ok := i["mcc-mnc"]; ok {
			v := flattenObjectFirewallGtpImsiMccMnc(i["mcc-mnc"], d, pre_append)
			tmp["mcc_mnc"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Imsi-MccMnc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msisdn_prefix"
		if _, ok := i["msisdn-prefix"]; ok {
			v := flattenObjectFirewallGtpImsiMsisdnPrefix(i["msisdn-prefix"], d, pre_append)
			tmp["msisdn_prefix"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Imsi-MsisdnPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "selection_mode"
		if _, ok := i["selection-mode"]; ok {
			v := flattenObjectFirewallGtpImsiSelectionMode(i["selection-mode"], d, pre_append)
			tmp["selection_mode"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Imsi-SelectionMode")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallGtpImsiAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpImsiApnmember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpImsiId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpImsiMccMnc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpImsiMsisdnPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpImsiSelectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "ms",
			2: "net",
			4: "vrf",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpImsiFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpInterfaceNotify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpInvalidReservedField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpInvalidSgsnsToLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpInvalidSgsns6ToLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIpFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIpPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallGtpIpPolicyAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-IpPolicy-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenObjectFirewallGtpIpPolicyDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-IpPolicy-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := i["dstaddr6"]; ok {
			v := flattenObjectFirewallGtpIpPolicyDstaddr6(i["dstaddr6"], d, pre_append)
			tmp["dstaddr6"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-IpPolicy-Dstaddr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallGtpIpPolicyId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-IpPolicy-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenObjectFirewallGtpIpPolicySrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-IpPolicy-Srcaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := i["srcaddr6"]; ok {
			v := flattenObjectFirewallGtpIpPolicySrcaddr6(i["srcaddr6"], d, pre_append)
			tmp["srcaddr6"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-IpPolicy-Srcaddr6")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallGtpIpPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpIpPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIpPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIpPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIpPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIpPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpLogFreq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpLogGtpuLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpLogImsiPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpLogMsisdnPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMaxMessageLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageFilterV0V1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageFilterV2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimit(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "create_aa_pdp_request"
	if _, ok := i["create-aa-pdp-request"]; ok {
		result["create_aa_pdp_request"] = flattenObjectFirewallGtpMessageRateLimitCreateAaPdpRequest(i["create-aa-pdp-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "create_aa_pdp_response"
	if _, ok := i["create-aa-pdp-response"]; ok {
		result["create_aa_pdp_response"] = flattenObjectFirewallGtpMessageRateLimitCreateAaPdpResponse(i["create-aa-pdp-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "create_mbms_request"
	if _, ok := i["create-mbms-request"]; ok {
		result["create_mbms_request"] = flattenObjectFirewallGtpMessageRateLimitCreateMbmsRequest(i["create-mbms-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "create_mbms_response"
	if _, ok := i["create-mbms-response"]; ok {
		result["create_mbms_response"] = flattenObjectFirewallGtpMessageRateLimitCreateMbmsResponse(i["create-mbms-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "create_pdp_request"
	if _, ok := i["create-pdp-request"]; ok {
		result["create_pdp_request"] = flattenObjectFirewallGtpMessageRateLimitCreatePdpRequest(i["create-pdp-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "create_pdp_response"
	if _, ok := i["create-pdp-response"]; ok {
		result["create_pdp_response"] = flattenObjectFirewallGtpMessageRateLimitCreatePdpResponse(i["create-pdp-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "delete_aa_pdp_request"
	if _, ok := i["delete-aa-pdp-request"]; ok {
		result["delete_aa_pdp_request"] = flattenObjectFirewallGtpMessageRateLimitDeleteAaPdpRequest(i["delete-aa-pdp-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "delete_aa_pdp_response"
	if _, ok := i["delete-aa-pdp-response"]; ok {
		result["delete_aa_pdp_response"] = flattenObjectFirewallGtpMessageRateLimitDeleteAaPdpResponse(i["delete-aa-pdp-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "delete_mbms_request"
	if _, ok := i["delete-mbms-request"]; ok {
		result["delete_mbms_request"] = flattenObjectFirewallGtpMessageRateLimitDeleteMbmsRequest(i["delete-mbms-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "delete_mbms_response"
	if _, ok := i["delete-mbms-response"]; ok {
		result["delete_mbms_response"] = flattenObjectFirewallGtpMessageRateLimitDeleteMbmsResponse(i["delete-mbms-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "delete_pdp_request"
	if _, ok := i["delete-pdp-request"]; ok {
		result["delete_pdp_request"] = flattenObjectFirewallGtpMessageRateLimitDeletePdpRequest(i["delete-pdp-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "delete_pdp_response"
	if _, ok := i["delete-pdp-response"]; ok {
		result["delete_pdp_response"] = flattenObjectFirewallGtpMessageRateLimitDeletePdpResponse(i["delete-pdp-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "echo_reponse"
	if _, ok := i["echo-reponse"]; ok {
		result["echo_reponse"] = flattenObjectFirewallGtpMessageRateLimitEchoReponse(i["echo-reponse"], d, pre_append)
	}

	pre_append = pre + ".0." + "echo_request"
	if _, ok := i["echo-request"]; ok {
		result["echo_request"] = flattenObjectFirewallGtpMessageRateLimitEchoRequest(i["echo-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "error_indication"
	if _, ok := i["error-indication"]; ok {
		result["error_indication"] = flattenObjectFirewallGtpMessageRateLimitErrorIndication(i["error-indication"], d, pre_append)
	}

	pre_append = pre + ".0." + "failure_report_request"
	if _, ok := i["failure-report-request"]; ok {
		result["failure_report_request"] = flattenObjectFirewallGtpMessageRateLimitFailureReportRequest(i["failure-report-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "failure_report_response"
	if _, ok := i["failure-report-response"]; ok {
		result["failure_report_response"] = flattenObjectFirewallGtpMessageRateLimitFailureReportResponse(i["failure-report-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_reloc_complete_ack"
	if _, ok := i["fwd-reloc-complete-ack"]; ok {
		result["fwd_reloc_complete_ack"] = flattenObjectFirewallGtpMessageRateLimitFwdRelocCompleteAck(i["fwd-reloc-complete-ack"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_relocation_complete"
	if _, ok := i["fwd-relocation-complete"]; ok {
		result["fwd_relocation_complete"] = flattenObjectFirewallGtpMessageRateLimitFwdRelocationComplete(i["fwd-relocation-complete"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_relocation_request"
	if _, ok := i["fwd-relocation-request"]; ok {
		result["fwd_relocation_request"] = flattenObjectFirewallGtpMessageRateLimitFwdRelocationRequest(i["fwd-relocation-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_relocation_response"
	if _, ok := i["fwd-relocation-response"]; ok {
		result["fwd_relocation_response"] = flattenObjectFirewallGtpMessageRateLimitFwdRelocationResponse(i["fwd-relocation-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_srns_context"
	if _, ok := i["fwd-srns-context"]; ok {
		result["fwd_srns_context"] = flattenObjectFirewallGtpMessageRateLimitFwdSrnsContext(i["fwd-srns-context"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_srns_context_ack"
	if _, ok := i["fwd-srns-context-ack"]; ok {
		result["fwd_srns_context_ack"] = flattenObjectFirewallGtpMessageRateLimitFwdSrnsContextAck(i["fwd-srns-context-ack"], d, pre_append)
	}

	pre_append = pre + ".0." + "g_pdu"
	if _, ok := i["g-pdu"]; ok {
		result["g_pdu"] = flattenObjectFirewallGtpMessageRateLimitGPdu(i["g-pdu"], d, pre_append)
	}

	pre_append = pre + ".0." + "identification_request"
	if _, ok := i["identification-request"]; ok {
		result["identification_request"] = flattenObjectFirewallGtpMessageRateLimitIdentificationRequest(i["identification-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "identification_response"
	if _, ok := i["identification-response"]; ok {
		result["identification_response"] = flattenObjectFirewallGtpMessageRateLimitIdentificationResponse(i["identification-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_de_reg_request"
	if _, ok := i["mbms-de-reg-request"]; ok {
		result["mbms_de_reg_request"] = flattenObjectFirewallGtpMessageRateLimitMbmsDeRegRequest(i["mbms-de-reg-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_de_reg_response"
	if _, ok := i["mbms-de-reg-response"]; ok {
		result["mbms_de_reg_response"] = flattenObjectFirewallGtpMessageRateLimitMbmsDeRegResponse(i["mbms-de-reg-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_notify_rej_request"
	if _, ok := i["mbms-notify-rej-request"]; ok {
		result["mbms_notify_rej_request"] = flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRejRequest(i["mbms-notify-rej-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_notify_rej_response"
	if _, ok := i["mbms-notify-rej-response"]; ok {
		result["mbms_notify_rej_response"] = flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRejResponse(i["mbms-notify-rej-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_notify_request"
	if _, ok := i["mbms-notify-request"]; ok {
		result["mbms_notify_request"] = flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRequest(i["mbms-notify-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_notify_response"
	if _, ok := i["mbms-notify-response"]; ok {
		result["mbms_notify_response"] = flattenObjectFirewallGtpMessageRateLimitMbmsNotifyResponse(i["mbms-notify-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_reg_request"
	if _, ok := i["mbms-reg-request"]; ok {
		result["mbms_reg_request"] = flattenObjectFirewallGtpMessageRateLimitMbmsRegRequest(i["mbms-reg-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_reg_response"
	if _, ok := i["mbms-reg-response"]; ok {
		result["mbms_reg_response"] = flattenObjectFirewallGtpMessageRateLimitMbmsRegResponse(i["mbms-reg-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_ses_start_request"
	if _, ok := i["mbms-ses-start-request"]; ok {
		result["mbms_ses_start_request"] = flattenObjectFirewallGtpMessageRateLimitMbmsSesStartRequest(i["mbms-ses-start-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_ses_start_response"
	if _, ok := i["mbms-ses-start-response"]; ok {
		result["mbms_ses_start_response"] = flattenObjectFirewallGtpMessageRateLimitMbmsSesStartResponse(i["mbms-ses-start-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_ses_stop_request"
	if _, ok := i["mbms-ses-stop-request"]; ok {
		result["mbms_ses_stop_request"] = flattenObjectFirewallGtpMessageRateLimitMbmsSesStopRequest(i["mbms-ses-stop-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "mbms_ses_stop_response"
	if _, ok := i["mbms-ses-stop-response"]; ok {
		result["mbms_ses_stop_response"] = flattenObjectFirewallGtpMessageRateLimitMbmsSesStopResponse(i["mbms-ses-stop-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "note_ms_request"
	if _, ok := i["note-ms-request"]; ok {
		result["note_ms_request"] = flattenObjectFirewallGtpMessageRateLimitNoteMsRequest(i["note-ms-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "note_ms_response"
	if _, ok := i["note-ms-response"]; ok {
		result["note_ms_response"] = flattenObjectFirewallGtpMessageRateLimitNoteMsResponse(i["note-ms-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdu_notify_rej_request"
	if _, ok := i["pdu-notify-rej-request"]; ok {
		result["pdu_notify_rej_request"] = flattenObjectFirewallGtpMessageRateLimitPduNotifyRejRequest(i["pdu-notify-rej-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdu_notify_rej_response"
	if _, ok := i["pdu-notify-rej-response"]; ok {
		result["pdu_notify_rej_response"] = flattenObjectFirewallGtpMessageRateLimitPduNotifyRejResponse(i["pdu-notify-rej-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdu_notify_request"
	if _, ok := i["pdu-notify-request"]; ok {
		result["pdu_notify_request"] = flattenObjectFirewallGtpMessageRateLimitPduNotifyRequest(i["pdu-notify-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdu_notify_response"
	if _, ok := i["pdu-notify-response"]; ok {
		result["pdu_notify_response"] = flattenObjectFirewallGtpMessageRateLimitPduNotifyResponse(i["pdu-notify-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "ran_info"
	if _, ok := i["ran-info"]; ok {
		result["ran_info"] = flattenObjectFirewallGtpMessageRateLimitRanInfo(i["ran-info"], d, pre_append)
	}

	pre_append = pre + ".0." + "relocation_cancel_request"
	if _, ok := i["relocation-cancel-request"]; ok {
		result["relocation_cancel_request"] = flattenObjectFirewallGtpMessageRateLimitRelocationCancelRequest(i["relocation-cancel-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "relocation_cancel_response"
	if _, ok := i["relocation-cancel-response"]; ok {
		result["relocation_cancel_response"] = flattenObjectFirewallGtpMessageRateLimitRelocationCancelResponse(i["relocation-cancel-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "send_route_request"
	if _, ok := i["send-route-request"]; ok {
		result["send_route_request"] = flattenObjectFirewallGtpMessageRateLimitSendRouteRequest(i["send-route-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "send_route_response"
	if _, ok := i["send-route-response"]; ok {
		result["send_route_response"] = flattenObjectFirewallGtpMessageRateLimitSendRouteResponse(i["send-route-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "sgsn_context_ack"
	if _, ok := i["sgsn-context-ack"]; ok {
		result["sgsn_context_ack"] = flattenObjectFirewallGtpMessageRateLimitSgsnContextAck(i["sgsn-context-ack"], d, pre_append)
	}

	pre_append = pre + ".0." + "sgsn_context_request"
	if _, ok := i["sgsn-context-request"]; ok {
		result["sgsn_context_request"] = flattenObjectFirewallGtpMessageRateLimitSgsnContextRequest(i["sgsn-context-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "sgsn_context_response"
	if _, ok := i["sgsn-context-response"]; ok {
		result["sgsn_context_response"] = flattenObjectFirewallGtpMessageRateLimitSgsnContextResponse(i["sgsn-context-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "support_ext_hdr_notify"
	if _, ok := i["support-ext-hdr-notify"]; ok {
		result["support_ext_hdr_notify"] = flattenObjectFirewallGtpMessageRateLimitSupportExtHdrNotify(i["support-ext-hdr-notify"], d, pre_append)
	}

	pre_append = pre + ".0." + "update_mbms_request"
	if _, ok := i["update-mbms-request"]; ok {
		result["update_mbms_request"] = flattenObjectFirewallGtpMessageRateLimitUpdateMbmsRequest(i["update-mbms-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "update_mbms_response"
	if _, ok := i["update-mbms-response"]; ok {
		result["update_mbms_response"] = flattenObjectFirewallGtpMessageRateLimitUpdateMbmsResponse(i["update-mbms-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "update_pdp_request"
	if _, ok := i["update-pdp-request"]; ok {
		result["update_pdp_request"] = flattenObjectFirewallGtpMessageRateLimitUpdatePdpRequest(i["update-pdp-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "update_pdp_response"
	if _, ok := i["update-pdp-response"]; ok {
		result["update_pdp_response"] = flattenObjectFirewallGtpMessageRateLimitUpdatePdpResponse(i["update-pdp-response"], d, pre_append)
	}

	pre_append = pre + ".0." + "version_not_support"
	if _, ok := i["version-not-support"]; ok {
		result["version_not_support"] = flattenObjectFirewallGtpMessageRateLimitVersionNotSupport(i["version-not-support"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallGtpMessageRateLimitCreateAaPdpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitCreateAaPdpResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitCreateMbmsRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitCreateMbmsResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitCreatePdpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitCreatePdpResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeleteAaPdpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeleteAaPdpResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeleteMbmsRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeleteMbmsResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeletePdpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitDeletePdpResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitEchoReponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitEchoRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitErrorIndication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFailureReportRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFailureReportResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdRelocCompleteAck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdRelocationComplete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdRelocationRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdRelocationResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdSrnsContext(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitFwdSrnsContextAck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitGPdu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitIdentificationRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitIdentificationResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsDeRegRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsDeRegResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRejRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRejResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsNotifyRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsNotifyResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsRegRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsRegResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsSesStartRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsSesStartResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsSesStopRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitMbmsSesStopResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitNoteMsRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitNoteMsResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitPduNotifyRejRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitPduNotifyRejResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitPduNotifyRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitPduNotifyResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitRanInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitRelocationCancelRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitRelocationCancelResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSendRouteRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSendRouteResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSgsnContextAck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSgsnContextRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSgsnContextResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitSupportExtHdrNotify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitUpdateMbmsRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitUpdateMbmsResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitUpdatePdpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitUpdatePdpResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitVersionNotSupport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV0(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "create_pdp_request"
	if _, ok := i["create-pdp-request"]; ok {
		result["create_pdp_request"] = flattenObjectFirewallGtpMessageRateLimitV0CreatePdpRequest(i["create-pdp-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "delete_pdp_request"
	if _, ok := i["delete-pdp-request"]; ok {
		result["delete_pdp_request"] = flattenObjectFirewallGtpMessageRateLimitV0DeletePdpRequest(i["delete-pdp-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "echo_request"
	if _, ok := i["echo-request"]; ok {
		result["echo_request"] = flattenObjectFirewallGtpMessageRateLimitV0EchoRequest(i["echo-request"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallGtpMessageRateLimitV0CreatePdpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV0DeletePdpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV0EchoRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "create_pdp_request"
	if _, ok := i["create-pdp-request"]; ok {
		result["create_pdp_request"] = flattenObjectFirewallGtpMessageRateLimitV1CreatePdpRequest(i["create-pdp-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "delete_pdp_request"
	if _, ok := i["delete-pdp-request"]; ok {
		result["delete_pdp_request"] = flattenObjectFirewallGtpMessageRateLimitV1DeletePdpRequest(i["delete-pdp-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "echo_request"
	if _, ok := i["echo-request"]; ok {
		result["echo_request"] = flattenObjectFirewallGtpMessageRateLimitV1EchoRequest(i["echo-request"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallGtpMessageRateLimitV1CreatePdpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV1DeletePdpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV1EchoRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "create_session_request"
	if _, ok := i["create-session-request"]; ok {
		result["create_session_request"] = flattenObjectFirewallGtpMessageRateLimitV2CreateSessionRequest(i["create-session-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "delete_session_request"
	if _, ok := i["delete-session-request"]; ok {
		result["delete_session_request"] = flattenObjectFirewallGtpMessageRateLimitV2DeleteSessionRequest(i["delete-session-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "echo_request"
	if _, ok := i["echo-request"]; ok {
		result["echo_request"] = flattenObjectFirewallGtpMessageRateLimitV2EchoRequest(i["echo-request"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallGtpMessageRateLimitV2CreateSessionRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV2DeleteSessionRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV2EchoRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMinMessageLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMissMustIe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpMonitorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "disable",
			1:  "enable",
			43: "vdom",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpNoipFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpNoipPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallGtpNoipPolicyAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-NoipPolicy-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end"
		if _, ok := i["end"]; ok {
			v := flattenObjectFirewallGtpNoipPolicyEnd(i["end"], d, pre_append)
			tmp["end"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-NoipPolicy-End")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallGtpNoipPolicyId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-NoipPolicy-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start"
		if _, ok := i["start"]; ok {
			v := flattenObjectFirewallGtpNoipPolicyStart(i["start"], d, pre_append)
			tmp["start"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-NoipPolicy-Start")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallGtpNoipPolicyType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-NoipPolicy-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallGtpNoipPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpNoipPolicyEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpNoipPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpNoipPolicyStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpNoipPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "etsi",
			1: "ietf",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpOutOfStateIe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpOutOfStateMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPerApnShaper(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apn"
		if _, ok := i["apn"]; ok {
			v := flattenObjectFirewallGtpPerApnShaperApn(i["apn"], d, pre_append)
			tmp["apn"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PerApnShaper-Apn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallGtpPerApnShaperId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PerApnShaper-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_limit"
		if _, ok := i["rate-limit"]; ok {
			v := flattenObjectFirewallGtpPerApnShaperRateLimit(i["rate-limit"], d, pre_append)
			tmp["rate_limit"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PerApnShaper-RateLimit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenObjectFirewallGtpPerApnShaperVersion(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PerApnShaper-Version")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallGtpPerApnShaperApn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPerApnShaperId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPerApnShaperRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPerApnShaperVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallGtpPolicyAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apn_sel_mode"
		if _, ok := i["apn-sel-mode"]; ok {
			v := flattenObjectFirewallGtpPolicyApnSelMode(i["apn-sel-mode"], d, pre_append)
			tmp["apn_sel_mode"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-ApnSelMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apnmember"
		if _, ok := i["apnmember"]; ok {
			v := flattenObjectFirewallGtpPolicyApnmember(i["apnmember"], d, pre_append)
			tmp["apnmember"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-Apnmember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallGtpPolicyId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imei"
		if _, ok := i["imei"]; ok {
			v := flattenObjectFirewallGtpPolicyImei(i["imei"], d, pre_append)
			tmp["imei"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-Imei")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imsi_prefix"
		if _, ok := i["imsi-prefix"]; ok {
			v := flattenObjectFirewallGtpPolicyImsiPrefix(i["imsi-prefix"], d, pre_append)
			tmp["imsi_prefix"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-ImsiPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_apn_restriction"
		if _, ok := i["max-apn-restriction"]; ok {
			v := flattenObjectFirewallGtpPolicyMaxApnRestriction(i["max-apn-restriction"], d, pre_append)
			tmp["max_apn_restriction"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-MaxApnRestriction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "messages"
		if _, ok := i["messages"]; ok {
			v := flattenObjectFirewallGtpPolicyMessages(i["messages"], d, pre_append)
			tmp["messages"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-Messages")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msisdn_prefix"
		if _, ok := i["msisdn-prefix"]; ok {
			v := flattenObjectFirewallGtpPolicyMsisdnPrefix(i["msisdn-prefix"], d, pre_append)
			tmp["msisdn_prefix"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-MsisdnPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rai"
		if _, ok := i["rai"]; ok {
			v := flattenObjectFirewallGtpPolicyRai(i["rai"], d, pre_append)
			tmp["rai"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-Rai")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rat_type"
		if _, ok := i["rat-type"]; ok {
			v := flattenObjectFirewallGtpPolicyRatType(i["rat-type"], d, pre_append)
			tmp["rat_type"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-RatType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uli"
		if _, ok := i["uli"]; ok {
			v := flattenObjectFirewallGtpPolicyUli(i["uli"], d, pre_append)
			tmp["uli"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-Policy-Uli")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallGtpPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyApnSelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "ms",
			2: "net",
			4: "vrf",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyApnmember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyImei(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyImsiPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyMaxApnRestriction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "all",
			1: "public-1",
			2: "public-2",
			3: "private-1",
			4: "private-2",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyMessages(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "create-req",
			2: "create-res",
			4: "update-req",
			8: "update-res",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyMsisdnPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyRai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyRatType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:   "any",
			2:   "utran",
			4:   "geran",
			8:   "wlan",
			16:  "gan",
			32:  "hspa",
			64:  "eutran",
			128: "virtual",
			256: "nbiot",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyUli(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyV2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallGtpPolicyV2Action(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apn_sel_mode"
		if _, ok := i["apn-sel-mode"]; ok {
			v := flattenObjectFirewallGtpPolicyV2ApnSelMode(i["apn-sel-mode"], d, pre_append)
			tmp["apn_sel_mode"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-ApnSelMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apnmember"
		if _, ok := i["apnmember"]; ok {
			v := flattenObjectFirewallGtpPolicyV2Apnmember(i["apnmember"], d, pre_append)
			tmp["apnmember"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-Apnmember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallGtpPolicyV2Id(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imsi_prefix"
		if _, ok := i["imsi-prefix"]; ok {
			v := flattenObjectFirewallGtpPolicyV2ImsiPrefix(i["imsi-prefix"], d, pre_append)
			tmp["imsi_prefix"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-ImsiPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_apn_restriction"
		if _, ok := i["max-apn-restriction"]; ok {
			v := flattenObjectFirewallGtpPolicyV2MaxApnRestriction(i["max-apn-restriction"], d, pre_append)
			tmp["max_apn_restriction"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-MaxApnRestriction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mei"
		if _, ok := i["mei"]; ok {
			v := flattenObjectFirewallGtpPolicyV2Mei(i["mei"], d, pre_append)
			tmp["mei"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-Mei")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "messages"
		if _, ok := i["messages"]; ok {
			v := flattenObjectFirewallGtpPolicyV2Messages(i["messages"], d, pre_append)
			tmp["messages"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-Messages")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msisdn_prefix"
		if _, ok := i["msisdn-prefix"]; ok {
			v := flattenObjectFirewallGtpPolicyV2MsisdnPrefix(i["msisdn-prefix"], d, pre_append)
			tmp["msisdn_prefix"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-MsisdnPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rat_type"
		if _, ok := i["rat-type"]; ok {
			v := flattenObjectFirewallGtpPolicyV2RatType(i["rat-type"], d, pre_append)
			tmp["rat_type"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-RatType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uli"
		if _, ok := i["uli"]; ok {
			v := flattenObjectFirewallGtpPolicyV2Uli(i["uli"], d, pre_append)
			tmp["uli"] = fortiAPISubPartPatch(v, "ObjectFirewallGtp-PolicyV2-Uli")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallGtpPolicyV2Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "deny",
			1: "allow",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyV2ApnSelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "ms",
			2: "net",
			4: "vrf",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyV2Apnmember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2ImsiPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2MaxApnRestriction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "all",
			1: "public-1",
			2: "public-2",
			3: "private-1",
			4: "private-2",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyV2Mei(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2Messages(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "create-ses-req",
			2: "create-ses-res",
			4: "modify-bearer-req",
			8: "modify-bearer-res",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyV2MsisdnPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPolicyV2RatType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:    "any",
			2:    "utran",
			4:    "geran",
			8:    "wlan",
			16:   "gan",
			32:   "hspa",
			64:   "eutran",
			128:  "virtual",
			256:  "nbiot",
			512:  "ltem",
			1024: "nr",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpPolicyV2Uli(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallGtpPortNotify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpRateLimitMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "per-profile",
			1: "per-stream",
			2: "per-apn",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpRateLimitedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpRateSamplingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpRemoveIfEchoExpires(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpRemoveIfRecoveryDiffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpReservedIe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpSendDeleteWhenTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpSendDeleteWhenTimeoutV2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpSpoofSrcAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpStateInvalidLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpSubSecondInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			100: "0.1",
			250: "0.25",
			500: "0.5",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpSubSecondSampling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpTrafficCountLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpTunnelLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpTunnelLimitLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpTunnelTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpUnknownVersionAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallGtpUserPlaneMessageRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpWarningThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallGtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_notify", flattenObjectFirewallGtpAddrNotify(o["addr-notify"], d, "addr_notify")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-notify"], "ObjectFirewallGtp-AddrNotify"); ok {
			if err = d.Set("addr_notify", vv); err != nil {
				return fmt.Errorf("Error reading addr_notify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_notify: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("apn", flattenObjectFirewallGtpApn(o["apn"], d, "apn")); err != nil {
			if vv, ok := fortiAPIPatch(o["apn"], "ObjectFirewallGtp-Apn"); ok {
				if err = d.Set("apn", vv); err != nil {
					return fmt.Errorf("Error reading apn: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading apn: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("apn"); ok {
			if err = d.Set("apn", flattenObjectFirewallGtpApn(o["apn"], d, "apn")); err != nil {
				if vv, ok := fortiAPIPatch(o["apn"], "ObjectFirewallGtp-Apn"); ok {
					if err = d.Set("apn", vv); err != nil {
						return fmt.Errorf("Error reading apn: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading apn: %v", err)
				}
			}
		}
	}

	if err = d.Set("apn_filter", flattenObjectFirewallGtpApnFilter(o["apn-filter"], d, "apn_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["apn-filter"], "ObjectFirewallGtp-ApnFilter"); ok {
			if err = d.Set("apn_filter", vv); err != nil {
				return fmt.Errorf("Error reading apn_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apn_filter: %v", err)
		}
	}

	if err = d.Set("authorized_ggsns", flattenObjectFirewallGtpAuthorizedGgsns(o["authorized-ggsns"], d, "authorized_ggsns")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorized-ggsns"], "ObjectFirewallGtp-AuthorizedGgsns"); ok {
			if err = d.Set("authorized_ggsns", vv); err != nil {
				return fmt.Errorf("Error reading authorized_ggsns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorized_ggsns: %v", err)
		}
	}

	if err = d.Set("authorized_ggsns6", flattenObjectFirewallGtpAuthorizedGgsns6(o["authorized-ggsns6"], d, "authorized_ggsns6")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorized-ggsns6"], "ObjectFirewallGtp-AuthorizedGgsns6"); ok {
			if err = d.Set("authorized_ggsns6", vv); err != nil {
				return fmt.Errorf("Error reading authorized_ggsns6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorized_ggsns6: %v", err)
		}
	}

	if err = d.Set("authorized_sgsns", flattenObjectFirewallGtpAuthorizedSgsns(o["authorized-sgsns"], d, "authorized_sgsns")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorized-sgsns"], "ObjectFirewallGtp-AuthorizedSgsns"); ok {
			if err = d.Set("authorized_sgsns", vv); err != nil {
				return fmt.Errorf("Error reading authorized_sgsns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorized_sgsns: %v", err)
		}
	}

	if err = d.Set("authorized_sgsns6", flattenObjectFirewallGtpAuthorizedSgsns6(o["authorized-sgsns6"], d, "authorized_sgsns6")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorized-sgsns6"], "ObjectFirewallGtp-AuthorizedSgsns6"); ok {
			if err = d.Set("authorized_sgsns6", vv); err != nil {
				return fmt.Errorf("Error reading authorized_sgsns6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorized_sgsns6: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallGtpComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallGtp-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("context_id", flattenObjectFirewallGtpContextId(o["context-id"], d, "context_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["context-id"], "ObjectFirewallGtp-ContextId"); ok {
			if err = d.Set("context_id", vv); err != nil {
				return fmt.Errorf("Error reading context_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading context_id: %v", err)
		}
	}

	if err = d.Set("control_plane_message_rate_limit", flattenObjectFirewallGtpControlPlaneMessageRateLimit(o["control-plane-message-rate-limit"], d, "control_plane_message_rate_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["control-plane-message-rate-limit"], "ObjectFirewallGtp-ControlPlaneMessageRateLimit"); ok {
			if err = d.Set("control_plane_message_rate_limit", vv); err != nil {
				return fmt.Errorf("Error reading control_plane_message_rate_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading control_plane_message_rate_limit: %v", err)
		}
	}

	if err = d.Set("default_apn_action", flattenObjectFirewallGtpDefaultApnAction(o["default-apn-action"], d, "default_apn_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-apn-action"], "ObjectFirewallGtp-DefaultApnAction"); ok {
			if err = d.Set("default_apn_action", vv); err != nil {
				return fmt.Errorf("Error reading default_apn_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_apn_action: %v", err)
		}
	}

	if err = d.Set("default_imsi_action", flattenObjectFirewallGtpDefaultImsiAction(o["default-imsi-action"], d, "default_imsi_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-imsi-action"], "ObjectFirewallGtp-DefaultImsiAction"); ok {
			if err = d.Set("default_imsi_action", vv); err != nil {
				return fmt.Errorf("Error reading default_imsi_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_imsi_action: %v", err)
		}
	}

	if err = d.Set("default_ip_action", flattenObjectFirewallGtpDefaultIpAction(o["default-ip-action"], d, "default_ip_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-ip-action"], "ObjectFirewallGtp-DefaultIpAction"); ok {
			if err = d.Set("default_ip_action", vv); err != nil {
				return fmt.Errorf("Error reading default_ip_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_ip_action: %v", err)
		}
	}

	if err = d.Set("default_noip_action", flattenObjectFirewallGtpDefaultNoipAction(o["default-noip-action"], d, "default_noip_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-noip-action"], "ObjectFirewallGtp-DefaultNoipAction"); ok {
			if err = d.Set("default_noip_action", vv); err != nil {
				return fmt.Errorf("Error reading default_noip_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_noip_action: %v", err)
		}
	}

	if err = d.Set("default_policy_action", flattenObjectFirewallGtpDefaultPolicyAction(o["default-policy-action"], d, "default_policy_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-policy-action"], "ObjectFirewallGtp-DefaultPolicyAction"); ok {
			if err = d.Set("default_policy_action", vv); err != nil {
				return fmt.Errorf("Error reading default_policy_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_policy_action: %v", err)
		}
	}

	if err = d.Set("denied_log", flattenObjectFirewallGtpDeniedLog(o["denied-log"], d, "denied_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["denied-log"], "ObjectFirewallGtp-DeniedLog"); ok {
			if err = d.Set("denied_log", vv); err != nil {
				return fmt.Errorf("Error reading denied_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading denied_log: %v", err)
		}
	}

	if err = d.Set("echo_request_interval", flattenObjectFirewallGtpEchoRequestInterval(o["echo-request-interval"], d, "echo_request_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["echo-request-interval"], "ObjectFirewallGtp-EchoRequestInterval"); ok {
			if err = d.Set("echo_request_interval", vv); err != nil {
				return fmt.Errorf("Error reading echo_request_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading echo_request_interval: %v", err)
		}
	}

	if err = d.Set("extension_log", flattenObjectFirewallGtpExtensionLog(o["extension-log"], d, "extension_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension-log"], "ObjectFirewallGtp-ExtensionLog"); ok {
			if err = d.Set("extension_log", vv); err != nil {
				return fmt.Errorf("Error reading extension_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension_log: %v", err)
		}
	}

	if err = d.Set("forwarded_log", flattenObjectFirewallGtpForwardedLog(o["forwarded-log"], d, "forwarded_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["forwarded-log"], "ObjectFirewallGtp-ForwardedLog"); ok {
			if err = d.Set("forwarded_log", vv); err != nil {
				return fmt.Errorf("Error reading forwarded_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forwarded_log: %v", err)
		}
	}

	if err = d.Set("global_tunnel_limit", flattenObjectFirewallGtpGlobalTunnelLimit(o["global-tunnel-limit"], d, "global_tunnel_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-tunnel-limit"], "ObjectFirewallGtp-GlobalTunnelLimit"); ok {
			if err = d.Set("global_tunnel_limit", vv); err != nil {
				return fmt.Errorf("Error reading global_tunnel_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_tunnel_limit: %v", err)
		}
	}

	if err = d.Set("gtp_in_gtp", flattenObjectFirewallGtpGtpInGtp(o["gtp-in-gtp"], d, "gtp_in_gtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-in-gtp"], "ObjectFirewallGtp-GtpInGtp"); ok {
			if err = d.Set("gtp_in_gtp", vv); err != nil {
				return fmt.Errorf("Error reading gtp_in_gtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_in_gtp: %v", err)
		}
	}

	if err = d.Set("gtpu_denied_log", flattenObjectFirewallGtpGtpuDeniedLog(o["gtpu-denied-log"], d, "gtpu_denied_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtpu-denied-log"], "ObjectFirewallGtp-GtpuDeniedLog"); ok {
			if err = d.Set("gtpu_denied_log", vv); err != nil {
				return fmt.Errorf("Error reading gtpu_denied_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtpu_denied_log: %v", err)
		}
	}

	if err = d.Set("gtpu_forwarded_log", flattenObjectFirewallGtpGtpuForwardedLog(o["gtpu-forwarded-log"], d, "gtpu_forwarded_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtpu-forwarded-log"], "ObjectFirewallGtp-GtpuForwardedLog"); ok {
			if err = d.Set("gtpu_forwarded_log", vv); err != nil {
				return fmt.Errorf("Error reading gtpu_forwarded_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtpu_forwarded_log: %v", err)
		}
	}

	if err = d.Set("gtpu_log_freq", flattenObjectFirewallGtpGtpuLogFreq(o["gtpu-log-freq"], d, "gtpu_log_freq")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtpu-log-freq"], "ObjectFirewallGtp-GtpuLogFreq"); ok {
			if err = d.Set("gtpu_log_freq", vv); err != nil {
				return fmt.Errorf("Error reading gtpu_log_freq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtpu_log_freq: %v", err)
		}
	}

	if err = d.Set("half_close_timeout", flattenObjectFirewallGtpHalfCloseTimeout(o["half-close-timeout"], d, "half_close_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["half-close-timeout"], "ObjectFirewallGtp-HalfCloseTimeout"); ok {
			if err = d.Set("half_close_timeout", vv); err != nil {
				return fmt.Errorf("Error reading half_close_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading half_close_timeout: %v", err)
		}
	}

	if err = d.Set("half_open_timeout", flattenObjectFirewallGtpHalfOpenTimeout(o["half-open-timeout"], d, "half_open_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["half-open-timeout"], "ObjectFirewallGtp-HalfOpenTimeout"); ok {
			if err = d.Set("half_open_timeout", vv); err != nil {
				return fmt.Errorf("Error reading half_open_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading half_open_timeout: %v", err)
		}
	}

	if err = d.Set("handover_group", flattenObjectFirewallGtpHandoverGroup(o["handover-group"], d, "handover_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["handover-group"], "ObjectFirewallGtp-HandoverGroup"); ok {
			if err = d.Set("handover_group", vv); err != nil {
				return fmt.Errorf("Error reading handover_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading handover_group: %v", err)
		}
	}

	if err = d.Set("handover_group6", flattenObjectFirewallGtpHandoverGroup6(o["handover-group6"], d, "handover_group6")); err != nil {
		if vv, ok := fortiAPIPatch(o["handover-group6"], "ObjectFirewallGtp-HandoverGroup6"); ok {
			if err = d.Set("handover_group6", vv); err != nil {
				return fmt.Errorf("Error reading handover_group6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading handover_group6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ie_remove_policy", flattenObjectFirewallGtpIeRemovePolicy(o["ie-remove-policy"], d, "ie_remove_policy")); err != nil {
			if vv, ok := fortiAPIPatch(o["ie-remove-policy"], "ObjectFirewallGtp-IeRemovePolicy"); ok {
				if err = d.Set("ie_remove_policy", vv); err != nil {
					return fmt.Errorf("Error reading ie_remove_policy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ie_remove_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ie_remove_policy"); ok {
			if err = d.Set("ie_remove_policy", flattenObjectFirewallGtpIeRemovePolicy(o["ie-remove-policy"], d, "ie_remove_policy")); err != nil {
				if vv, ok := fortiAPIPatch(o["ie-remove-policy"], "ObjectFirewallGtp-IeRemovePolicy"); ok {
					if err = d.Set("ie_remove_policy", vv); err != nil {
						return fmt.Errorf("Error reading ie_remove_policy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ie_remove_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("ie_remover", flattenObjectFirewallGtpIeRemover(o["ie-remover"], d, "ie_remover")); err != nil {
		if vv, ok := fortiAPIPatch(o["ie-remover"], "ObjectFirewallGtp-IeRemover"); ok {
			if err = d.Set("ie_remover", vv); err != nil {
				return fmt.Errorf("Error reading ie_remover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ie_remover: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ie_validation", flattenObjectFirewallGtpIeValidation(o["ie-validation"], d, "ie_validation")); err != nil {
			if vv, ok := fortiAPIPatch(o["ie-validation"], "ObjectFirewallGtp-IeValidation"); ok {
				if err = d.Set("ie_validation", vv); err != nil {
					return fmt.Errorf("Error reading ie_validation: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ie_validation: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ie_validation"); ok {
			if err = d.Set("ie_validation", flattenObjectFirewallGtpIeValidation(o["ie-validation"], d, "ie_validation")); err != nil {
				if vv, ok := fortiAPIPatch(o["ie-validation"], "ObjectFirewallGtp-IeValidation"); ok {
					if err = d.Set("ie_validation", vv); err != nil {
						return fmt.Errorf("Error reading ie_validation: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ie_validation: %v", err)
				}
			}
		}
	}

	if err = d.Set("ie_white_list_v0v1", flattenObjectFirewallGtpIeWhiteListV0V1(o["ie-white-list-v0v1"], d, "ie_white_list_v0v1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ie-white-list-v0v1"], "ObjectFirewallGtp-IeWhiteListV0V1"); ok {
			if err = d.Set("ie_white_list_v0v1", vv); err != nil {
				return fmt.Errorf("Error reading ie_white_list_v0v1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ie_white_list_v0v1: %v", err)
		}
	}

	if err = d.Set("ie_white_list_v2", flattenObjectFirewallGtpIeWhiteListV2(o["ie-white-list-v2"], d, "ie_white_list_v2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ie-white-list-v2"], "ObjectFirewallGtp-IeWhiteListV2"); ok {
			if err = d.Set("ie_white_list_v2", vv); err != nil {
				return fmt.Errorf("Error reading ie_white_list_v2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ie_white_list_v2: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("imsi", flattenObjectFirewallGtpImsi(o["imsi"], d, "imsi")); err != nil {
			if vv, ok := fortiAPIPatch(o["imsi"], "ObjectFirewallGtp-Imsi"); ok {
				if err = d.Set("imsi", vv); err != nil {
					return fmt.Errorf("Error reading imsi: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading imsi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imsi"); ok {
			if err = d.Set("imsi", flattenObjectFirewallGtpImsi(o["imsi"], d, "imsi")); err != nil {
				if vv, ok := fortiAPIPatch(o["imsi"], "ObjectFirewallGtp-Imsi"); ok {
					if err = d.Set("imsi", vv); err != nil {
						return fmt.Errorf("Error reading imsi: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading imsi: %v", err)
				}
			}
		}
	}

	if err = d.Set("imsi_filter", flattenObjectFirewallGtpImsiFilter(o["imsi-filter"], d, "imsi_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["imsi-filter"], "ObjectFirewallGtp-ImsiFilter"); ok {
			if err = d.Set("imsi_filter", vv); err != nil {
				return fmt.Errorf("Error reading imsi_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imsi_filter: %v", err)
		}
	}

	if err = d.Set("interface_notify", flattenObjectFirewallGtpInterfaceNotify(o["interface-notify"], d, "interface_notify")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-notify"], "ObjectFirewallGtp-InterfaceNotify"); ok {
			if err = d.Set("interface_notify", vv); err != nil {
				return fmt.Errorf("Error reading interface_notify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_notify: %v", err)
		}
	}

	if err = d.Set("invalid_reserved_field", flattenObjectFirewallGtpInvalidReservedField(o["invalid-reserved-field"], d, "invalid_reserved_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["invalid-reserved-field"], "ObjectFirewallGtp-InvalidReservedField"); ok {
			if err = d.Set("invalid_reserved_field", vv); err != nil {
				return fmt.Errorf("Error reading invalid_reserved_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invalid_reserved_field: %v", err)
		}
	}

	if err = d.Set("invalid_sgsns_to_log", flattenObjectFirewallGtpInvalidSgsnsToLog(o["invalid-sgsns-to-log"], d, "invalid_sgsns_to_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["invalid-sgsns-to-log"], "ObjectFirewallGtp-InvalidSgsnsToLog"); ok {
			if err = d.Set("invalid_sgsns_to_log", vv); err != nil {
				return fmt.Errorf("Error reading invalid_sgsns_to_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invalid_sgsns_to_log: %v", err)
		}
	}

	if err = d.Set("invalid_sgsns6_to_log", flattenObjectFirewallGtpInvalidSgsns6ToLog(o["invalid-sgsns6-to-log"], d, "invalid_sgsns6_to_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["invalid-sgsns6-to-log"], "ObjectFirewallGtp-InvalidSgsns6ToLog"); ok {
			if err = d.Set("invalid_sgsns6_to_log", vv); err != nil {
				return fmt.Errorf("Error reading invalid_sgsns6_to_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invalid_sgsns6_to_log: %v", err)
		}
	}

	if err = d.Set("ip_filter", flattenObjectFirewallGtpIpFilter(o["ip-filter"], d, "ip_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-filter"], "ObjectFirewallGtp-IpFilter"); ok {
			if err = d.Set("ip_filter", vv); err != nil {
				return fmt.Errorf("Error reading ip_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_filter: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_policy", flattenObjectFirewallGtpIpPolicy(o["ip-policy"], d, "ip_policy")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-policy"], "ObjectFirewallGtp-IpPolicy"); ok {
				if err = d.Set("ip_policy", vv); err != nil {
					return fmt.Errorf("Error reading ip_policy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_policy"); ok {
			if err = d.Set("ip_policy", flattenObjectFirewallGtpIpPolicy(o["ip-policy"], d, "ip_policy")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-policy"], "ObjectFirewallGtp-IpPolicy"); ok {
					if err = d.Set("ip_policy", vv); err != nil {
						return fmt.Errorf("Error reading ip_policy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_freq", flattenObjectFirewallGtpLogFreq(o["log-freq"], d, "log_freq")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-freq"], "ObjectFirewallGtp-LogFreq"); ok {
			if err = d.Set("log_freq", vv); err != nil {
				return fmt.Errorf("Error reading log_freq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_freq: %v", err)
		}
	}

	if err = d.Set("log_gtpu_limit", flattenObjectFirewallGtpLogGtpuLimit(o["log-gtpu-limit"], d, "log_gtpu_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-gtpu-limit"], "ObjectFirewallGtp-LogGtpuLimit"); ok {
			if err = d.Set("log_gtpu_limit", vv); err != nil {
				return fmt.Errorf("Error reading log_gtpu_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_gtpu_limit: %v", err)
		}
	}

	if err = d.Set("log_imsi_prefix", flattenObjectFirewallGtpLogImsiPrefix(o["log-imsi-prefix"], d, "log_imsi_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-imsi-prefix"], "ObjectFirewallGtp-LogImsiPrefix"); ok {
			if err = d.Set("log_imsi_prefix", vv); err != nil {
				return fmt.Errorf("Error reading log_imsi_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_imsi_prefix: %v", err)
		}
	}

	if err = d.Set("log_msisdn_prefix", flattenObjectFirewallGtpLogMsisdnPrefix(o["log-msisdn-prefix"], d, "log_msisdn_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-msisdn-prefix"], "ObjectFirewallGtp-LogMsisdnPrefix"); ok {
			if err = d.Set("log_msisdn_prefix", vv); err != nil {
				return fmt.Errorf("Error reading log_msisdn_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_msisdn_prefix: %v", err)
		}
	}

	if err = d.Set("max_message_length", flattenObjectFirewallGtpMaxMessageLength(o["max-message-length"], d, "max_message_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-message-length"], "ObjectFirewallGtp-MaxMessageLength"); ok {
			if err = d.Set("max_message_length", vv); err != nil {
				return fmt.Errorf("Error reading max_message_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_message_length: %v", err)
		}
	}

	if err = d.Set("message_filter_v0v1", flattenObjectFirewallGtpMessageFilterV0V1(o["message-filter-v0v1"], d, "message_filter_v0v1")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-filter-v0v1"], "ObjectFirewallGtp-MessageFilterV0V1"); ok {
			if err = d.Set("message_filter_v0v1", vv); err != nil {
				return fmt.Errorf("Error reading message_filter_v0v1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_filter_v0v1: %v", err)
		}
	}

	if err = d.Set("message_filter_v2", flattenObjectFirewallGtpMessageFilterV2(o["message-filter-v2"], d, "message_filter_v2")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-filter-v2"], "ObjectFirewallGtp-MessageFilterV2"); ok {
			if err = d.Set("message_filter_v2", vv); err != nil {
				return fmt.Errorf("Error reading message_filter_v2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_filter_v2: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("message_rate_limit", flattenObjectFirewallGtpMessageRateLimit(o["message-rate-limit"], d, "message_rate_limit")); err != nil {
			if vv, ok := fortiAPIPatch(o["message-rate-limit"], "ObjectFirewallGtp-MessageRateLimit"); ok {
				if err = d.Set("message_rate_limit", vv); err != nil {
					return fmt.Errorf("Error reading message_rate_limit: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading message_rate_limit: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("message_rate_limit"); ok {
			if err = d.Set("message_rate_limit", flattenObjectFirewallGtpMessageRateLimit(o["message-rate-limit"], d, "message_rate_limit")); err != nil {
				if vv, ok := fortiAPIPatch(o["message-rate-limit"], "ObjectFirewallGtp-MessageRateLimit"); ok {
					if err = d.Set("message_rate_limit", vv); err != nil {
						return fmt.Errorf("Error reading message_rate_limit: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading message_rate_limit: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("message_rate_limit_v0", flattenObjectFirewallGtpMessageRateLimitV0(o["message-rate-limit-v0"], d, "message_rate_limit_v0")); err != nil {
			if vv, ok := fortiAPIPatch(o["message-rate-limit-v0"], "ObjectFirewallGtp-MessageRateLimitV0"); ok {
				if err = d.Set("message_rate_limit_v0", vv); err != nil {
					return fmt.Errorf("Error reading message_rate_limit_v0: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading message_rate_limit_v0: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("message_rate_limit_v0"); ok {
			if err = d.Set("message_rate_limit_v0", flattenObjectFirewallGtpMessageRateLimitV0(o["message-rate-limit-v0"], d, "message_rate_limit_v0")); err != nil {
				if vv, ok := fortiAPIPatch(o["message-rate-limit-v0"], "ObjectFirewallGtp-MessageRateLimitV0"); ok {
					if err = d.Set("message_rate_limit_v0", vv); err != nil {
						return fmt.Errorf("Error reading message_rate_limit_v0: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading message_rate_limit_v0: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("message_rate_limit_v1", flattenObjectFirewallGtpMessageRateLimitV1(o["message-rate-limit-v1"], d, "message_rate_limit_v1")); err != nil {
			if vv, ok := fortiAPIPatch(o["message-rate-limit-v1"], "ObjectFirewallGtp-MessageRateLimitV1"); ok {
				if err = d.Set("message_rate_limit_v1", vv); err != nil {
					return fmt.Errorf("Error reading message_rate_limit_v1: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading message_rate_limit_v1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("message_rate_limit_v1"); ok {
			if err = d.Set("message_rate_limit_v1", flattenObjectFirewallGtpMessageRateLimitV1(o["message-rate-limit-v1"], d, "message_rate_limit_v1")); err != nil {
				if vv, ok := fortiAPIPatch(o["message-rate-limit-v1"], "ObjectFirewallGtp-MessageRateLimitV1"); ok {
					if err = d.Set("message_rate_limit_v1", vv); err != nil {
						return fmt.Errorf("Error reading message_rate_limit_v1: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading message_rate_limit_v1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("message_rate_limit_v2", flattenObjectFirewallGtpMessageRateLimitV2(o["message-rate-limit-v2"], d, "message_rate_limit_v2")); err != nil {
			if vv, ok := fortiAPIPatch(o["message-rate-limit-v2"], "ObjectFirewallGtp-MessageRateLimitV2"); ok {
				if err = d.Set("message_rate_limit_v2", vv); err != nil {
					return fmt.Errorf("Error reading message_rate_limit_v2: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading message_rate_limit_v2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("message_rate_limit_v2"); ok {
			if err = d.Set("message_rate_limit_v2", flattenObjectFirewallGtpMessageRateLimitV2(o["message-rate-limit-v2"], d, "message_rate_limit_v2")); err != nil {
				if vv, ok := fortiAPIPatch(o["message-rate-limit-v2"], "ObjectFirewallGtp-MessageRateLimitV2"); ok {
					if err = d.Set("message_rate_limit_v2", vv); err != nil {
						return fmt.Errorf("Error reading message_rate_limit_v2: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading message_rate_limit_v2: %v", err)
				}
			}
		}
	}

	if err = d.Set("min_message_length", flattenObjectFirewallGtpMinMessageLength(o["min-message-length"], d, "min_message_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-message-length"], "ObjectFirewallGtp-MinMessageLength"); ok {
			if err = d.Set("min_message_length", vv); err != nil {
				return fmt.Errorf("Error reading min_message_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_message_length: %v", err)
		}
	}

	if err = d.Set("miss_must_ie", flattenObjectFirewallGtpMissMustIe(o["miss-must-ie"], d, "miss_must_ie")); err != nil {
		if vv, ok := fortiAPIPatch(o["miss-must-ie"], "ObjectFirewallGtp-MissMustIe"); ok {
			if err = d.Set("miss_must_ie", vv); err != nil {
				return fmt.Errorf("Error reading miss_must_ie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading miss_must_ie: %v", err)
		}
	}

	if err = d.Set("monitor_mode", flattenObjectFirewallGtpMonitorMode(o["monitor-mode"], d, "monitor_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-mode"], "ObjectFirewallGtp-MonitorMode"); ok {
			if err = d.Set("monitor_mode", vv); err != nil {
				return fmt.Errorf("Error reading monitor_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_mode: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallGtpName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallGtp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("noip_filter", flattenObjectFirewallGtpNoipFilter(o["noip-filter"], d, "noip_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["noip-filter"], "ObjectFirewallGtp-NoipFilter"); ok {
			if err = d.Set("noip_filter", vv); err != nil {
				return fmt.Errorf("Error reading noip_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading noip_filter: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("noip_policy", flattenObjectFirewallGtpNoipPolicy(o["noip-policy"], d, "noip_policy")); err != nil {
			if vv, ok := fortiAPIPatch(o["noip-policy"], "ObjectFirewallGtp-NoipPolicy"); ok {
				if err = d.Set("noip_policy", vv); err != nil {
					return fmt.Errorf("Error reading noip_policy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading noip_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("noip_policy"); ok {
			if err = d.Set("noip_policy", flattenObjectFirewallGtpNoipPolicy(o["noip-policy"], d, "noip_policy")); err != nil {
				if vv, ok := fortiAPIPatch(o["noip-policy"], "ObjectFirewallGtp-NoipPolicy"); ok {
					if err = d.Set("noip_policy", vv); err != nil {
						return fmt.Errorf("Error reading noip_policy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading noip_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("out_of_state_ie", flattenObjectFirewallGtpOutOfStateIe(o["out-of-state-ie"], d, "out_of_state_ie")); err != nil {
		if vv, ok := fortiAPIPatch(o["out-of-state-ie"], "ObjectFirewallGtp-OutOfStateIe"); ok {
			if err = d.Set("out_of_state_ie", vv); err != nil {
				return fmt.Errorf("Error reading out_of_state_ie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading out_of_state_ie: %v", err)
		}
	}

	if err = d.Set("out_of_state_message", flattenObjectFirewallGtpOutOfStateMessage(o["out-of-state-message"], d, "out_of_state_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["out-of-state-message"], "ObjectFirewallGtp-OutOfStateMessage"); ok {
			if err = d.Set("out_of_state_message", vv); err != nil {
				return fmt.Errorf("Error reading out_of_state_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading out_of_state_message: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("per_apn_shaper", flattenObjectFirewallGtpPerApnShaper(o["per-apn-shaper"], d, "per_apn_shaper")); err != nil {
			if vv, ok := fortiAPIPatch(o["per-apn-shaper"], "ObjectFirewallGtp-PerApnShaper"); ok {
				if err = d.Set("per_apn_shaper", vv); err != nil {
					return fmt.Errorf("Error reading per_apn_shaper: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading per_apn_shaper: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("per_apn_shaper"); ok {
			if err = d.Set("per_apn_shaper", flattenObjectFirewallGtpPerApnShaper(o["per-apn-shaper"], d, "per_apn_shaper")); err != nil {
				if vv, ok := fortiAPIPatch(o["per-apn-shaper"], "ObjectFirewallGtp-PerApnShaper"); ok {
					if err = d.Set("per_apn_shaper", vv); err != nil {
						return fmt.Errorf("Error reading per_apn_shaper: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading per_apn_shaper: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("policy", flattenObjectFirewallGtpPolicy(o["policy"], d, "policy")); err != nil {
			if vv, ok := fortiAPIPatch(o["policy"], "ObjectFirewallGtp-Policy"); ok {
				if err = d.Set("policy", vv); err != nil {
					return fmt.Errorf("Error reading policy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("policy"); ok {
			if err = d.Set("policy", flattenObjectFirewallGtpPolicy(o["policy"], d, "policy")); err != nil {
				if vv, ok := fortiAPIPatch(o["policy"], "ObjectFirewallGtp-Policy"); ok {
					if err = d.Set("policy", vv); err != nil {
						return fmt.Errorf("Error reading policy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("policy_filter", flattenObjectFirewallGtpPolicyFilter(o["policy-filter"], d, "policy_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-filter"], "ObjectFirewallGtp-PolicyFilter"); ok {
			if err = d.Set("policy_filter", vv); err != nil {
				return fmt.Errorf("Error reading policy_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_filter: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("policy_v2", flattenObjectFirewallGtpPolicyV2(o["policy-v2"], d, "policy_v2")); err != nil {
			if vv, ok := fortiAPIPatch(o["policy-v2"], "ObjectFirewallGtp-PolicyV2"); ok {
				if err = d.Set("policy_v2", vv); err != nil {
					return fmt.Errorf("Error reading policy_v2: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading policy_v2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("policy_v2"); ok {
			if err = d.Set("policy_v2", flattenObjectFirewallGtpPolicyV2(o["policy-v2"], d, "policy_v2")); err != nil {
				if vv, ok := fortiAPIPatch(o["policy-v2"], "ObjectFirewallGtp-PolicyV2"); ok {
					if err = d.Set("policy_v2", vv); err != nil {
						return fmt.Errorf("Error reading policy_v2: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading policy_v2: %v", err)
				}
			}
		}
	}

	if err = d.Set("port_notify", flattenObjectFirewallGtpPortNotify(o["port-notify"], d, "port_notify")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-notify"], "ObjectFirewallGtp-PortNotify"); ok {
			if err = d.Set("port_notify", vv); err != nil {
				return fmt.Errorf("Error reading port_notify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_notify: %v", err)
		}
	}

	if err = d.Set("rate_limit_mode", flattenObjectFirewallGtpRateLimitMode(o["rate-limit-mode"], d, "rate_limit_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-limit-mode"], "ObjectFirewallGtp-RateLimitMode"); ok {
			if err = d.Set("rate_limit_mode", vv); err != nil {
				return fmt.Errorf("Error reading rate_limit_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_limit_mode: %v", err)
		}
	}

	if err = d.Set("rate_limited_log", flattenObjectFirewallGtpRateLimitedLog(o["rate-limited-log"], d, "rate_limited_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-limited-log"], "ObjectFirewallGtp-RateLimitedLog"); ok {
			if err = d.Set("rate_limited_log", vv); err != nil {
				return fmt.Errorf("Error reading rate_limited_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_limited_log: %v", err)
		}
	}

	if err = d.Set("rate_sampling_interval", flattenObjectFirewallGtpRateSamplingInterval(o["rate-sampling-interval"], d, "rate_sampling_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-sampling-interval"], "ObjectFirewallGtp-RateSamplingInterval"); ok {
			if err = d.Set("rate_sampling_interval", vv); err != nil {
				return fmt.Errorf("Error reading rate_sampling_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_sampling_interval: %v", err)
		}
	}

	if err = d.Set("remove_if_echo_expires", flattenObjectFirewallGtpRemoveIfEchoExpires(o["remove-if-echo-expires"], d, "remove_if_echo_expires")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-if-echo-expires"], "ObjectFirewallGtp-RemoveIfEchoExpires"); ok {
			if err = d.Set("remove_if_echo_expires", vv); err != nil {
				return fmt.Errorf("Error reading remove_if_echo_expires: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_if_echo_expires: %v", err)
		}
	}

	if err = d.Set("remove_if_recovery_differ", flattenObjectFirewallGtpRemoveIfRecoveryDiffer(o["remove-if-recovery-differ"], d, "remove_if_recovery_differ")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-if-recovery-differ"], "ObjectFirewallGtp-RemoveIfRecoveryDiffer"); ok {
			if err = d.Set("remove_if_recovery_differ", vv); err != nil {
				return fmt.Errorf("Error reading remove_if_recovery_differ: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_if_recovery_differ: %v", err)
		}
	}

	if err = d.Set("reserved_ie", flattenObjectFirewallGtpReservedIe(o["reserved-ie"], d, "reserved_ie")); err != nil {
		if vv, ok := fortiAPIPatch(o["reserved-ie"], "ObjectFirewallGtp-ReservedIe"); ok {
			if err = d.Set("reserved_ie", vv); err != nil {
				return fmt.Errorf("Error reading reserved_ie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reserved_ie: %v", err)
		}
	}

	if err = d.Set("send_delete_when_timeout", flattenObjectFirewallGtpSendDeleteWhenTimeout(o["send-delete-when-timeout"], d, "send_delete_when_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-delete-when-timeout"], "ObjectFirewallGtp-SendDeleteWhenTimeout"); ok {
			if err = d.Set("send_delete_when_timeout", vv); err != nil {
				return fmt.Errorf("Error reading send_delete_when_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_delete_when_timeout: %v", err)
		}
	}

	if err = d.Set("send_delete_when_timeout_v2", flattenObjectFirewallGtpSendDeleteWhenTimeoutV2(o["send-delete-when-timeout-v2"], d, "send_delete_when_timeout_v2")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-delete-when-timeout-v2"], "ObjectFirewallGtp-SendDeleteWhenTimeoutV2"); ok {
			if err = d.Set("send_delete_when_timeout_v2", vv); err != nil {
				return fmt.Errorf("Error reading send_delete_when_timeout_v2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_delete_when_timeout_v2: %v", err)
		}
	}

	if err = d.Set("spoof_src_addr", flattenObjectFirewallGtpSpoofSrcAddr(o["spoof-src-addr"], d, "spoof_src_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["spoof-src-addr"], "ObjectFirewallGtp-SpoofSrcAddr"); ok {
			if err = d.Set("spoof_src_addr", vv); err != nil {
				return fmt.Errorf("Error reading spoof_src_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spoof_src_addr: %v", err)
		}
	}

	if err = d.Set("state_invalid_log", flattenObjectFirewallGtpStateInvalidLog(o["state-invalid-log"], d, "state_invalid_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["state-invalid-log"], "ObjectFirewallGtp-StateInvalidLog"); ok {
			if err = d.Set("state_invalid_log", vv); err != nil {
				return fmt.Errorf("Error reading state_invalid_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading state_invalid_log: %v", err)
		}
	}

	if err = d.Set("sub_second_interval", flattenObjectFirewallGtpSubSecondInterval(o["sub-second-interval"], d, "sub_second_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["sub-second-interval"], "ObjectFirewallGtp-SubSecondInterval"); ok {
			if err = d.Set("sub_second_interval", vv); err != nil {
				return fmt.Errorf("Error reading sub_second_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sub_second_interval: %v", err)
		}
	}

	if err = d.Set("sub_second_sampling", flattenObjectFirewallGtpSubSecondSampling(o["sub-second-sampling"], d, "sub_second_sampling")); err != nil {
		if vv, ok := fortiAPIPatch(o["sub-second-sampling"], "ObjectFirewallGtp-SubSecondSampling"); ok {
			if err = d.Set("sub_second_sampling", vv); err != nil {
				return fmt.Errorf("Error reading sub_second_sampling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sub_second_sampling: %v", err)
		}
	}

	if err = d.Set("traffic_count_log", flattenObjectFirewallGtpTrafficCountLog(o["traffic-count-log"], d, "traffic_count_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-count-log"], "ObjectFirewallGtp-TrafficCountLog"); ok {
			if err = d.Set("traffic_count_log", vv); err != nil {
				return fmt.Errorf("Error reading traffic_count_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_count_log: %v", err)
		}
	}

	if err = d.Set("tunnel_limit", flattenObjectFirewallGtpTunnelLimit(o["tunnel-limit"], d, "tunnel_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-limit"], "ObjectFirewallGtp-TunnelLimit"); ok {
			if err = d.Set("tunnel_limit", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_limit: %v", err)
		}
	}

	if err = d.Set("tunnel_limit_log", flattenObjectFirewallGtpTunnelLimitLog(o["tunnel-limit-log"], d, "tunnel_limit_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-limit-log"], "ObjectFirewallGtp-TunnelLimitLog"); ok {
			if err = d.Set("tunnel_limit_log", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_limit_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_limit_log: %v", err)
		}
	}

	if err = d.Set("tunnel_timeout", flattenObjectFirewallGtpTunnelTimeout(o["tunnel-timeout"], d, "tunnel_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-timeout"], "ObjectFirewallGtp-TunnelTimeout"); ok {
			if err = d.Set("tunnel_timeout", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_timeout: %v", err)
		}
	}

	if err = d.Set("unknown_version_action", flattenObjectFirewallGtpUnknownVersionAction(o["unknown-version-action"], d, "unknown_version_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-version-action"], "ObjectFirewallGtp-UnknownVersionAction"); ok {
			if err = d.Set("unknown_version_action", vv); err != nil {
				return fmt.Errorf("Error reading unknown_version_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_version_action: %v", err)
		}
	}

	if err = d.Set("user_plane_message_rate_limit", flattenObjectFirewallGtpUserPlaneMessageRateLimit(o["user-plane-message-rate-limit"], d, "user_plane_message_rate_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-plane-message-rate-limit"], "ObjectFirewallGtp-UserPlaneMessageRateLimit"); ok {
			if err = d.Set("user_plane_message_rate_limit", vv); err != nil {
				return fmt.Errorf("Error reading user_plane_message_rate_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_plane_message_rate_limit: %v", err)
		}
	}

	if err = d.Set("warning_threshold", flattenObjectFirewallGtpWarningThreshold(o["warning-threshold"], d, "warning_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["warning-threshold"], "ObjectFirewallGtp-WarningThreshold"); ok {
			if err = d.Set("warning_threshold", vv); err != nil {
				return fmt.Errorf("Error reading warning_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warning_threshold: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpAddrNotify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpApn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectFirewallGtpApnAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apnmember"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apnmember"], _ = expandObjectFirewallGtpApnApnmember(d, i["apnmember"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectFirewallGtpApnId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "selection_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["selection-mode"], _ = expandObjectFirewallGtpApnSelectionMode(d, i["selection_mode"], pre_append)
		} else {
			tmp["selection-mode"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallGtpApnAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpApnApnmember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpApnId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpApnSelectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpApnFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpAuthorizedGgsns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpAuthorizedGgsns6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpAuthorizedSgsns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpAuthorizedSgsns6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpContextId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpControlPlaneMessageRateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpDefaultApnAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpDefaultImsiAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpDefaultIpAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpDefaultNoipAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpDefaultPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpDeniedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpEchoRequestInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpExtensionLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpForwardedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpGlobalTunnelLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpGtpInGtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpGtpuDeniedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpGtpuForwardedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpGtpuLogFreq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpHalfCloseTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpHalfOpenTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpHandoverGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpHandoverGroup6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeRemovePolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectFirewallGtpIeRemovePolicyId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_ies"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-ies"], _ = expandObjectFirewallGtpIeRemovePolicyRemoveIes(d, i["remove_ies"], pre_append)
		} else {
			tmp["remove-ies"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sgsn_addr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sgsn-addr"], _ = expandObjectFirewallGtpIeRemovePolicySgsnAddr(d, i["sgsn_addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sgsn_addr6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sgsn-addr6"], _ = expandObjectFirewallGtpIeRemovePolicySgsnAddr6(d, i["sgsn_addr6"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallGtpIeRemovePolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeRemovePolicyRemoveIes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpIeRemovePolicySgsnAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeRemovePolicySgsnAddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeRemover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "apn_restriction"
	if _, ok := d.GetOk(pre_append); ok {
		result["apn-restriction"], _ = expandObjectFirewallGtpIeValidationApnRestriction(d, i["apn_restriction"], pre_append)
	}
	pre_append = pre + ".0." + "charging_ID"
	if _, ok := d.GetOk(pre_append); ok {
		result["charging-ID"], _ = expandObjectFirewallGtpIeValidationChargingId(d, i["charging_ID"], pre_append)
	}
	pre_append = pre + ".0." + "charging_gateway_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["charging-gateway-addr"], _ = expandObjectFirewallGtpIeValidationChargingGatewayAddr(d, i["charging_gateway_addr"], pre_append)
	}
	pre_append = pre + ".0." + "end_user_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["end-user-addr"], _ = expandObjectFirewallGtpIeValidationEndUserAddr(d, i["end_user_addr"], pre_append)
	}
	pre_append = pre + ".0." + "gsn_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["gsn-addr"], _ = expandObjectFirewallGtpIeValidationGsnAddr(d, i["gsn_addr"], pre_append)
	}
	pre_append = pre + ".0." + "imei"
	if _, ok := d.GetOk(pre_append); ok {
		result["imei"], _ = expandObjectFirewallGtpIeValidationImei(d, i["imei"], pre_append)
	}
	pre_append = pre + ".0." + "imsi"
	if _, ok := d.GetOk(pre_append); ok {
		result["imsi"], _ = expandObjectFirewallGtpIeValidationImsi(d, i["imsi"], pre_append)
	}
	pre_append = pre + ".0." + "mm_context"
	if _, ok := d.GetOk(pre_append); ok {
		result["mm-context"], _ = expandObjectFirewallGtpIeValidationMmContext(d, i["mm_context"], pre_append)
	}
	pre_append = pre + ".0." + "ms_tzone"
	if _, ok := d.GetOk(pre_append); ok {
		result["ms-tzone"], _ = expandObjectFirewallGtpIeValidationMsTzone(d, i["ms_tzone"], pre_append)
	}
	pre_append = pre + ".0." + "ms_validated"
	if _, ok := d.GetOk(pre_append); ok {
		result["ms-validated"], _ = expandObjectFirewallGtpIeValidationMsValidated(d, i["ms_validated"], pre_append)
	}
	pre_append = pre + ".0." + "msisdn"
	if _, ok := d.GetOk(pre_append); ok {
		result["msisdn"], _ = expandObjectFirewallGtpIeValidationMsisdn(d, i["msisdn"], pre_append)
	}
	pre_append = pre + ".0." + "nsapi"
	if _, ok := d.GetOk(pre_append); ok {
		result["nsapi"], _ = expandObjectFirewallGtpIeValidationNsapi(d, i["nsapi"], pre_append)
	}
	pre_append = pre + ".0." + "pdp_context"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdp-context"], _ = expandObjectFirewallGtpIeValidationPdpContext(d, i["pdp_context"], pre_append)
	}
	pre_append = pre + ".0." + "qos_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["qos-profile"], _ = expandObjectFirewallGtpIeValidationQosProfile(d, i["qos_profile"], pre_append)
	}
	pre_append = pre + ".0." + "rai"
	if _, ok := d.GetOk(pre_append); ok {
		result["rai"], _ = expandObjectFirewallGtpIeValidationRai(d, i["rai"], pre_append)
	}
	pre_append = pre + ".0." + "rat_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["rat-type"], _ = expandObjectFirewallGtpIeValidationRatType(d, i["rat_type"], pre_append)
	}
	pre_append = pre + ".0." + "reordering_required"
	if _, ok := d.GetOk(pre_append); ok {
		result["reordering-required"], _ = expandObjectFirewallGtpIeValidationReorderingRequired(d, i["reordering_required"], pre_append)
	}
	pre_append = pre + ".0." + "selection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["selection-mode"], _ = expandObjectFirewallGtpIeValidationSelectionMode(d, i["selection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "uli"
	if _, ok := d.GetOk(pre_append); ok {
		result["uli"], _ = expandObjectFirewallGtpIeValidationUli(d, i["uli"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallGtpIeValidationApnRestriction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationChargingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationChargingGatewayAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationEndUserAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationGsnAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationImei(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationImsi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationMmContext(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationMsTzone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationMsValidated(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationMsisdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationNsapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationPdpContext(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationQosProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationRai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationRatType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationReorderingRequired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationSelectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationUli(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeWhiteListV0V1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeWhiteListV2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectFirewallGtpImsiAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apnmember"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apnmember"], _ = expandObjectFirewallGtpImsiApnmember(d, i["apnmember"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectFirewallGtpImsiId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcc_mnc"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mcc-mnc"], _ = expandObjectFirewallGtpImsiMccMnc(d, i["mcc_mnc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msisdn_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msisdn-prefix"], _ = expandObjectFirewallGtpImsiMsisdnPrefix(d, i["msisdn_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "selection_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["selection-mode"], _ = expandObjectFirewallGtpImsiSelectionMode(d, i["selection_mode"], pre_append)
		} else {
			tmp["selection-mode"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallGtpImsiAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsiApnmember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsiId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsiMccMnc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsiMsisdnPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsiSelectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpImsiFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpInterfaceNotify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpInvalidReservedField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpInvalidSgsnsToLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpInvalidSgsns6ToLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIpFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIpPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectFirewallGtpIpPolicyAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dstaddr"], _ = expandObjectFirewallGtpIpPolicyDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dstaddr6"], _ = expandObjectFirewallGtpIpPolicyDstaddr6(d, i["dstaddr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectFirewallGtpIpPolicyId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["srcaddr"], _ = expandObjectFirewallGtpIpPolicySrcaddr(d, i["srcaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["srcaddr6"], _ = expandObjectFirewallGtpIpPolicySrcaddr6(d, i["srcaddr6"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallGtpIpPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIpPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIpPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIpPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIpPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIpPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpLogFreq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpLogGtpuLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpLogImsiPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpLogMsisdnPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMaxMessageLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageFilterV0V1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageFilterV2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "create_aa_pdp_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["create-aa-pdp-request"], _ = expandObjectFirewallGtpMessageRateLimitCreateAaPdpRequest(d, i["create_aa_pdp_request"], pre_append)
	}
	pre_append = pre + ".0." + "create_aa_pdp_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["create-aa-pdp-response"], _ = expandObjectFirewallGtpMessageRateLimitCreateAaPdpResponse(d, i["create_aa_pdp_response"], pre_append)
	}
	pre_append = pre + ".0." + "create_mbms_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["create-mbms-request"], _ = expandObjectFirewallGtpMessageRateLimitCreateMbmsRequest(d, i["create_mbms_request"], pre_append)
	}
	pre_append = pre + ".0." + "create_mbms_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["create-mbms-response"], _ = expandObjectFirewallGtpMessageRateLimitCreateMbmsResponse(d, i["create_mbms_response"], pre_append)
	}
	pre_append = pre + ".0." + "create_pdp_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["create-pdp-request"], _ = expandObjectFirewallGtpMessageRateLimitCreatePdpRequest(d, i["create_pdp_request"], pre_append)
	}
	pre_append = pre + ".0." + "create_pdp_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["create-pdp-response"], _ = expandObjectFirewallGtpMessageRateLimitCreatePdpResponse(d, i["create_pdp_response"], pre_append)
	}
	pre_append = pre + ".0." + "delete_aa_pdp_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["delete-aa-pdp-request"], _ = expandObjectFirewallGtpMessageRateLimitDeleteAaPdpRequest(d, i["delete_aa_pdp_request"], pre_append)
	}
	pre_append = pre + ".0." + "delete_aa_pdp_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["delete-aa-pdp-response"], _ = expandObjectFirewallGtpMessageRateLimitDeleteAaPdpResponse(d, i["delete_aa_pdp_response"], pre_append)
	}
	pre_append = pre + ".0." + "delete_mbms_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["delete-mbms-request"], _ = expandObjectFirewallGtpMessageRateLimitDeleteMbmsRequest(d, i["delete_mbms_request"], pre_append)
	}
	pre_append = pre + ".0." + "delete_mbms_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["delete-mbms-response"], _ = expandObjectFirewallGtpMessageRateLimitDeleteMbmsResponse(d, i["delete_mbms_response"], pre_append)
	}
	pre_append = pre + ".0." + "delete_pdp_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["delete-pdp-request"], _ = expandObjectFirewallGtpMessageRateLimitDeletePdpRequest(d, i["delete_pdp_request"], pre_append)
	}
	pre_append = pre + ".0." + "delete_pdp_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["delete-pdp-response"], _ = expandObjectFirewallGtpMessageRateLimitDeletePdpResponse(d, i["delete_pdp_response"], pre_append)
	}
	pre_append = pre + ".0." + "echo_reponse"
	if _, ok := d.GetOk(pre_append); ok {
		result["echo-reponse"], _ = expandObjectFirewallGtpMessageRateLimitEchoReponse(d, i["echo_reponse"], pre_append)
	}
	pre_append = pre + ".0." + "echo_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["echo-request"], _ = expandObjectFirewallGtpMessageRateLimitEchoRequest(d, i["echo_request"], pre_append)
	}
	pre_append = pre + ".0." + "error_indication"
	if _, ok := d.GetOk(pre_append); ok {
		result["error-indication"], _ = expandObjectFirewallGtpMessageRateLimitErrorIndication(d, i["error_indication"], pre_append)
	}
	pre_append = pre + ".0." + "failure_report_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["failure-report-request"], _ = expandObjectFirewallGtpMessageRateLimitFailureReportRequest(d, i["failure_report_request"], pre_append)
	}
	pre_append = pre + ".0." + "failure_report_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["failure-report-response"], _ = expandObjectFirewallGtpMessageRateLimitFailureReportResponse(d, i["failure_report_response"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_reloc_complete_ack"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-reloc-complete-ack"], _ = expandObjectFirewallGtpMessageRateLimitFwdRelocCompleteAck(d, i["fwd_reloc_complete_ack"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_relocation_complete"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-relocation-complete"], _ = expandObjectFirewallGtpMessageRateLimitFwdRelocationComplete(d, i["fwd_relocation_complete"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_relocation_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-relocation-request"], _ = expandObjectFirewallGtpMessageRateLimitFwdRelocationRequest(d, i["fwd_relocation_request"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_relocation_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-relocation-response"], _ = expandObjectFirewallGtpMessageRateLimitFwdRelocationResponse(d, i["fwd_relocation_response"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_srns_context"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-srns-context"], _ = expandObjectFirewallGtpMessageRateLimitFwdSrnsContext(d, i["fwd_srns_context"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_srns_context_ack"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-srns-context-ack"], _ = expandObjectFirewallGtpMessageRateLimitFwdSrnsContextAck(d, i["fwd_srns_context_ack"], pre_append)
	}
	pre_append = pre + ".0." + "g_pdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["g-pdu"], _ = expandObjectFirewallGtpMessageRateLimitGPdu(d, i["g_pdu"], pre_append)
	}
	pre_append = pre + ".0." + "identification_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["identification-request"], _ = expandObjectFirewallGtpMessageRateLimitIdentificationRequest(d, i["identification_request"], pre_append)
	}
	pre_append = pre + ".0." + "identification_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["identification-response"], _ = expandObjectFirewallGtpMessageRateLimitIdentificationResponse(d, i["identification_response"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_de_reg_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-de-reg-request"], _ = expandObjectFirewallGtpMessageRateLimitMbmsDeRegRequest(d, i["mbms_de_reg_request"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_de_reg_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-de-reg-response"], _ = expandObjectFirewallGtpMessageRateLimitMbmsDeRegResponse(d, i["mbms_de_reg_response"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_notify_rej_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-notify-rej-request"], _ = expandObjectFirewallGtpMessageRateLimitMbmsNotifyRejRequest(d, i["mbms_notify_rej_request"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_notify_rej_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-notify-rej-response"], _ = expandObjectFirewallGtpMessageRateLimitMbmsNotifyRejResponse(d, i["mbms_notify_rej_response"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_notify_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-notify-request"], _ = expandObjectFirewallGtpMessageRateLimitMbmsNotifyRequest(d, i["mbms_notify_request"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_notify_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-notify-response"], _ = expandObjectFirewallGtpMessageRateLimitMbmsNotifyResponse(d, i["mbms_notify_response"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_reg_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-reg-request"], _ = expandObjectFirewallGtpMessageRateLimitMbmsRegRequest(d, i["mbms_reg_request"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_reg_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-reg-response"], _ = expandObjectFirewallGtpMessageRateLimitMbmsRegResponse(d, i["mbms_reg_response"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_ses_start_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-ses-start-request"], _ = expandObjectFirewallGtpMessageRateLimitMbmsSesStartRequest(d, i["mbms_ses_start_request"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_ses_start_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-ses-start-response"], _ = expandObjectFirewallGtpMessageRateLimitMbmsSesStartResponse(d, i["mbms_ses_start_response"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_ses_stop_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-ses-stop-request"], _ = expandObjectFirewallGtpMessageRateLimitMbmsSesStopRequest(d, i["mbms_ses_stop_request"], pre_append)
	}
	pre_append = pre + ".0." + "mbms_ses_stop_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["mbms-ses-stop-response"], _ = expandObjectFirewallGtpMessageRateLimitMbmsSesStopResponse(d, i["mbms_ses_stop_response"], pre_append)
	}
	pre_append = pre + ".0." + "note_ms_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["note-ms-request"], _ = expandObjectFirewallGtpMessageRateLimitNoteMsRequest(d, i["note_ms_request"], pre_append)
	}
	pre_append = pre + ".0." + "note_ms_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["note-ms-response"], _ = expandObjectFirewallGtpMessageRateLimitNoteMsResponse(d, i["note_ms_response"], pre_append)
	}
	pre_append = pre + ".0." + "pdu_notify_rej_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdu-notify-rej-request"], _ = expandObjectFirewallGtpMessageRateLimitPduNotifyRejRequest(d, i["pdu_notify_rej_request"], pre_append)
	}
	pre_append = pre + ".0." + "pdu_notify_rej_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdu-notify-rej-response"], _ = expandObjectFirewallGtpMessageRateLimitPduNotifyRejResponse(d, i["pdu_notify_rej_response"], pre_append)
	}
	pre_append = pre + ".0." + "pdu_notify_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdu-notify-request"], _ = expandObjectFirewallGtpMessageRateLimitPduNotifyRequest(d, i["pdu_notify_request"], pre_append)
	}
	pre_append = pre + ".0." + "pdu_notify_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdu-notify-response"], _ = expandObjectFirewallGtpMessageRateLimitPduNotifyResponse(d, i["pdu_notify_response"], pre_append)
	}
	pre_append = pre + ".0." + "ran_info"
	if _, ok := d.GetOk(pre_append); ok {
		result["ran-info"], _ = expandObjectFirewallGtpMessageRateLimitRanInfo(d, i["ran_info"], pre_append)
	}
	pre_append = pre + ".0." + "relocation_cancel_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["relocation-cancel-request"], _ = expandObjectFirewallGtpMessageRateLimitRelocationCancelRequest(d, i["relocation_cancel_request"], pre_append)
	}
	pre_append = pre + ".0." + "relocation_cancel_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["relocation-cancel-response"], _ = expandObjectFirewallGtpMessageRateLimitRelocationCancelResponse(d, i["relocation_cancel_response"], pre_append)
	}
	pre_append = pre + ".0." + "send_route_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["send-route-request"], _ = expandObjectFirewallGtpMessageRateLimitSendRouteRequest(d, i["send_route_request"], pre_append)
	}
	pre_append = pre + ".0." + "send_route_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["send-route-response"], _ = expandObjectFirewallGtpMessageRateLimitSendRouteResponse(d, i["send_route_response"], pre_append)
	}
	pre_append = pre + ".0." + "sgsn_context_ack"
	if _, ok := d.GetOk(pre_append); ok {
		result["sgsn-context-ack"], _ = expandObjectFirewallGtpMessageRateLimitSgsnContextAck(d, i["sgsn_context_ack"], pre_append)
	}
	pre_append = pre + ".0." + "sgsn_context_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["sgsn-context-request"], _ = expandObjectFirewallGtpMessageRateLimitSgsnContextRequest(d, i["sgsn_context_request"], pre_append)
	}
	pre_append = pre + ".0." + "sgsn_context_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["sgsn-context-response"], _ = expandObjectFirewallGtpMessageRateLimitSgsnContextResponse(d, i["sgsn_context_response"], pre_append)
	}
	pre_append = pre + ".0." + "support_ext_hdr_notify"
	if _, ok := d.GetOk(pre_append); ok {
		result["support-ext-hdr-notify"], _ = expandObjectFirewallGtpMessageRateLimitSupportExtHdrNotify(d, i["support_ext_hdr_notify"], pre_append)
	}
	pre_append = pre + ".0." + "update_mbms_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["update-mbms-request"], _ = expandObjectFirewallGtpMessageRateLimitUpdateMbmsRequest(d, i["update_mbms_request"], pre_append)
	}
	pre_append = pre + ".0." + "update_mbms_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["update-mbms-response"], _ = expandObjectFirewallGtpMessageRateLimitUpdateMbmsResponse(d, i["update_mbms_response"], pre_append)
	}
	pre_append = pre + ".0." + "update_pdp_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["update-pdp-request"], _ = expandObjectFirewallGtpMessageRateLimitUpdatePdpRequest(d, i["update_pdp_request"], pre_append)
	}
	pre_append = pre + ".0." + "update_pdp_response"
	if _, ok := d.GetOk(pre_append); ok {
		result["update-pdp-response"], _ = expandObjectFirewallGtpMessageRateLimitUpdatePdpResponse(d, i["update_pdp_response"], pre_append)
	}
	pre_append = pre + ".0." + "version_not_support"
	if _, ok := d.GetOk(pre_append); ok {
		result["version-not-support"], _ = expandObjectFirewallGtpMessageRateLimitVersionNotSupport(d, i["version_not_support"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallGtpMessageRateLimitCreateAaPdpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitCreateAaPdpResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitCreateMbmsRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitCreateMbmsResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitCreatePdpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitCreatePdpResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeleteAaPdpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeleteAaPdpResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeleteMbmsRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeleteMbmsResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeletePdpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitDeletePdpResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitEchoReponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitEchoRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitErrorIndication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFailureReportRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFailureReportResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdRelocCompleteAck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdRelocationComplete(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdRelocationRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdRelocationResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdSrnsContext(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitFwdSrnsContextAck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitGPdu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitIdentificationRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitIdentificationResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsDeRegRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsDeRegResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsNotifyRejRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsNotifyRejResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsNotifyRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsNotifyResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsRegRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsRegResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsSesStartRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsSesStartResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsSesStopRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitMbmsSesStopResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitNoteMsRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitNoteMsResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitPduNotifyRejRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitPduNotifyRejResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitPduNotifyRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitPduNotifyResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitRanInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitRelocationCancelRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitRelocationCancelResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSendRouteRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSendRouteResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSgsnContextAck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSgsnContextRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSgsnContextResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitSupportExtHdrNotify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitUpdateMbmsRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitUpdateMbmsResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitUpdatePdpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitUpdatePdpResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitVersionNotSupport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "create_pdp_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["create-pdp-request"], _ = expandObjectFirewallGtpMessageRateLimitV0CreatePdpRequest(d, i["create_pdp_request"], pre_append)
	}
	pre_append = pre + ".0." + "delete_pdp_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["delete-pdp-request"], _ = expandObjectFirewallGtpMessageRateLimitV0DeletePdpRequest(d, i["delete_pdp_request"], pre_append)
	}
	pre_append = pre + ".0." + "echo_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["echo-request"], _ = expandObjectFirewallGtpMessageRateLimitV0EchoRequest(d, i["echo_request"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallGtpMessageRateLimitV0CreatePdpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV0DeletePdpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV0EchoRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "create_pdp_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["create-pdp-request"], _ = expandObjectFirewallGtpMessageRateLimitV1CreatePdpRequest(d, i["create_pdp_request"], pre_append)
	}
	pre_append = pre + ".0." + "delete_pdp_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["delete-pdp-request"], _ = expandObjectFirewallGtpMessageRateLimitV1DeletePdpRequest(d, i["delete_pdp_request"], pre_append)
	}
	pre_append = pre + ".0." + "echo_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["echo-request"], _ = expandObjectFirewallGtpMessageRateLimitV1EchoRequest(d, i["echo_request"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallGtpMessageRateLimitV1CreatePdpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV1DeletePdpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV1EchoRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "create_session_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["create-session-request"], _ = expandObjectFirewallGtpMessageRateLimitV2CreateSessionRequest(d, i["create_session_request"], pre_append)
	}
	pre_append = pre + ".0." + "delete_session_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["delete-session-request"], _ = expandObjectFirewallGtpMessageRateLimitV2DeleteSessionRequest(d, i["delete_session_request"], pre_append)
	}
	pre_append = pre + ".0." + "echo_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["echo-request"], _ = expandObjectFirewallGtpMessageRateLimitV2EchoRequest(d, i["echo_request"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallGtpMessageRateLimitV2CreateSessionRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV2DeleteSessionRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV2EchoRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMinMessageLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMissMustIe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMonitorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpNoipFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpNoipPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectFirewallGtpNoipPolicyAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end"], _ = expandObjectFirewallGtpNoipPolicyEnd(d, i["end"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectFirewallGtpNoipPolicyId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start"], _ = expandObjectFirewallGtpNoipPolicyStart(d, i["start"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandObjectFirewallGtpNoipPolicyType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallGtpNoipPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpNoipPolicyEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpNoipPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpNoipPolicyStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpNoipPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpOutOfStateIe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpOutOfStateMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPerApnShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apn"], _ = expandObjectFirewallGtpPerApnShaperApn(d, i["apn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectFirewallGtpPerApnShaperId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_limit"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-limit"], _ = expandObjectFirewallGtpPerApnShaperRateLimit(d, i["rate_limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["version"], _ = expandObjectFirewallGtpPerApnShaperVersion(d, i["version"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallGtpPerApnShaperApn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPerApnShaperId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPerApnShaperRateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPerApnShaperVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectFirewallGtpPolicyAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apn_sel_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apn-sel-mode"], _ = expandObjectFirewallGtpPolicyApnSelMode(d, i["apn_sel_mode"], pre_append)
		} else {
			tmp["apn-sel-mode"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apnmember"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apnmember"], _ = expandObjectFirewallGtpPolicyApnmember(d, i["apnmember"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectFirewallGtpPolicyId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imei"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["imei"], _ = expandObjectFirewallGtpPolicyImei(d, i["imei"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imsi_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["imsi-prefix"], _ = expandObjectFirewallGtpPolicyImsiPrefix(d, i["imsi_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_apn_restriction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-apn-restriction"], _ = expandObjectFirewallGtpPolicyMaxApnRestriction(d, i["max_apn_restriction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "messages"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["messages"], _ = expandObjectFirewallGtpPolicyMessages(d, i["messages"], pre_append)
		} else {
			tmp["messages"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msisdn_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msisdn-prefix"], _ = expandObjectFirewallGtpPolicyMsisdnPrefix(d, i["msisdn_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rai"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rai"], _ = expandObjectFirewallGtpPolicyRai(d, i["rai"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rat_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rat-type"], _ = expandObjectFirewallGtpPolicyRatType(d, i["rat_type"], pre_append)
		} else {
			tmp["rat-type"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uli"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["uli"], _ = expandObjectFirewallGtpPolicyUli(d, i["uli"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallGtpPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyApnSelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyApnmember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyImei(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyImsiPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyMaxApnRestriction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyMessages(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyMsisdnPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyRai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyRatType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyUli(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectFirewallGtpPolicyV2Action(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apn_sel_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apn-sel-mode"], _ = expandObjectFirewallGtpPolicyV2ApnSelMode(d, i["apn_sel_mode"], pre_append)
		} else {
			tmp["apn-sel-mode"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apnmember"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apnmember"], _ = expandObjectFirewallGtpPolicyV2Apnmember(d, i["apnmember"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectFirewallGtpPolicyV2Id(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imsi_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["imsi-prefix"], _ = expandObjectFirewallGtpPolicyV2ImsiPrefix(d, i["imsi_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_apn_restriction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-apn-restriction"], _ = expandObjectFirewallGtpPolicyV2MaxApnRestriction(d, i["max_apn_restriction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mei"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mei"], _ = expandObjectFirewallGtpPolicyV2Mei(d, i["mei"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "messages"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["messages"], _ = expandObjectFirewallGtpPolicyV2Messages(d, i["messages"], pre_append)
		} else {
			tmp["messages"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msisdn_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msisdn-prefix"], _ = expandObjectFirewallGtpPolicyV2MsisdnPrefix(d, i["msisdn_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rat_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rat-type"], _ = expandObjectFirewallGtpPolicyV2RatType(d, i["rat_type"], pre_append)
		} else {
			tmp["rat-type"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uli"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["uli"], _ = expandObjectFirewallGtpPolicyV2Uli(d, i["uli"], pre_append)
		} else {
			tmp["uli"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallGtpPolicyV2Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2ApnSelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyV2Apnmember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2ImsiPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2MaxApnRestriction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2Mei(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2Messages(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyV2MsisdnPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPolicyV2RatType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPolicyV2Uli(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpPortNotify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpRateLimitMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpRateLimitedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpRateSamplingInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpRemoveIfEchoExpires(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpRemoveIfRecoveryDiffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpReservedIe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpSendDeleteWhenTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpSendDeleteWhenTimeoutV2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpSpoofSrcAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpStateInvalidLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpSubSecondInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpSubSecondSampling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpTrafficCountLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpTunnelLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpTunnelLimitLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpTunnelTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpUnknownVersionAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpUserPlaneMessageRateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpWarningThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallGtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_notify"); ok {
		t, err := expandObjectFirewallGtpAddrNotify(d, v, "addr_notify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-notify"] = t
		}
	}

	if v, ok := d.GetOk("apn"); ok {
		t, err := expandObjectFirewallGtpApn(d, v, "apn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn"] = t
		}
	}

	if v, ok := d.GetOk("apn_filter"); ok {
		t, err := expandObjectFirewallGtpApnFilter(d, v, "apn_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn-filter"] = t
		}
	}

	if v, ok := d.GetOk("authorized_ggsns"); ok {
		t, err := expandObjectFirewallGtpAuthorizedGgsns(d, v, "authorized_ggsns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized-ggsns"] = t
		}
	}

	if v, ok := d.GetOk("authorized_ggsns6"); ok {
		t, err := expandObjectFirewallGtpAuthorizedGgsns6(d, v, "authorized_ggsns6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized-ggsns6"] = t
		}
	}

	if v, ok := d.GetOk("authorized_sgsns"); ok {
		t, err := expandObjectFirewallGtpAuthorizedSgsns(d, v, "authorized_sgsns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized-sgsns"] = t
		}
	}

	if v, ok := d.GetOk("authorized_sgsns6"); ok {
		t, err := expandObjectFirewallGtpAuthorizedSgsns6(d, v, "authorized_sgsns6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized-sgsns6"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectFirewallGtpComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("context_id"); ok {
		t, err := expandObjectFirewallGtpContextId(d, v, "context_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["context-id"] = t
		}
	}

	if v, ok := d.GetOk("control_plane_message_rate_limit"); ok {
		t, err := expandObjectFirewallGtpControlPlaneMessageRateLimit(d, v, "control_plane_message_rate_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["control-plane-message-rate-limit"] = t
		}
	}

	if v, ok := d.GetOk("default_apn_action"); ok {
		t, err := expandObjectFirewallGtpDefaultApnAction(d, v, "default_apn_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-apn-action"] = t
		}
	}

	if v, ok := d.GetOk("default_imsi_action"); ok {
		t, err := expandObjectFirewallGtpDefaultImsiAction(d, v, "default_imsi_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-imsi-action"] = t
		}
	}

	if v, ok := d.GetOk("default_ip_action"); ok {
		t, err := expandObjectFirewallGtpDefaultIpAction(d, v, "default_ip_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-ip-action"] = t
		}
	}

	if v, ok := d.GetOk("default_noip_action"); ok {
		t, err := expandObjectFirewallGtpDefaultNoipAction(d, v, "default_noip_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-noip-action"] = t
		}
	}

	if v, ok := d.GetOk("default_policy_action"); ok {
		t, err := expandObjectFirewallGtpDefaultPolicyAction(d, v, "default_policy_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-policy-action"] = t
		}
	}

	if v, ok := d.GetOk("denied_log"); ok {
		t, err := expandObjectFirewallGtpDeniedLog(d, v, "denied_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["denied-log"] = t
		}
	}

	if v, ok := d.GetOk("echo_request_interval"); ok {
		t, err := expandObjectFirewallGtpEchoRequestInterval(d, v, "echo_request_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["echo-request-interval"] = t
		}
	}

	if v, ok := d.GetOk("extension_log"); ok {
		t, err := expandObjectFirewallGtpExtensionLog(d, v, "extension_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-log"] = t
		}
	}

	if v, ok := d.GetOk("forwarded_log"); ok {
		t, err := expandObjectFirewallGtpForwardedLog(d, v, "forwarded_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarded-log"] = t
		}
	}

	if v, ok := d.GetOk("global_tunnel_limit"); ok {
		t, err := expandObjectFirewallGtpGlobalTunnelLimit(d, v, "global_tunnel_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-tunnel-limit"] = t
		}
	}

	if v, ok := d.GetOk("gtp_in_gtp"); ok {
		t, err := expandObjectFirewallGtpGtpInGtp(d, v, "gtp_in_gtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-in-gtp"] = t
		}
	}

	if v, ok := d.GetOk("gtpu_denied_log"); ok {
		t, err := expandObjectFirewallGtpGtpuDeniedLog(d, v, "gtpu_denied_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtpu-denied-log"] = t
		}
	}

	if v, ok := d.GetOk("gtpu_forwarded_log"); ok {
		t, err := expandObjectFirewallGtpGtpuForwardedLog(d, v, "gtpu_forwarded_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtpu-forwarded-log"] = t
		}
	}

	if v, ok := d.GetOk("gtpu_log_freq"); ok {
		t, err := expandObjectFirewallGtpGtpuLogFreq(d, v, "gtpu_log_freq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtpu-log-freq"] = t
		}
	}

	if v, ok := d.GetOk("half_close_timeout"); ok {
		t, err := expandObjectFirewallGtpHalfCloseTimeout(d, v, "half_close_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["half-close-timeout"] = t
		}
	}

	if v, ok := d.GetOk("half_open_timeout"); ok {
		t, err := expandObjectFirewallGtpHalfOpenTimeout(d, v, "half_open_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["half-open-timeout"] = t
		}
	}

	if v, ok := d.GetOk("handover_group"); ok {
		t, err := expandObjectFirewallGtpHandoverGroup(d, v, "handover_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handover-group"] = t
		}
	}

	if v, ok := d.GetOk("handover_group6"); ok {
		t, err := expandObjectFirewallGtpHandoverGroup6(d, v, "handover_group6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handover-group6"] = t
		}
	}

	if v, ok := d.GetOk("ie_remove_policy"); ok {
		t, err := expandObjectFirewallGtpIeRemovePolicy(d, v, "ie_remove_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ie-remove-policy"] = t
		}
	}

	if v, ok := d.GetOk("ie_remover"); ok {
		t, err := expandObjectFirewallGtpIeRemover(d, v, "ie_remover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ie-remover"] = t
		}
	}

	if v, ok := d.GetOk("ie_validation"); ok {
		t, err := expandObjectFirewallGtpIeValidation(d, v, "ie_validation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ie-validation"] = t
		}
	}

	if v, ok := d.GetOk("ie_white_list_v0v1"); ok {
		t, err := expandObjectFirewallGtpIeWhiteListV0V1(d, v, "ie_white_list_v0v1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ie-white-list-v0v1"] = t
		}
	}

	if v, ok := d.GetOk("ie_white_list_v2"); ok {
		t, err := expandObjectFirewallGtpIeWhiteListV2(d, v, "ie_white_list_v2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ie-white-list-v2"] = t
		}
	}

	if v, ok := d.GetOk("imsi"); ok {
		t, err := expandObjectFirewallGtpImsi(d, v, "imsi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imsi"] = t
		}
	}

	if v, ok := d.GetOk("imsi_filter"); ok {
		t, err := expandObjectFirewallGtpImsiFilter(d, v, "imsi_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imsi-filter"] = t
		}
	}

	if v, ok := d.GetOk("interface_notify"); ok {
		t, err := expandObjectFirewallGtpInterfaceNotify(d, v, "interface_notify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-notify"] = t
		}
	}

	if v, ok := d.GetOk("invalid_reserved_field"); ok {
		t, err := expandObjectFirewallGtpInvalidReservedField(d, v, "invalid_reserved_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid-reserved-field"] = t
		}
	}

	if v, ok := d.GetOk("invalid_sgsns_to_log"); ok {
		t, err := expandObjectFirewallGtpInvalidSgsnsToLog(d, v, "invalid_sgsns_to_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid-sgsns-to-log"] = t
		}
	}

	if v, ok := d.GetOk("invalid_sgsns6_to_log"); ok {
		t, err := expandObjectFirewallGtpInvalidSgsns6ToLog(d, v, "invalid_sgsns6_to_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid-sgsns6-to-log"] = t
		}
	}

	if v, ok := d.GetOk("ip_filter"); ok {
		t, err := expandObjectFirewallGtpIpFilter(d, v, "ip_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-filter"] = t
		}
	}

	if v, ok := d.GetOk("ip_policy"); ok {
		t, err := expandObjectFirewallGtpIpPolicy(d, v, "ip_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-policy"] = t
		}
	}

	if v, ok := d.GetOk("log_freq"); ok {
		t, err := expandObjectFirewallGtpLogFreq(d, v, "log_freq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-freq"] = t
		}
	}

	if v, ok := d.GetOk("log_gtpu_limit"); ok {
		t, err := expandObjectFirewallGtpLogGtpuLimit(d, v, "log_gtpu_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-gtpu-limit"] = t
		}
	}

	if v, ok := d.GetOk("log_imsi_prefix"); ok {
		t, err := expandObjectFirewallGtpLogImsiPrefix(d, v, "log_imsi_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-imsi-prefix"] = t
		}
	}

	if v, ok := d.GetOk("log_msisdn_prefix"); ok {
		t, err := expandObjectFirewallGtpLogMsisdnPrefix(d, v, "log_msisdn_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-msisdn-prefix"] = t
		}
	}

	if v, ok := d.GetOk("max_message_length"); ok {
		t, err := expandObjectFirewallGtpMaxMessageLength(d, v, "max_message_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-message-length"] = t
		}
	}

	if v, ok := d.GetOk("message_filter_v0v1"); ok {
		t, err := expandObjectFirewallGtpMessageFilterV0V1(d, v, "message_filter_v0v1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-filter-v0v1"] = t
		}
	}

	if v, ok := d.GetOk("message_filter_v2"); ok {
		t, err := expandObjectFirewallGtpMessageFilterV2(d, v, "message_filter_v2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-filter-v2"] = t
		}
	}

	if v, ok := d.GetOk("message_rate_limit"); ok {
		t, err := expandObjectFirewallGtpMessageRateLimit(d, v, "message_rate_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-rate-limit"] = t
		}
	}

	if v, ok := d.GetOk("message_rate_limit_v0"); ok {
		t, err := expandObjectFirewallGtpMessageRateLimitV0(d, v, "message_rate_limit_v0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-rate-limit-v0"] = t
		}
	}

	if v, ok := d.GetOk("message_rate_limit_v1"); ok {
		t, err := expandObjectFirewallGtpMessageRateLimitV1(d, v, "message_rate_limit_v1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-rate-limit-v1"] = t
		}
	}

	if v, ok := d.GetOk("message_rate_limit_v2"); ok {
		t, err := expandObjectFirewallGtpMessageRateLimitV2(d, v, "message_rate_limit_v2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-rate-limit-v2"] = t
		}
	}

	if v, ok := d.GetOk("min_message_length"); ok {
		t, err := expandObjectFirewallGtpMinMessageLength(d, v, "min_message_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-message-length"] = t
		}
	}

	if v, ok := d.GetOk("miss_must_ie"); ok {
		t, err := expandObjectFirewallGtpMissMustIe(d, v, "miss_must_ie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["miss-must-ie"] = t
		}
	}

	if v, ok := d.GetOk("monitor_mode"); ok {
		t, err := expandObjectFirewallGtpMonitorMode(d, v, "monitor_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallGtpName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("noip_filter"); ok {
		t, err := expandObjectFirewallGtpNoipFilter(d, v, "noip_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["noip-filter"] = t
		}
	}

	if v, ok := d.GetOk("noip_policy"); ok {
		t, err := expandObjectFirewallGtpNoipPolicy(d, v, "noip_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["noip-policy"] = t
		}
	}

	if v, ok := d.GetOk("out_of_state_ie"); ok {
		t, err := expandObjectFirewallGtpOutOfStateIe(d, v, "out_of_state_ie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["out-of-state-ie"] = t
		}
	}

	if v, ok := d.GetOk("out_of_state_message"); ok {
		t, err := expandObjectFirewallGtpOutOfStateMessage(d, v, "out_of_state_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["out-of-state-message"] = t
		}
	}

	if v, ok := d.GetOk("per_apn_shaper"); ok {
		t, err := expandObjectFirewallGtpPerApnShaper(d, v, "per_apn_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-apn-shaper"] = t
		}
	}

	if v, ok := d.GetOk("policy"); ok {
		t, err := expandObjectFirewallGtpPolicy(d, v, "policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy"] = t
		}
	}

	if v, ok := d.GetOk("policy_filter"); ok {
		t, err := expandObjectFirewallGtpPolicyFilter(d, v, "policy_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-filter"] = t
		}
	}

	if v, ok := d.GetOk("policy_v2"); ok {
		t, err := expandObjectFirewallGtpPolicyV2(d, v, "policy_v2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-v2"] = t
		}
	}

	if v, ok := d.GetOk("port_notify"); ok {
		t, err := expandObjectFirewallGtpPortNotify(d, v, "port_notify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-notify"] = t
		}
	}

	if v, ok := d.GetOk("rate_limit_mode"); ok {
		t, err := expandObjectFirewallGtpRateLimitMode(d, v, "rate_limit_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-limit-mode"] = t
		}
	}

	if v, ok := d.GetOk("rate_limited_log"); ok {
		t, err := expandObjectFirewallGtpRateLimitedLog(d, v, "rate_limited_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-limited-log"] = t
		}
	}

	if v, ok := d.GetOk("rate_sampling_interval"); ok {
		t, err := expandObjectFirewallGtpRateSamplingInterval(d, v, "rate_sampling_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-sampling-interval"] = t
		}
	}

	if v, ok := d.GetOk("remove_if_echo_expires"); ok {
		t, err := expandObjectFirewallGtpRemoveIfEchoExpires(d, v, "remove_if_echo_expires")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-if-echo-expires"] = t
		}
	}

	if v, ok := d.GetOk("remove_if_recovery_differ"); ok {
		t, err := expandObjectFirewallGtpRemoveIfRecoveryDiffer(d, v, "remove_if_recovery_differ")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-if-recovery-differ"] = t
		}
	}

	if v, ok := d.GetOk("reserved_ie"); ok {
		t, err := expandObjectFirewallGtpReservedIe(d, v, "reserved_ie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserved-ie"] = t
		}
	}

	if v, ok := d.GetOk("send_delete_when_timeout"); ok {
		t, err := expandObjectFirewallGtpSendDeleteWhenTimeout(d, v, "send_delete_when_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-delete-when-timeout"] = t
		}
	}

	if v, ok := d.GetOk("send_delete_when_timeout_v2"); ok {
		t, err := expandObjectFirewallGtpSendDeleteWhenTimeoutV2(d, v, "send_delete_when_timeout_v2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-delete-when-timeout-v2"] = t
		}
	}

	if v, ok := d.GetOk("spoof_src_addr"); ok {
		t, err := expandObjectFirewallGtpSpoofSrcAddr(d, v, "spoof_src_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spoof-src-addr"] = t
		}
	}

	if v, ok := d.GetOk("state_invalid_log"); ok {
		t, err := expandObjectFirewallGtpStateInvalidLog(d, v, "state_invalid_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state-invalid-log"] = t
		}
	}

	if v, ok := d.GetOk("sub_second_interval"); ok {
		t, err := expandObjectFirewallGtpSubSecondInterval(d, v, "sub_second_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-second-interval"] = t
		}
	}

	if v, ok := d.GetOk("sub_second_sampling"); ok {
		t, err := expandObjectFirewallGtpSubSecondSampling(d, v, "sub_second_sampling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-second-sampling"] = t
		}
	}

	if v, ok := d.GetOk("traffic_count_log"); ok {
		t, err := expandObjectFirewallGtpTrafficCountLog(d, v, "traffic_count_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-count-log"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_limit"); ok {
		t, err := expandObjectFirewallGtpTunnelLimit(d, v, "tunnel_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-limit"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_limit_log"); ok {
		t, err := expandObjectFirewallGtpTunnelLimitLog(d, v, "tunnel_limit_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-limit-log"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_timeout"); ok {
		t, err := expandObjectFirewallGtpTunnelTimeout(d, v, "tunnel_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-timeout"] = t
		}
	}

	if v, ok := d.GetOk("unknown_version_action"); ok {
		t, err := expandObjectFirewallGtpUnknownVersionAction(d, v, "unknown_version_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-version-action"] = t
		}
	}

	if v, ok := d.GetOk("user_plane_message_rate_limit"); ok {
		t, err := expandObjectFirewallGtpUserPlaneMessageRateLimit(d, v, "user_plane_message_rate_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-plane-message-rate-limit"] = t
		}
	}

	if v, ok := d.GetOk("warning_threshold"); ok {
		t, err := expandObjectFirewallGtpWarningThreshold(d, v, "warning_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warning-threshold"] = t
		}
	}

	return &obj, nil
}
