---
subcategory: "Object CLI"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_cli_templategroup"
description: |-
  ObjectCli TemplateGroup
---

# fortimanager_object_cli_templategroup
ObjectCli TemplateGroup

## Example Usage

```hcl
resource "fortimanager_object_cli_template" "trname" {
  description = "This is a Terraform example"
  name        = "terr-cli-template"
  script      = "terr-script"
}

resource "fortimanager_object_cli_templategroup" "trname" {
  description = "This is a Terraform example"
  member      = ["terr-cli-template"]
  name        = "terr-cli-tplt-grp"
  depends_on = [
    fortimanager_object_cli_template.trname
  ]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `description` - Description.
* `member` - Member.
* `modification_time` - Modification-Time.
* `name` - Name.
* `option` - Option. Valid values: `sdwan-overlay`, `sdwan-manager`.

* `variables` - Variables.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCli TemplateGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_cli_templategroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
