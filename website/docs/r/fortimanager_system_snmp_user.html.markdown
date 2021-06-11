---
subcategory: "System SNMP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_snmp_user"
description: |-
  SNMP user configuration.
---

# fortimanager_system_snmp_user
SNMP user configuration.

## Example Usage

```hcl
resource "fortimanager_system_snmp_user" "trname" {
  auth_proto     = "sha"
  auth_pwd       = ["fortinet"]
  events         = ["disk_low", "ha_switch", "intf_ip_chg", "sys_reboot", "cpu_high", "mem_low", "log-alert", "log-rate", "log-data-rate", "lic-gbday", "lic-dev-quota", "cpu-high-exclude-nice"]
  name           = "terraform-tefv-snmpuser"
  queries        = "disable"
  security_level = "no-auth-no-priv"
}
```

## Argument Reference


The following arguments are supported:


* `auth_proto` - Authentication protocol. md5 - HMAC-MD5-96 authentication protocol. sha - HMAC-SHA-96 authentication protocol. Valid values: `md5`, `sha`.

* `auth_pwd` - Password for authentication protocol.
* `events` - SNMP notifications (traps) to send. disk_low - Disk usage too high. ha_switch - HA switch. intf_ip_chg - Interface IP address changed. sys_reboot - System reboot. cpu_high - CPU usage too high. mem_low - Available memory is low. log-alert - Log base alert message. log-rate - High incoming log rate detected. log-data-rate - High incoming log data rate detected. lic-gbday - High licensed log GB/day detected. lic-dev-quota - High licensed device quota detected. cpu-high-exclude-nice - CPU usage exclude NICE threshold. Valid values: `disk_low`, `ha_switch`, `intf_ip_chg`, `sys_reboot`, `cpu_high`, `mem_low`, `log-alert`, `log-rate`, `log-data-rate`, `lic-gbday`, `lic-dev-quota`, `cpu-high-exclude-nice`.

* `name` - SNMP user name.
* `notify_hosts` - Hosts to send notifications (traps) to.
* `notify_hosts6` - IPv6 hosts to send notifications (traps) to.
* `priv_proto` - Privacy (encryption) protocol. aes - CFB128-AES-128 symmetric encryption protocol. des - CBC-DES symmetric encryption protocol. Valid values: `aes`, `des`.

* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable queries for this user. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_port` - SNMPv3 query port.
* `security_level` - Security level for message authentication and encryption. no-auth-no-priv - Message with no authentication and no privacy (encryption). auth-no-priv - Message with authentication but no privacy (encryption). auth-priv - Message with authentication and privacy (encryption). Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SnmpUser can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_snmp_user.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

