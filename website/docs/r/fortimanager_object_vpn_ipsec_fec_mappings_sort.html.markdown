
---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ipsec_fec_mappings_sort"
description: |-
  Sort for object_vpn ipsec_fec_mappings.
---

# fortimanager_object_vpn_ipsec_fec_mappings_sort
Sort for object_vpn ipsec_fec_mappings.

```hcl
resource "fortimanager_object_vpn_ipsec_fec_mappings_sort" "test" {
	sortby        = "seqno"
	sortdirection = "descending"
}
		
```

The following arguments are supported:

* `sortby` - (Required) Sort security policies by the value, it currently supports "seqno".
* `sortdirection` - (Required) Sort dirction, supports "ascending" and "descending".
* `manual_order` - Manual order for resources you want to be sorted. Content must be the category of `sortby`. Available when `sortdirection` set to "manual".
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created. If set to 'True', then the value on state file will always be set to False to trigger the re-create operation for every terraform plan/apply. Otherwise, will set to the given value. 
* `comment` - Comment.

The following attributes are exported:

* `id` - an identifier for the resource.
* `sortby` - Sort security policies by the value, it currently supports "seqno".
* `sortdirection` - Sort dirction, supports "ascending" and "descending".
* `manual_order` - Manual order for resources you want to be sorted. Content must be the category of `sortby`. Available when `sortdirection` set to "manual".
* `status` - The parameter is read-only, it is used to indicate whether the order of the resource on FortiManager matches the terraform configuration, if matched, the value is empty(that means ""), otherwise the value is "unsorted", usually the modification outside of the terrform will cause that the status value is "unsorted".
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created. If set to 'True', then the value on state file will always be set to False to trigger the re-create operation for every terraform plan/apply. Otherwise, will set to the given value. 
* `comment` - Comment.
* `state_list` - The parameter is read-only, it is used to get the latest entry list. It will be updated after each terraform apply or terraform refresh.


~> **Note** Since the order changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state.

!> **Warning** This resource involves the priority shift of many entries, when using terraform apply to apply this resource, please ensure that the FortiManager is offline to avoid business interruption or unnecessary security risks.


