---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_acl_ingress"
description: |-
  Configure ingress ACL policies to be applied on managed FortiSwitch ports.
---

# fortimanager_object_switchcontroller_acl_ingress
Configure ingress ACL policies to be applied on managed FortiSwitch ports.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`action`: `fortimanager_object_switchcontroller_acl_ingress_action`
`classifier`: `fortimanager_object_switchcontroller_acl_ingress_classifier`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `action` - Action. The structure of `action` block is documented below.
* `classifier` - Classifier. The structure of `classifier` block is documented below.
* `description` - Description for the ACL policy.
* `fosid` - ACL ID.

The `action` block supports:

* `count` - Enable/disable count. Valid values: `disable`, `enable`.

* `drop` - Enable/disable drop. Valid values: `disable`, `enable`.


The `classifier` block supports:

* `dst_ip_prefix` - Destination IP address to be matched.
* `dst_mac` - Destination MAC address to be matched.
* `src_ip_prefix` - Source IP address to be matched.
* `src_mac` - Source MAC address to be matched.
* `vlan` - VLAN ID to be matched.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectSwitchController AclIngress can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_acl_ingress.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
