---
subcategory: "Object Webfilter"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_profile"
description: |-
  Configure Web filter profiles.
---

# fortimanager_object_webfilter_profile
Configure Web filter profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`antiphish`: `fortimanager_object_webfilter_profile_antiphish`
`file_filter`: `fortimanager_object_webfilter_profile_filefilter`
`ftgd_wf`: `fortimanager_object_webfilter_profile_ftgdwf`
`override`: `fortimanager_object_webfilter_profile_override`
`url_extraction`: `fortimanager_object_webfilter_profile_urlextraction`
`web`: `fortimanager_object_webfilter_profile_web`
`youtube_channel_filter`: `fortimanager_object_webfilter_profile_youtubechannelfilter`



## Example Usage

```hcl
resource "fortimanager_object_webfilter_profile" "trname" {
  comment                      = "This is a Terraform example"
  extended_log                 = "disable"
  feature_set                  = "flow"
  https_replacemsg             = "enable"
  log_all_url                  = "disable"
  name                         = "terr-webfilter-profile"
  options                      = ["js", "jscript"]
  ovrd_perm                    = ["bannedword-override"]
  post_action                  = "block"
  web_content_log              = "enable"
  web_extended_all_action_log  = "disable"
  web_filter_command_block_log = "enable"
  web_filter_cookie_log        = "enable"
  web_ftgd_err_log             = "enable"
  web_invalid_domain_log       = "enable"
  web_url_log                  = "enable"
  wisp                         = "disable"
  wisp_algorithm               = "auto-learning"
  youtube_channel_status       = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `antiphish` - Antiphish. The structure of `antiphish` block is documented below.
* `comment` - Optional comments.
* `extended_log` - Enable/disable extended logging for web filtering. Valid values: `disable`, `enable`.

* `file_filter` - File-Filter. The structure of `file_filter` block is documented below.
* `feature_set` - Flow/proxy feature set. Valid values: `proxy`, `flow`.

* `ftgd_wf` - Ftgd-Wf. The structure of `ftgd_wf` block is documented below.
* `https_replacemsg` - Enable replacement messages for HTTPS. Valid values: `disable`, `enable`.

* `inspection_mode` - Web filtering inspection mode. Valid values: `proxy`, `flow-based`, `dns`.

* `log_all_url` - Enable/disable logging all URLs visited. Valid values: `disable`, `enable`.

* `name` - Profile name.
* `options` - Options. Valid values: `block-invalid-url`, `jscript`, `js`, `vbs`, `unknown`, `wf-referer`, `https-scan`, `intrinsic`, `wf-cookie`, `per-user-bwl`, `activexfilter`, `cookiefilter`, `https-url-scan`, `javafilter`, `rangeblock`, `contenttype-check`.

* `override` - Override. The structure of `override` block is documented below.
* `ovrd_perm` - Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.

* `post_action` - Action taken for HTTP POST traffic. Valid values: `normal`, `comfort`, `block`.

* `replacemsg_group` - Replacement message group.
* `url_extraction` - Url-Extraction. The structure of `url_extraction` block is documented below.
* `web` - Web. The structure of `web` block is documented below.
* `web_antiphishing_log` - Enable/disable logging of AntiPhishing checks. Valid values: `disable`, `enable`.

* `web_content_log` - Enable/disable logging logging blocked web content. Valid values: `disable`, `enable`.

* `web_extended_all_action_log` - Enable/disable extended any filter action logging for web filtering. Valid values: `disable`, `enable`.

* `web_filter_activex_log` - Enable/disable logging ActiveX. Valid values: `disable`, `enable`.

* `web_filter_applet_log` - Enable/disable logging Java applets. Valid values: `disable`, `enable`.

* `web_filter_command_block_log` - Enable/disable logging blocked commands. Valid values: `disable`, `enable`.

* `web_filter_cookie_log` - Enable/disable logging cookie filtering. Valid values: `disable`, `enable`.

* `web_filter_cookie_removal_log` - Enable/disable logging blocked cookies. Valid values: `disable`, `enable`.

* `web_filter_js_log` - Enable/disable logging Java scripts. Valid values: `disable`, `enable`.

* `web_filter_jscript_log` - Enable/disable logging JScripts. Valid values: `disable`, `enable`.

* `web_filter_referer_log` - Enable/disable logging referrers. Valid values: `disable`, `enable`.

* `web_filter_unknown_log` - Enable/disable logging unknown scripts. Valid values: `disable`, `enable`.

* `web_filter_vbs_log` - Enable/disable logging VBS scripts. Valid values: `disable`, `enable`.

* `web_ftgd_err_log` - Enable/disable logging rating errors. Valid values: `disable`, `enable`.

* `web_ftgd_quota_usage` - Enable/disable logging daily quota usage. Valid values: `disable`, `enable`.

* `web_invalid_domain_log` - Enable/disable logging invalid domain names. Valid values: `disable`, `enable`.

* `web_url_log` - Enable/disable logging URL filtering. Valid values: `disable`, `enable`.

* `wisp` - Enable/disable web proxy WISP. Valid values: `disable`, `enable`.

* `wisp_algorithm` - WISP server selection algorithm. Valid values: `auto-learning`, `primary-secondary`, `round-robin`.

* `wisp_servers` - WISP servers.
* `youtube_channel_filter` - Youtube-Channel-Filter. The structure of `youtube_channel_filter` block is documented below.
* `youtube_channel_status` - YouTube channel filter status. Valid values: `disable`, `blacklist`, `whitelist`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `antiphish` block supports:

* `authentication` - Authentication methods. Valid values: `domain-controller`, `ldap`.

* `check_basic_auth` - Enable/disable checking of HTTP Basic Auth field for known credentials. Valid values: `disable`, `enable`.

* `check_uri` - Enable/disable checking of GET URI parameters for known credentials. Valid values: `disable`, `enable`.

* `check_username_only` - Enable/disable acting only on valid username credentials. Action will be taken for valid usernames regardless of password validity. Valid values: `disable`, `enable`.

* `custom_patterns` - Custom-Patterns. The structure of `custom_patterns` block is documented below.
* `default_action` - Action to be taken when there is no matching rule. Valid values: `log`, `block`, `exempt`.

* `domain_controller` - Domain for which to verify received credentials against.
* `inspection_entries` - Inspection-Entries. The structure of `inspection_entries` block is documented below.
* `ldap` - LDAP server for which to verify received credentials against.
* `max_body_len` - Maximum size of a POST body to check for credentials.
* `status` - Toggle AntiPhishing functionality. Valid values: `disable`, `enable`.


The `custom_patterns` block supports:

* `category` - Category that the pattern matches. Valid values: `username`, `password`.

* `pattern` - Target pattern.
* `type` - Pattern will be treated either as a regex pattern or literal string. Valid values: `regex`, `literal`.


The `inspection_entries` block supports:

* `action` - Action to be taken upon an AntiPhishing match. Valid values: `log`, `block`, `exempt`.

* `fortiguard_category` - FortiGuard category to match.
* `name` - Inspection target name.

The `file_filter` block supports:

* `entries` - Entries. The structure of `entries` block is documented below.
* `log` - Enable/disable file filter logging. Valid values: `disable`, `enable`.

* `scan_archive_contents` - Enable/disable file filter archive contents scan. Valid values: `disable`, `enable`.

* `status` - Enable/disable file filter. Valid values: `disable`, `enable`.


The `entries` block supports:

* `action` - Action taken for matched file. Valid values: `log`, `block`.

* `comment` - Comment.
* `direction` - Match files transmitted in the session's originating or reply direction. Valid values: `any`, `incoming`, `outgoing`.

* `encryption` - Encryption. Valid values: `any`, `yes`.

* `file_type` - Select file type.
* `filter` - Add a file filter.
* `password_protected` - Match password-protected files. Valid values: `any`, `yes`.

* `protocol` - Protocols to apply with. Valid values: `http`, `ftp`.


The `ftgd_wf` block supports:

* `exempt_quota` - Do not stop quota for these categories.
* `filters` - Filters. The structure of `filters` block is documented below.
* `max_quota_timeout` - Maximum FortiGuard quota used by single page view in seconds (excludes streams).
* `options` - Options for FortiGuard Web Filter. Valid values: `error-allow`, `http-err-detail`, `rate-image-urls`, `strict-blocking`, `rate-server-ip`, `redir-block`, `connect-request-bypass`, `log-all-url`, `ftgd-disable`.

* `ovrd` - Allow web filter profile overrides.
* `quota` - Quota. The structure of `quota` block is documented below.
* `rate_crl_urls` - Enable/disable rating CRL by URL. Valid values: `disable`, `enable`.

* `rate_css_urls` - Enable/disable rating CSS by URL. Valid values: `disable`, `enable`.

* `rate_image_urls` - Rate-Image-Urls. Valid values: `disable`, `enable`.

* `rate_javascript_urls` - Enable/disable rating JavaScript by URL. Valid values: `disable`, `enable`.


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

The `override` block supports:

* `ovrd_cookie` - Allow/deny browser-based (cookie) overrides. Valid values: `deny`, `allow`.

* `ovrd_dur` - Override duration.
* `ovrd_dur_mode` - Override duration mode. Valid values: `constant`, `ask`.

* `ovrd_scope` - Override scope. Valid values: `user`, `user-group`, `ip`, `ask`, `browser`.

* `ovrd_user_group` - User groups with permission to use the override.
* `profile` - Web filter profile with permission to create overrides.
* `profile_attribute` - Profile attribute to retrieve from the RADIUS server. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Vendor-Specific`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `profile_type` - Override profile type. Valid values: `list`, `radius`.


The `url_extraction` block supports:

* `redirect_header` - HTTP header name to use for client redirect on blocked requests
* `redirect_no_content` - Enable / Disable empty message-body entity in HTTP response Valid values: `disable`, `enable`.

* `redirect_url` - HTTP header value to use for client redirect on blocked requests
* `server_fqdn` - URL extraction server FQDN (fully qualified domain name)
* `status` - Enable URL Extraction Valid values: `disable`, `enable`.


The `web` block supports:

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


The `youtube_channel_filter` block supports:

* `channel_id` - YouTube channel ID to be filtered.
* `comment` - Comment.
* `id` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWebfilter Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
