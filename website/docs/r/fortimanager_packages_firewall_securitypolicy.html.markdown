---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_securitypolicy"
description: |-
  Configure NGFW IPv4/IPv6 application policies.
---

# fortimanager_packages_firewall_securitypolicy
Configure NGFW IPv4/IPv6 application policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Policy action (accept/deny). Valid values: `deny`, `accept`.

* `app_category` - Application category ID list.
* `app_group` - Application group names.
* `application` - Application ID list.
* `application_list` - Name of an existing Application list.
* `av_profile` - Name of an existing Antivirus profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Comment.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dstaddr` - Destination IPv4 address name and address group names. (`ver Controlled FortiOS >= 6.4`)
* `dstaddr_negate` - When enabled dstaddr/dstaddr6 specifies what the destination address must NOT be. Valid values: `disable`, `enable`.
 (`ver Controlled FortiOS >= 6.4`)
* `dstaddr4` - Destination IPv4 address name and address group names. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `dstaddr6` - Destination IPv6 address name and address group names.
* `dstintf` - Outgoing (egress) interface.
* `emailfilter_profile` - Name of an existing email filter profile.
* `enforce_default_app_port` - Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.

* `file_filter_profile` - Name of an existing file-filter profile. (`ver Controlled FortiOS >= 6.4`)
* `fsso_groups` - Names of FSSO groups.
* `global_label` - Label for the policy that appears when the GUI is in Global View mode.
* `groups` - Names of user groups that can authenticate with this policy.
* `icap_profile` - Name of an existing ICAP profile.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `disable`, `enable`.

* `internet_service_custom` - Custom Internet Service name.
* `internet_service_custom_group` - Custom Internet Service group name.
* `internet_service_group` - Internet Service group name.
* `internet_service_name` - Internet Service name. (`ver Controlled FortiOS >= 6.4`)
* `internet_service_id` - Internet Service ID. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `disable`, `enable`.

* `internet_service_src_custom` - Custom Internet Service source name.
* `internet_service_src_custom_group` - Custom Internet Service source group name.
* `internet_service_src_group` - Internet Service source group name.
* `internet_service_src_name` - Internet Service source name. (`ver Controlled FortiOS >= 6.4`)
* `internet_service_src_id` - Internet Service source ID. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `ips_sensor` - Name of an existing IPS sensor.
* `learning_mode` - Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated. Valid values: `disable`, `enable`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `disable`, `all`, `utm`.

* `logtraffic_start` - Record logs when a session starts. Valid values: `disable`, `enable`.
 (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `mms_profile` - Name of an existing MMS profile. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `name` - Policy name.
* `policyid` - Policy ID.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.

* `schedule` - Schedule name.
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy. Valid values: `disable`, `enable`.

* `service` - Service and service group names.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `srcaddr` - Source IPv4 address name and address group names. (`ver Controlled FortiOS >= 6.4`)
* `srcaddr_negate` - When enabled srcaddr/srcaddr6 specifies what the source address must NOT be. Valid values: `disable`, `enable`.
 (`ver Controlled FortiOS >= 6.4`)
* `srcaddr4` - Source IPv4 address name and address group names. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `srcaddr6` - Source IPv6 address name and address group names.
* `srcintf` - Incoming (ingress) interface.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `status` - Enable or disable this policy. Valid values: `disable`, `enable`.

* `url_category` - URL category ID list.
* `users` - Names of individual users that can authenticate with this policy.
* `utm_status` - Enable security profiles. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `voip_profile` - Name of an existing VoIP profile.
* `webfilter_profile` - Name of an existing Web filter profile.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallSecurityPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_securitypolicy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
