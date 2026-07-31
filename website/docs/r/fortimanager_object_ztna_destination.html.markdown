---
subcategory: "Object ZTNA"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ztna_destination"
description: |-
  Configure ZTNA destination.
---

# fortimanager_object_ztna_destination
Configure ZTNA destination.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `address` - Address or address group of the ZTNA destination.
* `conn_type` - Connection type. Valid values: `traffic-forwarding`, `ssh`.

* `domain` - Wildcard domain name of the ZTNA destination.
* `external_auth` - Enable/disable use of external browser as user-agent for SAML user authentication. Valid values: `disable`, `enable`.

* `mappedport` - Port for communicating with the real server.
* `name` - Destination name.
* `protocol` - Protocol type based on IANA numbers. Valid values: `TCP`, `UDP`, `ALL`.

* `saas_application` - SaaS application controlled by this ZTNA destination.
* `ssh_client_cert` - Configure access-proxy SSH client certificate profile.
* `ssh_host_key` - Configure host keys (one or more may be configured).
* `ssh_host_key_validation` - Enable/disable SSH host key validation. Valid values: `disable`, `enable`.

* `tunnel_encryption` - Tunnel encryption. Valid values: `disable`, `enable`.

* `type` - ZTNA destination type. Valid values: `on-premise`, `saas`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectZtna Destination can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_ztna_destination.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
