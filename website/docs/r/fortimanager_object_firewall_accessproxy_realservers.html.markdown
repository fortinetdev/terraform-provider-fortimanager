---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_accessproxy_realservers"
description: |-
  Select the SSL real servers that this Access Proxy will distribute traffic to.
---

# fortimanager_object_firewall_accessproxy_realservers
Select the SSL real servers that this Access Proxy will distribute traffic to.

~> This resource is a sub resource for variable `realservers` of resource `fortimanager_object_firewall_accessproxy`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_accessproxy_realservers" "trname" {
  access_proxy = fortimanager_object_firewall_accessproxy.trname.name
  fosid        = 1
  ip           = "12.3.4.5"
  port         = 65
  weight       = 23
  depends_on   = [fortimanager_object_firewall_accessproxy.trname]
}

resource "fortimanager_object_firewall_accessproxy" "trname" {
  name = "terr-accessproxy"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `access_proxy` - Access Proxy.

* `fosid` - Real server ID.
* `ip` - IP address of the real server.
* `port` - Port for communicating with the real server.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.

* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall AccessProxyRealservers can be imported using any of these accepted formats:
```
Set import_options = ["access_proxy=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_accessproxy_realservers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
