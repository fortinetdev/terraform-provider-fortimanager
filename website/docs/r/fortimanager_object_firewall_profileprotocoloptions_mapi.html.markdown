---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_profileprotocoloptions_mapi"
description: |-
  Configure MAPI protocol options.
---

# fortimanager_object_firewall_profileprotocoloptions_mapi
Configure MAPI protocol options.

~> This resource is a sub resource for variable `mapi` of resource `fortimanager_object_firewall_profileprotocoloptions`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_profileprotocoloptions_mapi" "trname" {
  profile_protocol_options = fortimanager_object_firewall_profileprotocoloptions.trname.name
  options                  = ["fragmail"]
  oversize_limit           = 200
  ports                    = [123, 124]
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

* `options` - One or more options that can be applied to the session. Valid values: `fragmail`, `oversize`, `no-content-summary`.

* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `ports` - Ports to scan for content (1 - 65535, default = 135).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `disable`, `enable`.

* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `disable`, `enable`.

* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall ProfileProtocolOptionsMapi can be imported using any of these accepted formats:
```
Set import_options = ["profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_profileprotocoloptions_mapi.labelname ObjectFirewallProfileProtocolOptionsMapi
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
