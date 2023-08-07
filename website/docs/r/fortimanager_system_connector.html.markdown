---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_connector"
description: |-
  Configure connector.
---

# fortimanager_system_connector
Configure connector.

## Example Usage

```hcl
resource "fortimanager_system_connector" "trname" {
  fsso_refresh_interval = "70"
  fsso_sess_timeout     = "70"
  px_refresh_interval   = "70"
  px_svr_timeout        = "70"
}
```

## Argument Reference


The following arguments are supported:


* `cloud_orchest_refresh_interval` - Cloud Orchestration refresh interval (300 - 1800 seconds).
* `conn_refresh_interval` - connector refresh interval (60 - 1800 seconds).
* `fsso_refresh_interval` - FSSO refresh interval (60 - 600 seconds).
* `fsso_sess_timeout` - FSSO session timeout (60 - 600 seconds).
* `px_refresh_interval` - pxGrid refresh interval (60 - 1800 seconds).
* `px_svr_timeout` - pxGrid server timeout (30 - 600 seconds).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Connector can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_connector.labelname SystemConnector
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

