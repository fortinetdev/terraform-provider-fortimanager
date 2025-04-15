---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_profileprotocoloptions_http"
description: |-
  Configure HTTP protocol options.
---

# fortimanager_object_firewall_profileprotocoloptions_http
Configure HTTP protocol options.

~> This resource is a sub resource for variable `http` of resource `fortimanager_object_firewall_profileprotocoloptions`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_profileprotocoloptions_http" "trname" {
  profile_protocol_options = fortimanager_object_firewall_profileprotocoloptions.trname.name
  address_ip_rating        = "enable"
  block_page_status_code   = 499
  comfort_amount           = 200
  depends_on               = [fortimanager_object_firewall_profileprotocoloptions.trname]
}

resource "fortimanager_object_firewall_profileprotocoloptions" "trname" {
  name = "terr-profileprotocoloptions"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile_protocol_options` - Profile Protocol Options.

* `address_ip_rating` - Enable/disable IP based URL rating. Valid values: `disable`, `enable`.

* `block_page_status_code` - Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `domain_fronting` - Configure HTTP domain fronting (default = block). Valid values: `block`, `monitor`, `allow`.

* `h2c` - Enable/disable h2c HTTP connection upgrade. Valid values: `disable`, `enable`.

* `http_09` - Configure action to take upon receipt of HTTP 0.9 request. Valid values: `block`, `allow`.

* `fortinet_bar` - Enable/disable Fortinet bar on HTML content. Valid values: `disable`, `enable`.

* `fortinet_bar_port` - Port for use by Fortinet Bar (1 - 65535, default = 8011).
* `http_policy` - Enable/disable HTTP policy check. Valid values: `disable`, `enable`.

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `chunkedbypass`, `clientcomfort`, `no-content-summary`, `servercomfort`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 80).
* `post_lang` - ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets). Valid values: `jisx0201`, `jisx0208`, `jisx0212`, `gb2312`, `ksc5601-ex`, `euc-jp`, `sjis`, `iso2022-jp`, `iso2022-jp-1`, `iso2022-jp-2`, `euc-cn`, `ces-gbk`, `hz`, `ces-big5`, `euc-kr`, `iso2022-jp-3`, `iso8859-1`, `tis620`, `cp874`, `cp1252`, `cp1251`.

* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `range_block` - Enable/disable blocking of partial downloads. Valid values: `disable`, `enable`.

* `retry_count` - Number of attempts to retry HTTP connection (0 - 100, default = 0).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
* `streaming_content_bypass` - Enable/disable bypassing of streaming content from buffering. Valid values: `disable`, `enable`.

* `strip_x_forwarded_for` - Enable/disable stripping of HTTP X-Forwarded-For header. Valid values: `disable`, `enable`.

* `switching_protocols` - Bypass from scanning, or block a connection that attempts to switch protocol. Valid values: `bypass`, `block`.

* `tcp_window_maximum` - Maximum dynamic TCP window size (default = 8MB).
* `tcp_window_minimum` - Minimum dynamic TCP window size (default = 128KB).
* `tcp_window_size` - Set TCP static window size (default = 256KB).
* `tcp_window_type` - Specify type of TCP window to use for this protocol. Valid values: `system`, `static`, `dynamic`.

* `tunnel_non_http` - Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `unknown_content_encoding` - Configure the action the FortiGate unit will take on unknown content-encoding. Valid values: `block`, `inspect`, `bypass`.

* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `best-effort`, `reject`, `tunnel`.

* `verify_dns_for_policy_matching` - Enable/disable verification of DNS for policy matching. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall ProfileProtocolOptionsHttp can be imported using any of these accepted formats:
```
Set import_options = ["profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_profileprotocoloptions_http.labelname ObjectFirewallProfileProtocolOptionsHttp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
