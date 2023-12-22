---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_profile_override"
description: |-
  Web Filter override settings.
---

# fortimanager_object_webfilter_profile_override
Web Filter override settings.

~> This resource is a sub resource for variable `override` of resource `fortimanager_object_webfilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `url_profile` - Url Profile.

* `ovrd_cookie` - Allow/deny browser-based (cookie) overrides. Valid values: `deny`, `allow`.

* `ovrd_dur` - Override duration.
* `ovrd_dur_mode` - Override duration mode. Valid values: `constant`, `ask`.

* `ovrd_scope` - Override scope. Valid values: `user`, `user-group`, `ip`, `ask`, `browser`.

* `ovrd_user_group` - User groups with permission to use the override.
* `profile` - Web filter profile with permission to create overrides.
* `profile_attribute` - Profile attribute to retrieve from the RADIUS server. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Vendor-Specific`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `profile_type` - Override profile type. Valid values: `list`, `radius`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWebfilter ProfileOverride can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_profile_override.labelname ObjectWebfilterProfileOverride
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
