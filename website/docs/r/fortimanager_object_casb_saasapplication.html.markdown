---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_saasapplication"
description: |-
  Configure CASB SaaS application.
---

# fortimanager_object_casb_saasapplication
Configure CASB SaaS application.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `input_attributes`: `fortimanager_object_casb_saasapplication_inputattributes`
>- `output_attributes`: `fortimanager_object_casb_saasapplication_outputattributes`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `casb_name` - SaaS application signature name.
* `description` - SaaS application description.
* `domains` - SaaS application domain list.
* `input_attributes` - Input-Attributes. The structure of `input_attributes` block is documented below.
* `name` - SaaS application name.
* `output_attributes` - Output-Attributes. The structure of `output_attributes` block is documented below.
* `status` - Enable/disable setting. Valid values: `disable`, `enable`.

* `type` - SaaS application type. Valid values: `built-in`, `customized`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `input_attributes` block supports:

* `attr_type` - CASB attribute type. Valid values: `tenant`.

* `default` - CASB attribute default value. Valid values: `string`, `string-list`.

* `description` - CASB attribute description.
* `fallback_input` - CASB attribute legacy input. Valid values: `disable`, `enable`.

* `name` - CASB attribute name.
* `required` - CASB attribute required. Valid values: `disable`, `enable`.

* `type` - CASB attribute format type. Valid values: `string`, `string-list`, `integer`, `integer-list`, `boolean`.


The `output_attributes` block supports:

* `attr_type` - CASB attribute type. Valid values: `tenant`.

* `description` - CASB attribute description.
* `name` - CASB attribute name.
* `required` - CASB attribute required. Valid values: `disable`, `enable`.

* `type` - CASB attribute format type. Valid values: `string`, `string-list`, `integer`, `integer-list`, `boolean`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCasb SaasApplication can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_saasapplication.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
