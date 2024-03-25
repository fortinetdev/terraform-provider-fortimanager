---
subcategory: "Object Switch-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_managedswitch"
description: |-
  Configure FortiSwitch devices that are managed by this FortiGate.
---

# fortimanager_object_switchcontroller_managedswitch
Configure FortiSwitch devices that are managed by this FortiGate.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `custom_command`: `fortimanager_object_switchcontroller_managedswitch_customcommand`
>- `dhcp_snooping_static_client`: `fortimanager_object_switchcontroller_managedswitch_dhcpsnoopingstaticclient`
>- `ports`: `fortimanager_object_switchcontroller_managedswitch_ports`
>- `route_offload_router`: `fortimanager_object_switchcontroller_managedswitch_routeoffloadrouter`
>- `vlan`: `fortimanager_object_switchcontroller_managedswitch_vlan`



## Example Usage

```hcl
resource "fortimanager_object_switchcontroller_managedswitch" "trname" {
  _platform               = "FS1D48T418000533"
  description             = "This is a Terraform example"
  dhcp_server_access_list = "enable"
  switch_id               = 3
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_platform` - _Platform.
* `custom_command` - Custom-Command. The structure of `custom_command` block is documented below.
* `description` - Description.
* `dhcp_server_access_list` - DHCP snooping server access list. Valid values: `disable`, `enable`, `global`.

* `dhcp_snooping_static_client` - Dhcp-Snooping-Static-Client. The structure of `dhcp_snooping_static_client` block is documented below.
* `firmware_provision` - Enable/disable provisioning of firmware to FortiSwitches on join connection. Valid values: `disable`, `enable`.

* `firmware_provision_latest` - Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.

* `firmware_provision_version` - Firmware version to provision to this FortiSwitch on bootup (major.minor.build, i.e. 6.2.1234).
* `l3_discovered` - L3-Discovered.
* `mclag_igmp_snooping_aware` - Enable/disable MCLAG IGMP-snooping awareness. Valid values: `disable`, `enable`.

* `mgmt_mode` - FortiLink management mode.
* `name` - Managed-switch name.
* `override_snmp_community` - Enable/disable overriding the global SNMP communities. Valid values: `disable`, `enable`.

* `override_snmp_sysinfo` - Enable/disable overriding the global SNMP system information. Valid values: `disable`, `enable`.

* `override_snmp_trap_threshold` - Enable/disable overriding the global SNMP trap threshold values. Valid values: `disable`, `enable`.

* `override_snmp_user` - Enable/disable overriding the global SNMP users. Valid values: `disable`, `enable`.

* `poe_detection_type` - Poe-Detection-Type.
* `ports` - Ports. The structure of `ports` block is documented below.
* `ptp_profile` - PTP profile configuration.
* `ptp_status` - Enable/disable PTP profile on this FortiSwitch. Valid values: `disable`, `enable`.

* `purdue_level` - Purdue Level of this FortiSwitch. Valid values: `1`, `2`, `3`, `4`, `5`, `1.5`, `2.5`, `3.5`, `5.5`.

* `qos_drop_policy` - Set QoS drop-policy. Valid values: `taildrop`, `random-early-detection`.

* `qos_red_probability` - Set QoS RED/WRED drop probability.
* `radius_nas_ip` - NAS-IP address.
* `radius_nas_ip_override` - Use locally defined NAS-IP. Valid values: `disable`, `enable`.

* `route_offload` - Enable/disable route offload on this FortiSwitch. Valid values: `disable`, `enable`.

* `route_offload_mclag` - Enable/disable route offload MCLAG on this FortiSwitch. Valid values: `disable`, `enable`.

