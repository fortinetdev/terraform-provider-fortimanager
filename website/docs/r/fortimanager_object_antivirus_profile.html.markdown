---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_antivirus_profile"
description: |-
  Configure AntiVirus profiles.
---

# fortimanager_object_antivirus_profile
Configure AntiVirus profiles.

## Example Usage

```hcl
resource "fortimanager_object_antivirus_profile" "trname" {
  analytics_db         = "disable"
  analytics_max_upload = 20
  av_block_log         = "disable"
  av_virus_log         = "disable"
  comment              = "tefv comment"
  extended_log         = "disable"
  ftgd_analytics       = "disable"
  mobile_malware_db    = "disable"
  name                 = "terr-antivirus-profile"
  scan_mode            = "default"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `analytics_accept_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox.
* `analytics_bl_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox.
* `analytics_db` - Enable/disable using the FortiSandbox signature database to supplement the AV signature databases. Valid values: `disable`, `enable`.

* `analytics_ignore_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox.
* `analytics_max_upload` - Maximum size of files that can be uploaded to FortiSandbox (1 - 395 MBytes, default = 10).
* `analytics_wl_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox.
* `av_block_log` - Enable/disable logging for AntiVirus file blocking. Valid values: `disable`, `enable`.

* `av_virus_log` - Enable/disable AntiVirus logging. Valid values: `disable`, `enable`.

* `cifs` - Cifs. The structure of `cifs` block is documented below.
* `comment` - Comment.
* `content_disarm` - Content-Disarm. The structure of `content_disarm` block is documented below.
* `ems_threat_feed` - Enable/disable use of EMS threat feed when performing AntiVirus scan. Valid values: `disable`, `enable`.

* `extended_log` - Enable/disable extended logging for antivirus. Valid values: `disable`, `enable`.

* `external_blocklist` - One or more external malware block lists.
* `external_blocklist_archive_scan` - Enable/disable external-blocklist archive scanning. Valid values: `disable`, `enable`.

* `external_blocklist_enable_all` - Enable/disable all external blocklists. Valid values: `disable`, `enable`.

* `feature_set` - Flow/proxy feature set. Valid values: `proxy`, `flow`.

* `ftgd_analytics` - Settings to control which files are uploaded to FortiSandbox. Valid values: `disable`, `suspicious`, `everything`.

* `ftp` - Ftp. The structure of `ftp` block is documented below.
* `http` - Http. The structure of `http` block is documented below.
* `imap` - Imap. The structure of `imap` block is documented below.
* `mapi` - Mapi. The structure of `mapi` block is documented below.
* `mobile_malware_db` - Enable/disable using the mobile malware signature database. Valid values: `disable`, `enable`.

* `nac_quar` - Nac-Quar. The structure of `nac_quar` block is documented below.
* `name` - Profile name.
* `nntp` - Nntp. The structure of `nntp` block is documented below.
* `outbreak_prevention_archive_scan` - Enable/disable outbreak-prevention archive scanning. Valid values: `disable`, `enable`.

* `outbreak_prevention` - Outbreak-Prevention. The structure of `outbreak_prevention` block is documented below.
* `pop3` - Pop3. The structure of `pop3` block is documented below.
* `replacemsg_group` - Replacement message group customized for this profile.
* `scan_mode` - Choose between default scan mode and legacy scan mode. Valid values: `quick`, `full`, `legacy`, `default`.

* `smtp` - Smtp. The structure of `smtp` block is documented below.
* `ssh` - Ssh. The structure of `ssh` block is documented below.

The `cifs` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `external_blocklist` - Enable external-blocklist. Valid values: `disable`, `monitor`, `block`.

* `options` - Enable/disable CIFS AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avmonitor`.

* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`, `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `content_disarm` block supports:

* `cover_page` - Enable/disable inserting a cover page into the disarmed document. Valid values: `disable`, `enable`.

* `detect_only` - Enable/disable only detect disarmable files, do not alter content. Valid values: `disable`, `enable`.

* `error_action` - Action to be taken if CDR engine encounters an unrecoverable error. Valid values: `block`, `log-only`, `ignore`.

* `office_action` - Enable/disable stripping of PowerPoint action events in Microsoft Office documents. Valid values: `disable`, `enable`.

* `office_dde` - Enable/disable stripping of Dynamic Data Exchange events in Microsoft Office documents. Valid values: `disable`, `enable`.

* `office_embed` - Enable/disable stripping of embedded objects in Microsoft Office documents. Valid values: `disable`, `enable`.

