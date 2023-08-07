---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_pblock_firewall_securitypolicy"
description: |-
  /pm/config/pblock/firewall/security-policy
---

# fortimanager_packages_pblock_firewall_securitypolicy
/pm/config/pblock/firewall/security-policy

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pblock` - Pblock.

* `_policy_block` - Assigned policy block.  When this attribute is set, the policy represent a policy block, and all other attributes are ignored. This attribute is not available when configuring policy inside a policy block.
* `action` - Policy action (accept/deny). Valid values: `deny`, `accept`.

* `app_category` - Application category ID list.
* `app_group` - Application group names.
* `application` - Application ID list.
* `application_list` - Name of an existing Application list.
* `av_profile` - Name of an existing Antivirus profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Comment.
* `dlp_profile` - Name of an existing DLP profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dstaddr4` - Destination IPv4 address name and address group names.
* `dstaddr` - Destination IPv4 address name and address group names.
* `dstaddr_negate` - When enabled dstaddr/dstaddr6 specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `dstaddr6` - Destination IPv6 address name and address group names.
* `dstintf` - Outgoing (egress) interface.
* `emailfilter_profile` - Name of an existing email filter profile.
* `enforce_default_app_port` - Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.

* `file_filter_profile` - Name of an existing file-filter profile.
* `fsso_groups` - Names of FSSO groups.
* `global_label` - Label for the policy that appears when the GUI is in Global View mode.
* `groups` - Names of user groups that can authenticate with this policy.
* `icap_profile` - Name of an existing ICAP profile.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `disable`, `enable`.

* `internet_service_custom` - Custom Internet Service name.
* `internet_service_custom_group` - Custom Internet Service group name.
* `internet_service_group` - Internet Service group name.
* `internet_service_id` - Internet Service ID.
* `internet_service_name` - Internet Service name.
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `disable`, `enable`.

* `internet_service_src_custom` - Custom Internet Service source name.
* `internet_service_src_custom_group` - Custom Internet Service source group name.
* `internet_service_src_group` - Internet Service source group name.
* `internet_service_src_id` - Internet Service source ID.
* `internet_service_src_name` - Internet Service source name.
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `internet_service6` - Enable/disable use of IPv6 Internet Services for this policy. If enabled, destination address, service and default application port enforcement are not used. Valid values: `disable`, `enable`.

* `internet_service6_custom` - Custom IPv6 Internet Service name.
* `internet_service6_custom_group` - Custom IPv6 Internet Service group name.
* `internet_service6_group` - Internet Service group name.
* `internet_service6_name` - IPv6 Internet Service name.
* `internet_service6_negate` - When enabled internet-service6 specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `internet_service6_src` - Enable/disable use of IPv6 Internet Services in source for this policy. If enabled, source address is not used. Valid values: `disable`, `enable`.

* `internet_service6_src_custom` - Custom IPv6 Internet Service source name.
* `internet_service6_src_custom_group` - Custom Internet Service6 source group name.
* `internet_service6_src_group` - Internet Service6 source group name.
* `internet_service6_src_name` - IPv6 Internet Service source name.
* `internet_service6_src_negate` - When enabled internet-service6-src specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `ips_sensor` - Name of an existing IPS sensor.
* `learning_mode` - Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated. Valid values: `disable`, `enable`.

* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `disable`, `all`, `utm`.

* `logtraffic_start` - Record logs when a session starts. Valid values: `disable`, `enable`.

* `mms_profile` - Name of an existing MMS profile.
* `name` - Policy name.
* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `nat64` - Enable/disable NAT64. Valid values: `disable`, `enable`.

* `policyid` - Policy ID.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.

* `schedule` - Schedule name.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy. Valid values: `disable`, `enable`.

* `service` - Service and service group names.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `srcaddr4` - Source IPv4 address name and address group names.
* `srcaddr` - Source IPv4 address name and address group names.
* `srcaddr_negate` - When enabled srcaddr/srcaddr6 specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `srcaddr6` - Source IPv6 address name and address group names.
* `srcintf` - Incoming (ingress) interface.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `status` - Enable or disable this policy. Valid values: `disable`, `enable`.

* `url_category` - URL category ID list.
* `users` - Names of individual users that can authenticate with this policy.
* `utm_status` - Enable security profiles. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `voip_profile` - Name of an existing VoIP profile.
* `webfilter_profile` - Name of an existing Web filter profile.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages PblockFirewallSecurityPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pblock=mypblock"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_pblock_firewall_securitypolicy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
