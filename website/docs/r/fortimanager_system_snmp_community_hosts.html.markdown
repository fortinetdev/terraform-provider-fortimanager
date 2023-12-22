---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_snmp_community_hosts"
description: |-
  Allow hosts configuration.
---

# fortimanager_system_snmp_community_hosts
Allow hosts configuration.

~> This resource is a sub resource for variable `hosts` of resource `fortimanager_system_snmp_community`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_system_snmp_community_hosts" "trname" {
  community  = fortimanager_system_snmp_community.trname.fosid
  fosid      = 3
  interface  = "port4"
  ip         = "192.168.0.1 255.255.255.255"
  depends_on = [fortimanager_system_snmp_community.trname]
}

resource "fortimanager_system_snmp_community" "trname" {
  fosid = "1"
  name  = "terraform-tefv-snmp3"
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

System SnmpCommunityHosts can be imported using any of these accepted formats:
```
Set import_options = ["community=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_snmp_community_hosts.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

