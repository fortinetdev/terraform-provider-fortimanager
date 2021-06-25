---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extendercontroller_template"
description: |-
  ObjectExtenderController Template
---

# fortimanager_object_extendercontroller_template
ObjectExtenderController Template Applies to `FortiManager >= 7.0 and Controlled FortiOS >= 6.4`.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `dataplan` - Dataplan.
* `description` - Description.
* `modem1_ifname` - Modem1_Ifname.
* `modem1_sim_profile` - Modem1_Sim_Profile.
* `modem2_ifname` - Modem2_Ifname.
* `modem2_sim_profile` - Modem2_Sim_Profile.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectExtenderController Template can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extendercontroller_template.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
