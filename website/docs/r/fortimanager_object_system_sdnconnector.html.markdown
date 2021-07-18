---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_sdnconnector"
description: |-
  Configure connection to SDN Connector.
---

# fortimanager_object_system_sdnconnector
Configure connection to SDN Connector.

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
* `api_key` - IBM cloud API key or service ID API key.
* `azure_region` - Azure server region. Valid values: `global`, `china`, `germany`, `usgov`, `local`.

* `client_id` - Azure client ID (application ID).
* `client_secret` - Azure client secret (application key).
* `compartment_id` - Compartment ID.
* `compute_generation` - Compute generation for IBM cloud infrastructure.
* `domain` - Domain name.
* `external_ip` - External-Ip. The structure of `external_ip` block is documented below.
* `gcp_project` - GCP project name.
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
* `oci_region` - OCI server region.
* `oci_region_type` - OCI region type. Valid values: `commercial`, `government`.

* `password` - Password of the remote SDN connector as login credentials.
* `private_key` - Private key of GCP service account.
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
* `vmx_image_url` - URL of web-hosted VMX image.
* `vmx_service_name` - VMX Service name.
* `vpc_id` - AWS VPC ID.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `external_ip` block supports:

* `name` - External IP name.

The `nic` block supports:

* `ip` - Ip. The structure of `ip` block is documented below.
* `name` - Network interface name.

The `ip` block supports:

* `name` - IP configuration name.
* `public_ip` - Public IP name.
* `resource_group` - Resource group of Azure public IP.

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
