---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_waf_profile_constraint"
description: |-
  WAF HTTP protocol restrictions.
---

# fortimanager_object_waf_profile_constraint
WAF HTTP protocol restrictions.

~> This resource is a sub resource for variable `constraint` of resource `fortimanager_object_waf_profile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`content_length`: `fortimanager_object_waf_profile_constraint_contentlength`
`exception`: `fortimanager_object_waf_profile_constraint_exception`
`header_length`: `fortimanager_object_waf_profile_constraint_headerlength`
`hostname`: `fortimanager_object_waf_profile_constraint_hostname`
`line_length`: `fortimanager_object_waf_profile_constraint_linelength`
`malformed`: `fortimanager_object_waf_profile_constraint_malformed`
`max_cookie`: `fortimanager_object_waf_profile_constraint_maxcookie`
`max_header_line`: `fortimanager_object_waf_profile_constraint_maxheaderline`
`max_range_segment`: `fortimanager_object_waf_profile_constraint_maxrangesegment`
`max_url_param`: `fortimanager_object_waf_profile_constraint_maxurlparam`
`method`: `fortimanager_object_waf_profile_constraint_method`
`param_length`: `fortimanager_object_waf_profile_constraint_paramlength`
`url_param_length`: `fortimanager_object_waf_profile_constraint_urlparamlength`
`version`: `fortimanager_object_waf_profile_constraint_version`



## Example Usage

```hcl
resource "fortimanager_object_waf_profile_constraint" "trname" {
  content_length {
    action = "allow"
    length = 120
    log    = "enable"
    status = "disable"
  }
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

* `content_length` - Content-Length. The structure of `content_length` block is documented below.
* `exception` - Exception. The structure of `exception` block is documented below.
* `header_length` - Header-Length. The structure of `header_length` block is documented below.
* `hostname` - Hostname. The structure of `hostname` block is documented below.
* `line_length` - Line-Length. The structure of `line_length` block is documented below.
* `malformed` - Malformed. The structure of `malformed` block is documented below.
* `max_cookie` - Max-Cookie. The structure of `max_cookie` block is documented below.
* `max_header_line` - Max-Header-Line. The structure of `max_header_line` block is documented below.
* `max_range_segment` - Max-Range-Segment. The structure of `max_range_segment` block is documented below.
* `max_url_param` - Max-Url-Param. The structure of `max_url_param` block is documented below.
* `method` - Method. The structure of `method` block is documented below.
* `param_length` - Param-Length. The structure of `param_length` block is documented below.
* `url_param_length` - Url-Param-Length. The structure of `url_param_length` block is documented below.
* `version` - Version. The structure of `version` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `content_length` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Length of HTTP content in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `exception` block supports:

* `address` - Host address.
* `content_length` - HTTP content length in request. Valid values: `disable`, `enable`.

* `header_length` - HTTP header length in request. Valid values: `disable`, `enable`.

* `hostname` - Enable/disable hostname check. Valid values: `disable`, `enable`.

* `id` - Exception ID.
* `line_length` - HTTP line length in request. Valid values: `disable`, `enable`.

* `malformed` - Enable/disable malformed HTTP request check. Valid values: `disable`, `enable`.

* `max_cookie` - Maximum number of cookies in HTTP request. Valid values: `disable`, `enable`.

* `max_header_line` - Maximum number of HTTP header line. Valid values: `disable`, `enable`.

* `max_range_segment` - Maximum number of range segments in HTTP range line. Valid values: `disable`, `enable`.

* `max_url_param` - Maximum number of parameters in URL. Valid values: `disable`, `enable`.

* `method` - Enable/disable HTTP method check. Valid values: `disable`, `enable`.

* `param_length` - Maximum length of parameter in URL, HTTP POST request or HTTP body. Valid values: `disable`, `enable`.

* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `disable`, `enable`.

* `url_param_length` - Maximum length of parameter in URL. Valid values: `disable`, `enable`.

* `version` - Enable/disable HTTP version check. Valid values: `disable`, `enable`.


The `header_length` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Length of HTTP header in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `hostname` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `line_length` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Length of HTTP line in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `malformed` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `max_cookie` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `max_cookie` - Maximum number of cookies in HTTP request (0 to 2147483647).
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `max_header_line` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `max_header_line` - Maximum number HTTP header lines (0 to 2147483647).
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `max_range_segment` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `max_range_segment` - Maximum number of range segments in HTTP range line (0 to 2147483647).
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `max_url_param` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `max_url_param` - Maximum number of parameters in URL (0 to 2147483647).
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `method` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `param_length` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Maximum length of parameter in URL, HTTP POST request or HTTP body in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `url_param_length` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Maximum length of URL parameter in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `version` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWaf ProfileConstraint can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_waf_profile_constraint.labelname ObjectWafProfileConstraint
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
