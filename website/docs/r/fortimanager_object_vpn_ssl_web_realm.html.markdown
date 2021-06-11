---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ssl_web_realm"
description: |-
  Realm.
---

# fortimanager_object_vpn_ssl_web_realm
Realm.

## Example Usage

```hcl
resource "fortimanager_object_vpn_ssl_web_realm" "trname" {
  max_concurrent_user = 5
  nas_ip              = "192.168.1.1"
  url_path            = "terr-vpn-ssl-web-realm"
  virtual_host_only   = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `login_page` - Replacement HTML for SSL-VPN login page.
* `max_concurrent_user` - Maximum concurrent users (0 - 65535, 0 means unlimited).
* `nas_ip` - IP address used as a NAS-IP to communicate with the RADIUS server.
* `radius_port` - RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
* `radius_server` - RADIUS server associated with realm.
* `url_path` - URL path to access SSL-VPN login page.
* `virtual_host` - Virtual host name for realm.
* `virtual_host_only` - Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{url_path}}.

## Import

ObjectVpn SslWebRealm can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ssl_web_realm.labelname {{url_path}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
