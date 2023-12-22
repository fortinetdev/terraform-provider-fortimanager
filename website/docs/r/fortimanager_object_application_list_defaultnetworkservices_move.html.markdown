---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_application_list_defaultnetworkservices_move"
description: |-
  Default network service entries.
---

# fortimanager_object_application_list_defaultnetworkservices_move
Default network service entries.

## Example Usage

```hcl
resource "fortimanager_object_application_list_defaultnetworkservices_move" "trname" {
  list                     = fortimanager_object_application_list.trname.name
  default_network_services = 2
  target                   = 1
  option                   = "before"
  depends_on               = [fortimanager_object_application_list_defaultnetworkservices.trname, fortimanager_object_application_list_defaultnetworkservices.trname2]
}

resource "fortimanager_object_application_list_defaultnetworkservices" "trname2" {
  list       = fortimanager_object_application_list.trname.name
  fosid      = 2
  depends_on = [fortimanager_object_application_list.trname]
}

resource "fortimanager_object_application_list_defaultnetworkservices" "trname" {
  list       = fortimanager_object_application_list.trname.name
  fosid      = 1
  depends_on = [fortimanager_object_application_list.trname]
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
* `default_network_services` - Default Network Services.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the default network services changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of default network servicess.
