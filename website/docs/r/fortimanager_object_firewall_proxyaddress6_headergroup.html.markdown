---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_proxyaddress6_headergroup"
description: |-
  ObjectFirewall ProxyAddress6HeaderGroup
---

# fortimanager_object_firewall_proxyaddress6_headergroup
ObjectFirewall ProxyAddress6HeaderGroup

~> This resource is a sub resource for variable `header_group` of resource `fortimanager_object_firewall_proxyaddress6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `proxy_address6` - Proxy Address6.

* `case_sensitivity` - Case-Sensitivity. Valid values: `disable`, `enable`.

* `header` - Header.
* `header_name` - Header-Name.
* `fosid` - Id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall ProxyAddress6HeaderGroup can be imported using any of these accepted formats:
```
Set import_options = ["proxy_address6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_proxyaddress6_headergroup.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
