---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_shapingprofile_shapingentries"
description: |-
  Define shaping entries of this shaping profile.
---

# fortimanager_object_firewall_shapingprofile_shapingentries
Define shaping entries of this shaping profile.

~> This resource is a sub resource for variable `shaping_entries` of resource `fortimanager_object_firewall_shapingprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_shapingprofile_shapingentries" "trname" {
  shaping_profile              = fortimanager_object_firewall_shapingprofile.trname.profile_name
  fosid                        = 1
  limit                        = 50
  max                          = 60
  maximum_bandwidth_percentage = 40
  depends_on                   = [fortimanager_object_firewall_shapingprofile.trname]
}

resource "fortimanager_object_firewall_shapingprofile" "trname" {
  profile_name = "terr-shapingprofile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `shaping_profile` - Shaping Profile.

* `burst_in_msec` - Number of bytes that can be burst at maximum-bandwidth speed. Formula: burst = maximum-bandwidth*burst-in-msec.
* `cburst_in_msec` - Number of bytes that can be burst as fast as the interface can transmit. Formula: cburst = maximum-bandwidth*cburst-in-msec.
* `class_id` - Class ID.
* `guaranteed_bandwidth_percentage` - Guaranteed bandwith in percentage.
* `fosid` - ID number.
* `limit` - Hard limit on the real queue size in packets.
* `max` - Average queue size in packets at which RED drop probability is maximal.
* `maximum_bandwidth_percentage` - Maximum bandwith in percentage.
* `min` - Average queue size in packets at which RED drop becomes a possibility.
* `priority` - Priority. Valid values: `low`, `medium`, `high`, `critical`, `top`.

* `red_probability` - Maximum probability (in percentage) for RED marking.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall ShapingProfileShapingEntries can be imported using any of these accepted formats:
```
Set import_options = ["shaping_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_shapingprofile_shapingentries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
