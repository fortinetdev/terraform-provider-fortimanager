---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_gtp_messageratelimitv0"
description: |-
  Message rate limiting for GTP version 0.
---

# fortimanager_object_firewall_gtp_messageratelimitv0
Message rate limiting for GTP version 0.

~> This resource is a sub resource for variable `message_rate_limit_v0` of resource `fortimanager_object_firewall_gtp`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `gtp` - Gtp.

* `create_pdp_request` - Rate limit (packets/s) for create PDP context request.
* `delete_pdp_request` - Rate limit (packets/s) for delete PDP context request.
* `echo_request` - Rate limit (packets/s) for echo request.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall GtpMessageRateLimitV0 can be imported using any of these accepted formats:
```
Set import_options = ["gtp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_gtp_messageratelimitv0.labelname ObjectFirewallGtpMessageRateLimitV0
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
