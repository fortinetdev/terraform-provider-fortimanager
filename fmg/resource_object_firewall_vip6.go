// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure virtual IP for IPv6.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallVip6() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallVip6Create,
		Read:   resourceObjectFirewallVip6Read,
		Update: resourceObjectFirewallVip6Update,
		Delete: resourceObjectFirewallVip6Delete,

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
			"add_nat64_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
						"add_nat64_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"embedded_ipv4_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"extip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"extport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_cookie_age": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"http_cookie_domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_cookie_domain_from_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"http_ip_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_ip_header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_multiplex": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_redirect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"https_cookie_secure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ipv4_mappedip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv4_mappedport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ldb_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mappedip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mappedport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_embryonic_connections": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
						"nat64": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nat66": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ndp_reply": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outlook_web_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"persistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"portforward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"realservers": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"healthcheck": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"holddown_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
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
										Computed: true,
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
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"translate_host": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"weight": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"server_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"src_filter": &schema.Schema{
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
							Computed: true,
						},
						"ssl_hsts_age": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ssl_hsts_include_subdomains": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_http_location_conversion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"websphere_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"embedded_ipv4_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_cookie_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_cookie_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_cookie_domain_from_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"http_ip_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_ip_header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_multiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_cookie_secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ipv4_mappedip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mappedip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_embryonic_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat66": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ndp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outlook_web_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realservers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"healthcheck": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"holddown_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
							Computed: true,
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
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"translate_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_filter": &schema.Schema{
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
				Computed: true,
			},
			"ssl_hsts_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssl_hsts_include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_http_location_conversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"websphere_server": &schema.Schema{
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

