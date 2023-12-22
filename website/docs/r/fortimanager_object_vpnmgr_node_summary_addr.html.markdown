---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpnmgr_node_summary_addr"
description: |-
  ObjectVpnmgr NodeSummaryAddr
---

# fortimanager_object_vpnmgr_node_summary_addr
ObjectVpnmgr NodeSummaryAddr

~> This resource is a sub resource for variable `summary_addr` of resource `fortimanager_object_vpnmgr_node`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_vpnmgr_node_summary_addr" "trname" {
  node       = fortimanager_object_vpnmgr_node.trname.fosid
  priority   = 2
  seq        = 1
  depends_on = [fortimanager_object_vpnmgr_node.trname]
}

resource "fortimanager_object_vpnmgr_node" "trname" {
  fosid = 5
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `node` - Node.

* `addr` - Addr.
* `priority` - Priority.
* `seq` - Seq.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq}}.

## Import

ObjectVpnmgr NodeSummaryAddr can be imported using any of these accepted formats:
```
Set import_options = ["node=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpnmgr_node_summary_addr.labelname {{seq}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
