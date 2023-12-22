---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpnmgr_node_protected_subnet"
description: |-
  ObjectVpnmgr NodeProtectedSubnet
---

# fortimanager_object_vpnmgr_node_protected_subnet
ObjectVpnmgr NodeProtectedSubnet

~> This resource is a sub resource for variable `protected_subnet` of resource `fortimanager_object_vpnmgr_node`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_vpnmgr_node_protected_subnet" "trname" {
  node       = fortimanager_object_vpnmgr_node.trname.fosid
  seq        = 1
  depends_on = [fortimanager_object_vpnmgr_node.trname]
}

resource "fortimanager_object_vpnmgr_node" "trname" {
  fosid = 3
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `node` - Node.

* `addr` - Addr.
* `seq` - Seq.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq}}.

## Import

ObjectVpnmgr NodeProtectedSubnet can be imported using any of these accepted formats:
```
Set import_options = ["node=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpnmgr_node_protected_subnet.labelname {{seq}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
