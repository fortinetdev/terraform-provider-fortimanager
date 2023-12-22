---
subcategory: "Object Dnsfilter"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dnsfilter_profile"
description: |-
  Configure DNS domain filter profiles.
---

# fortimanager_object_dnsfilter_profile
Configure DNS domain filter profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`dns_translation`: `fortimanager_object_dnsfilter_profile_dnstranslation`
`domain_filter`: `fortimanager_object_dnsfilter_profile_domainfilter`
`ftgd_dns`: `fortimanager_object_dnsfilter_profile_ftgddns`



## Example Usage

```hcl
resource "fortimanager_object_dnsfilter_profile" "trname" {
  block_action   = "redirect"
  block_botnet   = "disable"
  comment        = "terraform-tefv-comment"
  log_all_domain = "disable"
  name           = "terraform-tefv"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `block_action` - Action to take for blocked domains. Valid values: `block`, `redirect`.

* `block_botnet` - Enable/disable blocking botnet C&C DNS lookups. Valid values: `disable`, `enable`.

* `comment` - Comment.
* `dns_translation` - Dns-Translation. The structure of `dns_translation` block is documented below.
* `domain_filter` - Domain-Filter. The structure of `domain_filter` block is documented below.
* `external_ip_blocklist` - One or more external IP block lists.
* `ftgd_dns` - Ftgd-Dns. The structure of `ftgd_dns` block is documented below.
* `log_all_domain` - Enable/disable logging of all domains visited (detailed DNS logging). Valid values: `disable`, `enable`.

* `name` - Profile name.
* `redirect_portal` - IPv4 address of the SDNS redirect portal.
* `redirect_portal6` - IPv6 address of the SDNS redirect portal.
* `safe_search` - Enable/disable Google, Bing, and YouTube safe search. Valid values: `disable`, `enable`.

* `sdns_domain_log` - Enable/disable domain filtering and botnet domain logging. Valid values: `disable`, `enable`.

* `sdns_ftgd_err_log` - Enable/disable FortiGuard SDNS rating error logging. Valid values: `disable`, `enable`.

* `transparent_dns_database` - Transparent DNS database zones.
* `youtube_restrict` - Set safe search for YouTube restriction level. Valid values: `strict`, `moderate`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dns_translation` block supports:

* `addr_type` - DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `dst` - IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
* `dst6` - IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
* `id` - ID.
* `netmask` - If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
* `prefix` - If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
* `src` - IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
* `src6` - IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
* `status` - Enable/disable this DNS translation entry. Valid values: `disable`, `enable`.


The `domain_filter` block supports:

* `domain_filter_table` - DNS domain filter table ID.

The `ftgd_dns` block supports:

* `filters` - Filters. The structure of `filters` block is documented below.
* `options` - FortiGuard DNS filter options. Valid values: `error-allow`, `ftgd-disable`.


The `filters` block supports:

* `action` - Action to take for DNS requests matching the category. Valid values: `monitor`, `block`.

* `category` - Category number.
* `id` - ID number.
* `log` - Enable/disable DNS filter logging for this DNS profile. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDnsfilter Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dnsfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
