---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_gtp_policy"
description: |-
  Policy.
---

# fortimanager_object_firewall_gtp_policy
Policy.

~> This resource is a sub resource for variable `policy` of resource `fortimanager_object_firewall_gtp`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `gtp` - Gtp.

* `action` - Action. Valid values: `allow`, `deny`.

* `apn_sel_mode` - APN selection mode. Valid values: `ms`, `net`, `vrf`.

* `apnmember` - APN member.
* `fosid` - ID.
* `imei` - IMEI pattern.
* `imsi` - IMSI prefix.
* `imsi_prefix` - IMSI prefix.
* `max_apn_restriction` - Maximum APN restriction value. Valid values: `all`, `public-1`, `public-2`, `private-1`, `private-2`.

* `messages` - GTP messages. Valid values: `create-req`, `create-res`, `update-req`, `update-res`.

* `msisdn` - MSISDN prefix.
* `msisdn_prefix` - MSISDN prefix.
* `rai` - RAI pattern.
* `rat_type` - RAT Type. Valid values: `any`, `utran`, `geran`, `wlan`, `gan`, `hspa`, `eutran`, `virtual`, `nbiot`.

* `uli` - ULI pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall GtpPolicy can be imported using any of these accepted formats:
```
Set import_options = ["gtp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_gtp_policy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
