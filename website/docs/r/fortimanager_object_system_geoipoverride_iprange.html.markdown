---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_geoipoverride_iprange"
description: |-
  Table of IP ranges assigned to country.
---

# fortimanager_object_system_geoipoverride_iprange
Table of IP ranges assigned to country.

~> This resource is a sub resource for variable `ip_range` of resource `fortimanager_object_system_geoipoverride`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_geoipoverride_iprange" "trname" {
  geoip_override = fortimanager_object_system_geoipoverride.trname.name
  end_ip         = "10.160.2.27"
  fosid          = 1
  start_ip       = "10.160.2.25"
  depends_on     = [fortimanager_object_system_geoipoverride.trname]
}

resource "fortimanager_object_system_geoipoverride" "trname" {
  name = "ACN3"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `geoip_override` - Geoip Override.

* `end_ip` - Final IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
* `fosid` - ID number for individual entry in the IP-Range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectSystem GeoipOverrideIpRange can be imported using any of these accepted formats:
```
Set import_options = ["geoip_override=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_geoipoverride_iprange.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
