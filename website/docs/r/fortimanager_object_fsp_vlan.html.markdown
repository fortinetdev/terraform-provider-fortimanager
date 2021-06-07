---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fsp_vlan"
description: |-
  ObjectFsp Vlan
---

# fortimanager_object_fsp_vlan
ObjectFsp Vlan

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_dhcp_status` - _Dhcp-Status. Valid values: `disable`, `enable`.

* `color` - Color.
* `dhcp_server` - Dhcp-Server. The structure of `dhcp_server` block is documented below.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `interface` - Interface. The structure of `interface` block is documented below.
* `name` - Name.
* `vdom` - Vdom.
* `vlanid` - Vlanid.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dhcp_server` block supports:

* `auto_configuration` - Enable/disable auto configuration. Valid values: `disable`, `enable`.

* `auto_managed_status` - Enable/disable use of this DHCP server once this interface has been assigned an IP address from FortiIPAM. Valid values: `disable`, `enable`.

* `conflicted_ip_timeout` - Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.
* `ddns_auth` - DDNS authentication mode. Valid values: `disable`, `tsig`.

* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_keyname` - DDNS update key name.
* `ddns_server_ip` - DDNS server IP.
* `ddns_ttl` - TTL.
* `ddns_update` - Enable/disable DDNS update for DHCP. Valid values: `disable`, `enable`.

* `ddns_update_override` - Enable/disable DDNS update override for DHCP. Valid values: `disable`, `enable`.

* `ddns_zone` - Zone of your domain name (ex. DDNS.com).
* `default_gateway` - Default gateway IP address assigned by the DHCP server.
* `dhcp_settings_from_fortiipam` - Enable/disable populating of DHCP server settings from FortiIPAM. Valid values: `disable`, `enable`.

* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server4` - DNS server 4.
* `dns_service` - Options for assigning DNS servers to DHCP clients. Valid values: `default`, `specify`, `local`.

* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `enable` - Enable. Valid values: `disable`, `enable`.

* `exclude_range` - Exclude-Range. The structure of `exclude_range` block is documented below.
* `filename` - Name of the boot file on the TFTP server.
* `forticlient_on_net_status` - Enable/disable FortiClient-On-Net service for this DHCP server. Valid values: `disable`, `enable`.

* `id` - ID.
* `ip_mode` - Method used to assign client IP. Valid values: `range`, `usrgrp`.

* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `ipsec_lease_hold` - DHCP over IPsec leases expire this many seconds after tunnel down (0 to disable forced-expiry).
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `mac_acl_default_action` - MAC access control default action (allow or block assigning IP settings). Valid values: `assign`, `block`.

* `netmask` - Netmask assigned by the DHCP server.
* `next_server` - IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.
* `ntp_server1` - NTP server 1.
* `ntp_server2` - NTP server 2.
* `ntp_server3` - NTP server 3.
* `ntp_service` - Options for assigning Network Time Protocol (NTP) servers to DHCP clients. Valid values: `default`, `specify`, `local`.

* `option1` - Option1.
* `option2` - Option2.
* `option3` - Option3.
* `option4` - Option4.
* `option5` - Option5.
* `option6` - Option6.
* `options` - Options. The structure of `options` block is documented below.
* `reserved_address` - Reserved-Address. The structure of `reserved_address` block is documented below.
* `server_type` - DHCP server can be a normal DHCP server or an IPsec DHCP server. Valid values: `regular`, `ipsec`.

* `status` - Enable/disable this DHCP configuration. Valid values: `disable`, `enable`.

* `tftp_server` - One or more hostnames or IP addresses of the TFTP servers in quotes separated by spaces.
* `timezone` - Select the time zone to be assigned to DHCP clients. Valid values: `00`, `01`, `02`, `03`, `04`, `05`, `06`, `07`, `08`, `09`, `10`, `11`, `12`, `13`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `22`, `23`, `24`, `25`, `26`, `27`, `28`, `29`, `30`, `31`, `32`, `33`, `34`, `35`, `36`, `37`, `38`, `39`, `40`, `41`, `42`, `43`, `44`, `45`, `46`, `47`, `48`, `49`, `50`, `51`, `52`, `53`, `54`, `55`, `56`, `57`, `58`, `59`, `60`, `61`, `62`, `63`, `64`, `65`, `66`, `67`, `68`, `69`, `70`, `71`, `72`, `73`, `74`, `75`, `76`, `77`, `78`, `79`, `80`, `81`, `82`, `83`, `84`, `85`, `86`, `87`.

* `timezone_option` - Options for the DHCP server to set the client's time zone. Valid values: `disable`, `default`, `specify`.

* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.
* `wifi_ac_service` - Options for assigning WiFi Access Controllers to DHCP clients Valid values: `specify`, `local`.

* `wifi_ac1` - WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).
* `wifi_ac2` - WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).
* `wifi_ac3` - WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).
* `wins_server1` - WINS server 1.
* `wins_server2` - WINS server 2.

The `exclude_range` block supports:

* `end_ip` - End of IP range.
* `id` - ID.
* `start_ip` - Start of IP range.

The `ip_range` block supports:

* `end_ip` - End of IP range.
* `id` - ID.
* `start_ip` - Start of IP range.

The `options` block supports:

* `code` - DHCP option code.
* `id` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.

* `value` - DHCP option value.

The `reserved_address` block supports:

* `action` - Options for the DHCP server to configure the client with the reserved MAC address. Valid values: `assign`, `block`, `reserved`.

* `circuit_id` - Option 82 circuit-ID of the client that will get the reserved IP address.
* `circuit_id_type` - DHCP option type. Valid values: `hex`, `string`.

* `description` - Description.
* `id` - ID.
* `ip` - IP address to be reserved for the MAC address.
* `mac` - MAC address of the client that will get the reserved IP address.
* `remote_id` - Option 82 remote-ID of the client that will get the reserved IP address.
* `remote_id_type` - DHCP option type. Valid values: `hex`, `string`.

* `type` - DHCP reserved-address type. Valid values: `mac`, `option82`.


The `dynamic_mapping` block supports:

* `_dhcp_status` - _Dhcp-Status. Valid values: `disable`, `enable`.

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `dhcp_server` - Dhcp-Server. The structure of `dhcp_server` block is documented below.
* `interface` - Interface. The structure of `interface` block is documented below.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `dhcp_server` block supports:

