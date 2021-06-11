---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_replacemsggroup"
description: |-
  Configure replacement message groups.
---

# fortimanager_object_system_replacemsggroup
Configure replacement message groups.

## Example Usage

```hcl
resource "fortimanager_object_system_replacemsggroup" "trname" {
  comment = "terraform-comment"
  name    = "terraform-tefv"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `admin` - Admin. The structure of `admin` block is documented below.
* `alertmail` - Alertmail. The structure of `alertmail` block is documented below.
* `auth` - Auth. The structure of `auth` block is documented below.
* `comment` - Comment.
* `custom_message` - Custom-Message. The structure of `custom_message` block is documented below.
* `device_detection_portal` - Device-Detection-Portal. The structure of `device_detection_portal` block is documented below.
* `fortiguard_wf` - Fortiguard-Wf. The structure of `fortiguard_wf` block is documented below.
* `ftp` - Ftp. The structure of `ftp` block is documented below.
* `group_type` - Group type. Valid values: `default`, `utm`, `auth`, `ec`, `captive-portal`.

* `http` - Http. The structure of `http` block is documented below.
* `icap` - Icap. The structure of `icap` block is documented below.
* `mail` - Mail. The structure of `mail` block is documented below.
* `nac_quar` - Nac-Quar. The structure of `nac_quar` block is documented below.
* `name` - Group name.
* `nntp` - Nntp. The structure of `nntp` block is documented below.
* `spam` - Spam. The structure of `spam` block is documented below.
* `sslvpn` - Sslvpn. The structure of `sslvpn` block is documented below.
* `traffic_quota` - Traffic-Quota. The structure of `traffic_quota` block is documented below.
* `utm` - Utm. The structure of `utm` block is documented below.
* `webproxy` - Webproxy. The structure of `webproxy` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `admin` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `alertmail` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `auth` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `custom_message` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `device_detection_portal` block supports:

* `buffer` - Buffer.
* `format` - Format. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Msg-Type.

The `fortiguard_wf` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `ftp` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `http` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `icap` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `mail` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `nac_quar` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `nntp` block supports:

* `buffer` - Buffer.
* `format` - Format. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Msg-Type.

The `spam` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `sslvpn` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `traffic_quota` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `utm` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `webproxy` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem ReplacemsgGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_replacemsggroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
