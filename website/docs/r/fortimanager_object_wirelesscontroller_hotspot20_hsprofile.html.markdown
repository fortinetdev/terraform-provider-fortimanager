---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_hsprofile"
description: |-
  Configure hotspot profile.
---

# fortimanager_object_wirelesscontroller_hotspot20_hsprofile
Configure hotspot profile.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_hsprofile" "trname" {
  access_network_asra     = "disable"
  access_network_esr      = "disable"
  access_network_internet = "disable"
  access_network_type     = "private-network"
  access_network_uesa     = "disable"
  bss_transition          = "disable"
  deauth_request_timeout  = 60
  dgaf                    = "disable"
  gas_comeback_delay      = 500
  gas_fragmentation_limit = 1024
  hessid                  = "00:00:00:00:00:00"
  l2tif                   = "disable"
  name                    = "terr-wictl-hot20-hs-profile"
  pame_bi                 = "enable"
  proxy_arp               = "enable"
  venue_group             = "unspecified"
  venue_type              = "unspecified"
  wnm_sleep_mode          = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `n3gpp_plmn` - 3GPP PLMN name.
* `access_network_asra` - Enable/disable additional step required for access (ASRA). Valid values: `disable`, `enable`.

* `access_network_esr` - Enable/disable emergency services reachable (ESR). Valid values: `disable`, `enable`.

* `access_network_internet` - Enable/disable connectivity to the Internet. Valid values: `disable`, `enable`.

* `access_network_type` - Access network type. Valid values: `private-network`, `private-network-with-guest-access`, `chargeable-public-network`, `free-public-network`, `personal-device-network`, `emergency-services-only-network`, `test-or-experimental`, `wildcard`.

* `access_network_uesa` - Enable/disable unauthenticated emergency service accessible (UESA). Valid values: `disable`, `enable`.

* `anqp_domain_id` - ANQP Domain ID (0-65535).
* `bss_transition` - Enable/disable basic service set (BSS) transition Support. Valid values: `disable`, `enable`.

* `conn_cap` - Connection capability name.
* `deauth_request_timeout` - Deauthentication request timeout (in seconds).
* `dgaf` - Enable/disable downstream group-addressed forwarding (DGAF). Valid values: `disable`, `enable`.

* `domain_name` - Domain name.
* `gas_comeback_delay` - GAS comeback delay (0 or 100 - 4000 milliseconds, default = 500).
* `gas_fragmentation_limit` - GAS fragmentation limit (512 - 4096, default = 1024).
* `hessid` - Homogeneous extended service set identifier (HESSID).
* `ip_addr_type` - IP address type name.
* `l2tif` - Enable/disable Layer 2 traffic inspection and filtering. Valid values: `disable`, `enable`.

* `nai_realm` - NAI realm list name.
* `name` - Hotspot profile name.
* `network_auth` - Network authentication name.
* `oper_friendly_name` - Operator friendly name.
* `osu_provider` - Manually selected list of OSU provider(s).
* `osu_ssid` - Online sign up (OSU) SSID.
* `pame_bi` - Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI). Valid values: `disable`, `enable`.

* `proxy_arp` - Enable/disable Proxy ARP. Valid values: `disable`, `enable`.

* `qos_map` - QoS MAP set ID.
* `roaming_consortium` - Roaming consortium list name.
* `venue_group` - Venue group. Valid values: `unspecified`, `assembly`, `business`, `educational`, `factory`, `institutional`, `mercantile`, `residential`, `storage`, `utility`, `vehicular`, `outdoor`.

* `venue_name` - Venue name.
* `venue_type` - Venue type. Valid values: `unspecified`, `arena`, `stadium`, `passenger-terminal`, `amphitheater`, `amusement-park`, `place-of-worship`, `convention-center`, `library`, `museum`, `restaurant`, `theater`, `bar`, `coffee-shop`, `zoo-or-aquarium`, `emergency-center`, `doctor-office`, `bank`, `fire-station`, `police-station`, `post-office`, `professional-office`, `research-facility`, `attorney-office`, `primary-school`, `secondary-school`, `university-or-college`, `factory`, `hospital`, `long-term-care-facility`, `rehab-center`, `group-home`, `prison-or-jail`, `retail-store`, `grocery-market`, `auto-service-station`, `shopping-mall`, `gas-station`, `private`, `hotel-or-motel`, `dormitory`, `boarding-house`, `automobile`, `airplane`, `bus`, `ferry`, `ship-or-boat`, `train`, `motor-bike`, `muni-mesh-network`, `city-park`, `rest-area`, `traffic-control`, `bus-stop`, `kiosk`.

* `wan_metrics` - WAN metric name.
* `wnm_sleep_mode` - Enable/disable wireless network management (WNM) sleep mode. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20HsProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_hsprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
