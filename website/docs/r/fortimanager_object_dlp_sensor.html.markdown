---
subcategory: "Object DLP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dlp_sensor"
description: |-
  Configure DLP sensors.
---

# fortimanager_object_dlp_sensor
Configure DLP sensors.

## Example Usage

```hcl
resource "fortimanager_object_dlp_sensor" "trname" {
  dlp_log            = "enable"
  extended_log       = "disable"
  feature_set        = "flow"
  full_archive_proto = ["ftp", "http-get", "http-post", "imap", "mapi", "nntp"]
  nac_quar_log       = "disable"
  name               = "terr-dlp-sensor"
  summary_proto      = ["ftp", "http-get", "http-post", "imap", "mapi", "nntp"]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `entries` - Entries. The structure of `entries` block is documented below.
* `eval` - Expression to evaluate.
* `match_type` - Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`, `match-eval`.

* `dlp_log` - Enable/disable DLP logging. Valid values: `disable`, `enable`.

* `extended_log` - Enable/disable extended logging for data leak prevention. Valid values: `disable`, `enable`.

* `feature_set` - Flow/proxy feature set. Valid values: `proxy`, `flow`.

* `filter` - Filter. The structure of `filter` block is documented below.
* `flow_based` - Enable/disable flow-based DLP. Valid values: `disable`, `enable`.

* `full_archive_proto` - Protocols to always content archive. Valid values: `imap`, `smtp`, `pop3`, `ftp`, `nntp`, `mm1`, `mm3`, `mm4`, `mm7`, `mapi`, `aim`, `icq`, `msn`, `yahoo`, `http-get`, `http-post`, `cifs`, `ssh`.

* `nac_quar_log` - Enable/disable NAC quarantine logging. Valid values: `disable`, `enable`.

* `name` - Name of the DLP sensor.
* `options` - Configure DLP options. Valid values: `strict-file`.

* `replacemsg_group` - Replacement message group used by this DLP sensor.
* `summary_proto` - Protocols to always log summary. Valid values: `imap`, `smtp`, `pop3`, `ftp`, `nntp`, `mm1`, `mm3`, `mm4`, `mm7`, `mapi`, `aim`, `icq`, `msn`, `yahoo`, `http-get`, `http-post`, `cifs`, `ssh`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `count` - Count of dictionary matches to trigger sensor entry match (Dictionary might not be able to trigger more than once based on its 'repeat' option, 1 - 255, default = 1).
* `dictionary` - Select a DLP dictionary.
* `id` - ID.
* `status` - Enable/disable this entry. Valid values: `disable`, `enable`.


The `filter` block supports:

* `action` - Action to take with content that this DLP sensor matches. Valid values: `log-only`, `block`, `exempt`, `ban`, `ban-sender`, `quarantine-ip`, `quarantine-port`, `none`, `allow`.

* `archive` - Enable/disable DLP archiving. Valid values: `disable`, `enable`, `summary-only`.

* `company_identifier` - Enter a company identifier watermark to match. Only watermarks that your company has placed on the files are matched.
* `expiry` - Quarantine duration in days, hours, minutes format (dddhhmm).
* `file_size` - Match files this size or larger (0 - 4294967295 kbytes).
* `file_type` - Select the number of a DLP file pattern table to match.
* `filter_by` - Select the type of content to match. Valid values: `credit-card`, `ssn`, `regexp`, `file-type`, `file-size`, `fingerprint`, `watermark`, `encrypted`.

* `fp_sensitivity` - Select a DLP file pattern sensitivity to match.
* `id` - ID.
* `match_percentage` - Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.
* `name` - Filter name.
* `proto` - Check messages or files over one or more of these protocols. Valid values: `imap`, `smtp`, `pop3`, `ftp`, `nntp`, `mm1`, `mm3`, `mm4`, `mm7`, `mapi`, `aim`, `icq`, `msn`, `yahoo`, `http-get`, `http-post`, `cifs`, `ssh`.

* `regexp` - Enter a regular expression to match (max. 255 characters).
* `sensitivity` - Select a DLP file pattern sensitivity to match.
* `severity` - Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.

* `type` - Select whether to check the content of messages (an email message) or files (downloaded files or email attachments). Valid values: `file`, `message`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDlp Sensor can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dlp_sensor.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
