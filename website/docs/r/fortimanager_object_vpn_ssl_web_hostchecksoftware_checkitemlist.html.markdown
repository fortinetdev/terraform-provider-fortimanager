---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ssl_web_hostchecksoftware_checkitemlist"
description: |-
  Check item list.
---

# fortimanager_object_vpn_ssl_web_hostchecksoftware_checkitemlist
Check item list.

~> This resource is a sub resource for variable `check_item_list` of resource `fortimanager_object_vpn_ssl_web_hostchecksoftware`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_vpn_ssl_web_hostchecksoftware_checkitemlist" "trname" {
  host_check_software = fortimanager_object_vpn_ssl_web_hostchecksoftware.trname.name
  action              = "deny"
  fosid               = 1
  type                = "file"
  depends_on          = [fortimanager_object_vpn_ssl_web_hostchecksoftware.trname]
}

resource "fortimanager_object_vpn_ssl_web_hostchecksoftware" "trname" {
  name = "terr-software"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `host_check_software` - Host Check Software.

* `action` - Action. Valid values: `deny`, `require`.

* `fosid` - ID (0 - 4294967295).
* `md5s` - MD5 checksum.
* `target` - Target.
* `type` - Type. Valid values: `file`, `registry`, `process`.

* `version` - Version.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectVpn SslWebHostCheckSoftwareCheckItemList can be imported using any of these accepted formats:
```
Set import_options = ["host_check_software=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ssl_web_hostchecksoftware_checkitemlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
