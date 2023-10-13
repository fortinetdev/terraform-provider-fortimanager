// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure interfaces.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanInterfaceUpdate,
		Read:   resourceObjectFspVlanInterfaceRead,
		Update: resourceObjectFspVlanInterfaceUpdate,
		Delete: resourceObjectFspVlanInterfaceDelete,

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
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vlan_op_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ac_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aggregate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aggregate_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ap_discover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arpforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"atm_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_portal_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_auth_extension_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_measure_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd_desired_min_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bfd_detect_mult": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bfd_required_min_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"broadcast_forticlient_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"broadcast_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cli_conn_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ddns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_keyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ddns_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_sn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ddns_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dedicated_to": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_purdue_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultgw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"detected_peer_mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"detectprotocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"detectserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_access_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_identification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_identification_active_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_netscan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_user_identification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"devindex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp_broadcast_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_classless_route_addition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_client_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_relay_agent_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_relay_circuit_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_relay_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_relay_interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_link_selection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_request_all_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_renew_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp_smart_relay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disc_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disconnect_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dns_query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server_protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"drop_fragment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drop_overlapped_fragment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_ca_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"eap_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eap_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eap_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"eap_supplicant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_user_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"egress_cos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"egress_shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"endpoint_compliance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"estimated_downstream_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"estimated_upstream_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"explicit_ftp_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"explicit_web_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_action_on_extender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_alert_interfaces": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fail_alert_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fail_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_detect_option": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fdp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortiheartbeat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortilink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink_backup_link": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fortilink_neighbor_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink_split_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortilink_stacking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_domain": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"forward_error_correction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fp_anomaly": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fp_disable": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gateway_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"generic_receive_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gi_gk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gwaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gwdetect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"icmp_accept_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"icmp_send_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ident_accept": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"if_mdix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"if_media": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ike_saml_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"in_force_vlan_cos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ingress_cos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ingress_shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ingress_spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interconnect_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internal": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_managed_by_fortiipam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_sniffer_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipunnumbered": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"autoconf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cli_conn6_status": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dhcp6_client_options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dhcp6_information_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_prefix_delegation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_prefix_hint": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_prefix_hint_plt": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dhcp6_prefix_hint_vlt": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dhcp6_relay_interface_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp6_relay_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp6_relay_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_relay_source_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_relay_source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_relay_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmp6_send_redirect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface_identifier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_allowaccess": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip6_default_life": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip6_delegated_prefix_iaid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip6_delegated_prefix_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"autonomous_flag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"delegated_prefix_iaid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"onlink_flag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"prefix_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"rdnss": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"rdnss_service": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"subnet": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"upstream_interface": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"ip6_dns_server_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_extra_addr": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"ip6_hop_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip6_link_mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip6_manage_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_max_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip6_min_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip6_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_other_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_prefix_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"autonomous_flag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dnssl": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"onlink_flag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"preferred_life_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"prefix": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"rdnss": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"valid_life_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"ip6_prefix_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_reachable_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip6_retrans_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip6_send_adv": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_upstream_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nd_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nd_cga_modifier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nd_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nd_security_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"nd_timestamp_delta": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"nd_timestamp_fuzz": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ra_send_mtu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unique_autoconf_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrip6_link_local": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrrp_virtual_mac6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrrp6": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"accept_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"adv_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"preempt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"priority": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vrdst6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vrgrp": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"vrid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"vrip6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"l2forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2tp_client": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lacp_ha_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lacp_ha_slave": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lacp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lacp_speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"large_receive_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lcp_echo_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lcp_max_echo_fails": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"link_up_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"listen_forticlient_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lldp_network_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lldp_reception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lldp_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"macaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_subnetwork_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_egress_burst_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_egress_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"measured_downstream_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"measured_upstream_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mediatype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"min_links": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_links_down": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mtu_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mux_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ndiscforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netbios_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netflow_sampler": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"np_qos_profile": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"npu_fastpath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"out_force_vlan_cos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"outbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"padt_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"peer_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"phy_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ping_serv_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"poe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polling_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pppoe_unnumbered_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pptp_auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pptp_client": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pptp_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"pptp_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pptp_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pptp_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"preserve_session_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_captive_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pvc_atm_qos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pvc_chan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pvc_crc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pvc_pcr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pvc_scr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pvc_vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pvc_vlan_rx_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pvc_vlan_rx_op": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pvc_vlan_tx_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pvc_vlan_tx_op": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reachable_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"redundant_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacemsg_override_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"retransmission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ring_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ring_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sample_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sample_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondaryip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowaccess": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"detectprotocol": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"detectserver": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"gwdetect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ha_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"secip_relay_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ping_serv_status": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"seq": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"security_8021x_dynamic_vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"security_8021x_master": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_8021x_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_exempt_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_external_logout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_external_web": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"select_profile_30a_35b": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sflow_sampler": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sfp_dsl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sfp_dsl_adsl_fallback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sfp_dsl_autodetect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sfp_dsl_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"src_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stp_ha_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stp_ha_slave": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stpforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stpforward_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_priority_vlan_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"substitute_dst_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sw_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"swc_first_create": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"swc_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_controller_access_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_arp_inspection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_controller_dhcp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_dhcp_snooping_option82": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_dhcp_snooping_verify_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_dynamic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_controller_feature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_igmp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_igmp_snooping_fast_leave": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_igmp_snooping_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_iot_scanning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_learning_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"switch_controller_mgmt_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"switch_controller_nac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_controller_netflow_collect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_offload_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_offload_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_offloading": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_controller_offloading_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_controller_offloading_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_controller_radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_controller_rspan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_traffic_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tc_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_mss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"trunk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip_1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trust_ip_2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trust_ip_3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trust_ip6_1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip6_2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip6_3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vci": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vectoring": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vindex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlanforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vpi": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrrp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"adv_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ignore_default_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preempt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"proxy_arp": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrdst": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vrdst_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vrgrp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vrid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vrip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"vrrp_virtual_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wifi_5g_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_acl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_ap_band": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_auto_connect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_auto_save": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_broadcast_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_encrypt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_fragment_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wifi_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wifi_keyindex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wifi_mac_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_passphrase": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"wifi_radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_rts_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wifi_security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_usergroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wins_ip": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectFspVlanInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFspVlanInterface(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFspVlanInterface")

	return resourceObjectFspVlanInterfaceRead(d, m)
}

