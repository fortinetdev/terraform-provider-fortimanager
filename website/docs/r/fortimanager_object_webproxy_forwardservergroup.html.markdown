---
subcategory: "Object Web-Proxy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webproxy_forwardservergroup"
description: |-
  Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.
---

# fortimanager_object_webproxy_forwardservergroup
Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.

## Example Usage

```hcl
resource "fortimanager_object_webproxy_forwardservergroup" "trname" {
  affinity          = "enable"
  group_down_option = "block"
  ldb_method        = "active-passive"
  name              = "terr-web-proxy-forward-server-group"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `affinity` - Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `disable`, `enable`.

* `group_down_option` - Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.

* `ldb_method` - Load balance method: weighted or least-session. Valid values: `weighted`, `least-session`, `active-passive`.

* `name` - Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
* `server_list` - Server-List. The structure of `server_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server_list` block supports:

* `name` - Forward server name.
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWebProxy ForwardServerGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webproxy_forwardservergroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
