---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_urlfilter_entries_move"
description: |-
  URL filter entries.
---

# fortimanager_object_webfilter_urlfilter_entries_move
URL filter entries.

## Example Usage

```hcl
resource "fortimanager_object_webfilter_urlfilter_entries_move" "trname" {
  urlfilter  = fortimanager_object_webfilter_urlfilter.trname.fosid
  entries    = fortimanager_object_webfilter_urlfilter_entries.trname.fosid
  target     = fortimanager_object_webfilter_urlfilter_entries.trname2.fosid
  option     = "before"
  depends_on = [fortimanager_object_webfilter_urlfilter_entries.trname2, fortimanager_object_webfilter_urlfilter_entries.trname]
}

resource "fortimanager_object_webfilter_urlfilter_entries" "trname2" {
  fosid      = 4
  url        = "www.example.com/path/to/resource?param1=value1&param2=value2"
  urlfilter  = fortimanager_object_webfilter_urlfilter.trname.fosid
  depends_on = [fortimanager_object_webfilter_urlfilter.trname]
}

resource "fortimanager_object_webfilter_urlfilter_entries" "trname" {
  fosid      = 3
  url        = "www.example.com/path/to/resource?param1=value1&param3=value3"
  urlfilter  = fortimanager_object_webfilter_urlfilter.trname.fosid
  depends_on = [fortimanager_object_webfilter_urlfilter.trname]
}

resource "fortimanager_object_webfilter_urlfilter" "trname" {
  comment               = "This is a Terraform example"
  fosid                 = 1
  ip_addr_block         = "enable"
  name                  = "terr-webfilter-urlfilter"
  one_arm_ips_urlfilter = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `urlfilter` - Urlfilter.
* `entries` - Entries.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the entries changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of entry.
