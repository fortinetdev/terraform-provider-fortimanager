---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_domaincontroller"
description: |-
  Configure domain controller entries.
---

# fortimanager_object_user_domaincontroller
Configure domain controller entries.

## Example Usage

```hcl
resource "fortimanager_object_user_domaincontroller" "trname" {
  domain_name = "admin"
  ip_address  = "192.168.1.1"
  name        = "terr-user-domain-controller"
  port        = 445
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ad_mode` - Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `adlds_dn` - AD LDS distinguished name. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `adlds_ip_address` - AD LDS IPv4 address. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `adlds_ip6` - AD LDS IPv6 address. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `adlds_port` - Port number of AD LDS service (default = 389). (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `dns_srv_lookup` - Enable/disable DNS service lookup. Valid values: `disable`, `enable`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `domain_name` - Domain DNS name.
* `extra_server` - Extra-Server. The structure of `extra_server` block is documented below.
* `hostname` - Hostname of the server to connect to. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `interface` - Specify outgoing interface to reach server. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `ip_address` - Domain controller IP address.
* `ip6` - Domain controller IPv6 address. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `ldap_server` - LDAP server name.
* `name` - Domain controller entry name.
* `password` - Password for specified username. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `port` - Port to be used for communication with the domain controller (default = 445).
* `replication_port` - Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `source_ip6` - FortiGate IPv6 address to be used for communication with the domain controller. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `source_port` - Source port to be used for communication with the domain controller. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `username` - User name to sign in with. Must have proper permissions for service. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `extra_server` block supports:

* `id` - Server ID.
* `ip_address` - Domain controller IP address.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `source_port` - Source port to be used for communication with the domain controller. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser DomainController can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_domaincontroller.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
