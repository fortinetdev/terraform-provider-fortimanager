---
subcategory: "System Global"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_exec_workspace_action"
description: |-
  Workspace enables locking ADOMs, devices, or policy packages so that an administrator can prevent other administrators from making changes to the elements that they are working in.

---

# fortimanager_exec_workspace_action

Workspace enables locking ADOMs, devices, or policy packages so that an administrator can prevent other administrators from making changes to the elements that they are working in.


## Argument Reference


The following arguments are supported:


* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `action` - Lock or Commit/Unlock ADOMs, devices, or policy packages. Valid values: `lockbegin` lock, `lockend` Commit/Unlock.
* `target` - Lock an entire ADOM: keep the argument empty, a device: `dev`, a specific object : `obj` or a specific package: `pkg`.
* `param` - the target param will be locked or unlocked.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.
* `comment` - Comment.

See `Guides -> To Lock for Restricting Configuration Changes` for examples.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{"workspaceaction" + adomv + action + target + param}}.
