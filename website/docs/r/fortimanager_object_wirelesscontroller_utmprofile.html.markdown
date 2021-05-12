---
subcategory: "ObjectWireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_utmprofile"
description: |-
  Configure UTM (Unified Threat Management) profile.
---

# fortimanager_object_wirelesscontroller_utmprofile
Configure UTM (Unified Threat Management) profile.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `antivirus_profile` - AntiVirus profile name.
* `application_list` - Application control list name.
* `comment` - Comment.
* `ips_sensor` - IPS sensor name.
* `name` - UTM profile name.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.

* `utm_log` - Enable/disable UTM logging. Valid values: `disable`, `enable`.

* `webfilter_profile` - WebFilter profile name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController UtmProfile can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_utmprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.