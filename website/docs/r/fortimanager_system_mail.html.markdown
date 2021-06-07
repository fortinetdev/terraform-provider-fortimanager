---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_mail"
description: |-
  Alert emails.
---

# fortimanager_system_mail
Alert emails.

## Argument Reference


The following arguments are supported:


* `auth` - Enable authentication. disable - Disable authentication. enable - Enable authentication. Valid values: `disable`, `enable`.

* `fosid` - Mail Service ID.
* `passwd` - SMTP account password.
* `port` - SMTP server port.
* `secure_option` - Communication secure option. default - Try STARTTLS, proceed as plain text communication otherwise. none - Communication will be in plain text format. smtps - Communication will be protected by SMTPS. starttls - Communication will be protected by STARTTLS. Valid values: `default`, `none`, `smtps`, `starttls`.

* `server` - SMTP server.
* `user` - SMTP account username.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Mail can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_mail.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

