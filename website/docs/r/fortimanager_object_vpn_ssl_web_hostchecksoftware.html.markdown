---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ssl_web_hostchecksoftware"
description: |-
  SSL-VPN host check software.
---

# fortimanager_object_vpn_ssl_web_hostchecksoftware
SSL-VPN host check software.

## Example Usage

```hcl
resource "fortimanager_object_vpn_ssl_web_hostchecksoftware" "trname" {
  name    = "terr-vpn-ssl-web-host-check-software"
  os_type = "macos"
  type    = "av"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `check_item_list` - Check-Item-List. The structure of `check_item_list` block is documented below.
* `guid` - Globally unique ID.
* `name` - Name.
* `os_type` - OS type. Valid values: `macos`, `windows`.

* `type` - Type. Valid values: `av`, `fw`.

* `version` - Version.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `check_item_list` block supports:

* `action` - Action. Valid values: `deny`, `require`.

* `id` - ID (0 - 4294967295).
* `md5s` - MD5 checksum.
* `target` - Target.
* `type` - Type. Valid values: `file`, `registry`, `process`.

* `version` - Version.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn SslWebHostCheckSoftware can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ssl_web_hostchecksoftware.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
