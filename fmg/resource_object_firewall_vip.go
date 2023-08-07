// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure virtual IP for IPv4.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallVip() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallVipCreate,
		Read:   resourceObjectFirewallVipRead,
		Update: resourceObjectFirewallVipUpdate,
		Delete: resourceObjectFirewallVipDelete,

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
			"add_nat46_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"arp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_mapping_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"add_nat46_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"arp_reply": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dns_mapping_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"extaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"extintf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"extip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"extport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"gratuitous_arp_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"http_cookie_age": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"http_cookie_domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_cookie_domain_from_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_cookie_generation": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"http_cookie_path": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_cookie_share": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_ip_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_ip_header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_multiplex": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_multiplex_max_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"http_multiplex_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"http_redirect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_supported_max_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"https_cookie_secure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ipv6_mappedip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_mappedport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ldb_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mapped_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mappedip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mappedport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_embryonic_connections": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nat_source_vip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nat44": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nat46": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outlook_web_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"persistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"portforward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"portmapping_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"realservers": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"client_ip": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"health_check_proto": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"healthcheck": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"holddown_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"http_host": &schema.Schema{
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
									"max_connections": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"monitor": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"seq": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"translate_host": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"weight": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"server_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"src_filter": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"srcintf_filter": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ssl_accept_ffdhe_groups": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_cipher_suites": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cipher": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"priority": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"versions": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"ssl_client_fallback": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_client_rekey_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_client_renegotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_client_session_state_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_client_session_state_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_client_session_state_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_dh_bits": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_hpkp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_hpkp_age": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_hpkp_backup": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_hpkp_include_subdomains": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_hpkp_primary": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_hpkp_report_uri": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_hsts": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_hsts_age": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_hsts_include_subdomains": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_http_location_conversion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_http_match_host": &schema.Schema{
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
						},
						"ssl_pfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_send_empty_frags": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_server_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_server_max_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_server_min_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_server_renegotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_server_session_state_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_server_session_state_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_server_session_state_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weblogic_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"websphere_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"extaddr": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"extintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gratuitous_arp_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_cookie_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_cookie_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_cookie_domain_from_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_cookie_generation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_cookie_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_cookie_share": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_ip_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_ip_header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_multiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_multiplex_max_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_multiplex_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_supported_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_cookie_secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ipv6_mappedip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mapped_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mappedip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_embryonic_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nat_source_vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat44": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outlook_web_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"portforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portmapping_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"realservers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"healthcheck": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"holddown_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"http_host": &schema.Schema{
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
						"max_connections": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"seq": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"translate_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"src_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcintf_filter": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_accept_ffdhe_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_cipher_suites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"versions": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_client_fallback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_client_rekey_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_client_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_client_session_state_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_client_session_state_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_client_session_state_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_dh_bits": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_hpkp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_hpkp_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_hpkp_backup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_hpkp_include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_hpkp_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_hpkp_report_uri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_hsts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_hsts_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_hsts_include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_http_location_conversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_http_match_host": &schema.Schema{
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
			},
			"ssl_pfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_send_empty_frags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_server_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_cipher_suites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"versions": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_server_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_session_state_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_server_session_state_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_server_session_state_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weblogic_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"websphere_server": &schema.Schema{
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

func resourceObjectFirewallVipCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectFirewallVip(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallVip resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallVip(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallVip resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallVipRead(d, m)
}

func resourceObjectFirewallVipUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallVip(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallVip resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallVip(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallVip resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallVipRead(d, m)
}

func resourceObjectFirewallVipDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallVip(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallVip resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallVipRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallVip(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallVip resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallVip(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallVip resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallVipAddNat46Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDnsMappingTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectFirewallVipDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat46_route"
		if _, ok := i["add-nat46-route"]; ok {
			v := flattenObjectFirewallVipDynamicMappingAddNat46Route(i["add-nat46-route"], d, pre_append)
			tmp["add_nat46_route"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-AddNat46Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := i["arp-reply"]; ok {
			v := flattenObjectFirewallVipDynamicMappingArpReply(i["arp-reply"], d, pre_append)
			tmp["arp_reply"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-ArpReply")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenObjectFirewallVipDynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectFirewallVipDynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_mapping_ttl"
		if _, ok := i["dns-mapping-ttl"]; ok {
			v := flattenObjectFirewallVipDynamicMappingDnsMappingTtl(i["dns-mapping-ttl"], d, pre_append)
			tmp["dns_mapping_ttl"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-DnsMappingTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extaddr"
		if _, ok := i["extaddr"]; ok {
			v := flattenObjectFirewallVipDynamicMappingExtaddr(i["extaddr"], d, pre_append)
			tmp["extaddr"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Extaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extintf"
		if _, ok := i["extintf"]; ok {
			v := flattenObjectFirewallVipDynamicMappingExtintf(i["extintf"], d, pre_append)
			tmp["extintf"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Extintf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := i["extip"]; ok {
			v := flattenObjectFirewallVipDynamicMappingExtip(i["extip"], d, pre_append)
			tmp["extip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Extip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := i["extport"]; ok {
			v := flattenObjectFirewallVipDynamicMappingExtport(i["extport"], d, pre_append)
			tmp["extport"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Extport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gratuitous_arp_interval"
		if _, ok := i["gratuitous-arp-interval"]; ok {
			v := flattenObjectFirewallVipDynamicMappingGratuitousArpInterval(i["gratuitous-arp-interval"], d, pre_append)
			tmp["gratuitous_arp_interval"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-GratuitousArpInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpCookieAge(i["http-cookie-age"], d, pre_append)
			tmp["http_cookie_age"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpCookieAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpCookieDomain(i["http-cookie-domain"], d, pre_append)
			tmp["http_cookie_domain"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpCookieDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
			tmp["http_cookie_domain_from_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpCookieDomainFromHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
			tmp["http_cookie_generation"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpCookieGeneration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpCookiePath(i["http-cookie-path"], d, pre_append)
			tmp["http_cookie_path"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpCookiePath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpCookieShare(i["http-cookie-share"], d, pre_append)
			tmp["http_cookie_share"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpCookieShare")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header"
		if _, ok := i["http-ip-header"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpIpHeader(i["http-ip-header"], d, pre_append)
			tmp["http_ip_header"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpIpHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header_name"
		if _, ok := i["http-ip-header-name"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpIpHeaderName(i["http-ip-header-name"], d, pre_append)
			tmp["http_ip_header_name"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpIpHeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex"
		if _, ok := i["http-multiplex"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpMultiplex(i["http-multiplex"], d, pre_append)
			tmp["http_multiplex"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpMultiplex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex_max_request"
		if _, ok := i["http-multiplex-max-request"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpMultiplexMaxRequest(i["http-multiplex-max-request"], d, pre_append)
			tmp["http_multiplex_max_request"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpMultiplexMaxRequest")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex_ttl"
		if _, ok := i["http-multiplex-ttl"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpMultiplexTtl(i["http-multiplex-ttl"], d, pre_append)
			tmp["http_multiplex_ttl"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpMultiplexTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_redirect"
		if _, ok := i["http-redirect"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpRedirect(i["http-redirect"], d, pre_append)
			tmp["http_redirect"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpRedirect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_supported_max_version"
		if _, ok := i["http-supported-max-version"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpSupportedMaxVersion(i["http-supported-max-version"], d, pre_append)
			tmp["http_supported_max_version"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpSupportedMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {
			v := flattenObjectFirewallVipDynamicMappingHttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
			tmp["https_cookie_secure"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-HttpsCookieSecure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallVipDynamicMappingId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_mappedip"
		if _, ok := i["ipv6-mappedip"]; ok {
			v := flattenObjectFirewallVipDynamicMappingIpv6Mappedip(i["ipv6-mappedip"], d, pre_append)
			tmp["ipv6_mappedip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Ipv6Mappedip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_mappedport"
		if _, ok := i["ipv6-mappedport"]; ok {
			v := flattenObjectFirewallVipDynamicMappingIpv6Mappedport(i["ipv6-mappedport"], d, pre_append)
			tmp["ipv6_mappedport"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Ipv6Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenObjectFirewallVipDynamicMappingLdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mapped_addr"
		if _, ok := i["mapped-addr"]; ok {
			v := flattenObjectFirewallVipDynamicMappingMappedAddr(i["mapped-addr"], d, pre_append)
			tmp["mapped_addr"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-MappedAddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedip"
		if _, ok := i["mappedip"]; ok {
			v := flattenObjectFirewallVipDynamicMappingMappedip(i["mappedip"], d, pre_append)
			tmp["mappedip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Mappedip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenObjectFirewallVipDynamicMappingMappedport(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_embryonic_connections"
		if _, ok := i["max-embryonic-connections"]; ok {
			v := flattenObjectFirewallVipDynamicMappingMaxEmbryonicConnections(i["max-embryonic-connections"], d, pre_append)
			tmp["max_embryonic_connections"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-MaxEmbryonicConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenObjectFirewallVipDynamicMappingMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat_source_vip"
		if _, ok := i["nat-source-vip"]; ok {
			v := flattenObjectFirewallVipDynamicMappingNatSourceVip(i["nat-source-vip"], d, pre_append)
			tmp["nat_source_vip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-NatSourceVip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat44"
		if _, ok := i["nat44"]; ok {
			v := flattenObjectFirewallVipDynamicMappingNat44(i["nat44"], d, pre_append)
			tmp["nat44"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Nat44")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat46"
		if _, ok := i["nat46"]; ok {
			v := flattenObjectFirewallVipDynamicMappingNat46(i["nat46"], d, pre_append)
			tmp["nat46"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Nat46")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "outlook_web_access"
		if _, ok := i["outlook-web-access"]; ok {
			v := flattenObjectFirewallVipDynamicMappingOutlookWebAccess(i["outlook-web-access"], d, pre_append)
			tmp["outlook_web_access"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-OutlookWebAccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {
			v := flattenObjectFirewallVipDynamicMappingPersistence(i["persistence"], d, pre_append)
			tmp["persistence"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Persistence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portforward"
		if _, ok := i["portforward"]; ok {
			v := flattenObjectFirewallVipDynamicMappingPortforward(i["portforward"], d, pre_append)
			tmp["portforward"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Portforward")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portmapping_type"
		if _, ok := i["portmapping-type"]; ok {
			v := flattenObjectFirewallVipDynamicMappingPortmappingType(i["portmapping-type"], d, pre_append)
			tmp["portmapping_type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-PortmappingType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectFirewallVipDynamicMappingProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealservers(i["realservers"], d, pre_append)
			tmp["realservers"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Realservers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := i["server-type"]; ok {
			v := flattenObjectFirewallVipDynamicMappingServerType(i["server-type"], d, pre_append)
			tmp["server_type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-ServerType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenObjectFirewallVipDynamicMappingService(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_filter"
		if _, ok := i["src-filter"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSrcFilter(i["src-filter"], d, pre_append)
			tmp["src_filter"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SrcFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf_filter"
		if _, ok := i["srcintf-filter"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSrcintfFilter(i["srcintf-filter"], d, pre_append)
			tmp["srcintf_filter"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SrcintfFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_accept_ffdhe_groups"
		if _, ok := i["ssl-accept-ffdhe-groups"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslAcceptFfdheGroups(i["ssl-accept-ffdhe-groups"], d, pre_append)
			tmp["ssl_accept_ffdhe_groups"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslAcceptFfdheGroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslAlgorithm(i["ssl-algorithm"], d, pre_append)
			tmp["ssl_algorithm"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_certificate"
		if _, ok := i["ssl-certificate"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslCertificate(i["ssl-certificate"], d, pre_append)
			tmp["ssl_certificate"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
			tmp["ssl_cipher_suites"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslCipherSuites")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_fallback"
		if _, ok := i["ssl-client-fallback"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslClientFallback(i["ssl-client-fallback"], d, pre_append)
			tmp["ssl_client_fallback"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslClientFallback")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_rekey_count"
		if _, ok := i["ssl-client-rekey-count"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslClientRekeyCount(i["ssl-client-rekey-count"], d, pre_append)
			tmp["ssl_client_rekey_count"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslClientRekeyCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_renegotiation"
		if _, ok := i["ssl-client-renegotiation"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslClientRenegotiation(i["ssl-client-renegotiation"], d, pre_append)
			tmp["ssl_client_renegotiation"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslClientRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_max"
		if _, ok := i["ssl-client-session-state-max"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslClientSessionStateMax(i["ssl-client-session-state-max"], d, pre_append)
			tmp["ssl_client_session_state_max"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslClientSessionStateMax")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_timeout"
		if _, ok := i["ssl-client-session-state-timeout"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslClientSessionStateTimeout(i["ssl-client-session-state-timeout"], d, pre_append)
			tmp["ssl_client_session_state_timeout"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslClientSessionStateTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_type"
		if _, ok := i["ssl-client-session-state-type"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslClientSessionStateType(i["ssl-client-session-state-type"], d, pre_append)
			tmp["ssl_client_session_state_type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslClientSessionStateType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslDhBits(i["ssl-dh-bits"], d, pre_append)
			tmp["ssl_dh_bits"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslDhBits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp"
		if _, ok := i["ssl-hpkp"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHpkp(i["ssl-hpkp"], d, pre_append)
			tmp["ssl_hpkp"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHpkp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_age"
		if _, ok := i["ssl-hpkp-age"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHpkpAge(i["ssl-hpkp-age"], d, pre_append)
			tmp["ssl_hpkp_age"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHpkpAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_backup"
		if _, ok := i["ssl-hpkp-backup"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHpkpBackup(i["ssl-hpkp-backup"], d, pre_append)
			tmp["ssl_hpkp_backup"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHpkpBackup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_include_subdomains"
		if _, ok := i["ssl-hpkp-include-subdomains"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHpkpIncludeSubdomains(i["ssl-hpkp-include-subdomains"], d, pre_append)
			tmp["ssl_hpkp_include_subdomains"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHpkpIncludeSubdomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_primary"
		if _, ok := i["ssl-hpkp-primary"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHpkpPrimary(i["ssl-hpkp-primary"], d, pre_append)
			tmp["ssl_hpkp_primary"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHpkpPrimary")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_report_uri"
		if _, ok := i["ssl-hpkp-report-uri"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHpkpReportUri(i["ssl-hpkp-report-uri"], d, pre_append)
			tmp["ssl_hpkp_report_uri"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHpkpReportUri")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts"
		if _, ok := i["ssl-hsts"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHsts(i["ssl-hsts"], d, pre_append)
			tmp["ssl_hsts"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHsts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_age"
		if _, ok := i["ssl-hsts-age"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHstsAge(i["ssl-hsts-age"], d, pre_append)
			tmp["ssl_hsts_age"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHstsAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_include_subdomains"
		if _, ok := i["ssl-hsts-include-subdomains"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHstsIncludeSubdomains(i["ssl-hsts-include-subdomains"], d, pre_append)
			tmp["ssl_hsts_include_subdomains"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHstsIncludeSubdomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_location_conversion"
		if _, ok := i["ssl-http-location-conversion"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHttpLocationConversion(i["ssl-http-location-conversion"], d, pre_append)
			tmp["ssl_http_location_conversion"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHttpLocationConversion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_match_host"
		if _, ok := i["ssl-http-match-host"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslHttpMatchHost(i["ssl-http-match-host"], d, pre_append)
			tmp["ssl_http_match_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslHttpMatchHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslMinVersion(i["ssl-min-version"], d, pre_append)
			tmp["ssl_min_version"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_mode"
		if _, ok := i["ssl-mode"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslMode(i["ssl-mode"], d, pre_append)
			tmp["ssl_mode"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_pfs"
		if _, ok := i["ssl-pfs"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslPfs(i["ssl-pfs"], d, pre_append)
			tmp["ssl_pfs"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslPfs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_send_empty_frags"
		if _, ok := i["ssl-send-empty-frags"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslSendEmptyFrags(i["ssl-send-empty-frags"], d, pre_append)
			tmp["ssl_send_empty_frags"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslSendEmptyFrags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_algorithm"
		if _, ok := i["ssl-server-algorithm"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslServerAlgorithm(i["ssl-server-algorithm"], d, pre_append)
			tmp["ssl_server_algorithm"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslServerAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_max_version"
		if _, ok := i["ssl-server-max-version"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslServerMaxVersion(i["ssl-server-max-version"], d, pre_append)
			tmp["ssl_server_max_version"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslServerMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_min_version"
		if _, ok := i["ssl-server-min-version"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslServerMinVersion(i["ssl-server-min-version"], d, pre_append)
			tmp["ssl_server_min_version"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslServerMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_renegotiation"
		if _, ok := i["ssl-server-renegotiation"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslServerRenegotiation(i["ssl-server-renegotiation"], d, pre_append)
			tmp["ssl_server_renegotiation"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslServerRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_max"
		if _, ok := i["ssl-server-session-state-max"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslServerSessionStateMax(i["ssl-server-session-state-max"], d, pre_append)
			tmp["ssl_server_session_state_max"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslServerSessionStateMax")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_timeout"
		if _, ok := i["ssl-server-session-state-timeout"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslServerSessionStateTimeout(i["ssl-server-session-state-timeout"], d, pre_append)
			tmp["ssl_server_session_state_timeout"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslServerSessionStateTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_type"
		if _, ok := i["ssl-server-session-state-type"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslServerSessionStateType(i["ssl-server-session-state-type"], d, pre_append)
			tmp["ssl_server_session_state_type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-SslServerSessionStateType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFirewallVipDynamicMappingStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallVipDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenObjectFirewallVipDynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-Uuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weblogic_server"
		if _, ok := i["weblogic-server"]; ok {
			v := flattenObjectFirewallVipDynamicMappingWeblogicServer(i["weblogic-server"], d, pre_append)
			tmp["weblogic_server"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-WeblogicServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "websphere_server"
		if _, ok := i["websphere-server"]; ok {
			v := flattenObjectFirewallVipDynamicMappingWebsphereServer(i["websphere-server"], d, pre_append)
			tmp["websphere_server"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-DynamicMapping-WebsphereServer")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallVipDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallVipDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFirewallVipDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallVipDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingAddNat46Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingDnsMappingTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingExtaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingExtintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingExtip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFirewallVipDynamicMappingExtport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingGratuitousArpInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpIpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpIpHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpMultiplexMaxRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpMultiplexTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpSupportedMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingIpv6Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingIpv6Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingMappedAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingMappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipDynamicMappingMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingMaxEmbryonicConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingNatSourceVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingNat44(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingOutlookWebAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingPortforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingPortmappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversClientIp(i["client-ip"], d, pre_append)
			tmp["client_ip"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-ClientIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversHealthCheckProto(i["health-check-proto"], d, pre_append)
			tmp["health_check_proto"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-HealthCheckProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversHealthcheck(i["healthcheck"], d, pre_append)
			tmp["healthcheck"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-Healthcheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversSeq(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-Seq")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectFirewallVipDynamicMappingRealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-Realservers-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallVipDynamicMappingRealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipDynamicMappingRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSrcFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipDynamicMappingSrcintfFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipDynamicMappingSslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslCipherSuitesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-SslCipherSuites-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenObjectFirewallVipDynamicMappingSslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ObjectFirewallVipDynamicMapping-SslCipherSuites-Versions")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallVipDynamicMappingSslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslCipherSuitesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipDynamicMappingSslClientFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHpkp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHpkpAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHpkpBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHsts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHstsAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslPfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslServerMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslServerRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingSslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingWeblogicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipDynamicMappingWebsphereServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipExtaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipExtintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipExtip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipExtport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipGratuitousArpInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpIpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpIpHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpMultiplexMaxRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpMultiplexTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpSupportedMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipIpv6Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipIpv6Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipMappedAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipMappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipMaxEmbryonicConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipNatSourceVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipNat44(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipOutlookWebAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipPortforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipPortmappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenObjectFirewallVipRealserversAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {
			v := flattenObjectFirewallVipRealserversClientIp(i["client-ip"], d, pre_append)
			tmp["client_ip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-ClientIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {
			v := flattenObjectFirewallVipRealserversHealthcheck(i["healthcheck"], d, pre_append)
			tmp["healthcheck"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-Healthcheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenObjectFirewallVipRealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenObjectFirewallVipRealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallVipRealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFirewallVipRealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenObjectFirewallVipRealserversMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenObjectFirewallVipRealserversMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectFirewallVipRealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenObjectFirewallVipRealserversSeq(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-Seq")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFirewallVipRealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenObjectFirewallVipRealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallVipRealserversType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectFirewallVipRealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-Realservers-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallVipRealserversAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipRealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipSrcFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipSrcintfFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipSslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {
			v := flattenObjectFirewallVipSslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFirewallVipSslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallVipSslCipherSuitesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-SslCipherSuites-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenObjectFirewallVipSslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-SslCipherSuites-Versions")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallVipSslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslCipherSuitesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipSslClientFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHpkp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHpkpAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHpkpBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHsts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHstsAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslPfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslServerCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {
			v := flattenObjectFirewallVipSslServerCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-SslServerCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFirewallVipSslServerCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-SslServerCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenObjectFirewallVipSslServerCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ObjectFirewallVip-SslServerCipherSuites-Versions")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallVipSslServerCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslServerCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslServerCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVipSslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslServerMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslServerRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipSslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipWeblogicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipWebsphereServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallVip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("add_nat46_route", flattenObjectFirewallVipAddNat46Route(o["add-nat46-route"], d, "add_nat46_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-nat46-route"], "ObjectFirewallVip-AddNat46Route"); ok {
			if err = d.Set("add_nat46_route", vv); err != nil {
				return fmt.Errorf("Error reading add_nat46_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_nat46_route: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenObjectFirewallVipArpReply(o["arp-reply"], d, "arp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-reply"], "ObjectFirewallVip-ArpReply"); ok {
			if err = d.Set("arp_reply", vv); err != nil {
				return fmt.Errorf("Error reading arp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("color", flattenObjectFirewallVipColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallVip-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallVipComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallVip-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dns_mapping_ttl", flattenObjectFirewallVipDnsMappingTtl(o["dns-mapping-ttl"], d, "dns_mapping_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-mapping-ttl"], "ObjectFirewallVip-DnsMappingTtl"); ok {
			if err = d.Set("dns_mapping_ttl", vv); err != nil {
				return fmt.Errorf("Error reading dns_mapping_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_mapping_ttl: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectFirewallVipDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallVip-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectFirewallVipDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallVip-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("extaddr", flattenObjectFirewallVipExtaddr(o["extaddr"], d, "extaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["extaddr"], "ObjectFirewallVip-Extaddr"); ok {
			if err = d.Set("extaddr", vv); err != nil {
				return fmt.Errorf("Error reading extaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extaddr: %v", err)
		}
	}

	if err = d.Set("extintf", flattenObjectFirewallVipExtintf(o["extintf"], d, "extintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["extintf"], "ObjectFirewallVip-Extintf"); ok {
			if err = d.Set("extintf", vv); err != nil {
				return fmt.Errorf("Error reading extintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extintf: %v", err)
		}
	}

	if err = d.Set("extip", flattenObjectFirewallVipExtip(o["extip"], d, "extip")); err != nil {
		if vv, ok := fortiAPIPatch(o["extip"], "ObjectFirewallVip-Extip"); ok {
			if err = d.Set("extip", vv); err != nil {
				return fmt.Errorf("Error reading extip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if err = d.Set("extport", flattenObjectFirewallVipExtport(o["extport"], d, "extport")); err != nil {
		if vv, ok := fortiAPIPatch(o["extport"], "ObjectFirewallVip-Extport"); ok {
			if err = d.Set("extport", vv); err != nil {
				return fmt.Errorf("Error reading extport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("gratuitous_arp_interval", flattenObjectFirewallVipGratuitousArpInterval(o["gratuitous-arp-interval"], d, "gratuitous_arp_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["gratuitous-arp-interval"], "ObjectFirewallVip-GratuitousArpInterval"); ok {
			if err = d.Set("gratuitous_arp_interval", vv); err != nil {
				return fmt.Errorf("Error reading gratuitous_arp_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gratuitous_arp_interval: %v", err)
		}
	}

	if err = d.Set("http_cookie_age", flattenObjectFirewallVipHttpCookieAge(o["http-cookie-age"], d, "http_cookie_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-age"], "ObjectFirewallVip-HttpCookieAge"); ok {
			if err = d.Set("http_cookie_age", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_age: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain", flattenObjectFirewallVipHttpCookieDomain(o["http-cookie-domain"], d, "http_cookie_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain"], "ObjectFirewallVip-HttpCookieDomain"); ok {
			if err = d.Set("http_cookie_domain", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain_from_host", flattenObjectFirewallVipHttpCookieDomainFromHost(o["http-cookie-domain-from-host"], d, "http_cookie_domain_from_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain-from-host"], "ObjectFirewallVip-HttpCookieDomainFromHost"); ok {
			if err = d.Set("http_cookie_domain_from_host", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
		}
	}

	if err = d.Set("http_cookie_generation", flattenObjectFirewallVipHttpCookieGeneration(o["http-cookie-generation"], d, "http_cookie_generation")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-generation"], "ObjectFirewallVip-HttpCookieGeneration"); ok {
			if err = d.Set("http_cookie_generation", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_generation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_generation: %v", err)
		}
	}

	if err = d.Set("http_cookie_path", flattenObjectFirewallVipHttpCookiePath(o["http-cookie-path"], d, "http_cookie_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-path"], "ObjectFirewallVip-HttpCookiePath"); ok {
			if err = d.Set("http_cookie_path", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_path: %v", err)
		}
	}

	if err = d.Set("http_cookie_share", flattenObjectFirewallVipHttpCookieShare(o["http-cookie-share"], d, "http_cookie_share")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-share"], "ObjectFirewallVip-HttpCookieShare"); ok {
			if err = d.Set("http_cookie_share", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_share: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_share: %v", err)
		}
	}

	if err = d.Set("http_ip_header", flattenObjectFirewallVipHttpIpHeader(o["http-ip-header"], d, "http_ip_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-ip-header"], "ObjectFirewallVip-HttpIpHeader"); ok {
			if err = d.Set("http_ip_header", vv); err != nil {
				return fmt.Errorf("Error reading http_ip_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_ip_header: %v", err)
		}
	}

	if err = d.Set("http_ip_header_name", flattenObjectFirewallVipHttpIpHeaderName(o["http-ip-header-name"], d, "http_ip_header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-ip-header-name"], "ObjectFirewallVip-HttpIpHeaderName"); ok {
			if err = d.Set("http_ip_header_name", vv); err != nil {
				return fmt.Errorf("Error reading http_ip_header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_ip_header_name: %v", err)
		}
	}

	if err = d.Set("http_multiplex", flattenObjectFirewallVipHttpMultiplex(o["http-multiplex"], d, "http_multiplex")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-multiplex"], "ObjectFirewallVip-HttpMultiplex"); ok {
			if err = d.Set("http_multiplex", vv); err != nil {
				return fmt.Errorf("Error reading http_multiplex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_multiplex: %v", err)
		}
	}

	if err = d.Set("http_multiplex_max_request", flattenObjectFirewallVipHttpMultiplexMaxRequest(o["http-multiplex-max-request"], d, "http_multiplex_max_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-multiplex-max-request"], "ObjectFirewallVip-HttpMultiplexMaxRequest"); ok {
			if err = d.Set("http_multiplex_max_request", vv); err != nil {
				return fmt.Errorf("Error reading http_multiplex_max_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_multiplex_max_request: %v", err)
		}
	}

	if err = d.Set("http_multiplex_ttl", flattenObjectFirewallVipHttpMultiplexTtl(o["http-multiplex-ttl"], d, "http_multiplex_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-multiplex-ttl"], "ObjectFirewallVip-HttpMultiplexTtl"); ok {
			if err = d.Set("http_multiplex_ttl", vv); err != nil {
				return fmt.Errorf("Error reading http_multiplex_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_multiplex_ttl: %v", err)
		}
	}

	if err = d.Set("http_redirect", flattenObjectFirewallVipHttpRedirect(o["http-redirect"], d, "http_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-redirect"], "ObjectFirewallVip-HttpRedirect"); ok {
			if err = d.Set("http_redirect", vv); err != nil {
				return fmt.Errorf("Error reading http_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_redirect: %v", err)
		}
	}

	if err = d.Set("http_supported_max_version", flattenObjectFirewallVipHttpSupportedMaxVersion(o["http-supported-max-version"], d, "http_supported_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-supported-max-version"], "ObjectFirewallVip-HttpSupportedMaxVersion"); ok {
			if err = d.Set("http_supported_max_version", vv); err != nil {
				return fmt.Errorf("Error reading http_supported_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_supported_max_version: %v", err)
		}
	}

	if err = d.Set("https_cookie_secure", flattenObjectFirewallVipHttpsCookieSecure(o["https-cookie-secure"], d, "https_cookie_secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-cookie-secure"], "ObjectFirewallVip-HttpsCookieSecure"); ok {
			if err = d.Set("https_cookie_secure", vv); err != nil {
				return fmt.Errorf("Error reading https_cookie_secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_cookie_secure: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallVipId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallVip-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ipv6_mappedip", flattenObjectFirewallVipIpv6Mappedip(o["ipv6-mappedip"], d, "ipv6_mappedip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-mappedip"], "ObjectFirewallVip-Ipv6Mappedip"); ok {
			if err = d.Set("ipv6_mappedip", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_mappedip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_mappedip: %v", err)
		}
	}

	if err = d.Set("ipv6_mappedport", flattenObjectFirewallVipIpv6Mappedport(o["ipv6-mappedport"], d, "ipv6_mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-mappedport"], "ObjectFirewallVip-Ipv6Mappedport"); ok {
			if err = d.Set("ipv6_mappedport", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_mappedport: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenObjectFirewallVipLdbMethod(o["ldb-method"], d, "ldb_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldb-method"], "ObjectFirewallVip-LdbMethod"); ok {
			if err = d.Set("ldb_method", vv); err != nil {
				return fmt.Errorf("Error reading ldb_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("mapped_addr", flattenObjectFirewallVipMappedAddr(o["mapped-addr"], d, "mapped_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["mapped-addr"], "ObjectFirewallVip-MappedAddr"); ok {
			if err = d.Set("mapped_addr", vv); err != nil {
				return fmt.Errorf("Error reading mapped_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mapped_addr: %v", err)
		}
	}

	if err = d.Set("mappedip", flattenObjectFirewallVipMappedip(o["mappedip"], d, "mappedip")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedip"], "ObjectFirewallVip-Mappedip"); ok {
			if err = d.Set("mappedip", vv); err != nil {
				return fmt.Errorf("Error reading mappedip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedip: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenObjectFirewallVipMappedport(o["mappedport"], d, "mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedport"], "ObjectFirewallVip-Mappedport"); ok {
			if err = d.Set("mappedport", vv); err != nil {
				return fmt.Errorf("Error reading mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("max_embryonic_connections", flattenObjectFirewallVipMaxEmbryonicConnections(o["max-embryonic-connections"], d, "max_embryonic_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-embryonic-connections"], "ObjectFirewallVip-MaxEmbryonicConnections"); ok {
			if err = d.Set("max_embryonic_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_embryonic_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_embryonic_connections: %v", err)
		}
	}

	if err = d.Set("monitor", flattenObjectFirewallVipMonitor(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "ObjectFirewallVip-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallVipName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallVip-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat_source_vip", flattenObjectFirewallVipNatSourceVip(o["nat-source-vip"], d, "nat_source_vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-source-vip"], "ObjectFirewallVip-NatSourceVip"); ok {
			if err = d.Set("nat_source_vip", vv); err != nil {
				return fmt.Errorf("Error reading nat_source_vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_source_vip: %v", err)
		}
	}

	if err = d.Set("nat44", flattenObjectFirewallVipNat44(o["nat44"], d, "nat44")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat44"], "ObjectFirewallVip-Nat44"); ok {
			if err = d.Set("nat44", vv); err != nil {
				return fmt.Errorf("Error reading nat44: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat44: %v", err)
		}
	}

	if err = d.Set("nat46", flattenObjectFirewallVipNat46(o["nat46"], d, "nat46")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46"], "ObjectFirewallVip-Nat46"); ok {
			if err = d.Set("nat46", vv); err != nil {
				return fmt.Errorf("Error reading nat46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("outlook_web_access", flattenObjectFirewallVipOutlookWebAccess(o["outlook-web-access"], d, "outlook_web_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["outlook-web-access"], "ObjectFirewallVip-OutlookWebAccess"); ok {
			if err = d.Set("outlook_web_access", vv); err != nil {
				return fmt.Errorf("Error reading outlook_web_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outlook_web_access: %v", err)
		}
	}

	if err = d.Set("persistence", flattenObjectFirewallVipPersistence(o["persistence"], d, "persistence")); err != nil {
		if vv, ok := fortiAPIPatch(o["persistence"], "ObjectFirewallVip-Persistence"); ok {
			if err = d.Set("persistence", vv); err != nil {
				return fmt.Errorf("Error reading persistence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("portforward", flattenObjectFirewallVipPortforward(o["portforward"], d, "portforward")); err != nil {
		if vv, ok := fortiAPIPatch(o["portforward"], "ObjectFirewallVip-Portforward"); ok {
			if err = d.Set("portforward", vv); err != nil {
				return fmt.Errorf("Error reading portforward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portforward: %v", err)
		}
	}

	if err = d.Set("portmapping_type", flattenObjectFirewallVipPortmappingType(o["portmapping-type"], d, "portmapping_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["portmapping-type"], "ObjectFirewallVip-PortmappingType"); ok {
			if err = d.Set("portmapping_type", vv); err != nil {
				return fmt.Errorf("Error reading portmapping_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portmapping_type: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectFirewallVipProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectFirewallVip-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("realservers", flattenObjectFirewallVipRealservers(o["realservers"], d, "realservers")); err != nil {
			if vv, ok := fortiAPIPatch(o["realservers"], "ObjectFirewallVip-Realservers"); ok {
				if err = d.Set("realservers", vv); err != nil {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenObjectFirewallVipRealservers(o["realservers"], d, "realservers")); err != nil {
				if vv, ok := fortiAPIPatch(o["realservers"], "ObjectFirewallVip-Realservers"); ok {
					if err = d.Set("realservers", vv); err != nil {
						return fmt.Errorf("Error reading realservers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenObjectFirewallVipServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "ObjectFirewallVip-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("service", flattenObjectFirewallVipService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "ObjectFirewallVip-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("src_filter", flattenObjectFirewallVipSrcFilter(o["src-filter"], d, "src_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-filter"], "ObjectFirewallVip-SrcFilter"); ok {
			if err = d.Set("src_filter", vv); err != nil {
				return fmt.Errorf("Error reading src_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_filter: %v", err)
		}
	}

	if err = d.Set("srcintf_filter", flattenObjectFirewallVipSrcintfFilter(o["srcintf-filter"], d, "srcintf_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf-filter"], "ObjectFirewallVip-SrcintfFilter"); ok {
			if err = d.Set("srcintf_filter", vv); err != nil {
				return fmt.Errorf("Error reading srcintf_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf_filter: %v", err)
		}
	}

	if err = d.Set("ssl_accept_ffdhe_groups", flattenObjectFirewallVipSslAcceptFfdheGroups(o["ssl-accept-ffdhe-groups"], d, "ssl_accept_ffdhe_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-accept-ffdhe-groups"], "ObjectFirewallVip-SslAcceptFfdheGroups"); ok {
			if err = d.Set("ssl_accept_ffdhe_groups", vv); err != nil {
				return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenObjectFirewallVipSslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "ObjectFirewallVip-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenObjectFirewallVipSslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "ObjectFirewallVip-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenObjectFirewallVipSslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "ObjectFirewallVip-SslCipherSuites"); ok {
				if err = d.Set("ssl_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenObjectFirewallVipSslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "ObjectFirewallVip-SslCipherSuites"); ok {
					if err = d.Set("ssl_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_client_fallback", flattenObjectFirewallVipSslClientFallback(o["ssl-client-fallback"], d, "ssl_client_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-fallback"], "ObjectFirewallVip-SslClientFallback"); ok {
			if err = d.Set("ssl_client_fallback", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
		}
	}

	if err = d.Set("ssl_client_rekey_count", flattenObjectFirewallVipSslClientRekeyCount(o["ssl-client-rekey-count"], d, "ssl_client_rekey_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-rekey-count"], "ObjectFirewallVip-SslClientRekeyCount"); ok {
			if err = d.Set("ssl_client_rekey_count", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenObjectFirewallVipSslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-renegotiation"], "ObjectFirewallVip-SslClientRenegotiation"); ok {
			if err = d.Set("ssl_client_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_max", flattenObjectFirewallVipSslClientSessionStateMax(o["ssl-client-session-state-max"], d, "ssl_client_session_state_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-max"], "ObjectFirewallVip-SslClientSessionStateMax"); ok {
			if err = d.Set("ssl_client_session_state_max", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_timeout", flattenObjectFirewallVipSslClientSessionStateTimeout(o["ssl-client-session-state-timeout"], d, "ssl_client_session_state_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-timeout"], "ObjectFirewallVip-SslClientSessionStateTimeout"); ok {
			if err = d.Set("ssl_client_session_state_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_type", flattenObjectFirewallVipSslClientSessionStateType(o["ssl-client-session-state-type"], d, "ssl_client_session_state_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-type"], "ObjectFirewallVip-SslClientSessionStateType"); ok {
			if err = d.Set("ssl_client_session_state_type", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenObjectFirewallVipSslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "ObjectFirewallVip-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp", flattenObjectFirewallVipSslHpkp(o["ssl-hpkp"], d, "ssl_hpkp")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp"], "ObjectFirewallVip-SslHpkp"); ok {
			if err = d.Set("ssl_hpkp", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_age", flattenObjectFirewallVipSslHpkpAge(o["ssl-hpkp-age"], d, "ssl_hpkp_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-age"], "ObjectFirewallVip-SslHpkpAge"); ok {
			if err = d.Set("ssl_hpkp_age", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_backup", flattenObjectFirewallVipSslHpkpBackup(o["ssl-hpkp-backup"], d, "ssl_hpkp_backup")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-backup"], "ObjectFirewallVip-SslHpkpBackup"); ok {
			if err = d.Set("ssl_hpkp_backup", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_include_subdomains", flattenObjectFirewallVipSslHpkpIncludeSubdomains(o["ssl-hpkp-include-subdomains"], d, "ssl_hpkp_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-include-subdomains"], "ObjectFirewallVip-SslHpkpIncludeSubdomains"); ok {
			if err = d.Set("ssl_hpkp_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_primary", flattenObjectFirewallVipSslHpkpPrimary(o["ssl-hpkp-primary"], d, "ssl_hpkp_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-primary"], "ObjectFirewallVip-SslHpkpPrimary"); ok {
			if err = d.Set("ssl_hpkp_primary", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_report_uri", flattenObjectFirewallVipSslHpkpReportUri(o["ssl-hpkp-report-uri"], d, "ssl_hpkp_report_uri")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-report-uri"], "ObjectFirewallVip-SslHpkpReportUri"); ok {
			if err = d.Set("ssl_hpkp_report_uri", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
		}
	}

	if err = d.Set("ssl_hsts", flattenObjectFirewallVipSslHsts(o["ssl-hsts"], d, "ssl_hsts")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts"], "ObjectFirewallVip-SslHsts"); ok {
			if err = d.Set("ssl_hsts", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_age", flattenObjectFirewallVipSslHstsAge(o["ssl-hsts-age"], d, "ssl_hsts_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts-age"], "ObjectFirewallVip-SslHstsAge"); ok {
			if err = d.Set("ssl_hsts_age", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_include_subdomains", flattenObjectFirewallVipSslHstsIncludeSubdomains(o["ssl-hsts-include-subdomains"], d, "ssl_hsts_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts-include-subdomains"], "ObjectFirewallVip-SslHstsIncludeSubdomains"); ok {
			if err = d.Set("ssl_hsts_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_http_location_conversion", flattenObjectFirewallVipSslHttpLocationConversion(o["ssl-http-location-conversion"], d, "ssl_http_location_conversion")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-http-location-conversion"], "ObjectFirewallVip-SslHttpLocationConversion"); ok {
			if err = d.Set("ssl_http_location_conversion", vv); err != nil {
				return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
		}
	}

	if err = d.Set("ssl_http_match_host", flattenObjectFirewallVipSslHttpMatchHost(o["ssl-http-match-host"], d, "ssl_http_match_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-http-match-host"], "ObjectFirewallVip-SslHttpMatchHost"); ok {
			if err = d.Set("ssl_http_match_host", vv); err != nil {
				return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenObjectFirewallVipSslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "ObjectFirewallVip-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenObjectFirewallVipSslMinVersion(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "ObjectFirewallVip-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_mode", flattenObjectFirewallVipSslMode(o["ssl-mode"], d, "ssl_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mode"], "ObjectFirewallVip-SslMode"); ok {
			if err = d.Set("ssl_mode", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if err = d.Set("ssl_pfs", flattenObjectFirewallVipSslPfs(o["ssl-pfs"], d, "ssl_pfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-pfs"], "ObjectFirewallVip-SslPfs"); ok {
			if err = d.Set("ssl_pfs", vv); err != nil {
				return fmt.Errorf("Error reading ssl_pfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_pfs: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenObjectFirewallVipSslSendEmptyFrags(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-send-empty-frags"], "ObjectFirewallVip-SslSendEmptyFrags"); ok {
			if err = d.Set("ssl_send_empty_frags", vv); err != nil {
				return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("ssl_server_algorithm", flattenObjectFirewallVipSslServerAlgorithm(o["ssl-server-algorithm"], d, "ssl_server_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-algorithm"], "ObjectFirewallVip-SslServerAlgorithm"); ok {
			if err = d.Set("ssl_server_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_server_cipher_suites", flattenObjectFirewallVipSslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-server-cipher-suites"], "ObjectFirewallVip-SslServerCipherSuites"); ok {
				if err = d.Set("ssl_server_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server_cipher_suites"); ok {
			if err = d.Set("ssl_server_cipher_suites", flattenObjectFirewallVipSslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-server-cipher-suites"], "ObjectFirewallVip-SslServerCipherSuites"); ok {
					if err = d.Set("ssl_server_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_server_max_version", flattenObjectFirewallVipSslServerMaxVersion(o["ssl-server-max-version"], d, "ssl_server_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-max-version"], "ObjectFirewallVip-SslServerMaxVersion"); ok {
			if err = d.Set("ssl_server_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_min_version", flattenObjectFirewallVipSslServerMinVersion(o["ssl-server-min-version"], d, "ssl_server_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-min-version"], "ObjectFirewallVip-SslServerMinVersion"); ok {
			if err = d.Set("ssl_server_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_renegotiation", flattenObjectFirewallVipSslServerRenegotiation(o["ssl-server-renegotiation"], d, "ssl_server_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-renegotiation"], "ObjectFirewallVip-SslServerRenegotiation"); ok {
			if err = d.Set("ssl_server_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_max", flattenObjectFirewallVipSslServerSessionStateMax(o["ssl-server-session-state-max"], d, "ssl_server_session_state_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-max"], "ObjectFirewallVip-SslServerSessionStateMax"); ok {
			if err = d.Set("ssl_server_session_state_max", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_timeout", flattenObjectFirewallVipSslServerSessionStateTimeout(o["ssl-server-session-state-timeout"], d, "ssl_server_session_state_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-timeout"], "ObjectFirewallVip-SslServerSessionStateTimeout"); ok {
			if err = d.Set("ssl_server_session_state_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_type", flattenObjectFirewallVipSslServerSessionStateType(o["ssl-server-session-state-type"], d, "ssl_server_session_state_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-type"], "ObjectFirewallVip-SslServerSessionStateType"); ok {
			if err = d.Set("ssl_server_session_state_type", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFirewallVipStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFirewallVip-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallVipType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallVip-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectFirewallVipUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectFirewallVip-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("weblogic_server", flattenObjectFirewallVipWeblogicServer(o["weblogic-server"], d, "weblogic_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["weblogic-server"], "ObjectFirewallVip-WeblogicServer"); ok {
			if err = d.Set("weblogic_server", vv); err != nil {
				return fmt.Errorf("Error reading weblogic_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weblogic_server: %v", err)
		}
	}

	if err = d.Set("websphere_server", flattenObjectFirewallVipWebsphereServer(o["websphere-server"], d, "websphere_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["websphere-server"], "ObjectFirewallVip-WebsphereServer"); ok {
			if err = d.Set("websphere_server", vv); err != nil {
				return fmt.Errorf("Error reading websphere_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading websphere_server: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallVipFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallVipAddNat46Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDnsMappingTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallVipDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat46_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-nat46-route"], _ = expandObjectFirewallVipDynamicMappingAddNat46Route(d, i["add_nat46_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-reply"], _ = expandObjectFirewallVipDynamicMappingArpReply(d, i["arp_reply"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color"], _ = expandObjectFirewallVipDynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectFirewallVipDynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_mapping_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-mapping-ttl"], _ = expandObjectFirewallVipDynamicMappingDnsMappingTtl(d, i["dns_mapping_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extaddr"], _ = expandObjectFirewallVipDynamicMappingExtaddr(d, i["extaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extintf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extintf"], _ = expandObjectFirewallVipDynamicMappingExtintf(d, i["extintf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extip"], _ = expandObjectFirewallVipDynamicMappingExtip(d, i["extip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extport"], _ = expandObjectFirewallVipDynamicMappingExtport(d, i["extport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gratuitous_arp_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gratuitous-arp-interval"], _ = expandObjectFirewallVipDynamicMappingGratuitousArpInterval(d, i["gratuitous_arp_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-age"], _ = expandObjectFirewallVipDynamicMappingHttpCookieAge(d, i["http_cookie_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain"], _ = expandObjectFirewallVipDynamicMappingHttpCookieDomain(d, i["http_cookie_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain-from-host"], _ = expandObjectFirewallVipDynamicMappingHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-generation"], _ = expandObjectFirewallVipDynamicMappingHttpCookieGeneration(d, i["http_cookie_generation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-path"], _ = expandObjectFirewallVipDynamicMappingHttpCookiePath(d, i["http_cookie_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-share"], _ = expandObjectFirewallVipDynamicMappingHttpCookieShare(d, i["http_cookie_share"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-ip-header"], _ = expandObjectFirewallVipDynamicMappingHttpIpHeader(d, i["http_ip_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-ip-header-name"], _ = expandObjectFirewallVipDynamicMappingHttpIpHeaderName(d, i["http_ip_header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-multiplex"], _ = expandObjectFirewallVipDynamicMappingHttpMultiplex(d, i["http_multiplex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex_max_request"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-multiplex-max-request"], _ = expandObjectFirewallVipDynamicMappingHttpMultiplexMaxRequest(d, i["http_multiplex_max_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-multiplex-ttl"], _ = expandObjectFirewallVipDynamicMappingHttpMultiplexTtl(d, i["http_multiplex_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_redirect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-redirect"], _ = expandObjectFirewallVipDynamicMappingHttpRedirect(d, i["http_redirect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_supported_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-supported-max-version"], _ = expandObjectFirewallVipDynamicMappingHttpSupportedMaxVersion(d, i["http_supported_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-cookie-secure"], _ = expandObjectFirewallVipDynamicMappingHttpsCookieSecure(d, i["https_cookie_secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallVipDynamicMappingId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_mappedip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-mappedip"], _ = expandObjectFirewallVipDynamicMappingIpv6Mappedip(d, i["ipv6_mappedip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-mappedport"], _ = expandObjectFirewallVipDynamicMappingIpv6Mappedport(d, i["ipv6_mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandObjectFirewallVipDynamicMappingLdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mapped_addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mapped-addr"], _ = expandObjectFirewallVipDynamicMappingMappedAddr(d, i["mapped_addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedip"], _ = expandObjectFirewallVipDynamicMappingMappedip(d, i["mappedip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandObjectFirewallVipDynamicMappingMappedport(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_embryonic_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-embryonic-connections"], _ = expandObjectFirewallVipDynamicMappingMaxEmbryonicConnections(d, i["max_embryonic_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandObjectFirewallVipDynamicMappingMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat_source_vip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat-source-vip"], _ = expandObjectFirewallVipDynamicMappingNatSourceVip(d, i["nat_source_vip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat44"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat44"], _ = expandObjectFirewallVipDynamicMappingNat44(d, i["nat44"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat46"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat46"], _ = expandObjectFirewallVipDynamicMappingNat46(d, i["nat46"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "outlook_web_access"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["outlook-web-access"], _ = expandObjectFirewallVipDynamicMappingOutlookWebAccess(d, i["outlook_web_access"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["persistence"], _ = expandObjectFirewallVipDynamicMappingPersistence(d, i["persistence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portforward"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portforward"], _ = expandObjectFirewallVipDynamicMappingPortforward(d, i["portforward"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portmapping_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portmapping-type"], _ = expandObjectFirewallVipDynamicMappingPortmappingType(d, i["portmapping_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectFirewallVipDynamicMappingProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallVipDynamicMappingRealservers(d, i["realservers"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["realservers"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-type"], _ = expandObjectFirewallVipDynamicMappingServerType(d, i["server_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandObjectFirewallVipDynamicMappingService(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-filter"], _ = expandObjectFirewallVipDynamicMappingSrcFilter(d, i["src_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcintf-filter"], _ = expandObjectFirewallVipDynamicMappingSrcintfFilter(d, i["srcintf_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_accept_ffdhe_groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-accept-ffdhe-groups"], _ = expandObjectFirewallVipDynamicMappingSslAcceptFfdheGroups(d, i["ssl_accept_ffdhe_groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-algorithm"], _ = expandObjectFirewallVipDynamicMappingSslAlgorithm(d, i["ssl_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-certificate"], _ = expandObjectFirewallVipDynamicMappingSslCertificate(d, i["ssl_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallVipDynamicMappingSslCipherSuites(d, i["ssl_cipher_suites"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ssl-cipher-suites"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_fallback"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-fallback"], _ = expandObjectFirewallVipDynamicMappingSslClientFallback(d, i["ssl_client_fallback"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_rekey_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-rekey-count"], _ = expandObjectFirewallVipDynamicMappingSslClientRekeyCount(d, i["ssl_client_rekey_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-renegotiation"], _ = expandObjectFirewallVipDynamicMappingSslClientRenegotiation(d, i["ssl_client_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-max"], _ = expandObjectFirewallVipDynamicMappingSslClientSessionStateMax(d, i["ssl_client_session_state_max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-timeout"], _ = expandObjectFirewallVipDynamicMappingSslClientSessionStateTimeout(d, i["ssl_client_session_state_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-type"], _ = expandObjectFirewallVipDynamicMappingSslClientSessionStateType(d, i["ssl_client_session_state_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-dh-bits"], _ = expandObjectFirewallVipDynamicMappingSslDhBits(d, i["ssl_dh_bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp"], _ = expandObjectFirewallVipDynamicMappingSslHpkp(d, i["ssl_hpkp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-age"], _ = expandObjectFirewallVipDynamicMappingSslHpkpAge(d, i["ssl_hpkp_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_backup"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-backup"], _ = expandObjectFirewallVipDynamicMappingSslHpkpBackup(d, i["ssl_hpkp_backup"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_include_subdomains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-include-subdomains"], _ = expandObjectFirewallVipDynamicMappingSslHpkpIncludeSubdomains(d, i["ssl_hpkp_include_subdomains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_primary"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-primary"], _ = expandObjectFirewallVipDynamicMappingSslHpkpPrimary(d, i["ssl_hpkp_primary"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_report_uri"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-report-uri"], _ = expandObjectFirewallVipDynamicMappingSslHpkpReportUri(d, i["ssl_hpkp_report_uri"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts"], _ = expandObjectFirewallVipDynamicMappingSslHsts(d, i["ssl_hsts"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts-age"], _ = expandObjectFirewallVipDynamicMappingSslHstsAge(d, i["ssl_hsts_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_include_subdomains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts-include-subdomains"], _ = expandObjectFirewallVipDynamicMappingSslHstsIncludeSubdomains(d, i["ssl_hsts_include_subdomains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_location_conversion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-http-location-conversion"], _ = expandObjectFirewallVipDynamicMappingSslHttpLocationConversion(d, i["ssl_http_location_conversion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_match_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-http-match-host"], _ = expandObjectFirewallVipDynamicMappingSslHttpMatchHost(d, i["ssl_http_match_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandObjectFirewallVipDynamicMappingSslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-version"], _ = expandObjectFirewallVipDynamicMappingSslMinVersion(d, i["ssl_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-mode"], _ = expandObjectFirewallVipDynamicMappingSslMode(d, i["ssl_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_pfs"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-pfs"], _ = expandObjectFirewallVipDynamicMappingSslPfs(d, i["ssl_pfs"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_send_empty_frags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-send-empty-frags"], _ = expandObjectFirewallVipDynamicMappingSslSendEmptyFrags(d, i["ssl_send_empty_frags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-algorithm"], _ = expandObjectFirewallVipDynamicMappingSslServerAlgorithm(d, i["ssl_server_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-max-version"], _ = expandObjectFirewallVipDynamicMappingSslServerMaxVersion(d, i["ssl_server_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-min-version"], _ = expandObjectFirewallVipDynamicMappingSslServerMinVersion(d, i["ssl_server_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-renegotiation"], _ = expandObjectFirewallVipDynamicMappingSslServerRenegotiation(d, i["ssl_server_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-max"], _ = expandObjectFirewallVipDynamicMappingSslServerSessionStateMax(d, i["ssl_server_session_state_max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-timeout"], _ = expandObjectFirewallVipDynamicMappingSslServerSessionStateTimeout(d, i["ssl_server_session_state_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-type"], _ = expandObjectFirewallVipDynamicMappingSslServerSessionStateType(d, i["ssl_server_session_state_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFirewallVipDynamicMappingStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallVipDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandObjectFirewallVipDynamicMappingUuid(d, i["uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weblogic_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weblogic-server"], _ = expandObjectFirewallVipDynamicMappingWeblogicServer(d, i["weblogic_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "websphere_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["websphere-server"], _ = expandObjectFirewallVipDynamicMappingWebsphereServer(d, i["websphere_server"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVipDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallVipDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectFirewallVipDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVipDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingAddNat46Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingDnsMappingTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingExtaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingExtintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingExtip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingExtport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingGratuitousArpInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpIpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpIpHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpMultiplexMaxRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpMultiplexTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpSupportedMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingIpv6Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingIpv6Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingMappedAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingMappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVipDynamicMappingMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingMaxEmbryonicConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingNatSourceVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingNat44(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingOutlookWebAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingPersistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingPortforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingPortmappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandObjectFirewallVipDynamicMappingRealserversAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-ip"], _ = expandObjectFirewallVipDynamicMappingRealserversClientIp(d, i["client_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-proto"], _ = expandObjectFirewallVipDynamicMappingRealserversHealthCheckProto(d, i["health_check_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["healthcheck"], _ = expandObjectFirewallVipDynamicMappingRealserversHealthcheck(d, i["healthcheck"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandObjectFirewallVipDynamicMappingRealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandObjectFirewallVipDynamicMappingRealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallVipDynamicMappingRealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFirewallVipDynamicMappingRealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandObjectFirewallVipDynamicMappingRealserversMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandObjectFirewallVipDynamicMappingRealserversMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectFirewallVipDynamicMappingRealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandObjectFirewallVipDynamicMappingRealserversSeq(d, i["seq"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFirewallVipDynamicMappingRealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandObjectFirewallVipDynamicMappingRealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallVipDynamicMappingRealserversType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectFirewallVipDynamicMappingRealserversWeight(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVipDynamicMappingRealserversAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVipDynamicMappingRealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingRealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVipDynamicMappingSrcintfFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVipDynamicMappingSslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cipher"], _ = expandObjectFirewallVipDynamicMappingSslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallVipDynamicMappingSslCipherSuitesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFirewallVipDynamicMappingSslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandObjectFirewallVipDynamicMappingSslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVipDynamicMappingSslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslCipherSuitesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVipDynamicMappingSslClientFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHpkp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHpkpAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHpkpBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHsts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHstsAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslPfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslServerMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslServerRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingSslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingWeblogicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipDynamicMappingWebsphereServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipExtaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectFirewallVipExtintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipExtip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipExtport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipGratuitousArpInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpIpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpIpHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpMultiplexMaxRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpMultiplexTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpSupportedMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipIpv6Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipIpv6Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipMappedAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipMappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVipMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipMaxEmbryonicConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipNatSourceVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipNat44(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipOutlookWebAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipPersistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipPortforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipPortmappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandObjectFirewallVipRealserversAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-ip"], _ = expandObjectFirewallVipRealserversClientIp(d, i["client_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["healthcheck"], _ = expandObjectFirewallVipRealserversHealthcheck(d, i["healthcheck"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandObjectFirewallVipRealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandObjectFirewallVipRealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallVipRealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFirewallVipRealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandObjectFirewallVipRealserversMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandObjectFirewallVipRealserversMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectFirewallVipRealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandObjectFirewallVipRealserversSeq(d, i["seq"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFirewallVipRealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandObjectFirewallVipRealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallVipRealserversType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectFirewallVipRealserversWeight(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVipRealserversAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVipRealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipRealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectFirewallVipSrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVipSrcintfFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectFirewallVipSslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cipher"], _ = expandObjectFirewallVipSslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFirewallVipSslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallVipSslCipherSuitesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandObjectFirewallVipSslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVipSslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslCipherSuitesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVipSslClientFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHpkp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHpkpAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHpkpBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHsts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHstsAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslPfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslServerCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cipher"], _ = expandObjectFirewallVipSslServerCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFirewallVipSslServerCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandObjectFirewallVipSslServerCipherSuitesVersions(d, i["versions"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVipSslServerCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslServerCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslServerCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVipSslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslServerMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslServerRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipSslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipWeblogicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipWebsphereServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallVip(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_nat46_route"); ok || d.HasChange("add_nat46_route") {
		t, err := expandObjectFirewallVipAddNat46Route(d, v, "add_nat46_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat46-route"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok || d.HasChange("arp_reply") {
		t, err := expandObjectFirewallVipArpReply(d, v, "arp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectFirewallVipColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectFirewallVipComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dns_mapping_ttl"); ok || d.HasChange("dns_mapping_ttl") {
		t, err := expandObjectFirewallVipDnsMappingTtl(d, v, "dns_mapping_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mapping-ttl"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectFirewallVipDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("extaddr"); ok || d.HasChange("extaddr") {
		t, err := expandObjectFirewallVipExtaddr(d, v, "extaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extaddr"] = t
		}
	}

	if v, ok := d.GetOk("extintf"); ok || d.HasChange("extintf") {
		t, err := expandObjectFirewallVipExtintf(d, v, "extintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extintf"] = t
		}
	}

	if v, ok := d.GetOk("extip"); ok || d.HasChange("extip") {
		t, err := expandObjectFirewallVipExtip(d, v, "extip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extip"] = t
		}
	}

	if v, ok := d.GetOk("extport"); ok || d.HasChange("extport") {
		t, err := expandObjectFirewallVipExtport(d, v, "extport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extport"] = t
		}
	}

	if v, ok := d.GetOk("gratuitous_arp_interval"); ok || d.HasChange("gratuitous_arp_interval") {
		t, err := expandObjectFirewallVipGratuitousArpInterval(d, v, "gratuitous_arp_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gratuitous-arp-interval"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_age"); ok || d.HasChange("http_cookie_age") {
		t, err := expandObjectFirewallVipHttpCookieAge(d, v, "http_cookie_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain"); ok || d.HasChange("http_cookie_domain") {
		t, err := expandObjectFirewallVipHttpCookieDomain(d, v, "http_cookie_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain_from_host"); ok || d.HasChange("http_cookie_domain_from_host") {
		t, err := expandObjectFirewallVipHttpCookieDomainFromHost(d, v, "http_cookie_domain_from_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain-from-host"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_generation"); ok || d.HasChange("http_cookie_generation") {
		t, err := expandObjectFirewallVipHttpCookieGeneration(d, v, "http_cookie_generation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-generation"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_path"); ok || d.HasChange("http_cookie_path") {
		t, err := expandObjectFirewallVipHttpCookiePath(d, v, "http_cookie_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-path"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_share"); ok || d.HasChange("http_cookie_share") {
		t, err := expandObjectFirewallVipHttpCookieShare(d, v, "http_cookie_share")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-share"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header"); ok || d.HasChange("http_ip_header") {
		t, err := expandObjectFirewallVipHttpIpHeader(d, v, "http_ip_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header_name"); ok || d.HasChange("http_ip_header_name") {
		t, err := expandObjectFirewallVipHttpIpHeaderName(d, v, "http_ip_header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header-name"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex"); ok || d.HasChange("http_multiplex") {
		t, err := expandObjectFirewallVipHttpMultiplex(d, v, "http_multiplex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex_max_request"); ok || d.HasChange("http_multiplex_max_request") {
		t, err := expandObjectFirewallVipHttpMultiplexMaxRequest(d, v, "http_multiplex_max_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex-max-request"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex_ttl"); ok || d.HasChange("http_multiplex_ttl") {
		t, err := expandObjectFirewallVipHttpMultiplexTtl(d, v, "http_multiplex_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex-ttl"] = t
		}
	}

	if v, ok := d.GetOk("http_redirect"); ok || d.HasChange("http_redirect") {
		t, err := expandObjectFirewallVipHttpRedirect(d, v, "http_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-redirect"] = t
		}
	}

	if v, ok := d.GetOk("http_supported_max_version"); ok || d.HasChange("http_supported_max_version") {
		t, err := expandObjectFirewallVipHttpSupportedMaxVersion(d, v, "http_supported_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-supported-max-version"] = t
		}
	}

	if v, ok := d.GetOk("https_cookie_secure"); ok || d.HasChange("https_cookie_secure") {
		t, err := expandObjectFirewallVipHttpsCookieSecure(d, v, "https_cookie_secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-cookie-secure"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallVipId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_mappedip"); ok || d.HasChange("ipv6_mappedip") {
		t, err := expandObjectFirewallVipIpv6Mappedip(d, v, "ipv6_mappedip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-mappedip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_mappedport"); ok || d.HasChange("ipv6_mappedport") {
		t, err := expandObjectFirewallVipIpv6Mappedport(d, v, "ipv6_mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-mappedport"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok || d.HasChange("ldb_method") {
		t, err := expandObjectFirewallVipLdbMethod(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("mapped_addr"); ok || d.HasChange("mapped_addr") {
		t, err := expandObjectFirewallVipMappedAddr(d, v, "mapped_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapped-addr"] = t
		}
	}

	if v, ok := d.GetOk("mappedip"); ok || d.HasChange("mappedip") {
		t, err := expandObjectFirewallVipMappedip(d, v, "mappedip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedip"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok || d.HasChange("mappedport") {
		t, err := expandObjectFirewallVipMappedport(d, v, "mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("max_embryonic_connections"); ok || d.HasChange("max_embryonic_connections") {
		t, err := expandObjectFirewallVipMaxEmbryonicConnections(d, v, "max_embryonic_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-embryonic-connections"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandObjectFirewallVipMonitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallVipName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat_source_vip"); ok || d.HasChange("nat_source_vip") {
		t, err := expandObjectFirewallVipNatSourceVip(d, v, "nat_source_vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-source-vip"] = t
		}
	}

	if v, ok := d.GetOk("nat44"); ok || d.HasChange("nat44") {
		t, err := expandObjectFirewallVipNat44(d, v, "nat44")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat44"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok || d.HasChange("nat46") {
		t, err := expandObjectFirewallVipNat46(d, v, "nat46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("outlook_web_access"); ok || d.HasChange("outlook_web_access") {
		t, err := expandObjectFirewallVipOutlookWebAccess(d, v, "outlook_web_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outlook-web-access"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok || d.HasChange("persistence") {
		t, err := expandObjectFirewallVipPersistence(d, v, "persistence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("portforward"); ok || d.HasChange("portforward") {
		t, err := expandObjectFirewallVipPortforward(d, v, "portforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portforward"] = t
		}
	}

	if v, ok := d.GetOk("portmapping_type"); ok || d.HasChange("portmapping_type") {
		t, err := expandObjectFirewallVipPortmappingType(d, v, "portmapping_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portmapping-type"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectFirewallVipProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok || d.HasChange("realservers") {
		t, err := expandObjectFirewallVipRealservers(d, v, "realservers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandObjectFirewallVipServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandObjectFirewallVipService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("src_filter"); ok || d.HasChange("src_filter") {
		t, err := expandObjectFirewallVipSrcFilter(d, v, "src_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-filter"] = t
		}
	}

	if v, ok := d.GetOk("srcintf_filter"); ok || d.HasChange("srcintf_filter") {
		t, err := expandObjectFirewallVipSrcintfFilter(d, v, "srcintf_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf-filter"] = t
		}
	}

	if v, ok := d.GetOk("ssl_accept_ffdhe_groups"); ok || d.HasChange("ssl_accept_ffdhe_groups") {
		t, err := expandObjectFirewallVipSslAcceptFfdheGroups(d, v, "ssl_accept_ffdhe_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-accept-ffdhe-groups"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandObjectFirewallVipSslAlgorithm(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandObjectFirewallVipSslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandObjectFirewallVipSslCipherSuites(d, v, "ssl_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_fallback"); ok || d.HasChange("ssl_client_fallback") {
		t, err := expandObjectFirewallVipSslClientFallback(d, v, "ssl_client_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-fallback"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_rekey_count"); ok || d.HasChange("ssl_client_rekey_count") {
		t, err := expandObjectFirewallVipSslClientRekeyCount(d, v, "ssl_client_rekey_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-rekey-count"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok || d.HasChange("ssl_client_renegotiation") {
		t, err := expandObjectFirewallVipSslClientRenegotiation(d, v, "ssl_client_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_max"); ok || d.HasChange("ssl_client_session_state_max") {
		t, err := expandObjectFirewallVipSslClientSessionStateMax(d, v, "ssl_client_session_state_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_timeout"); ok || d.HasChange("ssl_client_session_state_timeout") {
		t, err := expandObjectFirewallVipSslClientSessionStateTimeout(d, v, "ssl_client_session_state_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_type"); ok || d.HasChange("ssl_client_session_state_type") {
		t, err := expandObjectFirewallVipSslClientSessionStateType(d, v, "ssl_client_session_state_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandObjectFirewallVipSslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp"); ok || d.HasChange("ssl_hpkp") {
		t, err := expandObjectFirewallVipSslHpkp(d, v, "ssl_hpkp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_age"); ok || d.HasChange("ssl_hpkp_age") {
		t, err := expandObjectFirewallVipSslHpkpAge(d, v, "ssl_hpkp_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_backup"); ok || d.HasChange("ssl_hpkp_backup") {
		t, err := expandObjectFirewallVipSslHpkpBackup(d, v, "ssl_hpkp_backup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-backup"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_include_subdomains"); ok || d.HasChange("ssl_hpkp_include_subdomains") {
		t, err := expandObjectFirewallVipSslHpkpIncludeSubdomains(d, v, "ssl_hpkp_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_primary"); ok || d.HasChange("ssl_hpkp_primary") {
		t, err := expandObjectFirewallVipSslHpkpPrimary(d, v, "ssl_hpkp_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-primary"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_report_uri"); ok || d.HasChange("ssl_hpkp_report_uri") {
		t, err := expandObjectFirewallVipSslHpkpReportUri(d, v, "ssl_hpkp_report_uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-report-uri"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts"); ok || d.HasChange("ssl_hsts") {
		t, err := expandObjectFirewallVipSslHsts(d, v, "ssl_hsts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_age"); ok || d.HasChange("ssl_hsts_age") {
		t, err := expandObjectFirewallVipSslHstsAge(d, v, "ssl_hsts_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_include_subdomains"); ok || d.HasChange("ssl_hsts_include_subdomains") {
		t, err := expandObjectFirewallVipSslHstsIncludeSubdomains(d, v, "ssl_hsts_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_location_conversion"); ok || d.HasChange("ssl_http_location_conversion") {
		t, err := expandObjectFirewallVipSslHttpLocationConversion(d, v, "ssl_http_location_conversion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-location-conversion"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_match_host"); ok || d.HasChange("ssl_http_match_host") {
		t, err := expandObjectFirewallVipSslHttpMatchHost(d, v, "ssl_http_match_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-match-host"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandObjectFirewallVipSslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandObjectFirewallVipSslMinVersion(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok || d.HasChange("ssl_mode") {
		t, err := expandObjectFirewallVipSslMode(d, v, "ssl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_pfs"); ok || d.HasChange("ssl_pfs") {
		t, err := expandObjectFirewallVipSslPfs(d, v, "ssl_pfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-pfs"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok || d.HasChange("ssl_send_empty_frags") {
		t, err := expandObjectFirewallVipSslSendEmptyFrags(d, v, "ssl_send_empty_frags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_algorithm"); ok || d.HasChange("ssl_server_algorithm") {
		t, err := expandObjectFirewallVipSslServerAlgorithm(d, v, "ssl_server_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_cipher_suites"); ok || d.HasChange("ssl_server_cipher_suites") {
		t, err := expandObjectFirewallVipSslServerCipherSuites(d, v, "ssl_server_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_max_version"); ok || d.HasChange("ssl_server_max_version") {
		t, err := expandObjectFirewallVipSslServerMaxVersion(d, v, "ssl_server_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_min_version"); ok || d.HasChange("ssl_server_min_version") {
		t, err := expandObjectFirewallVipSslServerMinVersion(d, v, "ssl_server_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_renegotiation"); ok || d.HasChange("ssl_server_renegotiation") {
		t, err := expandObjectFirewallVipSslServerRenegotiation(d, v, "ssl_server_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_max"); ok || d.HasChange("ssl_server_session_state_max") {
		t, err := expandObjectFirewallVipSslServerSessionStateMax(d, v, "ssl_server_session_state_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_timeout"); ok || d.HasChange("ssl_server_session_state_timeout") {
		t, err := expandObjectFirewallVipSslServerSessionStateTimeout(d, v, "ssl_server_session_state_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_type"); ok || d.HasChange("ssl_server_session_state_type") {
		t, err := expandObjectFirewallVipSslServerSessionStateType(d, v, "ssl_server_session_state_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFirewallVipStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallVipType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectFirewallVipUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("weblogic_server"); ok || d.HasChange("weblogic_server") {
		t, err := expandObjectFirewallVipWeblogicServer(d, v, "weblogic_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weblogic-server"] = t
		}
	}

	if v, ok := d.GetOk("websphere_server"); ok || d.HasChange("websphere_server") {
		t, err := expandObjectFirewallVipWebsphereServer(d, v, "websphere_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["websphere-server"] = t
		}
	}

	return &obj, nil
}
