---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_fortiguard"
description: |-
  Configure FortiGuard services.
---

# fortimanager_object_system_fortiguard
Configure FortiGuard services.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `antispam_cache` - Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance. Valid values: `disable`, `enable`.

* `antispam_cache_mpercent` - Maximum percent of FortiGate memory the antispam cache is allowed to use (1 - 15%).
* `antispam_cache_ttl` - Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
* `antispam_expiration` - Antispam-Expiration.
* `antispam_force_off` - Enable/disable turning off the FortiGuard antispam service. Valid values: `disable`, `enable`.

* `antispam_license` - Antispam-License.
* `antispam_timeout` - Antispam query time out (1 - 30 sec, default = 7).
* `anycast_sdns_server_ip` - IP address of the FortiGuard anycast DNS rating server.
* `anycast_sdns_server_port` - Port to connect to on the FortiGuard anycast DNS rating server.
* `auto_join_forticloud` - Automatically connect to and login to FortiCloud. Valid values: `disable`, `enable`.

* `ddns_server_ip` - IP address of the FortiDDNS server.
* `ddns_server_port` - Port used to communicate with FortiDDNS servers.
* `fortiguard_anycast` - Enable/disable use of FortiGuard's anycast network. Valid values: `disable`, `enable`.

* `fortiguard_anycast_source` - Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet. Valid values: `fortinet`, `aws`, `debug`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `load_balance_servers` - Number of servers to alternate between as first FortiGuard option.
* `outbreak_prevention_cache` - Enable/disable FortiGuard Virus Outbreak Prevention cache. Valid values: `disable`, `enable`.

* `outbreak_prevention_cache_mpercent` - Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
* `outbreak_prevention_cache_ttl` - Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
* `outbreak_prevention_expiration` - Outbreak-Prevention-Expiration.
* `outbreak_prevention_force_off` - Turn off FortiGuard Virus Outbreak Prevention service. Valid values: `disable`, `enable`.

* `outbreak_prevention_license` - Outbreak-Prevention-License.
* `outbreak_prevention_timeout` - FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
* `port` - Port used to communicate with the FortiGuard servers. Valid values: `53`, `80`, `443`, `8888`.

* `protocol` - Protocol used to communicate with the FortiGuard servers. Valid values: `udp`, `http`, `https`.

* `proxy_password` - Proxy user password.
* `proxy_server_ip` - IP address of the proxy server.
* `proxy_server_port` - Port used to communicate with the proxy server.
* `proxy_username` - Proxy user name.
* `sandbox_region` - Cloud sandbox region.
* `sdns_options` - Customization options for the FortiGuard DNS service. Valid values: `include-question-section`.

* `sdns_server_ip` - IP address of the FortiGuard DNS rating server.
* `sdns_server_port` - Port to connect to on the FortiGuard DNS rating server.
* `service_account_id` - Service account ID.
* `source_ip` - Source IPv4 address used to communicate with FortiGuard.
* `source_ip6` - Source IPv6 address used to communicate with FortiGuard.
* `update_server_location` - Signature update server location. Valid values: `any`, `usa`.

* `webfilter_cache` - Enable/disable FortiGuard web filter caching. Valid values: `disable`, `enable`.

* `webfilter_cache_ttl` - Time-to-live for web filter cache entries in seconds (300 - 86400).
* `webfilter_expiration` - Webfilter-Expiration.
* `webfilter_force_off` - Enable/disable turning off the FortiGuard web filtering service. Valid values: `disable`, `enable`.

* `webfilter_license` - Webfilter-License.
* `webfilter_timeout` - Web filter query time out (1 - 30 sec, default = 7).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem Fortiguard can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_fortiguard.labelname ObjectSystemFortiguard
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
