---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_sdnconnector_forwardingrule"
description: |-
  Configure GCP forwarding rule.
---

# fortimanager_object_system_sdnconnector_forwardingrule
Configure GCP forwarding rule.

~> This resource is a sub resource for variable `forwarding_rule` of resource `fortimanager_object_system_sdnconnector`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_sdnconnector_forwardingrule" "trname" {
  sdn_connector = fortimanager_object_system_sdnconnector.trname.name
  rule_name     = "terr-forwardingrule"
  depends_on    = [fortimanager_object_system_sdnconnector.trname]
}

resource "fortimanager_object_system_sdnconnector" "trname" {
  access_key = "key"
  region     = "local"
  name       = "terr-sdnconnector"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `sdn_connector` - Sdn Connector.

* `rule_name` - Forwarding rule name.
* `target` - Target instance name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{rule_name}}.

## Import

ObjectSystem SdnConnectorForwardingRule can be imported using any of these accepted formats:
```
Set import_options = ["sdn_connector=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_sdnconnector_forwardingrule.labelname {{rule_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
