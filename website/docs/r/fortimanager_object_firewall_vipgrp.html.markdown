---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_vipgrp"
description: |-
  Configure IPv4 virtual IP groups.
---

# fortimanager_object_firewall_vipgrp
Configure IPv4 virtual IP groups.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fortimanager_object_firewall_vipgrp_dynamic_mapping`



## Example Usage

```hcl
resource "fortimanager_object_firewall_vip" "trname" {
  arp_reply              = "enable"
  color                  = 1
  comment                = "This is a Terraform example"
  extintf                = "any"
  extip                  = "192.168.1.1"
  http_redirect          = "disable"
  name                   = "terr-firewall-vip"
  nat_source_vip         = "disable"
  portforward            = "disable"
  ssl_client_fallback    = "enable"
  ssl_server_algorithm   = "client"
  ssl_server_max_version = "client"
  ssl_server_min_version = "client"
  type                   = "static-nat"
}

resource "fortimanager_object_firewall_vipgrp" "trname" {
  color     = 3
  comments  = "This is a Terraform example"
  interface = "any"
  member    = ["terr-firewall-vip"]
  name      = "terr-firewall-vipgrp"
  depends_on = [
    fortimanager_object_firewall_vip.trname
  ]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comments` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `interface` - Interface.
* `member` - Member VIP objects of the group (Separate multiple objects with a space).
* `name` - VIP group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comments` - Comment.
* `interface` - Interface.
* `member` - Member VIP objects of the group (Separate multiple objects with a space).
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Vipgrp can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_vipgrp.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
