---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_csf_fabricconnector"
description: |-
  Fabric connector configuration.
---

# fortimanager_system_csf_fabricconnector
Fabric connector configuration.

## Argument Reference


The following arguments are supported:


* `accprofile` - Override access profile.
* `configuration_write_access` - Enable/disable downstream device write access to configuration. disable - Disable downstream device write access to configuration. enable - Enable downstream device write access to configuration. Valid values: `disable`, `enable`.

* `serial` - Serial.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{serial}}.

## Import

System CsfFabricConnector can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_csf_fabricconnector.labelname {{serial}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

