---
subcategory: "Object WAF"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_waf_profile"
description: |-
  Web application firewall configuration.
---

# fortimanager_object_waf_profile
Web application firewall configuration.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `address_list` - Address-List. The structure of `address_list` block is documented below.
* `comment` - Comment.
* `constraint` - Constraint. The structure of `constraint` block is documented below.
* `extended_log` - Enable/disable extended logging. Valid values: `disable`, `enable`.

* `external` - Disable/Enable external HTTP Inspection. Valid values: `disable`, `enable`.

* `method` - Method. The structure of `method` block is documented below.
* `name` - WAF Profile name.
* `signature` - Signature. The structure of `signature` block is documented below.
* `url_access` - Url-Access. The structure of `url_access` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `address_list` block supports:

* `blocked_address` - Blocked address.
* `blocked_log` - Enable/disable logging on blocked addresses. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.

* `trusted_address` - Trusted address.

The `constraint` block supports:

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


The `method` block supports:

* `default_allowed_methods` - Methods. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `others`, `connect`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `method_policy` - Method-Policy. The structure of `method_policy` block is documented below.
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.


The `method_policy` block supports:

* `address` - Host address.
* `allowed_methods` - Allowed Methods. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `others`, `connect`.

* `id` - HTTP method policy ID.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `disable`, `enable`.


The `signature` block supports:

* `credit_card_detection_threshold` - The minimum number of Credit cards to detect violation.
* `custom_signature` - Custom-Signature. The structure of `custom_signature` block is documented below.
* `disabled_signature` - Disabled signatures
* `disabled_sub_class` - Disabled signature subclasses.
* `main_class` - Main-Class. The structure of `main_class` block is documented below.

The `custom_signature` block supports:

* `action` - Action. Valid values: `allow`, `block`, `erase`.

* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable`, `enable`.

* `direction` - Traffic direction. Valid values: `request`, `response`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `name` - Signature name.
* `pattern` - Match pattern.
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.

* `target` - Match HTTP target. Valid values: `arg`, `arg-name`, `req-body`, `req-cookie`, `req-cookie-name`, `req-filename`, `req-header`, `req-header-name`, `req-raw-uri`, `req-uri`, `resp-body`, `resp-hdr`, `resp-status`.


The `main_class` block supports:

* `action` - Action. Valid values: `allow`, `block`, `erase`.

* `id` - Main signature class ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.


The `url_access` block supports:

* `access_pattern` - Access-Pattern. The structure of `access_pattern` block is documented below.
* `action` - Action. Valid values: `bypass`, `permit`, `block`.

* `address` - Host address.
* `id` - URL access ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.


The `access_pattern` block supports:

* `id` - URL access pattern ID.
* `negate` - Enable/disable match negation. Valid values: `disable`, `enable`.

* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `disable`, `enable`.

* `srcaddr` - Source address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWaf Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_waf_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
