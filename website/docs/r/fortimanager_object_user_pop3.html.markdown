---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_pop3"
description: |-
  POP3 server entry configuration.
---

# fortimanager_object_user_pop3
POP3 server entry configuration.

## Example Usage

```hcl
resource "fortimanager_object_user_pop3" "trname" {
  name                  = "terraform-tefv-pop"
  port                  = 8000
  secure                = "none"
  server                = "terraform"
  ssl_min_proto_version = "default"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `name` - POP3 server entry name.
* `port` - POP3 service port number.
* `secure` - SSL connection. Valid values: `none`, `starttls`, `pop3s`.

* `server` - {&lt;name_str|ip_str&gt;} server domain name or IP.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Pop3 can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_pop3.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
