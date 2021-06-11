---
subcategory: "Object WAF"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_waf_subclass"
description: |-
  Hidden table for datasource.
---

# fortimanager_object_waf_subclass
Hidden table for datasource.

## Example Usage

```hcl
resource "fortimanager_object_waf_subclass" "trname" {
  fosid = 1
  name  = "terr-waf-sub-class"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `fosid` - Signature subclass ID.
* `name` - Signature subclass name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWaf SubClass can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_waf_subclass.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
