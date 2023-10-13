---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_wtpprofile_lan"
description: |-
  WTP LAN port mapping.
---

# fortimanager_object_wirelesscontroller_wtpprofile_lan
WTP LAN port mapping.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wtp_profile` - Wtp Profile.

* `port_esl_mode` - ESL port mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port_esl_ssid` - Bridge ESL port to SSID.
* `port_mode` - LAN port mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port_ssid` - Bridge LAN port to SSID.
* `port1_mode` - LAN port 1 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port1_ssid` - Bridge LAN port 1 to SSID.
* `port2_mode` - LAN port 2 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port2_ssid` - Bridge LAN port 2 to SSID.
* `port3_mode` - LAN port 3 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port3_ssid` - Bridge LAN port 3 to SSID.
* `port4_mode` - LAN port 4 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port4_ssid` - Bridge LAN port 4 to SSID.
* `port5_mode` - LAN port 5 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port5_ssid` - Bridge LAN port 5 to SSID.
* `port6_mode` - LAN port 6 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port6_ssid` - Bridge LAN port 6 to SSID.
* `port7_mode` - LAN port 7 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port7_ssid` - Bridge LAN port 7 to SSID.
* `port8_mode` - LAN port 8 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port8_ssid` - Bridge LAN port 8 to SSID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWirelessController WtpProfileLan can be imported using any of these accepted formats:
```
Set import_options = ["wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_wtpprofile_lan.labelname ObjectWirelessControllerWtpProfileLan
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
