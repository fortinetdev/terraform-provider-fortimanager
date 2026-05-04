---
subcategory: "Object ZTNA"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ztna_webportalbookmark_llmsecureproxy"
description: |-
  ObjectZtna WebPortalBookmarkLlmSecureProxy
---

# fortimanager_object_ztna_webportalbookmark_llmsecureproxy
ObjectZtna WebPortalBookmarkLlmSecureProxy

~> This resource is a sub resource for variable `llm_secure_proxy` of resource `fortimanager_object_ztna_webportalbookmark`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `web_portal_bookmark` - Web Portal Bookmark.

* `all_llm_servers` - All-Llm-Servers. Valid values: `disable`, `enable`.

* `llm_servers` - Llm-Servers.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectZtna WebPortalBookmarkLlmSecureProxy can be imported using any of these accepted formats:
```
Set import_options = ["web_portal_bookmark=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_ztna_webportalbookmark_llmsecureproxy.labelname ObjectZtnaWebPortalBookmarkLlmSecureProxy
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
