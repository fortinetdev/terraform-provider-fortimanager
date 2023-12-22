---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_vlanpolicy"
description: |-
  Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.
---

# fortimanager_object_switchcontroller_vlanpolicy
Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.

## Example Usage

```hcl
resource "fortimanager_object_switchcontroller_vlanpolicy" "trname" {
  name              = "terr-vlanpolicy"
  allowed_vlans_all = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `allowed_vlans` - Allowed VLANs to be applied when using this VLAN policy.
* `allowed_vlans_all` - Enable/disable all defined VLANs when using this VLAN policy. Valid values: `disable`, `enable`.

* `description` - Description for the VLAN policy.
* `discard_mode` - Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.

* `name` - VLAN policy name.
* `untagged_vlans` - Untagged VLANs to be applied when using this VLAN policy.
* `vlan` - Native VLAN to be applied when using this VLAN policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController VlanPolicy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_vlanpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
