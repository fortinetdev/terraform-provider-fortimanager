---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_radius_dynamic_mapping_accountingserver"
description: |-
  Additional accounting servers.
---

# fortimanager_object_user_radius_dynamic_mapping_accountingserver
Additional accounting servers.

~> This resource is a sub resource for variable `accounting_server` of resource `fortimanager_object_user_radius_dynamic_mapping`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `radius` - Radius.
* `dynamic_mapping_name` - Dynamic Mapping Name.
* `dynamic_mapping_vdom` - Dynamic Mapping Vdom.

* `fosid` - ID (0 - 4294967295).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `port` - RADIUS accounting port number.
* `secret` - Secret key.
* `server` - {&lt;name_str|ip_str&gt;} Server CN domain name or IP.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectUser RadiusDynamicMappingAccountingServer can be imported using any of these accepted formats:
```
Set import_options = ["radius=YOUR_VALUE", "dynamic_mapping_name=YOUR_VALUE", "dynamic_mapping_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_radius_dynamic_mapping_accountingserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