* `auto_configuration` - Enable/disable auto configuration. Valid values: `disable`, `enable`.

* `auto_managed_status` - Enable/disable use of this DHCP server once this interface has been assigned an IP address from FortiIPAM. Valid values: `disable`, `enable`.

* `conflicted_ip_timeout` - Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.
* `ddns_auth` - DDNS authentication mode. Valid values: `disable`, `tsig`.

* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_keyname` - DDNS update key name.
* `ddns_server_ip` - DDNS server IP.
* `ddns_ttl` - TTL.
* `ddns_update` - Enable/disable DDNS update for DHCP. Valid values: `disable`, `enable`.

* `ddns_update_override` - Enable/disable DDNS update override for DHCP. Valid values: `disable`, `enable`.

* `ddns_zone` - Zone of your domain name (ex. DDNS.com).
* `default_gateway` - Default gateway IP address assigned by the DHCP server.
* `dhcp_settings_from_fortiipam` - Enable/disable populating of DHCP server settings from FortiIPAM. Valid values: `disable`, `enable`.

* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server4` - DNS server 4.
* `dns_service` - Options for assigning DNS servers to DHCP clients. Valid values: `default`, `specify`, `local`.

* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `enable` - Enable. Valid values: `disable`, `enable`.

* `exclude_range` - Exclude-Range. The structure of `exclude_range` block is documented below.
* `filename` - Name of the boot file on the TFTP server.
* `forticlient_on_net_status` - Enable/disable FortiClient-On-Net service for this DHCP server. Valid values: `disable`, `enable`.

* `id` - ID.
* `ip_mode` - Method used to assign client IP. Valid values: `range`, `usrgrp`.

* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `ipsec_lease_hold` - DHCP over IPsec leases expire this many seconds after tunnel down (0 to disable forced-expiry).
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `mac_acl_default_action` - MAC access control default action (allow or block assigning IP settings). Valid values: `assign`, `block`.

* `netmask` - Netmask assigned by the DHCP server.
* `next_server` - IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.
* `ntp_server1` - NTP server 1.
* `ntp_server2` - NTP server 2.
* `ntp_server3` - NTP server 3.
* `ntp_service` - Options for assigning Network Time Protocol (NTP) servers to DHCP clients. Valid values: `default`, `specify`, `local`.

* `option1` - Option1.
* `option2` - Option2.
* `option3` - Option3.
* `option4` - Option4.
* `option5` - Option5.
* `option6` - Option6.
* `options` - Options. The structure of `options` block is documented below.
* `reserved_address` - Reserved-Address. The structure of `reserved_address` block is documented below.
* `server_type` - DHCP server can be a normal DHCP server or an IPsec DHCP server. Valid values: `regular`, `ipsec`.

* `status` - Enable/disable this DHCP configuration. Valid values: `disable`, `enable`.

* `tftp_server` - One or more hostnames or IP addresses of the TFTP servers in quotes separated by spaces.
* `timezone` - Select the time zone to be assigned to DHCP clients. Valid values: `00`, `01`, `02`, `03`, `04`, `05`, `06`, `07`, `08`, `09`, `10`, `11`, `12`, `13`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `22`, `23`, `24`, `25`, `26`, `27`, `28`, `29`, `30`, `31`, `32`, `33`, `34`, `35`, `36`, `37`, `38`, `39`, `40`, `41`, `42`, `43`, `44`, `45`, `46`, `47`, `48`, `49`, `50`, `51`, `52`, `53`, `54`, `55`, `56`, `57`, `58`, `59`, `60`, `61`, `62`, `63`, `64`, `65`, `66`, `67`, `68`, `69`, `70`, `71`, `72`, `73`, `74`, `75`, `76`, `77`, `78`, `79`, `80`, `81`, `82`, `83`, `84`, `85`, `86`, `87`.

* `timezone_option` - Options for the DHCP server to set the client's time zone. Valid values: `disable`, `default`, `specify`.

* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.
* `wifi_ac_service` - Options for assigning WiFi Access Controllers to DHCP clients Valid values: `specify`, `local`.

* `wifi_ac1` - WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).
* `wifi_ac2` - WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).
* `wifi_ac3` - WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).
* `wins_server1` - WINS server 1.
* `wins_server2` - WINS server 2.

The `exclude_range` block supports:

* `end_ip` - End of IP range.
* `id` - ID.
* `start_ip` - Start of IP range.

The `ip_range` block supports:

* `end_ip` - End of IP range.
* `id` - ID.
* `start_ip` - Start of IP range.

The `options` block supports:

* `code` - DHCP option code.
* `id` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.

* `value` - DHCP option value.

The `reserved_address` block supports:

* `action` - Options for the DHCP server to configure the client with the reserved MAC address. Valid values: `assign`, `block`, `reserved`.

* `circuit_id` - Option 82 circuit-ID of the client that will get the reserved IP address.
* `circuit_id_type` - DHCP option type. Valid values: `hex`, `string`.

* `description` - Description.
* `id` - ID.
* `ip` - IP address to be reserved for the MAC address.
* `mac` - MAC address of the client that will get the reserved IP address.
* `remote_id` - Option 82 remote-ID of the client that will get the reserved IP address.
* `remote_id_type` - DHCP option type. Valid values: `hex`, `string`.

* `type` - DHCP reserved-address type. Valid values: `mac`, `option82`.


The `interface` block supports:

* `dhcp_relay_agent_option` - Dhcp-Relay-Agent-Option. Valid values: `disable`, `enable`.

* `dhcp_relay_ip` - Dhcp-Relay-Ip.
* `dhcp_relay_service` - Dhcp-Relay-Service. Valid values: `disable`, `enable`.

* `dhcp_relay_type` - Dhcp-Relay-Type. Valid values: `regular`, `ipsec`.

* `ip` - Ip.
* `ipv6` - Ipv6. The structure of `ipv6` block is documented below.
* `secondary_IP` - Secondary-Ip. Valid values: `disable`, `enable`.

* `secondaryip` - Secondaryip. The structure of `secondaryip` block is documented below.
* `vlanid` - Vlanid.

The `ipv6` block supports:

