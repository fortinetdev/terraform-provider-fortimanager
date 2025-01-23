---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_csf_trustedlist"
description: |-
  Pre-authorized and blocked security fabric nodes.
---

# fortimanager_system_csf_trustedlist
Pre-authorized and blocked security fabric nodes.

~> This resource is a sub resource for variable `trusted_list` of resource `fortimanager_system_csf`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:


* `action` - Security fabric authorization action. accept - Accept authorization request. deny - Deny authorization request. Valid values: `accept`, `deny`.

* `authorization_type` - Authorization type. serial - Verify downstream by serial number. certificate - Verify downstream by certificate. Valid values: `serial`, `certificate`.

* `certificate` - Certificate.
* `downstream_authorization` - Trust authorizations by this node&apos;s administrator. disable - Disable downstream authorization. enable - Enable downstream authorization. Valid values: `disable`, `enable`.

* `ha_members` - HA members.
* `index` - Index of the downstream in tree.
* `name` - Name.
* `serial` - Serial.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CsfTrustedList can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_csf_trustedlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

