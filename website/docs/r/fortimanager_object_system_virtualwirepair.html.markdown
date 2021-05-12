---
subcategory: "ObjectSystem"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_virtualwirepair"
description: |-
  Configure virtual wire pairs.
---

# fortimanager_object_system_virtualwirepair
Configure virtual wire pairs.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `member` - Interfaces belong to the virtual-wire-pair.
* `name` - Virtual-wire-pair name. Must be a unique interface name.
* `vlan_filter` - Set VLAN filters.
* `wildcard_vlan` - Enable/disable wildcard VLAN. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem VirtualWirePair can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_virtualwirepair.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
