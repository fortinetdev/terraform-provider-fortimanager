---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_wtpprofile_eslsesdongle"
description: |-
  ESL SES-imagotag dongle configuration.
---

# fortimanager_object_wirelesscontroller_wtpprofile_eslsesdongle
ESL SES-imagotag dongle configuration.

~> This resource is a sub resource for variable `esl_ses_dongle` of resource `fortimanager_object_wirelesscontroller_wtpprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_wtpprofile_eslsesdongle" "trname" {
  apc_addr_type = "ip"
  apc_ip        = "10.160.88.123"
  apc_port      = 67
  wtp_profile   = fortimanager_object_wirelesscontroller_wtpprofile.trname2.name
  depends_on    = [fortimanager_object_wirelesscontroller_wtpprofile.trname2]
}

resource "fortimanager_object_wirelesscontroller_wtpprofile" "trname2" {
  name = "terr-wtpprofile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wtp_profile` - Wtp Profile.

* `apc_addr_type` - ESL SES-imagotag APC address type (default = fqdn). Valid values: `fqdn`, `ip`.

* `apc_fqdn` - FQDN of ESL SES-imagotag Access Point Controller (APC).
* `apc_ip` - IP address of ESL SES-imagotag Access Point Controller (APC).
* `apc_port` - Port of ESL SES-imagotag Access Point Controller (APC).
* `coex_level` - ESL SES-imagotag dongle coexistence level (default = none). Valid values: `none`.

* `compliance_level` - Compliance levels for the ESL solution integration (default = compliance-level-2). Valid values: `compliance-level-2`.

* `esl_channel` - ESL SES-imagotag dongle channel (default = 127). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `127`, `-1`.

* `output_power` - ESL SES-imagotag dongle output power (default = A). Valid values: `a`, `b`, `c`, `d`, `e`, `f`, `g`, `h`.

* `scd_enable` - Enable/disable ESL SES-imagotag Serial Communication Daemon (SCD) (default = disable). Valid values: `disable`, `enable`.

* `tls_cert_verification` - Enable/disable TLS Certificate verification. (default = enable). Valid values: `disable`, `enable`.

* `tls_fqdn_verification` - Enable/disable TLS Certificate verification. (default = disable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWirelessController WtpProfileEslSesDongle can be imported using any of these accepted formats:
```
Set import_options = ["wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_wtpprofile_eslsesdongle.labelname ObjectWirelessControllerWtpProfileEslSesDongle
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
