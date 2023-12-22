---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_lldpprofile_customtlvs"
description: |-
  Configuration method to edit custom TLV entries.
---

# fortimanager_object_switchcontroller_lldpprofile_customtlvs
Configuration method to edit custom TLV entries.

~> This resource is a sub resource for variable `custom_tlvs` of resource `fortimanager_object_switchcontroller_lldpprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_switchcontroller_lldpprofile_customtlvs" "trname" {
  lldp_profile       = fortimanager_object_switchcontroller_lldpprofile.trname.name
  information_string = 400
  name               = "terr-customtlvs"
  depends_on         = [fortimanager_object_switchcontroller_lldpprofile.trname]
}

resource "fortimanager_object_switchcontroller_lldpprofile" "trname" {
  name = "terr-switch-controller-lldp-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `lldp_profile` - Lldp Profile.

* `information_string` - Organizationally defined information string (0 - 507 hexadecimal bytes).
* `name` - TLV name (not sent).
* `oui` - Organizationally unique identifier (OUI), a 3-byte hexadecimal number, for this TLV.
* `subtype` - Organizationally defined subtype (0 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController LldpProfileCustomTlvs can be imported using any of these accepted formats:
```
Set import_options = ["lldp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_lldpprofile_customtlvs.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
