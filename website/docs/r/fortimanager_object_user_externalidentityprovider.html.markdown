---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_externalidentityprovider"
description: |-
  Configure external identity provider.
---

# fortimanager_object_user_externalidentityprovider
Configure external identity provider.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `group_attr_name` - Group attribute name in authentication query.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `name` - External identity provider name.
* `port` - External identity provider service port number (0 to use default).
* `server_identity_check` - Enable/disable server's identity check against its certificate and subject alternative name(s). Valid values: `disable`, `enable`.

* `source_ip` - Use this IPv4/v6 address to connect to the external identity provider.
* `timeout` - Connection timeout value in seconds (default=5).
* `type` - External identity provider type. Valid values: `ms-graph`.

* `url` - Url.
* `user_attr_name` - User attribute name in authentication query.
* `version` - External identity API version. Valid values: `beta`, `v1.0`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser ExternalIdentityProvider can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_externalidentityprovider.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
