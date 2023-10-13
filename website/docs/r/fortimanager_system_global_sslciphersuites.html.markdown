---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_global_sslciphersuites"
description: |-
  Configure preferred SSL/TLS cipher suites
---

# fortimanager_system_global_sslciphersuites
Configure preferred SSL/TLS cipher suites

## Argument Reference


The following arguments are supported:


* `cipher` - Cipher name
* `priority` - SSL/TLS cipher suites priority.
* `version` - SSL/TLS version the cipher suite can be used with. tls1.2-or-below - TLS 1.2 or below. tls1.3 - TLS 1.3 Valid values: `tls1.2-or-below`, `tls1.3`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{cipher}}.

## Import

System GlobalSslCipherSuites can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_global_sslciphersuites.labelname {{cipher}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

