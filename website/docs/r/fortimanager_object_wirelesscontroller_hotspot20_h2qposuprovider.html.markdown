---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovider"
description: |-
  Configure online sign up (OSU) provider list.
---

# fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovider
Configure online sign up (OSU) provider list.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`friendly_name`: `fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovider_friendlyname`
`service_description`: `fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovider_servicedescription`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovider" "trname" {
  name       = "terr-wictl-hot20-heqp-osu-provider"
  osu_method = ["oma-dm"]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `friendly_name` - Friendly-Name. The structure of `friendly_name` block is documented below.
* `icon` - OSU provider icon.
* `name` - OSU provider ID.
* `osu_method` - OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.

* `osu_nai` - OSU NAI.
* `server_uri` - Server URI.
* `service_description` - Service-Description. The structure of `service_description` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `friendly_name` block supports:

* `friendly_name` - OSU provider friendly name.
* `index` - OSU provider friendly name index.
* `lang` - Language code.

The `service_description` block supports:

* `lang` - Language code.
* `service_description` - Service description.
* `service_id` - OSU service ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20H2QpOsuProvider can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovider.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
