---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_saasapplication_outputattributes"
description: |-
  SaaS application output attributes.
---

# fortimanager_object_casb_saasapplication_outputattributes
SaaS application output attributes.

~> This resource is a sub resource for variable `output_attributes` of resource `fortimanager_object_casb_saasapplication`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `saas_application` - Saas Application.

* `attr_type` - CASB attribute type. Valid values: `tenant`.

* `description` - CASB attribute description.
* `name` - CASB attribute name.
* `required` - CASB attribute required. Valid values: `disable`, `enable`.

* `type` - CASB attribute format type. Valid values: `string`, `string-list`, `integer`, `integer-list`, `boolean`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCasb SaasApplicationOutputAttributes can be imported using any of these accepted formats:
```
Set import_options = ["saas_application=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_saasapplication_outputattributes.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
