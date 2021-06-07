---
subcategory: ""
layout: "fortimanager"
page_title: "To Lock for Restricting Configuration Changes"
description: |-
  To Lock for Restricting Configuration Changes
---


# To Lock for Restricting Configuration Changes

Workspace enables locking ADOMs, devices, or policy packages so that an administrator can prevent other administrators from making changes to the elements that they are working in.

The provider currently supports `normal mode`. In Normal mode, ADOMs, or individual devices or policy packages must be locked before policy, object, or device changes can be made. Multiple administrators can lock devices and policy packages within a single, unlocked ADOM at the same time. When an individual device or policy package is locked, other administrators can only lock the ADOM that contains the locked device or policy package by disconnecting the administrator that locked it.

## Example

### Step 1 Set workspace_mode to normal

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

### Step 2 Wrap the resources that need to be locked with two fortimanager_exec_workspace_action resources

#### Example 1 Lock an entire ADOM

```hcl
resource "fortimanager_exec_workspace_action" "lockres" { # lock root ADOM
  scopetype      = "inherit"
  action         = "lockbegin"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
}

resource "fortimanager_object_firewall_vip" "trname" { # object resource
  scopetype  = "inherit"
  extintf    = "any"
  extip      = "1.1.1.1-2.1.1.1"
  mappedip   = ["12.1.1.1-13.1.1.1"]
  name       = "sgh"
  depends_on = [fortimanager_exec_workspace_action.lockres]
}

resource "fortimanager_exec_workspace_action" "unlockres" { # save change and unlock root ADOM
  scopetype      = "inherit"
  action         = "lockend"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
  depends_on     = [fortimanager_object_firewall_vip.trname]
}

```

#### Example 2 Lock global database

```hcl

resource "fortimanager_exec_workspace_action" "lockres" {
  scopetype      = "global"
  action         = "lockbegin"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
}

resource "fortimanager_object_firewall_vip" "trname1" {
  scopetype  = "global"
  extintf    = "any"
  extip      = "1.1.1.1-2.1.1.1"
  mappedip   = ["12.1.1.1-13.1.1.1"]
  name       = "sgh1"
  depends_on = [fortimanager_exec_workspace_action.lockres]
}

resource "fortimanager_object_firewall_vip" "trname2" {
  scopetype  = "global"
  extintf    = "any"
  extip      = "2.1.1.1-3.1.1.1"
  mappedip   = ["22.1.1.1-23.1.1.1"]
  name       = "sgh2"
  depends_on = [fortimanager_exec_workspace_action.lockres]
}

resource "fortimanager_exec_workspace_action" "unlockres" {
  scopetype      = "global"
  action         = "lockend"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
  depends_on = [
    fortimanager_object_firewall_vip.trname1,
    fortimanager_object_firewall_vip.trname2,
  ]
}

```

#### Example 3 Lock package `mypkg` in root ADOM

```hcl
resource "fortimanager_exec_workspace_action" "lockres" {
  action         = "lockbegin"

  scopetype      = "adom"
  adom           = "root"
  target         = "pkg"
  param          = "mypkg"

  comment        = ""
  force_recreate = uuid()
}

resource "fortimanager_packages_firewall_policy" "trname" {
  pkg        = "mypkg"

  action     = "accept"
  dstaddr    = "all"
  dstintf    = "any"
  name       = "s13"
  schedule   = "always"
  srcaddr    = "all"
  srcintf    = "any"
  service    = "ALL"
  status     = "enable"

  depends_on = [fortimanager_exec_workspace_action.lockres]
}

resource "fortimanager_exec_workspace_action" "unlockres" {
  action         = "lockend"

  scopetype      = "adom"
  adom           = "root"
  target         = "pkg"
  param          = "mypkg"

  comment        = ""
  force_recreate = uuid()

  depends_on     = [fortimanager_packages_firewall_policy.trname]
}

```

~> All resources that need to be executed in the lock must depend on the first fortimanager_exec_workspace_action. The second fortimanager_exec_workspace_action must contain dependencies on all resources that need to be executed in the lock. Usually these resources that need to be executed can be put in the terraform module, which can simplify the configuration.

## When the execution of a resource is unsuccessful

When a resource fails to execute, terraform will exit, which will cause the second fortimanager_exec_workspace_action to not be executed, and eventually the locked object cannot be unlocked. For example:

### Step 1 TF configuration
```hcl
provider "fortimanager" {
  hostname = "192.168.52.178"
  username = "admin"
  password = "admin"
  insecure = "true"

  scopetype = "adom"
  adom      = "root"

  logsession = true
  presession = ""
}

resource "fortimanager_exec_workspace_action" "lockres" {
  scopetype      = "global"
  action         = "lockbegin"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
}

resource "fortimanager_object_firewall_vip" "trname1" {
  scopetype  = "global"
  extintf    = "any"
  extip      = "331.1.1.1-2.1.1.1"  #<=================================
  mappedip   = ["12.1.1.1-13.1.1.1"]
  name       = "sgh1"
  depends_on = [fortimanager_exec_workspace_action.lockres]
}

resource "fortimanager_exec_workspace_action" "unlockres" {
  scopetype      = "global"
  action         = "lockend"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
  depends_on = [
    fortimanager_object_firewall_vip.trname1,
  ]
}

```

Note, we deliberately set a wrong ip for extip.

### Step 2 terraform plan

