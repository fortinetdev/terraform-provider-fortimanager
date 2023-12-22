---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_sdnconnector_routetable_route"
description: |-
  Configure Azure route.
---

# fortimanager_object_system_sdnconnector_routetable_route
Configure Azure route.

~> This resource is a sub resource for variable `route` of resource `fortimanager_object_system_sdnconnector_routetable`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `sdn_connector` - Sdn Connector.
* `route_table` - Route Table.

* `name` - Route name.
* `next_hop` - Next hop address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem SdnConnectorRouteTableRoute can be imported using any of these accepted formats:
```
Set import_options = ["sdn_connector=YOUR_VALUE", "route_table=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_sdnconnector_routetable_route.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
