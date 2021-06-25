---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_ldbmonitor"
description: |-
  Configure server load balancing health monitors.
---

# fortimanager_object_firewall_ldbmonitor
Configure server load balancing health monitors.

## Example Usage

```hcl
resource "fortimanager_object_firewall_ldbmonitor" "trname" {
  http_max_redirects = 0
  interval           = 10
  name               = "dfsd"
  port               = 0
  retry              = 3
  src_ip             = "0.0.0.0"
  timeout            = 2
  type               = "https"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `dns_match_ip` - Response IP expected from DNS server. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `dns_protocol` - Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `http_get` - URL used to send a GET request to check the health of an HTTP server.
* `http_match` - String to match the value expected in response to an HTTP-GET request.
* `http_max_redirects` - The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
* `interval` - Time between health checks (5 - 65635 sec, default = 10).
* `name` - Monitor name.
* `port` - Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
* `retry` - Number health check attempts before the server is considered down (1 - 255, default = 3).
* `src_ip` - Source IP for ldb-monitor. (`ver Controlled FortiOS >= 6.4`)
* `timeout` - Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
* `type` - Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP | HTTPS). Valid values: `ping`, `tcp`, `http`, `passive-sip`, `https`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall LdbMonitor can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_ldbmonitor.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
