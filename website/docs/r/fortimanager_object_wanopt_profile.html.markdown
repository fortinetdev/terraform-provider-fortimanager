---
subcategory: "Object WAN-Opt"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wanopt_profile"
description: |-
  Configure WAN optimization profiles.
---

# fortimanager_object_wanopt_profile
Configure WAN optimization profiles.

## Example Usage

```hcl
resource "fortimanager_object_wanopt_profile" "trname" {
  comments    = "This is a Terraform example"
  name        = "terr-wanopt-profile"
  transparent = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auth_group` - Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
* `cifs` - Cifs. The structure of `cifs` block is documented below.
* `comments` - Comment.
* `ftp` - Ftp. The structure of `ftp` block is documented below.
* `http` - Http. The structure of `http` block is documented below.
* `mapi` - Mapi. The structure of `mapi` block is documented below.
* `name` - Profile name.
* `tcp` - Tcp. The structure of `tcp` block is documented below.
* `transparent` - Enable/disable transparent mode. Valid values: `disable`, `enable`.


The `cifs` block supports:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `port` - Single port number or port number range for CIFS. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `prefer_chunking` - Select dynamic or fixed-size data chunking for WAN Optimization. Valid values: `dynamic`, `fix`.

* `protocol_opt` - Select Protocol specific optimitation or generic TCP optimization. Valid values: `protocol`, `tcp`.

* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.


The `ftp` block supports:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `port` - Single port number or port number range for FTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `prefer_chunking` - Select dynamic or fixed-size data chunking for WAN Optimization. Valid values: `dynamic`, `fix`.

* `protocol_opt` - Select Protocol specific optimitation or generic TCP optimization. Valid values: `protocol`, `tcp`.

* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `disable`, `enable`.

* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.


The `http` block supports:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `port` - Single port number or port number range for HTTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `prefer_chunking` - Select dynamic or fixed-size data chunking for WAN Optimization. Valid values: `dynamic`, `fix`.

* `protocol_opt` - Select Protocol specific optimitation or generic TCP optimization. Valid values: `protocol`, `tcp`.

* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `disable`, `enable`.

* `ssl_port` - Port on which to expect HTTPS traffic for SSL/TLS offloading.
* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_non_http` - Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.

* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `best-effort`, `reject`, `tunnel`.


The `mapi` block supports:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `port` - Single port number or port number range for MAPI. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.


The `tcp` block supports:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `byte_caching_opt` - Select whether TCP byte-caching uses system memory only or both memory and disk space. Valid values: `mem-only`, `mem-disk`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `port` - Port numbers or port number ranges for TCP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `disable`, `enable`.

* `ssl_port` - Port numbers or port number ranges on which to expect HTTPS traffic for SSL/TLS offloading.
* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWanopt Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wanopt_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
