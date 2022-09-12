---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_fortilinksettings"
description: |-
  Configure integrated FortiLink settings for FortiSwitch.
---

# fortimanager_object_switchcontroller_fortilinksettings
Configure integrated FortiLink settings for FortiSwitch.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `fortilink` - FortiLink interface to which this fortilink-setting belongs.
* `inactive_timer` - Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
* `link_down_flush` - Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.

* `nac_ports` - Nac-Ports. The structure of `nac_ports` block is documented below.
* `name` - FortiLink settings name.

The `nac_ports` block supports:

* `bounce_nac_port` - Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.

* `lan_segment` - Enable/disable LAN segment feature on the FortiLink interface. Valid values: `disabled`, `enabled`.

* `member_change` - Member-Change.
* `nac_lan_interface` - Configure NAC LAN interface.
* `nac_segment_vlans` - Configure NAC segment VLANs.
* `onboarding_vlan` - Default NAC Onboarding VLAN when NAC devices are discovered.
* `parent_key` - Parent-Key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController FortilinkSettings can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_fortilinksettings.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
