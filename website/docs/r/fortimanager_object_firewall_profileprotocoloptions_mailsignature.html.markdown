---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_profileprotocoloptions_mailsignature"
description: |-
  Configure Mail signature.
---

# fortimanager_object_firewall_profileprotocoloptions_mailsignature
Configure Mail signature.

~> This resource is a sub resource for variable `mail_signature` of resource `fortimanager_object_firewall_profileprotocoloptions`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_profileprotocoloptions_mailsignature" "trname" {
  profile_protocol_options = fortimanager_object_firewall_profileprotocoloptions.trname.name
  status                   = "disable"
  depends_on               = [fortimanager_object_firewall_profileprotocoloptions.trname]
}

resource "fortimanager_object_firewall_profileprotocoloptions" "trname" {
  name = "terr-profileprotocoloptions"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile_protocol_options` - Profile Protocol Options.

* `signature` - Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).
* `status` - Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall ProfileProtocolOptionsMailSignature can be imported using any of these accepted formats:
```
Set import_options = ["profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_profileprotocoloptions_mailsignature.labelname ObjectFirewallProfileProtocolOptionsMailSignature
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
