---
subcategory: "Object WAN-Opt"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wanopt_authgroup"
description: |-
  Configure WAN optimization authentication groups.
---

# fortimanager_object_wanopt_authgroup
Configure WAN optimization authentication groups.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auth_method` - Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.

* `cert` - Name of certificate to identify this peer.
* `name` - Auth-group name.
* `peer` - If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
* `peer_accept` - Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.

* `psk` - Pre-shared key used by the peers in this authentication group.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWanopt AuthGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wanopt_authgroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
