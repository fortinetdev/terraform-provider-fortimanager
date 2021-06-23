---
subcategory: "Packages Authentication"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_authentication_rule"
description: |-
  Configure Authentication Rules.
---

# fortimanager_packages_authentication_rule
Configure Authentication Rules.

## Example Usage

```hcl
resource "fortimanager_packages_authentication_rule" "labelname" {
  ip_based = "enable"
  name     = "ss"
  pkg      = "default"
  protocol = "http"
  srcaddr  = "all"
  status   = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `active_auth_method` - Select an active authentication method.
* `comments` - Comment.
* `dstaddr` - Select an IPv4 destination address from available options. Required for web proxy authentication.
* `dstaddr6` - Select an IPv6 destination address from available options. Required for web proxy authentication.
* `ip_based` - Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `disable`, `enable`.

* `name` - Authentication rule name.
* `protocol` - Select the protocol to use for authentication (default = http). Users connect to the FortiGate using this protocol and are asked to authenticate. Valid values: `http`, `ftp`, `socks`, `ssh`.

* `srcaddr` - Select an IPv4 source address from available options. Required for web proxy authentication.
* `srcaddr6` - Select an IPv6 source address. Required for web proxy authentication.
* `srcintf` - Incoming (ingress) interface.
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
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_authentication_rule.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
