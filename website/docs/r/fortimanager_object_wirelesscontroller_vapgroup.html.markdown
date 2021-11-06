---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_vapgroup"
description: |-
  Configure virtual Access Point (VAP) groups.
---

# fortimanager_object_wirelesscontroller_vapgroup
Configure virtual Access Point (VAP) groups.

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

resource "fortimanager_object_wirelesscontroller_vapgroup" "trname" {
  comment = "This is a Terraform example"
  name    = "terr-wictl-vap-group"
  vaps    = ["terr-wictl-vap"]
  depends_on = [
    fortimanager_object_wirelesscontroller_vap.trname
  ]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `name` - Group Name
* `vaps` - List of SSIDs to be included in the VAP group.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController VapGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_vapgroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
