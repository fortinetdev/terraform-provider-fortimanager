---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_cifs_domaincontroller"
description: |-
  Define known domain controller servers.
---

# fortimanager_object_cifs_domaincontroller
Define known domain controller servers.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `domain_name` - Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
* `ip` - IPv4 server address.
* `ip6` - IPv6 server address.
* `password` - Password for specified username.
* `port` - Port number of service. Port number 0 indicates automatic discovery.
* `server_name` - Name of the server to connect to.
* `username` - User name to sign in with. Must have proper permissions for service.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{server_name}}.

## Import

ObjectCifs DomainController can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_cifs_domaincontroller.labelname {{server_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
