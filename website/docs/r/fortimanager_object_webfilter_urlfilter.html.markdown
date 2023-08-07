---
subcategory: "Object Webfilter"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_urlfilter"
description: |-
  Configure URL filter lists.
---

# fortimanager_object_webfilter_urlfilter
Configure URL filter lists.

## Example Usage

```hcl
resource "fortimanager_object_webfilter_urlfilter" "trname" {
  comment               = "This is a Terraform example"
  fosid                 = 1
  ip_addr_block         = "enable"
  name                  = "terr-webfilter-urlfilter"
  one_arm_ips_urlfilter = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fosid` - ID.
* `ip_addr_block` - Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `disable`, `enable`.

* `ip4_mapped_ip6` - Enable/disable matching of IPv4 mapped IPv6 URLs. Valid values: `disable`, `enable`.

* `name` - Name of URL filter list.
* `one_arm_ips_urlfilter` - Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `action` - Action to take for URL filter matches. Valid values: `exempt`, `block`, `allow`, `monitor`, `pass`.

* `antiphish_action` - Action to take for AntiPhishing matches. Valid values: `block`, `log`.

* `dns_address_family` - Resolve IPv4 address, IPv6 address, or both from DNS server. Valid values: `ipv4`, `ipv6`, `both`.

* `exempt` - If action is set to exempt, select the security profile operations that exempt URLs skip. Separate multiple options with a space. Valid values: `av`, `web-content`, `activex-java-cookie`, `dlp`, `fortiguard`, `all`, `filepattern`, `pass`, `range-block`, `antiphish`.

* `id` - Id.
* `referrer_host` - Referrer host name.
* `status` - Enable/disable this URL filter. Valid values: `disable`, `enable`.

* `type` - Filter type (simple, regex, or wildcard). Valid values: `simple`, `regex`, `wildcard`.

* `url` - URL to be filtered.
* `web_proxy_profile` - Web proxy profile.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWebfilter Urlfilter can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_urlfilter.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
