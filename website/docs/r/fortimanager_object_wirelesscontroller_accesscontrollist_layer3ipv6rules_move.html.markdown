---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules_move"
description: |-
  AP ACL layer3 ipv6 rule list.
---

# fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules_move
AP ACL layer3 ipv6 rule list.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules_move" "trname" {
  access_control_list = fortimanager_object_wirelesscontroller_accesscontrollist.trname2.name
  layer3_ipv6_rules   = fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules.trname2.rule_id
  target              = fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules.trname.rule_id
  option              = "after"
  depends_on          = [fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules.trname2, fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules.trname]
}

resource "fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules" "trname2" {
  rule_id             = 13
  srcport             = 35
  access_control_list = fortimanager_object_wirelesscontroller_accesscontrollist.trname2.name
  depends_on          = [fortimanager_object_wirelesscontroller_accesscontrollist.trname2]
}

resource "fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules" "trname" {
  rule_id             = 12
  srcport             = 34
  access_control_list = fortimanager_object_wirelesscontroller_accesscontrollist.trname2.name
  depends_on          = [fortimanager_object_wirelesscontroller_accesscontrollist.trname2]
}

resource "fortimanager_object_wirelesscontroller_accesscontrollist" "trname2" {
  name = "terr-accesscontrollist"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `access_control_list` - Access Control List.
* `layer3_ipv6_rules` - Layer3 Ipv6 Rules.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{rule_id}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the layer3 ipv6 rules changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of layer3 ipv6 ruless.
