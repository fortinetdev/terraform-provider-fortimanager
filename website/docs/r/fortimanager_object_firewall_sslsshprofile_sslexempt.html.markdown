---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_sslsshprofile_sslexempt"
description: |-
  Servers to exempt from SSL inspection.
---

# fortimanager_object_firewall_sslsshprofile_sslexempt
Servers to exempt from SSL inspection.

~> This resource is a sub resource for variable `ssl_exempt` of resource `fortimanager_object_firewall_sslsshprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_sslsshprofile_sslexempt" "trname" {
  ssl_ssh_profile = fortimanager_object_firewall_sslsshprofile.trname.name
  address         = "FABRIC_DEVICE"
  fosid           = 4
  type            = "address"
  depends_on      = [fortimanager_object_firewall_sslsshprofile.trname]
}

resource "fortimanager_object_firewall_sslsshprofile" "trname" {
  name = "terr-sslsshprofile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `ssl_ssh_profile` - Ssl Ssh Profile.

* `address` - IPv4 address object.
* `address6` - IPv6 address object.
* `fortiguard_category` - FortiGuard category ID.
* `fosid` - ID number.
* `regex` - Exempt servers by regular expression.
* `type` - Type of address object (IPv4 or IPv6) or FortiGuard category. Valid values: `fortiguard-category`, `address`, `address6`, `wildcard-fqdn`, `regex`.

* `wildcard_fqdn` - Exempt servers by wildcard FQDN.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall SslSshProfileSslExempt can be imported using any of these accepted formats:
```
Set import_options = ["ssl_ssh_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_sslsshprofile_sslexempt.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
