---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_icap_profile_respmodforwardrules_headergroup"
description: |-
  HTTP header group.
---

# fortimanager_object_icap_profile_respmodforwardrules_headergroup
HTTP header group.

~> This resource is a sub resource for variable `header_group` of resource `fortimanager_object_icap_profile_respmodforwardrules`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.
* `respmod_forward_rules` - Respmod Forward Rules.

* `case_sensitivity` - Enable/disable case sensitivity when matching header. Valid values: `disable`, `enable`.

* `header` - HTTP header regular expression.
* `header_name` - HTTP header.
* `fosid` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectIcap ProfileRespmodForwardRulesHeaderGroup can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE", "respmod_forward_rules=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_icap_profile_respmodforwardrules_headergroup.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
