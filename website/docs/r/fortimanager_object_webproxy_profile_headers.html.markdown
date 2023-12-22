---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webproxy_profile_headers"
description: |-
  Configure HTTP forwarded requests headers.
---

# fortimanager_object_webproxy_profile_headers
Configure HTTP forwarded requests headers.

~> This resource is a sub resource for variable `headers` of resource `fortimanager_object_webproxy_profile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_webproxy_profile_headers" "trname" {
  fosid      = 1
  name       = "terr-headers"
  protocol   = ["https"]
  profile    = fortimanager_object_webproxy_profile.trname2.name
  depends_on = [fortimanager_object_webproxy_profile.trname2]
}

resource "fortimanager_object_webproxy_profile" "trname2" {
  name = "terr-web-proxy-profile2"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `action` - Action when the HTTP header is forwarded. Valid values: `add-to-request`, `add-to-response`, `remove-from-request`, `remove-from-response`.

* `add_option` - Configure options to append content to existing HTTP header or add new HTTP header. Valid values: `append`, `new-on-not-found`, `new`.

* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.

* `content` - HTTP header content.
* `dstaddr` - Destination address and address group names.
* `dstaddr6` - Destination address and address group names (IPv6).
* `fosid` - HTTP forwarded header id.
* `name` - HTTP forwarded header name.
* `protocol` - Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both). Valid values: `https`, `http`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWebProxy ProfileHeaders can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webproxy_profile_headers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
