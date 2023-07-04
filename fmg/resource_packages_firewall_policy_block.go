// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePackagesFirewallPolicyBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallPolicyBlockCreate,
		Read:   resourcePackagesFirewallPolicyBlockRead,
		Update: resourcePackagesFirewallPolicyBlockUpdate,
		Delete: resourcePackagesFirewallPolicyBlockDelete,

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
			"pblock": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"_policy_block": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anti_replay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_redirect_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"best_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"block_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_exempt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capture_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_eif": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cgn_eim": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cgn_log_server_grp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_resource_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cgn_session_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cifs_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_log_fields": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"decrypted_traffic_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"delay_tcp_npu_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_copy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"devices": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"diffserv_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"diffservcode_rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dscp_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dscp_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dscp_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dsri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_shaping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email_collect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_filter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"firewall_session_dirty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fixedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsso_agent_for_ntlm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsso_groups": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"geoip_anycast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geoip_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gtp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"icap_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"identity_based_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_custom_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_custom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_src_custom_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_src_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_src_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_src_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"learning_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_vip_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"natinbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"natip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"natoutbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_service_dynamic": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"network_service_src_dynamic": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"np_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntlm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntlm_enabled_browsers": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ntlm_guest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"outbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"passive_wan_health_measurement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_ip_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permit_any_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_stun_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pfcp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_expiry_date": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"poolname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"poolname6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacemsg_override_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reputation_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reputation_direction6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reputation_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reputation_minimum6": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rtp_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rtp_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedule_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sctp_filter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"send_deny_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spamfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sgt": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"sgt_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_vendor_mac": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_mirror_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_mss_receiver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_mss_sender": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_session_without_syn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_timeout_pid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout_send_rst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_shaper_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"udp_timeout_pid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"videofilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_cos_fwd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_cos_rev": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"voip_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpn_dst_node": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"seq": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"vpn_src_node": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"seq": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"vpntunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"waf_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wanopt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wanopt_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wanopt_passive_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wanopt_peer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wanopt_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webcache_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webproxy_forward_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webproxy_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ztna_ems_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ztna_geo_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ztna_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wsso": &schema.Schema{
				Type:     schema.TypeString,
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

func resourcePackagesFirewallPolicyBlockCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	pblock := d.Get("pblock").(string)
	paradict["pblock"] = pblock

	obj, err := getObjectPackagesFirewallPolicyBlock(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallPolicy resource while getting object: %v", err)
	}

	v, err := c.CreatePackagesFirewallPolicyBlock(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallPolicy resource: %v", err)
	}

	if v != nil && v["policyid"] != nil {
		if vidn, ok := v["policyid"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourcePackagesFirewallPolicyBlockRead(d, m)
		} else {
			return fmt.Errorf("Error creating PackagesFirewallPolicy resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallPolicyBlockRead(d, m)
}

func resourcePackagesFirewallPolicyBlockUpdate(d *schema.ResourceData, m interface{}) error {
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

	pblock := d.Get("pblock").(string)
	paradict["pblock"] = pblock

	obj, err := getObjectPackagesFirewallPolicyBlock(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallPolicyBlock(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallPolicyBlockRead(d, m)
}

func resourcePackagesFirewallPolicyBlockDelete(d *schema.ResourceData, m interface{}) error {
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

	pblock := d.Get("pblock").(string)
	paradict["pblock"] = pblock

	err = c.DeletePackagesFirewallPolicyBlock(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallPolicyBlockRead(d *schema.ResourceData, m interface{}) error {
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

	pblock := d.Get("pblock").(string)
	if pblock == "" {
		pblock = importOptionChecking(m.(*FortiClient).Cfg, "pblock")
		if err = d.Set("pblock", pblock); err != nil {
			return fmt.Errorf("Error set params pblock: %v", err)
		}
	}
	paradict["pblock"] = pblock

	o, err := c.ReadPackagesFirewallPolicyBlock(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallPolicyBlock(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallPolicy resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallPolicyBlockPolicyBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockAntiReplay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockAppCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockAppGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenPackagesFirewallPolicyBlockApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockAuthPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockAuthRedirectAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockBestRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockBlockNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockCaptivePortalExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockCapturePacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockCgnEif(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockCgnEim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockCgnLogServerGrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockCgnResourceQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockCgnSessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockCifsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockCustomLogFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDelayTcpNpuSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDiffservCopy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDevices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDiffservForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDiffservReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDisclaimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDlpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDscpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDscpNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDscpValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDsri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockDstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockDstaddr6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockDynamicShaping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockEmailCollect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockFec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockFileFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockFirewallSessionDirty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockFixedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockFsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockFssoAgentForNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockFssoGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockGeoipAnycast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockGeoipMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockGlobalLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockGtpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockHttpPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockIcapProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockIdentityBasedRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceSrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceSrcId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetService6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetService6Custom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockInternetService6CustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockInternetService6Group(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockInternetService6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockInternetService6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetService6Src(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockInternetService6SrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockInternetService6SrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockInternetService6SrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockInternetService6SrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockInternetService6SrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockLearningMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockLogtrafficStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockMatchVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockMatchVipOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockMmsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockNatinbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockNatip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockNatoutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockNetworkServiceDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockNetworkServiceSrcDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockNpAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockNtlmEnabledBrowsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockNtlmGuest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockOutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPassiveWanHealthMeasurement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPermitAnyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPermitStunHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPfcpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPolicyExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPolicyExpiryDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPolicyOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPoolname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockPoolname6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockRadiusMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockReputationDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockReputationDirection6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockReputationMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockReputationMinimum6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockRsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockRtpAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockRtpNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockScheduleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSctpFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSendDenyPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convttlfloat642String(v)
}

func flattenPackagesFirewallPolicyBlockSpamfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSgt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenPackagesFirewallPolicyBlockSgtCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSrcVendorMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockSrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockSrcaddr6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockSshFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSshPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSslMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSslMirrorIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockSslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockTcpMssReceiver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockTcpMssSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockTcpSessionWithoutSyn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockTcpTimeoutPid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockTimeoutSendRst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockTosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockTosNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockTrafficShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockTrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockUdpTimeoutPid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockUrlCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallPolicyBlockUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVideofilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVlanCosFwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVlanCosRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVlanFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVoipProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVpnDstNode(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenPackagesFirewallPolicyBlockVpnDstNodeHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "PackagesFirewallPolicy-VpnDstNode-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenPackagesFirewallPolicyBlockVpnDstNodeSeq(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "PackagesFirewallPolicy-VpnDstNode-Seq")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenPackagesFirewallPolicyBlockVpnDstNodeSubnet(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "PackagesFirewallPolicy-VpnDstNode-Subnet")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenPackagesFirewallPolicyBlockVpnDstNodeHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVpnDstNodeSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVpnDstNodeSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVpnSrcNode(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenPackagesFirewallPolicyBlockVpnSrcNodeHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "PackagesFirewallPolicy-VpnSrcNode-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenPackagesFirewallPolicyBlockVpnSrcNodeSeq(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "PackagesFirewallPolicy-VpnSrcNode-Seq")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenPackagesFirewallPolicyBlockVpnSrcNodeSubnet(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "PackagesFirewallPolicy-VpnSrcNode-Subnet")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenPackagesFirewallPolicyBlockVpnSrcNodeHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVpnSrcNodeSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVpnSrcNodeSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockVpntunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWafProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWanopt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWanoptDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWanoptPassiveOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWanoptPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWanoptProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWccp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWebcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWebcacheHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWebproxyForwardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWebproxyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockZtnaEmsTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockZtnaGeoTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockZtnaStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicyBlockWsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallPolicyBlock(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_policy_block", flattenPackagesFirewallPolicyBlockPolicyBlock(o["_policy_block"], d, "_policy_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["_policy_block"], "PackagesFirewallPolicy-PolicyBlock"); ok {
			if err = d.Set("_policy_block", vv); err != nil {
				return fmt.Errorf("Error reading _policy_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _policy_block: %v", err)
		}
	}

	if err = d.Set("action", flattenPackagesFirewallPolicyBlockAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "PackagesFirewallPolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("anti_replay", flattenPackagesFirewallPolicyBlockAntiReplay(o["anti-replay"], d, "anti_replay")); err != nil {
		if vv, ok := fortiAPIPatch(o["anti-replay"], "PackagesFirewallPolicy-AntiReplay"); ok {
			if err = d.Set("anti_replay", vv); err != nil {
				return fmt.Errorf("Error reading anti_replay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anti_replay: %v", err)
		}
	}

	if err = d.Set("app_category", flattenPackagesFirewallPolicyBlockAppCategory(o["app-category"], d, "app_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-category"], "PackagesFirewallPolicy-AppCategory"); ok {
			if err = d.Set("app_category", vv); err != nil {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("app_group", flattenPackagesFirewallPolicyBlockAppGroup(o["app-group"], d, "app_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-group"], "PackagesFirewallPolicy-AppGroup"); ok {
			if err = d.Set("app_group", vv); err != nil {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_group: %v", err)
		}
	}

	if err = d.Set("application", flattenPackagesFirewallPolicyBlockApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "PackagesFirewallPolicy-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("application_list", flattenPackagesFirewallPolicyBlockApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "PackagesFirewallPolicy-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenPackagesFirewallPolicyBlockAuthCert(o["auth-cert"], d, "auth_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-cert"], "PackagesFirewallPolicy-AuthCert"); ok {
			if err = d.Set("auth_cert", vv); err != nil {
				return fmt.Errorf("Error reading auth_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_path", flattenPackagesFirewallPolicyBlockAuthPath(o["auth-path"], d, "auth_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-path"], "PackagesFirewallPolicy-AuthPath"); ok {
			if err = d.Set("auth_path", vv); err != nil {
				return fmt.Errorf("Error reading auth_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_path: %v", err)
		}
	}

	if err = d.Set("auth_redirect_addr", flattenPackagesFirewallPolicyBlockAuthRedirectAddr(o["auth-redirect-addr"], d, "auth_redirect_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-redirect-addr"], "PackagesFirewallPolicy-AuthRedirectAddr"); ok {
			if err = d.Set("auth_redirect_addr", vv); err != nil {
				return fmt.Errorf("Error reading auth_redirect_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_redirect_addr: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenPackagesFirewallPolicyBlockAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-asic-offload"], "PackagesFirewallPolicy-AutoAsicOffload"); ok {
			if err = d.Set("auto_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading auto_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenPackagesFirewallPolicyBlockAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "PackagesFirewallPolicy-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("best_route", flattenPackagesFirewallPolicyBlockBestRoute(o["best-route"], d, "best_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["best-route"], "PackagesFirewallPolicy-BestRoute"); ok {
			if err = d.Set("best_route", vv); err != nil {
				return fmt.Errorf("Error reading best_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading best_route: %v", err)
		}
	}

	if err = d.Set("block_notification", flattenPackagesFirewallPolicyBlockBlockNotification(o["block-notification"], d, "block_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-notification"], "PackagesFirewallPolicy-BlockNotification"); ok {
			if err = d.Set("block_notification", vv); err != nil {
				return fmt.Errorf("Error reading block_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_notification: %v", err)
		}
	}

	if err = d.Set("captive_portal_exempt", flattenPackagesFirewallPolicyBlockCaptivePortalExempt(o["captive-portal-exempt"], d, "captive_portal_exempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-exempt"], "PackagesFirewallPolicy-CaptivePortalExempt"); ok {
			if err = d.Set("captive_portal_exempt", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_exempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_exempt: %v", err)
		}
	}

	if err = d.Set("capture_packet", flattenPackagesFirewallPolicyBlockCapturePacket(o["capture-packet"], d, "capture_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["capture-packet"], "PackagesFirewallPolicy-CapturePacket"); ok {
			if err = d.Set("capture_packet", vv); err != nil {
				return fmt.Errorf("Error reading capture_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capture_packet: %v", err)
		}
	}

	if err = d.Set("cgn_eif", flattenPackagesFirewallPolicyBlockCgnEif(o["cgn-eif"], d, "cgn_eif")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-eif"], "PackagesFirewallPolicy-CgnEif"); ok {
			if err = d.Set("cgn_eif", vv); err != nil {
				return fmt.Errorf("Error reading cgn_eif: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_eif: %v", err)
		}
	}

	if err = d.Set("cgn_eim", flattenPackagesFirewallPolicyBlockCgnEim(o["cgn-eim"], d, "cgn_eim")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-eim"], "PackagesFirewallPolicy-CgnEim"); ok {
			if err = d.Set("cgn_eim", vv); err != nil {
				return fmt.Errorf("Error reading cgn_eim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_eim: %v", err)
		}
	}

	if err = d.Set("cgn_log_server_grp", flattenPackagesFirewallPolicyBlockCgnLogServerGrp(o["cgn-log-server-grp"], d, "cgn_log_server_grp")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-log-server-grp"], "PackagesFirewallPolicy-CgnLogServerGrp"); ok {
			if err = d.Set("cgn_log_server_grp", vv); err != nil {
				return fmt.Errorf("Error reading cgn_log_server_grp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_log_server_grp: %v", err)
		}
	}

	if err = d.Set("cgn_resource_quota", flattenPackagesFirewallPolicyBlockCgnResourceQuota(o["cgn-resource-quota"], d, "cgn_resource_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-resource-quota"], "PackagesFirewallPolicy-CgnResourceQuota"); ok {
			if err = d.Set("cgn_resource_quota", vv); err != nil {
				return fmt.Errorf("Error reading cgn_resource_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_resource_quota: %v", err)
		}
	}

	if err = d.Set("cgn_session_quota", flattenPackagesFirewallPolicyBlockCgnSessionQuota(o["cgn-session-quota"], d, "cgn_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-session-quota"], "PackagesFirewallPolicy-CgnSessionQuota"); ok {
			if err = d.Set("cgn_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading cgn_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_session_quota: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenPackagesFirewallPolicyBlockCifsProfile(o["cifs-profile"], d, "cifs_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["cifs-profile"], "PackagesFirewallPolicy-CifsProfile"); ok {
			if err = d.Set("cifs_profile", vv); err != nil {
				return fmt.Errorf("Error reading cifs_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("comments", flattenPackagesFirewallPolicyBlockComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesFirewallPolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("custom_log_fields", flattenPackagesFirewallPolicyBlockCustomLogFields(o["custom-log-fields"], d, "custom_log_fields")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-log-fields"], "PackagesFirewallPolicy-CustomLogFields"); ok {
			if err = d.Set("custom_log_fields", vv); err != nil {
				return fmt.Errorf("Error reading custom_log_fields: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_log_fields: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenPackagesFirewallPolicyBlockDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["decrypted-traffic-mirror"], "PackagesFirewallPolicy-DecryptedTrafficMirror"); ok {
			if err = d.Set("decrypted_traffic_mirror", vv); err != nil {
				return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("delay_tcp_npu_session", flattenPackagesFirewallPolicyBlockDelayTcpNpuSession(o["delay-tcp-npu-session"], d, "delay_tcp_npu_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["delay-tcp-npu-session"], "PackagesFirewallPolicy-DelayTcpNpuSession"); ok {
			if err = d.Set("delay_tcp_npu_session", vv); err != nil {
				return fmt.Errorf("Error reading delay_tcp_npu_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delay_tcp_npu_session: %v", err)
		}
	}

	if err = d.Set("diffserv_copy", flattenPackagesFirewallPolicyBlockDiffservCopy(o["diffserv-copy"], d, "diffserv_copy")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-copy"], "PackagesFirewallPolicy-DiffservCopy"); ok {
			if err = d.Set("diffserv_copy", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_copy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_copy: %v", err)
		}
	}

	if err = d.Set("devices", flattenPackagesFirewallPolicyBlockDevices(o["devices"], d, "devices")); err != nil {
		if vv, ok := fortiAPIPatch(o["devices"], "PackagesFirewallPolicy-Devices"); ok {
			if err = d.Set("devices", vv); err != nil {
				return fmt.Errorf("Error reading devices: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devices: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenPackagesFirewallPolicyBlockDiffservForward(o["diffserv-forward"], d, "diffserv_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-forward"], "PackagesFirewallPolicy-DiffservForward"); ok {
			if err = d.Set("diffserv_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenPackagesFirewallPolicyBlockDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-reverse"], "PackagesFirewallPolicy-DiffservReverse"); ok {
			if err = d.Set("diffserv_reverse", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenPackagesFirewallPolicyBlockDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-forward"], "PackagesFirewallPolicy-DiffservcodeForward"); ok {
			if err = d.Set("diffservcode_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenPackagesFirewallPolicyBlockDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-rev"], "PackagesFirewallPolicy-DiffservcodeRev"); ok {
			if err = d.Set("diffservcode_rev", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("disclaimer", flattenPackagesFirewallPolicyBlockDisclaimer(o["disclaimer"], d, "disclaimer")); err != nil {
		if vv, ok := fortiAPIPatch(o["disclaimer"], "PackagesFirewallPolicy-Disclaimer"); ok {
			if err = d.Set("disclaimer", vv); err != nil {
				return fmt.Errorf("Error reading disclaimer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disclaimer: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenPackagesFirewallPolicyBlockDlpProfile(o["dlp-profile"], d, "dlp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile"], "PackagesFirewallPolicy-DlpProfile"); ok {
			if err = d.Set("dlp_profile", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenPackagesFirewallPolicyBlockDlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "PackagesFirewallPolicy-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenPackagesFirewallPolicyBlockDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnsfilter-profile"], "PackagesFirewallPolicy-DnsfilterProfile"); ok {
			if err = d.Set("dnsfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("dscp_match", flattenPackagesFirewallPolicyBlockDscpMatch(o["dscp-match"], d, "dscp_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-match"], "PackagesFirewallPolicy-DscpMatch"); ok {
			if err = d.Set("dscp_match", vv); err != nil {
				return fmt.Errorf("Error reading dscp_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_match: %v", err)
		}
	}

	if err = d.Set("dscp_negate", flattenPackagesFirewallPolicyBlockDscpNegate(o["dscp-negate"], d, "dscp_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-negate"], "PackagesFirewallPolicy-DscpNegate"); ok {
			if err = d.Set("dscp_negate", vv); err != nil {
				return fmt.Errorf("Error reading dscp_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_negate: %v", err)
		}
	}

	if err = d.Set("dscp_value", flattenPackagesFirewallPolicyBlockDscpValue(o["dscp-value"], d, "dscp_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-value"], "PackagesFirewallPolicy-DscpValue"); ok {
			if err = d.Set("dscp_value", vv); err != nil {
				return fmt.Errorf("Error reading dscp_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_value: %v", err)
		}
	}

	if err = d.Set("dsri", flattenPackagesFirewallPolicyBlockDsri(o["dsri"], d, "dsri")); err != nil {
		if vv, ok := fortiAPIPatch(o["dsri"], "PackagesFirewallPolicy-Dsri"); ok {
			if err = d.Set("dsri", vv); err != nil {
				return fmt.Errorf("Error reading dsri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesFirewallPolicyBlockDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesFirewallPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenPackagesFirewallPolicyBlockDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr-negate"], "PackagesFirewallPolicy-DstaddrNegate"); ok {
			if err = d.Set("dstaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenPackagesFirewallPolicyBlockDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "PackagesFirewallPolicy-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstaddr6_negate", flattenPackagesFirewallPolicyBlockDstaddr6Negate(o["dstaddr6-negate"], d, "dstaddr6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6-negate"], "PackagesFirewallPolicy-Dstaddr6Negate"); ok {
			if err = d.Set("dstaddr6_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6_negate: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenPackagesFirewallPolicyBlockDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "PackagesFirewallPolicy-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("dynamic_shaping", flattenPackagesFirewallPolicyBlockDynamicShaping(o["dynamic-shaping"], d, "dynamic_shaping")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-shaping"], "PackagesFirewallPolicy-DynamicShaping"); ok {
			if err = d.Set("dynamic_shaping", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_shaping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_shaping: %v", err)
		}
	}

	if err = d.Set("email_collect", flattenPackagesFirewallPolicyBlockEmailCollect(o["email-collect"], d, "email_collect")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-collect"], "PackagesFirewallPolicy-EmailCollect"); ok {
			if err = d.Set("email_collect", vv); err != nil {
				return fmt.Errorf("Error reading email_collect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_collect: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenPackagesFirewallPolicyBlockEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "PackagesFirewallPolicy-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("fec", flattenPackagesFirewallPolicyBlockFec(o["fec"], d, "fec")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec"], "PackagesFirewallPolicy-Fec"); ok {
			if err = d.Set("fec", vv); err != nil {
				return fmt.Errorf("Error reading fec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenPackagesFirewallPolicyBlockFileFilterProfile(o["file-filter-profile"], d, "file_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-filter-profile"], "PackagesFirewallPolicy-FileFilterProfile"); ok {
			if err = d.Set("file_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading file_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("firewall_session_dirty", flattenPackagesFirewallPolicyBlockFirewallSessionDirty(o["firewall-session-dirty"], d, "firewall_session_dirty")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-session-dirty"], "PackagesFirewallPolicy-FirewallSessionDirty"); ok {
			if err = d.Set("firewall_session_dirty", vv); err != nil {
				return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
		}
	}

	if err = d.Set("fixedport", flattenPackagesFirewallPolicyBlockFixedport(o["fixedport"], d, "fixedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["fixedport"], "PackagesFirewallPolicy-Fixedport"); ok {
			if err = d.Set("fixedport", vv); err != nil {
				return fmt.Errorf("Error reading fixedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fixedport: %v", err)
		}
	}

	if err = d.Set("fsso", flattenPackagesFirewallPolicyBlockFsso(o["fsso"], d, "fsso")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso"], "PackagesFirewallPolicy-Fsso"); ok {
			if err = d.Set("fsso", vv); err != nil {
				return fmt.Errorf("Error reading fsso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso: %v", err)
		}
	}

	if err = d.Set("fsso_agent_for_ntlm", flattenPackagesFirewallPolicyBlockFssoAgentForNtlm(o["fsso-agent-for-ntlm"], d, "fsso_agent_for_ntlm")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-agent-for-ntlm"], "PackagesFirewallPolicy-FssoAgentForNtlm"); ok {
			if err = d.Set("fsso_agent_for_ntlm", vv); err != nil {
				return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if err = d.Set("fsso_groups", flattenPackagesFirewallPolicyBlockFssoGroups(o["fsso-groups"], d, "fsso_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-groups"], "PackagesFirewallPolicy-FssoGroups"); ok {
			if err = d.Set("fsso_groups", vv); err != nil {
				return fmt.Errorf("Error reading fsso_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_groups: %v", err)
		}
	}

	if err = d.Set("geoip_anycast", flattenPackagesFirewallPolicyBlockGeoipAnycast(o["geoip-anycast"], d, "geoip_anycast")); err != nil {
		if vv, ok := fortiAPIPatch(o["geoip-anycast"], "PackagesFirewallPolicy-GeoipAnycast"); ok {
			if err = d.Set("geoip_anycast", vv); err != nil {
				return fmt.Errorf("Error reading geoip_anycast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading geoip_anycast: %v", err)
		}
	}

	if err = d.Set("geoip_match", flattenPackagesFirewallPolicyBlockGeoipMatch(o["geoip-match"], d, "geoip_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["geoip-match"], "PackagesFirewallPolicy-GeoipMatch"); ok {
			if err = d.Set("geoip_match", vv); err != nil {
				return fmt.Errorf("Error reading geoip_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading geoip_match: %v", err)
		}
	}

	if err = d.Set("global_label", flattenPackagesFirewallPolicyBlockGlobalLabel(o["global-label"], d, "global_label")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-label"], "PackagesFirewallPolicy-GlobalLabel"); ok {
			if err = d.Set("global_label", vv); err != nil {
				return fmt.Errorf("Error reading global_label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("groups", flattenPackagesFirewallPolicyBlockGroups(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "PackagesFirewallPolicy-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("gtp_profile", flattenPackagesFirewallPolicyBlockGtpProfile(o["gtp-profile"], d, "gtp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-profile"], "PackagesFirewallPolicy-GtpProfile"); ok {
			if err = d.Set("gtp_profile", vv); err != nil {
				return fmt.Errorf("Error reading gtp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_profile: %v", err)
		}
	}

	if err = d.Set("http_policy_redirect", flattenPackagesFirewallPolicyBlockHttpPolicyRedirect(o["http-policy-redirect"], d, "http_policy_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-policy-redirect"], "PackagesFirewallPolicy-HttpPolicyRedirect"); ok {
			if err = d.Set("http_policy_redirect", vv); err != nil {
				return fmt.Errorf("Error reading http_policy_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_policy_redirect: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenPackagesFirewallPolicyBlockIcapProfile(o["icap-profile"], d, "icap_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-profile"], "PackagesFirewallPolicy-IcapProfile"); ok {
			if err = d.Set("icap_profile", vv); err != nil {
				return fmt.Errorf("Error reading icap_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("identity_based_route", flattenPackagesFirewallPolicyBlockIdentityBasedRoute(o["identity-based-route"], d, "identity_based_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["identity-based-route"], "PackagesFirewallPolicy-IdentityBasedRoute"); ok {
			if err = d.Set("identity_based_route", vv); err != nil {
				return fmt.Errorf("Error reading identity_based_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading identity_based_route: %v", err)
		}
	}

	if err = d.Set("inbound", flattenPackagesFirewallPolicyBlockInbound(o["inbound"], d, "inbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound"], "PackagesFirewallPolicy-Inbound"); ok {
			if err = d.Set("inbound", vv); err != nil {
				return fmt.Errorf("Error reading inbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenPackagesFirewallPolicyBlockInspectionMode(o["inspection-mode"], d, "inspection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspection-mode"], "PackagesFirewallPolicy-InspectionMode"); ok {
			if err = d.Set("inspection_mode", vv); err != nil {
				return fmt.Errorf("Error reading inspection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenPackagesFirewallPolicyBlockInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service"], "PackagesFirewallPolicy-InternetService"); ok {
			if err = d.Set("internet_service", vv); err != nil {
				return fmt.Errorf("Error reading internet_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenPackagesFirewallPolicyBlockInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom"], "PackagesFirewallPolicy-InternetServiceCustom"); ok {
			if err = d.Set("internet_service_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", flattenPackagesFirewallPolicyBlockInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom-group"], "PackagesFirewallPolicy-InternetServiceCustomGroup"); ok {
			if err = d.Set("internet_service_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_group", flattenPackagesFirewallPolicyBlockInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-group"], "PackagesFirewallPolicy-InternetServiceGroup"); ok {
			if err = d.Set("internet_service_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_name", flattenPackagesFirewallPolicyBlockInternetServiceName(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-name"], "PackagesFirewallPolicy-InternetServiceName"); ok {
			if err = d.Set("internet_service_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("internet_service_id", flattenPackagesFirewallPolicyBlockInternetServiceId(o["internet-service-id"], d, "internet_service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-id"], "PackagesFirewallPolicy-InternetServiceId"); ok {
			if err = d.Set("internet_service_id", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_id: %v", err)
		}
	}

	if err = d.Set("internet_service_negate", flattenPackagesFirewallPolicyBlockInternetServiceNegate(o["internet-service-negate"], d, "internet_service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-negate"], "PackagesFirewallPolicy-InternetServiceNegate"); ok {
			if err = d.Set("internet_service_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_src", flattenPackagesFirewallPolicyBlockInternetServiceSrc(o["internet-service-src"], d, "internet_service_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src"], "PackagesFirewallPolicy-InternetServiceSrc"); ok {
			if err = d.Set("internet_service_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom", flattenPackagesFirewallPolicyBlockInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom"], "PackagesFirewallPolicy-InternetServiceSrcCustom"); ok {
			if err = d.Set("internet_service_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom_group", flattenPackagesFirewallPolicyBlockInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom-group"], "PackagesFirewallPolicy-InternetServiceSrcCustomGroup"); ok {
			if err = d.Set("internet_service_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_group", flattenPackagesFirewallPolicyBlockInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-group"], "PackagesFirewallPolicy-InternetServiceSrcGroup"); ok {
			if err = d.Set("internet_service_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_name", flattenPackagesFirewallPolicyBlockInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-name"], "PackagesFirewallPolicy-InternetServiceSrcName"); ok {
			if err = d.Set("internet_service_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_name: %v", err)
		}
	}

	if err = d.Set("internet_service_src_id", flattenPackagesFirewallPolicyBlockInternetServiceSrcId(o["internet-service-src-id"], d, "internet_service_src_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-id"], "PackagesFirewallPolicy-InternetServiceSrcId"); ok {
			if err = d.Set("internet_service_src_id", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_id: %v", err)
		}
	}

	if err = d.Set("internet_service_src_negate", flattenPackagesFirewallPolicyBlockInternetServiceSrcNegate(o["internet-service-src-negate"], d, "internet_service_src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-negate"], "PackagesFirewallPolicy-InternetServiceSrcNegate"); ok {
			if err = d.Set("internet_service_src_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
		}
	}

	if err = d.Set("internet_service6", flattenPackagesFirewallPolicyBlockInternetService6(o["internet-service6"], d, "internet_service6")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6"], "PackagesFirewallPolicy-InternetService6"); ok {
			if err = d.Set("internet_service6", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6: %v", err)
		}
	}

	if err = d.Set("internet_service6_custom", flattenPackagesFirewallPolicyBlockInternetService6Custom(o["internet-service6-custom"], d, "internet_service6_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-custom"], "PackagesFirewallPolicy-InternetService6Custom"); ok {
			if err = d.Set("internet_service6_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_custom: %v", err)
		}
	}

	if err = d.Set("internet_service6_custom_group", flattenPackagesFirewallPolicyBlockInternetService6CustomGroup(o["internet-service6-custom-group"], d, "internet_service6_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-custom-group"], "PackagesFirewallPolicy-InternetService6CustomGroup"); ok {
			if err = d.Set("internet_service6_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_group", flattenPackagesFirewallPolicyBlockInternetService6Group(o["internet-service6-group"], d, "internet_service6_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-group"], "PackagesFirewallPolicy-InternetService6Group"); ok {
			if err = d.Set("internet_service6_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_name", flattenPackagesFirewallPolicyBlockInternetService6Name(o["internet-service6-name"], d, "internet_service6_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-name"], "PackagesFirewallPolicy-InternetService6Name"); ok {
			if err = d.Set("internet_service6_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_name: %v", err)
		}
	}

	if err = d.Set("internet_service6_negate", flattenPackagesFirewallPolicyBlockInternetService6Negate(o["internet-service6-negate"], d, "internet_service6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-negate"], "PackagesFirewallPolicy-InternetService6Negate"); ok {
			if err = d.Set("internet_service6_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_negate: %v", err)
		}
	}

	if err = d.Set("internet_service6_src", flattenPackagesFirewallPolicyBlockInternetService6Src(o["internet-service6-src"], d, "internet_service6_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src"], "PackagesFirewallPolicy-InternetService6Src"); ok {
			if err = d.Set("internet_service6_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_custom", flattenPackagesFirewallPolicyBlockInternetService6SrcCustom(o["internet-service6-src-custom"], d, "internet_service6_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-custom"], "PackagesFirewallPolicy-InternetService6SrcCustom"); ok {
			if err = d.Set("internet_service6_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_custom_group", flattenPackagesFirewallPolicyBlockInternetService6SrcCustomGroup(o["internet-service6-src-custom-group"], d, "internet_service6_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-custom-group"], "PackagesFirewallPolicy-InternetService6SrcCustomGroup"); ok {
			if err = d.Set("internet_service6_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_group", flattenPackagesFirewallPolicyBlockInternetService6SrcGroup(o["internet-service6-src-group"], d, "internet_service6_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-group"], "PackagesFirewallPolicy-InternetService6SrcGroup"); ok {
			if err = d.Set("internet_service6_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_name", flattenPackagesFirewallPolicyBlockInternetService6SrcName(o["internet-service6-src-name"], d, "internet_service6_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-name"], "PackagesFirewallPolicy-InternetService6SrcName"); ok {
			if err = d.Set("internet_service6_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_negate", flattenPackagesFirewallPolicyBlockInternetService6SrcNegate(o["internet-service6-src-negate"], d, "internet_service6_src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-negate"], "PackagesFirewallPolicy-InternetService6SrcNegate"); ok {
			if err = d.Set("internet_service6_src_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_negate: %v", err)
		}
	}

	if err = d.Set("ippool", flattenPackagesFirewallPolicyBlockIppool(o["ippool"], d, "ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["ippool"], "PackagesFirewallPolicy-Ippool"); ok {
			if err = d.Set("ippool", vv); err != nil {
				return fmt.Errorf("Error reading ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenPackagesFirewallPolicyBlockIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "PackagesFirewallPolicy-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("label", flattenPackagesFirewallPolicyBlockLabel(o["label"], d, "label")); err != nil {
		if vv, ok := fortiAPIPatch(o["label"], "PackagesFirewallPolicy-Label"); ok {
			if err = d.Set("label", vv); err != nil {
				return fmt.Errorf("Error reading label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("learning_mode", flattenPackagesFirewallPolicyBlockLearningMode(o["learning-mode"], d, "learning_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["learning-mode"], "PackagesFirewallPolicy-LearningMode"); ok {
			if err = d.Set("learning_mode", vv); err != nil {
				return fmt.Errorf("Error reading learning_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learning_mode: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenPackagesFirewallPolicyBlockLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "PackagesFirewallPolicy-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenPackagesFirewallPolicyBlockLogtrafficStart(o["logtraffic-start"], d, "logtraffic_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic-start"], "PackagesFirewallPolicy-LogtrafficStart"); ok {
			if err = d.Set("logtraffic_start", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("match_vip", flattenPackagesFirewallPolicyBlockMatchVip(o["match-vip"], d, "match_vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-vip"], "PackagesFirewallPolicy-MatchVip"); ok {
			if err = d.Set("match_vip", vv); err != nil {
				return fmt.Errorf("Error reading match_vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_vip: %v", err)
		}
	}

	if err = d.Set("match_vip_only", flattenPackagesFirewallPolicyBlockMatchVipOnly(o["match-vip-only"], d, "match_vip_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-vip-only"], "PackagesFirewallPolicy-MatchVipOnly"); ok {
			if err = d.Set("match_vip_only", vv); err != nil {
				return fmt.Errorf("Error reading match_vip_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_vip_only: %v", err)
		}
	}

	if err = d.Set("mms_profile", flattenPackagesFirewallPolicyBlockMmsProfile(o["mms-profile"], d, "mms_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-profile"], "PackagesFirewallPolicy-MmsProfile"); ok {
			if err = d.Set("mms_profile", vv); err != nil {
				return fmt.Errorf("Error reading mms_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_profile: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesFirewallPolicyBlockName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesFirewallPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat", flattenPackagesFirewallPolicyBlockNat(o["nat"], d, "nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat"], "PackagesFirewallPolicy-Nat"); ok {
			if err = d.Set("nat", vv); err != nil {
				return fmt.Errorf("Error reading nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("nat46", flattenPackagesFirewallPolicyBlockNat46(o["nat46"], d, "nat46")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46"], "PackagesFirewallPolicy-Nat46"); ok {
			if err = d.Set("nat46", vv); err != nil {
				return fmt.Errorf("Error reading nat46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("nat64", flattenPackagesFirewallPolicyBlockNat64(o["nat64"], d, "nat64")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64"], "PackagesFirewallPolicy-Nat64"); ok {
			if err = d.Set("nat64", vv); err != nil {
				return fmt.Errorf("Error reading nat64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("natinbound", flattenPackagesFirewallPolicyBlockNatinbound(o["natinbound"], d, "natinbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["natinbound"], "PackagesFirewallPolicy-Natinbound"); ok {
			if err = d.Set("natinbound", vv); err != nil {
				return fmt.Errorf("Error reading natinbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading natinbound: %v", err)
		}
	}

	if err = d.Set("natip", flattenPackagesFirewallPolicyBlockNatip(o["natip"], d, "natip")); err != nil {
		if vv, ok := fortiAPIPatch(o["natip"], "PackagesFirewallPolicy-Natip"); ok {
			if err = d.Set("natip", vv); err != nil {
				return fmt.Errorf("Error reading natip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading natip: %v", err)
		}
	}

	if err = d.Set("natoutbound", flattenPackagesFirewallPolicyBlockNatoutbound(o["natoutbound"], d, "natoutbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["natoutbound"], "PackagesFirewallPolicy-Natoutbound"); ok {
			if err = d.Set("natoutbound", vv); err != nil {
				return fmt.Errorf("Error reading natoutbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading natoutbound: %v", err)
		}
	}

	if err = d.Set("network_service_dynamic", flattenPackagesFirewallPolicyBlockNetworkServiceDynamic(o["network-service-dynamic"], d, "network_service_dynamic")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-service-dynamic"], "PackagesFirewallPolicy-NetworkServiceDynamic"); ok {
			if err = d.Set("network_service_dynamic", vv); err != nil {
				return fmt.Errorf("Error reading network_service_dynamic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_service_dynamic: %v", err)
		}
	}

	if err = d.Set("network_service_src_dynamic", flattenPackagesFirewallPolicyBlockNetworkServiceSrcDynamic(o["network-service-src-dynamic"], d, "network_service_src_dynamic")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-service-src-dynamic"], "PackagesFirewallPolicy-NetworkServiceSrcDynamic"); ok {
			if err = d.Set("network_service_src_dynamic", vv); err != nil {
				return fmt.Errorf("Error reading network_service_src_dynamic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_service_src_dynamic: %v", err)
		}
	}

	if err = d.Set("np_acceleration", flattenPackagesFirewallPolicyBlockNpAcceleration(o["np-acceleration"], d, "np_acceleration")); err != nil {
		if vv, ok := fortiAPIPatch(o["np-acceleration"], "PackagesFirewallPolicy-NpAcceleration"); ok {
			if err = d.Set("np_acceleration", vv); err != nil {
				return fmt.Errorf("Error reading np_acceleration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading np_acceleration: %v", err)
		}
	}

	if err = d.Set("ntlm", flattenPackagesFirewallPolicyBlockNtlm(o["ntlm"], d, "ntlm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntlm"], "PackagesFirewallPolicy-Ntlm"); ok {
			if err = d.Set("ntlm", vv); err != nil {
				return fmt.Errorf("Error reading ntlm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntlm: %v", err)
		}
	}

	if err = d.Set("ntlm_enabled_browsers", flattenPackagesFirewallPolicyBlockNtlmEnabledBrowsers(o["ntlm-enabled-browsers"], d, "ntlm_enabled_browsers")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntlm-enabled-browsers"], "PackagesFirewallPolicy-NtlmEnabledBrowsers"); ok {
			if err = d.Set("ntlm_enabled_browsers", vv); err != nil {
				return fmt.Errorf("Error reading ntlm_enabled_browsers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntlm_enabled_browsers: %v", err)
		}
	}

	if err = d.Set("ntlm_guest", flattenPackagesFirewallPolicyBlockNtlmGuest(o["ntlm-guest"], d, "ntlm_guest")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntlm-guest"], "PackagesFirewallPolicy-NtlmGuest"); ok {
			if err = d.Set("ntlm_guest", vv); err != nil {
				return fmt.Errorf("Error reading ntlm_guest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntlm_guest: %v", err)
		}
	}

	if err = d.Set("outbound", flattenPackagesFirewallPolicyBlockOutbound(o["outbound"], d, "outbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbound"], "PackagesFirewallPolicy-Outbound"); ok {
			if err = d.Set("outbound", vv); err != nil {
				return fmt.Errorf("Error reading outbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbound: %v", err)
		}
	}

	if err = d.Set("passive_wan_health_measurement", flattenPackagesFirewallPolicyBlockPassiveWanHealthMeasurement(o["passive-wan-health-measurement"], d, "passive_wan_health_measurement")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive-wan-health-measurement"], "PackagesFirewallPolicy-PassiveWanHealthMeasurement"); ok {
			if err = d.Set("passive_wan_health_measurement", vv); err != nil {
				return fmt.Errorf("Error reading passive_wan_health_measurement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive_wan_health_measurement: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenPackagesFirewallPolicyBlockPerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-ip-shaper"], "PackagesFirewallPolicy-PerIpShaper"); ok {
			if err = d.Set("per_ip_shaper", vv); err != nil {
				return fmt.Errorf("Error reading per_ip_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("permit_any_host", flattenPackagesFirewallPolicyBlockPermitAnyHost(o["permit-any-host"], d, "permit_any_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-any-host"], "PackagesFirewallPolicy-PermitAnyHost"); ok {
			if err = d.Set("permit_any_host", vv); err != nil {
				return fmt.Errorf("Error reading permit_any_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_any_host: %v", err)
		}
	}

	if err = d.Set("permit_stun_host", flattenPackagesFirewallPolicyBlockPermitStunHost(o["permit-stun-host"], d, "permit_stun_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-stun-host"], "PackagesFirewallPolicy-PermitStunHost"); ok {
			if err = d.Set("permit_stun_host", vv); err != nil {
				return fmt.Errorf("Error reading permit_stun_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_stun_host: %v", err)
		}
	}

	if err = d.Set("pfcp_profile", flattenPackagesFirewallPolicyBlockPfcpProfile(o["pfcp-profile"], d, "pfcp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["pfcp-profile"], "PackagesFirewallPolicy-PfcpProfile"); ok {
			if err = d.Set("pfcp_profile", vv); err != nil {
				return fmt.Errorf("Error reading pfcp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pfcp_profile: %v", err)
		}
	}

	if err = d.Set("policy_expiry", flattenPackagesFirewallPolicyBlockPolicyExpiry(o["policy-expiry"], d, "policy_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-expiry"], "PackagesFirewallPolicy-PolicyExpiry"); ok {
			if err = d.Set("policy_expiry", vv); err != nil {
				return fmt.Errorf("Error reading policy_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_expiry: %v", err)
		}
	}

	if err = d.Set("policy_expiry_date", flattenPackagesFirewallPolicyBlockPolicyExpiryDate(o["policy-expiry-date"], d, "policy_expiry_date")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-expiry-date"], "PackagesFirewallPolicy-PolicyExpiryDate"); ok {
			if err = d.Set("policy_expiry_date", vv); err != nil {
				return fmt.Errorf("Error reading policy_expiry_date: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_expiry_date: %v", err)
		}
	}

	if err = d.Set("policy_offload", flattenPackagesFirewallPolicyBlockPolicyOffload(o["policy-offload"], d, "policy_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-offload"], "PackagesFirewallPolicy-PolicyOffload"); ok {
			if err = d.Set("policy_offload", vv); err != nil {
				return fmt.Errorf("Error reading policy_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_offload: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesFirewallPolicyBlockPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesFirewallPolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("poolname", flattenPackagesFirewallPolicyBlockPoolname(o["poolname"], d, "poolname")); err != nil {
		if vv, ok := fortiAPIPatch(o["poolname"], "PackagesFirewallPolicy-Poolname"); ok {
			if err = d.Set("poolname", vv); err != nil {
				return fmt.Errorf("Error reading poolname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poolname: %v", err)
		}
	}

	if err = d.Set("poolname6", flattenPackagesFirewallPolicyBlockPoolname6(o["poolname6"], d, "poolname6")); err != nil {
		if vv, ok := fortiAPIPatch(o["poolname6"], "PackagesFirewallPolicy-Poolname6"); ok {
			if err = d.Set("poolname6", vv); err != nil {
				return fmt.Errorf("Error reading poolname6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poolname6: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenPackagesFirewallPolicyBlockProfileGroup(o["profile-group"], d, "profile_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-group"], "PackagesFirewallPolicy-ProfileGroup"); ok {
			if err = d.Set("profile_group", vv); err != nil {
				return fmt.Errorf("Error reading profile_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenPackagesFirewallPolicyBlockProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-protocol-options"], "PackagesFirewallPolicy-ProfileProtocolOptions"); ok {
			if err = d.Set("profile_protocol_options", vv); err != nil {
				return fmt.Errorf("Error reading profile_protocol_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenPackagesFirewallPolicyBlockProfileType(o["profile-type"], d, "profile_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-type"], "PackagesFirewallPolicy-ProfileType"); ok {
			if err = d.Set("profile_type", vv); err != nil {
				return fmt.Errorf("Error reading profile_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_bypass", flattenPackagesFirewallPolicyBlockRadiusMacAuthBypass(o["radius-mac-auth-bypass"], d, "radius_mac_auth_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-bypass"], "PackagesFirewallPolicy-RadiusMacAuthBypass"); ok {
			if err = d.Set("radius_mac_auth_bypass", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("redirect_url", flattenPackagesFirewallPolicyBlockRedirectUrl(o["redirect-url"], d, "redirect_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-url"], "PackagesFirewallPolicy-RedirectUrl"); ok {
			if err = d.Set("redirect_url", vv); err != nil {
				return fmt.Errorf("Error reading redirect_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_url: %v", err)
		}
	}

	if err = d.Set("replacemsg_override_group", flattenPackagesFirewallPolicyBlockReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-override-group"], "PackagesFirewallPolicy-ReplacemsgOverrideGroup"); ok {
			if err = d.Set("replacemsg_override_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("reputation_direction", flattenPackagesFirewallPolicyBlockReputationDirection(o["reputation-direction"], d, "reputation_direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["reputation-direction"], "PackagesFirewallPolicy-ReputationDirection"); ok {
			if err = d.Set("reputation_direction", vv); err != nil {
				return fmt.Errorf("Error reading reputation_direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reputation_direction: %v", err)
		}
	}

	if err = d.Set("reputation_direction6", flattenPackagesFirewallPolicyBlockReputationDirection6(o["reputation-direction6"], d, "reputation_direction6")); err != nil {
		if vv, ok := fortiAPIPatch(o["reputation-direction6"], "PackagesFirewallPolicy-ReputationDirection6"); ok {
			if err = d.Set("reputation_direction6", vv); err != nil {
				return fmt.Errorf("Error reading reputation_direction6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reputation_direction6: %v", err)
		}
	}

	if err = d.Set("reputation_minimum", flattenPackagesFirewallPolicyBlockReputationMinimum(o["reputation-minimum"], d, "reputation_minimum")); err != nil {
		if vv, ok := fortiAPIPatch(o["reputation-minimum"], "PackagesFirewallPolicy-ReputationMinimum"); ok {
			if err = d.Set("reputation_minimum", vv); err != nil {
				return fmt.Errorf("Error reading reputation_minimum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reputation_minimum: %v", err)
		}
	}

	if err = d.Set("reputation_minimum6", flattenPackagesFirewallPolicyBlockReputationMinimum6(o["reputation-minimum6"], d, "reputation_minimum6")); err != nil {
		if vv, ok := fortiAPIPatch(o["reputation-minimum6"], "PackagesFirewallPolicy-ReputationMinimum6"); ok {
			if err = d.Set("reputation_minimum6", vv); err != nil {
				return fmt.Errorf("Error reading reputation_minimum6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reputation_minimum6: %v", err)
		}
	}

	if err = d.Set("rsso", flattenPackagesFirewallPolicyBlockRsso(o["rsso"], d, "rsso")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso"], "PackagesFirewallPolicy-Rsso"); ok {
			if err = d.Set("rsso", vv); err != nil {
				return fmt.Errorf("Error reading rsso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso: %v", err)
		}
	}

	if err = d.Set("rtp_addr", flattenPackagesFirewallPolicyBlockRtpAddr(o["rtp-addr"], d, "rtp_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtp-addr"], "PackagesFirewallPolicy-RtpAddr"); ok {
			if err = d.Set("rtp_addr", vv); err != nil {
				return fmt.Errorf("Error reading rtp_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtp_addr: %v", err)
		}
	}

	if err = d.Set("rtp_nat", flattenPackagesFirewallPolicyBlockRtpNat(o["rtp-nat"], d, "rtp_nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtp-nat"], "PackagesFirewallPolicy-RtpNat"); ok {
			if err = d.Set("rtp_nat", vv); err != nil {
				return fmt.Errorf("Error reading rtp_nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtp_nat: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenPackagesFirewallPolicyBlockScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-botnet-connections"], "PackagesFirewallPolicy-ScanBotnetConnections"); ok {
			if err = d.Set("scan_botnet_connections", vv); err != nil {
				return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("schedule", flattenPackagesFirewallPolicyBlockSchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "PackagesFirewallPolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("schedule_timeout", flattenPackagesFirewallPolicyBlockScheduleTimeout(o["schedule-timeout"], d, "schedule_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule-timeout"], "PackagesFirewallPolicy-ScheduleTimeout"); ok {
			if err = d.Set("schedule_timeout", vv); err != nil {
				return fmt.Errorf("Error reading schedule_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule_timeout: %v", err)
		}
	}

	if err = d.Set("sctp_filter_profile", flattenPackagesFirewallPolicyBlockSctpFilterProfile(o["sctp-filter-profile"], d, "sctp_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-filter-profile"], "PackagesFirewallPolicy-SctpFilterProfile"); ok {
			if err = d.Set("sctp_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", flattenPackagesFirewallPolicyBlockSendDenyPacket(o["send-deny-packet"], d, "send_deny_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-deny-packet"], "PackagesFirewallPolicy-SendDenyPacket"); ok {
			if err = d.Set("send_deny_packet", vv); err != nil {
				return fmt.Errorf("Error reading send_deny_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("service", flattenPackagesFirewallPolicyBlockService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "PackagesFirewallPolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenPackagesFirewallPolicyBlockServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-negate"], "PackagesFirewallPolicy-ServiceNegate"); ok {
			if err = d.Set("service_negate", vv); err != nil {
				return fmt.Errorf("Error reading service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("session_ttl", flattenPackagesFirewallPolicyBlockSessionTtl(o["session-ttl"], d, "session_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-ttl"], "PackagesFirewallPolicy-SessionTtl"); ok {
			if err = d.Set("session_ttl", vv); err != nil {
				return fmt.Errorf("Error reading session_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", flattenPackagesFirewallPolicyBlockSpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["spamfilter-profile"], "PackagesFirewallPolicy-SpamfilterProfile"); ok {
			if err = d.Set("spamfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading spamfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("sgt", flattenPackagesFirewallPolicyBlockSgt(o["sgt"], d, "sgt")); err != nil {
		if vv, ok := fortiAPIPatch(o["sgt"], "PackagesFirewallPolicy-Sgt"); ok {
			if err = d.Set("sgt", vv); err != nil {
				return fmt.Errorf("Error reading sgt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sgt: %v", err)
		}
	}

	if err = d.Set("sgt_check", flattenPackagesFirewallPolicyBlockSgtCheck(o["sgt-check"], d, "sgt_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["sgt-check"], "PackagesFirewallPolicy-SgtCheck"); ok {
			if err = d.Set("sgt_check", vv); err != nil {
				return fmt.Errorf("Error reading sgt_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sgt_check: %v", err)
		}
	}

	if err = d.Set("src_vendor_mac", flattenPackagesFirewallPolicyBlockSrcVendorMac(o["src-vendor-mac"], d, "src_vendor_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-vendor-mac"], "PackagesFirewallPolicy-SrcVendorMac"); ok {
			if err = d.Set("src_vendor_mac", vv); err != nil {
				return fmt.Errorf("Error reading src_vendor_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_vendor_mac: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesFirewallPolicyBlockSrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesFirewallPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenPackagesFirewallPolicyBlockSrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr-negate"], "PackagesFirewallPolicy-SrcaddrNegate"); ok {
			if err = d.Set("srcaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenPackagesFirewallPolicyBlockSrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "PackagesFirewallPolicy-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcaddr6_negate", flattenPackagesFirewallPolicyBlockSrcaddr6Negate(o["srcaddr6-negate"], d, "srcaddr6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6-negate"], "PackagesFirewallPolicy-Srcaddr6Negate"); ok {
			if err = d.Set("srcaddr6_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6_negate: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenPackagesFirewallPolicyBlockSrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "PackagesFirewallPolicy-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenPackagesFirewallPolicyBlockSshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-filter-profile"], "PackagesFirewallPolicy-SshFilterProfile"); ok {
			if err = d.Set("ssh_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("ssh_policy_redirect", flattenPackagesFirewallPolicyBlockSshPolicyRedirect(o["ssh-policy-redirect"], d, "ssh_policy_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-policy-redirect"], "PackagesFirewallPolicy-SshPolicyRedirect"); ok {
			if err = d.Set("ssh_policy_redirect", vv); err != nil {
				return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ssl_mirror", flattenPackagesFirewallPolicyBlockSslMirror(o["ssl-mirror"], d, "ssl_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mirror"], "PackagesFirewallPolicy-SslMirror"); ok {
			if err = d.Set("ssl_mirror", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mirror: %v", err)
		}
	}

	if err = d.Set("ssl_mirror_intf", flattenPackagesFirewallPolicyBlockSslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mirror-intf"], "PackagesFirewallPolicy-SslMirrorIntf"); ok {
			if err = d.Set("ssl_mirror_intf", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenPackagesFirewallPolicyBlockSslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-ssh-profile"], "PackagesFirewallPolicy-SslSshProfile"); ok {
			if err = d.Set("ssl_ssh_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallPolicyBlockStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tcp_mss_receiver", flattenPackagesFirewallPolicyBlockTcpMssReceiver(o["tcp-mss-receiver"], d, "tcp_mss_receiver")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-mss-receiver"], "PackagesFirewallPolicy-TcpMssReceiver"); ok {
			if err = d.Set("tcp_mss_receiver", vv); err != nil {
				return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
		}
	}

	if err = d.Set("tcp_mss_sender", flattenPackagesFirewallPolicyBlockTcpMssSender(o["tcp-mss-sender"], d, "tcp_mss_sender")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-mss-sender"], "PackagesFirewallPolicy-TcpMssSender"); ok {
			if err = d.Set("tcp_mss_sender", vv); err != nil {
				return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
		}
	}

	if err = d.Set("tcp_session_without_syn", flattenPackagesFirewallPolicyBlockTcpSessionWithoutSyn(o["tcp-session-without-syn"], d, "tcp_session_without_syn")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-session-without-syn"], "PackagesFirewallPolicy-TcpSessionWithoutSyn"); ok {
			if err = d.Set("tcp_session_without_syn", vv); err != nil {
				return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
		}
	}

	if err = d.Set("tcp_timeout_pid", flattenPackagesFirewallPolicyBlockTcpTimeoutPid(o["tcp-timeout-pid"], d, "tcp_timeout_pid")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-timeout-pid"], "PackagesFirewallPolicy-TcpTimeoutPid"); ok {
			if err = d.Set("tcp_timeout_pid", vv); err != nil {
				return fmt.Errorf("Error reading tcp_timeout_pid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_timeout_pid: %v", err)
		}
	}

	if err = d.Set("timeout_send_rst", flattenPackagesFirewallPolicyBlockTimeoutSendRst(o["timeout-send-rst"], d, "timeout_send_rst")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-send-rst"], "PackagesFirewallPolicy-TimeoutSendRst"); ok {
			if err = d.Set("timeout_send_rst", vv); err != nil {
				return fmt.Errorf("Error reading timeout_send_rst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_send_rst: %v", err)
		}
	}

	if err = d.Set("tos", flattenPackagesFirewallPolicyBlockTos(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "PackagesFirewallPolicy-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenPackagesFirewallPolicyBlockTosMask(o["tos-mask"], d, "tos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-mask"], "PackagesFirewallPolicy-TosMask"); ok {
			if err = d.Set("tos_mask", vv); err != nil {
				return fmt.Errorf("Error reading tos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("tos_negate", flattenPackagesFirewallPolicyBlockTosNegate(o["tos-negate"], d, "tos_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-negate"], "PackagesFirewallPolicy-TosNegate"); ok {
			if err = d.Set("tos_negate", vv); err != nil {
				return fmt.Errorf("Error reading tos_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_negate: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenPackagesFirewallPolicyBlockTrafficShaper(o["traffic-shaper"], d, "traffic_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper"], "PackagesFirewallPolicy-TrafficShaper"); ok {
			if err = d.Set("traffic_shaper", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenPackagesFirewallPolicyBlockTrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper-reverse"], "PackagesFirewallPolicy-TrafficShaperReverse"); ok {
			if err = d.Set("traffic_shaper_reverse", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("udp_timeout_pid", flattenPackagesFirewallPolicyBlockUdpTimeoutPid(o["udp-timeout-pid"], d, "udp_timeout_pid")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-timeout-pid"], "PackagesFirewallPolicy-UdpTimeoutPid"); ok {
			if err = d.Set("udp_timeout_pid", vv); err != nil {
				return fmt.Errorf("Error reading udp_timeout_pid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_timeout_pid: %v", err)
		}
	}

	if err = d.Set("url_category", flattenPackagesFirewallPolicyBlockUrlCategory(o["url-category"], d, "url_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-category"], "PackagesFirewallPolicy-UrlCategory"); ok {
			if err = d.Set("url_category", vv); err != nil {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_category: %v", err)
		}
	}

	if err = d.Set("users", flattenPackagesFirewallPolicyBlockUsers(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "PackagesFirewallPolicy-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenPackagesFirewallPolicyBlockUtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-status"], "PackagesFirewallPolicy-UtmStatus"); ok {
			if err = d.Set("utm_status", vv); err != nil {
				return fmt.Errorf("Error reading utm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenPackagesFirewallPolicyBlockUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesFirewallPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("videofilter_profile", flattenPackagesFirewallPolicyBlockVideofilterProfile(o["videofilter-profile"], d, "videofilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["videofilter-profile"], "PackagesFirewallPolicy-VideofilterProfile"); ok {
			if err = d.Set("videofilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading videofilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading videofilter_profile: %v", err)
		}
	}

	if err = d.Set("vlan_cos_fwd", flattenPackagesFirewallPolicyBlockVlanCosFwd(o["vlan-cos-fwd"], d, "vlan_cos_fwd")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-cos-fwd"], "PackagesFirewallPolicy-VlanCosFwd"); ok {
			if err = d.Set("vlan_cos_fwd", vv); err != nil {
				return fmt.Errorf("Error reading vlan_cos_fwd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_cos_fwd: %v", err)
		}
	}

	if err = d.Set("vlan_cos_rev", flattenPackagesFirewallPolicyBlockVlanCosRev(o["vlan-cos-rev"], d, "vlan_cos_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-cos-rev"], "PackagesFirewallPolicy-VlanCosRev"); ok {
			if err = d.Set("vlan_cos_rev", vv); err != nil {
				return fmt.Errorf("Error reading vlan_cos_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_cos_rev: %v", err)
		}
	}

	if err = d.Set("vlan_filter", flattenPackagesFirewallPolicyBlockVlanFilter(o["vlan-filter"], d, "vlan_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-filter"], "PackagesFirewallPolicy-VlanFilter"); ok {
			if err = d.Set("vlan_filter", vv); err != nil {
				return fmt.Errorf("Error reading vlan_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_filter: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenPackagesFirewallPolicyBlockVoipProfile(o["voip-profile"], d, "voip_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-profile"], "PackagesFirewallPolicy-VoipProfile"); ok {
			if err = d.Set("voip_profile", vv); err != nil {
				return fmt.Errorf("Error reading voip_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vpn_dst_node", flattenPackagesFirewallPolicyBlockVpnDstNode(o["vpn_dst_node"], d, "vpn_dst_node")); err != nil {
			if vv, ok := fortiAPIPatch(o["vpn_dst_node"], "PackagesFirewallPolicy-VpnDstNode"); ok {
				if err = d.Set("vpn_dst_node", vv); err != nil {
					return fmt.Errorf("Error reading vpn_dst_node: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vpn_dst_node: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vpn_dst_node"); ok {
			if err = d.Set("vpn_dst_node", flattenPackagesFirewallPolicyBlockVpnDstNode(o["vpn_dst_node"], d, "vpn_dst_node")); err != nil {
				if vv, ok := fortiAPIPatch(o["vpn_dst_node"], "PackagesFirewallPolicy-VpnDstNode"); ok {
					if err = d.Set("vpn_dst_node", vv); err != nil {
						return fmt.Errorf("Error reading vpn_dst_node: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vpn_dst_node: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("vpn_src_node", flattenPackagesFirewallPolicyBlockVpnSrcNode(o["vpn_src_node"], d, "vpn_src_node")); err != nil {
			if vv, ok := fortiAPIPatch(o["vpn_src_node"], "PackagesFirewallPolicy-VpnSrcNode"); ok {
				if err = d.Set("vpn_src_node", vv); err != nil {
					return fmt.Errorf("Error reading vpn_src_node: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vpn_src_node: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vpn_src_node"); ok {
			if err = d.Set("vpn_src_node", flattenPackagesFirewallPolicyBlockVpnSrcNode(o["vpn_src_node"], d, "vpn_src_node")); err != nil {
				if vv, ok := fortiAPIPatch(o["vpn_src_node"], "PackagesFirewallPolicy-VpnSrcNode"); ok {
					if err = d.Set("vpn_src_node", vv); err != nil {
						return fmt.Errorf("Error reading vpn_src_node: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vpn_src_node: %v", err)
				}
			}
		}
	}

	if err = d.Set("vpntunnel", flattenPackagesFirewallPolicyBlockVpntunnel(o["vpntunnel"], d, "vpntunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpntunnel"], "PackagesFirewallPolicy-Vpntunnel"); ok {
			if err = d.Set("vpntunnel", vv); err != nil {
				return fmt.Errorf("Error reading vpntunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpntunnel: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenPackagesFirewallPolicyBlockWafProfile(o["waf-profile"], d, "waf_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["waf-profile"], "PackagesFirewallPolicy-WafProfile"); ok {
			if err = d.Set("waf_profile", vv); err != nil {
				return fmt.Errorf("Error reading waf_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("wanopt", flattenPackagesFirewallPolicyBlockWanopt(o["wanopt"], d, "wanopt")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt"], "PackagesFirewallPolicy-Wanopt"); ok {
			if err = d.Set("wanopt", vv); err != nil {
				return fmt.Errorf("Error reading wanopt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt: %v", err)
		}
	}

	if err = d.Set("wanopt_detection", flattenPackagesFirewallPolicyBlockWanoptDetection(o["wanopt-detection"], d, "wanopt_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt-detection"], "PackagesFirewallPolicy-WanoptDetection"); ok {
			if err = d.Set("wanopt_detection", vv); err != nil {
				return fmt.Errorf("Error reading wanopt_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt_detection: %v", err)
		}
	}

	if err = d.Set("wanopt_passive_opt", flattenPackagesFirewallPolicyBlockWanoptPassiveOpt(o["wanopt-passive-opt"], d, "wanopt_passive_opt")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt-passive-opt"], "PackagesFirewallPolicy-WanoptPassiveOpt"); ok {
			if err = d.Set("wanopt_passive_opt", vv); err != nil {
				return fmt.Errorf("Error reading wanopt_passive_opt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt_passive_opt: %v", err)
		}
	}

	if err = d.Set("wanopt_peer", flattenPackagesFirewallPolicyBlockWanoptPeer(o["wanopt-peer"], d, "wanopt_peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt-peer"], "PackagesFirewallPolicy-WanoptPeer"); ok {
			if err = d.Set("wanopt_peer", vv); err != nil {
				return fmt.Errorf("Error reading wanopt_peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt_peer: %v", err)
		}
	}

	if err = d.Set("wanopt_profile", flattenPackagesFirewallPolicyBlockWanoptProfile(o["wanopt-profile"], d, "wanopt_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt-profile"], "PackagesFirewallPolicy-WanoptProfile"); ok {
			if err = d.Set("wanopt_profile", vv); err != nil {
				return fmt.Errorf("Error reading wanopt_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt_profile: %v", err)
		}
	}

	if err = d.Set("wccp", flattenPackagesFirewallPolicyBlockWccp(o["wccp"], d, "wccp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wccp"], "PackagesFirewallPolicy-Wccp"); ok {
			if err = d.Set("wccp", vv); err != nil {
				return fmt.Errorf("Error reading wccp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("webcache", flattenPackagesFirewallPolicyBlockWebcache(o["webcache"], d, "webcache")); err != nil {
		if vv, ok := fortiAPIPatch(o["webcache"], "PackagesFirewallPolicy-Webcache"); ok {
			if err = d.Set("webcache", vv); err != nil {
				return fmt.Errorf("Error reading webcache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webcache: %v", err)
		}
	}

	if err = d.Set("webcache_https", flattenPackagesFirewallPolicyBlockWebcacheHttps(o["webcache-https"], d, "webcache_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["webcache-https"], "PackagesFirewallPolicy-WebcacheHttps"); ok {
			if err = d.Set("webcache_https", vv); err != nil {
				return fmt.Errorf("Error reading webcache_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webcache_https: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenPackagesFirewallPolicyBlockWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "PackagesFirewallPolicy-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("webproxy_forward_server", flattenPackagesFirewallPolicyBlockWebproxyForwardServer(o["webproxy-forward-server"], d, "webproxy_forward_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["webproxy-forward-server"], "PackagesFirewallPolicy-WebproxyForwardServer"); ok {
			if err = d.Set("webproxy_forward_server", vv); err != nil {
				return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", flattenPackagesFirewallPolicyBlockWebproxyProfile(o["webproxy-profile"], d, "webproxy_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webproxy-profile"], "PackagesFirewallPolicy-WebproxyProfile"); ok {
			if err = d.Set("webproxy_profile", vv); err != nil {
				return fmt.Errorf("Error reading webproxy_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	if err = d.Set("ztna_ems_tag", flattenPackagesFirewallPolicyBlockZtnaEmsTag(o["ztna-ems-tag"], d, "ztna_ems_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-ems-tag"], "PackagesFirewallPolicy-ZtnaEmsTag"); ok {
			if err = d.Set("ztna_ems_tag", vv); err != nil {
				return fmt.Errorf("Error reading ztna_ems_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_ems_tag: %v", err)
		}
	}

	if err = d.Set("ztna_geo_tag", flattenPackagesFirewallPolicyBlockZtnaGeoTag(o["ztna-geo-tag"], d, "ztna_geo_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-geo-tag"], "PackagesFirewallPolicy-ZtnaGeoTag"); ok {
			if err = d.Set("ztna_geo_tag", vv); err != nil {
				return fmt.Errorf("Error reading ztna_geo_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_geo_tag: %v", err)
		}
	}

	if err = d.Set("ztna_status", flattenPackagesFirewallPolicyBlockZtnaStatus(o["ztna-status"], d, "ztna_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-status"], "PackagesFirewallPolicy-ZtnaStatus"); ok {
			if err = d.Set("ztna_status", vv); err != nil {
				return fmt.Errorf("Error reading ztna_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_status: %v", err)
		}
	}

	if err = d.Set("wsso", flattenPackagesFirewallPolicyBlockWsso(o["wsso"], d, "wsso")); err != nil {
		if vv, ok := fortiAPIPatch(o["wsso"], "PackagesFirewallPolicy-Wsso"); ok {
			if err = d.Set("wsso", vv); err != nil {
				return fmt.Errorf("Error reading wsso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wsso: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallPolicyBlockFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallPolicyBlockPolicyBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockAntiReplay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockAppCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockAppGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockAuthCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockAuthPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockAuthRedirectAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockAvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockBestRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockBlockNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockCaptivePortalExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockCapturePacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockCgnEif(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockCgnEim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockCgnLogServerGrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockCgnResourceQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockCgnSessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockCifsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockCustomLogFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDelayTcpNpuSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDiffservCopy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDevices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDiffservForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDiffservReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDiffservcodeForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDiffservcodeRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDisclaimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDlpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDscpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDscpNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDscpValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDsri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockDstaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockDstaddr6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockDynamicShaping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockEmailCollect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockFec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockFileFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockFirewallSessionDirty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockFixedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockFsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockFssoAgentForNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockFssoGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockGeoipAnycast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockGeoipMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockGlobalLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockGtpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockHttpPolicyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockIcapProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockIdentityBasedRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceSrcId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetService6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetService6Custom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockInternetService6CustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockInternetService6Group(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockInternetService6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockInternetService6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetService6Src(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockInternetService6SrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockInternetService6SrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockInternetService6SrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockInternetService6SrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockInternetService6SrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockLearningMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockLogtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockLogtrafficStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockMatchVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockMatchVipOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockMmsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockNatinbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockNatip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockNatoutbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockNetworkServiceDynamic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockNetworkServiceSrcDynamic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockNpAcceleration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockNtlmEnabledBrowsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockNtlmGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockOutbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPassiveWanHealthMeasurement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPermitAnyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPermitStunHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPfcpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPolicyExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPolicyExpiryDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPolicyOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPoolname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockPoolname6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockRadiusMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockReplacemsgOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockReputationDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockReputationDirection6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockReputationMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockReputationMinimum6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockRsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockRtpAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockRtpNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockScheduleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSctpFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSendDenyPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSpamfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSgt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallPolicyBlockSgtCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSrcVendorMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockSrcaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockSrcaddr6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockSshFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSshPolicyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSslMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSslMirrorIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockSslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockTcpMssReceiver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockTcpMssSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockTcpSessionWithoutSyn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockTcpTimeoutPid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockTimeoutSendRst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockTosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockTosNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockTrafficShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockTrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockUdpTimeoutPid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockUrlCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallPolicyBlockUtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVideofilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVlanCosFwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVlanCosRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVlanFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVoipProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVpnDstNode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandPackagesFirewallPolicyBlockVpnDstNodeHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandPackagesFirewallPolicyBlockVpnDstNodeSeq(d, i["seq"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandPackagesFirewallPolicyBlockVpnDstNodeSubnet(d, i["subnet"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandPackagesFirewallPolicyBlockVpnDstNodeHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVpnDstNodeSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVpnDstNodeSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVpnSrcNode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandPackagesFirewallPolicyBlockVpnSrcNodeHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandPackagesFirewallPolicyBlockVpnSrcNodeSeq(d, i["seq"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandPackagesFirewallPolicyBlockVpnSrcNodeSubnet(d, i["subnet"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandPackagesFirewallPolicyBlockVpnSrcNodeHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVpnSrcNodeSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVpnSrcNodeSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockVpntunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWafProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWanopt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWanoptDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWanoptPassiveOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWanoptPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWanoptProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWccp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWebcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWebcacheHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWebproxyForwardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWebproxyProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockZtnaEmsTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockZtnaGeoTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockZtnaStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicyBlockWsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallPolicyBlock(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_policy_block"); ok || d.HasChange("_policy_block") {
		t, err := expandPackagesFirewallPolicyBlockPolicyBlock(d, v, "_policy_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_policy_block"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandPackagesFirewallPolicyBlockAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("anti_replay"); ok || d.HasChange("anti_replay") {
		t, err := expandPackagesFirewallPolicyBlockAntiReplay(d, v, "anti_replay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anti-replay"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok || d.HasChange("app_category") {
		t, err := expandPackagesFirewallPolicyBlockAppCategory(d, v, "app_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok || d.HasChange("app_group") {
		t, err := expandPackagesFirewallPolicyBlockAppGroup(d, v, "app_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandPackagesFirewallPolicyBlockApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandPackagesFirewallPolicyBlockApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok || d.HasChange("auth_cert") {
		t, err := expandPackagesFirewallPolicyBlockAuthCert(d, v, "auth_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_path"); ok || d.HasChange("auth_path") {
		t, err := expandPackagesFirewallPolicyBlockAuthPath(d, v, "auth_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-path"] = t
		}
	}

	if v, ok := d.GetOk("auth_redirect_addr"); ok || d.HasChange("auth_redirect_addr") {
		t, err := expandPackagesFirewallPolicyBlockAuthRedirectAddr(d, v, "auth_redirect_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-redirect-addr"] = t
		}
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok || d.HasChange("auto_asic_offload") {
		t, err := expandPackagesFirewallPolicyBlockAutoAsicOffload(d, v, "auto_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandPackagesFirewallPolicyBlockAvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("best_route"); ok || d.HasChange("best_route") {
		t, err := expandPackagesFirewallPolicyBlockBestRoute(d, v, "best_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["best-route"] = t
		}
	}

	if v, ok := d.GetOk("block_notification"); ok || d.HasChange("block_notification") {
		t, err := expandPackagesFirewallPolicyBlockBlockNotification(d, v, "block_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-notification"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_exempt"); ok || d.HasChange("captive_portal_exempt") {
		t, err := expandPackagesFirewallPolicyBlockCaptivePortalExempt(d, v, "captive_portal_exempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-exempt"] = t
		}
	}

	if v, ok := d.GetOk("capture_packet"); ok || d.HasChange("capture_packet") {
		t, err := expandPackagesFirewallPolicyBlockCapturePacket(d, v, "capture_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capture-packet"] = t
		}
	}

	if v, ok := d.GetOk("cgn_eif"); ok || d.HasChange("cgn_eif") {
		t, err := expandPackagesFirewallPolicyBlockCgnEif(d, v, "cgn_eif")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-eif"] = t
		}
	}

	if v, ok := d.GetOk("cgn_eim"); ok || d.HasChange("cgn_eim") {
		t, err := expandPackagesFirewallPolicyBlockCgnEim(d, v, "cgn_eim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-eim"] = t
		}
	}

	if v, ok := d.GetOk("cgn_log_server_grp"); ok || d.HasChange("cgn_log_server_grp") {
		t, err := expandPackagesFirewallPolicyBlockCgnLogServerGrp(d, v, "cgn_log_server_grp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-log-server-grp"] = t
		}
	}

	if v, ok := d.GetOk("cgn_resource_quota"); ok || d.HasChange("cgn_resource_quota") {
		t, err := expandPackagesFirewallPolicyBlockCgnResourceQuota(d, v, "cgn_resource_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-resource-quota"] = t
		}
	}

	if v, ok := d.GetOk("cgn_session_quota"); ok || d.HasChange("cgn_session_quota") {
		t, err := expandPackagesFirewallPolicyBlockCgnSessionQuota(d, v, "cgn_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("cifs_profile"); ok || d.HasChange("cifs_profile") {
		t, err := expandPackagesFirewallPolicyBlockCifsProfile(d, v, "cifs_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandPackagesFirewallPolicyBlockComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("custom_log_fields"); ok || d.HasChange("custom_log_fields") {
		t, err := expandPackagesFirewallPolicyBlockCustomLogFields(d, v, "custom_log_fields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-log-fields"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok || d.HasChange("decrypted_traffic_mirror") {
		t, err := expandPackagesFirewallPolicyBlockDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("delay_tcp_npu_session"); ok || d.HasChange("delay_tcp_npu_session") {
		t, err := expandPackagesFirewallPolicyBlockDelayTcpNpuSession(d, v, "delay_tcp_npu_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay-tcp-npu-session"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_copy"); ok || d.HasChange("diffserv_copy") {
		t, err := expandPackagesFirewallPolicyBlockDiffservCopy(d, v, "diffserv_copy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-copy"] = t
		}
	}

	if v, ok := d.GetOk("devices"); ok || d.HasChange("devices") {
		t, err := expandPackagesFirewallPolicyBlockDevices(d, v, "devices")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devices"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok || d.HasChange("diffserv_forward") {
		t, err := expandPackagesFirewallPolicyBlockDiffservForward(d, v, "diffserv_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok || d.HasChange("diffserv_reverse") {
		t, err := expandPackagesFirewallPolicyBlockDiffservReverse(d, v, "diffserv_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok || d.HasChange("diffservcode_forward") {
		t, err := expandPackagesFirewallPolicyBlockDiffservcodeForward(d, v, "diffservcode_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok || d.HasChange("diffservcode_rev") {
		t, err := expandPackagesFirewallPolicyBlockDiffservcodeRev(d, v, "diffservcode_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOk("disclaimer"); ok || d.HasChange("disclaimer") {
		t, err := expandPackagesFirewallPolicyBlockDisclaimer(d, v, "disclaimer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile"); ok || d.HasChange("dlp_profile") {
		t, err := expandPackagesFirewallPolicyBlockDlpProfile(d, v, "dlp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandPackagesFirewallPolicyBlockDlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok || d.HasChange("dnsfilter_profile") {
		t, err := expandPackagesFirewallPolicyBlockDnsfilterProfile(d, v, "dnsfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dscp_match"); ok || d.HasChange("dscp_match") {
		t, err := expandPackagesFirewallPolicyBlockDscpMatch(d, v, "dscp_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-match"] = t
		}
	}

	if v, ok := d.GetOk("dscp_negate"); ok || d.HasChange("dscp_negate") {
		t, err := expandPackagesFirewallPolicyBlockDscpNegate(d, v, "dscp_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-negate"] = t
		}
	}

	if v, ok := d.GetOk("dscp_value"); ok || d.HasChange("dscp_value") {
		t, err := expandPackagesFirewallPolicyBlockDscpValue(d, v, "dscp_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-value"] = t
		}
	}

	if v, ok := d.GetOk("dsri"); ok || d.HasChange("dsri") {
		t, err := expandPackagesFirewallPolicyBlockDsri(d, v, "dsri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandPackagesFirewallPolicyBlockDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok || d.HasChange("dstaddr_negate") {
		t, err := expandPackagesFirewallPolicyBlockDstaddrNegate(d, v, "dstaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandPackagesFirewallPolicyBlockDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6_negate"); ok || d.HasChange("dstaddr6_negate") {
		t, err := expandPackagesFirewallPolicyBlockDstaddr6Negate(d, v, "dstaddr6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandPackagesFirewallPolicyBlockDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_shaping"); ok || d.HasChange("dynamic_shaping") {
		t, err := expandPackagesFirewallPolicyBlockDynamicShaping(d, v, "dynamic_shaping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-shaping"] = t
		}
	}

	if v, ok := d.GetOk("email_collect"); ok || d.HasChange("email_collect") {
		t, err := expandPackagesFirewallPolicyBlockEmailCollect(d, v, "email_collect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-collect"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok || d.HasChange("emailfilter_profile") {
		t, err := expandPackagesFirewallPolicyBlockEmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("fec"); ok || d.HasChange("fec") {
		t, err := expandPackagesFirewallPolicyBlockFec(d, v, "fec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec"] = t
		}
	}

	if v, ok := d.GetOk("file_filter_profile"); ok || d.HasChange("file_filter_profile") {
		t, err := expandPackagesFirewallPolicyBlockFileFilterProfile(d, v, "file_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("firewall_session_dirty"); ok || d.HasChange("firewall_session_dirty") {
		t, err := expandPackagesFirewallPolicyBlockFirewallSessionDirty(d, v, "firewall_session_dirty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-session-dirty"] = t
		}
	}

	if v, ok := d.GetOk("fixedport"); ok || d.HasChange("fixedport") {
		t, err := expandPackagesFirewallPolicyBlockFixedport(d, v, "fixedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixedport"] = t
		}
	}

	if v, ok := d.GetOk("fsso"); ok || d.HasChange("fsso") {
		t, err := expandPackagesFirewallPolicyBlockFsso(d, v, "fsso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso"] = t
		}
	}

	if v, ok := d.GetOk("fsso_agent_for_ntlm"); ok || d.HasChange("fsso_agent_for_ntlm") {
		t, err := expandPackagesFirewallPolicyBlockFssoAgentForNtlm(d, v, "fsso_agent_for_ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-agent-for-ntlm"] = t
		}
	}

	if v, ok := d.GetOk("fsso_groups"); ok || d.HasChange("fsso_groups") {
		t, err := expandPackagesFirewallPolicyBlockFssoGroups(d, v, "fsso_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-groups"] = t
		}
	}

	if v, ok := d.GetOk("geoip_anycast"); ok || d.HasChange("geoip_anycast") {
		t, err := expandPackagesFirewallPolicyBlockGeoipAnycast(d, v, "geoip_anycast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geoip-anycast"] = t
		}
	}

	if v, ok := d.GetOk("geoip_match"); ok || d.HasChange("geoip_match") {
		t, err := expandPackagesFirewallPolicyBlockGeoipMatch(d, v, "geoip_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geoip-match"] = t
		}
	}

	if v, ok := d.GetOk("global_label"); ok || d.HasChange("global_label") {
		t, err := expandPackagesFirewallPolicyBlockGlobalLabel(d, v, "global_label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-label"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandPackagesFirewallPolicyBlockGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("gtp_profile"); ok || d.HasChange("gtp_profile") {
		t, err := expandPackagesFirewallPolicyBlockGtpProfile(d, v, "gtp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-profile"] = t
		}
	}

	if v, ok := d.GetOk("http_policy_redirect"); ok || d.HasChange("http_policy_redirect") {
		t, err := expandPackagesFirewallPolicyBlockHttpPolicyRedirect(d, v, "http_policy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("icap_profile"); ok || d.HasChange("icap_profile") {
		t, err := expandPackagesFirewallPolicyBlockIcapProfile(d, v, "icap_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	}

	if v, ok := d.GetOk("identity_based_route"); ok || d.HasChange("identity_based_route") {
		t, err := expandPackagesFirewallPolicyBlockIdentityBasedRoute(d, v, "identity_based_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identity-based-route"] = t
		}
	}

	if v, ok := d.GetOk("inbound"); ok || d.HasChange("inbound") {
		t, err := expandPackagesFirewallPolicyBlockInbound(d, v, "inbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok || d.HasChange("inspection_mode") {
		t, err := expandPackagesFirewallPolicyBlockInspectionMode(d, v, "inspection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok || d.HasChange("internet_service") {
		t, err := expandPackagesFirewallPolicyBlockInternetService(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceCustom(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok || d.HasChange("internet_service_custom_group") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceCustomGroup(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok || d.HasChange("internet_service_group") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceGroup(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok || d.HasChange("internet_service_name") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceName(d, v, "internet_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok || d.HasChange("internet_service_id") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceId(d, v, "internet_service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_negate"); ok || d.HasChange("internet_service_negate") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceNegate(d, v, "internet_service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok || d.HasChange("internet_service_src") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceSrc(d, v, "internet_service_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok || d.HasChange("internet_service_src_custom") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceSrcCustom(d, v, "internet_service_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok || d.HasChange("internet_service_src_custom_group") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok || d.HasChange("internet_service_src_group") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceSrcGroup(d, v, "internet_service_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_name"); ok || d.HasChange("internet_service_src_name") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceSrcName(d, v, "internet_service_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_id"); ok || d.HasChange("internet_service_src_id") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceSrcId(d, v, "internet_service_src_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_negate"); ok || d.HasChange("internet_service_src_negate") {
		t, err := expandPackagesFirewallPolicyBlockInternetServiceSrcNegate(d, v, "internet_service_src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6"); ok || d.HasChange("internet_service6") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6(d, v, "internet_service6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom"); ok || d.HasChange("internet_service6_custom") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6Custom(d, v, "internet_service6_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom_group"); ok || d.HasChange("internet_service6_custom_group") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6CustomGroup(d, v, "internet_service6_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_group"); ok || d.HasChange("internet_service6_group") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6Group(d, v, "internet_service6_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_name"); ok || d.HasChange("internet_service6_name") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6Name(d, v, "internet_service6_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_negate"); ok || d.HasChange("internet_service6_negate") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6Negate(d, v, "internet_service6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src"); ok || d.HasChange("internet_service6_src") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6Src(d, v, "internet_service6_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom"); ok || d.HasChange("internet_service6_src_custom") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6SrcCustom(d, v, "internet_service6_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom_group"); ok || d.HasChange("internet_service6_src_custom_group") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6SrcCustomGroup(d, v, "internet_service6_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_group"); ok || d.HasChange("internet_service6_src_group") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6SrcGroup(d, v, "internet_service6_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_name"); ok || d.HasChange("internet_service6_src_name") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6SrcName(d, v, "internet_service6_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_negate"); ok || d.HasChange("internet_service6_src_negate") {
		t, err := expandPackagesFirewallPolicyBlockInternetService6SrcNegate(d, v, "internet_service6_src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("ippool"); ok || d.HasChange("ippool") {
		t, err := expandPackagesFirewallPolicyBlockIppool(d, v, "ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandPackagesFirewallPolicyBlockIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		t, err := expandPackagesFirewallPolicyBlockLabel(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("learning_mode"); ok || d.HasChange("learning_mode") {
		t, err := expandPackagesFirewallPolicyBlockLearningMode(d, v, "learning_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning-mode"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandPackagesFirewallPolicyBlockLogtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok || d.HasChange("logtraffic_start") {
		t, err := expandPackagesFirewallPolicyBlockLogtrafficStart(d, v, "logtraffic_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	}

	if v, ok := d.GetOk("match_vip"); ok || d.HasChange("match_vip") {
		t, err := expandPackagesFirewallPolicyBlockMatchVip(d, v, "match_vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-vip"] = t
		}
	}

	if v, ok := d.GetOk("match_vip_only"); ok || d.HasChange("match_vip_only") {
		t, err := expandPackagesFirewallPolicyBlockMatchVipOnly(d, v, "match_vip_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-vip-only"] = t
		}
	}

	if v, ok := d.GetOk("mms_profile"); ok || d.HasChange("mms_profile") {
		t, err := expandPackagesFirewallPolicyBlockMmsProfile(d, v, "mms_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-profile"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesFirewallPolicyBlockName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok || d.HasChange("nat") {
		t, err := expandPackagesFirewallPolicyBlockNat(d, v, "nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok || d.HasChange("nat46") {
		t, err := expandPackagesFirewallPolicyBlockNat46(d, v, "nat46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok || d.HasChange("nat64") {
		t, err := expandPackagesFirewallPolicyBlockNat64(d, v, "nat64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("natinbound"); ok || d.HasChange("natinbound") {
		t, err := expandPackagesFirewallPolicyBlockNatinbound(d, v, "natinbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natinbound"] = t
		}
	}

	if v, ok := d.GetOk("natip"); ok || d.HasChange("natip") {
		t, err := expandPackagesFirewallPolicyBlockNatip(d, v, "natip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natip"] = t
		}
	}

	if v, ok := d.GetOk("natoutbound"); ok || d.HasChange("natoutbound") {
		t, err := expandPackagesFirewallPolicyBlockNatoutbound(d, v, "natoutbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natoutbound"] = t
		}
	}

	if v, ok := d.GetOk("network_service_dynamic"); ok || d.HasChange("network_service_dynamic") {
		t, err := expandPackagesFirewallPolicyBlockNetworkServiceDynamic(d, v, "network_service_dynamic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-service-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("network_service_src_dynamic"); ok || d.HasChange("network_service_src_dynamic") {
		t, err := expandPackagesFirewallPolicyBlockNetworkServiceSrcDynamic(d, v, "network_service_src_dynamic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-service-src-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("np_acceleration"); ok || d.HasChange("np_acceleration") {
		t, err := expandPackagesFirewallPolicyBlockNpAcceleration(d, v, "np_acceleration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("ntlm"); ok || d.HasChange("ntlm") {
		t, err := expandPackagesFirewallPolicyBlockNtlm(d, v, "ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm"] = t
		}
	}

	if v, ok := d.GetOk("ntlm_enabled_browsers"); ok || d.HasChange("ntlm_enabled_browsers") {
		t, err := expandPackagesFirewallPolicyBlockNtlmEnabledBrowsers(d, v, "ntlm_enabled_browsers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm-enabled-browsers"] = t
		}
	}

	if v, ok := d.GetOk("ntlm_guest"); ok || d.HasChange("ntlm_guest") {
		t, err := expandPackagesFirewallPolicyBlockNtlmGuest(d, v, "ntlm_guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm-guest"] = t
		}
	}

	if v, ok := d.GetOk("outbound"); ok || d.HasChange("outbound") {
		t, err := expandPackagesFirewallPolicyBlockOutbound(d, v, "outbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbound"] = t
		}
	}

	if v, ok := d.GetOk("passive_wan_health_measurement"); ok || d.HasChange("passive_wan_health_measurement") {
		t, err := expandPackagesFirewallPolicyBlockPassiveWanHealthMeasurement(d, v, "passive_wan_health_measurement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-wan-health-measurement"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok || d.HasChange("per_ip_shaper") {
		t, err := expandPackagesFirewallPolicyBlockPerIpShaper(d, v, "per_ip_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("permit_any_host"); ok || d.HasChange("permit_any_host") {
		t, err := expandPackagesFirewallPolicyBlockPermitAnyHost(d, v, "permit_any_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-any-host"] = t
		}
	}

	if v, ok := d.GetOk("permit_stun_host"); ok || d.HasChange("permit_stun_host") {
		t, err := expandPackagesFirewallPolicyBlockPermitStunHost(d, v, "permit_stun_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-stun-host"] = t
		}
	}

	if v, ok := d.GetOk("pfcp_profile"); ok || d.HasChange("pfcp_profile") {
		t, err := expandPackagesFirewallPolicyBlockPfcpProfile(d, v, "pfcp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pfcp-profile"] = t
		}
	}

	if v, ok := d.GetOk("policy_expiry"); ok || d.HasChange("policy_expiry") {
		t, err := expandPackagesFirewallPolicyBlockPolicyExpiry(d, v, "policy_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-expiry"] = t
		}
	}

	if v, ok := d.GetOk("policy_expiry_date"); ok || d.HasChange("policy_expiry_date") {
		t, err := expandPackagesFirewallPolicyBlockPolicyExpiryDate(d, v, "policy_expiry_date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-expiry-date"] = t
		}
	}

	if v, ok := d.GetOk("policy_offload"); ok || d.HasChange("policy_offload") {
		t, err := expandPackagesFirewallPolicyBlockPolicyOffload(d, v, "policy_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-offload"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandPackagesFirewallPolicyBlockPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("poolname"); ok || d.HasChange("poolname") {
		t, err := expandPackagesFirewallPolicyBlockPoolname(d, v, "poolname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname"] = t
		}
	}

	if v, ok := d.GetOk("poolname6"); ok || d.HasChange("poolname6") {
		t, err := expandPackagesFirewallPolicyBlockPoolname6(d, v, "poolname6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname6"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok || d.HasChange("profile_group") {
		t, err := expandPackagesFirewallPolicyBlockProfileGroup(d, v, "profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok || d.HasChange("profile_protocol_options") {
		t, err := expandPackagesFirewallPolicyBlockProfileProtocolOptions(d, v, "profile_protocol_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok || d.HasChange("profile_type") {
		t, err := expandPackagesFirewallPolicyBlockProfileType(d, v, "profile_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_bypass"); ok || d.HasChange("radius_mac_auth_bypass") {
		t, err := expandPackagesFirewallPolicyBlockRadiusMacAuthBypass(d, v, "radius_mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("redirect_url"); ok || d.HasChange("redirect_url") {
		t, err := expandPackagesFirewallPolicyBlockRedirectUrl(d, v, "redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_override_group"); ok || d.HasChange("replacemsg_override_group") {
		t, err := expandPackagesFirewallPolicyBlockReplacemsgOverrideGroup(d, v, "replacemsg_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-override-group"] = t
		}
	}

	if v, ok := d.GetOk("reputation_direction"); ok || d.HasChange("reputation_direction") {
		t, err := expandPackagesFirewallPolicyBlockReputationDirection(d, v, "reputation_direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-direction"] = t
		}
	}

	if v, ok := d.GetOk("reputation_direction6"); ok || d.HasChange("reputation_direction6") {
		t, err := expandPackagesFirewallPolicyBlockReputationDirection6(d, v, "reputation_direction6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-direction6"] = t
		}
	}

	if v, ok := d.GetOk("reputation_minimum"); ok || d.HasChange("reputation_minimum") {
		t, err := expandPackagesFirewallPolicyBlockReputationMinimum(d, v, "reputation_minimum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-minimum"] = t
		}
	}

	if v, ok := d.GetOk("reputation_minimum6"); ok || d.HasChange("reputation_minimum6") {
		t, err := expandPackagesFirewallPolicyBlockReputationMinimum6(d, v, "reputation_minimum6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-minimum6"] = t
		}
	}

	if v, ok := d.GetOk("rsso"); ok || d.HasChange("rsso") {
		t, err := expandPackagesFirewallPolicyBlockRsso(d, v, "rsso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso"] = t
		}
	}

	if v, ok := d.GetOk("rtp_addr"); ok || d.HasChange("rtp_addr") {
		t, err := expandPackagesFirewallPolicyBlockRtpAddr(d, v, "rtp_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtp-addr"] = t
		}
	}

	if v, ok := d.GetOk("rtp_nat"); ok || d.HasChange("rtp_nat") {
		t, err := expandPackagesFirewallPolicyBlockRtpNat(d, v, "rtp_nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtp-nat"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok || d.HasChange("scan_botnet_connections") {
		t, err := expandPackagesFirewallPolicyBlockScanBotnetConnections(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandPackagesFirewallPolicyBlockSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("schedule_timeout"); ok || d.HasChange("schedule_timeout") {
		t, err := expandPackagesFirewallPolicyBlockScheduleTimeout(d, v, "schedule_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-timeout"] = t
		}
	}

	if v, ok := d.GetOk("sctp_filter_profile"); ok || d.HasChange("sctp_filter_profile") {
		t, err := expandPackagesFirewallPolicyBlockSctpFilterProfile(d, v, "sctp_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("send_deny_packet"); ok || d.HasChange("send_deny_packet") {
		t, err := expandPackagesFirewallPolicyBlockSendDenyPacket(d, v, "send_deny_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-deny-packet"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandPackagesFirewallPolicyBlockService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok || d.HasChange("service_negate") {
		t, err := expandPackagesFirewallPolicyBlockServiceNegate(d, v, "service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("session_ttl"); ok || d.HasChange("session_ttl") {
		t, err := expandPackagesFirewallPolicyBlockSessionTtl(d, v, "session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("spamfilter_profile"); ok || d.HasChange("spamfilter_profile") {
		t, err := expandPackagesFirewallPolicyBlockSpamfilterProfile(d, v, "spamfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("sgt"); ok || d.HasChange("sgt") {
		t, err := expandPackagesFirewallPolicyBlockSgt(d, v, "sgt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgt"] = t
		}
	}

	if v, ok := d.GetOk("sgt_check"); ok || d.HasChange("sgt_check") {
		t, err := expandPackagesFirewallPolicyBlockSgtCheck(d, v, "sgt_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgt-check"] = t
		}
	}

	if v, ok := d.GetOk("src_vendor_mac"); ok || d.HasChange("src_vendor_mac") {
		t, err := expandPackagesFirewallPolicyBlockSrcVendorMac(d, v, "src_vendor_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-vendor-mac"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandPackagesFirewallPolicyBlockSrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok || d.HasChange("srcaddr_negate") {
		t, err := expandPackagesFirewallPolicyBlockSrcaddrNegate(d, v, "srcaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandPackagesFirewallPolicyBlockSrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6_negate"); ok || d.HasChange("srcaddr6_negate") {
		t, err := expandPackagesFirewallPolicyBlockSrcaddr6Negate(d, v, "srcaddr6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandPackagesFirewallPolicyBlockSrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok || d.HasChange("ssh_filter_profile") {
		t, err := expandPackagesFirewallPolicyBlockSshFilterProfile(d, v, "ssh_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssh_policy_redirect"); ok || d.HasChange("ssh_policy_redirect") {
		t, err := expandPackagesFirewallPolicyBlockSshPolicyRedirect(d, v, "ssh_policy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mirror"); ok || d.HasChange("ssl_mirror") {
		t, err := expandPackagesFirewallPolicyBlockSslMirror(d, v, "ssl_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mirror_intf"); ok || d.HasChange("ssl_mirror_intf") {
		t, err := expandPackagesFirewallPolicyBlockSslMirrorIntf(d, v, "ssl_mirror_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror-intf"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok || d.HasChange("ssl_ssh_profile") {
		t, err := expandPackagesFirewallPolicyBlockSslSshProfile(d, v, "ssl_ssh_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesFirewallPolicyBlockStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss_receiver"); ok || d.HasChange("tcp_mss_receiver") {
		t, err := expandPackagesFirewallPolicyBlockTcpMssReceiver(d, v, "tcp_mss_receiver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-receiver"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss_sender"); ok || d.HasChange("tcp_mss_sender") {
		t, err := expandPackagesFirewallPolicyBlockTcpMssSender(d, v, "tcp_mss_sender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-sender"] = t
		}
	}

	if v, ok := d.GetOk("tcp_session_without_syn"); ok || d.HasChange("tcp_session_without_syn") {
		t, err := expandPackagesFirewallPolicyBlockTcpSessionWithoutSyn(d, v, "tcp_session_without_syn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-session-without-syn"] = t
		}
	}

	if v, ok := d.GetOk("tcp_timeout_pid"); ok || d.HasChange("tcp_timeout_pid") {
		t, err := expandPackagesFirewallPolicyBlockTcpTimeoutPid(d, v, "tcp_timeout_pid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-timeout-pid"] = t
		}
	}

	if v, ok := d.GetOk("timeout_send_rst"); ok || d.HasChange("timeout_send_rst") {
		t, err := expandPackagesFirewallPolicyBlockTimeoutSendRst(d, v, "timeout_send_rst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-send-rst"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandPackagesFirewallPolicyBlockTos(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok || d.HasChange("tos_mask") {
		t, err := expandPackagesFirewallPolicyBlockTosMask(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("tos_negate"); ok || d.HasChange("tos_negate") {
		t, err := expandPackagesFirewallPolicyBlockTosNegate(d, v, "tos_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-negate"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok || d.HasChange("traffic_shaper") {
		t, err := expandPackagesFirewallPolicyBlockTrafficShaper(d, v, "traffic_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok || d.HasChange("traffic_shaper_reverse") {
		t, err := expandPackagesFirewallPolicyBlockTrafficShaperReverse(d, v, "traffic_shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("udp_timeout_pid"); ok || d.HasChange("udp_timeout_pid") {
		t, err := expandPackagesFirewallPolicyBlockUdpTimeoutPid(d, v, "udp_timeout_pid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-timeout-pid"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok || d.HasChange("url_category") {
		t, err := expandPackagesFirewallPolicyBlockUrlCategory(d, v, "url_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandPackagesFirewallPolicyBlockUsers(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok || d.HasChange("utm_status") {
		t, err := expandPackagesFirewallPolicyBlockUtmStatus(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandPackagesFirewallPolicyBlockUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("videofilter_profile"); ok || d.HasChange("videofilter_profile") {
		t, err := expandPackagesFirewallPolicyBlockVideofilterProfile(d, v, "videofilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("vlan_cos_fwd"); ok || d.HasChange("vlan_cos_fwd") {
		t, err := expandPackagesFirewallPolicyBlockVlanCosFwd(d, v, "vlan_cos_fwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-fwd"] = t
		}
	}

	if v, ok := d.GetOk("vlan_cos_rev"); ok || d.HasChange("vlan_cos_rev") {
		t, err := expandPackagesFirewallPolicyBlockVlanCosRev(d, v, "vlan_cos_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-rev"] = t
		}
	}

	if v, ok := d.GetOk("vlan_filter"); ok || d.HasChange("vlan_filter") {
		t, err := expandPackagesFirewallPolicyBlockVlanFilter(d, v, "vlan_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-filter"] = t
		}
	}

	if v, ok := d.GetOk("voip_profile"); ok || d.HasChange("voip_profile") {
		t, err := expandPackagesFirewallPolicyBlockVoipProfile(d, v, "voip_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("vpn_dst_node"); ok || d.HasChange("vpn_dst_node") {
		t, err := expandPackagesFirewallPolicyBlockVpnDstNode(d, v, "vpn_dst_node")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn_dst_node"] = t
		}
	}

	if v, ok := d.GetOk("vpn_src_node"); ok || d.HasChange("vpn_src_node") {
		t, err := expandPackagesFirewallPolicyBlockVpnSrcNode(d, v, "vpn_src_node")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn_src_node"] = t
		}
	}

	if v, ok := d.GetOk("vpntunnel"); ok || d.HasChange("vpntunnel") {
		t, err := expandPackagesFirewallPolicyBlockVpntunnel(d, v, "vpntunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpntunnel"] = t
		}
	}

	if v, ok := d.GetOk("waf_profile"); ok || d.HasChange("waf_profile") {
		t, err := expandPackagesFirewallPolicyBlockWafProfile(d, v, "waf_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("wanopt"); ok || d.HasChange("wanopt") {
		t, err := expandPackagesFirewallPolicyBlockWanopt(d, v, "wanopt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_detection"); ok || d.HasChange("wanopt_detection") {
		t, err := expandPackagesFirewallPolicyBlockWanoptDetection(d, v, "wanopt_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-detection"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_passive_opt"); ok || d.HasChange("wanopt_passive_opt") {
		t, err := expandPackagesFirewallPolicyBlockWanoptPassiveOpt(d, v, "wanopt_passive_opt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-passive-opt"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_peer"); ok || d.HasChange("wanopt_peer") {
		t, err := expandPackagesFirewallPolicyBlockWanoptPeer(d, v, "wanopt_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-peer"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_profile"); ok || d.HasChange("wanopt_profile") {
		t, err := expandPackagesFirewallPolicyBlockWanoptProfile(d, v, "wanopt_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-profile"] = t
		}
	}

	if v, ok := d.GetOk("wccp"); ok || d.HasChange("wccp") {
		t, err := expandPackagesFirewallPolicyBlockWccp(d, v, "wccp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wccp"] = t
		}
	}

	if v, ok := d.GetOk("webcache"); ok || d.HasChange("webcache") {
		t, err := expandPackagesFirewallPolicyBlockWebcache(d, v, "webcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache"] = t
		}
	}

	if v, ok := d.GetOk("webcache_https"); ok || d.HasChange("webcache_https") {
		t, err := expandPackagesFirewallPolicyBlockWebcacheHttps(d, v, "webcache_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache-https"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandPackagesFirewallPolicyBlockWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_forward_server"); ok || d.HasChange("webproxy_forward_server") {
		t, err := expandPackagesFirewallPolicyBlockWebproxyForwardServer(d, v, "webproxy_forward_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-forward-server"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok || d.HasChange("webproxy_profile") {
		t, err := expandPackagesFirewallPolicyBlockWebproxyProfile(d, v, "webproxy_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-profile"] = t
		}
	}

	if v, ok := d.GetOk("ztna_ems_tag"); ok || d.HasChange("ztna_ems_tag") {
		t, err := expandPackagesFirewallPolicyBlockZtnaEmsTag(d, v, "ztna_ems_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-ems-tag"] = t
		}
	}

	if v, ok := d.GetOk("ztna_geo_tag"); ok || d.HasChange("ztna_geo_tag") {
		t, err := expandPackagesFirewallPolicyBlockZtnaGeoTag(d, v, "ztna_geo_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-geo-tag"] = t
		}
	}

	if v, ok := d.GetOk("ztna_status"); ok || d.HasChange("ztna_status") {
		t, err := expandPackagesFirewallPolicyBlockZtnaStatus(d, v, "ztna_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-status"] = t
		}
	}

	if v, ok := d.GetOk("wsso"); ok || d.HasChange("wsso") {
		t, err := expandPackagesFirewallPolicyBlockWsso(d, v, "wsso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wsso"] = t
		}
	}

	return &obj, nil
}
