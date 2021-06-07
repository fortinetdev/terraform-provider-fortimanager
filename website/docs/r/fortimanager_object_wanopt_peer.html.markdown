---
subcategory: "Object WAN-Opt"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wanopt_peer"
description: |-
  Configure WAN optimization peers.
---

# fortimanager_object_wanopt_peer
Configure WAN optimization peers.

## Example Usage

```hcl
resource "fortimanager_object_wanopt_peer" "labelname" {
  ip           = "20.0.0.20"
  peer_host_id = "ds"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ip` - Peer IP address.
* `peer_host_id` - Peer host ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{peer_host_id}}.

## Import

ObjectWanopt Peer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wanopt_peer.labelname {{peer_host_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
