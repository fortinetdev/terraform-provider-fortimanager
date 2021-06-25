---
subcategory: "Object Application"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_application_group"
description: |-
  Configure firewall application groups.
---

# fortimanager_object_application_group
Configure firewall application groups.

## Example Usage

```hcl
resource "fortimanager_object_application_group" "trname" {
  comment = "terraform-tefv-comment"
  name    = "terraform-tefv"
  type    = "application"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `application` - Application ID list.
* `behavior` - Application behavior filter. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `category` - Application category ID list.
* `comment` - Comment
* `name` - Application group name.
* `popularity` - Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `protocols` - Application protocol filter. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `risk` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `technology` - Application technology filter. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `type` - Application group type. Valid values: `application`, `category`.

* `vendor` - Application vendor filter. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectApplication Group can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_application_group.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
