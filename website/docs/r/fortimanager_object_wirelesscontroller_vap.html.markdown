---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_vap"
description: |-
  Configure Virtual Access Points (VAPs).
---

# fortimanager_object_wirelesscontroller_vap
Configure Virtual Access Points (VAPs).

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`dynamic_mapping`: `fortimanager_object_wirelesscontroller_vap_dynamic_mapping`
`mac_filter_list`: `fortimanager_object_wirelesscontroller_vap_macfilterlist`
`mpsk_key`: `fortimanager_object_wirelesscontroller_vap_mpskkey`
`portal_message_overrides`: `fortimanager_object_wirelesscontroller_vap_portalmessageoverrides`
`vlan_name`: `fortimanager_object_wirelesscontroller_vap_vlanname`
`vlan_pool`: `fortimanager_object_wirelesscontroller_vap_vlanpool`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_vap" "trname" {
  _centmgmt                           = "enable"
  _intf_allowaccess                   = ["http", "https", "ping", "ssh"]
  _intf_device_identification         = "enable"
  _intf_device_netscan                = "disable"
  _intf_dhcp6_relay_service           = "disable"
  _intf_dhcp6_relay_type              = "regular"
  _intf_dhcp_relay_service            = "disable"
  _intf_dhcp_relay_type               = "regular"
  _intf_listen_forticlient_connection = "disable"
  atf_weight                          = 20
  broadcast_ssid                      = "enable"
  broadcast_suppression               = ["arp-known", "dhcp-ucast", "dhcp-up"]
  bss_color_partial                   = "enable"
  dhcp_option43_insertion             = "enable"
  dhcp_option82_circuit_id_insertion  = "disable"
  dhcp_option82_insertion             = "disable"
  dhcp_option82_remote_id_insertion   = "disable"
  eap_reauth                          = "disable"
  eap_reauth_intv                     = 86400
  eapol_key_retries                   = "enable"
  encrypt                             = "AES"
  external_fast_roaming               = "disable"
  fast_bss_transition                 = "disable"
  fast_roaming                        = "enable"
  ft_mobility_domain                  = 1000
  ft_over_ds                          = "enable"
  ft_r0_key_lifetime                  = 480
  gtk_rekey                           = "disable"
  gtk_rekey_intv                      = 86400
  high_efficiency                     = "enable"
  igmp_snooping                       = "disable"
  intra_vap_privacy                   = "disable"
  ipv6_rules                          = ["drop-dhcp6c", "drop-dhcp6s", "drop-icmp6mld2", "drop-icmp6ra", "drop-icmp6rs", "drop-llmnr6", "drop-ns-dad", "ndp-proxy"]
  ldpc                                = "rxtx"
  local_authentication                = "disable"
  local_bridging                      = "disable"
  local_lan                           = "allow"
  local_standalone                    = "disable"
  mac_auth_bypass                     = "disable"
  mac_filter                          = "disable"
  mac_filter_policy_other             = "allow"
  me_disable_thresh                   = 32
  mesh_backhaul                       = "disable"
  mpsk                                = "disable"
  mu_mimo                             = "enable"
  multicast_enhance                   = "disable"
  multicast_rate                      = "0"
  name                                = "terr-wictl-vap"
  okc                                 = "enable"
  owe_transition                      = "disable"
  passphrase                          = ["fortinet"]
  pmf                                 = "disable"
  pmf_assoc_comeback_timeout          = 1
  pmf_sa_query_retry_timeout          = 2
  port_macauth                        = "disable"
  port_macauth_reauth_timeout         = 7200
  port_macauth_timeout                = 600
  probe_resp_suppression              = "disable"
  probe_resp_threshold                = "-80"
  ptk_rekey                           = "disable"
  ptk_rekey_intv                      = 86400
  quarantine                          = "enable"
  radio_2g_threshold                  = "-79"
  radio_5g_threshold                  = "-76"
  radio_sensitivity                   = "disable"
  radius_mac_auth                     = "disable"
  security                            = "wpa2-only-personal"
  security_obsolete_option            = "disable"
  split_tunneling                     = "disable"
  ssid                                = "fortinet"
  sticky_client_remove                = "disable"
  sticky_client_threshold_2g          = "-79"
  sticky_client_threshold_5g          = "-76"
  target_wake_time                    = "enable"
  tkip_counter_measure                = "enable"
  vlan_auto                           = "disable"
  vlan_pooling                        = "disable"
  voice_enterprise                    = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_centmgmt` - _Centmgmt. Valid values: `disable`, `enable`.

* `_dhcp_svr_id` - _Dhcp_Svr_Id.
* `_intf_allowaccess` - _Intf_Allowaccess. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `auto-ipsec`, `radius-acct`, `probe-response`, `capwap`.

* `_intf_device_access_list` - _Intf_Device-Access-List.
* `_intf_device_identification` - _Intf_Device-Identification. Valid values: `disable`, `enable`.

* `_intf_device_netscan` - _Intf_Device-Netscan. Valid values: `disable`, `enable`.

* `_intf_dhcp_relay_ip` - _Intf_Dhcp-Relay-Ip.
* `_intf_dhcp_relay_service` - _Intf_Dhcp-Relay-Service. Valid values: `disable`, `enable`.

* `_intf_dhcp_relay_type` - _Intf_Dhcp-Relay-Type. Valid values: `regular`, `ipsec`.

* `_intf_dhcp6_relay_ip` - _Intf_Dhcp6-Relay-Ip.
* `_intf_dhcp6_relay_service` - _Intf_Dhcp6-Relay-Service. Valid values: `disable`, `enable`.

* `_intf_dhcp6_relay_type` - _Intf_Dhcp6-Relay-Type. Valid values: `regular`.

* `_intf_ip` - _Intf_Ip.
* `_intf_ip6_address` - _Intf_Ip6-Address.
* `_intf_ip6_allowaccess` - _Intf_Ip6-Allowaccess. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `any`, `fgfm`, `capwap`.

* `_intf_listen_forticlient_connection` - _Intf_Listen-Forticlient-Connection. Valid values: `disable`, `enable`.

* `_is_factory_setting` - _Is_Factory_Setting. Valid values: `disable`, `enable`, `ext`.

* `access_control_list` - access-control-list profile name.
* `additional_akms` - Additional AKMs. Valid values: `akm6`.

* `acct_interim_interval` - WiFi RADIUS accounting interim interval (60 - 86400 sec, default = 0).
* `address_group` - Address group ID.
* `address_group_policy` - Configure MAC address filtering policy for MAC addresses that are in the address-group. Valid values: `disable`, `allow`, `deny`.

* `alias` - Alias.
* `antivirus_profile` - AntiVirus profile name.
* `application_detection_engine` - Enable/disable application detection engine (default = disable). Valid values: `disable`, `enable`.

* `application_dscp_marking` - Enable/disable application attribute based DSCP marking (default = disable). Valid values: `disable`, `enable`.

* `application_list` - Application control list name.
* `application_report_intv` - Application report interval (30 - 864000 sec, default = 120).
* `atf_weight` - Airtime weight in percentage (default = 20).
* `auth` - Authentication protocol. Valid values: `PSK`, `psk`, `RADIUS`, `radius`, `usergroup`.

* `auth_cert` - HTTPS server certificate.
* `auth_portal_addr` - Address of captive portal.
* `beacon_advertising` - Fortinet beacon advertising IE data   (default = empty). Valid values: `name`, `model`, `serial-number`.

* `broadcast_ssid` - Enable/disable broadcasting the SSID (default = enable). Valid values: `disable`, `enable`.

* `broadcast_suppression` - Optional suppression of broadcast messages. For example, you can keep DHCP messages, ARP broadcasts, and so on off of the wireless network. Valid values: `dhcp`, `arp`, `dhcp2`, `arp2`, `netbios-ns`, `netbios-ds`, `arp3`, `dhcp-up`, `dhcp-down`, `arp-known`, `arp-unknown`, `arp-reply`, `ipv6`, `dhcp-starvation`, `arp-poison`, `all-other-mc`, `all-other-bc`, `arp-proxy`, `dhcp-ucast`.

* `bss_color_partial` - Enable/disable 802.11ax partial BSS color (default = enable). Valid values: `disable`, `enable`.

* `bstm_disassociation_imminent` - Enable/disable forcing of disassociation after the BSTM request timer has been reached (default = enable). Valid values: `disable`, `enable`.

* `bstm_load_balancing_disassoc_timer` - Time interval for client to voluntarily leave AP before forcing a disassociation due to AP load-balancing (0 to 30, default = 10).
* `bstm_rssi_disassoc_timer` - Time interval for client to voluntarily leave AP before forcing a disassociation due to low RSSI (0 to 2000, default = 200).
* `captive_portal_ac_name` - Local-bridging captive portal ac-name.
* `captive_portal_auth_timeout` - Hard timeout - AP will always clear the session after timeout regardless of traffic (0 - 864000 sec, default = 0).
* `captive_portal_fw_accounting` - Enable/disable RADIUS accounting for captive portal firewall authentication session. Valid values: `disable`, `enable`.

* `dhcp_address_enforcement` - Enable/disable DHCP address enforcement (default = disable). Valid values: `disable`, `enable`.

* `captive_portal_macauth_radius_secret` - Secret key to access the macauth RADIUS server.
* `captive_portal_macauth_radius_server` - Captive portal external RADIUS server domain name or IP address.
* `captive_portal_radius_secret` - Secret key to access the RADIUS server.
* `captive_portal_radius_server` - Captive portal RADIUS server domain name or IP address.
* `captive_portal_session_timeout_interval` - Session timeout interval (0 - 864000 sec, default = 0).
* `dhcp_lease_time` - DHCP lease time in seconds for NAT IP address.
* `dhcp_option43_insertion` - Enable/disable insertion of DHCP option 43 (default = enable). Valid values: `disable`, `enable`.

* `dhcp_option82_circuit_id_insertion` - Enable/disable DHCP option 82 circuit-id insert (default = disable). Valid values: `disable`, `style-1`, `style-2`, `style-3`.

* `dhcp_option82_insertion` - Enable/disable DHCP option 82 insert (default = disable). Valid values: `disable`, `enable`.

* `dhcp_option82_remote_id_insertion` - Enable/disable DHCP option 82 remote-id insert (default = disable). Valid values: `disable`, `style-1`.

* `dynamic_vlan` - Enable/disable dynamic VLAN assignment. Valid values: `disable`, `enable`.

* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `eap_reauth` - Enable/disable EAP re-authentication for WPA-Enterprise security. Valid values: `disable`, `enable`.

* `eap_reauth_intv` - EAP re-authentication interval (1800 - 864000 sec, default = 86400).
* `eapol_key_retries` - Enable/disable retransmission of EAPOL-Key frames (message 3/4 and group message 1/2) (default = enable). Valid values: `disable`, `enable`.

* `encrypt` - Encryption protocol to use (only available when security is set to a WPA type). Valid values: `TKIP`, `AES`, `TKIP-AES`.

* `external_fast_roaming` - Enable/disable fast roaming or pre-authentication with external APs not managed by the FortiGate (default = disable). Valid values: `disable`, `enable`.

* `external_logout` - URL of external authentication logout server.
* `external_web` - URL of external authentication web server.
* `external_web_format` - URL query parameter detection (default = auto-detect). Valid values: `auto-detect`, `no-query-string`, `partial-query-string`.

* `fast_bss_transition` - Enable/disable 802.11r Fast BSS Transition (FT) (default = disable). Valid values: `disable`, `enable`.

* `fast_roaming` - Enable/disable fast-roaming, or pre-authentication, where supported by clients (default = disable). Valid values: `disable`, `enable`.

* `ft_mobility_domain` - Mobility domain identifier in FT (1 - 65535, default = 1000).
* `ft_over_ds` - Enable/disable FT over the Distribution System (DS). Valid values: `disable`, `enable`.

* `ft_r0_key_lifetime` - Lifetime of the PMK-R0 key in FT, 1-65535 minutes.
* `gas_comeback_delay` - GAS comeback delay (0 or 100 - 10000 milliseconds, default = 500).
* `gas_fragmentation_limit` - GAS fragmentation limit (512 - 4096, default = 1024).
* `gtk_rekey` - Enable/disable GTK rekey for WPA security. Valid values: `disable`, `enable`.

* `gtk_rekey_intv` - GTK rekey interval (1800 - 864000 sec, default = 86400).
* `high_efficiency` - Enable/disable 802.11ax high efficiency (default = enable). Valid values: `disable`, `enable`.

* `hotspot20_profile` - Hotspot 2.0 profile name.
* `igmp_snooping` - Enable/disable IGMP snooping. Valid values: `disable`, `enable`.

* `intra_vap_privacy` - Enable/disable blocking communication between clients on the same SSID (called intra-SSID privacy) (default = disable). Valid values: `disable`, `enable`.

* `ip` - IP address and subnet mask for the local standalone NAT subnet.
* `ips_sensor` - IPS sensor name.
* `ipv6_rules` - Optional rules of IPv6 packets. For example, you can keep RA, RS and so on off of the wireless network. Valid values: `drop-icmp6ra`, `drop-icmp6rs`, `drop-llmnr6`, `drop-icmp6mld2`, `drop-dhcp6s`, `drop-dhcp6c`, `ndp-proxy`, `drop-ns-dad`, `drop-ns-nondad`.

* `key` - WEP Key.
* `keyindex` - WEP key index (1 - 4).
* `l3_roaming` - Enable/disable layer 3 roaming (default = disable). Valid values: `disable`, `enable`.

* `l3_roaming_mode` - Select the way that layer 3 roaming traffic is passed (default = direct). Valid values: `direct`, `indirect`.

* `ldpc` - VAP low-density parity-check (LDPC) coding configuration. Valid values: `disable`, `tx`, `rx`, `rxtx`.

* `local_authentication` - Enable/disable AP local authentication. Valid values: `disable`, `enable`.

* `local_bridging` - Enable/disable bridging of wireless and Ethernet interfaces on the FortiAP (default = disable). Valid values: `disable`, `enable`.

* `local_lan` - Allow/deny traffic destined for a Class A, B, or C private IP address (default = allow). Valid values: `deny`, `allow`.

* `local_standalone` - Enable/disable AP local standalone (default = disable). Valid values: `disable`, `enable`.

* `local_standalone_dns` - Enable/disable AP local standalone DNS. Valid values: `disable`, `enable`.

* `local_standalone_dns_ip` - IPv4 addresses for the local standalone DNS.
* `local_standalone_nat` - Enable/disable AP local standalone NAT mode. Valid values: `disable`, `enable`.

* `mac_auth_bypass` - Enable/disable MAC authentication bypass. Valid values: `disable`, `enable`.

* `mac_called_station_delimiter` - MAC called station delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_calling_station_delimiter` - MAC calling station delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_case` - MAC case (default = uppercase). Valid values: `uppercase`, `lowercase`.

* `mac_filter` - Enable/disable MAC filtering to block wireless clients by mac address. Valid values: `disable`, `enable`.

* `mac_filter_list` - Mac-Filter-List. The structure of `mac_filter_list` block is documented below.
* `mac_filter_policy_other` - Allow or block clients with MAC addresses that are not in the filter list. Valid values: `deny`, `allow`.

* `mac_password_delimiter` - MAC authentication password delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_username_delimiter` - MAC authentication username delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `max_clients` - Maximum number of clients that can connect simultaneously to the VAP (default = 0, meaning no limitation).
* `max_clients_ap` - Maximum number of clients that can connect simultaneously to the VAP per AP radio (default = 0, meaning no limitation).
* `mbo` - Enable/disable Multiband Operation (default = disable). Valid values: `disable`, `enable`.

