---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_vipgrp64"
description: |-
  Configure IPv6 to IPv4 virtual IP groups.
---

# fortimanager_object_firewall_vipgrp64
Configure IPv6 to IPv4 virtual IP groups.

## Example Usage

```hcl
resource "fortimanager_object_firewall_vipgrp64" "trname" {
  color    = 5
  comments = "This is a Terraform example"
  member   = ["terr-firewall-vip64"]
  name     = "terr-firewall-vipgrp64"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comments` - Comment.
* `member` - Member VIP objects of the group (Separate multiple objects with a space).
* `name` - VIP64 group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Vipgrp64 can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_vipgrp64.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
