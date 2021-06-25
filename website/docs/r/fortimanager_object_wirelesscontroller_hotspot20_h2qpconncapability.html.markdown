---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_h2qpconncapability"
description: |-
  Configure connection capability.
---

# fortimanager_object_wirelesscontroller_hotspot20_h2qpconncapability
Configure connection capability.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_h2qpconncapability" "trname" {
  esp_port      = "closed"
  ftp_port      = "closed"
  http_port     = "open"
  icmp_port     = "closed"
  ikev2_port    = "closed"
  ikev2_xx_port = "closed"
  name          = "terr-wictl-hot20-heqp-conn-capblity"
  pptp_vpn_port = "closed"
  ssh_port      = "open"
  tls_port      = "closed"
  voip_tcp_port = "closed"
  voip_udp_port = "unknown"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `esp_port` - Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.

* `ftp_port` - Set FTP port service status. Valid values: `closed`, `open`, `unknown`.

* `http_port` - Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.

* `icmp_port` - Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.

* `ikev2_port` - Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.

* `ikev2_xx_port` - Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.

* `name` - Connection capability name.
* `pptp_vpn_port` - Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.

* `ssh_port` - Set SSH port service status. Valid values: `closed`, `open`, `unknown`.

* `tls_port` - Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.

* `voip_tcp_port` - Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.

* `voip_udp_port` - Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20H2QpConnCapability can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_h2qpconncapability.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