* `mbo_cell_data_conn_pref` - MBO cell data connection preference (0, 1, or 255, default = 1). Valid values: `excluded`, `prefer-not`, `prefer-use`.

* `me_disable_thresh` - Disable multicast enhancement when this many clients are receiving multicast traffic.
* `mesh_backhaul` - Enable/disable using this VAP as a WiFi mesh backhaul (default = disable). This entry is only available when security is set to a WPA type or open. Valid values: `disable`, `enable`.

* `mpsk` - Enable/disable multiple PSK authentication. Valid values: `disable`, `enable`.

* `mpsk_concurrent_clients` - Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
* `mpsk_key` - Mpsk-Key. The structure of `mpsk_key` block is documented below.
* `mpsk_profile` - MPSK profile name.
* `mu_mimo` - Enable/disable Multi-user MIMO (default = enable). Valid values: `disable`, `enable`.

* `multicast_enhance` - Enable/disable converting multicast to unicast to improve performance (default = disable). Valid values: `disable`, `enable`.

* `multicast_rate` - Multicast rate (0, 6000, 12000, or 24000 kbps, default = 0). Valid values: `0`, `6000`, `12000`, `24000`.

* `nac` - Enable/disable network access control. Valid values: `disable`, `enable`.

* `nac_profile` - NAC profile name.
* `name` - Virtual AP name.
* `neighbor_report_dual_band` - Enable/disable dual-band neighbor report (default = disable). Valid values: `disable`, `enable`.

