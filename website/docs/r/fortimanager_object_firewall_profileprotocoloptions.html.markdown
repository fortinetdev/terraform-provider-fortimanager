---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_profileprotocoloptions"
description: |-
  Configure protocol options.
---

# fortimanager_object_firewall_profileprotocoloptions
Configure protocol options.

## Example Usage

```hcl
resource "fortimanager_object_firewall_profileprotocoloptions" "trname" {
  comment = "terraform-comment"
  name    = "terraform-tefv"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `cifs` - Cifs. The structure of `cifs` block is documented below.
* `comment` - Optional comments.
* `dns` - Dns. The structure of `dns` block is documented below.
* `feature_set` - Feature-Set. Valid values: `proxy`, `flow`.

* `ftp` - Ftp. The structure of `ftp` block is documented below.
* `http` - Http. The structure of `http` block is documented below.
* `imap` - Imap. The structure of `imap` block is documented below.
* `mail_signature` - Mail-Signature. The structure of `mail_signature` block is documented below.
* `mapi` - Mapi. The structure of `mapi` block is documented below.
* `name` - Name.
* `nntp` - Nntp. The structure of `nntp` block is documented below.
* `oversize_log` - Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.

* `pop3` - Pop3. The structure of `pop3` block is documented below.
* `replacemsg_group` - Name of the replacement message group to be used
* `rpc_over_http` - Enable/disable inspection of RPC over HTTP. Valid values: `disable`, `enable`.

* `smtp` - Smtp. The structure of `smtp` block is documented below.
* `ssh` - Ssh. The structure of `ssh` block is documented below.
* `switching_protocols_log` - Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.


The `cifs` block supports:

* `domain_controller` - Domain for which to decrypt CIFS traffic.
* `file_filter` - File-Filter. The structure of `file_filter` block is documented below.
* `options` - One or more options that can be applied to the session. Valid values: `oversize`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 445).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `server_credential_type` - CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.

* `server_keytab` - Server-Keytab. The structure of `server_keytab` block is documented below.
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `tcp_window_maximum` - Maximum dynamic TCP window size (default = 8MB).
* `tcp_window_minimum` - Minimum dynamic TCP window size (default = 128KB).
* `tcp_window_size` - Set TCP static window size (default = 256KB).
* `tcp_window_type` - Specify type of TCP window to use for this protocol. Valid values: `system`, `static`, `dynamic`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).

The `file_filter` block supports:

* `entries` - Entries. The structure of `entries` block is documented below.
* `log` - Enable/disable file filter logging. Valid values: `disable`, `enable`.

* `status` - Enable/disable file filter. Valid values: `disable`, `enable`.


The `entries` block supports:

* `action` - Action taken for matched file. Valid values: `log`, `block`.

* `comment` - Comment.
* `direction` - Match files transmitted in the session's originating or reply direction. Valid values: `any`, `incoming`, `outgoing`.

* `file_type` - Select file type.
* `filter` - Add a file filter.
* `protocol` - Protocols to apply with. Valid values: `cifs`.


The `server_keytab` block supports:

* `keytab` - Base64 encoded keytab file containing credential of the server.
* `password` - Password for keytab.
* `principal` - Service principal.  For example, "host/cifsserver.example.com@example.com".

The `dns` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 53).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.


The `ftp` block supports:

* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `clientcomfort`, `no-content-summary`, `oversize`, `splice`, `bypass-rest-command`, `bypass-mode-command`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 21).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `system`, `static`, `dynamic`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).

The `http` block supports:

* `block_page_status_code` - Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `fortinet_bar` - Fortinet-Bar. Valid values: `disable`, `enable`.

* `fortinet_bar_port` - Fortinet-Bar-Port.
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
* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `best-effort`, `reject`, `tunnel`.


The `imap` block supports:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `fragmail`, `no-content-summary`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 143).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).

The `mail_signature` block supports:

* `signature` - Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).
* `status` - Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate. Valid values: `disable`, `enable`.


The `mapi` block supports:

* `options` - One or more options that can be applied to the session. Valid values: `fragmail`, `oversize`, `no-content-summary`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 135).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).

The `nntp` block supports:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `no-content-summary`, `splice`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 119).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).

The `pop3` block supports:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `fragmail`, `no-content-summary`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 110).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).

The `smtp` block supports:

* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `disable`, `enable`.

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `fragmail`, `no-content-summary`, `splice`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 25).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `server_busy` - Enable/disable SMTP server busy when server not available. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).

The `ssh` block supports:

* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `clientcomfort`, `servercomfort`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `system`, `static`, `dynamic`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ProfileProtocolOptions can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_profileprotocoloptions.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
