---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_accessproxy6_move"
description: |-
  Move IPv6 access proxy.
---

# fortimanager_object_firewall_accessproxy6_move
Move IPv6 access proxy.

## Example Usage

```hcl
resource "fortimanager_object_firewall_accessproxy6" "trname" {
  name = "terr-accessproxy66"
}

resource "fortimanager_object_firewall_accessproxy6" "trname2" {
  name = "terr-accessproxy67"
}

resource "fortimanager_object_firewall_accessproxy6_move" "trname2" {
  access_proxy6 = "terr-accessproxy66"
  target        = "terr-accessproxy67"
  option        = "before"
  depends_on    = [fortimanager_object_firewall_accessproxy6.trname, fortimanager_object_firewall_accessproxy6.trname2]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `access_proxy6` - Access Proxy6.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the access proxy6 changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of access proxy6s.