* `okc` - Enable/disable Opportunistic Key Caching (OKC) (default = enable). Valid values: `disable`, `enable`.

* `osen` - Enable/disable OSEN as part of key management (default = disable). Valid values: `disable`, `enable`.

* `owe_groups` - OWE-Groups. Valid values: `19`, `20`, `21`.

* `owe_transition` - Enable/disable OWE transition mode support. Valid values: `disable`, `enable`.

* `owe_transition_ssid` - OWE transition mode peer SSID.
* `passphrase` - WPA pre-shared key (PSK) to be used to authenticate WiFi users.
* `pmf` - Protected Management Frames (PMF) support (default = disable). Valid values: `disable`, `enable`, `optional`.

* `pmf_assoc_comeback_timeout` - Protected Management Frames (PMF) comeback maximum timeout (1-20 sec).
* `pmf_sa_query_retry_timeout` - Protected Management Frames (PMF) SA query retry timeout interval (1 - 5 100s of msec).
* `port_macauth` - Enable/disable LAN port MAC authentication (default = disable). Valid values: `disable`, `radius`, `address-group`.

* `port_macauth_reauth_timeout` - LAN port MAC authentication re-authentication timeout value (default = 7200 sec).
* `port_macauth_timeout` - LAN port MAC authentication idle timeout value (default = 600 sec).
* `portal_message_override_group` - Replacement message group for this VAP (only available when security is set to a captive portal type).
* `portal_message_overrides` - Portal-Message-Overrides. The structure of `portal_message_overrides` block is documented below.
* `portal_type` - Captive portal functionality. Configure how the captive portal authenticates users and whether it includes a disclaimer. Valid values: `auth`, `auth+disclaimer`, `disclaimer`, `email-collect`, `cmcc`, `cmcc-macauth`, `auth-mac`, `external-auth`.

* `primary_wag_profile` - Primary wireless access gateway profile name.
* `probe_resp_suppression` - Enable/disable probe response suppression (to ignore weak signals) (default = disable). Valid values: `disable`, `enable`.

* `probe_resp_threshold` - Minimum signal level/threshold in dBm required for the AP response to probe requests (-95 to -20, default = -80).
* `ptk_rekey` - Enable/disable PTK rekey for WPA-Enterprise security. Valid values: `disable`, `enable`.

* `ptk_rekey_intv` - PTK rekey interval (1800 - 864000 sec, default = 86400).
* `qos_profile` - Quality of service profile name.
* `quarantine` - Enable/disable station quarantine (default = enable). Valid values: `disable`, `enable`.

* `radio_2g_threshold` - Minimum signal level/threshold in dBm required for the AP response to receive a packet in 2.4G band (-95 to -20, default = -79).
* `radio_5g_threshold` - Minimum signal level/threshold in dBm required for the AP response to receive a packet in 5G band(-95 to -20, default = -76).
* `radio_sensitivity` - Enable/disable software radio sensitivity (to ignore weak signals) (default = disable). Valid values: `disable`, `enable`.

* `radius_mac_auth` - Enable/disable RADIUS-based MAC authentication of clients (default = disable). Valid values: `disable`, `enable`.

* `radius_mac_auth_block_interval` - Don't send RADIUS MAC auth request again if the client has been rejected within specific interval (0 or 30 - 864000 seconds, default = 0, 0 to disable blocking).
* `radius_mac_auth_server` - RADIUS-based MAC authentication server.
* `radius_mac_auth_usergroups` - Selective user groups that are permitted for RADIUS mac authentication.
* `radius_mac_mpsk_auth` - Enable/disable RADIUS-based MAC authentication of clients for MPSK authentication (default = disable). Valid values: `disable`, `enable`.

