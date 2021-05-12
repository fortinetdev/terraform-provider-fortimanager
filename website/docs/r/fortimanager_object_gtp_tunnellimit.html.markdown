---
subcategory: "ObjectGtp"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_gtp_tunnellimit"
description: |-
  GTP tunnel limiter.
---

# fortimanager_object_gtp_tunnellimit
GTP tunnel limiter.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `name` - Tunnel limiter name.
* `tunnel_limit` - Tunnel limit [1, 16000000].


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectGtp TunnelLimit can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_gtp_tunnellimit.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
