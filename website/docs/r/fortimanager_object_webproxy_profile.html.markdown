---
subcategory: "Object Web-Proxy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webproxy_profile"
description: |-
  Configure web proxy profiles.
---

# fortimanager_object_webproxy_profile
Configure web proxy profiles.

## Example Usage

```hcl
resource "fortimanager_object_webproxy_profile" "trname" {
  header_client_ip              = "pass"
  header_front_end_https        = "add"
  header_via_request            = "add"
  header_via_response           = "pass"
  header_x_authenticated_groups = "add"
  header_x_authenticated_user   = "add"
  header_x_forwarded_for        = "add"
  log_header_change             = "disable"
  name                          = "terr-web-proxy-profile"
  strip_encoding                = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `header_client_ip` - Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_front_end_https` - Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_via_request` - Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_via_response` - Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_x_authenticated_groups` - Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_x_authenticated_user` - Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_x_forwarded_client_cert` - Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_x_forwarded_for` - Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `headers` - Headers. The structure of `headers` block is documented below.
* `log_header_change` - Enable/disable logging HTTP header changes. Valid values: `disable`, `enable`.

* `name` - Profile name.
* `strip_encoding` - Enable/disable stripping unsupported encoding from the request header. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `headers` block supports:

* `action` - Action when the HTTP header is forwarded. Valid values: `add-to-request`, `add-to-response`, `remove-from-request`, `remove-from-response`.

* `add_option` - Configure options to append content to existing HTTP header or add new HTTP header. Valid values: `append`, `new-on-not-found`, `new`.

* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.

* `content` - HTTP header content.
* `dstaddr` - Destination address and address group names.
* `dstaddr6` - Destination address and address group names (IPv6).
* `id` - HTTP forwarded header id.
* `name` - HTTP forwarded header name.
* `protocol` - Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both). Valid values: `https`, `http`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWebProxy Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webproxy_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
