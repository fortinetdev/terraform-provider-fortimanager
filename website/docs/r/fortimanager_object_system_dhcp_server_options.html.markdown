---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_dhcp_server_options"
description: |-
  DHCP options.
---

# fortimanager_object_system_dhcp_server_options
DHCP options.

~> This resource is a sub resource for variable `options` of resource `fortimanager_object_system_dhcp_server`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_dhcp_server_options" "trname" {
  server     = fortimanager_object_system_dhcp_server.trname3.fosid
  fosid      = 34
  ip         = ["10.160.88.32"]
  type       = "ip"
  depends_on = [fortimanager_object_system_dhcp_server.trname3]
}

resource "fortimanager_object_system_dhcp_server" "trname3" {
  fosid = 4
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `server` - Server.

* `code` - DHCP option code.
* `fosid` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.

* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this option. Valid values: `disable`, `enable`.

* `uci_string` - One or more UCI strings in quotes separated by spaces.
* `value` - DHCP option value.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this option. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectSystem DhcpServerOptions can be imported using any of these accepted formats:
```
Set import_options = ["server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_dhcp_server_options.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
