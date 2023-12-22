---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_vap_vlanname"
description: |-
  Table for mapping VLAN name to VLAN ID.
---

# fortimanager_object_wirelesscontroller_vap_vlanname
Table for mapping VLAN name to VLAN ID.

~> This resource is a sub resource for variable `vlan_name` of resource `fortimanager_object_wirelesscontroller_vap`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_vap_vlanname" "trname" {
  name       = "terr-vlanname"
  vlan_id    = 1200
  vap        = fortimanager_object_wirelesscontroller_vap.trname34.name
  depends_on = [fortimanager_object_wirelesscontroller_vap.trname34]
}

resource "fortimanager_object_wirelesscontroller_vap" "trname34" {
  name = "terr-wictl-vap5"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vap` - Vap.

* `name` - VLAN name.
* `vlan_id` - VLAN ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController VapVlanName can be imported using any of these accepted formats:
```
Set import_options = ["vap=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_vap_vlanname.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
