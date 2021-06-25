---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_service_group"
description: |-
  Configure service groups.
---

# fortimanager_object_firewall_service_group
Configure service groups.

## Example Usage

```hcl
resource "fortimanager_object_firewall_service_group" "trname" {
  color   = 1
  comment = "terraform-comment"
  member  = "ALL"
  name    = "terraform-tefv-group"
  proxy   = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.
 (`ver Controlled FortiOS >= 6.4`)
* `global_object` - Global Object. (`ver Controlled FortiOS = 6.4`)
* `member` - Service objects contained within the group.
* `name` - Address group name.
* `proxy` - Enable/disable web proxy service group. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ServiceGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_service_group.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
