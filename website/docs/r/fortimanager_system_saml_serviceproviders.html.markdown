---
subcategory: "System Saml"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_saml_serviceproviders"
description: |-
  Authorized service providers.
---

# fortimanager_system_saml_serviceproviders
Authorized service providers.

## Argument Reference


The following arguments are supported:


* `idp_entity_id` - IDP Entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `name` - Name.
* `prefix` - Prefix.
* `sp_adom` - SP adom name.
* `sp_cert` - SP certificate name.
* `sp_entity_id` - SP Entity ID.
* `sp_profile` - SP profile name.
* `sp_single_logout_url` - SP single logout URL.
* `sp_single_sign_on_url` - SP single sign-on URL.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{idp_entity_id}}.

## Import

System SamlServiceProviders can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_saml_serviceproviders.labelname {{idp_entity_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