* `radius_mac_mpsk_timeout` - RADIUS MAC MPSK cache timeout interval (1800 - 864000, default = 86400).
* `radius_server` - RADIUS server to be used to authenticate WiFi users.
* `rates_11a` - Allowed data rates for 802.11a. Valid values: `1`, `1-basic`, `2`, `2-basic`, `5.5`, `5.5-basic`, `6`, `6-basic`, `9`, `9-basic`, `12`, `12-basic`, `18`, `18-basic`, `24`, `24-basic`, `36`, `36-basic`, `48`, `48-basic`, `54`, `54-basic`, `11`, `11-basic`.

* `rates_11ac_mcs_map` - Comma separated list of max supported VHT MCS for spatial streams 1 through 8.
* `rates_11ac_ss12` - Allowed data rates for 802.11ac/ax with 1 or 2 spatial streams. Valid values: `mcs0/1`, `mcs1/1`, `mcs2/1`, `mcs3/1`, `mcs4/1`, `mcs5/1`, `mcs6/1`, `mcs7/1`, `mcs8/1`, `mcs9/1`, `mcs0/2`, `mcs1/2`, `mcs2/2`, `mcs3/2`, `mcs4/2`, `mcs5/2`, `mcs6/2`, `mcs7/2`, `mcs8/2`, `mcs9/2`, `mcs10/1`, `mcs11/1`, `mcs10/2`, `mcs11/2`.

* `rates_11ac_ss34` - Allowed data rates for 802.11ac/ax with 3 or 4 spatial streams. Valid values: `mcs0/3`, `mcs1/3`, `mcs2/3`, `mcs3/3`, `mcs4/3`, `mcs5/3`, `mcs6/3`, `mcs7/3`, `mcs8/3`, `mcs9/3`, `mcs0/4`, `mcs1/4`, `mcs2/4`, `mcs3/4`, `mcs4/4`, `mcs5/4`, `mcs6/4`, `mcs7/4`, `mcs8/4`, `mcs9/4`, `mcs10/3`, `mcs11/3`, `mcs10/4`, `mcs11/4`.

* `rates_11ax_mcs_map` - Comma separated list of max supported HE MCS for spatial streams 1 through 8.
* `rates_11ax_ss12` - Allowed data rates for 802.11ax with 1 or 2 spatial streams. Valid values: `mcs0/1`, `mcs1/1`, `mcs2/1`, `mcs3/1`, `mcs4/1`, `mcs5/1`, `mcs6/1`, `mcs7/1`, `mcs8/1`, `mcs9/1`, `mcs10/1`, `mcs11/1`, `mcs0/2`, `mcs1/2`, `mcs2/2`, `mcs3/2`, `mcs4/2`, `mcs5/2`, `mcs6/2`, `mcs7/2`, `mcs8/2`, `mcs9/2`, `mcs10/2`, `mcs11/2`.

* `rates_11ax_ss34` - Allowed data rates for 802.11ax with 3 or 4 spatial streams. Valid values: `mcs0/3`, `mcs1/3`, `mcs2/3`, `mcs3/3`, `mcs4/3`, `mcs5/3`, `mcs6/3`, `mcs7/3`, `mcs8/3`, `mcs9/3`, `mcs10/3`, `mcs11/3`, `mcs0/4`, `mcs1/4`, `mcs2/4`, `mcs3/4`, `mcs4/4`, `mcs5/4`, `mcs6/4`, `mcs7/4`, `mcs8/4`, `mcs9/4`, `mcs10/4`, `mcs11/4`.

* `rates_11bg` - Allowed data rates for 802.11b/g. Valid values: `1`, `1-basic`, `2`, `2-basic`, `5.5`, `5.5-basic`, `6`, `6-basic`, `9`, `9-basic`, `12`, `12-basic`, `18`, `18-basic`, `24`, `24-basic`, `36`, `36-basic`, `48`, `48-basic`, `54`, `54-basic`, `11`, `11-basic`.

* `rates_11n_ss12` - Allowed data rates for 802.11n with 1 or 2 spatial streams. Valid values: `mcs0/1`, `mcs1/1`, `mcs2/1`, `mcs3/1`, `mcs4/1`, `mcs5/1`, `mcs6/1`, `mcs7/1`, `mcs8/2`, `mcs9/2`, `mcs10/2`, `mcs11/2`, `mcs12/2`, `mcs13/2`, `mcs14/2`, `mcs15/2`.

* `rates_11n_ss34` - Allowed data rates for 802.11n with 3 or 4 spatial streams. Valid values: `mcs16/3`, `mcs17/3`, `mcs18/3`, `mcs19/3`, `mcs20/3`, `mcs21/3`, `mcs22/3`, `mcs23/3`, `mcs24/4`, `mcs25/4`, `mcs26/4`, `mcs27/4`, `mcs28/4`, `mcs29/4`, `mcs30/4`, `mcs31/4`.

* `sae_groups` - SAE-Groups. Valid values: `1`, `2`, `5`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `27`, `28`, `29`, `30`, `31`.

* `sae_h2e_only` - Use hash-to-element-only mechanism for PWE derivation (default = disable). Valid values: `disable`, `enable`.

* `sae_password` - WPA3 SAE password to be used to authenticate WiFi users.
* `sae_pk` - Enable/disable WPA3 SAE-PK (default = disable). Valid values: `disable`, `enable`.

* `sae_private_key` - Private key used for WPA3 SAE-PK authentication.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.

* `schedule` - Firewall schedules for enabling this VAP on the FortiAP. This VAP will be enabled when at least one of the schedules is valid. Separate multiple schedule names with a space.
* `secondary_wag_profile` - Secondary wireless access gateway profile name.
* `security` - Security mode for the wireless interface (default = wpa2-only-personal). Valid values: `None`, `WEP64`, `wep64`, `WEP128`, `wep128`, `WPA_PSK`, `WPA_RADIUS`, `WPA`, `WPA2`, `WPA2_AUTO`, `open`, `wpa-personal`, `wpa-enterprise`, `captive-portal`, `wpa-only-personal`, `wpa-only-enterprise`, `wpa2-only-personal`, `wpa2-only-enterprise`, `wpa-personal+captive-portal`, `wpa-only-personal+captive-portal`, `wpa2-only-personal+captive-portal`, `osen`, `wpa3-enterprise`, `sae`, `sae-transition`, `owe`, `wpa3-sae`, `wpa3-sae-transition`.

* `security_exempt_list` - Optional security exempt list for captive portal authentication.
* `security_obsolete_option` - Enable/disable obsolete security options. Valid values: `disable`, `enable`.

* `security_redirect_url` - Optional URL for redirecting users after they pass captive portal authentication.
* `selected_usergroups` - Selective user groups that are permitted to authenticate.
* `split_tunneling` - Enable/disable split tunneling (default = disable). Valid values: `disable`, `enable`.

