---
subcategory: ""
layout: "fortimanager"
page_title: "To Lock for Restricting Configuration Changes"
description: |-
  To Lock for Restricting Configuration Changes
---


# To Lock for Restricting Configuration Changes

Workspace enables locking ADOMs, devices, or policy packages so that an administrator can prevent other administrators from making changes to the elements that they are working in.

The provider currently supports `Normal mode`. In Normal mode, ADOMs, or individual devices or policy packages must be locked before policy, object, or device changes can be made. Multiple administrators can lock devices and policy packages within a single, unlocked ADOM at the same time. When an individual device or policy package is locked, other administrators can only lock the ADOM that contains the locked device or policy package by disconnecting the administrator that locked it.

## Example:

### 1 Set workspace_mode to normal
```hcl
provider "fortimanager" {
  hostname     = "myfirewall"
  username     = "admin"
  password     = "admin"
  insecure     = "false"
  cabundlefile = "server.crt"

  scopetype = "adom"
  adom      = "root"
}

resource "fortimanager_system_global" "trname" {
  workspace_mode = "normal"
}
```
-> The resource `fortimanager_system_global` only needs to be executed once.

### 2 Wrap the resources that need to be locked with fortimanager_exec_workspace_action resources

```hcl
resource "fortimanager_exec_workspace_action" "lockres" {
  scopetype      = "global"
  adom           = "root"
  action         = "lockbegin"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
}

resource "fortimanager_object_firewall_vip" "trname1" {
  scopetype  = "inherit"
  extintf    = "any"
  extip      = "1.1.1.1-2.1.1.1"
  mappedip   = ["11.1.1.1-12.1.1.1"]
  name       = "vip1"
  depends_on = [fortimanager_exec_workspace_action.lockres]
}

resource "fortimanager_object_firewall_vip" "trname2" {
  scopetype  = "inherit"
  extintf    = "any"
  extip      = "3.1.1.1-4.1.1.1"
  mappedip   = ["13.1.1.1-14.1.1.1"]
  name       = "vip2"
  depends_on = [fortimanager_exec_workspace_action.lockres]
}

resource "fortimanager_exec_workspace_action" "unlockres" {
  scopetype      = "global"
  adom           = "root"
  action         = "lockend"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
  depends_on     = [
	  fortimanager_object_firewall_vip.trname1,
	  fortimanager_object_firewall_vip.trname2,
  ]
}

```

#### Hints
1. The beginning resource fortimanager_exec_workspace_action should set `action` = `lockbegin`. When terraform apply is executed, target will be locked.
2. The ending resource fortimanager_exec_workspace_action should set `action` = `lockend`. When terraform apply is executed, changes will be saved, target will be unlocked.
3. When terraform destroy is executed, the ending resource fortimanager_exec_workspace_action will lock the target, and the beginning resource fortimanager_exec_workspace_action will save changes and unlock target.
4. If an error occurs during execution after locking, currentlythe changes will be discarded, and you need to unlock it manually, for example: `diagnose system admin-session kill <sid>`.
5. All resources that need to be executed in the lock must add a dependency on the beginning fortimanager_exec_workspace_action. The dependency on all the above resources needs to be added to the ending resource fortimanager_exec_workspace_action. Usually the locked resources can be put in the terraform module, which can simplify the configuration.
