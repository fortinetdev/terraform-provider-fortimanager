---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_log_npuserver_serverinfo"
description: |-
  configure server info.
---

# fortimanager_object_log_npuserver_serverinfo
configure server info.

~> This resource is a sub resource for variable `server_info` of resource `fortimanager_object_log_npuserver`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_log_npuserver_serverinfo" "trname" {
  fosid       = 1
  dest_port   = 60
  ip_family   = "v4"
  ipv4_server = "34.5.6.9"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `dest_port` - set the dest port for the log packet
* `fosid` - server id.
* `ip_family` - set the version the IP address Valid values: `v4`, `v6`.

* `ipv4_server` - set the IPv4 address for the log server
* `ipv6_server` - set the IPv6 address for the log server
* `log_transport` - set transport protocol Valid values: `udp`, `tcp`.

* `source_port` - set the source port for the log packet
* `template_tx_timeout` - set the template tx timeout
* `vdom` - Interface connected to the log server is in this virtual domain (VDOM).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectLog NpuServerServerInfo can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_log_npuserver_serverinfo.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
