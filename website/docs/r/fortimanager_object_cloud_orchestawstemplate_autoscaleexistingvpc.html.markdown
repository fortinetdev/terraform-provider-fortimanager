---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_cloud_orchestawstemplate_autoscaleexistingvpc"
description: |-
  ObjectCloud OrchestAwstemplateAutoscaleExistingVpc
---

# fortimanager_object_cloud_orchestawstemplate_autoscaleexistingvpc
ObjectCloud OrchestAwstemplateAutoscaleExistingVpc

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

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
* `internal_balancer_dns_name` - Internal-Balancer-Dns-Name.
* `internal_balancing_options` - Internal-Balancing-Options. Valid values: `add a new internal load balancer`, `use a load balancer specified below`, `do not need one`.

* `internal_target_group_health_check_path` - Internal-Target-Group-Health-Check-Path.
* `key_pair_name` - Key-Pair-Name.
* `lifecycle_hook_timeout` - Lifecycle-Hook-Timeout.
* `loadbalancing_health_check_threshold` - Loadbalancing-Health-Check-Threshold.
* `loadbalancing_traffic_port` - Loadbalancing-Traffic-Port.
* `loadbalancing_traffic_protocol` - Loadbalancing-Traffic-Protocol. Valid values: `HTTPS`, `HTTP`, `TCP`.

* `name` - Name.
* `notification_email` - Notification-Email.
* `primary_election_timeout` - Primary-Election-Timeout.
* `private_subnet_route_table` - Private-Subnet-Route-Table.
* `private_subnet1` - Private-Subnet1.
* `private_subnet2` - Private-Subnet2.
* `public_subnet1` - Public-Subnet1.
* `public_subnet2` - Public-Subnet2.
* `resource_tag_prefix` - Resource-Tag-Prefix.
* `s3_bucket_name` - S3-Bucket-Name.
* `s3_key_prefix` - S3-Key-Prefix.
* `sync_recovery_count` - Sync-Recovery-Count.
* `terminate_unhealthy_vm` - Terminate-Unhealthy-Vm. Valid values: `no`, `yes`.

* `use_custom_asset_location` - Use-Custom-Asset-Location. Valid values: `no`, `yes`.

* `vpc_cidr` - Vpc-Cidr.
* `vpc_endpoint_id` - Vpc-Endpoint-Id.
* `vpc_id` - Vpc-Id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCloud OrchestAwstemplateAutoscaleExistingVpc can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_cloud_orchestawstemplate_autoscaleexistingvpc.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
