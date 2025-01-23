---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_addrgrp_dynamic_mapping"
description: |-
  Configure IPv4 address groups.
---

# fortimanager_object_firewall_addrgrp_dynamic_mapping
Configure IPv4 address groups.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_firewall_addrgrp`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `addrgrp` - Addrgrp.

* `_image_base64` - _Image-Base64.
* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `allow_routing` - Enable/disable use of this group in the static route configuration. Valid values: `disable`, `enable`.

* `category` - Category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `exclude` - Enable/disable address exclusion. Valid values: `disable`, `enable`.

* `exclude_member` - Address exclusion member.
* `fabric_object` - Fabric-Object. Valid values: `disable`, `enable`.

* `global_object` - Global-Object.
* `member` - Address objects contained within the group.
* `tags` - Tags.
* `type` - Type. Valid values: `default`, `array`, `folder`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility in the GUI. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectFirewall AddrgrpDynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["addrgrp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_addrgrp_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
