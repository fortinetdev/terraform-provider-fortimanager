---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_system_emailserver"
description: |-
  Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.
---

# fortimanager_systemp_system_emailserver
Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.

## Example Usage

```hcl
resource "fortimanager_systemp_system_emailserver" "trname" {
  devprof      = "default"
  authenticate = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `authenticate` - Enable/disable authentication. Valid values: `disable`, `enable`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `password` - SMTP server user password for authentication.
* `port` - SMTP server port.
* `reply_to` - Reply-To email address.
* `security` - Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.

* `server` - SMTP server IP address or hostname.
* `source_ip` - SMTP server IPv4 source IP.
* `source_ip6` - SMTP server IPv6 source IP.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`.

* `type` - Use FortiGuard Message service or custom email server. Valid values: `custom`.

* `username` - SMTP server user name for authentication.
* `validate_server` - Enable/disable validation of server certificate. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Systemp SystemEmailServer can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_system_emailserver.labelname SystempSystemEmailServer
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
