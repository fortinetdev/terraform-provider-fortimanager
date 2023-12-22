---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_profile_antiphish_inspectionentries"
description: |-
  AntiPhishing entries.
---

# fortimanager_object_webfilter_profile_antiphish_inspectionentries
AntiPhishing entries.

~> This resource is a sub resource for variable `inspection_entries` of resource `fortimanager_object_webfilter_profile_antiphish`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_webfilter_profile_antiphish_inspectionentries" "trname" {
  action     = "log"
  name       = "terr-inspectionentries"
  profile    = fortimanager_object_webfilter_profile.trname4.name
  depends_on = [fortimanager_object_webfilter_profile.trname4]
}

resource "fortimanager_object_webfilter_profile" "trname4" {
  name = "terr-webfilter-profile4"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `action` - Action to be taken upon an AntiPhishing match. Valid values: `log`, `block`, `exempt`.

* `fortiguard_category` - FortiGuard category to match.
* `name` - Inspection target name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWebfilter ProfileAntiphishInspectionEntries can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_profile_antiphish_inspectionentries.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
