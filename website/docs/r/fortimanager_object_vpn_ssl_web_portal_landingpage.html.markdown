---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ssl_web_portal_landingpage"
description: |-
  Landing page options.
---

# fortimanager_object_vpn_ssl_web_portal_landingpage
Landing page options.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `portal` - Portal.

* `form_data` - Form-Data. The structure of `form_data` block is documented below.
* `logout_url` - Landing page log out URL.
* `sso` - Single sign-on. Valid values: `disable`, `static`, `auto`.

* `sso_credential` - Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.

* `sso_password` - SSO password.
* `sso_username` - SSO user name.
* `url` - Landing page URL.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `form_data` block supports:

* `name` - Name.
* `value` - Value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectVpn SslWebPortalLandingPage can be imported using any of these accepted formats:
```
Set import_options = ["portal=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ssl_web_portal_landingpage.labelname ObjectVpnSslWebPortalLandingPage
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