```
# terraform apply
2021/06/06 04:53:14 [WARN] Log levels other than TRACE are currently unreliable, and are supported only for backward compatibility.
  Use TF_LOG=TRACE to see Terraform's internal logs.
  ----

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # fortimanager_exec_workspace_action.lockres will be created
  + resource "fortimanager_exec_workspace_action" "lockres" {
      + action         = "lockbegin"
      + force_recreate = (known after apply)
      + id             = (known after apply)
      + scopetype      = "global"
    }

  # fortimanager_exec_workspace_action.unlockres will be created
  + resource "fortimanager_exec_workspace_action" "unlockres" {
      + action         = "lockend"
      + force_recreate = (known after apply)
      + id             = (known after apply)
      + scopetype      = "global"
    }

  # fortimanager_object_firewall_vip.trname1 will be created
  + resource "fortimanager_object_firewall_vip" "trname1" {
      + arp_reply                        = (known after apply)
      + color                            = (known after apply)
      + comment                          = (known after apply)
      + dns_mapping_ttl                  = (known after apply)
      + dynamic_sort_subtable            = "false"
      + extaddr                          = (known after apply)
      + extintf                          = "any"
      + extip                            = "331.1.1.1-2.1.1.1"
      + extport                          = (known after apply)
      + fosid                            = (known after apply)
      + gratuitous_arp_interval          = (known after apply)

      ..........

      + ssl_server_session_state_max     = (known after apply)
      + ssl_server_session_state_timeout = (known after apply)
      + ssl_server_session_state_type    = (known after apply)
      + type                             = (known after apply)
      + uuid                             = (known after apply)
      + weblogic_server                  = (known after apply)
      + websphere_server                 = (known after apply)
    }

Plan: 3 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

fortimanager_exec_workspace_action.lockres: Creating...
fortimanager_exec_workspace_action.lockres: Creation complete after 0s [id=workspaceactiongloballock]
fortimanager_object_firewall_vip.trname1: Creating...

Error: Error creating ObjectFirewallVip resource:
err -9998: firewall/vip/sgh1/ : External ip range "0.0.0.0-2.1.1.1" is invalid

  on main.tf line 24, in resource "fortimanager_object_firewall_vip" "trname1":
  24: resource "fortimanager_object_firewall_vip" "trname1" {
```

We can find that the second fortimanager_exec_workspace_action is not executed, terraform has already exited. The locked object cannot be unlocked (Global database in the example).

### Step 3 Unlock the locked object

In step 2, we see the following

```
provider "fortimanager" {
  hostname = "192.168.52.178"
  username = "admin"
  password = "admin"
  insecure = "true"

  scopetype = "adom"
  adom      = "root"

  logsession = true
  presession = ""
}
```
When `logsession` is true, in step 2, the provider will write the login session string to the presession.txt file in the current execution directory:

```
# cat presession.txt
GceyOAi/Ft5mBRW1ZRDHWXFC/mkljWiWrk5JeW//SH4RpdzYQS48fZOhhFv6Bodk03AY4tYK5j10eV6HaZmujA==
```

Then let's create a new directory and create the following TF file:
```
# mkdir unlock
# cd unlock
# cat mait.tf
provider "fortimanager" {
  hostname = "192.168.52.178"
  username = "admin"
  password = "admin"
  insecure = "true"

  scopetype = "adom"
  adom      = "root"

  logsession = true
  presession = "N9QtGOCN1TZU+cL0fI9+vzjk9jxB3wrFAuXsW7PAkHLaY4F40rn9Rm9yVI25SDVt1LpXAQScL9/XHQCaCrN2qA=="
}

resource "fortimanager_exec_workspace_action" "unlockres" {
  scopetype      = "global"
  action         = "lockend"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
}
```
Set the `presession` here to the content saved in the previous pression.txt file.

```
# terraform apply
2021/06/06 05:18:45 [WARN] Log levels other than TRACE are currently unreliable, and are supported only for backward compatibility.
  Use TF_LOG=TRACE to see Terraform's internal logs.
  ----
fortimanager_exec_workspace_action.unlockres: Refreshing state... [id=workspaceactionglobalunlock]

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # fortimanager_exec_workspace_action.unlockres will be created
  + resource "fortimanager_exec_workspace_action" "unlockres" {
      + action         = "lockend"
      + force_recreate = (known after apply)
      + id             = (known after apply)
      + scopetype      = "global"
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

fortimanager_exec_workspace_action.unlockres: Creating...
fortimanager_exec_workspace_action.unlockres: Creation complete after 0s [id=workspaceactionglobalunlock]

Apply complete! Resources: 1 added, 0 changed, 1 destroyed.

```

So far we have repaired the damaged lock/unlock pair. Then we go back to the directory in step1, fix the error in the resource, and then re-execute terraform apply

## Hints

1. The first fortimanager_exec_workspace_action resource should include `action` = `lockbegin`. When terraform apply is executed, target will be locked. The second fortimanager_exec_workspace_action resource should include `action` = `lockend`. When terraform apply is executed, changes will be saved, target will be unlocked.
2. When terraform destroy is executed, the second fortimanager_exec_workspace_action resource will lock target, and the first fortimanager_exec_workspace_action resource will save changes and unlock target.
3. Since fortimanager_exec_workspace_action is an execution resource, after it is executed once, if we try to execute the resource again, it will be recognized by terraform as no change in state, so it will not be called, which will cause an execution error. In order to solve this problem, force_recreate = uuid() need to be set to force the resource to be recreated every time it is executed.

