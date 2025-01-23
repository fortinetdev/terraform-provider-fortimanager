---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_profile_datamaskcustomfields"
description: |-
  Customized datamask fields.
---

# fortimanager_system_admin_profile_datamaskcustomfields
Customized datamask fields.

~> This resource is a sub resource for variable `datamask_custom_fields` of resource `fortimanager_system_admin_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `profile` - Profile.

* `field_category` - Field categories. log - Log. fortiview - FortiView. alert - Event management. ueba - UEBA. all - All. Valid values: `log`, `fortiview`, `alert`, `ueba`, `all`.

* `field_name` - Field name.
* `field_status` - Field status. disable - Disable field. enable - Enable field. Valid values: `disable`, `enable`.

* `field_type` - Field type. string - String. ip - IP. mac - MAC address. email - Email address. unknown - Unknown. Valid values: `string`, `ip`, `mac`, `email`, `unknown`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{field_name}}.

## Import

System AdminProfileDatamaskCustomFields can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_profile_datamaskcustomfields.labelname {{field_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

