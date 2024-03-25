---
subcategory: "Object CLI"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_cli_template"
description: |-
  ObjectCli Template
---

# fortimanager_object_cli_template
ObjectCli Template

## Example Usage

```hcl
resource "fortimanager_object_cli_template" "trname" {
  description = "This is a Terraform example"
  name        = "terr-cli-template"
  script      = "terr-script"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `description` - Description.
* `modification_time` - Modification-Time.
* `name` - Name.
* `position` - Position. Valid values: `post-vdom-copy`, `prep-vdom-copy`.

* `provision` - Provision. Valid values: `disable`, `enable`.

* `script` - Script.
* `type` - Type. Valid values: `cli`, `jinja`.

* `variables` - Variables.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCli Template can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_cli_template.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
