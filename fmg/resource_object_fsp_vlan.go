// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFsp Vlan

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFspVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanCreate,
		Read:   resourceObjectFspVlanRead,
		Update: resourceObjectFspVlanUpdate,
		Delete: resourceObjectFspVlanDelete,

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
			"_dhcp_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
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
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vci_match": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vci_string": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
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
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vci_match": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vci_string": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
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
						},
						"option2": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"option3": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
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
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vci_match": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vci_string": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
									},
								},
							},
						},
						"reserved_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
									"mac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
								},
							},
						},
						"server_type": &schema.Schema{
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
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_configuration": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"auto_managed_status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"conflicted_ip_timeout": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ddns_auth": &schema.Schema{
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
									"ddns_server_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ddns_ttl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ddns_update": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ddns_update_override": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ddns_zone": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"default_gateway": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dhcp_settings_from_fortiipam": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dns_server1": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dns_server2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dns_server3": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dns_server4": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dns_service": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
												},
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"start_ip": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"vci_match": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"vci_string": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
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
												},
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"start_ip": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"vci_match": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"vci_string": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
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
									},
									"mac_acl_default_action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"netmask": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"next_server": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ntp_server1": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ntp_server2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ntp_server3": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ntp_service": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"option1": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
									},
									"option2": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
									},
									"option3": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
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
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"vci_match": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"vci_string": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
												},
											},
										},
									},
									"reserved_address": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
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
												},
												"mac": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
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
												},
											},
										},
									},
									"server_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"tftp_server": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
									},
									"timezone": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"timezone_option": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vci_match": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vci_string": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
									},
									"wifi_ac_service": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"wifi_ac1": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"wifi_ac2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"wifi_ac3": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"wins_server1": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"wins_server2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"interface": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
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
												"dhcp6_relay_ip": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"dhcp6_relay_service": &schema.Schema{
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
												},
												"detectprotocol": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
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
								},
							},
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_op_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"ap_discover": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"arpforward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"bandwidth_measure_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"dhcp_classless_route_addition": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_client_identifier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_relay_agent_option": &schema.Schema{
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
						},
						"dhcp_relay_ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"dhcp_relay_link_selection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_relay_request_all_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_relay_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_relay_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_renew_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						},
						"drop_fragment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"drop_overlapped_fragment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"eap_ca_cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
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
						},
						"eap_supplicant": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"eap_user_cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"egress_cos": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"egress_shaping_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"eip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"explicit_web_proxy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fail_action_on_extender": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"fail_detect_option": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
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
						},
						"fortilink_backup_link": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fortilink_neighbor_detect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortilink_split_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortilink_stacking": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"forward_domain": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"forward_error_correction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fp_anomaly": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"fp_disable": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
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
						},
						"icmp_redirect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"icmp_send_redirect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ident_accept": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"internal": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_managed_by_fortiipam": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipmac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ips_sniffer_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipunnumbered": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
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
									"dhcp6_relay_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"dhcp6_relay_service": &schema.Schema{
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
						"l2forward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"l2tp_client": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"lacp_ha_secondary": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"management_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"min_links_down": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"monitor_bandwidth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mtu_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"netbios_forward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"netflow_sampler": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"polling_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pppoe_unnumbered_negotiate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pptp_auth_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pptp_client": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pptp_password": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"pptp_server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"proxy_captive_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"sample_direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sample_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"scan_botnet_connections": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
									},
									"detectprotocol": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
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
						},
						"security_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"security_redirect_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"select_profile_30a_35b": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sflow_sampler": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sfp_dsl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sfp_dsl_adsl_fallback": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sfp_dsl_autodetect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sfp_dsl_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"speed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"spillover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"src_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"stp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"stp_ha_secondary": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"stp_ha_slave": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"stpforward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"stpforward_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"strip_priority_vlan_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subst": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"substitute_dst_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sw_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"switch_controller_arp_inspection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_dhcp_snooping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_dhcp_snooping_option82": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_dhcp_snooping_verify_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_dynamic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_feature": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_igmp_snooping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_igmp_snooping_fast_leave": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_igmp_snooping_proxy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_iot_scanning": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_learning_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"switch_controller_mgmt_vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"switch_controller_nac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_netflow_collect": &schema.Schema{
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
						},
						"switch_controller_source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_traffic_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"system_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"system_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"trust_ip6_2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"trust_ip6_3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"vindex": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vlan_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlanforward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vrdst": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
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
						"vrrp_virtual_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wccp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"wifi_5g_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wifi_acl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wifi_ap_band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wifi_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wifi_auto_connect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wifi_auto_save": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wifi_broadcast_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wifi_encrypt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wifi_fragment_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"wifi_key": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
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
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
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
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlanid": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceObjectFspVlanCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFspVlan(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlan resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFspVlan(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlan resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFspVlanRead(d, m)
}

func resourceObjectFspVlanUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFspVlan(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlan resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFspVlan(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFspVlanRead(d, m)
}

func resourceObjectFspVlanDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFspVlan(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFspVlan(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlan resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanDhcpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_configuration"
	if _, ok := i["auto-configuration"]; ok {
		result["auto_configuration"] = flattenObjectFspVlanDhcpServerAutoConfiguration(i["auto-configuration"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_managed_status"
	if _, ok := i["auto-managed-status"]; ok {
		result["auto_managed_status"] = flattenObjectFspVlanDhcpServerAutoManagedStatus(i["auto-managed-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "conflicted_ip_timeout"
	if _, ok := i["conflicted-ip-timeout"]; ok {
		result["conflicted_ip_timeout"] = flattenObjectFspVlanDhcpServerConflictedIpTimeout(i["conflicted-ip-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_auth"
	if _, ok := i["ddns-auth"]; ok {
		result["ddns_auth"] = flattenObjectFspVlanDhcpServerDdnsAuth(i["ddns-auth"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_key"
	if _, ok := i["ddns-key"]; ok {
		result["ddns_key"] = flattenObjectFspVlanDhcpServerDdnsKey(i["ddns-key"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_keyname"
	if _, ok := i["ddns-keyname"]; ok {
		result["ddns_keyname"] = flattenObjectFspVlanDhcpServerDdnsKeyname(i["ddns-keyname"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_server_ip"
	if _, ok := i["ddns-server-ip"]; ok {
		result["ddns_server_ip"] = flattenObjectFspVlanDhcpServerDdnsServerIp(i["ddns-server-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_ttl"
	if _, ok := i["ddns-ttl"]; ok {
		result["ddns_ttl"] = flattenObjectFspVlanDhcpServerDdnsTtl(i["ddns-ttl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_update"
	if _, ok := i["ddns-update"]; ok {
		result["ddns_update"] = flattenObjectFspVlanDhcpServerDdnsUpdate(i["ddns-update"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_update_override"
	if _, ok := i["ddns-update-override"]; ok {
		result["ddns_update_override"] = flattenObjectFspVlanDhcpServerDdnsUpdateOverride(i["ddns-update-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_zone"
	if _, ok := i["ddns-zone"]; ok {
		result["ddns_zone"] = flattenObjectFspVlanDhcpServerDdnsZone(i["ddns-zone"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_gateway"
	if _, ok := i["default-gateway"]; ok {
		result["default_gateway"] = flattenObjectFspVlanDhcpServerDefaultGateway(i["default-gateway"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_settings_from_fortiipam"
	if _, ok := i["dhcp-settings-from-fortiipam"]; ok {
		result["dhcp_settings_from_fortiipam"] = flattenObjectFspVlanDhcpServerDhcpSettingsFromFortiipam(i["dhcp-settings-from-fortiipam"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server1"
	if _, ok := i["dns-server1"]; ok {
		result["dns_server1"] = flattenObjectFspVlanDhcpServerDnsServer1(i["dns-server1"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server2"
	if _, ok := i["dns-server2"]; ok {
		result["dns_server2"] = flattenObjectFspVlanDhcpServerDnsServer2(i["dns-server2"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server3"
	if _, ok := i["dns-server3"]; ok {
		result["dns_server3"] = flattenObjectFspVlanDhcpServerDnsServer3(i["dns-server3"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server4"
	if _, ok := i["dns-server4"]; ok {
		result["dns_server4"] = flattenObjectFspVlanDhcpServerDnsServer4(i["dns-server4"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_service"
	if _, ok := i["dns-service"]; ok {
		result["dns_service"] = flattenObjectFspVlanDhcpServerDnsService(i["dns-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "domain"
	if _, ok := i["domain"]; ok {
		result["domain"] = flattenObjectFspVlanDhcpServerDomain(i["domain"], d, pre_append)
	}

	pre_append = pre + ".0." + "enable"
	if _, ok := i["enable"]; ok {
		result["enable"] = flattenObjectFspVlanDhcpServerEnable(i["enable"], d, pre_append)
	}

	pre_append = pre + ".0." + "exclude_range"
	if _, ok := i["exclude-range"]; ok {
		result["exclude_range"] = flattenObjectFspVlanDhcpServerExcludeRange(i["exclude-range"], d, pre_append)
	}

	pre_append = pre + ".0." + "filename"
	if _, ok := i["filename"]; ok {
		result["filename"] = flattenObjectFspVlanDhcpServerFilename(i["filename"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_on_net_status"
	if _, ok := i["forticlient-on-net-status"]; ok {
		result["forticlient_on_net_status"] = flattenObjectFspVlanDhcpServerForticlientOnNetStatus(i["forticlient-on-net-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenObjectFspVlanDhcpServerId(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_mode"
	if _, ok := i["ip-mode"]; ok {
		result["ip_mode"] = flattenObjectFspVlanDhcpServerIpMode(i["ip-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_range"
	if _, ok := i["ip-range"]; ok {
		result["ip_range"] = flattenObjectFspVlanDhcpServerIpRange(i["ip-range"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipsec_lease_hold"
	if _, ok := i["ipsec-lease-hold"]; ok {
		result["ipsec_lease_hold"] = flattenObjectFspVlanDhcpServerIpsecLeaseHold(i["ipsec-lease-hold"], d, pre_append)
	}

	pre_append = pre + ".0." + "lease_time"
	if _, ok := i["lease-time"]; ok {
		result["lease_time"] = flattenObjectFspVlanDhcpServerLeaseTime(i["lease-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_acl_default_action"
	if _, ok := i["mac-acl-default-action"]; ok {
		result["mac_acl_default_action"] = flattenObjectFspVlanDhcpServerMacAclDefaultAction(i["mac-acl-default-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "netmask"
	if _, ok := i["netmask"]; ok {
		result["netmask"] = flattenObjectFspVlanDhcpServerNetmask(i["netmask"], d, pre_append)
	}

	pre_append = pre + ".0." + "next_server"
	if _, ok := i["next-server"]; ok {
		result["next_server"] = flattenObjectFspVlanDhcpServerNextServer(i["next-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_server1"
	if _, ok := i["ntp-server1"]; ok {
		result["ntp_server1"] = flattenObjectFspVlanDhcpServerNtpServer1(i["ntp-server1"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_server2"
	if _, ok := i["ntp-server2"]; ok {
		result["ntp_server2"] = flattenObjectFspVlanDhcpServerNtpServer2(i["ntp-server2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_server3"
	if _, ok := i["ntp-server3"]; ok {
		result["ntp_server3"] = flattenObjectFspVlanDhcpServerNtpServer3(i["ntp-server3"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_service"
	if _, ok := i["ntp-service"]; ok {
		result["ntp_service"] = flattenObjectFspVlanDhcpServerNtpService(i["ntp-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "option1"
	if _, ok := i["option1"]; ok {
		result["option1"] = flattenObjectFspVlanDhcpServerOption1(i["option1"], d, pre_append)
	}

	pre_append = pre + ".0." + "option2"
	if _, ok := i["option2"]; ok {
		result["option2"] = flattenObjectFspVlanDhcpServerOption2(i["option2"], d, pre_append)
	}

	pre_append = pre + ".0." + "option3"
	if _, ok := i["option3"]; ok {
		result["option3"] = flattenObjectFspVlanDhcpServerOption3(i["option3"], d, pre_append)
	}

	pre_append = pre + ".0." + "option4"
	if _, ok := i["option4"]; ok {
		result["option4"] = flattenObjectFspVlanDhcpServerOption4(i["option4"], d, pre_append)
	}

	pre_append = pre + ".0." + "option5"
	if _, ok := i["option5"]; ok {
		result["option5"] = flattenObjectFspVlanDhcpServerOption5(i["option5"], d, pre_append)
	}

	pre_append = pre + ".0." + "option6"
	if _, ok := i["option6"]; ok {
		result["option6"] = flattenObjectFspVlanDhcpServerOption6(i["option6"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFspVlanDhcpServerOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "reserved_address"
	if _, ok := i["reserved-address"]; ok {
		result["reserved_address"] = flattenObjectFspVlanDhcpServerReservedAddress(i["reserved-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_type"
	if _, ok := i["server-type"]; ok {
		result["server_type"] = flattenObjectFspVlanDhcpServerServerType(i["server-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFspVlanDhcpServerStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tftp_server"
	if _, ok := i["tftp-server"]; ok {
		result["tftp_server"] = flattenObjectFspVlanDhcpServerTftpServer(i["tftp-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "timezone"
	if _, ok := i["timezone"]; ok {
		result["timezone"] = flattenObjectFspVlanDhcpServerTimezone(i["timezone"], d, pre_append)
	}

	pre_append = pre + ".0." + "timezone_option"
	if _, ok := i["timezone-option"]; ok {
		result["timezone_option"] = flattenObjectFspVlanDhcpServerTimezoneOption(i["timezone-option"], d, pre_append)
	}

	pre_append = pre + ".0." + "vci_match"
	if _, ok := i["vci-match"]; ok {
		result["vci_match"] = flattenObjectFspVlanDhcpServerVciMatch(i["vci-match"], d, pre_append)
	}

	pre_append = pre + ".0." + "vci_string"
	if _, ok := i["vci-string"]; ok {
		result["vci_string"] = flattenObjectFspVlanDhcpServerVciString(i["vci-string"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac_service"
	if _, ok := i["wifi-ac-service"]; ok {
		result["wifi_ac_service"] = flattenObjectFspVlanDhcpServerWifiAcService(i["wifi-ac-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac1"
	if _, ok := i["wifi-ac1"]; ok {
		result["wifi_ac1"] = flattenObjectFspVlanDhcpServerWifiAc1(i["wifi-ac1"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac2"
	if _, ok := i["wifi-ac2"]; ok {
		result["wifi_ac2"] = flattenObjectFspVlanDhcpServerWifiAc2(i["wifi-ac2"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac3"
	if _, ok := i["wifi-ac3"]; ok {
		result["wifi_ac3"] = flattenObjectFspVlanDhcpServerWifiAc3(i["wifi-ac3"], d, pre_append)
	}

	pre_append = pre + ".0." + "wins_server1"
	if _, ok := i["wins-server1"]; ok {
		result["wins_server1"] = flattenObjectFspVlanDhcpServerWinsServer1(i["wins-server1"], d, pre_append)
	}

	pre_append = pre + ".0." + "wins_server2"
	if _, ok := i["wins-server2"]; ok {
		result["wins_server2"] = flattenObjectFspVlanDhcpServerWinsServer2(i["wins-server2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFspVlanDhcpServerAutoConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerAutoManagedStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerConflictedIpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsKeyname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsUpdateOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDefaultGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDhcpSettingsFromFortiipam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDnsServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDnsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDhcpServerExcludeRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-VciString")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDhcpServerExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRangeVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRangeVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerForticlientOnNetStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDhcpServerIpRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-VciString")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDhcpServerIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRangeVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRangeVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerIpsecLeaseHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerMacAclDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNextServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNtpServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNtpServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNtpServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNtpService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOption1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerOption2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerOption3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerOption4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOption5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOption6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDhcpServerOptionsCode(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-Value")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-VciString")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDhcpServerOptionsCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerOptionsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerReservedAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDhcpServerReservedAddressAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressCircuitId(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := i["circuit-id-type"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressCircuitIdType(i["circuit-id-type"], d, pre_append)
			tmp["circuit_id_type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-CircuitIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressRemoteId(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := i["remote-id-type"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressRemoteIdType(i["remote-id-type"], d, pre_append)
			tmp["remote_id_type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-RemoteIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDhcpServerReservedAddressAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressCircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressCircuitIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressRemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressRemoteIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerTftpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerTimezoneOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerWifiAcService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerWifiAc1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerWifiAc2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerWifiAc3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_dhcp_status"
		if _, ok := i["_dhcp-status"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpStatus(i["_dhcp-status"], d, pre_append)
			tmp["_dhcp_status"] = fortiAPISubPartPatch(v, "ObjectFspVlan-DynamicMapping-DhcpStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectFspVlanDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectFspVlan-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_server"
		if _, ok := i["dhcp-server"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServer(i["dhcp-server"], d, pre_append)
			tmp["dhcp_server"] = fortiAPISubPartPatch(v, "ObjectFspVlan-DynamicMapping-DhcpServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectFspVlan-DynamicMapping-Interface")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFspVlanDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_configuration"
	if _, ok := i["auto-configuration"]; ok {
		result["auto_configuration"] = flattenObjectFspVlanDynamicMappingDhcpServerAutoConfiguration(i["auto-configuration"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_managed_status"
	if _, ok := i["auto-managed-status"]; ok {
		result["auto_managed_status"] = flattenObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus(i["auto-managed-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "conflicted_ip_timeout"
	if _, ok := i["conflicted-ip-timeout"]; ok {
		result["conflicted_ip_timeout"] = flattenObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout(i["conflicted-ip-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_auth"
	if _, ok := i["ddns-auth"]; ok {
		result["ddns_auth"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsAuth(i["ddns-auth"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_key"
	if _, ok := i["ddns-key"]; ok {
		result["ddns_key"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsKey(i["ddns-key"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_keyname"
	if _, ok := i["ddns-keyname"]; ok {
		result["ddns_keyname"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsKeyname(i["ddns-keyname"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_server_ip"
	if _, ok := i["ddns-server-ip"]; ok {
		result["ddns_server_ip"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsServerIp(i["ddns-server-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_ttl"
	if _, ok := i["ddns-ttl"]; ok {
		result["ddns_ttl"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsTtl(i["ddns-ttl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_update"
	if _, ok := i["ddns-update"]; ok {
		result["ddns_update"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdate(i["ddns-update"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_update_override"
	if _, ok := i["ddns-update-override"]; ok {
		result["ddns_update_override"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride(i["ddns-update-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_zone"
	if _, ok := i["ddns-zone"]; ok {
		result["ddns_zone"] = flattenObjectFspVlanDynamicMappingDhcpServerDdnsZone(i["ddns-zone"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_gateway"
	if _, ok := i["default-gateway"]; ok {
		result["default_gateway"] = flattenObjectFspVlanDynamicMappingDhcpServerDefaultGateway(i["default-gateway"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_settings_from_fortiipam"
	if _, ok := i["dhcp-settings-from-fortiipam"]; ok {
		result["dhcp_settings_from_fortiipam"] = flattenObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam(i["dhcp-settings-from-fortiipam"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server1"
	if _, ok := i["dns-server1"]; ok {
		result["dns_server1"] = flattenObjectFspVlanDynamicMappingDhcpServerDnsServer1(i["dns-server1"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server2"
	if _, ok := i["dns-server2"]; ok {
		result["dns_server2"] = flattenObjectFspVlanDynamicMappingDhcpServerDnsServer2(i["dns-server2"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server3"
	if _, ok := i["dns-server3"]; ok {
		result["dns_server3"] = flattenObjectFspVlanDynamicMappingDhcpServerDnsServer3(i["dns-server3"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server4"
	if _, ok := i["dns-server4"]; ok {
		result["dns_server4"] = flattenObjectFspVlanDynamicMappingDhcpServerDnsServer4(i["dns-server4"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_service"
	if _, ok := i["dns-service"]; ok {
		result["dns_service"] = flattenObjectFspVlanDynamicMappingDhcpServerDnsService(i["dns-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "domain"
	if _, ok := i["domain"]; ok {
		result["domain"] = flattenObjectFspVlanDynamicMappingDhcpServerDomain(i["domain"], d, pre_append)
	}

	pre_append = pre + ".0." + "enable"
	if _, ok := i["enable"]; ok {
		result["enable"] = flattenObjectFspVlanDynamicMappingDhcpServerEnable(i["enable"], d, pre_append)
	}

	pre_append = pre + ".0." + "exclude_range"
	if _, ok := i["exclude-range"]; ok {
		result["exclude_range"] = flattenObjectFspVlanDynamicMappingDhcpServerExcludeRange(i["exclude-range"], d, pre_append)
	}

	pre_append = pre + ".0." + "filename"
	if _, ok := i["filename"]; ok {
		result["filename"] = flattenObjectFspVlanDynamicMappingDhcpServerFilename(i["filename"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_on_net_status"
	if _, ok := i["forticlient-on-net-status"]; ok {
		result["forticlient_on_net_status"] = flattenObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus(i["forticlient-on-net-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenObjectFspVlanDynamicMappingDhcpServerId(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_mode"
	if _, ok := i["ip-mode"]; ok {
		result["ip_mode"] = flattenObjectFspVlanDynamicMappingDhcpServerIpMode(i["ip-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_range"
	if _, ok := i["ip-range"]; ok {
		result["ip_range"] = flattenObjectFspVlanDynamicMappingDhcpServerIpRange(i["ip-range"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipsec_lease_hold"
	if _, ok := i["ipsec-lease-hold"]; ok {
		result["ipsec_lease_hold"] = flattenObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold(i["ipsec-lease-hold"], d, pre_append)
	}

	pre_append = pre + ".0." + "lease_time"
	if _, ok := i["lease-time"]; ok {
		result["lease_time"] = flattenObjectFspVlanDynamicMappingDhcpServerLeaseTime(i["lease-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_acl_default_action"
	if _, ok := i["mac-acl-default-action"]; ok {
		result["mac_acl_default_action"] = flattenObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction(i["mac-acl-default-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "netmask"
	if _, ok := i["netmask"]; ok {
		result["netmask"] = flattenObjectFspVlanDynamicMappingDhcpServerNetmask(i["netmask"], d, pre_append)
	}

	pre_append = pre + ".0." + "next_server"
	if _, ok := i["next-server"]; ok {
		result["next_server"] = flattenObjectFspVlanDynamicMappingDhcpServerNextServer(i["next-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_server1"
	if _, ok := i["ntp-server1"]; ok {
		result["ntp_server1"] = flattenObjectFspVlanDynamicMappingDhcpServerNtpServer1(i["ntp-server1"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_server2"
	if _, ok := i["ntp-server2"]; ok {
		result["ntp_server2"] = flattenObjectFspVlanDynamicMappingDhcpServerNtpServer2(i["ntp-server2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_server3"
	if _, ok := i["ntp-server3"]; ok {
		result["ntp_server3"] = flattenObjectFspVlanDynamicMappingDhcpServerNtpServer3(i["ntp-server3"], d, pre_append)
	}

	pre_append = pre + ".0." + "ntp_service"
	if _, ok := i["ntp-service"]; ok {
		result["ntp_service"] = flattenObjectFspVlanDynamicMappingDhcpServerNtpService(i["ntp-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "option1"
	if _, ok := i["option1"]; ok {
		result["option1"] = flattenObjectFspVlanDynamicMappingDhcpServerOption1(i["option1"], d, pre_append)
	}

	pre_append = pre + ".0." + "option2"
	if _, ok := i["option2"]; ok {
		result["option2"] = flattenObjectFspVlanDynamicMappingDhcpServerOption2(i["option2"], d, pre_append)
	}

	pre_append = pre + ".0." + "option3"
	if _, ok := i["option3"]; ok {
		result["option3"] = flattenObjectFspVlanDynamicMappingDhcpServerOption3(i["option3"], d, pre_append)
	}

	pre_append = pre + ".0." + "option4"
	if _, ok := i["option4"]; ok {
		result["option4"] = flattenObjectFspVlanDynamicMappingDhcpServerOption4(i["option4"], d, pre_append)
	}

	pre_append = pre + ".0." + "option5"
	if _, ok := i["option5"]; ok {
		result["option5"] = flattenObjectFspVlanDynamicMappingDhcpServerOption5(i["option5"], d, pre_append)
	}

	pre_append = pre + ".0." + "option6"
	if _, ok := i["option6"]; ok {
		result["option6"] = flattenObjectFspVlanDynamicMappingDhcpServerOption6(i["option6"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFspVlanDynamicMappingDhcpServerOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "reserved_address"
	if _, ok := i["reserved-address"]; ok {
		result["reserved_address"] = flattenObjectFspVlanDynamicMappingDhcpServerReservedAddress(i["reserved-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_type"
	if _, ok := i["server-type"]; ok {
		result["server_type"] = flattenObjectFspVlanDynamicMappingDhcpServerServerType(i["server-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFspVlanDynamicMappingDhcpServerStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tftp_server"
	if _, ok := i["tftp-server"]; ok {
		result["tftp_server"] = flattenObjectFspVlanDynamicMappingDhcpServerTftpServer(i["tftp-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "timezone"
	if _, ok := i["timezone"]; ok {
		result["timezone"] = flattenObjectFspVlanDynamicMappingDhcpServerTimezone(i["timezone"], d, pre_append)
	}

	pre_append = pre + ".0." + "timezone_option"
	if _, ok := i["timezone-option"]; ok {
		result["timezone_option"] = flattenObjectFspVlanDynamicMappingDhcpServerTimezoneOption(i["timezone-option"], d, pre_append)
	}

	pre_append = pre + ".0." + "vci_match"
	if _, ok := i["vci-match"]; ok {
		result["vci_match"] = flattenObjectFspVlanDynamicMappingDhcpServerVciMatch(i["vci-match"], d, pre_append)
	}

	pre_append = pre + ".0." + "vci_string"
	if _, ok := i["vci-string"]; ok {
		result["vci_string"] = flattenObjectFspVlanDynamicMappingDhcpServerVciString(i["vci-string"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac_service"
	if _, ok := i["wifi-ac-service"]; ok {
		result["wifi_ac_service"] = flattenObjectFspVlanDynamicMappingDhcpServerWifiAcService(i["wifi-ac-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac1"
	if _, ok := i["wifi-ac1"]; ok {
		result["wifi_ac1"] = flattenObjectFspVlanDynamicMappingDhcpServerWifiAc1(i["wifi-ac1"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac2"
	if _, ok := i["wifi-ac2"]; ok {
		result["wifi_ac2"] = flattenObjectFspVlanDynamicMappingDhcpServerWifiAc2(i["wifi-ac2"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ac3"
	if _, ok := i["wifi-ac3"]; ok {
		result["wifi_ac3"] = flattenObjectFspVlanDynamicMappingDhcpServerWifiAc3(i["wifi-ac3"], d, pre_append)
	}

	pre_append = pre + ".0." + "wins_server1"
	if _, ok := i["wins-server1"]; ok {
		result["wins_server1"] = flattenObjectFspVlanDynamicMappingDhcpServerWinsServer1(i["wins-server1"], d, pre_append)
	}

	pre_append = pre + ".0." + "wins_server2"
	if _, ok := i["wins-server2"]; ok {
		result["wins_server2"] = flattenObjectFspVlanDynamicMappingDhcpServerWinsServer2(i["wins-server2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFspVlanDynamicMappingDhcpServerAutoConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsKeyname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDefaultGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-VciString")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-VciString")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNextServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsCode(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Value")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-VciString")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := i["circuit-id-type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType(i["circuit-id-type"], d, pre_append)
			tmp["circuit_id_type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-CircuitIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := i["remote-id-type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType(i["remote-id-type"], d, pre_append)
			tmp["remote_id_type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-RemoteIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerTftpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerTimezoneOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAcService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAc1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAc2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAc3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dhcp_relay_agent_option"
	if _, ok := i["dhcp-relay-agent-option"]; ok {
		result["dhcp_relay_agent_option"] = flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayAgentOption(i["dhcp-relay-agent-option"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_interface_select_method"
	if _, ok := i["dhcp-relay-interface-select-method"]; ok {
		result["dhcp_relay_interface_select_method"] = flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayInterfaceSelectMethod(i["dhcp-relay-interface-select-method"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_ip"
	if _, ok := i["dhcp-relay-ip"]; ok {
		result["dhcp_relay_ip"] = flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayIp(i["dhcp-relay-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_service"
	if _, ok := i["dhcp-relay-service"]; ok {
		result["dhcp_relay_service"] = flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayService(i["dhcp-relay-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_type"
	if _, ok := i["dhcp-relay-type"]; ok {
		result["dhcp_relay_type"] = flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayType(i["dhcp-relay-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenObjectFspVlanDynamicMappingInterfaceIp(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6"
	if _, ok := i["ipv6"]; ok {
		result["ipv6"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6(i["ipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "secondary_IP"
	if _, ok := i["secondary-IP"]; ok {
		result["secondary_IP"] = flattenObjectFspVlanDynamicMappingInterfaceSecondaryIp(i["secondary-IP"], d, pre_append)
	}

	pre_append = pre + ".0." + "secondaryip"
	if _, ok := i["secondaryip"]; ok {
		result["secondaryip"] = flattenObjectFspVlanDynamicMappingInterfaceSecondaryip(i["secondaryip"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlanid"
	if _, ok := i["vlanid"]; ok {
		result["vlanid"] = flattenObjectFspVlanDynamicMappingInterfaceVlanid(i["vlanid"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayAgentOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceDhcpRelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "autoconf"
	if _, ok := i["autoconf"]; ok {
		result["autoconf"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Autoconf(i["autoconf"], d, pre_append)
	}

	pre_append = pre + ".0." + "cli_conn6_status"
	if _, ok := i["cli-conn6-status"]; ok {
		result["cli_conn6_status"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6CliConn6Status(i["cli-conn6-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := i["dhcp6-client-options"]; ok {
		result["dhcp6_client_options"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6ClientOptions(i["dhcp6-client-options"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := i["dhcp6-information-request"]; ok {
		result["dhcp6_information_request"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6InformationRequest(i["dhcp6-information-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := i["dhcp6-prefix-delegation"]; ok {
		result["dhcp6_prefix_delegation"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixDelegation(i["dhcp6-prefix-delegation"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := i["dhcp6-prefix-hint"]; ok {
		result["dhcp6_prefix_hint"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHint(i["dhcp6-prefix-hint"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := i["dhcp6-prefix-hint-plt"]; ok {
		result["dhcp6_prefix_hint_plt"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintPlt(i["dhcp6-prefix-hint-plt"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := i["dhcp6-prefix-hint-vlt"]; ok {
		result["dhcp6_prefix_hint_vlt"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintVlt(i["dhcp6-prefix-hint-vlt"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := i["dhcp6-relay-ip"]; ok {
		result["dhcp6_relay_ip"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayIp(i["dhcp6-relay-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := i["dhcp6-relay-service"]; ok {
		result["dhcp6_relay_service"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayService(i["dhcp6-relay-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := i["dhcp6-relay-type"]; ok {
		result["dhcp6_relay_type"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayType(i["dhcp6-relay-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp6_send_redirect"
	if _, ok := i["icmp6-send-redirect"]; ok {
		result["icmp6_send_redirect"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Icmp6SendRedirect(i["icmp6-send-redirect"], d, pre_append)
	}

	pre_append = pre + ".0." + "interface_identifier"
	if _, ok := i["interface-identifier"]; ok {
		result["interface_identifier"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6InterfaceIdentifier(i["interface-identifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_address"
	if _, ok := i["ip6-address"]; ok {
		result["ip6_address"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Address(i["ip6-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := i["ip6-allowaccess"]; ok {
		result["ip6_allowaccess"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Allowaccess(i["ip6-allowaccess"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := i["ip6-default-life"]; ok {
		result["ip6_default_life"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DefaultLife(i["ip6-default-life"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_iaid"
	if _, ok := i["ip6-delegated-prefix-iaid"]; ok {
		result["ip6_delegated_prefix_iaid"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixIaid(i["ip6-delegated-prefix-iaid"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := i["ip6-delegated-prefix-list"]; ok {
		result["ip6_delegated_prefix_list"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixList(i["ip6-delegated-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := i["ip6-dns-server-override"]; ok {
		result["ip6_dns_server_override"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DnsServerOverride(i["ip6-dns-server-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := i["ip6-extra-addr"]; ok {
		result["ip6_extra_addr"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddr(i["ip6-extra-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := i["ip6-hop-limit"]; ok {
		result["ip6_hop_limit"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6HopLimit(i["ip6-hop-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := i["ip6-link-mtu"]; ok {
		result["ip6_link_mtu"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6LinkMtu(i["ip6-link-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := i["ip6-manage-flag"]; ok {
		result["ip6_manage_flag"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ManageFlag(i["ip6-manage-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := i["ip6-max-interval"]; ok {
		result["ip6_max_interval"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6MaxInterval(i["ip6-max-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := i["ip6-min-interval"]; ok {
		result["ip6_min_interval"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6MinInterval(i["ip6-min-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := i["ip6-mode"]; ok {
		result["ip6_mode"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Mode(i["ip6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := i["ip6-other-flag"]; ok {
		result["ip6_other_flag"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6OtherFlag(i["ip6-other-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := i["ip6-prefix-list"]; ok {
		result["ip6_prefix_list"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixList(i["ip6-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_mode"
	if _, ok := i["ip6-prefix-mode"]; ok {
		result["ip6_prefix_mode"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixMode(i["ip6-prefix-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := i["ip6-reachable-time"]; ok {
		result["ip6_reachable_time"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ReachableTime(i["ip6-reachable-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := i["ip6-retrans-time"]; ok {
		result["ip6_retrans_time"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6RetransTime(i["ip6-retrans-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := i["ip6-send-adv"]; ok {
		result["ip6_send_adv"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6SendAdv(i["ip6-send-adv"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := i["ip6-subnet"]; ok {
		result["ip6_subnet"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Subnet(i["ip6-subnet"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := i["ip6-upstream-interface"]; ok {
		result["ip6_upstream_interface"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6UpstreamInterface(i["ip6-upstream-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cert"
	if _, ok := i["nd-cert"]; ok {
		result["nd_cert"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdCert(i["nd-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := i["nd-cga-modifier"]; ok {
		result["nd_cga_modifier"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdCgaModifier(i["nd-cga-modifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_mode"
	if _, ok := i["nd-mode"]; ok {
		result["nd_mode"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdMode(i["nd-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := i["nd-security-level"]; ok {
		result["nd_security_level"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdSecurityLevel(i["nd-security-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := i["nd-timestamp-delta"]; ok {
		result["nd_timestamp_delta"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampDelta(i["nd-timestamp-delta"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := i["nd-timestamp-fuzz"]; ok {
		result["nd_timestamp_fuzz"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampFuzz(i["nd-timestamp-fuzz"], d, pre_append)
	}

	pre_append = pre + ".0." + "ra_send_mtu"
	if _, ok := i["ra-send-mtu"]; ok {
		result["ra_send_mtu"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6RaSendMtu(i["ra-send-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "unique_autoconf_addr"
	if _, ok := i["unique-autoconf-addr"]; ok {
		result["unique_autoconf_addr"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6UniqueAutoconfAddr(i["unique-autoconf-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := i["vrip6_link_local"]; ok {
		result["vrip6_link_local"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrip6LinkLocal(i["vrip6_link_local"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := i["vrrp-virtual-mac6"]; ok {
		result["vrrp_virtual_mac6"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6VrrpVirtualMac6(i["vrrp-virtual-mac6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp6"
	if _, ok := i["vrrp6"]; ok {
		result["vrrp6"] = flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6(i["vrrp6"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Autoconf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6CliConn6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6ClientOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6InformationRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixDelegation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintPlt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintVlt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Icmp6SendRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6InterfaceIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DefaultLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixIaid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := i["delegated-prefix-iaid"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid(i["delegated-prefix-iaid"], d, pre_append)
			tmp["delegated_prefix_iaid"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-DelegatedPrefixIaid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := i["prefix-id"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListPrefixId(i["prefix-id"], d, pre_append)
			tmp["prefix_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-PrefixId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnss(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := i["rdnss-service"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnssService(i["rdnss-service"], d, pre_append)
			tmp["rdnss_service"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-RdnssService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListSubnet(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := i["upstream-interface"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(i["upstream-interface"], d, pre_append)
			tmp["upstream_interface"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6DelegatedPrefixList-UpstreamInterface")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListPrefixId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnssService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6DnsServerOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddrPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6ExtraAddr-Prefix")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddrPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6HopLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6LinkMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ManageFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6MaxInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6MinInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6OtherFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListAutonomousFlag(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := i["dnssl"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListDnssl(i["dnssl"], d, pre_append)
			tmp["dnssl"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-Dnssl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListOnlinkFlag(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := i["preferred-life-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPreferredLifeTime(i["preferred-life-time"], d, pre_append)
			tmp["preferred_life_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-PreferredLifeTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListRdnss(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := i["valid-life-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListValidLifeTime(i["valid-life-time"], d, pre_append)
			tmp["valid_life_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Ip6PrefixList-ValidLifeTime")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListAutonomousFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListDnssl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListOnlinkFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPreferredLifeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListRdnss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListValidLifeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6ReachableTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6RetransTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6SendAdv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Ip6UpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdCgaModifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampDelta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampFuzz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6RaSendMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6UniqueAutoconfAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrip6LinkLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6VrrpVirtualMac6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AcceptMode(i["accept-mode"], d, pre_append)
			tmp["accept_mode"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-AcceptMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AdvInterval(i["adv-interval"], d, pre_append)
			tmp["adv_interval"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-AdvInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Preempt(i["preempt"], d, pre_append)
			tmp["preempt"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Preempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Priority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6StartTime(i["start-time"], d, pre_append)
			tmp["start_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-StartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Status(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := i["vrdst6"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrdst6(i["vrdst6"], d, pre_append)
			tmp["vrdst6"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Vrdst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrgrp(i["vrgrp"], d, pre_append)
			tmp["vrgrp"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Vrgrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrid(i["vrid"], d, pre_append)
			tmp["vrid"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Vrid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := i["vrip6"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrip6(i["vrip6"], d, pre_append)
			tmp["vrip6"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterfaceIpv6-Vrrp6-Vrip6")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AcceptMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AdvInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Preempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6StartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrdst6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryip(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipAllowaccess(i["allowaccess"], d, pre_append)
			tmp["allowaccess"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Allowaccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := i["detectprotocol"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipDetectprotocol(i["detectprotocol"], d, pre_append)
			tmp["detectprotocol"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Detectprotocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := i["detectserver"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipDetectserver(i["detectserver"], d, pre_append)
			tmp["detectserver"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Detectserver")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := i["gwdetect"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipGwdetect(i["gwdetect"], d, pre_append)
			tmp["gwdetect"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Gwdetect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipHaPriority(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := i["ping-serv-status"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipPingServStatus(i["ping-serv-status"], d, pre_append)
			tmp["ping_serv_status"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-PingServStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenObjectFspVlanDynamicMappingInterfaceSecondaryipSeq(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingInterface-Secondaryip-Seq")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipDetectprotocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipDetectserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipGwdetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipPingServStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceSecondaryipSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingInterfaceVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vlan_op_mode"
	if _, ok := i["vlan-op-mode"]; ok {
		result["vlan_op_mode"] = flattenObjectFspVlanInterfaceVlanOpMode(i["vlan-op-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ac_name"
	if _, ok := i["ac-name"]; ok {
		result["ac_name"] = flattenObjectFspVlanInterfaceAcName(i["ac-name"], d, pre_append)
	}

	pre_append = pre + ".0." + "aggregate"
	if _, ok := i["aggregate"]; ok {
		result["aggregate"] = flattenObjectFspVlanInterfaceAggregate(i["aggregate"], d, pre_append)
	}

	pre_append = pre + ".0." + "aggregate_type"
	if _, ok := i["aggregate-type"]; ok {
		result["aggregate_type"] = flattenObjectFspVlanInterfaceAggregateType(i["aggregate-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "algorithm"
	if _, ok := i["algorithm"]; ok {
		result["algorithm"] = flattenObjectFspVlanInterfaceAlgorithm(i["algorithm"], d, pre_append)
	}

	pre_append = pre + ".0." + "alias"
	if _, ok := i["alias"]; ok {
		result["alias"] = flattenObjectFspVlanInterfaceAlias(i["alias"], d, pre_append)
	}

	pre_append = pre + ".0." + "allowaccess"
	if _, ok := i["allowaccess"]; ok {
		result["allowaccess"] = flattenObjectFspVlanInterfaceAllowaccess(i["allowaccess"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_discover"
	if _, ok := i["ap-discover"]; ok {
		result["ap_discover"] = flattenObjectFspVlanInterfaceApDiscover(i["ap-discover"], d, pre_append)
	}

	pre_append = pre + ".0." + "arpforward"
	if _, ok := i["arpforward"]; ok {
		result["arpforward"] = flattenObjectFspVlanInterfaceArpforward(i["arpforward"], d, pre_append)
	}

	pre_append = pre + ".0." + "atm_protocol"
	if _, ok := i["atm-protocol"]; ok {
		result["atm_protocol"] = flattenObjectFspVlanInterfaceAtmProtocol(i["atm-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_cert"
	if _, ok := i["auth-cert"]; ok {
		result["auth_cert"] = flattenObjectFspVlanInterfaceAuthCert(i["auth-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_portal_addr"
	if _, ok := i["auth-portal-addr"]; ok {
		result["auth_portal_addr"] = flattenObjectFspVlanInterfaceAuthPortalAddr(i["auth-portal-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_type"
	if _, ok := i["auth-type"]; ok {
		result["auth_type"] = flattenObjectFspVlanInterfaceAuthType(i["auth-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_auth_extension_device"
	if _, ok := i["auto-auth-extension-device"]; ok {
		result["auto_auth_extension_device"] = flattenObjectFspVlanInterfaceAutoAuthExtensionDevice(i["auto-auth-extension-device"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_measure_time"
	if _, ok := i["bandwidth-measure-time"]; ok {
		result["bandwidth_measure_time"] = flattenObjectFspVlanInterfaceBandwidthMeasureTime(i["bandwidth-measure-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "bfd"
	if _, ok := i["bfd"]; ok {
		result["bfd"] = flattenObjectFspVlanInterfaceBfd(i["bfd"], d, pre_append)
	}

	pre_append = pre + ".0." + "bfd_desired_min_tx"
	if _, ok := i["bfd-desired-min-tx"]; ok {
		result["bfd_desired_min_tx"] = flattenObjectFspVlanInterfaceBfdDesiredMinTx(i["bfd-desired-min-tx"], d, pre_append)
	}

	pre_append = pre + ".0." + "bfd_detect_mult"
	if _, ok := i["bfd-detect-mult"]; ok {
		result["bfd_detect_mult"] = flattenObjectFspVlanInterfaceBfdDetectMult(i["bfd-detect-mult"], d, pre_append)
	}

	pre_append = pre + ".0." + "bfd_required_min_rx"
	if _, ok := i["bfd-required-min-rx"]; ok {
		result["bfd_required_min_rx"] = flattenObjectFspVlanInterfaceBfdRequiredMinRx(i["bfd-required-min-rx"], d, pre_append)
	}

	pre_append = pre + ".0." + "broadcast_forticlient_discovery"
	if _, ok := i["broadcast-forticlient-discovery"]; ok {
		result["broadcast_forticlient_discovery"] = flattenObjectFspVlanInterfaceBroadcastForticlientDiscovery(i["broadcast-forticlient-discovery"], d, pre_append)
	}

	pre_append = pre + ".0." + "broadcast_forward"
	if _, ok := i["broadcast-forward"]; ok {
		result["broadcast_forward"] = flattenObjectFspVlanInterfaceBroadcastForward(i["broadcast-forward"], d, pre_append)
	}

	pre_append = pre + ".0." + "captive_portal"
	if _, ok := i["captive-portal"]; ok {
		result["captive_portal"] = flattenObjectFspVlanInterfaceCaptivePortal(i["captive-portal"], d, pre_append)
	}

	pre_append = pre + ".0." + "cli_conn_status"
	if _, ok := i["cli-conn-status"]; ok {
		result["cli_conn_status"] = flattenObjectFspVlanInterfaceCliConnStatus(i["cli-conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "color"
	if _, ok := i["color"]; ok {
		result["color"] = flattenObjectFspVlanInterfaceColor(i["color"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns"
	if _, ok := i["ddns"]; ok {
		result["ddns"] = flattenObjectFspVlanInterfaceDdns(i["ddns"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_auth"
	if _, ok := i["ddns-auth"]; ok {
		result["ddns_auth"] = flattenObjectFspVlanInterfaceDdnsAuth(i["ddns-auth"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_domain"
	if _, ok := i["ddns-domain"]; ok {
		result["ddns_domain"] = flattenObjectFspVlanInterfaceDdnsDomain(i["ddns-domain"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_key"
	if _, ok := i["ddns-key"]; ok {
		result["ddns_key"] = flattenObjectFspVlanInterfaceDdnsKey(i["ddns-key"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_keyname"
	if _, ok := i["ddns-keyname"]; ok {
		result["ddns_keyname"] = flattenObjectFspVlanInterfaceDdnsKeyname(i["ddns-keyname"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_password"
	if _, ok := i["ddns-password"]; ok {
		result["ddns_password"] = flattenObjectFspVlanInterfaceDdnsPassword(i["ddns-password"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_server"
	if _, ok := i["ddns-server"]; ok {
		result["ddns_server"] = flattenObjectFspVlanInterfaceDdnsServer(i["ddns-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_server_ip"
	if _, ok := i["ddns-server-ip"]; ok {
		result["ddns_server_ip"] = flattenObjectFspVlanInterfaceDdnsServerIp(i["ddns-server-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_sn"
	if _, ok := i["ddns-sn"]; ok {
		result["ddns_sn"] = flattenObjectFspVlanInterfaceDdnsSn(i["ddns-sn"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_ttl"
	if _, ok := i["ddns-ttl"]; ok {
		result["ddns_ttl"] = flattenObjectFspVlanInterfaceDdnsTtl(i["ddns-ttl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_username"
	if _, ok := i["ddns-username"]; ok {
		result["ddns_username"] = flattenObjectFspVlanInterfaceDdnsUsername(i["ddns-username"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddns_zone"
	if _, ok := i["ddns-zone"]; ok {
		result["ddns_zone"] = flattenObjectFspVlanInterfaceDdnsZone(i["ddns-zone"], d, pre_append)
	}

	pre_append = pre + ".0." + "dedicated_to"
	if _, ok := i["dedicated-to"]; ok {
		result["dedicated_to"] = flattenObjectFspVlanInterfaceDedicatedTo(i["dedicated-to"], d, pre_append)
	}

	pre_append = pre + ".0." + "defaultgw"
	if _, ok := i["defaultgw"]; ok {
		result["defaultgw"] = flattenObjectFspVlanInterfaceDefaultgw(i["defaultgw"], d, pre_append)
	}

	pre_append = pre + ".0." + "description"
	if _, ok := i["description"]; ok {
		result["description"] = flattenObjectFspVlanInterfaceDescription(i["description"], d, pre_append)
	}

	pre_append = pre + ".0." + "detected_peer_mtu"
	if _, ok := i["detected-peer-mtu"]; ok {
		result["detected_peer_mtu"] = flattenObjectFspVlanInterfaceDetectedPeerMtu(i["detected-peer-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "detectprotocol"
	if _, ok := i["detectprotocol"]; ok {
		result["detectprotocol"] = flattenObjectFspVlanInterfaceDetectprotocol(i["detectprotocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "detectserver"
	if _, ok := i["detectserver"]; ok {
		result["detectserver"] = flattenObjectFspVlanInterfaceDetectserver(i["detectserver"], d, pre_append)
	}

	pre_append = pre + ".0." + "device_access_list"
	if _, ok := i["device-access-list"]; ok {
		result["device_access_list"] = flattenObjectFspVlanInterfaceDeviceAccessList(i["device-access-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "device_identification"
	if _, ok := i["device-identification"]; ok {
		result["device_identification"] = flattenObjectFspVlanInterfaceDeviceIdentification(i["device-identification"], d, pre_append)
	}

	pre_append = pre + ".0." + "device_identification_active_scan"
	if _, ok := i["device-identification-active-scan"]; ok {
		result["device_identification_active_scan"] = flattenObjectFspVlanInterfaceDeviceIdentificationActiveScan(i["device-identification-active-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "device_netscan"
	if _, ok := i["device-netscan"]; ok {
		result["device_netscan"] = flattenObjectFspVlanInterfaceDeviceNetscan(i["device-netscan"], d, pre_append)
	}

	pre_append = pre + ".0." + "device_user_identification"
	if _, ok := i["device-user-identification"]; ok {
		result["device_user_identification"] = flattenObjectFspVlanInterfaceDeviceUserIdentification(i["device-user-identification"], d, pre_append)
	}

	pre_append = pre + ".0." + "devindex"
	if _, ok := i["devindex"]; ok {
		result["devindex"] = flattenObjectFspVlanInterfaceDevindex(i["devindex"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_classless_route_addition"
	if _, ok := i["dhcp-classless-route-addition"]; ok {
		result["dhcp_classless_route_addition"] = flattenObjectFspVlanInterfaceDhcpClasslessRouteAddition(i["dhcp-classless-route-addition"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_client_identifier"
	if _, ok := i["dhcp-client-identifier"]; ok {
		result["dhcp_client_identifier"] = flattenObjectFspVlanInterfaceDhcpClientIdentifier(i["dhcp-client-identifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_agent_option"
	if _, ok := i["dhcp-relay-agent-option"]; ok {
		result["dhcp_relay_agent_option"] = flattenObjectFspVlanInterfaceDhcpRelayAgentOption(i["dhcp-relay-agent-option"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_interface"
	if _, ok := i["dhcp-relay-interface"]; ok {
		result["dhcp_relay_interface"] = flattenObjectFspVlanInterfaceDhcpRelayInterface(i["dhcp-relay-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_interface_select_method"
	if _, ok := i["dhcp-relay-interface-select-method"]; ok {
		result["dhcp_relay_interface_select_method"] = flattenObjectFspVlanInterfaceDhcpRelayInterfaceSelectMethod(i["dhcp-relay-interface-select-method"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_ip"
	if _, ok := i["dhcp-relay-ip"]; ok {
		result["dhcp_relay_ip"] = flattenObjectFspVlanInterfaceDhcpRelayIp(i["dhcp-relay-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_link_selection"
	if _, ok := i["dhcp-relay-link-selection"]; ok {
		result["dhcp_relay_link_selection"] = flattenObjectFspVlanInterfaceDhcpRelayLinkSelection(i["dhcp-relay-link-selection"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_request_all_server"
	if _, ok := i["dhcp-relay-request-all-server"]; ok {
		result["dhcp_relay_request_all_server"] = flattenObjectFspVlanInterfaceDhcpRelayRequestAllServer(i["dhcp-relay-request-all-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_service"
	if _, ok := i["dhcp-relay-service"]; ok {
		result["dhcp_relay_service"] = flattenObjectFspVlanInterfaceDhcpRelayService(i["dhcp-relay-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_relay_type"
	if _, ok := i["dhcp-relay-type"]; ok {
		result["dhcp_relay_type"] = flattenObjectFspVlanInterfaceDhcpRelayType(i["dhcp-relay-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp_renew_time"
	if _, ok := i["dhcp-renew-time"]; ok {
		result["dhcp_renew_time"] = flattenObjectFspVlanInterfaceDhcpRenewTime(i["dhcp-renew-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "disc_retry_timeout"
	if _, ok := i["disc-retry-timeout"]; ok {
		result["disc_retry_timeout"] = flattenObjectFspVlanInterfaceDiscRetryTimeout(i["disc-retry-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenObjectFspVlanInterfaceDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "distance"
	if _, ok := i["distance"]; ok {
		result["distance"] = flattenObjectFspVlanInterfaceDistance(i["distance"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_query"
	if _, ok := i["dns-query"]; ok {
		result["dns_query"] = flattenObjectFspVlanInterfaceDnsQuery(i["dns-query"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server_override"
	if _, ok := i["dns-server-override"]; ok {
		result["dns_server_override"] = flattenObjectFspVlanInterfaceDnsServerOverride(i["dns-server-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_server_protocol"
	if _, ok := i["dns-server-protocol"]; ok {
		result["dns_server_protocol"] = flattenObjectFspVlanInterfaceDnsServerProtocol(i["dns-server-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "drop_fragment"
	if _, ok := i["drop-fragment"]; ok {
		result["drop_fragment"] = flattenObjectFspVlanInterfaceDropFragment(i["drop-fragment"], d, pre_append)
	}

	pre_append = pre + ".0." + "drop_overlapped_fragment"
	if _, ok := i["drop-overlapped-fragment"]; ok {
		result["drop_overlapped_fragment"] = flattenObjectFspVlanInterfaceDropOverlappedFragment(i["drop-overlapped-fragment"], d, pre_append)
	}

	pre_append = pre + ".0." + "eap_ca_cert"
	if _, ok := i["eap-ca-cert"]; ok {
		result["eap_ca_cert"] = flattenObjectFspVlanInterfaceEapCaCert(i["eap-ca-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "eap_identity"
	if _, ok := i["eap-identity"]; ok {
		result["eap_identity"] = flattenObjectFspVlanInterfaceEapIdentity(i["eap-identity"], d, pre_append)
	}

	pre_append = pre + ".0." + "eap_method"
	if _, ok := i["eap-method"]; ok {
		result["eap_method"] = flattenObjectFspVlanInterfaceEapMethod(i["eap-method"], d, pre_append)
	}

	pre_append = pre + ".0." + "eap_password"
	if _, ok := i["eap-password"]; ok {
		result["eap_password"] = flattenObjectFspVlanInterfaceEapPassword(i["eap-password"], d, pre_append)
	}

	pre_append = pre + ".0." + "eap_supplicant"
	if _, ok := i["eap-supplicant"]; ok {
		result["eap_supplicant"] = flattenObjectFspVlanInterfaceEapSupplicant(i["eap-supplicant"], d, pre_append)
	}

	pre_append = pre + ".0." + "eap_user_cert"
	if _, ok := i["eap-user-cert"]; ok {
		result["eap_user_cert"] = flattenObjectFspVlanInterfaceEapUserCert(i["eap-user-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "egress_cos"
	if _, ok := i["egress-cos"]; ok {
		result["egress_cos"] = flattenObjectFspVlanInterfaceEgressCos(i["egress-cos"], d, pre_append)
	}

	pre_append = pre + ".0." + "egress_shaping_profile"
	if _, ok := i["egress-shaping-profile"]; ok {
		result["egress_shaping_profile"] = flattenObjectFspVlanInterfaceEgressShapingProfile(i["egress-shaping-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "eip"
	if _, ok := i["eip"]; ok {
		result["eip"] = flattenObjectFspVlanInterfaceEip(i["eip"], d, pre_append)
	}

	pre_append = pre + ".0." + "endpoint_compliance"
	if _, ok := i["endpoint-compliance"]; ok {
		result["endpoint_compliance"] = flattenObjectFspVlanInterfaceEndpointCompliance(i["endpoint-compliance"], d, pre_append)
	}

	pre_append = pre + ".0." + "estimated_downstream_bandwidth"
	if _, ok := i["estimated-downstream-bandwidth"]; ok {
		result["estimated_downstream_bandwidth"] = flattenObjectFspVlanInterfaceEstimatedDownstreamBandwidth(i["estimated-downstream-bandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "estimated_upstream_bandwidth"
	if _, ok := i["estimated-upstream-bandwidth"]; ok {
		result["estimated_upstream_bandwidth"] = flattenObjectFspVlanInterfaceEstimatedUpstreamBandwidth(i["estimated-upstream-bandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "explicit_ftp_proxy"
	if _, ok := i["explicit-ftp-proxy"]; ok {
		result["explicit_ftp_proxy"] = flattenObjectFspVlanInterfaceExplicitFtpProxy(i["explicit-ftp-proxy"], d, pre_append)
	}

	pre_append = pre + ".0." + "explicit_web_proxy"
	if _, ok := i["explicit-web-proxy"]; ok {
		result["explicit_web_proxy"] = flattenObjectFspVlanInterfaceExplicitWebProxy(i["explicit-web-proxy"], d, pre_append)
	}

	pre_append = pre + ".0." + "external"
	if _, ok := i["external"]; ok {
		result["external"] = flattenObjectFspVlanInterfaceExternal(i["external"], d, pre_append)
	}

	pre_append = pre + ".0." + "fail_action_on_extender"
	if _, ok := i["fail-action-on-extender"]; ok {
		result["fail_action_on_extender"] = flattenObjectFspVlanInterfaceFailActionOnExtender(i["fail-action-on-extender"], d, pre_append)
	}

	pre_append = pre + ".0." + "fail_alert_interfaces"
	if _, ok := i["fail-alert-interfaces"]; ok {
		result["fail_alert_interfaces"] = flattenObjectFspVlanInterfaceFailAlertInterfaces(i["fail-alert-interfaces"], d, pre_append)
	}

	pre_append = pre + ".0." + "fail_alert_method"
	if _, ok := i["fail-alert-method"]; ok {
		result["fail_alert_method"] = flattenObjectFspVlanInterfaceFailAlertMethod(i["fail-alert-method"], d, pre_append)
	}

	pre_append = pre + ".0." + "fail_detect"
	if _, ok := i["fail-detect"]; ok {
		result["fail_detect"] = flattenObjectFspVlanInterfaceFailDetect(i["fail-detect"], d, pre_append)
	}

	pre_append = pre + ".0." + "fail_detect_option"
	if _, ok := i["fail-detect-option"]; ok {
		result["fail_detect_option"] = flattenObjectFspVlanInterfaceFailDetectOption(i["fail-detect-option"], d, pre_append)
	}

	pre_append = pre + ".0." + "fdp"
	if _, ok := i["fdp"]; ok {
		result["fdp"] = flattenObjectFspVlanInterfaceFdp(i["fdp"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiheartbeat"
	if _, ok := i["fortiheartbeat"]; ok {
		result["fortiheartbeat"] = flattenObjectFspVlanInterfaceFortiheartbeat(i["fortiheartbeat"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortilink"
	if _, ok := i["fortilink"]; ok {
		result["fortilink"] = flattenObjectFspVlanInterfaceFortilink(i["fortilink"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortilink_backup_link"
	if _, ok := i["fortilink-backup-link"]; ok {
		result["fortilink_backup_link"] = flattenObjectFspVlanInterfaceFortilinkBackupLink(i["fortilink-backup-link"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortilink_neighbor_detect"
	if _, ok := i["fortilink-neighbor-detect"]; ok {
		result["fortilink_neighbor_detect"] = flattenObjectFspVlanInterfaceFortilinkNeighborDetect(i["fortilink-neighbor-detect"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortilink_split_interface"
	if _, ok := i["fortilink-split-interface"]; ok {
		result["fortilink_split_interface"] = flattenObjectFspVlanInterfaceFortilinkSplitInterface(i["fortilink-split-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortilink_stacking"
	if _, ok := i["fortilink-stacking"]; ok {
		result["fortilink_stacking"] = flattenObjectFspVlanInterfaceFortilinkStacking(i["fortilink-stacking"], d, pre_append)
	}

	pre_append = pre + ".0." + "forward_domain"
	if _, ok := i["forward-domain"]; ok {
		result["forward_domain"] = flattenObjectFspVlanInterfaceForwardDomain(i["forward-domain"], d, pre_append)
	}

	pre_append = pre + ".0." + "forward_error_correction"
	if _, ok := i["forward-error-correction"]; ok {
		result["forward_error_correction"] = flattenObjectFspVlanInterfaceForwardErrorCorrection(i["forward-error-correction"], d, pre_append)
	}

	pre_append = pre + ".0." + "fp_anomaly"
	if _, ok := i["fp-anomaly"]; ok {
		result["fp_anomaly"] = flattenObjectFspVlanInterfaceFpAnomaly(i["fp-anomaly"], d, pre_append)
	}

	pre_append = pre + ".0." + "fp_disable"
	if _, ok := i["fp-disable"]; ok {
		result["fp_disable"] = flattenObjectFspVlanInterfaceFpDisable(i["fp-disable"], d, pre_append)
	}

	pre_append = pre + ".0." + "gateway_address"
	if _, ok := i["gateway-address"]; ok {
		result["gateway_address"] = flattenObjectFspVlanInterfaceGatewayAddress(i["gateway-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "generic_receive_offload"
	if _, ok := i["generic-receive-offload"]; ok {
		result["generic_receive_offload"] = flattenObjectFspVlanInterfaceGenericReceiveOffload(i["generic-receive-offload"], d, pre_append)
	}

	pre_append = pre + ".0." + "gi_gk"
	if _, ok := i["gi-gk"]; ok {
		result["gi_gk"] = flattenObjectFspVlanInterfaceGiGk(i["gi-gk"], d, pre_append)
	}

	pre_append = pre + ".0." + "gwaddr"
	if _, ok := i["gwaddr"]; ok {
		result["gwaddr"] = flattenObjectFspVlanInterfaceGwaddr(i["gwaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "gwdetect"
	if _, ok := i["gwdetect"]; ok {
		result["gwdetect"] = flattenObjectFspVlanInterfaceGwdetect(i["gwdetect"], d, pre_append)
	}

	pre_append = pre + ".0." + "ha_priority"
	if _, ok := i["ha-priority"]; ok {
		result["ha_priority"] = flattenObjectFspVlanInterfaceHaPriority(i["ha-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_accept_redirect"
	if _, ok := i["icmp-accept-redirect"]; ok {
		result["icmp_accept_redirect"] = flattenObjectFspVlanInterfaceIcmpAcceptRedirect(i["icmp-accept-redirect"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_redirect"
	if _, ok := i["icmp-redirect"]; ok {
		result["icmp_redirect"] = flattenObjectFspVlanInterfaceIcmpRedirect(i["icmp-redirect"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_send_redirect"
	if _, ok := i["icmp-send-redirect"]; ok {
		result["icmp_send_redirect"] = flattenObjectFspVlanInterfaceIcmpSendRedirect(i["icmp-send-redirect"], d, pre_append)
	}

	pre_append = pre + ".0." + "ident_accept"
	if _, ok := i["ident-accept"]; ok {
		result["ident_accept"] = flattenObjectFspVlanInterfaceIdentAccept(i["ident-accept"], d, pre_append)
	}

	pre_append = pre + ".0." + "idle_timeout"
	if _, ok := i["idle-timeout"]; ok {
		result["idle_timeout"] = flattenObjectFspVlanInterfaceIdleTimeout(i["idle-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "if_mdix"
	if _, ok := i["if-mdix"]; ok {
		result["if_mdix"] = flattenObjectFspVlanInterfaceIfMdix(i["if-mdix"], d, pre_append)
	}

	pre_append = pre + ".0." + "if_media"
	if _, ok := i["if-media"]; ok {
		result["if_media"] = flattenObjectFspVlanInterfaceIfMedia(i["if-media"], d, pre_append)
	}

	pre_append = pre + ".0." + "ike_saml_server"
	if _, ok := i["ike-saml-server"]; ok {
		result["ike_saml_server"] = flattenObjectFspVlanInterfaceIkeSamlServer(i["ike-saml-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "in_force_vlan_cos"
	if _, ok := i["in-force-vlan-cos"]; ok {
		result["in_force_vlan_cos"] = flattenObjectFspVlanInterfaceInForceVlanCos(i["in-force-vlan-cos"], d, pre_append)
	}

	pre_append = pre + ".0." + "inbandwidth"
	if _, ok := i["inbandwidth"]; ok {
		result["inbandwidth"] = flattenObjectFspVlanInterfaceInbandwidth(i["inbandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "ingress_cos"
	if _, ok := i["ingress-cos"]; ok {
		result["ingress_cos"] = flattenObjectFspVlanInterfaceIngressCos(i["ingress-cos"], d, pre_append)
	}

	pre_append = pre + ".0." + "ingress_shaping_profile"
	if _, ok := i["ingress-shaping-profile"]; ok {
		result["ingress_shaping_profile"] = flattenObjectFspVlanInterfaceIngressShapingProfile(i["ingress-shaping-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "ingress_spillover_threshold"
	if _, ok := i["ingress-spillover-threshold"]; ok {
		result["ingress_spillover_threshold"] = flattenObjectFspVlanInterfaceIngressSpilloverThreshold(i["ingress-spillover-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "interconnect_profile"
	if _, ok := i["interconnect-profile"]; ok {
		result["interconnect_profile"] = flattenObjectFspVlanInterfaceInterconnectProfile(i["interconnect-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "internal"
	if _, ok := i["internal"]; ok {
		result["internal"] = flattenObjectFspVlanInterfaceInternal(i["internal"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenObjectFspVlanInterfaceIp(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_managed_by_fortiipam"
	if _, ok := i["ip-managed-by-fortiipam"]; ok {
		result["ip_managed_by_fortiipam"] = flattenObjectFspVlanInterfaceIpManagedByFortiipam(i["ip-managed-by-fortiipam"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipmac"
	if _, ok := i["ipmac"]; ok {
		result["ipmac"] = flattenObjectFspVlanInterfaceIpmac(i["ipmac"], d, pre_append)
	}

	pre_append = pre + ".0." + "ips_sniffer_mode"
	if _, ok := i["ips-sniffer-mode"]; ok {
		result["ips_sniffer_mode"] = flattenObjectFspVlanInterfaceIpsSnifferMode(i["ips-sniffer-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipunnumbered"
	if _, ok := i["ipunnumbered"]; ok {
		result["ipunnumbered"] = flattenObjectFspVlanInterfaceIpunnumbered(i["ipunnumbered"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6"
	if _, ok := i["ipv6"]; ok {
		result["ipv6"] = flattenObjectFspVlanInterfaceIpv6(i["ipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "l2forward"
	if _, ok := i["l2forward"]; ok {
		result["l2forward"] = flattenObjectFspVlanInterfaceL2Forward(i["l2forward"], d, pre_append)
	}

	pre_append = pre + ".0." + "l2tp_client"
	if _, ok := i["l2tp-client"]; ok {
		result["l2tp_client"] = flattenObjectFspVlanInterfaceL2TpClient(i["l2tp-client"], d, pre_append)
	}

	pre_append = pre + ".0." + "lacp_ha_secondary"
	if _, ok := i["lacp-ha-secondary"]; ok {
		result["lacp_ha_secondary"] = flattenObjectFspVlanInterfaceLacpHaSecondary(i["lacp-ha-secondary"], d, pre_append)
	}

	pre_append = pre + ".0." + "lacp_ha_slave"
	if _, ok := i["lacp-ha-slave"]; ok {
		result["lacp_ha_slave"] = flattenObjectFspVlanInterfaceLacpHaSlave(i["lacp-ha-slave"], d, pre_append)
	}

	pre_append = pre + ".0." + "lacp_mode"
	if _, ok := i["lacp-mode"]; ok {
		result["lacp_mode"] = flattenObjectFspVlanInterfaceLacpMode(i["lacp-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "lacp_speed"
	if _, ok := i["lacp-speed"]; ok {
		result["lacp_speed"] = flattenObjectFspVlanInterfaceLacpSpeed(i["lacp-speed"], d, pre_append)
	}

	pre_append = pre + ".0." + "large_receive_offload"
	if _, ok := i["large-receive-offload"]; ok {
		result["large_receive_offload"] = flattenObjectFspVlanInterfaceLargeReceiveOffload(i["large-receive-offload"], d, pre_append)
	}

	pre_append = pre + ".0." + "lcp_echo_interval"
	if _, ok := i["lcp-echo-interval"]; ok {
		result["lcp_echo_interval"] = flattenObjectFspVlanInterfaceLcpEchoInterval(i["lcp-echo-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "lcp_max_echo_fails"
	if _, ok := i["lcp-max-echo-fails"]; ok {
		result["lcp_max_echo_fails"] = flattenObjectFspVlanInterfaceLcpMaxEchoFails(i["lcp-max-echo-fails"], d, pre_append)
	}

	pre_append = pre + ".0." + "link_up_delay"
	if _, ok := i["link-up-delay"]; ok {
		result["link_up_delay"] = flattenObjectFspVlanInterfaceLinkUpDelay(i["link-up-delay"], d, pre_append)
	}

	pre_append = pre + ".0." + "listen_forticlient_connection"
	if _, ok := i["listen-forticlient-connection"]; ok {
		result["listen_forticlient_connection"] = flattenObjectFspVlanInterfaceListenForticlientConnection(i["listen-forticlient-connection"], d, pre_append)
	}

	pre_append = pre + ".0." + "lldp_network_policy"
	if _, ok := i["lldp-network-policy"]; ok {
		result["lldp_network_policy"] = flattenObjectFspVlanInterfaceLldpNetworkPolicy(i["lldp-network-policy"], d, pre_append)
	}

	pre_append = pre + ".0." + "lldp_reception"
	if _, ok := i["lldp-reception"]; ok {
		result["lldp_reception"] = flattenObjectFspVlanInterfaceLldpReception(i["lldp-reception"], d, pre_append)
	}

	pre_append = pre + ".0." + "lldp_transmission"
	if _, ok := i["lldp-transmission"]; ok {
		result["lldp_transmission"] = flattenObjectFspVlanInterfaceLldpTransmission(i["lldp-transmission"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectFspVlanInterfaceLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "macaddr"
	if _, ok := i["macaddr"]; ok {
		result["macaddr"] = flattenObjectFspVlanInterfaceMacaddr(i["macaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "managed_subnetwork_size"
	if _, ok := i["managed-subnetwork-size"]; ok {
		result["managed_subnetwork_size"] = flattenObjectFspVlanInterfaceManagedSubnetworkSize(i["managed-subnetwork-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "management_ip"
	if _, ok := i["management-ip"]; ok {
		result["management_ip"] = flattenObjectFspVlanInterfaceManagementIp(i["management-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_egress_burst_rate"
	if _, ok := i["max-egress-burst-rate"]; ok {
		result["max_egress_burst_rate"] = flattenObjectFspVlanInterfaceMaxEgressBurstRate(i["max-egress-burst-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_egress_rate"
	if _, ok := i["max-egress-rate"]; ok {
		result["max_egress_rate"] = flattenObjectFspVlanInterfaceMaxEgressRate(i["max-egress-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "measured_downstream_bandwidth"
	if _, ok := i["measured-downstream-bandwidth"]; ok {
		result["measured_downstream_bandwidth"] = flattenObjectFspVlanInterfaceMeasuredDownstreamBandwidth(i["measured-downstream-bandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "measured_upstream_bandwidth"
	if _, ok := i["measured-upstream-bandwidth"]; ok {
		result["measured_upstream_bandwidth"] = flattenObjectFspVlanInterfaceMeasuredUpstreamBandwidth(i["measured-upstream-bandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "mediatype"
	if _, ok := i["mediatype"]; ok {
		result["mediatype"] = flattenObjectFspVlanInterfaceMediatype(i["mediatype"], d, pre_append)
	}

	pre_append = pre + ".0." + "member"
	if _, ok := i["member"]; ok {
		result["member"] = flattenObjectFspVlanInterfaceMember(i["member"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_links"
	if _, ok := i["min-links"]; ok {
		result["min_links"] = flattenObjectFspVlanInterfaceMinLinks(i["min-links"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_links_down"
	if _, ok := i["min-links-down"]; ok {
		result["min_links_down"] = flattenObjectFspVlanInterfaceMinLinksDown(i["min-links-down"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenObjectFspVlanInterfaceMode(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "monitor_bandwidth"
	if _, ok := i["monitor-bandwidth"]; ok {
		result["monitor_bandwidth"] = flattenObjectFspVlanInterfaceMonitorBandwidth(i["monitor-bandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "mtu"
	if _, ok := i["mtu"]; ok {
		result["mtu"] = flattenObjectFspVlanInterfaceMtu(i["mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "mtu_override"
	if _, ok := i["mtu-override"]; ok {
		result["mtu_override"] = flattenObjectFspVlanInterfaceMtuOverride(i["mtu-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "mux_type"
	if _, ok := i["mux-type"]; ok {
		result["mux_type"] = flattenObjectFspVlanInterfaceMuxType(i["mux-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenObjectFspVlanInterfaceName(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "ndiscforward"
	if _, ok := i["ndiscforward"]; ok {
		result["ndiscforward"] = flattenObjectFspVlanInterfaceNdiscforward(i["ndiscforward"], d, pre_append)
	}

	pre_append = pre + ".0." + "netbios_forward"
	if _, ok := i["netbios-forward"]; ok {
		result["netbios_forward"] = flattenObjectFspVlanInterfaceNetbiosForward(i["netbios-forward"], d, pre_append)
	}

	pre_append = pre + ".0." + "netflow_sampler"
	if _, ok := i["netflow-sampler"]; ok {
		result["netflow_sampler"] = flattenObjectFspVlanInterfaceNetflowSampler(i["netflow-sampler"], d, pre_append)
	}

	pre_append = pre + ".0." + "np_qos_profile"
	if _, ok := i["np-qos-profile"]; ok {
		result["np_qos_profile"] = flattenObjectFspVlanInterfaceNpQosProfile(i["np-qos-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "npu_fastpath"
	if _, ok := i["npu-fastpath"]; ok {
		result["npu_fastpath"] = flattenObjectFspVlanInterfaceNpuFastpath(i["npu-fastpath"], d, pre_append)
	}

	pre_append = pre + ".0." + "nst"
	if _, ok := i["nst"]; ok {
		result["nst"] = flattenObjectFspVlanInterfaceNst(i["nst"], d, pre_append)
	}

	pre_append = pre + ".0." + "out_force_vlan_cos"
	if _, ok := i["out-force-vlan-cos"]; ok {
		result["out_force_vlan_cos"] = flattenObjectFspVlanInterfaceOutForceVlanCos(i["out-force-vlan-cos"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbandwidth"
	if _, ok := i["outbandwidth"]; ok {
		result["outbandwidth"] = flattenObjectFspVlanInterfaceOutbandwidth(i["outbandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "padt_retry_timeout"
	if _, ok := i["padt-retry-timeout"]; ok {
		result["padt_retry_timeout"] = flattenObjectFspVlanInterfacePadtRetryTimeout(i["padt-retry-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "password"
	if _, ok := i["password"]; ok {
		result["password"] = flattenObjectFspVlanInterfacePassword(i["password"], d, pre_append)
	}

	pre_append = pre + ".0." + "peer_interface"
	if _, ok := i["peer-interface"]; ok {
		result["peer_interface"] = flattenObjectFspVlanInterfacePeerInterface(i["peer-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "phy_mode"
	if _, ok := i["phy-mode"]; ok {
		result["phy_mode"] = flattenObjectFspVlanInterfacePhyMode(i["phy-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ping_serv_status"
	if _, ok := i["ping-serv-status"]; ok {
		result["ping_serv_status"] = flattenObjectFspVlanInterfacePingServStatus(i["ping-serv-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "poe"
	if _, ok := i["poe"]; ok {
		result["poe"] = flattenObjectFspVlanInterfacePoe(i["poe"], d, pre_append)
	}

	pre_append = pre + ".0." + "polling_interval"
	if _, ok := i["polling-interval"]; ok {
		result["polling_interval"] = flattenObjectFspVlanInterfacePollingInterval(i["polling-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "pppoe_unnumbered_negotiate"
	if _, ok := i["pppoe-unnumbered-negotiate"]; ok {
		result["pppoe_unnumbered_negotiate"] = flattenObjectFspVlanInterfacePppoeUnnumberedNegotiate(i["pppoe-unnumbered-negotiate"], d, pre_append)
	}

	pre_append = pre + ".0." + "pptp_auth_type"
	if _, ok := i["pptp-auth-type"]; ok {
		result["pptp_auth_type"] = flattenObjectFspVlanInterfacePptpAuthType(i["pptp-auth-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "pptp_client"
	if _, ok := i["pptp-client"]; ok {
		result["pptp_client"] = flattenObjectFspVlanInterfacePptpClient(i["pptp-client"], d, pre_append)
	}

	pre_append = pre + ".0." + "pptp_password"
	if _, ok := i["pptp-password"]; ok {
		result["pptp_password"] = flattenObjectFspVlanInterfacePptpPassword(i["pptp-password"], d, pre_append)
	}

	pre_append = pre + ".0." + "pptp_server_ip"
	if _, ok := i["pptp-server-ip"]; ok {
		result["pptp_server_ip"] = flattenObjectFspVlanInterfacePptpServerIp(i["pptp-server-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "pptp_timeout"
	if _, ok := i["pptp-timeout"]; ok {
		result["pptp_timeout"] = flattenObjectFspVlanInterfacePptpTimeout(i["pptp-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "pptp_user"
	if _, ok := i["pptp-user"]; ok {
		result["pptp_user"] = flattenObjectFspVlanInterfacePptpUser(i["pptp-user"], d, pre_append)
	}

	pre_append = pre + ".0." + "preserve_session_route"
	if _, ok := i["preserve-session-route"]; ok {
		result["preserve_session_route"] = flattenObjectFspVlanInterfacePreserveSessionRoute(i["preserve-session-route"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenObjectFspVlanInterfacePriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority_override"
	if _, ok := i["priority-override"]; ok {
		result["priority_override"] = flattenObjectFspVlanInterfacePriorityOverride(i["priority-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_captive_portal"
	if _, ok := i["proxy-captive-portal"]; ok {
		result["proxy_captive_portal"] = flattenObjectFspVlanInterfaceProxyCaptivePortal(i["proxy-captive-portal"], d, pre_append)
	}

	pre_append = pre + ".0." + "pvc_atm_qos"
	if _, ok := i["pvc-atm-qos"]; ok {
		result["pvc_atm_qos"] = flattenObjectFspVlanInterfacePvcAtmQos(i["pvc-atm-qos"], d, pre_append)
	}

	pre_append = pre + ".0." + "pvc_chan"
	if _, ok := i["pvc-chan"]; ok {
		result["pvc_chan"] = flattenObjectFspVlanInterfacePvcChan(i["pvc-chan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pvc_crc"
	if _, ok := i["pvc-crc"]; ok {
		result["pvc_crc"] = flattenObjectFspVlanInterfacePvcCrc(i["pvc-crc"], d, pre_append)
	}

	pre_append = pre + ".0." + "pvc_pcr"
	if _, ok := i["pvc-pcr"]; ok {
		result["pvc_pcr"] = flattenObjectFspVlanInterfacePvcPcr(i["pvc-pcr"], d, pre_append)
	}

	pre_append = pre + ".0." + "pvc_scr"
	if _, ok := i["pvc-scr"]; ok {
		result["pvc_scr"] = flattenObjectFspVlanInterfacePvcScr(i["pvc-scr"], d, pre_append)
	}

	pre_append = pre + ".0." + "pvc_vlan_id"
	if _, ok := i["pvc-vlan-id"]; ok {
		result["pvc_vlan_id"] = flattenObjectFspVlanInterfacePvcVlanId(i["pvc-vlan-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "pvc_vlan_rx_id"
	if _, ok := i["pvc-vlan-rx-id"]; ok {
		result["pvc_vlan_rx_id"] = flattenObjectFspVlanInterfacePvcVlanRxId(i["pvc-vlan-rx-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "pvc_vlan_rx_op"
	if _, ok := i["pvc-vlan-rx-op"]; ok {
		result["pvc_vlan_rx_op"] = flattenObjectFspVlanInterfacePvcVlanRxOp(i["pvc-vlan-rx-op"], d, pre_append)
	}

	pre_append = pre + ".0." + "pvc_vlan_tx_id"
	if _, ok := i["pvc-vlan-tx-id"]; ok {
		result["pvc_vlan_tx_id"] = flattenObjectFspVlanInterfacePvcVlanTxId(i["pvc-vlan-tx-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "pvc_vlan_tx_op"
	if _, ok := i["pvc-vlan-tx-op"]; ok {
		result["pvc_vlan_tx_op"] = flattenObjectFspVlanInterfacePvcVlanTxOp(i["pvc-vlan-tx-op"], d, pre_append)
	}

	pre_append = pre + ".0." + "reachable_time"
	if _, ok := i["reachable-time"]; ok {
		result["reachable_time"] = flattenObjectFspVlanInterfaceReachableTime(i["reachable-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_interface"
	if _, ok := i["redundant-interface"]; ok {
		result["redundant_interface"] = flattenObjectFspVlanInterfaceRedundantInterface(i["redundant-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "remote_ip"
	if _, ok := i["remote-ip"]; ok {
		result["remote_ip"] = flattenObjectFspVlanInterfaceRemoteIp(i["remote-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "replacemsg_override_group"
	if _, ok := i["replacemsg-override-group"]; ok {
		result["replacemsg_override_group"] = flattenObjectFspVlanInterfaceReplacemsgOverrideGroup(i["replacemsg-override-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "retransmission"
	if _, ok := i["retransmission"]; ok {
		result["retransmission"] = flattenObjectFspVlanInterfaceRetransmission(i["retransmission"], d, pre_append)
	}

	pre_append = pre + ".0." + "ring_rx"
	if _, ok := i["ring-rx"]; ok {
		result["ring_rx"] = flattenObjectFspVlanInterfaceRingRx(i["ring-rx"], d, pre_append)
	}

	pre_append = pre + ".0." + "ring_tx"
	if _, ok := i["ring-tx"]; ok {
		result["ring_tx"] = flattenObjectFspVlanInterfaceRingTx(i["ring-tx"], d, pre_append)
	}

	pre_append = pre + ".0." + "role"
	if _, ok := i["role"]; ok {
		result["role"] = flattenObjectFspVlanInterfaceRole(i["role"], d, pre_append)
	}

	pre_append = pre + ".0." + "sample_direction"
	if _, ok := i["sample-direction"]; ok {
		result["sample_direction"] = flattenObjectFspVlanInterfaceSampleDirection(i["sample-direction"], d, pre_append)
	}

	pre_append = pre + ".0." + "sample_rate"
	if _, ok := i["sample-rate"]; ok {
		result["sample_rate"] = flattenObjectFspVlanInterfaceSampleRate(i["sample-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_botnet_connections"
	if _, ok := i["scan-botnet-connections"]; ok {
		result["scan_botnet_connections"] = flattenObjectFspVlanInterfaceScanBotnetConnections(i["scan-botnet-connections"], d, pre_append)
	}

	pre_append = pre + ".0." + "secondary_IP"
	if _, ok := i["secondary-IP"]; ok {
		result["secondary_IP"] = flattenObjectFspVlanInterfaceSecondaryIp(i["secondary-IP"], d, pre_append)
	}

	pre_append = pre + ".0." + "secondaryip"
	if _, ok := i["secondaryip"]; ok {
		result["secondaryip"] = flattenObjectFspVlanInterfaceSecondaryip(i["secondaryip"], d, pre_append)
	}

	pre_append = pre + ".0." + "security_8021x_dynamic_vlan_id"
	if _, ok := i["security-8021x-dynamic-vlan-id"]; ok {
		result["security_8021x_dynamic_vlan_id"] = flattenObjectFspVlanInterfaceSecurity8021XDynamicVlanId(i["security-8021x-dynamic-vlan-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "security_8021x_master"
	if _, ok := i["security-8021x-master"]; ok {
		result["security_8021x_master"] = flattenObjectFspVlanInterfaceSecurity8021XMaster(i["security-8021x-master"], d, pre_append)
	}

	pre_append = pre + ".0." + "security_8021x_mode"
	if _, ok := i["security-8021x-mode"]; ok {
		result["security_8021x_mode"] = flattenObjectFspVlanInterfaceSecurity8021XMode(i["security-8021x-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "security_exempt_list"
	if _, ok := i["security-exempt-list"]; ok {
		result["security_exempt_list"] = flattenObjectFspVlanInterfaceSecurityExemptList(i["security-exempt-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "security_external_logout"
	if _, ok := i["security-external-logout"]; ok {
		result["security_external_logout"] = flattenObjectFspVlanInterfaceSecurityExternalLogout(i["security-external-logout"], d, pre_append)
	}

	pre_append = pre + ".0." + "security_external_web"
	if _, ok := i["security-external-web"]; ok {
		result["security_external_web"] = flattenObjectFspVlanInterfaceSecurityExternalWeb(i["security-external-web"], d, pre_append)
	}

	pre_append = pre + ".0." + "security_groups"
	if _, ok := i["security-groups"]; ok {
		result["security_groups"] = flattenObjectFspVlanInterfaceSecurityGroups(i["security-groups"], d, pre_append)
	}

	pre_append = pre + ".0." + "security_mac_auth_bypass"
	if _, ok := i["security-mac-auth-bypass"]; ok {
		result["security_mac_auth_bypass"] = flattenObjectFspVlanInterfaceSecurityMacAuthBypass(i["security-mac-auth-bypass"], d, pre_append)
	}

	pre_append = pre + ".0." + "security_mode"
	if _, ok := i["security-mode"]; ok {
		result["security_mode"] = flattenObjectFspVlanInterfaceSecurityMode(i["security-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "security_redirect_url"
	if _, ok := i["security-redirect-url"]; ok {
		result["security_redirect_url"] = flattenObjectFspVlanInterfaceSecurityRedirectUrl(i["security-redirect-url"], d, pre_append)
	}

	pre_append = pre + ".0." + "select_profile_30a_35b"
	if _, ok := i["select-profile-30a-35b"]; ok {
		result["select_profile_30a_35b"] = flattenObjectFspVlanInterfaceSelectProfile30A35B(i["select-profile-30a-35b"], d, pre_append)
	}

	pre_append = pre + ".0." + "service_name"
	if _, ok := i["service-name"]; ok {
		result["service_name"] = flattenObjectFspVlanInterfaceServiceName(i["service-name"], d, pre_append)
	}

	pre_append = pre + ".0." + "sflow_sampler"
	if _, ok := i["sflow-sampler"]; ok {
		result["sflow_sampler"] = flattenObjectFspVlanInterfaceSflowSampler(i["sflow-sampler"], d, pre_append)
	}

	pre_append = pre + ".0." + "sfp_dsl"
	if _, ok := i["sfp-dsl"]; ok {
		result["sfp_dsl"] = flattenObjectFspVlanInterfaceSfpDsl(i["sfp-dsl"], d, pre_append)
	}

	pre_append = pre + ".0." + "sfp_dsl_adsl_fallback"
	if _, ok := i["sfp-dsl-adsl-fallback"]; ok {
		result["sfp_dsl_adsl_fallback"] = flattenObjectFspVlanInterfaceSfpDslAdslFallback(i["sfp-dsl-adsl-fallback"], d, pre_append)
	}

	pre_append = pre + ".0." + "sfp_dsl_autodetect"
	if _, ok := i["sfp-dsl-autodetect"]; ok {
		result["sfp_dsl_autodetect"] = flattenObjectFspVlanInterfaceSfpDslAutodetect(i["sfp-dsl-autodetect"], d, pre_append)
	}

	pre_append = pre + ".0." + "sfp_dsl_mac"
	if _, ok := i["sfp-dsl-mac"]; ok {
		result["sfp_dsl_mac"] = flattenObjectFspVlanInterfaceSfpDslMac(i["sfp-dsl-mac"], d, pre_append)
	}

	pre_append = pre + ".0." + "speed"
	if _, ok := i["speed"]; ok {
		result["speed"] = flattenObjectFspVlanInterfaceSpeed(i["speed"], d, pre_append)
	}

	pre_append = pre + ".0." + "spillover_threshold"
	if _, ok := i["spillover-threshold"]; ok {
		result["spillover_threshold"] = flattenObjectFspVlanInterfaceSpilloverThreshold(i["spillover-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_check"
	if _, ok := i["src-check"]; ok {
		result["src_check"] = flattenObjectFspVlanInterfaceSrcCheck(i["src-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFspVlanInterfaceStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "stp"
	if _, ok := i["stp"]; ok {
		result["stp"] = flattenObjectFspVlanInterfaceStp(i["stp"], d, pre_append)
	}

	pre_append = pre + ".0." + "stp_ha_secondary"
	if _, ok := i["stp-ha-secondary"]; ok {
		result["stp_ha_secondary"] = flattenObjectFspVlanInterfaceStpHaSecondary(i["stp-ha-secondary"], d, pre_append)
	}

	pre_append = pre + ".0." + "stp_ha_slave"
	if _, ok := i["stp-ha-slave"]; ok {
		result["stp_ha_slave"] = flattenObjectFspVlanInterfaceStpHaSlave(i["stp-ha-slave"], d, pre_append)
	}

	pre_append = pre + ".0." + "stpforward"
	if _, ok := i["stpforward"]; ok {
		result["stpforward"] = flattenObjectFspVlanInterfaceStpforward(i["stpforward"], d, pre_append)
	}

	pre_append = pre + ".0." + "stpforward_mode"
	if _, ok := i["stpforward-mode"]; ok {
		result["stpforward_mode"] = flattenObjectFspVlanInterfaceStpforwardMode(i["stpforward-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "strip_priority_vlan_tag"
	if _, ok := i["strip-priority-vlan-tag"]; ok {
		result["strip_priority_vlan_tag"] = flattenObjectFspVlanInterfaceStripPriorityVlanTag(i["strip-priority-vlan-tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "subst"
	if _, ok := i["subst"]; ok {
		result["subst"] = flattenObjectFspVlanInterfaceSubst(i["subst"], d, pre_append)
	}

	pre_append = pre + ".0." + "substitute_dst_mac"
	if _, ok := i["substitute-dst-mac"]; ok {
		result["substitute_dst_mac"] = flattenObjectFspVlanInterfaceSubstituteDstMac(i["substitute-dst-mac"], d, pre_append)
	}

	pre_append = pre + ".0." + "sw_algorithm"
	if _, ok := i["sw-algorithm"]; ok {
		result["sw_algorithm"] = flattenObjectFspVlanInterfaceSwAlgorithm(i["sw-algorithm"], d, pre_append)
	}

	pre_append = pre + ".0." + "swc_first_create"
	if _, ok := i["swc-first-create"]; ok {
		result["swc_first_create"] = flattenObjectFspVlanInterfaceSwcFirstCreate(i["swc-first-create"], d, pre_append)
	}

	pre_append = pre + ".0." + "swc_vlan"
	if _, ok := i["swc-vlan"]; ok {
		result["swc_vlan"] = flattenObjectFspVlanInterfaceSwcVlan(i["swc-vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch"
	if _, ok := i["switch"]; ok {
		result["switch"] = flattenObjectFspVlanInterfaceSwitch(i["switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_access_vlan"
	if _, ok := i["switch-controller-access-vlan"]; ok {
		result["switch_controller_access_vlan"] = flattenObjectFspVlanInterfaceSwitchControllerAccessVlan(i["switch-controller-access-vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_arp_inspection"
	if _, ok := i["switch-controller-arp-inspection"]; ok {
		result["switch_controller_arp_inspection"] = flattenObjectFspVlanInterfaceSwitchControllerArpInspection(i["switch-controller-arp-inspection"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_auth"
	if _, ok := i["switch-controller-auth"]; ok {
		result["switch_controller_auth"] = flattenObjectFspVlanInterfaceSwitchControllerAuth(i["switch-controller-auth"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_dhcp_snooping"
	if _, ok := i["switch-controller-dhcp-snooping"]; ok {
		result["switch_controller_dhcp_snooping"] = flattenObjectFspVlanInterfaceSwitchControllerDhcpSnooping(i["switch-controller-dhcp-snooping"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_dhcp_snooping_option82"
	if _, ok := i["switch-controller-dhcp-snooping-option82"]; ok {
		result["switch_controller_dhcp_snooping_option82"] = flattenObjectFspVlanInterfaceSwitchControllerDhcpSnoopingOption82(i["switch-controller-dhcp-snooping-option82"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_dhcp_snooping_verify_mac"
	if _, ok := i["switch-controller-dhcp-snooping-verify-mac"]; ok {
		result["switch_controller_dhcp_snooping_verify_mac"] = flattenObjectFspVlanInterfaceSwitchControllerDhcpSnoopingVerifyMac(i["switch-controller-dhcp-snooping-verify-mac"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_dynamic"
	if _, ok := i["switch-controller-dynamic"]; ok {
		result["switch_controller_dynamic"] = flattenObjectFspVlanInterfaceSwitchControllerDynamic(i["switch-controller-dynamic"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_feature"
	if _, ok := i["switch-controller-feature"]; ok {
		result["switch_controller_feature"] = flattenObjectFspVlanInterfaceSwitchControllerFeature(i["switch-controller-feature"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_igmp_snooping"
	if _, ok := i["switch-controller-igmp-snooping"]; ok {
		result["switch_controller_igmp_snooping"] = flattenObjectFspVlanInterfaceSwitchControllerIgmpSnooping(i["switch-controller-igmp-snooping"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_igmp_snooping_fast_leave"
	if _, ok := i["switch-controller-igmp-snooping-fast-leave"]; ok {
		result["switch_controller_igmp_snooping_fast_leave"] = flattenObjectFspVlanInterfaceSwitchControllerIgmpSnoopingFastLeave(i["switch-controller-igmp-snooping-fast-leave"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_igmp_snooping_proxy"
	if _, ok := i["switch-controller-igmp-snooping-proxy"]; ok {
		result["switch_controller_igmp_snooping_proxy"] = flattenObjectFspVlanInterfaceSwitchControllerIgmpSnoopingProxy(i["switch-controller-igmp-snooping-proxy"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_iot_scanning"
	if _, ok := i["switch-controller-iot-scanning"]; ok {
		result["switch_controller_iot_scanning"] = flattenObjectFspVlanInterfaceSwitchControllerIotScanning(i["switch-controller-iot-scanning"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_learning_limit"
	if _, ok := i["switch-controller-learning-limit"]; ok {
		result["switch_controller_learning_limit"] = flattenObjectFspVlanInterfaceSwitchControllerLearningLimit(i["switch-controller-learning-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_mgmt_vlan"
	if _, ok := i["switch-controller-mgmt-vlan"]; ok {
		result["switch_controller_mgmt_vlan"] = flattenObjectFspVlanInterfaceSwitchControllerMgmtVlan(i["switch-controller-mgmt-vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_nac"
	if _, ok := i["switch-controller-nac"]; ok {
		result["switch_controller_nac"] = flattenObjectFspVlanInterfaceSwitchControllerNac(i["switch-controller-nac"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_netflow_collect"
	if _, ok := i["switch-controller-netflow-collect"]; ok {
		result["switch_controller_netflow_collect"] = flattenObjectFspVlanInterfaceSwitchControllerNetflowCollect(i["switch-controller-netflow-collect"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_radius_server"
	if _, ok := i["switch-controller-radius-server"]; ok {
		result["switch_controller_radius_server"] = flattenObjectFspVlanInterfaceSwitchControllerRadiusServer(i["switch-controller-radius-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_rspan_mode"
	if _, ok := i["switch-controller-rspan-mode"]; ok {
		result["switch_controller_rspan_mode"] = flattenObjectFspVlanInterfaceSwitchControllerRspanMode(i["switch-controller-rspan-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_source_ip"
	if _, ok := i["switch-controller-source-ip"]; ok {
		result["switch_controller_source_ip"] = flattenObjectFspVlanInterfaceSwitchControllerSourceIp(i["switch-controller-source-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_controller_traffic_policy"
	if _, ok := i["switch-controller-traffic-policy"]; ok {
		result["switch_controller_traffic_policy"] = flattenObjectFspVlanInterfaceSwitchControllerTrafficPolicy(i["switch-controller-traffic-policy"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_id"
	if _, ok := i["system-id"]; ok {
		result["system_id"] = flattenObjectFspVlanInterfaceSystemId(i["system-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_id_type"
	if _, ok := i["system-id-type"]; ok {
		result["system_id_type"] = flattenObjectFspVlanInterfaceSystemIdType(i["system-id-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "tc_mode"
	if _, ok := i["tc-mode"]; ok {
		result["tc_mode"] = flattenObjectFspVlanInterfaceTcMode(i["tc-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_mss"
	if _, ok := i["tcp-mss"]; ok {
		result["tcp_mss"] = flattenObjectFspVlanInterfaceTcpMss(i["tcp-mss"], d, pre_append)
	}

	pre_append = pre + ".0." + "trunk"
	if _, ok := i["trunk"]; ok {
		result["trunk"] = flattenObjectFspVlanInterfaceTrunk(i["trunk"], d, pre_append)
	}

	pre_append = pre + ".0." + "trust_ip_1"
	if _, ok := i["trust-ip-1"]; ok {
		result["trust_ip_1"] = flattenObjectFspVlanInterfaceTrustIp1(i["trust-ip-1"], d, pre_append)
	}

	pre_append = pre + ".0." + "trust_ip_2"
	if _, ok := i["trust-ip-2"]; ok {
		result["trust_ip_2"] = flattenObjectFspVlanInterfaceTrustIp2(i["trust-ip-2"], d, pre_append)
	}

	pre_append = pre + ".0." + "trust_ip_3"
	if _, ok := i["trust-ip-3"]; ok {
		result["trust_ip_3"] = flattenObjectFspVlanInterfaceTrustIp3(i["trust-ip-3"], d, pre_append)
	}

	pre_append = pre + ".0." + "trust_ip6_1"
	if _, ok := i["trust-ip6-1"]; ok {
		result["trust_ip6_1"] = flattenObjectFspVlanInterfaceTrustIp61(i["trust-ip6-1"], d, pre_append)
	}

	pre_append = pre + ".0." + "trust_ip6_2"
	if _, ok := i["trust-ip6-2"]; ok {
		result["trust_ip6_2"] = flattenObjectFspVlanInterfaceTrustIp62(i["trust-ip6-2"], d, pre_append)
	}

	pre_append = pre + ".0." + "trust_ip6_3"
	if _, ok := i["trust-ip6-3"]; ok {
		result["trust_ip6_3"] = flattenObjectFspVlanInterfaceTrustIp63(i["trust-ip6-3"], d, pre_append)
	}

	pre_append = pre + ".0." + "type"
	if _, ok := i["type"]; ok {
		result["type"] = flattenObjectFspVlanInterfaceType(i["type"], d, pre_append)
	}

	pre_append = pre + ".0." + "username"
	if _, ok := i["username"]; ok {
		result["username"] = flattenObjectFspVlanInterfaceUsername(i["username"], d, pre_append)
	}

	pre_append = pre + ".0." + "vci"
	if _, ok := i["vci"]; ok {
		result["vci"] = flattenObjectFspVlanInterfaceVci(i["vci"], d, pre_append)
	}

	pre_append = pre + ".0." + "vectoring"
	if _, ok := i["vectoring"]; ok {
		result["vectoring"] = flattenObjectFspVlanInterfaceVectoring(i["vectoring"], d, pre_append)
	}

	pre_append = pre + ".0." + "vindex"
	if _, ok := i["vindex"]; ok {
		result["vindex"] = flattenObjectFspVlanInterfaceVindex(i["vindex"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan_protocol"
	if _, ok := i["vlan-protocol"]; ok {
		result["vlan_protocol"] = flattenObjectFspVlanInterfaceVlanProtocol(i["vlan-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlanforward"
	if _, ok := i["vlanforward"]; ok {
		result["vlanforward"] = flattenObjectFspVlanInterfaceVlanforward(i["vlanforward"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlanid"
	if _, ok := i["vlanid"]; ok {
		result["vlanid"] = flattenObjectFspVlanInterfaceVlanid(i["vlanid"], d, pre_append)
	}

	pre_append = pre + ".0." + "vpi"
	if _, ok := i["vpi"]; ok {
		result["vpi"] = flattenObjectFspVlanInterfaceVpi(i["vpi"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrf"
	if _, ok := i["vrf"]; ok {
		result["vrf"] = flattenObjectFspVlanInterfaceVrf(i["vrf"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp"
	if _, ok := i["vrrp"]; ok {
		result["vrrp"] = flattenObjectFspVlanInterfaceVrrp(i["vrrp"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp_virtual_mac"
	if _, ok := i["vrrp-virtual-mac"]; ok {
		result["vrrp_virtual_mac"] = flattenObjectFspVlanInterfaceVrrpVirtualMac(i["vrrp-virtual-mac"], d, pre_append)
	}

	pre_append = pre + ".0." + "wccp"
	if _, ok := i["wccp"]; ok {
		result["wccp"] = flattenObjectFspVlanInterfaceWccp(i["wccp"], d, pre_append)
	}

	pre_append = pre + ".0." + "weight"
	if _, ok := i["weight"]; ok {
		result["weight"] = flattenObjectFspVlanInterfaceWeight(i["weight"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_5g_threshold"
	if _, ok := i["wifi-5g-threshold"]; ok {
		result["wifi_5g_threshold"] = flattenObjectFspVlanInterfaceWifi5GThreshold(i["wifi-5g-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_acl"
	if _, ok := i["wifi-acl"]; ok {
		result["wifi_acl"] = flattenObjectFspVlanInterfaceWifiAcl(i["wifi-acl"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ap_band"
	if _, ok := i["wifi-ap-band"]; ok {
		result["wifi_ap_band"] = flattenObjectFspVlanInterfaceWifiApBand(i["wifi-ap-band"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_auth"
	if _, ok := i["wifi-auth"]; ok {
		result["wifi_auth"] = flattenObjectFspVlanInterfaceWifiAuth(i["wifi-auth"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_auto_connect"
	if _, ok := i["wifi-auto-connect"]; ok {
		result["wifi_auto_connect"] = flattenObjectFspVlanInterfaceWifiAutoConnect(i["wifi-auto-connect"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_auto_save"
	if _, ok := i["wifi-auto-save"]; ok {
		result["wifi_auto_save"] = flattenObjectFspVlanInterfaceWifiAutoSave(i["wifi-auto-save"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_broadcast_ssid"
	if _, ok := i["wifi-broadcast-ssid"]; ok {
		result["wifi_broadcast_ssid"] = flattenObjectFspVlanInterfaceWifiBroadcastSsid(i["wifi-broadcast-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_encrypt"
	if _, ok := i["wifi-encrypt"]; ok {
		result["wifi_encrypt"] = flattenObjectFspVlanInterfaceWifiEncrypt(i["wifi-encrypt"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_fragment_threshold"
	if _, ok := i["wifi-fragment-threshold"]; ok {
		result["wifi_fragment_threshold"] = flattenObjectFspVlanInterfaceWifiFragmentThreshold(i["wifi-fragment-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_key"
	if _, ok := i["wifi-key"]; ok {
		result["wifi_key"] = flattenObjectFspVlanInterfaceWifiKey(i["wifi-key"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_keyindex"
	if _, ok := i["wifi-keyindex"]; ok {
		result["wifi_keyindex"] = flattenObjectFspVlanInterfaceWifiKeyindex(i["wifi-keyindex"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_mac_filter"
	if _, ok := i["wifi-mac-filter"]; ok {
		result["wifi_mac_filter"] = flattenObjectFspVlanInterfaceWifiMacFilter(i["wifi-mac-filter"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_passphrase"
	if _, ok := i["wifi-passphrase"]; ok {
		result["wifi_passphrase"] = flattenObjectFspVlanInterfaceWifiPassphrase(i["wifi-passphrase"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_radius_server"
	if _, ok := i["wifi-radius-server"]; ok {
		result["wifi_radius_server"] = flattenObjectFspVlanInterfaceWifiRadiusServer(i["wifi-radius-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_rts_threshold"
	if _, ok := i["wifi-rts-threshold"]; ok {
		result["wifi_rts_threshold"] = flattenObjectFspVlanInterfaceWifiRtsThreshold(i["wifi-rts-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_security"
	if _, ok := i["wifi-security"]; ok {
		result["wifi_security"] = flattenObjectFspVlanInterfaceWifiSecurity(i["wifi-security"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_ssid"
	if _, ok := i["wifi-ssid"]; ok {
		result["wifi_ssid"] = flattenObjectFspVlanInterfaceWifiSsid(i["wifi-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "wifi_usergroup"
	if _, ok := i["wifi-usergroup"]; ok {
		result["wifi_usergroup"] = flattenObjectFspVlanInterfaceWifiUsergroup(i["wifi-usergroup"], d, pre_append)
	}

	pre_append = pre + ".0." + "wins_ip"
	if _, ok := i["wins-ip"]; ok {
		result["wins_ip"] = flattenObjectFspVlanInterfaceWinsIp(i["wins-ip"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFspVlanInterfaceVlanOpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAggregate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAggregateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceApDiscover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceArpforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAtmProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAuthPortalAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceAutoAuthExtensionDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBandwidthMeasureTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBfdDetectMult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBroadcastForticlientDiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceBroadcastForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceCliConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsKeyname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceDdnsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDdnsZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDedicatedTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDefaultgw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDetectedPeerMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDetectprotocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceDetectserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDeviceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDeviceIdentification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDeviceIdentificationActiveScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDeviceNetscan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDeviceUserIdentification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDevindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpClasslessRouteAddition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpClientIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayAgentOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceDhcpRelayLinkSelection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayRequestAllServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDhcpRenewTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDiscRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDnsQuery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDnsServerOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDnsServerProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceDropFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceDropOverlappedFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEapCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceEapIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEapMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEapPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceEapSupplicant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEapUserCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceEgressCos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEgressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEndpointCompliance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEstimatedDownstreamBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceEstimatedUpstreamBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceExplicitFtpProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceExplicitWebProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFailActionOnExtender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFailAlertInterfaces(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFailAlertMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFailDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFailDetectOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceFdp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortiheartbeat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortilinkBackupLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortilinkNeighborDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortilinkSplitInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFortilinkStacking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceForwardDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceForwardErrorCorrection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceFpAnomaly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceFpDisable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceGatewayAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceGenericReceiveOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceGiGk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceGwaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceGwdetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIcmpAcceptRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIcmpRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIcmpSendRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIdentAccept(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIfMdix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIfMedia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIkeSamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceInForceVlanCos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceInbandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIngressCos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIngressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceInterconnectProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceInternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpManagedByFortiipam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpmac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpsSnifferMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpunnumbered(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "autoconf"
	if _, ok := i["autoconf"]; ok {
		result["autoconf"] = flattenObjectFspVlanInterfaceIpv6Autoconf(i["autoconf"], d, pre_append)
	}

	pre_append = pre + ".0." + "cli_conn6_status"
	if _, ok := i["cli-conn6-status"]; ok {
		result["cli_conn6_status"] = flattenObjectFspVlanInterfaceIpv6CliConn6Status(i["cli-conn6-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := i["dhcp6-client-options"]; ok {
		result["dhcp6_client_options"] = flattenObjectFspVlanInterfaceIpv6Dhcp6ClientOptions(i["dhcp6-client-options"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := i["dhcp6-information-request"]; ok {
		result["dhcp6_information_request"] = flattenObjectFspVlanInterfaceIpv6Dhcp6InformationRequest(i["dhcp6-information-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := i["dhcp6-prefix-delegation"]; ok {
		result["dhcp6_prefix_delegation"] = flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation(i["dhcp6-prefix-delegation"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := i["dhcp6-prefix-hint"]; ok {
		result["dhcp6_prefix_hint"] = flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHint(i["dhcp6-prefix-hint"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := i["dhcp6-prefix-hint-plt"]; ok {
		result["dhcp6_prefix_hint_plt"] = flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt(i["dhcp6-prefix-hint-plt"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := i["dhcp6-prefix-hint-vlt"]; ok {
		result["dhcp6_prefix_hint_vlt"] = flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt(i["dhcp6-prefix-hint-vlt"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := i["dhcp6-relay-ip"]; ok {
		result["dhcp6_relay_ip"] = flattenObjectFspVlanInterfaceIpv6Dhcp6RelayIp(i["dhcp6-relay-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := i["dhcp6-relay-service"]; ok {
		result["dhcp6_relay_service"] = flattenObjectFspVlanInterfaceIpv6Dhcp6RelayService(i["dhcp6-relay-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := i["dhcp6-relay-type"]; ok {
		result["dhcp6_relay_type"] = flattenObjectFspVlanInterfaceIpv6Dhcp6RelayType(i["dhcp6-relay-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp6_send_redirect"
	if _, ok := i["icmp6-send-redirect"]; ok {
		result["icmp6_send_redirect"] = flattenObjectFspVlanInterfaceIpv6Icmp6SendRedirect(i["icmp6-send-redirect"], d, pre_append)
	}

	pre_append = pre + ".0." + "interface_identifier"
	if _, ok := i["interface-identifier"]; ok {
		result["interface_identifier"] = flattenObjectFspVlanInterfaceIpv6InterfaceIdentifier(i["interface-identifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_address"
	if _, ok := i["ip6-address"]; ok {
		result["ip6_address"] = flattenObjectFspVlanInterfaceIpv6Ip6Address(i["ip6-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := i["ip6-allowaccess"]; ok {
		result["ip6_allowaccess"] = flattenObjectFspVlanInterfaceIpv6Ip6Allowaccess(i["ip6-allowaccess"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := i["ip6-default-life"]; ok {
		result["ip6_default_life"] = flattenObjectFspVlanInterfaceIpv6Ip6DefaultLife(i["ip6-default-life"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_iaid"
	if _, ok := i["ip6-delegated-prefix-iaid"]; ok {
		result["ip6_delegated_prefix_iaid"] = flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid(i["ip6-delegated-prefix-iaid"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := i["ip6-delegated-prefix-list"]; ok {
		result["ip6_delegated_prefix_list"] = flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(i["ip6-delegated-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := i["ip6-dns-server-override"]; ok {
		result["ip6_dns_server_override"] = flattenObjectFspVlanInterfaceIpv6Ip6DnsServerOverride(i["ip6-dns-server-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := i["ip6-extra-addr"]; ok {
		result["ip6_extra_addr"] = flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddr(i["ip6-extra-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := i["ip6-hop-limit"]; ok {
		result["ip6_hop_limit"] = flattenObjectFspVlanInterfaceIpv6Ip6HopLimit(i["ip6-hop-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := i["ip6-link-mtu"]; ok {
		result["ip6_link_mtu"] = flattenObjectFspVlanInterfaceIpv6Ip6LinkMtu(i["ip6-link-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := i["ip6-manage-flag"]; ok {
		result["ip6_manage_flag"] = flattenObjectFspVlanInterfaceIpv6Ip6ManageFlag(i["ip6-manage-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := i["ip6-max-interval"]; ok {
		result["ip6_max_interval"] = flattenObjectFspVlanInterfaceIpv6Ip6MaxInterval(i["ip6-max-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := i["ip6-min-interval"]; ok {
		result["ip6_min_interval"] = flattenObjectFspVlanInterfaceIpv6Ip6MinInterval(i["ip6-min-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := i["ip6-mode"]; ok {
		result["ip6_mode"] = flattenObjectFspVlanInterfaceIpv6Ip6Mode(i["ip6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := i["ip6-other-flag"]; ok {
		result["ip6_other_flag"] = flattenObjectFspVlanInterfaceIpv6Ip6OtherFlag(i["ip6-other-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := i["ip6-prefix-list"]; ok {
		result["ip6_prefix_list"] = flattenObjectFspVlanInterfaceIpv6Ip6PrefixList(i["ip6-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_mode"
	if _, ok := i["ip6-prefix-mode"]; ok {
		result["ip6_prefix_mode"] = flattenObjectFspVlanInterfaceIpv6Ip6PrefixMode(i["ip6-prefix-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := i["ip6-reachable-time"]; ok {
		result["ip6_reachable_time"] = flattenObjectFspVlanInterfaceIpv6Ip6ReachableTime(i["ip6-reachable-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := i["ip6-retrans-time"]; ok {
		result["ip6_retrans_time"] = flattenObjectFspVlanInterfaceIpv6Ip6RetransTime(i["ip6-retrans-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := i["ip6-send-adv"]; ok {
		result["ip6_send_adv"] = flattenObjectFspVlanInterfaceIpv6Ip6SendAdv(i["ip6-send-adv"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := i["ip6-subnet"]; ok {
		result["ip6_subnet"] = flattenObjectFspVlanInterfaceIpv6Ip6Subnet(i["ip6-subnet"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := i["ip6-upstream-interface"]; ok {
		result["ip6_upstream_interface"] = flattenObjectFspVlanInterfaceIpv6Ip6UpstreamInterface(i["ip6-upstream-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cert"
	if _, ok := i["nd-cert"]; ok {
		result["nd_cert"] = flattenObjectFspVlanInterfaceIpv6NdCert(i["nd-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := i["nd-cga-modifier"]; ok {
		result["nd_cga_modifier"] = flattenObjectFspVlanInterfaceIpv6NdCgaModifier(i["nd-cga-modifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_mode"
	if _, ok := i["nd-mode"]; ok {
		result["nd_mode"] = flattenObjectFspVlanInterfaceIpv6NdMode(i["nd-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := i["nd-security-level"]; ok {
		result["nd_security_level"] = flattenObjectFspVlanInterfaceIpv6NdSecurityLevel(i["nd-security-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := i["nd-timestamp-delta"]; ok {
		result["nd_timestamp_delta"] = flattenObjectFspVlanInterfaceIpv6NdTimestampDelta(i["nd-timestamp-delta"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := i["nd-timestamp-fuzz"]; ok {
		result["nd_timestamp_fuzz"] = flattenObjectFspVlanInterfaceIpv6NdTimestampFuzz(i["nd-timestamp-fuzz"], d, pre_append)
	}

	pre_append = pre + ".0." + "ra_send_mtu"
	if _, ok := i["ra-send-mtu"]; ok {
		result["ra_send_mtu"] = flattenObjectFspVlanInterfaceIpv6RaSendMtu(i["ra-send-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "unique_autoconf_addr"
	if _, ok := i["unique-autoconf-addr"]; ok {
		result["unique_autoconf_addr"] = flattenObjectFspVlanInterfaceIpv6UniqueAutoconfAddr(i["unique-autoconf-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := i["vrip6_link_local"]; ok {
		result["vrip6_link_local"] = flattenObjectFspVlanInterfaceIpv6Vrip6LinkLocal(i["vrip6_link_local"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := i["vrrp-virtual-mac6"]; ok {
		result["vrrp_virtual_mac6"] = flattenObjectFspVlanInterfaceIpv6VrrpVirtualMac6(i["vrrp-virtual-mac6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp6"
	if _, ok := i["vrrp6"]; ok {
		result["vrrp6"] = flattenObjectFspVlanInterfaceIpv6Vrrp6(i["vrrp6"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFspVlanInterfaceIpv6Autoconf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6CliConn6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6ClientOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6InformationRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Icmp6SendRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6InterfaceIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6DefaultLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := i["delegated-prefix-iaid"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid(i["delegated-prefix-iaid"], d, pre_append)
			tmp["delegated_prefix_iaid"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-DelegatedPrefixIaid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := i["prefix-id"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId(i["prefix-id"], d, pre_append)
			tmp["prefix_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-PrefixId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := i["rdnss-service"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService(i["rdnss-service"], d, pre_append)
			tmp["rdnss_service"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-RdnssService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := i["upstream-interface"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(i["upstream-interface"], d, pre_append)
			tmp["upstream_interface"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-UpstreamInterface")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DnsServerOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6ExtraAddr-Prefix")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6HopLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6LinkMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6ManageFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6MaxInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6MinInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6OtherFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := i["dnssl"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl(i["dnssl"], d, pre_append)
			tmp["dnssl"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-Dnssl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := i["preferred-life-time"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime(i["preferred-life-time"], d, pre_append)
			tmp["preferred_life_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-PreferredLifeTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := i["valid-life-time"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime(i["valid-life-time"], d, pre_append)
			tmp["valid_life_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-ValidLifeTime")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6ReachableTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6RetransTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6SendAdv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6UpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdCgaModifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdTimestampDelta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdTimestampFuzz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6RaSendMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6UniqueAutoconfAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrip6LinkLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6VrrpVirtualMac6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6AcceptMode(i["accept-mode"], d, pre_append)
			tmp["accept_mode"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-AcceptMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6AdvInterval(i["adv-interval"], d, pre_append)
			tmp["adv_interval"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-AdvInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Preempt(i["preempt"], d, pre_append)
			tmp["preempt"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Preempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Priority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6StartTime(i["start-time"], d, pre_append)
			tmp["start_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-StartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Status(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := i["vrdst6"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrdst6(i["vrdst6"], d, pre_append)
			tmp["vrdst6"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrdst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrgrp(i["vrgrp"], d, pre_append)
			tmp["vrgrp"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrgrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrid(i["vrid"], d, pre_append)
			tmp["vrid"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := i["vrip6"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrip6(i["vrip6"], d, pre_append)
			tmp["vrip6"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrip6")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6AcceptMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6AdvInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Preempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6StartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrdst6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceL2Forward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceL2TpClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLacpHaSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLacpHaSlave(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLacpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLacpSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLargeReceiveOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLcpEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLcpMaxEchoFails(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLinkUpDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceListenForticlientConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLldpNetworkPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLldpReception(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLldpTransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMacaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceManagedSubnetworkSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceManagementIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMaxEgressBurstRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMaxEgressRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMeasuredDownstreamBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMeasuredUpstreamBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMediatype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMinLinks(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMinLinksDown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMonitorBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMtuOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceMuxType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNdiscforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNetbiosForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNetflowSampler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNpQosProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNpuFastpath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceNst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceOutForceVlanCos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceOutbandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePadtRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfacePeerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePhyMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePingServStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePoe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePollingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePppoeUnnumberedNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePptpAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePptpClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePptpPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfacePptpServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePptpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePptpUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePreserveSessionRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePriorityOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceProxyCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcAtmQos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcCrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcPcr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcScr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcVlanRxId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcVlanRxOp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcVlanTxId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfacePvcVlanTxOp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceReachableTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRedundantInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRemoteIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRetransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRingRx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRingTx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSampleDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSampleRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryip(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanInterfaceSecondaryipAllowaccess(i["allowaccess"], d, pre_append)
			tmp["allowaccess"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Allowaccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := i["detectprotocol"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipDetectprotocol(i["detectprotocol"], d, pre_append)
			tmp["detectprotocol"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Detectprotocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := i["detectserver"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipDetectserver(i["detectserver"], d, pre_append)
			tmp["detectserver"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Detectserver")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := i["gwdetect"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipGwdetect(i["gwdetect"], d, pre_append)
			tmp["gwdetect"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Gwdetect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipHaPriority(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := i["ping-serv-status"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipPingServStatus(i["ping-serv-status"], d, pre_append)
			tmp["ping_serv_status"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-PingServStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenObjectFspVlanInterfaceSecondaryipSeq(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Secondaryip-Seq")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceSecondaryipAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceSecondaryipDetectprotocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceSecondaryipDetectserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipGwdetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipPingServStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurity8021XDynamicVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurity8021XMaster(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurity8021XMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityExemptList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityExternalLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityExternalWeb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecurityRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSelectProfile30A35B(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSflowSampler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSfpDsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSfpDslAdslFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSfpDslAutodetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSfpDslMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSrcCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStpHaSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStpHaSlave(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStpforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStpforwardMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceStripPriorityVlanTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSubst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSubstituteDstMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwcFirstCreate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwcVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerAccessVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerArpInspection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerDhcpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerDhcpSnoopingOption82(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerDhcpSnoopingVerifyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerFeature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerIgmpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerIgmpSnoopingFastLeave(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerIgmpSnoopingProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerIotScanning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerLearningLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerMgmtVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerNac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerNetflowCollect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerRspanMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSwitchControllerTrafficPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSystemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSystemIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTcMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTcpMss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrunk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp61(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp62(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceTrustIp63(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVci(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVectoring(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVlanProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVlanforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanInterfaceVrrpAcceptMode(i["accept-mode"], d, pre_append)
			tmp["accept_mode"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-AcceptMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpAdvInterval(i["adv-interval"], d, pre_append)
			tmp["adv_interval"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-AdvInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := i["ignore-default-route"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpIgnoreDefaultRoute(i["ignore-default-route"], d, pre_append)
			tmp["ignore_default_route"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-IgnoreDefaultRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpPreempt(i["preempt"], d, pre_append)
			tmp["preempt"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Preempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpStartTime(i["start-time"], d, pre_append)
			tmp["start_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-StartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVersion(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Version")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst"
		if _, ok := i["vrdst"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVrdst(i["vrdst"], d, pre_append)
			tmp["vrdst"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Vrdst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := i["vrdst-priority"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVrdstPriority(i["vrdst-priority"], d, pre_append)
			tmp["vrdst_priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-VrdstPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVrgrp(i["vrgrp"], d, pre_append)
			tmp["vrgrp"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Vrgrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVrid(i["vrid"], d, pre_append)
			tmp["vrid"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Vrid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip"
		if _, ok := i["vrip"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpVrip(i["vrip"], d, pre_append)
			tmp["vrip"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterface-Vrrp-Vrip")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceVrrpAcceptMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpAdvInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpIgnoreDefaultRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpPreempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpStartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrdst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceVrrpVrdstPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVirtualMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWccp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifi5GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiAcl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiApBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiAutoConnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiAutoSave(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiBroadcastSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiEncrypt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiFragmentThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceWifiKeyindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiMacFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiPassphrase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceWifiRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiRtsThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWifiUsergroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceWinsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_dhcp_status", flattenObjectFspVlanDhcpStatus(o["_dhcp-status"], d, "_dhcp_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dhcp-status"], "ObjectFspVlan-DhcpStatus"); ok {
			if err = d.Set("_dhcp_status", vv); err != nil {
				return fmt.Errorf("Error reading _dhcp_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dhcp_status: %v", err)
		}
	}

	if err = d.Set("color", flattenObjectFspVlanColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFspVlan-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dhcp_server", flattenObjectFspVlanDhcpServer(o["dhcp-server"], d, "dhcp_server")); err != nil {
			if vv, ok := fortiAPIPatch(o["dhcp-server"], "ObjectFspVlan-DhcpServer"); ok {
				if err = d.Set("dhcp_server", vv); err != nil {
					return fmt.Errorf("Error reading dhcp_server: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dhcp_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dhcp_server"); ok {
			if err = d.Set("dhcp_server", flattenObjectFspVlanDhcpServer(o["dhcp-server"], d, "dhcp_server")); err != nil {
				if vv, ok := fortiAPIPatch(o["dhcp-server"], "ObjectFspVlan-DhcpServer"); ok {
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
		if err = d.Set("dynamic_mapping", flattenObjectFspVlanDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFspVlan-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectFspVlanDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFspVlan-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("interface", flattenObjectFspVlanInterface(o["interface"], d, "interface")); err != nil {
			if vv, ok := fortiAPIPatch(o["interface"], "ObjectFspVlan-Interface"); ok {
				if err = d.Set("interface", vv); err != nil {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenObjectFspVlanInterface(o["interface"], d, "interface")); err != nil {
				if vv, ok := fortiAPIPatch(o["interface"], "ObjectFspVlan-Interface"); ok {
					if err = d.Set("interface", vv); err != nil {
						return fmt.Errorf("Error reading interface: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectFspVlanName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFspVlan-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vdom", flattenObjectFspVlanVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "ObjectFspVlan-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenObjectFspVlanVlanid(o["vlanid"], d, "vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlanid"], "ObjectFspVlan-Vlanid"); ok {
			if err = d.Set("vlanid", vv); err != nil {
				return fmt.Errorf("Error reading vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanDhcpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_configuration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-configuration"], _ = expandObjectFspVlanDhcpServerAutoConfiguration(d, i["auto_configuration"], pre_append)
	}
	pre_append = pre + ".0." + "auto_managed_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-managed-status"], _ = expandObjectFspVlanDhcpServerAutoManagedStatus(d, i["auto_managed_status"], pre_append)
	}
	pre_append = pre + ".0." + "conflicted_ip_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conflicted-ip-timeout"], _ = expandObjectFspVlanDhcpServerConflictedIpTimeout(d, i["conflicted_ip_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_auth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-auth"], _ = expandObjectFspVlanDhcpServerDdnsAuth(d, i["ddns_auth"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_key"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-key"], _ = expandObjectFspVlanDhcpServerDdnsKey(d, i["ddns_key"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_keyname"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-keyname"], _ = expandObjectFspVlanDhcpServerDdnsKeyname(d, i["ddns_keyname"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_server_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-server-ip"], _ = expandObjectFspVlanDhcpServerDdnsServerIp(d, i["ddns_server_ip"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_ttl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-ttl"], _ = expandObjectFspVlanDhcpServerDdnsTtl(d, i["ddns_ttl"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_update"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-update"], _ = expandObjectFspVlanDhcpServerDdnsUpdate(d, i["ddns_update"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_update_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-update-override"], _ = expandObjectFspVlanDhcpServerDdnsUpdateOverride(d, i["ddns_update_override"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_zone"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-zone"], _ = expandObjectFspVlanDhcpServerDdnsZone(d, i["ddns_zone"], pre_append)
	}
	pre_append = pre + ".0." + "default_gateway"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-gateway"], _ = expandObjectFspVlanDhcpServerDefaultGateway(d, i["default_gateway"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_settings_from_fortiipam"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-settings-from-fortiipam"], _ = expandObjectFspVlanDhcpServerDhcpSettingsFromFortiipam(d, i["dhcp_settings_from_fortiipam"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server1"], _ = expandObjectFspVlanDhcpServerDnsServer1(d, i["dns_server1"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server2"], _ = expandObjectFspVlanDhcpServerDnsServer2(d, i["dns_server2"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server3"], _ = expandObjectFspVlanDhcpServerDnsServer3(d, i["dns_server3"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server4"], _ = expandObjectFspVlanDhcpServerDnsServer4(d, i["dns_server4"], pre_append)
	}
	pre_append = pre + ".0." + "dns_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-service"], _ = expandObjectFspVlanDhcpServerDnsService(d, i["dns_service"], pre_append)
	}
	pre_append = pre + ".0." + "domain"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["domain"], _ = expandObjectFspVlanDhcpServerDomain(d, i["domain"], pre_append)
	}
	pre_append = pre + ".0." + "enable"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["enable"], _ = expandObjectFspVlanDhcpServerEnable(d, i["enable"], pre_append)
	}
	pre_append = pre + ".0." + "exclude_range"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["exclude-range"], _ = expandObjectFspVlanDhcpServerExcludeRange(d, i["exclude_range"], pre_append)
	} else {
		result["exclude-range"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "filename"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["filename"], _ = expandObjectFspVlanDhcpServerFilename(d, i["filename"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_on_net_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["forticlient-on-net-status"], _ = expandObjectFspVlanDhcpServerForticlientOnNetStatus(d, i["forticlient_on_net_status"], pre_append)
	}
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandObjectFspVlanDhcpServerId(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "ip_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip-mode"], _ = expandObjectFspVlanDhcpServerIpMode(d, i["ip_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip_range"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip-range"], _ = expandObjectFspVlanDhcpServerIpRange(d, i["ip_range"], pre_append)
	} else {
		result["ip-range"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ipsec_lease_hold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipsec-lease-hold"], _ = expandObjectFspVlanDhcpServerIpsecLeaseHold(d, i["ipsec_lease_hold"], pre_append)
	}
	pre_append = pre + ".0." + "lease_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lease-time"], _ = expandObjectFspVlanDhcpServerLeaseTime(d, i["lease_time"], pre_append)
	}
	pre_append = pre + ".0." + "mac_acl_default_action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-acl-default-action"], _ = expandObjectFspVlanDhcpServerMacAclDefaultAction(d, i["mac_acl_default_action"], pre_append)
	}
	pre_append = pre + ".0." + "netmask"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["netmask"], _ = expandObjectFspVlanDhcpServerNetmask(d, i["netmask"], pre_append)
	}
	pre_append = pre + ".0." + "next_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["next-server"], _ = expandObjectFspVlanDhcpServerNextServer(d, i["next_server"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_server1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-server1"], _ = expandObjectFspVlanDhcpServerNtpServer1(d, i["ntp_server1"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-server2"], _ = expandObjectFspVlanDhcpServerNtpServer2(d, i["ntp_server2"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_server3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-server3"], _ = expandObjectFspVlanDhcpServerNtpServer3(d, i["ntp_server3"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-service"], _ = expandObjectFspVlanDhcpServerNtpService(d, i["ntp_service"], pre_append)
	}
	pre_append = pre + ".0." + "option1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option1"], _ = expandObjectFspVlanDhcpServerOption1(d, i["option1"], pre_append)
	} else {
		result["option1"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "option2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option2"], _ = expandObjectFspVlanDhcpServerOption2(d, i["option2"], pre_append)
	} else {
		result["option2"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "option3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option3"], _ = expandObjectFspVlanDhcpServerOption3(d, i["option3"], pre_append)
	} else {
		result["option3"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "option4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option4"], _ = expandObjectFspVlanDhcpServerOption4(d, i["option4"], pre_append)
	}
	pre_append = pre + ".0." + "option5"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option5"], _ = expandObjectFspVlanDhcpServerOption5(d, i["option5"], pre_append)
	}
	pre_append = pre + ".0." + "option6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option6"], _ = expandObjectFspVlanDhcpServerOption6(d, i["option6"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectFspVlanDhcpServerOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "reserved_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["reserved-address"], _ = expandObjectFspVlanDhcpServerReservedAddress(d, i["reserved_address"], pre_append)
	} else {
		result["reserved-address"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "server_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-type"], _ = expandObjectFspVlanDhcpServerServerType(d, i["server_type"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectFspVlanDhcpServerStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tftp_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tftp-server"], _ = expandObjectFspVlanDhcpServerTftpServer(d, i["tftp_server"], pre_append)
	} else {
		result["tftp-server"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "timezone"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["timezone"], _ = expandObjectFspVlanDhcpServerTimezone(d, i["timezone"], pre_append)
	}
	pre_append = pre + ".0." + "timezone_option"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["timezone-option"], _ = expandObjectFspVlanDhcpServerTimezoneOption(d, i["timezone_option"], pre_append)
	}
	pre_append = pre + ".0." + "vci_match"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vci-match"], _ = expandObjectFspVlanDhcpServerVciMatch(d, i["vci_match"], pre_append)
	}
	pre_append = pre + ".0." + "vci_string"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vci-string"], _ = expandObjectFspVlanDhcpServerVciString(d, i["vci_string"], pre_append)
	} else {
		result["vci-string"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "wifi_ac_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac-service"], _ = expandObjectFspVlanDhcpServerWifiAcService(d, i["wifi_ac_service"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ac1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac1"], _ = expandObjectFspVlanDhcpServerWifiAc1(d, i["wifi_ac1"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ac2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac2"], _ = expandObjectFspVlanDhcpServerWifiAc2(d, i["wifi_ac2"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ac3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac3"], _ = expandObjectFspVlanDhcpServerWifiAc3(d, i["wifi_ac3"], pre_append)
	}
	pre_append = pre + ".0." + "wins_server1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wins-server1"], _ = expandObjectFspVlanDhcpServerWinsServer1(d, i["wins_server1"], pre_append)
	}
	pre_append = pre + ".0." + "wins_server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wins-server2"], _ = expandObjectFspVlanDhcpServerWinsServer2(d, i["wins_server2"], pre_append)
	}

	return result, nil
}

func expandObjectFspVlanDhcpServerAutoConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerAutoManagedStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerConflictedIpTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsKeyname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsUpdateOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDefaultGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDhcpSettingsFromFortiipam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDnsServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDnsService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerEnable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectFspVlanDhcpServerExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDhcpServerExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFspVlanDhcpServerExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDhcpServerExcludeRangeVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDhcpServerExcludeRangeVciString(d, i["vci_string"], pre_append)
		} else {
			tmp["vci-string"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerFilename(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerForticlientOnNetStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectFspVlanDhcpServerIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDhcpServerIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFspVlanDhcpServerIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDhcpServerIpRangeVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDhcpServerIpRangeVciString(d, i["vci_string"], pre_append)
		} else {
			tmp["vci-string"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDhcpServerIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRangeVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRangeVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerIpsecLeaseHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerMacAclDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNextServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNtpServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNtpServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNtpServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNtpService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOption1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerOption2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerOption3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerOption4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOption5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOption6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["code"], _ = expandObjectFspVlanDhcpServerOptionsCode(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDhcpServerOptionsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDhcpServerOptionsIp(d, i["ip"], pre_append)
		} else {
			tmp["ip"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFspVlanDhcpServerOptionsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectFspVlanDhcpServerOptionsValue(d, i["value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDhcpServerOptionsVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDhcpServerOptionsVciString(d, i["vci_string"], pre_append)
		} else {
			tmp["vci-string"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDhcpServerOptionsCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerOptionsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerReservedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectFspVlanDhcpServerReservedAddressAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id"], _ = expandObjectFspVlanDhcpServerReservedAddressCircuitId(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id-type"], _ = expandObjectFspVlanDhcpServerReservedAddressCircuitIdType(d, i["circuit_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectFspVlanDhcpServerReservedAddressDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDhcpServerReservedAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDhcpServerReservedAddressIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectFspVlanDhcpServerReservedAddressMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandObjectFspVlanDhcpServerReservedAddressRemoteId(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id-type"], _ = expandObjectFspVlanDhcpServerReservedAddressRemoteIdType(d, i["remote_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFspVlanDhcpServerReservedAddressType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDhcpServerReservedAddressAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressCircuitId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressCircuitIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressRemoteId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressRemoteIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerTftpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerTimezone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerTimezoneOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerWifiAcService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerWifiAc1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerWifiAc2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerWifiAc3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerWinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerWinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_dhcp_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_dhcp-status"], _ = expandObjectFspVlanDynamicMappingDhcpStatus(d, i["_dhcp_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_scope"], _ = expandObjectFspVlanDynamicMappingScope(d, i["_scope"], pre_append)
		} else {
			tmp["_scope"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-server"], _ = expandObjectFspVlanDynamicMappingDhcpServer(d, i["dhcp_server"], pre_append)
		} else {
			tmp["dhcp-server"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandObjectFspVlanDynamicMappingInterface(d, i["interface"], pre_append)
		} else {
			tmp["interface"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFspVlanDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectFspVlanDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_configuration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-configuration"], _ = expandObjectFspVlanDynamicMappingDhcpServerAutoConfiguration(d, i["auto_configuration"], pre_append)
	}
	pre_append = pre + ".0." + "auto_managed_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-managed-status"], _ = expandObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus(d, i["auto_managed_status"], pre_append)
	}
	pre_append = pre + ".0." + "conflicted_ip_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conflicted-ip-timeout"], _ = expandObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout(d, i["conflicted_ip_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_auth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-auth"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsAuth(d, i["ddns_auth"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_key"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-key"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsKey(d, i["ddns_key"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_keyname"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-keyname"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsKeyname(d, i["ddns_keyname"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_server_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-server-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsServerIp(d, i["ddns_server_ip"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_ttl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-ttl"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsTtl(d, i["ddns_ttl"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_update"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-update"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdate(d, i["ddns_update"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_update_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-update-override"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride(d, i["ddns_update_override"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_zone"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-zone"], _ = expandObjectFspVlanDynamicMappingDhcpServerDdnsZone(d, i["ddns_zone"], pre_append)
	}
	pre_append = pre + ".0." + "default_gateway"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-gateway"], _ = expandObjectFspVlanDynamicMappingDhcpServerDefaultGateway(d, i["default_gateway"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_settings_from_fortiipam"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-settings-from-fortiipam"], _ = expandObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam(d, i["dhcp_settings_from_fortiipam"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server1"], _ = expandObjectFspVlanDynamicMappingDhcpServerDnsServer1(d, i["dns_server1"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server2"], _ = expandObjectFspVlanDynamicMappingDhcpServerDnsServer2(d, i["dns_server2"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server3"], _ = expandObjectFspVlanDynamicMappingDhcpServerDnsServer3(d, i["dns_server3"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server4"], _ = expandObjectFspVlanDynamicMappingDhcpServerDnsServer4(d, i["dns_server4"], pre_append)
	}
	pre_append = pre + ".0." + "dns_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-service"], _ = expandObjectFspVlanDynamicMappingDhcpServerDnsService(d, i["dns_service"], pre_append)
	}
	pre_append = pre + ".0." + "domain"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["domain"], _ = expandObjectFspVlanDynamicMappingDhcpServerDomain(d, i["domain"], pre_append)
	}
	pre_append = pre + ".0." + "enable"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["enable"], _ = expandObjectFspVlanDynamicMappingDhcpServerEnable(d, i["enable"], pre_append)
	}
	pre_append = pre + ".0." + "exclude_range"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["exclude-range"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRange(d, i["exclude_range"], pre_append)
	} else {
		result["exclude-range"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "filename"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["filename"], _ = expandObjectFspVlanDynamicMappingDhcpServerFilename(d, i["filename"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_on_net_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["forticlient-on-net-status"], _ = expandObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus(d, i["forticlient_on_net_status"], pre_append)
	}
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerId(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "ip_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip-mode"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpMode(d, i["ip_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip_range"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip-range"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRange(d, i["ip_range"], pre_append)
	} else {
		result["ip-range"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ipsec_lease_hold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipsec-lease-hold"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold(d, i["ipsec_lease_hold"], pre_append)
	}
	pre_append = pre + ".0." + "lease_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lease-time"], _ = expandObjectFspVlanDynamicMappingDhcpServerLeaseTime(d, i["lease_time"], pre_append)
	}
	pre_append = pre + ".0." + "mac_acl_default_action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-acl-default-action"], _ = expandObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction(d, i["mac_acl_default_action"], pre_append)
	}
	pre_append = pre + ".0." + "netmask"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["netmask"], _ = expandObjectFspVlanDynamicMappingDhcpServerNetmask(d, i["netmask"], pre_append)
	}
	pre_append = pre + ".0." + "next_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["next-server"], _ = expandObjectFspVlanDynamicMappingDhcpServerNextServer(d, i["next_server"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_server1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-server1"], _ = expandObjectFspVlanDynamicMappingDhcpServerNtpServer1(d, i["ntp_server1"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-server2"], _ = expandObjectFspVlanDynamicMappingDhcpServerNtpServer2(d, i["ntp_server2"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_server3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-server3"], _ = expandObjectFspVlanDynamicMappingDhcpServerNtpServer3(d, i["ntp_server3"], pre_append)
	}
	pre_append = pre + ".0." + "ntp_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ntp-service"], _ = expandObjectFspVlanDynamicMappingDhcpServerNtpService(d, i["ntp_service"], pre_append)
	}
	pre_append = pre + ".0." + "option1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option1"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption1(d, i["option1"], pre_append)
	} else {
		result["option1"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "option2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option2"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption2(d, i["option2"], pre_append)
	} else {
		result["option2"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "option3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option3"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption3(d, i["option3"], pre_append)
	} else {
		result["option3"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "option4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option4"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption4(d, i["option4"], pre_append)
	}
	pre_append = pre + ".0." + "option5"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option5"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption5(d, i["option5"], pre_append)
	}
	pre_append = pre + ".0." + "option6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["option6"], _ = expandObjectFspVlanDynamicMappingDhcpServerOption6(d, i["option6"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "reserved_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["reserved-address"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddress(d, i["reserved_address"], pre_append)
	} else {
		result["reserved-address"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "server_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-type"], _ = expandObjectFspVlanDynamicMappingDhcpServerServerType(d, i["server_type"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectFspVlanDynamicMappingDhcpServerStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tftp_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tftp-server"], _ = expandObjectFspVlanDynamicMappingDhcpServerTftpServer(d, i["tftp_server"], pre_append)
	} else {
		result["tftp-server"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "timezone"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["timezone"], _ = expandObjectFspVlanDynamicMappingDhcpServerTimezone(d, i["timezone"], pre_append)
	}
	pre_append = pre + ".0." + "timezone_option"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["timezone-option"], _ = expandObjectFspVlanDynamicMappingDhcpServerTimezoneOption(d, i["timezone_option"], pre_append)
	}
	pre_append = pre + ".0." + "vci_match"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerVciMatch(d, i["vci_match"], pre_append)
	}
	pre_append = pre + ".0." + "vci_string"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerVciString(d, i["vci_string"], pre_append)
	} else {
		result["vci-string"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "wifi_ac_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac-service"], _ = expandObjectFspVlanDynamicMappingDhcpServerWifiAcService(d, i["wifi_ac_service"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ac1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac1"], _ = expandObjectFspVlanDynamicMappingDhcpServerWifiAc1(d, i["wifi_ac1"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ac2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac2"], _ = expandObjectFspVlanDynamicMappingDhcpServerWifiAc2(d, i["wifi_ac2"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ac3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ac3"], _ = expandObjectFspVlanDynamicMappingDhcpServerWifiAc3(d, i["wifi_ac3"], pre_append)
	}
	pre_append = pre + ".0." + "wins_server1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wins-server1"], _ = expandObjectFspVlanDynamicMappingDhcpServerWinsServer1(d, i["wins_server1"], pre_append)
	}
	pre_append = pre + ".0." + "wins_server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wins-server2"], _ = expandObjectFspVlanDynamicMappingDhcpServerWinsServer2(d, i["wins_server2"], pre_append)
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerAutoConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsKeyname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDefaultGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerEnable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString(d, i["vci_string"], pre_append)
		} else {
			tmp["vci-string"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerFilename(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciString(d, i["vci_string"], pre_append)
		} else {
			tmp["vci-string"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNextServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["code"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsCode(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsIp(d, i["ip"], pre_append)
		} else {
			tmp["ip"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsValue(d, i["value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsVciString(d, i["vci_string"], pre_append)
		} else {
			tmp["vci-string"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id-type"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType(d, i["circuit_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id-type"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType(d, i["remote_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerTftpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerTimezone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerTimezoneOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAcService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAc1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAc2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAc3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dhcp_relay_agent_option"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-agent-option"], _ = expandObjectFspVlanDynamicMappingInterfaceDhcpRelayAgentOption(d, i["dhcp_relay_agent_option"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_interface_select_method"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-interface-select-method"], _ = expandObjectFspVlanDynamicMappingInterfaceDhcpRelayInterfaceSelectMethod(d, i["dhcp_relay_interface_select_method"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-ip"], _ = expandObjectFspVlanDynamicMappingInterfaceDhcpRelayIp(d, i["dhcp_relay_ip"], pre_append)
	} else {
		result["dhcp-relay-ip"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "dhcp_relay_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-service"], _ = expandObjectFspVlanDynamicMappingInterfaceDhcpRelayService(d, i["dhcp_relay_service"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-type"], _ = expandObjectFspVlanDynamicMappingInterfaceDhcpRelayType(d, i["dhcp_relay_type"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandObjectFspVlanDynamicMappingInterfaceIp(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6(d, i["ipv6"], pre_append)
	} else {
		result["ipv6"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "secondary_IP"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["secondary-IP"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryIp(d, i["secondary_IP"], pre_append)
	}
	pre_append = pre + ".0." + "secondaryip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["secondaryip"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryip(d, i["secondaryip"], pre_append)
	} else {
		result["secondaryip"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "vlanid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlanid"], _ = expandObjectFspVlanDynamicMappingInterfaceVlanid(d, i["vlanid"], pre_append)
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceDhcpRelayAgentOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceDhcpRelayInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceDhcpRelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceDhcpRelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceDhcpRelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "autoconf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["autoconf"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Autoconf(d, i["autoconf"], pre_append)
	}
	pre_append = pre + ".0." + "cli_conn6_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cli-conn6-status"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6CliConn6Status(d, i["cli_conn6_status"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-client-options"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6ClientOptions(d, i["dhcp6_client_options"], pre_append)
	} else {
		result["dhcp6-client-options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-information-request"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6InformationRequest(d, i["dhcp6_information_request"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-delegation"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixDelegation(d, i["dhcp6_prefix_delegation"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHint(d, i["dhcp6_prefix_hint"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint-plt"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintPlt(d, i["dhcp6_prefix_hint_plt"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint-vlt"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintVlt(d, i["dhcp6_prefix_hint_vlt"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-ip"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayIp(d, i["dhcp6_relay_ip"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-service"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayService(d, i["dhcp6_relay_service"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-type"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayType(d, i["dhcp6_relay_type"], pre_append)
	}
	pre_append = pre + ".0." + "icmp6_send_redirect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp6-send-redirect"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Icmp6SendRedirect(d, i["icmp6_send_redirect"], pre_append)
	}
	pre_append = pre + ".0." + "interface_identifier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interface-identifier"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6InterfaceIdentifier(d, i["interface_identifier"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-address"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Address(d, i["ip6_address"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-allowaccess"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Allowaccess(d, i["ip6_allowaccess"], pre_append)
	} else {
		result["ip6-allowaccess"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-default-life"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DefaultLife(d, i["ip6_default_life"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_delegated_prefix_iaid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-delegated-prefix-iaid"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixIaid(d, i["ip6_delegated_prefix_iaid"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-delegated-prefix-list"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixList(d, i["ip6_delegated_prefix_list"], pre_append)
	} else {
		result["ip6-delegated-prefix-list"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-dns-server-override"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DnsServerOverride(d, i["ip6_dns_server_override"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-extra-addr"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddr(d, i["ip6_extra_addr"], pre_append)
	} else {
		result["ip6-extra-addr"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-hop-limit"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6HopLimit(d, i["ip6_hop_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-link-mtu"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6LinkMtu(d, i["ip6_link_mtu"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-manage-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ManageFlag(d, i["ip6_manage_flag"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-max-interval"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6MaxInterval(d, i["ip6_max_interval"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-min-interval"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6MinInterval(d, i["ip6_min_interval"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-mode"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Mode(d, i["ip6_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-other-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6OtherFlag(d, i["ip6_other_flag"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-prefix-list"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixList(d, i["ip6_prefix_list"], pre_append)
	} else {
		result["ip6-prefix-list"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_prefix_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-prefix-mode"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixMode(d, i["ip6_prefix_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-reachable-time"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ReachableTime(d, i["ip6_reachable_time"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-retrans-time"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6RetransTime(d, i["ip6_retrans_time"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-send-adv"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6SendAdv(d, i["ip6_send_adv"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-subnet"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Subnet(d, i["ip6_subnet"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-upstream-interface"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6UpstreamInterface(d, i["ip6_upstream_interface"], pre_append)
	}
	pre_append = pre + ".0." + "nd_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-cert"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdCert(d, i["nd_cert"], pre_append)
	}
	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-cga-modifier"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdCgaModifier(d, i["nd_cga_modifier"], pre_append)
	}
	pre_append = pre + ".0." + "nd_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-mode"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdMode(d, i["nd_mode"], pre_append)
	}
	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-security-level"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdSecurityLevel(d, i["nd_security_level"], pre_append)
	}
	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-timestamp-delta"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampDelta(d, i["nd_timestamp_delta"], pre_append)
	}
	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-timestamp-fuzz"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampFuzz(d, i["nd_timestamp_fuzz"], pre_append)
	}
	pre_append = pre + ".0." + "ra_send_mtu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ra-send-mtu"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6RaSendMtu(d, i["ra_send_mtu"], pre_append)
	}
	pre_append = pre + ".0." + "unique_autoconf_addr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unique-autoconf-addr"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6UniqueAutoconfAddr(d, i["unique_autoconf_addr"], pre_append)
	}
	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrip6_link_local"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrip6LinkLocal(d, i["vrip6_link_local"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrrp-virtual-mac6"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6VrrpVirtualMac6(d, i["vrrp_virtual_mac6"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrrp6"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6(d, i["vrrp6"], pre_append)
	} else {
		result["vrrp6"] = make([]string, 0)
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Autoconf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6CliConn6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6ClientOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6InformationRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixDelegation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintPlt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6PrefixHintVlt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Dhcp6RelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Icmp6SendRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6InterfaceIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Allowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DefaultLife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixIaid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["autonomous-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["delegated-prefix-iaid"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid(d, i["delegated_prefix_iaid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-id"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListPrefixId(d, i["prefix_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnss(d, i["rdnss"], pre_append)
		} else {
			tmp["rdnss"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss-service"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnssService(d, i["rdnss_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListSubnet(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["upstream-interface"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(d, i["upstream_interface"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListPrefixId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListRdnssService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6DnsServerOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddrPrefix(d, i["prefix"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ExtraAddrPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6HopLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6LinkMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ManageFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6MaxInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6MinInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6OtherFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["autonomous-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListAutonomousFlag(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dnssl"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListDnssl(d, i["dnssl"], pre_append)
		} else {
			tmp["dnssl"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListOnlinkFlag(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preferred-life-time"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPreferredLifeTime(d, i["preferred_life_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListRdnss(d, i["rdnss"], pre_append)
		} else {
			tmp["rdnss"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["valid-life-time"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListValidLifeTime(d, i["valid_life_time"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListAutonomousFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListDnssl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListOnlinkFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPreferredLifeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListRdnss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixListValidLifeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6PrefixMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6ReachableTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6RetransTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6SendAdv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Ip6UpstreamInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdCgaModifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdSecurityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampDelta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6NdTimestampFuzz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6RaSendMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6UniqueAutoconfAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrip6LinkLocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6VrrpVirtualMac6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["accept-mode"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AcceptMode(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-interval"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AdvInterval(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preempt"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Preempt(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Priority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-time"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6StartTime(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Status(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst6"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrdst6(d, i["vrdst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrgrp"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrgrp(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrid"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrid(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrip6"], _ = expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrip6(d, i["vrip6"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AcceptMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6AdvInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Preempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6StartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrdst6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceIpv6Vrrp6Vrip6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowaccess"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowaccess"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipAllowaccess(d, i["allowaccess"], pre_append)
		} else {
			tmp["allowaccess"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detectprotocol"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipDetectprotocol(d, i["detectprotocol"], pre_append)
		} else {
			tmp["detectprotocol"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detectserver"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipDetectserver(d, i["detectserver"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gwdetect"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipGwdetect(d, i["gwdetect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipHaPriority(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ping-serv-status"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipPingServStatus(d, i["ping_serv_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandObjectFspVlanDynamicMappingInterfaceSecondaryipSeq(d, i["seq"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipDetectprotocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipDetectserver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipGwdetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipPingServStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceSecondaryipSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingInterfaceVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vlan_op_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan-op-mode"], _ = expandObjectFspVlanInterfaceVlanOpMode(d, i["vlan_op_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ac_name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ac-name"], _ = expandObjectFspVlanInterfaceAcName(d, i["ac_name"], pre_append)
	}
	pre_append = pre + ".0." + "aggregate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["aggregate"], _ = expandObjectFspVlanInterfaceAggregate(d, i["aggregate"], pre_append)
	}
	pre_append = pre + ".0." + "aggregate_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["aggregate-type"], _ = expandObjectFspVlanInterfaceAggregateType(d, i["aggregate_type"], pre_append)
	}
	pre_append = pre + ".0." + "algorithm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["algorithm"], _ = expandObjectFspVlanInterfaceAlgorithm(d, i["algorithm"], pre_append)
	}
	pre_append = pre + ".0." + "alias"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["alias"], _ = expandObjectFspVlanInterfaceAlias(d, i["alias"], pre_append)
	}
	pre_append = pre + ".0." + "allowaccess"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["allowaccess"], _ = expandObjectFspVlanInterfaceAllowaccess(d, i["allowaccess"], pre_append)
	} else {
		result["allowaccess"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ap_discover"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ap-discover"], _ = expandObjectFspVlanInterfaceApDiscover(d, i["ap_discover"], pre_append)
	}
	pre_append = pre + ".0." + "arpforward"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["arpforward"], _ = expandObjectFspVlanInterfaceArpforward(d, i["arpforward"], pre_append)
	}
	pre_append = pre + ".0." + "atm_protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["atm-protocol"], _ = expandObjectFspVlanInterfaceAtmProtocol(d, i["atm_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "auth_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-cert"], _ = expandObjectFspVlanInterfaceAuthCert(d, i["auth_cert"], pre_append)
	}
	pre_append = pre + ".0." + "auth_portal_addr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-portal-addr"], _ = expandObjectFspVlanInterfaceAuthPortalAddr(d, i["auth_portal_addr"], pre_append)
	}
	pre_append = pre + ".0." + "auth_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-type"], _ = expandObjectFspVlanInterfaceAuthType(d, i["auth_type"], pre_append)
	}
	pre_append = pre + ".0." + "auto_auth_extension_device"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-auth-extension-device"], _ = expandObjectFspVlanInterfaceAutoAuthExtensionDevice(d, i["auto_auth_extension_device"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_measure_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bandwidth-measure-time"], _ = expandObjectFspVlanInterfaceBandwidthMeasureTime(d, i["bandwidth_measure_time"], pre_append)
	}
	pre_append = pre + ".0." + "bfd"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bfd"], _ = expandObjectFspVlanInterfaceBfd(d, i["bfd"], pre_append)
	}
	pre_append = pre + ".0." + "bfd_desired_min_tx"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bfd-desired-min-tx"], _ = expandObjectFspVlanInterfaceBfdDesiredMinTx(d, i["bfd_desired_min_tx"], pre_append)
	}
	pre_append = pre + ".0." + "bfd_detect_mult"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bfd-detect-mult"], _ = expandObjectFspVlanInterfaceBfdDetectMult(d, i["bfd_detect_mult"], pre_append)
	}
	pre_append = pre + ".0." + "bfd_required_min_rx"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bfd-required-min-rx"], _ = expandObjectFspVlanInterfaceBfdRequiredMinRx(d, i["bfd_required_min_rx"], pre_append)
	}
	pre_append = pre + ".0." + "broadcast_forticlient_discovery"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["broadcast-forticlient-discovery"], _ = expandObjectFspVlanInterfaceBroadcastForticlientDiscovery(d, i["broadcast_forticlient_discovery"], pre_append)
	}
	pre_append = pre + ".0." + "broadcast_forward"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["broadcast-forward"], _ = expandObjectFspVlanInterfaceBroadcastForward(d, i["broadcast_forward"], pre_append)
	}
	pre_append = pre + ".0." + "captive_portal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["captive-portal"], _ = expandObjectFspVlanInterfaceCaptivePortal(d, i["captive_portal"], pre_append)
	}
	pre_append = pre + ".0." + "cli_conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cli-conn-status"], _ = expandObjectFspVlanInterfaceCliConnStatus(d, i["cli_conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "color"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["color"], _ = expandObjectFspVlanInterfaceColor(d, i["color"], pre_append)
	}
	pre_append = pre + ".0." + "ddns"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns"], _ = expandObjectFspVlanInterfaceDdns(d, i["ddns"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_auth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-auth"], _ = expandObjectFspVlanInterfaceDdnsAuth(d, i["ddns_auth"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_domain"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-domain"], _ = expandObjectFspVlanInterfaceDdnsDomain(d, i["ddns_domain"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_key"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-key"], _ = expandObjectFspVlanInterfaceDdnsKey(d, i["ddns_key"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_keyname"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-keyname"], _ = expandObjectFspVlanInterfaceDdnsKeyname(d, i["ddns_keyname"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-password"], _ = expandObjectFspVlanInterfaceDdnsPassword(d, i["ddns_password"], pre_append)
	} else {
		result["ddns-password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ddns_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-server"], _ = expandObjectFspVlanInterfaceDdnsServer(d, i["ddns_server"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_server_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-server-ip"], _ = expandObjectFspVlanInterfaceDdnsServerIp(d, i["ddns_server_ip"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_sn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-sn"], _ = expandObjectFspVlanInterfaceDdnsSn(d, i["ddns_sn"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_ttl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-ttl"], _ = expandObjectFspVlanInterfaceDdnsTtl(d, i["ddns_ttl"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-username"], _ = expandObjectFspVlanInterfaceDdnsUsername(d, i["ddns_username"], pre_append)
	}
	pre_append = pre + ".0." + "ddns_zone"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ddns-zone"], _ = expandObjectFspVlanInterfaceDdnsZone(d, i["ddns_zone"], pre_append)
	}
	pre_append = pre + ".0." + "dedicated_to"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dedicated-to"], _ = expandObjectFspVlanInterfaceDedicatedTo(d, i["dedicated_to"], pre_append)
	}
	pre_append = pre + ".0." + "defaultgw"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["defaultgw"], _ = expandObjectFspVlanInterfaceDefaultgw(d, i["defaultgw"], pre_append)
	}
	pre_append = pre + ".0." + "description"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["description"], _ = expandObjectFspVlanInterfaceDescription(d, i["description"], pre_append)
	}
	pre_append = pre + ".0." + "detected_peer_mtu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["detected-peer-mtu"], _ = expandObjectFspVlanInterfaceDetectedPeerMtu(d, i["detected_peer_mtu"], pre_append)
	}
	pre_append = pre + ".0." + "detectprotocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["detectprotocol"], _ = expandObjectFspVlanInterfaceDetectprotocol(d, i["detectprotocol"], pre_append)
	} else {
		result["detectprotocol"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "detectserver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["detectserver"], _ = expandObjectFspVlanInterfaceDetectserver(d, i["detectserver"], pre_append)
	}
	pre_append = pre + ".0." + "device_access_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["device-access-list"], _ = expandObjectFspVlanInterfaceDeviceAccessList(d, i["device_access_list"], pre_append)
	}
	pre_append = pre + ".0." + "device_identification"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["device-identification"], _ = expandObjectFspVlanInterfaceDeviceIdentification(d, i["device_identification"], pre_append)
	}
	pre_append = pre + ".0." + "device_identification_active_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["device-identification-active-scan"], _ = expandObjectFspVlanInterfaceDeviceIdentificationActiveScan(d, i["device_identification_active_scan"], pre_append)
	}
	pre_append = pre + ".0." + "device_netscan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["device-netscan"], _ = expandObjectFspVlanInterfaceDeviceNetscan(d, i["device_netscan"], pre_append)
	}
	pre_append = pre + ".0." + "device_user_identification"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["device-user-identification"], _ = expandObjectFspVlanInterfaceDeviceUserIdentification(d, i["device_user_identification"], pre_append)
	}
	pre_append = pre + ".0." + "devindex"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["devindex"], _ = expandObjectFspVlanInterfaceDevindex(d, i["devindex"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_classless_route_addition"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-classless-route-addition"], _ = expandObjectFspVlanInterfaceDhcpClasslessRouteAddition(d, i["dhcp_classless_route_addition"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_client_identifier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-client-identifier"], _ = expandObjectFspVlanInterfaceDhcpClientIdentifier(d, i["dhcp_client_identifier"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_agent_option"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-agent-option"], _ = expandObjectFspVlanInterfaceDhcpRelayAgentOption(d, i["dhcp_relay_agent_option"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-interface"], _ = expandObjectFspVlanInterfaceDhcpRelayInterface(d, i["dhcp_relay_interface"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_interface_select_method"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-interface-select-method"], _ = expandObjectFspVlanInterfaceDhcpRelayInterfaceSelectMethod(d, i["dhcp_relay_interface_select_method"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-ip"], _ = expandObjectFspVlanInterfaceDhcpRelayIp(d, i["dhcp_relay_ip"], pre_append)
	} else {
		result["dhcp-relay-ip"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "dhcp_relay_link_selection"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-link-selection"], _ = expandObjectFspVlanInterfaceDhcpRelayLinkSelection(d, i["dhcp_relay_link_selection"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_request_all_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-request-all-server"], _ = expandObjectFspVlanInterfaceDhcpRelayRequestAllServer(d, i["dhcp_relay_request_all_server"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-service"], _ = expandObjectFspVlanInterfaceDhcpRelayService(d, i["dhcp_relay_service"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_relay_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-relay-type"], _ = expandObjectFspVlanInterfaceDhcpRelayType(d, i["dhcp_relay_type"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp_renew_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp-renew-time"], _ = expandObjectFspVlanInterfaceDhcpRenewTime(d, i["dhcp_renew_time"], pre_append)
	}
	pre_append = pre + ".0." + "disc_retry_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disc-retry-timeout"], _ = expandObjectFspVlanInterfaceDiscRetryTimeout(d, i["disc_retry_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandObjectFspVlanInterfaceDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "distance"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["distance"], _ = expandObjectFspVlanInterfaceDistance(d, i["distance"], pre_append)
	}
	pre_append = pre + ".0." + "dns_query"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-query"], _ = expandObjectFspVlanInterfaceDnsQuery(d, i["dns_query"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server-override"], _ = expandObjectFspVlanInterfaceDnsServerOverride(d, i["dns_server_override"], pre_append)
	}
	pre_append = pre + ".0." + "dns_server_protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-server-protocol"], _ = expandObjectFspVlanInterfaceDnsServerProtocol(d, i["dns_server_protocol"], pre_append)
	} else {
		result["dns-server-protocol"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "drop_fragment"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["drop-fragment"], _ = expandObjectFspVlanInterfaceDropFragment(d, i["drop_fragment"], pre_append)
	}
	pre_append = pre + ".0." + "drop_overlapped_fragment"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["drop-overlapped-fragment"], _ = expandObjectFspVlanInterfaceDropOverlappedFragment(d, i["drop_overlapped_fragment"], pre_append)
	}
	pre_append = pre + ".0." + "eap_ca_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["eap-ca-cert"], _ = expandObjectFspVlanInterfaceEapCaCert(d, i["eap_ca_cert"], pre_append)
	} else {
		result["eap-ca-cert"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "eap_identity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["eap-identity"], _ = expandObjectFspVlanInterfaceEapIdentity(d, i["eap_identity"], pre_append)
	}
	pre_append = pre + ".0." + "eap_method"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["eap-method"], _ = expandObjectFspVlanInterfaceEapMethod(d, i["eap_method"], pre_append)
	}
	pre_append = pre + ".0." + "eap_password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["eap-password"], _ = expandObjectFspVlanInterfaceEapPassword(d, i["eap_password"], pre_append)
	} else {
		result["eap-password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "eap_supplicant"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["eap-supplicant"], _ = expandObjectFspVlanInterfaceEapSupplicant(d, i["eap_supplicant"], pre_append)
	}
	pre_append = pre + ".0." + "eap_user_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["eap-user-cert"], _ = expandObjectFspVlanInterfaceEapUserCert(d, i["eap_user_cert"], pre_append)
	} else {
		result["eap-user-cert"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "egress_cos"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["egress-cos"], _ = expandObjectFspVlanInterfaceEgressCos(d, i["egress_cos"], pre_append)
	}
	pre_append = pre + ".0." + "egress_shaping_profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["egress-shaping-profile"], _ = expandObjectFspVlanInterfaceEgressShapingProfile(d, i["egress_shaping_profile"], pre_append)
	}
	pre_append = pre + ".0." + "eip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["eip"], _ = expandObjectFspVlanInterfaceEip(d, i["eip"], pre_append)
	}
	pre_append = pre + ".0." + "endpoint_compliance"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["endpoint-compliance"], _ = expandObjectFspVlanInterfaceEndpointCompliance(d, i["endpoint_compliance"], pre_append)
	}
	pre_append = pre + ".0." + "estimated_downstream_bandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["estimated-downstream-bandwidth"], _ = expandObjectFspVlanInterfaceEstimatedDownstreamBandwidth(d, i["estimated_downstream_bandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "estimated_upstream_bandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["estimated-upstream-bandwidth"], _ = expandObjectFspVlanInterfaceEstimatedUpstreamBandwidth(d, i["estimated_upstream_bandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "explicit_ftp_proxy"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["explicit-ftp-proxy"], _ = expandObjectFspVlanInterfaceExplicitFtpProxy(d, i["explicit_ftp_proxy"], pre_append)
	}
	pre_append = pre + ".0." + "explicit_web_proxy"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["explicit-web-proxy"], _ = expandObjectFspVlanInterfaceExplicitWebProxy(d, i["explicit_web_proxy"], pre_append)
	}
	pre_append = pre + ".0." + "external"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external"], _ = expandObjectFspVlanInterfaceExternal(d, i["external"], pre_append)
	}
	pre_append = pre + ".0." + "fail_action_on_extender"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fail-action-on-extender"], _ = expandObjectFspVlanInterfaceFailActionOnExtender(d, i["fail_action_on_extender"], pre_append)
	}
	pre_append = pre + ".0." + "fail_alert_interfaces"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fail-alert-interfaces"], _ = expandObjectFspVlanInterfaceFailAlertInterfaces(d, i["fail_alert_interfaces"], pre_append)
	}
	pre_append = pre + ".0." + "fail_alert_method"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fail-alert-method"], _ = expandObjectFspVlanInterfaceFailAlertMethod(d, i["fail_alert_method"], pre_append)
	}
	pre_append = pre + ".0." + "fail_detect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fail-detect"], _ = expandObjectFspVlanInterfaceFailDetect(d, i["fail_detect"], pre_append)
	}
	pre_append = pre + ".0." + "fail_detect_option"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fail-detect-option"], _ = expandObjectFspVlanInterfaceFailDetectOption(d, i["fail_detect_option"], pre_append)
	} else {
		result["fail-detect-option"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "fdp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fdp"], _ = expandObjectFspVlanInterfaceFdp(d, i["fdp"], pre_append)
	}
	pre_append = pre + ".0." + "fortiheartbeat"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiheartbeat"], _ = expandObjectFspVlanInterfaceFortiheartbeat(d, i["fortiheartbeat"], pre_append)
	}
	pre_append = pre + ".0." + "fortilink"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortilink"], _ = expandObjectFspVlanInterfaceFortilink(d, i["fortilink"], pre_append)
	}
	pre_append = pre + ".0." + "fortilink_backup_link"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortilink-backup-link"], _ = expandObjectFspVlanInterfaceFortilinkBackupLink(d, i["fortilink_backup_link"], pre_append)
	}
	pre_append = pre + ".0." + "fortilink_neighbor_detect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortilink-neighbor-detect"], _ = expandObjectFspVlanInterfaceFortilinkNeighborDetect(d, i["fortilink_neighbor_detect"], pre_append)
	}
	pre_append = pre + ".0." + "fortilink_split_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortilink-split-interface"], _ = expandObjectFspVlanInterfaceFortilinkSplitInterface(d, i["fortilink_split_interface"], pre_append)
	}
	pre_append = pre + ".0." + "fortilink_stacking"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortilink-stacking"], _ = expandObjectFspVlanInterfaceFortilinkStacking(d, i["fortilink_stacking"], pre_append)
	}
	pre_append = pre + ".0." + "forward_domain"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["forward-domain"], _ = expandObjectFspVlanInterfaceForwardDomain(d, i["forward_domain"], pre_append)
	}
	pre_append = pre + ".0." + "forward_error_correction"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["forward-error-correction"], _ = expandObjectFspVlanInterfaceForwardErrorCorrection(d, i["forward_error_correction"], pre_append)
	}
	pre_append = pre + ".0." + "fp_anomaly"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fp-anomaly"], _ = expandObjectFspVlanInterfaceFpAnomaly(d, i["fp_anomaly"], pre_append)
	} else {
		result["fp-anomaly"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "fp_disable"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fp-disable"], _ = expandObjectFspVlanInterfaceFpDisable(d, i["fp_disable"], pre_append)
	} else {
		result["fp-disable"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "gateway_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gateway-address"], _ = expandObjectFspVlanInterfaceGatewayAddress(d, i["gateway_address"], pre_append)
	}
	pre_append = pre + ".0." + "generic_receive_offload"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["generic-receive-offload"], _ = expandObjectFspVlanInterfaceGenericReceiveOffload(d, i["generic_receive_offload"], pre_append)
	}
	pre_append = pre + ".0." + "gi_gk"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gi-gk"], _ = expandObjectFspVlanInterfaceGiGk(d, i["gi_gk"], pre_append)
	}
	pre_append = pre + ".0." + "gwaddr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gwaddr"], _ = expandObjectFspVlanInterfaceGwaddr(d, i["gwaddr"], pre_append)
	}
	pre_append = pre + ".0." + "gwdetect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gwdetect"], _ = expandObjectFspVlanInterfaceGwdetect(d, i["gwdetect"], pre_append)
	}
	pre_append = pre + ".0." + "ha_priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ha-priority"], _ = expandObjectFspVlanInterfaceHaPriority(d, i["ha_priority"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_accept_redirect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-accept-redirect"], _ = expandObjectFspVlanInterfaceIcmpAcceptRedirect(d, i["icmp_accept_redirect"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_redirect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-redirect"], _ = expandObjectFspVlanInterfaceIcmpRedirect(d, i["icmp_redirect"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_send_redirect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-send-redirect"], _ = expandObjectFspVlanInterfaceIcmpSendRedirect(d, i["icmp_send_redirect"], pre_append)
	}
	pre_append = pre + ".0." + "ident_accept"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ident-accept"], _ = expandObjectFspVlanInterfaceIdentAccept(d, i["ident_accept"], pre_append)
	}
	pre_append = pre + ".0." + "idle_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["idle-timeout"], _ = expandObjectFspVlanInterfaceIdleTimeout(d, i["idle_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "if_mdix"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["if-mdix"], _ = expandObjectFspVlanInterfaceIfMdix(d, i["if_mdix"], pre_append)
	}
	pre_append = pre + ".0." + "if_media"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["if-media"], _ = expandObjectFspVlanInterfaceIfMedia(d, i["if_media"], pre_append)
	}
	pre_append = pre + ".0." + "ike_saml_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ike-saml-server"], _ = expandObjectFspVlanInterfaceIkeSamlServer(d, i["ike_saml_server"], pre_append)
	} else {
		result["ike-saml-server"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "in_force_vlan_cos"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["in-force-vlan-cos"], _ = expandObjectFspVlanInterfaceInForceVlanCos(d, i["in_force_vlan_cos"], pre_append)
	}
	pre_append = pre + ".0." + "inbandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inbandwidth"], _ = expandObjectFspVlanInterfaceInbandwidth(d, i["inbandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "ingress_cos"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ingress-cos"], _ = expandObjectFspVlanInterfaceIngressCos(d, i["ingress_cos"], pre_append)
	}
	pre_append = pre + ".0." + "ingress_shaping_profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ingress-shaping-profile"], _ = expandObjectFspVlanInterfaceIngressShapingProfile(d, i["ingress_shaping_profile"], pre_append)
	}
	pre_append = pre + ".0." + "ingress_spillover_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ingress-spillover-threshold"], _ = expandObjectFspVlanInterfaceIngressSpilloverThreshold(d, i["ingress_spillover_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "interconnect_profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interconnect-profile"], _ = expandObjectFspVlanInterfaceInterconnectProfile(d, i["interconnect_profile"], pre_append)
	}
	pre_append = pre + ".0." + "internal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["internal"], _ = expandObjectFspVlanInterfaceInternal(d, i["internal"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandObjectFspVlanInterfaceIp(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "ip_managed_by_fortiipam"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip-managed-by-fortiipam"], _ = expandObjectFspVlanInterfaceIpManagedByFortiipam(d, i["ip_managed_by_fortiipam"], pre_append)
	}
	pre_append = pre + ".0." + "ipmac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipmac"], _ = expandObjectFspVlanInterfaceIpmac(d, i["ipmac"], pre_append)
	}
	pre_append = pre + ".0." + "ips_sniffer_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ips-sniffer-mode"], _ = expandObjectFspVlanInterfaceIpsSnifferMode(d, i["ips_sniffer_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ipunnumbered"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipunnumbered"], _ = expandObjectFspVlanInterfaceIpunnumbered(d, i["ipunnumbered"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6"], _ = expandObjectFspVlanInterfaceIpv6(d, i["ipv6"], pre_append)
	} else {
		result["ipv6"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "l2forward"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l2forward"], _ = expandObjectFspVlanInterfaceL2Forward(d, i["l2forward"], pre_append)
	}
	pre_append = pre + ".0." + "l2tp_client"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l2tp-client"], _ = expandObjectFspVlanInterfaceL2TpClient(d, i["l2tp_client"], pre_append)
	}
	pre_append = pre + ".0." + "lacp_ha_secondary"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lacp-ha-secondary"], _ = expandObjectFspVlanInterfaceLacpHaSecondary(d, i["lacp_ha_secondary"], pre_append)
	}
	pre_append = pre + ".0." + "lacp_ha_slave"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lacp-ha-slave"], _ = expandObjectFspVlanInterfaceLacpHaSlave(d, i["lacp_ha_slave"], pre_append)
	}
	pre_append = pre + ".0." + "lacp_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lacp-mode"], _ = expandObjectFspVlanInterfaceLacpMode(d, i["lacp_mode"], pre_append)
	}
	pre_append = pre + ".0." + "lacp_speed"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lacp-speed"], _ = expandObjectFspVlanInterfaceLacpSpeed(d, i["lacp_speed"], pre_append)
	}
	pre_append = pre + ".0." + "large_receive_offload"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["large-receive-offload"], _ = expandObjectFspVlanInterfaceLargeReceiveOffload(d, i["large_receive_offload"], pre_append)
	}
	pre_append = pre + ".0." + "lcp_echo_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lcp-echo-interval"], _ = expandObjectFspVlanInterfaceLcpEchoInterval(d, i["lcp_echo_interval"], pre_append)
	}
	pre_append = pre + ".0." + "lcp_max_echo_fails"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lcp-max-echo-fails"], _ = expandObjectFspVlanInterfaceLcpMaxEchoFails(d, i["lcp_max_echo_fails"], pre_append)
	}
	pre_append = pre + ".0." + "link_up_delay"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["link-up-delay"], _ = expandObjectFspVlanInterfaceLinkUpDelay(d, i["link_up_delay"], pre_append)
	}
	pre_append = pre + ".0." + "listen_forticlient_connection"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["listen-forticlient-connection"], _ = expandObjectFspVlanInterfaceListenForticlientConnection(d, i["listen_forticlient_connection"], pre_append)
	}
	pre_append = pre + ".0." + "lldp_network_policy"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lldp-network-policy"], _ = expandObjectFspVlanInterfaceLldpNetworkPolicy(d, i["lldp_network_policy"], pre_append)
	}
	pre_append = pre + ".0." + "lldp_reception"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lldp-reception"], _ = expandObjectFspVlanInterfaceLldpReception(d, i["lldp_reception"], pre_append)
	}
	pre_append = pre + ".0." + "lldp_transmission"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lldp-transmission"], _ = expandObjectFspVlanInterfaceLldpTransmission(d, i["lldp_transmission"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectFspVlanInterfaceLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "macaddr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["macaddr"], _ = expandObjectFspVlanInterfaceMacaddr(d, i["macaddr"], pre_append)
	}
	pre_append = pre + ".0." + "managed_subnetwork_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["managed-subnetwork-size"], _ = expandObjectFspVlanInterfaceManagedSubnetworkSize(d, i["managed_subnetwork_size"], pre_append)
	}
	pre_append = pre + ".0." + "management_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["management-ip"], _ = expandObjectFspVlanInterfaceManagementIp(d, i["management_ip"], pre_append)
	}
	pre_append = pre + ".0." + "max_egress_burst_rate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-egress-burst-rate"], _ = expandObjectFspVlanInterfaceMaxEgressBurstRate(d, i["max_egress_burst_rate"], pre_append)
	}
	pre_append = pre + ".0." + "max_egress_rate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-egress-rate"], _ = expandObjectFspVlanInterfaceMaxEgressRate(d, i["max_egress_rate"], pre_append)
	}
	pre_append = pre + ".0." + "measured_downstream_bandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["measured-downstream-bandwidth"], _ = expandObjectFspVlanInterfaceMeasuredDownstreamBandwidth(d, i["measured_downstream_bandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "measured_upstream_bandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["measured-upstream-bandwidth"], _ = expandObjectFspVlanInterfaceMeasuredUpstreamBandwidth(d, i["measured_upstream_bandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "mediatype"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mediatype"], _ = expandObjectFspVlanInterfaceMediatype(d, i["mediatype"], pre_append)
	}
	pre_append = pre + ".0." + "member"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["member"], _ = expandObjectFspVlanInterfaceMember(d, i["member"], pre_append)
	}
	pre_append = pre + ".0." + "min_links"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-links"], _ = expandObjectFspVlanInterfaceMinLinks(d, i["min_links"], pre_append)
	}
	pre_append = pre + ".0." + "min_links_down"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-links-down"], _ = expandObjectFspVlanInterfaceMinLinksDown(d, i["min_links_down"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandObjectFspVlanInterfaceMode(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "monitor_bandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["monitor-bandwidth"], _ = expandObjectFspVlanInterfaceMonitorBandwidth(d, i["monitor_bandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "mtu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mtu"], _ = expandObjectFspVlanInterfaceMtu(d, i["mtu"], pre_append)
	}
	pre_append = pre + ".0." + "mtu_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mtu-override"], _ = expandObjectFspVlanInterfaceMtuOverride(d, i["mtu_override"], pre_append)
	}
	pre_append = pre + ".0." + "mux_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mux-type"], _ = expandObjectFspVlanInterfaceMuxType(d, i["mux_type"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandObjectFspVlanInterfaceName(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "ndiscforward"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ndiscforward"], _ = expandObjectFspVlanInterfaceNdiscforward(d, i["ndiscforward"], pre_append)
	}
	pre_append = pre + ".0." + "netbios_forward"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["netbios-forward"], _ = expandObjectFspVlanInterfaceNetbiosForward(d, i["netbios_forward"], pre_append)
	}
	pre_append = pre + ".0." + "netflow_sampler"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["netflow-sampler"], _ = expandObjectFspVlanInterfaceNetflowSampler(d, i["netflow_sampler"], pre_append)
	}
	pre_append = pre + ".0." + "np_qos_profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["np-qos-profile"], _ = expandObjectFspVlanInterfaceNpQosProfile(d, i["np_qos_profile"], pre_append)
	}
	pre_append = pre + ".0." + "npu_fastpath"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["npu-fastpath"], _ = expandObjectFspVlanInterfaceNpuFastpath(d, i["npu_fastpath"], pre_append)
	}
	pre_append = pre + ".0." + "nst"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nst"], _ = expandObjectFspVlanInterfaceNst(d, i["nst"], pre_append)
	}
	pre_append = pre + ".0." + "out_force_vlan_cos"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["out-force-vlan-cos"], _ = expandObjectFspVlanInterfaceOutForceVlanCos(d, i["out_force_vlan_cos"], pre_append)
	}
	pre_append = pre + ".0." + "outbandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbandwidth"], _ = expandObjectFspVlanInterfaceOutbandwidth(d, i["outbandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "padt_retry_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["padt-retry-timeout"], _ = expandObjectFspVlanInterfacePadtRetryTimeout(d, i["padt_retry_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password"], _ = expandObjectFspVlanInterfacePassword(d, i["password"], pre_append)
	} else {
		result["password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "peer_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["peer-interface"], _ = expandObjectFspVlanInterfacePeerInterface(d, i["peer_interface"], pre_append)
	}
	pre_append = pre + ".0." + "phy_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["phy-mode"], _ = expandObjectFspVlanInterfacePhyMode(d, i["phy_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ping_serv_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ping-serv-status"], _ = expandObjectFspVlanInterfacePingServStatus(d, i["ping_serv_status"], pre_append)
	}
	pre_append = pre + ".0." + "poe"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["poe"], _ = expandObjectFspVlanInterfacePoe(d, i["poe"], pre_append)
	}
	pre_append = pre + ".0." + "polling_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["polling-interval"], _ = expandObjectFspVlanInterfacePollingInterval(d, i["polling_interval"], pre_append)
	}
	pre_append = pre + ".0." + "pppoe_unnumbered_negotiate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pppoe-unnumbered-negotiate"], _ = expandObjectFspVlanInterfacePppoeUnnumberedNegotiate(d, i["pppoe_unnumbered_negotiate"], pre_append)
	}
	pre_append = pre + ".0." + "pptp_auth_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pptp-auth-type"], _ = expandObjectFspVlanInterfacePptpAuthType(d, i["pptp_auth_type"], pre_append)
	}
	pre_append = pre + ".0." + "pptp_client"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pptp-client"], _ = expandObjectFspVlanInterfacePptpClient(d, i["pptp_client"], pre_append)
	}
	pre_append = pre + ".0." + "pptp_password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pptp-password"], _ = expandObjectFspVlanInterfacePptpPassword(d, i["pptp_password"], pre_append)
	} else {
		result["pptp-password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "pptp_server_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pptp-server-ip"], _ = expandObjectFspVlanInterfacePptpServerIp(d, i["pptp_server_ip"], pre_append)
	}
	pre_append = pre + ".0." + "pptp_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pptp-timeout"], _ = expandObjectFspVlanInterfacePptpTimeout(d, i["pptp_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "pptp_user"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pptp-user"], _ = expandObjectFspVlanInterfacePptpUser(d, i["pptp_user"], pre_append)
	}
	pre_append = pre + ".0." + "preserve_session_route"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preserve-session-route"], _ = expandObjectFspVlanInterfacePreserveSessionRoute(d, i["preserve_session_route"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandObjectFspVlanInterfacePriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "priority_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority-override"], _ = expandObjectFspVlanInterfacePriorityOverride(d, i["priority_override"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_captive_portal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["proxy-captive-portal"], _ = expandObjectFspVlanInterfaceProxyCaptivePortal(d, i["proxy_captive_portal"], pre_append)
	}
	pre_append = pre + ".0." + "pvc_atm_qos"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pvc-atm-qos"], _ = expandObjectFspVlanInterfacePvcAtmQos(d, i["pvc_atm_qos"], pre_append)
	}
	pre_append = pre + ".0." + "pvc_chan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pvc-chan"], _ = expandObjectFspVlanInterfacePvcChan(d, i["pvc_chan"], pre_append)
	}
	pre_append = pre + ".0." + "pvc_crc"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pvc-crc"], _ = expandObjectFspVlanInterfacePvcCrc(d, i["pvc_crc"], pre_append)
	}
	pre_append = pre + ".0." + "pvc_pcr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pvc-pcr"], _ = expandObjectFspVlanInterfacePvcPcr(d, i["pvc_pcr"], pre_append)
	}
	pre_append = pre + ".0." + "pvc_scr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pvc-scr"], _ = expandObjectFspVlanInterfacePvcScr(d, i["pvc_scr"], pre_append)
	}
	pre_append = pre + ".0." + "pvc_vlan_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pvc-vlan-id"], _ = expandObjectFspVlanInterfacePvcVlanId(d, i["pvc_vlan_id"], pre_append)
	}
	pre_append = pre + ".0." + "pvc_vlan_rx_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pvc-vlan-rx-id"], _ = expandObjectFspVlanInterfacePvcVlanRxId(d, i["pvc_vlan_rx_id"], pre_append)
	}
	pre_append = pre + ".0." + "pvc_vlan_rx_op"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pvc-vlan-rx-op"], _ = expandObjectFspVlanInterfacePvcVlanRxOp(d, i["pvc_vlan_rx_op"], pre_append)
	}
	pre_append = pre + ".0." + "pvc_vlan_tx_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pvc-vlan-tx-id"], _ = expandObjectFspVlanInterfacePvcVlanTxId(d, i["pvc_vlan_tx_id"], pre_append)
	}
	pre_append = pre + ".0." + "pvc_vlan_tx_op"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pvc-vlan-tx-op"], _ = expandObjectFspVlanInterfacePvcVlanTxOp(d, i["pvc_vlan_tx_op"], pre_append)
	}
	pre_append = pre + ".0." + "reachable_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["reachable-time"], _ = expandObjectFspVlanInterfaceReachableTime(d, i["reachable_time"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-interface"], _ = expandObjectFspVlanInterfaceRedundantInterface(d, i["redundant_interface"], pre_append)
	}
	pre_append = pre + ".0." + "remote_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["remote-ip"], _ = expandObjectFspVlanInterfaceRemoteIp(d, i["remote_ip"], pre_append)
	}
	pre_append = pre + ".0." + "replacemsg_override_group"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["replacemsg-override-group"], _ = expandObjectFspVlanInterfaceReplacemsgOverrideGroup(d, i["replacemsg_override_group"], pre_append)
	}
	pre_append = pre + ".0." + "retransmission"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["retransmission"], _ = expandObjectFspVlanInterfaceRetransmission(d, i["retransmission"], pre_append)
	}
	pre_append = pre + ".0." + "ring_rx"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ring-rx"], _ = expandObjectFspVlanInterfaceRingRx(d, i["ring_rx"], pre_append)
	}
	pre_append = pre + ".0." + "ring_tx"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ring-tx"], _ = expandObjectFspVlanInterfaceRingTx(d, i["ring_tx"], pre_append)
	}
	pre_append = pre + ".0." + "role"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["role"], _ = expandObjectFspVlanInterfaceRole(d, i["role"], pre_append)
	}
	pre_append = pre + ".0." + "sample_direction"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sample-direction"], _ = expandObjectFspVlanInterfaceSampleDirection(d, i["sample_direction"], pre_append)
	}
	pre_append = pre + ".0." + "sample_rate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sample-rate"], _ = expandObjectFspVlanInterfaceSampleRate(d, i["sample_rate"], pre_append)
	}
	pre_append = pre + ".0." + "scan_botnet_connections"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-botnet-connections"], _ = expandObjectFspVlanInterfaceScanBotnetConnections(d, i["scan_botnet_connections"], pre_append)
	}
	pre_append = pre + ".0." + "secondary_IP"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["secondary-IP"], _ = expandObjectFspVlanInterfaceSecondaryIp(d, i["secondary_IP"], pre_append)
	}
	pre_append = pre + ".0." + "secondaryip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["secondaryip"], _ = expandObjectFspVlanInterfaceSecondaryip(d, i["secondaryip"], pre_append)
	} else {
		result["secondaryip"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "security_8021x_dynamic_vlan_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["security-8021x-dynamic-vlan-id"], _ = expandObjectFspVlanInterfaceSecurity8021XDynamicVlanId(d, i["security_8021x_dynamic_vlan_id"], pre_append)
	}
	pre_append = pre + ".0." + "security_8021x_master"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["security-8021x-master"], _ = expandObjectFspVlanInterfaceSecurity8021XMaster(d, i["security_8021x_master"], pre_append)
	}
	pre_append = pre + ".0." + "security_8021x_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["security-8021x-mode"], _ = expandObjectFspVlanInterfaceSecurity8021XMode(d, i["security_8021x_mode"], pre_append)
	}
	pre_append = pre + ".0." + "security_exempt_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["security-exempt-list"], _ = expandObjectFspVlanInterfaceSecurityExemptList(d, i["security_exempt_list"], pre_append)
	}
	pre_append = pre + ".0." + "security_external_logout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["security-external-logout"], _ = expandObjectFspVlanInterfaceSecurityExternalLogout(d, i["security_external_logout"], pre_append)
	}
	pre_append = pre + ".0." + "security_external_web"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["security-external-web"], _ = expandObjectFspVlanInterfaceSecurityExternalWeb(d, i["security_external_web"], pre_append)
	}
	pre_append = pre + ".0." + "security_groups"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["security-groups"], _ = expandObjectFspVlanInterfaceSecurityGroups(d, i["security_groups"], pre_append)
	}
	pre_append = pre + ".0." + "security_mac_auth_bypass"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["security-mac-auth-bypass"], _ = expandObjectFspVlanInterfaceSecurityMacAuthBypass(d, i["security_mac_auth_bypass"], pre_append)
	}
	pre_append = pre + ".0." + "security_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["security-mode"], _ = expandObjectFspVlanInterfaceSecurityMode(d, i["security_mode"], pre_append)
	}
	pre_append = pre + ".0." + "security_redirect_url"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["security-redirect-url"], _ = expandObjectFspVlanInterfaceSecurityRedirectUrl(d, i["security_redirect_url"], pre_append)
	}
	pre_append = pre + ".0." + "select_profile_30a_35b"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["select-profile-30a-35b"], _ = expandObjectFspVlanInterfaceSelectProfile30A35B(d, i["select_profile_30a_35b"], pre_append)
	}
	pre_append = pre + ".0." + "service_name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["service-name"], _ = expandObjectFspVlanInterfaceServiceName(d, i["service_name"], pre_append)
	}
	pre_append = pre + ".0." + "sflow_sampler"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sflow-sampler"], _ = expandObjectFspVlanInterfaceSflowSampler(d, i["sflow_sampler"], pre_append)
	}
	pre_append = pre + ".0." + "sfp_dsl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sfp-dsl"], _ = expandObjectFspVlanInterfaceSfpDsl(d, i["sfp_dsl"], pre_append)
	}
	pre_append = pre + ".0." + "sfp_dsl_adsl_fallback"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sfp-dsl-adsl-fallback"], _ = expandObjectFspVlanInterfaceSfpDslAdslFallback(d, i["sfp_dsl_adsl_fallback"], pre_append)
	}
	pre_append = pre + ".0." + "sfp_dsl_autodetect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sfp-dsl-autodetect"], _ = expandObjectFspVlanInterfaceSfpDslAutodetect(d, i["sfp_dsl_autodetect"], pre_append)
	}
	pre_append = pre + ".0." + "sfp_dsl_mac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sfp-dsl-mac"], _ = expandObjectFspVlanInterfaceSfpDslMac(d, i["sfp_dsl_mac"], pre_append)
	}
	pre_append = pre + ".0." + "speed"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["speed"], _ = expandObjectFspVlanInterfaceSpeed(d, i["speed"], pre_append)
	}
	pre_append = pre + ".0." + "spillover_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["spillover-threshold"], _ = expandObjectFspVlanInterfaceSpilloverThreshold(d, i["spillover_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "src_check"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-check"], _ = expandObjectFspVlanInterfaceSrcCheck(d, i["src_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectFspVlanInterfaceStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "stp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stp"], _ = expandObjectFspVlanInterfaceStp(d, i["stp"], pre_append)
	}
	pre_append = pre + ".0." + "stp_ha_secondary"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stp-ha-secondary"], _ = expandObjectFspVlanInterfaceStpHaSecondary(d, i["stp_ha_secondary"], pre_append)
	}
	pre_append = pre + ".0." + "stp_ha_slave"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stp-ha-slave"], _ = expandObjectFspVlanInterfaceStpHaSlave(d, i["stp_ha_slave"], pre_append)
	}
	pre_append = pre + ".0." + "stpforward"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stpforward"], _ = expandObjectFspVlanInterfaceStpforward(d, i["stpforward"], pre_append)
	}
	pre_append = pre + ".0." + "stpforward_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stpforward-mode"], _ = expandObjectFspVlanInterfaceStpforwardMode(d, i["stpforward_mode"], pre_append)
	}
	pre_append = pre + ".0." + "strip_priority_vlan_tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["strip-priority-vlan-tag"], _ = expandObjectFspVlanInterfaceStripPriorityVlanTag(d, i["strip_priority_vlan_tag"], pre_append)
	}
	pre_append = pre + ".0." + "subst"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["subst"], _ = expandObjectFspVlanInterfaceSubst(d, i["subst"], pre_append)
	}
	pre_append = pre + ".0." + "substitute_dst_mac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["substitute-dst-mac"], _ = expandObjectFspVlanInterfaceSubstituteDstMac(d, i["substitute_dst_mac"], pre_append)
	}
	pre_append = pre + ".0." + "sw_algorithm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sw-algorithm"], _ = expandObjectFspVlanInterfaceSwAlgorithm(d, i["sw_algorithm"], pre_append)
	}
	pre_append = pre + ".0." + "swc_first_create"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["swc-first-create"], _ = expandObjectFspVlanInterfaceSwcFirstCreate(d, i["swc_first_create"], pre_append)
	}
	pre_append = pre + ".0." + "swc_vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["swc-vlan"], _ = expandObjectFspVlanInterfaceSwcVlan(d, i["swc_vlan"], pre_append)
	}
	pre_append = pre + ".0." + "switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch"], _ = expandObjectFspVlanInterfaceSwitch(d, i["switch"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_access_vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-access-vlan"], _ = expandObjectFspVlanInterfaceSwitchControllerAccessVlan(d, i["switch_controller_access_vlan"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_arp_inspection"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-arp-inspection"], _ = expandObjectFspVlanInterfaceSwitchControllerArpInspection(d, i["switch_controller_arp_inspection"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_auth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-auth"], _ = expandObjectFspVlanInterfaceSwitchControllerAuth(d, i["switch_controller_auth"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_dhcp_snooping"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-dhcp-snooping"], _ = expandObjectFspVlanInterfaceSwitchControllerDhcpSnooping(d, i["switch_controller_dhcp_snooping"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_dhcp_snooping_option82"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-dhcp-snooping-option82"], _ = expandObjectFspVlanInterfaceSwitchControllerDhcpSnoopingOption82(d, i["switch_controller_dhcp_snooping_option82"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_dhcp_snooping_verify_mac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-dhcp-snooping-verify-mac"], _ = expandObjectFspVlanInterfaceSwitchControllerDhcpSnoopingVerifyMac(d, i["switch_controller_dhcp_snooping_verify_mac"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_dynamic"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-dynamic"], _ = expandObjectFspVlanInterfaceSwitchControllerDynamic(d, i["switch_controller_dynamic"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_feature"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-feature"], _ = expandObjectFspVlanInterfaceSwitchControllerFeature(d, i["switch_controller_feature"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_igmp_snooping"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-igmp-snooping"], _ = expandObjectFspVlanInterfaceSwitchControllerIgmpSnooping(d, i["switch_controller_igmp_snooping"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_igmp_snooping_fast_leave"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-igmp-snooping-fast-leave"], _ = expandObjectFspVlanInterfaceSwitchControllerIgmpSnoopingFastLeave(d, i["switch_controller_igmp_snooping_fast_leave"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_igmp_snooping_proxy"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-igmp-snooping-proxy"], _ = expandObjectFspVlanInterfaceSwitchControllerIgmpSnoopingProxy(d, i["switch_controller_igmp_snooping_proxy"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_iot_scanning"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-iot-scanning"], _ = expandObjectFspVlanInterfaceSwitchControllerIotScanning(d, i["switch_controller_iot_scanning"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_learning_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-learning-limit"], _ = expandObjectFspVlanInterfaceSwitchControllerLearningLimit(d, i["switch_controller_learning_limit"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_mgmt_vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-mgmt-vlan"], _ = expandObjectFspVlanInterfaceSwitchControllerMgmtVlan(d, i["switch_controller_mgmt_vlan"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_nac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-nac"], _ = expandObjectFspVlanInterfaceSwitchControllerNac(d, i["switch_controller_nac"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_netflow_collect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-netflow-collect"], _ = expandObjectFspVlanInterfaceSwitchControllerNetflowCollect(d, i["switch_controller_netflow_collect"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_radius_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-radius-server"], _ = expandObjectFspVlanInterfaceSwitchControllerRadiusServer(d, i["switch_controller_radius_server"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_rspan_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-rspan-mode"], _ = expandObjectFspVlanInterfaceSwitchControllerRspanMode(d, i["switch_controller_rspan_mode"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_source_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-source-ip"], _ = expandObjectFspVlanInterfaceSwitchControllerSourceIp(d, i["switch_controller_source_ip"], pre_append)
	}
	pre_append = pre + ".0." + "switch_controller_traffic_policy"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-controller-traffic-policy"], _ = expandObjectFspVlanInterfaceSwitchControllerTrafficPolicy(d, i["switch_controller_traffic_policy"], pre_append)
	}
	pre_append = pre + ".0." + "system_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-id"], _ = expandObjectFspVlanInterfaceSystemId(d, i["system_id"], pre_append)
	}
	pre_append = pre + ".0." + "system_id_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-id-type"], _ = expandObjectFspVlanInterfaceSystemIdType(d, i["system_id_type"], pre_append)
	}
	pre_append = pre + ".0." + "tc_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tc-mode"], _ = expandObjectFspVlanInterfaceTcMode(d, i["tc_mode"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_mss"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-mss"], _ = expandObjectFspVlanInterfaceTcpMss(d, i["tcp_mss"], pre_append)
	}
	pre_append = pre + ".0." + "trunk"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trunk"], _ = expandObjectFspVlanInterfaceTrunk(d, i["trunk"], pre_append)
	}
	pre_append = pre + ".0." + "trust_ip_1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trust-ip-1"], _ = expandObjectFspVlanInterfaceTrustIp1(d, i["trust_ip_1"], pre_append)
	}
	pre_append = pre + ".0." + "trust_ip_2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trust-ip-2"], _ = expandObjectFspVlanInterfaceTrustIp2(d, i["trust_ip_2"], pre_append)
	}
	pre_append = pre + ".0." + "trust_ip_3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trust-ip-3"], _ = expandObjectFspVlanInterfaceTrustIp3(d, i["trust_ip_3"], pre_append)
	}
	pre_append = pre + ".0." + "trust_ip6_1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trust-ip6-1"], _ = expandObjectFspVlanInterfaceTrustIp61(d, i["trust_ip6_1"], pre_append)
	}
	pre_append = pre + ".0." + "trust_ip6_2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trust-ip6-2"], _ = expandObjectFspVlanInterfaceTrustIp62(d, i["trust_ip6_2"], pre_append)
	}
	pre_append = pre + ".0." + "trust_ip6_3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trust-ip6-3"], _ = expandObjectFspVlanInterfaceTrustIp63(d, i["trust_ip6_3"], pre_append)
	}
	pre_append = pre + ".0." + "type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["type"], _ = expandObjectFspVlanInterfaceType(d, i["type"], pre_append)
	}
	pre_append = pre + ".0." + "username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username"], _ = expandObjectFspVlanInterfaceUsername(d, i["username"], pre_append)
	}
	pre_append = pre + ".0." + "vci"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vci"], _ = expandObjectFspVlanInterfaceVci(d, i["vci"], pre_append)
	}
	pre_append = pre + ".0." + "vectoring"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vectoring"], _ = expandObjectFspVlanInterfaceVectoring(d, i["vectoring"], pre_append)
	}
	pre_append = pre + ".0." + "vindex"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vindex"], _ = expandObjectFspVlanInterfaceVindex(d, i["vindex"], pre_append)
	}
	pre_append = pre + ".0." + "vlan_protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan-protocol"], _ = expandObjectFspVlanInterfaceVlanProtocol(d, i["vlan_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "vlanforward"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlanforward"], _ = expandObjectFspVlanInterfaceVlanforward(d, i["vlanforward"], pre_append)
	}
	pre_append = pre + ".0." + "vlanid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlanid"], _ = expandObjectFspVlanInterfaceVlanid(d, i["vlanid"], pre_append)
	}
	pre_append = pre + ".0." + "vpi"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vpi"], _ = expandObjectFspVlanInterfaceVpi(d, i["vpi"], pre_append)
	}
	pre_append = pre + ".0." + "vrf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrf"], _ = expandObjectFspVlanInterfaceVrf(d, i["vrf"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrrp"], _ = expandObjectFspVlanInterfaceVrrp(d, i["vrrp"], pre_append)
	} else {
		result["vrrp"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "vrrp_virtual_mac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrrp-virtual-mac"], _ = expandObjectFspVlanInterfaceVrrpVirtualMac(d, i["vrrp_virtual_mac"], pre_append)
	}
	pre_append = pre + ".0." + "wccp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wccp"], _ = expandObjectFspVlanInterfaceWccp(d, i["wccp"], pre_append)
	}
	pre_append = pre + ".0." + "weight"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["weight"], _ = expandObjectFspVlanInterfaceWeight(d, i["weight"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_5g_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-5g-threshold"], _ = expandObjectFspVlanInterfaceWifi5GThreshold(d, i["wifi_5g_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_acl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-acl"], _ = expandObjectFspVlanInterfaceWifiAcl(d, i["wifi_acl"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ap_band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ap-band"], _ = expandObjectFspVlanInterfaceWifiApBand(d, i["wifi_ap_band"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_auth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-auth"], _ = expandObjectFspVlanInterfaceWifiAuth(d, i["wifi_auth"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_auto_connect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-auto-connect"], _ = expandObjectFspVlanInterfaceWifiAutoConnect(d, i["wifi_auto_connect"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_auto_save"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-auto-save"], _ = expandObjectFspVlanInterfaceWifiAutoSave(d, i["wifi_auto_save"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_broadcast_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-broadcast-ssid"], _ = expandObjectFspVlanInterfaceWifiBroadcastSsid(d, i["wifi_broadcast_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_encrypt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-encrypt"], _ = expandObjectFspVlanInterfaceWifiEncrypt(d, i["wifi_encrypt"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_fragment_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-fragment-threshold"], _ = expandObjectFspVlanInterfaceWifiFragmentThreshold(d, i["wifi_fragment_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_key"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-key"], _ = expandObjectFspVlanInterfaceWifiKey(d, i["wifi_key"], pre_append)
	} else {
		result["wifi-key"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "wifi_keyindex"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-keyindex"], _ = expandObjectFspVlanInterfaceWifiKeyindex(d, i["wifi_keyindex"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_mac_filter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-mac-filter"], _ = expandObjectFspVlanInterfaceWifiMacFilter(d, i["wifi_mac_filter"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_passphrase"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-passphrase"], _ = expandObjectFspVlanInterfaceWifiPassphrase(d, i["wifi_passphrase"], pre_append)
	} else {
		result["wifi-passphrase"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "wifi_radius_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-radius-server"], _ = expandObjectFspVlanInterfaceWifiRadiusServer(d, i["wifi_radius_server"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_rts_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-rts-threshold"], _ = expandObjectFspVlanInterfaceWifiRtsThreshold(d, i["wifi_rts_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_security"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-security"], _ = expandObjectFspVlanInterfaceWifiSecurity(d, i["wifi_security"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-ssid"], _ = expandObjectFspVlanInterfaceWifiSsid(d, i["wifi_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "wifi_usergroup"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wifi-usergroup"], _ = expandObjectFspVlanInterfaceWifiUsergroup(d, i["wifi_usergroup"], pre_append)
	}
	pre_append = pre + ".0." + "wins_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["wins-ip"], _ = expandObjectFspVlanInterfaceWinsIp(d, i["wins_ip"], pre_append)
	}

	return result, nil
}

func expandObjectFspVlanInterfaceVlanOpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAggregate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAggregateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceApDiscover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceArpforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAtmProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAuthCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAuthPortalAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceAutoAuthExtensionDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBandwidthMeasureTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBfdDesiredMinTx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBfdDetectMult(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBfdRequiredMinRx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBroadcastForticlientDiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceBroadcastForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceCliConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsKeyname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceDdnsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDdnsZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDedicatedTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDefaultgw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDetectedPeerMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDetectprotocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceDetectserver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDeviceAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDeviceIdentification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDeviceIdentificationActiveScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDeviceNetscan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDeviceUserIdentification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDevindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpClasslessRouteAddition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpClientIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayAgentOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceDhcpRelayLinkSelection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayRequestAllServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDhcpRenewTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDiscRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDnsQuery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDnsServerOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDnsServerProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceDropFragment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceDropOverlappedFragment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEapCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceEapIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEapMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEapPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceEapSupplicant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEapUserCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceEgressCos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEgressShapingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEndpointCompliance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEstimatedDownstreamBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceEstimatedUpstreamBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceExplicitFtpProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceExplicitWebProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFailActionOnExtender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFailAlertInterfaces(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFailAlertMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFailDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFailDetectOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceFdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortiheartbeat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortilink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortilinkBackupLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortilinkNeighborDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortilinkSplitInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFortilinkStacking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceForwardDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceForwardErrorCorrection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceFpAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceFpDisable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceGatewayAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceGenericReceiveOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceGiGk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceGwaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceGwdetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIcmpAcceptRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIcmpRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIcmpSendRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIdentAccept(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIfMdix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIfMedia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIkeSamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceInForceVlanCos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceInbandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIngressCos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIngressShapingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceInterconnectProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceInternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpManagedByFortiipam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpmac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpsSnifferMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpunnumbered(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "autoconf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["autoconf"], _ = expandObjectFspVlanInterfaceIpv6Autoconf(d, i["autoconf"], pre_append)
	}
	pre_append = pre + ".0." + "cli_conn6_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cli-conn6-status"], _ = expandObjectFspVlanInterfaceIpv6CliConn6Status(d, i["cli_conn6_status"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-client-options"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6ClientOptions(d, i["dhcp6_client_options"], pre_append)
	} else {
		result["dhcp6-client-options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-information-request"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6InformationRequest(d, i["dhcp6_information_request"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-delegation"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation(d, i["dhcp6_prefix_delegation"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHint(d, i["dhcp6_prefix_hint"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint-plt"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt(d, i["dhcp6_prefix_hint_plt"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-prefix-hint-vlt"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt(d, i["dhcp6_prefix_hint_vlt"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-ip"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6RelayIp(d, i["dhcp6_relay_ip"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-service"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6RelayService(d, i["dhcp6_relay_service"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dhcp6-relay-type"], _ = expandObjectFspVlanInterfaceIpv6Dhcp6RelayType(d, i["dhcp6_relay_type"], pre_append)
	}
	pre_append = pre + ".0." + "icmp6_send_redirect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp6-send-redirect"], _ = expandObjectFspVlanInterfaceIpv6Icmp6SendRedirect(d, i["icmp6_send_redirect"], pre_append)
	}
	pre_append = pre + ".0." + "interface_identifier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interface-identifier"], _ = expandObjectFspVlanInterfaceIpv6InterfaceIdentifier(d, i["interface_identifier"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-address"], _ = expandObjectFspVlanInterfaceIpv6Ip6Address(d, i["ip6_address"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-allowaccess"], _ = expandObjectFspVlanInterfaceIpv6Ip6Allowaccess(d, i["ip6_allowaccess"], pre_append)
	} else {
		result["ip6-allowaccess"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-default-life"], _ = expandObjectFspVlanInterfaceIpv6Ip6DefaultLife(d, i["ip6_default_life"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_delegated_prefix_iaid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-delegated-prefix-iaid"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid(d, i["ip6_delegated_prefix_iaid"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-delegated-prefix-list"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(d, i["ip6_delegated_prefix_list"], pre_append)
	} else {
		result["ip6-delegated-prefix-list"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-dns-server-override"], _ = expandObjectFspVlanInterfaceIpv6Ip6DnsServerOverride(d, i["ip6_dns_server_override"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-extra-addr"], _ = expandObjectFspVlanInterfaceIpv6Ip6ExtraAddr(d, i["ip6_extra_addr"], pre_append)
	} else {
		result["ip6-extra-addr"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-hop-limit"], _ = expandObjectFspVlanInterfaceIpv6Ip6HopLimit(d, i["ip6_hop_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-link-mtu"], _ = expandObjectFspVlanInterfaceIpv6Ip6LinkMtu(d, i["ip6_link_mtu"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-manage-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6ManageFlag(d, i["ip6_manage_flag"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-max-interval"], _ = expandObjectFspVlanInterfaceIpv6Ip6MaxInterval(d, i["ip6_max_interval"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-min-interval"], _ = expandObjectFspVlanInterfaceIpv6Ip6MinInterval(d, i["ip6_min_interval"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-mode"], _ = expandObjectFspVlanInterfaceIpv6Ip6Mode(d, i["ip6_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-other-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6OtherFlag(d, i["ip6_other_flag"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-prefix-list"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixList(d, i["ip6_prefix_list"], pre_append)
	} else {
		result["ip6-prefix-list"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_prefix_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-prefix-mode"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixMode(d, i["ip6_prefix_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-reachable-time"], _ = expandObjectFspVlanInterfaceIpv6Ip6ReachableTime(d, i["ip6_reachable_time"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-retrans-time"], _ = expandObjectFspVlanInterfaceIpv6Ip6RetransTime(d, i["ip6_retrans_time"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-send-adv"], _ = expandObjectFspVlanInterfaceIpv6Ip6SendAdv(d, i["ip6_send_adv"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-subnet"], _ = expandObjectFspVlanInterfaceIpv6Ip6Subnet(d, i["ip6_subnet"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-upstream-interface"], _ = expandObjectFspVlanInterfaceIpv6Ip6UpstreamInterface(d, i["ip6_upstream_interface"], pre_append)
	}
	pre_append = pre + ".0." + "nd_cert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-cert"], _ = expandObjectFspVlanInterfaceIpv6NdCert(d, i["nd_cert"], pre_append)
	}
	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-cga-modifier"], _ = expandObjectFspVlanInterfaceIpv6NdCgaModifier(d, i["nd_cga_modifier"], pre_append)
	}
	pre_append = pre + ".0." + "nd_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-mode"], _ = expandObjectFspVlanInterfaceIpv6NdMode(d, i["nd_mode"], pre_append)
	}
	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-security-level"], _ = expandObjectFspVlanInterfaceIpv6NdSecurityLevel(d, i["nd_security_level"], pre_append)
	}
	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-timestamp-delta"], _ = expandObjectFspVlanInterfaceIpv6NdTimestampDelta(d, i["nd_timestamp_delta"], pre_append)
	}
	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nd-timestamp-fuzz"], _ = expandObjectFspVlanInterfaceIpv6NdTimestampFuzz(d, i["nd_timestamp_fuzz"], pre_append)
	}
	pre_append = pre + ".0." + "ra_send_mtu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ra-send-mtu"], _ = expandObjectFspVlanInterfaceIpv6RaSendMtu(d, i["ra_send_mtu"], pre_append)
	}
	pre_append = pre + ".0." + "unique_autoconf_addr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unique-autoconf-addr"], _ = expandObjectFspVlanInterfaceIpv6UniqueAutoconfAddr(d, i["unique_autoconf_addr"], pre_append)
	}
	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrip6_link_local"], _ = expandObjectFspVlanInterfaceIpv6Vrip6LinkLocal(d, i["vrip6_link_local"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrrp-virtual-mac6"], _ = expandObjectFspVlanInterfaceIpv6VrrpVirtualMac6(d, i["vrrp_virtual_mac6"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vrrp6"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6(d, i["vrrp6"], pre_append)
	} else {
		result["vrrp6"] = make([]string, 0)
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Autoconf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6CliConn6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6ClientOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6InformationRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Icmp6SendRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6InterfaceIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Allowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DefaultLife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["autonomous-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["delegated-prefix-iaid"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid(d, i["delegated_prefix_iaid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-id"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId(d, i["prefix_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss(d, i["rdnss"], pre_append)
		} else {
			tmp["rdnss"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss-service"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService(d, i["rdnss_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["upstream-interface"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(d, i["upstream_interface"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DnsServerOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ExtraAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix(d, i["prefix"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6HopLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6LinkMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ManageFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6MaxInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6MinInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6OtherFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["autonomous-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dnssl"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl(d, i["dnssl"], pre_append)
		} else {
			tmp["dnssl"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preferred-life-time"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime(d, i["preferred_life_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss(d, i["rdnss"], pre_append)
		} else {
			tmp["rdnss"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["valid-life-time"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime(d, i["valid_life_time"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ReachableTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6RetransTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6SendAdv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6UpstreamInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdCgaModifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdSecurityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdTimestampDelta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdTimestampFuzz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6RaSendMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6UniqueAutoconfAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrip6LinkLocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6VrrpVirtualMac6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["accept-mode"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6AcceptMode(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-interval"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6AdvInterval(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preempt"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Preempt(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Priority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-time"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6StartTime(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Status(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst6"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrdst6(d, i["vrdst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrgrp"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrgrp(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrid"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrid(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrip6"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrip6(d, i["vrip6"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6AcceptMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6AdvInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Preempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6StartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrdst6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrip6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceL2Forward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceL2TpClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLacpHaSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLacpHaSlave(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLacpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLacpSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLargeReceiveOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLcpEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLcpMaxEchoFails(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLinkUpDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceListenForticlientConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLldpNetworkPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLldpReception(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLldpTransmission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMacaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceManagedSubnetworkSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceManagementIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMaxEgressBurstRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMaxEgressRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMeasuredDownstreamBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMeasuredUpstreamBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMediatype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMinLinks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMinLinksDown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMonitorBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMtuOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceMuxType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNdiscforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNetbiosForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNetflowSampler(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNpQosProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNpuFastpath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceNst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceOutForceVlanCos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceOutbandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePadtRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfacePeerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePhyMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePingServStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePoe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePollingInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePppoeUnnumberedNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePptpAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePptpClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePptpPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfacePptpServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePptpTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePptpUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePreserveSessionRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePriorityOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceProxyCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcAtmQos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcCrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcPcr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcScr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcVlanRxId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcVlanRxOp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcVlanTxId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfacePvcVlanTxOp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceReachableTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRedundantInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRemoteIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceReplacemsgOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRetransmission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRingRx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRingTx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSampleDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSampleRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowaccess"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowaccess"], _ = expandObjectFspVlanInterfaceSecondaryipAllowaccess(d, i["allowaccess"], pre_append)
		} else {
			tmp["allowaccess"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detectprotocol"], _ = expandObjectFspVlanInterfaceSecondaryipDetectprotocol(d, i["detectprotocol"], pre_append)
		} else {
			tmp["detectprotocol"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detectserver"], _ = expandObjectFspVlanInterfaceSecondaryipDetectserver(d, i["detectserver"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gwdetect"], _ = expandObjectFspVlanInterfaceSecondaryipGwdetect(d, i["gwdetect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandObjectFspVlanInterfaceSecondaryipHaPriority(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanInterfaceSecondaryipId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanInterfaceSecondaryipIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ping-serv-status"], _ = expandObjectFspVlanInterfaceSecondaryipPingServStatus(d, i["ping_serv_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandObjectFspVlanInterfaceSecondaryipSeq(d, i["seq"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceSecondaryipAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceSecondaryipDetectprotocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceSecondaryipDetectserver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipGwdetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipPingServStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurity8021XDynamicVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurity8021XMaster(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurity8021XMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityExemptList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityExternalLogout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityExternalWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecurityRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSelectProfile30A35B(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSflowSampler(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSfpDsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSfpDslAdslFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSfpDslAutodetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSfpDslMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSrcCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStpHaSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStpHaSlave(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStpforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStpforwardMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceStripPriorityVlanTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSubst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSubstituteDstMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwcFirstCreate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwcVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerAccessVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerArpInspection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerDhcpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerDhcpSnoopingOption82(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerDhcpSnoopingVerifyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerDynamic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerFeature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerIgmpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerIgmpSnoopingFastLeave(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerIgmpSnoopingProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerIotScanning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerLearningLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerMgmtVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerNac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerNetflowCollect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerRspanMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSwitchControllerTrafficPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSystemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSystemIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTcMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTcpMss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrunk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp61(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp62(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceTrustIp63(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVci(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVectoring(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVlanProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVlanforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVpi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["accept-mode"], _ = expandObjectFspVlanInterfaceVrrpAcceptMode(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-interval"], _ = expandObjectFspVlanInterfaceVrrpAdvInterval(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ignore-default-route"], _ = expandObjectFspVlanInterfaceVrrpIgnoreDefaultRoute(d, i["ignore_default_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preempt"], _ = expandObjectFspVlanInterfaceVrrpPreempt(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFspVlanInterfaceVrrpPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-time"], _ = expandObjectFspVlanInterfaceVrrpStartTime(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFspVlanInterfaceVrrpStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandObjectFspVlanInterfaceVrrpVersion(d, i["version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst"], _ = expandObjectFspVlanInterfaceVrrpVrdst(d, i["vrdst"], pre_append)
		} else {
			tmp["vrdst"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst-priority"], _ = expandObjectFspVlanInterfaceVrrpVrdstPriority(d, i["vrdst_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrgrp"], _ = expandObjectFspVlanInterfaceVrrpVrgrp(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrid"], _ = expandObjectFspVlanInterfaceVrrpVrid(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrip"], _ = expandObjectFspVlanInterfaceVrrpVrip(d, i["vrip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceVrrpAcceptMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpAdvInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpIgnoreDefaultRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpPreempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpStartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrdst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceVrrpVrdstPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVirtualMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWccp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifi5GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiAcl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiApBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiAutoConnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiAutoSave(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiBroadcastSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiEncrypt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiFragmentThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceWifiKeyindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiMacFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceWifiRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiRtsThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWifiUsergroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceWinsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_dhcp_status"); ok || d.HasChange("_dhcp_status") {
		t, err := expandObjectFspVlanDhcpStatus(d, v, "_dhcp_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dhcp-status"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectFspVlanColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server"); ok || d.HasChange("dhcp_server") {
		t, err := expandObjectFspVlanDhcpServer(d, v, "dhcp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectFspVlanDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectFspVlanInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFspVlanName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandObjectFspVlanVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("vlanid"); ok || d.HasChange("vlanid") {
		t, err := expandObjectFspVlanVlanid(d, v, "vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	return &obj, nil
}
