---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ipsec_fec_mappings_move"
description: |-
  FEC redundancy mapping table.
---

# fortimanager_object_vpn_ipsec_fec_mappings_move
FEC redundancy mapping table.

## Example Usage

```hcl
resource "fortimanager_object_vpn_ipsec_fec_mappings_move" "trname" {
  fec        = fortimanager_object_vpn_ipsec_fec.trname2.name
  mappings   = fortimanager_object_vpn_ipsec_fec_mappings.trname1.seqno
  target     = fortimanager_object_vpn_ipsec_fec_mappings.trname2.seqno
  option     = "after"
  depends_on = [fortimanager_object_vpn_ipsec_fec_mappings.trname1, fortimanager_object_vpn_ipsec_fec_mappings.trname2]
}

resource "fortimanager_object_vpn_ipsec_fec_mappings" "trname2" {
  fec        = fortimanager_object_vpn_ipsec_fec.trname2.name
  seqno      = 4
  depends_on = [fortimanager_object_vpn_ipsec_fec.trname2]
}

resource "fortimanager_object_vpn_ipsec_fec_mappings" "trname1" {
  fec        = fortimanager_object_vpn_ipsec_fec.trname2.name
  seqno      = 3
  depends_on = [fortimanager_object_vpn_ipsec_fec.trname2]
}

resource "fortimanager_object_vpn_ipsec_fec" "trname2" {
  name = "terr-fec"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `fec` - Fec.
* `mappings` - Mappings.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seqno}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the mappings changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of mapping.