func resourceObjectFirewallVip6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectFirewallVip6(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallVip6 resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallVip6(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallVip6 resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallVip6Read(d, m)
}

func resourceObjectFirewallVip6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallVip6(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallVip6 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallVip6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallVip6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallVip6Read(d, m)
}

func resourceObjectFirewallVip6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallVip6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallVip6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallVip6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallVip6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallVip6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallVip6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallVip6 resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallVip6AddNat64Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6ArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Color(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallVip6DynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat64_route"
		if _, ok := i["add-nat64-route"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingAddNat64Route(i["add-nat64-route"], d, pre_append)
			tmp["add_nat64_route"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-AddNat64Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := i["arp-reply"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingArpReply(i["arp-reply"], d, pre_append)
			tmp["arp_reply"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-ArpReply")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embedded_ipv4_address"
		if _, ok := i["embedded-ipv4-address"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingEmbeddedIpv4Address(i["embedded-ipv4-address"], d, pre_append)
			tmp["embedded_ipv4_address"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-EmbeddedIpv4Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := i["extip"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingExtip(i["extip"], d, pre_append)
			tmp["extip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Extip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := i["extport"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingExtport(i["extport"], d, pre_append)
			tmp["extport"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Extport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpCookieAge(i["http-cookie-age"], d, pre_append)
			tmp["http_cookie_age"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpCookieAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpCookieDomain(i["http-cookie-domain"], d, pre_append)
			tmp["http_cookie_domain"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpCookieDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append)
			tmp["http_cookie_domain_from_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpCookieDomainFromHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpCookieGeneration(i["http-cookie-generation"], d, pre_append)
			tmp["http_cookie_generation"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpCookieGeneration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpCookiePath(i["http-cookie-path"], d, pre_append)
			tmp["http_cookie_path"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpCookiePath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpCookieShare(i["http-cookie-share"], d, pre_append)
			tmp["http_cookie_share"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpCookieShare")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header"
		if _, ok := i["http-ip-header"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpIpHeader(i["http-ip-header"], d, pre_append)
			tmp["http_ip_header"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpIpHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header_name"
		if _, ok := i["http-ip-header-name"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpIpHeaderName(i["http-ip-header-name"], d, pre_append)
			tmp["http_ip_header_name"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpIpHeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex"
		if _, ok := i["http-multiplex"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpMultiplex(i["http-multiplex"], d, pre_append)
			tmp["http_multiplex"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpMultiplex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_redirect"
		if _, ok := i["http-redirect"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpRedirect(i["http-redirect"], d, pre_append)
			tmp["http_redirect"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpRedirect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingHttpsCookieSecure(i["https-cookie-secure"], d, pre_append)
			tmp["https_cookie_secure"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-HttpsCookieSecure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_mappedip"
		if _, ok := i["ipv4-mappedip"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingIpv4Mappedip(i["ipv4-mappedip"], d, pre_append)
			tmp["ipv4_mappedip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Ipv4Mappedip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_mappedport"
		if _, ok := i["ipv4-mappedport"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingIpv4Mappedport(i["ipv4-mappedport"], d, pre_append)
			tmp["ipv4_mappedport"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Ipv4Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingLdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedip"
		if _, ok := i["mappedip"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingMappedip(i["mappedip"], d, pre_append)
			tmp["mappedip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Mappedip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingMappedport(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_embryonic_connections"
		if _, ok := i["max-embryonic-connections"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingMaxEmbryonicConnections(i["max-embryonic-connections"], d, pre_append)
			tmp["max_embryonic_connections"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-MaxEmbryonicConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat_source_vip"
		if _, ok := i["nat-source-vip"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingNatSourceVip(i["nat-source-vip"], d, pre_append)
			tmp["nat_source_vip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-NatSourceVip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64"
		if _, ok := i["nat64"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingNat64(i["nat64"], d, pre_append)
			tmp["nat64"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Nat64")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat66"
		if _, ok := i["nat66"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingNat66(i["nat66"], d, pre_append)
			tmp["nat66"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Nat66")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ndp_reply"
		if _, ok := i["ndp-reply"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingNdpReply(i["ndp-reply"], d, pre_append)
			tmp["ndp_reply"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-NdpReply")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "outlook_web_access"
		if _, ok := i["outlook-web-access"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingOutlookWebAccess(i["outlook-web-access"], d, pre_append)
			tmp["outlook_web_access"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-OutlookWebAccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingPersistence(i["persistence"], d, pre_append)
			tmp["persistence"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Persistence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portforward"
		if _, ok := i["portforward"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingPortforward(i["portforward"], d, pre_append)
			tmp["portforward"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Portforward")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealservers(i["realservers"], d, pre_append)
			tmp["realservers"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Realservers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := i["server-type"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingServerType(i["server-type"], d, pre_append)
			tmp["server_type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-ServerType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_filter"
		if _, ok := i["src-filter"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSrcFilter(i["src-filter"], d, pre_append)
			tmp["src_filter"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SrcFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_accept_ffdhe_groups"
		if _, ok := i["ssl-accept-ffdhe-groups"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslAcceptFfdheGroups(i["ssl-accept-ffdhe-groups"], d, pre_append)
			tmp["ssl_accept_ffdhe_groups"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslAcceptFfdheGroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslAlgorithm(i["ssl-algorithm"], d, pre_append)
			tmp["ssl_algorithm"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_certificate"
		if _, ok := i["ssl-certificate"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslCertificate(i["ssl-certificate"], d, pre_append)
			tmp["ssl_certificate"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslCipherSuites(i["ssl-cipher-suites"], d, pre_append)
			tmp["ssl_cipher_suites"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslCipherSuites")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_fallback"
		if _, ok := i["ssl-client-fallback"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslClientFallback(i["ssl-client-fallback"], d, pre_append)
			tmp["ssl_client_fallback"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslClientFallback")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_rekey_count"
		if _, ok := i["ssl-client-rekey-count"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslClientRekeyCount(i["ssl-client-rekey-count"], d, pre_append)
			tmp["ssl_client_rekey_count"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslClientRekeyCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_renegotiation"
		if _, ok := i["ssl-client-renegotiation"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslClientRenegotiation(i["ssl-client-renegotiation"], d, pre_append)
			tmp["ssl_client_renegotiation"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslClientRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_max"
		if _, ok := i["ssl-client-session-state-max"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslClientSessionStateMax(i["ssl-client-session-state-max"], d, pre_append)
			tmp["ssl_client_session_state_max"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslClientSessionStateMax")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_timeout"
		if _, ok := i["ssl-client-session-state-timeout"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslClientSessionStateTimeout(i["ssl-client-session-state-timeout"], d, pre_append)
			tmp["ssl_client_session_state_timeout"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslClientSessionStateTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_type"
		if _, ok := i["ssl-client-session-state-type"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslClientSessionStateType(i["ssl-client-session-state-type"], d, pre_append)
			tmp["ssl_client_session_state_type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslClientSessionStateType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslDhBits(i["ssl-dh-bits"], d, pre_append)
			tmp["ssl_dh_bits"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslDhBits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp"
		if _, ok := i["ssl-hpkp"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHpkp(i["ssl-hpkp"], d, pre_append)
			tmp["ssl_hpkp"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHpkp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_age"
		if _, ok := i["ssl-hpkp-age"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHpkpAge(i["ssl-hpkp-age"], d, pre_append)
			tmp["ssl_hpkp_age"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHpkpAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_backup"
		if _, ok := i["ssl-hpkp-backup"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHpkpBackup(i["ssl-hpkp-backup"], d, pre_append)
			tmp["ssl_hpkp_backup"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHpkpBackup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_include_subdomains"
		if _, ok := i["ssl-hpkp-include-subdomains"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHpkpIncludeSubdomains(i["ssl-hpkp-include-subdomains"], d, pre_append)
			tmp["ssl_hpkp_include_subdomains"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHpkpIncludeSubdomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_primary"
		if _, ok := i["ssl-hpkp-primary"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHpkpPrimary(i["ssl-hpkp-primary"], d, pre_append)
			tmp["ssl_hpkp_primary"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHpkpPrimary")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_report_uri"
		if _, ok := i["ssl-hpkp-report-uri"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHpkpReportUri(i["ssl-hpkp-report-uri"], d, pre_append)
			tmp["ssl_hpkp_report_uri"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHpkpReportUri")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts"
		if _, ok := i["ssl-hsts"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHsts(i["ssl-hsts"], d, pre_append)
			tmp["ssl_hsts"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHsts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_age"
		if _, ok := i["ssl-hsts-age"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHstsAge(i["ssl-hsts-age"], d, pre_append)
			tmp["ssl_hsts_age"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHstsAge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_include_subdomains"
		if _, ok := i["ssl-hsts-include-subdomains"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHstsIncludeSubdomains(i["ssl-hsts-include-subdomains"], d, pre_append)
			tmp["ssl_hsts_include_subdomains"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHstsIncludeSubdomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_location_conversion"
		if _, ok := i["ssl-http-location-conversion"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHttpLocationConversion(i["ssl-http-location-conversion"], d, pre_append)
			tmp["ssl_http_location_conversion"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHttpLocationConversion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_match_host"
		if _, ok := i["ssl-http-match-host"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslHttpMatchHost(i["ssl-http-match-host"], d, pre_append)
			tmp["ssl_http_match_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslHttpMatchHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslMinVersion(i["ssl-min-version"], d, pre_append)
			tmp["ssl_min_version"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_mode"
		if _, ok := i["ssl-mode"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslMode(i["ssl-mode"], d, pre_append)
			tmp["ssl_mode"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_pfs"
		if _, ok := i["ssl-pfs"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslPfs(i["ssl-pfs"], d, pre_append)
			tmp["ssl_pfs"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslPfs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_send_empty_frags"
		if _, ok := i["ssl-send-empty-frags"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslSendEmptyFrags(i["ssl-send-empty-frags"], d, pre_append)
			tmp["ssl_send_empty_frags"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslSendEmptyFrags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_algorithm"
		if _, ok := i["ssl-server-algorithm"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslServerAlgorithm(i["ssl-server-algorithm"], d, pre_append)
			tmp["ssl_server_algorithm"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslServerAlgorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_max_version"
		if _, ok := i["ssl-server-max-version"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslServerMaxVersion(i["ssl-server-max-version"], d, pre_append)
			tmp["ssl_server_max_version"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslServerMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_min_version"
		if _, ok := i["ssl-server-min-version"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslServerMinVersion(i["ssl-server-min-version"], d, pre_append)
			tmp["ssl_server_min_version"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslServerMinVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_renegotiation"
		if _, ok := i["ssl-server-renegotiation"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslServerRenegotiation(i["ssl-server-renegotiation"], d, pre_append)
			tmp["ssl_server_renegotiation"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslServerRenegotiation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_max"
		if _, ok := i["ssl-server-session-state-max"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslServerSessionStateMax(i["ssl-server-session-state-max"], d, pre_append)
			tmp["ssl_server_session_state_max"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslServerSessionStateMax")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_timeout"
		if _, ok := i["ssl-server-session-state-timeout"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslServerSessionStateTimeout(i["ssl-server-session-state-timeout"], d, pre_append)
			tmp["ssl_server_session_state_timeout"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslServerSessionStateTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_type"
		if _, ok := i["ssl-server-session-state-type"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslServerSessionStateType(i["ssl-server-session-state-type"], d, pre_append)
			tmp["ssl_server_session_state_type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-SslServerSessionStateType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-Uuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weblogic_server"
		if _, ok := i["weblogic-server"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingWeblogicServer(i["weblogic-server"], d, pre_append)
			tmp["weblogic_server"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-WeblogicServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "websphere_server"
		if _, ok := i["websphere-server"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingWebsphereServer(i["websphere-server"], d, pre_append)
			tmp["websphere_server"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-DynamicMapping-WebsphereServer")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallVip6DynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallVip6DynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallVip6DynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingAddNat64Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingEmbeddedIpv4Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingExtip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFirewallVip6DynamicMappingExtport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpIpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpIpHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingIpv4Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingIpv4Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingMappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFirewallVip6DynamicMappingMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingMaxEmbryonicConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingNatSourceVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingNat66(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingNdpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingOutlookWebAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingPersistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingPortforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversClientIp(i["client-ip"], d, pre_append)
			tmp["client_ip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-ClientIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversHealthcheck(i["healthcheck"], d, pre_append)
			tmp["healthcheck"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-Healthcheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingRealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallVip6DynamicMappingRealserversClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingRealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSrcFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVip6DynamicMappingSslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallVip6DynamicMappingSslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenObjectFirewallVip6DynamicMappingSslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6DynamicMapping-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallVip6DynamicMappingSslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVip6DynamicMappingSslClientFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHpkp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHpkpAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHpkpBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHsts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHstsAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslPfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslServerMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslServerRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingSslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingWeblogicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6DynamicMappingWebsphereServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6EmbeddedIpv4Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Extip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Extport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpCookieDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpCookiePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpCookieShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpIpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpIpHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpMultiplex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6HttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Ipv4Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Ipv4Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6LdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVip6Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6MaxEmbryonicConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Monitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6NatSourceVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Nat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Nat66(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6NdpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6OutlookWebAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Persistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Portforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Realservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {
			v := flattenObjectFirewallVip6RealserversClientIp(i["client-ip"], d, pre_append)
			tmp["client_ip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-ClientIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {
			v := flattenObjectFirewallVip6RealserversHealthcheck(i["healthcheck"], d, pre_append)
			tmp["healthcheck"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-Healthcheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenObjectFirewallVip6RealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {
			v := flattenObjectFirewallVip6RealserversHttpHost(i["http-host"], d, pre_append)
			tmp["http_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-HttpHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallVip6RealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFirewallVip6RealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenObjectFirewallVip6RealserversMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenObjectFirewallVip6RealserversMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectFirewallVip6RealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFirewallVip6RealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := i["translate-host"]; ok {
			v := flattenObjectFirewallVip6RealserversTranslateHost(i["translate-host"], d, pre_append)
			tmp["translate_host"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-TranslateHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectFirewallVip6RealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-Realservers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallVip6RealserversClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversHttpHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6RealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6ServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SrcFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVip6SslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallVip6SslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFirewallVip6SslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenObjectFirewallVip6SslCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-SslCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallVip6SslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVip6SslClientFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHpkp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHpkpAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHpkpBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHsts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHstsAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslPfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslServerCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallVip6SslServerCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-SslServerCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFirewallVip6SslServerCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-SslServerCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {
			v := flattenObjectFirewallVip6SslServerCipherSuitesVersions(i["versions"], d, pre_append)
			tmp["versions"] = fortiAPISubPartPatch(v, "ObjectFirewallVip6-SslServerCipherSuites-Versions")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallVip6SslServerCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslServerCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslServerCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVip6SslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslServerMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslServerRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6SslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6WeblogicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip6WebsphereServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallVip6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("add_nat64_route", flattenObjectFirewallVip6AddNat64Route(o["add-nat64-route"], d, "add_nat64_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-nat64-route"], "ObjectFirewallVip6-AddNat64Route"); ok {
			if err = d.Set("add_nat64_route", vv); err != nil {
				return fmt.Errorf("Error reading add_nat64_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_nat64_route: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenObjectFirewallVip6ArpReply(o["arp-reply"], d, "arp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-reply"], "ObjectFirewallVip6-ArpReply"); ok {
			if err = d.Set("arp_reply", vv); err != nil {
				return fmt.Errorf("Error reading arp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("color", flattenObjectFirewallVip6Color(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallVip6-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallVip6Comment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallVip6-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectFirewallVip6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallVip6-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectFirewallVip6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallVip6-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("embedded_ipv4_address", flattenObjectFirewallVip6EmbeddedIpv4Address(o["embedded-ipv4-address"], d, "embedded_ipv4_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["embedded-ipv4-address"], "ObjectFirewallVip6-EmbeddedIpv4Address"); ok {
			if err = d.Set("embedded_ipv4_address", vv); err != nil {
				return fmt.Errorf("Error reading embedded_ipv4_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading embedded_ipv4_address: %v", err)
		}
	}

	if err = d.Set("extip", flattenObjectFirewallVip6Extip(o["extip"], d, "extip")); err != nil {
		if vv, ok := fortiAPIPatch(o["extip"], "ObjectFirewallVip6-Extip"); ok {
			if err = d.Set("extip", vv); err != nil {
				return fmt.Errorf("Error reading extip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if err = d.Set("extport", flattenObjectFirewallVip6Extport(o["extport"], d, "extport")); err != nil {
		if vv, ok := fortiAPIPatch(o["extport"], "ObjectFirewallVip6-Extport"); ok {
			if err = d.Set("extport", vv); err != nil {
				return fmt.Errorf("Error reading extport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("http_cookie_age", flattenObjectFirewallVip6HttpCookieAge(o["http-cookie-age"], d, "http_cookie_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-age"], "ObjectFirewallVip6-HttpCookieAge"); ok {
			if err = d.Set("http_cookie_age", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_age: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain", flattenObjectFirewallVip6HttpCookieDomain(o["http-cookie-domain"], d, "http_cookie_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain"], "ObjectFirewallVip6-HttpCookieDomain"); ok {
			if err = d.Set("http_cookie_domain", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain_from_host", flattenObjectFirewallVip6HttpCookieDomainFromHost(o["http-cookie-domain-from-host"], d, "http_cookie_domain_from_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-domain-from-host"], "ObjectFirewallVip6-HttpCookieDomainFromHost"); ok {
			if err = d.Set("http_cookie_domain_from_host", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
		}
	}

	if err = d.Set("http_cookie_generation", flattenObjectFirewallVip6HttpCookieGeneration(o["http-cookie-generation"], d, "http_cookie_generation")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-generation"], "ObjectFirewallVip6-HttpCookieGeneration"); ok {
			if err = d.Set("http_cookie_generation", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_generation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_generation: %v", err)
		}
	}

	if err = d.Set("http_cookie_path", flattenObjectFirewallVip6HttpCookiePath(o["http-cookie-path"], d, "http_cookie_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-path"], "ObjectFirewallVip6-HttpCookiePath"); ok {
			if err = d.Set("http_cookie_path", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_path: %v", err)
		}
	}

	if err = d.Set("http_cookie_share", flattenObjectFirewallVip6HttpCookieShare(o["http-cookie-share"], d, "http_cookie_share")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-cookie-share"], "ObjectFirewallVip6-HttpCookieShare"); ok {
			if err = d.Set("http_cookie_share", vv); err != nil {
				return fmt.Errorf("Error reading http_cookie_share: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_cookie_share: %v", err)
		}
	}

	if err = d.Set("http_ip_header", flattenObjectFirewallVip6HttpIpHeader(o["http-ip-header"], d, "http_ip_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-ip-header"], "ObjectFirewallVip6-HttpIpHeader"); ok {
			if err = d.Set("http_ip_header", vv); err != nil {
				return fmt.Errorf("Error reading http_ip_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_ip_header: %v", err)
		}
	}

	if err = d.Set("http_ip_header_name", flattenObjectFirewallVip6HttpIpHeaderName(o["http-ip-header-name"], d, "http_ip_header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-ip-header-name"], "ObjectFirewallVip6-HttpIpHeaderName"); ok {
			if err = d.Set("http_ip_header_name", vv); err != nil {
				return fmt.Errorf("Error reading http_ip_header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_ip_header_name: %v", err)
		}
	}

	if err = d.Set("http_multiplex", flattenObjectFirewallVip6HttpMultiplex(o["http-multiplex"], d, "http_multiplex")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-multiplex"], "ObjectFirewallVip6-HttpMultiplex"); ok {
			if err = d.Set("http_multiplex", vv); err != nil {
				return fmt.Errorf("Error reading http_multiplex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_multiplex: %v", err)
		}
	}

	if err = d.Set("http_redirect", flattenObjectFirewallVip6HttpRedirect(o["http-redirect"], d, "http_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-redirect"], "ObjectFirewallVip6-HttpRedirect"); ok {
			if err = d.Set("http_redirect", vv); err != nil {
				return fmt.Errorf("Error reading http_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_redirect: %v", err)
		}
	}

	if err = d.Set("https_cookie_secure", flattenObjectFirewallVip6HttpsCookieSecure(o["https-cookie-secure"], d, "https_cookie_secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-cookie-secure"], "ObjectFirewallVip6-HttpsCookieSecure"); ok {
			if err = d.Set("https_cookie_secure", vv); err != nil {
				return fmt.Errorf("Error reading https_cookie_secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_cookie_secure: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallVip6Id(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallVip6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ipv4_mappedip", flattenObjectFirewallVip6Ipv4Mappedip(o["ipv4-mappedip"], d, "ipv4_mappedip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-mappedip"], "ObjectFirewallVip6-Ipv4Mappedip"); ok {
			if err = d.Set("ipv4_mappedip", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_mappedip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_mappedip: %v", err)
		}
	}

	if err = d.Set("ipv4_mappedport", flattenObjectFirewallVip6Ipv4Mappedport(o["ipv4-mappedport"], d, "ipv4_mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-mappedport"], "ObjectFirewallVip6-Ipv4Mappedport"); ok {
			if err = d.Set("ipv4_mappedport", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_mappedport: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenObjectFirewallVip6LdbMethod(o["ldb-method"], d, "ldb_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldb-method"], "ObjectFirewallVip6-LdbMethod"); ok {
			if err = d.Set("ldb_method", vv); err != nil {
				return fmt.Errorf("Error reading ldb_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("mappedip", flattenObjectFirewallVip6Mappedip(o["mappedip"], d, "mappedip")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedip"], "ObjectFirewallVip6-Mappedip"); ok {
			if err = d.Set("mappedip", vv); err != nil {
				return fmt.Errorf("Error reading mappedip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedip: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenObjectFirewallVip6Mappedport(o["mappedport"], d, "mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedport"], "ObjectFirewallVip6-Mappedport"); ok {
			if err = d.Set("mappedport", vv); err != nil {
				return fmt.Errorf("Error reading mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("max_embryonic_connections", flattenObjectFirewallVip6MaxEmbryonicConnections(o["max-embryonic-connections"], d, "max_embryonic_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-embryonic-connections"], "ObjectFirewallVip6-MaxEmbryonicConnections"); ok {
			if err = d.Set("max_embryonic_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_embryonic_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_embryonic_connections: %v", err)
		}
	}

	if err = d.Set("monitor", flattenObjectFirewallVip6Monitor(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "ObjectFirewallVip6-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallVip6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallVip6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat_source_vip", flattenObjectFirewallVip6NatSourceVip(o["nat-source-vip"], d, "nat_source_vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-source-vip"], "ObjectFirewallVip6-NatSourceVip"); ok {
			if err = d.Set("nat_source_vip", vv); err != nil {
				return fmt.Errorf("Error reading nat_source_vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_source_vip: %v", err)
		}
	}

	if err = d.Set("nat64", flattenObjectFirewallVip6Nat64(o["nat64"], d, "nat64")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64"], "ObjectFirewallVip6-Nat64"); ok {
			if err = d.Set("nat64", vv); err != nil {
				return fmt.Errorf("Error reading nat64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("nat66", flattenObjectFirewallVip6Nat66(o["nat66"], d, "nat66")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat66"], "ObjectFirewallVip6-Nat66"); ok {
			if err = d.Set("nat66", vv); err != nil {
				return fmt.Errorf("Error reading nat66: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat66: %v", err)
		}
	}

	if err = d.Set("ndp_reply", flattenObjectFirewallVip6NdpReply(o["ndp-reply"], d, "ndp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["ndp-reply"], "ObjectFirewallVip6-NdpReply"); ok {
			if err = d.Set("ndp_reply", vv); err != nil {
				return fmt.Errorf("Error reading ndp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ndp_reply: %v", err)
		}
	}

	if err = d.Set("outlook_web_access", flattenObjectFirewallVip6OutlookWebAccess(o["outlook-web-access"], d, "outlook_web_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["outlook-web-access"], "ObjectFirewallVip6-OutlookWebAccess"); ok {
			if err = d.Set("outlook_web_access", vv); err != nil {
				return fmt.Errorf("Error reading outlook_web_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outlook_web_access: %v", err)
		}
	}

	if err = d.Set("persistence", flattenObjectFirewallVip6Persistence(o["persistence"], d, "persistence")); err != nil {
		if vv, ok := fortiAPIPatch(o["persistence"], "ObjectFirewallVip6-Persistence"); ok {
			if err = d.Set("persistence", vv); err != nil {
				return fmt.Errorf("Error reading persistence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("portforward", flattenObjectFirewallVip6Portforward(o["portforward"], d, "portforward")); err != nil {
		if vv, ok := fortiAPIPatch(o["portforward"], "ObjectFirewallVip6-Portforward"); ok {
			if err = d.Set("portforward", vv); err != nil {
				return fmt.Errorf("Error reading portforward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portforward: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectFirewallVip6Protocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectFirewallVip6-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("realservers", flattenObjectFirewallVip6Realservers(o["realservers"], d, "realservers")); err != nil {
			if vv, ok := fortiAPIPatch(o["realservers"], "ObjectFirewallVip6-Realservers"); ok {
				if err = d.Set("realservers", vv); err != nil {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenObjectFirewallVip6Realservers(o["realservers"], d, "realservers")); err != nil {
				if vv, ok := fortiAPIPatch(o["realservers"], "ObjectFirewallVip6-Realservers"); ok {
					if err = d.Set("realservers", vv); err != nil {
						return fmt.Errorf("Error reading realservers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenObjectFirewallVip6ServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "ObjectFirewallVip6-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("src_filter", flattenObjectFirewallVip6SrcFilter(o["src-filter"], d, "src_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-filter"], "ObjectFirewallVip6-SrcFilter"); ok {
			if err = d.Set("src_filter", vv); err != nil {
				return fmt.Errorf("Error reading src_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_filter: %v", err)
		}
	}

	if err = d.Set("ssl_accept_ffdhe_groups", flattenObjectFirewallVip6SslAcceptFfdheGroups(o["ssl-accept-ffdhe-groups"], d, "ssl_accept_ffdhe_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-accept-ffdhe-groups"], "ObjectFirewallVip6-SslAcceptFfdheGroups"); ok {
			if err = d.Set("ssl_accept_ffdhe_groups", vv); err != nil {
				return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenObjectFirewallVip6SslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "ObjectFirewallVip6-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenObjectFirewallVip6SslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "ObjectFirewallVip6-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenObjectFirewallVip6SslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "ObjectFirewallVip6-SslCipherSuites"); ok {
				if err = d.Set("ssl_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenObjectFirewallVip6SslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "ObjectFirewallVip6-SslCipherSuites"); ok {
					if err = d.Set("ssl_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_client_fallback", flattenObjectFirewallVip6SslClientFallback(o["ssl-client-fallback"], d, "ssl_client_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-fallback"], "ObjectFirewallVip6-SslClientFallback"); ok {
			if err = d.Set("ssl_client_fallback", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
		}
	}

	if err = d.Set("ssl_client_rekey_count", flattenObjectFirewallVip6SslClientRekeyCount(o["ssl-client-rekey-count"], d, "ssl_client_rekey_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-rekey-count"], "ObjectFirewallVip6-SslClientRekeyCount"); ok {
			if err = d.Set("ssl_client_rekey_count", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenObjectFirewallVip6SslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-renegotiation"], "ObjectFirewallVip6-SslClientRenegotiation"); ok {
			if err = d.Set("ssl_client_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_max", flattenObjectFirewallVip6SslClientSessionStateMax(o["ssl-client-session-state-max"], d, "ssl_client_session_state_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-max"], "ObjectFirewallVip6-SslClientSessionStateMax"); ok {
			if err = d.Set("ssl_client_session_state_max", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_timeout", flattenObjectFirewallVip6SslClientSessionStateTimeout(o["ssl-client-session-state-timeout"], d, "ssl_client_session_state_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-timeout"], "ObjectFirewallVip6-SslClientSessionStateTimeout"); ok {
			if err = d.Set("ssl_client_session_state_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_type", flattenObjectFirewallVip6SslClientSessionStateType(o["ssl-client-session-state-type"], d, "ssl_client_session_state_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-session-state-type"], "ObjectFirewallVip6-SslClientSessionStateType"); ok {
			if err = d.Set("ssl_client_session_state_type", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenObjectFirewallVip6SslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "ObjectFirewallVip6-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp", flattenObjectFirewallVip6SslHpkp(o["ssl-hpkp"], d, "ssl_hpkp")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp"], "ObjectFirewallVip6-SslHpkp"); ok {
			if err = d.Set("ssl_hpkp", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_age", flattenObjectFirewallVip6SslHpkpAge(o["ssl-hpkp-age"], d, "ssl_hpkp_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-age"], "ObjectFirewallVip6-SslHpkpAge"); ok {
			if err = d.Set("ssl_hpkp_age", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_backup", flattenObjectFirewallVip6SslHpkpBackup(o["ssl-hpkp-backup"], d, "ssl_hpkp_backup")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-backup"], "ObjectFirewallVip6-SslHpkpBackup"); ok {
			if err = d.Set("ssl_hpkp_backup", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_include_subdomains", flattenObjectFirewallVip6SslHpkpIncludeSubdomains(o["ssl-hpkp-include-subdomains"], d, "ssl_hpkp_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-include-subdomains"], "ObjectFirewallVip6-SslHpkpIncludeSubdomains"); ok {
			if err = d.Set("ssl_hpkp_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_primary", flattenObjectFirewallVip6SslHpkpPrimary(o["ssl-hpkp-primary"], d, "ssl_hpkp_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-primary"], "ObjectFirewallVip6-SslHpkpPrimary"); ok {
			if err = d.Set("ssl_hpkp_primary", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_report_uri", flattenObjectFirewallVip6SslHpkpReportUri(o["ssl-hpkp-report-uri"], d, "ssl_hpkp_report_uri")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hpkp-report-uri"], "ObjectFirewallVip6-SslHpkpReportUri"); ok {
			if err = d.Set("ssl_hpkp_report_uri", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
		}
	}

	if err = d.Set("ssl_hsts", flattenObjectFirewallVip6SslHsts(o["ssl-hsts"], d, "ssl_hsts")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts"], "ObjectFirewallVip6-SslHsts"); ok {
			if err = d.Set("ssl_hsts", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_age", flattenObjectFirewallVip6SslHstsAge(o["ssl-hsts-age"], d, "ssl_hsts_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts-age"], "ObjectFirewallVip6-SslHstsAge"); ok {
			if err = d.Set("ssl_hsts_age", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_include_subdomains", flattenObjectFirewallVip6SslHstsIncludeSubdomains(o["ssl-hsts-include-subdomains"], d, "ssl_hsts_include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-hsts-include-subdomains"], "ObjectFirewallVip6-SslHstsIncludeSubdomains"); ok {
			if err = d.Set("ssl_hsts_include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_http_location_conversion", flattenObjectFirewallVip6SslHttpLocationConversion(o["ssl-http-location-conversion"], d, "ssl_http_location_conversion")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-http-location-conversion"], "ObjectFirewallVip6-SslHttpLocationConversion"); ok {
			if err = d.Set("ssl_http_location_conversion", vv); err != nil {
				return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
		}
	}

	if err = d.Set("ssl_http_match_host", flattenObjectFirewallVip6SslHttpMatchHost(o["ssl-http-match-host"], d, "ssl_http_match_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-http-match-host"], "ObjectFirewallVip6-SslHttpMatchHost"); ok {
			if err = d.Set("ssl_http_match_host", vv); err != nil {
				return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenObjectFirewallVip6SslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "ObjectFirewallVip6-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenObjectFirewallVip6SslMinVersion(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "ObjectFirewallVip6-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_mode", flattenObjectFirewallVip6SslMode(o["ssl-mode"], d, "ssl_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mode"], "ObjectFirewallVip6-SslMode"); ok {
			if err = d.Set("ssl_mode", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if err = d.Set("ssl_pfs", flattenObjectFirewallVip6SslPfs(o["ssl-pfs"], d, "ssl_pfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-pfs"], "ObjectFirewallVip6-SslPfs"); ok {
			if err = d.Set("ssl_pfs", vv); err != nil {
				return fmt.Errorf("Error reading ssl_pfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_pfs: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenObjectFirewallVip6SslSendEmptyFrags(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-send-empty-frags"], "ObjectFirewallVip6-SslSendEmptyFrags"); ok {
			if err = d.Set("ssl_send_empty_frags", vv); err != nil {
				return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("ssl_server_algorithm", flattenObjectFirewallVip6SslServerAlgorithm(o["ssl-server-algorithm"], d, "ssl_server_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-algorithm"], "ObjectFirewallVip6-SslServerAlgorithm"); ok {
			if err = d.Set("ssl_server_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_server_cipher_suites", flattenObjectFirewallVip6SslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-server-cipher-suites"], "ObjectFirewallVip6-SslServerCipherSuites"); ok {
				if err = d.Set("ssl_server_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server_cipher_suites"); ok {
			if err = d.Set("ssl_server_cipher_suites", flattenObjectFirewallVip6SslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-server-cipher-suites"], "ObjectFirewallVip6-SslServerCipherSuites"); ok {
					if err = d.Set("ssl_server_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_server_max_version", flattenObjectFirewallVip6SslServerMaxVersion(o["ssl-server-max-version"], d, "ssl_server_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-max-version"], "ObjectFirewallVip6-SslServerMaxVersion"); ok {
			if err = d.Set("ssl_server_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_min_version", flattenObjectFirewallVip6SslServerMinVersion(o["ssl-server-min-version"], d, "ssl_server_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-min-version"], "ObjectFirewallVip6-SslServerMinVersion"); ok {
			if err = d.Set("ssl_server_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_renegotiation", flattenObjectFirewallVip6SslServerRenegotiation(o["ssl-server-renegotiation"], d, "ssl_server_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-renegotiation"], "ObjectFirewallVip6-SslServerRenegotiation"); ok {
			if err = d.Set("ssl_server_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_max", flattenObjectFirewallVip6SslServerSessionStateMax(o["ssl-server-session-state-max"], d, "ssl_server_session_state_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-max"], "ObjectFirewallVip6-SslServerSessionStateMax"); ok {
			if err = d.Set("ssl_server_session_state_max", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_timeout", flattenObjectFirewallVip6SslServerSessionStateTimeout(o["ssl-server-session-state-timeout"], d, "ssl_server_session_state_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-timeout"], "ObjectFirewallVip6-SslServerSessionStateTimeout"); ok {
			if err = d.Set("ssl_server_session_state_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_type", flattenObjectFirewallVip6SslServerSessionStateType(o["ssl-server-session-state-type"], d, "ssl_server_session_state_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-session-state-type"], "ObjectFirewallVip6-SslServerSessionStateType"); ok {
			if err = d.Set("ssl_server_session_state_type", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallVip6Type(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallVip6-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectFirewallVip6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectFirewallVip6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("weblogic_server", flattenObjectFirewallVip6WeblogicServer(o["weblogic-server"], d, "weblogic_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["weblogic-server"], "ObjectFirewallVip6-WeblogicServer"); ok {
			if err = d.Set("weblogic_server", vv); err != nil {
				return fmt.Errorf("Error reading weblogic_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weblogic_server: %v", err)
		}
	}

	if err = d.Set("websphere_server", flattenObjectFirewallVip6WebsphereServer(o["websphere-server"], d, "websphere_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["websphere-server"], "ObjectFirewallVip6-WebsphereServer"); ok {
			if err = d.Set("websphere_server", vv); err != nil {
				return fmt.Errorf("Error reading websphere_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading websphere_server: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallVip6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallVip6AddNat64Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6ArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Color(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Comment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandObjectFirewallVip6DynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat64_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-nat64-route"], _ = expandObjectFirewallVip6DynamicMappingAddNat64Route(d, i["add_nat64_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-reply"], _ = expandObjectFirewallVip6DynamicMappingArpReply(d, i["arp_reply"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color"], _ = expandObjectFirewallVip6DynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectFirewallVip6DynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embedded_ipv4_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["embedded-ipv4-address"], _ = expandObjectFirewallVip6DynamicMappingEmbeddedIpv4Address(d, i["embedded_ipv4_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extip"], _ = expandObjectFirewallVip6DynamicMappingExtip(d, i["extip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extport"], _ = expandObjectFirewallVip6DynamicMappingExtport(d, i["extport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-age"], _ = expandObjectFirewallVip6DynamicMappingHttpCookieAge(d, i["http_cookie_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain"], _ = expandObjectFirewallVip6DynamicMappingHttpCookieDomain(d, i["http_cookie_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-domain-from-host"], _ = expandObjectFirewallVip6DynamicMappingHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-generation"], _ = expandObjectFirewallVip6DynamicMappingHttpCookieGeneration(d, i["http_cookie_generation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-path"], _ = expandObjectFirewallVip6DynamicMappingHttpCookiePath(d, i["http_cookie_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-cookie-share"], _ = expandObjectFirewallVip6DynamicMappingHttpCookieShare(d, i["http_cookie_share"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-ip-header"], _ = expandObjectFirewallVip6DynamicMappingHttpIpHeader(d, i["http_ip_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_ip_header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-ip-header-name"], _ = expandObjectFirewallVip6DynamicMappingHttpIpHeaderName(d, i["http_ip_header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_multiplex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-multiplex"], _ = expandObjectFirewallVip6DynamicMappingHttpMultiplex(d, i["http_multiplex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_redirect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-redirect"], _ = expandObjectFirewallVip6DynamicMappingHttpRedirect(d, i["http_redirect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-cookie-secure"], _ = expandObjectFirewallVip6DynamicMappingHttpsCookieSecure(d, i["https_cookie_secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallVip6DynamicMappingId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_mappedip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv4-mappedip"], _ = expandObjectFirewallVip6DynamicMappingIpv4Mappedip(d, i["ipv4_mappedip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv4-mappedport"], _ = expandObjectFirewallVip6DynamicMappingIpv4Mappedport(d, i["ipv4_mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandObjectFirewallVip6DynamicMappingLdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedip"], _ = expandObjectFirewallVip6DynamicMappingMappedip(d, i["mappedip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandObjectFirewallVip6DynamicMappingMappedport(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_embryonic_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-embryonic-connections"], _ = expandObjectFirewallVip6DynamicMappingMaxEmbryonicConnections(d, i["max_embryonic_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandObjectFirewallVip6DynamicMappingMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat_source_vip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat-source-vip"], _ = expandObjectFirewallVip6DynamicMappingNatSourceVip(d, i["nat_source_vip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat64"], _ = expandObjectFirewallVip6DynamicMappingNat64(d, i["nat64"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat66"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat66"], _ = expandObjectFirewallVip6DynamicMappingNat66(d, i["nat66"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ndp_reply"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ndp-reply"], _ = expandObjectFirewallVip6DynamicMappingNdpReply(d, i["ndp_reply"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "outlook_web_access"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["outlook-web-access"], _ = expandObjectFirewallVip6DynamicMappingOutlookWebAccess(d, i["outlook_web_access"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["persistence"], _ = expandObjectFirewallVip6DynamicMappingPersistence(d, i["persistence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portforward"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portforward"], _ = expandObjectFirewallVip6DynamicMappingPortforward(d, i["portforward"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectFirewallVip6DynamicMappingProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallVip6DynamicMappingRealservers(d, i["realservers"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["realservers"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-type"], _ = expandObjectFirewallVip6DynamicMappingServerType(d, i["server_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-filter"], _ = expandObjectFirewallVip6DynamicMappingSrcFilter(d, i["src_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_accept_ffdhe_groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-accept-ffdhe-groups"], _ = expandObjectFirewallVip6DynamicMappingSslAcceptFfdheGroups(d, i["ssl_accept_ffdhe_groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-algorithm"], _ = expandObjectFirewallVip6DynamicMappingSslAlgorithm(d, i["ssl_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-certificate"], _ = expandObjectFirewallVip6DynamicMappingSslCertificate(d, i["ssl_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallVip6DynamicMappingSslCipherSuites(d, i["ssl_cipher_suites"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ssl-cipher-suites"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_fallback"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-fallback"], _ = expandObjectFirewallVip6DynamicMappingSslClientFallback(d, i["ssl_client_fallback"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_rekey_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-rekey-count"], _ = expandObjectFirewallVip6DynamicMappingSslClientRekeyCount(d, i["ssl_client_rekey_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-renegotiation"], _ = expandObjectFirewallVip6DynamicMappingSslClientRenegotiation(d, i["ssl_client_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-max"], _ = expandObjectFirewallVip6DynamicMappingSslClientSessionStateMax(d, i["ssl_client_session_state_max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-timeout"], _ = expandObjectFirewallVip6DynamicMappingSslClientSessionStateTimeout(d, i["ssl_client_session_state_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_client_session_state_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-client-session-state-type"], _ = expandObjectFirewallVip6DynamicMappingSslClientSessionStateType(d, i["ssl_client_session_state_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-dh-bits"], _ = expandObjectFirewallVip6DynamicMappingSslDhBits(d, i["ssl_dh_bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp"], _ = expandObjectFirewallVip6DynamicMappingSslHpkp(d, i["ssl_hpkp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-age"], _ = expandObjectFirewallVip6DynamicMappingSslHpkpAge(d, i["ssl_hpkp_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_backup"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-backup"], _ = expandObjectFirewallVip6DynamicMappingSslHpkpBackup(d, i["ssl_hpkp_backup"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_include_subdomains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-include-subdomains"], _ = expandObjectFirewallVip6DynamicMappingSslHpkpIncludeSubdomains(d, i["ssl_hpkp_include_subdomains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_primary"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-primary"], _ = expandObjectFirewallVip6DynamicMappingSslHpkpPrimary(d, i["ssl_hpkp_primary"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hpkp_report_uri"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hpkp-report-uri"], _ = expandObjectFirewallVip6DynamicMappingSslHpkpReportUri(d, i["ssl_hpkp_report_uri"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts"], _ = expandObjectFirewallVip6DynamicMappingSslHsts(d, i["ssl_hsts"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_age"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts-age"], _ = expandObjectFirewallVip6DynamicMappingSslHstsAge(d, i["ssl_hsts_age"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_hsts_include_subdomains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-hsts-include-subdomains"], _ = expandObjectFirewallVip6DynamicMappingSslHstsIncludeSubdomains(d, i["ssl_hsts_include_subdomains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_location_conversion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-http-location-conversion"], _ = expandObjectFirewallVip6DynamicMappingSslHttpLocationConversion(d, i["ssl_http_location_conversion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_http_match_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-http-match-host"], _ = expandObjectFirewallVip6DynamicMappingSslHttpMatchHost(d, i["ssl_http_match_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandObjectFirewallVip6DynamicMappingSslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-version"], _ = expandObjectFirewallVip6DynamicMappingSslMinVersion(d, i["ssl_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-mode"], _ = expandObjectFirewallVip6DynamicMappingSslMode(d, i["ssl_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_pfs"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-pfs"], _ = expandObjectFirewallVip6DynamicMappingSslPfs(d, i["ssl_pfs"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_send_empty_frags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-send-empty-frags"], _ = expandObjectFirewallVip6DynamicMappingSslSendEmptyFrags(d, i["ssl_send_empty_frags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-algorithm"], _ = expandObjectFirewallVip6DynamicMappingSslServerAlgorithm(d, i["ssl_server_algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-max-version"], _ = expandObjectFirewallVip6DynamicMappingSslServerMaxVersion(d, i["ssl_server_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_min_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-min-version"], _ = expandObjectFirewallVip6DynamicMappingSslServerMinVersion(d, i["ssl_server_min_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_renegotiation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-renegotiation"], _ = expandObjectFirewallVip6DynamicMappingSslServerRenegotiation(d, i["ssl_server_renegotiation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-max"], _ = expandObjectFirewallVip6DynamicMappingSslServerSessionStateMax(d, i["ssl_server_session_state_max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-timeout"], _ = expandObjectFirewallVip6DynamicMappingSslServerSessionStateTimeout(d, i["ssl_server_session_state_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_session_state_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-session-state-type"], _ = expandObjectFirewallVip6DynamicMappingSslServerSessionStateType(d, i["ssl_server_session_state_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallVip6DynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandObjectFirewallVip6DynamicMappingUuid(d, i["uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weblogic_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weblogic-server"], _ = expandObjectFirewallVip6DynamicMappingWeblogicServer(d, i["weblogic_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "websphere_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["websphere-server"], _ = expandObjectFirewallVip6DynamicMappingWebsphereServer(d, i["websphere_server"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVip6DynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallVip6DynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectFirewallVip6DynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVip6DynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingAddNat64Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingEmbeddedIpv4Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingExtip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingExtport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpIpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpIpHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingIpv4Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingIpv4Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingMappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingMaxEmbryonicConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingNatSourceVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingNat66(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingNdpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingOutlookWebAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingPersistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingPortforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-ip"], _ = expandObjectFirewallVip6DynamicMappingRealserversClientIp(d, i["client_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["healthcheck"], _ = expandObjectFirewallVip6DynamicMappingRealserversHealthcheck(d, i["healthcheck"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandObjectFirewallVip6DynamicMappingRealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandObjectFirewallVip6DynamicMappingRealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallVip6DynamicMappingRealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFirewallVip6DynamicMappingRealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandObjectFirewallVip6DynamicMappingRealserversMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandObjectFirewallVip6DynamicMappingRealserversMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectFirewallVip6DynamicMappingRealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFirewallVip6DynamicMappingRealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandObjectFirewallVip6DynamicMappingRealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectFirewallVip6DynamicMappingRealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingRealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVip6DynamicMappingSslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandObjectFirewallVip6DynamicMappingSslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFirewallVip6DynamicMappingSslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandObjectFirewallVip6DynamicMappingSslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVip6DynamicMappingSslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVip6DynamicMappingSslClientFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHpkp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHpkpAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHpkpBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHsts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHstsAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslPfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslServerMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslServerRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingSslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingWeblogicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6DynamicMappingWebsphereServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6EmbeddedIpv4Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Extip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Extport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpCookieDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpCookiePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpCookieShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpIpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpIpHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpMultiplex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6HttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Ipv4Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Ipv4Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6LdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectFirewallVip6Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6MaxEmbryonicConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Monitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6NatSourceVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Nat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Nat66(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6NdpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6OutlookWebAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Persistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Portforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Protocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Realservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-ip"], _ = expandObjectFirewallVip6RealserversClientIp(d, i["client_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["healthcheck"], _ = expandObjectFirewallVip6RealserversHealthcheck(d, i["healthcheck"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandObjectFirewallVip6RealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-host"], _ = expandObjectFirewallVip6RealserversHttpHost(d, i["http_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallVip6RealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFirewallVip6RealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandObjectFirewallVip6RealserversMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandObjectFirewallVip6RealserversMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectFirewallVip6RealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFirewallVip6RealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["translate-host"], _ = expandObjectFirewallVip6RealserversTranslateHost(d, i["translate_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectFirewallVip6RealserversWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVip6RealserversClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversHttpHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6RealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6ServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVip6SslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandObjectFirewallVip6SslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFirewallVip6SslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandObjectFirewallVip6SslCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVip6SslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVip6SslClientFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHpkp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHpkpAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHpkpBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHsts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHstsAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslPfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslServerCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandObjectFirewallVip6SslServerCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFirewallVip6SslServerCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["versions"], _ = expandObjectFirewallVip6SslServerCipherSuitesVersions(d, i["versions"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVip6SslServerCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslServerCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslServerCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVip6SslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslServerMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslServerRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6SslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6WeblogicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip6WebsphereServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallVip6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_nat64_route"); ok || d.HasChange("add_nat64_route") {
		t, err := expandObjectFirewallVip6AddNat64Route(d, v, "add_nat64_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat64-route"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok || d.HasChange("arp_reply") {
		t, err := expandObjectFirewallVip6ArpReply(d, v, "arp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectFirewallVip6Color(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectFirewallVip6Comment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectFirewallVip6DynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("embedded_ipv4_address"); ok || d.HasChange("embedded_ipv4_address") {
		t, err := expandObjectFirewallVip6EmbeddedIpv4Address(d, v, "embedded_ipv4_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["embedded-ipv4-address"] = t
		}
	}

	if v, ok := d.GetOk("extip"); ok || d.HasChange("extip") {
		t, err := expandObjectFirewallVip6Extip(d, v, "extip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extip"] = t
		}
	}

	if v, ok := d.GetOk("extport"); ok || d.HasChange("extport") {
		t, err := expandObjectFirewallVip6Extport(d, v, "extport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extport"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_age"); ok || d.HasChange("http_cookie_age") {
		t, err := expandObjectFirewallVip6HttpCookieAge(d, v, "http_cookie_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain"); ok || d.HasChange("http_cookie_domain") {
		t, err := expandObjectFirewallVip6HttpCookieDomain(d, v, "http_cookie_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain_from_host"); ok || d.HasChange("http_cookie_domain_from_host") {
		t, err := expandObjectFirewallVip6HttpCookieDomainFromHost(d, v, "http_cookie_domain_from_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain-from-host"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_generation"); ok || d.HasChange("http_cookie_generation") {
		t, err := expandObjectFirewallVip6HttpCookieGeneration(d, v, "http_cookie_generation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-generation"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_path"); ok || d.HasChange("http_cookie_path") {
		t, err := expandObjectFirewallVip6HttpCookiePath(d, v, "http_cookie_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-path"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_share"); ok || d.HasChange("http_cookie_share") {
		t, err := expandObjectFirewallVip6HttpCookieShare(d, v, "http_cookie_share")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-share"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header"); ok || d.HasChange("http_ip_header") {
		t, err := expandObjectFirewallVip6HttpIpHeader(d, v, "http_ip_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header_name"); ok || d.HasChange("http_ip_header_name") {
		t, err := expandObjectFirewallVip6HttpIpHeaderName(d, v, "http_ip_header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header-name"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex"); ok || d.HasChange("http_multiplex") {
		t, err := expandObjectFirewallVip6HttpMultiplex(d, v, "http_multiplex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex"] = t
		}
	}

	if v, ok := d.GetOk("http_redirect"); ok || d.HasChange("http_redirect") {
		t, err := expandObjectFirewallVip6HttpRedirect(d, v, "http_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-redirect"] = t
		}
	}

	if v, ok := d.GetOk("https_cookie_secure"); ok || d.HasChange("https_cookie_secure") {
		t, err := expandObjectFirewallVip6HttpsCookieSecure(d, v, "https_cookie_secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-cookie-secure"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallVip6Id(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_mappedip"); ok || d.HasChange("ipv4_mappedip") {
		t, err := expandObjectFirewallVip6Ipv4Mappedip(d, v, "ipv4_mappedip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-mappedip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_mappedport"); ok || d.HasChange("ipv4_mappedport") {
		t, err := expandObjectFirewallVip6Ipv4Mappedport(d, v, "ipv4_mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-mappedport"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok || d.HasChange("ldb_method") {
		t, err := expandObjectFirewallVip6LdbMethod(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("mappedip"); ok || d.HasChange("mappedip") {
		t, err := expandObjectFirewallVip6Mappedip(d, v, "mappedip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedip"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok || d.HasChange("mappedport") {
		t, err := expandObjectFirewallVip6Mappedport(d, v, "mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("max_embryonic_connections"); ok || d.HasChange("max_embryonic_connections") {
		t, err := expandObjectFirewallVip6MaxEmbryonicConnections(d, v, "max_embryonic_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-embryonic-connections"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandObjectFirewallVip6Monitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallVip6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat_source_vip"); ok || d.HasChange("nat_source_vip") {
		t, err := expandObjectFirewallVip6NatSourceVip(d, v, "nat_source_vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-source-vip"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok || d.HasChange("nat64") {
		t, err := expandObjectFirewallVip6Nat64(d, v, "nat64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("nat66"); ok || d.HasChange("nat66") {
		t, err := expandObjectFirewallVip6Nat66(d, v, "nat66")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat66"] = t
		}
	}

	if v, ok := d.GetOk("ndp_reply"); ok || d.HasChange("ndp_reply") {
		t, err := expandObjectFirewallVip6NdpReply(d, v, "ndp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ndp-reply"] = t
		}
	}

	if v, ok := d.GetOk("outlook_web_access"); ok || d.HasChange("outlook_web_access") {
		t, err := expandObjectFirewallVip6OutlookWebAccess(d, v, "outlook_web_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outlook-web-access"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok || d.HasChange("persistence") {
		t, err := expandObjectFirewallVip6Persistence(d, v, "persistence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("portforward"); ok || d.HasChange("portforward") {
		t, err := expandObjectFirewallVip6Portforward(d, v, "portforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portforward"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectFirewallVip6Protocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok || d.HasChange("realservers") {
		t, err := expandObjectFirewallVip6Realservers(d, v, "realservers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandObjectFirewallVip6ServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("src_filter"); ok || d.HasChange("src_filter") {
		t, err := expandObjectFirewallVip6SrcFilter(d, v, "src_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-filter"] = t
		}
	}

	if v, ok := d.GetOk("ssl_accept_ffdhe_groups"); ok || d.HasChange("ssl_accept_ffdhe_groups") {
		t, err := expandObjectFirewallVip6SslAcceptFfdheGroups(d, v, "ssl_accept_ffdhe_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-accept-ffdhe-groups"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandObjectFirewallVip6SslAlgorithm(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandObjectFirewallVip6SslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandObjectFirewallVip6SslCipherSuites(d, v, "ssl_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_fallback"); ok || d.HasChange("ssl_client_fallback") {
		t, err := expandObjectFirewallVip6SslClientFallback(d, v, "ssl_client_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-fallback"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_rekey_count"); ok || d.HasChange("ssl_client_rekey_count") {
		t, err := expandObjectFirewallVip6SslClientRekeyCount(d, v, "ssl_client_rekey_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-rekey-count"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok || d.HasChange("ssl_client_renegotiation") {
		t, err := expandObjectFirewallVip6SslClientRenegotiation(d, v, "ssl_client_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_max"); ok || d.HasChange("ssl_client_session_state_max") {
		t, err := expandObjectFirewallVip6SslClientSessionStateMax(d, v, "ssl_client_session_state_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_timeout"); ok || d.HasChange("ssl_client_session_state_timeout") {
		t, err := expandObjectFirewallVip6SslClientSessionStateTimeout(d, v, "ssl_client_session_state_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_type"); ok || d.HasChange("ssl_client_session_state_type") {
		t, err := expandObjectFirewallVip6SslClientSessionStateType(d, v, "ssl_client_session_state_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandObjectFirewallVip6SslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp"); ok || d.HasChange("ssl_hpkp") {
		t, err := expandObjectFirewallVip6SslHpkp(d, v, "ssl_hpkp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_age"); ok || d.HasChange("ssl_hpkp_age") {
		t, err := expandObjectFirewallVip6SslHpkpAge(d, v, "ssl_hpkp_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_backup"); ok || d.HasChange("ssl_hpkp_backup") {
		t, err := expandObjectFirewallVip6SslHpkpBackup(d, v, "ssl_hpkp_backup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-backup"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_include_subdomains"); ok || d.HasChange("ssl_hpkp_include_subdomains") {
		t, err := expandObjectFirewallVip6SslHpkpIncludeSubdomains(d, v, "ssl_hpkp_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_primary"); ok || d.HasChange("ssl_hpkp_primary") {
		t, err := expandObjectFirewallVip6SslHpkpPrimary(d, v, "ssl_hpkp_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-primary"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_report_uri"); ok || d.HasChange("ssl_hpkp_report_uri") {
		t, err := expandObjectFirewallVip6SslHpkpReportUri(d, v, "ssl_hpkp_report_uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-report-uri"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts"); ok || d.HasChange("ssl_hsts") {
		t, err := expandObjectFirewallVip6SslHsts(d, v, "ssl_hsts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_age"); ok || d.HasChange("ssl_hsts_age") {
		t, err := expandObjectFirewallVip6SslHstsAge(d, v, "ssl_hsts_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_include_subdomains"); ok || d.HasChange("ssl_hsts_include_subdomains") {
		t, err := expandObjectFirewallVip6SslHstsIncludeSubdomains(d, v, "ssl_hsts_include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_location_conversion"); ok || d.HasChange("ssl_http_location_conversion") {
		t, err := expandObjectFirewallVip6SslHttpLocationConversion(d, v, "ssl_http_location_conversion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-location-conversion"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_match_host"); ok || d.HasChange("ssl_http_match_host") {
		t, err := expandObjectFirewallVip6SslHttpMatchHost(d, v, "ssl_http_match_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-match-host"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandObjectFirewallVip6SslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandObjectFirewallVip6SslMinVersion(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok || d.HasChange("ssl_mode") {
		t, err := expandObjectFirewallVip6SslMode(d, v, "ssl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_pfs"); ok || d.HasChange("ssl_pfs") {
		t, err := expandObjectFirewallVip6SslPfs(d, v, "ssl_pfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-pfs"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok || d.HasChange("ssl_send_empty_frags") {
		t, err := expandObjectFirewallVip6SslSendEmptyFrags(d, v, "ssl_send_empty_frags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_algorithm"); ok || d.HasChange("ssl_server_algorithm") {
		t, err := expandObjectFirewallVip6SslServerAlgorithm(d, v, "ssl_server_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_cipher_suites"); ok || d.HasChange("ssl_server_cipher_suites") {
		t, err := expandObjectFirewallVip6SslServerCipherSuites(d, v, "ssl_server_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_max_version"); ok || d.HasChange("ssl_server_max_version") {
		t, err := expandObjectFirewallVip6SslServerMaxVersion(d, v, "ssl_server_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_min_version"); ok || d.HasChange("ssl_server_min_version") {
		t, err := expandObjectFirewallVip6SslServerMinVersion(d, v, "ssl_server_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_renegotiation"); ok || d.HasChange("ssl_server_renegotiation") {
		t, err := expandObjectFirewallVip6SslServerRenegotiation(d, v, "ssl_server_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_max"); ok || d.HasChange("ssl_server_session_state_max") {
		t, err := expandObjectFirewallVip6SslServerSessionStateMax(d, v, "ssl_server_session_state_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_timeout"); ok || d.HasChange("ssl_server_session_state_timeout") {
		t, err := expandObjectFirewallVip6SslServerSessionStateTimeout(d, v, "ssl_server_session_state_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_type"); ok || d.HasChange("ssl_server_session_state_type") {
		t, err := expandObjectFirewallVip6SslServerSessionStateType(d, v, "ssl_server_session_state_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallVip6Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectFirewallVip6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("weblogic_server"); ok || d.HasChange("weblogic_server") {
		t, err := expandObjectFirewallVip6WeblogicServer(d, v, "weblogic_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weblogic-server"] = t
		}
	}

	if v, ok := d.GetOk("websphere_server"); ok || d.HasChange("websphere_server") {
		t, err := expandObjectFirewallVip6WebsphereServer(d, v, "websphere_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["websphere-server"] = t
		}
	}

	return &obj, nil
}
