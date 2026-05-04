---
subcategory: "Object GTP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_gtp_rattimeoutprofile"
description: |-
  RAT timeout profile
---

# fortimanager_object_gtp_rattimeoutprofile
RAT timeout profile

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `eutran_timeout` - Established eutran timeout in seconds (default = 0).
* `gan_timeout` - Established gan timeout in seconds (default = 0).
* `geran_timeout` - Established geran timeout in seconds (default = 0).
* `hspa_timeout` - Established hspa timeout in seconds (default = 0).
* `ltem_timeout` - Established ltem timeout in seconds (default = 0).
* `name` - RAT timeout profile name.
* `nbiot_timeout` - Established nbiot timeout in seconds (default = 0).
* `nr_timeout` - Established nr timeout in seconds (default = 0).
* `utran_timeout` - Established utran timeout in seconds (default = 0).
* `virtual_timeout` - Established virtual timeout in seconds (default = 0).
* `wlan_timeout` - Established wlan timeout in seconds (default = 0).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectGtp RatTimeoutProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_gtp_rattimeoutprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
