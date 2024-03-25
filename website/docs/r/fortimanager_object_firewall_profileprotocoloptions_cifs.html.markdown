---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_profileprotocoloptions_cifs"
description: |-
  Configure CIFS protocol options.
---

# fortimanager_object_firewall_profileprotocoloptions_cifs
Configure CIFS protocol options.

~> This resource is a sub resource for variable `cifs` of resource `fortimanager_object_firewall_profileprotocoloptions`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `file_filter`: `fortimanager_object_firewall_profileprotocoloptions_cifs_filefilter`
>- `server_keytab`: `fortimanager_object_firewall_profileprotocoloptions_cifs_serverkeytab`



## Example Usage

```hcl
resource "fortimanager_object_firewall_profileprotocoloptions_cifs" "trname" {
  profile_protocol_options = fortimanager_object_firewall_profileprotocoloptions.trname.name
  options                  = ["oversize"]
  oversize_limit           = 200
  ports                    = [445]
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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall ProfileProtocolOptionsCifs can be imported using any of these accepted formats:
```
Set import_options = ["profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_profileprotocoloptions_cifs.labelname ObjectFirewallProfileProtocolOptionsCifs
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
