---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_system_snmp_user"
description: |-
  SNMP user configuration.
---

# fortimanager_systemp_system_snmp_user
SNMP user configuration.

## Example Usage

```hcl
resource "fortimanager_systemp_system_snmp_user" "trname" {
  devprof    = "default"
  auth_proto = "sha512"
  name       = "terr-user"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `auth_proto` - Authentication protocol. Valid values: `md5`, `sha`, `sha224`, `sha256`, `sha384`, `sha512`.

* `auth_pwd` - Password for authentication protocol.
* `events` - SNMP notifications (traps) to send. Valid values: `cpu-high`, `mem-low`, `log-full`, `intf-ip`, `vpn-tun-up`, `vpn-tun-down`, `ha-switch`, `fm-conf-change`, `ips-signature`, `ips-anomaly`, `temperature-high`, `voltage-alert`, `av-virus`, `av-oversize`, `av-pattern`, `av-fragmented`, `ha-hb-failure`, `fan-failure`, `ha-member-up`, `ha-member-down`, `ent-conf-change`, `av-conserve`, `av-bypass`, `av-oversize-passed`, `av-oversize-blocked`, `ips-pkg-update`, `fm-if-change`, `power-supply-failure`, `amc-bypass`, `faz-disconnect`, `bgp-established`, `bgp-backward-transition`, `wc-ap-up`, `wc-ap-down`, `fswctl-session-up`, `fswctl-session-down`, `ips-fail-open`, `load-balance-real-server-down`, `device-new`, `enter-intf-bypass`, `exit-intf-bypass`, `per-cpu-high`, `power-blade-down`, `confsync_failure`, `dhcp`, `pool-usage`, `power-redundancy-degrade`, `power-redundancy-failure`.

* `ha_direct` - Enable/disable direct management of HA cluster members. Valid values: `disable`, `enable`.

* `mib_view` - SNMP access control MIB view.
* `name` - SNMP user name.
* `notify_hosts` - SNMP managers to send notifications (traps) to.
* `notify_hosts6` - IPv6 SNMP managers to send notifications (traps) to.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.

* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.

* `query_port` - SNMPv3 query port (default = 161).
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.

* `source_ip` - Source IP for SNMP trap.
* `source_ipv6` - Source IPv6 for SNMP trap.
* `status` - Enable/disable this SNMP user. Valid values: `disable`, `enable`.

* `trap_lport` - SNMPv3 local trap port (default = 162).
* `trap_rport` - SNMPv3 trap remote port (default = 162).
* `trap_status` - Enable/disable traps for this SNMP user. Valid values: `disable`, `enable`.

* `vdoms` - SNMP access control VDOMs.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Systemp SystemSnmpUser can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_system_snmp_user.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
