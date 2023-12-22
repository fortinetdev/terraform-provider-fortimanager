---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_antivirus_profile_contentdisarm"
description: |-
  AV Content Disarm and Reconstruction settings.
---

# fortimanager_object_antivirus_profile_contentdisarm
AV Content Disarm and Reconstruction settings.

~> This resource is a sub resource for variable `content_disarm` of resource `fortimanager_object_antivirus_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

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



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectAntivirus ProfileContentDisarm can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_antivirus_profile_contentdisarm.labelname ObjectAntivirusProfileContentDisarm
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
