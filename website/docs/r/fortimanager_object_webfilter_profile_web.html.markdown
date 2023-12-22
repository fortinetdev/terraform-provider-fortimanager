---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_profile_web"
description: |-
  Web content filtering settings.
---

# fortimanager_object_webfilter_profile_web
Web content filtering settings.

~> This resource is a sub resource for variable `web` of resource `fortimanager_object_webfilter_profile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_webfilter_profile_web" "trname" {
  safe_search      = ["url"]
  youtube_restrict = "strict"
  depends_on       = [fortimanager_object_webfilter_profile.trname8]
  profile          = fortimanager_object_webfilter_profile.trname8.name
}

resource "fortimanager_object_webfilter_profile" "trname8" {
  name = "terr-webfilter-profile8"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `allowlist` - FortiGuard allowlist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.

* `blocklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blocklist. Valid values: `disable`, `enable`.

* `blacklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blacklist. Valid values: `disable`, `enable`.

* `bword_table` - Banned word table ID.
* `bword_threshold` - Banned word score threshold.
* `content_header_list` - Content header list.
* `keyword_match` - Search keywords to log when match is found.
* `log_search` - Enable/disable logging all search phrases. Valid values: `disable`, `enable`.

* `safe_search` - Safe search type. Valid values: `google`, `yahoo`, `bing`, `url`, `header`.

* `urlfilter_table` - URL filter table ID.
* `vimeo_restrict` - Set Vimeo-restrict ("7" = don't show mature content, "134" = don't show unrated and mature content). A value of cookie "content_rating".
* `whitelist` - FortiGuard whitelist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.

* `youtube_restrict` - YouTube EDU filter level. Valid values: `strict`, `none`, `moderate`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWebfilter ProfileWeb can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_profile_web.labelname ObjectWebfilterProfileWeb
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