* `office_hylink` - Enable/disable stripping of hyperlinks in Microsoft Office documents. Valid values: `disable`, `enable`.

* `office_linked` - Enable/disable stripping of linked objects in Microsoft Office documents. Valid values: `disable`, `enable`.

* `office_macro` - Enable/disable stripping of macros in Microsoft Office documents. Valid values: `disable`, `enable`.

* `original_file_destination` - Destination to send original file if active content is removed. Valid values: `fortisandbox`, `quarantine`, `discard`.

* `pdf_act_form` - Enable/disable stripping of PDF document actions that submit data to other targets. Valid values: `disable`, `enable`.

* `pdf_act_gotor` - Enable/disable stripping of PDF document actions that access other PDF documents. Valid values: `disable`, `enable`.

* `pdf_act_java` - Enable/disable stripping of PDF document actions that execute JavaScript code. Valid values: `disable`, `enable`.

* `pdf_act_launch` - Enable/disable stripping of PDF document actions that launch other applications. Valid values: `disable`, `enable`.

* `pdf_act_movie` - Enable/disable stripping of PDF document actions that play a movie. Valid values: `disable`, `enable`.

* `pdf_act_sound` - Enable/disable stripping of PDF document actions that play a sound. Valid values: `disable`, `enable`.

* `pdf_embedfile` - Enable/disable stripping of embedded files in PDF documents. Valid values: `disable`, `enable`.

* `pdf_hyperlink` - Enable/disable stripping of hyperlinks from PDF documents. Valid values: `disable`, `enable`.

* `pdf_javacode` - Enable/disable stripping of JavaScript code in PDF documents. Valid values: `disable`, `enable`.


The `ftp` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `external_blocklist` - Enable external-blocklist. Valid values: `disable`, `monitor`, `block`.

* `options` - Enable/disable FTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `file-filter`, `quarantine`, `avquery`, `avmonitor`.

* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`, `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `http` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol. Valid values: `disable`, `enable`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `external_blocklist` - Enable external-blocklist. Valid values: `disable`, `monitor`, `block`.

* `options` - Enable/disable HTTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `file-filter`, `quarantine`, `avquery`, `avmonitor`, `strict-file`.

* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`, `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `imap` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol. Valid values: `disable`, `enable`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.

* `external_blocklist` - Enable external-blocklist. Valid values: `disable`, `monitor`, `block`.

* `options` - Enable/disable IMAP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `file-filter`, `quarantine`, `avquery`, `avmonitor`.

* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`, `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `mapi` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.

* `external_blocklist` - Enable external-blocklist. Valid values: `disable`, `monitor`, `block`.

* `options` - Enable/disable MAPI AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avquery`, `avmonitor`.

* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`, `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `nac_quar` block supports:

* `expiry` - Duration of quarantine.
* `infected` - Enable/Disable quarantining infected hosts to the banned user list. Valid values: `none`, `quar-src-ip`, `quar-interface`.

* `log` - Enable/disable AntiVirus quarantine logging. Valid values: `disable`, `enable`.


The `nntp` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `external_blocklist` - Enable external-blocklist. Valid values: `disable`, `monitor`, `block`.

* `options` - Enable/disable NNTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `file-filter`, `quarantine`, `avquery`, `avmonitor`.

* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`, `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `outbreak_prevention` block supports:

* `external_blocklist` - Enable/disable external malware blocklist. Valid values: `disable`, `enable`.

* `ftgd_service` - Enable/disable FortiGuard Virus outbreak prevention service. Valid values: `disable`, `enable`.


The `pop3` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol. Valid values: `disable`, `enable`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.

* `external_blocklist` - Enable external-blocklist. Valid values: `disable`, `monitor`, `block`.

* `options` - Enable/disable POP3 AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `file-filter`, `quarantine`, `avquery`, `avmonitor`.

* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`, `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `smtp` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol. Valid values: `disable`, `enable`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.

* `external_blocklist` - Enable external-blocklist. Valid values: `disable`, `monitor`, `block`.

* `options` - Enable/disable SMTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `file-filter`, `quarantine`, `avquery`, `avmonitor`.

* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`, `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `ssh` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `fileslimit`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `external_blocklist` - Enable external-blocklist. Valid values: `disable`, `monitor`, `block`.

* `options` - Enable/disable SFTP and SCP AntiVirus scanning, monitoring, and quarantine. Valid values: `avmonitor`, `quarantine`, `scan`.

* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`, `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectAntivirus Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_antivirus_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
