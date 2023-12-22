---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_antivirus_profile_nacquar"
description: |-
  Configure AntiVirus quarantine settings.
---

# fortimanager_object_antivirus_profile_nacquar
Configure AntiVirus quarantine settings.

~> This resource is a sub resource for variable `nac_quar` of resource `fortimanager_object_antivirus_profile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_antivirus_profile_nacquar" "trname" {
  infected   = "none"
  log        = "enable"
  profile    = fortimanager_object_antivirus_profile.trname.name
  depends_on = [fortimanager_object_antivirus_profile.trname]
}

resource "fortimanager_object_antivirus_profile" "trname" {
  name = "terr-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `expiry` - Duration of quarantine.
* `infected` - Enable/Disable quarantining infected hosts to the banned user list. Valid values: `none`, `quar-src-ip`, `quar-interface`.

* `log` - Enable/disable AntiVirus quarantine logging. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectAntivirus ProfileNacQuar can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_antivirus_profile_nacquar.labelname ObjectAntivirusProfileNacQuar
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
