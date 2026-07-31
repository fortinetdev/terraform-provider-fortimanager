---
subcategory: "Object ZTNA"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ztna_webportalbookmark_bookmarks"
description: |-
  Bookmark table.
---

# fortimanager_object_ztna_webportalbookmark_bookmarks
Bookmark table.

~> This resource is a sub resource for variable `bookmarks` of resource `fortimanager_object_ztna_webportalbookmark`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `web_portal_bookmark` - Web Portal Bookmark.

* `apptype` - Application type. Valid values: `web`, `telnet`, `ssh`, `ftp`, `smb`, `vnc`, `rdp`, `sftp`.

* `color_depth` - Color depth per pixel. Valid values: `8`, `16`, `32`.

* `description` - Description.
* `domain` - Login domain.
* `folder` - Network shared file folder parameter.
* `height` - Screen height (range from 0 - 65535, default = 0).
* `host` - Host name/IP parameter.
* `keyboard_layout` - Keyboard layout. Valid values: `da`, `de`, `de-ch`, `en-uk`, `en-us`, `es`, `fi`, `fr`, `fr-be`, `fr-ca`, `fr-ch`, `hr`, `hu`, `it`, `ja`, `lt`, `mk`, `no`, `pt`, `pt-br`, `ru`, `sl`, `sv`, `ar-101`, `ar-102`, `ar-102-azerty`, `can-mul`, `cz`, `cz-qwerty`, `cz-pr`, `nl`, `de-ibm`, `en-uk-ext`, `en-us-dvorak`, `es-var`, `fi-sami`, `hu-101`, `it-142`, `ko`, `lt-ibm`, `lt-std`, `lav-std`, `lav-leg`, `mk-std`, `no-sami`, `pol-214`, `pol-pr`, `pt-br-abnt2`, `ru-mne`, `ru-t`, `sv-sami`, `tuk`, `tur-f`, `tur-q`, `zh-sym-sg-us`, `zh-sym-us`, `zh-tr-hk`, `zh-tr-mo`, `zh-tr-us`, `fr-apple`, `la-am`, `ja-106`.

* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `logon_password` - Logon password.
* `logon_user` - Logon user.
* `name` - Bookmark name.
* `port` - Remote port.
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `preconnection_id` - The numeric ID of the RDP source (0-4294967295).
* `restricted_admin` - Enable/disable restricted admin mode for RDP. Valid values: `disable`, `enable`.

* `security` - Security mode for RDP connection (default = any). Valid values: `rdp`, `nla`, `tls`, `any`.

* `send_preconnection_id` - Enable/disable sending of preconnection ID. Valid values: `disable`, `enable`.

* `sso` - Single sign-on. Valid values: `disable`, `enable`.

* `url` - URL parameter.
* `verify_cert` - Enable/disable certificate verification of the real server. Valid values: `disable`, `enable`.

* `vnc_keyboard_layout` - Keyboard layout. Valid values: `da`, `de`, `de-ch`, `en-uk`, `es`, `fi`, `fr`, `fr-be`, `it`, `no`, `pt`, `sv`, `nl`, `en-uk-ext`, `it-142`, `pt-br-abnt2`, `default`, `fr-ca-mul`, `gd`, `us-intl`.

* `width` - Screen width (range from 0 - 65535, default = 0).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectZtna WebPortalBookmarkBookmarks can be imported using any of these accepted formats:
```
Set import_options = ["web_portal_bookmark=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_ztna_webportalbookmark_bookmarks.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
