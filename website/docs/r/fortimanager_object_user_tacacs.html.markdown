---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_tacacs"
description: |-
  Configure TACACS+ server entries.
---

# fortimanager_object_user_tacacs
Configure TACACS+ server entries.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fortimanager_object_user_tacacs_dynamic_mapping`



## Example Usage

```hcl
resource "fortimanager_object_user_tacacs" "trname" {
  authen_type             = "auto"
  authorization           = "disable"
  interface               = "port10"
  interface_select_method = "auto"
  key                     = ["fortinet"]
  name                    = "terr-user-tacacs"
  port                    = 49
  secondary_key           = ["fortinet"]
  server                  = "192.168.1.1"
  tertiary_key            = ["fortinet"]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `authen_type` - Allowed authentication protocols/methods. Valid values: `auto`, `ascii`, `pap`, `chap`, `mschap`.

* `authorization` - Enable/disable TACACS+ authorization. Valid values: `disable`, `enable`.

* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `key` - Key to access the primary server.
* `name` - TACACS+ server entry name.
* `port` - Port number of the TACACS+ server.
* `secondary_key` - Key to access the secondary server.
* `secondary_server` - Secondary TACACS+ server CN domain name or IP address.
* `server` - Primary TACACS+ server CN domain name or IP address.
* `source_ip` - source IP for communications to TACACS+ server.
* `status_ttl` - Time for which server reachability is cached so that when a server is unreachable, it will not be retried for at least this period of time (0 = cache disabled, default = 300).
* `tertiary_key` - Key to access the tertiary server.
* `tertiary_server` - Tertiary TACACS+ server CN domain name or IP address.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

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

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Tacacs can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_tacacs.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
