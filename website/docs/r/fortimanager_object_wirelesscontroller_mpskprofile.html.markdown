---
subcategory: "ObjectWireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_mpskprofile"
description: |-
  Configure MPSK profile.
---

# fortimanager_object_wirelesscontroller_mpskprofile
Configure MPSK profile.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `mpsk_concurrent_clients` - Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
* `mpsk_group` - Mpsk-Group. The structure of `mpsk_group` block is documented below.
* `name` - MPSK profile name.
* `ssid` - SSID of the VAP in which the MPSK profile is configured.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `mpsk_group` block supports:

* `mpsk_key` - Mpsk-Key. The structure of `mpsk_key` block is documented below.
* `name` - MPSK group name.
* `vlan_id` - Optional VLAN ID.
* `vlan_type` - MPSK group VLAN options. Valid values: `no-vlan`, `fixed-vlan`.


The `mpsk_key` block supports:

* `comment` - Comment.
* `concurrent_client_limit_type` - MPSK client limit type options. Valid values: `default`, `unlimited`, `specified`.

* `concurrent_clients` - Number of clients that can connect using this pre-shared key (1 - 65535, default is 256).
* `mac` - MAC address.
* `mpsk_schedules` - Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid.
* `name` - Pre-shared key name.
* `passphrase` - WPA Pre-shared key.
* `pmk` - WPA PMK.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController MpskProfile can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_mpskprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.