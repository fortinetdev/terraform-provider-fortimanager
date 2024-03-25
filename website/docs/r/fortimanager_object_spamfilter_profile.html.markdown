---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_spamfilter_profile"
description: |-
  Configure AntiSpam profiles.
---

# fortimanager_object_spamfilter_profile
Configure AntiSpam profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `gmail`: `fortimanager_object_spamfilter_profile_gmail`
>- `imap`: `fortimanager_object_spamfilter_profile_imap`
>- `mapi`: `fortimanager_object_spamfilter_profile_mapi`
>- `msn_hotmail`: `fortimanager_object_spamfilter_profile_msnhotmail`
>- `pop3`: `fortimanager_object_spamfilter_profile_pop3`
>- `smtp`: `fortimanager_object_spamfilter_profile_smtp`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `external` - Enable/disable external Email inspection. Valid values: `disable`, `enable`.

* `flow_based` - Enable/disable flow-based spam filtering. Valid values: `disable`, `enable`.

* `gmail` - Gmail. The structure of `gmail` block is documented below.
* `imap` - Imap. The structure of `imap` block is documented below.
* `mapi` - Mapi. The structure of `mapi` block is documented below.
* `msn_hotmail` - Msn-Hotmail. The structure of `msn_hotmail` block is documented below.
* `name` - Profile name.
* `options` - Options. Valid values: `bannedword`, `spamemailbwl`, `spamfsip`, `spamfssubmit`, `spamfschksum`, `spamfsurl`, `spamhelodns`, `spamipbwl`, `spamraddrdns`, `spamrbl`, `spamhdrcheck`, `spamfsphish`, `spambwl`.

* `pop3` - Pop3. The structure of `pop3` block is documented below.
* `replacemsg_group` - Replacement message group.
* `smtp` - Smtp. The structure of `smtp` block is documented below.
* `spam_bwl_table` - Anti-spam black/white list table ID.
* `spam_bword_table` - Anti-spam banned word table ID.
* `spam_bword_threshold` - Spam banned word threshold.
* `spam_filtering` - Enable/disable spam filtering. Valid values: `disable`, `enable`.

* `spam_iptrust_table` - Anti-spam IP trust table ID.
* `spam_log` - Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.

* `spam_log_fortiguard_response` - Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.

* `spam_mheader_table` - Anti-spam MIME header table ID.
* `spam_rbl_table` - Anti-spam DNSBL table ID.

The `gmail` block supports:

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.


The `imap` block supports:

* `action` - Action for spam email. Valid values: `pass`, `tag`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.


The `mapi` block supports:

* `action` - Action for spam email. Valid values: `pass`, `discard`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.


The `msn_hotmail` block supports:

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.


The `pop3` block supports:

* `action` - Action for spam email. Valid values: `pass`, `tag`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.


The `smtp` block supports:

* `action` - Action for spam email. Valid values: `pass`, `tag`, `discard`.

* `hdrip` - Enable/disable SMTP email header IP checks for spamfsip, spamrbl and spambwl filters. Valid values: `disable`, `enable`.

* `local_override` - Enable/disable local filter to override SMTP remote check result. Valid values: `disable`, `enable`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSpamfilter Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_spamfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
