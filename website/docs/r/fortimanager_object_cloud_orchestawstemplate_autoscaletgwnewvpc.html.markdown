---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_cloud_orchestawstemplate_autoscaletgwnewvpc"
description: |-
  ObjectCloud OrchestAwstemplateAutoscaleTgwNewVpc
---

# fortimanager_object_cloud_orchestawstemplate_autoscaletgwnewvpc
ObjectCloud OrchestAwstemplateAutoscaleTgwNewVpc

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `availability_zones` - Availability-Zones.
* `bgp_asn` - Bgp-Asn.
* `custom_asset_container` - Custom-Asset-Container.
* `custom_asset_directory` - Custom-Asset-Directory.
* `custom_identifier` - Custom-Identifier.
* `faz_autoscale_admin_password` - Faz-Autoscale-Admin-Password.
* `faz_autoscale_admin_username` - Faz-Autoscale-Admin-Username.
* `faz_custom_private_ipaddress` - Faz-Custom-Private-Ipaddress.
* `faz_instance_type` - Faz-Instance-Type. Valid values: `h1.2xlarge`, `h1.4xlarge`, `h1.8xlarge`, `m5.large`, `m5.xlarge`, `m5.2xlarge`, `m5.4xlarge`, `m5.12xlarge`, `t2.medium`, `t2.large`, `t2.xlarge`.

* `faz_integration_options` - Faz-Integration-Options. Valid values: `no`, `yes`.

* `faz_version` - Faz-Version.
* `fgt_admin_cidr` - Fgt-Admin-Cidr.
* `fgt_admin_port` - Fgt-Admin-Port.
* `fgt_instance_type` - Fgt-Instance-Type. Valid values: `t2.small`, `c5.large`, `c5.xlarge`, `c5.2xlarge`, `c5.4xlarge`, `c5.9xlarge`.

* `fgt_psk_secret` - Fgt-Psk-Secret.
* `fgtasg_cool_down` - Fgtasg-Cool-Down.
* `fgtasg_desired_capacity_byol` - Fgtasg-Desired-Capacity-Byol.
* `fgtasg_desired_capacity_payg` - Fgtasg-Desired-Capacity-Payg.
* `fgtasg_health_check_grace_period` - Fgtasg-Health-Check-Grace-Period.
* `fgtasg_max_size_byol` - Fgtasg-Max-Size-Byol.
* `fgtasg_max_size_payg` - Fgtasg-Max-Size-Payg.
* `fgtasg_min_size_byol` - Fgtasg-Min-Size-Byol.
* `fgtasg_min_size_payg` - Fgtasg-Min-Size-Payg.
* `fgtasg_scale_in_threshold` - Fgtasg-Scale-In-Threshold.
* `fgtasg_scale_out_threshold` - Fgtasg-Scale-Out-Threshold.
* `fos_version` - Fos-Version.
* `get_license_grace_period` - Get-License-Grace-Period.
* `heartbeat_delay_allowance` - Heartbeat-Delay-Allowance.
* `heartbeat_interval` - Heartbeat-Interval.
* `heartbeat_loss_count` - Heartbeat-Loss-Count.
* `key_pair_name` - Key-Pair-Name.
* `lifecycle_hook_timeout` - Lifecycle-Hook-Timeout.
* `name` - Name.
* `notification_email` - Notification-Email.
* `primary_election_timeout` - Primary-Election-Timeout.
* `public_subnet1_cidr` - Public-Subnet1-Cidr.
* `public_subnet2_cidr` - Public-Subnet2-Cidr.
* `resource_tag_prefix` - Resource-Tag-Prefix.
* `s3_bucket_name` - S3-Bucket-Name.
* `s3_key_prefix` - S3-Key-Prefix.
* `sync_recovery_count` - Sync-Recovery-Count.
* `terminate_unhealthy_vm` - Terminate-Unhealthy-Vm. Valid values: `no`, `yes`.

* `transit_gateway_id` - Transit-Gateway-Id.
* `transit_gateway_support_options` - Transit-Gateway-Support-Options. Valid values: `create one`, `use an existing one`.

* `use_custom_asset_location` - Use-Custom-Asset-Location. Valid values: `no`, `yes`.

* `vpc_cidr` - Vpc-Cidr.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCloud OrchestAwstemplateAutoscaleTgwNewVpc can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_cloud_orchestawstemplate_autoscaletgwnewvpc.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
