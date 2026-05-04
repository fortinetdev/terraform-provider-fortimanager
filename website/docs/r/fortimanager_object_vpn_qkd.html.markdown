---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_qkd"
description: |-
  Configure Quantum Key Distribution servers
---

# fortimanager_object_vpn_qkd
Configure Quantum Key Distribution servers

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `certificate` - Names of up to 4 certificates to offer to the KME.
* `comment` - Comment.
* `fosid` - Quantum Key Distribution ID assigned by the KME.
* `name` - Quantum Key Distribution configuration name.
* `peer` - Authenticate Quantum Key Device's certificate with the peer/peergrp.
* `port` - Port to connect to on the KME.
* `server` - IPv4, IPv6 or DNS address of the KME.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn Qkd can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_qkd.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
