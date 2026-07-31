---
subcategory: "Object Wireless Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_lwprofile"
description: |-
  Configure LoRaWAN profile.
---

# fortimanager_object_wirelesscontroller_lwprofile
Configure LoRaWAN profile.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `cups_api_key` - CUPS API key of LoRaWAN device.
* `cups_server` - CUPS (Configuration and Update Server) domain name or IP address of LoRaWAN device.
* `cups_server_port` - CUPS Port value of LoRaWAN device.
* `lw_protocol` - Configure LoRaWAN protocol (default = basics-station) Valid values: `basics-station`, `packet-forwarder`.

* `name` - LoRaWAN profile name.
* `tc_api_key` - TC API key of LoRaWAN device.
* `tc_server` - TC (Traffic Controller) domain name or IP address of LoRaWAN device.
* `tc_server_port` - TC Port value of LoRaWAN device.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController LwProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_lwprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
