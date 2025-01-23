// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFsp VlanDynamicMapping

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanDynamicMappingCreate,
		Read:   resourceObjectFspVlanDynamicMappingRead,
		Update: resourceObjectFspVlanDynamicMappingUpdate,
		Delete: resourceObjectFspVlanDynamicMappingDelete,

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
			"_dhcp_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"_scope": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dhcp_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_configuration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_managed_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"conflicted_ip_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ddns_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ddns_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ddns_keyname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ddns_server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ddns_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ddns_update": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ddns_update_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ddns_zone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"default_gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp_settings_from_fortiipam": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_server1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_server2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_server3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_server4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"exclude_range": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"lease_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"uci_match": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"uci_string": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"vci_match": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vci_string": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"filename": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"forticlient_on_net_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_range": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"lease_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"uci_match": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"uci_string": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"vci_match": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vci_string": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"ipsec_lease_hold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lease_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mac_acl_default_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"netmask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"next_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ntp_server1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ntp_server2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ntp_server3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ntp_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"option1": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"option2": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"option3": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"option4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"option5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"option6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"uci_match": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"uci_string": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vci_match": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vci_string": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"relay_agent": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"reserved_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"circuit_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"circuit_id_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"remote_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_id_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"server_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"shared_subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tftp_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"timezone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"timezone_option": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"wifi_ac_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wifi_ac1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wifi_ac2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wifi_ac3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wins_server1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wins_server2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dhcp_relay_agent_option": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_relay_interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_relay_ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dhcp_relay_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_relay_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip": &schema.Schema{
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
									},
									"dhcp6_prefix_delegation": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dhcp6_prefix_hint": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dhcp6_prefix_hint_plt": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"dhcp6_prefix_hint_vlt": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
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
									},
									"dhcp6_relay_source_interface": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dhcp6_relay_source_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dhcp6_relay_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"icmp6_send_redirect": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"interface_identifier": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ip6_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
									"ip6_max_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip6_min_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip6_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ip6_other_flag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
									"ip6_subnet": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
									"unique_autoconf_addr": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vrip6_link_local": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vrrp_virtual_mac6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vrrp6": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accept_mode": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"adv_interval": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"ignore_default_route": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"preempt": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"priority": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"status": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
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
												},
											},
										},
									},
								},
							},
						},
						"secondary_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"vlanid": &schema.Schema{
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
									},
									"adv_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ignore_default_route": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"preempt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"priority": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
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
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
								},
							},
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectFspVlanDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFspVlanDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanDynamicMapping resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFspVlanDynamicMapping(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanDynamicMapping resource: %v", err)
	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectFspVlanDynamicMappingRead(d, m)
}

func resourceObjectFspVlanDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFspVlanDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDynamicMapping resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFspVlanDynamicMapping(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectFspVlanDynamicMappingRead(d, m)
}

func resourceObjectFspVlanDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFspVlanDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["vlan"] = vlan

	o, err := c.ReadObjectFspVlanDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanDynamicMappingDhcpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFspVlanDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFspVlanDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServer2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_configuration"
	if _, ok := i["auto-configuration"]; ok {
		result["auto_configuration"] = flattenObjectFspVlanDynamicMappingDhcpServerAutoConfiguration2edl(i["auto-configuration"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_managed_status"
	if _, ok := i["auto-managed-status"]; ok {
		result["auto_managed_status"] = flattenObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus2edl(i["auto-managed-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "conflicted_ip_timeout"
	if _, ok := i["conflicted-ip-timeout"]; ok {
		result["conflicted_ip_timeout"] = flattenObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout2edl(i["conflicted-ip-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_auth"
	if _, ok := i["ddns-auth"]; ok {
		result["ddns_auth"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsAuth2edl(i["ddns-auth"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_key"
	if _, ok := i["ddns-key"]; ok {
		result["ddns_key"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsKey2edl(i["ddns-key"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_keyname"
	if _, ok := i["ddns-keyname"]; ok {
		result["ddns_keyname"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsKeyname2edl(i["ddns-keyname"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_server_ip"
	if _, ok := i["ddns-server-ip"]; ok {
		result["ddns_server_ip"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsServerIp2edl(i["ddns-server-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_ttl"
	if _, ok := i["ddns-ttl"]; ok {
		result["ddns_ttl"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsTtl2edl(i["ddns-ttl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_update"
	if _, ok := i["ddns-update"]; ok {
		result["ddns_update"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdate2edl(i["ddns-update"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_update_override"
	if _, ok := i["ddns-update-override"]; ok {
		result["ddns_update_override"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride2edl(i["ddns-update-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_zone"
	if _, ok := i["ddns-zone"]; ok {
		result["ddns_zone"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsZone2edl(i["ddns-zone"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_gateway"
	if _, ok := i["default-gateway"]; ok {
		result["default_gateway"] = flattenObjectFspVlanDynamicMappingDhcpServerDefaultGateway2edl(i["default-gateway"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_settings_from_fortiipam"
	if _, ok := i["dhcp-settings-from-fortiipam"]; ok {
		result["dhcp_settings_from_fortiipam"] = flattenObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam2edl(i["dhcp-settings-from-fortiipam"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server1"
	if _, ok := i["dns-server1"]; ok {
		result["dns_server1"] = flattenObjectFspVlanDynamicMappingDhcpServerDnsServer12edl(i["dns-server1"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server2"
	if _, ok := i["dns-server2"]; ok {
		result["dns_server2"] = flattenObjectFspVlanDynamicMappingDhcpServerDnsServer22edl(i["dns-server2"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server3"
	if _, ok := i["dns-server3"]; ok {
		result["dns_server3"] = flattenObjectFspVlanDynamicMappingDhcpServerDnsServer32edl(i["dns-server3"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server4"
	if _, ok := i["dns-server4"]; ok {
		result["dns_server4"] = flattenObjectFspVlanDynamicMappingDhcpServerDnsServer42edl(i["dns-server4"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_service"
	if _, ok := i["dns-service"]; ok {
		result["dns_service"] = flattenObjectFspVlanDynamicMappingDhcpServerDnsService2edl(i["dns-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "domain"
	if _, ok := i["domain"]; ok {
		result["domain"] = flattenObjectFspVlanDynamicMappingDhcpServerDomain2edl(i["domain"], d, pre_append)
	}

	pre_append = pre + ".0." + "enable"
	if _, ok := i["enable"]; ok {
		result["enable"] = flattenObjectFspVlanDynamicMappingDhcpServerEnable2edl(i["enable"], d, pre_append)
	}

	pre_append = pre + ".0." + "exclude_range"
	if _, ok := i["exclude-range"]; ok {
		result["exclude_range"] = flattenObjectFspVlanDynamicMappingDhcpServerExcludeRange2edl(i["exclude-range"], d, pre_append)
	}

	pre_append = pre + ".0." + "filename"
	if _, ok := i["filename"]; ok {
		result["filename"] = flattenObjectFspVlanDynamicMappingDhcpServerFilename2edl(i["filename"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_on_net_status"
	if _, ok := i["forticlient-on-net-status"]; ok {
		result["forticlient_on_net_status"] = flattenObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus2edl(i["forticlient-on-net-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenObjectFspVlanDynamicMappingDhcpServerId2edl(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_mode"
	if _, ok := i["ip-mode"]; ok {
		result["ip_mode"] = flattenObjectFspVlanDynamicMappingDhcpServerIpMode2edl(i["ip-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_range"
	if _, ok := i["ip-range"]; ok {
		result["ip_range"] = flattenObjectFspVlanDynamicMappingDhcpServerIpRange2edl(i["ip-range"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipsec_lease_hold"
	if _, ok := i["ipsec-lease-hold"]; ok {
		result["ipsec_lease_hold"] = flattenObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold2edl(i["ipsec-lease-hold"], d, pre_append)
	}

	pre_append = pre + ".0." + "lease_time"
	if _, ok := i["lease-time"]; ok {
		result["lease_time"] = flattenObjectFspVlanDynamicMappingDhcpServerLeaseTime2edl(i["lease-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_acl_default_action"
	if _, ok := i["mac-acl-default-action"]; ok {
		result["mac_acl_default_action"] = flattenObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction2edl(i["mac-acl-default-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "netmask"
	if _, ok := i["netmask"]; ok {
		result["netmask"] = flattenObjectFspVlanDynamicMappingDhcpServerNetmask2edl(i["netmask"], d, pre_append)
	}

	pre_append = pre + ".0." + "next_server"
	if _, ok := i["next-server"]; ok {
		result["next_server"] = flattenObjectFspVlanDynamicMappingDhcpServerNextServer2edl(i["next-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_server1"
	if _, ok := i["ntp-server1"]; ok {
		result["ntp_server1"] = flattenObjectFspVlanDynamicMappingDhcpServerNtpServer12edl(i["ntp-server1"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_server2"
	if _, ok := i["ntp-server2"]; ok {
		result["ntp_server2"] = flattenObjectFspVlanDynamicMappingDhcpServerNtpServer22edl(i["ntp-server2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_server3"
	if _, ok := i["ntp-server3"]; ok {
		result["ntp_server3"] = flattenObjectFspVlanDynamicMappingDhcpServerNtpServer32edl(i["ntp-server3"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_service"
	if _, ok := i["ntp-service"]; ok {
		result["ntp_service"] = flattenObjectFspVlanDynamicMappingDhcpServerNtpService2edl(i["ntp-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "option1"
	if _, ok := i["option1"]; ok {
		result["option1"] = flattenObjectFspVlanDynamicMappingDhcpServerOption12edl(i["option1"], d, pre_append)
	}

	pre_append = pre + ".0." + "option2"
	if _, ok := i["option2"]; ok {
		result["option2"] = flattenObjectFspVlanDynamicMappingDhcpServerOption22edl(i["option2"], d, pre_append)
	}

	pre_append = pre + ".0." + "option3"
	if _, ok := i["option3"]; ok {
		result["option3"] = flattenObjectFspVlanDynamicMappingDhcpServerOption32edl(i["option3"], d, pre_append)
	}

	pre_append = pre + ".0." + "option4"
	if _, ok := i["option4"]; ok {
		result["option4"] = flattenObjectFspVlanDynamicMappingDhcpServerOption42edl(i["option4"], d, pre_append)
	}

	pre_append = pre + ".0." + "option5"
	if _, ok := i["option5"]; ok {
		result["option5"] = flattenObjectFspVlanDynamicMappingDhcpServerOption52edl(i["option5"], d, pre_append)
	}

	pre_append = pre + ".0." + "option6"
	if _, ok := i["option6"]; ok {
		result["option6"] = flattenObjectFspVlanDynamicMappingDhcpServerOption62edl(i["option6"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFspVlanDynamicMappingDhcpServerOptions2edl(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "relay_agent"
	if _, ok := i["relay-agent"]; ok {
		result["relay_agent"] = flattenObjectFspVlanDynamicMappingDhcpServerRelayAgent2edl(i["relay-agent"], d, pre_append)
	}

	pre_append = pre + ".0." + "reserved_address"
	if _, ok := i["reserved-address"]; ok {
		result["reserved_address"] = flattenObjectFspVlanDynamicMappingDhcpServerReservedAddress2edl(i["reserved-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_type"
	if _, ok := i["server-type"]; ok {
		result["server_type"] = flattenObjectFspVlanDynamicMappingDhcpServerServerType2edl(i["server-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "shared_subnet"
	if _, ok := i["shared-subnet"]; ok {
		result["shared_subnet"] = flattenObjectFspVlanDynamicMappingDhcpServerSharedSubnet2edl(i["shared-subnet"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFspVlanDynamicMappingDhcpServerStatus2edl(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tftp_server"
	if _, ok := i["tftp-server"]; ok {
		result["tftp_server"] = flattenObjectFspVlanDynamicMappingDhcpServerTftpServer2edl(i["tftp-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "timezone"
	if _, ok := i["timezone"]; ok {
		result["timezone"] = flattenObjectFspVlanDynamicMappingDhcpServerTimezone2edl(i["timezone"], d, pre_append)
	}

	pre_append = pre + ".0." + "timezone_option"
	if _, ok := i["timezone-option"]; ok {
		result["timezone_option"] = flattenObjectFspVlanDynamicMappingDhcpServerTimezoneOption2edl(i["timezone-option"], d, pre_append)
	}

	pre_append = pre + ".0." + "vci_match"
	if _, ok := i["vci-match"]; ok {
		result["vci_match"] = flattenObjectFspVlanDynamicMappingDhcpServerVciMatch2edl(i["vci-match"], d, pre_append)
	}

	pre_append = pre + ".0." + "vci_string"
	if _, ok := i["vci-string"]; ok {
		result["vci_string"] = flattenObjectFspVlanDynamicMappingDhcpServerVciString2edl(i["vci-string"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac_service"
	if _, ok := i["wifi-ac-service"]; ok {
		result["wifi_ac_service"] = flattenObjectFspVlanDynamicMappingDhcpServerWifiAcService2edl(i["wifi-ac-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac1"
	if _, ok := i["wifi-ac1"]; ok {
		result["wifi_ac1"] = flattenObjectFspVlanDynamicMappingDhcpServerWifiAc12edl(i["wifi-ac1"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac2"
	if _, ok := i["wifi-ac2"]; ok {
		result["wifi_ac2"] = flattenObjectFspVlanDynamicMappingDhcpServerWifiAc22edl(i["wifi-ac2"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac3"
	if _, ok := i["wifi-ac3"]; ok {
		result["wifi_ac3"] = flattenObjectFspVlanDynamicMappingDhcpServerWifiAc32edl(i["wifi-ac3"], d, pre_append)
	}

	pre_append = pre + ".0." + "wins_server1"
	if _, ok := i["wins-server1"]; ok {
		result["wins_server1"] = flattenObjectFspVlanDynamicMappingDhcpServerWinsServer12edl(i["wins-server1"], d, pre_append)
	}

	pre_append = pre + ".0." + "wins_server2"
	if _, ok := i["wins-server2"]; ok {
		result["wins_server2"] = flattenObjectFspVlanDynamicMappingDhcpServerWinsServer22edl(i["wins-server2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFspVlanDynamicMappingDhcpServerAutoConfiguration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsKeyname2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsZone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDefaultGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerEnable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRange2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp2edl(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := i["lease-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime2edl(i["lease-time"], d, pre_append)
			tmp["lease_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-LeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp2edl(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch2edl(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString2edl(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch2edl(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString2edl(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerFilename2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRange2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp2edl(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := i["lease-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeLeaseTime2edl(i["lease-time"], d, pre_append)
			tmp["lease_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-LeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp2edl(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeUciMatch2edl(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeUciString2edl(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch2edl(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciString2edl(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeUciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeUciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNetmask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNextServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpServer32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption52edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptions2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := i["code"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsCode2edl(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsUciMatch2edl(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsUciString2edl(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Value")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch2edl(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciString2edl(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsCode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsUciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsUciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerRelayAgent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddress2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId2edl(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := i["circuit-id-type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType2edl(i["circuit-id-type"], d, pre_append)
			tmp["circuit_id_type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-CircuitIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription2edl(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressMac2edl(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId2edl(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := i["remote-id-type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType2edl(i["remote-id-type"], d, pre_append)
			tmp["remote_id_type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-RemoteIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerServerType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerSharedSubnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerTftpServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerTimezone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerTimezoneOption2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAcService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAc12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAc22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAc32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWinsServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWinsServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterface2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dhcp_relay_agent_option"
	if _, ok := i["dhcp-relay-agent-option"]; ok {
		result["dhcp_relay_agent_option"] = flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayAgentOption2edl(i["dhcp-relay-agent-option"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_interface_select_method"
	if _, ok := i["dhcp-relay-interface-select-method"]; ok {
		result["dhcp_relay_interface_select_method"] = flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayInterfaceSelectMethod2edl(i["dhcp-relay-interface-select-method"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_ip"
	if _, ok := i["dhcp-relay-ip"]; ok {
		result["dhcp_relay_ip"] = flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayIp2edl(i["dhcp-relay-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_service"
	if _, ok := i["dhcp-relay-service"]; ok {
		result["dhcp_relay_service"] = flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayService2edl(i["dhcp-relay-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_type"
	if _, ok := i["dhcp-relay-type"]; ok {
		result["dhcp_relay_type"] = flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayType2edl(i["dhcp-relay-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenObjectFspVlanDynamicMappingInterfaceIp2edl(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6"
	if _, ok := i["ipv6"]; ok {
		result["ipv6"] = flattenObjectFspVlanDynamicMappingInterfaceIpv62edl(i["ipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "secondary_ip"
	if _, ok := i["secondary-IP"]; ok {
		result["secondary_ip"] = flattenObjectFspVlanDynamicMappingInterfaceSecondaryIp2edl(i["secondary-IP"], d, pre_append)
	}

	pre_append = pre + ".0." + "secondaryip"
	if _, ok := i["secondaryip"]; ok {
		result["secondaryip"] = flattenObjectFspVlanDynamicMappingInterfaceSecondaryip2edl(i["secondaryip"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlanid"
	if _, ok := i["vlanid"]; ok {
		result["vlanid"] = flattenObjectFspVlanDynamicMappingInterfaceVlanid2edl(i["vlanid"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp"
	if _, ok := i["vrrp"]; ok {
		result["vrrp"] = flattenObjectFspVlanDynamicMappingInterfaceVrrp2edl(i["vrrp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayAgentOption2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayInterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv62edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "autoconf"
	if _, ok := i["autoconf"]; ok {
		result["autoconf"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Autoconf2edl(i["autoconf"], d, pre_append)
	}

	pre_append = pre + ".0." + "cli_conn6_status"
	if _, ok := i["cli-conn6-status"]; ok {
		result["cli_conn6_status"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6CliConn6Status2edl(i["cli-conn6-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := i["dhcp6-client-options"]; ok {
		result["dhcp6_client_options"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6ClientOptions2edl(i["dhcp6-client-options"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := i["dhcp6-information-request"]; ok {
		result["dhcp6_information_request"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6InformationRequest2edl(i["dhcp6-information-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := i["dhcp6-prefix-delegation"]; ok {
		result["dhcp6_prefix_delegation"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixDelegation2edl(i["dhcp6-prefix-delegation"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := i["dhcp6-prefix-hint"]; ok {
		result["dhcp6_prefix_hint"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHint2edl(i["dhcp6-prefix-hint"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := i["dhcp6-prefix-hint-plt"]; ok {
		result["dhcp6_prefix_hint_plt"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintPlt2edl(i["dhcp6-prefix-hint-plt"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := i["dhcp6-prefix-hint-vlt"]; ok {
		result["dhcp6_prefix_hint_vlt"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintVlt2edl(i["dhcp6-prefix-hint-vlt"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_interface_id"
	if _, ok := i["dhcp6-relay-interface-id"]; ok {
		result["dhcp6_relay_interface_id"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayInterfaceId2edl(i["dhcp6-relay-interface-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := i["dhcp6-relay-ip"]; ok {
		result["dhcp6_relay_ip"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayIp2edl(i["dhcp6-relay-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := i["dhcp6-relay-service"]; ok {
		result["dhcp6_relay_service"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayService2edl(i["dhcp6-relay-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_source_interface"
	if _, ok := i["dhcp6-relay-source-interface"]; ok {
		result["dhcp6_relay_source_interface"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelaySourceInterface2edl(i["dhcp6-relay-source-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_source_ip"
	if _, ok := i["dhcp6-relay-source-ip"]; ok {
		result["dhcp6_relay_source_ip"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelaySourceIp2edl(i["dhcp6-relay-source-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := i["dhcp6-relay-type"]; ok {
		result["dhcp6_relay_type"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayType2edl(i["dhcp6-relay-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp6_send_redirect"
	if _, ok := i["icmp6-send-redirect"]; ok {
		result["icmp6_send_redirect"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Icmp6SendRedirect2edl(i["icmp6-send-redirect"], d, pre_append)
	}

	pre_append = pre + ".0." + "interface_identifier"
	if _, ok := i["interface-identifier"]; ok {
		result["interface_identifier"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6InterfaceIdentifier2edl(i["interface-identifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_address"
	if _, ok := i["ip6-address"]; ok {
		result["ip6_address"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Address2edl(i["ip6-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := i["ip6-allowaccess"]; ok {
		result["ip6_allowaccess"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Allowaccess2edl(i["ip6-allowaccess"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := i["ip6-default-life"]; ok {
		result["ip6_default_life"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DefaultLife2edl(i["ip6-default-life"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_iaid"
	if _, ok := i["ip6-delegated-prefix-iaid"]; ok {
		result["ip6_delegated_prefix_iaid"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixIaid2edl(i["ip6-delegated-prefix-iaid"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := i["ip6-delegated-prefix-list"]; ok {
		result["ip6_delegated_prefix_list"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixList2edl(i["ip6-delegated-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := i["ip6-dns-server-override"]; ok {
		result["ip6_dns_server_override"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DnsServerOverride2edl(i["ip6-dns-server-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := i["ip6-extra-addr"]; ok {
		result["ip6_extra_addr"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddr2edl(i["ip6-extra-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := i["ip6-hop-limit"]; ok {
		result["ip6_hop_limit"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6HopLimit2edl(i["ip6-hop-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := i["ip6-link-mtu"]; ok {
		result["ip6_link_mtu"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6LinkMtu2edl(i["ip6-link-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := i["ip6-manage-flag"]; ok {
		result["ip6_manage_flag"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ManageFlag2edl(i["ip6-manage-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := i["ip6-max-interval"]; ok {
		result["ip6_max_interval"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6MaxInterval2edl(i["ip6-max-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := i["ip6-min-interval"]; ok {
		result["ip6_min_interval"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6MinInterval2edl(i["ip6-min-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := i["ip6-mode"]; ok {
		result["ip6_mode"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Mode2edl(i["ip6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := i["ip6-other-flag"]; ok {
		result["ip6_other_flag"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6OtherFlag2edl(i["ip6-other-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := i["ip6-prefix-list"]; ok {
		result["ip6_prefix_list"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixList2edl(i["ip6-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_mode"
	if _, ok := i["ip6-prefix-mode"]; ok {
		result["ip6_prefix_mode"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixMode2edl(i["ip6-prefix-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := i["ip6-reachable-time"]; ok {
		result["ip6_reachable_time"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ReachableTime2edl(i["ip6-reachable-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := i["ip6-retrans-time"]; ok {
		result["ip6_retrans_time"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6RetransTime2edl(i["ip6-retrans-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := i["ip6-send-adv"]; ok {
		result["ip6_send_adv"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6SendAdv2edl(i["ip6-send-adv"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := i["ip6-subnet"]; ok {
		result["ip6_subnet"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Subnet2edl(i["ip6-subnet"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := i["ip6-upstream-interface"]; ok {
		result["ip6_upstream_interface"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6UpstreamInterface2edl(i["ip6-upstream-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cert"
	if _, ok := i["nd-cert"]; ok {
		result["nd_cert"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdCert2edl(i["nd-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := i["nd-cga-modifier"]; ok {
		result["nd_cga_modifier"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdCgaModifier2edl(i["nd-cga-modifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_mode"
	if _, ok := i["nd-mode"]; ok {
		result["nd_mode"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdMode2edl(i["nd-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := i["nd-security-level"]; ok {
		result["nd_security_level"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdSecurityLevel2edl(i["nd-security-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := i["nd-timestamp-delta"]; ok {
		result["nd_timestamp_delta"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampDelta2edl(i["nd-timestamp-delta"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := i["nd-timestamp-fuzz"]; ok {
		result["nd_timestamp_fuzz"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampFuzz2edl(i["nd-timestamp-fuzz"], d, pre_append)
	}

	pre_append = pre + ".0." + "ra_send_mtu"
	if _, ok := i["ra-send-mtu"]; ok {
		result["ra_send_mtu"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6RaSendMtu2edl(i["ra-send-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "unique_autoconf_addr"
	if _, ok := i["unique-autoconf-addr"]; ok {
		result["unique_autoconf_addr"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6UniqueAutoconfAddr2edl(i["unique-autoconf-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := i["vrip6_link_local"]; ok {
		result["vrip6_link_local"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrip6LinkLocal2edl(i["vrip6_link_local"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := i["vrrp-virtual-mac6"]; ok {
		result["vrrp_virtual_mac6"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6VrrpVirtualMac62edl(i["vrrp-virtual-mac6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp6"
	if _, ok := i["vrrp6"]; ok {
		result["vrrp6"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp62edl(i["vrrp6"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Autoconf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6CliConn6Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6ClientOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6InformationRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixDelegation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHint2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintPlt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintVlt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayInterfaceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelaySourceInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelaySourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Icmp6SendRedirect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6InterfaceIdentifier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Allowaccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DefaultLife2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixIaid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := i["delegated-prefix-iaid"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(i["delegated-prefix-iaid"], d, pre_append)
			tmp["delegated_prefix_iaid"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-DelegatedPrefixIaid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := i["prefix-id"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(i["prefix-id"], d, pre_append)
			tmp["prefix_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-PrefixId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := i["rdnss-service"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(i["rdnss-service"], d, pre_append)
			tmp["rdnss_service"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-RdnssService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := i["upstream-interface"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(i["upstream-interface"], d, pre_append)
			tmp["upstream_interface"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-UpstreamInterface")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DnsServerOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddr2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddrPrefix2edl(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6ExtraAddr-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddrPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6HopLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6LinkMtu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ManageFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6MaxInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6MinInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6OtherFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := i["dnssl"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListDnssl2edl(i["dnssl"], d, pre_append)
			tmp["dnssl"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-Dnssl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := i["preferred-life-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(i["preferred-life-time"], d, pre_append)
			tmp["preferred_life_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-PreferredLifeTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPrefix2edl(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListRdnss2edl(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := i["valid-life-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListValidLifeTime2edl(i["valid-life-time"], d, pre_append)
			tmp["valid_life_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-ValidLifeTime")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListDnssl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListRdnss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListValidLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ReachableTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6RetransTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6SendAdv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6UpstreamInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdCgaModifier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdSecurityLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampDelta2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampFuzz2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6RaSendMtu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6UniqueAutoconfAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrip6LinkLocal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6VrrpVirtualMac62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp62edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AcceptMode2edl(i["accept-mode"], d, pre_append)
			tmp["accept_mode"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-AcceptMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AdvInterval2edl(i["adv-interval"], d, pre_append)
			tmp["adv_interval"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-AdvInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := i["ignore-default-route"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6IgnoreDefaultRoute2edl(i["ignore-default-route"], d, pre_append)
			tmp["ignore_default_route"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-IgnoreDefaultRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Preempt2edl(i["preempt"], d, pre_append)
			tmp["preempt"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Preempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Priority2edl(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6StartTime2edl(i["start-time"], d, pre_append)
			tmp["start_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-StartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Status2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := i["vrdst6"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrdst62edl(i["vrdst6"], d, pre_append)
			tmp["vrdst6"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Vrdst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrgrp2edl(i["vrgrp"], d, pre_append)
			tmp["vrgrp"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Vrgrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrid2edl(i["vrid"], d, pre_append)
			tmp["vrid"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Vrid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := i["vrip6"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrip62edl(i["vrip6"], d, pre_append)
			tmp["vrip6"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Vrip6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AcceptMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AdvInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6IgnoreDefaultRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Preempt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Priority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6StartTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrdst62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrgrp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrip62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryip2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipAllowaccess2edl(i["allowaccess"], d, pre_append)
			tmp["allowaccess"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Allowaccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := i["detectprotocol"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipDetectprotocol2edl(i["detectprotocol"], d, pre_append)
			tmp["detectprotocol"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Detectprotocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := i["detectserver"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipDetectserver2edl(i["detectserver"], d, pre_append)
			tmp["detectserver"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Detectserver")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := i["gwdetect"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipGwdetect2edl(i["gwdetect"], d, pre_append)
			tmp["gwdetect"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Gwdetect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipHaPriority2edl(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secip_relay_ip"
		if _, ok := i["secip-relay-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipSecipRelayIp2edl(i["secip-relay-ip"], d, pre_append)
			tmp["secip_relay_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-SecipRelayIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := i["ping-serv-status"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipPingServStatus2edl(i["ping-serv-status"], d, pre_append)
			tmp["ping_serv_status"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-PingServStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipSeq2edl(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Seq")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipAllowaccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipDetectprotocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipDetectserver2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipGwdetect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipHaPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipSecipRelayIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipPingServStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipSeq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVlanid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrp2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpAcceptMode2edl(i["accept-mode"], d, pre_append)
			tmp["accept_mode"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-AcceptMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpAdvInterval2edl(i["adv-interval"], d, pre_append)
			tmp["adv_interval"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-AdvInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := i["ignore-default-route"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpIgnoreDefaultRoute2edl(i["ignore-default-route"], d, pre_append)
			tmp["ignore_default_route"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-IgnoreDefaultRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpPreempt2edl(i["preempt"], d, pre_append)
			tmp["preempt"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-Preempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpPriority2edl(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy_arp"
		if _, ok := i["proxy-arp"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpProxyArp2edl(i["proxy-arp"], d, pre_append)
			tmp["proxy_arp"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-ProxyArp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpStartTime2edl(i["start-time"], d, pre_append)
			tmp["start_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-StartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpStatus2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpVersion2edl(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-Version")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst"
		if _, ok := i["vrdst"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpVrdst2edl(i["vrdst"], d, pre_append)
			tmp["vrdst"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-Vrdst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := i["vrdst-priority"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpVrdstPriority2edl(i["vrdst-priority"], d, pre_append)
			tmp["vrdst_priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-VrdstPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpVrgrp2edl(i["vrgrp"], d, pre_append)
			tmp["vrgrp"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-Vrgrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpVrid2edl(i["vrid"], d, pre_append)
			tmp["vrid"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-Vrid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip"
		if _, ok := i["vrip"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpVrip2edl(i["vrip"], d, pre_append)
			tmp["vrip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Vrrp-Vrip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpAcceptMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpAdvInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpIgnoreDefaultRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpPreempt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpProxyArp2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpProxyArpId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceVrrp-ProxyArp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceVrrpProxyArpIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceVrrp-ProxyArp-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpProxyArpId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpProxyArpIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpStartTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpVrdst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpVrdstPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpVrgrp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpVrid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVrrpVrip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_dhcp_status", flattenObjectFspVlanDynamicMappingDhcpStatus2edl(o["_dhcp-status"], d, "_dhcp_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dhcp-status"], "ObjectFspVlanDynamicMapping-DhcpStatus"); ok {
			if err = d.Set("_dhcp_status", vv); err != nil {
				return fmt.Errorf("Error reading _dhcp_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dhcp_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectFspVlanDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectFspVlanDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectFspVlanDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectFspVlanDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dhcp_server", flattenObjectFspVlanDynamicMappingDhcpServer2edl(o["dhcp-server"], d, "dhcp_server")); err != nil {
			if vv, ok := fortiAPIPatch(o["dhcp-server"], "ObjectFspVlanDynamicMapping-DhcpServer"); ok {
				if err = d.Set("dhcp_server", vv); err != nil {
					return fmt.Errorf("Error reading dhcp_server: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dhcp_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dhcp_server"); ok {
			if err = d.Set("dhcp_server", flattenObjectFspVlanDynamicMappingDhcpServer2edl(o["dhcp-server"], d, "dhcp_server")); err != nil {
				if vv, ok := fortiAPIPatch(o["dhcp-server"], "ObjectFspVlanDynamicMapping-DhcpServer"); ok {
					if err = d.Set("dhcp_server", vv); err != nil {
						return fmt.Errorf("Error reading dhcp_server: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dhcp_server: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("interface", flattenObjectFspVlanDynamicMappingInterface2edl(o["interface"], d, "interface")); err != nil {
			if vv, ok := fortiAPIPatch(o["interface"], "ObjectFspVlanDynamicMapping-Interface"); ok {
				if err = d.Set("interface", vv); err != nil {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenObjectFspVlanDynamicMappingInterface2edl(o["interface"], d, "interface")); err != nil {
				if vv, ok := fortiAPIPatch(o["interface"], "ObjectFspVlanDynamicMapping-Interface"); ok {
					if err = d.Set("interface", vv); err != nil {
						return fmt.Errorf("Error reading interface: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectFspVlanDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanDynamicMappingDhcpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFspVlanDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectFspVlanDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_configuration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-configuration"], _ = expandObjectFspVlanDynamicMappingDhcpServerAutoConfiguration2edl(d, i["auto_configuration"], pre_append)
	}
	pre_append = pre + ".0." + "auto_managed_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-managed-status"], _ = expandObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus2edl(d, i["auto_managed_status"], pre_append)
	}
	pre_append = pre + ".0." + "conflicted_ip_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conflicted-ip-timeout"], _ = expandObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout2edl(d, i["conflicted_ip_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_auth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-auth"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsAuth2edl(d, i["ddns_auth"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_key"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-key"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsKey2edl(d, i["ddns_key"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_keyname"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-keyname"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsKeyname2edl(d, i["ddns_keyname"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_server_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-server-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsServerIp2edl(d, i["ddns_server_ip"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_ttl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-ttl"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsTtl2edl(d, i["ddns_ttl"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_update"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-update"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdate2edl(d, i["ddns_update"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_update_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-update-override"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride2edl(d, i["ddns_update_override"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_zone"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-zone"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsZone2edl(d, i["ddns_zone"], pre_append)
	}
	pre_append = pre + ".0." + "default_gateway"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-gateway"], _ = expandObjectFspVlanDynamicMappingDhcpServerDefaultGateway2edl(d, i["default_gateway"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_settings_from_fortiipam"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-settings-from-fortiipam"], _ = expandObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam2edl(d, i["dhcp_settings_from_fortiipam"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server1"], _ = expandObjectFspVlanDynamicMappingDhcpServerDnsServer12edl(d, i["dns_server1"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server2"], _ = expandObjectFspVlanDynamicMappingDhcpServerDnsServer22edl(d, i["dns_server2"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server3"], _ = expandObjectFspVlanDynamicMappingDhcpServerDnsServer32edl(d, i["dns_server3"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server4"], _ = expandObjectFspVlanDynamicMappingDhcpServerDnsServer42edl(d, i["dns_server4"], pre_append)
	}
	pre_append = pre + ".0." + "dns_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-service"], _ = expandObjectFspVlanDynamicMappingDhcpServerDnsService2edl(d, i["dns_service"], pre_append)
	}
	pre_append = pre + ".0." + "domain"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["domain"], _ = expandObjectFspVlanDynamicMappingDhcpServerDomain2edl(d, i["domain"], pre_append)
	}
	pre_append = pre + ".0." + "enable"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["enable"], _ = expandObjectFspVlanDynamicMappingDhcpServerEnable2edl(d, i["enable"], pre_append)
	}
	pre_append = pre + ".0." + "exclude_range"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerExcludeRange2edl(d, i["exclude_range"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["exclude-range"] = t
		}
	}
	pre_append = pre + ".0." + "filename"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["filename"], _ = expandObjectFspVlanDynamicMappingDhcpServerFilename2edl(d, i["filename"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_on_net_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["forticlient-on-net-status"], _ = expandObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus2edl(d, i["forticlient_on_net_status"], pre_append)
	}
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerId2edl(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "ip_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip-mode"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpMode2edl(d, i["ip_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip_range"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerIpRange2edl(d, i["ip_range"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ip-range"] = t
		}
	}
	pre_append = pre + ".0." + "ipsec_lease_hold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipsec-lease-hold"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold2edl(d, i["ipsec_lease_hold"], pre_append)
	}
	pre_append = pre + ".0." + "lease_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lease-time"], _ = expandObjectFspVlanDynamicMappingDhcpServerLeaseTime2edl(d, i["lease_time"], pre_append)
	}
	pre_append = pre + ".0." + "mac_acl_default_action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-acl-default-action"], _ = expandObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction2edl(d, i["mac_acl_default_action"], pre_append)
	}
	pre_append = pre + ".0." + "netmask"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["netmask"], _ = expandObjectFspVlanDynamicMappingDhcpServerNetmask2edl(d, i["netmask"], pre_append)
	}
	pre_append = pre + ".0." + "next_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["next-server"], _ = expandObjectFspVlanDynamicMappingDhcpServerNextServer2edl(d, i["next_server"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_server1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-server1"], _ = expandObjectFspVlanDynamicMappingDhcpServerNtpServer12edl(d, i["ntp_server1"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-server2"], _ = expandObjectFspVlanDynamicMappingDhcpServerNtpServer22edl(d, i["ntp_server2"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_server3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-server3"], _ = expandObjectFspVlanDynamicMappingDhcpServerNtpServer32edl(d, i["ntp_server3"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-service"], _ = expandObjectFspVlanDynamicMappingDhcpServerNtpService2edl(d, i["ntp_service"], pre_append)
	}
	pre_append = pre + ".0." + "option1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option1"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption12edl(d, i["option1"], pre_append)
	}
	pre_append = pre + ".0." + "option2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option2"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption22edl(d, i["option2"], pre_append)
	}
	pre_append = pre + ".0." + "option3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option3"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption32edl(d, i["option3"], pre_append)
	}
	pre_append = pre + ".0." + "option4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option4"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption42edl(d, i["option4"], pre_append)
	}
	pre_append = pre + ".0." + "option5"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option5"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption52edl(d, i["option5"], pre_append)
	}
	pre_append = pre + ".0." + "option6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option6"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption62edl(d, i["option6"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerOptions2edl(d, i["options"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["options"] = t
		}
	}
	pre_append = pre + ".0." + "relay_agent"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["relay-agent"], _ = expandObjectFspVlanDynamicMappingDhcpServerRelayAgent2edl(d, i["relay_agent"], pre_append)
	}
	pre_append = pre + ".0." + "reserved_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddress2edl(d, i["reserved_address"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["reserved-address"] = t
		}
	}
	pre_append = pre + ".0." + "server_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-type"], _ = expandObjectFspVlanDynamicMappingDhcpServerServerType2edl(d, i["server_type"], pre_append)
	}
	pre_append = pre + ".0." + "shared_subnet"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["shared-subnet"], _ = expandObjectFspVlanDynamicMappingDhcpServerSharedSubnet2edl(d, i["shared_subnet"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectFspVlanDynamicMappingDhcpServerStatus2edl(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tftp_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tftp-server"], _ = expandObjectFspVlanDynamicMappingDhcpServerTftpServer2edl(d, i["tftp_server"], pre_append)
	}
	pre_append = pre + ".0." + "timezone"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["timezone"], _ = expandObjectFspVlanDynamicMappingDhcpServerTimezone2edl(d, i["timezone"], pre_append)
	}
	pre_append = pre + ".0." + "timezone_option"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["timezone-option"], _ = expandObjectFspVlanDynamicMappingDhcpServerTimezoneOption2edl(d, i["timezone_option"], pre_append)
	}
	pre_append = pre + ".0." + "vci_match"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerVciMatch2edl(d, i["vci_match"], pre_append)
	}
	pre_append = pre + ".0." + "vci_string"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerVciString2edl(d, i["vci_string"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ac_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac-service"], _ = expandObjectFspVlanDynamicMappingDhcpServerWifiAcService2edl(d, i["wifi_ac_service"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ac1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac1"], _ = expandObjectFspVlanDynamicMappingDhcpServerWifiAc12edl(d, i["wifi_ac1"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ac2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac2"], _ = expandObjectFspVlanDynamicMappingDhcpServerWifiAc22edl(d, i["wifi_ac2"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ac3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac3"], _ = expandObjectFspVlanDynamicMappingDhcpServerWifiAc32edl(d, i["wifi_ac3"], pre_append)
	}
	pre_append = pre + ".0." + "wins_server1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wins-server1"], _ = expandObjectFspVlanDynamicMappingDhcpServerWinsServer12edl(d, i["wins_server1"], pre_append)
	}
	pre_append = pre + ".0." + "wins_server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wins-server2"], _ = expandObjectFspVlanDynamicMappingDhcpServerWinsServer22edl(d, i["wins_server2"], pre_append)
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerAutoConfiguration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsKeyname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsZone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDefaultGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerEnable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp2edl(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lease-time"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime2edl(d, i["lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp2edl(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch2edl(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString2edl(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch2edl(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString2edl(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerFilename2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp2edl(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lease-time"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeLeaseTime2edl(d, i["lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp2edl(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeUciMatch2edl(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeUciString2edl(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch2edl(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciString2edl(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeUciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeUciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNetmask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNextServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpServer32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["code"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsCode2edl(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsUciMatch2edl(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsUciString2edl(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsValue2edl(d, i["value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch2edl(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsVciString2edl(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsUciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsUciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerRelayAgent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId2edl(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id-type"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType2edl(d, i["circuit_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription2edl(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressMac2edl(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId2edl(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id-type"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType2edl(d, i["remote_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressType2edl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerServerType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerSharedSubnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerTftpServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerTimezone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerTimezoneOption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAcService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAc12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAc22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAc32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWinsServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWinsServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dhcp_relay_agent_option"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-agent-option"], _ = expandObjectFspVlanDynamicMappingInterfaceDhcpRelayAgentOption2edl(d, i["dhcp_relay_agent_option"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_interface_select_method"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-interface-select-method"], _ = expandObjectFspVlanDynamicMappingInterfaceDhcpRelayInterfaceSelectMethod2edl(d, i["dhcp_relay_interface_select_method"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-ip"], _ = expandObjectFspVlanDynamicMappingInterfaceDhcpRelayIp2edl(d, i["dhcp_relay_ip"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-service"], _ = expandObjectFspVlanDynamicMappingInterfaceDhcpRelayService2edl(d, i["dhcp_relay_service"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-type"], _ = expandObjectFspVlanDynamicMappingInterfaceDhcpRelayType2edl(d, i["dhcp_relay_type"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandObjectFspVlanDynamicMappingInterfaceIp2edl(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingInterfaceIpv62edl(d, i["ipv6"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ipv6"] = t
		}
	}
	pre_append = pre + ".0." + "secondary_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["secondary-IP"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryIp2edl(d, i["secondary_ip"], pre_append)
	}
	pre_append = pre + ".0." + "secondaryip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingInterfaceSecondaryip2edl(d, i["secondaryip"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["secondaryip"] = t
		}
	}
	pre_append = pre + ".0." + "vlanid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlanid"], _ = expandObjectFspVlanDynamicMappingInterfaceVlanid2edl(d, i["vlanid"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingInterfaceVrrp2edl(d, i["vrrp"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["vrrp"] = t
		}
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceDhcpRelayAgentOption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceDhcpRelayInterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceDhcpRelayIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceDhcpRelayService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceDhcpRelayType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "autoconf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["autoconf"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Autoconf2edl(d, i["autoconf"], pre_append)
	}
	pre_append = pre + ".0." + "cli_conn6_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cli-conn6-status"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6CliConn6Status2edl(d, i["cli_conn6_status"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-client-options"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6ClientOptions2edl(d, i["dhcp6_client_options"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-information-request"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6InformationRequest2edl(d, i["dhcp6_information_request"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-delegation"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixDelegation2edl(d, i["dhcp6_prefix_delegation"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHint2edl(d, i["dhcp6_prefix_hint"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint-plt"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintPlt2edl(d, i["dhcp6_prefix_hint_plt"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint-vlt"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintVlt2edl(d, i["dhcp6_prefix_hint_vlt"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_interface_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-interface-id"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayInterfaceId2edl(d, i["dhcp6_relay_interface_id"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-ip"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayIp2edl(d, i["dhcp6_relay_ip"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-service"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayService2edl(d, i["dhcp6_relay_service"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_source_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-source-interface"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelaySourceInterface2edl(d, i["dhcp6_relay_source_interface"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_source_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-source-ip"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelaySourceIp2edl(d, i["dhcp6_relay_source_ip"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-type"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayType2edl(d, i["dhcp6_relay_type"], pre_append)
	}
	pre_append = pre + ".0." + "icmp6_send_redirect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp6-send-redirect"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Icmp6SendRedirect2edl(d, i["icmp6_send_redirect"], pre_append)
	}
	pre_append = pre + ".0." + "interface_identifier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interface-identifier"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6InterfaceIdentifier2edl(d, i["interface_identifier"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-address"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Address2edl(d, i["ip6_address"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-allowaccess"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Allowaccess2edl(d, i["ip6_allowaccess"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-default-life"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DefaultLife2edl(d, i["ip6_default_life"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_delegated_prefix_iaid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-delegated-prefix-iaid"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixIaid2edl(d, i["ip6_delegated_prefix_iaid"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixList2edl(d, i["ip6_delegated_prefix_list"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ip6-delegated-prefix-list"] = t
		}
	}
	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-dns-server-override"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DnsServerOverride2edl(d, i["ip6_dns_server_override"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddr2edl(d, i["ip6_extra_addr"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ip6-extra-addr"] = t
		}
	}
	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-hop-limit"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6HopLimit2edl(d, i["ip6_hop_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-link-mtu"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6LinkMtu2edl(d, i["ip6_link_mtu"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-manage-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ManageFlag2edl(d, i["ip6_manage_flag"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-max-interval"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6MaxInterval2edl(d, i["ip6_max_interval"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-min-interval"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6MinInterval2edl(d, i["ip6_min_interval"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-mode"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Mode2edl(d, i["ip6_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-other-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6OtherFlag2edl(d, i["ip6_other_flag"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixList2edl(d, i["ip6_prefix_list"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ip6-prefix-list"] = t
		}
	}
	pre_append = pre + ".0." + "ip6_prefix_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-prefix-mode"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixMode2edl(d, i["ip6_prefix_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-reachable-time"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ReachableTime2edl(d, i["ip6_reachable_time"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-retrans-time"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6RetransTime2edl(d, i["ip6_retrans_time"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-send-adv"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6SendAdv2edl(d, i["ip6_send_adv"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-subnet"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Subnet2edl(d, i["ip6_subnet"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-upstream-interface"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6UpstreamInterface2edl(d, i["ip6_upstream_interface"], pre_append)
	}
	pre_append = pre + ".0." + "nd_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-cert"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdCert2edl(d, i["nd_cert"], pre_append)
	}
	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-cga-modifier"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdCgaModifier2edl(d, i["nd_cga_modifier"], pre_append)
	}
	pre_append = pre + ".0." + "nd_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-mode"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdMode2edl(d, i["nd_mode"], pre_append)
	}
	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-security-level"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdSecurityLevel2edl(d, i["nd_security_level"], pre_append)
	}
	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-timestamp-delta"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampDelta2edl(d, i["nd_timestamp_delta"], pre_append)
	}
	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-timestamp-fuzz"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampFuzz2edl(d, i["nd_timestamp_fuzz"], pre_append)
	}
	pre_append = pre + ".0." + "ra_send_mtu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ra-send-mtu"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6RaSendMtu2edl(d, i["ra_send_mtu"], pre_append)
	}
	pre_append = pre + ".0." + "unique_autoconf_addr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unique-autoconf-addr"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6UniqueAutoconfAddr2edl(d, i["unique_autoconf_addr"], pre_append)
	}
	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrip6_link_local"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrip6LinkLocal2edl(d, i["vrip6_link_local"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrrp-virtual-mac6"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6VrrpVirtualMac62edl(d, i["vrrp_virtual_mac6"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp62edl(d, i["vrrp6"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["vrrp6"] = t
		}
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Autoconf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6CliConn6Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6ClientOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6InformationRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixDelegation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHint2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintPlt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintVlt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayInterfaceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelaySourceInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelaySourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Icmp6SendRedirect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6InterfaceIdentifier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Allowaccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DefaultLife2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixIaid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["autonomous-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["delegated-prefix-iaid"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(d, i["delegated_prefix_iaid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-id"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(d, i["prefix_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(d, i["rdnss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss-service"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(d, i["rdnss_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["upstream-interface"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(d, i["upstream_interface"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DnsServerOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["prefix"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddrPrefix2edl(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddrPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6HopLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6LinkMtu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ManageFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6MaxInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6MinInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6OtherFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["autonomous-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dnssl"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListDnssl2edl(d, i["dnssl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preferred-life-time"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(d, i["preferred_life_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPrefix2edl(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListRdnss2edl(d, i["rdnss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["valid-life-time"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListValidLifeTime2edl(d, i["valid_life_time"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListDnssl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListRdnss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListValidLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ReachableTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6RetransTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6SendAdv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6UpstreamInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdCgaModifier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdSecurityLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampDelta2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampFuzz2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6RaSendMtu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6UniqueAutoconfAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrip6LinkLocal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6VrrpVirtualMac62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["accept-mode"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AcceptMode2edl(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-interval"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AdvInterval2edl(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ignore-default-route"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6IgnoreDefaultRoute2edl(d, i["ignore_default_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preempt"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Preempt2edl(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Priority2edl(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-time"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6StartTime2edl(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Status2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst6"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrdst62edl(d, i["vrdst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrgrp"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrgrp2edl(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrid"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrid2edl(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrip6"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrip62edl(d, i["vrip6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AcceptMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AdvInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6IgnoreDefaultRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Preempt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Priority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6StartTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrdst62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrgrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrip62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["allowaccess"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipAllowaccess2edl(d, i["allowaccess"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detectprotocol"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipDetectprotocol2edl(d, i["detectprotocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detectserver"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipDetectserver2edl(d, i["detectserver"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gwdetect"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipGwdetect2edl(d, i["gwdetect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipHaPriority2edl(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secip_relay_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secip-relay-ip"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipSecipRelayIp2edl(d, i["secip_relay_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ping-serv-status"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipPingServStatus2edl(d, i["ping_serv_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipSeq2edl(d, i["seq"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipAllowaccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipDetectprotocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipDetectserver2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipGwdetect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipHaPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipSecipRelayIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipPingServStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipSeq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVlanid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["accept-mode"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpAcceptMode2edl(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-interval"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpAdvInterval2edl(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ignore-default-route"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpIgnoreDefaultRoute2edl(d, i["ignore_default_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preempt"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpPreempt2edl(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpPriority2edl(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy_arp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFspVlanDynamicMappingInterfaceVrrpProxyArp2edl(d, i["proxy_arp"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["proxy-arp"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-time"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpStartTime2edl(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpStatus2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpVersion2edl(d, i["version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpVrdst2edl(d, i["vrdst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst-priority"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpVrdstPriority2edl(d, i["vrdst_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrgrp"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpVrgrp2edl(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrid"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpVrid2edl(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrip"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpVrip2edl(d, i["vrip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpAcceptMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpAdvInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpIgnoreDefaultRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpPreempt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpProxyArp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpProxyArpId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDynamicMappingInterfaceVrrpProxyArpIp2edl(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpProxyArpId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpProxyArpIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpStartTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpVrdst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpVrdstPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpVrgrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpVrid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVrrpVrip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_dhcp_status"); ok || d.HasChange("_dhcp_status") {
		t, err := expandObjectFspVlanDynamicMappingDhcpStatus2edl(d, v, "_dhcp_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dhcp-status"] = t
		}
	}

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectFspVlanDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server"); ok || d.HasChange("dhcp_server") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServer2edl(d, v, "dhcp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectFspVlanDynamicMappingInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	return &obj, nil
}
