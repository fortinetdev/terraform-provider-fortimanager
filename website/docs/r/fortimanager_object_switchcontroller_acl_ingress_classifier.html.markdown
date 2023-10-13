---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_acl_ingress_classifier"
description: |-
  ACL classifiers.
---

# fortimanager_object_switchcontroller_acl_ingress_classifier
ACL classifiers.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `ingress` - Ingress.

* `dst_ip_prefix` - Destination IP address to be matched.
* `dst_mac` - Destination MAC address to be matched.
* `src_ip_prefix` - Source IP address to be matched.
* `src_mac` - Source MAC address to be matched.
* `vlan` - VLAN ID to be matched.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSwitchController AclIngressClassifier can be imported using any of these accepted formats:
```
Set import_options = ["ingress=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_acl_ingress_classifier.labelname ObjectSwitchControllerAclIngressClassifier
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
