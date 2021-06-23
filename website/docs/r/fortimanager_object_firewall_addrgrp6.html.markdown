---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_addrgrp6"
description: |-
  Configure IPv6 address groups.
---

# fortimanager_object_firewall_addrgrp6
Configure IPv6 address groups.

## Example Usage

```hcl
resource "fortimanager_object_firewall_address6" "trname" {
  color     = 1
  comment   = "This is a Terraform example"
  country   = "US"
  end_ip    = "2001:192:168:1::10"
  end_mac   = "00:00:00:00:00:00"
  host      = "::"
  host_type = "any"
  ip6       = "::/0"
  name      = "terr-firewall-address6"
  start_ip  = "2001:192:168:1::1"
  start_mac = "00:00:00:00:00:00"
  type      = "iprange"
}

resource "fortimanager_object_firewall_addrgrp6" "trname" {
  comment = "terraform-comment"
  member  = "terr-firewall-address6"
  name    = "terraform-addrgrp6"
  depends_on = [
    fortimanager_object_firewall_address6.trname
  ]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_image_base64` - _Image-Base64.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `global_object` - Global Object.
* `member` - Address objects contained within the group.
* `name` - IPv6 address group name.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address group6 visibility in the GUI. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_image_base64` - _Image-Base64.
* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `fabric_object` - Fabric-Object. Valid values: `disable`, `enable`.

* `global_object` - Global-Object.
* `member` - Address objects contained within the group.
* `tags` - Tags.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address group6 visibility in the GUI. Valid values: `disable`, `enable`.


The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Addrgrp6 can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_addrgrp6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
