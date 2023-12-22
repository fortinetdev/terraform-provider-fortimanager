---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_waf_profile_urlaccess_accesspattern"
description: |-
  URL access pattern.
---

# fortimanager_object_waf_profile_urlaccess_accesspattern
URL access pattern.

~> This resource is a sub resource for variable `access_pattern` of resource `fortimanager_object_waf_profile_urlaccess`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_waf_profile_urlaccess_accesspattern" "trname" {
  profile    = fortimanager_object_waf_profile.trname.name
  url_access = fortimanager_object_waf_profile_urlaccess.trname2.fosid
  fosid      = 1
  negate     = "enable"
  depends_on = [fortimanager_object_waf_profile_urlaccess.trname2]
}

resource "fortimanager_object_waf_profile_urlaccess" "trname2" {
  fosid      = 13
  profile    = fortimanager_object_waf_profile.trname.name
  depends_on = [fortimanager_object_waf_profile.trname]
}

resource "fortimanager_object_waf_profile" "trname" {
  name = "terr-waf-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.
* `url_access` - Url Access.

* `fosid` - URL access pattern ID.
* `negate` - Enable/disable match negation. Valid values: `disable`, `enable`.

* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `disable`, `enable`.

* `srcaddr` - Source address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWaf ProfileUrlAccessAccessPattern can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE", "url_access=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_waf_profile_urlaccess_accesspattern.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
