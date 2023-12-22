---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ipsec_fec"
description: |-
  Configure Forward Error Correction (FEC) mapping profiles.
---

# fortimanager_object_vpn_ipsec_fec
Configure Forward Error Correction (FEC) mapping profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`mappings`: `fortimanager_object_vpn_ipsec_fec_mappings`



## Example Usage

```hcl
resource "fortimanager_object_vpn_ipsec_fec" "trname" {
  name = "terr-fec"
  mappings {
    bandwidth_bi_threshold = 12
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `mappings` - Mappings. The structure of `mappings` block is documented below.
* `name` - Profile name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `mappings` block supports:

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
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn IpsecFec can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ipsec_fec.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
