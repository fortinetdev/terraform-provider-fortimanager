---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_log_npuserver_servergroup"
description: |-
  create server group.
---

# fortimanager_object_log_npuserver_servergroup
create server group.

~> This resource is a sub resource for variable `server_group` of resource `fortimanager_object_log_npuserver`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `group_name` - server group name.
* `log_format` - Set the log format Valid values: `syslog`, `netflow`.

* `log_gen_event` - Enable/disbale generating event for Per-Mapping log Valid values: `disable`, `enable`.

* `log_mode` - Set the log mode Valid values: `per-session`, `per-nat-mapping`, `per-session-ending`.

* `log_tx_mode` - Configure log transmit mode. Valid values: `multicast`, `roundrobin`.

* `log_user_info` - Enable/disbale logging user information. Valid values: `disable`, `enable`.

* `server_number` - server number in this group.
* `server_start_id` - the start id of the continuous server series in this group,[1,16].
* `sw_log_flags` - Set flags for software logging via driver. Valid values: `tcp-udp-only`, `enable-all-log`, `disable-all-log`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{group_name}}.

## Import

ObjectLog NpuServerServerGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_log_npuserver_servergroup.labelname {{group_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
