---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_vap_macfilterlist"
description: |-
  Create a list of MAC addresses for MAC address filtering.
---

# fortimanager_object_wirelesscontroller_vap_macfilterlist
Create a list of MAC addresses for MAC address filtering.

~> This resource is a sub resource for variable `mac_filter_list` of resource `fortimanager_object_wirelesscontroller_vap`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_vap_macfilterlist" "trname" {
  fosid             = 1
  mac               = "4a:7e:1e:d2:9b:86"
  mac_filter_policy = "allow"
  depends_on        = [fortimanager_object_wirelesscontroller_vap.trname]
  vap               = fortimanager_object_wirelesscontroller_vap.trname.name
}

resource "fortimanager_object_wirelesscontroller_vap" "trname" {
  name = "terr-wictl-vap2"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vap` - Vap.

* `fosid` - ID.
* `mac` - MAC address.
* `mac_filter_policy` - Deny or allow the client with this MAC address. Valid values: `deny`, `allow`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWirelessController VapMacFilterList can be imported using any of these accepted formats:
```
Set import_options = ["vap=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_vap_macfilterlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
