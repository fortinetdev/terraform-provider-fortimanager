---
subcategory: "System Alert"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_alertemail"
description: |-
  Configure alertemail.
---

# fortimanager_system_alertemail
Configure alertemail.

## Example Usage

```hcl
resource "fortimanager_system_alertemail" "trname" {
  fromaddress = "tefv@tefv.com"
  fromname    = "terr-sys-alertemail"
  smtpport    = "2000"
  smtpserver  = "192.168.1.2"
}
```

## Argument Reference


The following arguments are supported:


* `authentication` - Enable/disable authentication. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fromaddress` - SMTP from address.
* `fromname` - SMTP from user.
* `smtppassword` - SMTP server password.
* `smtpport` - SMTP server port.
* `smtpserver` - SMTP server address.
* `smtpuser` - SMTP server user.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Alertemail can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_alertemail.labelname SystemAlertemail
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

