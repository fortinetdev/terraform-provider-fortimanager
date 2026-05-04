---
subcategory: "Object ZTNA"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ztna_webproxy_apigateway6_realservers"
description: |-
  Select the real servers that this Access Proxy will distribute traffic to.
---

# fortimanager_object_ztna_webproxy_apigateway6_realservers
Select the real servers that this Access Proxy will distribute traffic to.

~> This resource is a sub resource for variable `realservers` of resource `fortimanager_object_ztna_webproxy_apigateway6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `web_proxy` - Web Proxy.
* `api_gateway6` - Api Gateway6.

* `addr_type` - Type of address. Valid values: `fqdn`, `ip`.

* `address` - Address or address group of the real server.
* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`.

* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping`, `http`, `tcp-connect`.

* `holddown_interval` - Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds). Valid values: `disable`, `enable`.

* `http_host` - HTTP server domain name in HTTP header.
* `fosid` - Real server ID.
* `ip` - IPv6 address of the real server.
* `port` - Port for communicating with the real server.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.

* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `disable`, `enable`.

* `verify_cert` - Enable/disable certificate verification of the real server. Valid values: `disable`, `enable`.

* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectZtna WebProxyApiGateway6Realservers can be imported using any of these accepted formats:
```
Set import_options = ["web_proxy=YOUR_VALUE", "api_gateway6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_ztna_webproxy_apigateway6_realservers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
