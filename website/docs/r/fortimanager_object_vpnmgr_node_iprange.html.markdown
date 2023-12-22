---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpnmgr_node_iprange"
description: |-
  ObjectVpnmgr NodeIpRange
---

# fortimanager_object_vpnmgr_node_iprange
ObjectVpnmgr NodeIpRange

~> This resource is a sub resource for variable `ip_range` of resource `fortimanager_object_vpnmgr_node`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_vpnmgr_node_iprange" "trname" {
  node       = fortimanager_object_vpnmgr_node.trname.fosid
  end_ip     = "10.160.88.52"
  fosid      = 1
  start_ip   = "10.160.88.40"
  depends_on = [fortimanager_object_vpnmgr_node.trname]
}

resource "fortimanager_object_vpnmgr_node" "trname" {
  fosid = 1
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `node` - Node.

* `end_ip` - End-Ip.
* `fosid` - Id.
* `start_ip` - Start-Ip.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectVpnmgr NodeIpRange can be imported using any of these accepted formats:
```
Set import_options = ["node=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpnmgr_node_iprange.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
