---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_trafficclass"
description: |-
  Configure names for shaping classes.
---

# fortimanager_object_firewall_trafficclass
Configure names for shaping classes.

## Example Usage

```hcl
resource "fortimanager_object_firewall_trafficclass" "trname" {
  class_id   = 15
  class_name = "33233"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `class_id` - Class ID to be named.
* `class_name` - Define the name for this class-id.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `disable`, `enable`.

* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{class_id}}.

## Import

ObjectFirewall TrafficClass can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_trafficclass.labelname {{class_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
