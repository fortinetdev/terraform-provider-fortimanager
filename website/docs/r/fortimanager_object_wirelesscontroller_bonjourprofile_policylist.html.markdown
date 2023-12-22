---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_bonjourprofile_policylist"
description: |-
  Bonjour policy list.
---

# fortimanager_object_wirelesscontroller_bonjourprofile_policylist
Bonjour policy list.

~> This resource is a sub resource for variable `policy_list` of resource `fortimanager_object_wirelesscontroller_bonjourprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_bonjourprofile_policylist" "trname" {
  bonjour_profile = fortimanager_object_wirelesscontroller_bonjourprofile.trname.name
  policy_id       = 1
  services        = ["ftp"]
  to_vlan         = 2000
  depends_on      = [fortimanager_object_wirelesscontroller_bonjourprofile.trname]
}

resource "fortimanager_object_wirelesscontroller_bonjourprofile" "trname" {
  name = "teset"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `bonjour_profile` - Bonjour Profile.

* `description` - Description.
* `from_vlan` - VLAN ID from which the Bonjour service is advertised (0 - 4094, default = 0).
* `policy_id` - Policy ID.
* `services` - Bonjour services for the VLAN connecting to the Bonjour network. Valid values: `airplay`, `afp`, `bit-torrent`, `ftp`, `ichat`, `itunes`, `printers`, `samba`, `scanners`, `ssh`, `chromecast`, `all`.

* `to_vlan` - VLAN ID to which the Bonjour service is made available (0 - 4094, default = all).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policy_id}}.

## Import

ObjectWirelessController BonjourProfilePolicyList can be imported using any of these accepted formats:
```
Set import_options = ["bonjour_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_bonjourprofile_policylist.labelname {{policy_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
