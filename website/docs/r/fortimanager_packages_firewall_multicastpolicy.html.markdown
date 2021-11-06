---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_multicastpolicy"
description: |-
  Configure multicast NAT policies.
---

# fortimanager_packages_firewall_multicastpolicy
Configure multicast NAT policies.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_multicastpolicy" "labelname" {
  action            = "accept"
  auto_asic_offload = "enable"
  dnat              = "0.0.0.0"
  dstaddr           = ["all"]
  dstintf           = ["any"]
  end_port          = 0
  fosid             = 1
  logtraffic        = "disable"
  name              = "1"
  pkg               = "default"
  protocol          = 0
  snat              = "disable"
  srcaddr           = ["all"]
  srcintf           = ["any"]
  start_port        = 0
  status            = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Accept or deny traffic matching the policy. Valid values: `deny`, `accept`.

* `auto_asic_offload` - Enable/disable offloading policy traffic for hardware acceleration. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `dnat` - IPv4 DNAT address used for multicast destination addresses.
* `dstaddr` - Destination address objects.
* `dstintf` - Destination interface name.
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `fosid` - Policy ID ((0 - 4294967294).
* `logtraffic` - Enable/disable logging traffic accepted by this policy. Valid values: `disable`, `enable`.

* `name` - Policy name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
* `snat` - Enable/disable substitution of the outgoing interface IP address for the original source IP address (called source NAT or SNAT). Valid values: `disable`, `enable`.

* `snat_ip` - IPv4 address to be used as the source address for NATed traffic.
* `srcaddr` - Source address objects.
* `srcintf` - Source interface name.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `status` - Enable/disable this policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Packages FirewallMulticastPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_multicastpolicy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
