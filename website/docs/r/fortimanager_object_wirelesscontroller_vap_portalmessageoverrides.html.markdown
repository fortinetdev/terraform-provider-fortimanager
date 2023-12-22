---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_vap_portalmessageoverrides"
description: |-
  Individual message overrides.
---

# fortimanager_object_wirelesscontroller_vap_portalmessageoverrides
Individual message overrides.

~> This resource is a sub resource for variable `portal_message_overrides` of resource `fortimanager_object_wirelesscontroller_vap`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_vap_portalmessageoverrides" "trname" {
  auth_disclaimer_page   = 1
  auth_login_failed_page = 2
  vap                    = fortimanager_object_wirelesscontroller_vap.trname.name
  depends_on             = [fortimanager_object_wirelesscontroller_vap.trname]
}

resource "fortimanager_object_wirelesscontroller_vap" "trname" {
  name = "terr-wictl-vap6"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vap` - Vap.

* `auth_disclaimer_page` - Override auth-disclaimer-page message with message from portal-message-overrides group.
* `auth_login_failed_page` - Override auth-login-failed-page message with message from portal-message-overrides group.
* `auth_login_page` - Override auth-login-page message with message from portal-message-overrides group.
* `auth_reject_page` - Override auth-reject-page message with message from portal-message-overrides group.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWirelessController VapPortalMessageOverrides can be imported using any of these accepted formats:
```
Set import_options = ["vap=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_vap_portalmessageoverrides.labelname ObjectWirelessControllerVapPortalMessageOverrides
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
