---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_cloud_orchestaws"
description: |-
  ObjectCloud OrchestAws
---

# fortimanager_object_cloud_orchestaws
ObjectCloud OrchestAws

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `connector` - Connector.
* `name` - Name.
* `region_name` - Region-Name. Valid values: `us-east-1`, `us-east-2`, `us-west-1`, `us-west-2`, `eu-west-1`, `eu-west-2`, `eu-west-3`, `eu-north-1`, `eu-south-1`, `eu-south-2`, `eu-central-1`, `eu-central-2`, `ca-central-1`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-4`, `ap-south-1`, `ap-south-2`, `ap-northeast-1`, `ap-northeast-2`, `ap-northeast-3`, `af-south-1`, `me-central-1`, `me-south-1`, `sa-east-1`, `ap-east-1`, `us-gov-east-1`, `us-gov-west-1`.

* `template_configuration` - Template-Configuration.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCloud OrchestAws can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_cloud_orchestaws.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
