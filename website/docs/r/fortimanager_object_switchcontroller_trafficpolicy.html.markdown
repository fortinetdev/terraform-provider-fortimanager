---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_trafficpolicy"
description: |-
  Configure FortiSwitch traffic policy.
---

# fortimanager_object_switchcontroller_trafficpolicy
Configure FortiSwitch traffic policy.

## Example Usage

```hcl
resource "fortimanager_object_switchcontroller_trafficpolicy" "trname" {
  fosid            = 3
  name             = "terr-trafficpolicy"
  policer_status   = "enable"
  guaranteed_burst = 230000
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `cos_queue` - COS queue(0 - 7), or unset to disable.
* `description` - Description of the traffic policy.
* `guaranteed_bandwidth` - Guaranteed bandwidth in kbps (max value = 524287000).
* `guaranteed_burst` - Guaranteed burst size in bytes (max value = 4294967295).
* `fosid` - FSW Policer id
* `maximum_burst` - Maximum burst size in bytes (max value = 4294967295).
* `name` - Traffic policy name.
* `policer_status` - Enable/disable policer config on the traffic policy. Valid values: `disable`, `enable`.

* `type` - Type. Valid values: `ingress`, `egress`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController TrafficPolicy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_trafficpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
