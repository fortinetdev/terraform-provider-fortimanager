---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_casbprofile"
description: |-
  ObjectFirewall CasbProfile
---

# fortimanager_object_firewall_casbprofile
ObjectFirewall CasbProfile

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `name` - Name.
* `saas_application` - Saas-Application. The structure of `saas_application` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `saas_application` block supports:

* `access_rule` - Access-Rule. The structure of `access_rule` block is documented below.
* `custom_control` - Custom-Control. The structure of `custom_control` block is documented below.
* `domain_control` - Domain-Control. Valid values: `disable`, `enable`.

* `domain_control_domains` - Domain-Control-Domains.
* `log` - Log. Valid values: `disable`, `enable`.

* `name` - Name.
* `safe_search` - Safe-Search. Valid values: `disable`, `enable`.

* `safe_search_control` - Safe-Search-Control.
* `tenant_control` - Tenant-Control. Valid values: `disable`, `enable`.

* `tenant_control_tenants` - Tenant-Control-Tenants.

The `access_rule` block supports:

* `action` - Action. Valid values: `block`, `monitor`, `bypass`.

* `bypass` - Bypass. Valid values: `av`, `dlp`, `web-filter`, `file-filter`, `video-filter`.

* `name` - Name.

The `custom_control` block supports:

* `name` - Name.
* `option` - Option. The structure of `option` block is documented below.

The `option` block supports:

* `name` - Name.
* `user_input` - User-Input.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall CasbProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_casbprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
