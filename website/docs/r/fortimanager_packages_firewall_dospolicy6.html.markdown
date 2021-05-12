---
subcategory: "Packages"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_dospolicy6"
description: |-
  Configure IPv6 DoS policies.
---

# fortimanager_packages_firewall_dospolicy6
Configure IPv6 DoS policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `anomaly` - Anomaly. The structure of `anomaly` block is documented below.
* `comments` - Comment.
* `dstaddr` - Destination address name from available addresses.
* `interface` - Incoming interface name from available interfaces.
* `name` - Policy name.
* `policyid` - Policy ID.
* `service` - Service object from available options.
* `srcaddr` - Source address name from available addresses.
* `status` - Enable/disable this policy. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `anomaly` block supports:

* `action` - Action taken when the threshold is reached. Valid values: `pass`, `block`, `proxy`.

* `log` - Enable/disable anomaly logging. Valid values: `disable`, `enable`.

* `name` - Anomaly name.
* `quarantine` - Quarantine method. Valid values: `none`, `attacker`, `both`, `interface`.

* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable`, `enable`.

* `status` - Enable/disable this anomaly. Valid values: `disable`, `enable`.

* `threshold` - Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.
* `thresholddefault` - Threshold(Default).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallDosPolicy6 can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_dospolicy6.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
