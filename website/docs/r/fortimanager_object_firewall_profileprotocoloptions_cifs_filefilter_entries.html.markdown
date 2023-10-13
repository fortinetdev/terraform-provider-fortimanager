---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_profileprotocoloptions_cifs_filefilter_entries"
description: |-
  File filter entries.
---

# fortimanager_object_firewall_profileprotocoloptions_cifs_filefilter_entries
File filter entries.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile_protocol_options` - Profile Protocol Options.

* `action` - Action taken for matched file. Valid values: `log`, `block`.

* `comment` - Comment.
* `direction` - Match files transmitted in the session's originating or reply direction. Valid values: `any`, `incoming`, `outgoing`.

* `file_type` - Select file type.
* `filter` - Add a file filter.
* `protocol` - Protocols to apply with. Valid values: `cifs`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{filter}}.

## Import

ObjectFirewall ProfileProtocolOptionsCifsFileFilterEntries can be imported using any of these accepted formats:
```
Set import_options = ["profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_profileprotocoloptions_cifs_filefilter_entries.labelname {{filter}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
