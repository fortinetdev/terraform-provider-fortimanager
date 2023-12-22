---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fmg_fabric_authorization_template"
description: |-
  ObjectFmg FabricAuthorizationTemplate
---

# fortimanager_object_fmg_fabric_authorization_template
ObjectFmg FabricAuthorizationTemplate

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`platforms`: `fortimanager_object_fmg_fabric_authorization_template_platforms`



## Example Usage

```hcl
resource "fortimanager_object_fmg_fabric_authorization_template" "trname" {
  name              = "terr-template"
  description       = "This is a Terraform example"
  switch_controller = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `description` - Description.
* `extender_controller` - Extender-Controller. Valid values: `disable`, `enable`.

* `name` - Name.
* `platforms` - Platforms. The structure of `platforms` block is documented below.
* `switch_controller` - Switch-Controller. Valid values: `disable`, `enable`.

* `wireless_controller` - Wireless-Controller. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `platforms` block supports:

* `count` - Count.
* `extension_type` - Extension-Type. Valid values: `wan-extension`, `lan-extension`.

* `fortilink` - Fortilink.
* `prefix` - Prefix.
* `type` - Type. Valid values: `ap`, `extender`, `switch`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFmg FabricAuthorizationTemplate can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fmg_fabric_authorization_template.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
