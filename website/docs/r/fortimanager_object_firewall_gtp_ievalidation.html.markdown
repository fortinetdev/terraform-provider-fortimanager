---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_gtp_ievalidation"
description: |-
  IE validation.
---

# fortimanager_object_firewall_gtp_ievalidation
IE validation.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `gtp` - Gtp.

* `apn_restriction` - Validate APN restriction. Valid values: `disable`, `enable`.

* `charging_id` - Validate charging ID. Valid values: `disable`, `enable`.

* `charging_gateway_addr` - Validate charging gateway address. Valid values: `disable`, `enable`.

* `end_user_addr` - Validate end user address. Valid values: `disable`, `enable`.

* `gsn_addr` - Validate GSN address. Valid values: `disable`, `enable`.

* `imei` - Validate IMEI(SV). Valid values: `disable`, `enable`.

* `imsi` - Validate IMSI. Valid values: `disable`, `enable`.

* `mm_context` - Validate MM context. Valid values: `disable`, `enable`.

* `ms_tzone` - Validate MS time zone. Valid values: `disable`, `enable`.

* `ms_validated` - Validate MS validated. Valid values: `disable`, `enable`.

* `msisdn` - Validate MSISDN. Valid values: `disable`, `enable`.

* `nsapi` - Validate NSAPI. Valid values: `disable`, `enable`.

* `pdp_context` - Validate PDP context. Valid values: `disable`, `enable`.

* `qos_profile` - Validate Quality of Service(QoS) profile. Valid values: `disable`, `enable`.

* `rai` - Validate RAI. Valid values: `disable`, `enable`.

* `rat_type` - Validate RAT type. Valid values: `disable`, `enable`.

* `reordering_required` - Validate re-ordering required. Valid values: `disable`, `enable`.

* `selection_mode` - Validate selection mode. Valid values: `disable`, `enable`.

* `uli` - Validate user location information. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall GtpIeValidation can be imported using any of these accepted formats:
```
Set import_options = ["gtp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_gtp_ievalidation.labelname ObjectFirewallGtpIeValidation
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