* `autoconf` - Enable/disable address auto config. Valid values: `disable`, `enable`.

* `cli_conn6_status` - Cli-Conn6-Status.
* `dhcp6_client_options` - Dhcp6-Client-Options. Valid values: `rapid`, `iapd`, `iana`, `dns`, `dnsname`.

* `dhcp6_information_request` - Enable/disable DHCPv6 information request. Valid values: `disable`, `enable`.

* `dhcp6_prefix_delegation` - Enable/disable DHCPv6 prefix delegation. Valid values: `disable`, `enable`.

* `dhcp6_prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `dhcp6_prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `dhcp6_prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).
* `dhcp6_relay_ip` - DHCPv6 relay IP address.
* `dhcp6_relay_service` - Enable/disable DHCPv6 relay. Valid values: `disable`, `enable`.

* `dhcp6_relay_type` - DHCPv6 relay type. Valid values: `regular`.

* `icmp6_send_redirect` - Icmp6-Send-Redirect. Valid values: `disable`, `enable`.

* `interface_identifier` - IPv6 interface identifier.
* `ip6_address` - Primary IPv6 address prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
* `ip6_allowaccess` - Allow management access to the interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `capwap`, `fabric`.

* `ip6_default_life` - Default life (sec).
* `ip6_delegated_prefix_list` - Ip6-Delegated-Prefix-List. The structure of `ip6_delegated_prefix_list` block is documented below.
* `ip6_dns_server_override` - Enable/disable using the DNS server acquired by DHCP. Valid values: `disable`, `enable`.

* `ip6_extra_addr` - Ip6-Extra-Addr. The structure of `ip6_extra_addr` block is documented below.
* `ip6_hop_limit` - Hop limit (0 means unspecified).
* `ip6_link_mtu` - IPv6 link MTU.
* `ip6_manage_flag` - Enable/disable the managed flag. Valid values: `disable`, `enable`.

* `ip6_max_interval` - IPv6 maximum interval (4 to 1800 sec).
* `ip6_min_interval` - IPv6 minimum interval (3 to 1350 sec).
* `ip6_mode` - Addressing mode (static, DHCP, delegated). Valid values: `static`, `dhcp`, `pppoe`, `delegated`.

* `ip6_other_flag` - Enable/disable the other IPv6 flag. Valid values: `disable`, `enable`.

* `ip6_prefix_list` - Ip6-Prefix-List. The structure of `ip6_prefix_list` block is documented below.
* `ip6_prefix_mode` - Ip6-Prefix-Mode. Valid values: `dhcp6`, `ra`.

* `ip6_reachable_time` - IPv6 reachable time (milliseconds; 0 means unspecified).
* `ip6_retrans_time` - IPv6 retransmit time (milliseconds; 0 means unspecified).
* `ip6_send_adv` - Enable/disable sending advertisements about the interface. Valid values: `disable`, `enable`.

* `ip6_subnet` - Subnet to routing prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
* `ip6_upstream_interface` - Interface name providing delegated information.
* `nd_cert` - Neighbor discovery certificate.
* `nd_cga_modifier` - Neighbor discovery CGA modifier.
* `nd_mode` - Neighbor discovery mode. Valid values: `basic`, `SEND-compatible`.

* `nd_security_level` - Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).
* `nd_timestamp_delta` - Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).
* `nd_timestamp_fuzz` - Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).
* `unique_autoconf_addr` - Enable/disable unique auto config address. Valid values: `disable`, `enable`.

* `vrip6_link_local` - Link-local IPv6 address of virtual router.
* `vrrp_virtual_mac6` - Enable/disable virtual MAC for VRRP. Valid values: `disable`, `enable`.

* `vrrp6` - Vrrp6. The structure of `vrrp6` block is documented below.

The `ip6_delegated_prefix_list` block supports:

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `disable`, `enable`.

* `onlink_flag` - Enable/disable the onlink flag. Valid values: `disable`, `enable`.

* `prefix_id` - Prefix ID.
* `rdnss` - Recursive DNS server option.
* `rdnss_service` - Recursive DNS service option. Valid values: `delegated`, `default`, `specify`.

* `subnet` - Add subnet ID to routing prefix.
* `upstream_interface` - Name of the interface that provides delegated information.

The `ip6_extra_addr` block supports:

* `prefix` - IPv6 address prefix.

The `ip6_prefix_list` block supports:

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `disable`, `enable`.

* `dnssl` - DNS search list option.
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `disable`, `enable`.

* `preferred_life_time` - Preferred life time (sec).
* `prefix` - IPv6 prefix.
* `rdnss` - Recursive DNS server option.
* `valid_life_time` - Valid life time (sec).

The `vrrp6` block supports:

* `accept_mode` - Enable/disable accept mode. Valid values: `disable`, `enable`.

* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `preempt` - Enable/disable preempt mode. Valid values: `disable`, `enable`.

* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable VRRP. Valid values: `disable`, `enable`.

* `vrdst6` - Monitor the route to this destination.
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip6` - IPv6 address of the virtual router.

The `secondaryip` block supports:

* `allowaccess` - Management access settings for the secondary IP address. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `auto-ipsec`, `radius-acct`, `probe-response`, `capwap`, `dnp`, `ftm`, `fabric`.

* `detectprotocol` - Protocols used to detect the server. Valid values: `ping`, `tcp-echo`, `udp-echo`.

* `detectserver` - Gateway's ping server for this IP.
* `gwdetect` - Enable/disable detect gateway alive for first. Valid values: `disable`, `enable`.

* `ha_priority` - HA election priority for the PING server.
* `id` - ID.
* `ip` - Secondary IP address of the interface.
* `ping_serv_status` - Ping-Serv-Status.
* `seq` - Seq.

The `interface` block supports:

* `ac_name` - PPPoE server name.
* `aggregate` - Aggregate.
* `algorithm` - Frame distribution algorithm. Valid values: `L2`, `L3`, `L4`.

* `alias` - Alias will be displayed with the interface name to make it easier to distinguish.
* `allowaccess` - Permitted types of management access to this interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `auto-ipsec`, `radius-acct`, `probe-response`, `capwap`, `dnp`, `ftm`, `fabric`.

* `ap_discover` - Enable/disable automatic registration of unknown FortiAP devices. Valid values: `disable`, `enable`.

