---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv4rules"
description: |-
  AP ACL layer3 ipv4 rule list.
---

# fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv4rules
AP ACL layer3 ipv4 rule list.

~> This resource is a sub resource for variable `layer3_ipv4_rules` of resource `fortimanager_object_wirelesscontroller_accesscontrollist`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv4rules" "trname" {
  access_control_list = fortimanager_object_wirelesscontroller_accesscontrollist.trname.name
  protocol            = 24
  rule_id             = 12
  srcport             = 23
  depends_on          = [fortimanager_object_wirelesscontroller_accesscontrollist.trname]
}

resource "fortimanager_object_wirelesscontroller_accesscontrollist" "trname" {
  name = "terr-layer3ipv4rules"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `access_control_list` - Access Control List.

* `action` - Policy action (allow | deny). Valid values: `allow`, `deny`.

* `comment` - Description.
* `dstaddr` - Destination IP address (any | local-LAN | IPv4 address[/&lt;network mask | mask length&gt;], default = any).
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `rule_id` - Rule ID (1 - 65535).
* `srcaddr` - Source IP address (any | local-LAN | IPv4 address[/&lt;network mask | mask length&gt;], default = any).
* `srcport` - Source port (0 - 65535, default = 0, meaning any).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{rule_id}}.

## Import

ObjectWirelessController AccessControlListLayer3Ipv4Rules can be imported using any of these accepted formats:
```
Set import_options = ["access_control_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv4rules.labelname {{rule_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
