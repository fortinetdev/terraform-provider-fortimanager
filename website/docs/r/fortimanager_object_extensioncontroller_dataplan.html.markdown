---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extensioncontroller_dataplan"
description: |-
  FortiExtender dataplan configuration.
---

# fortimanager_object_extensioncontroller_dataplan
FortiExtender dataplan configuration.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `apn` - APN configuration.
* `auth_type` - Authentication type. Valid values: `none`, `pap`, `chap`.

* `billing_date` - Billing day of the month (1 - 31).
* `capacity` - Capacity in MB (0 - 102400000).
* `carrier` - Carrier configuration.
* `iccid` - ICCID configuration.
* `modem_id` - Dataplan's modem specifics, if any. Valid values: `all`, `modem1`, `modem2`.

* `monthly_fee` - Monthly fee of dataplan (0 - 100000, in local currency).
* `name` - FortiExtender data plan name.
* `overage` - Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.

* `password` - Password.
* `pdn` - PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.

* `preferred_subnet` - Preferred subnet mask (0 - 32).
* `private_network` - Enable/disable dataplan private network support. Valid values: `disable`, `enable`.

* `signal_period` - Signal period (600 to 18000 seconds).
* `signal_threshold` - Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
* `slot` - SIM slot configuration. Valid values: `sim1`, `sim2`.

* `type` - Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.

* `username` - Username.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectExtensionController Dataplan can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extensioncontroller_dataplan.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
