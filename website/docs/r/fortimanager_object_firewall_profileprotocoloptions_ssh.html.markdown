---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_profileprotocoloptions_ssh"
description: |-
  Configure SFTP and SCP protocol options.
---

# fortimanager_object_firewall_profileprotocoloptions_ssh
Configure SFTP and SCP protocol options.

~> This resource is a sub resource for variable `ssh` of resource `fortimanager_object_firewall_profileprotocoloptions`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_profileprotocoloptions_ssh" "trname" {
  profile_protocol_options = fortimanager_object_firewall_profileprotocoloptions.trname.name
  comfort_amount           = 200
  comfort_interval         = 20
  oversize_limit           = 200
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
* `id` - an identifier for the resource.

## Import

ObjectFirewall ProfileProtocolOptionsSsh can be imported using any of these accepted formats:
```
Set import_options = ["profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_profileprotocoloptions_ssh.labelname ObjectFirewallProfileProtocolOptionsSsh
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
