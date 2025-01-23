---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_group_dynamic_mapping"
description: |-
  Configure user groups.
---

# fortimanager_object_user_group_dynamic_mapping
Configure user groups.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_user_group`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `guest`: `fortimanager_object_user_group_dynamic_mapping_guest`
>- `match`: `fortimanager_object_user_group_dynamic_mapping_match`
>- `sslvpn_os_check_list`: `fortimanager_object_user_group_dynamic_mapping_sslvpnoschecklist`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `group` - Group.

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `auth_concurrent_override` - Enable/disable overriding the global number of concurrent authentication sessions for this user group. Valid values: `disable`, `enable`.

* `auth_concurrent_value` - Maximum number of concurrent authenticated connections per user (0 - 100).
* `authtimeout` - Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.
* `company` - Set the action for the company guest user field. Valid values: `optional`, `mandatory`, `disabled`.

* `email` - Enable/disable the guest user email address field. Valid values: `disable`, `enable`.

* `expire` - Time in seconds before guest user accounts expire. (1 - 31536000 sec)
* `expire_type` - Determine when the expiration countdown begins. Valid values: `immediately`, `first-successful-login`.

* `group_type` - Set the group to be for firewall authentication, FSSO, RSSO, or guest users. Valid values: `firewall`, `directory-service`, `fsso-service`, `guest`, `rsso`.

* `guest` - Guest. The structure of `guest` block is documented below.
* `http_digest_realm` - Realm attribute for MD5-digest authentication.
* `fosid` - Id.
* `ldap_memberof` - Ldap-Memberof.
* `logic_type` - Logic-Type. Valid values: `or`, `and`.

* `match` - Match. The structure of `match` block is documented below.
* `max_accounts` - Maximum number of guest accounts that can be created for this group (0 means unlimited).
* `member` - Names of users, peers, LDAP severs, or RADIUS servers to add to the user group.
* `mobile_phone` - Enable/disable the guest user mobile phone number field. Valid values: `disable`, `enable`.

* `multiple_guest_add` - Enable/disable addition of multiple guests. Valid values: `disable`, `enable`.

* `password` - Guest user password type. Valid values: `auto-generate`, `specify`, `disable`.

* `redir_url` - Redir-Url.
* `sms_custom_server` - SMS server.
* `sms_server` - Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.

* `sponsor` - Set the action for the sponsor guest user field. Valid values: `optional`, `mandatory`, `disabled`.

* `sslvpn_bookmarks_group` - Sslvpn-Bookmarks-Group.
* `sslvpn_cache_cleaner` - Sslvpn-Cache-Cleaner. Valid values: `disable`, `enable`.

* `sslvpn_client_check` - Sslvpn-Client-Check. Valid values: `forticlient`, `forticlient-av`, `forticlient-fw`, `3rdAV`, `3rdFW`.

* `sslvpn_ftp` - Sslvpn-Ftp. Valid values: `disable`, `enable`.

* `sslvpn_http` - Sslvpn-Http. Valid values: `disable`, `enable`.

* `sslvpn_os_check` - Sslvpn-Os-Check. Valid values: `disable`, `enable`.

* `sslvpn_os_check_list` - Sslvpn-Os-Check-List. The structure of `sslvpn_os_check_list` block is documented below.
* `sslvpn_portal` - Sslvpn-Portal.
* `sslvpn_portal_heading` - Sslvpn-Portal-Heading.
* `sslvpn_rdp` - Sslvpn-Rdp. Valid values: `disable`, `enable`.

* `sslvpn_samba` - Sslvpn-Samba. Valid values: `disable`, `enable`.

* `sslvpn_split_tunneling` - Sslvpn-Split-Tunneling. Valid values: `disable`, `enable`.

* `sslvpn_ssh` - Sslvpn-Ssh. Valid values: `disable`, `enable`.

* `sslvpn_telnet` - Sslvpn-Telnet. Valid values: `disable`, `enable`.

* `sslvpn_tunnel` - Sslvpn-Tunnel. Valid values: `disable`, `enable`.

* `sslvpn_tunnel_endip` - Sslvpn-Tunnel-Endip.
* `sslvpn_tunnel_ip_mode` - Sslvpn-Tunnel-Ip-Mode. Valid values: `range`, `usrgrp`.

* `sslvpn_tunnel_startip` - Sslvpn-Tunnel-Startip.
* `sslvpn_virtual_desktop` - Sslvpn-Virtual-Desktop. Valid values: `disable`, `enable`.

* `sslvpn_vnc` - Sslvpn-Vnc. Valid values: `disable`, `enable`.

* `sslvpn_webapp` - Sslvpn-Webapp. Valid values: `disable`, `enable`.

* `sso_attribute_value` - Name of the RADIUS user group that this local user group represents.
* `user_id` - Guest user ID type. Valid values: `email`, `auto-generate`, `specify`.

* `user_name` - Enable/disable the guest user name entry. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `guest` block supports:

* `comment` - Comment.
* `company` - Set the action for the company guest user field.
* `email` - Email.
* `expiration` - Expire time.
* `group` - Group.
* `id` - Guest ID.
* `mobile_phone` - Mobile phone.
* `name` - Guest name.
* `password` - Guest password.
* `sponsor` - Set the action for the sponsor guest user field.
* `user_id` - Guest ID.

The `match` block supports:

* `_gui_meta` - _Gui_Meta.
* `group_name` - Name of matching user or group on remote authentication server.
* `id` - ID.
* `server_name` - Name of remote auth server.

The `sslvpn_os_check_list` block supports:

* `action` - Action. Valid values: `allow`, `check-up-to-date`, `deny`.

* `latest_patch_level` - Latest-Patch-Level.
* `name` - Name.
* `tolerance` - Tolerance.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectUser GroupDynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_group_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
