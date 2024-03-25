---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_workflow_approvalmatrix"
description: |-
  workflow approval matrix.
---

# fortimanager_system_workflow_approvalmatrix
workflow approval matrix.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `approver`: `fortimanager_system_workflow_approvalmatrix_approver`



## Example Usage

```hcl
resource "fortimanager_system_workflow_approvalmatrix" "trname" {
  adom_name = "terraform"
  notify    = "terraform-notify"
}
```

## Argument Reference


The following arguments are supported:


* `adom_name` - Adom Name
* `approver` - Approver. The structure of `approver` block is documented below.
* `mail_server` - Notify mail server id.
* `notify` - Notify users
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `approver` block supports:

* `member` - Member of approver.
* `seq_num` - Entry number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{adom_name}}.

## Import

System WorkflowApprovalMatrix can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_workflow_approvalmatrix.labelname {{adom_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