* `arpforward` - Enable/disable ARP forwarding. Valid values: `disable`, `enable`.

* `atm_protocol` - ATM protocol. Valid values: `none`, `ipoa`.

* `auth_type` - PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.

* `auto_auth_extension_device` - Enable/disable automatic authorization of dedicated Fortinet extension device on this interface. Valid values: `disable`, `enable`.

* `bandwidth_measure_time` - Bandwidth measure time
* `bfd` - Bidirectional Forwarding Detection (BFD) settings. Valid values: `global`, `enable`, `disable`.

* `bfd_desired_min_tx` - BFD desired minimal transmit interval.
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal receive interval.
* `broadcast_forticlient_discovery` - Broadcast-Forticlient-Discovery. Valid values: `disable`, `enable`.

* `broadcast_forward` - Enable/disable broadcast forwarding. Valid values: `disable`, `enable`.

* `captive_portal` - Enable/disable captive portal.
* `cli_conn_status` - Cli-Conn-Status.
* `color` - Color of icon on the GUI.
* `ddns` - Ddns. Valid values: `disable`, `enable`.

* `ddns_auth` - Ddns-Auth. Valid values: `disable`, `tsig`.

* `ddns_domain` - Ddns-Domain.
* `ddns_key` - Ddns-Key.
* `ddns_keyname` - Ddns-Keyname.
* `ddns_password` - Ddns-Password.
* `ddns_server` - Ddns-Server. Valid values: `dhs.org`, `dyndns.org`, `dyns.net`, `tzo.com`, `ods.org`, `vavic.com`, `now.net.cn`, `dipdns.net`, `easydns.com`, `genericDDNS`.

* `ddns_server_ip` - Ddns-Server-Ip.
* `ddns_sn` - Ddns-Sn.
* `ddns_ttl` - Ddns-Ttl.
* `ddns_username` - Ddns-Username.
* `ddns_zone` - Ddns-Zone.
* `dedicated_to` - Configure interface for single purpose. Valid values: `none`, `management`.

* `defaultgw` - Enable to get the gateway IP from the DHCP or PPPoE server. Valid values: `disable`, `enable`.

* `description` - Description.
* `detected_peer_mtu` - Detected-Peer-Mtu.
* `detectprotocol` - Protocols used to detect the server. Valid values: `ping`, `tcp-echo`, `udp-echo`.

* `detectserver` - Gateway's ping server for this IP.
* `device_access_list` - Device-Access-List.
* `device_identification` - Enable/disable passively gathering of device identity information about the devices on the network connected to this interface. Valid values: `disable`, `enable`.

* `device_identification_active_scan` - Device-Identification-Active-Scan. Valid values: `disable`, `enable`.

* `device_netscan` - Device-Netscan. Valid values: `disable`, `enable`.

* `device_user_identification` - Enable/disable passive gathering of user identity information about users on this interface. Valid values: `disable`, `enable`.

* `devindex` - Devindex.
* `dhcp_client_identifier` - DHCP client identifier.
* `dhcp_relay_agent_option` - Enable/disable DHCP relay agent option. Valid values: `disable`, `enable`.

