---
subcategory: "Object Webfilter"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_ftgdlocalrating"
description: |-
  Configure local FortiGuard Web Filter local ratings.
---

# fortimanager_object_webfilter_ftgdlocalrating
Configure local FortiGuard Web Filter local ratings.

## Example Usage

```hcl
resource "fortimanager_object_webfilter_ftgdlocalrating" "trname" {
  comment = "This is a Terraform example"
  rating  = "0"
  status  = "enable"
  url     = "terr-webfilter-ftgd-local-rating"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `rating` - Local rating.
* `status` - Enable/disable local rating. Valid values: `disable`, `enable`.

* `url` - URL to rate locally.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{url}}.

## Import

ObjectWebfilter FtgdLocalRating can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_ftgdlocalrating.labelname {{url}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
