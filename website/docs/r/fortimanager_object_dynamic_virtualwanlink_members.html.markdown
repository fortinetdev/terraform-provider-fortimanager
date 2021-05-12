---
subcategory: "ObjectDynamic"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dynamic_virtualwanlink_members"
description: |-
  ObjectDynamic VirtualWanLinkMembers
---

# fortimanager_object_dynamic_virtualwanlink_members
ObjectDynamic VirtualWanLinkMembers

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `cost` - Cost.
* `detect_failtime` - Detect-Failtime.
* `detect_http_get` - Detect-Http-Get.
* `detect_http_match` - Detect-Http-Match.
* `detect_http_port` - Detect-Http-Port.
* `detect_interval` - Detect-Interval.
* `detect_protocol` - Detect-Protocol. Valid values: `ping`, `tcp-echo`, `udp-echo`, `http`.

* `detect_recoverytime` - Detect-Recoverytime.
* `detect_server` - Detect-Server.
* `detect_timeout` - Detect-Timeout.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `gateway` - Gateway.
* `gateway6` - Gateway6.
* `ingress_spillover_threshold` - Ingress-Spillover-Threshold.
* `interface` - Interface.
* `name` - Name.
* `priority` - Priority.
* `source` - Source.
* `source6` - Source6.
* `spillover_threshold` - Spillover-Threshold.
* `status` - Status. Valid values: `disable`, `enable`.

* `volume_ratio` - Volume-Ratio.
* `weight` - Weight.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `comment` - Comment.
* `cost` - Cost.
* `detect_failtime` - Detect-Failtime.
* `detect_http_get` - Detect-Http-Get.
* `detect_http_match` - Detect-Http-Match.
* `detect_http_port` - Detect-Http-Port.
* `detect_interval` - Detect-Interval.
* `detect_protocol` - Detect-Protocol. Valid values: `ping`, `tcp-echo`, `udp-echo`, `http`.

* `detect_recoverytime` - Detect-Recoverytime.
* `detect_server` - Detect-Server.
* `detect_timeout` - Detect-Timeout.
* `gateway` - Gateway.
* `gateway6` - Gateway6.
* `ingress_spillover_threshold` - Ingress-Spillover-Threshold.
* `priority` - Priority.
* `source` - Source.
* `source6` - Source6.
* `spillover_threshold` - Spillover-Threshold.
* `status` - Status. Valid values: `disable`, `enable`.

* `volume_ratio` - Volume-Ratio.
* `weight` - Weight.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDynamic VirtualWanLinkMembers can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dynamic_virtualwanlink_members.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
