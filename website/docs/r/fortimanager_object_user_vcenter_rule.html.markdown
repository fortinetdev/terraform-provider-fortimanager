---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_vcenter_rule"
description: |-
  ObjectUser VcenterRule
---

# fortimanager_object_user_vcenter_rule
ObjectUser VcenterRule

~> This resource is a sub resource for variable `rule` of resource `fortimanager_object_user_vcenter`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_user_vcenter_rule" "trname" {
  vcenter    = fortimanager_object_user_vcenter.trname.name
  name       = "terr-rule"
  depends_on = [fortimanager_object_user_vcenter.trname]
}

resource "fortimanager_object_user_vcenter" "trname" {
  name   = "terr-vcenter"
  server = "192.168.1.1"
  user   = "admin"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vcenter` - Vcenter.

* `name` - Name.
* `rule` - Rule.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser VcenterRule can be imported using any of these accepted formats:
```
Set import_options = ["vcenter=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_vcenter_rule.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
