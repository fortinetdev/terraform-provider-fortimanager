---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_router_accesslist6_rule"
description: |-
  Rule.
---

# fortimanager_object_router_accesslist6_rule
Rule.

~> This resource is a sub resource for variable `rule` of resource `fortimanager_object_router_accesslist6`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_router_accesslist6_rule" "trname" {
  access_list6 = fortimanager_object_router_accesslist6.trname.name
  action       = "permit"
  fosid        = 1
  depends_on   = [fortimanager_object_router_accesslist6.trname]
}

resource "fortimanager_object_router_accesslist6" "trname" {
  name = "terr-router-accesslist6"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `access_list6` - Access List6.

* `action` - Permit or deny this IP address and netmask prefix. Valid values: `permit`, `deny`.

* `exact_match` - Enable/disable exact prefix match. Valid values: `disable`, `enable`.

* `flags` - Flags.
* `fosid` - Rule ID.
* `prefix6` - IPv6 prefix to define regular filter criteria, such as "any" or subnets.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectRouter AccessList6Rule can be imported using any of these accepted formats:
```
Set import_options = ["access_list6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_router_accesslist6_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
