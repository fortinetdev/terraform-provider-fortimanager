---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_profile_ftgdwf"
description: |-
  FortiGuard Web Filter settings.
---

# fortimanager_object_webfilter_profile_ftgdwf
FortiGuard Web Filter settings.

~> This resource is a sub resource for variable `ftgd_wf` of resource `fortimanager_object_webfilter_profile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `filters`: `fortimanager_object_webfilter_profile_ftgdwf_filters`
>- `quota`: `fortimanager_object_webfilter_profile_ftgdwf_quota`
>- `risk`: `fortimanager_object_webfilter_profile_ftgdwf_risk`



## Example Usage

```hcl
resource "fortimanager_object_webfilter_profile_ftgdwf" "trname" {
  rate_css_urls = "disable"
  rate_crl_urls = "enable"
  profile       = fortimanager_object_webfilter_profile.trname5.name
  depends_on    = [fortimanager_object_webfilter_profile.trname5]
}

resource "fortimanager_object_webfilter_profile" "trname5" {
  name = "terr-webfilter-profile5"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `exempt_quota` - Do not stop quota for these categories.
* `filters` - Filters. The structure of `filters` block is documented below.
* `max_quota_timeout` - Maximum FortiGuard quota used by single page view in seconds (excludes streams).
* `options` - Options for FortiGuard Web Filter. Valid values: `error-allow`, `http-err-detail`, `rate-image-urls`, `strict-blocking`, `rate-server-ip`, `redir-block`, `connect-request-bypass`, `log-all-url`, `ftgd-disable`.

* `ovrd` - Allow web filter profile overrides.
* `quota` - Quota. The structure of `quota` block is documented below.
* `rate_crl_urls` - Enable/disable rating CRL by URL. Valid values: `disable`, `enable`.

* `rate_css_urls` - Enable/disable rating CSS by URL. Valid values: `disable`, `enable`.

* `rate_image_urls` - Enable/disable rating images by URL. Valid values: `disable`, `enable`.

* `rate_javascript_urls` - Enable/disable rating JavaScript by URL. Valid values: `disable`, `enable`.

* `risk` - Risk. The structure of `risk` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filters` block supports:

* `action` - Action to take for matches. Valid values: `block`, `monitor`, `warning`, `authenticate`.

* `auth_usr_grp` - Groups with permission to authenticate.
* `category` - Categories and groups the filter examines.
* `id` - ID number.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `override_replacemsg` - Override replacement message.
* `warn_duration` - Duration of warnings.
* `warning_duration_type` - Re-display warning after closing browser or after a timeout. Valid values: `session`, `timeout`.

* `warning_prompt` - Warning prompts in each category or each domain. Valid values: `per-domain`, `per-category`.


The `quota` block supports:

* `category` - FortiGuard categories to apply quota to (category action must be set to monitor).
* `duration` - Duration of quota.
* `id` - ID number.
* `override_replacemsg` - Override replacement message.
* `type` - Quota type. Valid values: `time`, `traffic`.

* `unit` - Traffic quota unit of measurement. Valid values: `B`, `KB`, `MB`, `GB`.

* `value` - Traffic quota value.

The `risk` block supports:

* `action` - Action to take for matches. Valid values: `block`, `monitor`.

* `id` - ID number.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `risk_level` - Risk level to be examined.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWebfilter ProfileFtgdWf can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_profile_ftgdwf.labelname ObjectWebfilterProfileFtgdWf
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
