---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_system_ntp"
description: |-
  Configure system NTP information.
---

# fortimanager_systemp_system_ntp
Configure system NTP information.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ntpserver`: `fortimanager_systemp_system_ntp_ntpserver`



## Example Usage

```hcl
resource "fortimanager_systemp_system_ntp" "trname" {
  devprof        = "default"
  authentication = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `authentication` - Enable/disable authentication. Valid values: `disable`, `enable`.

* `interface` - <i>Support meta variable</i> FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services.
* `key` - Key for authentication.
* `key_id` - Key ID for authentication.
* `key_type` - Key type for authentication (MD5, SHA1). Valid values: `MD5`, `SHA1`.

* `ntpserver` - Ntpserver. The structure of `ntpserver` block is documented below.
* `ntpsync` - Enable/disable setting the FortiGate system time by synchronizing with an NTP Server. Valid values: `disable`, `enable`.

* `server_mode` - Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server. Valid values: `disable`, `enable`.

* `source_ip` - <i>Support meta variable</i> Source IP address for communication to the NTP server.
* `source_ip6` - Source IPv6 address for communication to the NTP server.
* `syncinterval` - NTP synchronization interval (1 - 1440 min).
* `type` - Use the FortiGuard NTP server or any other available NTP Server. Valid values: `fortiguard`, `custom`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ntpserver` block supports:

* `authentication` - Enable/disable MD5(NTPv3)/SHA1(NTPv4) authentication. Valid values: `disable`, `enable`.

* `id` - NTP server ID.
* `ip_type` - Choose to connect to IPv4 or/and IPv6 NTP server. Valid values: `IPv6`, `IPv4`, `Both`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `key` - Key for MD5(NTPv3)/SHA1(NTPv4) authentication.
* `key_id` - Key ID for authentication.
* `key_type` - Select NTP authentication type. Valid values: `SHA1`, `SHA256`, `MD5`.

* `ntpv3` - Enable to use NTPv3 instead of NTPv4. Valid values: `disable`, `enable`.

* `server` - IP address or hostname of the NTP Server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Systemp SystemNtp can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_system_ntp.labelname SystempSystemNtp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
