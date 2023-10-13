// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Virtual Access Points (VAPs).

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerVap() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerVapCreate,
		Read:   resourceObjectWirelessControllerVapRead,
		Update: resourceObjectWirelessControllerVapUpdate,
		Delete: resourceObjectWirelessControllerVapDelete,

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
			"_centmgmt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_dhcp_svr_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"_intf_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_intf_device_access_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"_intf_device_identification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_device_netscan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp_relay_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp_relay_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp_relay_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp6_relay_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"_intf_dhcp6_relay_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_dhcp6_relay_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_intf_ip6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"_intf_ip6_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_intf_listen_forticlient_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"_is_factory_setting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_control_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"additional_akms": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"acct_interim_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"address_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_group_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"antivirus_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_detection_engine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_dscp_marking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_report_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"atf_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth": &schema.Schema{
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
			"beacon_advertising": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"broadcast_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"broadcast_suppression": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bss_color_partial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bstm_disassociation_imminent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bstm_load_balancing_disassoc_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bstm_rssi_disassoc_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ac_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captive_portal_auth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"captive_portal_fw_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_address_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_macauth_radius_secret": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"captive_portal_macauth_radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captive_portal_radius_secret": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"captive_portal_radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captive_portal_session_timeout_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp_lease_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp_option43_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_circuit_id_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_remote_id_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_centmgmt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_dhcp_svr_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"_intf_allowaccess": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"_intf_device_access_list": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"_intf_device_identification": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_device_netscan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp_relay_ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp_relay_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp_relay_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp6_relay_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"_intf_dhcp6_relay_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_dhcp6_relay_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_intf_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"_intf_ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"_intf_ip6_allowaccess": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"_intf_listen_forticlient_connection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"_is_factory_setting": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"access_control_list": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"acct_interim_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"additional_akms": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"address_group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"address_group_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"alias": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"antivirus_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"application_detection_engine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"application_dscp_marking": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"application_list": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"application_report_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"atf_weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auth": &schema.Schema{
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
						"beacon_advertising": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"broadcast_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"broadcast_suppression": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"bss_color_partial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bstm_disassociation_imminent": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bstm_load_balancing_disassoc_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bstm_rssi_disassoc_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"captive_portal_ac_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"captive_portal_auth_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"captive_portal_fw_accounting": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"captive_portal_macauth_radius_secret": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"captive_portal_macauth_radius_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"captive_portal_radius_secret": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"captive_portal_radius_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"captive_portal_session_timeout_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"client_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dhcp_address_enforcement": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_lease_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dhcp_option43_insertion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp_option82_circuit_id_insertion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp_option82_insertion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp_option82_remote_id_insertion": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dynamic_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"eap_reauth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"eap_reauth_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"eapol_key_retries": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"encrypt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_fast_roaming": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_logout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_web": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_web_format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fast_bss_transition": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fast_roaming": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ft_mobility_domain": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ft_over_ds": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ft_r0_key_lifetime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"gas_comeback_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"gas_fragmentation_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"gtk_rekey": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gtk_rekey_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"high_efficiency": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hotspot20_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"igmp_snooping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"intra_vap_privacy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ips_sensor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6_rules": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"keyindex": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"l3_roaming": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"l3_roaming_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ldpc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_bridging": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_lan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_standalone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_standalone_dns": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_standalone_dns_ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"local_standalone_nat": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_switching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_auth_bypass": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_called_station_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_calling_station_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_case": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_filter_policy_other": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_password_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_username_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_clients_ap": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mbo": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mbo_cell_data_conn_pref": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"me_disable_thresh": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mesh_backhaul": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mpsk": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mpsk_concurrent_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mpsk_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mu_mimo": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"multicast_enhance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"multicast_rate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nac_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"neighbor_report_dual_band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"okc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"osen": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"owe_groups": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"owe_transition": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"owe_transition_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"passphrase": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"pmf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pmf_assoc_comeback_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pmf_sa_query_retry_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"port_macauth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port_macauth_reauth_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"port_macauth_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"portal_message_override_group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"portal_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"primary_wag_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"probe_resp_suppression": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"probe_resp_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ptk_rekey": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ptk_rekey_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"qos_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radio_2g_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radio_5g_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radio_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radius_mac_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radius_mac_auth_block_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"radius_mac_auth_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"radius_mac_auth_usergroups": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"radius_mac_mpsk_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radius_mac_mpsk_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"radius_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rates_11a": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11ac_mcs_map": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rates_11ac_ss12": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11ac_ss34": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11ax_mcs_map": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rates_11ax_ss12": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11ax_ss34": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11bg": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11n_ss12": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rates_11n_ss34": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sae_groups": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sae_h2e_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sae_password": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sae_pk": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sae_private_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"scan_botnet_connections": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"schedule": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"secondary_wag_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"security": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"security_exempt_list": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"security_obsolete_option": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"security_redirect_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"selected_usergroups": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"split_tunneling": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sticky_client_remove": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sticky_client_threshold_2g": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sticky_client_threshold_5g": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sticky_client_threshold_6g": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"target_wake_time": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tkip_counter_measure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_echo_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tunnel_fallback_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"usergroup": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"utm_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"utm_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"utm_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_auto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan_pooling": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlanid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"voice_enterprise": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"webfilter_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"eap_reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_reauth_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_key_retries": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_fast_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_logout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_web": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_web_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fast_bss_transition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fast_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ft_mobility_domain": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ft_over_ds": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ft_r0_key_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gas_comeback_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gas_fragmentation_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gtk_rekey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtk_rekey_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"high_efficiency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hotspot20_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"igmp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intra_vap_privacy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_rules": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"keyindex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"l3_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l3_roaming_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldpc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_bridging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_lan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_standalone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_standalone_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_standalone_dns_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"local_standalone_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_called_station_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_calling_station_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_filter_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_filter_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mac_filter_policy_other": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_password_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_username_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_clients_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbo_cell_data_conn_pref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"me_disable_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mesh_backhaul": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpsk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpsk_concurrent_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mpsk_key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"concurrent_clients": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mpsk_schedules": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"passphrase": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mpsk_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mu_mimo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_enhance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_rate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"neighbor_report_dual_band": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"okc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"osen": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owe_groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"owe_transition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owe_transition_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"passphrase": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"pmf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pmf_assoc_comeback_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pmf_sa_query_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port_macauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_macauth_reauth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port_macauth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"portal_message_override_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"portal_message_overrides": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_disclaimer_page": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_login_failed_page": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_login_page": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_reject_page": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"portal_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_wag_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"probe_resp_suppression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probe_resp_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ptk_rekey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ptk_rekey_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"qos_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quarantine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_2g_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_5g_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_auth_block_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"radius_mac_auth_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"radius_mac_auth_usergroups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"radius_mac_mpsk_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_mpsk_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11a": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11ac_mcs_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11ac_ss12": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11ac_ss34": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11ax_mcs_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11ax_ss12": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11ax_ss34": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11bg": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11n_ss12": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rates_11n_ss34": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sae_groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sae_h2e_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sae_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sae_pk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sae_private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secondary_wag_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_exempt_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_obsolete_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"selected_usergroups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky_client_remove": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky_client_threshold_2g": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky_client_threshold_5g": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky_client_threshold_6g": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_wake_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tkip_counter_measure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_echo_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tunnel_fallback_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"usergroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"utm_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utm_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_auto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"vlan_pool": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_wtp_group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"vlan_pooling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"voice_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
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

func resourceObjectWirelessControllerVapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectWirelessControllerVap(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerVap resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerVap(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerVap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerVapRead(d, m)
}

func resourceObjectWirelessControllerVapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerVap(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVap resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerVap(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerVapRead(d, m)
}

func resourceObjectWirelessControllerVapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerVap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerVap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerVapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerVap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerVap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVap resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerVapCentmgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDhcpSvrId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntfAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapIntfDeviceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntfDeviceIdentification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntfDeviceNetscan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntfDhcpRelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapIntfDhcpRelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntfDhcpRelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntfDhcp6RelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntfDhcp6RelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntfDhcp6RelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntfIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenObjectWirelessControllerVapIntfIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntfIp6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapIntfListenForticlientConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIsFactorySetting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapAccessControlList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapAdditionalAkms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapAddressGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapAddressGroupPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapAntivirusProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapApplicationDetectionEngine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapApplicationDscpMarking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapApplicationReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapAtfWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapAuthPortalAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapBeaconAdvertising(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapBroadcastSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapBroadcastSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapBssColorPartial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapBstmDisassociationImminent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapBstmLoadBalancingDisassocTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapBstmRssiDisassocTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapCaptivePortalAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapCaptivePortalAuthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapCaptivePortalFwAccounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDhcpAddressEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapCaptivePortalMacauthRadiusSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapCaptivePortalMacauthRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapCaptivePortalRadiusSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapCaptivePortalRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapCaptivePortalSessionTimeoutInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDhcpLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDhcpOption43Insertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDhcpOption82CircuitIdInsertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDhcpOption82Insertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDhcpOption82RemoteIdInsertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_centmgmt"
		if _, ok := i["_centmgmt"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingCentmgmt(i["_centmgmt"], d, pre_append)
			tmp["_centmgmt"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Centmgmt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_dhcp_svr_id"
		if _, ok := i["_dhcp_svr_id"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingDhcpSvrId(i["_dhcp_svr_id"], d, pre_append)
			tmp["_dhcp_svr_id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-DhcpSvrId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_allowaccess"
		if _, ok := i["_intf_allowaccess"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfAllowaccess(i["_intf_allowaccess"], d, pre_append)
			tmp["_intf_allowaccess"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfAllowaccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_access_list"
		if _, ok := i["_intf_device-access-list"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfDeviceAccessList(i["_intf_device-access-list"], d, pre_append)
			tmp["_intf_device_access_list"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfDeviceAccessList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_identification"
		if _, ok := i["_intf_device-identification"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfDeviceIdentification(i["_intf_device-identification"], d, pre_append)
			tmp["_intf_device_identification"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfDeviceIdentification")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_netscan"
		if _, ok := i["_intf_device-netscan"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfDeviceNetscan(i["_intf_device-netscan"], d, pre_append)
			tmp["_intf_device_netscan"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfDeviceNetscan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_ip"
		if _, ok := i["_intf_dhcp-relay-ip"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayIp(i["_intf_dhcp-relay-ip"], d, pre_append)
			tmp["_intf_dhcp_relay_ip"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfDhcpRelayIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_service"
		if _, ok := i["_intf_dhcp-relay-service"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayService(i["_intf_dhcp-relay-service"], d, pre_append)
			tmp["_intf_dhcp_relay_service"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfDhcpRelayService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_type"
		if _, ok := i["_intf_dhcp-relay-type"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayType(i["_intf_dhcp-relay-type"], d, pre_append)
			tmp["_intf_dhcp_relay_type"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfDhcpRelayType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_ip"
		if _, ok := i["_intf_dhcp6-relay-ip"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayIp(i["_intf_dhcp6-relay-ip"], d, pre_append)
			tmp["_intf_dhcp6_relay_ip"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfDhcp6RelayIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_service"
		if _, ok := i["_intf_dhcp6-relay-service"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayService(i["_intf_dhcp6-relay-service"], d, pre_append)
			tmp["_intf_dhcp6_relay_service"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfDhcp6RelayService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_type"
		if _, ok := i["_intf_dhcp6-relay-type"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayType(i["_intf_dhcp6-relay-type"], d, pre_append)
			tmp["_intf_dhcp6_relay_type"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfDhcp6RelayType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip"
		if _, ok := i["_intf_ip"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfIp(i["_intf_ip"], d, pre_append)
			tmp["_intf_ip"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip6_address"
		if _, ok := i["_intf_ip6-address"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfIp6Address(i["_intf_ip6-address"], d, pre_append)
			tmp["_intf_ip6_address"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfIp6Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip6_allowaccess"
		if _, ok := i["_intf_ip6-allowaccess"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfIp6Allowaccess(i["_intf_ip6-allowaccess"], d, pre_append)
			tmp["_intf_ip6_allowaccess"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfIp6Allowaccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_listen_forticlient_connection"
		if _, ok := i["_intf_listen-forticlient-connection"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntfListenForticlientConnection(i["_intf_listen-forticlient-connection"], d, pre_append)
			tmp["_intf_listen_forticlient_connection"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntfListenForticlientConnection")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_is_factory_setting"
		if _, ok := i["_is_factory_setting"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIsFactorySetting(i["_is_factory_setting"], d, pre_append)
			tmp["_is_factory_setting"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IsFactorySetting")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_control_list"
		if _, ok := i["access-control-list"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAccessControlList(i["access-control-list"], d, pre_append)
			tmp["access_control_list"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-AccessControlList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_interim_interval"
		if _, ok := i["acct-interim-interval"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAcctInterimInterval(i["acct-interim-interval"], d, pre_append)
			tmp["acct_interim_interval"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-AcctInterimInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_akms"
		if _, ok := i["additional-akms"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAdditionalAkms(i["additional-akms"], d, pre_append)
			tmp["additional_akms"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-AdditionalAkms")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address_group"
		if _, ok := i["address-group"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAddressGroup(i["address-group"], d, pre_append)
			tmp["address_group"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-AddressGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address_group_policy"
		if _, ok := i["address-group-policy"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAddressGroupPolicy(i["address-group-policy"], d, pre_append)
			tmp["address_group_policy"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-AddressGroupPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alias"
		if _, ok := i["alias"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAlias(i["alias"], d, pre_append)
			tmp["alias"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Alias")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antivirus_profile"
		if _, ok := i["antivirus-profile"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAntivirusProfile(i["antivirus-profile"], d, pre_append)
			tmp["antivirus_profile"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-AntivirusProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_detection_engine"
		if _, ok := i["application-detection-engine"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingApplicationDetectionEngine(i["application-detection-engine"], d, pre_append)
			tmp["application_detection_engine"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ApplicationDetectionEngine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_dscp_marking"
		if _, ok := i["application-dscp-marking"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingApplicationDscpMarking(i["application-dscp-marking"], d, pre_append)
			tmp["application_dscp_marking"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ApplicationDscpMarking")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_list"
		if _, ok := i["application-list"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingApplicationList(i["application-list"], d, pre_append)
			tmp["application_list"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ApplicationList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_report_intv"
		if _, ok := i["application-report-intv"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingApplicationReportIntv(i["application-report-intv"], d, pre_append)
			tmp["application_report_intv"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ApplicationReportIntv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "atf_weight"
		if _, ok := i["atf-weight"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAtfWeight(i["atf-weight"], d, pre_append)
			tmp["atf_weight"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-AtfWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth"
		if _, ok := i["auth"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAuth(i["auth"], d, pre_append)
			tmp["auth"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Auth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_cert"
		if _, ok := i["auth-cert"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAuthCert(i["auth-cert"], d, pre_append)
			tmp["auth_cert"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-AuthCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_portal_addr"
		if _, ok := i["auth-portal-addr"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingAuthPortalAddr(i["auth-portal-addr"], d, pre_append)
			tmp["auth_portal_addr"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-AuthPortalAddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "beacon_advertising"
		if _, ok := i["beacon-advertising"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingBeaconAdvertising(i["beacon-advertising"], d, pre_append)
			tmp["beacon_advertising"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-BeaconAdvertising")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "broadcast_ssid"
		if _, ok := i["broadcast-ssid"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingBroadcastSsid(i["broadcast-ssid"], d, pre_append)
			tmp["broadcast_ssid"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-BroadcastSsid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "broadcast_suppression"
		if _, ok := i["broadcast-suppression"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingBroadcastSuppression(i["broadcast-suppression"], d, pre_append)
			tmp["broadcast_suppression"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-BroadcastSuppression")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bss_color_partial"
		if _, ok := i["bss-color-partial"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingBssColorPartial(i["bss-color-partial"], d, pre_append)
			tmp["bss_color_partial"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-BssColorPartial")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_disassociation_imminent"
		if _, ok := i["bstm-disassociation-imminent"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingBstmDisassociationImminent(i["bstm-disassociation-imminent"], d, pre_append)
			tmp["bstm_disassociation_imminent"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-BstmDisassociationImminent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_load_balancing_disassoc_timer"
		if _, ok := i["bstm-load-balancing-disassoc-timer"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer(i["bstm-load-balancing-disassoc-timer"], d, pre_append)
			tmp["bstm_load_balancing_disassoc_timer"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-BstmLoadBalancingDisassocTimer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_rssi_disassoc_timer"
		if _, ok := i["bstm-rssi-disassoc-timer"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingBstmRssiDisassocTimer(i["bstm-rssi-disassoc-timer"], d, pre_append)
			tmp["bstm_rssi_disassoc_timer"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-BstmRssiDisassocTimer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_ac_name"
		if _, ok := i["captive-portal-ac-name"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingCaptivePortalAcName(i["captive-portal-ac-name"], d, pre_append)
			tmp["captive_portal_ac_name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-CaptivePortalAcName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_auth_timeout"
		if _, ok := i["captive-portal-auth-timeout"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout(i["captive-portal-auth-timeout"], d, pre_append)
			tmp["captive_portal_auth_timeout"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-CaptivePortalAuthTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_fw_accounting"
		if _, ok := i["captive-portal-fw-accounting"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingCaptivePortalFwAccounting(i["captive-portal-fw-accounting"], d, pre_append)
			tmp["captive_portal_fw_accounting"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-CaptivePortalFwAccounting")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_macauth_radius_secret"
		if _, ok := i["captive-portal-macauth-radius-secret"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusSecret(i["captive-portal-macauth-radius-secret"], d, pre_append)
			tmp["captive_portal_macauth_radius_secret"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-CaptivePortalMacauthRadiusSecret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_macauth_radius_server"
		if _, ok := i["captive-portal-macauth-radius-server"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer(i["captive-portal-macauth-radius-server"], d, pre_append)
			tmp["captive_portal_macauth_radius_server"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-CaptivePortalMacauthRadiusServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_radius_secret"
		if _, ok := i["captive-portal-radius-secret"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusSecret(i["captive-portal-radius-secret"], d, pre_append)
			tmp["captive_portal_radius_secret"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-CaptivePortalRadiusSecret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_radius_server"
		if _, ok := i["captive-portal-radius-server"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusServer(i["captive-portal-radius-server"], d, pre_append)
			tmp["captive_portal_radius_server"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-CaptivePortalRadiusServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_session_timeout_interval"
		if _, ok := i["captive-portal-session-timeout-interval"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval(i["captive-portal-session-timeout-interval"], d, pre_append)
			tmp["captive_portal_session_timeout_interval"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-CaptivePortalSessionTimeoutInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_count"
		if _, ok := i["client-count"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingClientCount(i["client-count"], d, pre_append)
			tmp["client_count"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ClientCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_address_enforcement"
		if _, ok := i["dhcp-address-enforcement"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingDhcpAddressEnforcement(i["dhcp-address-enforcement"], d, pre_append)
			tmp["dhcp_address_enforcement"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-DhcpAddressEnforcement")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_lease_time"
		if _, ok := i["dhcp-lease-time"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingDhcpLeaseTime(i["dhcp-lease-time"], d, pre_append)
			tmp["dhcp_lease_time"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-DhcpLeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option43_insertion"
		if _, ok := i["dhcp-option43-insertion"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingDhcpOption43Insertion(i["dhcp-option43-insertion"], d, pre_append)
			tmp["dhcp_option43_insertion"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-DhcpOption43Insertion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_circuit_id_insertion"
		if _, ok := i["dhcp-option82-circuit-id-insertion"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion(i["dhcp-option82-circuit-id-insertion"], d, pre_append)
			tmp["dhcp_option82_circuit_id_insertion"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-DhcpOption82CircuitIdInsertion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_insertion"
		if _, ok := i["dhcp-option82-insertion"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingDhcpOption82Insertion(i["dhcp-option82-insertion"], d, pre_append)
			tmp["dhcp_option82_insertion"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-DhcpOption82Insertion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_remote_id_insertion"
		if _, ok := i["dhcp-option82-remote-id-insertion"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion(i["dhcp-option82-remote-id-insertion"], d, pre_append)
			tmp["dhcp_option82_remote_id_insertion"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-DhcpOption82RemoteIdInsertion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dynamic_vlan"
		if _, ok := i["dynamic-vlan"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingDynamicVlan(i["dynamic-vlan"], d, pre_append)
			tmp["dynamic_vlan"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-DynamicVlan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_reauth"
		if _, ok := i["eap-reauth"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingEapReauth(i["eap-reauth"], d, pre_append)
			tmp["eap_reauth"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-EapReauth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_reauth_intv"
		if _, ok := i["eap-reauth-intv"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingEapReauthIntv(i["eap-reauth-intv"], d, pre_append)
			tmp["eap_reauth_intv"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-EapReauthIntv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eapol_key_retries"
		if _, ok := i["eapol-key-retries"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingEapolKeyRetries(i["eapol-key-retries"], d, pre_append)
			tmp["eapol_key_retries"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-EapolKeyRetries")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encrypt"
		if _, ok := i["encrypt"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingEncrypt(i["encrypt"], d, pre_append)
			tmp["encrypt"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Encrypt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_fast_roaming"
		if _, ok := i["external-fast-roaming"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingExternalFastRoaming(i["external-fast-roaming"], d, pre_append)
			tmp["external_fast_roaming"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ExternalFastRoaming")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_logout"
		if _, ok := i["external-logout"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingExternalLogout(i["external-logout"], d, pre_append)
			tmp["external_logout"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ExternalLogout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_web"
		if _, ok := i["external-web"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingExternalWeb(i["external-web"], d, pre_append)
			tmp["external_web"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ExternalWeb")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_web_format"
		if _, ok := i["external-web-format"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingExternalWebFormat(i["external-web-format"], d, pre_append)
			tmp["external_web_format"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ExternalWebFormat")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fast_bss_transition"
		if _, ok := i["fast-bss-transition"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingFastBssTransition(i["fast-bss-transition"], d, pre_append)
			tmp["fast_bss_transition"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-FastBssTransition")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fast_roaming"
		if _, ok := i["fast-roaming"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingFastRoaming(i["fast-roaming"], d, pre_append)
			tmp["fast_roaming"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-FastRoaming")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_mobility_domain"
		if _, ok := i["ft-mobility-domain"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingFtMobilityDomain(i["ft-mobility-domain"], d, pre_append)
			tmp["ft_mobility_domain"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-FtMobilityDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_over_ds"
		if _, ok := i["ft-over-ds"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingFtOverDs(i["ft-over-ds"], d, pre_append)
			tmp["ft_over_ds"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-FtOverDs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_r0_key_lifetime"
		if _, ok := i["ft-r0-key-lifetime"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingFtR0KeyLifetime(i["ft-r0-key-lifetime"], d, pre_append)
			tmp["ft_r0_key_lifetime"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-FtR0KeyLifetime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gas_comeback_delay"
		if _, ok := i["gas-comeback-delay"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingGasComebackDelay(i["gas-comeback-delay"], d, pre_append)
			tmp["gas_comeback_delay"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-GasComebackDelay")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gas_fragmentation_limit"
		if _, ok := i["gas-fragmentation-limit"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingGasFragmentationLimit(i["gas-fragmentation-limit"], d, pre_append)
			tmp["gas_fragmentation_limit"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-GasFragmentationLimit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gtk_rekey"
		if _, ok := i["gtk-rekey"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingGtkRekey(i["gtk-rekey"], d, pre_append)
			tmp["gtk_rekey"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-GtkRekey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gtk_rekey_intv"
		if _, ok := i["gtk-rekey-intv"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingGtkRekeyIntv(i["gtk-rekey-intv"], d, pre_append)
			tmp["gtk_rekey_intv"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-GtkRekeyIntv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "high_efficiency"
		if _, ok := i["high-efficiency"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingHighEfficiency(i["high-efficiency"], d, pre_append)
			tmp["high_efficiency"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-HighEfficiency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hotspot20_profile"
		if _, ok := i["hotspot20-profile"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingHotspot20Profile(i["hotspot20-profile"], d, pre_append)
			tmp["hotspot20_profile"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Hotspot20Profile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if _, ok := i["igmp-snooping"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIgmpSnooping(i["igmp-snooping"], d, pre_append)
			tmp["igmp_snooping"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IgmpSnooping")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intra_vap_privacy"
		if _, ok := i["intra-vap-privacy"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIntraVapPrivacy(i["intra-vap-privacy"], d, pre_append)
			tmp["intra_vap_privacy"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IntraVapPrivacy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ips_sensor"
		if _, ok := i["ips-sensor"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIpsSensor(i["ips-sensor"], d, pre_append)
			tmp["ips_sensor"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-IpsSensor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_rules"
		if _, ok := i["ipv6-rules"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingIpv6Rules(i["ipv6-rules"], d, pre_append)
			tmp["ipv6_rules"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Ipv6Rules")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingKey(i["key"], d, pre_append)
			tmp["key"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Key")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyindex"
		if _, ok := i["keyindex"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingKeyindex(i["keyindex"], d, pre_append)
			tmp["keyindex"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Keyindex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "l3_roaming"
		if _, ok := i["l3-roaming"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingL3Roaming(i["l3-roaming"], d, pre_append)
			tmp["l3_roaming"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-L3Roaming")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "l3_roaming_mode"
		if _, ok := i["l3-roaming-mode"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingL3RoamingMode(i["l3-roaming-mode"], d, pre_append)
			tmp["l3_roaming_mode"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-L3RoamingMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldpc"
		if _, ok := i["ldpc"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingLdpc(i["ldpc"], d, pre_append)
			tmp["ldpc"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Ldpc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_authentication"
		if _, ok := i["local-authentication"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingLocalAuthentication(i["local-authentication"], d, pre_append)
			tmp["local_authentication"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-LocalAuthentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_bridging"
		if _, ok := i["local-bridging"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingLocalBridging(i["local-bridging"], d, pre_append)
			tmp["local_bridging"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-LocalBridging")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_lan"
		if _, ok := i["local-lan"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingLocalLan(i["local-lan"], d, pre_append)
			tmp["local_lan"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-LocalLan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone"
		if _, ok := i["local-standalone"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingLocalStandalone(i["local-standalone"], d, pre_append)
			tmp["local_standalone"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-LocalStandalone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_dns"
		if _, ok := i["local-standalone-dns"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneDns(i["local-standalone-dns"], d, pre_append)
			tmp["local_standalone_dns"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-LocalStandaloneDns")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_dns_ip"
		if _, ok := i["local-standalone-dns-ip"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneDnsIp(i["local-standalone-dns-ip"], d, pre_append)
			tmp["local_standalone_dns_ip"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-LocalStandaloneDnsIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_nat"
		if _, ok := i["local-standalone-nat"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneNat(i["local-standalone-nat"], d, pre_append)
			tmp["local_standalone_nat"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-LocalStandaloneNat")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_switching"
		if _, ok := i["local-switching"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingLocalSwitching(i["local-switching"], d, pre_append)
			tmp["local_switching"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-LocalSwitching")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_auth_bypass"
		if _, ok := i["mac-auth-bypass"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMacAuthBypass(i["mac-auth-bypass"], d, pre_append)
			tmp["mac_auth_bypass"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MacAuthBypass")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_called_station_delimiter"
		if _, ok := i["mac-called-station-delimiter"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMacCalledStationDelimiter(i["mac-called-station-delimiter"], d, pre_append)
			tmp["mac_called_station_delimiter"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MacCalledStationDelimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_calling_station_delimiter"
		if _, ok := i["mac-calling-station-delimiter"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMacCallingStationDelimiter(i["mac-calling-station-delimiter"], d, pre_append)
			tmp["mac_calling_station_delimiter"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MacCallingStationDelimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_case"
		if _, ok := i["mac-case"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMacCase(i["mac-case"], d, pre_append)
			tmp["mac_case"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MacCase")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter"
		if _, ok := i["mac-filter"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMacFilter(i["mac-filter"], d, pre_append)
			tmp["mac_filter"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MacFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy_other"
		if _, ok := i["mac-filter-policy-other"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMacFilterPolicyOther(i["mac-filter-policy-other"], d, pre_append)
			tmp["mac_filter_policy_other"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MacFilterPolicyOther")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_password_delimiter"
		if _, ok := i["mac-password-delimiter"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMacPasswordDelimiter(i["mac-password-delimiter"], d, pre_append)
			tmp["mac_password_delimiter"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MacPasswordDelimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_username_delimiter"
		if _, ok := i["mac-username-delimiter"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMacUsernameDelimiter(i["mac-username-delimiter"], d, pre_append)
			tmp["mac_username_delimiter"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MacUsernameDelimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_clients"
		if _, ok := i["max-clients"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMaxClients(i["max-clients"], d, pre_append)
			tmp["max_clients"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MaxClients")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_clients_ap"
		if _, ok := i["max-clients-ap"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMaxClientsAp(i["max-clients-ap"], d, pre_append)
			tmp["max_clients_ap"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MaxClientsAp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mbo"
		if _, ok := i["mbo"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMbo(i["mbo"], d, pre_append)
			tmp["mbo"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Mbo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mbo_cell_data_conn_pref"
		if _, ok := i["mbo-cell-data-conn-pref"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMboCellDataConnPref(i["mbo-cell-data-conn-pref"], d, pre_append)
			tmp["mbo_cell_data_conn_pref"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MboCellDataConnPref")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "me_disable_thresh"
		if _, ok := i["me-disable-thresh"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMeDisableThresh(i["me-disable-thresh"], d, pre_append)
			tmp["me_disable_thresh"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MeDisableThresh")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mesh_backhaul"
		if _, ok := i["mesh-backhaul"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMeshBackhaul(i["mesh-backhaul"], d, pre_append)
			tmp["mesh_backhaul"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MeshBackhaul")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk"
		if _, ok := i["mpsk"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMpsk(i["mpsk"], d, pre_append)
			tmp["mpsk"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Mpsk")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_concurrent_clients"
		if _, ok := i["mpsk-concurrent-clients"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMpskConcurrentClients(i["mpsk-concurrent-clients"], d, pre_append)
			tmp["mpsk_concurrent_clients"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MpskConcurrentClients")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_profile"
		if _, ok := i["mpsk-profile"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMpskProfile(i["mpsk-profile"], d, pre_append)
			tmp["mpsk_profile"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MpskProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mu_mimo"
		if _, ok := i["mu-mimo"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMuMimo(i["mu-mimo"], d, pre_append)
			tmp["mu_mimo"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MuMimo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_enhance"
		if _, ok := i["multicast-enhance"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMulticastEnhance(i["multicast-enhance"], d, pre_append)
			tmp["multicast_enhance"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MulticastEnhance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_rate"
		if _, ok := i["multicast-rate"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingMulticastRate(i["multicast-rate"], d, pre_append)
			tmp["multicast_rate"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-MulticastRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nac"
		if _, ok := i["nac"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingNac(i["nac"], d, pre_append)
			tmp["nac"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Nac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nac_profile"
		if _, ok := i["nac-profile"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingNacProfile(i["nac-profile"], d, pre_append)
			tmp["nac_profile"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-NacProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_report_dual_band"
		if _, ok := i["neighbor-report-dual-band"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingNeighborReportDualBand(i["neighbor-report-dual-band"], d, pre_append)
			tmp["neighbor_report_dual_band"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-NeighborReportDualBand")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "okc"
		if _, ok := i["okc"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingOkc(i["okc"], d, pre_append)
			tmp["okc"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Okc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "osen"
		if _, ok := i["osen"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingOsen(i["osen"], d, pre_append)
			tmp["osen"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Osen")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_groups"
		if _, ok := i["owe-groups"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingOweGroups(i["owe-groups"], d, pre_append)
			tmp["owe_groups"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-OweGroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_transition"
		if _, ok := i["owe-transition"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingOweTransition(i["owe-transition"], d, pre_append)
			tmp["owe_transition"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-OweTransition")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_transition_ssid"
		if _, ok := i["owe-transition-ssid"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingOweTransitionSsid(i["owe-transition-ssid"], d, pre_append)
			tmp["owe_transition_ssid"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-OweTransitionSsid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := i["passphrase"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPassphrase(i["passphrase"], d, pre_append)
			tmp["passphrase"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Passphrase")
			c := d.Get(pre_append).(*schema.Set)
			if c.Len() > 0 {
				tmp["passphrase"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf"
		if _, ok := i["pmf"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPmf(i["pmf"], d, pre_append)
			tmp["pmf"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Pmf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf_assoc_comeback_timeout"
		if _, ok := i["pmf-assoc-comeback-timeout"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPmfAssocComebackTimeout(i["pmf-assoc-comeback-timeout"], d, pre_append)
			tmp["pmf_assoc_comeback_timeout"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-PmfAssocComebackTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf_sa_query_retry_timeout"
		if _, ok := i["pmf-sa-query-retry-timeout"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout(i["pmf-sa-query-retry-timeout"], d, pre_append)
			tmp["pmf_sa_query_retry_timeout"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-PmfSaQueryRetryTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth"
		if _, ok := i["port-macauth"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPortMacauth(i["port-macauth"], d, pre_append)
			tmp["port_macauth"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-PortMacauth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth_reauth_timeout"
		if _, ok := i["port-macauth-reauth-timeout"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPortMacauthReauthTimeout(i["port-macauth-reauth-timeout"], d, pre_append)
			tmp["port_macauth_reauth_timeout"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-PortMacauthReauthTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth_timeout"
		if _, ok := i["port-macauth-timeout"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPortMacauthTimeout(i["port-macauth-timeout"], d, pre_append)
			tmp["port_macauth_timeout"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-PortMacauthTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal_message_override_group"
		if _, ok := i["portal-message-override-group"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPortalMessageOverrideGroup(i["portal-message-override-group"], d, pre_append)
			tmp["portal_message_override_group"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-PortalMessageOverrideGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal_type"
		if _, ok := i["portal-type"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPortalType(i["portal-type"], d, pre_append)
			tmp["portal_type"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-PortalType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "primary_wag_profile"
		if _, ok := i["primary-wag-profile"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPrimaryWagProfile(i["primary-wag-profile"], d, pre_append)
			tmp["primary_wag_profile"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-PrimaryWagProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_resp_suppression"
		if _, ok := i["probe-resp-suppression"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingProbeRespSuppression(i["probe-resp-suppression"], d, pre_append)
			tmp["probe_resp_suppression"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ProbeRespSuppression")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_resp_threshold"
		if _, ok := i["probe-resp-threshold"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingProbeRespThreshold(i["probe-resp-threshold"], d, pre_append)
			tmp["probe_resp_threshold"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ProbeRespThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptk_rekey"
		if _, ok := i["ptk-rekey"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPtkRekey(i["ptk-rekey"], d, pre_append)
			tmp["ptk_rekey"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-PtkRekey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptk_rekey_intv"
		if _, ok := i["ptk-rekey-intv"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingPtkRekeyIntv(i["ptk-rekey-intv"], d, pre_append)
			tmp["ptk_rekey_intv"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-PtkRekeyIntv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_profile"
		if _, ok := i["qos-profile"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingQosProfile(i["qos-profile"], d, pre_append)
			tmp["qos_profile"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-QosProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_2g_threshold"
		if _, ok := i["radio-2g-threshold"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRadio2GThreshold(i["radio-2g-threshold"], d, pre_append)
			tmp["radio_2g_threshold"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Radio2GThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_5g_threshold"
		if _, ok := i["radio-5g-threshold"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRadio5GThreshold(i["radio-5g-threshold"], d, pre_append)
			tmp["radio_5g_threshold"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Radio5GThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_sensitivity"
		if _, ok := i["radio-sensitivity"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRadioSensitivity(i["radio-sensitivity"], d, pre_append)
			tmp["radio_sensitivity"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-RadioSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth"
		if _, ok := i["radius-mac-auth"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuth(i["radius-mac-auth"], d, pre_append)
			tmp["radius_mac_auth"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-RadiusMacAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_block_interval"
		if _, ok := i["radius-mac-auth-block-interval"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval(i["radius-mac-auth-block-interval"], d, pre_append)
			tmp["radius_mac_auth_block_interval"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-RadiusMacAuthBlockInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_server"
		if _, ok := i["radius-mac-auth-server"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthServer(i["radius-mac-auth-server"], d, pre_append)
			tmp["radius_mac_auth_server"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-RadiusMacAuthServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_usergroups"
		if _, ok := i["radius-mac-auth-usergroups"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups(i["radius-mac-auth-usergroups"], d, pre_append)
			tmp["radius_mac_auth_usergroups"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-RadiusMacAuthUsergroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_mpsk_auth"
		if _, ok := i["radius-mac-mpsk-auth"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRadiusMacMpskAuth(i["radius-mac-mpsk-auth"], d, pre_append)
			tmp["radius_mac_mpsk_auth"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-RadiusMacMpskAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_mpsk_timeout"
		if _, ok := i["radius-mac-mpsk-timeout"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRadiusMacMpskTimeout(i["radius-mac-mpsk-timeout"], d, pre_append)
			tmp["radius_mac_mpsk_timeout"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-RadiusMacMpskTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_server"
		if _, ok := i["radius-server"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRadiusServer(i["radius-server"], d, pre_append)
			tmp["radius_server"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-RadiusServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11a"
		if _, ok := i["rates-11a"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRates11A(i["rates-11a"], d, pre_append)
			tmp["rates_11a"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Rates11A")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_mcs_map"
		if _, ok := i["rates-11ac-mcs-map"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRates11AcMcsMap(i["rates-11ac-mcs-map"], d, pre_append)
			tmp["rates_11ac_mcs_map"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Rates11AcMcsMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_ss12"
		if _, ok := i["rates-11ac-ss12"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRates11AcSs12(i["rates-11ac-ss12"], d, pre_append)
			tmp["rates_11ac_ss12"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Rates11AcSs12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_ss34"
		if _, ok := i["rates-11ac-ss34"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRates11AcSs34(i["rates-11ac-ss34"], d, pre_append)
			tmp["rates_11ac_ss34"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Rates11AcSs34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_mcs_map"
		if _, ok := i["rates-11ax-mcs-map"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRates11AxMcsMap(i["rates-11ax-mcs-map"], d, pre_append)
			tmp["rates_11ax_mcs_map"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Rates11AxMcsMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_ss12"
		if _, ok := i["rates-11ax-ss12"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRates11AxSs12(i["rates-11ax-ss12"], d, pre_append)
			tmp["rates_11ax_ss12"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Rates11AxSs12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_ss34"
		if _, ok := i["rates-11ax-ss34"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRates11AxSs34(i["rates-11ax-ss34"], d, pre_append)
			tmp["rates_11ax_ss34"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Rates11AxSs34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11bg"
		if _, ok := i["rates-11bg"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRates11Bg(i["rates-11bg"], d, pre_append)
			tmp["rates_11bg"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Rates11Bg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11n_ss12"
		if _, ok := i["rates-11n-ss12"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRates11NSs12(i["rates-11n-ss12"], d, pre_append)
			tmp["rates_11n_ss12"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Rates11NSs12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11n_ss34"
		if _, ok := i["rates-11n-ss34"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingRates11NSs34(i["rates-11n-ss34"], d, pre_append)
			tmp["rates_11n_ss34"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Rates11NSs34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_groups"
		if _, ok := i["sae-groups"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSaeGroups(i["sae-groups"], d, pre_append)
			tmp["sae_groups"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SaeGroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_h2e_only"
		if _, ok := i["sae-h2e-only"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSaeH2EOnly(i["sae-h2e-only"], d, pre_append)
			tmp["sae_h2e_only"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SaeH2EOnly")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_password"
		if _, ok := i["sae-password"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSaePassword(i["sae-password"], d, pre_append)
			tmp["sae_password"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SaePassword")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_pk"
		if _, ok := i["sae-pk"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSaePk(i["sae-pk"], d, pre_append)
			tmp["sae_pk"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SaePk")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_private_key"
		if _, ok := i["sae-private-key"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSaePrivateKey(i["sae-private-key"], d, pre_append)
			tmp["sae_private_key"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SaePrivateKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scan_botnet_connections"
		if _, ok := i["scan-botnet-connections"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingScanBotnetConnections(i["scan-botnet-connections"], d, pre_append)
			tmp["scan_botnet_connections"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-ScanBotnetConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "schedule"
		if _, ok := i["schedule"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSchedule(i["schedule"], d, pre_append)
			tmp["schedule"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Schedule")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_wag_profile"
		if _, ok := i["secondary-wag-profile"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSecondaryWagProfile(i["secondary-wag-profile"], d, pre_append)
			tmp["secondary_wag_profile"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SecondaryWagProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := i["security"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSecurity(i["security"], d, pre_append)
			tmp["security"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Security")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_exempt_list"
		if _, ok := i["security-exempt-list"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSecurityExemptList(i["security-exempt-list"], d, pre_append)
			tmp["security_exempt_list"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SecurityExemptList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_obsolete_option"
		if _, ok := i["security-obsolete-option"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSecurityObsoleteOption(i["security-obsolete-option"], d, pre_append)
			tmp["security_obsolete_option"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SecurityObsoleteOption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_redirect_url"
		if _, ok := i["security-redirect-url"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSecurityRedirectUrl(i["security-redirect-url"], d, pre_append)
			tmp["security_redirect_url"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SecurityRedirectUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "selected_usergroups"
		if _, ok := i["selected-usergroups"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSelectedUsergroups(i["selected-usergroups"], d, pre_append)
			tmp["selected_usergroups"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SelectedUsergroups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_tunneling"
		if _, ok := i["split-tunneling"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSplitTunneling(i["split-tunneling"], d, pre_append)
			tmp["split_tunneling"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-SplitTunneling")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssid"
		if _, ok := i["ssid"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingSsid(i["ssid"], d, pre_append)
			tmp["ssid"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Ssid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_remove"
		if _, ok := i["sticky-client-remove"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingStickyClientRemove(i["sticky-client-remove"], d, pre_append)
			tmp["sticky_client_remove"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-StickyClientRemove")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_2g"
		if _, ok := i["sticky-client-threshold-2g"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold2G(i["sticky-client-threshold-2g"], d, pre_append)
			tmp["sticky_client_threshold_2g"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-StickyClientThreshold2G")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_5g"
		if _, ok := i["sticky-client-threshold-5g"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold5G(i["sticky-client-threshold-5g"], d, pre_append)
			tmp["sticky_client_threshold_5g"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-StickyClientThreshold5G")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_6g"
		if _, ok := i["sticky-client-threshold-6g"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold6G(i["sticky-client-threshold-6g"], d, pre_append)
			tmp["sticky_client_threshold_6g"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-StickyClientThreshold6G")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target_wake_time"
		if _, ok := i["target-wake-time"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingTargetWakeTime(i["target-wake-time"], d, pre_append)
			tmp["target_wake_time"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-TargetWakeTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tkip_counter_measure"
		if _, ok := i["tkip-counter-measure"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingTkipCounterMeasure(i["tkip-counter-measure"], d, pre_append)
			tmp["tkip_counter_measure"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-TkipCounterMeasure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_echo_interval"
		if _, ok := i["tunnel-echo-interval"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingTunnelEchoInterval(i["tunnel-echo-interval"], d, pre_append)
			tmp["tunnel_echo_interval"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-TunnelEchoInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_fallback_interval"
		if _, ok := i["tunnel-fallback-interval"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingTunnelFallbackInterval(i["tunnel-fallback-interval"], d, pre_append)
			tmp["tunnel_fallback_interval"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-TunnelFallbackInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "usergroup"
		if _, ok := i["usergroup"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingUsergroup(i["usergroup"], d, pre_append)
			tmp["usergroup"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Usergroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_log"
		if _, ok := i["utm-log"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingUtmLog(i["utm-log"], d, pre_append)
			tmp["utm_log"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-UtmLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_profile"
		if _, ok := i["utm-profile"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingUtmProfile(i["utm-profile"], d, pre_append)
			tmp["utm_profile"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-UtmProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_status"
		if _, ok := i["utm-status"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingUtmStatus(i["utm-status"], d, pre_append)
			tmp["utm_status"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-UtmStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Vdom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_auto"
		if _, ok := i["vlan-auto"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingVlanAuto(i["vlan-auto"], d, pre_append)
			tmp["vlan_auto"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-VlanAuto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_pooling"
		if _, ok := i["vlan-pooling"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingVlanPooling(i["vlan-pooling"], d, pre_append)
			tmp["vlan_pooling"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-VlanPooling")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlanid"
		if _, ok := i["vlanid"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingVlanid(i["vlanid"], d, pre_append)
			tmp["vlanid"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-Vlanid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "voice_enterprise"
		if _, ok := i["voice-enterprise"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingVoiceEnterprise(i["voice-enterprise"], d, pre_append)
			tmp["voice_enterprise"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-VoiceEnterprise")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "webfilter_profile"
		if _, ok := i["webfilter-profile"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingWebfilterProfile(i["webfilter-profile"], d, pre_append)
			tmp["webfilter_profile"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-DynamicMapping-WebfilterProfile")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerVapDynamicMappingCentmgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpSvrId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDeviceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDeviceIdentification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDeviceNetscan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenObjectWirelessControllerVapDynamicMappingIntfIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfIp6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingIntfListenForticlientConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIsFactorySetting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerVapDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVapDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVapDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerVapDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAccessControlList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAdditionalAkms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingAddressGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAddressGroupPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAntivirusProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingApplicationDetectionEngine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingApplicationDscpMarking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingApplicationReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAtfWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAuthPortalAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingBeaconAdvertising(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingBroadcastSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingBroadcastSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingBssColorPartial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingBstmDisassociationImminent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingBstmRssiDisassocTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalFwAccounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingClientCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpAddressEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpOption43Insertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpOption82Insertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDynamicVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingEapReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingEapReauthIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingEapolKeyRetries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingEncrypt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingExternalFastRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingExternalLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingExternalWeb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingExternalWebFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingFastBssTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingFastRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingFtMobilityDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingFtOverDs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingFtR0KeyLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingGasComebackDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingGasFragmentationLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingGtkRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingGtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingHighEfficiency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingHotspot20Profile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIgmpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntraVapPrivacy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIpv6Rules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingKeyindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingL3Roaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingL3RoamingMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLdpc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalBridging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalLan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalStandalone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneDnsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalSwitching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacCalledStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacCallingStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacCase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacFilterPolicyOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMaxClientsAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMbo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMboCellDataConnPref(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMeDisableThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMeshBackhaul(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMpsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMpskConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMpskProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMuMimo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMulticastEnhance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMulticastRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingNac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingNacProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingNeighborReportDualBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingOkc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingOsen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingOweGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingOweTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingOweTransitionSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPassphrase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingPmf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPmfAssocComebackTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPortMacauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPortMacauthReauthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPortMacauthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPortalMessageOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPortalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPrimaryWagProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingProbeRespSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingProbeRespThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPtkRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingQosProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadio2GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadio5GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadioSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacMpskAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacMpskTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRates11A(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AcMcsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AcSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AcSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AxMcsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AxSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AxSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11Bg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11NSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11NSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingSaeGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingSaeH2EOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSaePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingSaePk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSaePrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSecondaryWagProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSecurityExemptList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSecurityObsoleteOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSecurityRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSelectedUsergroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingStickyClientRemove(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold2G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold5G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold6G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingTargetWakeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingTkipCounterMeasure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingTunnelEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingTunnelFallbackInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingUsergroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingUtmLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingUtmProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingVlanAuto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingVlanPooling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingVoiceEnterprise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapEapReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapEapReauthIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapEapolKeyRetries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapEncrypt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapExternalFastRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapExternalLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapExternalWeb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapExternalWebFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapFastBssTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapFastRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapFtMobilityDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapFtOverDs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapFtR0KeyLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapGasComebackDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapGasFragmentationLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapGtkRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapGtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapHighEfficiency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapHotspot20Profile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIgmpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIntraVapPrivacy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapIpv6Rules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapKeyindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapL3Roaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapL3RoamingMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapLdpc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapLocalAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapLocalBridging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapLocalLan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapLocalStandalone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapLocalStandaloneDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapLocalStandaloneDnsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapLocalStandaloneNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacCalledStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacCallingStationDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacCase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacFilterList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerVapMacFilterListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-MacFilterList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectWirelessControllerVapMacFilterListMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-MacFilterList-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy"
		if _, ok := i["mac-filter-policy"]; ok {
			v := flattenObjectWirelessControllerVapMacFilterListMacFilterPolicy(i["mac-filter-policy"], d, pre_append)
			tmp["mac_filter_policy"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-MacFilterList-MacFilterPolicy")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerVapMacFilterListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacFilterListMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacFilterListMacFilterPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacFilterPolicyOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMaxClientsAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMbo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMboCellDataConnPref(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMeDisableThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMeshBackhaul(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMpsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMpskConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMpskKey(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectWirelessControllerVapMpskKeyComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-MpskKey-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := i["concurrent-clients"]; ok {
			v := flattenObjectWirelessControllerVapMpskKeyConcurrentClients(i["concurrent-clients"], d, pre_append)
			tmp["concurrent_clients"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-MpskKey-ConcurrentClients")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_name"
		if _, ok := i["key-name"]; ok {
			v := flattenObjectWirelessControllerVapMpskKeyKeyName(i["key-name"], d, pre_append)
			tmp["key_name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-MpskKey-KeyName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := i["mpsk-schedules"]; ok {
			v := flattenObjectWirelessControllerVapMpskKeyMpskSchedules(i["mpsk-schedules"], d, pre_append)
			tmp["mpsk_schedules"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-MpskKey-MpskSchedules")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := i["passphrase"]; ok {
			v := flattenObjectWirelessControllerVapMpskKeyPassphrase(i["passphrase"], d, pre_append)
			tmp["passphrase"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-MpskKey-Passphrase")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerVapMpskKeyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMpskKeyConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMpskKeyKeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMpskKeyMpskSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMpskKeyPassphrase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapMpskProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMuMimo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMulticastEnhance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMulticastRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapNac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapNacProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapNeighborReportDualBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapOkc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapOsen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapOweGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapOweTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapOweTransitionSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPassphrase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapPmf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPmfAssocComebackTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPmfSaQueryRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortMacauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortMacauthReauthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortMacauthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortalMessageOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortalMessageOverrides(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_disclaimer_page"
	if _, ok := i["auth-disclaimer-page"]; ok {
		result["auth_disclaimer_page"] = flattenObjectWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(i["auth-disclaimer-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_login_failed_page"
	if _, ok := i["auth-login-failed-page"]; ok {
		result["auth_login_failed_page"] = flattenObjectWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(i["auth-login-failed-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_login_page"
	if _, ok := i["auth-login-page"]; ok {
		result["auth_login_page"] = flattenObjectWirelessControllerVapPortalMessageOverridesAuthLoginPage(i["auth-login-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_reject_page"
	if _, ok := i["auth-reject-page"]; ok {
		result["auth_reject_page"] = flattenObjectWirelessControllerVapPortalMessageOverridesAuthRejectPage(i["auth-reject-page"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortalMessageOverridesAuthLoginPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortalMessageOverridesAuthRejectPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPrimaryWagProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapProbeRespSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapProbeRespThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPtkRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapQosProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRadio2GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRadio5GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRadioSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRadiusMacAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRadiusMacAuthBlockInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRadiusMacAuthServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRadiusMacAuthUsergroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapRadiusMacMpskAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRadiusMacMpskTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRates11A(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapRates11AcMcsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRates11AcSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapRates11AcSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapRates11AxMcsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapRates11AxSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapRates11AxSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapRates11Bg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapRates11NSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapRates11NSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapSaeGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapSaeH2EOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapSaePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapSaePk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapSaePrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapSecondaryWagProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapSecurityExemptList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapSecurityObsoleteOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapSecurityRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapSelectedUsergroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapSplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapStickyClientRemove(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapStickyClientThreshold2G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapStickyClientThreshold5G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapStickyClientThreshold6G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapTargetWakeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapTkipCounterMeasure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapTunnelEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapTunnelFallbackInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapUsergroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapUtmLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapUtmProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVlanAuto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVlanName(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerVapVlanNameName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-VlanName-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if _, ok := i["vlan-id"]; ok {
			v := flattenObjectWirelessControllerVapVlanNameVlanId(i["vlan-id"], d, pre_append)
			tmp["vlan_id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-VlanName-VlanId")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerVapVlanNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVlanNameVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVlanPool(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_wtp_group"
		if _, ok := i["_wtp-group"]; ok {
			v := flattenObjectWirelessControllerVapVlanPoolWtpGroup(i["_wtp-group"], d, pre_append)
			tmp["_wtp_group"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-VlanPool-WtpGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWirelessControllerVapVlanPoolId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVap-VlanPool-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerVapVlanPoolWtpGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVlanPoolId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVlanPooling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVoiceEnterprise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerVap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_centmgmt", flattenObjectWirelessControllerVapCentmgmt(o["_centmgmt"], d, "_centmgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["_centmgmt"], "ObjectWirelessControllerVap-Centmgmt"); ok {
			if err = d.Set("_centmgmt", vv); err != nil {
				return fmt.Errorf("Error reading _centmgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _centmgmt: %v", err)
		}
	}

	if err = d.Set("_dhcp_svr_id", flattenObjectWirelessControllerVapDhcpSvrId(o["_dhcp_svr_id"], d, "_dhcp_svr_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dhcp_svr_id"], "ObjectWirelessControllerVap-DhcpSvrId"); ok {
			if err = d.Set("_dhcp_svr_id", vv); err != nil {
				return fmt.Errorf("Error reading _dhcp_svr_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dhcp_svr_id: %v", err)
		}
	}

	if err = d.Set("_intf_allowaccess", flattenObjectWirelessControllerVapIntfAllowaccess(o["_intf_allowaccess"], d, "_intf_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_allowaccess"], "ObjectWirelessControllerVap-IntfAllowaccess"); ok {
			if err = d.Set("_intf_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading _intf_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_allowaccess: %v", err)
		}
	}

	if err = d.Set("_intf_device_access_list", flattenObjectWirelessControllerVapIntfDeviceAccessList(o["_intf_device-access-list"], d, "_intf_device_access_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_device-access-list"], "ObjectWirelessControllerVap-IntfDeviceAccessList"); ok {
			if err = d.Set("_intf_device_access_list", vv); err != nil {
				return fmt.Errorf("Error reading _intf_device_access_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_device_access_list: %v", err)
		}
	}

	if err = d.Set("_intf_device_identification", flattenObjectWirelessControllerVapIntfDeviceIdentification(o["_intf_device-identification"], d, "_intf_device_identification")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_device-identification"], "ObjectWirelessControllerVap-IntfDeviceIdentification"); ok {
			if err = d.Set("_intf_device_identification", vv); err != nil {
				return fmt.Errorf("Error reading _intf_device_identification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_device_identification: %v", err)
		}
	}

	if err = d.Set("_intf_device_netscan", flattenObjectWirelessControllerVapIntfDeviceNetscan(o["_intf_device-netscan"], d, "_intf_device_netscan")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_device-netscan"], "ObjectWirelessControllerVap-IntfDeviceNetscan"); ok {
			if err = d.Set("_intf_device_netscan", vv); err != nil {
				return fmt.Errorf("Error reading _intf_device_netscan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_device_netscan: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp_relay_ip", flattenObjectWirelessControllerVapIntfDhcpRelayIp(o["_intf_dhcp-relay-ip"], d, "_intf_dhcp_relay_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp-relay-ip"], "ObjectWirelessControllerVap-IntfDhcpRelayIp"); ok {
			if err = d.Set("_intf_dhcp_relay_ip", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp_relay_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp_relay_ip: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp_relay_service", flattenObjectWirelessControllerVapIntfDhcpRelayService(o["_intf_dhcp-relay-service"], d, "_intf_dhcp_relay_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp-relay-service"], "ObjectWirelessControllerVap-IntfDhcpRelayService"); ok {
			if err = d.Set("_intf_dhcp_relay_service", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp_relay_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp_relay_service: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp_relay_type", flattenObjectWirelessControllerVapIntfDhcpRelayType(o["_intf_dhcp-relay-type"], d, "_intf_dhcp_relay_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp-relay-type"], "ObjectWirelessControllerVap-IntfDhcpRelayType"); ok {
			if err = d.Set("_intf_dhcp_relay_type", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp_relay_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp_relay_type: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp6_relay_ip", flattenObjectWirelessControllerVapIntfDhcp6RelayIp(o["_intf_dhcp6-relay-ip"], d, "_intf_dhcp6_relay_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp6-relay-ip"], "ObjectWirelessControllerVap-IntfDhcp6RelayIp"); ok {
			if err = d.Set("_intf_dhcp6_relay_ip", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp6_relay_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp6_relay_ip: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp6_relay_service", flattenObjectWirelessControllerVapIntfDhcp6RelayService(o["_intf_dhcp6-relay-service"], d, "_intf_dhcp6_relay_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp6-relay-service"], "ObjectWirelessControllerVap-IntfDhcp6RelayService"); ok {
			if err = d.Set("_intf_dhcp6_relay_service", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp6_relay_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp6_relay_service: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp6_relay_type", flattenObjectWirelessControllerVapIntfDhcp6RelayType(o["_intf_dhcp6-relay-type"], d, "_intf_dhcp6_relay_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp6-relay-type"], "ObjectWirelessControllerVap-IntfDhcp6RelayType"); ok {
			if err = d.Set("_intf_dhcp6_relay_type", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp6_relay_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp6_relay_type: %v", err)
		}
	}

	if err = d.Set("_intf_ip", flattenObjectWirelessControllerVapIntfIp(o["_intf_ip"], d, "_intf_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_ip"], "ObjectWirelessControllerVap-IntfIp"); ok {
			if err = d.Set("_intf_ip", vv); err != nil {
				return fmt.Errorf("Error reading _intf_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_ip: %v", err)
		}
	}

	if err = d.Set("_intf_ip6_address", flattenObjectWirelessControllerVapIntfIp6Address(o["_intf_ip6-address"], d, "_intf_ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_ip6-address"], "ObjectWirelessControllerVap-IntfIp6Address"); ok {
			if err = d.Set("_intf_ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading _intf_ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_ip6_address: %v", err)
		}
	}

	if err = d.Set("_intf_ip6_allowaccess", flattenObjectWirelessControllerVapIntfIp6Allowaccess(o["_intf_ip6-allowaccess"], d, "_intf_ip6_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_ip6-allowaccess"], "ObjectWirelessControllerVap-IntfIp6Allowaccess"); ok {
			if err = d.Set("_intf_ip6_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading _intf_ip6_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_ip6_allowaccess: %v", err)
		}
	}

	if err = d.Set("_intf_listen_forticlient_connection", flattenObjectWirelessControllerVapIntfListenForticlientConnection(o["_intf_listen-forticlient-connection"], d, "_intf_listen_forticlient_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_listen-forticlient-connection"], "ObjectWirelessControllerVap-IntfListenForticlientConnection"); ok {
			if err = d.Set("_intf_listen_forticlient_connection", vv); err != nil {
				return fmt.Errorf("Error reading _intf_listen_forticlient_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_listen_forticlient_connection: %v", err)
		}
	}

	if err = d.Set("_is_factory_setting", flattenObjectWirelessControllerVapIsFactorySetting(o["_is_factory_setting"], d, "_is_factory_setting")); err != nil {
		if vv, ok := fortiAPIPatch(o["_is_factory_setting"], "ObjectWirelessControllerVap-IsFactorySetting"); ok {
			if err = d.Set("_is_factory_setting", vv); err != nil {
				return fmt.Errorf("Error reading _is_factory_setting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _is_factory_setting: %v", err)
		}
	}

	if err = d.Set("access_control_list", flattenObjectWirelessControllerVapAccessControlList(o["access-control-list"], d, "access_control_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-control-list"], "ObjectWirelessControllerVap-AccessControlList"); ok {
			if err = d.Set("access_control_list", vv); err != nil {
				return fmt.Errorf("Error reading access_control_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_control_list: %v", err)
		}
	}

	if err = d.Set("additional_akms", flattenObjectWirelessControllerVapAdditionalAkms(o["additional-akms"], d, "additional_akms")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-akms"], "ObjectWirelessControllerVap-AdditionalAkms"); ok {
			if err = d.Set("additional_akms", vv); err != nil {
				return fmt.Errorf("Error reading additional_akms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_akms: %v", err)
		}
	}

	if err = d.Set("acct_interim_interval", flattenObjectWirelessControllerVapAcctInterimInterval(o["acct-interim-interval"], d, "acct_interim_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-interim-interval"], "ObjectWirelessControllerVap-AcctInterimInterval"); ok {
			if err = d.Set("acct_interim_interval", vv); err != nil {
				return fmt.Errorf("Error reading acct_interim_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_interim_interval: %v", err)
		}
	}

	if err = d.Set("address_group", flattenObjectWirelessControllerVapAddressGroup(o["address-group"], d, "address_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-group"], "ObjectWirelessControllerVap-AddressGroup"); ok {
			if err = d.Set("address_group", vv); err != nil {
				return fmt.Errorf("Error reading address_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_group: %v", err)
		}
	}

	if err = d.Set("address_group_policy", flattenObjectWirelessControllerVapAddressGroupPolicy(o["address-group-policy"], d, "address_group_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-group-policy"], "ObjectWirelessControllerVap-AddressGroupPolicy"); ok {
			if err = d.Set("address_group_policy", vv); err != nil {
				return fmt.Errorf("Error reading address_group_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_group_policy: %v", err)
		}
	}

	if err = d.Set("alias", flattenObjectWirelessControllerVapAlias(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "ObjectWirelessControllerVap-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("antivirus_profile", flattenObjectWirelessControllerVapAntivirusProfile(o["antivirus-profile"], d, "antivirus_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["antivirus-profile"], "ObjectWirelessControllerVap-AntivirusProfile"); ok {
			if err = d.Set("antivirus_profile", vv); err != nil {
				return fmt.Errorf("Error reading antivirus_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antivirus_profile: %v", err)
		}
	}

	if err = d.Set("application_detection_engine", flattenObjectWirelessControllerVapApplicationDetectionEngine(o["application-detection-engine"], d, "application_detection_engine")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-detection-engine"], "ObjectWirelessControllerVap-ApplicationDetectionEngine"); ok {
			if err = d.Set("application_detection_engine", vv); err != nil {
				return fmt.Errorf("Error reading application_detection_engine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_detection_engine: %v", err)
		}
	}

	if err = d.Set("application_dscp_marking", flattenObjectWirelessControllerVapApplicationDscpMarking(o["application-dscp-marking"], d, "application_dscp_marking")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-dscp-marking"], "ObjectWirelessControllerVap-ApplicationDscpMarking"); ok {
			if err = d.Set("application_dscp_marking", vv); err != nil {
				return fmt.Errorf("Error reading application_dscp_marking: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_dscp_marking: %v", err)
		}
	}

	if err = d.Set("application_list", flattenObjectWirelessControllerVapApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "ObjectWirelessControllerVap-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("application_report_intv", flattenObjectWirelessControllerVapApplicationReportIntv(o["application-report-intv"], d, "application_report_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-report-intv"], "ObjectWirelessControllerVap-ApplicationReportIntv"); ok {
			if err = d.Set("application_report_intv", vv); err != nil {
				return fmt.Errorf("Error reading application_report_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_report_intv: %v", err)
		}
	}

	if err = d.Set("atf_weight", flattenObjectWirelessControllerVapAtfWeight(o["atf-weight"], d, "atf_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["atf-weight"], "ObjectWirelessControllerVap-AtfWeight"); ok {
			if err = d.Set("atf_weight", vv); err != nil {
				return fmt.Errorf("Error reading atf_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading atf_weight: %v", err)
		}
	}

	if err = d.Set("auth", flattenObjectWirelessControllerVapAuth(o["auth"], d, "auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth"], "ObjectWirelessControllerVap-Auth"); ok {
			if err = d.Set("auth", vv); err != nil {
				return fmt.Errorf("Error reading auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenObjectWirelessControllerVapAuthCert(o["auth-cert"], d, "auth_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-cert"], "ObjectWirelessControllerVap-AuthCert"); ok {
			if err = d.Set("auth_cert", vv); err != nil {
				return fmt.Errorf("Error reading auth_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_portal_addr", flattenObjectWirelessControllerVapAuthPortalAddr(o["auth-portal-addr"], d, "auth_portal_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal-addr"], "ObjectWirelessControllerVap-AuthPortalAddr"); ok {
			if err = d.Set("auth_portal_addr", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal_addr: %v", err)
		}
	}

	if err = d.Set("beacon_advertising", flattenObjectWirelessControllerVapBeaconAdvertising(o["beacon-advertising"], d, "beacon_advertising")); err != nil {
		if vv, ok := fortiAPIPatch(o["beacon-advertising"], "ObjectWirelessControllerVap-BeaconAdvertising"); ok {
			if err = d.Set("beacon_advertising", vv); err != nil {
				return fmt.Errorf("Error reading beacon_advertising: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beacon_advertising: %v", err)
		}
	}

	if err = d.Set("broadcast_ssid", flattenObjectWirelessControllerVapBroadcastSsid(o["broadcast-ssid"], d, "broadcast_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast-ssid"], "ObjectWirelessControllerVap-BroadcastSsid"); ok {
			if err = d.Set("broadcast_ssid", vv); err != nil {
				return fmt.Errorf("Error reading broadcast_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast_ssid: %v", err)
		}
	}

	if err = d.Set("broadcast_suppression", flattenObjectWirelessControllerVapBroadcastSuppression(o["broadcast-suppression"], d, "broadcast_suppression")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast-suppression"], "ObjectWirelessControllerVap-BroadcastSuppression"); ok {
			if err = d.Set("broadcast_suppression", vv); err != nil {
				return fmt.Errorf("Error reading broadcast_suppression: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast_suppression: %v", err)
		}
	}

	if err = d.Set("bss_color_partial", flattenObjectWirelessControllerVapBssColorPartial(o["bss-color-partial"], d, "bss_color_partial")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-color-partial"], "ObjectWirelessControllerVap-BssColorPartial"); ok {
			if err = d.Set("bss_color_partial", vv); err != nil {
				return fmt.Errorf("Error reading bss_color_partial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_color_partial: %v", err)
		}
	}

	if err = d.Set("bstm_disassociation_imminent", flattenObjectWirelessControllerVapBstmDisassociationImminent(o["bstm-disassociation-imminent"], d, "bstm_disassociation_imminent")); err != nil {
		if vv, ok := fortiAPIPatch(o["bstm-disassociation-imminent"], "ObjectWirelessControllerVap-BstmDisassociationImminent"); ok {
			if err = d.Set("bstm_disassociation_imminent", vv); err != nil {
				return fmt.Errorf("Error reading bstm_disassociation_imminent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bstm_disassociation_imminent: %v", err)
		}
	}

	if err = d.Set("bstm_load_balancing_disassoc_timer", flattenObjectWirelessControllerVapBstmLoadBalancingDisassocTimer(o["bstm-load-balancing-disassoc-timer"], d, "bstm_load_balancing_disassoc_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["bstm-load-balancing-disassoc-timer"], "ObjectWirelessControllerVap-BstmLoadBalancingDisassocTimer"); ok {
			if err = d.Set("bstm_load_balancing_disassoc_timer", vv); err != nil {
				return fmt.Errorf("Error reading bstm_load_balancing_disassoc_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bstm_load_balancing_disassoc_timer: %v", err)
		}
	}

	if err = d.Set("bstm_rssi_disassoc_timer", flattenObjectWirelessControllerVapBstmRssiDisassocTimer(o["bstm-rssi-disassoc-timer"], d, "bstm_rssi_disassoc_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["bstm-rssi-disassoc-timer"], "ObjectWirelessControllerVap-BstmRssiDisassocTimer"); ok {
			if err = d.Set("bstm_rssi_disassoc_timer", vv); err != nil {
				return fmt.Errorf("Error reading bstm_rssi_disassoc_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bstm_rssi_disassoc_timer: %v", err)
		}
	}

	if err = d.Set("captive_portal_ac_name", flattenObjectWirelessControllerVapCaptivePortalAcName(o["captive-portal-ac-name"], d, "captive_portal_ac_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-ac-name"], "ObjectWirelessControllerVap-CaptivePortalAcName"); ok {
			if err = d.Set("captive_portal_ac_name", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_ac_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_ac_name: %v", err)
		}
	}

	if err = d.Set("captive_portal_auth_timeout", flattenObjectWirelessControllerVapCaptivePortalAuthTimeout(o["captive-portal-auth-timeout"], d, "captive_portal_auth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-auth-timeout"], "ObjectWirelessControllerVap-CaptivePortalAuthTimeout"); ok {
			if err = d.Set("captive_portal_auth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_auth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_auth_timeout: %v", err)
		}
	}

	if err = d.Set("captive_portal_fw_accounting", flattenObjectWirelessControllerVapCaptivePortalFwAccounting(o["captive-portal-fw-accounting"], d, "captive_portal_fw_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-fw-accounting"], "ObjectWirelessControllerVap-CaptivePortalFwAccounting"); ok {
			if err = d.Set("captive_portal_fw_accounting", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_fw_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_fw_accounting: %v", err)
		}
	}

	if err = d.Set("dhcp_address_enforcement", flattenObjectWirelessControllerVapDhcpAddressEnforcement(o["dhcp-address-enforcement"], d, "dhcp_address_enforcement")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-address-enforcement"], "ObjectWirelessControllerVap-DhcpAddressEnforcement"); ok {
			if err = d.Set("dhcp_address_enforcement", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_address_enforcement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_address_enforcement: %v", err)
		}
	}

	if err = d.Set("captive_portal_macauth_radius_secret", flattenObjectWirelessControllerVapCaptivePortalMacauthRadiusSecret(o["captive-portal-macauth-radius-secret"], d, "captive_portal_macauth_radius_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-macauth-radius-secret"], "ObjectWirelessControllerVap-CaptivePortalMacauthRadiusSecret"); ok {
			if err = d.Set("captive_portal_macauth_radius_secret", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_macauth_radius_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_macauth_radius_secret: %v", err)
		}
	}

	if err = d.Set("captive_portal_macauth_radius_server", flattenObjectWirelessControllerVapCaptivePortalMacauthRadiusServer(o["captive-portal-macauth-radius-server"], d, "captive_portal_macauth_radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-macauth-radius-server"], "ObjectWirelessControllerVap-CaptivePortalMacauthRadiusServer"); ok {
			if err = d.Set("captive_portal_macauth_radius_server", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_macauth_radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_macauth_radius_server: %v", err)
		}
	}

	if err = d.Set("captive_portal_radius_secret", flattenObjectWirelessControllerVapCaptivePortalRadiusSecret(o["captive-portal-radius-secret"], d, "captive_portal_radius_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-radius-secret"], "ObjectWirelessControllerVap-CaptivePortalRadiusSecret"); ok {
			if err = d.Set("captive_portal_radius_secret", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_radius_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_radius_secret: %v", err)
		}
	}

	if err = d.Set("captive_portal_radius_server", flattenObjectWirelessControllerVapCaptivePortalRadiusServer(o["captive-portal-radius-server"], d, "captive_portal_radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-radius-server"], "ObjectWirelessControllerVap-CaptivePortalRadiusServer"); ok {
			if err = d.Set("captive_portal_radius_server", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_radius_server: %v", err)
		}
	}

	if err = d.Set("captive_portal_session_timeout_interval", flattenObjectWirelessControllerVapCaptivePortalSessionTimeoutInterval(o["captive-portal-session-timeout-interval"], d, "captive_portal_session_timeout_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-session-timeout-interval"], "ObjectWirelessControllerVap-CaptivePortalSessionTimeoutInterval"); ok {
			if err = d.Set("captive_portal_session_timeout_interval", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_session_timeout_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_session_timeout_interval: %v", err)
		}
	}

	if err = d.Set("dhcp_lease_time", flattenObjectWirelessControllerVapDhcpLeaseTime(o["dhcp-lease-time"], d, "dhcp_lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-lease-time"], "ObjectWirelessControllerVap-DhcpLeaseTime"); ok {
			if err = d.Set("dhcp_lease_time", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_lease_time: %v", err)
		}
	}

	if err = d.Set("dhcp_option43_insertion", flattenObjectWirelessControllerVapDhcpOption43Insertion(o["dhcp-option43-insertion"], d, "dhcp_option43_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option43-insertion"], "ObjectWirelessControllerVap-DhcpOption43Insertion"); ok {
			if err = d.Set("dhcp_option43_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option43_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option43_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_circuit_id_insertion", flattenObjectWirelessControllerVapDhcpOption82CircuitIdInsertion(o["dhcp-option82-circuit-id-insertion"], d, "dhcp_option82_circuit_id_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-circuit-id-insertion"], "ObjectWirelessControllerVap-DhcpOption82CircuitIdInsertion"); ok {
			if err = d.Set("dhcp_option82_circuit_id_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_circuit_id_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_circuit_id_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_insertion", flattenObjectWirelessControllerVapDhcpOption82Insertion(o["dhcp-option82-insertion"], d, "dhcp_option82_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-insertion"], "ObjectWirelessControllerVap-DhcpOption82Insertion"); ok {
			if err = d.Set("dhcp_option82_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_remote_id_insertion", flattenObjectWirelessControllerVapDhcpOption82RemoteIdInsertion(o["dhcp-option82-remote-id-insertion"], d, "dhcp_option82_remote_id_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-remote-id-insertion"], "ObjectWirelessControllerVap-DhcpOption82RemoteIdInsertion"); ok {
			if err = d.Set("dhcp_option82_remote_id_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_remote_id_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_remote_id_insertion: %v", err)
		}
	}

	if err = d.Set("dynamic_vlan", flattenObjectWirelessControllerVapDynamicVlan(o["dynamic-vlan"], d, "dynamic_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-vlan"], "ObjectWirelessControllerVap-DynamicVlan"); ok {
			if err = d.Set("dynamic_vlan", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_vlan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectWirelessControllerVapDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectWirelessControllerVap-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectWirelessControllerVapDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectWirelessControllerVap-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("eap_reauth", flattenObjectWirelessControllerVapEapReauth(o["eap-reauth"], d, "eap_reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-reauth"], "ObjectWirelessControllerVap-EapReauth"); ok {
			if err = d.Set("eap_reauth", vv); err != nil {
				return fmt.Errorf("Error reading eap_reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_reauth: %v", err)
		}
	}

	if err = d.Set("eap_reauth_intv", flattenObjectWirelessControllerVapEapReauthIntv(o["eap-reauth-intv"], d, "eap_reauth_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-reauth-intv"], "ObjectWirelessControllerVap-EapReauthIntv"); ok {
			if err = d.Set("eap_reauth_intv", vv); err != nil {
				return fmt.Errorf("Error reading eap_reauth_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_reauth_intv: %v", err)
		}
	}

	if err = d.Set("eapol_key_retries", flattenObjectWirelessControllerVapEapolKeyRetries(o["eapol-key-retries"], d, "eapol_key_retries")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-key-retries"], "ObjectWirelessControllerVap-EapolKeyRetries"); ok {
			if err = d.Set("eapol_key_retries", vv); err != nil {
				return fmt.Errorf("Error reading eapol_key_retries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_key_retries: %v", err)
		}
	}

	if err = d.Set("encrypt", flattenObjectWirelessControllerVapEncrypt(o["encrypt"], d, "encrypt")); err != nil {
		if vv, ok := fortiAPIPatch(o["encrypt"], "ObjectWirelessControllerVap-Encrypt"); ok {
			if err = d.Set("encrypt", vv); err != nil {
				return fmt.Errorf("Error reading encrypt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encrypt: %v", err)
		}
	}

	if err = d.Set("external_fast_roaming", flattenObjectWirelessControllerVapExternalFastRoaming(o["external-fast-roaming"], d, "external_fast_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-fast-roaming"], "ObjectWirelessControllerVap-ExternalFastRoaming"); ok {
			if err = d.Set("external_fast_roaming", vv); err != nil {
				return fmt.Errorf("Error reading external_fast_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_fast_roaming: %v", err)
		}
	}

	if err = d.Set("external_logout", flattenObjectWirelessControllerVapExternalLogout(o["external-logout"], d, "external_logout")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-logout"], "ObjectWirelessControllerVap-ExternalLogout"); ok {
			if err = d.Set("external_logout", vv); err != nil {
				return fmt.Errorf("Error reading external_logout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_logout: %v", err)
		}
	}

	if err = d.Set("external_web", flattenObjectWirelessControllerVapExternalWeb(o["external-web"], d, "external_web")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-web"], "ObjectWirelessControllerVap-ExternalWeb"); ok {
			if err = d.Set("external_web", vv); err != nil {
				return fmt.Errorf("Error reading external_web: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_web: %v", err)
		}
	}

	if err = d.Set("external_web_format", flattenObjectWirelessControllerVapExternalWebFormat(o["external-web-format"], d, "external_web_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-web-format"], "ObjectWirelessControllerVap-ExternalWebFormat"); ok {
			if err = d.Set("external_web_format", vv); err != nil {
				return fmt.Errorf("Error reading external_web_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_web_format: %v", err)
		}
	}

	if err = d.Set("fast_bss_transition", flattenObjectWirelessControllerVapFastBssTransition(o["fast-bss-transition"], d, "fast_bss_transition")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-bss-transition"], "ObjectWirelessControllerVap-FastBssTransition"); ok {
			if err = d.Set("fast_bss_transition", vv); err != nil {
				return fmt.Errorf("Error reading fast_bss_transition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_bss_transition: %v", err)
		}
	}

	if err = d.Set("fast_roaming", flattenObjectWirelessControllerVapFastRoaming(o["fast-roaming"], d, "fast_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-roaming"], "ObjectWirelessControllerVap-FastRoaming"); ok {
			if err = d.Set("fast_roaming", vv); err != nil {
				return fmt.Errorf("Error reading fast_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_roaming: %v", err)
		}
	}

	if err = d.Set("ft_mobility_domain", flattenObjectWirelessControllerVapFtMobilityDomain(o["ft-mobility-domain"], d, "ft_mobility_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["ft-mobility-domain"], "ObjectWirelessControllerVap-FtMobilityDomain"); ok {
			if err = d.Set("ft_mobility_domain", vv); err != nil {
				return fmt.Errorf("Error reading ft_mobility_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ft_mobility_domain: %v", err)
		}
	}

	if err = d.Set("ft_over_ds", flattenObjectWirelessControllerVapFtOverDs(o["ft-over-ds"], d, "ft_over_ds")); err != nil {
		if vv, ok := fortiAPIPatch(o["ft-over-ds"], "ObjectWirelessControllerVap-FtOverDs"); ok {
			if err = d.Set("ft_over_ds", vv); err != nil {
				return fmt.Errorf("Error reading ft_over_ds: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ft_over_ds: %v", err)
		}
	}

	if err = d.Set("ft_r0_key_lifetime", flattenObjectWirelessControllerVapFtR0KeyLifetime(o["ft-r0-key-lifetime"], d, "ft_r0_key_lifetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["ft-r0-key-lifetime"], "ObjectWirelessControllerVap-FtR0KeyLifetime"); ok {
			if err = d.Set("ft_r0_key_lifetime", vv); err != nil {
				return fmt.Errorf("Error reading ft_r0_key_lifetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ft_r0_key_lifetime: %v", err)
		}
	}

	if err = d.Set("gas_comeback_delay", flattenObjectWirelessControllerVapGasComebackDelay(o["gas-comeback-delay"], d, "gas_comeback_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["gas-comeback-delay"], "ObjectWirelessControllerVap-GasComebackDelay"); ok {
			if err = d.Set("gas_comeback_delay", vv); err != nil {
				return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
		}
	}

	if err = d.Set("gas_fragmentation_limit", flattenObjectWirelessControllerVapGasFragmentationLimit(o["gas-fragmentation-limit"], d, "gas_fragmentation_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["gas-fragmentation-limit"], "ObjectWirelessControllerVap-GasFragmentationLimit"); ok {
			if err = d.Set("gas_fragmentation_limit", vv); err != nil {
				return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
		}
	}

	if err = d.Set("gtk_rekey", flattenObjectWirelessControllerVapGtkRekey(o["gtk-rekey"], d, "gtk_rekey")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtk-rekey"], "ObjectWirelessControllerVap-GtkRekey"); ok {
			if err = d.Set("gtk_rekey", vv); err != nil {
				return fmt.Errorf("Error reading gtk_rekey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtk_rekey: %v", err)
		}
	}

	if err = d.Set("gtk_rekey_intv", flattenObjectWirelessControllerVapGtkRekeyIntv(o["gtk-rekey-intv"], d, "gtk_rekey_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtk-rekey-intv"], "ObjectWirelessControllerVap-GtkRekeyIntv"); ok {
			if err = d.Set("gtk_rekey_intv", vv); err != nil {
				return fmt.Errorf("Error reading gtk_rekey_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtk_rekey_intv: %v", err)
		}
	}

	if err = d.Set("high_efficiency", flattenObjectWirelessControllerVapHighEfficiency(o["high-efficiency"], d, "high_efficiency")); err != nil {
		if vv, ok := fortiAPIPatch(o["high-efficiency"], "ObjectWirelessControllerVap-HighEfficiency"); ok {
			if err = d.Set("high_efficiency", vv); err != nil {
				return fmt.Errorf("Error reading high_efficiency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading high_efficiency: %v", err)
		}
	}

	if err = d.Set("hotspot20_profile", flattenObjectWirelessControllerVapHotspot20Profile(o["hotspot20-profile"], d, "hotspot20_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["hotspot20-profile"], "ObjectWirelessControllerVap-Hotspot20Profile"); ok {
			if err = d.Set("hotspot20_profile", vv); err != nil {
				return fmt.Errorf("Error reading hotspot20_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hotspot20_profile: %v", err)
		}
	}

	if err = d.Set("igmp_snooping", flattenObjectWirelessControllerVapIgmpSnooping(o["igmp-snooping"], d, "igmp_snooping")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmp-snooping"], "ObjectWirelessControllerVap-IgmpSnooping"); ok {
			if err = d.Set("igmp_snooping", vv); err != nil {
				return fmt.Errorf("Error reading igmp_snooping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmp_snooping: %v", err)
		}
	}

	if err = d.Set("intra_vap_privacy", flattenObjectWirelessControllerVapIntraVapPrivacy(o["intra-vap-privacy"], d, "intra_vap_privacy")); err != nil {
		if vv, ok := fortiAPIPatch(o["intra-vap-privacy"], "ObjectWirelessControllerVap-IntraVapPrivacy"); ok {
			if err = d.Set("intra_vap_privacy", vv); err != nil {
				return fmt.Errorf("Error reading intra_vap_privacy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intra_vap_privacy: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectWirelessControllerVapIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectWirelessControllerVap-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenObjectWirelessControllerVapIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "ObjectWirelessControllerVap-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ipv6_rules", flattenObjectWirelessControllerVapIpv6Rules(o["ipv6-rules"], d, "ipv6_rules")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-rules"], "ObjectWirelessControllerVap-Ipv6Rules"); ok {
			if err = d.Set("ipv6_rules", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_rules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_rules: %v", err)
		}
	}

	if err = d.Set("key", flattenObjectWirelessControllerVapKey(o["key"], d, "key")); err != nil {
		if vv, ok := fortiAPIPatch(o["key"], "ObjectWirelessControllerVap-Key"); ok {
			if err = d.Set("key", vv); err != nil {
				return fmt.Errorf("Error reading key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	if err = d.Set("keyindex", flattenObjectWirelessControllerVapKeyindex(o["keyindex"], d, "keyindex")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyindex"], "ObjectWirelessControllerVap-Keyindex"); ok {
			if err = d.Set("keyindex", vv); err != nil {
				return fmt.Errorf("Error reading keyindex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyindex: %v", err)
		}
	}

	if err = d.Set("l3_roaming", flattenObjectWirelessControllerVapL3Roaming(o["l3-roaming"], d, "l3_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["l3-roaming"], "ObjectWirelessControllerVap-L3Roaming"); ok {
			if err = d.Set("l3_roaming", vv); err != nil {
				return fmt.Errorf("Error reading l3_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l3_roaming: %v", err)
		}
	}

	if err = d.Set("l3_roaming_mode", flattenObjectWirelessControllerVapL3RoamingMode(o["l3-roaming-mode"], d, "l3_roaming_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["l3-roaming-mode"], "ObjectWirelessControllerVap-L3RoamingMode"); ok {
			if err = d.Set("l3_roaming_mode", vv); err != nil {
				return fmt.Errorf("Error reading l3_roaming_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l3_roaming_mode: %v", err)
		}
	}

	if err = d.Set("ldpc", flattenObjectWirelessControllerVapLdpc(o["ldpc"], d, "ldpc")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldpc"], "ObjectWirelessControllerVap-Ldpc"); ok {
			if err = d.Set("ldpc", vv); err != nil {
				return fmt.Errorf("Error reading ldpc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldpc: %v", err)
		}
	}

	if err = d.Set("local_authentication", flattenObjectWirelessControllerVapLocalAuthentication(o["local-authentication"], d, "local_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-authentication"], "ObjectWirelessControllerVap-LocalAuthentication"); ok {
			if err = d.Set("local_authentication", vv); err != nil {
				return fmt.Errorf("Error reading local_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_authentication: %v", err)
		}
	}

	if err = d.Set("local_bridging", flattenObjectWirelessControllerVapLocalBridging(o["local-bridging"], d, "local_bridging")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-bridging"], "ObjectWirelessControllerVap-LocalBridging"); ok {
			if err = d.Set("local_bridging", vv); err != nil {
				return fmt.Errorf("Error reading local_bridging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_bridging: %v", err)
		}
	}

	if err = d.Set("local_lan", flattenObjectWirelessControllerVapLocalLan(o["local-lan"], d, "local_lan")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-lan"], "ObjectWirelessControllerVap-LocalLan"); ok {
			if err = d.Set("local_lan", vv); err != nil {
				return fmt.Errorf("Error reading local_lan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_lan: %v", err)
		}
	}

	if err = d.Set("local_standalone", flattenObjectWirelessControllerVapLocalStandalone(o["local-standalone"], d, "local_standalone")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone"], "ObjectWirelessControllerVap-LocalStandalone"); ok {
			if err = d.Set("local_standalone", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone: %v", err)
		}
	}

	if err = d.Set("local_standalone_dns", flattenObjectWirelessControllerVapLocalStandaloneDns(o["local-standalone-dns"], d, "local_standalone_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone-dns"], "ObjectWirelessControllerVap-LocalStandaloneDns"); ok {
			if err = d.Set("local_standalone_dns", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone_dns: %v", err)
		}
	}

	if err = d.Set("local_standalone_dns_ip", flattenObjectWirelessControllerVapLocalStandaloneDnsIp(o["local-standalone-dns-ip"], d, "local_standalone_dns_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone-dns-ip"], "ObjectWirelessControllerVap-LocalStandaloneDnsIp"); ok {
			if err = d.Set("local_standalone_dns_ip", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone_dns_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone_dns_ip: %v", err)
		}
	}

	if err = d.Set("local_standalone_nat", flattenObjectWirelessControllerVapLocalStandaloneNat(o["local-standalone-nat"], d, "local_standalone_nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone-nat"], "ObjectWirelessControllerVap-LocalStandaloneNat"); ok {
			if err = d.Set("local_standalone_nat", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone_nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone_nat: %v", err)
		}
	}

	if err = d.Set("mac_auth_bypass", flattenObjectWirelessControllerVapMacAuthBypass(o["mac-auth-bypass"], d, "mac_auth_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-auth-bypass"], "ObjectWirelessControllerVap-MacAuthBypass"); ok {
			if err = d.Set("mac_auth_bypass", vv); err != nil {
				return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("mac_called_station_delimiter", flattenObjectWirelessControllerVapMacCalledStationDelimiter(o["mac-called-station-delimiter"], d, "mac_called_station_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-called-station-delimiter"], "ObjectWirelessControllerVap-MacCalledStationDelimiter"); ok {
			if err = d.Set("mac_called_station_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_calling_station_delimiter", flattenObjectWirelessControllerVapMacCallingStationDelimiter(o["mac-calling-station-delimiter"], d, "mac_calling_station_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-calling-station-delimiter"], "ObjectWirelessControllerVap-MacCallingStationDelimiter"); ok {
			if err = d.Set("mac_calling_station_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_case", flattenObjectWirelessControllerVapMacCase(o["mac-case"], d, "mac_case")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-case"], "ObjectWirelessControllerVap-MacCase"); ok {
			if err = d.Set("mac_case", vv); err != nil {
				return fmt.Errorf("Error reading mac_case: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_case: %v", err)
		}
	}

	if err = d.Set("mac_filter", flattenObjectWirelessControllerVapMacFilter(o["mac-filter"], d, "mac_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-filter"], "ObjectWirelessControllerVap-MacFilter"); ok {
			if err = d.Set("mac_filter", vv); err != nil {
				return fmt.Errorf("Error reading mac_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_filter: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mac_filter_list", flattenObjectWirelessControllerVapMacFilterList(o["mac-filter-list"], d, "mac_filter_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["mac-filter-list"], "ObjectWirelessControllerVap-MacFilterList"); ok {
				if err = d.Set("mac_filter_list", vv); err != nil {
					return fmt.Errorf("Error reading mac_filter_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mac_filter_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mac_filter_list"); ok {
			if err = d.Set("mac_filter_list", flattenObjectWirelessControllerVapMacFilterList(o["mac-filter-list"], d, "mac_filter_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["mac-filter-list"], "ObjectWirelessControllerVap-MacFilterList"); ok {
					if err = d.Set("mac_filter_list", vv); err != nil {
						return fmt.Errorf("Error reading mac_filter_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mac_filter_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("mac_filter_policy_other", flattenObjectWirelessControllerVapMacFilterPolicyOther(o["mac-filter-policy-other"], d, "mac_filter_policy_other")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-filter-policy-other"], "ObjectWirelessControllerVap-MacFilterPolicyOther"); ok {
			if err = d.Set("mac_filter_policy_other", vv); err != nil {
				return fmt.Errorf("Error reading mac_filter_policy_other: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_filter_policy_other: %v", err)
		}
	}

	if err = d.Set("mac_password_delimiter", flattenObjectWirelessControllerVapMacPasswordDelimiter(o["mac-password-delimiter"], d, "mac_password_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-password-delimiter"], "ObjectWirelessControllerVap-MacPasswordDelimiter"); ok {
			if err = d.Set("mac_password_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_username_delimiter", flattenObjectWirelessControllerVapMacUsernameDelimiter(o["mac-username-delimiter"], d, "mac_username_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-username-delimiter"], "ObjectWirelessControllerVap-MacUsernameDelimiter"); ok {
			if err = d.Set("mac_username_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenObjectWirelessControllerVapMaxClients(o["max-clients"], d, "max_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients"], "ObjectWirelessControllerVap-MaxClients"); ok {
			if err = d.Set("max_clients", vv); err != nil {
				return fmt.Errorf("Error reading max_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("max_clients_ap", flattenObjectWirelessControllerVapMaxClientsAp(o["max-clients-ap"], d, "max_clients_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients-ap"], "ObjectWirelessControllerVap-MaxClientsAp"); ok {
			if err = d.Set("max_clients_ap", vv); err != nil {
				return fmt.Errorf("Error reading max_clients_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients_ap: %v", err)
		}
	}

	if err = d.Set("mbo", flattenObjectWirelessControllerVapMbo(o["mbo"], d, "mbo")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbo"], "ObjectWirelessControllerVap-Mbo"); ok {
			if err = d.Set("mbo", vv); err != nil {
				return fmt.Errorf("Error reading mbo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbo: %v", err)
		}
	}

	if err = d.Set("mbo_cell_data_conn_pref", flattenObjectWirelessControllerVapMboCellDataConnPref(o["mbo-cell-data-conn-pref"], d, "mbo_cell_data_conn_pref")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbo-cell-data-conn-pref"], "ObjectWirelessControllerVap-MboCellDataConnPref"); ok {
			if err = d.Set("mbo_cell_data_conn_pref", vv); err != nil {
				return fmt.Errorf("Error reading mbo_cell_data_conn_pref: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbo_cell_data_conn_pref: %v", err)
		}
	}

	if err = d.Set("me_disable_thresh", flattenObjectWirelessControllerVapMeDisableThresh(o["me-disable-thresh"], d, "me_disable_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["me-disable-thresh"], "ObjectWirelessControllerVap-MeDisableThresh"); ok {
			if err = d.Set("me_disable_thresh", vv); err != nil {
				return fmt.Errorf("Error reading me_disable_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading me_disable_thresh: %v", err)
		}
	}

	if err = d.Set("mesh_backhaul", flattenObjectWirelessControllerVapMeshBackhaul(o["mesh-backhaul"], d, "mesh_backhaul")); err != nil {
		if vv, ok := fortiAPIPatch(o["mesh-backhaul"], "ObjectWirelessControllerVap-MeshBackhaul"); ok {
			if err = d.Set("mesh_backhaul", vv); err != nil {
				return fmt.Errorf("Error reading mesh_backhaul: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mesh_backhaul: %v", err)
		}
	}

	if err = d.Set("mpsk", flattenObjectWirelessControllerVapMpsk(o["mpsk"], d, "mpsk")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk"], "ObjectWirelessControllerVap-Mpsk"); ok {
			if err = d.Set("mpsk", vv); err != nil {
				return fmt.Errorf("Error reading mpsk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk: %v", err)
		}
	}

	if err = d.Set("mpsk_concurrent_clients", flattenObjectWirelessControllerVapMpskConcurrentClients(o["mpsk-concurrent-clients"], d, "mpsk_concurrent_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-concurrent-clients"], "ObjectWirelessControllerVap-MpskConcurrentClients"); ok {
			if err = d.Set("mpsk_concurrent_clients", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mpsk_key", flattenObjectWirelessControllerVapMpskKey(o["mpsk-key"], d, "mpsk_key")); err != nil {
			if vv, ok := fortiAPIPatch(o["mpsk-key"], "ObjectWirelessControllerVap-MpskKey"); ok {
				if err = d.Set("mpsk_key", vv); err != nil {
					return fmt.Errorf("Error reading mpsk_key: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mpsk_key: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mpsk_key"); ok {
			if err = d.Set("mpsk_key", flattenObjectWirelessControllerVapMpskKey(o["mpsk-key"], d, "mpsk_key")); err != nil {
				if vv, ok := fortiAPIPatch(o["mpsk-key"], "ObjectWirelessControllerVap-MpskKey"); ok {
					if err = d.Set("mpsk_key", vv); err != nil {
						return fmt.Errorf("Error reading mpsk_key: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mpsk_key: %v", err)
				}
			}
		}
	}

	if err = d.Set("mpsk_profile", flattenObjectWirelessControllerVapMpskProfile(o["mpsk-profile"], d, "mpsk_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-profile"], "ObjectWirelessControllerVap-MpskProfile"); ok {
			if err = d.Set("mpsk_profile", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_profile: %v", err)
		}
	}

	if err = d.Set("mu_mimo", flattenObjectWirelessControllerVapMuMimo(o["mu-mimo"], d, "mu_mimo")); err != nil {
		if vv, ok := fortiAPIPatch(o["mu-mimo"], "ObjectWirelessControllerVap-MuMimo"); ok {
			if err = d.Set("mu_mimo", vv); err != nil {
				return fmt.Errorf("Error reading mu_mimo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mu_mimo: %v", err)
		}
	}

	if err = d.Set("multicast_enhance", flattenObjectWirelessControllerVapMulticastEnhance(o["multicast-enhance"], d, "multicast_enhance")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-enhance"], "ObjectWirelessControllerVap-MulticastEnhance"); ok {
			if err = d.Set("multicast_enhance", vv); err != nil {
				return fmt.Errorf("Error reading multicast_enhance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_enhance: %v", err)
		}
	}

	if err = d.Set("multicast_rate", flattenObjectWirelessControllerVapMulticastRate(o["multicast-rate"], d, "multicast_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-rate"], "ObjectWirelessControllerVap-MulticastRate"); ok {
			if err = d.Set("multicast_rate", vv); err != nil {
				return fmt.Errorf("Error reading multicast_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_rate: %v", err)
		}
	}

	if err = d.Set("nac", flattenObjectWirelessControllerVapNac(o["nac"], d, "nac")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac"], "ObjectWirelessControllerVap-Nac"); ok {
			if err = d.Set("nac", vv); err != nil {
				return fmt.Errorf("Error reading nac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac: %v", err)
		}
	}

	if err = d.Set("nac_profile", flattenObjectWirelessControllerVapNacProfile(o["nac-profile"], d, "nac_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-profile"], "ObjectWirelessControllerVap-NacProfile"); ok {
			if err = d.Set("nac_profile", vv); err != nil {
				return fmt.Errorf("Error reading nac_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_profile: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerVapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerVap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("neighbor_report_dual_band", flattenObjectWirelessControllerVapNeighborReportDualBand(o["neighbor-report-dual-band"], d, "neighbor_report_dual_band")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-report-dual-band"], "ObjectWirelessControllerVap-NeighborReportDualBand"); ok {
			if err = d.Set("neighbor_report_dual_band", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_report_dual_band: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_report_dual_band: %v", err)
		}
	}

	if err = d.Set("okc", flattenObjectWirelessControllerVapOkc(o["okc"], d, "okc")); err != nil {
		if vv, ok := fortiAPIPatch(o["okc"], "ObjectWirelessControllerVap-Okc"); ok {
			if err = d.Set("okc", vv); err != nil {
				return fmt.Errorf("Error reading okc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading okc: %v", err)
		}
	}

	if err = d.Set("osen", flattenObjectWirelessControllerVapOsen(o["osen"], d, "osen")); err != nil {
		if vv, ok := fortiAPIPatch(o["osen"], "ObjectWirelessControllerVap-Osen"); ok {
			if err = d.Set("osen", vv); err != nil {
				return fmt.Errorf("Error reading osen: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osen: %v", err)
		}
	}

	if err = d.Set("owe_groups", flattenObjectWirelessControllerVapOweGroups(o["owe-groups"], d, "owe_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["owe-groups"], "ObjectWirelessControllerVap-OweGroups"); ok {
			if err = d.Set("owe_groups", vv); err != nil {
				return fmt.Errorf("Error reading owe_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owe_groups: %v", err)
		}
	}

	if err = d.Set("owe_transition", flattenObjectWirelessControllerVapOweTransition(o["owe-transition"], d, "owe_transition")); err != nil {
		if vv, ok := fortiAPIPatch(o["owe-transition"], "ObjectWirelessControllerVap-OweTransition"); ok {
			if err = d.Set("owe_transition", vv); err != nil {
				return fmt.Errorf("Error reading owe_transition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owe_transition: %v", err)
		}
	}

	if err = d.Set("owe_transition_ssid", flattenObjectWirelessControllerVapOweTransitionSsid(o["owe-transition-ssid"], d, "owe_transition_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["owe-transition-ssid"], "ObjectWirelessControllerVap-OweTransitionSsid"); ok {
			if err = d.Set("owe_transition_ssid", vv); err != nil {
				return fmt.Errorf("Error reading owe_transition_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owe_transition_ssid: %v", err)
		}
	}

	if err = d.Set("pmf", flattenObjectWirelessControllerVapPmf(o["pmf"], d, "pmf")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmf"], "ObjectWirelessControllerVap-Pmf"); ok {
			if err = d.Set("pmf", vv); err != nil {
				return fmt.Errorf("Error reading pmf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmf: %v", err)
		}
	}

	if err = d.Set("pmf_assoc_comeback_timeout", flattenObjectWirelessControllerVapPmfAssocComebackTimeout(o["pmf-assoc-comeback-timeout"], d, "pmf_assoc_comeback_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmf-assoc-comeback-timeout"], "ObjectWirelessControllerVap-PmfAssocComebackTimeout"); ok {
			if err = d.Set("pmf_assoc_comeback_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pmf_assoc_comeback_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmf_assoc_comeback_timeout: %v", err)
		}
	}

	if err = d.Set("pmf_sa_query_retry_timeout", flattenObjectWirelessControllerVapPmfSaQueryRetryTimeout(o["pmf-sa-query-retry-timeout"], d, "pmf_sa_query_retry_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmf-sa-query-retry-timeout"], "ObjectWirelessControllerVap-PmfSaQueryRetryTimeout"); ok {
			if err = d.Set("pmf_sa_query_retry_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pmf_sa_query_retry_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmf_sa_query_retry_timeout: %v", err)
		}
	}

	if err = d.Set("port_macauth", flattenObjectWirelessControllerVapPortMacauth(o["port-macauth"], d, "port_macauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-macauth"], "ObjectWirelessControllerVap-PortMacauth"); ok {
			if err = d.Set("port_macauth", vv); err != nil {
				return fmt.Errorf("Error reading port_macauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_macauth: %v", err)
		}
	}

	if err = d.Set("port_macauth_reauth_timeout", flattenObjectWirelessControllerVapPortMacauthReauthTimeout(o["port-macauth-reauth-timeout"], d, "port_macauth_reauth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-macauth-reauth-timeout"], "ObjectWirelessControllerVap-PortMacauthReauthTimeout"); ok {
			if err = d.Set("port_macauth_reauth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading port_macauth_reauth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_macauth_reauth_timeout: %v", err)
		}
	}

	if err = d.Set("port_macauth_timeout", flattenObjectWirelessControllerVapPortMacauthTimeout(o["port-macauth-timeout"], d, "port_macauth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-macauth-timeout"], "ObjectWirelessControllerVap-PortMacauthTimeout"); ok {
			if err = d.Set("port_macauth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading port_macauth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_macauth_timeout: %v", err)
		}
	}

	if err = d.Set("portal_message_override_group", flattenObjectWirelessControllerVapPortalMessageOverrideGroup(o["portal-message-override-group"], d, "portal_message_override_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["portal-message-override-group"], "ObjectWirelessControllerVap-PortalMessageOverrideGroup"); ok {
			if err = d.Set("portal_message_override_group", vv); err != nil {
				return fmt.Errorf("Error reading portal_message_override_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portal_message_override_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("portal_message_overrides", flattenObjectWirelessControllerVapPortalMessageOverrides(o["portal-message-overrides"], d, "portal_message_overrides")); err != nil {
			if vv, ok := fortiAPIPatch(o["portal-message-overrides"], "ObjectWirelessControllerVap-PortalMessageOverrides"); ok {
				if err = d.Set("portal_message_overrides", vv); err != nil {
					return fmt.Errorf("Error reading portal_message_overrides: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading portal_message_overrides: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("portal_message_overrides"); ok {
			if err = d.Set("portal_message_overrides", flattenObjectWirelessControllerVapPortalMessageOverrides(o["portal-message-overrides"], d, "portal_message_overrides")); err != nil {
				if vv, ok := fortiAPIPatch(o["portal-message-overrides"], "ObjectWirelessControllerVap-PortalMessageOverrides"); ok {
					if err = d.Set("portal_message_overrides", vv); err != nil {
						return fmt.Errorf("Error reading portal_message_overrides: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading portal_message_overrides: %v", err)
				}
			}
		}
	}

	if err = d.Set("portal_type", flattenObjectWirelessControllerVapPortalType(o["portal-type"], d, "portal_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["portal-type"], "ObjectWirelessControllerVap-PortalType"); ok {
			if err = d.Set("portal_type", vv); err != nil {
				return fmt.Errorf("Error reading portal_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portal_type: %v", err)
		}
	}

	if err = d.Set("primary_wag_profile", flattenObjectWirelessControllerVapPrimaryWagProfile(o["primary-wag-profile"], d, "primary_wag_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-wag-profile"], "ObjectWirelessControllerVap-PrimaryWagProfile"); ok {
			if err = d.Set("primary_wag_profile", vv); err != nil {
				return fmt.Errorf("Error reading primary_wag_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_wag_profile: %v", err)
		}
	}

	if err = d.Set("probe_resp_suppression", flattenObjectWirelessControllerVapProbeRespSuppression(o["probe-resp-suppression"], d, "probe_resp_suppression")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-resp-suppression"], "ObjectWirelessControllerVap-ProbeRespSuppression"); ok {
			if err = d.Set("probe_resp_suppression", vv); err != nil {
				return fmt.Errorf("Error reading probe_resp_suppression: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_resp_suppression: %v", err)
		}
	}

	if err = d.Set("probe_resp_threshold", flattenObjectWirelessControllerVapProbeRespThreshold(o["probe-resp-threshold"], d, "probe_resp_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-resp-threshold"], "ObjectWirelessControllerVap-ProbeRespThreshold"); ok {
			if err = d.Set("probe_resp_threshold", vv); err != nil {
				return fmt.Errorf("Error reading probe_resp_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_resp_threshold: %v", err)
		}
	}

	if err = d.Set("ptk_rekey", flattenObjectWirelessControllerVapPtkRekey(o["ptk-rekey"], d, "ptk_rekey")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptk-rekey"], "ObjectWirelessControllerVap-PtkRekey"); ok {
			if err = d.Set("ptk_rekey", vv); err != nil {
				return fmt.Errorf("Error reading ptk_rekey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptk_rekey: %v", err)
		}
	}

	if err = d.Set("ptk_rekey_intv", flattenObjectWirelessControllerVapPtkRekeyIntv(o["ptk-rekey-intv"], d, "ptk_rekey_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptk-rekey-intv"], "ObjectWirelessControllerVap-PtkRekeyIntv"); ok {
			if err = d.Set("ptk_rekey_intv", vv); err != nil {
				return fmt.Errorf("Error reading ptk_rekey_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptk_rekey_intv: %v", err)
		}
	}

	if err = d.Set("qos_profile", flattenObjectWirelessControllerVapQosProfile(o["qos-profile"], d, "qos_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-profile"], "ObjectWirelessControllerVap-QosProfile"); ok {
			if err = d.Set("qos_profile", vv); err != nil {
				return fmt.Errorf("Error reading qos_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_profile: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenObjectWirelessControllerVapQuarantine(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "ObjectWirelessControllerVap-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("radio_2g_threshold", flattenObjectWirelessControllerVapRadio2GThreshold(o["radio-2g-threshold"], d, "radio_2g_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-2g-threshold"], "ObjectWirelessControllerVap-Radio2GThreshold"); ok {
			if err = d.Set("radio_2g_threshold", vv); err != nil {
				return fmt.Errorf("Error reading radio_2g_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_2g_threshold: %v", err)
		}
	}

	if err = d.Set("radio_5g_threshold", flattenObjectWirelessControllerVapRadio5GThreshold(o["radio-5g-threshold"], d, "radio_5g_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-5g-threshold"], "ObjectWirelessControllerVap-Radio5GThreshold"); ok {
			if err = d.Set("radio_5g_threshold", vv); err != nil {
				return fmt.Errorf("Error reading radio_5g_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_5g_threshold: %v", err)
		}
	}

	if err = d.Set("radio_sensitivity", flattenObjectWirelessControllerVapRadioSensitivity(o["radio-sensitivity"], d, "radio_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-sensitivity"], "ObjectWirelessControllerVap-RadioSensitivity"); ok {
			if err = d.Set("radio_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading radio_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_sensitivity: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth", flattenObjectWirelessControllerVapRadiusMacAuth(o["radius-mac-auth"], d, "radius_mac_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth"], "ObjectWirelessControllerVap-RadiusMacAuth"); ok {
			if err = d.Set("radius_mac_auth", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_block_interval", flattenObjectWirelessControllerVapRadiusMacAuthBlockInterval(o["radius-mac-auth-block-interval"], d, "radius_mac_auth_block_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-block-interval"], "ObjectWirelessControllerVap-RadiusMacAuthBlockInterval"); ok {
			if err = d.Set("radius_mac_auth_block_interval", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_block_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_block_interval: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_server", flattenObjectWirelessControllerVapRadiusMacAuthServer(o["radius-mac-auth-server"], d, "radius_mac_auth_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-server"], "ObjectWirelessControllerVap-RadiusMacAuthServer"); ok {
			if err = d.Set("radius_mac_auth_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_server: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_usergroups", flattenObjectWirelessControllerVapRadiusMacAuthUsergroups(o["radius-mac-auth-usergroups"], d, "radius_mac_auth_usergroups")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-usergroups"], "ObjectWirelessControllerVap-RadiusMacAuthUsergroups"); ok {
			if err = d.Set("radius_mac_auth_usergroups", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_usergroups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_usergroups: %v", err)
		}
	}

	if err = d.Set("radius_mac_mpsk_auth", flattenObjectWirelessControllerVapRadiusMacMpskAuth(o["radius-mac-mpsk-auth"], d, "radius_mac_mpsk_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-mpsk-auth"], "ObjectWirelessControllerVap-RadiusMacMpskAuth"); ok {
			if err = d.Set("radius_mac_mpsk_auth", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_mpsk_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_mpsk_auth: %v", err)
		}
	}

	if err = d.Set("radius_mac_mpsk_timeout", flattenObjectWirelessControllerVapRadiusMacMpskTimeout(o["radius-mac-mpsk-timeout"], d, "radius_mac_mpsk_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-mpsk-timeout"], "ObjectWirelessControllerVap-RadiusMacMpskTimeout"); ok {
			if err = d.Set("radius_mac_mpsk_timeout", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_mpsk_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_mpsk_timeout: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenObjectWirelessControllerVapRadiusServer(o["radius-server"], d, "radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-server"], "ObjectWirelessControllerVap-RadiusServer"); ok {
			if err = d.Set("radius_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("rates_11a", flattenObjectWirelessControllerVapRates11A(o["rates-11a"], d, "rates_11a")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11a"], "ObjectWirelessControllerVap-Rates11A"); ok {
			if err = d.Set("rates_11a", vv); err != nil {
				return fmt.Errorf("Error reading rates_11a: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11a: %v", err)
		}
	}

	if err = d.Set("rates_11ac_mcs_map", flattenObjectWirelessControllerVapRates11AcMcsMap(o["rates-11ac-mcs-map"], d, "rates_11ac_mcs_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ac-mcs-map"], "ObjectWirelessControllerVap-Rates11AcMcsMap"); ok {
			if err = d.Set("rates_11ac_mcs_map", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ac_mcs_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ac_mcs_map: %v", err)
		}
	}

	if err = d.Set("rates_11ac_ss12", flattenObjectWirelessControllerVapRates11AcSs12(o["rates-11ac-ss12"], d, "rates_11ac_ss12")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ac-ss12"], "ObjectWirelessControllerVap-Rates11AcSs12"); ok {
			if err = d.Set("rates_11ac_ss12", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ac_ss12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ac_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11ac_ss34", flattenObjectWirelessControllerVapRates11AcSs34(o["rates-11ac-ss34"], d, "rates_11ac_ss34")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ac-ss34"], "ObjectWirelessControllerVap-Rates11AcSs34"); ok {
			if err = d.Set("rates_11ac_ss34", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ac_ss34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ac_ss34: %v", err)
		}
	}

	if err = d.Set("rates_11ax_mcs_map", flattenObjectWirelessControllerVapRates11AxMcsMap(o["rates-11ax-mcs-map"], d, "rates_11ax_mcs_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ax-mcs-map"], "ObjectWirelessControllerVap-Rates11AxMcsMap"); ok {
			if err = d.Set("rates_11ax_mcs_map", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ax_mcs_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ax_mcs_map: %v", err)
		}
	}

	if err = d.Set("rates_11ax_ss12", flattenObjectWirelessControllerVapRates11AxSs12(o["rates-11ax-ss12"], d, "rates_11ax_ss12")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ax-ss12"], "ObjectWirelessControllerVap-Rates11AxSs12"); ok {
			if err = d.Set("rates_11ax_ss12", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ax_ss12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ax_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11ax_ss34", flattenObjectWirelessControllerVapRates11AxSs34(o["rates-11ax-ss34"], d, "rates_11ax_ss34")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ax-ss34"], "ObjectWirelessControllerVap-Rates11AxSs34"); ok {
			if err = d.Set("rates_11ax_ss34", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ax_ss34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ax_ss34: %v", err)
		}
	}

	if err = d.Set("rates_11bg", flattenObjectWirelessControllerVapRates11Bg(o["rates-11bg"], d, "rates_11bg")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11bg"], "ObjectWirelessControllerVap-Rates11Bg"); ok {
			if err = d.Set("rates_11bg", vv); err != nil {
				return fmt.Errorf("Error reading rates_11bg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11bg: %v", err)
		}
	}

	if err = d.Set("rates_11n_ss12", flattenObjectWirelessControllerVapRates11NSs12(o["rates-11n-ss12"], d, "rates_11n_ss12")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11n-ss12"], "ObjectWirelessControllerVap-Rates11NSs12"); ok {
			if err = d.Set("rates_11n_ss12", vv); err != nil {
				return fmt.Errorf("Error reading rates_11n_ss12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11n_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11n_ss34", flattenObjectWirelessControllerVapRates11NSs34(o["rates-11n-ss34"], d, "rates_11n_ss34")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11n-ss34"], "ObjectWirelessControllerVap-Rates11NSs34"); ok {
			if err = d.Set("rates_11n_ss34", vv); err != nil {
				return fmt.Errorf("Error reading rates_11n_ss34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11n_ss34: %v", err)
		}
	}

	if err = d.Set("sae_groups", flattenObjectWirelessControllerVapSaeGroups(o["sae-groups"], d, "sae_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-groups"], "ObjectWirelessControllerVap-SaeGroups"); ok {
			if err = d.Set("sae_groups", vv); err != nil {
				return fmt.Errorf("Error reading sae_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_groups: %v", err)
		}
	}

	if err = d.Set("sae_h2e_only", flattenObjectWirelessControllerVapSaeH2EOnly(o["sae-h2e-only"], d, "sae_h2e_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-h2e-only"], "ObjectWirelessControllerVap-SaeH2EOnly"); ok {
			if err = d.Set("sae_h2e_only", vv); err != nil {
				return fmt.Errorf("Error reading sae_h2e_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_h2e_only: %v", err)
		}
	}

	if err = d.Set("sae_password", flattenObjectWirelessControllerVapSaePassword(o["sae-password"], d, "sae_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-password"], "ObjectWirelessControllerVap-SaePassword"); ok {
			if err = d.Set("sae_password", vv); err != nil {
				return fmt.Errorf("Error reading sae_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_password: %v", err)
		}
	}

	if err = d.Set("sae_pk", flattenObjectWirelessControllerVapSaePk(o["sae-pk"], d, "sae_pk")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-pk"], "ObjectWirelessControllerVap-SaePk"); ok {
			if err = d.Set("sae_pk", vv); err != nil {
				return fmt.Errorf("Error reading sae_pk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_pk: %v", err)
		}
	}

	if err = d.Set("sae_private_key", flattenObjectWirelessControllerVapSaePrivateKey(o["sae-private-key"], d, "sae_private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-private-key"], "ObjectWirelessControllerVap-SaePrivateKey"); ok {
			if err = d.Set("sae_private_key", vv); err != nil {
				return fmt.Errorf("Error reading sae_private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_private_key: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenObjectWirelessControllerVapScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-botnet-connections"], "ObjectWirelessControllerVap-ScanBotnetConnections"); ok {
			if err = d.Set("scan_botnet_connections", vv); err != nil {
				return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("schedule", flattenObjectWirelessControllerVapSchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "ObjectWirelessControllerVap-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("secondary_wag_profile", flattenObjectWirelessControllerVapSecondaryWagProfile(o["secondary-wag-profile"], d, "secondary_wag_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-wag-profile"], "ObjectWirelessControllerVap-SecondaryWagProfile"); ok {
			if err = d.Set("secondary_wag_profile", vv); err != nil {
				return fmt.Errorf("Error reading secondary_wag_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_wag_profile: %v", err)
		}
	}

	if err = d.Set("security", flattenObjectWirelessControllerVapSecurity(o["security"], d, "security")); err != nil {
		if vv, ok := fortiAPIPatch(o["security"], "ObjectWirelessControllerVap-Security"); ok {
			if err = d.Set("security", vv); err != nil {
				return fmt.Errorf("Error reading security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("security_exempt_list", flattenObjectWirelessControllerVapSecurityExemptList(o["security-exempt-list"], d, "security_exempt_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-exempt-list"], "ObjectWirelessControllerVap-SecurityExemptList"); ok {
			if err = d.Set("security_exempt_list", vv); err != nil {
				return fmt.Errorf("Error reading security_exempt_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_exempt_list: %v", err)
		}
	}

	if err = d.Set("security_obsolete_option", flattenObjectWirelessControllerVapSecurityObsoleteOption(o["security-obsolete-option"], d, "security_obsolete_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-obsolete-option"], "ObjectWirelessControllerVap-SecurityObsoleteOption"); ok {
			if err = d.Set("security_obsolete_option", vv); err != nil {
				return fmt.Errorf("Error reading security_obsolete_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_obsolete_option: %v", err)
		}
	}

	if err = d.Set("security_redirect_url", flattenObjectWirelessControllerVapSecurityRedirectUrl(o["security-redirect-url"], d, "security_redirect_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-redirect-url"], "ObjectWirelessControllerVap-SecurityRedirectUrl"); ok {
			if err = d.Set("security_redirect_url", vv); err != nil {
				return fmt.Errorf("Error reading security_redirect_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_redirect_url: %v", err)
		}
	}

	if err = d.Set("selected_usergroups", flattenObjectWirelessControllerVapSelectedUsergroups(o["selected-usergroups"], d, "selected_usergroups")); err != nil {
		if vv, ok := fortiAPIPatch(o["selected-usergroups"], "ObjectWirelessControllerVap-SelectedUsergroups"); ok {
			if err = d.Set("selected_usergroups", vv); err != nil {
				return fmt.Errorf("Error reading selected_usergroups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading selected_usergroups: %v", err)
		}
	}

	if err = d.Set("split_tunneling", flattenObjectWirelessControllerVapSplitTunneling(o["split-tunneling"], d, "split_tunneling")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling"], "ObjectWirelessControllerVap-SplitTunneling"); ok {
			if err = d.Set("split_tunneling", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling: %v", err)
		}
	}

	if err = d.Set("ssid", flattenObjectWirelessControllerVapSsid(o["ssid"], d, "ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssid"], "ObjectWirelessControllerVap-Ssid"); ok {
			if err = d.Set("ssid", vv); err != nil {
				return fmt.Errorf("Error reading ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssid: %v", err)
		}
	}

	if err = d.Set("sticky_client_remove", flattenObjectWirelessControllerVapStickyClientRemove(o["sticky-client-remove"], d, "sticky_client_remove")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-remove"], "ObjectWirelessControllerVap-StickyClientRemove"); ok {
			if err = d.Set("sticky_client_remove", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_remove: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_remove: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_2g", flattenObjectWirelessControllerVapStickyClientThreshold2G(o["sticky-client-threshold-2g"], d, "sticky_client_threshold_2g")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-threshold-2g"], "ObjectWirelessControllerVap-StickyClientThreshold2G"); ok {
			if err = d.Set("sticky_client_threshold_2g", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_threshold_2g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_threshold_2g: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_5g", flattenObjectWirelessControllerVapStickyClientThreshold5G(o["sticky-client-threshold-5g"], d, "sticky_client_threshold_5g")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-threshold-5g"], "ObjectWirelessControllerVap-StickyClientThreshold5G"); ok {
			if err = d.Set("sticky_client_threshold_5g", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_threshold_5g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_threshold_5g: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_6g", flattenObjectWirelessControllerVapStickyClientThreshold6G(o["sticky-client-threshold-6g"], d, "sticky_client_threshold_6g")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-threshold-6g"], "ObjectWirelessControllerVap-StickyClientThreshold6G"); ok {
			if err = d.Set("sticky_client_threshold_6g", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_threshold_6g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_threshold_6g: %v", err)
		}
	}

	if err = d.Set("target_wake_time", flattenObjectWirelessControllerVapTargetWakeTime(o["target-wake-time"], d, "target_wake_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["target-wake-time"], "ObjectWirelessControllerVap-TargetWakeTime"); ok {
			if err = d.Set("target_wake_time", vv); err != nil {
				return fmt.Errorf("Error reading target_wake_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target_wake_time: %v", err)
		}
	}

	if err = d.Set("tkip_counter_measure", flattenObjectWirelessControllerVapTkipCounterMeasure(o["tkip-counter-measure"], d, "tkip_counter_measure")); err != nil {
		if vv, ok := fortiAPIPatch(o["tkip-counter-measure"], "ObjectWirelessControllerVap-TkipCounterMeasure"); ok {
			if err = d.Set("tkip_counter_measure", vv); err != nil {
				return fmt.Errorf("Error reading tkip_counter_measure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tkip_counter_measure: %v", err)
		}
	}

	if err = d.Set("tunnel_echo_interval", flattenObjectWirelessControllerVapTunnelEchoInterval(o["tunnel-echo-interval"], d, "tunnel_echo_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-echo-interval"], "ObjectWirelessControllerVap-TunnelEchoInterval"); ok {
			if err = d.Set("tunnel_echo_interval", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_echo_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_echo_interval: %v", err)
		}
	}

	if err = d.Set("tunnel_fallback_interval", flattenObjectWirelessControllerVapTunnelFallbackInterval(o["tunnel-fallback-interval"], d, "tunnel_fallback_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-fallback-interval"], "ObjectWirelessControllerVap-TunnelFallbackInterval"); ok {
			if err = d.Set("tunnel_fallback_interval", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_fallback_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_fallback_interval: %v", err)
		}
	}

	if err = d.Set("usergroup", flattenObjectWirelessControllerVapUsergroup(o["usergroup"], d, "usergroup")); err != nil {
		if vv, ok := fortiAPIPatch(o["usergroup"], "ObjectWirelessControllerVap-Usergroup"); ok {
			if err = d.Set("usergroup", vv); err != nil {
				return fmt.Errorf("Error reading usergroup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usergroup: %v", err)
		}
	}

	if err = d.Set("utm_log", flattenObjectWirelessControllerVapUtmLog(o["utm-log"], d, "utm_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-log"], "ObjectWirelessControllerVap-UtmLog"); ok {
			if err = d.Set("utm_log", vv); err != nil {
				return fmt.Errorf("Error reading utm_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_log: %v", err)
		}
	}

	if err = d.Set("utm_profile", flattenObjectWirelessControllerVapUtmProfile(o["utm-profile"], d, "utm_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-profile"], "ObjectWirelessControllerVap-UtmProfile"); ok {
			if err = d.Set("utm_profile", vv); err != nil {
				return fmt.Errorf("Error reading utm_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_profile: %v", err)
		}
	}

	if err = d.Set("vdom", flattenObjectWirelessControllerVapVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "ObjectWirelessControllerVap-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenObjectWirelessControllerVapUtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-status"], "ObjectWirelessControllerVap-UtmStatus"); ok {
			if err = d.Set("utm_status", vv); err != nil {
				return fmt.Errorf("Error reading utm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("vlan_auto", flattenObjectWirelessControllerVapVlanAuto(o["vlan-auto"], d, "vlan_auto")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-auto"], "ObjectWirelessControllerVap-VlanAuto"); ok {
			if err = d.Set("vlan_auto", vv); err != nil {
				return fmt.Errorf("Error reading vlan_auto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_auto: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vlan_name", flattenObjectWirelessControllerVapVlanName(o["vlan-name"], d, "vlan_name")); err != nil {
			if vv, ok := fortiAPIPatch(o["vlan-name"], "ObjectWirelessControllerVap-VlanName"); ok {
				if err = d.Set("vlan_name", vv); err != nil {
					return fmt.Errorf("Error reading vlan_name: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vlan_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan_name"); ok {
			if err = d.Set("vlan_name", flattenObjectWirelessControllerVapVlanName(o["vlan-name"], d, "vlan_name")); err != nil {
				if vv, ok := fortiAPIPatch(o["vlan-name"], "ObjectWirelessControllerVap-VlanName"); ok {
					if err = d.Set("vlan_name", vv); err != nil {
						return fmt.Errorf("Error reading vlan_name: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vlan_name: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("vlan_pool", flattenObjectWirelessControllerVapVlanPool(o["vlan-pool"], d, "vlan_pool")); err != nil {
			if vv, ok := fortiAPIPatch(o["vlan-pool"], "ObjectWirelessControllerVap-VlanPool"); ok {
				if err = d.Set("vlan_pool", vv); err != nil {
					return fmt.Errorf("Error reading vlan_pool: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vlan_pool: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan_pool"); ok {
			if err = d.Set("vlan_pool", flattenObjectWirelessControllerVapVlanPool(o["vlan-pool"], d, "vlan_pool")); err != nil {
				if vv, ok := fortiAPIPatch(o["vlan-pool"], "ObjectWirelessControllerVap-VlanPool"); ok {
					if err = d.Set("vlan_pool", vv); err != nil {
						return fmt.Errorf("Error reading vlan_pool: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vlan_pool: %v", err)
				}
			}
		}
	}

	if err = d.Set("vlan_pooling", flattenObjectWirelessControllerVapVlanPooling(o["vlan-pooling"], d, "vlan_pooling")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-pooling"], "ObjectWirelessControllerVap-VlanPooling"); ok {
			if err = d.Set("vlan_pooling", vv); err != nil {
				return fmt.Errorf("Error reading vlan_pooling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_pooling: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenObjectWirelessControllerVapVlanid(o["vlanid"], d, "vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlanid"], "ObjectWirelessControllerVap-Vlanid"); ok {
			if err = d.Set("vlanid", vv); err != nil {
				return fmt.Errorf("Error reading vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("voice_enterprise", flattenObjectWirelessControllerVapVoiceEnterprise(o["voice-enterprise"], d, "voice_enterprise")); err != nil {
		if vv, ok := fortiAPIPatch(o["voice-enterprise"], "ObjectWirelessControllerVap-VoiceEnterprise"); ok {
			if err = d.Set("voice_enterprise", vv); err != nil {
				return fmt.Errorf("Error reading voice_enterprise: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voice_enterprise: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenObjectWirelessControllerVapWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "ObjectWirelessControllerVap-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerVapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerVapCentmgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDhcpSvrId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapIntfDeviceAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfDeviceIdentification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfDeviceNetscan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfDhcpRelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapIntfDhcpRelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfDhcpRelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfDhcp6RelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfDhcp6RelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfDhcp6RelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntfIp6Allowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapIntfListenForticlientConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIsFactorySetting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapAccessControlList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapAdditionalAkms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapAddressGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapAddressGroupPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapAntivirusProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapApplicationDetectionEngine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapApplicationDscpMarking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapApplicationReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapAtfWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapAuthCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapAuthPortalAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapBeaconAdvertising(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapBroadcastSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapBroadcastSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapBssColorPartial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapBstmDisassociationImminent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapBstmLoadBalancingDisassocTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapBstmRssiDisassocTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapCaptivePortalAcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapCaptivePortalAuthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapCaptivePortalFwAccounting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDhcpAddressEnforcement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapCaptivePortalMacauthRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapCaptivePortalMacauthRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapCaptivePortalRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapCaptivePortalRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapCaptivePortalSessionTimeoutInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDhcpLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDhcpOption43Insertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDhcpOption82CircuitIdInsertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDhcpOption82Insertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDhcpOption82RemoteIdInsertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_centmgmt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_centmgmt"], _ = expandObjectWirelessControllerVapDynamicMappingCentmgmt(d, i["_centmgmt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_dhcp_svr_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_dhcp_svr_id"], _ = expandObjectWirelessControllerVapDynamicMappingDhcpSvrId(d, i["_dhcp_svr_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_allowaccess"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_allowaccess"], _ = expandObjectWirelessControllerVapDynamicMappingIntfAllowaccess(d, i["_intf_allowaccess"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_access_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_device-access-list"], _ = expandObjectWirelessControllerVapDynamicMappingIntfDeviceAccessList(d, i["_intf_device_access_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_identification"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_device-identification"], _ = expandObjectWirelessControllerVapDynamicMappingIntfDeviceIdentification(d, i["_intf_device_identification"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_device_netscan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_device-netscan"], _ = expandObjectWirelessControllerVapDynamicMappingIntfDeviceNetscan(d, i["_intf_device_netscan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp-relay-ip"], _ = expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayIp(d, i["_intf_dhcp_relay_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp-relay-service"], _ = expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayService(d, i["_intf_dhcp_relay_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp_relay_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp-relay-type"], _ = expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayType(d, i["_intf_dhcp_relay_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp6-relay-ip"], _ = expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayIp(d, i["_intf_dhcp6_relay_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp6-relay-service"], _ = expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayService(d, i["_intf_dhcp6_relay_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_dhcp6_relay_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_dhcp6-relay-type"], _ = expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayType(d, i["_intf_dhcp6_relay_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_ip"], _ = expandObjectWirelessControllerVapDynamicMappingIntfIp(d, i["_intf_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip6_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_ip6-address"], _ = expandObjectWirelessControllerVapDynamicMappingIntfIp6Address(d, i["_intf_ip6_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_ip6_allowaccess"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_ip6-allowaccess"], _ = expandObjectWirelessControllerVapDynamicMappingIntfIp6Allowaccess(d, i["_intf_ip6_allowaccess"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_intf_listen_forticlient_connection"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_intf_listen-forticlient-connection"], _ = expandObjectWirelessControllerVapDynamicMappingIntfListenForticlientConnection(d, i["_intf_listen_forticlient_connection"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_is_factory_setting"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_is_factory_setting"], _ = expandObjectWirelessControllerVapDynamicMappingIsFactorySetting(d, i["_is_factory_setting"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectWirelessControllerVapDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_control_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["access-control-list"], _ = expandObjectWirelessControllerVapDynamicMappingAccessControlList(d, i["access_control_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_interim_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["acct-interim-interval"], _ = expandObjectWirelessControllerVapDynamicMappingAcctInterimInterval(d, i["acct_interim_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_akms"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-akms"], _ = expandObjectWirelessControllerVapDynamicMappingAdditionalAkms(d, i["additional_akms"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address-group"], _ = expandObjectWirelessControllerVapDynamicMappingAddressGroup(d, i["address_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address_group_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address-group-policy"], _ = expandObjectWirelessControllerVapDynamicMappingAddressGroupPolicy(d, i["address_group_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alias"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["alias"], _ = expandObjectWirelessControllerVapDynamicMappingAlias(d, i["alias"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antivirus_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["antivirus-profile"], _ = expandObjectWirelessControllerVapDynamicMappingAntivirusProfile(d, i["antivirus_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_detection_engine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application-detection-engine"], _ = expandObjectWirelessControllerVapDynamicMappingApplicationDetectionEngine(d, i["application_detection_engine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_dscp_marking"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application-dscp-marking"], _ = expandObjectWirelessControllerVapDynamicMappingApplicationDscpMarking(d, i["application_dscp_marking"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application-list"], _ = expandObjectWirelessControllerVapDynamicMappingApplicationList(d, i["application_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_report_intv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application-report-intv"], _ = expandObjectWirelessControllerVapDynamicMappingApplicationReportIntv(d, i["application_report_intv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "atf_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["atf-weight"], _ = expandObjectWirelessControllerVapDynamicMappingAtfWeight(d, i["atf_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth"], _ = expandObjectWirelessControllerVapDynamicMappingAuth(d, i["auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-cert"], _ = expandObjectWirelessControllerVapDynamicMappingAuthCert(d, i["auth_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_portal_addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-portal-addr"], _ = expandObjectWirelessControllerVapDynamicMappingAuthPortalAddr(d, i["auth_portal_addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "beacon_advertising"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["beacon-advertising"], _ = expandObjectWirelessControllerVapDynamicMappingBeaconAdvertising(d, i["beacon_advertising"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "broadcast_ssid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["broadcast-ssid"], _ = expandObjectWirelessControllerVapDynamicMappingBroadcastSsid(d, i["broadcast_ssid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "broadcast_suppression"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["broadcast-suppression"], _ = expandObjectWirelessControllerVapDynamicMappingBroadcastSuppression(d, i["broadcast_suppression"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bss_color_partial"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bss-color-partial"], _ = expandObjectWirelessControllerVapDynamicMappingBssColorPartial(d, i["bss_color_partial"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_disassociation_imminent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bstm-disassociation-imminent"], _ = expandObjectWirelessControllerVapDynamicMappingBstmDisassociationImminent(d, i["bstm_disassociation_imminent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_load_balancing_disassoc_timer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bstm-load-balancing-disassoc-timer"], _ = expandObjectWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer(d, i["bstm_load_balancing_disassoc_timer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bstm_rssi_disassoc_timer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bstm-rssi-disassoc-timer"], _ = expandObjectWirelessControllerVapDynamicMappingBstmRssiDisassocTimer(d, i["bstm_rssi_disassoc_timer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_ac_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-ac-name"], _ = expandObjectWirelessControllerVapDynamicMappingCaptivePortalAcName(d, i["captive_portal_ac_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_auth_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-auth-timeout"], _ = expandObjectWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout(d, i["captive_portal_auth_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_fw_accounting"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-fw-accounting"], _ = expandObjectWirelessControllerVapDynamicMappingCaptivePortalFwAccounting(d, i["captive_portal_fw_accounting"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_macauth_radius_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-macauth-radius-secret"], _ = expandObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusSecret(d, i["captive_portal_macauth_radius_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_macauth_radius_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-macauth-radius-server"], _ = expandObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer(d, i["captive_portal_macauth_radius_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_radius_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-radius-secret"], _ = expandObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusSecret(d, i["captive_portal_radius_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_radius_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-radius-server"], _ = expandObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusServer(d, i["captive_portal_radius_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "captive_portal_session_timeout_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["captive-portal-session-timeout-interval"], _ = expandObjectWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval(d, i["captive_portal_session_timeout_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-count"], _ = expandObjectWirelessControllerVapDynamicMappingClientCount(d, i["client_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_address_enforcement"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-address-enforcement"], _ = expandObjectWirelessControllerVapDynamicMappingDhcpAddressEnforcement(d, i["dhcp_address_enforcement"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-lease-time"], _ = expandObjectWirelessControllerVapDynamicMappingDhcpLeaseTime(d, i["dhcp_lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option43_insertion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-option43-insertion"], _ = expandObjectWirelessControllerVapDynamicMappingDhcpOption43Insertion(d, i["dhcp_option43_insertion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_circuit_id_insertion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-option82-circuit-id-insertion"], _ = expandObjectWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion(d, i["dhcp_option82_circuit_id_insertion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_insertion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-option82-insertion"], _ = expandObjectWirelessControllerVapDynamicMappingDhcpOption82Insertion(d, i["dhcp_option82_insertion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_option82_remote_id_insertion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-option82-remote-id-insertion"], _ = expandObjectWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion(d, i["dhcp_option82_remote_id_insertion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dynamic_vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dynamic-vlan"], _ = expandObjectWirelessControllerVapDynamicMappingDynamicVlan(d, i["dynamic_vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_reauth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["eap-reauth"], _ = expandObjectWirelessControllerVapDynamicMappingEapReauth(d, i["eap_reauth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_reauth_intv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["eap-reauth-intv"], _ = expandObjectWirelessControllerVapDynamicMappingEapReauthIntv(d, i["eap_reauth_intv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eapol_key_retries"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["eapol-key-retries"], _ = expandObjectWirelessControllerVapDynamicMappingEapolKeyRetries(d, i["eapol_key_retries"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encrypt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["encrypt"], _ = expandObjectWirelessControllerVapDynamicMappingEncrypt(d, i["encrypt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_fast_roaming"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-fast-roaming"], _ = expandObjectWirelessControllerVapDynamicMappingExternalFastRoaming(d, i["external_fast_roaming"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_logout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-logout"], _ = expandObjectWirelessControllerVapDynamicMappingExternalLogout(d, i["external_logout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_web"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-web"], _ = expandObjectWirelessControllerVapDynamicMappingExternalWeb(d, i["external_web"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_web_format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-web-format"], _ = expandObjectWirelessControllerVapDynamicMappingExternalWebFormat(d, i["external_web_format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fast_bss_transition"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fast-bss-transition"], _ = expandObjectWirelessControllerVapDynamicMappingFastBssTransition(d, i["fast_bss_transition"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fast_roaming"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fast-roaming"], _ = expandObjectWirelessControllerVapDynamicMappingFastRoaming(d, i["fast_roaming"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_mobility_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ft-mobility-domain"], _ = expandObjectWirelessControllerVapDynamicMappingFtMobilityDomain(d, i["ft_mobility_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_over_ds"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ft-over-ds"], _ = expandObjectWirelessControllerVapDynamicMappingFtOverDs(d, i["ft_over_ds"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ft_r0_key_lifetime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ft-r0-key-lifetime"], _ = expandObjectWirelessControllerVapDynamicMappingFtR0KeyLifetime(d, i["ft_r0_key_lifetime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gas_comeback_delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gas-comeback-delay"], _ = expandObjectWirelessControllerVapDynamicMappingGasComebackDelay(d, i["gas_comeback_delay"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gas_fragmentation_limit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gas-fragmentation-limit"], _ = expandObjectWirelessControllerVapDynamicMappingGasFragmentationLimit(d, i["gas_fragmentation_limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gtk_rekey"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gtk-rekey"], _ = expandObjectWirelessControllerVapDynamicMappingGtkRekey(d, i["gtk_rekey"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gtk_rekey_intv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gtk-rekey-intv"], _ = expandObjectWirelessControllerVapDynamicMappingGtkRekeyIntv(d, i["gtk_rekey_intv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "high_efficiency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["high-efficiency"], _ = expandObjectWirelessControllerVapDynamicMappingHighEfficiency(d, i["high_efficiency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hotspot20_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hotspot20-profile"], _ = expandObjectWirelessControllerVapDynamicMappingHotspot20Profile(d, i["hotspot20_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["igmp-snooping"], _ = expandObjectWirelessControllerVapDynamicMappingIgmpSnooping(d, i["igmp_snooping"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intra_vap_privacy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["intra-vap-privacy"], _ = expandObjectWirelessControllerVapDynamicMappingIntraVapPrivacy(d, i["intra_vap_privacy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectWirelessControllerVapDynamicMappingIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ips_sensor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ips-sensor"], _ = expandObjectWirelessControllerVapDynamicMappingIpsSensor(d, i["ips_sensor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_rules"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-rules"], _ = expandObjectWirelessControllerVapDynamicMappingIpv6Rules(d, i["ipv6_rules"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key"], _ = expandObjectWirelessControllerVapDynamicMappingKey(d, i["key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyindex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keyindex"], _ = expandObjectWirelessControllerVapDynamicMappingKeyindex(d, i["keyindex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "l3_roaming"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["l3-roaming"], _ = expandObjectWirelessControllerVapDynamicMappingL3Roaming(d, i["l3_roaming"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "l3_roaming_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["l3-roaming-mode"], _ = expandObjectWirelessControllerVapDynamicMappingL3RoamingMode(d, i["l3_roaming_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldpc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldpc"], _ = expandObjectWirelessControllerVapDynamicMappingLdpc(d, i["ldpc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-authentication"], _ = expandObjectWirelessControllerVapDynamicMappingLocalAuthentication(d, i["local_authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_bridging"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-bridging"], _ = expandObjectWirelessControllerVapDynamicMappingLocalBridging(d, i["local_bridging"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_lan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-lan"], _ = expandObjectWirelessControllerVapDynamicMappingLocalLan(d, i["local_lan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-standalone"], _ = expandObjectWirelessControllerVapDynamicMappingLocalStandalone(d, i["local_standalone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_dns"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-standalone-dns"], _ = expandObjectWirelessControllerVapDynamicMappingLocalStandaloneDns(d, i["local_standalone_dns"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_dns_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-standalone-dns-ip"], _ = expandObjectWirelessControllerVapDynamicMappingLocalStandaloneDnsIp(d, i["local_standalone_dns_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_standalone_nat"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-standalone-nat"], _ = expandObjectWirelessControllerVapDynamicMappingLocalStandaloneNat(d, i["local_standalone_nat"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_switching"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-switching"], _ = expandObjectWirelessControllerVapDynamicMappingLocalSwitching(d, i["local_switching"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_auth_bypass"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-auth-bypass"], _ = expandObjectWirelessControllerVapDynamicMappingMacAuthBypass(d, i["mac_auth_bypass"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_called_station_delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-called-station-delimiter"], _ = expandObjectWirelessControllerVapDynamicMappingMacCalledStationDelimiter(d, i["mac_called_station_delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_calling_station_delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-calling-station-delimiter"], _ = expandObjectWirelessControllerVapDynamicMappingMacCallingStationDelimiter(d, i["mac_calling_station_delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_case"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-case"], _ = expandObjectWirelessControllerVapDynamicMappingMacCase(d, i["mac_case"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-filter"], _ = expandObjectWirelessControllerVapDynamicMappingMacFilter(d, i["mac_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy_other"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-filter-policy-other"], _ = expandObjectWirelessControllerVapDynamicMappingMacFilterPolicyOther(d, i["mac_filter_policy_other"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_password_delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-password-delimiter"], _ = expandObjectWirelessControllerVapDynamicMappingMacPasswordDelimiter(d, i["mac_password_delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_username_delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-username-delimiter"], _ = expandObjectWirelessControllerVapDynamicMappingMacUsernameDelimiter(d, i["mac_username_delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_clients"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-clients"], _ = expandObjectWirelessControllerVapDynamicMappingMaxClients(d, i["max_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_clients_ap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-clients-ap"], _ = expandObjectWirelessControllerVapDynamicMappingMaxClientsAp(d, i["max_clients_ap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mbo"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mbo"], _ = expandObjectWirelessControllerVapDynamicMappingMbo(d, i["mbo"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mbo_cell_data_conn_pref"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mbo-cell-data-conn-pref"], _ = expandObjectWirelessControllerVapDynamicMappingMboCellDataConnPref(d, i["mbo_cell_data_conn_pref"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "me_disable_thresh"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["me-disable-thresh"], _ = expandObjectWirelessControllerVapDynamicMappingMeDisableThresh(d, i["me_disable_thresh"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mesh_backhaul"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mesh-backhaul"], _ = expandObjectWirelessControllerVapDynamicMappingMeshBackhaul(d, i["mesh_backhaul"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk"], _ = expandObjectWirelessControllerVapDynamicMappingMpsk(d, i["mpsk"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-concurrent-clients"], _ = expandObjectWirelessControllerVapDynamicMappingMpskConcurrentClients(d, i["mpsk_concurrent_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-profile"], _ = expandObjectWirelessControllerVapDynamicMappingMpskProfile(d, i["mpsk_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mu_mimo"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mu-mimo"], _ = expandObjectWirelessControllerVapDynamicMappingMuMimo(d, i["mu_mimo"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_enhance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["multicast-enhance"], _ = expandObjectWirelessControllerVapDynamicMappingMulticastEnhance(d, i["multicast_enhance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_rate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["multicast-rate"], _ = expandObjectWirelessControllerVapDynamicMappingMulticastRate(d, i["multicast_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nac"], _ = expandObjectWirelessControllerVapDynamicMappingNac(d, i["nac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nac_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nac-profile"], _ = expandObjectWirelessControllerVapDynamicMappingNacProfile(d, i["nac_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_report_dual_band"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["neighbor-report-dual-band"], _ = expandObjectWirelessControllerVapDynamicMappingNeighborReportDualBand(d, i["neighbor_report_dual_band"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "okc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["okc"], _ = expandObjectWirelessControllerVapDynamicMappingOkc(d, i["okc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "osen"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["osen"], _ = expandObjectWirelessControllerVapDynamicMappingOsen(d, i["osen"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["owe-groups"], _ = expandObjectWirelessControllerVapDynamicMappingOweGroups(d, i["owe_groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_transition"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["owe-transition"], _ = expandObjectWirelessControllerVapDynamicMappingOweTransition(d, i["owe_transition"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "owe_transition_ssid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["owe-transition-ssid"], _ = expandObjectWirelessControllerVapDynamicMappingOweTransitionSsid(d, i["owe_transition_ssid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passphrase"], _ = expandObjectWirelessControllerVapDynamicMappingPassphrase(d, i["passphrase"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pmf"], _ = expandObjectWirelessControllerVapDynamicMappingPmf(d, i["pmf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf_assoc_comeback_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pmf-assoc-comeback-timeout"], _ = expandObjectWirelessControllerVapDynamicMappingPmfAssocComebackTimeout(d, i["pmf_assoc_comeback_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmf_sa_query_retry_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pmf-sa-query-retry-timeout"], _ = expandObjectWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout(d, i["pmf_sa_query_retry_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-macauth"], _ = expandObjectWirelessControllerVapDynamicMappingPortMacauth(d, i["port_macauth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth_reauth_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-macauth-reauth-timeout"], _ = expandObjectWirelessControllerVapDynamicMappingPortMacauthReauthTimeout(d, i["port_macauth_reauth_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_macauth_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-macauth-timeout"], _ = expandObjectWirelessControllerVapDynamicMappingPortMacauthTimeout(d, i["port_macauth_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal_message_override_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portal-message-override-group"], _ = expandObjectWirelessControllerVapDynamicMappingPortalMessageOverrideGroup(d, i["portal_message_override_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portal-type"], _ = expandObjectWirelessControllerVapDynamicMappingPortalType(d, i["portal_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "primary_wag_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["primary-wag-profile"], _ = expandObjectWirelessControllerVapDynamicMappingPrimaryWagProfile(d, i["primary_wag_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_resp_suppression"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-resp-suppression"], _ = expandObjectWirelessControllerVapDynamicMappingProbeRespSuppression(d, i["probe_resp_suppression"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_resp_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-resp-threshold"], _ = expandObjectWirelessControllerVapDynamicMappingProbeRespThreshold(d, i["probe_resp_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptk_rekey"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ptk-rekey"], _ = expandObjectWirelessControllerVapDynamicMappingPtkRekey(d, i["ptk_rekey"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptk_rekey_intv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ptk-rekey-intv"], _ = expandObjectWirelessControllerVapDynamicMappingPtkRekeyIntv(d, i["ptk_rekey_intv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["qos-profile"], _ = expandObjectWirelessControllerVapDynamicMappingQosProfile(d, i["qos_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandObjectWirelessControllerVapDynamicMappingQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_2g_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radio-2g-threshold"], _ = expandObjectWirelessControllerVapDynamicMappingRadio2GThreshold(d, i["radio_2g_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_5g_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radio-5g-threshold"], _ = expandObjectWirelessControllerVapDynamicMappingRadio5GThreshold(d, i["radio_5g_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radio_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radio-sensitivity"], _ = expandObjectWirelessControllerVapDynamicMappingRadioSensitivity(d, i["radio_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-auth"], _ = expandObjectWirelessControllerVapDynamicMappingRadiusMacAuth(d, i["radius_mac_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_block_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-auth-block-interval"], _ = expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval(d, i["radius_mac_auth_block_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-auth-server"], _ = expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthServer(d, i["radius_mac_auth_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_auth_usergroups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-auth-usergroups"], _ = expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups(d, i["radius_mac_auth_usergroups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_mpsk_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-mpsk-auth"], _ = expandObjectWirelessControllerVapDynamicMappingRadiusMacMpskAuth(d, i["radius_mac_mpsk_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_mac_mpsk_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-mac-mpsk-timeout"], _ = expandObjectWirelessControllerVapDynamicMappingRadiusMacMpskTimeout(d, i["radius_mac_mpsk_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-server"], _ = expandObjectWirelessControllerVapDynamicMappingRadiusServer(d, i["radius_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11a"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11a"], _ = expandObjectWirelessControllerVapDynamicMappingRates11A(d, i["rates_11a"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_mcs_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ac-mcs-map"], _ = expandObjectWirelessControllerVapDynamicMappingRates11AcMcsMap(d, i["rates_11ac_mcs_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_ss12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ac-ss12"], _ = expandObjectWirelessControllerVapDynamicMappingRates11AcSs12(d, i["rates_11ac_ss12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ac_ss34"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ac-ss34"], _ = expandObjectWirelessControllerVapDynamicMappingRates11AcSs34(d, i["rates_11ac_ss34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_mcs_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ax-mcs-map"], _ = expandObjectWirelessControllerVapDynamicMappingRates11AxMcsMap(d, i["rates_11ax_mcs_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_ss12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ax-ss12"], _ = expandObjectWirelessControllerVapDynamicMappingRates11AxSs12(d, i["rates_11ax_ss12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11ax_ss34"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11ax-ss34"], _ = expandObjectWirelessControllerVapDynamicMappingRates11AxSs34(d, i["rates_11ax_ss34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11bg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11bg"], _ = expandObjectWirelessControllerVapDynamicMappingRates11Bg(d, i["rates_11bg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11n_ss12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11n-ss12"], _ = expandObjectWirelessControllerVapDynamicMappingRates11NSs12(d, i["rates_11n_ss12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rates_11n_ss34"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rates-11n-ss34"], _ = expandObjectWirelessControllerVapDynamicMappingRates11NSs34(d, i["rates_11n_ss34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-groups"], _ = expandObjectWirelessControllerVapDynamicMappingSaeGroups(d, i["sae_groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_h2e_only"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-h2e-only"], _ = expandObjectWirelessControllerVapDynamicMappingSaeH2EOnly(d, i["sae_h2e_only"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-password"], _ = expandObjectWirelessControllerVapDynamicMappingSaePassword(d, i["sae_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_pk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-pk"], _ = expandObjectWirelessControllerVapDynamicMappingSaePk(d, i["sae_pk"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_private_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-private-key"], _ = expandObjectWirelessControllerVapDynamicMappingSaePrivateKey(d, i["sae_private_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scan_botnet_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["scan-botnet-connections"], _ = expandObjectWirelessControllerVapDynamicMappingScanBotnetConnections(d, i["scan_botnet_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "schedule"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["schedule"], _ = expandObjectWirelessControllerVapDynamicMappingSchedule(d, i["schedule"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_wag_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secondary-wag-profile"], _ = expandObjectWirelessControllerVapDynamicMappingSecondaryWagProfile(d, i["secondary_wag_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security"], _ = expandObjectWirelessControllerVapDynamicMappingSecurity(d, i["security"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_exempt_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-exempt-list"], _ = expandObjectWirelessControllerVapDynamicMappingSecurityExemptList(d, i["security_exempt_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_obsolete_option"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-obsolete-option"], _ = expandObjectWirelessControllerVapDynamicMappingSecurityObsoleteOption(d, i["security_obsolete_option"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_redirect_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-redirect-url"], _ = expandObjectWirelessControllerVapDynamicMappingSecurityRedirectUrl(d, i["security_redirect_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "selected_usergroups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["selected-usergroups"], _ = expandObjectWirelessControllerVapDynamicMappingSelectedUsergroups(d, i["selected_usergroups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_tunneling"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["split-tunneling"], _ = expandObjectWirelessControllerVapDynamicMappingSplitTunneling(d, i["split_tunneling"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssid"], _ = expandObjectWirelessControllerVapDynamicMappingSsid(d, i["ssid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_remove"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sticky-client-remove"], _ = expandObjectWirelessControllerVapDynamicMappingStickyClientRemove(d, i["sticky_client_remove"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_2g"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sticky-client-threshold-2g"], _ = expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold2G(d, i["sticky_client_threshold_2g"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_5g"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sticky-client-threshold-5g"], _ = expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold5G(d, i["sticky_client_threshold_5g"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_client_threshold_6g"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sticky-client-threshold-6g"], _ = expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold6G(d, i["sticky_client_threshold_6g"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target_wake_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target-wake-time"], _ = expandObjectWirelessControllerVapDynamicMappingTargetWakeTime(d, i["target_wake_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tkip_counter_measure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tkip-counter-measure"], _ = expandObjectWirelessControllerVapDynamicMappingTkipCounterMeasure(d, i["tkip_counter_measure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_echo_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tunnel-echo-interval"], _ = expandObjectWirelessControllerVapDynamicMappingTunnelEchoInterval(d, i["tunnel_echo_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_fallback_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tunnel-fallback-interval"], _ = expandObjectWirelessControllerVapDynamicMappingTunnelFallbackInterval(d, i["tunnel_fallback_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "usergroup"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["usergroup"], _ = expandObjectWirelessControllerVapDynamicMappingUsergroup(d, i["usergroup"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["utm-log"], _ = expandObjectWirelessControllerVapDynamicMappingUtmLog(d, i["utm_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["utm-profile"], _ = expandObjectWirelessControllerVapDynamicMappingUtmProfile(d, i["utm_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utm_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["utm-status"], _ = expandObjectWirelessControllerVapDynamicMappingUtmStatus(d, i["utm_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectWirelessControllerVapDynamicMappingVdom(d, i["vdom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_auto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-auto"], _ = expandObjectWirelessControllerVapDynamicMappingVlanAuto(d, i["vlan_auto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_pooling"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-pooling"], _ = expandObjectWirelessControllerVapDynamicMappingVlanPooling(d, i["vlan_pooling"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlanid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlanid"], _ = expandObjectWirelessControllerVapDynamicMappingVlanid(d, i["vlanid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "voice_enterprise"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["voice-enterprise"], _ = expandObjectWirelessControllerVapDynamicMappingVoiceEnterprise(d, i["voice_enterprise"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "webfilter_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["webfilter-profile"], _ = expandObjectWirelessControllerVapDynamicMappingWebfilterProfile(d, i["webfilter_profile"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerVapDynamicMappingCentmgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpSvrId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDeviceAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDeviceIdentification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDeviceNetscan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfIp6Allowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfListenForticlientConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIsFactorySetting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectWirelessControllerVapDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectWirelessControllerVapDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerVapDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAccessControlList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAdditionalAkms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingAddressGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAddressGroupPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAntivirusProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingApplicationDetectionEngine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingApplicationDscpMarking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingApplicationReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAtfWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAuthCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAuthPortalAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingBeaconAdvertising(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingBroadcastSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingBroadcastSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingBssColorPartial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingBstmDisassociationImminent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingBstmRssiDisassocTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalAcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalFwAccounting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingClientCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpAddressEnforcement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpOption43Insertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpOption82Insertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDynamicVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingEapReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingEapReauthIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingEapolKeyRetries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingEncrypt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingExternalFastRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingExternalLogout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingExternalWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingExternalWebFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingFastBssTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingFastRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingFtMobilityDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingFtOverDs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingFtR0KeyLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingGasComebackDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingGasFragmentationLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingGtkRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingGtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingHighEfficiency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingHotspot20Profile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIgmpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntraVapPrivacy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIpv6Rules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingKeyindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingL3Roaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingL3RoamingMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLdpc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalBridging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalStandalone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalStandaloneDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalStandaloneDnsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalStandaloneNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalSwitching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacCalledStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacCallingStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacCase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacFilterPolicyOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMaxClientsAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMbo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMboCellDataConnPref(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMeDisableThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMeshBackhaul(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMpsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMpskConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMpskProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMuMimo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMulticastEnhance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMulticastRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingNac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingNacProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingNeighborReportDualBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingOkc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingOsen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingOweGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingOweTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingOweTransitionSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingPmf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPmfAssocComebackTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPortMacauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPortMacauthReauthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPortMacauthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPortalMessageOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPortalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPrimaryWagProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingProbeRespSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingProbeRespThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPtkRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingQosProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadio2GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadio5GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadioSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacMpskAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacMpskTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11A(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AcMcsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AcSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AcSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AxMcsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AxSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AxSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11Bg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11NSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11NSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingSaeGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingSaeH2EOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSaePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingSaePk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSaePrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSecondaryWagProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSecurityExemptList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSecurityObsoleteOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSecurityRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSelectedUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingStickyClientRemove(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold2G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold5G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold6G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingTargetWakeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingTkipCounterMeasure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingTunnelEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingTunnelFallbackInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingUsergroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingUtmLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingUtmProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingUtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingVlanAuto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingVlanPooling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingVoiceEnterprise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapEapReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapEapReauthIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapEapolKeyRetries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapEncrypt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapExternalFastRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapExternalLogout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapExternalWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapExternalWebFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapFastBssTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapFastRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapFtMobilityDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapFtOverDs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapFtR0KeyLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapGasComebackDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapGasFragmentationLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapGtkRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapGtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapHighEfficiency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapHotspot20Profile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIgmpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIntraVapPrivacy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapIpv6Rules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapKeyindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapL3Roaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapL3RoamingMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapLdpc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapLocalAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapLocalBridging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapLocalLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapLocalStandalone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapLocalStandaloneDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapLocalStandaloneDnsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapLocalStandaloneNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacCalledStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacCallingStationDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacCase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacFilterList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandObjectWirelessControllerVapMacFilterListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectWirelessControllerVapMacFilterListMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-filter-policy"], _ = expandObjectWirelessControllerVapMacFilterListMacFilterPolicy(d, i["mac_filter_policy"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerVapMacFilterListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacFilterListMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacFilterListMacFilterPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacFilterPolicyOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMaxClientsAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMbo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMboCellDataConnPref(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMeDisableThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMeshBackhaul(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMpsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMpskConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMpskKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectWirelessControllerVapMpskKeyComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["concurrent-clients"], _ = expandObjectWirelessControllerVapMpskKeyConcurrentClients(d, i["concurrent_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-name"], _ = expandObjectWirelessControllerVapMpskKeyKeyName(d, i["key_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-schedules"], _ = expandObjectWirelessControllerVapMpskKeyMpskSchedules(d, i["mpsk_schedules"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passphrase"], _ = expandObjectWirelessControllerVapMpskKeyPassphrase(d, i["passphrase"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerVapMpskKeyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMpskKeyConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMpskKeyKeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMpskKeyMpskSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMpskKeyPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapMpskProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMuMimo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMulticastEnhance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMulticastRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapNac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapNacProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapNeighborReportDualBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapOkc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapOsen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapOweGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapOweTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapOweTransitionSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapPmf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPmfAssocComebackTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPmfSaQueryRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortMacauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortMacauthReauthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortMacauthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortalMessageOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortalMessageOverrides(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_disclaimer_page"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-disclaimer-page"], _ = expandObjectWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(d, i["auth_disclaimer_page"], pre_append)
	}
	pre_append = pre + ".0." + "auth_login_failed_page"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-login-failed-page"], _ = expandObjectWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(d, i["auth_login_failed_page"], pre_append)
	}
	pre_append = pre + ".0." + "auth_login_page"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-login-page"], _ = expandObjectWirelessControllerVapPortalMessageOverridesAuthLoginPage(d, i["auth_login_page"], pre_append)
	}
	pre_append = pre + ".0." + "auth_reject_page"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auth-reject-page"], _ = expandObjectWirelessControllerVapPortalMessageOverridesAuthRejectPage(d, i["auth_reject_page"], pre_append)
	}

	return result, nil
}

func expandObjectWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortalMessageOverridesAuthLoginPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortalMessageOverridesAuthRejectPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPrimaryWagProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapProbeRespSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapProbeRespThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPtkRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapQosProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRadio2GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRadio5GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRadioSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRadiusMacAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRadiusMacAuthBlockInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRadiusMacAuthServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRadiusMacAuthUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapRadiusMacMpskAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRadiusMacMpskTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRates11A(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapRates11AcMcsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRates11AcSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapRates11AcSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapRates11AxMcsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapRates11AxSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapRates11AxSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapRates11Bg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapRates11NSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapRates11NSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapSaeGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapSaeH2EOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapSaePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapSaePk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapSaePrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectWirelessControllerVapSecondaryWagProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapSecurityExemptList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapSecurityObsoleteOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapSecurityRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapSelectedUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapSplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapStickyClientRemove(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapStickyClientThreshold2G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapStickyClientThreshold5G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapStickyClientThreshold6G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapTargetWakeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapTkipCounterMeasure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapTunnelEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapTunnelFallbackInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapUsergroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapUtmLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapUtmProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapUtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVlanAuto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVlanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectWirelessControllerVapVlanNameName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-id"], _ = expandObjectWirelessControllerVapVlanNameVlanId(d, i["vlan_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerVapVlanNameName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVlanNameVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVlanPool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_wtp_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_wtp-group"], _ = expandObjectWirelessControllerVapVlanPoolWtpGroup(d, i["_wtp_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWirelessControllerVapVlanPoolId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerVapVlanPoolWtpGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVlanPoolId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVlanPooling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVoiceEnterprise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerVap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_centmgmt"); ok || d.HasChange("_centmgmt") {
		t, err := expandObjectWirelessControllerVapCentmgmt(d, v, "_centmgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_centmgmt"] = t
		}
	}

	if v, ok := d.GetOk("_dhcp_svr_id"); ok || d.HasChange("_dhcp_svr_id") {
		t, err := expandObjectWirelessControllerVapDhcpSvrId(d, v, "_dhcp_svr_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dhcp_svr_id"] = t
		}
	}

	if v, ok := d.GetOk("_intf_allowaccess"); ok || d.HasChange("_intf_allowaccess") {
		t, err := expandObjectWirelessControllerVapIntfAllowaccess(d, v, "_intf_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("_intf_device_access_list"); ok || d.HasChange("_intf_device_access_list") {
		t, err := expandObjectWirelessControllerVapIntfDeviceAccessList(d, v, "_intf_device_access_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_device-access-list"] = t
		}
	}

	if v, ok := d.GetOk("_intf_device_identification"); ok || d.HasChange("_intf_device_identification") {
		t, err := expandObjectWirelessControllerVapIntfDeviceIdentification(d, v, "_intf_device_identification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_device-identification"] = t
		}
	}

	if v, ok := d.GetOk("_intf_device_netscan"); ok || d.HasChange("_intf_device_netscan") {
		t, err := expandObjectWirelessControllerVapIntfDeviceNetscan(d, v, "_intf_device_netscan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_device-netscan"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp_relay_ip"); ok || d.HasChange("_intf_dhcp_relay_ip") {
		t, err := expandObjectWirelessControllerVapIntfDhcpRelayIp(d, v, "_intf_dhcp_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp_relay_service"); ok || d.HasChange("_intf_dhcp_relay_service") {
		t, err := expandObjectWirelessControllerVapIntfDhcpRelayService(d, v, "_intf_dhcp_relay_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp-relay-service"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp_relay_type"); ok || d.HasChange("_intf_dhcp_relay_type") {
		t, err := expandObjectWirelessControllerVapIntfDhcpRelayType(d, v, "_intf_dhcp_relay_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp-relay-type"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp6_relay_ip"); ok || d.HasChange("_intf_dhcp6_relay_ip") {
		t, err := expandObjectWirelessControllerVapIntfDhcp6RelayIp(d, v, "_intf_dhcp6_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp6-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp6_relay_service"); ok || d.HasChange("_intf_dhcp6_relay_service") {
		t, err := expandObjectWirelessControllerVapIntfDhcp6RelayService(d, v, "_intf_dhcp6_relay_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp6-relay-service"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp6_relay_type"); ok || d.HasChange("_intf_dhcp6_relay_type") {
		t, err := expandObjectWirelessControllerVapIntfDhcp6RelayType(d, v, "_intf_dhcp6_relay_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp6-relay-type"] = t
		}
	}

	if v, ok := d.GetOk("_intf_ip"); ok || d.HasChange("_intf_ip") {
		t, err := expandObjectWirelessControllerVapIntfIp(d, v, "_intf_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_ip"] = t
		}
	}

	if v, ok := d.GetOk("_intf_ip6_address"); ok || d.HasChange("_intf_ip6_address") {
		t, err := expandObjectWirelessControllerVapIntfIp6Address(d, v, "_intf_ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("_intf_ip6_allowaccess"); ok || d.HasChange("_intf_ip6_allowaccess") {
		t, err := expandObjectWirelessControllerVapIntfIp6Allowaccess(d, v, "_intf_ip6_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_ip6-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("_intf_listen_forticlient_connection"); ok || d.HasChange("_intf_listen_forticlient_connection") {
		t, err := expandObjectWirelessControllerVapIntfListenForticlientConnection(d, v, "_intf_listen_forticlient_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_listen-forticlient-connection"] = t
		}
	}

	if v, ok := d.GetOk("_is_factory_setting"); ok || d.HasChange("_is_factory_setting") {
		t, err := expandObjectWirelessControllerVapIsFactorySetting(d, v, "_is_factory_setting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_is_factory_setting"] = t
		}
	}

	if v, ok := d.GetOk("access_control_list"); ok || d.HasChange("access_control_list") {
		t, err := expandObjectWirelessControllerVapAccessControlList(d, v, "access_control_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-control-list"] = t
		}
	}

	if v, ok := d.GetOk("additional_akms"); ok || d.HasChange("additional_akms") {
		t, err := expandObjectWirelessControllerVapAdditionalAkms(d, v, "additional_akms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-akms"] = t
		}
	}

	if v, ok := d.GetOk("acct_interim_interval"); ok || d.HasChange("acct_interim_interval") {
		t, err := expandObjectWirelessControllerVapAcctInterimInterval(d, v, "acct_interim_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-interim-interval"] = t
		}
	}

	if v, ok := d.GetOk("address_group"); ok || d.HasChange("address_group") {
		t, err := expandObjectWirelessControllerVapAddressGroup(d, v, "address_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-group"] = t
		}
	}

	if v, ok := d.GetOk("address_group_policy"); ok || d.HasChange("address_group_policy") {
		t, err := expandObjectWirelessControllerVapAddressGroupPolicy(d, v, "address_group_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-group-policy"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok || d.HasChange("alias") {
		t, err := expandObjectWirelessControllerVapAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("antivirus_profile"); ok || d.HasChange("antivirus_profile") {
		t, err := expandObjectWirelessControllerVapAntivirusProfile(d, v, "antivirus_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antivirus-profile"] = t
		}
	}

	if v, ok := d.GetOk("application_detection_engine"); ok || d.HasChange("application_detection_engine") {
		t, err := expandObjectWirelessControllerVapApplicationDetectionEngine(d, v, "application_detection_engine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-detection-engine"] = t
		}
	}

	if v, ok := d.GetOk("application_dscp_marking"); ok || d.HasChange("application_dscp_marking") {
		t, err := expandObjectWirelessControllerVapApplicationDscpMarking(d, v, "application_dscp_marking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-dscp-marking"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandObjectWirelessControllerVapApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("application_report_intv"); ok || d.HasChange("application_report_intv") {
		t, err := expandObjectWirelessControllerVapApplicationReportIntv(d, v, "application_report_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("atf_weight"); ok || d.HasChange("atf_weight") {
		t, err := expandObjectWirelessControllerVapAtfWeight(d, v, "atf_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["atf-weight"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok || d.HasChange("auth") {
		t, err := expandObjectWirelessControllerVapAuth(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok || d.HasChange("auth_cert") {
		t, err := expandObjectWirelessControllerVapAuthCert(d, v, "auth_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal_addr"); ok || d.HasChange("auth_portal_addr") {
		t, err := expandObjectWirelessControllerVapAuthPortalAddr(d, v, "auth_portal_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal-addr"] = t
		}
	}

	if v, ok := d.GetOk("beacon_advertising"); ok || d.HasChange("beacon_advertising") {
		t, err := expandObjectWirelessControllerVapBeaconAdvertising(d, v, "beacon_advertising")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-advertising"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_ssid"); ok || d.HasChange("broadcast_ssid") {
		t, err := expandObjectWirelessControllerVapBroadcastSsid(d, v, "broadcast_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-ssid"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_suppression"); ok || d.HasChange("broadcast_suppression") {
		t, err := expandObjectWirelessControllerVapBroadcastSuppression(d, v, "broadcast_suppression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-suppression"] = t
		}
	}

	if v, ok := d.GetOk("bss_color_partial"); ok || d.HasChange("bss_color_partial") {
		t, err := expandObjectWirelessControllerVapBssColorPartial(d, v, "bss_color_partial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color-partial"] = t
		}
	}

	if v, ok := d.GetOk("bstm_disassociation_imminent"); ok || d.HasChange("bstm_disassociation_imminent") {
		t, err := expandObjectWirelessControllerVapBstmDisassociationImminent(d, v, "bstm_disassociation_imminent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-disassociation-imminent"] = t
		}
	}

	if v, ok := d.GetOk("bstm_load_balancing_disassoc_timer"); ok || d.HasChange("bstm_load_balancing_disassoc_timer") {
		t, err := expandObjectWirelessControllerVapBstmLoadBalancingDisassocTimer(d, v, "bstm_load_balancing_disassoc_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-load-balancing-disassoc-timer"] = t
		}
	}

	if v, ok := d.GetOk("bstm_rssi_disassoc_timer"); ok || d.HasChange("bstm_rssi_disassoc_timer") {
		t, err := expandObjectWirelessControllerVapBstmRssiDisassocTimer(d, v, "bstm_rssi_disassoc_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-rssi-disassoc-timer"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ac_name"); ok || d.HasChange("captive_portal_ac_name") {
		t, err := expandObjectWirelessControllerVapCaptivePortalAcName(d, v, "captive_portal_ac_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ac-name"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_auth_timeout"); ok || d.HasChange("captive_portal_auth_timeout") {
		t, err := expandObjectWirelessControllerVapCaptivePortalAuthTimeout(d, v, "captive_portal_auth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-auth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_fw_accounting"); ok || d.HasChange("captive_portal_fw_accounting") {
		t, err := expandObjectWirelessControllerVapCaptivePortalFwAccounting(d, v, "captive_portal_fw_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-fw-accounting"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_address_enforcement"); ok || d.HasChange("dhcp_address_enforcement") {
		t, err := expandObjectWirelessControllerVapDhcpAddressEnforcement(d, v, "dhcp_address_enforcement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-address-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_macauth_radius_secret"); ok || d.HasChange("captive_portal_macauth_radius_secret") {
		t, err := expandObjectWirelessControllerVapCaptivePortalMacauthRadiusSecret(d, v, "captive_portal_macauth_radius_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-macauth-radius-secret"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_macauth_radius_server"); ok || d.HasChange("captive_portal_macauth_radius_server") {
		t, err := expandObjectWirelessControllerVapCaptivePortalMacauthRadiusServer(d, v, "captive_portal_macauth_radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-macauth-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_radius_secret"); ok || d.HasChange("captive_portal_radius_secret") {
		t, err := expandObjectWirelessControllerVapCaptivePortalRadiusSecret(d, v, "captive_portal_radius_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-radius-secret"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_radius_server"); ok || d.HasChange("captive_portal_radius_server") {
		t, err := expandObjectWirelessControllerVapCaptivePortalRadiusServer(d, v, "captive_portal_radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_session_timeout_interval"); ok || d.HasChange("captive_portal_session_timeout_interval") {
		t, err := expandObjectWirelessControllerVapCaptivePortalSessionTimeoutInterval(d, v, "captive_portal_session_timeout_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-session-timeout-interval"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_lease_time"); ok || d.HasChange("dhcp_lease_time") {
		t, err := expandObjectWirelessControllerVapDhcpLeaseTime(d, v, "dhcp_lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-lease-time"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option43_insertion"); ok || d.HasChange("dhcp_option43_insertion") {
		t, err := expandObjectWirelessControllerVapDhcpOption43Insertion(d, v, "dhcp_option43_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option43-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_circuit_id_insertion"); ok || d.HasChange("dhcp_option82_circuit_id_insertion") {
		t, err := expandObjectWirelessControllerVapDhcpOption82CircuitIdInsertion(d, v, "dhcp_option82_circuit_id_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-circuit-id-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_insertion"); ok || d.HasChange("dhcp_option82_insertion") {
		t, err := expandObjectWirelessControllerVapDhcpOption82Insertion(d, v, "dhcp_option82_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_remote_id_insertion"); ok || d.HasChange("dhcp_option82_remote_id_insertion") {
		t, err := expandObjectWirelessControllerVapDhcpOption82RemoteIdInsertion(d, v, "dhcp_option82_remote_id_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-remote-id-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_vlan"); ok || d.HasChange("dynamic_vlan") {
		t, err := expandObjectWirelessControllerVapDynamicVlan(d, v, "dynamic_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-vlan"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectWirelessControllerVapDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("eap_reauth"); ok || d.HasChange("eap_reauth") {
		t, err := expandObjectWirelessControllerVapEapReauth(d, v, "eap_reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-reauth"] = t
		}
	}

	if v, ok := d.GetOk("eap_reauth_intv"); ok || d.HasChange("eap_reauth_intv") {
		t, err := expandObjectWirelessControllerVapEapReauthIntv(d, v, "eap_reauth_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-reauth-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_key_retries"); ok || d.HasChange("eapol_key_retries") {
		t, err := expandObjectWirelessControllerVapEapolKeyRetries(d, v, "eapol_key_retries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-key-retries"] = t
		}
	}

	if v, ok := d.GetOk("encrypt"); ok || d.HasChange("encrypt") {
		t, err := expandObjectWirelessControllerVapEncrypt(d, v, "encrypt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encrypt"] = t
		}
	}

	if v, ok := d.GetOk("external_fast_roaming"); ok || d.HasChange("external_fast_roaming") {
		t, err := expandObjectWirelessControllerVapExternalFastRoaming(d, v, "external_fast_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-fast-roaming"] = t
		}
	}

	if v, ok := d.GetOk("external_logout"); ok || d.HasChange("external_logout") {
		t, err := expandObjectWirelessControllerVapExternalLogout(d, v, "external_logout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-logout"] = t
		}
	}

	if v, ok := d.GetOk("external_web"); ok || d.HasChange("external_web") {
		t, err := expandObjectWirelessControllerVapExternalWeb(d, v, "external_web")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-web"] = t
		}
	}

	if v, ok := d.GetOk("external_web_format"); ok || d.HasChange("external_web_format") {
		t, err := expandObjectWirelessControllerVapExternalWebFormat(d, v, "external_web_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-web-format"] = t
		}
	}

	if v, ok := d.GetOk("fast_bss_transition"); ok || d.HasChange("fast_bss_transition") {
		t, err := expandObjectWirelessControllerVapFastBssTransition(d, v, "fast_bss_transition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-bss-transition"] = t
		}
	}

	if v, ok := d.GetOk("fast_roaming"); ok || d.HasChange("fast_roaming") {
		t, err := expandObjectWirelessControllerVapFastRoaming(d, v, "fast_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-roaming"] = t
		}
	}

	if v, ok := d.GetOk("ft_mobility_domain"); ok || d.HasChange("ft_mobility_domain") {
		t, err := expandObjectWirelessControllerVapFtMobilityDomain(d, v, "ft_mobility_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-mobility-domain"] = t
		}
	}

	if v, ok := d.GetOk("ft_over_ds"); ok || d.HasChange("ft_over_ds") {
		t, err := expandObjectWirelessControllerVapFtOverDs(d, v, "ft_over_ds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-over-ds"] = t
		}
	}

	if v, ok := d.GetOk("ft_r0_key_lifetime"); ok || d.HasChange("ft_r0_key_lifetime") {
		t, err := expandObjectWirelessControllerVapFtR0KeyLifetime(d, v, "ft_r0_key_lifetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-r0-key-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("gas_comeback_delay"); ok || d.HasChange("gas_comeback_delay") {
		t, err := expandObjectWirelessControllerVapGasComebackDelay(d, v, "gas_comeback_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-comeback-delay"] = t
		}
	}

	if v, ok := d.GetOk("gas_fragmentation_limit"); ok || d.HasChange("gas_fragmentation_limit") {
		t, err := expandObjectWirelessControllerVapGasFragmentationLimit(d, v, "gas_fragmentation_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-fragmentation-limit"] = t
		}
	}

	if v, ok := d.GetOk("gtk_rekey"); ok || d.HasChange("gtk_rekey") {
		t, err := expandObjectWirelessControllerVapGtkRekey(d, v, "gtk_rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtk-rekey"] = t
		}
	}

	if v, ok := d.GetOk("gtk_rekey_intv"); ok || d.HasChange("gtk_rekey_intv") {
		t, err := expandObjectWirelessControllerVapGtkRekeyIntv(d, v, "gtk_rekey_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtk-rekey-intv"] = t
		}
	}

	if v, ok := d.GetOk("high_efficiency"); ok || d.HasChange("high_efficiency") {
		t, err := expandObjectWirelessControllerVapHighEfficiency(d, v, "high_efficiency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high-efficiency"] = t
		}
	}

	if v, ok := d.GetOk("hotspot20_profile"); ok || d.HasChange("hotspot20_profile") {
		t, err := expandObjectWirelessControllerVapHotspot20Profile(d, v, "hotspot20_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hotspot20-profile"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping"); ok || d.HasChange("igmp_snooping") {
		t, err := expandObjectWirelessControllerVapIgmpSnooping(d, v, "igmp_snooping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("intra_vap_privacy"); ok || d.HasChange("intra_vap_privacy") {
		t, err := expandObjectWirelessControllerVapIntraVapPrivacy(d, v, "intra_vap_privacy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intra-vap-privacy"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectWirelessControllerVapIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandObjectWirelessControllerVapIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_rules"); ok || d.HasChange("ipv6_rules") {
		t, err := expandObjectWirelessControllerVapIpv6Rules(d, v, "ipv6_rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-rules"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandObjectWirelessControllerVapKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("keyindex"); ok || d.HasChange("keyindex") {
		t, err := expandObjectWirelessControllerVapKeyindex(d, v, "keyindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyindex"] = t
		}
	}

	if v, ok := d.GetOk("l3_roaming"); ok || d.HasChange("l3_roaming") {
		t, err := expandObjectWirelessControllerVapL3Roaming(d, v, "l3_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-roaming"] = t
		}
	}

	if v, ok := d.GetOk("l3_roaming_mode"); ok || d.HasChange("l3_roaming_mode") {
		t, err := expandObjectWirelessControllerVapL3RoamingMode(d, v, "l3_roaming_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-roaming-mode"] = t
		}
	}

	if v, ok := d.GetOk("ldpc"); ok || d.HasChange("ldpc") {
		t, err := expandObjectWirelessControllerVapLdpc(d, v, "ldpc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldpc"] = t
		}
	}

	if v, ok := d.GetOk("local_authentication"); ok || d.HasChange("local_authentication") {
		t, err := expandObjectWirelessControllerVapLocalAuthentication(d, v, "local_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-authentication"] = t
		}
	}

	if v, ok := d.GetOk("local_bridging"); ok || d.HasChange("local_bridging") {
		t, err := expandObjectWirelessControllerVapLocalBridging(d, v, "local_bridging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-bridging"] = t
		}
	}

	if v, ok := d.GetOk("local_lan"); ok || d.HasChange("local_lan") {
		t, err := expandObjectWirelessControllerVapLocalLan(d, v, "local_lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-lan"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone"); ok || d.HasChange("local_standalone") {
		t, err := expandObjectWirelessControllerVapLocalStandalone(d, v, "local_standalone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_dns"); ok || d.HasChange("local_standalone_dns") {
		t, err := expandObjectWirelessControllerVapLocalStandaloneDns(d, v, "local_standalone_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-dns"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_dns_ip"); ok || d.HasChange("local_standalone_dns_ip") {
		t, err := expandObjectWirelessControllerVapLocalStandaloneDnsIp(d, v, "local_standalone_dns_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-dns-ip"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_nat"); ok || d.HasChange("local_standalone_nat") {
		t, err := expandObjectWirelessControllerVapLocalStandaloneNat(d, v, "local_standalone_nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-nat"] = t
		}
	}

	if v, ok := d.GetOk("mac_auth_bypass"); ok || d.HasChange("mac_auth_bypass") {
		t, err := expandObjectWirelessControllerVapMacAuthBypass(d, v, "mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("mac_called_station_delimiter"); ok || d.HasChange("mac_called_station_delimiter") {
		t, err := expandObjectWirelessControllerVapMacCalledStationDelimiter(d, v, "mac_called_station_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-called-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_calling_station_delimiter"); ok || d.HasChange("mac_calling_station_delimiter") {
		t, err := expandObjectWirelessControllerVapMacCallingStationDelimiter(d, v, "mac_calling_station_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-calling-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_case"); ok || d.HasChange("mac_case") {
		t, err := expandObjectWirelessControllerVapMacCase(d, v, "mac_case")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-case"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter"); ok || d.HasChange("mac_filter") {
		t, err := expandObjectWirelessControllerVapMacFilter(d, v, "mac_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter_list"); ok || d.HasChange("mac_filter_list") {
		t, err := expandObjectWirelessControllerVapMacFilterList(d, v, "mac_filter_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter-list"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter_policy_other"); ok || d.HasChange("mac_filter_policy_other") {
		t, err := expandObjectWirelessControllerVapMacFilterPolicyOther(d, v, "mac_filter_policy_other")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter-policy-other"] = t
		}
	}

	if v, ok := d.GetOk("mac_password_delimiter"); ok || d.HasChange("mac_password_delimiter") {
		t, err := expandObjectWirelessControllerVapMacPasswordDelimiter(d, v, "mac_password_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-password-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_username_delimiter"); ok || d.HasChange("mac_username_delimiter") {
		t, err := expandObjectWirelessControllerVapMacUsernameDelimiter(d, v, "mac_username_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-username-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("max_clients"); ok || d.HasChange("max_clients") {
		t, err := expandObjectWirelessControllerVapMaxClients(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("max_clients_ap"); ok || d.HasChange("max_clients_ap") {
		t, err := expandObjectWirelessControllerVapMaxClientsAp(d, v, "max_clients_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients-ap"] = t
		}
	}

	if v, ok := d.GetOk("mbo"); ok || d.HasChange("mbo") {
		t, err := expandObjectWirelessControllerVapMbo(d, v, "mbo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbo"] = t
		}
	}

	if v, ok := d.GetOk("mbo_cell_data_conn_pref"); ok || d.HasChange("mbo_cell_data_conn_pref") {
		t, err := expandObjectWirelessControllerVapMboCellDataConnPref(d, v, "mbo_cell_data_conn_pref")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbo-cell-data-conn-pref"] = t
		}
	}

	if v, ok := d.GetOk("me_disable_thresh"); ok || d.HasChange("me_disable_thresh") {
		t, err := expandObjectWirelessControllerVapMeDisableThresh(d, v, "me_disable_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["me-disable-thresh"] = t
		}
	}

	if v, ok := d.GetOk("mesh_backhaul"); ok || d.HasChange("mesh_backhaul") {
		t, err := expandObjectWirelessControllerVapMeshBackhaul(d, v, "mesh_backhaul")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-backhaul"] = t
		}
	}

	if v, ok := d.GetOk("mpsk"); ok || d.HasChange("mpsk") {
		t, err := expandObjectWirelessControllerVapMpsk(d, v, "mpsk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_concurrent_clients"); ok || d.HasChange("mpsk_concurrent_clients") {
		t, err := expandObjectWirelessControllerVapMpskConcurrentClients(d, v, "mpsk_concurrent_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-concurrent-clients"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_key"); ok || d.HasChange("mpsk_key") {
		t, err := expandObjectWirelessControllerVapMpskKey(d, v, "mpsk_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-key"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_profile"); ok || d.HasChange("mpsk_profile") {
		t, err := expandObjectWirelessControllerVapMpskProfile(d, v, "mpsk_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-profile"] = t
		}
	}

	if v, ok := d.GetOk("mu_mimo"); ok || d.HasChange("mu_mimo") {
		t, err := expandObjectWirelessControllerVapMuMimo(d, v, "mu_mimo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mu-mimo"] = t
		}
	}

	if v, ok := d.GetOk("multicast_enhance"); ok || d.HasChange("multicast_enhance") {
		t, err := expandObjectWirelessControllerVapMulticastEnhance(d, v, "multicast_enhance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-enhance"] = t
		}
	}

	if v, ok := d.GetOk("multicast_rate"); ok || d.HasChange("multicast_rate") {
		t, err := expandObjectWirelessControllerVapMulticastRate(d, v, "multicast_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-rate"] = t
		}
	}

	if v, ok := d.GetOk("nac"); ok || d.HasChange("nac") {
		t, err := expandObjectWirelessControllerVapNac(d, v, "nac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac"] = t
		}
	}

	if v, ok := d.GetOk("nac_profile"); ok || d.HasChange("nac_profile") {
		t, err := expandObjectWirelessControllerVapNacProfile(d, v, "nac_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-profile"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerVapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_report_dual_band"); ok || d.HasChange("neighbor_report_dual_band") {
		t, err := expandObjectWirelessControllerVapNeighborReportDualBand(d, v, "neighbor_report_dual_band")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-report-dual-band"] = t
		}
	}

	if v, ok := d.GetOk("okc"); ok || d.HasChange("okc") {
		t, err := expandObjectWirelessControllerVapOkc(d, v, "okc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["okc"] = t
		}
	}

	if v, ok := d.GetOk("osen"); ok || d.HasChange("osen") {
		t, err := expandObjectWirelessControllerVapOsen(d, v, "osen")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osen"] = t
		}
	}

	if v, ok := d.GetOk("owe_groups"); ok || d.HasChange("owe_groups") {
		t, err := expandObjectWirelessControllerVapOweGroups(d, v, "owe_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-groups"] = t
		}
	}

	if v, ok := d.GetOk("owe_transition"); ok || d.HasChange("owe_transition") {
		t, err := expandObjectWirelessControllerVapOweTransition(d, v, "owe_transition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-transition"] = t
		}
	}

	if v, ok := d.GetOk("owe_transition_ssid"); ok || d.HasChange("owe_transition_ssid") {
		t, err := expandObjectWirelessControllerVapOweTransitionSsid(d, v, "owe_transition_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-transition-ssid"] = t
		}
	}

	if v, ok := d.GetOk("passphrase"); ok || d.HasChange("passphrase") {
		t, err := expandObjectWirelessControllerVapPassphrase(d, v, "passphrase")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passphrase"] = t
		}
	}

	if v, ok := d.GetOk("pmf"); ok || d.HasChange("pmf") {
		t, err := expandObjectWirelessControllerVapPmf(d, v, "pmf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf"] = t
		}
	}

	if v, ok := d.GetOk("pmf_assoc_comeback_timeout"); ok || d.HasChange("pmf_assoc_comeback_timeout") {
		t, err := expandObjectWirelessControllerVapPmfAssocComebackTimeout(d, v, "pmf_assoc_comeback_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf-assoc-comeback-timeout"] = t
		}
	}

	if v, ok := d.GetOk("pmf_sa_query_retry_timeout"); ok || d.HasChange("pmf_sa_query_retry_timeout") {
		t, err := expandObjectWirelessControllerVapPmfSaQueryRetryTimeout(d, v, "pmf_sa_query_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf-sa-query-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth"); ok || d.HasChange("port_macauth") {
		t, err := expandObjectWirelessControllerVapPortMacauth(d, v, "port_macauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth_reauth_timeout"); ok || d.HasChange("port_macauth_reauth_timeout") {
		t, err := expandObjectWirelessControllerVapPortMacauthReauthTimeout(d, v, "port_macauth_reauth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth-reauth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth_timeout"); ok || d.HasChange("port_macauth_timeout") {
		t, err := expandObjectWirelessControllerVapPortMacauthTimeout(d, v, "port_macauth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("portal_message_override_group"); ok || d.HasChange("portal_message_override_group") {
		t, err := expandObjectWirelessControllerVapPortalMessageOverrideGroup(d, v, "portal_message_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-override-group"] = t
		}
	}

	if v, ok := d.GetOk("portal_message_overrides"); ok || d.HasChange("portal_message_overrides") {
		t, err := expandObjectWirelessControllerVapPortalMessageOverrides(d, v, "portal_message_overrides")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-overrides"] = t
		}
	}

	if v, ok := d.GetOk("portal_type"); ok || d.HasChange("portal_type") {
		t, err := expandObjectWirelessControllerVapPortalType(d, v, "portal_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-type"] = t
		}
	}

	if v, ok := d.GetOk("primary_wag_profile"); ok || d.HasChange("primary_wag_profile") {
		t, err := expandObjectWirelessControllerVapPrimaryWagProfile(d, v, "primary_wag_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-wag-profile"] = t
		}
	}

	if v, ok := d.GetOk("probe_resp_suppression"); ok || d.HasChange("probe_resp_suppression") {
		t, err := expandObjectWirelessControllerVapProbeRespSuppression(d, v, "probe_resp_suppression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-resp-suppression"] = t
		}
	}

	if v, ok := d.GetOk("probe_resp_threshold"); ok || d.HasChange("probe_resp_threshold") {
		t, err := expandObjectWirelessControllerVapProbeRespThreshold(d, v, "probe_resp_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-resp-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ptk_rekey"); ok || d.HasChange("ptk_rekey") {
		t, err := expandObjectWirelessControllerVapPtkRekey(d, v, "ptk_rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptk-rekey"] = t
		}
	}

	if v, ok := d.GetOk("ptk_rekey_intv"); ok || d.HasChange("ptk_rekey_intv") {
		t, err := expandObjectWirelessControllerVapPtkRekeyIntv(d, v, "ptk_rekey_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptk-rekey-intv"] = t
		}
	}

	if v, ok := d.GetOk("qos_profile"); ok || d.HasChange("qos_profile") {
		t, err := expandObjectWirelessControllerVapQosProfile(d, v, "qos_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-profile"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandObjectWirelessControllerVapQuarantine(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("radio_2g_threshold"); ok || d.HasChange("radio_2g_threshold") {
		t, err := expandObjectWirelessControllerVapRadio2GThreshold(d, v, "radio_2g_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("radio_5g_threshold"); ok || d.HasChange("radio_5g_threshold") {
		t, err := expandObjectWirelessControllerVapRadio5GThreshold(d, v, "radio_5g_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-5g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("radio_sensitivity"); ok || d.HasChange("radio_sensitivity") {
		t, err := expandObjectWirelessControllerVapRadioSensitivity(d, v, "radio_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth"); ok || d.HasChange("radius_mac_auth") {
		t, err := expandObjectWirelessControllerVapRadiusMacAuth(d, v, "radius_mac_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_block_interval"); ok || d.HasChange("radius_mac_auth_block_interval") {
		t, err := expandObjectWirelessControllerVapRadiusMacAuthBlockInterval(d, v, "radius_mac_auth_block_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-block-interval"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_server"); ok || d.HasChange("radius_mac_auth_server") {
		t, err := expandObjectWirelessControllerVapRadiusMacAuthServer(d, v, "radius_mac_auth_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-server"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_usergroups"); ok || d.HasChange("radius_mac_auth_usergroups") {
		t, err := expandObjectWirelessControllerVapRadiusMacAuthUsergroups(d, v, "radius_mac_auth_usergroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_mpsk_auth"); ok || d.HasChange("radius_mac_mpsk_auth") {
		t, err := expandObjectWirelessControllerVapRadiusMacMpskAuth(d, v, "radius_mac_mpsk_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-mpsk-auth"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_mpsk_timeout"); ok || d.HasChange("radius_mac_mpsk_timeout") {
		t, err := expandObjectWirelessControllerVapRadiusMacMpskTimeout(d, v, "radius_mac_mpsk_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-mpsk-timeout"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok || d.HasChange("radius_server") {
		t, err := expandObjectWirelessControllerVapRadiusServer(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("rates_11a"); ok || d.HasChange("rates_11a") {
		t, err := expandObjectWirelessControllerVapRates11A(d, v, "rates_11a")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11a"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_mcs_map"); ok || d.HasChange("rates_11ac_mcs_map") {
		t, err := expandObjectWirelessControllerVapRates11AcMcsMap(d, v, "rates_11ac_mcs_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-mcs-map"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_ss12"); ok || d.HasChange("rates_11ac_ss12") {
		t, err := expandObjectWirelessControllerVapRates11AcSs12(d, v, "rates_11ac_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_ss34"); ok || d.HasChange("rates_11ac_ss34") {
		t, err := expandObjectWirelessControllerVapRates11AcSs34(d, v, "rates_11ac_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-ss34"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ax_mcs_map"); ok || d.HasChange("rates_11ax_mcs_map") {
		t, err := expandObjectWirelessControllerVapRates11AxMcsMap(d, v, "rates_11ax_mcs_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-mcs-map"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ax_ss12"); ok || d.HasChange("rates_11ax_ss12") {
		t, err := expandObjectWirelessControllerVapRates11AxSs12(d, v, "rates_11ax_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ax_ss34"); ok || d.HasChange("rates_11ax_ss34") {
		t, err := expandObjectWirelessControllerVapRates11AxSs34(d, v, "rates_11ax_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-ss34"] = t
		}
	}

	if v, ok := d.GetOk("rates_11bg"); ok || d.HasChange("rates_11bg") {
		t, err := expandObjectWirelessControllerVapRates11Bg(d, v, "rates_11bg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11bg"] = t
		}
	}

	if v, ok := d.GetOk("rates_11n_ss12"); ok || d.HasChange("rates_11n_ss12") {
		t, err := expandObjectWirelessControllerVapRates11NSs12(d, v, "rates_11n_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11n-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11n_ss34"); ok || d.HasChange("rates_11n_ss34") {
		t, err := expandObjectWirelessControllerVapRates11NSs34(d, v, "rates_11n_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11n-ss34"] = t
		}
	}

	if v, ok := d.GetOk("sae_groups"); ok || d.HasChange("sae_groups") {
		t, err := expandObjectWirelessControllerVapSaeGroups(d, v, "sae_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-groups"] = t
		}
	}

	if v, ok := d.GetOk("sae_h2e_only"); ok || d.HasChange("sae_h2e_only") {
		t, err := expandObjectWirelessControllerVapSaeH2EOnly(d, v, "sae_h2e_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-h2e-only"] = t
		}
	}

	if v, ok := d.GetOk("sae_password"); ok || d.HasChange("sae_password") {
		t, err := expandObjectWirelessControllerVapSaePassword(d, v, "sae_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-password"] = t
		}
	}

	if v, ok := d.GetOk("sae_pk"); ok || d.HasChange("sae_pk") {
		t, err := expandObjectWirelessControllerVapSaePk(d, v, "sae_pk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-pk"] = t
		}
	}

	if v, ok := d.GetOk("sae_private_key"); ok || d.HasChange("sae_private_key") {
		t, err := expandObjectWirelessControllerVapSaePrivateKey(d, v, "sae_private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-private-key"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok || d.HasChange("scan_botnet_connections") {
		t, err := expandObjectWirelessControllerVapScanBotnetConnections(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandObjectWirelessControllerVapSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("secondary_wag_profile"); ok || d.HasChange("secondary_wag_profile") {
		t, err := expandObjectWirelessControllerVapSecondaryWagProfile(d, v, "secondary_wag_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-wag-profile"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok || d.HasChange("security") {
		t, err := expandObjectWirelessControllerVapSecurity(d, v, "security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("security_exempt_list"); ok || d.HasChange("security_exempt_list") {
		t, err := expandObjectWirelessControllerVapSecurityExemptList(d, v, "security_exempt_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-exempt-list"] = t
		}
	}

	if v, ok := d.GetOk("security_obsolete_option"); ok || d.HasChange("security_obsolete_option") {
		t, err := expandObjectWirelessControllerVapSecurityObsoleteOption(d, v, "security_obsolete_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-obsolete-option"] = t
		}
	}

	if v, ok := d.GetOk("security_redirect_url"); ok || d.HasChange("security_redirect_url") {
		t, err := expandObjectWirelessControllerVapSecurityRedirectUrl(d, v, "security_redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("selected_usergroups"); ok || d.HasChange("selected_usergroups") {
		t, err := expandObjectWirelessControllerVapSelectedUsergroups(d, v, "selected_usergroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selected-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling"); ok || d.HasChange("split_tunneling") {
		t, err := expandObjectWirelessControllerVapSplitTunneling(d, v, "split_tunneling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("ssid"); ok || d.HasChange("ssid") {
		t, err := expandObjectWirelessControllerVapSsid(d, v, "ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_remove"); ok || d.HasChange("sticky_client_remove") {
		t, err := expandObjectWirelessControllerVapStickyClientRemove(d, v, "sticky_client_remove")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-remove"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_2g"); ok || d.HasChange("sticky_client_threshold_2g") {
		t, err := expandObjectWirelessControllerVapStickyClientThreshold2G(d, v, "sticky_client_threshold_2g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-2g"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_5g"); ok || d.HasChange("sticky_client_threshold_5g") {
		t, err := expandObjectWirelessControllerVapStickyClientThreshold5G(d, v, "sticky_client_threshold_5g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-5g"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_6g"); ok || d.HasChange("sticky_client_threshold_6g") {
		t, err := expandObjectWirelessControllerVapStickyClientThreshold6G(d, v, "sticky_client_threshold_6g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-6g"] = t
		}
	}

	if v, ok := d.GetOk("target_wake_time"); ok || d.HasChange("target_wake_time") {
		t, err := expandObjectWirelessControllerVapTargetWakeTime(d, v, "target_wake_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-wake-time"] = t
		}
	}

	if v, ok := d.GetOk("tkip_counter_measure"); ok || d.HasChange("tkip_counter_measure") {
		t, err := expandObjectWirelessControllerVapTkipCounterMeasure(d, v, "tkip_counter_measure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tkip-counter-measure"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_echo_interval"); ok || d.HasChange("tunnel_echo_interval") {
		t, err := expandObjectWirelessControllerVapTunnelEchoInterval(d, v, "tunnel_echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-echo-interval"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_fallback_interval"); ok || d.HasChange("tunnel_fallback_interval") {
		t, err := expandObjectWirelessControllerVapTunnelFallbackInterval(d, v, "tunnel_fallback_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-fallback-interval"] = t
		}
	}

	if v, ok := d.GetOk("usergroup"); ok || d.HasChange("usergroup") {
		t, err := expandObjectWirelessControllerVapUsergroup(d, v, "usergroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usergroup"] = t
		}
	}

	if v, ok := d.GetOk("utm_log"); ok || d.HasChange("utm_log") {
		t, err := expandObjectWirelessControllerVapUtmLog(d, v, "utm_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-log"] = t
		}
	}

	if v, ok := d.GetOk("utm_profile"); ok || d.HasChange("utm_profile") {
		t, err := expandObjectWirelessControllerVapUtmProfile(d, v, "utm_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-profile"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandObjectWirelessControllerVapVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok || d.HasChange("utm_status") {
		t, err := expandObjectWirelessControllerVapUtmStatus(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("vlan_auto"); ok || d.HasChange("vlan_auto") {
		t, err := expandObjectWirelessControllerVapVlanAuto(d, v, "vlan_auto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-auto"] = t
		}
	}

	if v, ok := d.GetOk("vlan_name"); ok || d.HasChange("vlan_name") {
		t, err := expandObjectWirelessControllerVapVlanName(d, v, "vlan_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-name"] = t
		}
	}

	if v, ok := d.GetOk("vlan_pool"); ok || d.HasChange("vlan_pool") {
		t, err := expandObjectWirelessControllerVapVlanPool(d, v, "vlan_pool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-pool"] = t
		}
	}

	if v, ok := d.GetOk("vlan_pooling"); ok || d.HasChange("vlan_pooling") {
		t, err := expandObjectWirelessControllerVapVlanPooling(d, v, "vlan_pooling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-pooling"] = t
		}
	}

	if v, ok := d.GetOk("vlanid"); ok || d.HasChange("vlanid") {
		t, err := expandObjectWirelessControllerVapVlanid(d, v, "vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	if v, ok := d.GetOk("voice_enterprise"); ok || d.HasChange("voice_enterprise") {
		t, err := expandObjectWirelessControllerVapVoiceEnterprise(d, v, "voice_enterprise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice-enterprise"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandObjectWirelessControllerVapWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	return &obj, nil
}
