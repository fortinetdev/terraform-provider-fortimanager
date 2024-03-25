---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_profile_writepasswdprofiles"
description: |-
  Profile list.
---

# fortimanager_system_admin_profile_writepasswdprofiles
Profile list.

~> This resource is a sub resource for variable `write_passwd_profiles` of resource `fortimanager_system_admin_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `profile` - Profile.

* `profileid` - Profile ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{profileid}}.

## Import

System AdminProfileWritePasswdProfiles can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_profile_writepasswdprofiles.labelname {{profileid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

