---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_ldap_dynamic_mapping"
description: |-
  Configure LDAP server entries.
---

# fortimanager_object_user_ldap_dynamic_mapping
Configure LDAP server entries.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_user_ldap`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `ldap` - Ldap.

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `account_key_cert_field` - Define subject identity field in certificate for user access right checking. Valid values: `othername`, `rfc822name`, `dnsname`.

* `account_key_filter` - Account key filter, using the UPN as the search filter.
* `account_key_name` - Account-Key-Name.
* `account_key_processing` - Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.

* `account_key_upn_san` - Define SAN in certificate for user principle name matching. Valid values: `othername`, `rfc822name`, `dnsname`.

* `antiphish` - Antiphish. Valid values: `disable`, `enable`.

* `ca_cert` - CA certificate name.
* `client_cert` - Client-Cert.
* `client_cert_auth` - Client-Cert-Auth. Valid values: `disable`, `enable`.

* `cnid` - Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
* `dn` - Distinguished name used to look up entries on the LDAP server.
* `filter` - Filter.
* `group` - Group.
* `group_filter` - Filter used for group matching.
* `group_member_check` - Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.

* `group_object_filter` - Filter used for group searching.
* `group_object_search_base` - Group-Object-Search-Base.
* `group_search_base` - Search base used for group searching.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `max_connections` - Max-Connections.
* `member_attr` - Name of attribute from which to get group membership.
* `obtain_user_info` - Enable/disable obtaining of user information. Valid values: `disable`, `enable`.

* `password` - Password for initial binding.
* `password_attr` - Password-Attr.
* `password_expiry_warning` - Enable/disable password expiry warnings. Valid values: `disable`, `enable`.

* `password_renewal` - Enable/disable online password renewal. Valid values: `disable`, `enable`.

* `port` - Port to be used for communication with the LDAP server (default = 389).
* `retrieve_protection_profile` - Retrieve-Protection-Profile.
* `search_type` - Search type. Valid values: `nested`, `recursive`.

* `secondary_server` - Secondary LDAP server CN domain name or IP.
* `secure` - Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.

* `server` - LDAP server CN domain name or IP.
* `server_identity_check` - Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `disable`, `enable`.

* `source_ip` - Source IP for communications to LDAP server.
* `source_ip_interface` - Source interface for communication with the LDAP server.
* `source_port` - Source-Port.
* `ssl_max_proto_version` - Ssl-Max-Proto-Version. Valid values: `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1`, `TLSv1-3`.

* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`.

* `status_ttl` - Time for which server reachability is cached so that when a server is unreachable, it will not be retried for at least this period of time (0 = cache disabled, default = 300).
* `tertiary_server` - Tertiary LDAP server CN domain name or IP.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.

* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.

* `two_factor_filter` - Filter used to synchronize users to FortiToken Cloud.
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.

* `type` - Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.

* `user_info_exchange_server` - MS Exchange server from which to fetch user information.
* `username` - Username (full DN) for initial binding.
* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectUser LdapDynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["ldap=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_ldap_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
