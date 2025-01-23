---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dlp_sensor_filter"
description: |-
  Set up DLP filters for this sensor.
---

# fortimanager_object_dlp_sensor_filter
Set up DLP filters for this sensor.

~> This resource is a sub resource for variable `filter` of resource `fortimanager_object_dlp_sensor`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `sensor` - Sensor.

* `action` - Action to take with content that this DLP sensor matches. Valid values: `log-only`, `block`, `exempt`, `ban`, `ban-sender`, `quarantine-ip`, `quarantine-port`, `none`, `allow`.

* `archive` - Enable/disable DLP archiving. Valid values: `disable`, `enable`, `summary-only`.

* `company_identifier` - Enter a company identifier watermark to match. Only watermarks that your company has placed on the files are matched.
* `expiry` - Quarantine duration in days, hours, minutes format (dddhhmm).
* `file_size` - Match files this size or larger (0 - 4294967295 kbytes).
* `file_type` - Select the number of a DLP file pattern table to match.
* `filter_by` - Select the type of content to match. Valid values: `credit-card`, `ssn`, `regexp`, `file-type`, `file-size`, `fingerprint`, `watermark`, `encrypted`.

* `fp_sensitivity` - Select a DLP file pattern sensitivity to match.
* `fosid` - ID.
* `match_percentage` - Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.
* `name` - Filter name.
* `proto` - Check messages or files over one or more of these protocols. Valid values: `imap`, `smtp`, `pop3`, `ftp`, `nntp`, `mm1`, `mm3`, `mm4`, `mm7`, `mapi`, `aim`, `icq`, `msn`, `yahoo`, `http-get`, `http-post`, `cifs`, `ssh`.

* `regexp` - Enter a regular expression to match (max. 255 characters).
* `sensitivity` - Select a DLP file pattern sensitivity to match.
* `severity` - Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.

* `type` - Select whether to check the content of messages (an email message) or files (downloaded files or email attachments). Valid values: `file`, `message`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectDlp SensorFilter can be imported using any of these accepted formats:
```
Set import_options = ["sensor=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dlp_sensor_filter.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
