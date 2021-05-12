---
subcategory: "Securityconsole"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_securityconsole_import_dev_objs"
description: |-
  Import objects from device to ADOM, or from ADOM to Global.
---

# fortimanager_securityconsole_import_dev_objs
Import objects from device to ADOM, or from ADOM to Global.

## Argument Reference


The following arguments are supported:


* `add_mappings` - Automatically add required dynamic mappings for the device during the search stages. When used in policy_search action, add dynamic interface and zone mappings; in obj_search action, add ADOM object mappings. This attribute is not available for the do action. Valid values: `disable`, `enable`.

* `fmgadom` - Source ADOM name.
* `dst_name` - Name of the policy package where the objects are to be imported. If the package does not already exist in the database, a new one will be created.
* `dst_parent` - Path to the folder for the target package. If the package is to be placed in root, leave this field blank.
* `if_all_objs` - If_All_Objs. Valid values: `none`, `all`, `filter`.

* `if_all_policy` - If_All_Policy. Valid values: `disable`, `enable`.

* `import_action` - do - Perform the policy and object import. policy_search - Preprocess and scan through device database to gather information about policies that need to be imported. Can automatically add interface and zone mappings if add_mappings is enabled. obj_search - Preprocess and scan through device database to collect objects that are required to be imported. Can automatically add object mappings if add_mappings is enabled. Valid values: `do`, `policy_search`, `obj_search`.

* `name` - Source device name.
* `position` - Position. Valid values: `bottom`, `top`.

* `vdom` - Vdom.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Securityconsole ImportDevObjs can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_securityconsole_import_dev_objs.labelname SecurityconsoleImportDevObjs
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
