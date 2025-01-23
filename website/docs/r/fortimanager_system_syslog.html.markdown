---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_syslog"
description: |-
  Syslog servers.
---

# fortimanager_system_syslog
Syslog servers.

## Example Usage

```hcl
resource "fortimanager_system_syslog" "trname" {
  ip   = "192.168.1.1"
  name = "terr-sys-syslog"
}
```

## Argument Reference


The following arguments are supported:


* `ip` - Syslog server IP address or hostname.
* `local_cert` - Select local certificate used for secure connection.
* `name` - Syslog server name.
* `peer_cert_cn` - Certificate common name of syslog server. null or &apos;-&apos; means not check certificate CN of syslog server
* `port` - Syslog server port.
* `reliable` - Enable/disable reliable connection with syslog server. disable - Disable reliable connection with syslog server. enable - Enable reliable connection with syslog server. Valid values: `disable`, `enable`.

* `secure_connection` - Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: `disable`, `enable`.

* `ssl_protocol` - set the lowest SSL protocol version for connection to syslog server. follow-global-ssl-protocol - Follow system.global.global-ssl-protocol setting (default). sslv3 - set SSLv3 as the lowest version. tlsv1.0 - set TLSv1.0 as the lowest version. tlsv1.1 - set TLSv1.1 as the lowest version. tlsv1.2 - set TLSv1.2 as the lowest version. tlsv1.3 - set TLSv1.3 as the lowest version. Valid values: `follow-global-ssl-protocol`, `sslv3`, `tlsv1.0`, `tlsv1.1`, `tlsv1.2`, `tlsv1.3`.



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

