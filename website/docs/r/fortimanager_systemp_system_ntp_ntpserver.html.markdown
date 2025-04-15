---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_system_ntp_ntpserver"
description: |-
  Configure the FortiGate to connect to any available third-party NTP server.
---

# fortimanager_systemp_system_ntp_ntpserver
Configure the FortiGate to connect to any available third-party NTP server.

~> This resource is a sub resource for variable `ntpserver` of resource `fortimanager_systemp_system_ntp`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `authentication` - Enable/disable MD5(NTPv3)/SHA1(NTPv4) authentication. Valid values: `disable`, `enable`.

* `fosid` - NTP server ID.
* `ip_type` - Choose to connect to IPv4 or/and IPv6 NTP server. Valid values: `IPv6`, `IPv4`, `Both`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `key` - Key for MD5(NTPv3)/SHA1(NTPv4) authentication.
* `key_id` - Key ID for authentication.
* `key_type` - Select NTP authentication type. Valid values: `SHA1`, `SHA256`, `MD5`.

* `ntpv3` - Enable to use NTPv3 instead of NTPv4. Valid values: `disable`, `enable`.

* `server` - IP address or hostname of the NTP Server.
* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Systemp SystemNtpNtpserver can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_system_ntp_ntpserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
