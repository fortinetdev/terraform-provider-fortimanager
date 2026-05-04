---
subcategory: "Object ZTNA"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ztna_serviceconnector"
description: |-
  ObjectZtna ServiceConnector
---

# fortimanager_object_ztna_serviceconnector
ObjectZtna ServiceConnector

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `certificate` - Certificate.
* `connection_mode` - Connection-Mode. Valid values: `forward`, `reverse`.

* `encryption` - Encryption. Valid values: `disable`, `enable`.

* `forward_address` - Forward-Address.
* `forward_destination_cn` - Forward-Destination-Cn.
* `forward_port` - Forward-Port.
* `health_check_interval` - Health-Check-Interval.
* `log` - Log. Valid values: `disable`, `enable`.

* `name` - Name.
* `relay_dev_info` - Relay-Dev-Info. Valid values: `disable`, `enable`.

* `relay_user_info` - Relay-User-Info. Valid values: `disable`, `enable`.

* `ssl_max_version` - Ssl-Max-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Ssl-Min-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `status` - Status. Valid values: `disable`, `enable`.

* `trusted_ca` - Trusted-Ca.
* `url_map` - Url-Map.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectZtna ServiceConnector can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_ztna_serviceconnector.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
