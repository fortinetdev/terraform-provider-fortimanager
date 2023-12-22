---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_sdnconnector_routetable"
description: |-
  Configure Azure route table.
---

# fortimanager_object_system_sdnconnector_routetable
Configure Azure route table.

~> This resource is a sub resource for variable `route_table` of resource `fortimanager_object_system_sdnconnector`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`route`: `fortimanager_object_system_sdnconnector_routetable_route`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `sdn_connector` - Sdn Connector.

* `name` - Route table name.
* `resource_group` - Resource group of Azure route table.
* `route` - Route. The structure of `route` block is documented below.
* `subscription_id` - Subscription ID of Azure route table.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `route` block supports:

* `name` - Route name.
* `next_hop` - Next hop address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem SdnConnectorRouteTable can be imported using any of these accepted formats:
```
Set import_options = ["sdn_connector=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_sdnconnector_routetable.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