* `ssid` - IEEE 802.11 service set identifier (SSID) for the wireless interface. Users who wish to use the wireless network must configure their computers to access this SSID name.
* `sticky_client_remove` - Enable/disable sticky client remove to maintain good signal level clients in SSID. (default = disable). Valid values: `disable`, `enable`.

* `sticky_client_threshold_2g` - Minimum signal level/threshold in dBm required for the 2G client to be serviced by the AP (-95 to -20, default = -79).
* `sticky_client_threshold_5g` - Minimum signal level/threshold in dBm required for the 5G client to be serviced by the AP (-95 to -20, default = -76).
* `sticky_client_threshold_6g` - Minimum signal level/threshold in dBm required for the 6G client to be serviced by the AP (-95 to -20, default = -76).
* `target_wake_time` - Enable/disable 802.11ax target wake time (default = enable). Valid values: `disable`, `enable`.

* `tkip_counter_measure` - Enable/disable TKIP counter measure. Valid values: `disable`, `enable`.

* `tunnel_echo_interval` - The time interval to send echo to both primary and secondary tunnel peers (1 - 65535 sec, default = 300).
* `tunnel_fallback_interval` - The time interval for secondary tunnel to fall back to primary tunnel (0 - 65535 sec, default = 7200).
* `usergroup` - Firewall user group to be used to authenticate WiFi users.
* `utm_log` - Enable/disable UTM logging. Valid values: `disable`, `enable`.

* `utm_profile` - UTM profile name.
* `vdom` - Name of the VDOM that the Virtual AP has been added to.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the VAP. Valid values: `disable`, `enable`.

* `vlan_auto` - Enable/disable automatic management of SSID VLAN interface. Valid values: `disable`, `enable`.

* `vlan_name` - Vlan-Name. The structure of `vlan_name` block is documented below.
* `vlan_pool` - Vlan-Pool. The structure of `vlan_pool` block is documented below.
* `vlan_pooling` - Enable/disable VLAN pooling, to allow grouping of multiple wireless controller VLANs into VLAN pools (default = disable). When set to wtp-group, VLAN pooling occurs with VLAN assignment by wtp-group. Valid values: `wtp-group`, `round-robin`, `hash`, `disable`.

* `vlanid` - Optional VLAN ID.
* `voice_enterprise` - Enable/disable 802.11k and 802.11v assisted Voice-Enterprise roaming (default = disable). Valid values: `disable`, `enable`.

