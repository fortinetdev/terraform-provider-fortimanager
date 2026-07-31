---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_networkservicedynamic"
description: |-
  Configure Dynamic Network Services.
---

# fortimanager_object_firewall_networkservicedynamic
Configure Dynamic Network Services.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `disable`, `enable`.

* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.

* `filter` - Match criteria filter.
* `fosid` - Dynamic Network Service ID.
* `name` - Dynamic Network Service name.
* `sdn` - SDN connector name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall NetworkServiceDynamic can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_networkservicedynamic.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
