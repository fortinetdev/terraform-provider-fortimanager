---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_urlfilter_entries"
description: |-
  URL filter entries.
---

# fortimanager_object_webfilter_urlfilter_entries
URL filter entries.

~> This resource is a sub resource for variable `entries` of resource `fortimanager_object_webfilter_urlfilter`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_webfilter_urlfilter_entries" "trname" {
  fosid      = 3
  url        = "www.example.com/path/to/resource?param1=value1&param2=value2"
  depends_on = [fortimanager_object_webfilter_urlfilter.trname]
  urlfilter  = fortimanager_object_webfilter_urlfilter.trname.fosid
}

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
* `urlfilter` - Urlfilter.

* `action` - Action to take for URL filter matches. Valid values: `exempt`, `block`, `allow`, `monitor`, `pass`.

* `antiphish_action` - Action to take for AntiPhishing matches. Valid values: `block`, `log`.

* `dns_address_family` - Resolve IPv4 address, IPv6 address, or both from DNS server. Valid values: `ipv4`, `ipv6`, `both`.

* `exempt` - If action is set to exempt, select the security profile operations that exempt URLs skip. Separate multiple options with a space. Valid values: `av`, `web-content`, `activex-java-cookie`, `dlp`, `fortiguard`, `all`, `filepattern`, `pass`, `range-block`, `antiphish`.

* `fosid` - Id.
* `referrer_host` - Referrer host name.
* `status` - Enable/disable this URL filter. Valid values: `disable`, `enable`.

* `type` - Filter type (simple, regex, or wildcard). Valid values: `simple`, `regex`, `wildcard`.

* `url` - URL to be filtered.
* `web_proxy_profile` - Web proxy profile.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWebfilter UrlfilterEntries can be imported using any of these accepted formats:
```
Set import_options = ["urlfilter=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_urlfilter_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
