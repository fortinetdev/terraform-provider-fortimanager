---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_deviceaccesslist"
description: |-
  Configure device access control lists.
---

# fortimanager_object_user_deviceaccesslist
Configure device access control lists.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`device_list`: `fortimanager_object_user_deviceaccesslist_devicelist`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `default_action` - Accept or deny unknown/unspecified devices. Valid values: `deny`, `accept`.

* `device_list` - Device-List. The structure of `device_list` block is documented below.
* `name` - Device access list name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `device_list` block supports:

* `action` - Allow or block device. Valid values: `deny`, `accept`.

* `device` - Firewall device or device group.
* `id` - Entry ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser DeviceAccessList can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_deviceaccesslist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
