---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_accessproxy6_apigateway_realservers"
description: |-
  Select the real servers that this Access Proxy will distribute traffic to.
---

# fortimanager_object_firewall_accessproxy6_apigateway_realservers
Select the real servers that this Access Proxy will distribute traffic to.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `access_proxy6` - Access Proxy6.
* `api_gateway` - Api Gateway.

* `addr_type` - Type of address. Valid values: `fqdn`, `ip`.

* `address` - Address or address group of the real server.
* `domain` - Wildcard domain name of the real server.
* `external_auth` - Enable/disable use of external browser as user-agent for SAML user authentication. Valid values: `disable`, `enable`.

* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`.

* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping`, `http`, `tcp-connect`.

* `holddown_interval` - Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds). Valid values: `disable`, `enable`.

* `http_host` - HTTP server domain name in HTTP header.
* `fosid` - Real server ID.
* `ip` - IP address of the real server.
* `mappedport` - Port for communicating with the real server.
* `port` - Port for communicating with the real server.
* `ssh_client_cert` - Set access-proxy SSH client certificate profile.
* `ssh_host_key` - One or more server host key.
* `ssh_host_key_validation` - Enable/disable SSH real server host key validation. Valid values: `disable`, `enable`.

* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.

* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `disable`, `enable`.

* `tunnel_encryption` - Tunnel encryption. Valid values: `disable`, `enable`.

* `type` - TCP forwarding server type. Valid values: `tcp-forwarding`, `ssh`.

* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall AccessProxy6ApiGatewayRealservers can be imported using any of these accepted formats:
```
Set import_options = ["access_proxy6=YOUR_VALUE", "api_gateway=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_accessproxy6_apigateway_realservers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
