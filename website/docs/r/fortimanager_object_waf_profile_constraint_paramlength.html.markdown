---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_waf_profile_constraint_paramlength"
description: |-
  Maximum length of parameter in URL, HTTP POST request or HTTP body.
---

# fortimanager_object_waf_profile_constraint_paramlength
Maximum length of parameter in URL, HTTP POST request or HTTP body.

~> This resource is a sub resource for variable `param_length` of resource `fortimanager_object_waf_profile_constraint`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_waf_profile_constraint_paramlength" "trname" {
  action     = "block"
  length     = 120
  log        = "enable"
  profile    = fortimanager_object_waf_profile.trname.name
  depends_on = [fortimanager_object_waf_profile.trname]
}

resource "fortimanager_object_waf_profile" "trname" {
  name = "terr-waf-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Maximum length of parameter in URL, HTTP POST request or HTTP body in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWaf ProfileConstraintParamLength can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_waf_profile_constraint_paramlength.labelname ObjectWafProfileConstraintParamLength
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
