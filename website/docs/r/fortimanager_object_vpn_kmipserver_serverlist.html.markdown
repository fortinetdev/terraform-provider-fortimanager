---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_kmipserver_serverlist"
description: |-
  KMIP server list.
---

# fortimanager_object_vpn_kmipserver_serverlist
KMIP server list.

~> This resource is a sub resource for variable `server_list` of resource `fortimanager_object_vpn_kmipserver`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `kmip_server` - Kmip Server.

* `cert` - Client certificate to use for connectivity to the KMIP server.
* `fosid` - ID
* `port` - KMIP server port.
* `server` - KMIP server FQDN or IP address.
* `status` - Enable/disable KMIP server. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectVpn KmipServerServerList can be imported using any of these accepted formats:
```
Set import_options = ["kmip_server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_kmipserver_serverlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
