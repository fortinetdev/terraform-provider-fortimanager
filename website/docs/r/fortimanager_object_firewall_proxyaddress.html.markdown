---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_proxyaddress"
description: |-
  Configure web proxy address.
---

# fortimanager_object_firewall_proxyaddress
Configure web proxy address.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `header_group`: `fortimanager_object_firewall_proxyaddress_headergroup`
>- `tagging`: `fortimanager_object_firewall_proxyaddress_tagging`



## Example Usage

```hcl
resource "fortimanager_object_firewall_proxyaddress" "trname" {
  case_sensitivity = "disable"
  color            = 3
  name             = "tefv1"
  referrer         = "disable"
  type             = "url"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_image_base64` - _Image-Base64.
* `application` - SaaS application.
* `case_sensitivity` - Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.

* `category` - FortiGuard category ID.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `comment` - Optional comments.
* `header` - HTTP header name as a regular expression.
* `header_group` - Header-Group. The structure of `header_group` block is documented below.
* `header_name` - Name of HTTP header.
* `host` - Address object for the host.
* `host_regex` - Host name as a regular expression.
* `method` - HTTP request methods to be used. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `connect`.

* `name` - Address name.
* `path` - URL path as a regular expression.
* `query` - Match the query part of the URL as a regular expression.
* `referrer` - Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `disable`, `enable`.

* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `type` - Proxy address type. Valid values: `host-regex`, `url`, `category`, `method`, `ua`, `header`, `src-advanced`, `dst-advanced`.

* `ua` - Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.

* `ua_max_ver` - Maximum version of the user agent specified in dotted notation. For example, use 120 with the ua field set to "chrome" to require Google Chrome's maximum version must be 120.
* `ua_min_ver` - Minimum version of the user agent specified in dotted notation. For example, use 90.0.1 with the ua field set to "chrome" to require Google Chrome's minimum version must be 90.0.1.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable visibility of the object in the GUI. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `header_group` block supports:

* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable`, `enable`.

* `header` - HTTP header regular expression.
* `header_name` - HTTP header.
* `id` - ID.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ProxyAddress can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_proxyaddress.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
