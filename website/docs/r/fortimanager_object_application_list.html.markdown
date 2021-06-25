---
subcategory: "Object Application"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_application_list"
description: |-
  Configure application control lists.
---

# fortimanager_object_application_list
Configure application control lists.

## Example Usage

```hcl
resource "fortimanager_object_application_list" "trname" {
  app_replacemsg             = "enable"
  comment                    = "terraform-tefv-comment"
  deep_app_inspection        = "enable"
  extended_log               = "disable"
  name                       = "terraform-tefv"
  other_application_action   = "pass"
  other_application_log      = "disable"
  unknown_application_action = "pass"
  unknown_application_log    = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `app_replacemsg` - Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.

* `comment` - comments
* `control_default_network_services` - Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.

* `deep_app_inspection` - Enable/disable deep application inspection. Valid values: `disable`, `enable`.

* `default_network_services` - Default-Network-Services. The structure of `default_network_services` block is documented below.
* `enforce_default_app_port` - Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.

* `entries` - Entries. The structure of `entries` block is documented below.
* `extended_log` - Enable/disable extended logging. Valid values: `disable`, `enable`.

* `force_inclusion_ssl_di_sigs` - Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.

* `name` - List name.
* `options` - Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.

* `other_application_action` - Action for other applications. Valid values: `pass`, `block`.

* `other_application_log` - Enable/disable logging for other applications. Valid values: `disable`, `enable`.

* `p2p_block_list` - P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `p2p_black_list` - P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
 (`ver Controlled FortiOS <= 6.4`)
* `replacemsg_group` - Replacement message group.
* `unknown_application_action` - Pass or block traffic from unknown applications. Valid values: `pass`, `block`.

* `unknown_application_log` - Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `default_network_services` block supports:

* `id` - Entry ID.
* `port` - Port number.
* `services` - Network protocols. Valid values: `http`, `ssh`, `telnet`, `ftp`, `dns`, `smtp`, `pop3`, `imap`, `snmp`, `nntp`, `https`.

* `violation_action` - Action for protocols not white listed under selected port. Valid values: `block`, `monitor`, `pass`.


The `entries` block supports:

* `action` - Pass or block traffic, or reset connection for traffic from this application. Valid values: `pass`, `block`, `reset`.

* `application` - ID of allowed applications.
* `behavior` - Application behavior filter.
* `category` - Category ID list.
* `exclusion` - ID of excluded applications.
* `id` - Entry ID.
* `log` - Enable/disable logging for this application list. Valid values: `disable`, `enable`.

* `log_packet` - Enable/disable packet logging. Valid values: `disable`, `enable`.

* `parameters` - Parameters. The structure of `parameters` block is documented below.
* `per_ip_shaper` - Per-IP traffic shaper.
* `popularity` - Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.

* `protocols` - Application protocol filter.
* `quarantine` - Quarantine method. Valid values: `none`, `attacker`.

* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable`, `enable`.

* `rate_count` - Count of the rate.
* `rate_duration` - Duration (sec) of the rate.
* `rate_mode` - Rate limit mode. Valid values: `periodical`, `continuous`.

* `rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`, `dhcp-client-mac`, `dns-domain`.

* `risk` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).
* `session_ttl` - Session TTL (0 = default).
* `shaper` - Traffic shaper.
* `shaper_reverse` - Reverse traffic shaper.
* `sub_category` - Application Sub-category ID list. (`ver Controlled FortiOS <= 6.4`)
* `technology` - Application technology filter.
* `vendor` - Application vendor filter.

The `parameters` block supports:

* `id` - Parameter ID.
* `members` - Members. The structure of `members` block is documented below. (`ver Controlled FortiOS >= 6.4`)
* `value` - Parameter value. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)

The `members` block supports (`ver Controlled FortiOS >= 6.4`):

* `id` - Parameter.
* `name` - Parameter name.
* `value` - Parameter value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectApplication List can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_application_list.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
