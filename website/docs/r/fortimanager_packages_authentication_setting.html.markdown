---
subcategory: "Packages Authentication"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_authentication_setting"
description: |-
  Configure authentication setting.
---

# fortimanager_packages_authentication_setting
Configure authentication setting.

## Example Usage

```hcl
resource "fortimanager_packages_authentication_setting" "labelname" {
  auth_https              = "enable"
  captive_portal_port     = 7830
  captive_portal_ssl_port = 7831
  captive_portal_type     = "fqdn"
  pkg                     = "default"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `active_auth_scheme` - Active authentication method (scheme name).
* `auth_https` - Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `disable`, `enable`.

* `captive_portal` - Captive portal host name.
* `captive_portal_ip` - Captive portal IP address.
* `captive_portal_ip6` - Captive portal IPv6 address.
* `captive_portal_port` - Captive portal port number (1 - 65535, default = 7830).
* `captive_portal_ssl_port` - Captive portal SSL port number (1 - 65535, default = 7831).
* `captive_portal_type` - Captive portal type. Valid values: `fqdn`, `ip`.

* `captive_portal6` - IPv6 captive portal host name.
* `dev_range` - Address range for the IP based device query.
* `rewrite_https_port` - Rewrite-Https-Port.
* `sso_auth_scheme` - Single-Sign-On authentication method (scheme name).
* `user_cert_ca` - CA certificate used for client certificate verification.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Packages AuthenticationSetting can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_authentication_setting.labelname PackagesAuthenticationSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
