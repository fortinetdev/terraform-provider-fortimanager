---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_system_global"
description: |-
  Configure global attributes.
---

# fortimanager_systemp_system_global
Configure global attributes.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `admin_https_redirect` - Enable/disable redirection of HTTP administration access to HTTPS. Valid values: `disable`, `enable`.

* `admin_port` - Administrative access port for HTTP. (1 - 65535, default = 80).
* `admin_scp` - Enable/disable using SCP to download the system configuration. You can use SCP as an alternative method for backing up the configuration. Valid values: `disable`, `enable`.

* `admin_sport` - Administrative access port for HTTPS. (1 - 65535, default = 443).
* `admin_ssh_port` - Administrative access port for SSH. (1 - 65535, default = 22).
* `admin_ssh_v1` - Enable/disable SSH v1 compatibility. Valid values: `disable`, `enable`.

* `admin_telnet_port` - Administrative access port for TELNET. (1 - 65535, default = 23).
* `admintimeout` - Number of minutes before an idle administrator session times out (5 - 480 minutes (8 hours), default = 5). A shorter idle timeout is more secure.
* `gui_ipv6` - Enable/disable IPv6 settings on the GUI. Valid values: `disable`, `enable`.

* `gui_lines_per_page` - Number of lines to display per page for web administration.
* `gui_theme` - Color scheme for the administration GUI. Valid values: `blue`, `green`, `melongene`, `red`, `mariner`, `neutrino`.

* `language` - GUI display language. Valid values: `english`, `simch`, `japanese`, `korean`, `spanish`, `trach`, `french`, `portuguese`.

* `switch_controller` - Enable/disable switch controller feature. Switch controller allows you to manage FortiSwitch from the FortiGate itself. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Systemp SystemGlobal can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_system_global.labelname SystempSystemGlobal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
