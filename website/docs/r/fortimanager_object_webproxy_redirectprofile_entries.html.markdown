---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webproxy_redirectprofile_entries"
description: |-
  ObjectWebProxy RedirectProfileEntries
---

# fortimanager_object_webproxy_redirectprofile_entries
ObjectWebProxy RedirectProfileEntries

~> This resource is a sub resource for variable `entries` of resource `fortimanager_object_webproxy_redirectprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `redirect_profile` - Redirect Profile.

* `fosid` - Id.
* `redirect_code` - Redirect-Code. Valid values: `auto`, `301`, `302`, `303`, `307`, `308`.

* `redirect_url` - Redirect-Url.
* `type` - Type. Valid values: `wildcard`, `regex`, `simple`.

* `url` - Url.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWebProxy RedirectProfileEntries can be imported using any of these accepted formats:
```
Set import_options = ["redirect_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webproxy_redirectprofile_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
