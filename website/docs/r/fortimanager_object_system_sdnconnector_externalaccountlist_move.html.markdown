---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_sdnconnector_externalaccountlist_move"
description: |-
  Move AWS external account list.
---

# fortimanager_object_system_sdnconnector_externalaccountlist_move
Move AWS external account list.

## Example Usage

```hcl
resource "fortimanager_object_system_sdnconnector_externalaccountlist_move" "trname" {
  sdn_connector         = fortimanager_object_system_sdnconnector.trname.name
  external_account_list = fortimanager_object_system_sdnconnector_externalaccountlist.trname3.external_id
  target                = fortimanager_object_system_sdnconnector_externalaccountlist.trname2.external_id
  option                = "after"
  depends_on            = [fortimanager_object_system_sdnconnector_externalaccountlist.trname3, fortimanager_object_system_sdnconnector_externalaccountlist.trname2]
}

resource "fortimanager_object_system_sdnconnector_externalaccountlist" "trname3" {
  sdn_connector = fortimanager_object_system_sdnconnector.trname.name
  external_id   = 13
  role_arn      = 13
  depends_on    = [fortimanager_object_system_sdnconnector.trname]
}

resource "fortimanager_object_system_sdnconnector_externalaccountlist" "trname2" {
  sdn_connector = fortimanager_object_system_sdnconnector.trname.name
  external_id   = 14
  role_arn      = 14
  depends_on    = [fortimanager_object_system_sdnconnector.trname]
}

resource "fortimanager_object_system_sdnconnector" "trname" {
  access_key = "key"
  region     = "remote"
  name       = "terr-sdnconnector"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `sdn_connector` - Sdn Connector.
* `external_account_list` - External Account List.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{role_arn}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the external account list changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of external account lists.
