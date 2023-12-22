---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_application_list_defaultnetworkservices"
description: |-
  Default network service entries.
---

# fortimanager_object_application_list_defaultnetworkservices
Default network service entries.

~> This resource is a sub resource for variable `default_network_services` of resource `fortimanager_object_application_list`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_application_list_defaultnetworkservices" "trname" {
  services         = ["dns"]
  violation_action = "pass"
  list             = fortimanager_object_application_list.trname.name
  fosid            = 1
  depends_on       = [fortimanager_object_application_list.trname]
}

resource "fortimanager_object_application_list" "trname" {
  name = "terr-application-list"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `list` - List.

* `fosid` - Entry ID.
* `port` - Port number.
* `services` - Network protocols. Valid values: `http`, `ssh`, `telnet`, `ftp`, `dns`, `smtp`, `pop3`, `imap`, `snmp`, `nntp`, `https`.

* `violation_action` - Action for protocols not white listed under selected port. Valid values: `block`, `monitor`, `pass`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectApplication ListDefaultNetworkServices can be imported using any of these accepted formats:
```
Set import_options = ["list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_application_list_defaultnetworkservices.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
