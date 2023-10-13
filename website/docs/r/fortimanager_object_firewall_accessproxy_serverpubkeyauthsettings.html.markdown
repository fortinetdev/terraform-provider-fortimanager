---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_accessproxy_serverpubkeyauthsettings"
description: |-
  Server SSH public key authentication settings.
---

# fortimanager_object_firewall_accessproxy_serverpubkeyauthsettings
Server SSH public key authentication settings.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `access_proxy` - Access Proxy.

* `auth_ca` - Name of the SSH server public key authentication CA.
* `cert_extension` - Cert-Extension. The structure of `cert_extension` block is documented below.
* `permit_agent_forwarding` - Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `disable`, `enable`.

* `permit_port_forwarding` - Enable/disable appending permit-port-forwarding certificate extension. Valid values: `disable`, `enable`.

* `permit_pty` - Enable/disable appending permit-pty certificate extension. Valid values: `disable`, `enable`.

* `permit_user_rc` - Enable/disable appending permit-user-rc certificate extension. Valid values: `disable`, `enable`.

* `permit_x11_forwarding` - Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `disable`, `enable`.

* `source_address` - Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `cert_extension` block supports:

* `critical` - Critical option. Valid values: `no`, `yes`.

* `data` - Name of certificate extension.
* `name` - Name of certificate extension.
* `type` - Type of certificate extension. Valid values: `fixed`, `user`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall AccessProxyServerPubkeyAuthSettings can be imported using any of these accepted formats:
```
Set import_options = ["access_proxy=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_accessproxy_serverpubkeyauthsettings.labelname ObjectFirewallAccessProxyServerPubkeyAuthSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
