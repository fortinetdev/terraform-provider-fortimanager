---
subcategory: "Object Web-Proxy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webproxy_forwardserver"
description: |-
  Configure forward-server addresses.
---

# fortimanager_object_webproxy_forwardserver
Configure forward-server addresses.

## Example Usage

```hcl
resource "fortimanager_object_webproxy_forwardserver" "trname" {
  addr_type          = "ip"
  comment            = "This is a Terraform example"
  healthcheck        = "disable"
  ip                 = "192.168.1.1"
  monitor            = "http://www.google.com"
  name               = "terr-web-proxy-forward-server"
  password           = ["fortinet"]
  port               = 3128
  server_down_option = "block"
  username           = "admin"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `addr_type` - Address type of the forwarding proxy server: IP or FQDN. Valid values: `fqdn`, `ip`.

* `comment` - Comment.
* `fqdn` - Forward server Fully Qualified Domain Name (FQDN).
* `healthcheck` - Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.

* `ip` - Forward proxy server IP address.
* `monitor` - URL for forward server health check monitoring (default = http://www.google.com).
* `name` - Server name.
* `password` - HTTP authentication password.
* `port` - Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
* `server_down_option` - Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.

* `username` - HTTP authentication user name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWebProxy ForwardServer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webproxy_forwardserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
