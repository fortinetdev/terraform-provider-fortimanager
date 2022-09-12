---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_syslogprofile"
description: |-
  Configure Wireless Termination Points (WTP) system log server profile.
---

# fortimanager_object_wirelesscontroller_syslogprofile
Configure Wireless Termination Points (WTP) system log server profile.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `log_level` - Lowest level of log messages that FortiAP units send to this server (default = information). Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.

* `name` - WTP system log server profile name.
* `server_addr_type` - Syslog server address type (default = ip). Valid values: `fqdn`, `ip`.

* `server_fqdn` - FQDN of syslog server that FortiAP units send log messages to.
* `server_ip` - IP address of syslog server that FortiAP units send log messages to.
* `server_port` - Port number of syslog server that FortiAP units send log messages to (default = 514).
* `server_status` - Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController SyslogProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_syslogprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
