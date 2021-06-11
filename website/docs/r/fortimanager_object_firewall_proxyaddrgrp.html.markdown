---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_proxyaddrgrp"
description: |-
  Configure web proxy address group.
---

# fortimanager_object_firewall_proxyaddrgrp
Configure web proxy address group.

## Example Usage

```hcl
resource "fortimanager_object_firewall_proxyaddress" "trname" {
  case_sensitivity = "disable"
  color            = 1
  comment          = "This is a Terraform example"
  host             = "all"
  name             = "terr-firewall-proxy-address"
  type             = "ua"
  ua               = ["chrome"]
}

resource "fortimanager_object_firewall_proxyaddrgrp" "trname" {
  comment = "This is a Terraform example"
  member  = "terr-firewall-proxy-address"
  name    = "terr-firewall-proxy-addrgrp"
  type    = "dst"
  depends_on = [
    fortimanager_object_firewall_proxyaddress.trname
  ]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_image_base64` - _Image-Base64.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `comment` - Optional comments.
* `member` - Members of address group.
* `name` - Address group name.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `type` - Source or destination address group type. Valid values: `src`, `dst`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ProxyAddrgrp can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_proxyaddrgrp.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
