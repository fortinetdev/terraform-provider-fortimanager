---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_pfcp"
description: |-
  Configure PFCP.
---

# fortimanager_object_firewall_pfcp
Configure PFCP.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `denied_log` - Enable/disable logging denied PFCP packets. Valid values: `disable`, `enable`.

* `forwarded_log` - Enable/disable logging forwarded PFCP packets. Valid values: `disable`, `enable`.

* `invalid_reserved_field` - Allow or deny invalid reserved field in PFCP header packets. Valid values: `deny`, `allow`.

* `log_freq` - Logging frequency of PFCP packets.
* `max_message_length` - Maximum message length.
* `message_filter` - PFCP message filter.
* `min_message_length` - Minimum message length.
* `monitor_mode` - PFCP monitor mode. Valid values: `disable`, `enable`, `vdom`.

* `name` - PFCP profile name.
* `pfcp_timeout` - Set PFCP timeout (in seconds).
* `traffic_count_log` - Enable/disable logging session traffic counter. Valid values: `disable`, `enable`.

* `unknown_version` - Allow or deny unknown version packets. Valid values: `deny`, `allow`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Pfcp can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_pfcp.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
