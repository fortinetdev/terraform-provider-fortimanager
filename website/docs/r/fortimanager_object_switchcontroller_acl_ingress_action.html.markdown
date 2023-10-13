---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_acl_ingress_action"
description: |-
  ACL actions.
---

# fortimanager_object_switchcontroller_acl_ingress_action
ACL actions.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `ingress` - Ingress.

* `fmgcount` - Enable/disable count. Valid values: `disable`, `enable`.

* `drop` - Enable/disable drop. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSwitchController AclIngressAction can be imported using any of these accepted formats:
```
Set import_options = ["ingress=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_acl_ingress_action.labelname ObjectSwitchControllerAclIngressAction
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
