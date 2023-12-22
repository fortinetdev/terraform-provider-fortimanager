---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extendercontroller_extenderprofile_lanextension_backhaul"
description: |-
  LAN extension backhaul tunnel configuration.
---

# fortimanager_object_extendercontroller_extenderprofile_lanextension_backhaul
LAN extension backhaul tunnel configuration.

~> This resource is a sub resource for variable `backhaul` of resource `fortimanager_object_extendercontroller_extenderprofile_lanextension`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_extendercontroller_extenderprofile_lanextension_backhaul" "trname" {
  extender_profile = fortimanager_object_extendercontroller_extenderprofile.trname.name
  name             = "default"
  port             = "port1"
  role             = "secondary"
  weight           = 23
  depends_on       = [fortimanager_object_extendercontroller_extenderprofile.trname]
}

resource "fortimanager_object_extendercontroller_extenderprofile_lanextension" "trname" {
  extender_profile = fortimanager_object_extendercontroller_extenderprofile.trname.name
  depends_on       = [fortimanager_object_extendercontroller_extenderprofile.trname]
}

resource "fortimanager_object_extendercontroller_extenderprofile" "trname" {
  name = "terr-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.

* `name` - FortiExtender LAN extension backhaul name
* `port` - FortiExtender uplink port. Valid values: `wan`, `lte1`, `lte2`, `port1`, `port2`, `port3`, `port4`, `port5`, `sfp`.

* `role` - FortiExtender uplink port. Valid values: `primary`, `secondary`.

* `weight` - WRR weight parameter


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectExtenderController ExtenderProfileLanExtensionBackhaul can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extendercontroller_extenderprofile_lanextension_backhaul.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
