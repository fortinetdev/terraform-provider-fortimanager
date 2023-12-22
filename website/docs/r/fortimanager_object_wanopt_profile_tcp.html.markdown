---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wanopt_profile_tcp"
description: |-
  Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features.
---

# fortimanager_object_wanopt_profile_tcp
Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features.

~> This resource is a sub resource for variable `tcp` of resource `fortimanager_object_wanopt_profile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wanopt_profile_tcp" "trname" {
  byte_caching = "enable"
  port         = 12
  ssl          = "disable"
  depends_on   = [fortimanager_object_wanopt_profile.trname5]
  profile      = fortimanager_object_wanopt_profile.trname5.name
}

resource "fortimanager_object_wanopt_profile" "trname5" {
  name = "terr-wanopt-profile5"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `byte_caching` - Enable/disable byte-caching for HTTP. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `byte_caching_opt` - Select whether TCP byte-caching uses system memory only or both memory and disk space. Valid values: `mem-only`, `mem-disk`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `port` - Single port number or port number range for TCP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `ssl` - Enable/disable SSL/TLS offloading. Valid values: `disable`, `enable`.

* `ssl_port` - Port on which to expect HTTPS traffic for SSL/TLS offloading.
* `status` - Enable/disable HTTP WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWanopt ProfileTcp can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wanopt_profile_tcp.labelname ObjectWanoptProfileTcp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