* `webfilter_profile` - WebFilter profile name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_centmgmt` - _Centmgmt. Valid values: `disable`, `enable`.

* `_dhcp_svr_id` - _Dhcp_Svr_Id.
* `_intf_allowaccess` - _Intf_Allowaccess. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `auto-ipsec`, `radius-acct`, `probe-response`, `capwap`.

* `_intf_device_access_list` - _Intf_Device-Access-List.
* `_intf_device_identification` - _Intf_Device-Identification. Valid values: `disable`, `enable`.

* `_intf_device_netscan` - _Intf_Device-Netscan. Valid values: `disable`, `enable`.

* `_intf_dhcp_relay_ip` - _Intf_Dhcp-Relay-Ip.
* `_intf_dhcp_relay_service` - _Intf_Dhcp-Relay-Service. Valid values: `disable`, `enable`.

* `_intf_dhcp_relay_type` - _Intf_Dhcp-Relay-Type. Valid values: `regular`, `ipsec`.

* `_intf_dhcp6_relay_ip` - _Intf_Dhcp6-Relay-Ip.
* `_intf_dhcp6_relay_service` - _Intf_Dhcp6-Relay-Service. Valid values: `disable`, `enable`.

* `_intf_dhcp6_relay_type` - _Intf_Dhcp6-Relay-Type. Valid values: `regular`.

* `_intf_ip` - _Intf_Ip.
* `_intf_ip6_address` - _Intf_Ip6-Address.
* `_intf_ip6_allowaccess` - _Intf_Ip6-Allowaccess. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `any`, `fgfm`, `capwap`.

* `_intf_listen_forticlient_connection` - _Intf_Listen-Forticlient-Connection. Valid values: `disable`, `enable`.

* `_is_factory_setting` - _Is_Factory_Setting. Valid values: `disable`, `enable`, `ext`.

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `access_control_list` - Access-Control-List.
* `acct_interim_interval` - WiFi RADIUS accounting interim interval (60 - 86400 sec, default = 0).
* `additional_akms` - Additional-Akms. Valid values: `akm6`.

* `address_group` - Address group ID.
* `address_group_policy` - Address-Group-Policy. Valid values: `disable`, `allow`, `deny`.

* `alias` - Alias.
* `antivirus_profile` - AntiVirus profile name.
* `application_detection_engine` - Application-Detection-Engine. Valid values: `disable`, `enable`.

* `application_dscp_marking` - Enable/disable application attribute based DSCP marking (default = disable). Valid values: `disable`, `enable`.

* `application_list` - Application control list name.
* `application_report_intv` - Application-Report-Intv.
* `atf_weight` - Airtime weight in percentage (default = 20).
* `auth` - Authentication protocol. Valid values: `PSK`, `psk`, `RADIUS`, `radius`, `usergroup`.

* `auth_cert` - HTTPS server certificate.
* `auth_portal_addr` - Address of captive portal.
* `beacon_advertising` - Fortinet beacon advertising IE data   (default = empty). Valid values: `name`, `model`, `serial-number`.

* `broadcast_ssid` - Enable/disable broadcasting the SSID (default = enable). Valid values: `disable`, `enable`.

* `broadcast_suppression` - Optional suppression of broadcast messages. For example, you can keep DHCP messages, ARP broadcasts, and so on off of the wireless network. Valid values: `dhcp`, `arp`, `dhcp2`, `arp2`, `netbios-ns`, `netbios-ds`, `arp3`, `dhcp-up`, `dhcp-down`, `arp-known`, `arp-unknown`, `arp-reply`, `ipv6`, `dhcp-starvation`, `arp-poison`, `all-other-mc`, `all-other-bc`, `arp-proxy`, `dhcp-ucast`.

* `bss_color_partial` - Bss-Color-Partial. Valid values: `disable`, `enable`.

* `bstm_disassociation_imminent` - Bstm-Disassociation-Imminent. Valid values: `disable`, `enable`.

* `bstm_load_balancing_disassoc_timer` - Bstm-Load-Balancing-Disassoc-Timer.
* `bstm_rssi_disassoc_timer` - Bstm-Rssi-Disassoc-Timer.
* `captive_portal_ac_name` - Local-bridging captive portal ac-name.
* `captive_portal_auth_timeout` - Captive-Portal-Auth-Timeout.
* `captive_portal_fw_accounting` - Enable/disable RADIUS accounting for captive portal firewall authentication session. Valid values: `disable`, `enable`.

* `captive_portal_macauth_radius_secret` - Secret key to access the macauth RADIUS server.
* `captive_portal_macauth_radius_server` - Captive portal external RADIUS server domain name or IP address.
* `captive_portal_radius_secret` - Secret key to access the RADIUS server.
* `captive_portal_radius_server` - Captive portal RADIUS server domain name or IP address.
* `captive_portal_session_timeout_interval` - Session timeout interval (0 - 864000 sec, default = 0).
* `client_count` - Client-Count.
* `dhcp_address_enforcement` - Dhcp-Address-Enforcement. Valid values: `disable`, `enable`.

* `dhcp_lease_time` - DHCP lease time in seconds for NAT IP address.
* `dhcp_option43_insertion` - Dhcp-Option43-Insertion. Valid values: `disable`, `enable`.

* `dhcp_option82_circuit_id_insertion` - Enable/disable DHCP option 82 circuit-id insert (default = disable). Valid values: `disable`, `style-1`, `style-2`, `style-3`.

* `dhcp_option82_insertion` - Enable/disable DHCP option 82 insert (default = disable). Valid values: `disable`, `enable`.

* `dhcp_option82_remote_id_insertion` - Enable/disable DHCP option 82 remote-id insert (default = disable). Valid values: `disable`, `style-1`.

* `dynamic_vlan` - Enable/disable dynamic VLAN assignment. Valid values: `disable`, `enable`.

* `eap_reauth` - Enable/disable EAP re-authentication for WPA-Enterprise security. Valid values: `disable`, `enable`.

* `eap_reauth_intv` - EAP re-authentication interval (1800 - 864000 sec, default = 86400).
* `eapol_key_retries` - Enable/disable retransmission of EAPOL-Key frames (message 3/4 and group message 1/2) (default = enable). Valid values: `disable`, `enable`.

* `encrypt` - Encryption protocol to use (only available when security is set to a WPA type). Valid values: `TKIP`, `AES`, `TKIP-AES`.

* `external_fast_roaming` - Enable/disable fast roaming or pre-authentication with external APs not managed by the FortiGate (default = disable). Valid values: `disable`, `enable`.

* `external_logout` - URL of external authentication logout server.
* `external_web` - URL of external authentication web server.
* `external_web_format` - URL query parameter detection (default = auto-detect). Valid values: `auto-detect`, `no-query-string`, `partial-query-string`.

* `fast_bss_transition` - Enable/disable 802.11r Fast BSS Transition (FT) (default = disable). Valid values: `disable`, `enable`.

* `fast_roaming` - Enable/disable fast-roaming, or pre-authentication, where supported by clients (default = disable). Valid values: `disable`, `enable`.

* `ft_mobility_domain` - Mobility domain identifier in FT (1 - 65535, default = 1000).
* `ft_over_ds` - Enable/disable FT over the Distribution System (DS). Valid values: `disable`, `enable`.

* `ft_r0_key_lifetime` - Lifetime of the PMK-R0 key in FT, 1-65535 minutes.
* `gas_comeback_delay` - Gas-Comeback-Delay.
* `gas_fragmentation_limit` - Gas-Fragmentation-Limit.
* `gtk_rekey` - Enable/disable GTK rekey for WPA security. Valid values: `disable`, `enable`.

* `gtk_rekey_intv` - GTK rekey interval (1800 - 864000 sec, default = 86400).
* `high_efficiency` - Enable/disable 802.11ax high efficiency (default = enable). Valid values: `disable`, `enable`.

* `hotspot20_profile` - Hotspot 2.0 profile name.
* `igmp_snooping` - Igmp-Snooping. Valid values: `disable`, `enable`.

* `intra_vap_privacy` - Enable/disable blocking communication between clients on the same SSID (called intra-SSID privacy) (default = disable). Valid values: `disable`, `enable`.

* `ip` - IP address and subnet mask for the local standalone NAT subnet.
* `ips_sensor` - IPS sensor name.
* `ipv6_rules` - Ipv6-Rules. Valid values: `drop-icmp6ra`, `drop-icmp6rs`, `drop-llmnr6`, `drop-icmp6mld2`, `drop-dhcp6s`, `drop-dhcp6c`, `ndp-proxy`, `drop-ns-dad`, `drop-ns-nondad`.

* `key` - WEP Key.
* `keyindex` - WEP key index (1 - 4).
* `l3_roaming` - L3-Roaming. Valid values: `disable`, `enable`.

* `l3_roaming_mode` - Select the way that layer 3 roaming traffic is passed (default = direct). Valid values: `direct`, `indirect`.

* `ldpc` - VAP low-density parity-check (LDPC) coding configuration. Valid values: `disable`, `tx`, `rx`, `rxtx`.

* `local_authentication` - Enable/disable AP local authentication. Valid values: `disable`, `enable`.

* `local_bridging` - Enable/disable bridging of wireless and Ethernet interfaces on the FortiAP (default = disable). Valid values: `disable`, `enable`.

* `local_lan` - Allow/deny traffic destined for a Class A, B, or C private IP address (default = allow). Valid values: `deny`, `allow`.

* `local_standalone` - Enable/disable AP local standalone (default = disable). Valid values: `disable`, `enable`.

* `local_standalone_dns` - Enable/disable AP local standalone DNS. Valid values: `disable`, `enable`.

* `local_standalone_dns_ip` - IPv4 addresses for the local standalone DNS.
* `local_standalone_nat` - Enable/disable AP local standalone NAT mode. Valid values: `disable`, `enable`.

* `local_switching` - Local-Switching. Valid values: `disable`, `enable`.

* `mac_auth_bypass` - Enable/disable MAC authentication bypass. Valid values: `disable`, `enable`.

* `mac_called_station_delimiter` - Mac-Called-Station-Delimiter. Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_calling_station_delimiter` - Mac-Calling-Station-Delimiter. Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_case` - Mac-Case. Valid values: `uppercase`, `lowercase`.

* `mac_filter` - Enable/disable MAC filtering to block wireless clients by mac address. Valid values: `disable`, `enable`.

* `mac_filter_policy_other` - Allow or block clients with MAC addresses that are not in the filter list. Valid values: `deny`, `allow`.

* `mac_password_delimiter` - Mac-Password-Delimiter. Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_username_delimiter` - Mac-Username-Delimiter. Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `max_clients` - Maximum number of clients that can connect simultaneously to the VAP (default = 0, meaning no limitation).
* `max_clients_ap` - Maximum number of clients that can connect simultaneously to the VAP per AP radio (default = 0, meaning no limitation).
* `mbo` - Mbo. Valid values: `disable`, `enable`.

* `mbo_cell_data_conn_pref` - Mbo-Cell-Data-Conn-Pref. Valid values: `excluded`, `prefer-not`, `prefer-use`.

* `me_disable_thresh` - Disable multicast enhancement when this many clients are receiving multicast traffic.
* `mesh_backhaul` - Enable/disable using this VAP as a WiFi mesh backhaul (default = disable). This entry is only available when security is set to a WPA type or open. Valid values: `disable`, `enable`.

* `mpsk` - Enable/disable multiple PSK authentication. Valid values: `disable`, `enable`.

* `mpsk_concurrent_clients` - Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
* `mpsk_profile` - Mpsk-Profile.
* `mu_mimo` - Enable/disable Multi-user MIMO (default = enable). Valid values: `disable`, `enable`.

