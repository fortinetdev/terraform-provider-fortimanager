---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_geoipoverride_ip6range"
description: |-
  Table of IPv6 ranges assigned to country.
---

# fortimanager_object_system_geoipoverride_ip6range
Table of IPv6 ranges assigned to country.

~> This resource is a sub resource for variable `ip6_range` of resource `fortimanager_object_system_geoipoverride`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_geoipoverride_ip6range" "trname" {
  geoip_override = fortimanager_object_system_geoipoverride.trname2.name
  end_ip         = "2001:db8:85a3::8a2e:470:7334"
  fosid          = 2
  start_ip       = "2001:db8:85a3::8a2e:370:7334"
  depends_on     = [fortimanager_object_system_geoipoverride.trname2]
}

resource "fortimanager_object_system_geoipoverride" "trname2" {
  name = "ACN2"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `geoip_override` - Geoip Override.

* `end_ip` - Ending IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `fosid` - ID of individual entry in the IPv6 range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectSystem GeoipOverrideIp6Range can be imported using any of these accepted formats:
```
Set import_options = ["geoip_override=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_geoipoverride_ip6range.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
