---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_gtp_policyv2"
description: |-
  Apply allow or deny action to each GTPv2-c packet.
---

# fortimanager_object_firewall_gtp_policyv2
Apply allow or deny action to each GTPv2-c packet.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `gtp` - Gtp.

* `action` - Action. Valid values: `deny`, `allow`.

* `apn_sel_mode` - APN selection mode. Valid values: `ms`, `net`, `vrf`.

* `apnmember` - APN member.
* `fosid` - ID.
* `imsi_prefix` - IMSI prefix.
* `max_apn_restriction` - Maximum APN restriction value. Valid values: `all`, `public-1`, `public-2`, `private-1`, `private-2`.

* `mei` - MEI pattern.
* `messages` - GTP messages. Valid values: `create-ses-req`, `create-ses-res`, `modify-bearer-req`, `modify-bearer-res`.

* `msisdn_prefix` - MSISDN prefix.
* `rat_type` - RAT Type. Valid values: `any`, `utran`, `geran`, `wlan`, `gan`, `hspa`, `eutran`, `virtual`, `nbiot`, `ltem`, `nr`.

* `uli` - GTPv2 ULI patterns (in order of CGI SAI RAI TAI ECGI LAI).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall GtpPolicyV2 can be imported using any of these accepted formats:
```
Set import_options = ["gtp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_gtp_policyv2.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
