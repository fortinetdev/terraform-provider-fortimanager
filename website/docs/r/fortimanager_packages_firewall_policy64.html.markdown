---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_policy64"
description: |-
  Configure IPv6 to IPv4 policies.
---

# fortimanager_packages_firewall_policy64
Configure IPv6 to IPv4 policies.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_policy64" "labelname" {
  action           = "deny"
  dstaddr          = "all"
  dstintf          = "any"
  fixedport        = "disable"
  ippool           = "disable"
  logtraffic       = "disable"
  logtraffic_start = "disable"
  name             = "s"
  permit_any_host  = "disable"
  pkg              = "default"
  policyid         = 1
  schedule         = "always"
  service          = "ALL"
  srcaddr          = "all"
  srcintf          = "any"
  status           = "enable"
  tcp_mss_receiver = 0
  tcp_mss_sender   = 0
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Policy action. Valid values: `deny`, `accept`.

* `comments` - Comment.
* `dstaddr` - Destination address name.
* `dstintf` - Destination interface name.
* `fixedport` - Enable/disable policy fixed port. Valid values: `disable`, `enable`.

* `ippool` - Enable/disable policy64 IP pool. Valid values: `disable`, `enable`.

* `logtraffic` - Enable/disable policy log traffic. Valid values: `disable`, `enable`.

* `logtraffic_start` - Record logs when a session starts and ends. Valid values: `disable`, `enable`.

* `name` - Policy name.
* `per_ip_shaper` - Per-IP traffic shaper.
* `permit_any_host` - Enable/disable permit any host in. Valid values: `disable`, `enable`.

* `policyid` - Policy ID (0 - 4294967294).
* `poolname` - Policy IP pool names.
* `schedule` - Schedule name.
* `service` - Service name.
* `srcaddr` - Source address name.
* `srcintf` - Source interface name.
* `status` - Enable/disable policy status. Valid values: `disable`, `enable`.

* `tcp_mss_receiver` - TCP MSS value of receiver.
* `tcp_mss_sender` - TCP MSS value of sender.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallPolicy64 can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_policy64.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
