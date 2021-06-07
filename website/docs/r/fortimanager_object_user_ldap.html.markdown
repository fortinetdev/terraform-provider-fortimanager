---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_ldap"
description: |-
  Configure LDAP server entries.
---

# fortimanager_object_user_ldap
Configure LDAP server entries.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `account_key_filter` - Account key filter, using the UPN as the search filter.
* `account_key_processing` - Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.

* `ca_cert` - CA certificate name.
* `cnid` - Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
* `dn` - Distinguished name used to look up entries on the LDAP server.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `group_filter` - Filter used for group matching.
* `group_member_check` - Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.

* `group_object_filter` - Filter used for group searching.
* `group_search_base` - Search base used for group searching.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `member_attr` - Name of attribute from which to get group membership.
* `name` - LDAP server entry name.
* `obtain_user_info` - Enable/disable obtaining of user information. Valid values: `disable`, `enable`.

* `password` - Password for initial binding.
* `password_expiry_warning` - Enable/disable password expiry warnings. Valid values: `disable`, `enable`.

* `password_renewal` - Enable/disable online password renewal. Valid values: `disable`, `enable`.

* `port` - Port to be used for communication with the LDAP server (default = 389).
* `search_type` - Search type. Valid values: `nested`, `recursive`.

* `secondary_server` - Secondary LDAP server CN domain name or IP.
* `secure` - Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.

* `server` - LDAP server CN domain name or IP.
* `server_identity_check` - Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `disable`, `enable`.

* `source_ip` - Source IP for communications to LDAP server.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`.

* `tertiary_server` - Tertiary LDAP server CN domain name or IP.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.

* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.

* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.

* `type` - Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.

* `user_info_exchange_server` - MS Exchange server from which to fetch user information.
* `username` - Username (full DN) for initial binding.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `account_key_filter` - Account key filter, using the UPN as the search filter.
* `account_key_name` - Account-Key-Name.
* `account_key_processing` - Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.

* `antiphish` - Antiphish. Valid values: `disable`, `enable`.

* `ca_cert` - CA certificate name.
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
* `source_port` - Source-Port.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`.

* `tertiary_server` - Tertiary LDAP server CN domain name or IP.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.

* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.

* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.

* `type` - Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.

* `user_info_exchange_server` - MS Exchange server from which to fetch user information.
* `username` - Username (full DN) for initial binding.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Ldap can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_ldap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
