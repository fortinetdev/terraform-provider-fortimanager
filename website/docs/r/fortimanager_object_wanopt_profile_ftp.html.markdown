---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wanopt_profile_ftp"
description: |-
  Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features.
---

# fortimanager_object_wanopt_profile_ftp
Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features.

~> This resource is a sub resource for variable `ftp` of resource `fortimanager_object_wanopt_profile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wanopt_profile_ftp" "trname" {
  byte_caching = "enable"
  log_traffic  = "enable"
  profile      = fortimanager_object_wanopt_profile.trname2.name
  depends_on   = [fortimanager_object_wanopt_profile.trname2]
}

resource "fortimanager_object_wanopt_profile" "trname2" {
  name = "terr-wanopt-profile2"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `byte_caching` - Enable/disable byte-caching for HTTP. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `port` - Single port number or port number range for FTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `prefer_chunking` - Select dynamic or fixed-size data chunking for HTTP WAN Optimization. Valid values: `dynamic`, `fix`.

* `protocol_opt` - Select Protocol specific optimitation or generic TCP optimization. Valid values: `protocol`, `tcp`.

* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `disable`, `enable`.

* `status` - Enable/disable HTTP WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWanopt ProfileFtp can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wanopt_profile_ftp.labelname ObjectWanoptProfileFtp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