* `multicast_enhance` - Enable/disable converting multicast to unicast to improve performance (default = disable). Valid values: `disable`, `enable`.

* `multicast_rate` - Multicast rate (0, 6000, 12000, or 24000 kbps, default = 0). Valid values: `0`, `6000`, `12000`, `24000`.

* `nac` - Nac. Valid values: `disable`, `enable`.

* `nac_profile` - Nac-Profile.
* `neighbor_report_dual_band` - Neighbor-Report-Dual-Band. Valid values: `disable`, `enable`.

* `okc` - Enable/disable Opportunistic Key Caching (OKC) (default = enable). Valid values: `disable`, `enable`.

* `osen` - Enable/disable OSEN as part of key management (default = disable). Valid values: `disable`, `enable`.

* `owe_groups` - OWE-Groups. Valid values: `19`, `20`, `21`.

* `owe_transition` - Enable/disable OWE transition mode support. Valid values: `disable`, `enable`.

* `owe_transition_ssid` - OWE transition mode peer SSID.
* `passphrase` - WPA pre-shared key (PSK) to be used to authenticate WiFi users.
* `pmf` - Protected Management Frames (PMF) support (default = disable). Valid values: `disable`, `enable`, `optional`.

* `pmf_assoc_comeback_timeout` - Protected Management Frames (PMF) comeback maximum timeout (1-20 sec).
* `pmf_sa_query_retry_timeout` - Protected Management Frames (PMF) SA query retry timeout interval (1 - 5 100s of msec).
* `port_macauth` - Port-Macauth. Valid values: `disable`, `radius`, `address-group`.

* `port_macauth_reauth_timeout` - Port-Macauth-Reauth-Timeout.
* `port_macauth_timeout` - Port-Macauth-Timeout.
* `portal_message_override_group` - Replacement message group for this VAP (only available when security is set to a captive portal type).
* `portal_type` - Captive portal functionality. Configure how the captive portal authenticates users and whether it includes a disclaimer. Valid values: `auth`, `auth+disclaimer`, `disclaimer`, `email-collect`, `cmcc`, `cmcc-macauth`, `auth-mac`, `external-auth`.

* `primary_wag_profile` - Primary wireless access gateway profile name.
* `probe_resp_suppression` - Enable/disable probe response suppression (to ignore weak signals) (default = disable). Valid values: `disable`, `enable`.

* `probe_resp_threshold` - Minimum signal level/threshold in dBm required for the AP response to probe requests (-95 to -20, default = -80).
* `ptk_rekey` - Enable/disable PTK rekey for WPA-Enterprise security. Valid values: `disable`, `enable`.

* `ptk_rekey_intv` - PTK rekey interval (1800 - 864000 sec, default = 86400).
* `qos_profile` - Quality of service profile name.
* `quarantine` - Enable/disable station quarantine (default = enable). Valid values: `disable`, `enable`.

* `radio_2g_threshold` - Minimum signal level/threshold in dBm required for the AP response to receive a packet in 2.4G band (-95 to -20, default = -79).
* `radio_5g_threshold` - Minimum signal level/threshold in dBm required for the AP response to receive a packet in 5G band(-95 to -20, default = -76).
* `radio_sensitivity` - Enable/disable software radio sensitivity (to ignore weak signals) (default = disable). Valid values: `disable`, `enable`.

* `radius_mac_auth` - Enable/disable RADIUS-based MAC authentication of clients (default = disable). Valid values: `disable`, `enable`.

* `radius_mac_auth_block_interval` - Don't send RADIUS MAC auth request again if the client has been rejected within specific interval (0 or 30 - 864000 seconds, default = 0, 0 to disable blocking).
* `radius_mac_auth_server` - RADIUS-based MAC authentication server.
* `radius_mac_auth_usergroups` - Selective user groups that are permitted for RADIUS mac authentication.
* `radius_mac_mpsk_auth` - Enable/disable RADIUS-based MAC authentication of clients for MPSK authentication (default = disable). Valid values: `disable`, `enable`.

* `radius_mac_mpsk_timeout` - RADIUS MAC MPSK cache timeout interval (1800 - 864000, default = 86400).
* `radius_server` - RADIUS server to be used to authenticate WiFi users.
* `rates_11a` - Allowed data rates for 802.11a. Valid values: `1`, `1-basic`, `2`, `2-basic`, `5.5`, `5.5-basic`, `6`, `6-basic`, `9`, `9-basic`, `12`, `12-basic`, `18`, `18-basic`, `24`, `24-basic`, `36`, `36-basic`, `48`, `48-basic`, `54`, `54-basic`, `11`, `11-basic`.

* `rates_11ac_mcs_map` - Comma separated list of max supported VHT MCS for spatial streams 1 through 8.
* `rates_11ac_ss12` - Allowed data rates for 802.11ac/ax with 1 or 2 spatial streams. Valid values: `mcs0/1`, `mcs1/1`, `mcs2/1`, `mcs3/1`, `mcs4/1`, `mcs5/1`, `mcs6/1`, `mcs7/1`, `mcs8/1`, `mcs9/1`, `mcs0/2`, `mcs1/2`, `mcs2/2`, `mcs3/2`, `mcs4/2`, `mcs5/2`, `mcs6/2`, `mcs7/2`, `mcs8/2`, `mcs9/2`, `mcs10/1`, `mcs11/1`, `mcs10/2`, `mcs11/2`.

* `rates_11ac_ss34` - Allowed data rates for 802.11ac/ax with 3 or 4 spatial streams. Valid values: `mcs0/3`, `mcs1/3`, `mcs2/3`, `mcs3/3`, `mcs4/3`, `mcs5/3`, `mcs6/3`, `mcs7/3`, `mcs8/3`, `mcs9/3`, `mcs0/4`, `mcs1/4`, `mcs2/4`, `mcs3/4`, `mcs4/4`, `mcs5/4`, `mcs6/4`, `mcs7/4`, `mcs8/4`, `mcs9/4`, `mcs10/3`, `mcs11/3`, `mcs10/4`, `mcs11/4`.

* `rates_11ax_mcs_map` - Comma separated list of max supported HE MCS for spatial streams 1 through 8.
* `rates_11ax_ss12` - Allowed data rates for 802.11ax with 1 or 2 spatial streams. Valid values: `mcs0/1`, `mcs1/1`, `mcs2/1`, `mcs3/1`, `mcs4/1`, `mcs5/1`, `mcs6/1`, `mcs7/1`, `mcs8/1`, `mcs9/1`, `mcs10/1`, `mcs11/1`, `mcs0/2`, `mcs1/2`, `mcs2/2`, `mcs3/2`, `mcs4/2`, `mcs5/2`, `mcs6/2`, `mcs7/2`, `mcs8/2`, `mcs9/2`, `mcs10/2`, `mcs11/2`.

