---
subcategory: "ObjectFirewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_profilegroup"
description: |-
  Configure profile groups.
---

# fortimanager_object_firewall_profilegroup
Configure profile groups.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `application_list` - Name of an existing Application list.
* `av_profile` - Name of an existing Antivirus profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `emailfilter_profile` - Name of an existing email filter profile.
* `file_filter_profile` - Name of an existing file-filter profile.
* `icap_profile` - Name of an existing ICAP profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `name` - Profile group name.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `voip_profile` - Name of an existing VoIP profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `webfilter_profile` - Name of an existing Web filter profile.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ProfileGroup can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_profilegroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
