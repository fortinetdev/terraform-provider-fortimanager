---
subcategory: "Object Switch Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_securitypolicy_localaccess"
description: |-
  Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch units.
---

# fortimanager_object_switchcontroller_securitypolicy_localaccess
Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch units.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `internal_allowaccess` - Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.

* `mgmt_allowaccess` - Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.

* `name` - Policy name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController SecurityPolicyLocalAccess can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_securitypolicy_localaccess.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
