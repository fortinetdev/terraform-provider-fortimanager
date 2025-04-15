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

func resourceObjectWirelessControllerVapDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerVapDynamicMappingCreate,
		Read:   resourceObjectWirelessControllerVapDynamicMappingRead,
		Update: resourceObjectWirelessControllerVapDynamicMappingUpdate,
		Delete: resourceObjectWirelessControllerVapDynamicMappingDelete,

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
			"vap": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"n80211k": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n80211v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"captive_portal_macauth_radius_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"captive_portal_macauth_radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captive_portal_radius_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
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
				Computed: true,
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
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
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
			"roaming_acct_interim_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"sae_hnp_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sae_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectWirelessControllerVapDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	vap := d.Get("vap").(string)
	paradict["vap"] = vap

	obj, err := getObjectObjectWirelessControllerVapDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerVapDynamicMapping resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerVapDynamicMapping(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerVapDynamicMapping resource: %v", err)
	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectWirelessControllerVapDynamicMappingRead(d, m)
}

func resourceObjectWirelessControllerVapDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	vap := d.Get("vap").(string)
	paradict["vap"] = vap

	obj, err := getObjectObjectWirelessControllerVapDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVapDynamicMapping resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerVapDynamicMapping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVapDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectWirelessControllerVapDynamicMappingRead(d, m)
}

func resourceObjectWirelessControllerVapDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	vap := d.Get("vap").(string)
	paradict["vap"] = vap

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerVapDynamicMapping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerVapDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerVapDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	vap := d.Get("vap").(string)
	if vap == "" {
		vap = importOptionChecking(m.(*FortiClient).Cfg, "vap")
		if vap == "" {
			return fmt.Errorf("Parameter vap is missing")
		}
		if err = d.Set("vap", vap); err != nil {
			return fmt.Errorf("Error set params vap: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["vap"] = vap

	o, err := c.ReadObjectWirelessControllerVapDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVapDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerVapDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVapDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerVapDynamicMapping80211K2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMapping80211V2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCentmgmt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpSvrId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfAllowaccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDeviceAccessList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDeviceIdentification2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDeviceNetscan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenObjectWirelessControllerVapDynamicMappingIntfIp6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntfIp6Allowaccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingIntfListenForticlientConnection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIsFactorySetting2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerVapDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVapDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectWirelessControllerVapDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerVapDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerVapDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAccessControlList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingAcctInterimInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAdditionalAkms2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingAddressGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingAddressGroupPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAlias2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAntivirusProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingApplicationDetectionEngine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingApplicationDscpMarking2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingApplicationList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingApplicationReportIntv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAtfWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingAuthCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingAuthPortalAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingBeaconAdvertising2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingBroadcastSsid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingBroadcastSuppression2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingBssColorPartial2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingBstmDisassociationImminent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingBstmRssiDisassocTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalAcName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalFwAccounting2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingClientCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpAddressEnforcement2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpOption43Insertion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpOption82Insertion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingDynamicVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingEapReauth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingEapReauthIntv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingEapolKeyRetries2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingEncrypt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingExternalFastRoaming2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingExternalLogout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingExternalWeb2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingExternalWebFormat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingFastBssTransition2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingFastRoaming2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingFtMobilityDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingFtOverDs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingFtR0KeyLifetime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingGasComebackDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingGasFragmentationLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingGtkRekey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingGtkRekeyIntv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingHighEfficiency2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingHotspot20Profile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIgmpSnooping2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIntraVapPrivacy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingIpsSensor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingIpv6Rules2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingKeyindex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingL3Roaming2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingL3RoamingMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLdpc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalBridging2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalLan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalStandalone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneDns2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneDnsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneNat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingLocalSwitching2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacAuthBypass2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacCalledStationDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacCallingStationDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacCase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacFilterPolicyOther2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacPasswordDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMacUsernameDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMaxClients2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMaxClientsAp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMbo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMboCellDataConnPref2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMeDisableThresh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMeshBackhaul2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMpsk2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMpskConcurrentClients2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMpskProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingMuMimo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMulticastEnhance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingMulticastRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingNac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingNacProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingNeighborReportDualBand2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingOkc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingOsen2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingOweGroups2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingOweTransition2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingOweTransitionSsid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPmf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPmfAssocComebackTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPortMacauth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPortMacauthReauthTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPortMacauthTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPortalMessageOverrideGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPortalType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPrimaryWagProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingProbeRespSuppression2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingProbeRespThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPtkRekey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingPtkRekeyIntv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingQosProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadio2GThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadio5GThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadioSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacMpskAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusMacMpskTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRadiusServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRates11A2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AcMcsMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AcSs122edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AcSs342edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AxMcsMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AxSs122edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11AxSs342edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11Bg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11NSs122edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRates11NSs342edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingRoamingAcctInterimUpdate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSaeGroups2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerVapDynamicMappingSaeH2EOnly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSaeHnpOnly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSaePk2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSaePrivateKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingScanBotnetConnections2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSchedule2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingSecondaryWagProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingSecurity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSecurityExemptList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSecurityObsoleteOption2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSecurityRedirectUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSelectedUsergroups2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingSplitTunneling2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingSsid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingStickyClientRemove2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold2G2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold5G2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold6G2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingTargetWakeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingTkipCounterMeasure2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingTunnelEchoInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingTunnelFallbackInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingUsergroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingUtmLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingUtmProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingUtmStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerVapDynamicMappingVlanAuto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingVlanPooling2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingVlanid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingVoiceEnterprise2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapDynamicMappingWebfilterProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectObjectWirelessControllerVapDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("n80211k", flattenObjectWirelessControllerVapDynamicMapping80211K2edl(o["80211k"], d, "n80211k")); err != nil {
		if vv, ok := fortiAPIPatch(o["80211k"], "ObjectWirelessControllerVapDynamicMapping-80211K"); ok {
			if err = d.Set("n80211k", vv); err != nil {
				return fmt.Errorf("Error reading n80211k: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n80211k: %v", err)
		}
	}

	if err = d.Set("n80211v", flattenObjectWirelessControllerVapDynamicMapping80211V2edl(o["80211v"], d, "n80211v")); err != nil {
		if vv, ok := fortiAPIPatch(o["80211v"], "ObjectWirelessControllerVapDynamicMapping-80211V"); ok {
			if err = d.Set("n80211v", vv); err != nil {
				return fmt.Errorf("Error reading n80211v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n80211v: %v", err)
		}
	}

	if err = d.Set("_centmgmt", flattenObjectWirelessControllerVapDynamicMappingCentmgmt2edl(o["_centmgmt"], d, "_centmgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["_centmgmt"], "ObjectWirelessControllerVapDynamicMapping-Centmgmt"); ok {
			if err = d.Set("_centmgmt", vv); err != nil {
				return fmt.Errorf("Error reading _centmgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _centmgmt: %v", err)
		}
	}

	if err = d.Set("_dhcp_svr_id", flattenObjectWirelessControllerVapDynamicMappingDhcpSvrId2edl(o["_dhcp_svr_id"], d, "_dhcp_svr_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dhcp_svr_id"], "ObjectWirelessControllerVapDynamicMapping-DhcpSvrId"); ok {
			if err = d.Set("_dhcp_svr_id", vv); err != nil {
				return fmt.Errorf("Error reading _dhcp_svr_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dhcp_svr_id: %v", err)
		}
	}

	if err = d.Set("_intf_allowaccess", flattenObjectWirelessControllerVapDynamicMappingIntfAllowaccess2edl(o["_intf_allowaccess"], d, "_intf_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_allowaccess"], "ObjectWirelessControllerVapDynamicMapping-IntfAllowaccess"); ok {
			if err = d.Set("_intf_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading _intf_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_allowaccess: %v", err)
		}
	}

	if err = d.Set("_intf_device_access_list", flattenObjectWirelessControllerVapDynamicMappingIntfDeviceAccessList2edl(o["_intf_device-access-list"], d, "_intf_device_access_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_device-access-list"], "ObjectWirelessControllerVapDynamicMapping-IntfDeviceAccessList"); ok {
			if err = d.Set("_intf_device_access_list", vv); err != nil {
				return fmt.Errorf("Error reading _intf_device_access_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_device_access_list: %v", err)
		}
	}

	if err = d.Set("_intf_device_identification", flattenObjectWirelessControllerVapDynamicMappingIntfDeviceIdentification2edl(o["_intf_device-identification"], d, "_intf_device_identification")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_device-identification"], "ObjectWirelessControllerVapDynamicMapping-IntfDeviceIdentification"); ok {
			if err = d.Set("_intf_device_identification", vv); err != nil {
				return fmt.Errorf("Error reading _intf_device_identification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_device_identification: %v", err)
		}
	}

	if err = d.Set("_intf_device_netscan", flattenObjectWirelessControllerVapDynamicMappingIntfDeviceNetscan2edl(o["_intf_device-netscan"], d, "_intf_device_netscan")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_device-netscan"], "ObjectWirelessControllerVapDynamicMapping-IntfDeviceNetscan"); ok {
			if err = d.Set("_intf_device_netscan", vv); err != nil {
				return fmt.Errorf("Error reading _intf_device_netscan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_device_netscan: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp_relay_ip", flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayIp2edl(o["_intf_dhcp-relay-ip"], d, "_intf_dhcp_relay_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp-relay-ip"], "ObjectWirelessControllerVapDynamicMapping-IntfDhcpRelayIp"); ok {
			if err = d.Set("_intf_dhcp_relay_ip", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp_relay_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp_relay_ip: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp_relay_service", flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayService2edl(o["_intf_dhcp-relay-service"], d, "_intf_dhcp_relay_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp-relay-service"], "ObjectWirelessControllerVapDynamicMapping-IntfDhcpRelayService"); ok {
			if err = d.Set("_intf_dhcp_relay_service", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp_relay_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp_relay_service: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp_relay_type", flattenObjectWirelessControllerVapDynamicMappingIntfDhcpRelayType2edl(o["_intf_dhcp-relay-type"], d, "_intf_dhcp_relay_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp-relay-type"], "ObjectWirelessControllerVapDynamicMapping-IntfDhcpRelayType"); ok {
			if err = d.Set("_intf_dhcp_relay_type", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp_relay_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp_relay_type: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp6_relay_ip", flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayIp2edl(o["_intf_dhcp6-relay-ip"], d, "_intf_dhcp6_relay_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp6-relay-ip"], "ObjectWirelessControllerVapDynamicMapping-IntfDhcp6RelayIp"); ok {
			if err = d.Set("_intf_dhcp6_relay_ip", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp6_relay_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp6_relay_ip: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp6_relay_service", flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayService2edl(o["_intf_dhcp6-relay-service"], d, "_intf_dhcp6_relay_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp6-relay-service"], "ObjectWirelessControllerVapDynamicMapping-IntfDhcp6RelayService"); ok {
			if err = d.Set("_intf_dhcp6_relay_service", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp6_relay_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp6_relay_service: %v", err)
		}
	}

	if err = d.Set("_intf_dhcp6_relay_type", flattenObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayType2edl(o["_intf_dhcp6-relay-type"], d, "_intf_dhcp6_relay_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_dhcp6-relay-type"], "ObjectWirelessControllerVapDynamicMapping-IntfDhcp6RelayType"); ok {
			if err = d.Set("_intf_dhcp6_relay_type", vv); err != nil {
				return fmt.Errorf("Error reading _intf_dhcp6_relay_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_dhcp6_relay_type: %v", err)
		}
	}

	if err = d.Set("_intf_ip", flattenObjectWirelessControllerVapDynamicMappingIntfIp2edl(o["_intf_ip"], d, "_intf_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_ip"], "ObjectWirelessControllerVapDynamicMapping-IntfIp"); ok {
			if err = d.Set("_intf_ip", vv); err != nil {
				return fmt.Errorf("Error reading _intf_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_ip: %v", err)
		}
	}

	if err = d.Set("_intf_ip6_address", flattenObjectWirelessControllerVapDynamicMappingIntfIp6Address2edl(o["_intf_ip6-address"], d, "_intf_ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_ip6-address"], "ObjectWirelessControllerVapDynamicMapping-IntfIp6Address"); ok {
			if err = d.Set("_intf_ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading _intf_ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_ip6_address: %v", err)
		}
	}

	if err = d.Set("_intf_ip6_allowaccess", flattenObjectWirelessControllerVapDynamicMappingIntfIp6Allowaccess2edl(o["_intf_ip6-allowaccess"], d, "_intf_ip6_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_ip6-allowaccess"], "ObjectWirelessControllerVapDynamicMapping-IntfIp6Allowaccess"); ok {
			if err = d.Set("_intf_ip6_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading _intf_ip6_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_ip6_allowaccess: %v", err)
		}
	}

	if err = d.Set("_intf_listen_forticlient_connection", flattenObjectWirelessControllerVapDynamicMappingIntfListenForticlientConnection2edl(o["_intf_listen-forticlient-connection"], d, "_intf_listen_forticlient_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["_intf_listen-forticlient-connection"], "ObjectWirelessControllerVapDynamicMapping-IntfListenForticlientConnection"); ok {
			if err = d.Set("_intf_listen_forticlient_connection", vv); err != nil {
				return fmt.Errorf("Error reading _intf_listen_forticlient_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _intf_listen_forticlient_connection: %v", err)
		}
	}

	if err = d.Set("_is_factory_setting", flattenObjectWirelessControllerVapDynamicMappingIsFactorySetting2edl(o["_is_factory_setting"], d, "_is_factory_setting")); err != nil {
		if vv, ok := fortiAPIPatch(o["_is_factory_setting"], "ObjectWirelessControllerVapDynamicMapping-IsFactorySetting"); ok {
			if err = d.Set("_is_factory_setting", vv); err != nil {
				return fmt.Errorf("Error reading _is_factory_setting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _is_factory_setting: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectWirelessControllerVapDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectWirelessControllerVapDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectWirelessControllerVapDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectWirelessControllerVapDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("access_control_list", flattenObjectWirelessControllerVapDynamicMappingAccessControlList2edl(o["access-control-list"], d, "access_control_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-control-list"], "ObjectWirelessControllerVapDynamicMapping-AccessControlList"); ok {
			if err = d.Set("access_control_list", vv); err != nil {
				return fmt.Errorf("Error reading access_control_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_control_list: %v", err)
		}
	}

	if err = d.Set("acct_interim_interval", flattenObjectWirelessControllerVapDynamicMappingAcctInterimInterval2edl(o["acct-interim-interval"], d, "acct_interim_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-interim-interval"], "ObjectWirelessControllerVapDynamicMapping-AcctInterimInterval"); ok {
			if err = d.Set("acct_interim_interval", vv); err != nil {
				return fmt.Errorf("Error reading acct_interim_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_interim_interval: %v", err)
		}
	}

	if err = d.Set("additional_akms", flattenObjectWirelessControllerVapDynamicMappingAdditionalAkms2edl(o["additional-akms"], d, "additional_akms")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-akms"], "ObjectWirelessControllerVapDynamicMapping-AdditionalAkms"); ok {
			if err = d.Set("additional_akms", vv); err != nil {
				return fmt.Errorf("Error reading additional_akms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_akms: %v", err)
		}
	}

	if err = d.Set("address_group", flattenObjectWirelessControllerVapDynamicMappingAddressGroup2edl(o["address-group"], d, "address_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-group"], "ObjectWirelessControllerVapDynamicMapping-AddressGroup"); ok {
			if err = d.Set("address_group", vv); err != nil {
				return fmt.Errorf("Error reading address_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_group: %v", err)
		}
	}

	if err = d.Set("address_group_policy", flattenObjectWirelessControllerVapDynamicMappingAddressGroupPolicy2edl(o["address-group-policy"], d, "address_group_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-group-policy"], "ObjectWirelessControllerVapDynamicMapping-AddressGroupPolicy"); ok {
			if err = d.Set("address_group_policy", vv); err != nil {
				return fmt.Errorf("Error reading address_group_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_group_policy: %v", err)
		}
	}

	if err = d.Set("alias", flattenObjectWirelessControllerVapDynamicMappingAlias2edl(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "ObjectWirelessControllerVapDynamicMapping-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("antivirus_profile", flattenObjectWirelessControllerVapDynamicMappingAntivirusProfile2edl(o["antivirus-profile"], d, "antivirus_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["antivirus-profile"], "ObjectWirelessControllerVapDynamicMapping-AntivirusProfile"); ok {
			if err = d.Set("antivirus_profile", vv); err != nil {
				return fmt.Errorf("Error reading antivirus_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antivirus_profile: %v", err)
		}
	}

	if err = d.Set("application_detection_engine", flattenObjectWirelessControllerVapDynamicMappingApplicationDetectionEngine2edl(o["application-detection-engine"], d, "application_detection_engine")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-detection-engine"], "ObjectWirelessControllerVapDynamicMapping-ApplicationDetectionEngine"); ok {
			if err = d.Set("application_detection_engine", vv); err != nil {
				return fmt.Errorf("Error reading application_detection_engine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_detection_engine: %v", err)
		}
	}

	if err = d.Set("application_dscp_marking", flattenObjectWirelessControllerVapDynamicMappingApplicationDscpMarking2edl(o["application-dscp-marking"], d, "application_dscp_marking")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-dscp-marking"], "ObjectWirelessControllerVapDynamicMapping-ApplicationDscpMarking"); ok {
			if err = d.Set("application_dscp_marking", vv); err != nil {
				return fmt.Errorf("Error reading application_dscp_marking: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_dscp_marking: %v", err)
		}
	}

	if err = d.Set("application_list", flattenObjectWirelessControllerVapDynamicMappingApplicationList2edl(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "ObjectWirelessControllerVapDynamicMapping-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("application_report_intv", flattenObjectWirelessControllerVapDynamicMappingApplicationReportIntv2edl(o["application-report-intv"], d, "application_report_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-report-intv"], "ObjectWirelessControllerVapDynamicMapping-ApplicationReportIntv"); ok {
			if err = d.Set("application_report_intv", vv); err != nil {
				return fmt.Errorf("Error reading application_report_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_report_intv: %v", err)
		}
	}

	if err = d.Set("atf_weight", flattenObjectWirelessControllerVapDynamicMappingAtfWeight2edl(o["atf-weight"], d, "atf_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["atf-weight"], "ObjectWirelessControllerVapDynamicMapping-AtfWeight"); ok {
			if err = d.Set("atf_weight", vv); err != nil {
				return fmt.Errorf("Error reading atf_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading atf_weight: %v", err)
		}
	}

	if err = d.Set("auth", flattenObjectWirelessControllerVapDynamicMappingAuth2edl(o["auth"], d, "auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth"], "ObjectWirelessControllerVapDynamicMapping-Auth"); ok {
			if err = d.Set("auth", vv); err != nil {
				return fmt.Errorf("Error reading auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenObjectWirelessControllerVapDynamicMappingAuthCert2edl(o["auth-cert"], d, "auth_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-cert"], "ObjectWirelessControllerVapDynamicMapping-AuthCert"); ok {
			if err = d.Set("auth_cert", vv); err != nil {
				return fmt.Errorf("Error reading auth_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_portal_addr", flattenObjectWirelessControllerVapDynamicMappingAuthPortalAddr2edl(o["auth-portal-addr"], d, "auth_portal_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal-addr"], "ObjectWirelessControllerVapDynamicMapping-AuthPortalAddr"); ok {
			if err = d.Set("auth_portal_addr", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal_addr: %v", err)
		}
	}

	if err = d.Set("beacon_advertising", flattenObjectWirelessControllerVapDynamicMappingBeaconAdvertising2edl(o["beacon-advertising"], d, "beacon_advertising")); err != nil {
		if vv, ok := fortiAPIPatch(o["beacon-advertising"], "ObjectWirelessControllerVapDynamicMapping-BeaconAdvertising"); ok {
			if err = d.Set("beacon_advertising", vv); err != nil {
				return fmt.Errorf("Error reading beacon_advertising: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beacon_advertising: %v", err)
		}
	}

	if err = d.Set("broadcast_ssid", flattenObjectWirelessControllerVapDynamicMappingBroadcastSsid2edl(o["broadcast-ssid"], d, "broadcast_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast-ssid"], "ObjectWirelessControllerVapDynamicMapping-BroadcastSsid"); ok {
			if err = d.Set("broadcast_ssid", vv); err != nil {
				return fmt.Errorf("Error reading broadcast_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast_ssid: %v", err)
		}
	}

	if err = d.Set("broadcast_suppression", flattenObjectWirelessControllerVapDynamicMappingBroadcastSuppression2edl(o["broadcast-suppression"], d, "broadcast_suppression")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast-suppression"], "ObjectWirelessControllerVapDynamicMapping-BroadcastSuppression"); ok {
			if err = d.Set("broadcast_suppression", vv); err != nil {
				return fmt.Errorf("Error reading broadcast_suppression: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast_suppression: %v", err)
		}
	}

	if err = d.Set("bss_color_partial", flattenObjectWirelessControllerVapDynamicMappingBssColorPartial2edl(o["bss-color-partial"], d, "bss_color_partial")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-color-partial"], "ObjectWirelessControllerVapDynamicMapping-BssColorPartial"); ok {
			if err = d.Set("bss_color_partial", vv); err != nil {
				return fmt.Errorf("Error reading bss_color_partial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_color_partial: %v", err)
		}
	}

	if err = d.Set("bstm_disassociation_imminent", flattenObjectWirelessControllerVapDynamicMappingBstmDisassociationImminent2edl(o["bstm-disassociation-imminent"], d, "bstm_disassociation_imminent")); err != nil {
		if vv, ok := fortiAPIPatch(o["bstm-disassociation-imminent"], "ObjectWirelessControllerVapDynamicMapping-BstmDisassociationImminent"); ok {
			if err = d.Set("bstm_disassociation_imminent", vv); err != nil {
				return fmt.Errorf("Error reading bstm_disassociation_imminent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bstm_disassociation_imminent: %v", err)
		}
	}

	if err = d.Set("bstm_load_balancing_disassoc_timer", flattenObjectWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer2edl(o["bstm-load-balancing-disassoc-timer"], d, "bstm_load_balancing_disassoc_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["bstm-load-balancing-disassoc-timer"], "ObjectWirelessControllerVapDynamicMapping-BstmLoadBalancingDisassocTimer"); ok {
			if err = d.Set("bstm_load_balancing_disassoc_timer", vv); err != nil {
				return fmt.Errorf("Error reading bstm_load_balancing_disassoc_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bstm_load_balancing_disassoc_timer: %v", err)
		}
	}

	if err = d.Set("bstm_rssi_disassoc_timer", flattenObjectWirelessControllerVapDynamicMappingBstmRssiDisassocTimer2edl(o["bstm-rssi-disassoc-timer"], d, "bstm_rssi_disassoc_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["bstm-rssi-disassoc-timer"], "ObjectWirelessControllerVapDynamicMapping-BstmRssiDisassocTimer"); ok {
			if err = d.Set("bstm_rssi_disassoc_timer", vv); err != nil {
				return fmt.Errorf("Error reading bstm_rssi_disassoc_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bstm_rssi_disassoc_timer: %v", err)
		}
	}

	if err = d.Set("captive_portal_ac_name", flattenObjectWirelessControllerVapDynamicMappingCaptivePortalAcName2edl(o["captive-portal-ac-name"], d, "captive_portal_ac_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-ac-name"], "ObjectWirelessControllerVapDynamicMapping-CaptivePortalAcName"); ok {
			if err = d.Set("captive_portal_ac_name", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_ac_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_ac_name: %v", err)
		}
	}

	if err = d.Set("captive_portal_auth_timeout", flattenObjectWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout2edl(o["captive-portal-auth-timeout"], d, "captive_portal_auth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-auth-timeout"], "ObjectWirelessControllerVapDynamicMapping-CaptivePortalAuthTimeout"); ok {
			if err = d.Set("captive_portal_auth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_auth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_auth_timeout: %v", err)
		}
	}

	if err = d.Set("captive_portal_fw_accounting", flattenObjectWirelessControllerVapDynamicMappingCaptivePortalFwAccounting2edl(o["captive-portal-fw-accounting"], d, "captive_portal_fw_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-fw-accounting"], "ObjectWirelessControllerVapDynamicMapping-CaptivePortalFwAccounting"); ok {
			if err = d.Set("captive_portal_fw_accounting", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_fw_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_fw_accounting: %v", err)
		}
	}

	if err = d.Set("captive_portal_macauth_radius_server", flattenObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer2edl(o["captive-portal-macauth-radius-server"], d, "captive_portal_macauth_radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-macauth-radius-server"], "ObjectWirelessControllerVapDynamicMapping-CaptivePortalMacauthRadiusServer"); ok {
			if err = d.Set("captive_portal_macauth_radius_server", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_macauth_radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_macauth_radius_server: %v", err)
		}
	}

	if err = d.Set("captive_portal_radius_server", flattenObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusServer2edl(o["captive-portal-radius-server"], d, "captive_portal_radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-radius-server"], "ObjectWirelessControllerVapDynamicMapping-CaptivePortalRadiusServer"); ok {
			if err = d.Set("captive_portal_radius_server", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_radius_server: %v", err)
		}
	}

	if err = d.Set("captive_portal_session_timeout_interval", flattenObjectWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval2edl(o["captive-portal-session-timeout-interval"], d, "captive_portal_session_timeout_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-session-timeout-interval"], "ObjectWirelessControllerVapDynamicMapping-CaptivePortalSessionTimeoutInterval"); ok {
			if err = d.Set("captive_portal_session_timeout_interval", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_session_timeout_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_session_timeout_interval: %v", err)
		}
	}

	if err = d.Set("client_count", flattenObjectWirelessControllerVapDynamicMappingClientCount2edl(o["client-count"], d, "client_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-count"], "ObjectWirelessControllerVapDynamicMapping-ClientCount"); ok {
			if err = d.Set("client_count", vv); err != nil {
				return fmt.Errorf("Error reading client_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_count: %v", err)
		}
	}

	if err = d.Set("dhcp_address_enforcement", flattenObjectWirelessControllerVapDynamicMappingDhcpAddressEnforcement2edl(o["dhcp-address-enforcement"], d, "dhcp_address_enforcement")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-address-enforcement"], "ObjectWirelessControllerVapDynamicMapping-DhcpAddressEnforcement"); ok {
			if err = d.Set("dhcp_address_enforcement", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_address_enforcement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_address_enforcement: %v", err)
		}
	}

	if err = d.Set("dhcp_lease_time", flattenObjectWirelessControllerVapDynamicMappingDhcpLeaseTime2edl(o["dhcp-lease-time"], d, "dhcp_lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-lease-time"], "ObjectWirelessControllerVapDynamicMapping-DhcpLeaseTime"); ok {
			if err = d.Set("dhcp_lease_time", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_lease_time: %v", err)
		}
	}

	if err = d.Set("dhcp_option43_insertion", flattenObjectWirelessControllerVapDynamicMappingDhcpOption43Insertion2edl(o["dhcp-option43-insertion"], d, "dhcp_option43_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option43-insertion"], "ObjectWirelessControllerVapDynamicMapping-DhcpOption43Insertion"); ok {
			if err = d.Set("dhcp_option43_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option43_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option43_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_circuit_id_insertion", flattenObjectWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion2edl(o["dhcp-option82-circuit-id-insertion"], d, "dhcp_option82_circuit_id_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-circuit-id-insertion"], "ObjectWirelessControllerVapDynamicMapping-DhcpOption82CircuitIdInsertion"); ok {
			if err = d.Set("dhcp_option82_circuit_id_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_circuit_id_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_circuit_id_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_insertion", flattenObjectWirelessControllerVapDynamicMappingDhcpOption82Insertion2edl(o["dhcp-option82-insertion"], d, "dhcp_option82_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-insertion"], "ObjectWirelessControllerVapDynamicMapping-DhcpOption82Insertion"); ok {
			if err = d.Set("dhcp_option82_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_remote_id_insertion", flattenObjectWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion2edl(o["dhcp-option82-remote-id-insertion"], d, "dhcp_option82_remote_id_insertion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-remote-id-insertion"], "ObjectWirelessControllerVapDynamicMapping-DhcpOption82RemoteIdInsertion"); ok {
			if err = d.Set("dhcp_option82_remote_id_insertion", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_remote_id_insertion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_remote_id_insertion: %v", err)
		}
	}

	if err = d.Set("dynamic_vlan", flattenObjectWirelessControllerVapDynamicMappingDynamicVlan2edl(o["dynamic-vlan"], d, "dynamic_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-vlan"], "ObjectWirelessControllerVapDynamicMapping-DynamicVlan"); ok {
			if err = d.Set("dynamic_vlan", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_vlan: %v", err)
		}
	}

	if err = d.Set("eap_reauth", flattenObjectWirelessControllerVapDynamicMappingEapReauth2edl(o["eap-reauth"], d, "eap_reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-reauth"], "ObjectWirelessControllerVapDynamicMapping-EapReauth"); ok {
			if err = d.Set("eap_reauth", vv); err != nil {
				return fmt.Errorf("Error reading eap_reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_reauth: %v", err)
		}
	}

	if err = d.Set("eap_reauth_intv", flattenObjectWirelessControllerVapDynamicMappingEapReauthIntv2edl(o["eap-reauth-intv"], d, "eap_reauth_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-reauth-intv"], "ObjectWirelessControllerVapDynamicMapping-EapReauthIntv"); ok {
			if err = d.Set("eap_reauth_intv", vv); err != nil {
				return fmt.Errorf("Error reading eap_reauth_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_reauth_intv: %v", err)
		}
	}

	if err = d.Set("eapol_key_retries", flattenObjectWirelessControllerVapDynamicMappingEapolKeyRetries2edl(o["eapol-key-retries"], d, "eapol_key_retries")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-key-retries"], "ObjectWirelessControllerVapDynamicMapping-EapolKeyRetries"); ok {
			if err = d.Set("eapol_key_retries", vv); err != nil {
				return fmt.Errorf("Error reading eapol_key_retries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_key_retries: %v", err)
		}
	}

	if err = d.Set("encrypt", flattenObjectWirelessControllerVapDynamicMappingEncrypt2edl(o["encrypt"], d, "encrypt")); err != nil {
		if vv, ok := fortiAPIPatch(o["encrypt"], "ObjectWirelessControllerVapDynamicMapping-Encrypt"); ok {
			if err = d.Set("encrypt", vv); err != nil {
				return fmt.Errorf("Error reading encrypt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encrypt: %v", err)
		}
	}

	if err = d.Set("external_fast_roaming", flattenObjectWirelessControllerVapDynamicMappingExternalFastRoaming2edl(o["external-fast-roaming"], d, "external_fast_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-fast-roaming"], "ObjectWirelessControllerVapDynamicMapping-ExternalFastRoaming"); ok {
			if err = d.Set("external_fast_roaming", vv); err != nil {
				return fmt.Errorf("Error reading external_fast_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_fast_roaming: %v", err)
		}
	}

	if err = d.Set("external_logout", flattenObjectWirelessControllerVapDynamicMappingExternalLogout2edl(o["external-logout"], d, "external_logout")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-logout"], "ObjectWirelessControllerVapDynamicMapping-ExternalLogout"); ok {
			if err = d.Set("external_logout", vv); err != nil {
				return fmt.Errorf("Error reading external_logout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_logout: %v", err)
		}
	}

	if err = d.Set("external_web", flattenObjectWirelessControllerVapDynamicMappingExternalWeb2edl(o["external-web"], d, "external_web")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-web"], "ObjectWirelessControllerVapDynamicMapping-ExternalWeb"); ok {
			if err = d.Set("external_web", vv); err != nil {
				return fmt.Errorf("Error reading external_web: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_web: %v", err)
		}
	}

	if err = d.Set("external_web_format", flattenObjectWirelessControllerVapDynamicMappingExternalWebFormat2edl(o["external-web-format"], d, "external_web_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-web-format"], "ObjectWirelessControllerVapDynamicMapping-ExternalWebFormat"); ok {
			if err = d.Set("external_web_format", vv); err != nil {
				return fmt.Errorf("Error reading external_web_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_web_format: %v", err)
		}
	}

	if err = d.Set("fast_bss_transition", flattenObjectWirelessControllerVapDynamicMappingFastBssTransition2edl(o["fast-bss-transition"], d, "fast_bss_transition")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-bss-transition"], "ObjectWirelessControllerVapDynamicMapping-FastBssTransition"); ok {
			if err = d.Set("fast_bss_transition", vv); err != nil {
				return fmt.Errorf("Error reading fast_bss_transition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_bss_transition: %v", err)
		}
	}

	if err = d.Set("fast_roaming", flattenObjectWirelessControllerVapDynamicMappingFastRoaming2edl(o["fast-roaming"], d, "fast_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-roaming"], "ObjectWirelessControllerVapDynamicMapping-FastRoaming"); ok {
			if err = d.Set("fast_roaming", vv); err != nil {
				return fmt.Errorf("Error reading fast_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_roaming: %v", err)
		}
	}

	if err = d.Set("ft_mobility_domain", flattenObjectWirelessControllerVapDynamicMappingFtMobilityDomain2edl(o["ft-mobility-domain"], d, "ft_mobility_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["ft-mobility-domain"], "ObjectWirelessControllerVapDynamicMapping-FtMobilityDomain"); ok {
			if err = d.Set("ft_mobility_domain", vv); err != nil {
				return fmt.Errorf("Error reading ft_mobility_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ft_mobility_domain: %v", err)
		}
	}

	if err = d.Set("ft_over_ds", flattenObjectWirelessControllerVapDynamicMappingFtOverDs2edl(o["ft-over-ds"], d, "ft_over_ds")); err != nil {
		if vv, ok := fortiAPIPatch(o["ft-over-ds"], "ObjectWirelessControllerVapDynamicMapping-FtOverDs"); ok {
			if err = d.Set("ft_over_ds", vv); err != nil {
				return fmt.Errorf("Error reading ft_over_ds: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ft_over_ds: %v", err)
		}
	}

	if err = d.Set("ft_r0_key_lifetime", flattenObjectWirelessControllerVapDynamicMappingFtR0KeyLifetime2edl(o["ft-r0-key-lifetime"], d, "ft_r0_key_lifetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["ft-r0-key-lifetime"], "ObjectWirelessControllerVapDynamicMapping-FtR0KeyLifetime"); ok {
			if err = d.Set("ft_r0_key_lifetime", vv); err != nil {
				return fmt.Errorf("Error reading ft_r0_key_lifetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ft_r0_key_lifetime: %v", err)
		}
	}

	if err = d.Set("gas_comeback_delay", flattenObjectWirelessControllerVapDynamicMappingGasComebackDelay2edl(o["gas-comeback-delay"], d, "gas_comeback_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["gas-comeback-delay"], "ObjectWirelessControllerVapDynamicMapping-GasComebackDelay"); ok {
			if err = d.Set("gas_comeback_delay", vv); err != nil {
				return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
		}
	}

	if err = d.Set("gas_fragmentation_limit", flattenObjectWirelessControllerVapDynamicMappingGasFragmentationLimit2edl(o["gas-fragmentation-limit"], d, "gas_fragmentation_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["gas-fragmentation-limit"], "ObjectWirelessControllerVapDynamicMapping-GasFragmentationLimit"); ok {
			if err = d.Set("gas_fragmentation_limit", vv); err != nil {
				return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
		}
	}

	if err = d.Set("gtk_rekey", flattenObjectWirelessControllerVapDynamicMappingGtkRekey2edl(o["gtk-rekey"], d, "gtk_rekey")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtk-rekey"], "ObjectWirelessControllerVapDynamicMapping-GtkRekey"); ok {
			if err = d.Set("gtk_rekey", vv); err != nil {
				return fmt.Errorf("Error reading gtk_rekey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtk_rekey: %v", err)
		}
	}

	if err = d.Set("gtk_rekey_intv", flattenObjectWirelessControllerVapDynamicMappingGtkRekeyIntv2edl(o["gtk-rekey-intv"], d, "gtk_rekey_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtk-rekey-intv"], "ObjectWirelessControllerVapDynamicMapping-GtkRekeyIntv"); ok {
			if err = d.Set("gtk_rekey_intv", vv); err != nil {
				return fmt.Errorf("Error reading gtk_rekey_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtk_rekey_intv: %v", err)
		}
	}

	if err = d.Set("high_efficiency", flattenObjectWirelessControllerVapDynamicMappingHighEfficiency2edl(o["high-efficiency"], d, "high_efficiency")); err != nil {
		if vv, ok := fortiAPIPatch(o["high-efficiency"], "ObjectWirelessControllerVapDynamicMapping-HighEfficiency"); ok {
			if err = d.Set("high_efficiency", vv); err != nil {
				return fmt.Errorf("Error reading high_efficiency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading high_efficiency: %v", err)
		}
	}

	if err = d.Set("hotspot20_profile", flattenObjectWirelessControllerVapDynamicMappingHotspot20Profile2edl(o["hotspot20-profile"], d, "hotspot20_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["hotspot20-profile"], "ObjectWirelessControllerVapDynamicMapping-Hotspot20Profile"); ok {
			if err = d.Set("hotspot20_profile", vv); err != nil {
				return fmt.Errorf("Error reading hotspot20_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hotspot20_profile: %v", err)
		}
	}

	if err = d.Set("igmp_snooping", flattenObjectWirelessControllerVapDynamicMappingIgmpSnooping2edl(o["igmp-snooping"], d, "igmp_snooping")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmp-snooping"], "ObjectWirelessControllerVapDynamicMapping-IgmpSnooping"); ok {
			if err = d.Set("igmp_snooping", vv); err != nil {
				return fmt.Errorf("Error reading igmp_snooping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmp_snooping: %v", err)
		}
	}

	if err = d.Set("intra_vap_privacy", flattenObjectWirelessControllerVapDynamicMappingIntraVapPrivacy2edl(o["intra-vap-privacy"], d, "intra_vap_privacy")); err != nil {
		if vv, ok := fortiAPIPatch(o["intra-vap-privacy"], "ObjectWirelessControllerVapDynamicMapping-IntraVapPrivacy"); ok {
			if err = d.Set("intra_vap_privacy", vv); err != nil {
				return fmt.Errorf("Error reading intra_vap_privacy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intra_vap_privacy: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectWirelessControllerVapDynamicMappingIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectWirelessControllerVapDynamicMapping-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenObjectWirelessControllerVapDynamicMappingIpsSensor2edl(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "ObjectWirelessControllerVapDynamicMapping-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ipv6_rules", flattenObjectWirelessControllerVapDynamicMappingIpv6Rules2edl(o["ipv6-rules"], d, "ipv6_rules")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-rules"], "ObjectWirelessControllerVapDynamicMapping-Ipv6Rules"); ok {
			if err = d.Set("ipv6_rules", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_rules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_rules: %v", err)
		}
	}

	if err = d.Set("keyindex", flattenObjectWirelessControllerVapDynamicMappingKeyindex2edl(o["keyindex"], d, "keyindex")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyindex"], "ObjectWirelessControllerVapDynamicMapping-Keyindex"); ok {
			if err = d.Set("keyindex", vv); err != nil {
				return fmt.Errorf("Error reading keyindex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyindex: %v", err)
		}
	}

	if err = d.Set("l3_roaming", flattenObjectWirelessControllerVapDynamicMappingL3Roaming2edl(o["l3-roaming"], d, "l3_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["l3-roaming"], "ObjectWirelessControllerVapDynamicMapping-L3Roaming"); ok {
			if err = d.Set("l3_roaming", vv); err != nil {
				return fmt.Errorf("Error reading l3_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l3_roaming: %v", err)
		}
	}

	if err = d.Set("l3_roaming_mode", flattenObjectWirelessControllerVapDynamicMappingL3RoamingMode2edl(o["l3-roaming-mode"], d, "l3_roaming_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["l3-roaming-mode"], "ObjectWirelessControllerVapDynamicMapping-L3RoamingMode"); ok {
			if err = d.Set("l3_roaming_mode", vv); err != nil {
				return fmt.Errorf("Error reading l3_roaming_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l3_roaming_mode: %v", err)
		}
	}

	if err = d.Set("ldpc", flattenObjectWirelessControllerVapDynamicMappingLdpc2edl(o["ldpc"], d, "ldpc")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldpc"], "ObjectWirelessControllerVapDynamicMapping-Ldpc"); ok {
			if err = d.Set("ldpc", vv); err != nil {
				return fmt.Errorf("Error reading ldpc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldpc: %v", err)
		}
	}

	if err = d.Set("local_authentication", flattenObjectWirelessControllerVapDynamicMappingLocalAuthentication2edl(o["local-authentication"], d, "local_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-authentication"], "ObjectWirelessControllerVapDynamicMapping-LocalAuthentication"); ok {
			if err = d.Set("local_authentication", vv); err != nil {
				return fmt.Errorf("Error reading local_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_authentication: %v", err)
		}
	}

	if err = d.Set("local_bridging", flattenObjectWirelessControllerVapDynamicMappingLocalBridging2edl(o["local-bridging"], d, "local_bridging")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-bridging"], "ObjectWirelessControllerVapDynamicMapping-LocalBridging"); ok {
			if err = d.Set("local_bridging", vv); err != nil {
				return fmt.Errorf("Error reading local_bridging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_bridging: %v", err)
		}
	}

	if err = d.Set("local_lan", flattenObjectWirelessControllerVapDynamicMappingLocalLan2edl(o["local-lan"], d, "local_lan")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-lan"], "ObjectWirelessControllerVapDynamicMapping-LocalLan"); ok {
			if err = d.Set("local_lan", vv); err != nil {
				return fmt.Errorf("Error reading local_lan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_lan: %v", err)
		}
	}

	if err = d.Set("local_standalone", flattenObjectWirelessControllerVapDynamicMappingLocalStandalone2edl(o["local-standalone"], d, "local_standalone")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone"], "ObjectWirelessControllerVapDynamicMapping-LocalStandalone"); ok {
			if err = d.Set("local_standalone", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone: %v", err)
		}
	}

	if err = d.Set("local_standalone_dns", flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneDns2edl(o["local-standalone-dns"], d, "local_standalone_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone-dns"], "ObjectWirelessControllerVapDynamicMapping-LocalStandaloneDns"); ok {
			if err = d.Set("local_standalone_dns", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone_dns: %v", err)
		}
	}

	if err = d.Set("local_standalone_dns_ip", flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneDnsIp2edl(o["local-standalone-dns-ip"], d, "local_standalone_dns_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone-dns-ip"], "ObjectWirelessControllerVapDynamicMapping-LocalStandaloneDnsIp"); ok {
			if err = d.Set("local_standalone_dns_ip", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone_dns_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone_dns_ip: %v", err)
		}
	}

	if err = d.Set("local_standalone_nat", flattenObjectWirelessControllerVapDynamicMappingLocalStandaloneNat2edl(o["local-standalone-nat"], d, "local_standalone_nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-standalone-nat"], "ObjectWirelessControllerVapDynamicMapping-LocalStandaloneNat"); ok {
			if err = d.Set("local_standalone_nat", vv); err != nil {
				return fmt.Errorf("Error reading local_standalone_nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_standalone_nat: %v", err)
		}
	}

	if err = d.Set("local_switching", flattenObjectWirelessControllerVapDynamicMappingLocalSwitching2edl(o["local-switching"], d, "local_switching")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-switching"], "ObjectWirelessControllerVapDynamicMapping-LocalSwitching"); ok {
			if err = d.Set("local_switching", vv); err != nil {
				return fmt.Errorf("Error reading local_switching: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_switching: %v", err)
		}
	}

	if err = d.Set("mac_auth_bypass", flattenObjectWirelessControllerVapDynamicMappingMacAuthBypass2edl(o["mac-auth-bypass"], d, "mac_auth_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-auth-bypass"], "ObjectWirelessControllerVapDynamicMapping-MacAuthBypass"); ok {
			if err = d.Set("mac_auth_bypass", vv); err != nil {
				return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("mac_called_station_delimiter", flattenObjectWirelessControllerVapDynamicMappingMacCalledStationDelimiter2edl(o["mac-called-station-delimiter"], d, "mac_called_station_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-called-station-delimiter"], "ObjectWirelessControllerVapDynamicMapping-MacCalledStationDelimiter"); ok {
			if err = d.Set("mac_called_station_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_calling_station_delimiter", flattenObjectWirelessControllerVapDynamicMappingMacCallingStationDelimiter2edl(o["mac-calling-station-delimiter"], d, "mac_calling_station_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-calling-station-delimiter"], "ObjectWirelessControllerVapDynamicMapping-MacCallingStationDelimiter"); ok {
			if err = d.Set("mac_calling_station_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_case", flattenObjectWirelessControllerVapDynamicMappingMacCase2edl(o["mac-case"], d, "mac_case")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-case"], "ObjectWirelessControllerVapDynamicMapping-MacCase"); ok {
			if err = d.Set("mac_case", vv); err != nil {
				return fmt.Errorf("Error reading mac_case: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_case: %v", err)
		}
	}

	if err = d.Set("mac_filter", flattenObjectWirelessControllerVapDynamicMappingMacFilter2edl(o["mac-filter"], d, "mac_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-filter"], "ObjectWirelessControllerVapDynamicMapping-MacFilter"); ok {
			if err = d.Set("mac_filter", vv); err != nil {
				return fmt.Errorf("Error reading mac_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_filter: %v", err)
		}
	}

	if err = d.Set("mac_filter_policy_other", flattenObjectWirelessControllerVapDynamicMappingMacFilterPolicyOther2edl(o["mac-filter-policy-other"], d, "mac_filter_policy_other")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-filter-policy-other"], "ObjectWirelessControllerVapDynamicMapping-MacFilterPolicyOther"); ok {
			if err = d.Set("mac_filter_policy_other", vv); err != nil {
				return fmt.Errorf("Error reading mac_filter_policy_other: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_filter_policy_other: %v", err)
		}
	}

	if err = d.Set("mac_password_delimiter", flattenObjectWirelessControllerVapDynamicMappingMacPasswordDelimiter2edl(o["mac-password-delimiter"], d, "mac_password_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-password-delimiter"], "ObjectWirelessControllerVapDynamicMapping-MacPasswordDelimiter"); ok {
			if err = d.Set("mac_password_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_username_delimiter", flattenObjectWirelessControllerVapDynamicMappingMacUsernameDelimiter2edl(o["mac-username-delimiter"], d, "mac_username_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-username-delimiter"], "ObjectWirelessControllerVapDynamicMapping-MacUsernameDelimiter"); ok {
			if err = d.Set("mac_username_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenObjectWirelessControllerVapDynamicMappingMaxClients2edl(o["max-clients"], d, "max_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients"], "ObjectWirelessControllerVapDynamicMapping-MaxClients"); ok {
			if err = d.Set("max_clients", vv); err != nil {
				return fmt.Errorf("Error reading max_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("max_clients_ap", flattenObjectWirelessControllerVapDynamicMappingMaxClientsAp2edl(o["max-clients-ap"], d, "max_clients_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients-ap"], "ObjectWirelessControllerVapDynamicMapping-MaxClientsAp"); ok {
			if err = d.Set("max_clients_ap", vv); err != nil {
				return fmt.Errorf("Error reading max_clients_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients_ap: %v", err)
		}
	}

	if err = d.Set("mbo", flattenObjectWirelessControllerVapDynamicMappingMbo2edl(o["mbo"], d, "mbo")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbo"], "ObjectWirelessControllerVapDynamicMapping-Mbo"); ok {
			if err = d.Set("mbo", vv); err != nil {
				return fmt.Errorf("Error reading mbo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbo: %v", err)
		}
	}

	if err = d.Set("mbo_cell_data_conn_pref", flattenObjectWirelessControllerVapDynamicMappingMboCellDataConnPref2edl(o["mbo-cell-data-conn-pref"], d, "mbo_cell_data_conn_pref")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbo-cell-data-conn-pref"], "ObjectWirelessControllerVapDynamicMapping-MboCellDataConnPref"); ok {
			if err = d.Set("mbo_cell_data_conn_pref", vv); err != nil {
				return fmt.Errorf("Error reading mbo_cell_data_conn_pref: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbo_cell_data_conn_pref: %v", err)
		}
	}

	if err = d.Set("me_disable_thresh", flattenObjectWirelessControllerVapDynamicMappingMeDisableThresh2edl(o["me-disable-thresh"], d, "me_disable_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["me-disable-thresh"], "ObjectWirelessControllerVapDynamicMapping-MeDisableThresh"); ok {
			if err = d.Set("me_disable_thresh", vv); err != nil {
				return fmt.Errorf("Error reading me_disable_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading me_disable_thresh: %v", err)
		}
	}

	if err = d.Set("mesh_backhaul", flattenObjectWirelessControllerVapDynamicMappingMeshBackhaul2edl(o["mesh-backhaul"], d, "mesh_backhaul")); err != nil {
		if vv, ok := fortiAPIPatch(o["mesh-backhaul"], "ObjectWirelessControllerVapDynamicMapping-MeshBackhaul"); ok {
			if err = d.Set("mesh_backhaul", vv); err != nil {
				return fmt.Errorf("Error reading mesh_backhaul: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mesh_backhaul: %v", err)
		}
	}

	if err = d.Set("mpsk", flattenObjectWirelessControllerVapDynamicMappingMpsk2edl(o["mpsk"], d, "mpsk")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk"], "ObjectWirelessControllerVapDynamicMapping-Mpsk"); ok {
			if err = d.Set("mpsk", vv); err != nil {
				return fmt.Errorf("Error reading mpsk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk: %v", err)
		}
	}

	if err = d.Set("mpsk_concurrent_clients", flattenObjectWirelessControllerVapDynamicMappingMpskConcurrentClients2edl(o["mpsk-concurrent-clients"], d, "mpsk_concurrent_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-concurrent-clients"], "ObjectWirelessControllerVapDynamicMapping-MpskConcurrentClients"); ok {
			if err = d.Set("mpsk_concurrent_clients", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
		}
	}

	if err = d.Set("mpsk_profile", flattenObjectWirelessControllerVapDynamicMappingMpskProfile2edl(o["mpsk-profile"], d, "mpsk_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-profile"], "ObjectWirelessControllerVapDynamicMapping-MpskProfile"); ok {
			if err = d.Set("mpsk_profile", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_profile: %v", err)
		}
	}

	if err = d.Set("mu_mimo", flattenObjectWirelessControllerVapDynamicMappingMuMimo2edl(o["mu-mimo"], d, "mu_mimo")); err != nil {
		if vv, ok := fortiAPIPatch(o["mu-mimo"], "ObjectWirelessControllerVapDynamicMapping-MuMimo"); ok {
			if err = d.Set("mu_mimo", vv); err != nil {
				return fmt.Errorf("Error reading mu_mimo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mu_mimo: %v", err)
		}
	}

	if err = d.Set("multicast_enhance", flattenObjectWirelessControllerVapDynamicMappingMulticastEnhance2edl(o["multicast-enhance"], d, "multicast_enhance")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-enhance"], "ObjectWirelessControllerVapDynamicMapping-MulticastEnhance"); ok {
			if err = d.Set("multicast_enhance", vv); err != nil {
				return fmt.Errorf("Error reading multicast_enhance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_enhance: %v", err)
		}
	}

	if err = d.Set("multicast_rate", flattenObjectWirelessControllerVapDynamicMappingMulticastRate2edl(o["multicast-rate"], d, "multicast_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-rate"], "ObjectWirelessControllerVapDynamicMapping-MulticastRate"); ok {
			if err = d.Set("multicast_rate", vv); err != nil {
				return fmt.Errorf("Error reading multicast_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_rate: %v", err)
		}
	}

	if err = d.Set("nac", flattenObjectWirelessControllerVapDynamicMappingNac2edl(o["nac"], d, "nac")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac"], "ObjectWirelessControllerVapDynamicMapping-Nac"); ok {
			if err = d.Set("nac", vv); err != nil {
				return fmt.Errorf("Error reading nac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac: %v", err)
		}
	}

	if err = d.Set("nac_profile", flattenObjectWirelessControllerVapDynamicMappingNacProfile2edl(o["nac-profile"], d, "nac_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-profile"], "ObjectWirelessControllerVapDynamicMapping-NacProfile"); ok {
			if err = d.Set("nac_profile", vv); err != nil {
				return fmt.Errorf("Error reading nac_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_profile: %v", err)
		}
	}

	if err = d.Set("neighbor_report_dual_band", flattenObjectWirelessControllerVapDynamicMappingNeighborReportDualBand2edl(o["neighbor-report-dual-band"], d, "neighbor_report_dual_band")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-report-dual-band"], "ObjectWirelessControllerVapDynamicMapping-NeighborReportDualBand"); ok {
			if err = d.Set("neighbor_report_dual_band", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_report_dual_band: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_report_dual_band: %v", err)
		}
	}

	if err = d.Set("okc", flattenObjectWirelessControllerVapDynamicMappingOkc2edl(o["okc"], d, "okc")); err != nil {
		if vv, ok := fortiAPIPatch(o["okc"], "ObjectWirelessControllerVapDynamicMapping-Okc"); ok {
			if err = d.Set("okc", vv); err != nil {
				return fmt.Errorf("Error reading okc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading okc: %v", err)
		}
	}

	if err = d.Set("osen", flattenObjectWirelessControllerVapDynamicMappingOsen2edl(o["osen"], d, "osen")); err != nil {
		if vv, ok := fortiAPIPatch(o["osen"], "ObjectWirelessControllerVapDynamicMapping-Osen"); ok {
			if err = d.Set("osen", vv); err != nil {
				return fmt.Errorf("Error reading osen: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osen: %v", err)
		}
	}

	if err = d.Set("owe_groups", flattenObjectWirelessControllerVapDynamicMappingOweGroups2edl(o["owe-groups"], d, "owe_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["owe-groups"], "ObjectWirelessControllerVapDynamicMapping-OweGroups"); ok {
			if err = d.Set("owe_groups", vv); err != nil {
				return fmt.Errorf("Error reading owe_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owe_groups: %v", err)
		}
	}

	if err = d.Set("owe_transition", flattenObjectWirelessControllerVapDynamicMappingOweTransition2edl(o["owe-transition"], d, "owe_transition")); err != nil {
		if vv, ok := fortiAPIPatch(o["owe-transition"], "ObjectWirelessControllerVapDynamicMapping-OweTransition"); ok {
			if err = d.Set("owe_transition", vv); err != nil {
				return fmt.Errorf("Error reading owe_transition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owe_transition: %v", err)
		}
	}

	if err = d.Set("owe_transition_ssid", flattenObjectWirelessControllerVapDynamicMappingOweTransitionSsid2edl(o["owe-transition-ssid"], d, "owe_transition_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["owe-transition-ssid"], "ObjectWirelessControllerVapDynamicMapping-OweTransitionSsid"); ok {
			if err = d.Set("owe_transition_ssid", vv); err != nil {
				return fmt.Errorf("Error reading owe_transition_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owe_transition_ssid: %v", err)
		}
	}

	if err = d.Set("pmf", flattenObjectWirelessControllerVapDynamicMappingPmf2edl(o["pmf"], d, "pmf")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmf"], "ObjectWirelessControllerVapDynamicMapping-Pmf"); ok {
			if err = d.Set("pmf", vv); err != nil {
				return fmt.Errorf("Error reading pmf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmf: %v", err)
		}
	}

	if err = d.Set("pmf_assoc_comeback_timeout", flattenObjectWirelessControllerVapDynamicMappingPmfAssocComebackTimeout2edl(o["pmf-assoc-comeback-timeout"], d, "pmf_assoc_comeback_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmf-assoc-comeback-timeout"], "ObjectWirelessControllerVapDynamicMapping-PmfAssocComebackTimeout"); ok {
			if err = d.Set("pmf_assoc_comeback_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pmf_assoc_comeback_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmf_assoc_comeback_timeout: %v", err)
		}
	}

	if err = d.Set("pmf_sa_query_retry_timeout", flattenObjectWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout2edl(o["pmf-sa-query-retry-timeout"], d, "pmf_sa_query_retry_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmf-sa-query-retry-timeout"], "ObjectWirelessControllerVapDynamicMapping-PmfSaQueryRetryTimeout"); ok {
			if err = d.Set("pmf_sa_query_retry_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pmf_sa_query_retry_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmf_sa_query_retry_timeout: %v", err)
		}
	}

	if err = d.Set("port_macauth", flattenObjectWirelessControllerVapDynamicMappingPortMacauth2edl(o["port-macauth"], d, "port_macauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-macauth"], "ObjectWirelessControllerVapDynamicMapping-PortMacauth"); ok {
			if err = d.Set("port_macauth", vv); err != nil {
				return fmt.Errorf("Error reading port_macauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_macauth: %v", err)
		}
	}

	if err = d.Set("port_macauth_reauth_timeout", flattenObjectWirelessControllerVapDynamicMappingPortMacauthReauthTimeout2edl(o["port-macauth-reauth-timeout"], d, "port_macauth_reauth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-macauth-reauth-timeout"], "ObjectWirelessControllerVapDynamicMapping-PortMacauthReauthTimeout"); ok {
			if err = d.Set("port_macauth_reauth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading port_macauth_reauth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_macauth_reauth_timeout: %v", err)
		}
	}

	if err = d.Set("port_macauth_timeout", flattenObjectWirelessControllerVapDynamicMappingPortMacauthTimeout2edl(o["port-macauth-timeout"], d, "port_macauth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-macauth-timeout"], "ObjectWirelessControllerVapDynamicMapping-PortMacauthTimeout"); ok {
			if err = d.Set("port_macauth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading port_macauth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_macauth_timeout: %v", err)
		}
	}

	if err = d.Set("portal_message_override_group", flattenObjectWirelessControllerVapDynamicMappingPortalMessageOverrideGroup2edl(o["portal-message-override-group"], d, "portal_message_override_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["portal-message-override-group"], "ObjectWirelessControllerVapDynamicMapping-PortalMessageOverrideGroup"); ok {
			if err = d.Set("portal_message_override_group", vv); err != nil {
				return fmt.Errorf("Error reading portal_message_override_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portal_message_override_group: %v", err)
		}
	}

	if err = d.Set("portal_type", flattenObjectWirelessControllerVapDynamicMappingPortalType2edl(o["portal-type"], d, "portal_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["portal-type"], "ObjectWirelessControllerVapDynamicMapping-PortalType"); ok {
			if err = d.Set("portal_type", vv); err != nil {
				return fmt.Errorf("Error reading portal_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portal_type: %v", err)
		}
	}

	if err = d.Set("primary_wag_profile", flattenObjectWirelessControllerVapDynamicMappingPrimaryWagProfile2edl(o["primary-wag-profile"], d, "primary_wag_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-wag-profile"], "ObjectWirelessControllerVapDynamicMapping-PrimaryWagProfile"); ok {
			if err = d.Set("primary_wag_profile", vv); err != nil {
				return fmt.Errorf("Error reading primary_wag_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_wag_profile: %v", err)
		}
	}

	if err = d.Set("probe_resp_suppression", flattenObjectWirelessControllerVapDynamicMappingProbeRespSuppression2edl(o["probe-resp-suppression"], d, "probe_resp_suppression")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-resp-suppression"], "ObjectWirelessControllerVapDynamicMapping-ProbeRespSuppression"); ok {
			if err = d.Set("probe_resp_suppression", vv); err != nil {
				return fmt.Errorf("Error reading probe_resp_suppression: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_resp_suppression: %v", err)
		}
	}

	if err = d.Set("probe_resp_threshold", flattenObjectWirelessControllerVapDynamicMappingProbeRespThreshold2edl(o["probe-resp-threshold"], d, "probe_resp_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-resp-threshold"], "ObjectWirelessControllerVapDynamicMapping-ProbeRespThreshold"); ok {
			if err = d.Set("probe_resp_threshold", vv); err != nil {
				return fmt.Errorf("Error reading probe_resp_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_resp_threshold: %v", err)
		}
	}

	if err = d.Set("ptk_rekey", flattenObjectWirelessControllerVapDynamicMappingPtkRekey2edl(o["ptk-rekey"], d, "ptk_rekey")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptk-rekey"], "ObjectWirelessControllerVapDynamicMapping-PtkRekey"); ok {
			if err = d.Set("ptk_rekey", vv); err != nil {
				return fmt.Errorf("Error reading ptk_rekey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptk_rekey: %v", err)
		}
	}

	if err = d.Set("ptk_rekey_intv", flattenObjectWirelessControllerVapDynamicMappingPtkRekeyIntv2edl(o["ptk-rekey-intv"], d, "ptk_rekey_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptk-rekey-intv"], "ObjectWirelessControllerVapDynamicMapping-PtkRekeyIntv"); ok {
			if err = d.Set("ptk_rekey_intv", vv); err != nil {
				return fmt.Errorf("Error reading ptk_rekey_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptk_rekey_intv: %v", err)
		}
	}

	if err = d.Set("qos_profile", flattenObjectWirelessControllerVapDynamicMappingQosProfile2edl(o["qos-profile"], d, "qos_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-profile"], "ObjectWirelessControllerVapDynamicMapping-QosProfile"); ok {
			if err = d.Set("qos_profile", vv); err != nil {
				return fmt.Errorf("Error reading qos_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_profile: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenObjectWirelessControllerVapDynamicMappingQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "ObjectWirelessControllerVapDynamicMapping-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("radio_2g_threshold", flattenObjectWirelessControllerVapDynamicMappingRadio2GThreshold2edl(o["radio-2g-threshold"], d, "radio_2g_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-2g-threshold"], "ObjectWirelessControllerVapDynamicMapping-Radio2GThreshold"); ok {
			if err = d.Set("radio_2g_threshold", vv); err != nil {
				return fmt.Errorf("Error reading radio_2g_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_2g_threshold: %v", err)
		}
	}

	if err = d.Set("radio_5g_threshold", flattenObjectWirelessControllerVapDynamicMappingRadio5GThreshold2edl(o["radio-5g-threshold"], d, "radio_5g_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-5g-threshold"], "ObjectWirelessControllerVapDynamicMapping-Radio5GThreshold"); ok {
			if err = d.Set("radio_5g_threshold", vv); err != nil {
				return fmt.Errorf("Error reading radio_5g_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_5g_threshold: %v", err)
		}
	}

	if err = d.Set("radio_sensitivity", flattenObjectWirelessControllerVapDynamicMappingRadioSensitivity2edl(o["radio-sensitivity"], d, "radio_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-sensitivity"], "ObjectWirelessControllerVapDynamicMapping-RadioSensitivity"); ok {
			if err = d.Set("radio_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading radio_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_sensitivity: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth", flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuth2edl(o["radius-mac-auth"], d, "radius_mac_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth"], "ObjectWirelessControllerVapDynamicMapping-RadiusMacAuth"); ok {
			if err = d.Set("radius_mac_auth", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_block_interval", flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval2edl(o["radius-mac-auth-block-interval"], d, "radius_mac_auth_block_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-block-interval"], "ObjectWirelessControllerVapDynamicMapping-RadiusMacAuthBlockInterval"); ok {
			if err = d.Set("radius_mac_auth_block_interval", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_block_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_block_interval: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_server", flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthServer2edl(o["radius-mac-auth-server"], d, "radius_mac_auth_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-server"], "ObjectWirelessControllerVapDynamicMapping-RadiusMacAuthServer"); ok {
			if err = d.Set("radius_mac_auth_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_server: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_usergroups", flattenObjectWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups2edl(o["radius-mac-auth-usergroups"], d, "radius_mac_auth_usergroups")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-auth-usergroups"], "ObjectWirelessControllerVapDynamicMapping-RadiusMacAuthUsergroups"); ok {
			if err = d.Set("radius_mac_auth_usergroups", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_auth_usergroups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_auth_usergroups: %v", err)
		}
	}

	if err = d.Set("radius_mac_mpsk_auth", flattenObjectWirelessControllerVapDynamicMappingRadiusMacMpskAuth2edl(o["radius-mac-mpsk-auth"], d, "radius_mac_mpsk_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-mpsk-auth"], "ObjectWirelessControllerVapDynamicMapping-RadiusMacMpskAuth"); ok {
			if err = d.Set("radius_mac_mpsk_auth", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_mpsk_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_mpsk_auth: %v", err)
		}
	}

	if err = d.Set("radius_mac_mpsk_timeout", flattenObjectWirelessControllerVapDynamicMappingRadiusMacMpskTimeout2edl(o["radius-mac-mpsk-timeout"], d, "radius_mac_mpsk_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-mac-mpsk-timeout"], "ObjectWirelessControllerVapDynamicMapping-RadiusMacMpskTimeout"); ok {
			if err = d.Set("radius_mac_mpsk_timeout", vv); err != nil {
				return fmt.Errorf("Error reading radius_mac_mpsk_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_mac_mpsk_timeout: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenObjectWirelessControllerVapDynamicMappingRadiusServer2edl(o["radius-server"], d, "radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-server"], "ObjectWirelessControllerVapDynamicMapping-RadiusServer"); ok {
			if err = d.Set("radius_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("rates_11a", flattenObjectWirelessControllerVapDynamicMappingRates11A2edl(o["rates-11a"], d, "rates_11a")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11a"], "ObjectWirelessControllerVapDynamicMapping-Rates11A"); ok {
			if err = d.Set("rates_11a", vv); err != nil {
				return fmt.Errorf("Error reading rates_11a: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11a: %v", err)
		}
	}

	if err = d.Set("rates_11ac_mcs_map", flattenObjectWirelessControllerVapDynamicMappingRates11AcMcsMap2edl(o["rates-11ac-mcs-map"], d, "rates_11ac_mcs_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ac-mcs-map"], "ObjectWirelessControllerVapDynamicMapping-Rates11AcMcsMap"); ok {
			if err = d.Set("rates_11ac_mcs_map", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ac_mcs_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ac_mcs_map: %v", err)
		}
	}

	if err = d.Set("rates_11ac_ss12", flattenObjectWirelessControllerVapDynamicMappingRates11AcSs122edl(o["rates-11ac-ss12"], d, "rates_11ac_ss12")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ac-ss12"], "ObjectWirelessControllerVapDynamicMapping-Rates11AcSs12"); ok {
			if err = d.Set("rates_11ac_ss12", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ac_ss12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ac_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11ac_ss34", flattenObjectWirelessControllerVapDynamicMappingRates11AcSs342edl(o["rates-11ac-ss34"], d, "rates_11ac_ss34")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ac-ss34"], "ObjectWirelessControllerVapDynamicMapping-Rates11AcSs34"); ok {
			if err = d.Set("rates_11ac_ss34", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ac_ss34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ac_ss34: %v", err)
		}
	}

	if err = d.Set("rates_11ax_mcs_map", flattenObjectWirelessControllerVapDynamicMappingRates11AxMcsMap2edl(o["rates-11ax-mcs-map"], d, "rates_11ax_mcs_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ax-mcs-map"], "ObjectWirelessControllerVapDynamicMapping-Rates11AxMcsMap"); ok {
			if err = d.Set("rates_11ax_mcs_map", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ax_mcs_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ax_mcs_map: %v", err)
		}
	}

	if err = d.Set("rates_11ax_ss12", flattenObjectWirelessControllerVapDynamicMappingRates11AxSs122edl(o["rates-11ax-ss12"], d, "rates_11ax_ss12")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ax-ss12"], "ObjectWirelessControllerVapDynamicMapping-Rates11AxSs12"); ok {
			if err = d.Set("rates_11ax_ss12", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ax_ss12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ax_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11ax_ss34", flattenObjectWirelessControllerVapDynamicMappingRates11AxSs342edl(o["rates-11ax-ss34"], d, "rates_11ax_ss34")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11ax-ss34"], "ObjectWirelessControllerVapDynamicMapping-Rates11AxSs34"); ok {
			if err = d.Set("rates_11ax_ss34", vv); err != nil {
				return fmt.Errorf("Error reading rates_11ax_ss34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11ax_ss34: %v", err)
		}
	}

	if err = d.Set("rates_11bg", flattenObjectWirelessControllerVapDynamicMappingRates11Bg2edl(o["rates-11bg"], d, "rates_11bg")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11bg"], "ObjectWirelessControllerVapDynamicMapping-Rates11Bg"); ok {
			if err = d.Set("rates_11bg", vv); err != nil {
				return fmt.Errorf("Error reading rates_11bg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11bg: %v", err)
		}
	}

	if err = d.Set("rates_11n_ss12", flattenObjectWirelessControllerVapDynamicMappingRates11NSs122edl(o["rates-11n-ss12"], d, "rates_11n_ss12")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11n-ss12"], "ObjectWirelessControllerVapDynamicMapping-Rates11NSs12"); ok {
			if err = d.Set("rates_11n_ss12", vv); err != nil {
				return fmt.Errorf("Error reading rates_11n_ss12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11n_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11n_ss34", flattenObjectWirelessControllerVapDynamicMappingRates11NSs342edl(o["rates-11n-ss34"], d, "rates_11n_ss34")); err != nil {
		if vv, ok := fortiAPIPatch(o["rates-11n-ss34"], "ObjectWirelessControllerVapDynamicMapping-Rates11NSs34"); ok {
			if err = d.Set("rates_11n_ss34", vv); err != nil {
				return fmt.Errorf("Error reading rates_11n_ss34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rates_11n_ss34: %v", err)
		}
	}

	if err = d.Set("roaming_acct_interim_update", flattenObjectWirelessControllerVapDynamicMappingRoamingAcctInterimUpdate2edl(o["roaming-acct-interim-update"], d, "roaming_acct_interim_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["roaming-acct-interim-update"], "ObjectWirelessControllerVapDynamicMapping-RoamingAcctInterimUpdate"); ok {
			if err = d.Set("roaming_acct_interim_update", vv); err != nil {
				return fmt.Errorf("Error reading roaming_acct_interim_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roaming_acct_interim_update: %v", err)
		}
	}

	if err = d.Set("sae_groups", flattenObjectWirelessControllerVapDynamicMappingSaeGroups2edl(o["sae-groups"], d, "sae_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-groups"], "ObjectWirelessControllerVapDynamicMapping-SaeGroups"); ok {
			if err = d.Set("sae_groups", vv); err != nil {
				return fmt.Errorf("Error reading sae_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_groups: %v", err)
		}
	}

	if err = d.Set("sae_h2e_only", flattenObjectWirelessControllerVapDynamicMappingSaeH2EOnly2edl(o["sae-h2e-only"], d, "sae_h2e_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-h2e-only"], "ObjectWirelessControllerVapDynamicMapping-SaeH2EOnly"); ok {
			if err = d.Set("sae_h2e_only", vv); err != nil {
				return fmt.Errorf("Error reading sae_h2e_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_h2e_only: %v", err)
		}
	}

	if err = d.Set("sae_hnp_only", flattenObjectWirelessControllerVapDynamicMappingSaeHnpOnly2edl(o["sae-hnp-only"], d, "sae_hnp_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-hnp-only"], "ObjectWirelessControllerVapDynamicMapping-SaeHnpOnly"); ok {
			if err = d.Set("sae_hnp_only", vv); err != nil {
				return fmt.Errorf("Error reading sae_hnp_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_hnp_only: %v", err)
		}
	}

	if err = d.Set("sae_pk", flattenObjectWirelessControllerVapDynamicMappingSaePk2edl(o["sae-pk"], d, "sae_pk")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-pk"], "ObjectWirelessControllerVapDynamicMapping-SaePk"); ok {
			if err = d.Set("sae_pk", vv); err != nil {
				return fmt.Errorf("Error reading sae_pk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_pk: %v", err)
		}
	}

	if err = d.Set("sae_private_key", flattenObjectWirelessControllerVapDynamicMappingSaePrivateKey2edl(o["sae-private-key"], d, "sae_private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-private-key"], "ObjectWirelessControllerVapDynamicMapping-SaePrivateKey"); ok {
			if err = d.Set("sae_private_key", vv); err != nil {
				return fmt.Errorf("Error reading sae_private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_private_key: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenObjectWirelessControllerVapDynamicMappingScanBotnetConnections2edl(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-botnet-connections"], "ObjectWirelessControllerVapDynamicMapping-ScanBotnetConnections"); ok {
			if err = d.Set("scan_botnet_connections", vv); err != nil {
				return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("schedule", flattenObjectWirelessControllerVapDynamicMappingSchedule2edl(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "ObjectWirelessControllerVapDynamicMapping-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("secondary_wag_profile", flattenObjectWirelessControllerVapDynamicMappingSecondaryWagProfile2edl(o["secondary-wag-profile"], d, "secondary_wag_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-wag-profile"], "ObjectWirelessControllerVapDynamicMapping-SecondaryWagProfile"); ok {
			if err = d.Set("secondary_wag_profile", vv); err != nil {
				return fmt.Errorf("Error reading secondary_wag_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_wag_profile: %v", err)
		}
	}

	if err = d.Set("security", flattenObjectWirelessControllerVapDynamicMappingSecurity2edl(o["security"], d, "security")); err != nil {
		if vv, ok := fortiAPIPatch(o["security"], "ObjectWirelessControllerVapDynamicMapping-Security"); ok {
			if err = d.Set("security", vv); err != nil {
				return fmt.Errorf("Error reading security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("security_exempt_list", flattenObjectWirelessControllerVapDynamicMappingSecurityExemptList2edl(o["security-exempt-list"], d, "security_exempt_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-exempt-list"], "ObjectWirelessControllerVapDynamicMapping-SecurityExemptList"); ok {
			if err = d.Set("security_exempt_list", vv); err != nil {
				return fmt.Errorf("Error reading security_exempt_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_exempt_list: %v", err)
		}
	}

	if err = d.Set("security_obsolete_option", flattenObjectWirelessControllerVapDynamicMappingSecurityObsoleteOption2edl(o["security-obsolete-option"], d, "security_obsolete_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-obsolete-option"], "ObjectWirelessControllerVapDynamicMapping-SecurityObsoleteOption"); ok {
			if err = d.Set("security_obsolete_option", vv); err != nil {
				return fmt.Errorf("Error reading security_obsolete_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_obsolete_option: %v", err)
		}
	}

	if err = d.Set("security_redirect_url", flattenObjectWirelessControllerVapDynamicMappingSecurityRedirectUrl2edl(o["security-redirect-url"], d, "security_redirect_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-redirect-url"], "ObjectWirelessControllerVapDynamicMapping-SecurityRedirectUrl"); ok {
			if err = d.Set("security_redirect_url", vv); err != nil {
				return fmt.Errorf("Error reading security_redirect_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_redirect_url: %v", err)
		}
	}

	if err = d.Set("selected_usergroups", flattenObjectWirelessControllerVapDynamicMappingSelectedUsergroups2edl(o["selected-usergroups"], d, "selected_usergroups")); err != nil {
		if vv, ok := fortiAPIPatch(o["selected-usergroups"], "ObjectWirelessControllerVapDynamicMapping-SelectedUsergroups"); ok {
			if err = d.Set("selected_usergroups", vv); err != nil {
				return fmt.Errorf("Error reading selected_usergroups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading selected_usergroups: %v", err)
		}
	}

	if err = d.Set("split_tunneling", flattenObjectWirelessControllerVapDynamicMappingSplitTunneling2edl(o["split-tunneling"], d, "split_tunneling")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling"], "ObjectWirelessControllerVapDynamicMapping-SplitTunneling"); ok {
			if err = d.Set("split_tunneling", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling: %v", err)
		}
	}

	if err = d.Set("ssid", flattenObjectWirelessControllerVapDynamicMappingSsid2edl(o["ssid"], d, "ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssid"], "ObjectWirelessControllerVapDynamicMapping-Ssid"); ok {
			if err = d.Set("ssid", vv); err != nil {
				return fmt.Errorf("Error reading ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssid: %v", err)
		}
	}

	if err = d.Set("sticky_client_remove", flattenObjectWirelessControllerVapDynamicMappingStickyClientRemove2edl(o["sticky-client-remove"], d, "sticky_client_remove")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-remove"], "ObjectWirelessControllerVapDynamicMapping-StickyClientRemove"); ok {
			if err = d.Set("sticky_client_remove", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_remove: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_remove: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_2g", flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold2G2edl(o["sticky-client-threshold-2g"], d, "sticky_client_threshold_2g")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-threshold-2g"], "ObjectWirelessControllerVapDynamicMapping-StickyClientThreshold2G"); ok {
			if err = d.Set("sticky_client_threshold_2g", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_threshold_2g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_threshold_2g: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_5g", flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold5G2edl(o["sticky-client-threshold-5g"], d, "sticky_client_threshold_5g")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-threshold-5g"], "ObjectWirelessControllerVapDynamicMapping-StickyClientThreshold5G"); ok {
			if err = d.Set("sticky_client_threshold_5g", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_threshold_5g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_threshold_5g: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_6g", flattenObjectWirelessControllerVapDynamicMappingStickyClientThreshold6G2edl(o["sticky-client-threshold-6g"], d, "sticky_client_threshold_6g")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-client-threshold-6g"], "ObjectWirelessControllerVapDynamicMapping-StickyClientThreshold6G"); ok {
			if err = d.Set("sticky_client_threshold_6g", vv); err != nil {
				return fmt.Errorf("Error reading sticky_client_threshold_6g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_client_threshold_6g: %v", err)
		}
	}

	if err = d.Set("target_wake_time", flattenObjectWirelessControllerVapDynamicMappingTargetWakeTime2edl(o["target-wake-time"], d, "target_wake_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["target-wake-time"], "ObjectWirelessControllerVapDynamicMapping-TargetWakeTime"); ok {
			if err = d.Set("target_wake_time", vv); err != nil {
				return fmt.Errorf("Error reading target_wake_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target_wake_time: %v", err)
		}
	}

	if err = d.Set("tkip_counter_measure", flattenObjectWirelessControllerVapDynamicMappingTkipCounterMeasure2edl(o["tkip-counter-measure"], d, "tkip_counter_measure")); err != nil {
		if vv, ok := fortiAPIPatch(o["tkip-counter-measure"], "ObjectWirelessControllerVapDynamicMapping-TkipCounterMeasure"); ok {
			if err = d.Set("tkip_counter_measure", vv); err != nil {
				return fmt.Errorf("Error reading tkip_counter_measure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tkip_counter_measure: %v", err)
		}
	}

	if err = d.Set("tunnel_echo_interval", flattenObjectWirelessControllerVapDynamicMappingTunnelEchoInterval2edl(o["tunnel-echo-interval"], d, "tunnel_echo_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-echo-interval"], "ObjectWirelessControllerVapDynamicMapping-TunnelEchoInterval"); ok {
			if err = d.Set("tunnel_echo_interval", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_echo_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_echo_interval: %v", err)
		}
	}

	if err = d.Set("tunnel_fallback_interval", flattenObjectWirelessControllerVapDynamicMappingTunnelFallbackInterval2edl(o["tunnel-fallback-interval"], d, "tunnel_fallback_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-fallback-interval"], "ObjectWirelessControllerVapDynamicMapping-TunnelFallbackInterval"); ok {
			if err = d.Set("tunnel_fallback_interval", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_fallback_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_fallback_interval: %v", err)
		}
	}

	if err = d.Set("usergroup", flattenObjectWirelessControllerVapDynamicMappingUsergroup2edl(o["usergroup"], d, "usergroup")); err != nil {
		if vv, ok := fortiAPIPatch(o["usergroup"], "ObjectWirelessControllerVapDynamicMapping-Usergroup"); ok {
			if err = d.Set("usergroup", vv); err != nil {
				return fmt.Errorf("Error reading usergroup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usergroup: %v", err)
		}
	}

	if err = d.Set("utm_log", flattenObjectWirelessControllerVapDynamicMappingUtmLog2edl(o["utm-log"], d, "utm_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-log"], "ObjectWirelessControllerVapDynamicMapping-UtmLog"); ok {
			if err = d.Set("utm_log", vv); err != nil {
				return fmt.Errorf("Error reading utm_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_log: %v", err)
		}
	}

	if err = d.Set("utm_profile", flattenObjectWirelessControllerVapDynamicMappingUtmProfile2edl(o["utm-profile"], d, "utm_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-profile"], "ObjectWirelessControllerVapDynamicMapping-UtmProfile"); ok {
			if err = d.Set("utm_profile", vv); err != nil {
				return fmt.Errorf("Error reading utm_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_profile: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenObjectWirelessControllerVapDynamicMappingUtmStatus2edl(o["utm-status"], d, "utm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-status"], "ObjectWirelessControllerVapDynamicMapping-UtmStatus"); ok {
			if err = d.Set("utm_status", vv); err != nil {
				return fmt.Errorf("Error reading utm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("vdom", flattenObjectWirelessControllerVapDynamicMappingVdom2edl(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "ObjectWirelessControllerVapDynamicMapping-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vlan_auto", flattenObjectWirelessControllerVapDynamicMappingVlanAuto2edl(o["vlan-auto"], d, "vlan_auto")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-auto"], "ObjectWirelessControllerVapDynamicMapping-VlanAuto"); ok {
			if err = d.Set("vlan_auto", vv); err != nil {
				return fmt.Errorf("Error reading vlan_auto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_auto: %v", err)
		}
	}

	if err = d.Set("vlan_pooling", flattenObjectWirelessControllerVapDynamicMappingVlanPooling2edl(o["vlan-pooling"], d, "vlan_pooling")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-pooling"], "ObjectWirelessControllerVapDynamicMapping-VlanPooling"); ok {
			if err = d.Set("vlan_pooling", vv); err != nil {
				return fmt.Errorf("Error reading vlan_pooling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_pooling: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenObjectWirelessControllerVapDynamicMappingVlanid2edl(o["vlanid"], d, "vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlanid"], "ObjectWirelessControllerVapDynamicMapping-Vlanid"); ok {
			if err = d.Set("vlanid", vv); err != nil {
				return fmt.Errorf("Error reading vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("voice_enterprise", flattenObjectWirelessControllerVapDynamicMappingVoiceEnterprise2edl(o["voice-enterprise"], d, "voice_enterprise")); err != nil {
		if vv, ok := fortiAPIPatch(o["voice-enterprise"], "ObjectWirelessControllerVapDynamicMapping-VoiceEnterprise"); ok {
			if err = d.Set("voice_enterprise", vv); err != nil {
				return fmt.Errorf("Error reading voice_enterprise: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voice_enterprise: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenObjectWirelessControllerVapDynamicMappingWebfilterProfile2edl(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "ObjectWirelessControllerVapDynamicMapping-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerVapDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerVapDynamicMapping80211K2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMapping80211V2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCentmgmt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpSvrId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfAllowaccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDeviceAccessList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDeviceIdentification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDeviceNetscan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfIp6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfIp6Allowaccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingIntfListenForticlientConnection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIsFactorySetting2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectWirelessControllerVapDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectWirelessControllerVapDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerVapDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAccessControlList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingAcctInterimInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAdditionalAkms2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingAddressGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingAddressGroupPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAlias2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAntivirusProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingApplicationDetectionEngine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingApplicationDscpMarking2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingApplicationList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingApplicationReportIntv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAtfWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingAuthCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingAuthPortalAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingBeaconAdvertising2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingBroadcastSsid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingBroadcastSuppression2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingBssColorPartial2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingBstmDisassociationImminent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingBstmRssiDisassocTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalAcName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalFwAccounting2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingClientCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpAddressEnforcement2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpOption43Insertion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpOption82Insertion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingDynamicVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingEapReauth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingEapReauthIntv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingEapolKeyRetries2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingEncrypt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingExternalFastRoaming2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingExternalLogout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingExternalWeb2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingExternalWebFormat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingFastBssTransition2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingFastRoaming2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingFtMobilityDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingFtOverDs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingFtR0KeyLifetime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingGasComebackDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingGasFragmentationLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingGtkRekey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingGtkRekeyIntv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingHighEfficiency2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingHotspot20Profile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIgmpSnooping2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIntraVapPrivacy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingIpsSensor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingIpv6Rules2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingKeyindex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingL3Roaming2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingL3RoamingMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLdpc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalBridging2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalLan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalStandalone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalStandaloneDns2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalStandaloneDnsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalStandaloneNat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingLocalSwitching2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacAuthBypass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacCalledStationDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacCallingStationDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacCase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacFilterPolicyOther2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacPasswordDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMacUsernameDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMaxClients2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMaxClientsAp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMbo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMboCellDataConnPref2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMeDisableThresh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMeshBackhaul2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMpsk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMpskConcurrentClients2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMpskProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingMuMimo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMulticastEnhance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingMulticastRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingNac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingNacProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingNeighborReportDualBand2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingOkc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingOsen2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingOweGroups2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingOweTransition2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingOweTransitionSsid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPassphrase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingPmf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPmfAssocComebackTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPortMacauth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPortMacauthReauthTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPortMacauthTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPortalMessageOverrideGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPortalType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPrimaryWagProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingProbeRespSuppression2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingProbeRespThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPtkRekey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingPtkRekeyIntv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingQosProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadio2GThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadio5GThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadioSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacMpskAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusMacMpskTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRadiusServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11A2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AcMcsMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AcSs122edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AcSs342edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AxMcsMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AxSs122edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11AxSs342edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11Bg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11NSs122edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRates11NSs342edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingRoamingAcctInterimUpdate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSaeGroups2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingSaeH2EOnly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSaeHnpOnly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSaePassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerVapDynamicMappingSaePk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSaePrivateKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingScanBotnetConnections2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSchedule2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingSecondaryWagProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingSecurity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSecurityExemptList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSecurityObsoleteOption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSecurityRedirectUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSelectedUsergroups2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingSplitTunneling2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingSsid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingStickyClientRemove2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold2G2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold5G2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold6G2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingTargetWakeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingTkipCounterMeasure2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingTunnelEchoInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingTunnelFallbackInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingUsergroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingUtmLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingUtmProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingUtmStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerVapDynamicMappingVlanAuto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingVlanPooling2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingVlanid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingVoiceEnterprise2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapDynamicMappingWebfilterProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectObjectWirelessControllerVapDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n80211k"); ok || d.HasChange("n80211k") {
		t, err := expandObjectWirelessControllerVapDynamicMapping80211K2edl(d, v, "n80211k")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["80211k"] = t
		}
	}

	if v, ok := d.GetOk("n80211v"); ok || d.HasChange("n80211v") {
		t, err := expandObjectWirelessControllerVapDynamicMapping80211V2edl(d, v, "n80211v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["80211v"] = t
		}
	}

	if v, ok := d.GetOk("_centmgmt"); ok || d.HasChange("_centmgmt") {
		t, err := expandObjectWirelessControllerVapDynamicMappingCentmgmt2edl(d, v, "_centmgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_centmgmt"] = t
		}
	}

	if v, ok := d.GetOk("_dhcp_svr_id"); ok || d.HasChange("_dhcp_svr_id") {
		t, err := expandObjectWirelessControllerVapDynamicMappingDhcpSvrId2edl(d, v, "_dhcp_svr_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dhcp_svr_id"] = t
		}
	}

	if v, ok := d.GetOk("_intf_allowaccess"); ok || d.HasChange("_intf_allowaccess") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfAllowaccess2edl(d, v, "_intf_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("_intf_device_access_list"); ok || d.HasChange("_intf_device_access_list") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfDeviceAccessList2edl(d, v, "_intf_device_access_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_device-access-list"] = t
		}
	}

	if v, ok := d.GetOk("_intf_device_identification"); ok || d.HasChange("_intf_device_identification") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfDeviceIdentification2edl(d, v, "_intf_device_identification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_device-identification"] = t
		}
	}

	if v, ok := d.GetOk("_intf_device_netscan"); ok || d.HasChange("_intf_device_netscan") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfDeviceNetscan2edl(d, v, "_intf_device_netscan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_device-netscan"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp_relay_ip"); ok || d.HasChange("_intf_dhcp_relay_ip") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayIp2edl(d, v, "_intf_dhcp_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp_relay_service"); ok || d.HasChange("_intf_dhcp_relay_service") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayService2edl(d, v, "_intf_dhcp_relay_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp-relay-service"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp_relay_type"); ok || d.HasChange("_intf_dhcp_relay_type") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfDhcpRelayType2edl(d, v, "_intf_dhcp_relay_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp-relay-type"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp6_relay_ip"); ok || d.HasChange("_intf_dhcp6_relay_ip") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayIp2edl(d, v, "_intf_dhcp6_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp6-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp6_relay_service"); ok || d.HasChange("_intf_dhcp6_relay_service") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayService2edl(d, v, "_intf_dhcp6_relay_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp6-relay-service"] = t
		}
	}

	if v, ok := d.GetOk("_intf_dhcp6_relay_type"); ok || d.HasChange("_intf_dhcp6_relay_type") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfDhcp6RelayType2edl(d, v, "_intf_dhcp6_relay_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_dhcp6-relay-type"] = t
		}
	}

	if v, ok := d.GetOk("_intf_ip"); ok || d.HasChange("_intf_ip") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfIp2edl(d, v, "_intf_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_ip"] = t
		}
	}

	if v, ok := d.GetOk("_intf_ip6_address"); ok || d.HasChange("_intf_ip6_address") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfIp6Address2edl(d, v, "_intf_ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("_intf_ip6_allowaccess"); ok || d.HasChange("_intf_ip6_allowaccess") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfIp6Allowaccess2edl(d, v, "_intf_ip6_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_ip6-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("_intf_listen_forticlient_connection"); ok || d.HasChange("_intf_listen_forticlient_connection") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntfListenForticlientConnection2edl(d, v, "_intf_listen_forticlient_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_intf_listen-forticlient-connection"] = t
		}
	}

	if v, ok := d.GetOk("_is_factory_setting"); ok || d.HasChange("_is_factory_setting") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIsFactorySetting2edl(d, v, "_is_factory_setting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_is_factory_setting"] = t
		}
	}

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectWirelessControllerVapDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("access_control_list"); ok || d.HasChange("access_control_list") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAccessControlList2edl(d, v, "access_control_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-control-list"] = t
		}
	}

	if v, ok := d.GetOk("acct_interim_interval"); ok || d.HasChange("acct_interim_interval") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAcctInterimInterval2edl(d, v, "acct_interim_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-interim-interval"] = t
		}
	}

	if v, ok := d.GetOk("additional_akms"); ok || d.HasChange("additional_akms") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAdditionalAkms2edl(d, v, "additional_akms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-akms"] = t
		}
	}

	if v, ok := d.GetOk("address_group"); ok || d.HasChange("address_group") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAddressGroup2edl(d, v, "address_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-group"] = t
		}
	}

	if v, ok := d.GetOk("address_group_policy"); ok || d.HasChange("address_group_policy") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAddressGroupPolicy2edl(d, v, "address_group_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-group-policy"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok || d.HasChange("alias") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAlias2edl(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("antivirus_profile"); ok || d.HasChange("antivirus_profile") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAntivirusProfile2edl(d, v, "antivirus_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antivirus-profile"] = t
		}
	}

	if v, ok := d.GetOk("application_detection_engine"); ok || d.HasChange("application_detection_engine") {
		t, err := expandObjectWirelessControllerVapDynamicMappingApplicationDetectionEngine2edl(d, v, "application_detection_engine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-detection-engine"] = t
		}
	}

	if v, ok := d.GetOk("application_dscp_marking"); ok || d.HasChange("application_dscp_marking") {
		t, err := expandObjectWirelessControllerVapDynamicMappingApplicationDscpMarking2edl(d, v, "application_dscp_marking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-dscp-marking"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandObjectWirelessControllerVapDynamicMappingApplicationList2edl(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("application_report_intv"); ok || d.HasChange("application_report_intv") {
		t, err := expandObjectWirelessControllerVapDynamicMappingApplicationReportIntv2edl(d, v, "application_report_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("atf_weight"); ok || d.HasChange("atf_weight") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAtfWeight2edl(d, v, "atf_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["atf-weight"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok || d.HasChange("auth") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAuth2edl(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok || d.HasChange("auth_cert") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAuthCert2edl(d, v, "auth_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal_addr"); ok || d.HasChange("auth_portal_addr") {
		t, err := expandObjectWirelessControllerVapDynamicMappingAuthPortalAddr2edl(d, v, "auth_portal_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal-addr"] = t
		}
	}

	if v, ok := d.GetOk("beacon_advertising"); ok || d.HasChange("beacon_advertising") {
		t, err := expandObjectWirelessControllerVapDynamicMappingBeaconAdvertising2edl(d, v, "beacon_advertising")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-advertising"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_ssid"); ok || d.HasChange("broadcast_ssid") {
		t, err := expandObjectWirelessControllerVapDynamicMappingBroadcastSsid2edl(d, v, "broadcast_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-ssid"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_suppression"); ok || d.HasChange("broadcast_suppression") {
		t, err := expandObjectWirelessControllerVapDynamicMappingBroadcastSuppression2edl(d, v, "broadcast_suppression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-suppression"] = t
		}
	}

	if v, ok := d.GetOk("bss_color_partial"); ok || d.HasChange("bss_color_partial") {
		t, err := expandObjectWirelessControllerVapDynamicMappingBssColorPartial2edl(d, v, "bss_color_partial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color-partial"] = t
		}
	}

	if v, ok := d.GetOk("bstm_disassociation_imminent"); ok || d.HasChange("bstm_disassociation_imminent") {
		t, err := expandObjectWirelessControllerVapDynamicMappingBstmDisassociationImminent2edl(d, v, "bstm_disassociation_imminent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-disassociation-imminent"] = t
		}
	}

	if v, ok := d.GetOk("bstm_load_balancing_disassoc_timer"); ok || d.HasChange("bstm_load_balancing_disassoc_timer") {
		t, err := expandObjectWirelessControllerVapDynamicMappingBstmLoadBalancingDisassocTimer2edl(d, v, "bstm_load_balancing_disassoc_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-load-balancing-disassoc-timer"] = t
		}
	}

	if v, ok := d.GetOk("bstm_rssi_disassoc_timer"); ok || d.HasChange("bstm_rssi_disassoc_timer") {
		t, err := expandObjectWirelessControllerVapDynamicMappingBstmRssiDisassocTimer2edl(d, v, "bstm_rssi_disassoc_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-rssi-disassoc-timer"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ac_name"); ok || d.HasChange("captive_portal_ac_name") {
		t, err := expandObjectWirelessControllerVapDynamicMappingCaptivePortalAcName2edl(d, v, "captive_portal_ac_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ac-name"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_auth_timeout"); ok || d.HasChange("captive_portal_auth_timeout") {
		t, err := expandObjectWirelessControllerVapDynamicMappingCaptivePortalAuthTimeout2edl(d, v, "captive_portal_auth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-auth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_fw_accounting"); ok || d.HasChange("captive_portal_fw_accounting") {
		t, err := expandObjectWirelessControllerVapDynamicMappingCaptivePortalFwAccounting2edl(d, v, "captive_portal_fw_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-fw-accounting"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_macauth_radius_secret"); ok || d.HasChange("captive_portal_macauth_radius_secret") {
		t, err := expandObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusSecret2edl(d, v, "captive_portal_macauth_radius_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-macauth-radius-secret"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_macauth_radius_server"); ok || d.HasChange("captive_portal_macauth_radius_server") {
		t, err := expandObjectWirelessControllerVapDynamicMappingCaptivePortalMacauthRadiusServer2edl(d, v, "captive_portal_macauth_radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-macauth-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_radius_secret"); ok || d.HasChange("captive_portal_radius_secret") {
		t, err := expandObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusSecret2edl(d, v, "captive_portal_radius_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-radius-secret"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_radius_server"); ok || d.HasChange("captive_portal_radius_server") {
		t, err := expandObjectWirelessControllerVapDynamicMappingCaptivePortalRadiusServer2edl(d, v, "captive_portal_radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_session_timeout_interval"); ok || d.HasChange("captive_portal_session_timeout_interval") {
		t, err := expandObjectWirelessControllerVapDynamicMappingCaptivePortalSessionTimeoutInterval2edl(d, v, "captive_portal_session_timeout_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-session-timeout-interval"] = t
		}
	}

	if v, ok := d.GetOk("client_count"); ok || d.HasChange("client_count") {
		t, err := expandObjectWirelessControllerVapDynamicMappingClientCount2edl(d, v, "client_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-count"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_address_enforcement"); ok || d.HasChange("dhcp_address_enforcement") {
		t, err := expandObjectWirelessControllerVapDynamicMappingDhcpAddressEnforcement2edl(d, v, "dhcp_address_enforcement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-address-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_lease_time"); ok || d.HasChange("dhcp_lease_time") {
		t, err := expandObjectWirelessControllerVapDynamicMappingDhcpLeaseTime2edl(d, v, "dhcp_lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-lease-time"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option43_insertion"); ok || d.HasChange("dhcp_option43_insertion") {
		t, err := expandObjectWirelessControllerVapDynamicMappingDhcpOption43Insertion2edl(d, v, "dhcp_option43_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option43-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_circuit_id_insertion"); ok || d.HasChange("dhcp_option82_circuit_id_insertion") {
		t, err := expandObjectWirelessControllerVapDynamicMappingDhcpOption82CircuitIdInsertion2edl(d, v, "dhcp_option82_circuit_id_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-circuit-id-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_insertion"); ok || d.HasChange("dhcp_option82_insertion") {
		t, err := expandObjectWirelessControllerVapDynamicMappingDhcpOption82Insertion2edl(d, v, "dhcp_option82_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_remote_id_insertion"); ok || d.HasChange("dhcp_option82_remote_id_insertion") {
		t, err := expandObjectWirelessControllerVapDynamicMappingDhcpOption82RemoteIdInsertion2edl(d, v, "dhcp_option82_remote_id_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-remote-id-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_vlan"); ok || d.HasChange("dynamic_vlan") {
		t, err := expandObjectWirelessControllerVapDynamicMappingDynamicVlan2edl(d, v, "dynamic_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-vlan"] = t
		}
	}

	if v, ok := d.GetOk("eap_reauth"); ok || d.HasChange("eap_reauth") {
		t, err := expandObjectWirelessControllerVapDynamicMappingEapReauth2edl(d, v, "eap_reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-reauth"] = t
		}
	}

	if v, ok := d.GetOk("eap_reauth_intv"); ok || d.HasChange("eap_reauth_intv") {
		t, err := expandObjectWirelessControllerVapDynamicMappingEapReauthIntv2edl(d, v, "eap_reauth_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-reauth-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_key_retries"); ok || d.HasChange("eapol_key_retries") {
		t, err := expandObjectWirelessControllerVapDynamicMappingEapolKeyRetries2edl(d, v, "eapol_key_retries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-key-retries"] = t
		}
	}

	if v, ok := d.GetOk("encrypt"); ok || d.HasChange("encrypt") {
		t, err := expandObjectWirelessControllerVapDynamicMappingEncrypt2edl(d, v, "encrypt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encrypt"] = t
		}
	}

	if v, ok := d.GetOk("external_fast_roaming"); ok || d.HasChange("external_fast_roaming") {
		t, err := expandObjectWirelessControllerVapDynamicMappingExternalFastRoaming2edl(d, v, "external_fast_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-fast-roaming"] = t
		}
	}

	if v, ok := d.GetOk("external_logout"); ok || d.HasChange("external_logout") {
		t, err := expandObjectWirelessControllerVapDynamicMappingExternalLogout2edl(d, v, "external_logout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-logout"] = t
		}
	}

	if v, ok := d.GetOk("external_web"); ok || d.HasChange("external_web") {
		t, err := expandObjectWirelessControllerVapDynamicMappingExternalWeb2edl(d, v, "external_web")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-web"] = t
		}
	}

	if v, ok := d.GetOk("external_web_format"); ok || d.HasChange("external_web_format") {
		t, err := expandObjectWirelessControllerVapDynamicMappingExternalWebFormat2edl(d, v, "external_web_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-web-format"] = t
		}
	}

	if v, ok := d.GetOk("fast_bss_transition"); ok || d.HasChange("fast_bss_transition") {
		t, err := expandObjectWirelessControllerVapDynamicMappingFastBssTransition2edl(d, v, "fast_bss_transition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-bss-transition"] = t
		}
	}

	if v, ok := d.GetOk("fast_roaming"); ok || d.HasChange("fast_roaming") {
		t, err := expandObjectWirelessControllerVapDynamicMappingFastRoaming2edl(d, v, "fast_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-roaming"] = t
		}
	}

	if v, ok := d.GetOk("ft_mobility_domain"); ok || d.HasChange("ft_mobility_domain") {
		t, err := expandObjectWirelessControllerVapDynamicMappingFtMobilityDomain2edl(d, v, "ft_mobility_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-mobility-domain"] = t
		}
	}

	if v, ok := d.GetOk("ft_over_ds"); ok || d.HasChange("ft_over_ds") {
		t, err := expandObjectWirelessControllerVapDynamicMappingFtOverDs2edl(d, v, "ft_over_ds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-over-ds"] = t
		}
	}

	if v, ok := d.GetOk("ft_r0_key_lifetime"); ok || d.HasChange("ft_r0_key_lifetime") {
		t, err := expandObjectWirelessControllerVapDynamicMappingFtR0KeyLifetime2edl(d, v, "ft_r0_key_lifetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-r0-key-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("gas_comeback_delay"); ok || d.HasChange("gas_comeback_delay") {
		t, err := expandObjectWirelessControllerVapDynamicMappingGasComebackDelay2edl(d, v, "gas_comeback_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-comeback-delay"] = t
		}
	}

	if v, ok := d.GetOk("gas_fragmentation_limit"); ok || d.HasChange("gas_fragmentation_limit") {
		t, err := expandObjectWirelessControllerVapDynamicMappingGasFragmentationLimit2edl(d, v, "gas_fragmentation_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-fragmentation-limit"] = t
		}
	}

	if v, ok := d.GetOk("gtk_rekey"); ok || d.HasChange("gtk_rekey") {
		t, err := expandObjectWirelessControllerVapDynamicMappingGtkRekey2edl(d, v, "gtk_rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtk-rekey"] = t
		}
	}

	if v, ok := d.GetOk("gtk_rekey_intv"); ok || d.HasChange("gtk_rekey_intv") {
		t, err := expandObjectWirelessControllerVapDynamicMappingGtkRekeyIntv2edl(d, v, "gtk_rekey_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtk-rekey-intv"] = t
		}
	}

	if v, ok := d.GetOk("high_efficiency"); ok || d.HasChange("high_efficiency") {
		t, err := expandObjectWirelessControllerVapDynamicMappingHighEfficiency2edl(d, v, "high_efficiency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high-efficiency"] = t
		}
	}

	if v, ok := d.GetOk("hotspot20_profile"); ok || d.HasChange("hotspot20_profile") {
		t, err := expandObjectWirelessControllerVapDynamicMappingHotspot20Profile2edl(d, v, "hotspot20_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hotspot20-profile"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping"); ok || d.HasChange("igmp_snooping") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIgmpSnooping2edl(d, v, "igmp_snooping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("intra_vap_privacy"); ok || d.HasChange("intra_vap_privacy") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIntraVapPrivacy2edl(d, v, "intra_vap_privacy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intra-vap-privacy"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIpsSensor2edl(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_rules"); ok || d.HasChange("ipv6_rules") {
		t, err := expandObjectWirelessControllerVapDynamicMappingIpv6Rules2edl(d, v, "ipv6_rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-rules"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandObjectWirelessControllerVapDynamicMappingKey2edl(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("keyindex"); ok || d.HasChange("keyindex") {
		t, err := expandObjectWirelessControllerVapDynamicMappingKeyindex2edl(d, v, "keyindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyindex"] = t
		}
	}

	if v, ok := d.GetOk("l3_roaming"); ok || d.HasChange("l3_roaming") {
		t, err := expandObjectWirelessControllerVapDynamicMappingL3Roaming2edl(d, v, "l3_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-roaming"] = t
		}
	}

	if v, ok := d.GetOk("l3_roaming_mode"); ok || d.HasChange("l3_roaming_mode") {
		t, err := expandObjectWirelessControllerVapDynamicMappingL3RoamingMode2edl(d, v, "l3_roaming_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-roaming-mode"] = t
		}
	}

	if v, ok := d.GetOk("ldpc"); ok || d.HasChange("ldpc") {
		t, err := expandObjectWirelessControllerVapDynamicMappingLdpc2edl(d, v, "ldpc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldpc"] = t
		}
	}

	if v, ok := d.GetOk("local_authentication"); ok || d.HasChange("local_authentication") {
		t, err := expandObjectWirelessControllerVapDynamicMappingLocalAuthentication2edl(d, v, "local_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-authentication"] = t
		}
	}

	if v, ok := d.GetOk("local_bridging"); ok || d.HasChange("local_bridging") {
		t, err := expandObjectWirelessControllerVapDynamicMappingLocalBridging2edl(d, v, "local_bridging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-bridging"] = t
		}
	}

	if v, ok := d.GetOk("local_lan"); ok || d.HasChange("local_lan") {
		t, err := expandObjectWirelessControllerVapDynamicMappingLocalLan2edl(d, v, "local_lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-lan"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone"); ok || d.HasChange("local_standalone") {
		t, err := expandObjectWirelessControllerVapDynamicMappingLocalStandalone2edl(d, v, "local_standalone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_dns"); ok || d.HasChange("local_standalone_dns") {
		t, err := expandObjectWirelessControllerVapDynamicMappingLocalStandaloneDns2edl(d, v, "local_standalone_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-dns"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_dns_ip"); ok || d.HasChange("local_standalone_dns_ip") {
		t, err := expandObjectWirelessControllerVapDynamicMappingLocalStandaloneDnsIp2edl(d, v, "local_standalone_dns_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-dns-ip"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_nat"); ok || d.HasChange("local_standalone_nat") {
		t, err := expandObjectWirelessControllerVapDynamicMappingLocalStandaloneNat2edl(d, v, "local_standalone_nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-nat"] = t
		}
	}

	if v, ok := d.GetOk("local_switching"); ok || d.HasChange("local_switching") {
		t, err := expandObjectWirelessControllerVapDynamicMappingLocalSwitching2edl(d, v, "local_switching")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-switching"] = t
		}
	}

	if v, ok := d.GetOk("mac_auth_bypass"); ok || d.HasChange("mac_auth_bypass") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMacAuthBypass2edl(d, v, "mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("mac_called_station_delimiter"); ok || d.HasChange("mac_called_station_delimiter") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMacCalledStationDelimiter2edl(d, v, "mac_called_station_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-called-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_calling_station_delimiter"); ok || d.HasChange("mac_calling_station_delimiter") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMacCallingStationDelimiter2edl(d, v, "mac_calling_station_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-calling-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_case"); ok || d.HasChange("mac_case") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMacCase2edl(d, v, "mac_case")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-case"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter"); ok || d.HasChange("mac_filter") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMacFilter2edl(d, v, "mac_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter_policy_other"); ok || d.HasChange("mac_filter_policy_other") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMacFilterPolicyOther2edl(d, v, "mac_filter_policy_other")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter-policy-other"] = t
		}
	}

	if v, ok := d.GetOk("mac_password_delimiter"); ok || d.HasChange("mac_password_delimiter") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMacPasswordDelimiter2edl(d, v, "mac_password_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-password-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_username_delimiter"); ok || d.HasChange("mac_username_delimiter") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMacUsernameDelimiter2edl(d, v, "mac_username_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-username-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("max_clients"); ok || d.HasChange("max_clients") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMaxClients2edl(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("max_clients_ap"); ok || d.HasChange("max_clients_ap") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMaxClientsAp2edl(d, v, "max_clients_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients-ap"] = t
		}
	}

	if v, ok := d.GetOk("mbo"); ok || d.HasChange("mbo") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMbo2edl(d, v, "mbo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbo"] = t
		}
	}

	if v, ok := d.GetOk("mbo_cell_data_conn_pref"); ok || d.HasChange("mbo_cell_data_conn_pref") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMboCellDataConnPref2edl(d, v, "mbo_cell_data_conn_pref")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbo-cell-data-conn-pref"] = t
		}
	}

	if v, ok := d.GetOk("me_disable_thresh"); ok || d.HasChange("me_disable_thresh") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMeDisableThresh2edl(d, v, "me_disable_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["me-disable-thresh"] = t
		}
	}

	if v, ok := d.GetOk("mesh_backhaul"); ok || d.HasChange("mesh_backhaul") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMeshBackhaul2edl(d, v, "mesh_backhaul")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-backhaul"] = t
		}
	}

	if v, ok := d.GetOk("mpsk"); ok || d.HasChange("mpsk") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMpsk2edl(d, v, "mpsk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_concurrent_clients"); ok || d.HasChange("mpsk_concurrent_clients") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMpskConcurrentClients2edl(d, v, "mpsk_concurrent_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-concurrent-clients"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_profile"); ok || d.HasChange("mpsk_profile") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMpskProfile2edl(d, v, "mpsk_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-profile"] = t
		}
	}

	if v, ok := d.GetOk("mu_mimo"); ok || d.HasChange("mu_mimo") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMuMimo2edl(d, v, "mu_mimo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mu-mimo"] = t
		}
	}

	if v, ok := d.GetOk("multicast_enhance"); ok || d.HasChange("multicast_enhance") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMulticastEnhance2edl(d, v, "multicast_enhance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-enhance"] = t
		}
	}

	if v, ok := d.GetOk("multicast_rate"); ok || d.HasChange("multicast_rate") {
		t, err := expandObjectWirelessControllerVapDynamicMappingMulticastRate2edl(d, v, "multicast_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-rate"] = t
		}
	}

	if v, ok := d.GetOk("nac"); ok || d.HasChange("nac") {
		t, err := expandObjectWirelessControllerVapDynamicMappingNac2edl(d, v, "nac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac"] = t
		}
	}

	if v, ok := d.GetOk("nac_profile"); ok || d.HasChange("nac_profile") {
		t, err := expandObjectWirelessControllerVapDynamicMappingNacProfile2edl(d, v, "nac_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-profile"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_report_dual_band"); ok || d.HasChange("neighbor_report_dual_band") {
		t, err := expandObjectWirelessControllerVapDynamicMappingNeighborReportDualBand2edl(d, v, "neighbor_report_dual_band")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-report-dual-band"] = t
		}
	}

	if v, ok := d.GetOk("okc"); ok || d.HasChange("okc") {
		t, err := expandObjectWirelessControllerVapDynamicMappingOkc2edl(d, v, "okc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["okc"] = t
		}
	}

	if v, ok := d.GetOk("osen"); ok || d.HasChange("osen") {
		t, err := expandObjectWirelessControllerVapDynamicMappingOsen2edl(d, v, "osen")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osen"] = t
		}
	}

	if v, ok := d.GetOk("owe_groups"); ok || d.HasChange("owe_groups") {
		t, err := expandObjectWirelessControllerVapDynamicMappingOweGroups2edl(d, v, "owe_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-groups"] = t
		}
	}

	if v, ok := d.GetOk("owe_transition"); ok || d.HasChange("owe_transition") {
		t, err := expandObjectWirelessControllerVapDynamicMappingOweTransition2edl(d, v, "owe_transition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-transition"] = t
		}
	}

	if v, ok := d.GetOk("owe_transition_ssid"); ok || d.HasChange("owe_transition_ssid") {
		t, err := expandObjectWirelessControllerVapDynamicMappingOweTransitionSsid2edl(d, v, "owe_transition_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-transition-ssid"] = t
		}
	}

	if v, ok := d.GetOk("passphrase"); ok || d.HasChange("passphrase") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPassphrase2edl(d, v, "passphrase")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passphrase"] = t
		}
	}

	if v, ok := d.GetOk("pmf"); ok || d.HasChange("pmf") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPmf2edl(d, v, "pmf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf"] = t
		}
	}

	if v, ok := d.GetOk("pmf_assoc_comeback_timeout"); ok || d.HasChange("pmf_assoc_comeback_timeout") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPmfAssocComebackTimeout2edl(d, v, "pmf_assoc_comeback_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf-assoc-comeback-timeout"] = t
		}
	}

	if v, ok := d.GetOk("pmf_sa_query_retry_timeout"); ok || d.HasChange("pmf_sa_query_retry_timeout") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPmfSaQueryRetryTimeout2edl(d, v, "pmf_sa_query_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf-sa-query-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth"); ok || d.HasChange("port_macauth") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPortMacauth2edl(d, v, "port_macauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth_reauth_timeout"); ok || d.HasChange("port_macauth_reauth_timeout") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPortMacauthReauthTimeout2edl(d, v, "port_macauth_reauth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth-reauth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth_timeout"); ok || d.HasChange("port_macauth_timeout") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPortMacauthTimeout2edl(d, v, "port_macauth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("portal_message_override_group"); ok || d.HasChange("portal_message_override_group") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPortalMessageOverrideGroup2edl(d, v, "portal_message_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-override-group"] = t
		}
	}

	if v, ok := d.GetOk("portal_type"); ok || d.HasChange("portal_type") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPortalType2edl(d, v, "portal_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-type"] = t
		}
	}

	if v, ok := d.GetOk("primary_wag_profile"); ok || d.HasChange("primary_wag_profile") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPrimaryWagProfile2edl(d, v, "primary_wag_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-wag-profile"] = t
		}
	}

	if v, ok := d.GetOk("probe_resp_suppression"); ok || d.HasChange("probe_resp_suppression") {
		t, err := expandObjectWirelessControllerVapDynamicMappingProbeRespSuppression2edl(d, v, "probe_resp_suppression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-resp-suppression"] = t
		}
	}

	if v, ok := d.GetOk("probe_resp_threshold"); ok || d.HasChange("probe_resp_threshold") {
		t, err := expandObjectWirelessControllerVapDynamicMappingProbeRespThreshold2edl(d, v, "probe_resp_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-resp-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ptk_rekey"); ok || d.HasChange("ptk_rekey") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPtkRekey2edl(d, v, "ptk_rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptk-rekey"] = t
		}
	}

	if v, ok := d.GetOk("ptk_rekey_intv"); ok || d.HasChange("ptk_rekey_intv") {
		t, err := expandObjectWirelessControllerVapDynamicMappingPtkRekeyIntv2edl(d, v, "ptk_rekey_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptk-rekey-intv"] = t
		}
	}

	if v, ok := d.GetOk("qos_profile"); ok || d.HasChange("qos_profile") {
		t, err := expandObjectWirelessControllerVapDynamicMappingQosProfile2edl(d, v, "qos_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-profile"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandObjectWirelessControllerVapDynamicMappingQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("radio_2g_threshold"); ok || d.HasChange("radio_2g_threshold") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRadio2GThreshold2edl(d, v, "radio_2g_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("radio_5g_threshold"); ok || d.HasChange("radio_5g_threshold") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRadio5GThreshold2edl(d, v, "radio_5g_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-5g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("radio_sensitivity"); ok || d.HasChange("radio_sensitivity") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRadioSensitivity2edl(d, v, "radio_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth"); ok || d.HasChange("radius_mac_auth") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRadiusMacAuth2edl(d, v, "radius_mac_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_block_interval"); ok || d.HasChange("radius_mac_auth_block_interval") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthBlockInterval2edl(d, v, "radius_mac_auth_block_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-block-interval"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_server"); ok || d.HasChange("radius_mac_auth_server") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthServer2edl(d, v, "radius_mac_auth_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-server"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_usergroups"); ok || d.HasChange("radius_mac_auth_usergroups") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRadiusMacAuthUsergroups2edl(d, v, "radius_mac_auth_usergroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_mpsk_auth"); ok || d.HasChange("radius_mac_mpsk_auth") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRadiusMacMpskAuth2edl(d, v, "radius_mac_mpsk_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-mpsk-auth"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_mpsk_timeout"); ok || d.HasChange("radius_mac_mpsk_timeout") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRadiusMacMpskTimeout2edl(d, v, "radius_mac_mpsk_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-mpsk-timeout"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok || d.HasChange("radius_server") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRadiusServer2edl(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("rates_11a"); ok || d.HasChange("rates_11a") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRates11A2edl(d, v, "rates_11a")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11a"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_mcs_map"); ok || d.HasChange("rates_11ac_mcs_map") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRates11AcMcsMap2edl(d, v, "rates_11ac_mcs_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-mcs-map"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_ss12"); ok || d.HasChange("rates_11ac_ss12") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRates11AcSs122edl(d, v, "rates_11ac_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_ss34"); ok || d.HasChange("rates_11ac_ss34") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRates11AcSs342edl(d, v, "rates_11ac_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-ss34"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ax_mcs_map"); ok || d.HasChange("rates_11ax_mcs_map") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRates11AxMcsMap2edl(d, v, "rates_11ax_mcs_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-mcs-map"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ax_ss12"); ok || d.HasChange("rates_11ax_ss12") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRates11AxSs122edl(d, v, "rates_11ax_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ax_ss34"); ok || d.HasChange("rates_11ax_ss34") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRates11AxSs342edl(d, v, "rates_11ax_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-ss34"] = t
		}
	}

	if v, ok := d.GetOk("rates_11bg"); ok || d.HasChange("rates_11bg") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRates11Bg2edl(d, v, "rates_11bg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11bg"] = t
		}
	}

	if v, ok := d.GetOk("rates_11n_ss12"); ok || d.HasChange("rates_11n_ss12") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRates11NSs122edl(d, v, "rates_11n_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11n-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11n_ss34"); ok || d.HasChange("rates_11n_ss34") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRates11NSs342edl(d, v, "rates_11n_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11n-ss34"] = t
		}
	}

	if v, ok := d.GetOk("roaming_acct_interim_update"); ok || d.HasChange("roaming_acct_interim_update") {
		t, err := expandObjectWirelessControllerVapDynamicMappingRoamingAcctInterimUpdate2edl(d, v, "roaming_acct_interim_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming-acct-interim-update"] = t
		}
	}

	if v, ok := d.GetOk("sae_groups"); ok || d.HasChange("sae_groups") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSaeGroups2edl(d, v, "sae_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-groups"] = t
		}
	}

	if v, ok := d.GetOk("sae_h2e_only"); ok || d.HasChange("sae_h2e_only") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSaeH2EOnly2edl(d, v, "sae_h2e_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-h2e-only"] = t
		}
	}

	if v, ok := d.GetOk("sae_hnp_only"); ok || d.HasChange("sae_hnp_only") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSaeHnpOnly2edl(d, v, "sae_hnp_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-hnp-only"] = t
		}
	}

	if v, ok := d.GetOk("sae_password"); ok || d.HasChange("sae_password") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSaePassword2edl(d, v, "sae_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-password"] = t
		}
	}

	if v, ok := d.GetOk("sae_pk"); ok || d.HasChange("sae_pk") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSaePk2edl(d, v, "sae_pk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-pk"] = t
		}
	}

	if v, ok := d.GetOk("sae_private_key"); ok || d.HasChange("sae_private_key") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSaePrivateKey2edl(d, v, "sae_private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-private-key"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok || d.HasChange("scan_botnet_connections") {
		t, err := expandObjectWirelessControllerVapDynamicMappingScanBotnetConnections2edl(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSchedule2edl(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("secondary_wag_profile"); ok || d.HasChange("secondary_wag_profile") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSecondaryWagProfile2edl(d, v, "secondary_wag_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-wag-profile"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok || d.HasChange("security") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSecurity2edl(d, v, "security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("security_exempt_list"); ok || d.HasChange("security_exempt_list") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSecurityExemptList2edl(d, v, "security_exempt_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-exempt-list"] = t
		}
	}

	if v, ok := d.GetOk("security_obsolete_option"); ok || d.HasChange("security_obsolete_option") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSecurityObsoleteOption2edl(d, v, "security_obsolete_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-obsolete-option"] = t
		}
	}

	if v, ok := d.GetOk("security_redirect_url"); ok || d.HasChange("security_redirect_url") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSecurityRedirectUrl2edl(d, v, "security_redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("selected_usergroups"); ok || d.HasChange("selected_usergroups") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSelectedUsergroups2edl(d, v, "selected_usergroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selected-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling"); ok || d.HasChange("split_tunneling") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSplitTunneling2edl(d, v, "split_tunneling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("ssid"); ok || d.HasChange("ssid") {
		t, err := expandObjectWirelessControllerVapDynamicMappingSsid2edl(d, v, "ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_remove"); ok || d.HasChange("sticky_client_remove") {
		t, err := expandObjectWirelessControllerVapDynamicMappingStickyClientRemove2edl(d, v, "sticky_client_remove")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-remove"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_2g"); ok || d.HasChange("sticky_client_threshold_2g") {
		t, err := expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold2G2edl(d, v, "sticky_client_threshold_2g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-2g"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_5g"); ok || d.HasChange("sticky_client_threshold_5g") {
		t, err := expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold5G2edl(d, v, "sticky_client_threshold_5g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-5g"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_6g"); ok || d.HasChange("sticky_client_threshold_6g") {
		t, err := expandObjectWirelessControllerVapDynamicMappingStickyClientThreshold6G2edl(d, v, "sticky_client_threshold_6g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-6g"] = t
		}
	}

	if v, ok := d.GetOk("target_wake_time"); ok || d.HasChange("target_wake_time") {
		t, err := expandObjectWirelessControllerVapDynamicMappingTargetWakeTime2edl(d, v, "target_wake_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-wake-time"] = t
		}
	}

	if v, ok := d.GetOk("tkip_counter_measure"); ok || d.HasChange("tkip_counter_measure") {
		t, err := expandObjectWirelessControllerVapDynamicMappingTkipCounterMeasure2edl(d, v, "tkip_counter_measure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tkip-counter-measure"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_echo_interval"); ok || d.HasChange("tunnel_echo_interval") {
		t, err := expandObjectWirelessControllerVapDynamicMappingTunnelEchoInterval2edl(d, v, "tunnel_echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-echo-interval"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_fallback_interval"); ok || d.HasChange("tunnel_fallback_interval") {
		t, err := expandObjectWirelessControllerVapDynamicMappingTunnelFallbackInterval2edl(d, v, "tunnel_fallback_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-fallback-interval"] = t
		}
	}

	if v, ok := d.GetOk("usergroup"); ok || d.HasChange("usergroup") {
		t, err := expandObjectWirelessControllerVapDynamicMappingUsergroup2edl(d, v, "usergroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usergroup"] = t
		}
	}

	if v, ok := d.GetOk("utm_log"); ok || d.HasChange("utm_log") {
		t, err := expandObjectWirelessControllerVapDynamicMappingUtmLog2edl(d, v, "utm_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-log"] = t
		}
	}

	if v, ok := d.GetOk("utm_profile"); ok || d.HasChange("utm_profile") {
		t, err := expandObjectWirelessControllerVapDynamicMappingUtmProfile2edl(d, v, "utm_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-profile"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok || d.HasChange("utm_status") {
		t, err := expandObjectWirelessControllerVapDynamicMappingUtmStatus2edl(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandObjectWirelessControllerVapDynamicMappingVdom2edl(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("vlan_auto"); ok || d.HasChange("vlan_auto") {
		t, err := expandObjectWirelessControllerVapDynamicMappingVlanAuto2edl(d, v, "vlan_auto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-auto"] = t
		}
	}

	if v, ok := d.GetOk("vlan_pooling"); ok || d.HasChange("vlan_pooling") {
		t, err := expandObjectWirelessControllerVapDynamicMappingVlanPooling2edl(d, v, "vlan_pooling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-pooling"] = t
		}
	}

	if v, ok := d.GetOk("vlanid"); ok || d.HasChange("vlanid") {
		t, err := expandObjectWirelessControllerVapDynamicMappingVlanid2edl(d, v, "vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	if v, ok := d.GetOk("voice_enterprise"); ok || d.HasChange("voice_enterprise") {
		t, err := expandObjectWirelessControllerVapDynamicMappingVoiceEnterprise2edl(d, v, "voice_enterprise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice-enterprise"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandObjectWirelessControllerVapDynamicMappingWebfilterProfile2edl(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	return &obj, nil
}
