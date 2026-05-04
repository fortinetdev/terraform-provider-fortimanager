---
subcategory: "Object Isolator"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_isolator_profile_entries"
description: |-
  ObjectIsolator ProfileEntries
---

# fortimanager_object_isolator_profile_entries
ObjectIsolator ProfileEntries

~> This resource is a sub resource for variable `entries` of resource `fortimanager_object_isolator_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `action` - Action. Valid values: `isolate`, `freeze`, `block`, `allow`.

* `copy_paste` - Copy-Paste. Valid values: `disable`, `enable`.

* `fosid` - Id.
* `proxy_address` - Proxy-Address.
* `right_click` - Right-Click. Valid values: `disable`, `enable`.

* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectIsolator ProfileEntries can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_isolator_profile_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
