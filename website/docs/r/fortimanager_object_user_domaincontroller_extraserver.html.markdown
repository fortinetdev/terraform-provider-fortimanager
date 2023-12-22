---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_domaincontroller_extraserver"
description: |-
  extra servers.
---

# fortimanager_object_user_domaincontroller_extraserver
extra servers.

~> This resource is a sub resource for variable `extra_server` of resource `fortimanager_object_user_domaincontroller`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `domain_controller` - Domain Controller.

* `fosid` - Server ID.
* `ip_address` - Domain controller IP address.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller.
* `source_port` - Source port to be used for communication with the domain controller.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectUser DomainControllerExtraServer can be imported using any of these accepted formats:
```
Set import_options = ["domain_controller=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_domaincontroller_extraserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
