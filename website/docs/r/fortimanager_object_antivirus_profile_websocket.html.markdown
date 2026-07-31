---
subcategory: "Object Antivirus"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_antivirus_profile_websocket"
description: |-
  Configure WEBSOCKET AntiVirus options.
---

# fortimanager_object_antivirus_profile_websocket
Configure WEBSOCKET AntiVirus options.

~> This resource is a sub resource for variable `websocket` of resource `fortimanager_object_antivirus_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable/disable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `external_blocklist` - Enable/disable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable/disable scanning of files by FortiNDR. Valid values: `disable`, `monitor`, `block`.

* `fortisandbox` - Enable/disable scanning of files by FortiSandbox. Valid values: `disable`, `monitor`, `block`.

* `malware_stream` - Enable/disable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable/disable virus outbreak prevention service. Valid values: `disable`, `monitor`, `block`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectAntivirus ProfileWebsocket can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_antivirus_profile_websocket.labelname ObjectAntivirusProfileWebsocket
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
