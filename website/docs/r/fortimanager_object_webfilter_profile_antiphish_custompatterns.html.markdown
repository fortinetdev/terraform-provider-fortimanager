---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_profile_antiphish_custompatterns"
description: |-
  Custom username and password regex patterns.
---

# fortimanager_object_webfilter_profile_antiphish_custompatterns
Custom username and password regex patterns.

~> This resource is a sub resource for variable `custom_patterns` of resource `fortimanager_object_webfilter_profile_antiphish`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_webfilter_profile_antiphish_custompatterns" "trname" {
  category   = "username"
  pattern    = "tset"
  profile    = fortimanager_object_webfilter_profile.trname5.name
  depends_on = [fortimanager_object_webfilter_profile.trname5]
}

resource "fortimanager_object_webfilter_profile" "trname5" {
  name = "terr-webfilter-profile5"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `category` - Category that the pattern matches. Valid values: `username`, `password`.

* `pattern` - Target pattern.
* `type` - Pattern will be treated either as a regex pattern or literal string. Valid values: `regex`, `literal`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{pattern}}.

## Import

ObjectWebfilter ProfileAntiphishCustomPatterns can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_profile_antiphish_custompatterns.labelname {{pattern}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
