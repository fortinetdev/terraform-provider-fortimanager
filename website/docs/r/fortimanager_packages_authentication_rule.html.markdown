---
subcategory: "Packages"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_authentication_rule"
description: |-
  Configure Authentication Rules.
---

# fortimanager_packages_authentication_rule
Configure Authentication Rules.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `active_auth_method` - Select an active authentication method.
* `comments` - Comment.
* `ip_based` - Enable/disable IP-based authentication. When enabled, previously authenticated users from the same IP address will be exempted. Valid values: `disable`, `enable`.

* `name` - Authentication rule name.
* `protocol` - Authentication is required for the selected protocol (default = HTTP). Valid values: `http`, `ftp`, `socks`, `ssh`.

* `srcaddr` - Authentication is required for the selected IPv4 source address.
* `srcaddr6` - Authentication is required for the selected IPv6 source address.
* `sso_auth_method` - Select a single-sign on (SSO) authentication method.
* `status` - Enable/disable this authentication rule. Valid values: `disable`, `enable`.

* `transaction_based` - Enable/disable transaction based authentication (default = disable). Valid values: `disable`, `enable`.

* `web_auth_cookie` - Enable/disable Web authentication cookies (default = disable). Valid values: `disable`, `enable`.

* `web_portal` - Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Packages AuthenticationRule can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_authentication_rule.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
