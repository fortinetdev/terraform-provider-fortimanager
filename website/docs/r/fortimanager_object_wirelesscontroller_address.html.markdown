---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_address"
description: |-
  Configure the client with its MAC address.
---

# fortimanager_object_wirelesscontroller_address
Configure the client with its MAC address.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_address" "trname" {
  fosid  = 1
  mac    = "4a:7e:1e:d2:9b:86"
  policy = "allow"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `fosid` - ID.
* `mac` - MAC address.
* `policy` - Allow or block the client with this MAC address. Valid values: `deny`, `allow`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWirelessController Address can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_address.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