* `route_offload_router` - Route-Offload-Router. The structure of `route_offload_router` block is documented below.
* `switch_dhcp_opt43_key` - DHCP option43 key.
* `switch_id` - Managed-switch id.
* `tdr_supported` - Tdr-Supported.
* `tunnel_discovered` - Tunnel-Discovered.
* `vlan` - Vlan. The structure of `vlan` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_command` block supports:

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Names of commands to be pushed to this FortiSwitch device, as configured under config switch-controller custom-command.

The `dhcp_snooping_static_client` block supports:

* `ip` - Client static IP address.
* `mac` - Client MAC address.
* `name` - Client name.
* `port` - Interface name.
* `vlan` - VLAN name.

The `ports` block supports:

* `access_mode` - Access mode of the port. Valid values: `normal`, `nac`.

* `acl_group` - ACL groups on this port.
* `aggregator_mode` - LACP member select mode. Valid values: `bandwidth`, `count`.

* `allowed_vlans` - Configure switch port tagged vlans
* `allowed_vlans_all` - Enable/disable all defined vlans on this port. Valid values: `disable`, `enable`.

* `arp_inspection_trust` - Trusted or untrusted dynamic ARP inspection. Valid values: `untrusted`, `trusted`.

* `authenticated_port` - Authenticated-Port.
* `bundle` - Enable/disable Link Aggregation Group (LAG) bundling for non-FortiLink interfaces. Valid values: `disable`, `enable`.

* `description` - Description for port.
* `dhcp_snoop_option82_override` - Dhcp-Snoop-Option82-Override. The structure of `dhcp_snoop_option82_override` block is documented below.
* `dhcp_snoop_option82_trust` - Enable/disable allowance of DHCP with option-82 on untrusted interface. Valid values: `disable`, `enable`.

* `dhcp_snooping` - Trusted or untrusted DHCP-snooping interface. Valid values: `trusted`, `untrusted`.

* `discard_mode` - Configure discard mode for port. Valid values: `none`, `all-untagged`, `all-tagged`.

* `dsl_profile` - DSL policy configuration.
* `edge_port` - Enable/disable this interface as an edge port, bridging connections between workstations and/or computers. Valid values: `disable`, `enable`.

* `encrypted_port` - Encrypted-Port.
* `fec_capable` - FEC capable.
* `fec_state` - State of forward error correction. Valid values: `disabled`, `cl74`, `cl91`.

* `flap_duration` - Period over which flap events are calculated (seconds).
* `flap_rate` - Number of stage change events needed within flap-duration.
* `flap_timeout` - Flap guard disabling protection (min).
* `flapguard` - Enable/disable flap guard. Valid values: `disable`, `enable`.

* `flow_control` - Flow control direction. Valid values: `disable`, `tx`, `rx`, `both`.

* `fortiswitch_acls` - ACLs on this port.
* `igmp_snooping_flood_reports` - Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled. Valid values: `disable`, `enable`.

* `igmp_snooping` - Set IGMP snooping mode for the physical port interface. Valid values: `disable`, `enable`.

* `igmps_flood_reports` - Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled. Valid values: `disable`, `enable`.

* `igmps_flood_traffic` - Enable/disable flooding of IGMP snooping traffic to this interface. Valid values: `disable`, `enable`.

* `interface_tags` - Tag(s) associated with the interface for various features including virtual port pool, dynamic port policy.
* `ip_source_guard` - Enable/disable IP source guard. Valid values: `disable`, `enable`.

* `isl_peer_device_sn` - Isl-Peer-Device-Sn.
* `lacp_speed` - end Link Aggregation Control Protocol (LACP) messages every 30 seconds (slow) or every second (fast). Valid values: `slow`, `fast`.

* `learning_limit` - Limit the number of dynamic MAC addresses on this Port (1 - 128, 0 = no limit, default).
* `link_status` - Link-Status. Valid values: `down`, `up`.

* `lldp_profile` - LLDP port TLV profile.
* `lldp_status` - LLDP transmit and receive status. Valid values: `disable`, `rx-only`, `tx-only`, `tx-rx`.

* `loop_guard` - Enable/disable loop-guard on this interface, an STP optimization used to prevent network loops. Valid values: `disabled`, `enabled`.

* `loop_guard_timeout` - Loop-guard timeout (0 - 120 min, default = 45).
* `matched_dpp_intf_tags` - Matched interface tags in the dynamic port policy.
* `matched_dpp_policy` - Matched child policy in the dynamic port policy.
* `max_bundle` - Maximum size of LAG bundle (1 - 24, default = 24)
* `mcast_snooping_flood_traffic` - Enable/disable flooding of IGMP snooping traffic to this interface. Valid values: `disable`, `enable`.

* `mclag` - Enable/disable multi-chassis link aggregation (MCLAG). Valid values: `disable`, `enable`.

* `mclag_icl_port` - Mclag-Icl-Port.
* `media_type` - Media-Type.
* `member_withdrawal_behavior` - Port behavior after it withdraws because of loss of control packets. Valid values: `forward`, `block`.

* `members` - Aggregated LAG bundle interfaces.
* `min_bundle` - Minimum size of LAG bundle (1 - 24, default = 1)
* `mode` - LACP mode: ignore and do not send control messages, or negotiate 802.3ad aggregation passively or actively. Valid values: `static`, `lacp-passive`, `lacp-active`.

* `p2p_port` - P2P-Port.
* `packet_sample_rate` - Packet sampling rate (0 - 99999 p/sec).
* `packet_sampler` - Enable/disable packet sampling on this interface. Valid values: `disabled`, `enabled`.

* `pause_meter` - Configure ingress pause metering rate, in kbps (default = 0, disabled).
* `pause_meter_resume` - Resume threshold for resuming traffic on ingress port. Valid values: `25%`, `50%`, `75%`.

* `poe_max_power` - Poe-Max-Power.
* `poe_mode_bt_cabable` - PoE mode IEEE 802.3BT capable.
* `poe_port_mode` - Configure PoE port mode. Valid values: `ieee802-3af`, `ieee802-3at`, `ieee802-3bt`.

* `poe_port_power` - Configure PoE port power. Valid values: `normal`, `perpetual`, `perpetual-fast`.

* `poe_port_priority` - Configure PoE port priority. Valid values: `critical-priority`, `high-priority`, `low-priority`, `medium-priority`.

* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection. Valid values: `disable`, `enable`.

* `poe_standard` - Poe-Standard.
* `poe_status` - Enable/disable PoE status. Valid values: `disable`, `enable`.

* `port_name` - Switch port name.
* `port_owner` - Switch port name.
* `port_policy` - Switch controller dynamic port policy from available options.
* `port_security_policy` - Switch controller authentication policy to apply to this managed switch from available options.
* `port_selection_criteria` - Algorithm for aggregate port selection. Valid values: `src-mac`, `dst-mac`, `src-dst-mac`, `src-ip`, `dst-ip`, `src-dst-ip`.

* `ptp_status` - Enable/disable PTP policy on this FortiSwitch port. Valid values: `disable`, `enable`.

* `qos_policy` - Switch controller QoS policy from available options.
* `restricted_auth_port` - Restricted-Auth-Port.
* `rpvst_port` - Enable/disable inter-operability with rapid PVST on this interface. Valid values: `disabled`, `enabled`.

* `sample_direction` - Packet sampling direction. Valid values: `rx`, `tx`, `both`.

* `sflow_counter_interval` - sFlow sampling counter polling interval (0 - 255 sec).
* `sflow_sample_rate` - sFlow sampler sample rate (0 - 99999 p/sec).
* `sflow_sampler` - Enable/disable sFlow protocol on this interface. Valid values: `disabled`, `enabled`.

* `status` - Switch port admin status: up or down. Valid values: `down`, `up`.

* `sticky_mac` - Enable or disable sticky-mac on the interface. Valid values: `disable`, `enable`.

* `stp_bpdu_guard` - Enable/disable STP BPDU guard on this interface. Valid values: `disabled`, `enabled`.

* `stp_bpdu_guard_timeout` - BPDU Guard disabling protection (0 - 120 min).
* `stp_root_guard` - Enable/disable STP root guard on this interface. Valid values: `disabled`, `enabled`.

* `stp_state` - Enable/disable Spanning Tree Protocol (STP) on this interface. Valid values: `disabled`, `enabled`.

* `trunk_member` - Trunk member.
* `type` - Interface type: physical or trunk port. Valid values: `physical`, `trunk`.

* `untagged_vlans` - Configure switch port untagged vlans
* `vlan` - Assign switch ports to a VLAN.

The `dhcp_snoop_option82_override` block supports:

* `circuit_id` - Circuit ID string.
* `remote_id` - Remote ID string.
* `vlan_name` - DHCP snooping option 82 VLAN.

The `route_offload_router` block supports:

* `router_ip` - Router IP address.
* `vlan_name` - VLAN name.

The `vlan` block supports:

* `assignment_priority` - 802.1x Radius (Tunnel-Private-Group-Id) VLANID assign-by-name priority. A smaller value has a higher priority.
* `vlan_name` - VLAN name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{switch_id}}.

## Import

ObjectSwitchController ManagedSwitch can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_managedswitch.labelname {{switch_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
