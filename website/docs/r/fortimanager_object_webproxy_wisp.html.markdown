---
subcategory: "ObjectWeb-Proxy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webproxy_wisp"
description: |-
  Configure Wireless Internet service provider (WISP) servers.
---

# fortimanager_object_webproxy_wisp
Configure Wireless Internet service provider (WISP) servers.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `max_connections` - Maximum number of web proxy WISP connections (4 - 4096, default = 64).
* `name` - Server name.
* `outgoing_ip` - WISP outgoing IP address.
* `server_ip` - WISP server IP address.
* `server_port` - WISP server port (1 - 65535, default = 15868).
* `timeout` - Period of time before WISP requests time out (1 - 15 sec, default = 5).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWebProxy Wisp can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webproxy_wisp.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
