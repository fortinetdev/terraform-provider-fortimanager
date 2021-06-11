---
subcategory: "System Global"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_ha"
description: |-
  HA configuration.
---

# fortimanager_system_ha
HA configuration.

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
* `file_quota` - File quota in MB (2048 - 20480).
* `hb_interval` - Heartbeat interval (1 - 255).
* `hb_lost_threshold` - Heartbeat lost threshold (1 - 255).
* `local_cert` - set the ha local certificate.
* `mode` - Mode. standalone - Standalone. primary - Primary. secondary - Secondary. Valid values: `standalone`, `primary`, `secondary`.

* `password` - Group password.
* `peer` - Peer. The structure of `peer` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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

