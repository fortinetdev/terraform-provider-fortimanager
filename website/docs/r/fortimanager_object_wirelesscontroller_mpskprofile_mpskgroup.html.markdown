---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup"
description: |-
  List of multiple PSK groups.
---

# fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup
List of multiple PSK groups.

~> This resource is a sub resource for variable `mpsk_group` of resource `fortimanager_object_wirelesscontroller_mpskprofile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `mpsk_key`: `fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup" "trname" {
  name         = "tst"
  vlan_id      = 1300
  vlan_type    = "no-vlan"
  mpsk_profile = fortimanager_object_wirelesscontroller_mpskprofile.trname.name
  depends_on   = [fortimanager_object_wirelesscontroller_mpskprofile.trname]
}

resource "fortimanager_object_wirelesscontroller_mpskprofile" "trname" {
  name = "terr-wictl-mpsk-profile67"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `mpsk_profile` - Mpsk Profile.

* `mpsk_key` - Mpsk-Key. The structure of `mpsk_key` block is documented below.
* `name` - MPSK group name.
* `vlan_id` - Optional VLAN ID.
* `vlan_type` - MPSK group VLAN options. Valid values: `no-vlan`, `fixed-vlan`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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

ObjectWirelessController MpskProfileMpskGroup can be imported using any of these accepted formats:
```
Set import_options = ["mpsk_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
