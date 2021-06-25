---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_decryptedtrafficmirror"
description: |-
  Configure decrypted traffic mirror.
---

# fortimanager_object_firewall_decryptedtrafficmirror
Configure decrypted traffic mirror. Applies to `Controlled FortiOS >= 6.4`.

## Example Usage

```hcl
resource "fortimanager_object_firewall_decryptedtrafficmirror" "trname" {
  dstmac         = "ff:ff:ff:ff:ff:ff"
  interface      = "any"
  name           = "terr-firewall-mirror"
  traffic_source = "both"
  traffic_type   = ["ssh", "ssl"]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `dstmac` - Set destination MAC address for mirrored traffic.
* `interface` - Decrypted traffic mirror interface
* `name` - Name.
* `traffic_source` - Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.

* `traffic_type` - Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall DecryptedTrafficMirror can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_decryptedtrafficmirror.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
