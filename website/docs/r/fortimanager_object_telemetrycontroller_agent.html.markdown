---
subcategory: "Object Telemetry Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_telemetrycontroller_agent"
description: |-
  Configure FortiTelemetry agents managed by a FortiGate unit.
---

# fortimanager_object_telemetrycontroller_agent
Configure FortiTelemetry agents managed by a FortiGate unit.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `agent_id` - Agent ID.
* `agent_profile` - Name of an existing agent profile.
* `alias` - Alias used in display for ease of distinguishing agents.
* `authz` - Authorization status of this agent. Valid values: `authorized`, `rejected`, `unauthorized`.

* `comment` - Comment.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{agent_id}}.

## Import

ObjectTelemetryController Agent can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_telemetrycontroller_agent.labelname {{agent_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
