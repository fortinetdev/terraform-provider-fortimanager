---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_waf_profile_method_methodpolicy"
description: |-
  HTTP method policy.
---

# fortimanager_object_waf_profile_method_methodpolicy
HTTP method policy.

~> This resource is a sub resource for variable `method_policy` of resource `fortimanager_object_waf_profile_method`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_waf_profile_method_methodpolicy" "trname" {
  fosid      = 1
  pattern    = "https://www.example.com/path/to/resource?param1=value1&param2=value2"
  regex      = "enable"
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

* `address` - Host address.
* `allowed_methods` - Allowed Methods. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `others`, `connect`.

* `fosid` - HTTP method policy ID.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWaf ProfileMethodMethodPolicy can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_waf_profile_method_methodpolicy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
