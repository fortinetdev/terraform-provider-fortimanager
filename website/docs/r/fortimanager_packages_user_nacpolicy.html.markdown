---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_user_nacpolicy"
description: |-
  Configure NAC policy matching pattern to identify matching NAC devices.
---

# fortimanager_packages_user_nacpolicy
Configure NAC policy matching pattern to identify matching NAC devices.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `category` - Category of NAC policy. Valid values: `device`, `firewall-user`, `ems-tag`.

* `description` - Description for the NAC policy matching pattern.
* `ems_tag` - NAC policy matching EMS tag.
* `family` - NAC policy matching family.
* `host` - NAC policy matching host.
* `hw_vendor` - NAC policy matching hardware vendor.
* `hw_version` - NAC policy matching hardware version.
* `mac` - NAC policy matching MAC address.
* `name` - NAC policy name.
* `os` - NAC policy matching operating system.
* `severity` - NAC policy matching devices vulnerability severity lists.
* `src` - NAC policy matching source.
* `ssid_policy` - SSID policy to be applied on the matched NAC policy.
* `status` - Enable/disable NAC policy. Valid values: `disable`, `enable`.

* `sw_version` - NAC policy matching software version.
* `type` - NAC policy matching type.
* `user` - NAC policy matching user.
* `user_group` - NAC policy matching user group.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Packages UserNacPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_user_nacpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