* `dhcp_relay_interface` - Specify outgoing interface to reach server.
* `dhcp_relay_interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `dhcp_relay_ip` - DHCP relay IP address.
* `dhcp_relay_request_all_server` - Dhcp-Relay-Request-All-Server. Valid values: `disable`, `enable`.

* `dhcp_relay_service` - Enable/disable allowing this interface to act as a DHCP relay. Valid values: `disable`, `enable`.

* `dhcp_relay_type` - DHCP relay type (regular or IPsec). Valid values: `regular`, `ipsec`.

* `dhcp_renew_time` - DHCP renew time in seconds (300-604800), 0 means use the renew time provided by the server.
* `disc_retry_timeout` - Time in seconds to wait before retrying to start a PPPoE discovery, 0 means no timeout.
* `disconnect_threshold` - Time in milliseconds to wait before sending a notification that this interface is down or disconnected.
* `distance` - Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
* `dns_query` - Dns-Query. Valid values: `disable`, `recursive`, `non-recursive`.

* `dns_server_override` - Enable/disable use DNS acquired by DHCP or PPPoE. Valid values: `disable`, `enable`.

* `drop_fragment` - Enable/disable drop fragment packets. Valid values: `disable`, `enable`.

* `drop_overlapped_fragment` - Enable/disable drop overlapped fragment packets. Valid values: `disable`, `enable`.

* `egress_cos` - Override outgoing CoS in user VLAN tag. Valid values: `disable`, `cos0`, `cos1`, `cos2`, `cos3`, `cos4`, `cos5`, `cos6`, `cos7`.

* `egress_shaping_profile` - Outgoing traffic shaping profile.
* `eip` - Eip.
* `endpoint_compliance` - Endpoint-Compliance. Valid values: `disable`, `enable`.

* `estimated_downstream_bandwidth` - Estimated maximum downstream bandwidth (kbps). Used to estimate link utilization.
* `estimated_upstream_bandwidth` - Estimated maximum upstream bandwidth (kbps). Used to estimate link utilization.
* `explicit_ftp_proxy` - Enable/disable the explicit FTP proxy on this interface. Valid values: `disable`, `enable`.

* `explicit_web_proxy` - Enable/disable the explicit web proxy on this interface. Valid values: `disable`, `enable`.

* `external` - Enable/disable identifying the interface as an external interface (which usually means it's connected to the Internet). Valid values: `disable`, `enable`.

* `fail_action_on_extender` - Action on extender when interface fail . Valid values: `soft-restart`, `hard-restart`, `reboot`.

* `fail_alert_interfaces` - Names of the FortiGate interfaces to which the link failure alert is sent.
* `fail_alert_method` - Select link-failed-signal or link-down method to alert about a failed link. Valid values: `link-failed-signal`, `link-down`.

* `fail_detect` - Enable/disable fail detection features for this interface. Valid values: `disable`, `enable`.

* `fail_detect_option` - Options for detecting that this interface has failed. Valid values: `detectserver`, `link-down`.

* `fdp` - Fdp. Valid values: `disable`, `enable`.

* `fortiheartbeat` - Fortiheartbeat. Valid values: `disable`, `enable`.

* `fortilink` - Enable FortiLink to dedicate this interface to manage other Fortinet devices. Valid values: `disable`, `enable`.

* `fortilink_backup_link` - Fortilink-Backup-Link.
* `fortilink_neighbor_detect` - Protocol for FortiGate neighbor discovery. Valid values: `lldp`, `fortilink`.

* `fortilink_split_interface` - Enable/disable FortiLink split interface to connect member link to different FortiSwitch in stack for uplink redundancy. Valid values: `disable`, `enable`.

* `fortilink_stacking` - Enable/disable FortiLink switch-stacking on this interface. Valid values: `disable`, `enable`.

* `forward_domain` - Transparent mode forward domain.
* `forward_error_correction` - Enable/disable forward error correction (FEC Clause 91). Valid values: `disable`, `enable`, `rs-fec`, `base-r-fec`, `fec-cl91`, `fec-cl74`.

* `fp_anomaly` - Fp-Anomaly. Valid values: `drop_tcp_fin_noack`, `pass_winnuke`, `pass_tcpland`, `pass_udpland`, `pass_icmpland`, `pass_ipland`, `pass_iprr`, `pass_ipssrr`, `pass_iplsrr`, `pass_ipstream`, `pass_ipsecurity`, `pass_iptimestamp`, `pass_ipunknown_option`, `pass_ipunknown_prot`, `pass_icmp_frag`, `pass_tcp_no_flag`, `pass_tcp_fin_noack`, `drop_winnuke`, `drop_tcpland`, `drop_udpland`, `drop_icmpland`, `drop_ipland`, `drop_iprr`, `drop_ipssrr`, `drop_iplsrr`, `drop_ipstream`, `drop_ipsecurity`, `drop_iptimestamp`, `drop_ipunknown_option`, `drop_ipunknown_prot`, `drop_icmp_frag`, `drop_tcp_no_flag`.

* `fp_disable` - Fp-Disable. Valid values: `all`, `ipsec`, `none`.

* `gateway_address` - Gateway address
* `gi_gk` - Enable/disable Gi Gatekeeper. Valid values: `disable`, `enable`.

* `gwaddr` - Gateway address
* `gwdetect` - Enable/disable detect gateway alive for first. Valid values: `disable`, `enable`.

* `ha_priority` - HA election priority for the PING server.
* `icmp_accept_redirect` - Enable/disable ICMP accept redirect. Valid values: `disable`, `enable`.

* `icmp_redirect` - Icmp-Redirect. Valid values: `disable`, `enable`.

* `icmp_send_redirect` - Enable/disable ICMP send redirect. Valid values: `disable`, `enable`.

* `ident_accept` - Enable/disable authentication for this interface. Valid values: `disable`, `enable`.

* `idle_timeout` - PPPoE auto disconnect after idle timeout seconds, 0 means no timeout.
* `if_mdix` - If-Mdix. Valid values: `auto`, `normal`, `crossover`.

* `if_media` - If-Media. Valid values: `auto`, `copper`, `fiber`.

* `in_force_vlan_cos` - In-Force-Vlan-Cos.
* `inbandwidth` - Bandwidth limit for incoming traffic (0 - 16776000 kbps), 0 means unlimited.
* `ingress_cos` - Override incoming CoS in user VLAN tag on VLAN interface or assign a priority VLAN tag on physical interface. Valid values: `disable`, `cos0`, `cos1`, `cos2`, `cos3`, `cos4`, `cos5`, `cos6`, `cos7`.

* `ingress_shaping_profile` - Incoming traffic shaping profile.
* `ingress_spillover_threshold` - Ingress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `internal` - Implicitly created.
* `ip` - Interface IPv4 address and subnet mask, syntax: X.X.X.X/24.
* `ip_managed_by_fortiipam` - Enable/disable automatic IP address assignment of this interface by FortiIPAM. Valid values: `disable`, `enable`.

* `ipmac` - Enable/disable IP/MAC binding. Valid values: `disable`, `enable`.

* `ips_sniffer_mode` - Enable/disable the use of this interface as a one-armed sniffer. Valid values: `disable`, `enable`.

* `ipunnumbered` - Unnumbered IP used for PPPoE interfaces for which no unique local address is provided.
* `ipv6` - Ipv6. The structure of `ipv6` block is documented below.
* `l2forward` - Enable/disable l2 forwarding. Valid values: `disable`, `enable`.

* `l2tp_client` - Enable/disable this interface as a Layer 2 Tunnelling Protocol (L2TP) client. Valid values: `disable`, `enable`.

* `lacp_ha_slave` - LACP HA slave. Valid values: `disable`, `enable`.

* `lacp_mode` - LACP mode. Valid values: `static`, `passive`, `active`.

* `lacp_speed` - How often the interface sends LACP messages. Valid values: `slow`, `fast`.

* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect.
* `link_up_delay` - Number of milliseconds to wait before considering a link is up.
* `listen_forticlient_connection` - Listen-Forticlient-Connection. Valid values: `disable`, `enable`.

* `lldp_network_policy` - LLDP-MED network policy profile.
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception. Valid values: `disable`, `enable`, `vdom`.

* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission. Valid values: `enable`, `disable`, `vdom`.

* `log` - Log. Valid values: `disable`, `enable`.

* `macaddr` - Change the interface's MAC address.
* `managed_subnetwork_size` - Number of IP addresses to be allocated by FortiIPAM and used by this FortiGate unit's DHCP server settings. Valid values: `256`, `512`, `1024`, `2048`, `4096`, `8192`, `16384`, `32768`, `65536`.

* `management_ip` - High Availability in-band management IP address of this interface.
* `max_egress_burst_rate` - Max-Egress-Burst-Rate.
* `max_egress_rate` - Max-Egress-Rate.
* `measured_downstream_bandwidth` - Measured downstream bandwidth (kbps).
* `measured_upstream_bandwidth` - Measured upstream bandwidth (kbps).
* `mediatype` - Select SFP media interface type Valid values: `serdes-sfp`, `sgmii-sfp`, `cfp2-sr10`, `cfp2-lr4`, `serdes-copper-sfp`, `sr`, `cr`, `lr`, `qsfp28-sr4`, `qsfp28-lr4`, `qsfp28-cr4`, `sr4`, `cr4`, `lr4`.

* `member` - Physical interfaces that belong to the aggregate or redundant interface.
* `min_links` - Minimum number of aggregated ports that must be up.
* `min_links_down` - Action to take when less than the configured minimum number of links are active. Valid values: `operational`, `administrative`.

* `mode` - Addressing mode (static, DHCP, PPPoE). Valid values: `static`, `dhcp`, `pppoe`, `pppoa`, `ipoa`, `eoa`.

* `monitor_bandwidth` - Enable monitoring bandwidth on this interface. Valid values: `disable`, `enable`.

* `mtu` - MTU value for this interface.
* `mtu_override` - Enable to set a custom MTU for this interface. Valid values: `disable`, `enable`.

* `mux_type` - Multiplexer type Valid values: `llc-encaps`, `vc-encaps`.

* `name` - Name.
* `ndiscforward` - Enable/disable NDISC forwarding. Valid values: `disable`, `enable`.

* `netbios_forward` - Enable/disable NETBIOS forwarding. Valid values: `disable`, `enable`.

* `netflow_sampler` - Enable/disable NetFlow on this interface and set the data that NetFlow collects (rx, tx, or both). Valid values: `disable`, `tx`, `rx`, `both`.

* `np_qos_profile` - Np-Qos-Profile.
* `npu_fastpath` - Npu-Fastpath. Valid values: `disable`, `enable`.

* `nst` - Nst. Valid values: `disable`, `enable`.

* `out_force_vlan_cos` - Out-Force-Vlan-Cos.
* `outbandwidth` - Bandwidth limit for outgoing traffic (0 - 16776000 kbps), 0 means unlimited.
* `padt_retry_timeout` - PPPoE Active Discovery Terminate (PADT) used to terminate sessions after an idle time.
* `password` - PPPoE account's password.
* `peer_interface` - Peer-Interface.
* `phy_mode` - DSL physical mode. Valid values: `auto`, `adsl`, `vdsl`.

* `ping_serv_status` - Ping-Serv-Status.
* `poe` - Enable/disable PoE status. Valid values: `disable`, `enable`.

* `polling_interval` - sFlow polling interval (1 - 255 sec).
* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation. Valid values: `disable`, `enable`.

* `pptp_auth_type` - PPTP authentication type. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.

* `pptp_client` - Enable/disable PPTP client. Valid values: `disable`, `enable`.

* `pptp_password` - PPTP password.
* `pptp_server_ip` - PPTP server IP address.
* `pptp_timeout` - Idle timer in minutes (0 for disabled).
* `pptp_user` - PPTP user name.
* `preserve_session_route` - Enable/disable preservation of session route when dirty. Valid values: `disable`, `enable`.

* `priority` - Priority of learned routes.
* `priority_override` - Enable/disable fail back to higher priority port once recovered. Valid values: `disable`, `enable`.

* `proxy_captive_portal` - Enable/disable proxy captive portal on this interface. Valid values: `disable`, `enable`.

* `redundant_interface` - Redundant-Interface.
* `remote_ip` - Remote IP address of tunnel.
* `replacemsg_override_group` - Replacement message override group.
* `retransmission` - Enable/disable DSL retransmission. Valid values: `disable`, `enable`.

* `ring_rx` - RX ring size.
* `ring_tx` - TX ring size.
* `role` - Interface role. Valid values: `lan`, `wan`, `dmz`, `undefined`.

* `sample_direction` - Data that NetFlow collects (rx, tx, or both). Valid values: `rx`, `tx`, `both`.

* `sample_rate` - sFlow sample rate (10 - 99999).
* `scan_botnet_connections` - Scan-Botnet-Connections. Valid values: `disable`, `block`, `monitor`.

* `secondary_IP` - Enable/disable adding a secondary IP to this interface. Valid values: `disable`, `enable`.

* `secondaryip` - Secondaryip. The structure of `secondaryip` block is documented below.
* `security_8021x_dynamic_vlan_id` - VLAN ID for virtual switch.
* `security_8021x_master` - 802.1X master virtual-switch.
* `security_8021x_mode` - 802.1X mode. Valid values: `default`, `dynamic-vlan`, `fallback`, `slave`.

* `security_exempt_list` - Name of security-exempt-list.
* `security_external_logout` - URL of external authentication logout server.
* `security_external_web` - URL of external authentication web server.
* `security_groups` - User groups that can authenticate with the captive portal.
* `security_mac_auth_bypass` - Enable/disable MAC authentication bypass. Valid values: `disable`, `enable`, `mac-auth-only`.

* `security_mode` - Turn on captive portal authentication for this interface. Valid values: `none`, `captive-portal`, `802.1X`.

* `security_redirect_url` - URL redirection after disclaimer/authentication.
* `service_name` - PPPoE service name.
* `sflow_sampler` - Enable/disable sFlow on this interface. Valid values: `disable`, `enable`.

* `speed` - Interface speed. The default setting and the options available depend on the interface hardware. Valid values: `auto`, `10full`, `10half`, `100full`, `100half`, `1000full`, `1000half`, `10000full`, `1000auto`, `10000auto`, `40000full`, `100Gfull`, `25000full`, `40000auto`, `25000auto`, `100Gauto`.

* `spillover_threshold` - Egress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `src_check` - Enable/disable source IP check. Valid values: `disable`, `enable`.

* `status` - Bring the interface up or shut the interface down. Valid values: `down`, `up`.

* `stp` - Enable/disable STP. Valid values: `disable`, `enable`.

* `stp_ha_slave` - Control STP behaviour on HA slave. Valid values: `disable`, `enable`, `priority-adjust`.

* `stpforward` - Enable/disable STP forwarding. Valid values: `disable`, `enable`.

* `stpforward_mode` - Configure STP forwarding mode. Valid values: `rpl-all-ext-id`, `rpl-bridge-ext-id`, `rpl-nothing`.

* `strip_priority_vlan_tag` - Strip-Priority-Vlan-Tag. Valid values: `disable`, `enable`.

* `subst` - Enable to always send packets from this interface to a destination MAC address. Valid values: `disable`, `enable`.

* `substitute_dst_mac` - Destination MAC address that all packets are sent to from this interface.
* `swc_first_create` - Initial create for switch-controller VLANs.
* `swc_vlan` - Swc-Vlan.
* `switch` - Switch.
* `switch_controller_access_vlan` - Block FortiSwitch port-to-port traffic. Valid values: `disable`, `enable`.

* `switch_controller_arp_inspection` - Enable/disable FortiSwitch ARP inspection. Valid values: `disable`, `enable`.

* `switch_controller_auth` - Switch-Controller-Auth. Valid values: `radius`, `usergroup`.

* `switch_controller_dhcp_snooping` - Switch controller DHCP snooping. Valid values: `disable`, `enable`.

* `switch_controller_dhcp_snooping_option82` - Switch controller DHCP snooping option82. Valid values: `disable`, `enable`.

* `switch_controller_dhcp_snooping_verify_mac` - Switch controller DHCP snooping verify MAC. Valid values: `disable`, `enable`.

* `switch_controller_feature` - Interface's purpose when assigning traffic (read only). Valid values: `none`, `default-vlan`, `quarantine`, `sniffer`, `voice`, `camera`, `rspan`, `video`, `nac`.

* `switch_controller_igmp_snooping` - Switch controller IGMP snooping. Valid values: `disable`, `enable`.

* `switch_controller_igmp_snooping_fast_leave` - Switch controller IGMP snooping fast-leave. Valid values: `disable`, `enable`.

* `switch_controller_igmp_snooping_proxy` - Switch controller IGMP snooping proxy. Valid values: `disable`, `enable`.

* `switch_controller_iot_scanning` - Enable/disable managed FortiSwitch IoT scanning. Valid values: `disable`, `enable`.

* `switch_controller_learning_limit` - Limit the number of dynamic MAC addresses on this VLAN (1 - 128, 0 = no limit, default).
* `switch_controller_mgmt_vlan` - VLAN to use for FortiLink management purposes.
* `switch_controller_nac` - Integrated NAC settings for managed FortiSwitch.
* `switch_controller_radius_server` - Switch-Controller-Radius-Server.
* `switch_controller_rspan_mode` - Stop Layer2 MAC learning and interception of BPDUs and other packets on this interface. Valid values: `disable`, `enable`.

* `switch_controller_source_ip` - Source IP address used in FortiLink over L3 connections. Valid values: `outbound`, `fixed`.

* `switch_controller_traffic_policy` - Switch controller traffic policy for the VLAN.
* `tc_mode` - DSL transfer mode. Valid values: `ptm`, `atm`.

* `tcp_mss` - TCP maximum segment size. 0 means do not change segment size.
* `trunk` - Enable/disable VLAN trunk. Valid values: `disable`, `enable`.

* `trust_ip_1` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_2` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_3` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip6_1` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_2` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_3` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `type` - Interface type. Valid values: `physical`, `vlan`, `aggregate`, `redundant`, `tunnel`, `wireless`, `vdom-link`, `loopback`, `switch`, `hard-switch`, `hdlc`, `vap-switch`, `wl-mesh`, `fortilink`, `switch-vlan`, `fctrl-trunk`, `tdm`, `fext-wan`, `vxlan`, `emac-vlan`, `geneve`, `ssl`.

