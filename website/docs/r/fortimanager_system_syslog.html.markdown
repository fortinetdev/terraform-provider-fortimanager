---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_syslog"
description: |-
  Syslog servers.
---

# fortimanager_system_syslog
Syslog servers.

## Argument Reference


The following arguments are supported:


* `ip` - Syslog server IP address or hostname.
* `name` - Syslog server name.
* `port` - Syslog server port.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Syslog can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_syslog.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

