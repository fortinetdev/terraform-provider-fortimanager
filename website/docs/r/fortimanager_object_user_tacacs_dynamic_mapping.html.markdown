---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_tacacs_dynamic_mapping"
description: |-
  Configure TACACS+ server entries.
---

# fortimanager_object_user_tacacs_dynamic_mapping
Configure TACACS+ server entries.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_user_tacacs`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `tacacs` - Tacacs.

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `authen_type` - Allowed authentication protocols/methods. Valid values: `auto`, `ascii`, `pap`, `chap`, `mschap`.

* `authorization` - Enable/disable TACACS+ authorization. Valid values: `disable`, `enable`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `key` - Key to access the primary server.
* `port` - Port number of the TACACS+ server.
* `secondary_key` - Key to access the secondary server.
* `secondary_server` - Secondary TACACS+ server CN domain name or IP address.
* `server` - Primary TACACS+ server CN domain name or IP address.
* `source_ip` - source IP for communications to TACACS+ server.
* `status_ttl` - Time for which server reachability is cached so that when a server is unreachable, it will not be retried for at least this period of time (0 = cache disabled, default = 300).
* `tertiary_key` - Key to access the tertiary server.
* `tertiary_server` - Tertiary TACACS+ server CN domain name or IP address.
* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectUser TacacsDynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["tacacs=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_tacacs_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
