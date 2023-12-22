---
subcategory: "System Global"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_ha"
description: |-
  HA configuration.
---

# fortimanager_system_ha
HA configuration.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`monitored_interfaces`: `fortimanager_system_ha_monitoredinterfaces`
`monitored_ips`: `fortimanager_system_ha_monitoredips`
`peer`: `fortimanager_system_ha_peer`



## Example Usage

```hcl
resource "fortimanager_system_ha" "trname" {
  clusterid  = "2"
  file_quota = "2048"
  mode       = "standalone"
  password   = ["fortinet"]
}
```

## Argument Reference


The following arguments are supported:


* `clusterid` - Cluster ID range (1 - 64).
* `failover_mode` - HA failover mode. manual - Manual Failove vrrp - Use VRRP Valid values: `manual`, `vrrp`.

* `file_quota` - File quota in MB (2048 - 20480).
* `hb_interval` - Heartbeat interval (1 - 255).
* `hb_lost_threshold` - Heartbeat lost threshold (1 - 255).
* `local_cert` - set the ha local certificate.
* `mode` - Mode. standalone - Standalone. primary - Primary. secondary - Secondary. Valid values: `standalone`, `primary`, `secondary`.

* `monitored_interfaces` - Monitored-Interfaces. The structure of `monitored_interfaces` block is documented below.
* `monitored_ips` - Monitored-Ips. The structure of `monitored_ips` block is documented below.
* `password` - Group password.
* `peer` - Peer. The structure of `peer` block is documented below.
* `priority` - Runtime priority [1 (lowest) - 253 (highest)]
* `unicast` - Use unitcast for VRRP message. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `vip` - Virtual IP.
* `vip_interface` - vip interface.
* `vrrp_adv_interval` - VRRP advert interval [1 - 30 seconnds]
* `vrrp_interface` - VRRP and vip interface.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `monitored_interfaces` block supports:

* `interface_name` - Interface name.

The `monitored_ips` block supports:

* `id` - Id.
* `interface` - Interface name.
* `ip` - IP address.

The `peer` block supports:

* `id` - Id.
* `ip` - IP address of peer.
* `ip6` - IP address (V6) of peer.
* `serial_number` - Serial number of peer.
* `status` - Peer admin status. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ha can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_ha.labelname SystemHa
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

