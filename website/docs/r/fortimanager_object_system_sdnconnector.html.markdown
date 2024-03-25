---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_sdnconnector"
description: |-
  Configure connection to SDN Connector.
---

# fortimanager_object_system_sdnconnector
Configure connection to SDN Connector.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `compartment_list`: `fortimanager_object_system_sdnconnector_compartmentlist`
>- `external_account_list`: `fortimanager_object_system_sdnconnector_externalaccountlist`
>- `external_ip`: `fortimanager_object_system_sdnconnector_externalip`
>- `forwarding_rule`: `fortimanager_object_system_sdnconnector_forwardingrule`
>- `gcp_project_list`: `fortimanager_object_system_sdnconnector_gcpprojectlist`
>- `nic`: `fortimanager_object_system_sdnconnector_nic`
>- `oci_region_list`: `fortimanager_object_system_sdnconnector_ociregionlist`
>- `route`: `fortimanager_object_system_sdnconnector_route`
>- `route_table`: `fortimanager_object_system_sdnconnector_routetable`



## Example Usage

```hcl
resource "fortimanager_object_system_sdnconnector" "trname" {
  api_key            = ["fortinet"]
  compute_generation = 2
  ibm_region         = "us-south"
  ibm_region_gen1    = "us-south"
  ibm_region_gen2    = "us-south"
  name               = "terr-system-sdn-connector"
  password           = ["fortinet"]
  rest_interface     = "mgmt"
  rest_password      = ["fortinet"]
  rest_sport         = 9443
  rest_ssl           = "enable"
  server             = "192.168.1.1"
  status             = "enable"
  type               = "aci"
  username           = "terraform-tefv"
  vcenter_password   = ["fortinet"]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_local_cert` - _Local_Cert.
* `access_key` - AWS / ACS access key ID.
* `alt_resource_ip` - Enable/disable AWS alternative resource IP. Valid values: `disable`, `enable`.

* `api_key` - IBM cloud API key or service ID API key.
* `azure_region` - Azure server region. Valid values: `global`, `china`, `germany`, `usgov`, `local`.

* `client_id` - Azure client ID (application ID).
* `client_secret` - Azure client secret (application key).
* `compartment_list` - Compartment-List. The structure of `compartment_list` block is documented below.
* `compartment_id` - Compartment ID.
* `compute_generation` - Compute generation for IBM cloud infrastructure.
* `domain` - Domain name.
* `external_account_list` - External-Account-List. The structure of `external_account_list` block is documented below.
* `external_ip` - External-Ip. The structure of `external_ip` block is documented below.
* `forwarding_rule` - Forwarding-Rule. The structure of `forwarding_rule` block is documented below.
* `gcp_project` - GCP project name.
* `gcp_project_list` - Gcp-Project-List. The structure of `gcp_project_list` block is documented below.
* `group_name` - Group name of computers.
* `ha_status` - Enable/disable use for FortiGate HA service. Valid values: `disable`, `enable`.

* `ibm_region` - IBM cloud region name. Valid values: `us-south`, `us-east`, `germany`, `great-britain`, `japan`, `australia`.

* `ibm_region_gen1` - Ibm-Region-Gen1. Valid values: `us-south`, `us-east`, `germany`, `great-britain`, `japan`, `australia`.

* `ibm_region_gen2` - Ibm-Region-Gen2. Valid values: `us-south`, `us-east`, `great-britain`.

* `key_passwd` - Private key password.
* `last_update` - Last-Update.
* `login_endpoint` - Azure Stack login endpoint.
* `name` - SDN connector name.
* `nic` - Nic. The structure of `nic` block is documented below.
* `nsx_cert_fingerprint` - NSX certificate fingerprint.
* `oci_cert` - OCI certificate.
* `oci_fingerprint` - Oci-Fingerprint.
* `oci_region_list` - Oci-Region-List. The structure of `oci_region_list` block is documented below.
* `oci_region` - OCI server region.
* `oci_region_type` - OCI region type. Valid values: `commercial`, `government`.

* `password` - Password of the remote SDN connector as login credentials.
* `private_key` - Private key of GCP service account.
* `proxy` - SDN proxy.
* `region` - AWS / ACS region name.
* `resource_group` - Azure resource group.
* `resource_url` - Azure Stack resource URL.
* `rest_interface` - Interface name for REST service to listen on. Valid values: `mgmt`, `sync`.

* `rest_password` - Password for REST service.
* `rest_sport` - REST service access port (1 - 65535).
* `rest_ssl` - Rest-Ssl. Valid values: `disable`, `enable`.

* `route` - Route. The structure of `route` block is documented below.
* `route_table` - Route-Table. The structure of `route_table` block is documented below.
* `secret_key` - AWS / ACS secret access key.
* `secret_token` - Secret token of Kubernetes service account.
* `server` - Server address of the remote SDN connector.
* `server_ca_cert` - Trust only those servers whose certificate is directly/indirectly signed by this certificate.
* `server_cert` - Trust servers that contain this certificate only.
* `server_list` - Server address list of the remote SDN connector.
* `server_port` - Port number of the remote SDN connector.
* `service_account` - GCP service account email.
* `status` - Enable/disable connection to the remote SDN connector. Valid values: `disable`, `enable`.

* `subscription_id` - Azure subscription ID.
* `tenant_id` - Tenant ID (directory ID).
* `type` - Type of SDN connector. Valid values: `aci`, `aws`, `nsx`, `nuage`, `azure`, `gcp`, `oci`, `openstack`, `kubernetes`, `vmware`, `acs`, `alicloud`, `sepm`, `aci-direct`, `ibm`.

* `update_interval` - Dynamic object update interval (30 - 3600 sec, default = 60, 0 = disabled).
* `updating` - Updating.
* `use_metadata_iam` - Enable/disable using IAM role from metadata to call API. Valid values: `disable`, `enable`.

* `user_id` - User ID.
* `username` - Username of the remote SDN connector as login credentials.
* `vcenter_password` - vCenter server password for NSX quarantine.
* `vcenter_server` - vCenter server address for NSX quarantine.
* `vcenter_username` - vCenter server username for NSX quarantine.
* `verify_certificate` - Enable/disable server certificate verification. Valid values: `disable`, `enable`.

* `vmx_image_url` - URL of web-hosted VMX image.
* `vmx_service_name` - VMX Service name.
* `vpc_id` - AWS VPC ID.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `compartment_list` block supports:

* `compartment_id` - OCI compartment ID.

The `external_account_list` block supports:

* `external_id` - AWS external ID.
* `region_list` - AWS region name list.
* `role_arn` - AWS role ARN to assume.

The `external_ip` block supports:

* `name` - External IP name.

The `forwarding_rule` block supports:

* `rule_name` - Forwarding rule name.
* `target` - Target instance name.

The `gcp_project_list` block supports:

* `gcp_zone_list` - Configure GCP zone list.
* `id` - GCP project ID.

The `nic` block supports:

* `ip` - Ip. The structure of `ip` block is documented below.
* `name` - Network interface name.

The `ip` block supports:

* `name` - IP configuration name.
* `public_ip` - Public IP name.
* `resource_group` - Resource group of Azure public IP.

The `oci_region_list` block supports:

* `region` - OCI region.

The `route` block supports:

* `name` - Route name.

The `route_table` block supports:

* `name` - Route table name.
* `resource_group` - Resource group of Azure route table.
* `route` - Route. The structure of `route` block is documented below.
* `subscription_id` - Subscription ID of Azure route table.

The `route` block supports:

* `name` - Route name.
* `next_hop` - Next hop address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem SdnConnector can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_sdnconnector.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
