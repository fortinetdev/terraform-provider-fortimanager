---
subcategory: "Object ZTNA"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ztna_webportal"
description: |-
  Configure ztna web-portal.
---

# fortimanager_object_ztna_webportal
Configure ztna web-portal.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auth_portal` - Enable/disable authentication portal. Valid values: `disable`, `enable`.

* `auth_rule` - Authentication Rule.
* `auth_virtual_host` - Virtual host for authentication portal.
* `clipboard` - Enable to support RDP/VPC clipboard functionality. Valid values: `disable`, `enable`.

* `cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `customize_forticlient_download_url` - Enable support of customized download URL for FortiClient. Valid values: `disable`, `enable`.

* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `default_window_height` - Screen height (range from 0 - 65535, default = 768).
* `default_window_width` - Screen width (range from 0 - 65535, default = 1024).
* `display_bookmark` - Enable to display the web portal bookmark widget. Valid values: `disable`, `enable`.

* `display_history` - Enable to display the web portal user login history widget. Valid values: `disable`, `enable`.

* `display_status` - Enable to display the web portal status widget. Valid values: `disable`, `enable`.

* `focus_bookmark` - Enable to prioritize the placement of the bookmark section over the quick-connection section in the ztna web-portal. Valid values: `disable`, `enable`.

* `forticlient_download` - Enable/disable download option for FortiClient. Valid values: `disable`, `enable`.

* `forticlient_download_method` - Forticlient-Download-Method. Valid values: `direct`, `ssl-vpn`.

* `heading` - Web portal heading message.
* `host` - Virtual or real host name.
* `log_blocked_traffic` - Enable/disable logging of blocked traffic. Valid values: `disable`, `enable`.

* `macos_forticlient_download_url` - Download URL for Mac FortiClient.
* `name` - ZTNA proxy name.
* `policy_auth_sso` - Enable policy sso authentication. Valid values: `disable`, `enable`.

* `theme` - Web portal color scheme. Valid values: `melongene`, `mariner`, `neutrino`, `jade`, `graphite`, `dark-matter`, `onyx`, `eclipse`, `jet-stream`, `security-fabric`.

* `vip` - Virtual IP name.
* `vip6` - Virtual IPv6 name.
* `windows_forticlient_download_url` - Download URL for Windows FortiClient.
* `bookmarks` - Bookmarks.
* `llm_profile` - Llm-Profile.
* `llm_proxy` - Llm-Proxy. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectZtna WebPortal can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_ztna_webportal.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
