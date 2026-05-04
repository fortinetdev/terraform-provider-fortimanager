---
subcategory: "Object Isolator"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_isolator_profile"
description: |-
  ObjectIsolator Profile
---

# fortimanager_object_isolator_profile
ObjectIsolator Profile

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fortimanager_object_isolator_profile_entries`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comments` - Comments.
* `disclaimer` - Disclaimer.
* `entries` - Entries. The structure of `entries` block is documented below.
* `name` - Name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `action` - Action. Valid values: `isolate`, `freeze`, `block`, `allow`.

* `copy_paste` - Copy-Paste. Valid values: `disable`, `enable`.

* `id` - Id.
* `proxy_address` - Proxy-Address.
* `right_click` - Right-Click. Valid values: `disable`, `enable`.

* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectIsolator Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_isolator_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
