---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fmg_device_blueprint"
description: |-
  ObjectFmg DeviceBlueprint
---

# fortimanager_object_fmg_device_blueprint
ObjectFmg DeviceBlueprint

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auth_template` - Auth-Template.
* `cliprofs` - Cliprofs.
* `description` - Description.
* `dev_group` - Dev-Group.
* `enforce_device_config` - Enforce-Device-Config. Valid values: `disable`, `enable`.

* `folder` - Folder.
* `ha_config` - Ha-Config. Valid values: `disable`, `enable`.

* `ha_hbdev` - Ha-Hbdev.
* `ha_monitor` - Ha-Monitor.
* `ha_password` - Ha-Password.
* `linked_to_model` - Linked-To-Model. Valid values: `disable`, `enable`.

* `name` - Name.
* `pkg` - Pkg.
* `platform` - Platform.
* `port_provisioning` - Port-Provisioning.
* `prefer_img_ver` - Prefer-Img-Ver.
* `prerun_cliprof` - Prerun-Cliprof.
* `prov_type` - Prov-Type. Valid values: `none`, `templates`, `template-group`.

* `sdwan_management` - Sdwan-Management. Valid values: `disable`, `enable`.

* `split_switch_port` - Split-Switch-Port. Valid values: `disable`, `enable`.

* `template_group` - Template-Group.
* `templates` - Templates.
* `vm_log_disk` - Vm-Log-Disk. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFmg DeviceBlueprint can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fmg_device_blueprint.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
