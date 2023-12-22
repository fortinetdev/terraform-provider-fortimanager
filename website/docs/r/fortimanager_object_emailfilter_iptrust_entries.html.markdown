---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_emailfilter_iptrust_entries"
description: |-
  Spam filter trusted IP addresses.
---

# fortimanager_object_emailfilter_iptrust_entries
Spam filter trusted IP addresses.

~> This resource is a sub resource for variable `entries` of resource `fortimanager_object_emailfilter_iptrust`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_emailfilter_iptrust_entries" "trname" {
  iptrust    = fortimanager_object_emailfilter_iptrust.trname.fosid
  addr_type  = "ipv4"
  fosid      = 1
  ip4_subnet = "32.25.23.0 255.255.255.0"
  depends_on = [fortimanager_object_emailfilter_iptrust.trname]
}

resource "fortimanager_object_emailfilter_iptrust" "trname" {
  comment = "This is a Terraform example"
  fosid   = 1
  name    = "terr-emailfilter-iptrust"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `iptrust` - Iptrust.

* `addr_type` - Type of address. Valid values: `ipv4`, `ipv6`.

* `fosid` - Trusted IP entry ID.
* `ip4_subnet` - IPv4 network address or network address/subnet mask bits.
* `ip6_subnet` - IPv6 network address/subnet mask bits.
* `status` - Enable/disable status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectEmailfilter IptrustEntries can be imported using any of these accepted formats:
```
Set import_options = ["iptrust=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_emailfilter_iptrust_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
