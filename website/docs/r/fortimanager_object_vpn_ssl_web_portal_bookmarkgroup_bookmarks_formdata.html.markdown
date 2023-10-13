---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ssl_web_portal_bookmarkgroup_bookmarks_formdata"
description: |-
  Form data.
---

# fortimanager_object_vpn_ssl_web_portal_bookmarkgroup_bookmarks_formdata
Form data.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `portal` - Portal.
* `bookmark_group` - Bookmark Group.
* `bookmarks` - Bookmarks.

* `name` - Name.
* `value` - Value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn SslWebPortalBookmarkGroupBookmarksFormData can be imported using any of these accepted formats:
```
Set import_options = ["portal=YOUR_VALUE", "bookmark_group=YOUR_VALUE", "bookmarks=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ssl_web_portal_bookmarkgroup_bookmarks_formdata.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
