---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_accessproxyvirtualhost"
description: |-
  Configure Access Proxy virtual hosts.
---

# fortimanager_object_firewall_accessproxyvirtualhost
Configure Access Proxy virtual hosts.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `host` - The host name.
* `host_type` - Type of host pattern. Valid values: `sub-string`, `wildcard`.

* `name` - Virtual host name.
* `replacemsg_group` - Access-proxy-virtual-host replacement message override group.
* `ssl_certificate` - SSL certificate for this host.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall AccessProxyVirtualHost can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_accessproxyvirtualhost.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
