---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_wtpprofile_denymaclist"
description: |-
  List of MAC addresses that are denied access to this WTP, FortiAP, or AP.
---

# fortimanager_object_wirelesscontroller_wtpprofile_denymaclist
List of MAC addresses that are denied access to this WTP, FortiAP, or AP.

~> This resource is a sub resource for variable `deny_mac_list` of resource `fortimanager_object_wirelesscontroller_wtpprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_wtpprofile_denymaclist" "trname" {
  fosid       = 1
  mac         = "4a:7e:1e:d2:9b:86"
  wtp_profile = fortimanager_object_wirelesscontroller_wtpprofile.trname.name
  depends_on  = [fortimanager_object_wirelesscontroller_wtpprofile.trname]
}

resource "fortimanager_object_wirelesscontroller_wtpprofile" "trname" {
  name = "terr-wtpprofile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wtp_profile` - Wtp Profile.

* `fosid` - ID.
* `mac` - A WiFi device with this MAC address is denied access to this WTP, FortiAP or AP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWirelessController WtpProfileDenyMacList can be imported using any of these accepted formats:
```
Set import_options = ["wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_wtpprofile_denymaclist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