* `rates_11ax_ss34` - Allowed data rates for 802.11ax with 3 or 4 spatial streams. Valid values: `mcs0/3`, `mcs1/3`, `mcs2/3`, `mcs3/3`, `mcs4/3`, `mcs5/3`, `mcs6/3`, `mcs7/3`, `mcs8/3`, `mcs9/3`, `mcs10/3`, `mcs11/3`, `mcs0/4`, `mcs1/4`, `mcs2/4`, `mcs3/4`, `mcs4/4`, `mcs5/4`, `mcs6/4`, `mcs7/4`, `mcs8/4`, `mcs9/4`, `mcs10/4`, `mcs11/4`.

* `rates_11bg` - Allowed data rates for 802.11b/g. Valid values: `1`, `1-basic`, `2`, `2-basic`, `5.5`, `5.5-basic`, `6`, `6-basic`, `9`, `9-basic`, `12`, `12-basic`, `18`, `18-basic`, `24`, `24-basic`, `36`, `36-basic`, `48`, `48-basic`, `54`, `54-basic`, `11`, `11-basic`.

* `rates_11n_ss12` - Allowed data rates for 802.11n with 1 or 2 spatial streams. Valid values: `mcs0/1`, `mcs1/1`, `mcs2/1`, `mcs3/1`, `mcs4/1`, `mcs5/1`, `mcs6/1`, `mcs7/1`, `mcs8/2`, `mcs9/2`, `mcs10/2`, `mcs11/2`, `mcs12/2`, `mcs13/2`, `mcs14/2`, `mcs15/2`.

* `rates_11n_ss34` - Allowed data rates for 802.11n with 3 or 4 spatial streams. Valid values: `mcs16/3`, `mcs17/3`, `mcs18/3`, `mcs19/3`, `mcs20/3`, `mcs21/3`, `mcs22/3`, `mcs23/3`, `mcs24/4`, `mcs25/4`, `mcs26/4`, `mcs27/4`, `mcs28/4`, `mcs29/4`, `mcs30/4`, `mcs31/4`.

* `sae_groups` - SAE-Groups. Valid values: `1`, `2`, `5`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `27`, `28`, `29`, `30`, `31`.

* `sae_h2e_only` - Use hash-to-element-only mechanism for PWE derivation (default = disable). Valid values: `disable`, `enable`.

* `sae_password` - WPA3 SAE password to be used to authenticate WiFi users.
* `sae_pk` - Enable/disable WPA3 SAE-PK (default = disable). Valid values: `disable`, `enable`.

* `sae_private_key` - Private key used for WPA3 SAE-PK authentication.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.

* `schedule` - Firewall schedules for enabling this VAP on the FortiAP. This VAP will be enabled when at least one of the schedules is valid. Separate multiple schedule names with a space.
* `secondary_wag_profile` - Secondary wireless access gateway profile name.
* `security` - Security mode for the wireless interface (default = wpa2-only-personal). Valid values: `None`, `WEP64`, `wep64`, `WEP128`, `wep128`, `WPA_PSK`, `WPA_RADIUS`, `WPA`, `WPA2`, `WPA2_AUTO`, `open`, `wpa-personal`, `wpa-enterprise`, `captive-portal`, `wpa-only-personal`, `wpa-only-enterprise`, `wpa2-only-personal`, `wpa2-only-enterprise`, `wpa-personal+captive-portal`, `wpa-only-personal+captive-portal`, `wpa2-only-personal+captive-portal`, `osen`, `wpa3-enterprise`, `sae`, `sae-transition`, `owe`, `wpa3-sae`, `wpa3-sae-transition`.

* `security_exempt_list` - Optional security exempt list for captive portal authentication.
* `security_obsolete_option` - Enable/disable obsolete security options. Valid values: `disable`, `enable`.

* `security_redirect_url` - Optional URL for redirecting users after they pass captive portal authentication.
* `selected_usergroups` - Selective user groups that are permitted to authenticate.
* `split_tunneling` - Enable/disable split tunneling (default = disable). Valid values: `disable`, `enable`.

* `ssid` - IEEE 802.11 service set identifier (SSID) for the wireless interface. Users who wish to use the wireless network must configure their computers to access this SSID name.
* `sticky_client_remove` - Sticky-Client-Remove. Valid values: `disable`, `enable`.

* `sticky_client_threshold_2g` - Sticky-Client-Threshold-2G.
* `sticky_client_threshold_5g` - Sticky-Client-Threshold-5G.
* `sticky_client_threshold_6g` - Minimum signal level/threshold in dBm required for the 6G client to be serviced by the AP (-95 to -20, default = -76).
* `target_wake_time` - Enable/disable 802.11ax target wake time (default = enable). Valid values: `disable`, `enable`.

* `tkip_counter_measure` - Enable/disable TKIP counter measure. Valid values: `disable`, `enable`.

* `tunnel_echo_interval` - The time interval to send echo to both primary and secondary tunnel peers (1 - 65535 sec, default = 300).
* `tunnel_fallback_interval` - The time interval for secondary tunnel to fall back to primary tunnel (0 - 65535 sec, default = 7200).
* `usergroup` - Firewall user group to be used to authenticate WiFi users.
* `utm_log` - Enable/disable UTM logging. Valid values: `disable`, `enable`.

* `utm_profile` - UTM profile name.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the VAP. Valid values: `disable`, `enable`.

* `vdom` - Vdom.
* `vlan_auto` - Enable/disable automatic management of SSID VLAN interface. Valid values: `disable`, `enable`.

* `vlan_pooling` - Enable/disable VLAN pooling, to allow grouping of multiple wireless controller VLANs into VLAN pools (default = disable). When set to wtp-group, VLAN pooling occurs with VLAN assignment by wtp-group. Valid values: `wtp-group`, `round-robin`, `hash`, `disable`.

* `vlanid` - Optional VLAN ID.
* `voice_enterprise` - Enable/disable 802.11k and 802.11v assisted Voice-Enterprise roaming (default = disable). Valid values: `disable`, `enable`.

* `webfilter_profile` - WebFilter profile name.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `mac_filter_list` block supports:

* `id` - ID.
* `mac` - MAC address.
* `mac_filter_policy` - Deny or allow the client with this MAC address. Valid values: `deny`, `allow`.


The `mpsk_key` block supports:

* `comment` - Comment.
* `concurrent_clients` - Number of clients that can connect using this pre-shared key.
* `key_name` - Pre-shared key name.
* `mpsk_schedules` - Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid.
* `passphrase` - WPA Pre-shared key.

The `portal_message_overrides` block supports:

* `auth_disclaimer_page` - Override auth-disclaimer-page message with message from portal-message-overrides group.
* `auth_login_failed_page` - Override auth-login-failed-page message with message from portal-message-overrides group.
* `auth_login_page` - Override auth-login-page message with message from portal-message-overrides group.
* `auth_reject_page` - Override auth-reject-page message with message from portal-message-overrides group.

The `vlan_name` block supports:

* `name` - VLAN name.
* `vlan_id` - VLAN ID.

The `vlan_pool` block supports:

* `_wtp_group` - _Wtp-Group.
* `id` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Vap can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_vap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
