---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_snmp_community_hosts6"
description: |-
  Allow hosts configuration for IPv6.
---

# fortimanager_system_snmp_community_hosts6
Allow hosts configuration for IPv6.

~> This resource is a sub resource for variable `hosts6` of resource `fortimanager_system_snmp_community`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_system_snmp_community_hosts6" "trname" {
  community  = fortimanager_system_snmp_community.trname4.fosid
  fosid      = 2
  interface  = "port2"
  ip         = "2001:db8:85a3::8a2e:370:7334/128"
  depends_on = [fortimanager_system_snmp_community.trname4]
}

resource "fortimanager_system_snmp_community" "trname4" {
  fosid = "2"
  name  = "terraform-tefv-snmp4"
}
```

## Argument Reference


The following arguments are supported:

* `community` - Community.

* `fosid` - Host entry ID.
* `interface` - Allow interface name.
* `ip` - Allow host IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SnmpCommunityHosts6 can be imported using any of these accepted formats:
```
Set import_options = ["community=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_snmp_community_hosts6.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

