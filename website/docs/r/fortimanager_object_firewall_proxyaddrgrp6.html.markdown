---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_proxyaddrgrp6"
description: |-
  ObjectFirewall ProxyAddrgrp6
---

# fortimanager_object_firewall_proxyaddrgrp6
ObjectFirewall ProxyAddrgrp6

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `tagging`: `fortimanager_object_firewall_proxyaddrgrp6_tagging`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `color` - Color.
* `comment` - Comment.
* `logic_type` - Logic-Type. Valid values: `or`, `and`.

* `member` - Member.
* `name` - Name.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `type` - Type. Valid values: `src`, `dst`.

* `uuid` - Uuid.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `tagging` block supports:

* `category` - Category.
* `name` - Name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ProxyAddrgrp6 can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_proxyaddrgrp6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
