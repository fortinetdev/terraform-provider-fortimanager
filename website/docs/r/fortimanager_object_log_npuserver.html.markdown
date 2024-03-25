---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_log_npuserver"
description: |-
  Configure all the log servers and create the server groups.
---

# fortimanager_object_log_npuserver
Configure all the log servers and create the server groups.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `server_group`: `fortimanager_object_log_npuserver_servergroup`
>- `server_info`: `fortimanager_object_log_npuserver_serverinfo`



## Example Usage

```hcl
resource "fortimanager_object_log_npuserver" "trname" {
  log_processing = "no-drop"
  server_info {
    id          = 2
    dest_port   = 60
    ip_family   = "v4"
    ipv4_server = "34.5.6.9"
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `enforce_seq_order` - sw session netflow logs will be delivered in strict order if the option is enabled. Please do not switch the option while sw traffic is passing through. Valid values: `disable`, `enable`.

* `log_processing` - configure log processed by host to drop or no drop. Valid values: `may-drop`, `no-drop`.

* `log_processor` - configure the log module. Valid values: `hardware`, `host`.

* `netflow_ver` - configure the netfow verson. Valid values: `v9`, `v10`.

* `server_group` - Server-Group. The structure of `server_group` block is documented below.
* `server_info` - Server-Info. The structure of `server_info` block is documented below.
* `syslog_facility` - configure the syslog facility.
* `syslog_severity` - configure the syslog severity.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server_group` block supports:

* `group_name` - server group name.
* `log_format` - Set the log format Valid values: `syslog`, `netflow`.

* `log_gen_event` - Enable/disbale generating event for Per-Mapping log Valid values: `disable`, `enable`.

* `log_mode` - Set the log mode Valid values: `per-session`, `per-nat-mapping`, `per-session-ending`.

* `log_tx_mode` - Configure log transmit mode. Valid values: `multicast`, `roundrobin`.

* `log_user_info` - Enable/disbale logging user information. Valid values: `disable`, `enable`.

* `server_number` - server number in this group.
* `server_start_id` - the start id of the continuous server series in this group,[1,16].
* `sw_log_flags` - Set flags for software logging via driver. Valid values: `tcp-udp-only`, `enable-all-log`, `disable-all-log`.


The `server_info` block supports:

* `dest_port` - set the dest port for the log packet
* `id` - server id.
* `ip_family` - set the version the IP address Valid values: `v4`, `v6`.

* `ipv4_server` - set the IPv4 address for the log server
* `ipv6_server` - set the IPv6 address for the log server
* `log_transport` - set transport protocol Valid values: `udp`, `tcp`.

* `source_port` - set the source port for the log packet
* `template_tx_timeout` - set the template tx timeout
* `vdom` - Interface connected to the log server is in this virtual domain (VDOM).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectLog NpuServer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_log_npuserver.labelname ObjectLogNpuServer
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