func resourceObjectFspVlanInterfaceDelete(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	err = c.DeleteObjectFspVlanInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	paradict["vlan"] = vlan

	o, err := c.ReadObjectFspVlanInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterface resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanInterfaceVlanOpMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAcName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAggregate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAggregateType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAlgorithm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAlias2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAllowaccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceApDiscover2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceArpforward2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAtmProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAuthCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceAuthPortalAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAuthType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAutoAuthExtensionDevice2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBandwidthMeasureTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBfd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBfdDesiredMinTx2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBfdDetectMult2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBfdRequiredMinRx2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBroadcastForticlientDiscovery2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBroadcastForward2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceCaptivePortal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceCliConnStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceColor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdns2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsKeyname2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsPassword2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceDdnsServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsSn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsUsername2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsZone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDedicatedTo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDefaultPurdueLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDefaultgw2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDetectedPeerMtu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDetectprotocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceDetectserver2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDeviceAccessList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDeviceIdentification2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDeviceIdentificationActiveScan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDeviceNetscan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDeviceUserIdentification2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDevindex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpBroadcastFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpClasslessRouteAddition2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpClientIdentifier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayAgentOption2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayCircuitId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceDhcpRelayInterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceDhcpRelayLinkSelection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayRequestAllServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelaySourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRenewTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpSmartRelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDiscRetryTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDisconnectThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDistance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDnsQuery2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDnsServerOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDnsServerProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceDropFragment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDropOverlappedFragment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEapCaCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceEapIdentity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEapMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEapPassword2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceEapSupplicant2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEapUserCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceEgressCos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEgressShapingProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceEip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEndpointCompliance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEstimatedDownstreamBandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEstimatedUpstreamBandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceExplicitFtpProxy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceExplicitWebProxy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceExternal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFailActionOnExtender2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFailAlertInterfaces2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFailAlertMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFailDetect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFailDetectOption2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceFdp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortiheartbeat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortilink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortilinkBackupLink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortilinkNeighborDetect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortilinkSplitInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortilinkStacking2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceForwardDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceForwardErrorCorrection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFpAnomaly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceFpDisable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceGatewayAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceGenericReceiveOffload2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceGiGk2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceGwaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceGwdetect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceHaPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIcmpAcceptRedirect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIcmpRedirect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIcmpSendRedirect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIdentAccept2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIdleTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIfMdix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIfMedia2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIkeSamlServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceInForceVlanCos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceInbandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIngressCos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIngressShapingProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceIngressSpilloverThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceInterconnectProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceInternal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceIpManagedByFortiipam2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpmac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpsSnifferMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpunnumbered2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv62edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "autoconf"
	if _, ok := i["autoconf"]; ok {
		result["autoconf"] = flattenObjectFspVlanInterfaceIpv6Autoconf2edl(i["autoconf"], d, pre_append)
	}

	pre_append = pre + ".0." + "cli_conn6_status"
	if _, ok := i["cli-conn6-status"]; ok {
		result["cli_conn6_status"] = flattenObjectFspVlanInterfaceIpv6CliConn6Status2edl(i["cli-conn6-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := i["dhcp6-client-options"]; ok {
		result["dhcp6_client_options"] = flattenObjectFspVlanInterfaceIpv6Dhcp6ClientOptions2edl(i["dhcp6-client-options"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := i["dhcp6-information-request"]; ok {
		result["dhcp6_information_request"] = flattenObjectFspVlanInterfaceIpv6Dhcp6InformationRequest2edl(i["dhcp6-information-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := i["dhcp6-prefix-delegation"]; ok {
		result["dhcp6_prefix_delegation"] = flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation2edl(i["dhcp6-prefix-delegation"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := i["dhcp6-prefix-hint"]; ok {
		result["dhcp6_prefix_hint"] = flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHint2edl(i["dhcp6-prefix-hint"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := i["dhcp6-prefix-hint-plt"]; ok {
		result["dhcp6_prefix_hint_plt"] = flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt2edl(i["dhcp6-prefix-hint-plt"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := i["dhcp6-prefix-hint-vlt"]; ok {
		result["dhcp6_prefix_hint_vlt"] = flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt2edl(i["dhcp6-prefix-hint-vlt"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_interface_id"
	if _, ok := i["dhcp6-relay-interface-id"]; ok {
		result["dhcp6_relay_interface_id"] = flattenObjectFspVlanInterfaceIpv6Dhcp6RelayInterfaceId2edl(i["dhcp6-relay-interface-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := i["dhcp6-relay-ip"]; ok {
		result["dhcp6_relay_ip"] = flattenObjectFspVlanInterfaceIpv6Dhcp6RelayIp2edl(i["dhcp6-relay-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := i["dhcp6-relay-service"]; ok {
		result["dhcp6_relay_service"] = flattenObjectFspVlanInterfaceIpv6Dhcp6RelayService2edl(i["dhcp6-relay-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_source_interface"
	if _, ok := i["dhcp6-relay-source-interface"]; ok {
		result["dhcp6_relay_source_interface"] = flattenObjectFspVlanInterfaceIpv6Dhcp6RelaySourceInterface2edl(i["dhcp6-relay-source-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_source_ip"
	if _, ok := i["dhcp6-relay-source-ip"]; ok {
		result["dhcp6_relay_source_ip"] = flattenObjectFspVlanInterfaceIpv6Dhcp6RelaySourceIp2edl(i["dhcp6-relay-source-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := i["dhcp6-relay-type"]; ok {
		result["dhcp6_relay_type"] = flattenObjectFspVlanInterfaceIpv6Dhcp6RelayType2edl(i["dhcp6-relay-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp6_send_redirect"
	if _, ok := i["icmp6-send-redirect"]; ok {
		result["icmp6_send_redirect"] = flattenObjectFspVlanInterfaceIpv6Icmp6SendRedirect2edl(i["icmp6-send-redirect"], d, pre_append)
	}

	pre_append = pre + ".0." + "interface_identifier"
	if _, ok := i["interface-identifier"]; ok {
		result["interface_identifier"] = flattenObjectFspVlanInterfaceIpv6InterfaceIdentifier2edl(i["interface-identifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_address"
	if _, ok := i["ip6-address"]; ok {
		result["ip6_address"] = flattenObjectFspVlanInterfaceIpv6Ip6Address2edl(i["ip6-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := i["ip6-allowaccess"]; ok {
		result["ip6_allowaccess"] = flattenObjectFspVlanInterfaceIpv6Ip6Allowaccess2edl(i["ip6-allowaccess"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := i["ip6-default-life"]; ok {
		result["ip6_default_life"] = flattenObjectFspVlanInterfaceIpv6Ip6DefaultLife2edl(i["ip6-default-life"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_iaid"
	if _, ok := i["ip6-delegated-prefix-iaid"]; ok {
		result["ip6_delegated_prefix_iaid"] = flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid2edl(i["ip6-delegated-prefix-iaid"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := i["ip6-delegated-prefix-list"]; ok {
		result["ip6_delegated_prefix_list"] = flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList2edl(i["ip6-delegated-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := i["ip6-dns-server-override"]; ok {
		result["ip6_dns_server_override"] = flattenObjectFspVlanInterfaceIpv6Ip6DnsServerOverride2edl(i["ip6-dns-server-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := i["ip6-extra-addr"]; ok {
		result["ip6_extra_addr"] = flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddr2edl(i["ip6-extra-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := i["ip6-hop-limit"]; ok {
		result["ip6_hop_limit"] = flattenObjectFspVlanInterfaceIpv6Ip6HopLimit2edl(i["ip6-hop-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := i["ip6-link-mtu"]; ok {
		result["ip6_link_mtu"] = flattenObjectFspVlanInterfaceIpv6Ip6LinkMtu2edl(i["ip6-link-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := i["ip6-manage-flag"]; ok {
		result["ip6_manage_flag"] = flattenObjectFspVlanInterfaceIpv6Ip6ManageFlag2edl(i["ip6-manage-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := i["ip6-max-interval"]; ok {
		result["ip6_max_interval"] = flattenObjectFspVlanInterfaceIpv6Ip6MaxInterval2edl(i["ip6-max-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := i["ip6-min-interval"]; ok {
		result["ip6_min_interval"] = flattenObjectFspVlanInterfaceIpv6Ip6MinInterval2edl(i["ip6-min-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := i["ip6-mode"]; ok {
		result["ip6_mode"] = flattenObjectFspVlanInterfaceIpv6Ip6Mode2edl(i["ip6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := i["ip6-other-flag"]; ok {
		result["ip6_other_flag"] = flattenObjectFspVlanInterfaceIpv6Ip6OtherFlag2edl(i["ip6-other-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := i["ip6-prefix-list"]; ok {
		result["ip6_prefix_list"] = flattenObjectFspVlanInterfaceIpv6Ip6PrefixList2edl(i["ip6-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_mode"
	if _, ok := i["ip6-prefix-mode"]; ok {
		result["ip6_prefix_mode"] = flattenObjectFspVlanInterfaceIpv6Ip6PrefixMode2edl(i["ip6-prefix-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := i["ip6-reachable-time"]; ok {
		result["ip6_reachable_time"] = flattenObjectFspVlanInterfaceIpv6Ip6ReachableTime2edl(i["ip6-reachable-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := i["ip6-retrans-time"]; ok {
		result["ip6_retrans_time"] = flattenObjectFspVlanInterfaceIpv6Ip6RetransTime2edl(i["ip6-retrans-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := i["ip6-send-adv"]; ok {
		result["ip6_send_adv"] = flattenObjectFspVlanInterfaceIpv6Ip6SendAdv2edl(i["ip6-send-adv"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := i["ip6-subnet"]; ok {
		result["ip6_subnet"] = flattenObjectFspVlanInterfaceIpv6Ip6Subnet2edl(i["ip6-subnet"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := i["ip6-upstream-interface"]; ok {
		result["ip6_upstream_interface"] = flattenObjectFspVlanInterfaceIpv6Ip6UpstreamInterface2edl(i["ip6-upstream-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cert"
	if _, ok := i["nd-cert"]; ok {
		result["nd_cert"] = flattenObjectFspVlanInterfaceIpv6NdCert2edl(i["nd-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := i["nd-cga-modifier"]; ok {
		result["nd_cga_modifier"] = flattenObjectFspVlanInterfaceIpv6NdCgaModifier2edl(i["nd-cga-modifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_mode"
	if _, ok := i["nd-mode"]; ok {
		result["nd_mode"] = flattenObjectFspVlanInterfaceIpv6NdMode2edl(i["nd-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := i["nd-security-level"]; ok {
		result["nd_security_level"] = flattenObjectFspVlanInterfaceIpv6NdSecurityLevel2edl(i["nd-security-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := i["nd-timestamp-delta"]; ok {
		result["nd_timestamp_delta"] = flattenObjectFspVlanInterfaceIpv6NdTimestampDelta2edl(i["nd-timestamp-delta"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := i["nd-timestamp-fuzz"]; ok {
		result["nd_timestamp_fuzz"] = flattenObjectFspVlanInterfaceIpv6NdTimestampFuzz2edl(i["nd-timestamp-fuzz"], d, pre_append)
	}

	pre_append = pre + ".0." + "ra_send_mtu"
	if _, ok := i["ra-send-mtu"]; ok {
		result["ra_send_mtu"] = flattenObjectFspVlanInterfaceIpv6RaSendMtu2edl(i["ra-send-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "unique_autoconf_addr"
	if _, ok := i["unique-autoconf-addr"]; ok {
		result["unique_autoconf_addr"] = flattenObjectFspVlanInterfaceIpv6UniqueAutoconfAddr2edl(i["unique-autoconf-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := i["vrip6_link_local"]; ok {
		result["vrip6_link_local"] = flattenObjectFspVlanInterfaceIpv6Vrip6LinkLocal2edl(i["vrip6_link_local"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := i["vrrp-virtual-mac6"]; ok {
		result["vrrp_virtual_mac6"] = flattenObjectFspVlanInterfaceIpv6VrrpVirtualMac62edl(i["vrrp-virtual-mac6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp6"
	if _, ok := i["vrrp6"]; ok {
		result["vrrp6"] = flattenObjectFspVlanInterfaceIpv6Vrrp62edl(i["vrrp6"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFspVlanInterfaceIpv6Autoconf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6CliConn6Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6ClientOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6InformationRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHint2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayInterfaceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelaySourceInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelaySourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Icmp6SendRedirect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6InterfaceIdentifier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Allowaccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6DefaultLife2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := i["delegated-prefix-iaid"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(i["delegated-prefix-iaid"], d, pre_append)
			tmp["delegated_prefix_iaid"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-DelegatedPrefixIaid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := i["prefix-id"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(i["prefix-id"], d, pre_append)
			tmp["prefix_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-PrefixId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := i["rdnss-service"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(i["rdnss-service"], d, pre_append)
			tmp["rdnss_service"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-RdnssService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := i["upstream-interface"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(i["upstream-interface"], d, pre_append)
			tmp["upstream_interface"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-UpstreamInterface")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DnsServerOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddr2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix2edl(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6ExtraAddr-Prefix")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6HopLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6LinkMtu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6ManageFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6MaxInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6MinInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6OtherFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := i["dnssl"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl2edl(i["dnssl"], d, pre_append)
			tmp["dnssl"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-Dnssl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := i["preferred-life-time"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(i["preferred-life-time"], d, pre_append)
			tmp["preferred_life_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-PreferredLifeTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix2edl(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss2edl(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := i["valid-life-time"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime2edl(i["valid-life-time"], d, pre_append)
			tmp["valid_life_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-ValidLifeTime")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6ReachableTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6RetransTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6SendAdv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6UpstreamInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceIpv6NdCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdCgaModifier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdSecurityLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdTimestampDelta2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdTimestampFuzz2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6RaSendMtu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6UniqueAutoconfAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrip6LinkLocal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6VrrpVirtualMac62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp62edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := i["accept-mode"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6AcceptMode2edl(i["accept-mode"], d, pre_append)
			tmp["accept_mode"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-AcceptMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6AdvInterval2edl(i["adv-interval"], d, pre_append)
			tmp["adv_interval"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-AdvInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Preempt2edl(i["preempt"], d, pre_append)
			tmp["preempt"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Preempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Priority2edl(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6StartTime2edl(i["start-time"], d, pre_append)
			tmp["start_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-StartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Status2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := i["vrdst6"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrdst62edl(i["vrdst6"], d, pre_append)
			tmp["vrdst6"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrdst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrgrp2edl(i["vrgrp"], d, pre_append)
			tmp["vrgrp"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrgrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrid2edl(i["vrid"], d, pre_append)
			tmp["vrid"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := i["vrip6"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrip62edl(i["vrip6"], d, pre_append)
			tmp["vrip6"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrip6")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6AcceptMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6AdvInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Preempt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Priority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6StartTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrdst62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrgrp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrip62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceL2Forward2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceL2TpClient2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLacpHaSecondary2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLacpHaSlave2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLacpMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLacpSpeed2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLargeReceiveOffload2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLcpEchoInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLcpMaxEchoFails2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLinkUpDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceListenForticlientConnection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLldpNetworkPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceLldpReception2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLldpTransmission2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMacaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceManagedSubnetworkSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceManagementIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceMaxEgressBurstRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMaxEgressRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMeasuredDownstreamBandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMeasuredUpstreamBandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMediatype2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMinLinks2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMinLinksDown2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMonitorBandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMtu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMtuOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMuxType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNdiscforward2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNetbiosForward2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNetflowSampler2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNpQosProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNpuFastpath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceOutForceVlanCos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceOutbandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePadtRetryTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePassword2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfacePeerInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePhyMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePingServStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePoe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePollingInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePppoeUnnumberedNegotiate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePptpAuthType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePptpClient2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePptpPassword2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfacePptpServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePptpTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePptpUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePreserveSessionRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePriorityOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceProxyCaptivePortal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcAtmQos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcChan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcCrc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcPcr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcScr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcVlanId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcVlanRxId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcVlanRxOp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcVlanTxId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcVlanTxOp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceReachableTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRedundantInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRemoteIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceReplacemsgOverrideGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRetransmission2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRingRx2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRingTx2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSampleDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSampleRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceScanBotnetConnections2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryip2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowaccess"
		if _, ok := i["allowaccess"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipAllowaccess2edl(i["allowaccess"], d, pre_append)
			tmp["allowaccess"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Allowaccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := i["detectprotocol"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipDetectprotocol2edl(i["detectprotocol"], d, pre_append)
			tmp["detectprotocol"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Detectprotocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := i["detectserver"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipDetectserver2edl(i["detectserver"], d, pre_append)
			tmp["detectserver"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Detectserver")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := i["gwdetect"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipGwdetect2edl(i["gwdetect"], d, pre_append)
			tmp["gwdetect"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Gwdetect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipHaPriority2edl(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secip_relay_ip"
		if _, ok := i["secip-relay-ip"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipSecipRelayIp2edl(i["secip-relay-ip"], d, pre_append)
			tmp["secip_relay_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-SecipRelayIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := i["ping-serv-status"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipPingServStatus2edl(i["ping-serv-status"], d, pre_append)
			tmp["ping_serv_status"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-PingServStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipSeq2edl(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Seq")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceSecondaryipAllowaccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceSecondaryipDetectprotocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceSecondaryipDetectserver2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipGwdetect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipHaPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipSecipRelayIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipPingServStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipSeq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurity8021XDynamicVlanId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurity8021XMaster2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurity8021XMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityExemptList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceSecurityExternalLogout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityExternalWeb2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityGroups2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityMacAuthBypass2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityRedirectUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSelectProfile30A35B2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceServiceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSflowSampler2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSfpDsl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSfpDslAdslFallback2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSfpDslAutodetect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSfpDslMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSpeed2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSpilloverThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSrcCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStpHaSecondary2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStpHaSlave2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStpforward2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStpforwardMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStripPriorityVlanTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSubst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSubstituteDstMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwAlgorithm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwcFirstCreate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwcVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerAccessVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerArpInspection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerDhcpSnooping2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerDhcpSnoopingOption822edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerDhcpSnoopingVerifyMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerDynamic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceSwitchControllerFeature2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerIgmpSnooping2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerIgmpSnoopingFastLeave2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerIgmpSnoopingProxy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerIotScanning2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerLearningLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerMgmtVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerNac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceSwitchControllerNetflowCollect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerOffload2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerOffloadGw2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerOffloadIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerOffloading2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerOffloadingGw2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerOffloadingIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerRadiusServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerRspanMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerTrafficPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceSystemId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSystemIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTcMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTcpMss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrunk2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp612edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp622edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp632edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceUsername2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVci2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVectoring2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVindex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVlanProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVlanforward2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVlanid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVpi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrp2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := i["accept-mode"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpAcceptMode2edl(i["accept-mode"], d, pre_append)
			tmp["accept_mode"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-AcceptMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpAdvInterval2edl(i["adv-interval"], d, pre_append)
			tmp["adv_interval"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-AdvInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := i["ignore-default-route"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpIgnoreDefaultRoute2edl(i["ignore-default-route"], d, pre_append)
			tmp["ignore_default_route"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-IgnoreDefaultRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpPreempt2edl(i["preempt"], d, pre_append)
			tmp["preempt"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Preempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpPriority2edl(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy_arp"
		if _, ok := i["proxy-arp"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpProxyArp2edl(i["proxy-arp"], d, pre_append)
			tmp["proxy_arp"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-ProxyArp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpStartTime2edl(i["start-time"], d, pre_append)
			tmp["start_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-StartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpStatus2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVersion2edl(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Version")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst"
		if _, ok := i["vrdst"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVrdst2edl(i["vrdst"], d, pre_append)
			tmp["vrdst"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Vrdst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := i["vrdst-priority"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVrdstPriority2edl(i["vrdst-priority"], d, pre_append)
			tmp["vrdst_priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-VrdstPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVrgrp2edl(i["vrgrp"], d, pre_append)
			tmp["vrgrp"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Vrgrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVrid2edl(i["vrid"], d, pre_append)
			tmp["vrid"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Vrid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip"
		if _, ok := i["vrip"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVrip2edl(i["vrip"], d, pre_append)
			tmp["vrip"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Vrip")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceVrrpAcceptMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpAdvInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpIgnoreDefaultRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpPreempt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpProxyArp2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanInterfaceVrrpProxyArpId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceVrrp-ProxyArp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpProxyArpIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceVrrp-ProxyArp-Ip")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceVrrpProxyArpId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpProxyArpIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpStartTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrdst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceVrrpVrdstPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrgrp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVirtualMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWccp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifi5GThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiAcl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiApBand2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiAutoConnect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiAutoSave2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiBroadcastSsid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiDnsServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiDnsServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiEncrypt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiFragmentThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceWifiKeyindex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiMacFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiPassphrase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceWifiRadiusServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiRtsThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiSecurity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiSsid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiUsergroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWinsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("vlan_op_mode", flattenObjectFspVlanInterfaceVlanOpMode2edl(o["vlan-op-mode"], d, "vlan_op_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-op-mode"], "ObjectFspVlanInterface-VlanOpMode"); ok {
			if err = d.Set("vlan_op_mode", vv); err != nil {
				return fmt.Errorf("Error reading vlan_op_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_op_mode: %v", err)
		}
	}

	if err = d.Set("ac_name", flattenObjectFspVlanInterfaceAcName2edl(o["ac-name"], d, "ac_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ac-name"], "ObjectFspVlanInterface-AcName"); ok {
			if err = d.Set("ac_name", vv); err != nil {
				return fmt.Errorf("Error reading ac_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ac_name: %v", err)
		}
	}

	if err = d.Set("aggregate", flattenObjectFspVlanInterfaceAggregate2edl(o["aggregate"], d, "aggregate")); err != nil {
		if vv, ok := fortiAPIPatch(o["aggregate"], "ObjectFspVlanInterface-Aggregate"); ok {
			if err = d.Set("aggregate", vv); err != nil {
				return fmt.Errorf("Error reading aggregate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aggregate: %v", err)
		}
	}

	if err = d.Set("aggregate_type", flattenObjectFspVlanInterfaceAggregateType2edl(o["aggregate-type"], d, "aggregate_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["aggregate-type"], "ObjectFspVlanInterface-AggregateType"); ok {
			if err = d.Set("aggregate_type", vv); err != nil {
				return fmt.Errorf("Error reading aggregate_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aggregate_type: %v", err)
		}
	}

	if err = d.Set("algorithm", flattenObjectFspVlanInterfaceAlgorithm2edl(o["algorithm"], d, "algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["algorithm"], "ObjectFspVlanInterface-Algorithm"); ok {
			if err = d.Set("algorithm", vv); err != nil {
				return fmt.Errorf("Error reading algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading algorithm: %v", err)
		}
	}

	if err = d.Set("alias", flattenObjectFspVlanInterfaceAlias2edl(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "ObjectFspVlanInterface-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenObjectFspVlanInterfaceAllowaccess2edl(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "ObjectFspVlanInterface-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("ap_discover", flattenObjectFspVlanInterfaceApDiscover2edl(o["ap-discover"], d, "ap_discover")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-discover"], "ObjectFspVlanInterface-ApDiscover"); ok {
			if err = d.Set("ap_discover", vv); err != nil {
				return fmt.Errorf("Error reading ap_discover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_discover: %v", err)
		}
	}

	if err = d.Set("arpforward", flattenObjectFspVlanInterfaceArpforward2edl(o["arpforward"], d, "arpforward")); err != nil {
		if vv, ok := fortiAPIPatch(o["arpforward"], "ObjectFspVlanInterface-Arpforward"); ok {
			if err = d.Set("arpforward", vv); err != nil {
				return fmt.Errorf("Error reading arpforward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arpforward: %v", err)
		}
	}

	if err = d.Set("atm_protocol", flattenObjectFspVlanInterfaceAtmProtocol2edl(o["atm-protocol"], d, "atm_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["atm-protocol"], "ObjectFspVlanInterface-AtmProtocol"); ok {
			if err = d.Set("atm_protocol", vv); err != nil {
				return fmt.Errorf("Error reading atm_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading atm_protocol: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenObjectFspVlanInterfaceAuthCert2edl(o["auth-cert"], d, "auth_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-cert"], "ObjectFspVlanInterface-AuthCert"); ok {
			if err = d.Set("auth_cert", vv); err != nil {
				return fmt.Errorf("Error reading auth_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_portal_addr", flattenObjectFspVlanInterfaceAuthPortalAddr2edl(o["auth-portal-addr"], d, "auth_portal_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal-addr"], "ObjectFspVlanInterface-AuthPortalAddr"); ok {
			if err = d.Set("auth_portal_addr", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal_addr: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenObjectFspVlanInterfaceAuthType2edl(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "ObjectFspVlanInterface-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("auto_auth_extension_device", flattenObjectFspVlanInterfaceAutoAuthExtensionDevice2edl(o["auto-auth-extension-device"], d, "auto_auth_extension_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-auth-extension-device"], "ObjectFspVlanInterface-AutoAuthExtensionDevice"); ok {
			if err = d.Set("auto_auth_extension_device", vv); err != nil {
				return fmt.Errorf("Error reading auto_auth_extension_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_auth_extension_device: %v", err)
		}
	}

	if err = d.Set("bandwidth_measure_time", flattenObjectFspVlanInterfaceBandwidthMeasureTime2edl(o["bandwidth-measure-time"], d, "bandwidth_measure_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-measure-time"], "ObjectFspVlanInterface-BandwidthMeasureTime"); ok {
			if err = d.Set("bandwidth_measure_time", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_measure_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_measure_time: %v", err)
		}
	}

	if err = d.Set("bfd", flattenObjectFspVlanInterfaceBfd2edl(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "ObjectFspVlanInterface-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("bfd_desired_min_tx", flattenObjectFspVlanInterfaceBfdDesiredMinTx2edl(o["bfd-desired-min-tx"], d, "bfd_desired_min_tx")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-desired-min-tx"], "ObjectFspVlanInterface-BfdDesiredMinTx"); ok {
			if err = d.Set("bfd_desired_min_tx", vv); err != nil {
				return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
		}
	}

	if err = d.Set("bfd_detect_mult", flattenObjectFspVlanInterfaceBfdDetectMult2edl(o["bfd-detect-mult"], d, "bfd_detect_mult")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-detect-mult"], "ObjectFspVlanInterface-BfdDetectMult"); ok {
			if err = d.Set("bfd_detect_mult", vv); err != nil {
				return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
		}
	}

	if err = d.Set("bfd_required_min_rx", flattenObjectFspVlanInterfaceBfdRequiredMinRx2edl(o["bfd-required-min-rx"], d, "bfd_required_min_rx")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-required-min-rx"], "ObjectFspVlanInterface-BfdRequiredMinRx"); ok {
			if err = d.Set("bfd_required_min_rx", vv); err != nil {
				return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
		}
	}

	if err = d.Set("broadcast_forticlient_discovery", flattenObjectFspVlanInterfaceBroadcastForticlientDiscovery2edl(o["broadcast-forticlient-discovery"], d, "broadcast_forticlient_discovery")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast-forticlient-discovery"], "ObjectFspVlanInterface-BroadcastForticlientDiscovery"); ok {
			if err = d.Set("broadcast_forticlient_discovery", vv); err != nil {
				return fmt.Errorf("Error reading broadcast_forticlient_discovery: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast_forticlient_discovery: %v", err)
		}
	}

	if err = d.Set("broadcast_forward", flattenObjectFspVlanInterfaceBroadcastForward2edl(o["broadcast-forward"], d, "broadcast_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast-forward"], "ObjectFspVlanInterface-BroadcastForward"); ok {
			if err = d.Set("broadcast_forward", vv); err != nil {
				return fmt.Errorf("Error reading broadcast_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast_forward: %v", err)
		}
	}

	if err = d.Set("captive_portal", flattenObjectFspVlanInterfaceCaptivePortal2edl(o["captive-portal"], d, "captive_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal"], "ObjectFspVlanInterface-CaptivePortal"); ok {
			if err = d.Set("captive_portal", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal: %v", err)
		}
	}

	if err = d.Set("cli_conn_status", flattenObjectFspVlanInterfaceCliConnStatus2edl(o["cli-conn-status"], d, "cli_conn_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-conn-status"], "ObjectFspVlanInterface-CliConnStatus"); ok {
			if err = d.Set("cli_conn_status", vv); err != nil {
				return fmt.Errorf("Error reading cli_conn_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_conn_status: %v", err)
		}
	}

	if err = d.Set("color", flattenObjectFspVlanInterfaceColor2edl(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFspVlanInterface-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("ddns", flattenObjectFspVlanInterfaceDdns2edl(o["ddns"], d, "ddns")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns"], "ObjectFspVlanInterface-Ddns"); ok {
			if err = d.Set("ddns", vv); err != nil {
				return fmt.Errorf("Error reading ddns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns: %v", err)
		}
	}

	if err = d.Set("ddns_auth", flattenObjectFspVlanInterfaceDdnsAuth2edl(o["ddns-auth"], d, "ddns_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-auth"], "ObjectFspVlanInterface-DdnsAuth"); ok {
			if err = d.Set("ddns_auth", vv); err != nil {
				return fmt.Errorf("Error reading ddns_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_domain", flattenObjectFspVlanInterfaceDdnsDomain2edl(o["ddns-domain"], d, "ddns_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-domain"], "ObjectFspVlanInterface-DdnsDomain"); ok {
			if err = d.Set("ddns_domain", vv); err != nil {
				return fmt.Errorf("Error reading ddns_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_domain: %v", err)
		}
	}

	if err = d.Set("ddns_key", flattenObjectFspVlanInterfaceDdnsKey2edl(o["ddns-key"], d, "ddns_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-key"], "ObjectFspVlanInterface-DdnsKey"); ok {
			if err = d.Set("ddns_key", vv); err != nil {
				return fmt.Errorf("Error reading ddns_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_key: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", flattenObjectFspVlanInterfaceDdnsKeyname2edl(o["ddns-keyname"], d, "ddns_keyname")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-keyname"], "ObjectFspVlanInterface-DdnsKeyname"); ok {
			if err = d.Set("ddns_keyname", vv); err != nil {
				return fmt.Errorf("Error reading ddns_keyname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_password", flattenObjectFspVlanInterfaceDdnsPassword2edl(o["ddns-password"], d, "ddns_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-password"], "ObjectFspVlanInterface-DdnsPassword"); ok {
			if err = d.Set("ddns_password", vv); err != nil {
				return fmt.Errorf("Error reading ddns_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_password: %v", err)
		}
	}

	if err = d.Set("ddns_server", flattenObjectFspVlanInterfaceDdnsServer2edl(o["ddns-server"], d, "ddns_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server"], "ObjectFspVlanInterface-DdnsServer"); ok {
			if err = d.Set("ddns_server", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenObjectFspVlanInterfaceDdnsServerIp2edl(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-ip"], "ObjectFspVlanInterface-DdnsServerIp"); ok {
			if err = d.Set("ddns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_sn", flattenObjectFspVlanInterfaceDdnsSn2edl(o["ddns-sn"], d, "ddns_sn")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-sn"], "ObjectFspVlanInterface-DdnsSn"); ok {
			if err = d.Set("ddns_sn", vv); err != nil {
				return fmt.Errorf("Error reading ddns_sn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_sn: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", flattenObjectFspVlanInterfaceDdnsTtl2edl(o["ddns-ttl"], d, "ddns_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-ttl"], "ObjectFspVlanInterface-DdnsTtl"); ok {
			if err = d.Set("ddns_ttl", vv); err != nil {
				return fmt.Errorf("Error reading ddns_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("ddns_username", flattenObjectFspVlanInterfaceDdnsUsername2edl(o["ddns-username"], d, "ddns_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-username"], "ObjectFspVlanInterface-DdnsUsername"); ok {
			if err = d.Set("ddns_username", vv); err != nil {
				return fmt.Errorf("Error reading ddns_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_username: %v", err)
		}
	}

	if err = d.Set("ddns_zone", flattenObjectFspVlanInterfaceDdnsZone2edl(o["ddns-zone"], d, "ddns_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-zone"], "ObjectFspVlanInterface-DdnsZone"); ok {
			if err = d.Set("ddns_zone", vv); err != nil {
				return fmt.Errorf("Error reading ddns_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("dedicated_to", flattenObjectFspVlanInterfaceDedicatedTo2edl(o["dedicated-to"], d, "dedicated_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["dedicated-to"], "ObjectFspVlanInterface-DedicatedTo"); ok {
			if err = d.Set("dedicated_to", vv); err != nil {
				return fmt.Errorf("Error reading dedicated_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dedicated_to: %v", err)
		}
	}

	if err = d.Set("default_purdue_level", flattenObjectFspVlanInterfaceDefaultPurdueLevel2edl(o["default-purdue-level"], d, "default_purdue_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-purdue-level"], "ObjectFspVlanInterface-DefaultPurdueLevel"); ok {
			if err = d.Set("default_purdue_level", vv); err != nil {
				return fmt.Errorf("Error reading default_purdue_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_purdue_level: %v", err)
		}
	}

	if err = d.Set("defaultgw", flattenObjectFspVlanInterfaceDefaultgw2edl(o["defaultgw"], d, "defaultgw")); err != nil {
		if vv, ok := fortiAPIPatch(o["defaultgw"], "ObjectFspVlanInterface-Defaultgw"); ok {
			if err = d.Set("defaultgw", vv); err != nil {
				return fmt.Errorf("Error reading defaultgw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading defaultgw: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectFspVlanInterfaceDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectFspVlanInterface-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("detected_peer_mtu", flattenObjectFspVlanInterfaceDetectedPeerMtu2edl(o["detected-peer-mtu"], d, "detected_peer_mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["detected-peer-mtu"], "ObjectFspVlanInterface-DetectedPeerMtu"); ok {
			if err = d.Set("detected_peer_mtu", vv); err != nil {
				return fmt.Errorf("Error reading detected_peer_mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detected_peer_mtu: %v", err)
		}
	}

	if err = d.Set("detectprotocol", flattenObjectFspVlanInterfaceDetectprotocol2edl(o["detectprotocol"], d, "detectprotocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["detectprotocol"], "ObjectFspVlanInterface-Detectprotocol"); ok {
			if err = d.Set("detectprotocol", vv); err != nil {
				return fmt.Errorf("Error reading detectprotocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detectprotocol: %v", err)
		}
	}

	if err = d.Set("detectserver", flattenObjectFspVlanInterfaceDetectserver2edl(o["detectserver"], d, "detectserver")); err != nil {
		if vv, ok := fortiAPIPatch(o["detectserver"], "ObjectFspVlanInterface-Detectserver"); ok {
			if err = d.Set("detectserver", vv); err != nil {
				return fmt.Errorf("Error reading detectserver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detectserver: %v", err)
		}
	}

	if err = d.Set("device_access_list", flattenObjectFspVlanInterfaceDeviceAccessList2edl(o["device-access-list"], d, "device_access_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-access-list"], "ObjectFspVlanInterface-DeviceAccessList"); ok {
			if err = d.Set("device_access_list", vv); err != nil {
				return fmt.Errorf("Error reading device_access_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_access_list: %v", err)
		}
	}

	if err = d.Set("device_identification", flattenObjectFspVlanInterfaceDeviceIdentification2edl(o["device-identification"], d, "device_identification")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-identification"], "ObjectFspVlanInterface-DeviceIdentification"); ok {
			if err = d.Set("device_identification", vv); err != nil {
				return fmt.Errorf("Error reading device_identification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_identification: %v", err)
		}
	}

	if err = d.Set("device_identification_active_scan", flattenObjectFspVlanInterfaceDeviceIdentificationActiveScan2edl(o["device-identification-active-scan"], d, "device_identification_active_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-identification-active-scan"], "ObjectFspVlanInterface-DeviceIdentificationActiveScan"); ok {
			if err = d.Set("device_identification_active_scan", vv); err != nil {
				return fmt.Errorf("Error reading device_identification_active_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_identification_active_scan: %v", err)
		}
	}

	if err = d.Set("device_netscan", flattenObjectFspVlanInterfaceDeviceNetscan2edl(o["device-netscan"], d, "device_netscan")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-netscan"], "ObjectFspVlanInterface-DeviceNetscan"); ok {
			if err = d.Set("device_netscan", vv); err != nil {
				return fmt.Errorf("Error reading device_netscan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_netscan: %v", err)
		}
	}

	if err = d.Set("device_user_identification", flattenObjectFspVlanInterfaceDeviceUserIdentification2edl(o["device-user-identification"], d, "device_user_identification")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-user-identification"], "ObjectFspVlanInterface-DeviceUserIdentification"); ok {
			if err = d.Set("device_user_identification", vv); err != nil {
				return fmt.Errorf("Error reading device_user_identification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_user_identification: %v", err)
		}
	}

	if err = d.Set("devindex", flattenObjectFspVlanInterfaceDevindex2edl(o["devindex"], d, "devindex")); err != nil {
		if vv, ok := fortiAPIPatch(o["devindex"], "ObjectFspVlanInterface-Devindex"); ok {
			if err = d.Set("devindex", vv); err != nil {
				return fmt.Errorf("Error reading devindex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devindex: %v", err)
		}
	}

	if err = d.Set("dhcp_broadcast_flag", flattenObjectFspVlanInterfaceDhcpBroadcastFlag2edl(o["dhcp-broadcast-flag"], d, "dhcp_broadcast_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-broadcast-flag"], "ObjectFspVlanInterface-DhcpBroadcastFlag"); ok {
			if err = d.Set("dhcp_broadcast_flag", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_broadcast_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_broadcast_flag: %v", err)
		}
	}

	if err = d.Set("dhcp_classless_route_addition", flattenObjectFspVlanInterfaceDhcpClasslessRouteAddition2edl(o["dhcp-classless-route-addition"], d, "dhcp_classless_route_addition")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-classless-route-addition"], "ObjectFspVlanInterface-DhcpClasslessRouteAddition"); ok {
			if err = d.Set("dhcp_classless_route_addition", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_classless_route_addition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_classless_route_addition: %v", err)
		}
	}

	if err = d.Set("dhcp_client_identifier", flattenObjectFspVlanInterfaceDhcpClientIdentifier2edl(o["dhcp-client-identifier"], d, "dhcp_client_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-client-identifier"], "ObjectFspVlanInterface-DhcpClientIdentifier"); ok {
			if err = d.Set("dhcp_client_identifier", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_client_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_client_identifier: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_agent_option", flattenObjectFspVlanInterfaceDhcpRelayAgentOption2edl(o["dhcp-relay-agent-option"], d, "dhcp_relay_agent_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-relay-agent-option"], "ObjectFspVlanInterface-DhcpRelayAgentOption"); ok {
			if err = d.Set("dhcp_relay_agent_option", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_relay_agent_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_relay_agent_option: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_circuit_id", flattenObjectFspVlanInterfaceDhcpRelayCircuitId2edl(o["dhcp-relay-circuit-id"], d, "dhcp_relay_circuit_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-relay-circuit-id"], "ObjectFspVlanInterface-DhcpRelayCircuitId"); ok {
			if err = d.Set("dhcp_relay_circuit_id", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_relay_circuit_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_relay_circuit_id: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_interface", flattenObjectFspVlanInterfaceDhcpRelayInterface2edl(o["dhcp-relay-interface"], d, "dhcp_relay_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-relay-interface"], "ObjectFspVlanInterface-DhcpRelayInterface"); ok {
			if err = d.Set("dhcp_relay_interface", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_relay_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_relay_interface: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_interface_select_method", flattenObjectFspVlanInterfaceDhcpRelayInterfaceSelectMethod2edl(o["dhcp-relay-interface-select-method"], d, "dhcp_relay_interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-relay-interface-select-method"], "ObjectFspVlanInterface-DhcpRelayInterfaceSelectMethod"); ok {
			if err = d.Set("dhcp_relay_interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_relay_interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_relay_interface_select_method: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_ip", flattenObjectFspVlanInterfaceDhcpRelayIp2edl(o["dhcp-relay-ip"], d, "dhcp_relay_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-relay-ip"], "ObjectFspVlanInterface-DhcpRelayIp"); ok {
			if err = d.Set("dhcp_relay_ip", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_relay_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_relay_ip: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_link_selection", flattenObjectFspVlanInterfaceDhcpRelayLinkSelection2edl(o["dhcp-relay-link-selection"], d, "dhcp_relay_link_selection")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-relay-link-selection"], "ObjectFspVlanInterface-DhcpRelayLinkSelection"); ok {
			if err = d.Set("dhcp_relay_link_selection", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_relay_link_selection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_relay_link_selection: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_request_all_server", flattenObjectFspVlanInterfaceDhcpRelayRequestAllServer2edl(o["dhcp-relay-request-all-server"], d, "dhcp_relay_request_all_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-relay-request-all-server"], "ObjectFspVlanInterface-DhcpRelayRequestAllServer"); ok {
			if err = d.Set("dhcp_relay_request_all_server", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_relay_request_all_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_relay_request_all_server: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_service", flattenObjectFspVlanInterfaceDhcpRelayService2edl(o["dhcp-relay-service"], d, "dhcp_relay_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-relay-service"], "ObjectFspVlanInterface-DhcpRelayService"); ok {
			if err = d.Set("dhcp_relay_service", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_relay_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_relay_service: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_source_ip", flattenObjectFspVlanInterfaceDhcpRelaySourceIp2edl(o["dhcp-relay-source-ip"], d, "dhcp_relay_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-relay-source-ip"], "ObjectFspVlanInterface-DhcpRelaySourceIp"); ok {
			if err = d.Set("dhcp_relay_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_relay_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_relay_source_ip: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_type", flattenObjectFspVlanInterfaceDhcpRelayType2edl(o["dhcp-relay-type"], d, "dhcp_relay_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-relay-type"], "ObjectFspVlanInterface-DhcpRelayType"); ok {
			if err = d.Set("dhcp_relay_type", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_relay_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_relay_type: %v", err)
		}
	}

	if err = d.Set("dhcp_renew_time", flattenObjectFspVlanInterfaceDhcpRenewTime2edl(o["dhcp-renew-time"], d, "dhcp_renew_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-renew-time"], "ObjectFspVlanInterface-DhcpRenewTime"); ok {
			if err = d.Set("dhcp_renew_time", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_renew_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_renew_time: %v", err)
		}
	}

	if err = d.Set("dhcp_smart_relay", flattenObjectFspVlanInterfaceDhcpSmartRelay2edl(o["dhcp-smart-relay"], d, "dhcp_smart_relay")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-smart-relay"], "ObjectFspVlanInterface-DhcpSmartRelay"); ok {
			if err = d.Set("dhcp_smart_relay", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_smart_relay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_smart_relay: %v", err)
		}
	}

	if err = d.Set("disc_retry_timeout", flattenObjectFspVlanInterfaceDiscRetryTimeout2edl(o["disc-retry-timeout"], d, "disc_retry_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["disc-retry-timeout"], "ObjectFspVlanInterface-DiscRetryTimeout"); ok {
			if err = d.Set("disc_retry_timeout", vv); err != nil {
				return fmt.Errorf("Error reading disc_retry_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disc_retry_timeout: %v", err)
		}
	}

	if err = d.Set("disconnect_threshold", flattenObjectFspVlanInterfaceDisconnectThreshold2edl(o["disconnect-threshold"], d, "disconnect_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["disconnect-threshold"], "ObjectFspVlanInterface-DisconnectThreshold"); ok {
			if err = d.Set("disconnect_threshold", vv); err != nil {
				return fmt.Errorf("Error reading disconnect_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disconnect_threshold: %v", err)
		}
	}

	if err = d.Set("distance", flattenObjectFspVlanInterfaceDistance2edl(o["distance"], d, "distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance"], "ObjectFspVlanInterface-Distance"); ok {
			if err = d.Set("distance", vv); err != nil {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("dns_query", flattenObjectFspVlanInterfaceDnsQuery2edl(o["dns-query"], d, "dns_query")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-query"], "ObjectFspVlanInterface-DnsQuery"); ok {
			if err = d.Set("dns_query", vv); err != nil {
				return fmt.Errorf("Error reading dns_query: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_query: %v", err)
		}
	}

	if err = d.Set("dns_server_override", flattenObjectFspVlanInterfaceDnsServerOverride2edl(o["dns-server-override"], d, "dns_server_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server-override"], "ObjectFspVlanInterface-DnsServerOverride"); ok {
			if err = d.Set("dns_server_override", vv); err != nil {
				return fmt.Errorf("Error reading dns_server_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server_override: %v", err)
		}
	}

	if err = d.Set("dns_server_protocol", flattenObjectFspVlanInterfaceDnsServerProtocol2edl(o["dns-server-protocol"], d, "dns_server_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server-protocol"], "ObjectFspVlanInterface-DnsServerProtocol"); ok {
			if err = d.Set("dns_server_protocol", vv); err != nil {
				return fmt.Errorf("Error reading dns_server_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server_protocol: %v", err)
		}
	}

	if err = d.Set("drop_fragment", flattenObjectFspVlanInterfaceDropFragment2edl(o["drop-fragment"], d, "drop_fragment")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop-fragment"], "ObjectFspVlanInterface-DropFragment"); ok {
			if err = d.Set("drop_fragment", vv); err != nil {
				return fmt.Errorf("Error reading drop_fragment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop_fragment: %v", err)
		}
	}

	if err = d.Set("drop_overlapped_fragment", flattenObjectFspVlanInterfaceDropOverlappedFragment2edl(o["drop-overlapped-fragment"], d, "drop_overlapped_fragment")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop-overlapped-fragment"], "ObjectFspVlanInterface-DropOverlappedFragment"); ok {
			if err = d.Set("drop_overlapped_fragment", vv); err != nil {
				return fmt.Errorf("Error reading drop_overlapped_fragment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop_overlapped_fragment: %v", err)
		}
	}

	if err = d.Set("eap_ca_cert", flattenObjectFspVlanInterfaceEapCaCert2edl(o["eap-ca-cert"], d, "eap_ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-ca-cert"], "ObjectFspVlanInterface-EapCaCert"); ok {
			if err = d.Set("eap_ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading eap_ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_ca_cert: %v", err)
		}
	}

	if err = d.Set("eap_identity", flattenObjectFspVlanInterfaceEapIdentity2edl(o["eap-identity"], d, "eap_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-identity"], "ObjectFspVlanInterface-EapIdentity"); ok {
			if err = d.Set("eap_identity", vv); err != nil {
				return fmt.Errorf("Error reading eap_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_identity: %v", err)
		}
	}

	if err = d.Set("eap_method", flattenObjectFspVlanInterfaceEapMethod2edl(o["eap-method"], d, "eap_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-method"], "ObjectFspVlanInterface-EapMethod"); ok {
			if err = d.Set("eap_method", vv); err != nil {
				return fmt.Errorf("Error reading eap_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_method: %v", err)
		}
	}

	if err = d.Set("eap_password", flattenObjectFspVlanInterfaceEapPassword2edl(o["eap-password"], d, "eap_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-password"], "ObjectFspVlanInterface-EapPassword"); ok {
			if err = d.Set("eap_password", vv); err != nil {
				return fmt.Errorf("Error reading eap_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_password: %v", err)
		}
	}

	if err = d.Set("eap_supplicant", flattenObjectFspVlanInterfaceEapSupplicant2edl(o["eap-supplicant"], d, "eap_supplicant")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-supplicant"], "ObjectFspVlanInterface-EapSupplicant"); ok {
			if err = d.Set("eap_supplicant", vv); err != nil {
				return fmt.Errorf("Error reading eap_supplicant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_supplicant: %v", err)
		}
	}

	if err = d.Set("eap_user_cert", flattenObjectFspVlanInterfaceEapUserCert2edl(o["eap-user-cert"], d, "eap_user_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-user-cert"], "ObjectFspVlanInterface-EapUserCert"); ok {
			if err = d.Set("eap_user_cert", vv); err != nil {
				return fmt.Errorf("Error reading eap_user_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_user_cert: %v", err)
		}
	}

	if err = d.Set("egress_cos", flattenObjectFspVlanInterfaceEgressCos2edl(o["egress-cos"], d, "egress_cos")); err != nil {
		if vv, ok := fortiAPIPatch(o["egress-cos"], "ObjectFspVlanInterface-EgressCos"); ok {
			if err = d.Set("egress_cos", vv); err != nil {
				return fmt.Errorf("Error reading egress_cos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading egress_cos: %v", err)
		}
	}

	if err = d.Set("egress_shaping_profile", flattenObjectFspVlanInterfaceEgressShapingProfile2edl(o["egress-shaping-profile"], d, "egress_shaping_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["egress-shaping-profile"], "ObjectFspVlanInterface-EgressShapingProfile"); ok {
			if err = d.Set("egress_shaping_profile", vv); err != nil {
				return fmt.Errorf("Error reading egress_shaping_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading egress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("eip", flattenObjectFspVlanInterfaceEip2edl(o["eip"], d, "eip")); err != nil {
		if vv, ok := fortiAPIPatch(o["eip"], "ObjectFspVlanInterface-Eip"); ok {
			if err = d.Set("eip", vv); err != nil {
				return fmt.Errorf("Error reading eip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eip: %v", err)
		}
	}

	if err = d.Set("endpoint_compliance", flattenObjectFspVlanInterfaceEndpointCompliance2edl(o["endpoint-compliance"], d, "endpoint_compliance")); err != nil {
		if vv, ok := fortiAPIPatch(o["endpoint-compliance"], "ObjectFspVlanInterface-EndpointCompliance"); ok {
			if err = d.Set("endpoint_compliance", vv); err != nil {
				return fmt.Errorf("Error reading endpoint_compliance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endpoint_compliance: %v", err)
		}
	}

	if err = d.Set("estimated_downstream_bandwidth", flattenObjectFspVlanInterfaceEstimatedDownstreamBandwidth2edl(o["estimated-downstream-bandwidth"], d, "estimated_downstream_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["estimated-downstream-bandwidth"], "ObjectFspVlanInterface-EstimatedDownstreamBandwidth"); ok {
			if err = d.Set("estimated_downstream_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading estimated_downstream_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading estimated_downstream_bandwidth: %v", err)
		}
	}

	if err = d.Set("estimated_upstream_bandwidth", flattenObjectFspVlanInterfaceEstimatedUpstreamBandwidth2edl(o["estimated-upstream-bandwidth"], d, "estimated_upstream_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["estimated-upstream-bandwidth"], "ObjectFspVlanInterface-EstimatedUpstreamBandwidth"); ok {
			if err = d.Set("estimated_upstream_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading estimated_upstream_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading estimated_upstream_bandwidth: %v", err)
		}
	}

	if err = d.Set("explicit_ftp_proxy", flattenObjectFspVlanInterfaceExplicitFtpProxy2edl(o["explicit-ftp-proxy"], d, "explicit_ftp_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["explicit-ftp-proxy"], "ObjectFspVlanInterface-ExplicitFtpProxy"); ok {
			if err = d.Set("explicit_ftp_proxy", vv); err != nil {
				return fmt.Errorf("Error reading explicit_ftp_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading explicit_ftp_proxy: %v", err)
		}
	}

	if err = d.Set("explicit_web_proxy", flattenObjectFspVlanInterfaceExplicitWebProxy2edl(o["explicit-web-proxy"], d, "explicit_web_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["explicit-web-proxy"], "ObjectFspVlanInterface-ExplicitWebProxy"); ok {
			if err = d.Set("explicit_web_proxy", vv); err != nil {
				return fmt.Errorf("Error reading explicit_web_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading explicit_web_proxy: %v", err)
		}
	}

	if err = d.Set("external", flattenObjectFspVlanInterfaceExternal2edl(o["external"], d, "external")); err != nil {
		if vv, ok := fortiAPIPatch(o["external"], "ObjectFspVlanInterface-External"); ok {
			if err = d.Set("external", vv); err != nil {
				return fmt.Errorf("Error reading external: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("fail_action_on_extender", flattenObjectFspVlanInterfaceFailActionOnExtender2edl(o["fail-action-on-extender"], d, "fail_action_on_extender")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-action-on-extender"], "ObjectFspVlanInterface-FailActionOnExtender"); ok {
			if err = d.Set("fail_action_on_extender", vv); err != nil {
				return fmt.Errorf("Error reading fail_action_on_extender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_action_on_extender: %v", err)
		}
	}

	if err = d.Set("fail_alert_interfaces", flattenObjectFspVlanInterfaceFailAlertInterfaces2edl(o["fail-alert-interfaces"], d, "fail_alert_interfaces")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-alert-interfaces"], "ObjectFspVlanInterface-FailAlertInterfaces"); ok {
			if err = d.Set("fail_alert_interfaces", vv); err != nil {
				return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
		}
	}

	if err = d.Set("fail_alert_method", flattenObjectFspVlanInterfaceFailAlertMethod2edl(o["fail-alert-method"], d, "fail_alert_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-alert-method"], "ObjectFspVlanInterface-FailAlertMethod"); ok {
			if err = d.Set("fail_alert_method", vv); err != nil {
				return fmt.Errorf("Error reading fail_alert_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_alert_method: %v", err)
		}
	}

	if err = d.Set("fail_detect", flattenObjectFspVlanInterfaceFailDetect2edl(o["fail-detect"], d, "fail_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-detect"], "ObjectFspVlanInterface-FailDetect"); ok {
			if err = d.Set("fail_detect", vv); err != nil {
				return fmt.Errorf("Error reading fail_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_detect: %v", err)
		}
	}

	if err = d.Set("fail_detect_option", flattenObjectFspVlanInterfaceFailDetectOption2edl(o["fail-detect-option"], d, "fail_detect_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-detect-option"], "ObjectFspVlanInterface-FailDetectOption"); ok {
			if err = d.Set("fail_detect_option", vv); err != nil {
				return fmt.Errorf("Error reading fail_detect_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_detect_option: %v", err)
		}
	}

	if err = d.Set("fdp", flattenObjectFspVlanInterfaceFdp2edl(o["fdp"], d, "fdp")); err != nil {
		if vv, ok := fortiAPIPatch(o["fdp"], "ObjectFspVlanInterface-Fdp"); ok {
			if err = d.Set("fdp", vv); err != nil {
				return fmt.Errorf("Error reading fdp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fdp: %v", err)
		}
	}

	if err = d.Set("fortiheartbeat", flattenObjectFspVlanInterfaceFortiheartbeat2edl(o["fortiheartbeat"], d, "fortiheartbeat")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiheartbeat"], "ObjectFspVlanInterface-Fortiheartbeat"); ok {
			if err = d.Set("fortiheartbeat", vv); err != nil {
				return fmt.Errorf("Error reading fortiheartbeat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiheartbeat: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenObjectFspVlanInterfaceFortilink2edl(o["fortilink"], d, "fortilink")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink"], "ObjectFspVlanInterface-Fortilink"); ok {
			if err = d.Set("fortilink", vv); err != nil {
				return fmt.Errorf("Error reading fortilink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("fortilink_backup_link", flattenObjectFspVlanInterfaceFortilinkBackupLink2edl(o["fortilink-backup-link"], d, "fortilink_backup_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink-backup-link"], "ObjectFspVlanInterface-FortilinkBackupLink"); ok {
			if err = d.Set("fortilink_backup_link", vv); err != nil {
				return fmt.Errorf("Error reading fortilink_backup_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink_backup_link: %v", err)
		}
	}

	if err = d.Set("fortilink_neighbor_detect", flattenObjectFspVlanInterfaceFortilinkNeighborDetect2edl(o["fortilink-neighbor-detect"], d, "fortilink_neighbor_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink-neighbor-detect"], "ObjectFspVlanInterface-FortilinkNeighborDetect"); ok {
			if err = d.Set("fortilink_neighbor_detect", vv); err != nil {
				return fmt.Errorf("Error reading fortilink_neighbor_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink_neighbor_detect: %v", err)
		}
	}

	if err = d.Set("fortilink_split_interface", flattenObjectFspVlanInterfaceFortilinkSplitInterface2edl(o["fortilink-split-interface"], d, "fortilink_split_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink-split-interface"], "ObjectFspVlanInterface-FortilinkSplitInterface"); ok {
			if err = d.Set("fortilink_split_interface", vv); err != nil {
				return fmt.Errorf("Error reading fortilink_split_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink_split_interface: %v", err)
		}
	}

	if err = d.Set("fortilink_stacking", flattenObjectFspVlanInterfaceFortilinkStacking2edl(o["fortilink-stacking"], d, "fortilink_stacking")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink-stacking"], "ObjectFspVlanInterface-FortilinkStacking"); ok {
			if err = d.Set("fortilink_stacking", vv); err != nil {
				return fmt.Errorf("Error reading fortilink_stacking: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink_stacking: %v", err)
		}
	}

	if err = d.Set("forward_domain", flattenObjectFspVlanInterfaceForwardDomain2edl(o["forward-domain"], d, "forward_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-domain"], "ObjectFspVlanInterface-ForwardDomain"); ok {
			if err = d.Set("forward_domain", vv); err != nil {
				return fmt.Errorf("Error reading forward_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_domain: %v", err)
		}
	}

	if err = d.Set("forward_error_correction", flattenObjectFspVlanInterfaceForwardErrorCorrection2edl(o["forward-error-correction"], d, "forward_error_correction")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-error-correction"], "ObjectFspVlanInterface-ForwardErrorCorrection"); ok {
			if err = d.Set("forward_error_correction", vv); err != nil {
				return fmt.Errorf("Error reading forward_error_correction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_error_correction: %v", err)
		}
	}

	if err = d.Set("fp_anomaly", flattenObjectFspVlanInterfaceFpAnomaly2edl(o["fp-anomaly"], d, "fp_anomaly")); err != nil {
		if vv, ok := fortiAPIPatch(o["fp-anomaly"], "ObjectFspVlanInterface-FpAnomaly"); ok {
			if err = d.Set("fp_anomaly", vv); err != nil {
				return fmt.Errorf("Error reading fp_anomaly: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fp_anomaly: %v", err)
		}
	}

	if err = d.Set("fp_disable", flattenObjectFspVlanInterfaceFpDisable2edl(o["fp-disable"], d, "fp_disable")); err != nil {
		if vv, ok := fortiAPIPatch(o["fp-disable"], "ObjectFspVlanInterface-FpDisable"); ok {
			if err = d.Set("fp_disable", vv); err != nil {
				return fmt.Errorf("Error reading fp_disable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fp_disable: %v", err)
		}
	}

	if err = d.Set("gateway_address", flattenObjectFspVlanInterfaceGatewayAddress2edl(o["gateway-address"], d, "gateway_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway-address"], "ObjectFspVlanInterface-GatewayAddress"); ok {
			if err = d.Set("gateway_address", vv); err != nil {
				return fmt.Errorf("Error reading gateway_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway_address: %v", err)
		}
	}

	if err = d.Set("generic_receive_offload", flattenObjectFspVlanInterfaceGenericReceiveOffload2edl(o["generic-receive-offload"], d, "generic_receive_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["generic-receive-offload"], "ObjectFspVlanInterface-GenericReceiveOffload"); ok {
			if err = d.Set("generic_receive_offload", vv); err != nil {
				return fmt.Errorf("Error reading generic_receive_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading generic_receive_offload: %v", err)
		}
	}

	if err = d.Set("gi_gk", flattenObjectFspVlanInterfaceGiGk2edl(o["gi-gk"], d, "gi_gk")); err != nil {
		if vv, ok := fortiAPIPatch(o["gi-gk"], "ObjectFspVlanInterface-GiGk"); ok {
			if err = d.Set("gi_gk", vv); err != nil {
				return fmt.Errorf("Error reading gi_gk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gi_gk: %v", err)
		}
	}

	if err = d.Set("gwaddr", flattenObjectFspVlanInterfaceGwaddr2edl(o["gwaddr"], d, "gwaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["gwaddr"], "ObjectFspVlanInterface-Gwaddr"); ok {
			if err = d.Set("gwaddr", vv); err != nil {
				return fmt.Errorf("Error reading gwaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gwaddr: %v", err)
		}
	}

	if err = d.Set("gwdetect", flattenObjectFspVlanInterfaceGwdetect2edl(o["gwdetect"], d, "gwdetect")); err != nil {
		if vv, ok := fortiAPIPatch(o["gwdetect"], "ObjectFspVlanInterface-Gwdetect"); ok {
			if err = d.Set("gwdetect", vv); err != nil {
				return fmt.Errorf("Error reading gwdetect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gwdetect: %v", err)
		}
	}

	if err = d.Set("ha_priority", flattenObjectFspVlanInterfaceHaPriority2edl(o["ha-priority"], d, "ha_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-priority"], "ObjectFspVlanInterface-HaPriority"); ok {
			if err = d.Set("ha_priority", vv); err != nil {
				return fmt.Errorf("Error reading ha_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("icmp_accept_redirect", flattenObjectFspVlanInterfaceIcmpAcceptRedirect2edl(o["icmp-accept-redirect"], d, "icmp_accept_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-accept-redirect"], "ObjectFspVlanInterface-IcmpAcceptRedirect"); ok {
			if err = d.Set("icmp_accept_redirect", vv); err != nil {
				return fmt.Errorf("Error reading icmp_accept_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_accept_redirect: %v", err)
		}
	}

	if err = d.Set("icmp_redirect", flattenObjectFspVlanInterfaceIcmpRedirect2edl(o["icmp-redirect"], d, "icmp_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-redirect"], "ObjectFspVlanInterface-IcmpRedirect"); ok {
			if err = d.Set("icmp_redirect", vv); err != nil {
				return fmt.Errorf("Error reading icmp_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_redirect: %v", err)
		}
	}

	if err = d.Set("icmp_send_redirect", flattenObjectFspVlanInterfaceIcmpSendRedirect2edl(o["icmp-send-redirect"], d, "icmp_send_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-send-redirect"], "ObjectFspVlanInterface-IcmpSendRedirect"); ok {
			if err = d.Set("icmp_send_redirect", vv); err != nil {
				return fmt.Errorf("Error reading icmp_send_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_send_redirect: %v", err)
		}
	}

	if err = d.Set("ident_accept", flattenObjectFspVlanInterfaceIdentAccept2edl(o["ident-accept"], d, "ident_accept")); err != nil {
		if vv, ok := fortiAPIPatch(o["ident-accept"], "ObjectFspVlanInterface-IdentAccept"); ok {
			if err = d.Set("ident_accept", vv); err != nil {
				return fmt.Errorf("Error reading ident_accept: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ident_accept: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenObjectFspVlanInterfaceIdleTimeout2edl(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle-timeout"], "ObjectFspVlanInterface-IdleTimeout"); ok {
			if err = d.Set("idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("if_mdix", flattenObjectFspVlanInterfaceIfMdix2edl(o["if-mdix"], d, "if_mdix")); err != nil {
		if vv, ok := fortiAPIPatch(o["if-mdix"], "ObjectFspVlanInterface-IfMdix"); ok {
			if err = d.Set("if_mdix", vv); err != nil {
				return fmt.Errorf("Error reading if_mdix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading if_mdix: %v", err)
		}
	}

	if err = d.Set("if_media", flattenObjectFspVlanInterfaceIfMedia2edl(o["if-media"], d, "if_media")); err != nil {
		if vv, ok := fortiAPIPatch(o["if-media"], "ObjectFspVlanInterface-IfMedia"); ok {
			if err = d.Set("if_media", vv); err != nil {
				return fmt.Errorf("Error reading if_media: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading if_media: %v", err)
		}
	}

	if err = d.Set("ike_saml_server", flattenObjectFspVlanInterfaceIkeSamlServer2edl(o["ike-saml-server"], d, "ike_saml_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-saml-server"], "ObjectFspVlanInterface-IkeSamlServer"); ok {
			if err = d.Set("ike_saml_server", vv); err != nil {
				return fmt.Errorf("Error reading ike_saml_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_saml_server: %v", err)
		}
	}

	if err = d.Set("in_force_vlan_cos", flattenObjectFspVlanInterfaceInForceVlanCos2edl(o["in-force-vlan-cos"], d, "in_force_vlan_cos")); err != nil {
		if vv, ok := fortiAPIPatch(o["in-force-vlan-cos"], "ObjectFspVlanInterface-InForceVlanCos"); ok {
			if err = d.Set("in_force_vlan_cos", vv); err != nil {
				return fmt.Errorf("Error reading in_force_vlan_cos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading in_force_vlan_cos: %v", err)
		}
	}

	if err = d.Set("inbandwidth", flattenObjectFspVlanInterfaceInbandwidth2edl(o["inbandwidth"], d, "inbandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbandwidth"], "ObjectFspVlanInterface-Inbandwidth"); ok {
			if err = d.Set("inbandwidth", vv); err != nil {
				return fmt.Errorf("Error reading inbandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbandwidth: %v", err)
		}
	}

	if err = d.Set("ingress_cos", flattenObjectFspVlanInterfaceIngressCos2edl(o["ingress-cos"], d, "ingress_cos")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-cos"], "ObjectFspVlanInterface-IngressCos"); ok {
			if err = d.Set("ingress_cos", vv); err != nil {
				return fmt.Errorf("Error reading ingress_cos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_cos: %v", err)
		}
	}

	if err = d.Set("ingress_shaping_profile", flattenObjectFspVlanInterfaceIngressShapingProfile2edl(o["ingress-shaping-profile"], d, "ingress_shaping_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-shaping-profile"], "ObjectFspVlanInterface-IngressShapingProfile"); ok {
			if err = d.Set("ingress_shaping_profile", vv); err != nil {
				return fmt.Errorf("Error reading ingress_shaping_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("ingress_spillover_threshold", flattenObjectFspVlanInterfaceIngressSpilloverThreshold2edl(o["ingress-spillover-threshold"], d, "ingress_spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-spillover-threshold"], "ObjectFspVlanInterface-IngressSpilloverThreshold"); ok {
			if err = d.Set("ingress_spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
		}
	}

	if err = d.Set("interconnect_profile", flattenObjectFspVlanInterfaceInterconnectProfile2edl(o["interconnect-profile"], d, "interconnect_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["interconnect-profile"], "ObjectFspVlanInterface-InterconnectProfile"); ok {
			if err = d.Set("interconnect_profile", vv); err != nil {
				return fmt.Errorf("Error reading interconnect_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interconnect_profile: %v", err)
		}
	}

	if err = d.Set("internal", flattenObjectFspVlanInterfaceInternal2edl(o["internal"], d, "internal")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal"], "ObjectFspVlanInterface-Internal"); ok {
			if err = d.Set("internal", vv); err != nil {
				return fmt.Errorf("Error reading internal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectFspVlanInterfaceIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectFspVlanInterface-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip_managed_by_fortiipam", flattenObjectFspVlanInterfaceIpManagedByFortiipam2edl(o["ip-managed-by-fortiipam"], d, "ip_managed_by_fortiipam")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-managed-by-fortiipam"], "ObjectFspVlanInterface-IpManagedByFortiipam"); ok {
			if err = d.Set("ip_managed_by_fortiipam", vv); err != nil {
				return fmt.Errorf("Error reading ip_managed_by_fortiipam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_managed_by_fortiipam: %v", err)
		}
	}

	if err = d.Set("ipmac", flattenObjectFspVlanInterfaceIpmac2edl(o["ipmac"], d, "ipmac")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipmac"], "ObjectFspVlanInterface-Ipmac"); ok {
			if err = d.Set("ipmac", vv); err != nil {
				return fmt.Errorf("Error reading ipmac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipmac: %v", err)
		}
	}

	if err = d.Set("ips_sniffer_mode", flattenObjectFspVlanInterfaceIpsSnifferMode2edl(o["ips-sniffer-mode"], d, "ips_sniffer_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sniffer-mode"], "ObjectFspVlanInterface-IpsSnifferMode"); ok {
			if err = d.Set("ips_sniffer_mode", vv); err != nil {
				return fmt.Errorf("Error reading ips_sniffer_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sniffer_mode: %v", err)
		}
	}

	if err = d.Set("ipunnumbered", flattenObjectFspVlanInterfaceIpunnumbered2edl(o["ipunnumbered"], d, "ipunnumbered")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipunnumbered"], "ObjectFspVlanInterface-Ipunnumbered"); ok {
			if err = d.Set("ipunnumbered", vv); err != nil {
				return fmt.Errorf("Error reading ipunnumbered: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipunnumbered: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipv6", flattenObjectFspVlanInterfaceIpv62edl(o["ipv6"], d, "ipv6")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipv6"], "ObjectFspVlanInterface-Ipv6"); ok {
				if err = d.Set("ipv6", vv); err != nil {
					return fmt.Errorf("Error reading ipv6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6"); ok {
			if err = d.Set("ipv6", flattenObjectFspVlanInterfaceIpv62edl(o["ipv6"], d, "ipv6")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipv6"], "ObjectFspVlanInterface-Ipv6"); ok {
					if err = d.Set("ipv6", vv); err != nil {
						return fmt.Errorf("Error reading ipv6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipv6: %v", err)
				}
			}
		}
	}

	if err = d.Set("l2forward", flattenObjectFspVlanInterfaceL2Forward2edl(o["l2forward"], d, "l2forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["l2forward"], "ObjectFspVlanInterface-L2Forward"); ok {
			if err = d.Set("l2forward", vv); err != nil {
				return fmt.Errorf("Error reading l2forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l2forward: %v", err)
		}
	}

	if err = d.Set("l2tp_client", flattenObjectFspVlanInterfaceL2TpClient2edl(o["l2tp-client"], d, "l2tp_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["l2tp-client"], "ObjectFspVlanInterface-L2TpClient"); ok {
			if err = d.Set("l2tp_client", vv); err != nil {
				return fmt.Errorf("Error reading l2tp_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l2tp_client: %v", err)
		}
	}

	if err = d.Set("lacp_ha_secondary", flattenObjectFspVlanInterfaceLacpHaSecondary2edl(o["lacp-ha-secondary"], d, "lacp_ha_secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["lacp-ha-secondary"], "ObjectFspVlanInterface-LacpHaSecondary"); ok {
			if err = d.Set("lacp_ha_secondary", vv); err != nil {
				return fmt.Errorf("Error reading lacp_ha_secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lacp_ha_secondary: %v", err)
		}
	}

	if err = d.Set("lacp_ha_slave", flattenObjectFspVlanInterfaceLacpHaSlave2edl(o["lacp-ha-slave"], d, "lacp_ha_slave")); err != nil {
		if vv, ok := fortiAPIPatch(o["lacp-ha-slave"], "ObjectFspVlanInterface-LacpHaSlave"); ok {
			if err = d.Set("lacp_ha_slave", vv); err != nil {
				return fmt.Errorf("Error reading lacp_ha_slave: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lacp_ha_slave: %v", err)
		}
	}

	if err = d.Set("lacp_mode", flattenObjectFspVlanInterfaceLacpMode2edl(o["lacp-mode"], d, "lacp_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["lacp-mode"], "ObjectFspVlanInterface-LacpMode"); ok {
			if err = d.Set("lacp_mode", vv); err != nil {
				return fmt.Errorf("Error reading lacp_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lacp_mode: %v", err)
		}
	}

	if err = d.Set("lacp_speed", flattenObjectFspVlanInterfaceLacpSpeed2edl(o["lacp-speed"], d, "lacp_speed")); err != nil {
		if vv, ok := fortiAPIPatch(o["lacp-speed"], "ObjectFspVlanInterface-LacpSpeed"); ok {
			if err = d.Set("lacp_speed", vv); err != nil {
				return fmt.Errorf("Error reading lacp_speed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lacp_speed: %v", err)
		}
	}

	if err = d.Set("large_receive_offload", flattenObjectFspVlanInterfaceLargeReceiveOffload2edl(o["large-receive-offload"], d, "large_receive_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["large-receive-offload"], "ObjectFspVlanInterface-LargeReceiveOffload"); ok {
			if err = d.Set("large_receive_offload", vv); err != nil {
				return fmt.Errorf("Error reading large_receive_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading large_receive_offload: %v", err)
		}
	}

	if err = d.Set("lcp_echo_interval", flattenObjectFspVlanInterfaceLcpEchoInterval2edl(o["lcp-echo-interval"], d, "lcp_echo_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["lcp-echo-interval"], "ObjectFspVlanInterface-LcpEchoInterval"); ok {
			if err = d.Set("lcp_echo_interval", vv); err != nil {
				return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
		}
	}

	if err = d.Set("lcp_max_echo_fails", flattenObjectFspVlanInterfaceLcpMaxEchoFails2edl(o["lcp-max-echo-fails"], d, "lcp_max_echo_fails")); err != nil {
		if vv, ok := fortiAPIPatch(o["lcp-max-echo-fails"], "ObjectFspVlanInterface-LcpMaxEchoFails"); ok {
			if err = d.Set("lcp_max_echo_fails", vv); err != nil {
				return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
		}
	}

	if err = d.Set("link_up_delay", flattenObjectFspVlanInterfaceLinkUpDelay2edl(o["link-up-delay"], d, "link_up_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-up-delay"], "ObjectFspVlanInterface-LinkUpDelay"); ok {
			if err = d.Set("link_up_delay", vv); err != nil {
				return fmt.Errorf("Error reading link_up_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_up_delay: %v", err)
		}
	}

	if err = d.Set("listen_forticlient_connection", flattenObjectFspVlanInterfaceListenForticlientConnection2edl(o["listen-forticlient-connection"], d, "listen_forticlient_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["listen-forticlient-connection"], "ObjectFspVlanInterface-ListenForticlientConnection"); ok {
			if err = d.Set("listen_forticlient_connection", vv); err != nil {
				return fmt.Errorf("Error reading listen_forticlient_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading listen_forticlient_connection: %v", err)
		}
	}

	if err = d.Set("lldp_network_policy", flattenObjectFspVlanInterfaceLldpNetworkPolicy2edl(o["lldp-network-policy"], d, "lldp_network_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-network-policy"], "ObjectFspVlanInterface-LldpNetworkPolicy"); ok {
			if err = d.Set("lldp_network_policy", vv); err != nil {
				return fmt.Errorf("Error reading lldp_network_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_network_policy: %v", err)
		}
	}

	if err = d.Set("lldp_reception", flattenObjectFspVlanInterfaceLldpReception2edl(o["lldp-reception"], d, "lldp_reception")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-reception"], "ObjectFspVlanInterface-LldpReception"); ok {
			if err = d.Set("lldp_reception", vv); err != nil {
				return fmt.Errorf("Error reading lldp_reception: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_reception: %v", err)
		}
	}

	if err = d.Set("lldp_transmission", flattenObjectFspVlanInterfaceLldpTransmission2edl(o["lldp-transmission"], d, "lldp_transmission")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-transmission"], "ObjectFspVlanInterface-LldpTransmission"); ok {
			if err = d.Set("lldp_transmission", vv); err != nil {
				return fmt.Errorf("Error reading lldp_transmission: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_transmission: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectFspVlanInterfaceLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectFspVlanInterface-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("macaddr", flattenObjectFspVlanInterfaceMacaddr2edl(o["macaddr"], d, "macaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["macaddr"], "ObjectFspVlanInterface-Macaddr"); ok {
			if err = d.Set("macaddr", vv); err != nil {
				return fmt.Errorf("Error reading macaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading macaddr: %v", err)
		}
	}

	if err = d.Set("managed_subnetwork_size", flattenObjectFspVlanInterfaceManagedSubnetworkSize2edl(o["managed-subnetwork-size"], d, "managed_subnetwork_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["managed-subnetwork-size"], "ObjectFspVlanInterface-ManagedSubnetworkSize"); ok {
			if err = d.Set("managed_subnetwork_size", vv); err != nil {
				return fmt.Errorf("Error reading managed_subnetwork_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading managed_subnetwork_size: %v", err)
		}
	}

	if err = d.Set("management_ip", flattenObjectFspVlanInterfaceManagementIp2edl(o["management-ip"], d, "management_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["management-ip"], "ObjectFspVlanInterface-ManagementIp"); ok {
			if err = d.Set("management_ip", vv); err != nil {
				return fmt.Errorf("Error reading management_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("max_egress_burst_rate", flattenObjectFspVlanInterfaceMaxEgressBurstRate2edl(o["max-egress-burst-rate"], d, "max_egress_burst_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-egress-burst-rate"], "ObjectFspVlanInterface-MaxEgressBurstRate"); ok {
			if err = d.Set("max_egress_burst_rate", vv); err != nil {
				return fmt.Errorf("Error reading max_egress_burst_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_egress_burst_rate: %v", err)
		}
	}

	if err = d.Set("max_egress_rate", flattenObjectFspVlanInterfaceMaxEgressRate2edl(o["max-egress-rate"], d, "max_egress_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-egress-rate"], "ObjectFspVlanInterface-MaxEgressRate"); ok {
			if err = d.Set("max_egress_rate", vv); err != nil {
				return fmt.Errorf("Error reading max_egress_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_egress_rate: %v", err)
		}
	}

	if err = d.Set("measured_downstream_bandwidth", flattenObjectFspVlanInterfaceMeasuredDownstreamBandwidth2edl(o["measured-downstream-bandwidth"], d, "measured_downstream_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["measured-downstream-bandwidth"], "ObjectFspVlanInterface-MeasuredDownstreamBandwidth"); ok {
			if err = d.Set("measured_downstream_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading measured_downstream_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading measured_downstream_bandwidth: %v", err)
		}
	}

	if err = d.Set("measured_upstream_bandwidth", flattenObjectFspVlanInterfaceMeasuredUpstreamBandwidth2edl(o["measured-upstream-bandwidth"], d, "measured_upstream_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["measured-upstream-bandwidth"], "ObjectFspVlanInterface-MeasuredUpstreamBandwidth"); ok {
			if err = d.Set("measured_upstream_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading measured_upstream_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading measured_upstream_bandwidth: %v", err)
		}
	}

	if err = d.Set("mediatype", flattenObjectFspVlanInterfaceMediatype2edl(o["mediatype"], d, "mediatype")); err != nil {
		if vv, ok := fortiAPIPatch(o["mediatype"], "ObjectFspVlanInterface-Mediatype"); ok {
			if err = d.Set("mediatype", vv); err != nil {
				return fmt.Errorf("Error reading mediatype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mediatype: %v", err)
		}
	}

	if err = d.Set("member", flattenObjectFspVlanInterfaceMember2edl(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "ObjectFspVlanInterface-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("min_links", flattenObjectFspVlanInterfaceMinLinks2edl(o["min-links"], d, "min_links")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-links"], "ObjectFspVlanInterface-MinLinks"); ok {
			if err = d.Set("min_links", vv); err != nil {
				return fmt.Errorf("Error reading min_links: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_links: %v", err)
		}
	}

	if err = d.Set("min_links_down", flattenObjectFspVlanInterfaceMinLinksDown2edl(o["min-links-down"], d, "min_links_down")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-links-down"], "ObjectFspVlanInterface-MinLinksDown"); ok {
			if err = d.Set("min_links_down", vv); err != nil {
				return fmt.Errorf("Error reading min_links_down: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_links_down: %v", err)
		}
	}

	if err = d.Set("mode", flattenObjectFspVlanInterfaceMode2edl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "ObjectFspVlanInterface-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("monitor_bandwidth", flattenObjectFspVlanInterfaceMonitorBandwidth2edl(o["monitor-bandwidth"], d, "monitor_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-bandwidth"], "ObjectFspVlanInterface-MonitorBandwidth"); ok {
			if err = d.Set("monitor_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading monitor_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_bandwidth: %v", err)
		}
	}

	if err = d.Set("mtu", flattenObjectFspVlanInterfaceMtu2edl(o["mtu"], d, "mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtu"], "ObjectFspVlanInterface-Mtu"); ok {
			if err = d.Set("mtu", vv); err != nil {
				return fmt.Errorf("Error reading mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("mtu_override", flattenObjectFspVlanInterfaceMtuOverride2edl(o["mtu-override"], d, "mtu_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtu-override"], "ObjectFspVlanInterface-MtuOverride"); ok {
			if err = d.Set("mtu_override", vv); err != nil {
				return fmt.Errorf("Error reading mtu_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtu_override: %v", err)
		}
	}

	if err = d.Set("mux_type", flattenObjectFspVlanInterfaceMuxType2edl(o["mux-type"], d, "mux_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["mux-type"], "ObjectFspVlanInterface-MuxType"); ok {
			if err = d.Set("mux_type", vv); err != nil {
				return fmt.Errorf("Error reading mux_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mux_type: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFspVlanInterfaceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFspVlanInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ndiscforward", flattenObjectFspVlanInterfaceNdiscforward2edl(o["ndiscforward"], d, "ndiscforward")); err != nil {
		if vv, ok := fortiAPIPatch(o["ndiscforward"], "ObjectFspVlanInterface-Ndiscforward"); ok {
			if err = d.Set("ndiscforward", vv); err != nil {
				return fmt.Errorf("Error reading ndiscforward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ndiscforward: %v", err)
		}
	}

	if err = d.Set("netbios_forward", flattenObjectFspVlanInterfaceNetbiosForward2edl(o["netbios-forward"], d, "netbios_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["netbios-forward"], "ObjectFspVlanInterface-NetbiosForward"); ok {
			if err = d.Set("netbios_forward", vv); err != nil {
				return fmt.Errorf("Error reading netbios_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netbios_forward: %v", err)
		}
	}

	if err = d.Set("netflow_sampler", flattenObjectFspVlanInterfaceNetflowSampler2edl(o["netflow-sampler"], d, "netflow_sampler")); err != nil {
		if vv, ok := fortiAPIPatch(o["netflow-sampler"], "ObjectFspVlanInterface-NetflowSampler"); ok {
			if err = d.Set("netflow_sampler", vv); err != nil {
				return fmt.Errorf("Error reading netflow_sampler: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netflow_sampler: %v", err)
		}
	}

	if err = d.Set("np_qos_profile", flattenObjectFspVlanInterfaceNpQosProfile2edl(o["np-qos-profile"], d, "np_qos_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["np-qos-profile"], "ObjectFspVlanInterface-NpQosProfile"); ok {
			if err = d.Set("np_qos_profile", vv); err != nil {
				return fmt.Errorf("Error reading np_qos_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading np_qos_profile: %v", err)
		}
	}

	if err = d.Set("npu_fastpath", flattenObjectFspVlanInterfaceNpuFastpath2edl(o["npu-fastpath"], d, "npu_fastpath")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-fastpath"], "ObjectFspVlanInterface-NpuFastpath"); ok {
			if err = d.Set("npu_fastpath", vv); err != nil {
				return fmt.Errorf("Error reading npu_fastpath: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_fastpath: %v", err)
		}
	}

	if err = d.Set("nst", flattenObjectFspVlanInterfaceNst2edl(o["nst"], d, "nst")); err != nil {
		if vv, ok := fortiAPIPatch(o["nst"], "ObjectFspVlanInterface-Nst"); ok {
			if err = d.Set("nst", vv); err != nil {
				return fmt.Errorf("Error reading nst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nst: %v", err)
		}
	}

	if err = d.Set("out_force_vlan_cos", flattenObjectFspVlanInterfaceOutForceVlanCos2edl(o["out-force-vlan-cos"], d, "out_force_vlan_cos")); err != nil {
		if vv, ok := fortiAPIPatch(o["out-force-vlan-cos"], "ObjectFspVlanInterface-OutForceVlanCos"); ok {
			if err = d.Set("out_force_vlan_cos", vv); err != nil {
				return fmt.Errorf("Error reading out_force_vlan_cos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading out_force_vlan_cos: %v", err)
		}
	}

	if err = d.Set("outbandwidth", flattenObjectFspVlanInterfaceOutbandwidth2edl(o["outbandwidth"], d, "outbandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbandwidth"], "ObjectFspVlanInterface-Outbandwidth"); ok {
			if err = d.Set("outbandwidth", vv); err != nil {
				return fmt.Errorf("Error reading outbandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbandwidth: %v", err)
		}
	}

	if err = d.Set("padt_retry_timeout", flattenObjectFspVlanInterfacePadtRetryTimeout2edl(o["padt-retry-timeout"], d, "padt_retry_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["padt-retry-timeout"], "ObjectFspVlanInterface-PadtRetryTimeout"); ok {
			if err = d.Set("padt_retry_timeout", vv); err != nil {
				return fmt.Errorf("Error reading padt_retry_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading padt_retry_timeout: %v", err)
		}
	}

	if err = d.Set("password", flattenObjectFspVlanInterfacePassword2edl(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "ObjectFspVlanInterface-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("peer_interface", flattenObjectFspVlanInterfacePeerInterface2edl(o["peer-interface"], d, "peer_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-interface"], "ObjectFspVlanInterface-PeerInterface"); ok {
			if err = d.Set("peer_interface", vv); err != nil {
				return fmt.Errorf("Error reading peer_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_interface: %v", err)
		}
	}

	if err = d.Set("phy_mode", flattenObjectFspVlanInterfacePhyMode2edl(o["phy-mode"], d, "phy_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["phy-mode"], "ObjectFspVlanInterface-PhyMode"); ok {
			if err = d.Set("phy_mode", vv); err != nil {
				return fmt.Errorf("Error reading phy_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phy_mode: %v", err)
		}
	}

	if err = d.Set("ping_serv_status", flattenObjectFspVlanInterfacePingServStatus2edl(o["ping-serv-status"], d, "ping_serv_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ping-serv-status"], "ObjectFspVlanInterface-PingServStatus"); ok {
			if err = d.Set("ping_serv_status", vv); err != nil {
				return fmt.Errorf("Error reading ping_serv_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ping_serv_status: %v", err)
		}
	}

	if err = d.Set("poe", flattenObjectFspVlanInterfacePoe2edl(o["poe"], d, "poe")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe"], "ObjectFspVlanInterface-Poe"); ok {
			if err = d.Set("poe", vv); err != nil {
				return fmt.Errorf("Error reading poe: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe: %v", err)
		}
	}

	if err = d.Set("polling_interval", flattenObjectFspVlanInterfacePollingInterval2edl(o["polling-interval"], d, "polling_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["polling-interval"], "ObjectFspVlanInterface-PollingInterval"); ok {
			if err = d.Set("polling_interval", vv); err != nil {
				return fmt.Errorf("Error reading polling_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polling_interval: %v", err)
		}
	}

	if err = d.Set("pppoe_unnumbered_negotiate", flattenObjectFspVlanInterfacePppoeUnnumberedNegotiate2edl(o["pppoe-unnumbered-negotiate"], d, "pppoe_unnumbered_negotiate")); err != nil {
		if vv, ok := fortiAPIPatch(o["pppoe-unnumbered-negotiate"], "ObjectFspVlanInterface-PppoeUnnumberedNegotiate"); ok {
			if err = d.Set("pppoe_unnumbered_negotiate", vv); err != nil {
				return fmt.Errorf("Error reading pppoe_unnumbered_negotiate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pppoe_unnumbered_negotiate: %v", err)
		}
	}

	if err = d.Set("pptp_auth_type", flattenObjectFspVlanInterfacePptpAuthType2edl(o["pptp-auth-type"], d, "pptp_auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["pptp-auth-type"], "ObjectFspVlanInterface-PptpAuthType"); ok {
			if err = d.Set("pptp_auth_type", vv); err != nil {
				return fmt.Errorf("Error reading pptp_auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pptp_auth_type: %v", err)
		}
	}

	if err = d.Set("pptp_client", flattenObjectFspVlanInterfacePptpClient2edl(o["pptp-client"], d, "pptp_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["pptp-client"], "ObjectFspVlanInterface-PptpClient"); ok {
			if err = d.Set("pptp_client", vv); err != nil {
				return fmt.Errorf("Error reading pptp_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pptp_client: %v", err)
		}
	}

	if err = d.Set("pptp_password", flattenObjectFspVlanInterfacePptpPassword2edl(o["pptp-password"], d, "pptp_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["pptp-password"], "ObjectFspVlanInterface-PptpPassword"); ok {
			if err = d.Set("pptp_password", vv); err != nil {
				return fmt.Errorf("Error reading pptp_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pptp_password: %v", err)
		}
	}

	if err = d.Set("pptp_server_ip", flattenObjectFspVlanInterfacePptpServerIp2edl(o["pptp-server-ip"], d, "pptp_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["pptp-server-ip"], "ObjectFspVlanInterface-PptpServerIp"); ok {
			if err = d.Set("pptp_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading pptp_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pptp_server_ip: %v", err)
		}
	}

	if err = d.Set("pptp_timeout", flattenObjectFspVlanInterfacePptpTimeout2edl(o["pptp-timeout"], d, "pptp_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pptp-timeout"], "ObjectFspVlanInterface-PptpTimeout"); ok {
			if err = d.Set("pptp_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pptp_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pptp_timeout: %v", err)
		}
	}

	if err = d.Set("pptp_user", flattenObjectFspVlanInterfacePptpUser2edl(o["pptp-user"], d, "pptp_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["pptp-user"], "ObjectFspVlanInterface-PptpUser"); ok {
			if err = d.Set("pptp_user", vv); err != nil {
				return fmt.Errorf("Error reading pptp_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pptp_user: %v", err)
		}
	}

	if err = d.Set("preserve_session_route", flattenObjectFspVlanInterfacePreserveSessionRoute2edl(o["preserve-session-route"], d, "preserve_session_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["preserve-session-route"], "ObjectFspVlanInterface-PreserveSessionRoute"); ok {
			if err = d.Set("preserve_session_route", vv); err != nil {
				return fmt.Errorf("Error reading preserve_session_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preserve_session_route: %v", err)
		}
	}

	if err = d.Set("priority", flattenObjectFspVlanInterfacePriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ObjectFspVlanInterface-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("priority_override", flattenObjectFspVlanInterfacePriorityOverride2edl(o["priority-override"], d, "priority_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-override"], "ObjectFspVlanInterface-PriorityOverride"); ok {
			if err = d.Set("priority_override", vv); err != nil {
				return fmt.Errorf("Error reading priority_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_override: %v", err)
		}
	}

	if err = d.Set("proxy_captive_portal", flattenObjectFspVlanInterfaceProxyCaptivePortal2edl(o["proxy-captive-portal"], d, "proxy_captive_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-captive-portal"], "ObjectFspVlanInterface-ProxyCaptivePortal"); ok {
			if err = d.Set("proxy_captive_portal", vv); err != nil {
				return fmt.Errorf("Error reading proxy_captive_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_captive_portal: %v", err)
		}
	}

	if err = d.Set("pvc_atm_qos", flattenObjectFspVlanInterfacePvcAtmQos2edl(o["pvc-atm-qos"], d, "pvc_atm_qos")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvc-atm-qos"], "ObjectFspVlanInterface-PvcAtmQos"); ok {
			if err = d.Set("pvc_atm_qos", vv); err != nil {
				return fmt.Errorf("Error reading pvc_atm_qos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvc_atm_qos: %v", err)
		}
	}

	if err = d.Set("pvc_chan", flattenObjectFspVlanInterfacePvcChan2edl(o["pvc-chan"], d, "pvc_chan")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvc-chan"], "ObjectFspVlanInterface-PvcChan"); ok {
			if err = d.Set("pvc_chan", vv); err != nil {
				return fmt.Errorf("Error reading pvc_chan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvc_chan: %v", err)
		}
	}

	if err = d.Set("pvc_crc", flattenObjectFspVlanInterfacePvcCrc2edl(o["pvc-crc"], d, "pvc_crc")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvc-crc"], "ObjectFspVlanInterface-PvcCrc"); ok {
			if err = d.Set("pvc_crc", vv); err != nil {
				return fmt.Errorf("Error reading pvc_crc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvc_crc: %v", err)
		}
	}

	if err = d.Set("pvc_pcr", flattenObjectFspVlanInterfacePvcPcr2edl(o["pvc-pcr"], d, "pvc_pcr")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvc-pcr"], "ObjectFspVlanInterface-PvcPcr"); ok {
			if err = d.Set("pvc_pcr", vv); err != nil {
				return fmt.Errorf("Error reading pvc_pcr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvc_pcr: %v", err)
		}
	}

	if err = d.Set("pvc_scr", flattenObjectFspVlanInterfacePvcScr2edl(o["pvc-scr"], d, "pvc_scr")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvc-scr"], "ObjectFspVlanInterface-PvcScr"); ok {
			if err = d.Set("pvc_scr", vv); err != nil {
				return fmt.Errorf("Error reading pvc_scr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvc_scr: %v", err)
		}
	}

	if err = d.Set("pvc_vlan_id", flattenObjectFspVlanInterfacePvcVlanId2edl(o["pvc-vlan-id"], d, "pvc_vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvc-vlan-id"], "ObjectFspVlanInterface-PvcVlanId"); ok {
			if err = d.Set("pvc_vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading pvc_vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvc_vlan_id: %v", err)
		}
	}

	if err = d.Set("pvc_vlan_rx_id", flattenObjectFspVlanInterfacePvcVlanRxId2edl(o["pvc-vlan-rx-id"], d, "pvc_vlan_rx_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvc-vlan-rx-id"], "ObjectFspVlanInterface-PvcVlanRxId"); ok {
			if err = d.Set("pvc_vlan_rx_id", vv); err != nil {
				return fmt.Errorf("Error reading pvc_vlan_rx_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvc_vlan_rx_id: %v", err)
		}
	}

	if err = d.Set("pvc_vlan_rx_op", flattenObjectFspVlanInterfacePvcVlanRxOp2edl(o["pvc-vlan-rx-op"], d, "pvc_vlan_rx_op")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvc-vlan-rx-op"], "ObjectFspVlanInterface-PvcVlanRxOp"); ok {
			if err = d.Set("pvc_vlan_rx_op", vv); err != nil {
				return fmt.Errorf("Error reading pvc_vlan_rx_op: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvc_vlan_rx_op: %v", err)
		}
	}

	if err = d.Set("pvc_vlan_tx_id", flattenObjectFspVlanInterfacePvcVlanTxId2edl(o["pvc-vlan-tx-id"], d, "pvc_vlan_tx_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvc-vlan-tx-id"], "ObjectFspVlanInterface-PvcVlanTxId"); ok {
			if err = d.Set("pvc_vlan_tx_id", vv); err != nil {
				return fmt.Errorf("Error reading pvc_vlan_tx_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvc_vlan_tx_id: %v", err)
		}
	}

	if err = d.Set("pvc_vlan_tx_op", flattenObjectFspVlanInterfacePvcVlanTxOp2edl(o["pvc-vlan-tx-op"], d, "pvc_vlan_tx_op")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvc-vlan-tx-op"], "ObjectFspVlanInterface-PvcVlanTxOp"); ok {
			if err = d.Set("pvc_vlan_tx_op", vv); err != nil {
				return fmt.Errorf("Error reading pvc_vlan_tx_op: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvc_vlan_tx_op: %v", err)
		}
	}

	if err = d.Set("reachable_time", flattenObjectFspVlanInterfaceReachableTime2edl(o["reachable-time"], d, "reachable_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["reachable-time"], "ObjectFspVlanInterface-ReachableTime"); ok {
			if err = d.Set("reachable_time", vv); err != nil {
				return fmt.Errorf("Error reading reachable_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reachable_time: %v", err)
		}
	}

	if err = d.Set("redundant_interface", flattenObjectFspVlanInterfaceRedundantInterface2edl(o["redundant-interface"], d, "redundant_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant-interface"], "ObjectFspVlanInterface-RedundantInterface"); ok {
			if err = d.Set("redundant_interface", vv); err != nil {
				return fmt.Errorf("Error reading redundant_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant_interface: %v", err)
		}
	}

	if err = d.Set("remote_ip", flattenObjectFspVlanInterfaceRemoteIp2edl(o["remote-ip"], d, "remote_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-ip"], "ObjectFspVlanInterface-RemoteIp"); ok {
			if err = d.Set("remote_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_ip: %v", err)
		}
	}

	if err = d.Set("replacemsg_override_group", flattenObjectFspVlanInterfaceReplacemsgOverrideGroup2edl(o["replacemsg-override-group"], d, "replacemsg_override_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-override-group"], "ObjectFspVlanInterface-ReplacemsgOverrideGroup"); ok {
			if err = d.Set("replacemsg_override_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("retransmission", flattenObjectFspVlanInterfaceRetransmission2edl(o["retransmission"], d, "retransmission")); err != nil {
		if vv, ok := fortiAPIPatch(o["retransmission"], "ObjectFspVlanInterface-Retransmission"); ok {
			if err = d.Set("retransmission", vv); err != nil {
				return fmt.Errorf("Error reading retransmission: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retransmission: %v", err)
		}
	}

	if err = d.Set("ring_rx", flattenObjectFspVlanInterfaceRingRx2edl(o["ring-rx"], d, "ring_rx")); err != nil {
		if vv, ok := fortiAPIPatch(o["ring-rx"], "ObjectFspVlanInterface-RingRx"); ok {
			if err = d.Set("ring_rx", vv); err != nil {
				return fmt.Errorf("Error reading ring_rx: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ring_rx: %v", err)
		}
	}

	if err = d.Set("ring_tx", flattenObjectFspVlanInterfaceRingTx2edl(o["ring-tx"], d, "ring_tx")); err != nil {
		if vv, ok := fortiAPIPatch(o["ring-tx"], "ObjectFspVlanInterface-RingTx"); ok {
			if err = d.Set("ring_tx", vv); err != nil {
				return fmt.Errorf("Error reading ring_tx: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ring_tx: %v", err)
		}
	}

	if err = d.Set("role", flattenObjectFspVlanInterfaceRole2edl(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "ObjectFspVlanInterface-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("sample_direction", flattenObjectFspVlanInterfaceSampleDirection2edl(o["sample-direction"], d, "sample_direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["sample-direction"], "ObjectFspVlanInterface-SampleDirection"); ok {
			if err = d.Set("sample_direction", vv); err != nil {
				return fmt.Errorf("Error reading sample_direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sample_direction: %v", err)
		}
	}

	if err = d.Set("sample_rate", flattenObjectFspVlanInterfaceSampleRate2edl(o["sample-rate"], d, "sample_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["sample-rate"], "ObjectFspVlanInterface-SampleRate"); ok {
			if err = d.Set("sample_rate", vv); err != nil {
				return fmt.Errorf("Error reading sample_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sample_rate: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenObjectFspVlanInterfaceScanBotnetConnections2edl(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-botnet-connections"], "ObjectFspVlanInterface-ScanBotnetConnections"); ok {
			if err = d.Set("scan_botnet_connections", vv); err != nil {
				return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("secondary_ip", flattenObjectFspVlanInterfaceSecondaryIp2edl(o["secondary-IP"], d, "secondary_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-IP"], "ObjectFspVlanInterface-SecondaryIp"); ok {
			if err = d.Set("secondary_ip", vv); err != nil {
				return fmt.Errorf("Error reading secondary_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("secondaryip", flattenObjectFspVlanInterfaceSecondaryip2edl(o["secondaryip"], d, "secondaryip")); err != nil {
			if vv, ok := fortiAPIPatch(o["secondaryip"], "ObjectFspVlanInterface-Secondaryip"); ok {
				if err = d.Set("secondaryip", vv); err != nil {
					return fmt.Errorf("Error reading secondaryip: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading secondaryip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("secondaryip"); ok {
			if err = d.Set("secondaryip", flattenObjectFspVlanInterfaceSecondaryip2edl(o["secondaryip"], d, "secondaryip")); err != nil {
				if vv, ok := fortiAPIPatch(o["secondaryip"], "ObjectFspVlanInterface-Secondaryip"); ok {
					if err = d.Set("secondaryip", vv); err != nil {
						return fmt.Errorf("Error reading secondaryip: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading secondaryip: %v", err)
				}
			}
		}
	}

	if err = d.Set("security_8021x_dynamic_vlan_id", flattenObjectFspVlanInterfaceSecurity8021XDynamicVlanId2edl(o["security-8021x-dynamic-vlan-id"], d, "security_8021x_dynamic_vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-8021x-dynamic-vlan-id"], "ObjectFspVlanInterface-Security8021XDynamicVlanId"); ok {
			if err = d.Set("security_8021x_dynamic_vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading security_8021x_dynamic_vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_8021x_dynamic_vlan_id: %v", err)
		}
	}

	if err = d.Set("security_8021x_master", flattenObjectFspVlanInterfaceSecurity8021XMaster2edl(o["security-8021x-master"], d, "security_8021x_master")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-8021x-master"], "ObjectFspVlanInterface-Security8021XMaster"); ok {
			if err = d.Set("security_8021x_master", vv); err != nil {
				return fmt.Errorf("Error reading security_8021x_master: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_8021x_master: %v", err)
		}
	}

	if err = d.Set("security_8021x_mode", flattenObjectFspVlanInterfaceSecurity8021XMode2edl(o["security-8021x-mode"], d, "security_8021x_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-8021x-mode"], "ObjectFspVlanInterface-Security8021XMode"); ok {
			if err = d.Set("security_8021x_mode", vv); err != nil {
				return fmt.Errorf("Error reading security_8021x_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_8021x_mode: %v", err)
		}
	}

	if err = d.Set("security_exempt_list", flattenObjectFspVlanInterfaceSecurityExemptList2edl(o["security-exempt-list"], d, "security_exempt_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-exempt-list"], "ObjectFspVlanInterface-SecurityExemptList"); ok {
			if err = d.Set("security_exempt_list", vv); err != nil {
				return fmt.Errorf("Error reading security_exempt_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_exempt_list: %v", err)
		}
	}

	if err = d.Set("security_external_logout", flattenObjectFspVlanInterfaceSecurityExternalLogout2edl(o["security-external-logout"], d, "security_external_logout")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-external-logout"], "ObjectFspVlanInterface-SecurityExternalLogout"); ok {
			if err = d.Set("security_external_logout", vv); err != nil {
				return fmt.Errorf("Error reading security_external_logout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_external_logout: %v", err)
		}
	}

	if err = d.Set("security_external_web", flattenObjectFspVlanInterfaceSecurityExternalWeb2edl(o["security-external-web"], d, "security_external_web")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-external-web"], "ObjectFspVlanInterface-SecurityExternalWeb"); ok {
			if err = d.Set("security_external_web", vv); err != nil {
				return fmt.Errorf("Error reading security_external_web: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_external_web: %v", err)
		}
	}

	if err = d.Set("security_groups", flattenObjectFspVlanInterfaceSecurityGroups2edl(o["security-groups"], d, "security_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-groups"], "ObjectFspVlanInterface-SecurityGroups"); ok {
			if err = d.Set("security_groups", vv); err != nil {
				return fmt.Errorf("Error reading security_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_groups: %v", err)
		}
	}

	if err = d.Set("security_mac_auth_bypass", flattenObjectFspVlanInterfaceSecurityMacAuthBypass2edl(o["security-mac-auth-bypass"], d, "security_mac_auth_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-mac-auth-bypass"], "ObjectFspVlanInterface-SecurityMacAuthBypass"); ok {
			if err = d.Set("security_mac_auth_bypass", vv); err != nil {
				return fmt.Errorf("Error reading security_mac_auth_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenObjectFspVlanInterfaceSecurityMode2edl(o["security-mode"], d, "security_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-mode"], "ObjectFspVlanInterface-SecurityMode"); ok {
			if err = d.Set("security_mode", vv); err != nil {
				return fmt.Errorf("Error reading security_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("security_redirect_url", flattenObjectFspVlanInterfaceSecurityRedirectUrl2edl(o["security-redirect-url"], d, "security_redirect_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-redirect-url"], "ObjectFspVlanInterface-SecurityRedirectUrl"); ok {
			if err = d.Set("security_redirect_url", vv); err != nil {
				return fmt.Errorf("Error reading security_redirect_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_redirect_url: %v", err)
		}
	}

	if err = d.Set("select_profile_30a_35b", flattenObjectFspVlanInterfaceSelectProfile30A35B2edl(o["select-profile-30a-35b"], d, "select_profile_30a_35b")); err != nil {
		if vv, ok := fortiAPIPatch(o["select-profile-30a-35b"], "ObjectFspVlanInterface-SelectProfile30A35B"); ok {
			if err = d.Set("select_profile_30a_35b", vv); err != nil {
				return fmt.Errorf("Error reading select_profile_30a_35b: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading select_profile_30a_35b: %v", err)
		}
	}

	if err = d.Set("service_name", flattenObjectFspVlanInterfaceServiceName2edl(o["service-name"], d, "service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-name"], "ObjectFspVlanInterface-ServiceName"); ok {
			if err = d.Set("service_name", vv); err != nil {
				return fmt.Errorf("Error reading service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_name: %v", err)
		}
	}

	if err = d.Set("sflow_sampler", flattenObjectFspVlanInterfaceSflowSampler2edl(o["sflow-sampler"], d, "sflow_sampler")); err != nil {
		if vv, ok := fortiAPIPatch(o["sflow-sampler"], "ObjectFspVlanInterface-SflowSampler"); ok {
			if err = d.Set("sflow_sampler", vv); err != nil {
				return fmt.Errorf("Error reading sflow_sampler: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sflow_sampler: %v", err)
		}
	}

	if err = d.Set("sfp_dsl", flattenObjectFspVlanInterfaceSfpDsl2edl(o["sfp-dsl"], d, "sfp_dsl")); err != nil {
		if vv, ok := fortiAPIPatch(o["sfp-dsl"], "ObjectFspVlanInterface-SfpDsl"); ok {
			if err = d.Set("sfp_dsl", vv); err != nil {
				return fmt.Errorf("Error reading sfp_dsl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sfp_dsl: %v", err)
		}
	}

	if err = d.Set("sfp_dsl_adsl_fallback", flattenObjectFspVlanInterfaceSfpDslAdslFallback2edl(o["sfp-dsl-adsl-fallback"], d, "sfp_dsl_adsl_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["sfp-dsl-adsl-fallback"], "ObjectFspVlanInterface-SfpDslAdslFallback"); ok {
			if err = d.Set("sfp_dsl_adsl_fallback", vv); err != nil {
				return fmt.Errorf("Error reading sfp_dsl_adsl_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sfp_dsl_adsl_fallback: %v", err)
		}
	}

	if err = d.Set("sfp_dsl_autodetect", flattenObjectFspVlanInterfaceSfpDslAutodetect2edl(o["sfp-dsl-autodetect"], d, "sfp_dsl_autodetect")); err != nil {
		if vv, ok := fortiAPIPatch(o["sfp-dsl-autodetect"], "ObjectFspVlanInterface-SfpDslAutodetect"); ok {
			if err = d.Set("sfp_dsl_autodetect", vv); err != nil {
				return fmt.Errorf("Error reading sfp_dsl_autodetect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sfp_dsl_autodetect: %v", err)
		}
	}

	if err = d.Set("sfp_dsl_mac", flattenObjectFspVlanInterfaceSfpDslMac2edl(o["sfp-dsl-mac"], d, "sfp_dsl_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["sfp-dsl-mac"], "ObjectFspVlanInterface-SfpDslMac"); ok {
			if err = d.Set("sfp_dsl_mac", vv); err != nil {
				return fmt.Errorf("Error reading sfp_dsl_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sfp_dsl_mac: %v", err)
		}
	}

	if err = d.Set("speed", flattenObjectFspVlanInterfaceSpeed2edl(o["speed"], d, "speed")); err != nil {
		if vv, ok := fortiAPIPatch(o["speed"], "ObjectFspVlanInterface-Speed"); ok {
			if err = d.Set("speed", vv); err != nil {
				return fmt.Errorf("Error reading speed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading speed: %v", err)
		}
	}

	if err = d.Set("spillover_threshold", flattenObjectFspVlanInterfaceSpilloverThreshold2edl(o["spillover-threshold"], d, "spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spillover-threshold"], "ObjectFspVlanInterface-SpilloverThreshold"); ok {
			if err = d.Set("spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spillover_threshold: %v", err)
		}
	}

	if err = d.Set("src_check", flattenObjectFspVlanInterfaceSrcCheck2edl(o["src-check"], d, "src_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-check"], "ObjectFspVlanInterface-SrcCheck"); ok {
			if err = d.Set("src_check", vv); err != nil {
				return fmt.Errorf("Error reading src_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_check: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFspVlanInterfaceStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFspVlanInterface-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("stp", flattenObjectFspVlanInterfaceStp2edl(o["stp"], d, "stp")); err != nil {
		if vv, ok := fortiAPIPatch(o["stp"], "ObjectFspVlanInterface-Stp"); ok {
			if err = d.Set("stp", vv); err != nil {
				return fmt.Errorf("Error reading stp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stp: %v", err)
		}
	}

	if err = d.Set("stp_ha_secondary", flattenObjectFspVlanInterfaceStpHaSecondary2edl(o["stp-ha-secondary"], d, "stp_ha_secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["stp-ha-secondary"], "ObjectFspVlanInterface-StpHaSecondary"); ok {
			if err = d.Set("stp_ha_secondary", vv); err != nil {
				return fmt.Errorf("Error reading stp_ha_secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stp_ha_secondary: %v", err)
		}
	}

	if err = d.Set("stp_ha_slave", flattenObjectFspVlanInterfaceStpHaSlave2edl(o["stp-ha-slave"], d, "stp_ha_slave")); err != nil {
		if vv, ok := fortiAPIPatch(o["stp-ha-slave"], "ObjectFspVlanInterface-StpHaSlave"); ok {
			if err = d.Set("stp_ha_slave", vv); err != nil {
				return fmt.Errorf("Error reading stp_ha_slave: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stp_ha_slave: %v", err)
		}
	}

	if err = d.Set("stpforward", flattenObjectFspVlanInterfaceStpforward2edl(o["stpforward"], d, "stpforward")); err != nil {
		if vv, ok := fortiAPIPatch(o["stpforward"], "ObjectFspVlanInterface-Stpforward"); ok {
			if err = d.Set("stpforward", vv); err != nil {
				return fmt.Errorf("Error reading stpforward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stpforward: %v", err)
		}
	}

	if err = d.Set("stpforward_mode", flattenObjectFspVlanInterfaceStpforwardMode2edl(o["stpforward-mode"], d, "stpforward_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["stpforward-mode"], "ObjectFspVlanInterface-StpforwardMode"); ok {
			if err = d.Set("stpforward_mode", vv); err != nil {
				return fmt.Errorf("Error reading stpforward_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stpforward_mode: %v", err)
		}
	}

	if err = d.Set("strip_priority_vlan_tag", flattenObjectFspVlanInterfaceStripPriorityVlanTag2edl(o["strip-priority-vlan-tag"], d, "strip_priority_vlan_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["strip-priority-vlan-tag"], "ObjectFspVlanInterface-StripPriorityVlanTag"); ok {
			if err = d.Set("strip_priority_vlan_tag", vv); err != nil {
				return fmt.Errorf("Error reading strip_priority_vlan_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strip_priority_vlan_tag: %v", err)
		}
	}

	if err = d.Set("subst", flattenObjectFspVlanInterfaceSubst2edl(o["subst"], d, "subst")); err != nil {
		if vv, ok := fortiAPIPatch(o["subst"], "ObjectFspVlanInterface-Subst"); ok {
			if err = d.Set("subst", vv); err != nil {
				return fmt.Errorf("Error reading subst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subst: %v", err)
		}
	}

	if err = d.Set("substitute_dst_mac", flattenObjectFspVlanInterfaceSubstituteDstMac2edl(o["substitute-dst-mac"], d, "substitute_dst_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["substitute-dst-mac"], "ObjectFspVlanInterface-SubstituteDstMac"); ok {
			if err = d.Set("substitute_dst_mac", vv); err != nil {
				return fmt.Errorf("Error reading substitute_dst_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading substitute_dst_mac: %v", err)
		}
	}

	if err = d.Set("sw_algorithm", flattenObjectFspVlanInterfaceSwAlgorithm2edl(o["sw-algorithm"], d, "sw_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["sw-algorithm"], "ObjectFspVlanInterface-SwAlgorithm"); ok {
			if err = d.Set("sw_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading sw_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sw_algorithm: %v", err)
		}
	}

	if err = d.Set("swc_first_create", flattenObjectFspVlanInterfaceSwcFirstCreate2edl(o["swc-first-create"], d, "swc_first_create")); err != nil {
		if vv, ok := fortiAPIPatch(o["swc-first-create"], "ObjectFspVlanInterface-SwcFirstCreate"); ok {
			if err = d.Set("swc_first_create", vv); err != nil {
				return fmt.Errorf("Error reading swc_first_create: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading swc_first_create: %v", err)
		}
	}

	if err = d.Set("swc_vlan", flattenObjectFspVlanInterfaceSwcVlan2edl(o["swc-vlan"], d, "swc_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["swc-vlan"], "ObjectFspVlanInterface-SwcVlan"); ok {
			if err = d.Set("swc_vlan", vv); err != nil {
				return fmt.Errorf("Error reading swc_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading swc_vlan: %v", err)
		}
	}

	if err = d.Set("switch", flattenObjectFspVlanInterfaceSwitch2edl(o["switch"], d, "switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch"], "ObjectFspVlanInterface-Switch"); ok {
			if err = d.Set("switch", vv); err != nil {
				return fmt.Errorf("Error reading switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch: %v", err)
		}
	}

	if err = d.Set("switch_controller_access_vlan", flattenObjectFspVlanInterfaceSwitchControllerAccessVlan2edl(o["switch-controller-access-vlan"], d, "switch_controller_access_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-access-vlan"], "ObjectFspVlanInterface-SwitchControllerAccessVlan"); ok {
			if err = d.Set("switch_controller_access_vlan", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_access_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_access_vlan: %v", err)
		}
	}

	if err = d.Set("switch_controller_arp_inspection", flattenObjectFspVlanInterfaceSwitchControllerArpInspection2edl(o["switch-controller-arp-inspection"], d, "switch_controller_arp_inspection")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-arp-inspection"], "ObjectFspVlanInterface-SwitchControllerArpInspection"); ok {
			if err = d.Set("switch_controller_arp_inspection", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_arp_inspection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_arp_inspection: %v", err)
		}
	}

	if err = d.Set("switch_controller_auth", flattenObjectFspVlanInterfaceSwitchControllerAuth2edl(o["switch-controller-auth"], d, "switch_controller_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-auth"], "ObjectFspVlanInterface-SwitchControllerAuth"); ok {
			if err = d.Set("switch_controller_auth", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_auth: %v", err)
		}
	}

	if err = d.Set("switch_controller_dhcp_snooping", flattenObjectFspVlanInterfaceSwitchControllerDhcpSnooping2edl(o["switch-controller-dhcp-snooping"], d, "switch_controller_dhcp_snooping")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-dhcp-snooping"], "ObjectFspVlanInterface-SwitchControllerDhcpSnooping"); ok {
			if err = d.Set("switch_controller_dhcp_snooping", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_dhcp_snooping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_dhcp_snooping: %v", err)
		}
	}

	if err = d.Set("switch_controller_dhcp_snooping_option82", flattenObjectFspVlanInterfaceSwitchControllerDhcpSnoopingOption822edl(o["switch-controller-dhcp-snooping-option82"], d, "switch_controller_dhcp_snooping_option82")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-dhcp-snooping-option82"], "ObjectFspVlanInterface-SwitchControllerDhcpSnoopingOption82"); ok {
			if err = d.Set("switch_controller_dhcp_snooping_option82", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_dhcp_snooping_option82: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_dhcp_snooping_option82: %v", err)
		}
	}

	if err = d.Set("switch_controller_dhcp_snooping_verify_mac", flattenObjectFspVlanInterfaceSwitchControllerDhcpSnoopingVerifyMac2edl(o["switch-controller-dhcp-snooping-verify-mac"], d, "switch_controller_dhcp_snooping_verify_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-dhcp-snooping-verify-mac"], "ObjectFspVlanInterface-SwitchControllerDhcpSnoopingVerifyMac"); ok {
			if err = d.Set("switch_controller_dhcp_snooping_verify_mac", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_dhcp_snooping_verify_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_dhcp_snooping_verify_mac: %v", err)
		}
	}

	if err = d.Set("switch_controller_dynamic", flattenObjectFspVlanInterfaceSwitchControllerDynamic2edl(o["switch-controller-dynamic"], d, "switch_controller_dynamic")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-dynamic"], "ObjectFspVlanInterface-SwitchControllerDynamic"); ok {
			if err = d.Set("switch_controller_dynamic", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_dynamic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_dynamic: %v", err)
		}
	}

	if err = d.Set("switch_controller_feature", flattenObjectFspVlanInterfaceSwitchControllerFeature2edl(o["switch-controller-feature"], d, "switch_controller_feature")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-feature"], "ObjectFspVlanInterface-SwitchControllerFeature"); ok {
			if err = d.Set("switch_controller_feature", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_feature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_feature: %v", err)
		}
	}

	if err = d.Set("switch_controller_igmp_snooping", flattenObjectFspVlanInterfaceSwitchControllerIgmpSnooping2edl(o["switch-controller-igmp-snooping"], d, "switch_controller_igmp_snooping")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-igmp-snooping"], "ObjectFspVlanInterface-SwitchControllerIgmpSnooping"); ok {
			if err = d.Set("switch_controller_igmp_snooping", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_igmp_snooping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_igmp_snooping: %v", err)
		}
	}

	if err = d.Set("switch_controller_igmp_snooping_fast_leave", flattenObjectFspVlanInterfaceSwitchControllerIgmpSnoopingFastLeave2edl(o["switch-controller-igmp-snooping-fast-leave"], d, "switch_controller_igmp_snooping_fast_leave")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-igmp-snooping-fast-leave"], "ObjectFspVlanInterface-SwitchControllerIgmpSnoopingFastLeave"); ok {
			if err = d.Set("switch_controller_igmp_snooping_fast_leave", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_igmp_snooping_fast_leave: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_igmp_snooping_fast_leave: %v", err)
		}
	}

	if err = d.Set("switch_controller_igmp_snooping_proxy", flattenObjectFspVlanInterfaceSwitchControllerIgmpSnoopingProxy2edl(o["switch-controller-igmp-snooping-proxy"], d, "switch_controller_igmp_snooping_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-igmp-snooping-proxy"], "ObjectFspVlanInterface-SwitchControllerIgmpSnoopingProxy"); ok {
			if err = d.Set("switch_controller_igmp_snooping_proxy", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_igmp_snooping_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_igmp_snooping_proxy: %v", err)
		}
	}

	if err = d.Set("switch_controller_iot_scanning", flattenObjectFspVlanInterfaceSwitchControllerIotScanning2edl(o["switch-controller-iot-scanning"], d, "switch_controller_iot_scanning")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-iot-scanning"], "ObjectFspVlanInterface-SwitchControllerIotScanning"); ok {
			if err = d.Set("switch_controller_iot_scanning", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_iot_scanning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_iot_scanning: %v", err)
		}
	}

	if err = d.Set("switch_controller_learning_limit", flattenObjectFspVlanInterfaceSwitchControllerLearningLimit2edl(o["switch-controller-learning-limit"], d, "switch_controller_learning_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-learning-limit"], "ObjectFspVlanInterface-SwitchControllerLearningLimit"); ok {
			if err = d.Set("switch_controller_learning_limit", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_learning_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_learning_limit: %v", err)
		}
	}

	if err = d.Set("switch_controller_mgmt_vlan", flattenObjectFspVlanInterfaceSwitchControllerMgmtVlan2edl(o["switch-controller-mgmt-vlan"], d, "switch_controller_mgmt_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-mgmt-vlan"], "ObjectFspVlanInterface-SwitchControllerMgmtVlan"); ok {
			if err = d.Set("switch_controller_mgmt_vlan", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_mgmt_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_mgmt_vlan: %v", err)
		}
	}

	if err = d.Set("switch_controller_nac", flattenObjectFspVlanInterfaceSwitchControllerNac2edl(o["switch-controller-nac"], d, "switch_controller_nac")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-nac"], "ObjectFspVlanInterface-SwitchControllerNac"); ok {
			if err = d.Set("switch_controller_nac", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_nac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_nac: %v", err)
		}
	}

	if err = d.Set("switch_controller_netflow_collect", flattenObjectFspVlanInterfaceSwitchControllerNetflowCollect2edl(o["switch-controller-netflow-collect"], d, "switch_controller_netflow_collect")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-netflow-collect"], "ObjectFspVlanInterface-SwitchControllerNetflowCollect"); ok {
			if err = d.Set("switch_controller_netflow_collect", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_netflow_collect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_netflow_collect: %v", err)
		}
	}

	if err = d.Set("switch_controller_offload", flattenObjectFspVlanInterfaceSwitchControllerOffload2edl(o["switch-controller-offload"], d, "switch_controller_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-offload"], "ObjectFspVlanInterface-SwitchControllerOffload"); ok {
			if err = d.Set("switch_controller_offload", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_offload: %v", err)
		}
	}

	if err = d.Set("switch_controller_offload_gw", flattenObjectFspVlanInterfaceSwitchControllerOffloadGw2edl(o["switch-controller-offload-gw"], d, "switch_controller_offload_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-offload-gw"], "ObjectFspVlanInterface-SwitchControllerOffloadGw"); ok {
			if err = d.Set("switch_controller_offload_gw", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_offload_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_offload_gw: %v", err)
		}
	}

	if err = d.Set("switch_controller_offload_ip", flattenObjectFspVlanInterfaceSwitchControllerOffloadIp2edl(o["switch-controller-offload-ip"], d, "switch_controller_offload_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-offload-ip"], "ObjectFspVlanInterface-SwitchControllerOffloadIp"); ok {
			if err = d.Set("switch_controller_offload_ip", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_offload_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_offload_ip: %v", err)
		}
	}

	if err = d.Set("switch_controller_offloading", flattenObjectFspVlanInterfaceSwitchControllerOffloading2edl(o["switch-controller-offloading"], d, "switch_controller_offloading")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-offloading"], "ObjectFspVlanInterface-SwitchControllerOffloading"); ok {
			if err = d.Set("switch_controller_offloading", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_offloading: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_offloading: %v", err)
		}
	}

	if err = d.Set("switch_controller_offloading_gw", flattenObjectFspVlanInterfaceSwitchControllerOffloadingGw2edl(o["switch-controller-offloading-gw"], d, "switch_controller_offloading_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-offloading-gw"], "ObjectFspVlanInterface-SwitchControllerOffloadingGw"); ok {
			if err = d.Set("switch_controller_offloading_gw", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_offloading_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_offloading_gw: %v", err)
		}
	}

	if err = d.Set("switch_controller_offloading_ip", flattenObjectFspVlanInterfaceSwitchControllerOffloadingIp2edl(o["switch-controller-offloading-ip"], d, "switch_controller_offloading_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-offloading-ip"], "ObjectFspVlanInterface-SwitchControllerOffloadingIp"); ok {
			if err = d.Set("switch_controller_offloading_ip", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_offloading_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_offloading_ip: %v", err)
		}
	}

	if err = d.Set("switch_controller_radius_server", flattenObjectFspVlanInterfaceSwitchControllerRadiusServer2edl(o["switch-controller-radius-server"], d, "switch_controller_radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-radius-server"], "ObjectFspVlanInterface-SwitchControllerRadiusServer"); ok {
			if err = d.Set("switch_controller_radius_server", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_radius_server: %v", err)
		}
	}

	if err = d.Set("switch_controller_rspan_mode", flattenObjectFspVlanInterfaceSwitchControllerRspanMode2edl(o["switch-controller-rspan-mode"], d, "switch_controller_rspan_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-rspan-mode"], "ObjectFspVlanInterface-SwitchControllerRspanMode"); ok {
			if err = d.Set("switch_controller_rspan_mode", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_rspan_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_rspan_mode: %v", err)
		}
	}

	if err = d.Set("switch_controller_source_ip", flattenObjectFspVlanInterfaceSwitchControllerSourceIp2edl(o["switch-controller-source-ip"], d, "switch_controller_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-source-ip"], "ObjectFspVlanInterface-SwitchControllerSourceIp"); ok {
			if err = d.Set("switch_controller_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_source_ip: %v", err)
		}
	}

	if err = d.Set("switch_controller_traffic_policy", flattenObjectFspVlanInterfaceSwitchControllerTrafficPolicy2edl(o["switch-controller-traffic-policy"], d, "switch_controller_traffic_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-traffic-policy"], "ObjectFspVlanInterface-SwitchControllerTrafficPolicy"); ok {
			if err = d.Set("switch_controller_traffic_policy", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_traffic_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_traffic_policy: %v", err)
		}
	}

	if err = d.Set("system_id", flattenObjectFspVlanInterfaceSystemId2edl(o["system-id"], d, "system_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-id"], "ObjectFspVlanInterface-SystemId"); ok {
			if err = d.Set("system_id", vv); err != nil {
				return fmt.Errorf("Error reading system_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_id: %v", err)
		}
	}

	if err = d.Set("system_id_type", flattenObjectFspVlanInterfaceSystemIdType2edl(o["system-id-type"], d, "system_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-id-type"], "ObjectFspVlanInterface-SystemIdType"); ok {
			if err = d.Set("system_id_type", vv); err != nil {
				return fmt.Errorf("Error reading system_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_id_type: %v", err)
		}
	}

	if err = d.Set("tc_mode", flattenObjectFspVlanInterfaceTcMode2edl(o["tc-mode"], d, "tc_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["tc-mode"], "ObjectFspVlanInterface-TcMode"); ok {
			if err = d.Set("tc_mode", vv); err != nil {
				return fmt.Errorf("Error reading tc_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tc_mode: %v", err)
		}
	}

	if err = d.Set("tcp_mss", flattenObjectFspVlanInterfaceTcpMss2edl(o["tcp-mss"], d, "tcp_mss")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-mss"], "ObjectFspVlanInterface-TcpMss"); ok {
			if err = d.Set("tcp_mss", vv); err != nil {
				return fmt.Errorf("Error reading tcp_mss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_mss: %v", err)
		}
	}

	if err = d.Set("trunk", flattenObjectFspVlanInterfaceTrunk2edl(o["trunk"], d, "trunk")); err != nil {
		if vv, ok := fortiAPIPatch(o["trunk"], "ObjectFspVlanInterface-Trunk"); ok {
			if err = d.Set("trunk", vv); err != nil {
				return fmt.Errorf("Error reading trunk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trunk: %v", err)
		}
	}

	if err = d.Set("trust_ip_1", flattenObjectFspVlanInterfaceTrustIp12edl(o["trust-ip-1"], d, "trust_ip_1")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-ip-1"], "ObjectFspVlanInterface-TrustIp1"); ok {
			if err = d.Set("trust_ip_1", vv); err != nil {
				return fmt.Errorf("Error reading trust_ip_1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_ip_1: %v", err)
		}
	}

	if err = d.Set("trust_ip_2", flattenObjectFspVlanInterfaceTrustIp22edl(o["trust-ip-2"], d, "trust_ip_2")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-ip-2"], "ObjectFspVlanInterface-TrustIp2"); ok {
			if err = d.Set("trust_ip_2", vv); err != nil {
				return fmt.Errorf("Error reading trust_ip_2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_ip_2: %v", err)
		}
	}

	if err = d.Set("trust_ip_3", flattenObjectFspVlanInterfaceTrustIp32edl(o["trust-ip-3"], d, "trust_ip_3")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-ip-3"], "ObjectFspVlanInterface-TrustIp3"); ok {
			if err = d.Set("trust_ip_3", vv); err != nil {
				return fmt.Errorf("Error reading trust_ip_3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_ip_3: %v", err)
		}
	}

	if err = d.Set("trust_ip6_1", flattenObjectFspVlanInterfaceTrustIp612edl(o["trust-ip6-1"], d, "trust_ip6_1")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-ip6-1"], "ObjectFspVlanInterface-TrustIp61"); ok {
			if err = d.Set("trust_ip6_1", vv); err != nil {
				return fmt.Errorf("Error reading trust_ip6_1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_ip6_1: %v", err)
		}
	}

	if err = d.Set("trust_ip6_2", flattenObjectFspVlanInterfaceTrustIp622edl(o["trust-ip6-2"], d, "trust_ip6_2")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-ip6-2"], "ObjectFspVlanInterface-TrustIp62"); ok {
			if err = d.Set("trust_ip6_2", vv); err != nil {
				return fmt.Errorf("Error reading trust_ip6_2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_ip6_2: %v", err)
		}
	}

	if err = d.Set("trust_ip6_3", flattenObjectFspVlanInterfaceTrustIp632edl(o["trust-ip6-3"], d, "trust_ip6_3")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-ip6-3"], "ObjectFspVlanInterface-TrustIp63"); ok {
			if err = d.Set("trust_ip6_3", vv); err != nil {
				return fmt.Errorf("Error reading trust_ip6_3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_ip6_3: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFspVlanInterfaceType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFspVlanInterface-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username", flattenObjectFspVlanInterfaceUsername2edl(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ObjectFspVlanInterface-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("vci", flattenObjectFspVlanInterfaceVci2edl(o["vci"], d, "vci")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci"], "ObjectFspVlanInterface-Vci"); ok {
			if err = d.Set("vci", vv); err != nil {
				return fmt.Errorf("Error reading vci: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci: %v", err)
		}
	}

	if err = d.Set("vectoring", flattenObjectFspVlanInterfaceVectoring2edl(o["vectoring"], d, "vectoring")); err != nil {
		if vv, ok := fortiAPIPatch(o["vectoring"], "ObjectFspVlanInterface-Vectoring"); ok {
			if err = d.Set("vectoring", vv); err != nil {
				return fmt.Errorf("Error reading vectoring: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vectoring: %v", err)
		}
	}

	if err = d.Set("vindex", flattenObjectFspVlanInterfaceVindex2edl(o["vindex"], d, "vindex")); err != nil {
		if vv, ok := fortiAPIPatch(o["vindex"], "ObjectFspVlanInterface-Vindex"); ok {
			if err = d.Set("vindex", vv); err != nil {
				return fmt.Errorf("Error reading vindex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vindex: %v", err)
		}
	}

	if err = d.Set("vlan_protocol", flattenObjectFspVlanInterfaceVlanProtocol2edl(o["vlan-protocol"], d, "vlan_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-protocol"], "ObjectFspVlanInterface-VlanProtocol"); ok {
			if err = d.Set("vlan_protocol", vv); err != nil {
				return fmt.Errorf("Error reading vlan_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_protocol: %v", err)
		}
	}

	if err = d.Set("vlanforward", flattenObjectFspVlanInterfaceVlanforward2edl(o["vlanforward"], d, "vlanforward")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlanforward"], "ObjectFspVlanInterface-Vlanforward"); ok {
			if err = d.Set("vlanforward", vv); err != nil {
				return fmt.Errorf("Error reading vlanforward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlanforward: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenObjectFspVlanInterfaceVlanid2edl(o["vlanid"], d, "vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlanid"], "ObjectFspVlanInterface-Vlanid"); ok {
			if err = d.Set("vlanid", vv); err != nil {
				return fmt.Errorf("Error reading vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("vpi", flattenObjectFspVlanInterfaceVpi2edl(o["vpi"], d, "vpi")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpi"], "ObjectFspVlanInterface-Vpi"); ok {
			if err = d.Set("vpi", vv); err != nil {
				return fmt.Errorf("Error reading vpi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpi: %v", err)
		}
	}

	if err = d.Set("vrf", flattenObjectFspVlanInterfaceVrf2edl(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "ObjectFspVlanInterface-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vrrp", flattenObjectFspVlanInterfaceVrrp2edl(o["vrrp"], d, "vrrp")); err != nil {
			if vv, ok := fortiAPIPatch(o["vrrp"], "ObjectFspVlanInterface-Vrrp"); ok {
				if err = d.Set("vrrp", vv); err != nil {
					return fmt.Errorf("Error reading vrrp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vrrp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrrp"); ok {
			if err = d.Set("vrrp", flattenObjectFspVlanInterfaceVrrp2edl(o["vrrp"], d, "vrrp")); err != nil {
				if vv, ok := fortiAPIPatch(o["vrrp"], "ObjectFspVlanInterface-Vrrp"); ok {
					if err = d.Set("vrrp", vv); err != nil {
						return fmt.Errorf("Error reading vrrp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vrrp: %v", err)
				}
			}
		}
	}

	if err = d.Set("vrrp_virtual_mac", flattenObjectFspVlanInterfaceVrrpVirtualMac2edl(o["vrrp-virtual-mac"], d, "vrrp_virtual_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrrp-virtual-mac"], "ObjectFspVlanInterface-VrrpVirtualMac"); ok {
			if err = d.Set("vrrp_virtual_mac", vv); err != nil {
				return fmt.Errorf("Error reading vrrp_virtual_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrrp_virtual_mac: %v", err)
		}
	}

	if err = d.Set("wccp", flattenObjectFspVlanInterfaceWccp2edl(o["wccp"], d, "wccp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wccp"], "ObjectFspVlanInterface-Wccp"); ok {
			if err = d.Set("wccp", vv); err != nil {
				return fmt.Errorf("Error reading wccp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectFspVlanInterfaceWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectFspVlanInterface-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("wifi_5g_threshold", flattenObjectFspVlanInterfaceWifi5GThreshold2edl(o["wifi-5g-threshold"], d, "wifi_5g_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-5g-threshold"], "ObjectFspVlanInterface-Wifi5GThreshold"); ok {
			if err = d.Set("wifi_5g_threshold", vv); err != nil {
				return fmt.Errorf("Error reading wifi_5g_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_5g_threshold: %v", err)
		}
	}

	if err = d.Set("wifi_acl", flattenObjectFspVlanInterfaceWifiAcl2edl(o["wifi-acl"], d, "wifi_acl")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-acl"], "ObjectFspVlanInterface-WifiAcl"); ok {
			if err = d.Set("wifi_acl", vv); err != nil {
				return fmt.Errorf("Error reading wifi_acl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_acl: %v", err)
		}
	}

	if err = d.Set("wifi_ap_band", flattenObjectFspVlanInterfaceWifiApBand2edl(o["wifi-ap-band"], d, "wifi_ap_band")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ap-band"], "ObjectFspVlanInterface-WifiApBand"); ok {
			if err = d.Set("wifi_ap_band", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ap_band: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ap_band: %v", err)
		}
	}

	if err = d.Set("wifi_auth", flattenObjectFspVlanInterfaceWifiAuth2edl(o["wifi-auth"], d, "wifi_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-auth"], "ObjectFspVlanInterface-WifiAuth"); ok {
			if err = d.Set("wifi_auth", vv); err != nil {
				return fmt.Errorf("Error reading wifi_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_auth: %v", err)
		}
	}

	if err = d.Set("wifi_auto_connect", flattenObjectFspVlanInterfaceWifiAutoConnect2edl(o["wifi-auto-connect"], d, "wifi_auto_connect")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-auto-connect"], "ObjectFspVlanInterface-WifiAutoConnect"); ok {
			if err = d.Set("wifi_auto_connect", vv); err != nil {
				return fmt.Errorf("Error reading wifi_auto_connect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_auto_connect: %v", err)
		}
	}

	if err = d.Set("wifi_auto_save", flattenObjectFspVlanInterfaceWifiAutoSave2edl(o["wifi-auto-save"], d, "wifi_auto_save")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-auto-save"], "ObjectFspVlanInterface-WifiAutoSave"); ok {
			if err = d.Set("wifi_auto_save", vv); err != nil {
				return fmt.Errorf("Error reading wifi_auto_save: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_auto_save: %v", err)
		}
	}

	if err = d.Set("wifi_broadcast_ssid", flattenObjectFspVlanInterfaceWifiBroadcastSsid2edl(o["wifi-broadcast-ssid"], d, "wifi_broadcast_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-broadcast-ssid"], "ObjectFspVlanInterface-WifiBroadcastSsid"); ok {
			if err = d.Set("wifi_broadcast_ssid", vv); err != nil {
				return fmt.Errorf("Error reading wifi_broadcast_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_broadcast_ssid: %v", err)
		}
	}

	if err = d.Set("wifi_dns_server1", flattenObjectFspVlanInterfaceWifiDnsServer12edl(o["wifi-dns-server1"], d, "wifi_dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-dns-server1"], "ObjectFspVlanInterface-WifiDnsServer1"); ok {
			if err = d.Set("wifi_dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading wifi_dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_dns_server1: %v", err)
		}
	}

	if err = d.Set("wifi_dns_server2", flattenObjectFspVlanInterfaceWifiDnsServer22edl(o["wifi-dns-server2"], d, "wifi_dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-dns-server2"], "ObjectFspVlanInterface-WifiDnsServer2"); ok {
			if err = d.Set("wifi_dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading wifi_dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_dns_server2: %v", err)
		}
	}

	if err = d.Set("wifi_encrypt", flattenObjectFspVlanInterfaceWifiEncrypt2edl(o["wifi-encrypt"], d, "wifi_encrypt")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-encrypt"], "ObjectFspVlanInterface-WifiEncrypt"); ok {
			if err = d.Set("wifi_encrypt", vv); err != nil {
				return fmt.Errorf("Error reading wifi_encrypt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_encrypt: %v", err)
		}
	}

	if err = d.Set("wifi_fragment_threshold", flattenObjectFspVlanInterfaceWifiFragmentThreshold2edl(o["wifi-fragment-threshold"], d, "wifi_fragment_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-fragment-threshold"], "ObjectFspVlanInterface-WifiFragmentThreshold"); ok {
			if err = d.Set("wifi_fragment_threshold", vv); err != nil {
				return fmt.Errorf("Error reading wifi_fragment_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_fragment_threshold: %v", err)
		}
	}

	if err = d.Set("wifi_gateway", flattenObjectFspVlanInterfaceWifiGateway2edl(o["wifi-gateway"], d, "wifi_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-gateway"], "ObjectFspVlanInterface-WifiGateway"); ok {
			if err = d.Set("wifi_gateway", vv); err != nil {
				return fmt.Errorf("Error reading wifi_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_gateway: %v", err)
		}
	}

	if err = d.Set("wifi_key", flattenObjectFspVlanInterfaceWifiKey2edl(o["wifi-key"], d, "wifi_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-key"], "ObjectFspVlanInterface-WifiKey"); ok {
			if err = d.Set("wifi_key", vv); err != nil {
				return fmt.Errorf("Error reading wifi_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_key: %v", err)
		}
	}

	if err = d.Set("wifi_keyindex", flattenObjectFspVlanInterfaceWifiKeyindex2edl(o["wifi-keyindex"], d, "wifi_keyindex")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-keyindex"], "ObjectFspVlanInterface-WifiKeyindex"); ok {
			if err = d.Set("wifi_keyindex", vv); err != nil {
				return fmt.Errorf("Error reading wifi_keyindex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_keyindex: %v", err)
		}
	}

	if err = d.Set("wifi_mac_filter", flattenObjectFspVlanInterfaceWifiMacFilter2edl(o["wifi-mac-filter"], d, "wifi_mac_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-mac-filter"], "ObjectFspVlanInterface-WifiMacFilter"); ok {
			if err = d.Set("wifi_mac_filter", vv); err != nil {
				return fmt.Errorf("Error reading wifi_mac_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_mac_filter: %v", err)
		}
	}

	if err = d.Set("wifi_radius_server", flattenObjectFspVlanInterfaceWifiRadiusServer2edl(o["wifi-radius-server"], d, "wifi_radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-radius-server"], "ObjectFspVlanInterface-WifiRadiusServer"); ok {
			if err = d.Set("wifi_radius_server", vv); err != nil {
				return fmt.Errorf("Error reading wifi_radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_radius_server: %v", err)
		}
	}

	if err = d.Set("wifi_rts_threshold", flattenObjectFspVlanInterfaceWifiRtsThreshold2edl(o["wifi-rts-threshold"], d, "wifi_rts_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-rts-threshold"], "ObjectFspVlanInterface-WifiRtsThreshold"); ok {
			if err = d.Set("wifi_rts_threshold", vv); err != nil {
				return fmt.Errorf("Error reading wifi_rts_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_rts_threshold: %v", err)
		}
	}

	if err = d.Set("wifi_security", flattenObjectFspVlanInterfaceWifiSecurity2edl(o["wifi-security"], d, "wifi_security")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-security"], "ObjectFspVlanInterface-WifiSecurity"); ok {
			if err = d.Set("wifi_security", vv); err != nil {
				return fmt.Errorf("Error reading wifi_security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_security: %v", err)
		}
	}

	if err = d.Set("wifi_ssid", flattenObjectFspVlanInterfaceWifiSsid2edl(o["wifi-ssid"], d, "wifi_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ssid"], "ObjectFspVlanInterface-WifiSsid"); ok {
			if err = d.Set("wifi_ssid", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ssid: %v", err)
		}
	}

	if err = d.Set("wifi_usergroup", flattenObjectFspVlanInterfaceWifiUsergroup2edl(o["wifi-usergroup"], d, "wifi_usergroup")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-usergroup"], "ObjectFspVlanInterface-WifiUsergroup"); ok {
			if err = d.Set("wifi_usergroup", vv); err != nil {
				return fmt.Errorf("Error reading wifi_usergroup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_usergroup: %v", err)
		}
	}

	if err = d.Set("wins_ip", flattenObjectFspVlanInterfaceWinsIp2edl(o["wins-ip"], d, "wins_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-ip"], "ObjectFspVlanInterface-WinsIp"); ok {
			if err = d.Set("wins_ip", vv); err != nil {
				return fmt.Errorf("Error reading wins_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_ip: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanInterfaceVlanOpMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAcName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAggregate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAggregateType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAlgorithm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAlias2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAllowaccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceApDiscover2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceArpforward2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAtmProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAuthCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAuthPortalAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAuthType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAutoAuthExtensionDevice2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBandwidthMeasureTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBfd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBfdDesiredMinTx2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBfdDetectMult2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBfdRequiredMinRx2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBroadcastForticlientDiscovery2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBroadcastForward2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceCaptivePortal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceCliConnStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceColor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdns2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsKeyname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceDdnsServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsSn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsUsername2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsZone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDedicatedTo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDefaultPurdueLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDefaultgw2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDetectedPeerMtu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDetectprotocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceDetectserver2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDeviceAccessList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDeviceIdentification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDeviceIdentificationActiveScan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDeviceNetscan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDeviceUserIdentification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDevindex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpBroadcastFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpClasslessRouteAddition2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpClientIdentifier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayAgentOption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayCircuitId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayInterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceDhcpRelayLinkSelection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayRequestAllServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelaySourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRenewTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpSmartRelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDiscRetryTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDisconnectThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDistance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDnsQuery2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDnsServerOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDnsServerProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceDropFragment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDropOverlappedFragment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEapCaCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceEapIdentity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEapMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEapPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceEapSupplicant2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEapUserCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceEgressCos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEgressShapingProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEndpointCompliance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEstimatedDownstreamBandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEstimatedUpstreamBandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceExplicitFtpProxy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceExplicitWebProxy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceExternal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFailActionOnExtender2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFailAlertInterfaces2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFailAlertMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFailDetect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFailDetectOption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceFdp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortiheartbeat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortilink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortilinkBackupLink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortilinkNeighborDetect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortilinkSplitInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortilinkStacking2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceForwardDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceForwardErrorCorrection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFpAnomaly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceFpDisable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceGatewayAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceGenericReceiveOffload2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceGiGk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceGwaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceGwdetect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceHaPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIcmpAcceptRedirect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIcmpRedirect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIcmpSendRedirect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIdentAccept2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIdleTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIfMdix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIfMedia2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIkeSamlServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceInForceVlanCos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceInbandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIngressCos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIngressShapingProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIngressSpilloverThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceInterconnectProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceInternal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpManagedByFortiipam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpmac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpsSnifferMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpunnumbered2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "autoconf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["autoconf"], _ = expandObjectFspVlanInterfaceIpv6Autoconf2edl(d, i["autoconf"], pre_append)
	}
	pre_append = pre + ".0." + "cli_conn6_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cli-conn6-status"], _ = expandObjectFspVlanInterfaceIpv6CliConn6Status2edl(d, i["cli_conn6_status"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-client-options"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6ClientOptions2edl(d, i["dhcp6_client_options"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-information-request"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6InformationRequest2edl(d, i["dhcp6_information_request"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-delegation"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation2edl(d, i["dhcp6_prefix_delegation"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHint2edl(d, i["dhcp6_prefix_hint"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint-plt"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt2edl(d, i["dhcp6_prefix_hint_plt"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint-vlt"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt2edl(d, i["dhcp6_prefix_hint_vlt"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_interface_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-interface-id"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6RelayInterfaceId2edl(d, i["dhcp6_relay_interface_id"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-ip"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6RelayIp2edl(d, i["dhcp6_relay_ip"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-service"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6RelayService2edl(d, i["dhcp6_relay_service"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_source_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-source-interface"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6RelaySourceInterface2edl(d, i["dhcp6_relay_source_interface"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_source_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-source-ip"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6RelaySourceIp2edl(d, i["dhcp6_relay_source_ip"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-type"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6RelayType2edl(d, i["dhcp6_relay_type"], pre_append)
	}
	pre_append = pre + ".0." + "icmp6_send_redirect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp6-send-redirect"], _ = expandObjectFspVlanInterfaceIpv6Icmp6SendRedirect2edl(d, i["icmp6_send_redirect"], pre_append)
	}
	pre_append = pre + ".0." + "interface_identifier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interface-identifier"], _ = expandObjectFspVlanInterfaceIpv6InterfaceIdentifier2edl(d, i["interface_identifier"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-address"], _ = expandObjectFspVlanInterfaceIpv6Ip6Address2edl(d, i["ip6_address"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-allowaccess"], _ = expandObjectFspVlanInterfaceIpv6Ip6Allowaccess2edl(d, i["ip6_allowaccess"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-default-life"], _ = expandObjectFspVlanInterfaceIpv6Ip6DefaultLife2edl(d, i["ip6_default_life"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_delegated_prefix_iaid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-delegated-prefix-iaid"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid2edl(d, i["ip6_delegated_prefix_iaid"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList2edl(d, i["ip6_delegated_prefix_list"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ip6-delegated-prefix-list"] = t
		}
	}
	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-dns-server-override"], _ = expandObjectFspVlanInterfaceIpv6Ip6DnsServerOverride2edl(d, i["ip6_dns_server_override"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6ExtraAddr2edl(d, i["ip6_extra_addr"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ip6-extra-addr"] = t
		}
	}
	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-hop-limit"], _ = expandObjectFspVlanInterfaceIpv6Ip6HopLimit2edl(d, i["ip6_hop_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-link-mtu"], _ = expandObjectFspVlanInterfaceIpv6Ip6LinkMtu2edl(d, i["ip6_link_mtu"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-manage-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6ManageFlag2edl(d, i["ip6_manage_flag"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-max-interval"], _ = expandObjectFspVlanInterfaceIpv6Ip6MaxInterval2edl(d, i["ip6_max_interval"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-min-interval"], _ = expandObjectFspVlanInterfaceIpv6Ip6MinInterval2edl(d, i["ip6_min_interval"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-mode"], _ = expandObjectFspVlanInterfaceIpv6Ip6Mode2edl(d, i["ip6_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-other-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6OtherFlag2edl(d, i["ip6_other_flag"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6PrefixList2edl(d, i["ip6_prefix_list"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ip6-prefix-list"] = t
		}
	}
	pre_append = pre + ".0." + "ip6_prefix_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-prefix-mode"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixMode2edl(d, i["ip6_prefix_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-reachable-time"], _ = expandObjectFspVlanInterfaceIpv6Ip6ReachableTime2edl(d, i["ip6_reachable_time"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-retrans-time"], _ = expandObjectFspVlanInterfaceIpv6Ip6RetransTime2edl(d, i["ip6_retrans_time"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-send-adv"], _ = expandObjectFspVlanInterfaceIpv6Ip6SendAdv2edl(d, i["ip6_send_adv"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-subnet"], _ = expandObjectFspVlanInterfaceIpv6Ip6Subnet2edl(d, i["ip6_subnet"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-upstream-interface"], _ = expandObjectFspVlanInterfaceIpv6Ip6UpstreamInterface2edl(d, i["ip6_upstream_interface"], pre_append)
	}
	pre_append = pre + ".0." + "nd_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-cert"], _ = expandObjectFspVlanInterfaceIpv6NdCert2edl(d, i["nd_cert"], pre_append)
	}
	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-cga-modifier"], _ = expandObjectFspVlanInterfaceIpv6NdCgaModifier2edl(d, i["nd_cga_modifier"], pre_append)
	}
	pre_append = pre + ".0." + "nd_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-mode"], _ = expandObjectFspVlanInterfaceIpv6NdMode2edl(d, i["nd_mode"], pre_append)
	}
	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-security-level"], _ = expandObjectFspVlanInterfaceIpv6NdSecurityLevel2edl(d, i["nd_security_level"], pre_append)
	}
	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-timestamp-delta"], _ = expandObjectFspVlanInterfaceIpv6NdTimestampDelta2edl(d, i["nd_timestamp_delta"], pre_append)
	}
	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-timestamp-fuzz"], _ = expandObjectFspVlanInterfaceIpv6NdTimestampFuzz2edl(d, i["nd_timestamp_fuzz"], pre_append)
	}
	pre_append = pre + ".0." + "ra_send_mtu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ra-send-mtu"], _ = expandObjectFspVlanInterfaceIpv6RaSendMtu2edl(d, i["ra_send_mtu"], pre_append)
	}
	pre_append = pre + ".0." + "unique_autoconf_addr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unique-autoconf-addr"], _ = expandObjectFspVlanInterfaceIpv6UniqueAutoconfAddr2edl(d, i["unique_autoconf_addr"], pre_append)
	}
	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrip6_link_local"], _ = expandObjectFspVlanInterfaceIpv6Vrip6LinkLocal2edl(d, i["vrip6_link_local"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrrp-virtual-mac6"], _ = expandObjectFspVlanInterfaceIpv6VrrpVirtualMac62edl(d, i["vrrp_virtual_mac6"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp62edl(d, i["vrrp6"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["vrrp6"] = t
		}
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Autoconf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6CliConn6Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6ClientOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6InformationRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHint2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayInterfaceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelaySourceInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelaySourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Icmp6SendRedirect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6InterfaceIdentifier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Allowaccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DefaultLife2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["autonomous-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["delegated-prefix-iaid"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(d, i["delegated_prefix_iaid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-id"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(d, i["prefix_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(d, i["rdnss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss-service"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(d, i["rdnss_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["upstream-interface"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(d, i["upstream_interface"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DnsServerOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ExtraAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix2edl(d, i["prefix"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6HopLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6LinkMtu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ManageFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6MaxInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6MinInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6OtherFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["autonomous-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dnssl"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl2edl(d, i["dnssl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preferred-life-time"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(d, i["preferred_life_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix2edl(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss2edl(d, i["rdnss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["valid-life-time"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime2edl(d, i["valid_life_time"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ReachableTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6RetransTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6SendAdv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6UpstreamInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdCgaModifier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdSecurityLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdTimestampDelta2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdTimestampFuzz2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6RaSendMtu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6UniqueAutoconfAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrip6LinkLocal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6VrrpVirtualMac62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["accept-mode"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6AcceptMode2edl(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-interval"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6AdvInterval2edl(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preempt"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Preempt2edl(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Priority2edl(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-time"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6StartTime2edl(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Status2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst6"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrdst62edl(d, i["vrdst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrgrp"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrgrp2edl(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrid"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrid2edl(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrip6"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrip62edl(d, i["vrip6"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6AcceptMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6AdvInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Preempt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Priority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6StartTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrdst62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrgrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrip62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceL2Forward2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceL2TpClient2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLacpHaSecondary2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLacpHaSlave2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLacpMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLacpSpeed2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLargeReceiveOffload2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLcpEchoInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLcpMaxEchoFails2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLinkUpDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceListenForticlientConnection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLldpNetworkPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLldpReception2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLldpTransmission2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMacaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceManagedSubnetworkSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceManagementIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMaxEgressBurstRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMaxEgressRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMeasuredDownstreamBandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMeasuredUpstreamBandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMediatype2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMinLinks2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMinLinksDown2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMonitorBandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMtu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMtuOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMuxType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNdiscforward2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNetbiosForward2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNetflowSampler2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNpQosProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNpuFastpath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceOutForceVlanCos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceOutbandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePadtRetryTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfacePeerInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePhyMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePingServStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePoe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePollingInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePppoeUnnumberedNegotiate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePptpAuthType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePptpClient2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePptpPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfacePptpServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePptpTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePptpUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePreserveSessionRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePriorityOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceProxyCaptivePortal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcAtmQos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcChan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcCrc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcPcr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcScr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcVlanId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcVlanRxId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcVlanRxOp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcVlanTxId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcVlanTxOp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceReachableTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRedundantInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRemoteIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceReplacemsgOverrideGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRetransmission2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRingRx2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRingTx2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSampleDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSampleRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceScanBotnetConnections2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowaccess"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowaccess"], _ = expandObjectFspVlanInterfaceSecondaryipAllowaccess2edl(d, i["allowaccess"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detectprotocol"], _ = expandObjectFspVlanInterfaceSecondaryipDetectprotocol2edl(d, i["detectprotocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detectserver"], _ = expandObjectFspVlanInterfaceSecondaryipDetectserver2edl(d, i["detectserver"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gwdetect"], _ = expandObjectFspVlanInterfaceSecondaryipGwdetect2edl(d, i["gwdetect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandObjectFspVlanInterfaceSecondaryipHaPriority2edl(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanInterfaceSecondaryipId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanInterfaceSecondaryipIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secip_relay_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secip-relay-ip"], _ = expandObjectFspVlanInterfaceSecondaryipSecipRelayIp2edl(d, i["secip_relay_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ping-serv-status"], _ = expandObjectFspVlanInterfaceSecondaryipPingServStatus2edl(d, i["ping_serv_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandObjectFspVlanInterfaceSecondaryipSeq2edl(d, i["seq"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceSecondaryipAllowaccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceSecondaryipDetectprotocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceSecondaryipDetectserver2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipGwdetect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipHaPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipSecipRelayIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipPingServStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipSeq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurity8021XDynamicVlanId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurity8021XMaster2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurity8021XMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityExemptList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityExternalLogout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityExternalWeb2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityGroups2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityMacAuthBypass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityRedirectUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSelectProfile30A35B2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceServiceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSflowSampler2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSfpDsl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSfpDslAdslFallback2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSfpDslAutodetect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSfpDslMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSpeed2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSpilloverThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSrcCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStpHaSecondary2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStpHaSlave2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStpforward2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStpforwardMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStripPriorityVlanTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSubst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSubstituteDstMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwAlgorithm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwcFirstCreate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwcVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerAccessVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerArpInspection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerDhcpSnooping2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerDhcpSnoopingOption822edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerDhcpSnoopingVerifyMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerDynamic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerFeature2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerIgmpSnooping2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerIgmpSnoopingFastLeave2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerIgmpSnoopingProxy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerIotScanning2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerLearningLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerMgmtVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerNac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerNetflowCollect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerOffload2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerOffloadGw2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerOffloadIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerOffloading2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerOffloadingGw2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerOffloadingIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerRadiusServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerRspanMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerTrafficPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSystemId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSystemIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTcMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTcpMss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrunk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp612edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp622edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp632edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceUsername2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVci2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVectoring2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVindex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVlanProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVlanforward2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVlanid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVpi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["accept-mode"], _ = expandObjectFspVlanInterfaceVrrpAcceptMode2edl(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-interval"], _ = expandObjectFspVlanInterfaceVrrpAdvInterval2edl(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ignore-default-route"], _ = expandObjectFspVlanInterfaceVrrpIgnoreDefaultRoute2edl(d, i["ignore_default_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preempt"], _ = expandObjectFspVlanInterfaceVrrpPreempt2edl(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFspVlanInterfaceVrrpPriority2edl(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy_arp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFspVlanInterfaceVrrpProxyArp2edl(d, i["proxy_arp"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["proxy-arp"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-time"], _ = expandObjectFspVlanInterfaceVrrpStartTime2edl(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFspVlanInterfaceVrrpStatus2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandObjectFspVlanInterfaceVrrpVersion2edl(d, i["version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst"], _ = expandObjectFspVlanInterfaceVrrpVrdst2edl(d, i["vrdst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst-priority"], _ = expandObjectFspVlanInterfaceVrrpVrdstPriority2edl(d, i["vrdst_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrgrp"], _ = expandObjectFspVlanInterfaceVrrpVrgrp2edl(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrid"], _ = expandObjectFspVlanInterfaceVrrpVrid2edl(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrip"], _ = expandObjectFspVlanInterfaceVrrpVrip2edl(d, i["vrip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceVrrpAcceptMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpAdvInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpIgnoreDefaultRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpPreempt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpProxyArp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanInterfaceVrrpProxyArpId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanInterfaceVrrpProxyArpIp2edl(d, i["ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceVrrpProxyArpId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpProxyArpIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpStartTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrdst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceVrrpVrdstPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrgrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVirtualMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWccp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifi5GThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiAcl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiApBand2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiAutoConnect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiAutoSave2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiBroadcastSsid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiDnsServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiDnsServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiEncrypt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiFragmentThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceWifiKeyindex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiMacFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiPassphrase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceWifiRadiusServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiRtsThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiSecurity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiSsid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiUsergroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWinsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vlan_op_mode"); ok || d.HasChange("vlan_op_mode") {
		t, err := expandObjectFspVlanInterfaceVlanOpMode2edl(d, v, "vlan_op_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-op-mode"] = t
		}
	}

	if v, ok := d.GetOk("ac_name"); ok || d.HasChange("ac_name") {
		t, err := expandObjectFspVlanInterfaceAcName2edl(d, v, "ac_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-name"] = t
		}
	}

	if v, ok := d.GetOk("aggregate"); ok || d.HasChange("aggregate") {
		t, err := expandObjectFspVlanInterfaceAggregate2edl(d, v, "aggregate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate"] = t
		}
	}

	if v, ok := d.GetOk("aggregate_type"); ok || d.HasChange("aggregate_type") {
		t, err := expandObjectFspVlanInterfaceAggregateType2edl(d, v, "aggregate_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-type"] = t
		}
	}

	if v, ok := d.GetOk("algorithm"); ok || d.HasChange("algorithm") {
		t, err := expandObjectFspVlanInterfaceAlgorithm2edl(d, v, "algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["algorithm"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok || d.HasChange("alias") {
		t, err := expandObjectFspVlanInterfaceAlias2edl(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandObjectFspVlanInterfaceAllowaccess2edl(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("ap_discover"); ok || d.HasChange("ap_discover") {
		t, err := expandObjectFspVlanInterfaceApDiscover2edl(d, v, "ap_discover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-discover"] = t
		}
	}

	if v, ok := d.GetOk("arpforward"); ok || d.HasChange("arpforward") {
		t, err := expandObjectFspVlanInterfaceArpforward2edl(d, v, "arpforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arpforward"] = t
		}
	}

	if v, ok := d.GetOk("atm_protocol"); ok || d.HasChange("atm_protocol") {
		t, err := expandObjectFspVlanInterfaceAtmProtocol2edl(d, v, "atm_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["atm-protocol"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok || d.HasChange("auth_cert") {
		t, err := expandObjectFspVlanInterfaceAuthCert2edl(d, v, "auth_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal_addr"); ok || d.HasChange("auth_portal_addr") {
		t, err := expandObjectFspVlanInterfaceAuthPortalAddr2edl(d, v, "auth_portal_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal-addr"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandObjectFspVlanInterfaceAuthType2edl(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("auto_auth_extension_device"); ok || d.HasChange("auto_auth_extension_device") {
		t, err := expandObjectFspVlanInterfaceAutoAuthExtensionDevice2edl(d, v, "auto_auth_extension_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-auth-extension-device"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_measure_time"); ok || d.HasChange("bandwidth_measure_time") {
		t, err := expandObjectFspVlanInterfaceBandwidthMeasureTime2edl(d, v, "bandwidth_measure_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-measure-time"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandObjectFspVlanInterfaceBfd2edl(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("bfd_desired_min_tx"); ok || d.HasChange("bfd_desired_min_tx") {
		t, err := expandObjectFspVlanInterfaceBfdDesiredMinTx2edl(d, v, "bfd_desired_min_tx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-desired-min-tx"] = t
		}
	}

	if v, ok := d.GetOk("bfd_detect_mult"); ok || d.HasChange("bfd_detect_mult") {
		t, err := expandObjectFspVlanInterfaceBfdDetectMult2edl(d, v, "bfd_detect_mult")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-detect-mult"] = t
		}
	}

	if v, ok := d.GetOk("bfd_required_min_rx"); ok || d.HasChange("bfd_required_min_rx") {
		t, err := expandObjectFspVlanInterfaceBfdRequiredMinRx2edl(d, v, "bfd_required_min_rx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-required-min-rx"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_forticlient_discovery"); ok || d.HasChange("broadcast_forticlient_discovery") {
		t, err := expandObjectFspVlanInterfaceBroadcastForticlientDiscovery2edl(d, v, "broadcast_forticlient_discovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-forticlient-discovery"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_forward"); ok || d.HasChange("broadcast_forward") {
		t, err := expandObjectFspVlanInterfaceBroadcastForward2edl(d, v, "broadcast_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-forward"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal"); ok || d.HasChange("captive_portal") {
		t, err := expandObjectFspVlanInterfaceCaptivePortal2edl(d, v, "captive_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("cli_conn_status"); ok || d.HasChange("cli_conn_status") {
		t, err := expandObjectFspVlanInterfaceCliConnStatus2edl(d, v, "cli_conn_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-conn-status"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectFspVlanInterfaceColor2edl(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("ddns"); ok || d.HasChange("ddns") {
		t, err := expandObjectFspVlanInterfaceDdns2edl(d, v, "ddns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns"] = t
		}
	}

	if v, ok := d.GetOk("ddns_auth"); ok || d.HasChange("ddns_auth") {
		t, err := expandObjectFspVlanInterfaceDdnsAuth2edl(d, v, "ddns_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-auth"] = t
		}
	}

	if v, ok := d.GetOk("ddns_domain"); ok || d.HasChange("ddns_domain") {
		t, err := expandObjectFspVlanInterfaceDdnsDomain2edl(d, v, "ddns_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-domain"] = t
		}
	}

	if v, ok := d.GetOk("ddns_key"); ok || d.HasChange("ddns_key") {
		t, err := expandObjectFspVlanInterfaceDdnsKey2edl(d, v, "ddns_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-key"] = t
		}
	}

	if v, ok := d.GetOk("ddns_keyname"); ok || d.HasChange("ddns_keyname") {
		t, err := expandObjectFspVlanInterfaceDdnsKeyname2edl(d, v, "ddns_keyname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-keyname"] = t
		}
	}

	if v, ok := d.GetOk("ddns_password"); ok || d.HasChange("ddns_password") {
		t, err := expandObjectFspVlanInterfaceDdnsPassword2edl(d, v, "ddns_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-password"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server"); ok || d.HasChange("ddns_server") {
		t, err := expandObjectFspVlanInterfaceDdnsServer2edl(d, v, "ddns_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok || d.HasChange("ddns_server_ip") {
		t, err := expandObjectFspVlanInterfaceDdnsServerIp2edl(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_sn"); ok || d.HasChange("ddns_sn") {
		t, err := expandObjectFspVlanInterfaceDdnsSn2edl(d, v, "ddns_sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-sn"] = t
		}
	}

	if v, ok := d.GetOk("ddns_ttl"); ok || d.HasChange("ddns_ttl") {
		t, err := expandObjectFspVlanInterfaceDdnsTtl2edl(d, v, "ddns_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-ttl"] = t
		}
	}

	if v, ok := d.GetOk("ddns_username"); ok || d.HasChange("ddns_username") {
		t, err := expandObjectFspVlanInterfaceDdnsUsername2edl(d, v, "ddns_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-username"] = t
		}
	}

	if v, ok := d.GetOk("ddns_zone"); ok || d.HasChange("ddns_zone") {
		t, err := expandObjectFspVlanInterfaceDdnsZone2edl(d, v, "ddns_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-zone"] = t
		}
	}

	if v, ok := d.GetOk("dedicated_to"); ok || d.HasChange("dedicated_to") {
		t, err := expandObjectFspVlanInterfaceDedicatedTo2edl(d, v, "dedicated_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dedicated-to"] = t
		}
	}

	if v, ok := d.GetOk("default_purdue_level"); ok || d.HasChange("default_purdue_level") {
		t, err := expandObjectFspVlanInterfaceDefaultPurdueLevel2edl(d, v, "default_purdue_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-purdue-level"] = t
		}
	}

	if v, ok := d.GetOk("defaultgw"); ok || d.HasChange("defaultgw") {
		t, err := expandObjectFspVlanInterfaceDefaultgw2edl(d, v, "defaultgw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["defaultgw"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectFspVlanInterfaceDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("detected_peer_mtu"); ok || d.HasChange("detected_peer_mtu") {
		t, err := expandObjectFspVlanInterfaceDetectedPeerMtu2edl(d, v, "detected_peer_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detected-peer-mtu"] = t
		}
	}

	if v, ok := d.GetOk("detectprotocol"); ok || d.HasChange("detectprotocol") {
		t, err := expandObjectFspVlanInterfaceDetectprotocol2edl(d, v, "detectprotocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detectprotocol"] = t
		}
	}

	if v, ok := d.GetOk("detectserver"); ok || d.HasChange("detectserver") {
		t, err := expandObjectFspVlanInterfaceDetectserver2edl(d, v, "detectserver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detectserver"] = t
		}
	}

	if v, ok := d.GetOk("device_access_list"); ok || d.HasChange("device_access_list") {
		t, err := expandObjectFspVlanInterfaceDeviceAccessList2edl(d, v, "device_access_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-access-list"] = t
		}
	}

	if v, ok := d.GetOk("device_identification"); ok || d.HasChange("device_identification") {
		t, err := expandObjectFspVlanInterfaceDeviceIdentification2edl(d, v, "device_identification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-identification"] = t
		}
	}

	if v, ok := d.GetOk("device_identification_active_scan"); ok || d.HasChange("device_identification_active_scan") {
		t, err := expandObjectFspVlanInterfaceDeviceIdentificationActiveScan2edl(d, v, "device_identification_active_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-identification-active-scan"] = t
		}
	}

	if v, ok := d.GetOk("device_netscan"); ok || d.HasChange("device_netscan") {
		t, err := expandObjectFspVlanInterfaceDeviceNetscan2edl(d, v, "device_netscan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-netscan"] = t
		}
	}

	if v, ok := d.GetOk("device_user_identification"); ok || d.HasChange("device_user_identification") {
		t, err := expandObjectFspVlanInterfaceDeviceUserIdentification2edl(d, v, "device_user_identification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-user-identification"] = t
		}
	}

	if v, ok := d.GetOk("devindex"); ok || d.HasChange("devindex") {
		t, err := expandObjectFspVlanInterfaceDevindex2edl(d, v, "devindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devindex"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_broadcast_flag"); ok || d.HasChange("dhcp_broadcast_flag") {
		t, err := expandObjectFspVlanInterfaceDhcpBroadcastFlag2edl(d, v, "dhcp_broadcast_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-broadcast-flag"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_classless_route_addition"); ok || d.HasChange("dhcp_classless_route_addition") {
		t, err := expandObjectFspVlanInterfaceDhcpClasslessRouteAddition2edl(d, v, "dhcp_classless_route_addition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-classless-route-addition"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_client_identifier"); ok || d.HasChange("dhcp_client_identifier") {
		t, err := expandObjectFspVlanInterfaceDhcpClientIdentifier2edl(d, v, "dhcp_client_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-client-identifier"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_agent_option"); ok || d.HasChange("dhcp_relay_agent_option") {
		t, err := expandObjectFspVlanInterfaceDhcpRelayAgentOption2edl(d, v, "dhcp_relay_agent_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-agent-option"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_circuit_id"); ok || d.HasChange("dhcp_relay_circuit_id") {
		t, err := expandObjectFspVlanInterfaceDhcpRelayCircuitId2edl(d, v, "dhcp_relay_circuit_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-circuit-id"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_interface"); ok || d.HasChange("dhcp_relay_interface") {
		t, err := expandObjectFspVlanInterfaceDhcpRelayInterface2edl(d, v, "dhcp_relay_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-interface"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_interface_select_method"); ok || d.HasChange("dhcp_relay_interface_select_method") {
		t, err := expandObjectFspVlanInterfaceDhcpRelayInterfaceSelectMethod2edl(d, v, "dhcp_relay_interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_ip"); ok || d.HasChange("dhcp_relay_ip") {
		t, err := expandObjectFspVlanInterfaceDhcpRelayIp2edl(d, v, "dhcp_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_link_selection"); ok || d.HasChange("dhcp_relay_link_selection") {
		t, err := expandObjectFspVlanInterfaceDhcpRelayLinkSelection2edl(d, v, "dhcp_relay_link_selection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-link-selection"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_request_all_server"); ok || d.HasChange("dhcp_relay_request_all_server") {
		t, err := expandObjectFspVlanInterfaceDhcpRelayRequestAllServer2edl(d, v, "dhcp_relay_request_all_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-request-all-server"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_service"); ok || d.HasChange("dhcp_relay_service") {
		t, err := expandObjectFspVlanInterfaceDhcpRelayService2edl(d, v, "dhcp_relay_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-service"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_source_ip"); ok || d.HasChange("dhcp_relay_source_ip") {
		t, err := expandObjectFspVlanInterfaceDhcpRelaySourceIp2edl(d, v, "dhcp_relay_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_type"); ok || d.HasChange("dhcp_relay_type") {
		t, err := expandObjectFspVlanInterfaceDhcpRelayType2edl(d, v, "dhcp_relay_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-type"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_renew_time"); ok || d.HasChange("dhcp_renew_time") {
		t, err := expandObjectFspVlanInterfaceDhcpRenewTime2edl(d, v, "dhcp_renew_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-renew-time"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_smart_relay"); ok || d.HasChange("dhcp_smart_relay") {
		t, err := expandObjectFspVlanInterfaceDhcpSmartRelay2edl(d, v, "dhcp_smart_relay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-smart-relay"] = t
		}
	}

	if v, ok := d.GetOk("disc_retry_timeout"); ok || d.HasChange("disc_retry_timeout") {
		t, err := expandObjectFspVlanInterfaceDiscRetryTimeout2edl(d, v, "disc_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disc-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("disconnect_threshold"); ok || d.HasChange("disconnect_threshold") {
		t, err := expandObjectFspVlanInterfaceDisconnectThreshold2edl(d, v, "disconnect_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disconnect-threshold"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandObjectFspVlanInterfaceDistance2edl(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("dns_query"); ok || d.HasChange("dns_query") {
		t, err := expandObjectFspVlanInterfaceDnsQuery2edl(d, v, "dns_query")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-query"] = t
		}
	}

	if v, ok := d.GetOk("dns_server_override"); ok || d.HasChange("dns_server_override") {
		t, err := expandObjectFspVlanInterfaceDnsServerOverride2edl(d, v, "dns_server_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server-override"] = t
		}
	}

	if v, ok := d.GetOk("dns_server_protocol"); ok || d.HasChange("dns_server_protocol") {
		t, err := expandObjectFspVlanInterfaceDnsServerProtocol2edl(d, v, "dns_server_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server-protocol"] = t
		}
	}

	if v, ok := d.GetOk("drop_fragment"); ok || d.HasChange("drop_fragment") {
		t, err := expandObjectFspVlanInterfaceDropFragment2edl(d, v, "drop_fragment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-fragment"] = t
		}
	}

	if v, ok := d.GetOk("drop_overlapped_fragment"); ok || d.HasChange("drop_overlapped_fragment") {
		t, err := expandObjectFspVlanInterfaceDropOverlappedFragment2edl(d, v, "drop_overlapped_fragment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-overlapped-fragment"] = t
		}
	}

	if v, ok := d.GetOk("eap_ca_cert"); ok || d.HasChange("eap_ca_cert") {
		t, err := expandObjectFspVlanInterfaceEapCaCert2edl(d, v, "eap_ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("eap_identity"); ok || d.HasChange("eap_identity") {
		t, err := expandObjectFspVlanInterfaceEapIdentity2edl(d, v, "eap_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-identity"] = t
		}
	}

	if v, ok := d.GetOk("eap_method"); ok || d.HasChange("eap_method") {
		t, err := expandObjectFspVlanInterfaceEapMethod2edl(d, v, "eap_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-method"] = t
		}
	}

	if v, ok := d.GetOk("eap_password"); ok || d.HasChange("eap_password") {
		t, err := expandObjectFspVlanInterfaceEapPassword2edl(d, v, "eap_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-password"] = t
		}
	}

	if v, ok := d.GetOk("eap_supplicant"); ok || d.HasChange("eap_supplicant") {
		t, err := expandObjectFspVlanInterfaceEapSupplicant2edl(d, v, "eap_supplicant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-supplicant"] = t
		}
	}

	if v, ok := d.GetOk("eap_user_cert"); ok || d.HasChange("eap_user_cert") {
		t, err := expandObjectFspVlanInterfaceEapUserCert2edl(d, v, "eap_user_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-user-cert"] = t
		}
	}

	if v, ok := d.GetOk("egress_cos"); ok || d.HasChange("egress_cos") {
		t, err := expandObjectFspVlanInterfaceEgressCos2edl(d, v, "egress_cos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-cos"] = t
		}
	}

	if v, ok := d.GetOk("egress_shaping_profile"); ok || d.HasChange("egress_shaping_profile") {
		t, err := expandObjectFspVlanInterfaceEgressShapingProfile2edl(d, v, "egress_shaping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-shaping-profile"] = t
		}
	}

	if v, ok := d.GetOk("eip"); ok || d.HasChange("eip") {
		t, err := expandObjectFspVlanInterfaceEip2edl(d, v, "eip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eip"] = t
		}
	}

	if v, ok := d.GetOk("endpoint_compliance"); ok || d.HasChange("endpoint_compliance") {
		t, err := expandObjectFspVlanInterfaceEndpointCompliance2edl(d, v, "endpoint_compliance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endpoint-compliance"] = t
		}
	}

	if v, ok := d.GetOk("estimated_downstream_bandwidth"); ok || d.HasChange("estimated_downstream_bandwidth") {
		t, err := expandObjectFspVlanInterfaceEstimatedDownstreamBandwidth2edl(d, v, "estimated_downstream_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["estimated-downstream-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("estimated_upstream_bandwidth"); ok || d.HasChange("estimated_upstream_bandwidth") {
		t, err := expandObjectFspVlanInterfaceEstimatedUpstreamBandwidth2edl(d, v, "estimated_upstream_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["estimated-upstream-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("explicit_ftp_proxy"); ok || d.HasChange("explicit_ftp_proxy") {
		t, err := expandObjectFspVlanInterfaceExplicitFtpProxy2edl(d, v, "explicit_ftp_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["explicit-ftp-proxy"] = t
		}
	}

	if v, ok := d.GetOk("explicit_web_proxy"); ok || d.HasChange("explicit_web_proxy") {
		t, err := expandObjectFspVlanInterfaceExplicitWebProxy2edl(d, v, "explicit_web_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["explicit-web-proxy"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok || d.HasChange("external") {
		t, err := expandObjectFspVlanInterfaceExternal2edl(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("fail_action_on_extender"); ok || d.HasChange("fail_action_on_extender") {
		t, err := expandObjectFspVlanInterfaceFailActionOnExtender2edl(d, v, "fail_action_on_extender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-action-on-extender"] = t
		}
	}

	if v, ok := d.GetOk("fail_alert_interfaces"); ok || d.HasChange("fail_alert_interfaces") {
		t, err := expandObjectFspVlanInterfaceFailAlertInterfaces2edl(d, v, "fail_alert_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-alert-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("fail_alert_method"); ok || d.HasChange("fail_alert_method") {
		t, err := expandObjectFspVlanInterfaceFailAlertMethod2edl(d, v, "fail_alert_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-alert-method"] = t
		}
	}

	if v, ok := d.GetOk("fail_detect"); ok || d.HasChange("fail_detect") {
		t, err := expandObjectFspVlanInterfaceFailDetect2edl(d, v, "fail_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-detect"] = t
		}
	}

	if v, ok := d.GetOk("fail_detect_option"); ok || d.HasChange("fail_detect_option") {
		t, err := expandObjectFspVlanInterfaceFailDetectOption2edl(d, v, "fail_detect_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-detect-option"] = t
		}
	}

	if v, ok := d.GetOk("fdp"); ok || d.HasChange("fdp") {
		t, err := expandObjectFspVlanInterfaceFdp2edl(d, v, "fdp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fdp"] = t
		}
	}

	if v, ok := d.GetOk("fortiheartbeat"); ok || d.HasChange("fortiheartbeat") {
		t, err := expandObjectFspVlanInterfaceFortiheartbeat2edl(d, v, "fortiheartbeat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiheartbeat"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok || d.HasChange("fortilink") {
		t, err := expandObjectFspVlanInterfaceFortilink2edl(d, v, "fortilink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("fortilink_backup_link"); ok || d.HasChange("fortilink_backup_link") {
		t, err := expandObjectFspVlanInterfaceFortilinkBackupLink2edl(d, v, "fortilink_backup_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink-backup-link"] = t
		}
	}

	if v, ok := d.GetOk("fortilink_neighbor_detect"); ok || d.HasChange("fortilink_neighbor_detect") {
		t, err := expandObjectFspVlanInterfaceFortilinkNeighborDetect2edl(d, v, "fortilink_neighbor_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink-neighbor-detect"] = t
		}
	}

	if v, ok := d.GetOk("fortilink_split_interface"); ok || d.HasChange("fortilink_split_interface") {
		t, err := expandObjectFspVlanInterfaceFortilinkSplitInterface2edl(d, v, "fortilink_split_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink-split-interface"] = t
		}
	}

	if v, ok := d.GetOk("fortilink_stacking"); ok || d.HasChange("fortilink_stacking") {
		t, err := expandObjectFspVlanInterfaceFortilinkStacking2edl(d, v, "fortilink_stacking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink-stacking"] = t
		}
	}

	if v, ok := d.GetOk("forward_domain"); ok || d.HasChange("forward_domain") {
		t, err := expandObjectFspVlanInterfaceForwardDomain2edl(d, v, "forward_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-domain"] = t
		}
	}

	if v, ok := d.GetOk("forward_error_correction"); ok || d.HasChange("forward_error_correction") {
		t, err := expandObjectFspVlanInterfaceForwardErrorCorrection2edl(d, v, "forward_error_correction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-error-correction"] = t
		}
	}

	if v, ok := d.GetOk("fp_anomaly"); ok || d.HasChange("fp_anomaly") {
		t, err := expandObjectFspVlanInterfaceFpAnomaly2edl(d, v, "fp_anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fp-anomaly"] = t
		}
	}

	if v, ok := d.GetOk("fp_disable"); ok || d.HasChange("fp_disable") {
		t, err := expandObjectFspVlanInterfaceFpDisable2edl(d, v, "fp_disable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fp-disable"] = t
		}
	}

	if v, ok := d.GetOk("gateway_address"); ok || d.HasChange("gateway_address") {
		t, err := expandObjectFspVlanInterfaceGatewayAddress2edl(d, v, "gateway_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway-address"] = t
		}
	}

	if v, ok := d.GetOk("generic_receive_offload"); ok || d.HasChange("generic_receive_offload") {
		t, err := expandObjectFspVlanInterfaceGenericReceiveOffload2edl(d, v, "generic_receive_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["generic-receive-offload"] = t
		}
	}

	if v, ok := d.GetOk("gi_gk"); ok || d.HasChange("gi_gk") {
		t, err := expandObjectFspVlanInterfaceGiGk2edl(d, v, "gi_gk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gi-gk"] = t
		}
	}

	if v, ok := d.GetOk("gwaddr"); ok || d.HasChange("gwaddr") {
		t, err := expandObjectFspVlanInterfaceGwaddr2edl(d, v, "gwaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gwaddr"] = t
		}
	}

	if v, ok := d.GetOk("gwdetect"); ok || d.HasChange("gwdetect") {
		t, err := expandObjectFspVlanInterfaceGwdetect2edl(d, v, "gwdetect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gwdetect"] = t
		}
	}

	if v, ok := d.GetOk("ha_priority"); ok || d.HasChange("ha_priority") {
		t, err := expandObjectFspVlanInterfaceHaPriority2edl(d, v, "ha_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-priority"] = t
		}
	}

	if v, ok := d.GetOk("icmp_accept_redirect"); ok || d.HasChange("icmp_accept_redirect") {
		t, err := expandObjectFspVlanInterfaceIcmpAcceptRedirect2edl(d, v, "icmp_accept_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-accept-redirect"] = t
		}
	}

	if v, ok := d.GetOk("icmp_redirect"); ok || d.HasChange("icmp_redirect") {
		t, err := expandObjectFspVlanInterfaceIcmpRedirect2edl(d, v, "icmp_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-redirect"] = t
		}
	}

	if v, ok := d.GetOk("icmp_send_redirect"); ok || d.HasChange("icmp_send_redirect") {
		t, err := expandObjectFspVlanInterfaceIcmpSendRedirect2edl(d, v, "icmp_send_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-send-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ident_accept"); ok || d.HasChange("ident_accept") {
		t, err := expandObjectFspVlanInterfaceIdentAccept2edl(d, v, "ident_accept")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ident-accept"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok || d.HasChange("idle_timeout") {
		t, err := expandObjectFspVlanInterfaceIdleTimeout2edl(d, v, "idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("if_mdix"); ok || d.HasChange("if_mdix") {
		t, err := expandObjectFspVlanInterfaceIfMdix2edl(d, v, "if_mdix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["if-mdix"] = t
		}
	}

	if v, ok := d.GetOk("if_media"); ok || d.HasChange("if_media") {
		t, err := expandObjectFspVlanInterfaceIfMedia2edl(d, v, "if_media")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["if-media"] = t
		}
	}

	if v, ok := d.GetOk("ike_saml_server"); ok || d.HasChange("ike_saml_server") {
		t, err := expandObjectFspVlanInterfaceIkeSamlServer2edl(d, v, "ike_saml_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-saml-server"] = t
		}
	}

	if v, ok := d.GetOk("in_force_vlan_cos"); ok || d.HasChange("in_force_vlan_cos") {
		t, err := expandObjectFspVlanInterfaceInForceVlanCos2edl(d, v, "in_force_vlan_cos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["in-force-vlan-cos"] = t
		}
	}

	if v, ok := d.GetOk("inbandwidth"); ok || d.HasChange("inbandwidth") {
		t, err := expandObjectFspVlanInterfaceInbandwidth2edl(d, v, "inbandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("ingress_cos"); ok || d.HasChange("ingress_cos") {
		t, err := expandObjectFspVlanInterfaceIngressCos2edl(d, v, "ingress_cos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-cos"] = t
		}
	}

	if v, ok := d.GetOk("ingress_shaping_profile"); ok || d.HasChange("ingress_shaping_profile") {
		t, err := expandObjectFspVlanInterfaceIngressShapingProfile2edl(d, v, "ingress_shaping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-shaping-profile"] = t
		}
	}

	if v, ok := d.GetOk("ingress_spillover_threshold"); ok || d.HasChange("ingress_spillover_threshold") {
		t, err := expandObjectFspVlanInterfaceIngressSpilloverThreshold2edl(d, v, "ingress_spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("interconnect_profile"); ok || d.HasChange("interconnect_profile") {
		t, err := expandObjectFspVlanInterfaceInterconnectProfile2edl(d, v, "interconnect_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interconnect-profile"] = t
		}
	}

	if v, ok := d.GetOk("internal"); ok || d.HasChange("internal") {
		t, err := expandObjectFspVlanInterfaceInternal2edl(d, v, "internal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectFspVlanInterfaceIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip_managed_by_fortiipam"); ok || d.HasChange("ip_managed_by_fortiipam") {
		t, err := expandObjectFspVlanInterfaceIpManagedByFortiipam2edl(d, v, "ip_managed_by_fortiipam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-managed-by-fortiipam"] = t
		}
	}

	if v, ok := d.GetOk("ipmac"); ok || d.HasChange("ipmac") {
		t, err := expandObjectFspVlanInterfaceIpmac2edl(d, v, "ipmac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipmac"] = t
		}
	}

	if v, ok := d.GetOk("ips_sniffer_mode"); ok || d.HasChange("ips_sniffer_mode") {
		t, err := expandObjectFspVlanInterfaceIpsSnifferMode2edl(d, v, "ips_sniffer_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sniffer-mode"] = t
		}
	}

	if v, ok := d.GetOk("ipunnumbered"); ok || d.HasChange("ipunnumbered") {
		t, err := expandObjectFspVlanInterfaceIpunnumbered2edl(d, v, "ipunnumbered")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipunnumbered"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandObjectFspVlanInterfaceIpv62edl(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("l2forward"); ok || d.HasChange("l2forward") {
		t, err := expandObjectFspVlanInterfaceL2Forward2edl(d, v, "l2forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2forward"] = t
		}
	}

	if v, ok := d.GetOk("l2tp_client"); ok || d.HasChange("l2tp_client") {
		t, err := expandObjectFspVlanInterfaceL2TpClient2edl(d, v, "l2tp_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2tp-client"] = t
		}
	}

	if v, ok := d.GetOk("lacp_ha_secondary"); ok || d.HasChange("lacp_ha_secondary") {
		t, err := expandObjectFspVlanInterfaceLacpHaSecondary2edl(d, v, "lacp_ha_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-ha-secondary"] = t
		}
	}

	if v, ok := d.GetOk("lacp_ha_slave"); ok || d.HasChange("lacp_ha_slave") {
		t, err := expandObjectFspVlanInterfaceLacpHaSlave2edl(d, v, "lacp_ha_slave")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-ha-slave"] = t
		}
	}

	if v, ok := d.GetOk("lacp_mode"); ok || d.HasChange("lacp_mode") {
		t, err := expandObjectFspVlanInterfaceLacpMode2edl(d, v, "lacp_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-mode"] = t
		}
	}

	if v, ok := d.GetOk("lacp_speed"); ok || d.HasChange("lacp_speed") {
		t, err := expandObjectFspVlanInterfaceLacpSpeed2edl(d, v, "lacp_speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-speed"] = t
		}
	}

	if v, ok := d.GetOk("large_receive_offload"); ok || d.HasChange("large_receive_offload") {
		t, err := expandObjectFspVlanInterfaceLargeReceiveOffload2edl(d, v, "large_receive_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["large-receive-offload"] = t
		}
	}

	if v, ok := d.GetOk("lcp_echo_interval"); ok || d.HasChange("lcp_echo_interval") {
		t, err := expandObjectFspVlanInterfaceLcpEchoInterval2edl(d, v, "lcp_echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lcp-echo-interval"] = t
		}
	}

	if v, ok := d.GetOk("lcp_max_echo_fails"); ok || d.HasChange("lcp_max_echo_fails") {
		t, err := expandObjectFspVlanInterfaceLcpMaxEchoFails2edl(d, v, "lcp_max_echo_fails")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lcp-max-echo-fails"] = t
		}
	}

	if v, ok := d.GetOk("link_up_delay"); ok || d.HasChange("link_up_delay") {
		t, err := expandObjectFspVlanInterfaceLinkUpDelay2edl(d, v, "link_up_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-up-delay"] = t
		}
	}

	if v, ok := d.GetOk("listen_forticlient_connection"); ok || d.HasChange("listen_forticlient_connection") {
		t, err := expandObjectFspVlanInterfaceListenForticlientConnection2edl(d, v, "listen_forticlient_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["listen-forticlient-connection"] = t
		}
	}

	if v, ok := d.GetOk("lldp_network_policy"); ok || d.HasChange("lldp_network_policy") {
		t, err := expandObjectFspVlanInterfaceLldpNetworkPolicy2edl(d, v, "lldp_network_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-network-policy"] = t
		}
	}

	if v, ok := d.GetOk("lldp_reception"); ok || d.HasChange("lldp_reception") {
		t, err := expandObjectFspVlanInterfaceLldpReception2edl(d, v, "lldp_reception")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-reception"] = t
		}
	}

	if v, ok := d.GetOk("lldp_transmission"); ok || d.HasChange("lldp_transmission") {
		t, err := expandObjectFspVlanInterfaceLldpTransmission2edl(d, v, "lldp_transmission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-transmission"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectFspVlanInterfaceLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("macaddr"); ok || d.HasChange("macaddr") {
		t, err := expandObjectFspVlanInterfaceMacaddr2edl(d, v, "macaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macaddr"] = t
		}
	}

	if v, ok := d.GetOk("managed_subnetwork_size"); ok || d.HasChange("managed_subnetwork_size") {
		t, err := expandObjectFspVlanInterfaceManagedSubnetworkSize2edl(d, v, "managed_subnetwork_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["managed-subnetwork-size"] = t
		}
	}

	if v, ok := d.GetOk("management_ip"); ok || d.HasChange("management_ip") {
		t, err := expandObjectFspVlanInterfaceManagementIp2edl(d, v, "management_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-ip"] = t
		}
	}

	if v, ok := d.GetOk("max_egress_burst_rate"); ok || d.HasChange("max_egress_burst_rate") {
		t, err := expandObjectFspVlanInterfaceMaxEgressBurstRate2edl(d, v, "max_egress_burst_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-egress-burst-rate"] = t
		}
	}

	if v, ok := d.GetOk("max_egress_rate"); ok || d.HasChange("max_egress_rate") {
		t, err := expandObjectFspVlanInterfaceMaxEgressRate2edl(d, v, "max_egress_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-egress-rate"] = t
		}
	}

	if v, ok := d.GetOk("measured_downstream_bandwidth"); ok || d.HasChange("measured_downstream_bandwidth") {
		t, err := expandObjectFspVlanInterfaceMeasuredDownstreamBandwidth2edl(d, v, "measured_downstream_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["measured-downstream-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("measured_upstream_bandwidth"); ok || d.HasChange("measured_upstream_bandwidth") {
		t, err := expandObjectFspVlanInterfaceMeasuredUpstreamBandwidth2edl(d, v, "measured_upstream_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["measured-upstream-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("mediatype"); ok || d.HasChange("mediatype") {
		t, err := expandObjectFspVlanInterfaceMediatype2edl(d, v, "mediatype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mediatype"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandObjectFspVlanInterfaceMember2edl(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("min_links"); ok || d.HasChange("min_links") {
		t, err := expandObjectFspVlanInterfaceMinLinks2edl(d, v, "min_links")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-links"] = t
		}
	}

	if v, ok := d.GetOk("min_links_down"); ok || d.HasChange("min_links_down") {
		t, err := expandObjectFspVlanInterfaceMinLinksDown2edl(d, v, "min_links_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-links-down"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandObjectFspVlanInterfaceMode2edl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("monitor_bandwidth"); ok || d.HasChange("monitor_bandwidth") {
		t, err := expandObjectFspVlanInterfaceMonitorBandwidth2edl(d, v, "monitor_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("mtu"); ok || d.HasChange("mtu") {
		t, err := expandObjectFspVlanInterfaceMtu2edl(d, v, "mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu"] = t
		}
	}

	if v, ok := d.GetOk("mtu_override"); ok || d.HasChange("mtu_override") {
		t, err := expandObjectFspVlanInterfaceMtuOverride2edl(d, v, "mtu_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu-override"] = t
		}
	}

	if v, ok := d.GetOk("mux_type"); ok || d.HasChange("mux_type") {
		t, err := expandObjectFspVlanInterfaceMuxType2edl(d, v, "mux_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mux-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFspVlanInterfaceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ndiscforward"); ok || d.HasChange("ndiscforward") {
		t, err := expandObjectFspVlanInterfaceNdiscforward2edl(d, v, "ndiscforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ndiscforward"] = t
		}
	}

	if v, ok := d.GetOk("netbios_forward"); ok || d.HasChange("netbios_forward") {
		t, err := expandObjectFspVlanInterfaceNetbiosForward2edl(d, v, "netbios_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netbios-forward"] = t
		}
	}

	if v, ok := d.GetOk("netflow_sampler"); ok || d.HasChange("netflow_sampler") {
		t, err := expandObjectFspVlanInterfaceNetflowSampler2edl(d, v, "netflow_sampler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netflow-sampler"] = t
		}
	}

	if v, ok := d.GetOk("np_qos_profile"); ok || d.HasChange("np_qos_profile") {
		t, err := expandObjectFspVlanInterfaceNpQosProfile2edl(d, v, "np_qos_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np-qos-profile"] = t
		}
	}

	if v, ok := d.GetOk("npu_fastpath"); ok || d.HasChange("npu_fastpath") {
		t, err := expandObjectFspVlanInterfaceNpuFastpath2edl(d, v, "npu_fastpath")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-fastpath"] = t
		}
	}

	if v, ok := d.GetOk("nst"); ok || d.HasChange("nst") {
		t, err := expandObjectFspVlanInterfaceNst2edl(d, v, "nst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nst"] = t
		}
	}

	if v, ok := d.GetOk("out_force_vlan_cos"); ok || d.HasChange("out_force_vlan_cos") {
		t, err := expandObjectFspVlanInterfaceOutForceVlanCos2edl(d, v, "out_force_vlan_cos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["out-force-vlan-cos"] = t
		}
	}

	if v, ok := d.GetOk("outbandwidth"); ok || d.HasChange("outbandwidth") {
		t, err := expandObjectFspVlanInterfaceOutbandwidth2edl(d, v, "outbandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("padt_retry_timeout"); ok || d.HasChange("padt_retry_timeout") {
		t, err := expandObjectFspVlanInterfacePadtRetryTimeout2edl(d, v, "padt_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padt-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectFspVlanInterfacePassword2edl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("peer_interface"); ok || d.HasChange("peer_interface") {
		t, err := expandObjectFspVlanInterfacePeerInterface2edl(d, v, "peer_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-interface"] = t
		}
	}

	if v, ok := d.GetOk("phy_mode"); ok || d.HasChange("phy_mode") {
		t, err := expandObjectFspVlanInterfacePhyMode2edl(d, v, "phy_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phy-mode"] = t
		}
	}

	if v, ok := d.GetOk("ping_serv_status"); ok || d.HasChange("ping_serv_status") {
		t, err := expandObjectFspVlanInterfacePingServStatus2edl(d, v, "ping_serv_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ping-serv-status"] = t
		}
	}

	if v, ok := d.GetOk("poe"); ok || d.HasChange("poe") {
		t, err := expandObjectFspVlanInterfacePoe2edl(d, v, "poe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe"] = t
		}
	}

	if v, ok := d.GetOk("polling_interval"); ok || d.HasChange("polling_interval") {
		t, err := expandObjectFspVlanInterfacePollingInterval2edl(d, v, "polling_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polling-interval"] = t
		}
	}

	if v, ok := d.GetOk("pppoe_unnumbered_negotiate"); ok || d.HasChange("pppoe_unnumbered_negotiate") {
		t, err := expandObjectFspVlanInterfacePppoeUnnumberedNegotiate2edl(d, v, "pppoe_unnumbered_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pppoe-unnumbered-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("pptp_auth_type"); ok || d.HasChange("pptp_auth_type") {
		t, err := expandObjectFspVlanInterfacePptpAuthType2edl(d, v, "pptp_auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-auth-type"] = t
		}
	}

	if v, ok := d.GetOk("pptp_client"); ok || d.HasChange("pptp_client") {
		t, err := expandObjectFspVlanInterfacePptpClient2edl(d, v, "pptp_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-client"] = t
		}
	}

	if v, ok := d.GetOk("pptp_password"); ok || d.HasChange("pptp_password") {
		t, err := expandObjectFspVlanInterfacePptpPassword2edl(d, v, "pptp_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-password"] = t
		}
	}

	if v, ok := d.GetOk("pptp_server_ip"); ok || d.HasChange("pptp_server_ip") {
		t, err := expandObjectFspVlanInterfacePptpServerIp2edl(d, v, "pptp_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("pptp_timeout"); ok || d.HasChange("pptp_timeout") {
		t, err := expandObjectFspVlanInterfacePptpTimeout2edl(d, v, "pptp_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-timeout"] = t
		}
	}

	if v, ok := d.GetOk("pptp_user"); ok || d.HasChange("pptp_user") {
		t, err := expandObjectFspVlanInterfacePptpUser2edl(d, v, "pptp_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-user"] = t
		}
	}

	if v, ok := d.GetOk("preserve_session_route"); ok || d.HasChange("preserve_session_route") {
		t, err := expandObjectFspVlanInterfacePreserveSessionRoute2edl(d, v, "preserve_session_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preserve-session-route"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandObjectFspVlanInterfacePriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("priority_override"); ok || d.HasChange("priority_override") {
		t, err := expandObjectFspVlanInterfacePriorityOverride2edl(d, v, "priority_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-override"] = t
		}
	}

	if v, ok := d.GetOk("proxy_captive_portal"); ok || d.HasChange("proxy_captive_portal") {
		t, err := expandObjectFspVlanInterfaceProxyCaptivePortal2edl(d, v, "proxy_captive_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("pvc_atm_qos"); ok || d.HasChange("pvc_atm_qos") {
		t, err := expandObjectFspVlanInterfacePvcAtmQos2edl(d, v, "pvc_atm_qos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvc-atm-qos"] = t
		}
	}

	if v, ok := d.GetOk("pvc_chan"); ok || d.HasChange("pvc_chan") {
		t, err := expandObjectFspVlanInterfacePvcChan2edl(d, v, "pvc_chan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvc-chan"] = t
		}
	}

	if v, ok := d.GetOk("pvc_crc"); ok || d.HasChange("pvc_crc") {
		t, err := expandObjectFspVlanInterfacePvcCrc2edl(d, v, "pvc_crc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvc-crc"] = t
		}
	}

	if v, ok := d.GetOk("pvc_pcr"); ok || d.HasChange("pvc_pcr") {
		t, err := expandObjectFspVlanInterfacePvcPcr2edl(d, v, "pvc_pcr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvc-pcr"] = t
		}
	}

	if v, ok := d.GetOk("pvc_scr"); ok || d.HasChange("pvc_scr") {
		t, err := expandObjectFspVlanInterfacePvcScr2edl(d, v, "pvc_scr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvc-scr"] = t
		}
	}

	if v, ok := d.GetOk("pvc_vlan_id"); ok || d.HasChange("pvc_vlan_id") {
		t, err := expandObjectFspVlanInterfacePvcVlanId2edl(d, v, "pvc_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvc-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("pvc_vlan_rx_id"); ok || d.HasChange("pvc_vlan_rx_id") {
		t, err := expandObjectFspVlanInterfacePvcVlanRxId2edl(d, v, "pvc_vlan_rx_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvc-vlan-rx-id"] = t
		}
	}

	if v, ok := d.GetOk("pvc_vlan_rx_op"); ok || d.HasChange("pvc_vlan_rx_op") {
		t, err := expandObjectFspVlanInterfacePvcVlanRxOp2edl(d, v, "pvc_vlan_rx_op")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvc-vlan-rx-op"] = t
		}
	}

	if v, ok := d.GetOk("pvc_vlan_tx_id"); ok || d.HasChange("pvc_vlan_tx_id") {
		t, err := expandObjectFspVlanInterfacePvcVlanTxId2edl(d, v, "pvc_vlan_tx_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvc-vlan-tx-id"] = t
		}
	}

	if v, ok := d.GetOk("pvc_vlan_tx_op"); ok || d.HasChange("pvc_vlan_tx_op") {
		t, err := expandObjectFspVlanInterfacePvcVlanTxOp2edl(d, v, "pvc_vlan_tx_op")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvc-vlan-tx-op"] = t
		}
	}

	if v, ok := d.GetOk("reachable_time"); ok || d.HasChange("reachable_time") {
		t, err := expandObjectFspVlanInterfaceReachableTime2edl(d, v, "reachable_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reachable-time"] = t
		}
	}

	if v, ok := d.GetOk("redundant_interface"); ok || d.HasChange("redundant_interface") {
		t, err := expandObjectFspVlanInterfaceRedundantInterface2edl(d, v, "redundant_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-interface"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip"); ok || d.HasChange("remote_ip") {
		t, err := expandObjectFspVlanInterfaceRemoteIp2edl(d, v, "remote_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_override_group"); ok || d.HasChange("replacemsg_override_group") {
		t, err := expandObjectFspVlanInterfaceReplacemsgOverrideGroup2edl(d, v, "replacemsg_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-override-group"] = t
		}
	}

	if v, ok := d.GetOk("retransmission"); ok || d.HasChange("retransmission") {
		t, err := expandObjectFspVlanInterfaceRetransmission2edl(d, v, "retransmission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retransmission"] = t
		}
	}

	if v, ok := d.GetOk("ring_rx"); ok || d.HasChange("ring_rx") {
		t, err := expandObjectFspVlanInterfaceRingRx2edl(d, v, "ring_rx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ring-rx"] = t
		}
	}

	if v, ok := d.GetOk("ring_tx"); ok || d.HasChange("ring_tx") {
		t, err := expandObjectFspVlanInterfaceRingTx2edl(d, v, "ring_tx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ring-tx"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandObjectFspVlanInterfaceRole2edl(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("sample_direction"); ok || d.HasChange("sample_direction") {
		t, err := expandObjectFspVlanInterfaceSampleDirection2edl(d, v, "sample_direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sample-direction"] = t
		}
	}

	if v, ok := d.GetOk("sample_rate"); ok || d.HasChange("sample_rate") {
		t, err := expandObjectFspVlanInterfaceSampleRate2edl(d, v, "sample_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sample-rate"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok || d.HasChange("scan_botnet_connections") {
		t, err := expandObjectFspVlanInterfaceScanBotnetConnections2edl(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("secondary_ip"); ok || d.HasChange("secondary_ip") {
		t, err := expandObjectFspVlanInterfaceSecondaryIp2edl(d, v, "secondary_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-IP"] = t
		}
	}

	if v, ok := d.GetOk("secondaryip"); ok || d.HasChange("secondaryip") {
		t, err := expandObjectFspVlanInterfaceSecondaryip2edl(d, v, "secondaryip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondaryip"] = t
		}
	}

	if v, ok := d.GetOk("security_8021x_dynamic_vlan_id"); ok || d.HasChange("security_8021x_dynamic_vlan_id") {
		t, err := expandObjectFspVlanInterfaceSecurity8021XDynamicVlanId2edl(d, v, "security_8021x_dynamic_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-8021x-dynamic-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("security_8021x_master"); ok || d.HasChange("security_8021x_master") {
		t, err := expandObjectFspVlanInterfaceSecurity8021XMaster2edl(d, v, "security_8021x_master")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-8021x-master"] = t
		}
	}

	if v, ok := d.GetOk("security_8021x_mode"); ok || d.HasChange("security_8021x_mode") {
		t, err := expandObjectFspVlanInterfaceSecurity8021XMode2edl(d, v, "security_8021x_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-8021x-mode"] = t
		}
	}

	if v, ok := d.GetOk("security_exempt_list"); ok || d.HasChange("security_exempt_list") {
		t, err := expandObjectFspVlanInterfaceSecurityExemptList2edl(d, v, "security_exempt_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-exempt-list"] = t
		}
	}

	if v, ok := d.GetOk("security_external_logout"); ok || d.HasChange("security_external_logout") {
		t, err := expandObjectFspVlanInterfaceSecurityExternalLogout2edl(d, v, "security_external_logout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-external-logout"] = t
		}
	}

	if v, ok := d.GetOk("security_external_web"); ok || d.HasChange("security_external_web") {
		t, err := expandObjectFspVlanInterfaceSecurityExternalWeb2edl(d, v, "security_external_web")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-external-web"] = t
		}
	}

	if v, ok := d.GetOk("security_groups"); ok || d.HasChange("security_groups") {
		t, err := expandObjectFspVlanInterfaceSecurityGroups2edl(d, v, "security_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-groups"] = t
		}
	}

	if v, ok := d.GetOk("security_mac_auth_bypass"); ok || d.HasChange("security_mac_auth_bypass") {
		t, err := expandObjectFspVlanInterfaceSecurityMacAuthBypass2edl(d, v, "security_mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok || d.HasChange("security_mode") {
		t, err := expandObjectFspVlanInterfaceSecurityMode2edl(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("security_redirect_url"); ok || d.HasChange("security_redirect_url") {
		t, err := expandObjectFspVlanInterfaceSecurityRedirectUrl2edl(d, v, "security_redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("select_profile_30a_35b"); ok || d.HasChange("select_profile_30a_35b") {
		t, err := expandObjectFspVlanInterfaceSelectProfile30A35B2edl(d, v, "select_profile_30a_35b")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["select-profile-30a-35b"] = t
		}
	}

	if v, ok := d.GetOk("service_name"); ok || d.HasChange("service_name") {
		t, err := expandObjectFspVlanInterfaceServiceName2edl(d, v, "service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-name"] = t
		}
	}

	if v, ok := d.GetOk("sflow_sampler"); ok || d.HasChange("sflow_sampler") {
		t, err := expandObjectFspVlanInterfaceSflowSampler2edl(d, v, "sflow_sampler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sflow-sampler"] = t
		}
	}

	if v, ok := d.GetOk("sfp_dsl"); ok || d.HasChange("sfp_dsl") {
		t, err := expandObjectFspVlanInterfaceSfpDsl2edl(d, v, "sfp_dsl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sfp-dsl"] = t
		}
	}

	if v, ok := d.GetOk("sfp_dsl_adsl_fallback"); ok || d.HasChange("sfp_dsl_adsl_fallback") {
		t, err := expandObjectFspVlanInterfaceSfpDslAdslFallback2edl(d, v, "sfp_dsl_adsl_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sfp-dsl-adsl-fallback"] = t
		}
	}

	if v, ok := d.GetOk("sfp_dsl_autodetect"); ok || d.HasChange("sfp_dsl_autodetect") {
		t, err := expandObjectFspVlanInterfaceSfpDslAutodetect2edl(d, v, "sfp_dsl_autodetect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sfp-dsl-autodetect"] = t
		}
	}

	if v, ok := d.GetOk("sfp_dsl_mac"); ok || d.HasChange("sfp_dsl_mac") {
		t, err := expandObjectFspVlanInterfaceSfpDslMac2edl(d, v, "sfp_dsl_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sfp-dsl-mac"] = t
		}
	}

	if v, ok := d.GetOk("speed"); ok || d.HasChange("speed") {
		t, err := expandObjectFspVlanInterfaceSpeed2edl(d, v, "speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speed"] = t
		}
	}

	if v, ok := d.GetOk("spillover_threshold"); ok || d.HasChange("spillover_threshold") {
		t, err := expandObjectFspVlanInterfaceSpilloverThreshold2edl(d, v, "spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("src_check"); ok || d.HasChange("src_check") {
		t, err := expandObjectFspVlanInterfaceSrcCheck2edl(d, v, "src_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-check"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFspVlanInterfaceStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("stp"); ok || d.HasChange("stp") {
		t, err := expandObjectFspVlanInterfaceStp2edl(d, v, "stp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp"] = t
		}
	}

	if v, ok := d.GetOk("stp_ha_secondary"); ok || d.HasChange("stp_ha_secondary") {
		t, err := expandObjectFspVlanInterfaceStpHaSecondary2edl(d, v, "stp_ha_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-ha-secondary"] = t
		}
	}

	if v, ok := d.GetOk("stp_ha_slave"); ok || d.HasChange("stp_ha_slave") {
		t, err := expandObjectFspVlanInterfaceStpHaSlave2edl(d, v, "stp_ha_slave")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-ha-slave"] = t
		}
	}

	if v, ok := d.GetOk("stpforward"); ok || d.HasChange("stpforward") {
		t, err := expandObjectFspVlanInterfaceStpforward2edl(d, v, "stpforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stpforward"] = t
		}
	}

	if v, ok := d.GetOk("stpforward_mode"); ok || d.HasChange("stpforward_mode") {
		t, err := expandObjectFspVlanInterfaceStpforwardMode2edl(d, v, "stpforward_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stpforward-mode"] = t
		}
	}

	if v, ok := d.GetOk("strip_priority_vlan_tag"); ok || d.HasChange("strip_priority_vlan_tag") {
		t, err := expandObjectFspVlanInterfaceStripPriorityVlanTag2edl(d, v, "strip_priority_vlan_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-priority-vlan-tag"] = t
		}
	}

	if v, ok := d.GetOk("subst"); ok || d.HasChange("subst") {
		t, err := expandObjectFspVlanInterfaceSubst2edl(d, v, "subst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subst"] = t
		}
	}

	if v, ok := d.GetOk("substitute_dst_mac"); ok || d.HasChange("substitute_dst_mac") {
		t, err := expandObjectFspVlanInterfaceSubstituteDstMac2edl(d, v, "substitute_dst_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["substitute-dst-mac"] = t
		}
	}

	if v, ok := d.GetOk("sw_algorithm"); ok || d.HasChange("sw_algorithm") {
		t, err := expandObjectFspVlanInterfaceSwAlgorithm2edl(d, v, "sw_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("swc_first_create"); ok || d.HasChange("swc_first_create") {
		t, err := expandObjectFspVlanInterfaceSwcFirstCreate2edl(d, v, "swc_first_create")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["swc-first-create"] = t
		}
	}

	if v, ok := d.GetOk("swc_vlan"); ok || d.HasChange("swc_vlan") {
		t, err := expandObjectFspVlanInterfaceSwcVlan2edl(d, v, "swc_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["swc-vlan"] = t
		}
	}

	if v, ok := d.GetOk("switch"); ok || d.HasChange("switch") {
		t, err := expandObjectFspVlanInterfaceSwitch2edl(d, v, "switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_access_vlan"); ok || d.HasChange("switch_controller_access_vlan") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerAccessVlan2edl(d, v, "switch_controller_access_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-access-vlan"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_arp_inspection"); ok || d.HasChange("switch_controller_arp_inspection") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerArpInspection2edl(d, v, "switch_controller_arp_inspection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-arp-inspection"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_auth"); ok || d.HasChange("switch_controller_auth") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerAuth2edl(d, v, "switch_controller_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-auth"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_dhcp_snooping"); ok || d.HasChange("switch_controller_dhcp_snooping") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerDhcpSnooping2edl(d, v, "switch_controller_dhcp_snooping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-dhcp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_dhcp_snooping_option82"); ok || d.HasChange("switch_controller_dhcp_snooping_option82") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerDhcpSnoopingOption822edl(d, v, "switch_controller_dhcp_snooping_option82")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-dhcp-snooping-option82"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_dhcp_snooping_verify_mac"); ok || d.HasChange("switch_controller_dhcp_snooping_verify_mac") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerDhcpSnoopingVerifyMac2edl(d, v, "switch_controller_dhcp_snooping_verify_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-dhcp-snooping-verify-mac"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_dynamic"); ok || d.HasChange("switch_controller_dynamic") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerDynamic2edl(d, v, "switch_controller_dynamic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_feature"); ok || d.HasChange("switch_controller_feature") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerFeature2edl(d, v, "switch_controller_feature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-feature"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_igmp_snooping"); ok || d.HasChange("switch_controller_igmp_snooping") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerIgmpSnooping2edl(d, v, "switch_controller_igmp_snooping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-igmp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_igmp_snooping_fast_leave"); ok || d.HasChange("switch_controller_igmp_snooping_fast_leave") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerIgmpSnoopingFastLeave2edl(d, v, "switch_controller_igmp_snooping_fast_leave")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-igmp-snooping-fast-leave"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_igmp_snooping_proxy"); ok || d.HasChange("switch_controller_igmp_snooping_proxy") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerIgmpSnoopingProxy2edl(d, v, "switch_controller_igmp_snooping_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-igmp-snooping-proxy"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_iot_scanning"); ok || d.HasChange("switch_controller_iot_scanning") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerIotScanning2edl(d, v, "switch_controller_iot_scanning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-iot-scanning"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_learning_limit"); ok || d.HasChange("switch_controller_learning_limit") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerLearningLimit2edl(d, v, "switch_controller_learning_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-learning-limit"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_mgmt_vlan"); ok || d.HasChange("switch_controller_mgmt_vlan") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerMgmtVlan2edl(d, v, "switch_controller_mgmt_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-mgmt-vlan"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_nac"); ok || d.HasChange("switch_controller_nac") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerNac2edl(d, v, "switch_controller_nac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-nac"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_netflow_collect"); ok || d.HasChange("switch_controller_netflow_collect") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerNetflowCollect2edl(d, v, "switch_controller_netflow_collect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-netflow-collect"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_offload"); ok || d.HasChange("switch_controller_offload") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerOffload2edl(d, v, "switch_controller_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-offload"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_offload_gw"); ok || d.HasChange("switch_controller_offload_gw") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerOffloadGw2edl(d, v, "switch_controller_offload_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-offload-gw"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_offload_ip"); ok || d.HasChange("switch_controller_offload_ip") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerOffloadIp2edl(d, v, "switch_controller_offload_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-offload-ip"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_offloading"); ok || d.HasChange("switch_controller_offloading") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerOffloading2edl(d, v, "switch_controller_offloading")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-offloading"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_offloading_gw"); ok || d.HasChange("switch_controller_offloading_gw") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerOffloadingGw2edl(d, v, "switch_controller_offloading_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-offloading-gw"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_offloading_ip"); ok || d.HasChange("switch_controller_offloading_ip") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerOffloadingIp2edl(d, v, "switch_controller_offloading_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-offloading-ip"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_radius_server"); ok || d.HasChange("switch_controller_radius_server") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerRadiusServer2edl(d, v, "switch_controller_radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_rspan_mode"); ok || d.HasChange("switch_controller_rspan_mode") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerRspanMode2edl(d, v, "switch_controller_rspan_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-rspan-mode"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_source_ip"); ok || d.HasChange("switch_controller_source_ip") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerSourceIp2edl(d, v, "switch_controller_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_traffic_policy"); ok || d.HasChange("switch_controller_traffic_policy") {
		t, err := expandObjectFspVlanInterfaceSwitchControllerTrafficPolicy2edl(d, v, "switch_controller_traffic_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-traffic-policy"] = t
		}
	}

	if v, ok := d.GetOk("system_id"); ok || d.HasChange("system_id") {
		t, err := expandObjectFspVlanInterfaceSystemId2edl(d, v, "system_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-id"] = t
		}
	}

	if v, ok := d.GetOk("system_id_type"); ok || d.HasChange("system_id_type") {
		t, err := expandObjectFspVlanInterfaceSystemIdType2edl(d, v, "system_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-id-type"] = t
		}
	}

	if v, ok := d.GetOk("tc_mode"); ok || d.HasChange("tc_mode") {
		t, err := expandObjectFspVlanInterfaceTcMode2edl(d, v, "tc_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tc-mode"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss"); ok || d.HasChange("tcp_mss") {
		t, err := expandObjectFspVlanInterfaceTcpMss2edl(d, v, "tcp_mss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss"] = t
		}
	}

	if v, ok := d.GetOk("trunk"); ok || d.HasChange("trunk") {
		t, err := expandObjectFspVlanInterfaceTrunk2edl(d, v, "trunk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trunk"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_1"); ok || d.HasChange("trust_ip_1") {
		t, err := expandObjectFspVlanInterfaceTrustIp12edl(d, v, "trust_ip_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-1"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_2"); ok || d.HasChange("trust_ip_2") {
		t, err := expandObjectFspVlanInterfaceTrustIp22edl(d, v, "trust_ip_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-2"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_3"); ok || d.HasChange("trust_ip_3") {
		t, err := expandObjectFspVlanInterfaceTrustIp32edl(d, v, "trust_ip_3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-3"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip6_1"); ok || d.HasChange("trust_ip6_1") {
		t, err := expandObjectFspVlanInterfaceTrustIp612edl(d, v, "trust_ip6_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip6-1"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip6_2"); ok || d.HasChange("trust_ip6_2") {
		t, err := expandObjectFspVlanInterfaceTrustIp622edl(d, v, "trust_ip6_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip6-2"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip6_3"); ok || d.HasChange("trust_ip6_3") {
		t, err := expandObjectFspVlanInterfaceTrustIp632edl(d, v, "trust_ip6_3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip6-3"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFspVlanInterfaceType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandObjectFspVlanInterfaceUsername2edl(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("vci"); ok || d.HasChange("vci") {
		t, err := expandObjectFspVlanInterfaceVci2edl(d, v, "vci")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci"] = t
		}
	}

	if v, ok := d.GetOk("vectoring"); ok || d.HasChange("vectoring") {
		t, err := expandObjectFspVlanInterfaceVectoring2edl(d, v, "vectoring")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vectoring"] = t
		}
	}

	if v, ok := d.GetOk("vindex"); ok || d.HasChange("vindex") {
		t, err := expandObjectFspVlanInterfaceVindex2edl(d, v, "vindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vindex"] = t
		}
	}

	if v, ok := d.GetOk("vlan_protocol"); ok || d.HasChange("vlan_protocol") {
		t, err := expandObjectFspVlanInterfaceVlanProtocol2edl(d, v, "vlan_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-protocol"] = t
		}
	}

	if v, ok := d.GetOk("vlanforward"); ok || d.HasChange("vlanforward") {
		t, err := expandObjectFspVlanInterfaceVlanforward2edl(d, v, "vlanforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanforward"] = t
		}
	}

	if v, ok := d.GetOk("vlanid"); ok || d.HasChange("vlanid") {
		t, err := expandObjectFspVlanInterfaceVlanid2edl(d, v, "vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	if v, ok := d.GetOk("vpi"); ok || d.HasChange("vpi") {
		t, err := expandObjectFspVlanInterfaceVpi2edl(d, v, "vpi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpi"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandObjectFspVlanInterfaceVrf2edl(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	if v, ok := d.GetOk("vrrp"); ok || d.HasChange("vrrp") {
		t, err := expandObjectFspVlanInterfaceVrrp2edl(d, v, "vrrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp"] = t
		}
	}

	if v, ok := d.GetOk("vrrp_virtual_mac"); ok || d.HasChange("vrrp_virtual_mac") {
		t, err := expandObjectFspVlanInterfaceVrrpVirtualMac2edl(d, v, "vrrp_virtual_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp-virtual-mac"] = t
		}
	}

	if v, ok := d.GetOk("wccp"); ok || d.HasChange("wccp") {
		t, err := expandObjectFspVlanInterfaceWccp2edl(d, v, "wccp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wccp"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectFspVlanInterfaceWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("wifi_5g_threshold"); ok || d.HasChange("wifi_5g_threshold") {
		t, err := expandObjectFspVlanInterfaceWifi5GThreshold2edl(d, v, "wifi_5g_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-5g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("wifi_acl"); ok || d.HasChange("wifi_acl") {
		t, err := expandObjectFspVlanInterfaceWifiAcl2edl(d, v, "wifi_acl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-acl"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ap_band"); ok || d.HasChange("wifi_ap_band") {
		t, err := expandObjectFspVlanInterfaceWifiApBand2edl(d, v, "wifi_ap_band")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ap-band"] = t
		}
	}

	if v, ok := d.GetOk("wifi_auth"); ok || d.HasChange("wifi_auth") {
		t, err := expandObjectFspVlanInterfaceWifiAuth2edl(d, v, "wifi_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-auth"] = t
		}
	}

	if v, ok := d.GetOk("wifi_auto_connect"); ok || d.HasChange("wifi_auto_connect") {
		t, err := expandObjectFspVlanInterfaceWifiAutoConnect2edl(d, v, "wifi_auto_connect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-auto-connect"] = t
		}
	}

	if v, ok := d.GetOk("wifi_auto_save"); ok || d.HasChange("wifi_auto_save") {
		t, err := expandObjectFspVlanInterfaceWifiAutoSave2edl(d, v, "wifi_auto_save")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-auto-save"] = t
		}
	}

	if v, ok := d.GetOk("wifi_broadcast_ssid"); ok || d.HasChange("wifi_broadcast_ssid") {
		t, err := expandObjectFspVlanInterfaceWifiBroadcastSsid2edl(d, v, "wifi_broadcast_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-broadcast-ssid"] = t
		}
	}

	if v, ok := d.GetOk("wifi_dns_server1"); ok || d.HasChange("wifi_dns_server1") {
		t, err := expandObjectFspVlanInterfaceWifiDnsServer12edl(d, v, "wifi_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("wifi_dns_server2"); ok || d.HasChange("wifi_dns_server2") {
		t, err := expandObjectFspVlanInterfaceWifiDnsServer22edl(d, v, "wifi_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("wifi_encrypt"); ok || d.HasChange("wifi_encrypt") {
		t, err := expandObjectFspVlanInterfaceWifiEncrypt2edl(d, v, "wifi_encrypt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-encrypt"] = t
		}
	}

	if v, ok := d.GetOk("wifi_fragment_threshold"); ok || d.HasChange("wifi_fragment_threshold") {
		t, err := expandObjectFspVlanInterfaceWifiFragmentThreshold2edl(d, v, "wifi_fragment_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-fragment-threshold"] = t
		}
	}

	if v, ok := d.GetOk("wifi_gateway"); ok || d.HasChange("wifi_gateway") {
		t, err := expandObjectFspVlanInterfaceWifiGateway2edl(d, v, "wifi_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-gateway"] = t
		}
	}

	if v, ok := d.GetOk("wifi_key"); ok || d.HasChange("wifi_key") {
		t, err := expandObjectFspVlanInterfaceWifiKey2edl(d, v, "wifi_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-key"] = t
		}
	}

	if v, ok := d.GetOk("wifi_keyindex"); ok || d.HasChange("wifi_keyindex") {
		t, err := expandObjectFspVlanInterfaceWifiKeyindex2edl(d, v, "wifi_keyindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-keyindex"] = t
		}
	}

	if v, ok := d.GetOk("wifi_mac_filter"); ok || d.HasChange("wifi_mac_filter") {
		t, err := expandObjectFspVlanInterfaceWifiMacFilter2edl(d, v, "wifi_mac_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-mac-filter"] = t
		}
	}

	if v, ok := d.GetOk("wifi_passphrase"); ok || d.HasChange("wifi_passphrase") {
		t, err := expandObjectFspVlanInterfaceWifiPassphrase2edl(d, v, "wifi_passphrase")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-passphrase"] = t
		}
	}

	if v, ok := d.GetOk("wifi_radius_server"); ok || d.HasChange("wifi_radius_server") {
		t, err := expandObjectFspVlanInterfaceWifiRadiusServer2edl(d, v, "wifi_radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("wifi_rts_threshold"); ok || d.HasChange("wifi_rts_threshold") {
		t, err := expandObjectFspVlanInterfaceWifiRtsThreshold2edl(d, v, "wifi_rts_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-rts-threshold"] = t
		}
	}

	if v, ok := d.GetOk("wifi_security"); ok || d.HasChange("wifi_security") {
		t, err := expandObjectFspVlanInterfaceWifiSecurity2edl(d, v, "wifi_security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-security"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ssid"); ok || d.HasChange("wifi_ssid") {
		t, err := expandObjectFspVlanInterfaceWifiSsid2edl(d, v, "wifi_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ssid"] = t
		}
	}

	if v, ok := d.GetOk("wifi_usergroup"); ok || d.HasChange("wifi_usergroup") {
		t, err := expandObjectFspVlanInterfaceWifiUsergroup2edl(d, v, "wifi_usergroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-usergroup"] = t
		}
	}

	if v, ok := d.GetOk("wins_ip"); ok || d.HasChange("wins_ip") {
		t, err := expandObjectFspVlanInterfaceWinsIp2edl(d, v, "wins_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-ip"] = t
		}
	}

	return &obj, nil
}
