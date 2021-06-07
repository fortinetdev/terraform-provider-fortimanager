---
subcategory: "Object Switch-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_securitypolicy_8021x"
description: |-
  Configure 802.1x MAC Authentication Bypass (MAB) policies.
---

# fortimanager_object_switchcontroller_securitypolicy_8021x
Configure 802.1x MAC Authentication Bypass (MAB) policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auth_fail_vlan` - Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.

* `auth_fail_vlan_id` - VLAN ID on which authentication failed.
* `auth_fail_vlanid` - VLAN ID on which authentication failed.
* `authserver_timeout_period` - Authentication server timeout period (3 - 15 sec, default = 3).
* `authserver_timeout_vlan` - Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable. Valid values: `disable`, `enable`.

* `authserver_timeout_vlanid` - Authentication server timeout VLAN name.
* `eap_auto_untagged_vlans` - Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.

* `eap_passthru` - Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.

* `framevid_apply` - Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.

* `guest_auth_delay` - Guest authentication delay (1 - 900  sec, default = 30).
* `guest_vlan` - Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.

* `guest_vlan_id` - Guest VLAN name.
* `guest_vlanid` - Guest VLAN ID.
* `mac_auth_bypass` - Enable/disable MAB for this policy. Valid values: `disable`, `enable`.

* `name` - Policy name.
* `open_auth` - Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.

* `policy_type` - Policy type. Valid values: `802.1X`.

* `radius_timeout_overwrite` - Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.

* `security_mode` - Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.

* `user_group` - Name of user-group to assign to this MAC Authentication Bypass (MAB) policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController SecurityPolicy8021X can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_securitypolicy_8021x.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
