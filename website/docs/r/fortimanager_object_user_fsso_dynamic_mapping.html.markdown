---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_fsso_dynamic_mapping"
description: |-
  Configure Fortinet Single Sign On (FSSO) agents.
---

# fortimanager_object_user_fsso_dynamic_mapping
Configure Fortinet Single Sign On (FSSO) agents.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_user_fsso`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `fsso` - Fsso.

* `_gui_meta` - _Gui_Meta.
* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `group_poll_interval` - Interval in minutes within to fetch groups from FSSO server, or unset to disable.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ldap_poll` - Enable/disable automatic fetching of groups from LDAP server. Valid values: `disable`, `enable`.

* `ldap_poll_filter` - Filter used to fetch groups.
* `ldap_poll_interval` - Interval in minutes within to fetch groups from LDAP server.
* `ldap_server` - LDAP server to get group information.
* `logon_timeout` - Interval in minutes to keep logons after FSSO server down.
* `password` - Password of the first FSSO collector agent.
* `password2` - Password of the second FSSO collector agent.
* `password3` - Password of the third FSSO collector agent.
* `password4` - Password of the fourth FSSO collector agent.
* `password5` - Password of the fifth FSSO collector agent.
* `port` - Port of the first FSSO collector agent.
* `port2` - Port of the second FSSO collector agent.
* `port3` - Port of the third FSSO collector agent.
* `port4` - Port of the fourth FSSO collector agent.
* `port5` - Port of the fifth FSSO collector agent.
* `server` - Domain name or IP address of the first FSSO collector agent.
* `server2` - Domain name or IP address of the second FSSO collector agent.
* `server3` - Domain name or IP address of the third FSSO collector agent.
* `server4` - Domain name or IP address of the fourth FSSO collector agent.
* `server5` - Domain name or IP address of the fifth FSSO collector agent.
* `sni` - Sni.
* `source_ip` - Source IP for communications to FSSO agent.
* `source_ip6` - IPv6 source for communications to FSSO agent.
* `ssl` - Enable/disable use of SSL. Valid values: `disable`, `enable`.

* `ssl_server_host_ip_check` - Enable/disable server host/IP verification. Valid values: `disable`, `enable`.

* `ssl_trusted_cert` - Trusted server certificate or CA certificate.
* `type` - Server type. Valid values: `default`, `fortiems`, `fortinac`, `fortiems-cloud`.

* `user_info_server` - LDAP server to get user information.
* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectUser FssoDynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["fsso=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_fsso_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
