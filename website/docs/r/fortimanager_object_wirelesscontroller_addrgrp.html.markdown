---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_addrgrp"
description: |-
  Configure the MAC address group.
---

# fortimanager_object_wirelesscontroller_addrgrp
Configure the MAC address group.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_addrgrp" "trname" {
  fosid          = 1
  default_policy = "deny"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `addresses` - Manually selected group of addresses.
* `default_policy` - Allow or block the clients with MAC addresses that are not in the group. Valid values: `deny`, `allow`.

* `fosid` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWirelessController Addrgrp can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_addrgrp.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
