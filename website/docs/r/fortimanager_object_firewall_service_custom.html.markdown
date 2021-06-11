---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_service_custom"
description: |-
  Configure custom services.
---

# fortimanager_object_firewall_service_custom
Configure custom services.

## Example Usage

```hcl
resource "fortimanager_object_firewall_service_custom" "trname" {
  app_service_type = "disable"
  color            = 1
  comment          = "comment"
  helper           = "auto"
  name             = "terraform-tefv"
  protocol         = "ALL"
  proxy            = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `app_category` - Application category ID.
* `app_service_type` - Application service type. Valid values: `disable`, `app-id`, `app-category`.

* `application` - Application ID.
* `category` - Service category.
* `check_reset_range` - Configure the type of ICMP error message verification. Valid values: `disable`, `default`, `strict`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `fqdn` - Fully qualified domain name.
* `global_object` - Global Object.
* `helper` - Helper name. Valid values: `disable`, `auto`, `ftp`, `tftp`, `ras`, `h323`, `tns`, `mms`, `sip`, `pptp`, `rtsp`, `dns-udp`, `dns-tcp`, `pmap`, `rsh`, `dcerpc`, `mgcp`, `gtp-c`, `gtp-u`, `gtp-b`.

* `icmpcode` - ICMP code.
* `icmptype` - ICMP type.
* `iprange` - Start and end of the IP range associated with service.
* `name` - Custom service name.
* `protocol` - Protocol type based on IANA numbers. Valid values: `ICMP`, `IP`, `TCP/UDP/SCTP`, `ICMP6`, `HTTP`, `FTP`, `CONNECT`, `SOCKS`, `ALL`, `SOCKS-TCP`, `SOCKS-UDP`.

* `protocol_number` - IP protocol number.
* `proxy` - Enable/disable web proxy service. Valid values: `disable`, `enable`.

* `sctp_portrange` - Multiple SCTP port ranges.
* `session_ttl` - Session TTL (300 - 2764800, 0 = default).
* `tcp_halfclose_timer` - Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
* `tcp_halfopen_timer` - Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
* `tcp_portrange` - Multiple TCP port ranges.
* `tcp_timewait_timer` - Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
* `udp_idle_timer` - UDP half close timeout (0 - 86400 sec, 0 = default).
* `udp_portrange` - Multiple UDP port ranges.
* `visibility` - Enable/disable the visibility of the service on the GUI. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ServiceCustom can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_service_custom.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
