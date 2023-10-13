---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_waf_profile_signature"
description: |-
  WAF signatures.
---

# fortimanager_object_waf_profile_signature
WAF signatures.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `credit_card_detection_threshold` - The minimum number of Credit cards to detect violation.
* `custom_signature` - Custom-Signature. The structure of `custom_signature` block is documented below.
* `disabled_signature` - Disabled signatures
* `disabled_sub_class` - Disabled signature subclasses.
* `main_class` - Main-Class. The structure of `main_class` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_signature` block supports:

* `action` - Action. Valid values: `allow`, `block`, `erase`.

* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable`, `enable`.

* `direction` - Traffic direction. Valid values: `request`, `response`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `name` - Signature name.
* `pattern` - Match pattern.
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.

* `target` - Match HTTP target. Valid values: `arg`, `arg-name`, `req-body`, `req-cookie`, `req-cookie-name`, `req-filename`, `req-header`, `req-header-name`, `req-raw-uri`, `req-uri`, `resp-body`, `resp-hdr`, `resp-status`.


The `main_class` block supports:

* `action` - Action. Valid values: `allow`, `block`, `erase`.

* `id` - Main signature class ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWaf ProfileSignature can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_waf_profile_signature.labelname ObjectWafProfileSignature
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
