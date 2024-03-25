---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_multicastpolicy6"
description: |-
  Configure IPv6 multicast NAT policies.
---

# fortimanager_packages_firewall_multicastpolicy6
Configure IPv6 multicast NAT policies.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_multicastpolicy6" "labelname" {
  action            = "accept"
  auto_asic_offload = "enable"
  dstaddr           = ["all"]
  dstintf           = ["any"]
  end_port          = 65535
  fosid             = 1
  logtraffic        = "disable"
  name              = "1"
  pkg               = "default"
  protocol          = 0
  srcaddr           = ["all"]
  srcintf           = ["any"]
  start_port        = 1
  status            = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `action` - Accept or deny traffic matching the policy. Valid values: `deny`, `accept`.

* `auto_asic_offload` - Enable/disable offloading policy traffic for hardware acceleration. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `dstaddr` - IPv6 destination address name.
* `dstintf` - IPv6 destination interface name.
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
* `fosid` - Policy ID (0 - 4294967294).
* `ips_sensor` - Name of an existing IPS sensor.
* `logtraffic` - Enable/disable logging traffic accepted by this policy. Valid values: `disable`, `enable`.

* `name` - Policy name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
* `srcaddr` - IPv6 source address name.
* `srcintf` - IPv6 source interface name.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `status` - Enable/disable this policy. Valid values: `disable`, `enable`.

* `utm_status` - Enable to add an IPS security profile to the policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Packages FirewallMulticastPolicy6 can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_multicastpolicy6.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
