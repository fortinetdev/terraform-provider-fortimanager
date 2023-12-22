---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_securityexemptlist_rule"
description: |-
  Configure rules for exempting users from captive portal authentication.
---

# fortimanager_object_user_securityexemptlist_rule
Configure rules for exempting users from captive portal authentication.

~> This resource is a sub resource for variable `rule` of resource `fortimanager_object_user_securityexemptlist`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_user_securityexemptlist_rule" "trname" {
  security_exempt_list = fortimanager_object_user_securityexemptlist.trname.name
  fosid                = 1
  depends_on           = [fortimanager_object_user_securityexemptlist.trname]
}

resource "fortimanager_object_user_securityexemptlist" "trname" {
  name = "terr-securityexemptlist"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `security_exempt_list` - Security Exempt List.

* `devices` - Devices or device groups.
* `dstaddr` - Destination addresses or address groups.
* `fosid` - ID.
* `service` - Destination services.
* `srcaddr` - Source addresses or address groups.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectUser SecurityExemptListRule can be imported using any of these accepted formats:
```
Set import_options = ["security_exempt_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_securityexemptlist_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
