---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_internetservicename"
description: |-
  Define internet service names.
---

# fortimanager_object_firewall_internetservicename
Define internet service names.

## Example Usage

```hcl
resource "fortimanager_object_firewall_internetservicename" "trname" {
  country_id          = 1
  internet_service_id = "65536"
  name                = "terr-firewall-int-svs-name"
  type                = "location"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `city_id` - City ID.
* `country_id` - Country or Area ID.
* `internet_service_id` - Internet Service ID.
* `name` - Internet Service name.
* `region_id` - Region ID.
* `type` - Internet Service name type. Valid values: `default`, `location`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall InternetServiceName can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_internetservicename.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
