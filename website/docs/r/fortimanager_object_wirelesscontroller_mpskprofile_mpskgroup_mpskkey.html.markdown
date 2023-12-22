---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey"
description: |-
  List of multiple PSK entries.
---

# fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey
List of multiple PSK entries.

~> This resource is a sub resource for variable `mpsk_key` of resource `fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey" "trname" {
  name         = "terr-mpskkey"
  passphrase   = ["psph"]
  mpsk_profile = fortimanager_object_wirelesscontroller_mpskprofile.trname.name
  mpsk_group   = fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup.trname.name
  depends_on   = [fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup.trname]
}

resource "fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup" "trname" {
  name         = "tst"
  mpsk_profile = fortimanager_object_wirelesscontroller_mpskprofile.trname.name
  depends_on   = [fortimanager_object_wirelesscontroller_mpskprofile.trname]
}

resource "fortimanager_object_wirelesscontroller_mpskprofile" "trname" {
  name = "terr-wictl-mpsk-profile2"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `mpsk_profile` - Mpsk Profile.
* `mpsk_group` - Mpsk Group.

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

ObjectWirelessController MpskProfileMpskGroupMpskKey can be imported using any of these accepted formats:
```
Set import_options = ["mpsk_profile=YOUR_VALUE", "mpsk_group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
