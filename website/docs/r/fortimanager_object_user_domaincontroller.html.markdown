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

* `adlds_dn` - AD LDS distinguished name.
* `adlds_ip_address` - AD LDS IPv4 address.
* `adlds_ip6` - AD LDS IPv6 address.
* `adlds_port` - Port number of AD LDS service (default = 389).
* `dns_srv_lookup` - Enable/disable DNS service lookup. Valid values: `disable`, `enable`.

* `domain_name` - Domain DNS name.
* `extra_server` - Extra-Server. The structure of `extra_server` block is documented below.
* `hostname` - Hostname of the server to connect to.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ip_address` - Domain controller IP address.
* `ip6` - Domain controller IPv6 address.
* `ldap_server` - LDAP server name.
* `name` - Domain controller entry name.
* `password` - Password for specified username.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `replication_port` - Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller.
* `source_ip6` - FortiGate IPv6 address to be used for communication with the domain controller.
* `source_port` - Source port to be used for communication with the domain controller.
* `username` - User name to sign in with. Must have proper permissions for service.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `extra_server` block supports:

* `id` - Server ID.
* `ip_address` - Domain controller IP address.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller.
* `source_port` - Source port to be used for communication with the domain controller.


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
