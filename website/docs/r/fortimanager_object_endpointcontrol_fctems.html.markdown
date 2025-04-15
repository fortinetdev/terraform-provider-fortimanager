---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_endpointcontrol_fctems"
description: |-
  Configure FortiClient Enterprise Management Server (EMS) entries.
---

# fortimanager_object_endpointcontrol_fctems
Configure FortiClient Enterprise Management Server (EMS) entries.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ca_cn_info` - Ca-Cn-Info.
* `admin_password` - FortiClient EMS admin password.
* `admin_username` - FortiClient EMS admin username.
* `call_timeout` - FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
* `certificate` - FortiClient EMS certificate.
* `capabilities` - List of EMS capabilities. Valid values: `fabric-auth`, `silent-approval`, `websocket`, `websocket-malware`, `push-ca-certs`.

* `certificate_fingerprint` - EMS certificate fingerprint.
* `cloud_authentication_access_key` - FortiClient EMS Cloud multitenancy access key
* `cloud_server_type` - Cloud server type. Valid values: `production`, `alpha`, `beta`.

* `dirty_reason` - Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.

* `ems_id` - EMS ID in order (1 - 5)
* `fortinetone_cloud_authentication` - Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `disable`, `enable`.

* `https_port` - FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `name` - FortiClient Enterprise Management Server (EMS) name.
* `out_of_sync_threshold` - Outdated resource threshold in seconds (10 - 3600, default = 180).
* `serial_number` - FortiClient EMS Serial Number.
* `preserve_ssl_session` - Enable/disable preservation of EMS SSL session connection. WARNING: Most users should not touch this setting! Valid values: `disable`, `enable`.

* `pull_avatars` - Enable/disable pulling avatars from EMS. Valid values: `disable`, `enable`.

* `pull_malware_hash` - Enable/disable pulling FortiClient malware hash from EMS. Valid values: `disable`, `enable`.

* `pull_sysinfo` - Enable/disable pulling SysInfo from EMS. Valid values: `disable`, `enable`.

* `pull_tags` - Enable/disable pulling FortiClient user tags from EMS. Valid values: `disable`, `enable`.

* `pull_vulnerabilities` - Enable/disable pulling vulnerabilities from EMS. Valid values: `disable`, `enable`.

* `send_tags_to_all_vdoms` - Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `disable`, `enable`.

* `server` - FortiClient EMS FQDN or IPv4 address.
* `source_ip` - REST API call source IP.
* `status` - Enable or disable this EMS configuration. Valid values: `disable`, `enable`.

* `trust_ca_cn` - Trust-Ca-Cn. Valid values: `disable`, `enable`.

* `verified_cn` - EMS certificate CN.
* `verifying_ca` - Lowest CA cert on Fortigate in verified EMS cert chain.
* `tenant_id` - EMS Tenant ID.
* `status_check_interval` - FortiClient EMS call timeout in seconds (1 - 120 seconds, default = 5).
* `websocket_override` - Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectEndpointControl Fctems can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_endpointcontrol_fctems.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