* `username` - Username of the PPPoE account, provided by your ISP.
* `vci` - Virtual Channel ID
* `vectoring` - Enable/disable DSL vectoring. Valid values: `disable`, `enable`.

* `vindex` - Vindex.
* `vlan_protocol` - Ethernet protocol of VLAN. Valid values: `8021q`, `8021ad`.

* `vlanforward` - Enable/disable traffic forwarding between VLANs on this interface. Valid values: `disable`, `enable`.

* `vlanid` - VLAN ID (1 - 4094).
* `vpi` - Virtual Path ID
* `vrf` - Virtual Routing Forwarding ID.
* `vrrp` - Vrrp. The structure of `vrrp` block is documented below.
* `vrrp_virtual_mac` - Enable/disable use of virtual MAC for VRRP. Valid values: `disable`, `enable`.

* `wccp` - Enable/disable WCCP on this interface. Used for encapsulated WCCP communication between WCCP clients and servers. Valid values: `disable`, `enable`.

* `weight` - Default weight for static routes (if route has no weight configured).
* `wifi_5g_threshold` - Minimal signal strength to be considered as a good 5G AP.
* `wifi_acl` - Access control for MAC addresses in the MAC list. Valid values: `deny`, `allow`.

* `wifi_ap_band` - How to select the AP to connect. Valid values: `any`, `5g-preferred`, `5g-only`.

