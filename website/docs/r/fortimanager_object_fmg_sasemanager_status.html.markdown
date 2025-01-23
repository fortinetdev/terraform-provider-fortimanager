---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fmg_sasemanager_status"
description: |-
  ObjectFmg SaseManagerStatus
---

# fortimanager_object_fmg_sasemanager_status
ObjectFmg SaseManagerStatus

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `forticlient_ver` - Forticlient-Ver.
* `forticloud_id` - Forticloud-Id.
* `license_type` - License-Type. Valid values: `standard_license`, `advanced_license`, `comprehensive_license`.

* `spa_hubs` - Spa-Hubs.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFmg SaseManagerStatus can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fmg_sasemanager_status.labelname ObjectFmgSaseManagerStatus
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
