---
subcategory: "Object Extension Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extensioncontroller_extendervap"
description: |-
  FortiExtender wifi vap configuration.
---

# fortimanager_object_extensioncontroller_extendervap
FortiExtender wifi vap configuration.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `http`, `ssh`, `telnet`, `snmp`, `https`, `ping`.

* `auth_server_address` - Wi-Fi Authentication Server Address (IPv4 format).
* `auth_server_port` - Wi-Fi Authentication Server Port.
* `auth_server_secret` - Wi-Fi Authentication Server Secret.
* `broadcast_ssid` - Wi-Fi broadcast SSID enable / disable. Valid values: `disable`, `enable`.

* `bss_color_partial` - Wi-Fi 802.11AX bss color partial enable / disable, default = enable. Valid values: `disable`, `enable`.

* `dtim` - Wi-Fi DTIM (1 - 255) default = 1.
* `end_ip` - End ip address.
* `ip_address` - Extender ip address.
* `max_clients` - Wi-Fi max clients (0 - 512), default = 0 (no limit)
* `mu_mimo` - Wi-Fi multi-user MIMO enable / disable, default = enable. Valid values: `disable`, `enable`.

* `name` - Wi-Fi VAP name.
* `passphrase` - Wi-Fi passphrase.
* `pmf` - Wi-Fi pmf enable/disable, default = disable. Valid values: `disabled`, `optional`, `required`.

* `rts_threshold` - Wi-Fi RTS Threshold (256 - 2347), default = 2347 (RTS/CTS disabled).
* `sae_password` - Wi-Fi SAE Password.
* `security` - Wi-Fi security. Valid values: `OPEN`, `WPA2-Personal`, `WPA-WPA2-Personal`, `WPA3-SAE`, `WPA3-SAE-Transition`, `WPA2-Enterprise`, `WPA3-Enterprise-only`, `WPA3-Enterprise-transition`, `WPA3-Enterprise-192-bit`.

* `security_exempt_list` - Name of security exempt list.
* `security_external_web` - URL of external authentication web server.
* `security_groups` - User groups that can authenticate with the captive portal.
* `security_mode` - Turn on captive portal authentication for this Wi-Fi interface. Valid values: `none`, `captive-portal`.

* `security_redirect_url` - Optional URL for redirecting users after they pass captive portal authentication.
* `ssid` - Wi-Fi SSID.
* `start_ip` - Start ip address.
* `target_wake_time` - Wi-Fi 802.11AX target wake time enable / disable, default = enable. Valid values: `disable`, `enable`.

* `type` - Wi-Fi VAP type local-vap / lan-extension-vap. Valid values: `local-vap`, `lan-ext-vap`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectExtensionController ExtenderVap can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extensioncontroller_extendervap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
