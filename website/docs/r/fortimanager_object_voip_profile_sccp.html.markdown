---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_voip_profile_sccp"
description: |-
  SCCP.
---

# fortimanager_object_voip_profile_sccp
SCCP.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `block_mcast` - Enable/disable block multicast RTP connections. Valid values: `disable`, `enable`.

* `log_call_summary` - Enable/disable log summary of SCCP calls. Valid values: `disable`, `enable`.

* `log_violations` - Enable/disable logging of SCCP violations. Valid values: `disable`, `enable`.

* `max_calls` - Maximum calls per minute per SCCP client (max 65535).
* `status` - Enable/disable SCCP. Valid values: `disable`, `enable`.

* `verify_header` - Enable/disable verify SCCP header content. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectVoip ProfileSccp can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_voip_profile_sccp.labelname ObjectVoipProfileSccp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
