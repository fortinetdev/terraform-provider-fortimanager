---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ztna_trafficforwardproxy_urlroute"
description: |-
  ObjectZtna TrafficForwardProxyUrlRoute
---

# fortimanager_object_ztna_trafficforwardproxy_urlroute
ObjectZtna TrafficForwardProxyUrlRoute

~> This resource is a sub resource for variable `url_route` of resource `fortimanager_object_ztna_trafficforwardproxy`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `traffic_forward_proxy` - Traffic Forward Proxy.

* `name` - Name.
* `service_connector` - Service-Connector.
* `url_pattern` - Url-Pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectZtna TrafficForwardProxyUrlRoute can be imported using any of these accepted formats:
```
Set import_options = ["traffic_forward_proxy=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_ztna_trafficforwardproxy_urlroute.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
