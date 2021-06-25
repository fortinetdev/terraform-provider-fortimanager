---
subcategory: "Object Emailfilter"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_emailfilter_fortishield"
description: |-
  Configure FortiGuard - AntiSpam.
---

# fortimanager_object_emailfilter_fortishield
Configure FortiGuard - AntiSpam.

## Example Usage

```hcl
resource "fortimanager_object_emailfilter_fortishield" "trname" {
  spam_submit_force   = "enable"
  spam_submit_srv     = "www.nospammer1.net"
  spam_submit_txt2htm = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `spam_submit_force` - Enable/disable force insertion of a new mime entity for the submission text. Valid values: `disable`, `enable`.

* `spam_submit_srv` - Hostname of the spam submission server.
* `spam_submit_txt2htm` - Enable/disable conversion of text email to HTML email. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectEmailfilter Fortishield can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_emailfilter_fortishield.labelname ObjectEmailfilterFortishield
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
