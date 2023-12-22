---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_system_centralmanagement_serverlist"
description: |-
  Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers.
---

# fortimanager_systemp_system_centralmanagement_serverlist
Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers.

~> This resource is a sub resource for variable `server_list` of resource `fortimanager_systemp_system_centralmanagement`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_systemp_system_centralmanagement_serverlist" "trname" {
  devprof        = "default"
  fosid          = 5
  server_address = "10.160.2.8"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `addr_type` - Indicate whether the FortiGate communicates with the override server using an IPv4 address, an IPv6 address or a FQDN. Valid values: `fqdn`, `ipv4`, `ipv6`.

* `fqdn` - FQDN address of override server.
* `fosid` - ID.
* `server_address` - IPv4 address of override server.
* `server_address6` - IPv6 address of override server.
* `server_type` - FortiGuard service type. Valid values: `update`, `rating`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Systemp SystemCentralManagementServerList can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_system_centralmanagement_serverlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
