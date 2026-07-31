---
subcategory: "Object Switch Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_securitypolicy_admin"
description: |-
  Configure fortiswitch's admin security-policy.
---

# fortimanager_object_switchcontroller_securitypolicy_admin
Configure fortiswitch's admin security-policy.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auto` - Automatically set based on the host ip connected via the Fortilink interface. Valid values: `disable`, `enable`.

* `ip6_trusthost1` - Trusted IPv6 host.
* `ip6_trusthost10` - Trusted IPv6 host.
* `ip6_trusthost2` - Trusted IPv6 host.
* `ip6_trusthost3` - Trusted IPv6 host.
* `ip6_trusthost4` - Trusted IPv6 host.
* `ip6_trusthost5` - Trusted IPv6 host.
* `ip6_trusthost6` - Trusted IPv6 host.
* `ip6_trusthost7` - Trusted IPv6 host.
* `ip6_trusthost8` - Trusted IPv6 host.
* `ip6_trusthost9` - Trusted IPv6 host.
* `name` - Policy name.
* `trusthost1` - Trusted IPv4 host.
* `trusthost10` - Trusted IPv4 host.
* `trusthost2` - Trusted IPv4 host.
* `trusthost3` - Trusted IPv4 host.
* `trusthost4` - Trusted IPv4 host.
* `trusthost5` - Trusted IPv4 host.
* `trusthost6` - Trusted IPv4 host.
* `trusthost7` - Trusted IPv4 host.
* `trusthost8` - Trusted IPv4 host.
* `trusthost9` - Trusted IPv4 host.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController SecurityPolicyAdmin can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_securitypolicy_admin.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
