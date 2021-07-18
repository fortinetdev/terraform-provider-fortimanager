---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_interfacepolicy"
description: |-
  Configure IPv4 interface policies.
---

# fortimanager_packages_firewall_interfacepolicy
Configure IPv4 interface policies.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_interfacepolicy" "labelname" {
  address_type               = "ipv4"
  application_list_status    = "disable"
  av_profile_status          = "disable"
  dlp_sensor_status          = "disable"
  dsri                       = "disable"
  dstaddr                    = "all"
  emailfilter_profile_status = "disable"
  interface                  = "1-A14"
  ips_sensor_status          = "disable"
  logtraffic                 = "utm"
  pkg                        = "default"
  policyid                   = 1
  service                    = "ALL"
  srcaddr                    = "all"
  status                     = "enable"
  webfilter_profile_status   = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `address_type` - Address-Type. Valid values: `ipv4`, `ipv6`.

* `application_list` - Application list name.
* `application_list_status` - Enable/disable application control. Valid values: `disable`, `enable`.

* `av_profile` - Antivirus profile.
* `av_profile_status` - Enable/disable antivirus. Valid values: `disable`, `enable`.

* `comments` - Comments.
* `dlp_sensor` - DLP sensor name.
* `dlp_sensor_status` - Enable/disable DLP. Valid values: `disable`, `enable`.

* `dsri` - Enable/disable DSRI. Valid values: `disable`, `enable`.

* `dstaddr` - Address object to limit traffic monitoring to network traffic sent to the specified address or range.
* `emailfilter_profile` - Email filter profile.
* `emailfilter_profile_status` - Enable/disable email filter. Valid values: `disable`, `enable`.

* `interface` - Monitored interface name from available interfaces.
* `ips_sensor` - IPS sensor name.
* `ips_sensor_status` - Enable/disable IPS. Valid values: `disable`, `enable`.

* `label` - Label.
* `logtraffic` - Logging type to be used in this policy (Options: all | utm | disable, Default: utm). Valid values: `disable`, `all`, `utm`.

* `policyid` - Policy ID (0 - 4294967295).
* `service` - Service object from available options.
* `srcaddr` - Address object to limit traffic monitoring to network traffic sent from the specified address or range.
* `status` - Enable/disable this policy. Valid values: `disable`, `enable`.

* `webfilter_profile` - Web filter profile.
* `webfilter_profile_status` - Enable/disable web filtering. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallInterfacePolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_interfacepolicy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
