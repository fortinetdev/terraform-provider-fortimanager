---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_log_maildomain"
description: |-
  FortiMail domain setting.
---

# fortimanager_system_log_maildomain
FortiMail domain setting.

## Argument Reference


The following arguments are supported:


* `devices` - Devices for domain to vdom mapping
* `domain` - FortiMail domain
* `fosid` - ID of FortiMail domain.
* `vdom` - Virtual Domain name mapping to FortiMail domain


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogMailDomain can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_log_maildomain.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

