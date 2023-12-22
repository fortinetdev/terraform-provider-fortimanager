---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dnsfilter_profile_ftgddns_filters"
description: |-
  FortiGuard DNS domain filters.
---

# fortimanager_object_dnsfilter_profile_ftgddns_filters
FortiGuard DNS domain filters.

~> This resource is a sub resource for variable `filters` of resource `fortimanager_object_dnsfilter_profile_ftgddns`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_dnsfilter_profile_ftgddns_filters" "trname" {
  profile    = fortimanager_object_dnsfilter_profile.trname.name
  category   = 2
  fosid      = 1
  log        = "enable"
  depends_on = [fortimanager_object_dnsfilter_profile.trname]
}

resource "fortimanager_object_dnsfilter_profile" "trname" {
  name = "terr-dnsfilter-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `action` - Action to take for DNS requests matching the category. Valid values: `monitor`, `block`.

* `category` - Category number.
* `fosid` - ID number.
* `log` - Enable/disable DNS filter logging for this DNS profile. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectDnsfilter ProfileFtgdDnsFilters can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dnsfilter_profile_ftgddns_filters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
