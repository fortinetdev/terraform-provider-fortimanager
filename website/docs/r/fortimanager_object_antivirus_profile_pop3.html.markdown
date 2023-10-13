---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_antivirus_profile_pop3"
description: |-
  Configure POP3 AntiVirus options.
---

# fortimanager_object_antivirus_profile_pop3
Configure POP3 AntiVirus options.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol. Valid values: `disable`, `enable`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.

* `external_blocklist` - Enable external-blocklist. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable/disable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `options` - Enable/disable POP3 AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `file-filter`, `quarantine`, `avquery`, `avmonitor`.

* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectAntivirus ProfilePop3 can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_antivirus_profile_pop3.labelname ObjectAntivirusProfilePop3
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
