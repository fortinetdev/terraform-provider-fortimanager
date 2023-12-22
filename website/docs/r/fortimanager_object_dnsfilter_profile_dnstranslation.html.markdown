---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dnsfilter_profile_dnstranslation"
description: |-
  DNS translation settings.
---

# fortimanager_object_dnsfilter_profile_dnstranslation
DNS translation settings.

~> This resource is a sub resource for variable `dns_translation` of resource `fortimanager_object_dnsfilter_profile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_dnsfilter_profile_dnstranslation" "trname" {
  dst        = "38.53.6.21"
  fosid      = 1
  src        = "34.53.6.21"
  status     = "disable"
  profile    = fortimanager_object_dnsfilter_profile.trname.name
  depends_on = [fortimanager_object_dnsfilter_profile.trname]
}

resource "fortimanager_object_dnsfilter_profile" "trname" {
  name = "terr-dnsfilter-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `addr_type` - DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `dst` - IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
* `dst6` - IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
* `fosid` - ID.
* `netmask` - If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
* `prefix` - If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
* `src` - IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
* `src6` - IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
* `status` - Enable/disable this DNS translation entry. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectDnsfilter ProfileDnsTranslation can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dnsfilter_profile_dnstranslation.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
