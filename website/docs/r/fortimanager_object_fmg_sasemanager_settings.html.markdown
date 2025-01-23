---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fmg_sasemanager_settings"
description: |-
  ObjectFmg SaseManagerSettings
---

# fortimanager_object_fmg_sasemanager_settings
ObjectFmg SaseManagerSettings

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `address` - Address.
* `profile_group` - Profile-Group.
* `sync_address` - Sync-Address. Valid values: `disable`, `specify`, `all`.

* `sync_profile_group` - Sync-Profile-Group. Valid values: `disable`, `specify`, `all`.

* `sync_user` - Sync-User. Valid values: `disable`, `specify`, `all`.

* `user` - User.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFmg SaseManagerSettings can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fmg_sasemanager_settings.labelname ObjectFmgSaseManagerSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
