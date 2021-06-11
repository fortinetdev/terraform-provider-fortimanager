---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_log_customfield"
description: |-
  Configure custom log fields.
---

# fortimanager_object_log_customfield
Configure custom log fields.

## Example Usage

```hcl
resource "fortimanager_object_log_customfield" "trname" {
  fosid = "terr-log-custom-field"
  name  = "terr-log-custom"
  value = "Terraform tefv"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `fosid` - field ID &lt;string&gt;.
* `name` - Field name (max: 15 characters).
* `value` - Field value (max: 15 characters).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectLog CustomField can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_log_customfield.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
