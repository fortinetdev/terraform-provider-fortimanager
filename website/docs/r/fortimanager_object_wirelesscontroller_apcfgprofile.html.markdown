---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_apcfgprofile"
description: |-
  Configure AP local configuration profiles.
---

# fortimanager_object_wirelesscontroller_apcfgprofile
Configure AP local configuration profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`command_list`: `fortimanager_object_wirelesscontroller_apcfgprofile_commandlist`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_apcfgprofile" "trname" {
  command_list {
    id           = 1
    name         = "terr-apcfgprofile"
    type         = "password"
    passwd_value = ["123"]
  }
  comment = "This is a Terraform example"
  name    = "terr-apcfgprofile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ac_ip` - IP address of the validation controller that AP must be able to join after applying AP local configuration.
* `ac_port` - Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
* `ac_timer` - Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
* `ac_type` - Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.

* `ap_family` - FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.

* `command_list` - Command-List. The structure of `command_list` block is documented below.
* `comment` - Comment.
* `name` - AP local configuration profile name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `command_list` block supports:

* `id` - Command ID.
* `name` - AP local configuration command name.
* `passwd_value` - AP local configuration command password value.
* `type` - The command type (default = non-password). Valid values: `non-password`, `password`.

* `value` - AP local configuration command value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController ApcfgProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_apcfgprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
