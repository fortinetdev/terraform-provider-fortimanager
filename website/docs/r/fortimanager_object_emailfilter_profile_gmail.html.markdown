---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_emailfilter_profile_gmail"
description: |-
  Gmail.
---

# fortimanager_object_emailfilter_profile_gmail
Gmail.

~> This resource is a sub resource for variable `gmail` of resource `fortimanager_object_emailfilter_profile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_emailfilter_profile_gmail" "trname" {
  log_all    = "enable"
  depends_on = [fortimanager_object_emailfilter_profile.trname]
  profile    = fortimanager_object_emailfilter_profile.trname.name
}

resource "fortimanager_object_emailfilter_profile" "trname" {
  name = "terr-emailfilter-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectEmailfilter ProfileGmail can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_emailfilter_profile_gmail.labelname ObjectEmailfilterProfileGmail
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
