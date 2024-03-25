---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_accessproxysshclientcert_certextension"
description: |-
  Configure certificate extension for user certificate.
---

# fortimanager_object_firewall_accessproxysshclientcert_certextension
Configure certificate extension for user certificate.

~> This resource is a sub resource for variable `cert_extension` of resource `fortimanager_object_firewall_accessproxysshclientcert`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `access_proxy_ssh_client_cert` - Access Proxy Ssh Client Cert.

* `critical` - Critical option. Valid values: `no`, `yes`.

* `data` - Data of certificate extension.
* `name` - Name of certificate extension.
* `type` - Type of certificate extension. Valid values: `fixed`, `user`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall AccessProxySshClientCertCertExtension can be imported using any of these accepted formats:
```
Set import_options = ["access_proxy_ssh_client_cert=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_accessproxysshclientcert_certextension.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
