---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_ptp_profile"
description: |-
  Global PTP profile.
---

# fortimanager_object_switchcontroller_ptp_profile
Global PTP profile.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `description` - Description.
* `domain` - Configure PTP domain value (0 - 255, default = 254).
* `mode` - Select PTP mode. Valid values: `transparent-e2e`, `transparent-p2p`.

* `name` - Profile name.
* `pdelay_req_interval` - Configure PTP peer delay request interval. Valid values: `1sec`, `2sec`, `4sec`, `8sec`, `16sec`, `32sec`.

* `ptp_profile` - Configure PTP power profile. Valid values: `C37.238-2017`.

* `transport` - Configure PTP transport mode. Valid values: `l2-mcast`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController PtpProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_ptp_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
