---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_profile_urlextraction"
description: |-
  Configure URL Extraction
---

# fortimanager_object_webfilter_profile_urlextraction
Configure URL Extraction

~> This resource is a sub resource for variable `url_extraction` of resource `fortimanager_object_webfilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `redirect_header` - HTTP header name to use for client redirect on blocked requests
* `redirect_no_content` - Enable / Disable empty message-body entity in HTTP response Valid values: `disable`, `enable`.

* `redirect_url` - HTTP header value to use for client redirect on blocked requests
* `server_fqdn` - URL extraction server FQDN (fully qualified domain name)
* `status` - Enable URL Extraction Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWebfilter ProfileUrlExtraction can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_profile_urlextraction.labelname ObjectWebfilterProfileUrlExtraction
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
