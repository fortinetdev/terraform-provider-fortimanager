---
subcategory: "Object Emailfilter"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_emailfilter_profile"
description: |-
  Configure Email Filter profiles.
---

# fortimanager_object_emailfilter_profile
Configure Email Filter profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`file_filter`: `fortimanager_object_emailfilter_profile_filefilter`
`gmail`: `fortimanager_object_emailfilter_profile_gmail`
`imap`: `fortimanager_object_emailfilter_profile_imap`
`mapi`: `fortimanager_object_emailfilter_profile_mapi`
`msn_hotmail`: `fortimanager_object_emailfilter_profile_msnhotmail`
`other_webmails`: `fortimanager_object_emailfilter_profile_otherwebmails`
`pop3`: `fortimanager_object_emailfilter_profile_pop3`
`smtp`: `fortimanager_object_emailfilter_profile_smtp`



## Example Usage

```hcl
resource "fortimanager_object_emailfilter_profile" "trname" {
  comment                      = "This is a Terraform example"
  external                     = "disable"
  feature_set                  = "flow"
  name                         = "terr-emailfilter-profile"
  options                      = ["bannedword", "spambwl"]
  spam_bword_threshold         = 10
  spam_filtering               = "disable"
  spam_log                     = "enable"
  spam_log_fortiguard_response = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `external` - Enable/disable external Email inspection. Valid values: `disable`, `enable`.

* `file_filter` - File-Filter. The structure of `file_filter` block is documented below.
* `feature_set` - Flow/proxy feature set. Valid values: `proxy`, `flow`.

* `gmail` - Gmail. The structure of `gmail` block is documented below.
* `imap` - Imap. The structure of `imap` block is documented below.
* `mapi` - Mapi. The structure of `mapi` block is documented below.
* `msn_hotmail` - Msn-Hotmail. The structure of `msn_hotmail` block is documented below.
* `name` - Profile name.
* `options` - Options. Valid values: `bannedword`, `spambwl`, `spamfsip`, `spamfssubmit`, `spamfschksum`, `spamfsurl`, `spamhelodns`, `spamraddrdns`, `spamrbl`, `spamhdrcheck`, `spamfsphish`.

* `other_webmails` - Other-Webmails. The structure of `other_webmails` block is documented below.
* `pop3` - Pop3. The structure of `pop3` block is documented below.
* `replacemsg_group` - Replacement message group.
* `smtp` - Smtp. The structure of `smtp` block is documented below.
* `spam_bal_table` - Anti-spam block/allow list table ID.
* `spam_bwl_table` - Anti-spam black/white list table ID.
* `spam_bword_table` - Anti-spam banned word table ID.
* `spam_bword_threshold` - Spam banned word threshold.
* `spam_filtering` - Enable/disable spam filtering. Valid values: `disable`, `enable`.

* `spam_iptrust_table` - Anti-spam IP trust table ID.
* `spam_log` - Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.

* `spam_log_fortiguard_response` - Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.

* `spam_mheader_table` - Anti-spam MIME header table ID.
* `spam_rbl_table` - Anti-spam DNSBL table ID.

The `file_filter` block supports:

* `entries` - Entries. The structure of `entries` block is documented below.
* `log` - Enable/disable file filter logging. Valid values: `disable`, `enable`.

* `scan_archive_contents` - Enable/disable file filter archive contents scan. Valid values: `disable`, `enable`.

* `status` - Enable/disable file filter. Valid values: `disable`, `enable`.


The `entries` block supports:

* `action` - Action taken for matched file. Valid values: `log`, `block`.

* `comment` - Comment.
* `encryption` - Encryption. Valid values: `any`, `yes`.

* `file_type` - Select file type.
* `filter` - Add a file filter.
* `password_protected` - Match password-protected files. Valid values: `any`, `yes`.

* `protocol` - Protocols to apply with. Valid values: `smtp`, `imap`, `pop3`.


The `gmail` block supports:

* `log` - Log. Valid values: `disable`, `enable`.

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.


The `imap` block supports:

* `action` - Action for spam email. Valid values: `pass`, `tag`.

* `log` - Log. Valid values: `disable`, `enable`.

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.

* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.


The `mapi` block supports:

* `action` - Action for spam email. Valid values: `pass`, `discard`.

* `log` - Log. Valid values: `disable`, `enable`.

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.


The `msn_hotmail` block supports:

* `log` - Log. Valid values: `disable`, `enable`.

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.


The `other_webmails` block supports:

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.


The `pop3` block supports:

* `action` - Action for spam email. Valid values: `pass`, `tag`.

* `log` - Log. Valid values: `disable`, `enable`.

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.

* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.


The `smtp` block supports:

* `action` - Action for spam email. Valid values: `pass`, `tag`, `discard`.

* `hdrip` - Enable/disable SMTP email header IP checks for spamfsip, spamrbl and spambwl filters. Valid values: `disable`, `enable`.

* `local_override` - Enable/disable local filter to override SMTP remote check result. Valid values: `disable`, `enable`.

* `log` - Log. Valid values: `disable`, `enable`.

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.

* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectEmailfilter Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_emailfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
