---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_workflow_approvalmatrix_approver"
description: |-
  Approver.
---

# fortimanager_system_workflow_approvalmatrix_approver
Approver.

## Argument Reference


The following arguments are supported:

* `approval_matrix` - Approval Matrix.

* `member` - Member of approver.
* `seq_num` - Entry number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

System WorkflowApprovalMatrixApprover can be imported using any of these accepted formats:
```
Set import_options = ["approval_matrix=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_workflow_approvalmatrix_approver.labelname {{seq_num}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