* `wifi_auth` - WiFi authentication. Valid values: `PSK`, `RADIUS`, `radius`, `usergroup`.

* `wifi_auto_connect` - Enable/disable WiFi network auto connect. Valid values: `disable`, `enable`.

* `wifi_auto_save` - Enable/disable WiFi network automatic save. Valid values: `disable`, `enable`.

* `wifi_broadcast_ssid` - Enable/disable SSID broadcast in the beacon. Valid values: `disable`, `enable`.

* `wifi_encrypt` - Data encryption. Valid values: `TKIP`, `AES`.

* `wifi_fragment_threshold` - WiFi fragment threshold (800 - 2346).
* `wifi_key` - WiFi WEP Key.
* `wifi_keyindex` - WEP key index (1 - 4).
* `wifi_mac_filter` - Enable/disable MAC filter status. Valid values: `disable`, `enable`.

* `wifi_passphrase` - WiFi pre-shared key for WPA.
* `wifi_radius_server` - WiFi RADIUS server for WPA.
* `wifi_rts_threshold` - WiFi RTS threshold (256 - 2346).
* `wifi_security` - Wireless access security of SSID. Valid values: `None`, `WEP64`, `wep64`, `WEP128`, `wep128`, `WPA_PSK`, `WPA_RADIUS`, `WPA`, `WPA2`, `WPA2_AUTO`, `open`, `wpa-personal`, `wpa-enterprise`, `wpa-only-personal`, `wpa-only-enterprise`, `wpa2-only-personal`, `wpa2-only-enterprise`.

* `wifi_ssid` - IEEE 802.11 Service Set Identifier.
* `wifi_usergroup` - WiFi user group for WPA.
* `wins_ip` - WINS server IP.

The `ipv6` block supports:

* `autoconf` - Enable/disable address auto config. Valid values: `disable`, `enable`.

