---
subcategory: "System Log"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_logfetch_serversettings"
description: |-
  Log-fetch server settings.
---

# fortimanager_system_logfetch_serversettings
Log-fetch server settings.

## Example Usage

```hcl
resource "fortimanager_system_logfetch_serversettings" "trname" {
  max_sessions    = "10"
  session_timeout = "200"
}
```

## Argument Reference


The following arguments are supported:


* `max_conn_per_session` - Max concurrent file download connections per session.
* `max_sessions` - Max concurrent fetch sessions.
* `session_timeout` - Fetch session timeout in minute.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogFetchServerSettings can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_logfetch_serversettings.labelname SystemLogFetchServerSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

