---
subcategory: "Object ICAP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_icap_profile"
description: |-
  Configure ICAP profiles.
---

# fortimanager_object_icap_profile
Configure ICAP profiles.

## Example Usage

```hcl
resource "fortimanager_object_icap_profile" "trname" {
  methods                  = ["delete", "get", "head", "options", "other", "post", "put", "trace"]
  name                     = "terr-icap-profile"
  preview                  = "disable"
  preview_data_length      = 500
  request                  = "disable"
  respmod_default_action   = "forward"
  response                 = "disable"
  response_req_hdr         = "disable"
  streaming_content_bypass = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `chunk_encap` - Enable/disable chunked encapsulation (default = disable). Valid values: `disable`, `enable`.

* `extension_feature` - Enable/disable ICAP extension features. Valid values: `scan-progress`.

* `icap_block_log` - Enable/disable UTM log when infection found (default = disable). Valid values: `disable`, `enable`.

* `icap_headers` - Icap-Headers. The structure of `icap_headers` block is documented below.
* `methods` - The allowed HTTP methods that will be sent to ICAP server for further processing. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `other`.

* `name` - ICAP profile name.
* `preview` - Enable/disable preview of data to ICAP server. Valid values: `disable`, `enable`.

* `preview_data_length` - Preview data length to be sent to ICAP server.
* `replacemsg_group` - Replacement message group.
* `request` - Enable/disable whether an HTTP request is passed to an ICAP server. Valid values: `disable`, `enable`.

* `request_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP request. Valid values: `error`, `bypass`.

* `request_path` - Path component of the ICAP URI that identifies the HTTP request processing service.
* `request_server` - ICAP server to use for an HTTP request.
* `respmod_default_action` - Default action to ICAP response modification (respmod) processing. Valid values: `bypass`, `forward`.

* `respmod_forward_rules` - Respmod-Forward-Rules. The structure of `respmod_forward_rules` block is documented below.
* `response` - Enable/disable whether an HTTP response is passed to an ICAP server. Valid values: `disable`, `enable`.

* `response_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP response. Valid values: `error`, `bypass`.

* `response_path` - Path component of the ICAP URI that identifies the HTTP response processing service.
* `response_req_hdr` - Enable/disable addition of req-hdr for ICAP response modification (respmod) processing. Valid values: `disable`, `enable`.

* `response_server` - ICAP server to use for an HTTP response.
* `scan_progress_interval` - Scan progress interval value.
* `streaming_content_bypass` - Enable/disable bypassing of ICAP server for streaming content. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `icap_headers` block supports:

* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.

* `content` - HTTP header content.
* `id` - HTTP forwarded header ID.
* `name` - HTTP forwarded header name.

The `respmod_forward_rules` block supports:

* `action` - Action to be taken for ICAP server. Valid values: `bypass`, `forward`.

* `header_group` - Header-Group. The structure of `header_group` block is documented below.
* `host` - Address object for the host.
* `http_resp_status_code` - HTTP response status code.
* `name` - Address name.

The `header_group` block supports:

* `case_sensitivity` - Enable/disable case sensitivity when matching header. Valid values: `disable`, `enable`.

* `header` - HTTP header regular expression.
* `header_name` - HTTP header.
* `id` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectIcap Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_icap_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
