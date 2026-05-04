---
subcategory: "Systemp"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_system_dns"
description: |-
  Configure DNS.
---

# fortimanager_systemp_system_dns
Configure DNS.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `alt_primary` - Alternate primary DNS server. This is not used as a failover DNS server.
* `alt_secondary` - Alternate secondary DNS server. This is not used as a failover DNS server.
* `cache_notfound_responses` - Enable/disable response from the DNS server when a record is not in cache. Valid values: `disable`, `enable`.

* `dns_cache_limit` - Maximum number of records in the DNS cache.
* `dns_cache_ttl` - Duration in seconds that the DNS cache retains information.
* `dns_over_tls` - Enable/disable/enforce DNS over TLS. Valid values: `disable`, `enable`, `enforce`.

* `domain` - <i>Support meta variable</i> Search suffix list for hostname lookup.
* `fqdn_cache_ttl` - FQDN cache time to live in seconds (0 - 86400, default = 0).
* `fqdn_max_refresh` - FQDN cache maximum refresh time in seconds (3600 - 86400, default = 3600).
* `fqdn_min_refresh` - FQDN cache minimum refresh time in seconds (10 - 3600, default = 60).
* `hostname_limit` - Limit of the number of hostname table entries (0 - 50000).
* `hostname_ttl` - TTL of hostname table entries (60 - 86400).
* `interface` - <i>Support meta variable</i> Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ip6_primary` - Primary DNS server IPv6 address.
* `ip6_secondary` - Secondary DNS server IPv6 address.
* `log` - Local DNS log setting. Valid values: `disable`, `error`, `all`.

* `primary` - <i>Support meta variable</i> Primary DNS server IP address.
* `protocol` - DNS transport protocols. Valid values: `cleartext`, `dot`, `doh`.

* `retry` - Number of times to retry (0 - 5).
* `root_servers` - Configure up to two preferred servers that serve the DNS root zone (default uses all 13 root servers).
* `secondary` - <i>Support meta variable</i> Secondary DNS server IP address.
* `server_hostname` - DNS server host name list.
* `server_select_method` - Specify how configured servers are prioritized. Valid values: `least-rtt`, `failover`.

* `source_ip` - IP address used by the DNS server as its source IP.
* `source_ip_interface` - IP address of the specified interface as the source IP address.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `timeout` - DNS query timeout interval in seconds (1 - 10).
* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Systemp SystemDns can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_system_dns.labelname SystempSystemDns
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
