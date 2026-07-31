---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_proxyaddress6"
description: |-
  ObjectFirewall ProxyAddress6
---

# fortimanager_object_firewall_proxyaddress6
ObjectFirewall ProxyAddress6

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `header_group`: `fortimanager_object_firewall_proxyaddress6_headergroup`
>- `tagging`: `fortimanager_object_firewall_proxyaddress6_tagging`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `application` - Application.
* `case_sensitivity` - Case-Sensitivity. Valid values: `disable`, `enable`.

* `category` - Category.
* `color` - Color.
* `comment` - Comment.
* `custom_tags` - Custom tags.
* `display_with` - Display object with first tag, all tags, or just the icon. Valid values: `all-tags`, `first-tag-only`, `icon-and-color`.

* `header` - Header.
* `header_group` - Header-Group. The structure of `header_group` block is documented below.
* `header_name` - Header-Name.
* `host` - Host.
* `host_regex` - Host-Regex.
* `llm_servers` - Llm-Servers.
* `method` - Method. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `connect`.

* `name` - Name.
* `path` - Path.
* `post_arg` - Post-Arg. Valid values: `disable`, `enable`.

* `query` - Query.
* `referrer` - Referrer. Valid values: `disable`, `enable`.

* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `type` - Type. Valid values: `host-regex`, `url`, `category`, `method`, `ua`, `header`, `src-advanced`, `dst-advanced`, `url-list`, `saas`, `response-header`.

* `ua` - Ua. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`, `ie`, `edge`.

* `ua_max_ver` - Ua-Max-Ver.
* `ua_min_ver` - Ua-Min-Ver.
* `url_list` - Url-List.
* `uuid` - Uuid.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `header_group` block supports:

* `case_sensitivity` - Case-Sensitivity. Valid values: `disable`, `enable`.

* `header` - Header.
* `header_name` - Header-Name.
* `id` - Id.

The `tagging` block supports:

* `category` - Category.
* `name` - Name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ProxyAddress6 can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_proxyaddress6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
