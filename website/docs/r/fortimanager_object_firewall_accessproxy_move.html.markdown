---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_accessproxy_move"
description: |-
  Move Access Proxy.
---

# fortimanager_object_firewall_accessproxy_move
Move Access Proxy.

## Example Usage

```hcl
resource "fortimanager_object_firewall_accessproxy" "trname" {
  name = "accessproxy4"
}

resource "fortimanager_object_firewall_accessproxy" "trname2" {
  name = "accessproxy5"
}

resource "fortimanager_object_firewall_accessproxy_move" "trname2" {
  access_proxy = "accessproxy5"
  target       = "accessproxy4"
  option       = "before"
  depends_on   = [fortimanager_object_firewall_accessproxy.trname, fortimanager_object_firewall_accessproxy.trname]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `access_proxy` - Access Proxy.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the access proxy changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of access proxies.