* `cli_conn6_status` - Cli-Conn6-Status.
* `dhcp6_client_options` - Dhcp6-Client-Options. Valid values: `rapid`, `iapd`, `iana`, `dns`, `dnsname`.

* `dhcp6_information_request` - Enable/disable DHCPv6 information request. Valid values: `disable`, `enable`.

* `dhcp6_prefix_delegation` - Enable/disable DHCPv6 prefix delegation. Valid values: `disable`, `enable`.

* `dhcp6_prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `dhcp6_prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `dhcp6_prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).
* `dhcp6_relay_ip` - DHCPv6 relay IP address.
* `dhcp6_relay_service` - Enable/disable DHCPv6 relay. Valid values: `disable`, `enable`.

* `dhcp6_relay_type` - DHCPv6 relay type. Valid values: `regular`.

* `icmp6_send_redirect` - Icmp6-Send-Redirect. Valid values: `disable`, `enable`.

* `interface_identifier` - IPv6 interface identifier.
* `ip6_address` - Primary IPv6 address prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
* `ip6_allowaccess` - Allow management access to the interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `capwap`, `fabric`.

* `ip6_default_life` - Default life (sec).
* `ip6_delegated_prefix_list` - Ip6-Delegated-Prefix-List. The structure of `ip6_delegated_prefix_list` block is documented below.
* `ip6_dns_server_override` - Enable/disable using the DNS server acquired by DHCP. Valid values: `disable`, `enable`.

* `ip6_extra_addr` - Ip6-Extra-Addr. The structure of `ip6_extra_addr` block is documented below.
* `ip6_hop_limit` - Hop limit (0 means unspecified).
* `ip6_link_mtu` - IPv6 link MTU.
* `ip6_manage_flag` - Enable/disable the managed flag. Valid values: `disable`, `enable`.

* `ip6_max_interval` - IPv6 maximum interval (4 to 1800 sec).
* `ip6_min_interval` - IPv6 minimum interval (3 to 1350 sec).
* `ip6_mode` - Addressing mode (static, DHCP, delegated). Valid values: `static`, `dhcp`, `pppoe`, `delegated`.

* `ip6_other_flag` - Enable/disable the other IPv6 flag. Valid values: `disable`, `enable`.

* `ip6_prefix_list` - Ip6-Prefix-List. The structure of `ip6_prefix_list` block is documented below.
* `ip6_prefix_mode` - Ip6-Prefix-Mode. Valid values: `dhcp6`, `ra`.

* `ip6_reachable_time` - IPv6 reachable time (milliseconds; 0 means unspecified).
* `ip6_retrans_time` - IPv6 retransmit time (milliseconds; 0 means unspecified).
* `ip6_send_adv` - Enable/disable sending advertisements about the interface. Valid values: `disable`, `enable`.

* `ip6_subnet` - Subnet to routing prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
* `ip6_upstream_interface` - Interface name providing delegated information.
* `nd_cert` - Neighbor discovery certificate.
* `nd_cga_modifier` - Neighbor discovery CGA modifier.
* `nd_mode` - Neighbor discovery mode. Valid values: `basic`, `SEND-compatible`.

* `nd_security_level` - Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).
* `nd_timestamp_delta` - Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).
* `nd_timestamp_fuzz` - Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).
* `unique_autoconf_addr` - Enable/disable unique auto config address. Valid values: `disable`, `enable`.

* `vrip6_link_local` - Link-local IPv6 address of virtual router.
* `vrrp_virtual_mac6` - Enable/disable virtual MAC for VRRP. Valid values: `disable`, `enable`.

* `vrrp6` - Vrrp6. The structure of `vrrp6` block is documented below.

The `ip6_delegated_prefix_list` block supports:

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `disable`, `enable`.

* `onlink_flag` - Enable/disable the onlink flag. Valid values: `disable`, `enable`.

* `prefix_id` - Prefix ID.
* `rdnss` - Recursive DNS server option.
* `rdnss_service` - Recursive DNS service option. Valid values: `delegated`, `default`, `specify`.

* `subnet` - Add subnet ID to routing prefix.
* `upstream_interface` - Name of the interface that provides delegated information.

The `ip6_extra_addr` block supports:

* `prefix` - IPv6 address prefix.

The `ip6_prefix_list` block supports:

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `disable`, `enable`.

* `dnssl` - DNS search list option.
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `disable`, `enable`.

* `preferred_life_time` - Preferred life time (sec).
* `prefix` - IPv6 prefix.
* `rdnss` - Recursive DNS server option.
* `valid_life_time` - Valid life time (sec).

The `vrrp6` block supports:

* `accept_mode` - Enable/disable accept mode. Valid values: `disable`, `enable`.

* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `preempt` - Enable/disable preempt mode. Valid values: `disable`, `enable`.

* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable VRRP. Valid values: `disable`, `enable`.

* `vrdst6` - Monitor the route to this destination.
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip6` - IPv6 address of the virtual router.

The `secondaryip` block supports:

* `allowaccess` - Management access settings for the secondary IP address. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `auto-ipsec`, `radius-acct`, `probe-response`, `capwap`, `dnp`, `ftm`, `fabric`.

* `detectprotocol` - Protocols used to detect the server. Valid values: `ping`, `tcp-echo`, `udp-echo`.

* `detectserver` - Gateway's ping server for this IP.
* `gwdetect` - Enable/disable detect gateway alive for first. Valid values: `disable`, `enable`.

* `ha_priority` - HA election priority for the PING server.
* `id` - ID.
* `ip` - Secondary IP address of the interface.
* `ping_serv_status` - Ping-Serv-Status.
* `seq` - Seq.

The `vrrp` block supports:

* `accept_mode` - Enable/disable accept mode. Valid values: `disable`, `enable`.

* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `ignore_default_route` - Enable/disable ignoring of default route when checking destination. Valid values: `disable`, `enable`.

* `preempt` - Enable/disable preempt mode. Valid values: `disable`, `enable`.

* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable this VRRP configuration. Valid values: `disable`, `enable`.

* `version` - VRRP version. Valid values: `2`, `3`.

* `vrdst` - Monitor the route to this destination.
* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip` - IP address of the virtual router.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFsp Vlan can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fsp_vlan.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
