---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ipsec_fec_mappings"
description: |-
  FEC redundancy mapping table.
---

# fortimanager_object_vpn_ipsec_fec_mappings
FEC redundancy mapping table.

~> This resource is a sub resource for variable `mappings` of resource `fortimanager_object_vpn_ipsec_fec`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_vpn_ipsec_fec_mappings" "trname" {
  fec        = fortimanager_object_vpn_ipsec_fec.trname.name
  seqno      = 2
  depends_on = [fortimanager_object_vpn_ipsec_fec.trname]
}

resource "fortimanager_object_vpn_ipsec_fec" "trname" {
  name = "terr-fec"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `fec` - Fec.

* `bandwidth_bi_threshold` - Apply FEC parameters when available bi-bandwidth is &gt;= threshold (kbps, 0 means no threshold).
* `bandwidth_down_threshold` - Apply FEC parameters when available down bandwidth is &gt;= threshold (kbps, 0 means no threshold).
* `bandwidth_up_threshold` - Apply FEC parameters when available up bandwidth is &gt;= threshold (kbps, 0 means no threshold).
* `base` - Number of base FEC packets (1 - 20).
* `latency_threshold` - Apply FEC parameters when latency is &lt;= threshold (0 means no threshold).
* `packet_loss_threshold` - Apply FEC parameters when packet loss is &gt;= threshold (0 - 100, 0 means no threshold).
* `redundant` - Number of redundant FEC packets (1 - 5).
* `seqno` - Sequence number (1 - 64).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seqno}}.

## Import

ObjectVpn IpsecFecMappings can be imported using any of these accepted formats:
```
Set import_options = ["fec=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ipsec_fec_mappings.labelname {{seqno}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
