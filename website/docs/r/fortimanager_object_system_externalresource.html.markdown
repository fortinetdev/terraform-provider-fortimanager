---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_externalresource"
description: |-
  Configure external resource.
---

# fortimanager_object_system_externalresource
Configure external resource.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fortimanager_object_system_externalresource_dynamic_mapping`



## Example Usage

```hcl
resource "fortimanager_object_system_externalresource" "trname" {
  name     = "terr-externalresource"
  category = 192
  comment  = "This is a Terraform example"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `address_comment_field` - JSON Path to address description in generic address entry.
* `address_data_field` - JSON Path to address data in generic address entry.
* `address_name_field` - JSON Path to address name in generic address entry.
* `category` - User resource category.
* `client_cert` - Client certificate name.
* `client_cert_auth` - Enable/disable using client certificate for TLS authentication. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `name` - External resource name.
* `namespace` - Generic external connector address namespace.
* `object_array_path` - JSON Path to array of generic addresses in resource.
* `password` - HTTP basic authentication password.
* `refresh_rate` - Time interval to refresh external resource (1 - 43200 min, default = 5 min).
* `resource` - URI of external resource.
* `server_identity_check` - Certificate verification option. Valid values: `none`, `basic`, `full`.

* `source_ip` - Source IPv4 address used to communicate with server.
* `status` - Enable/disable user resource. Valid values: `disable`, `enable`.

* `type` - User resource type. Valid values: `category`, `address`, `domain`, `malware`.

* `update_method` - External resource update method. Valid values: `feed`, `push`.

* `user_agent` - Override HTTP User-Agent header used when retrieving this external resource.
* `username` - HTTP basic authentication user name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `source_ip` - Source-Ip.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem ExternalResource can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_externalresource.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
