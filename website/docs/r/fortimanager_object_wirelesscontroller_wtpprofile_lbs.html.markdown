---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_wtpprofile_lbs"
description: |-
  Set various location based service (LBS) options.
---

# fortimanager_object_wirelesscontroller_wtpprofile_lbs
Set various location based service (LBS) options.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wtp_profile` - Wtp Profile.

* `aeroscout` - Enable/disable AeroScout Real Time Location Service (RTLS) support (default = disable). Valid values: `disable`, `enable`.

* `aeroscout_ap_mac` - Use BSSID or board MAC address as AP MAC address in AeroScout AP messages (default = bssid). Valid values: `bssid`, `board-mac`.

* `aeroscout_mmu_report` - Enable/disable compounded AeroScout tag and MU report (default = enable). Valid values: `disable`, `enable`.

* `aeroscout_mu` - Enable/disable AeroScout Mobile Unit (MU) support (default = disable). Valid values: `disable`, `enable`.

* `aeroscout_mu_factor` - AeroScout MU mode dilution factor (default = 20).
* `aeroscout_mu_timeout` - AeroScout MU mode timeout (0 - 65535 sec, default = 5).
* `aeroscout_server_ip` - IP address of AeroScout server.
* `aeroscout_server_port` - AeroScout server UDP listening port.
* `ekahau_blink_mode` - Enable/disable Ekahau blink mode (now known as AiRISTA Flow) to track and locate WiFi tags (default = disable). Valid values: `disable`, `enable`.

* `ekahau_tag` - WiFi frame MAC address or WiFi Tag.
* `erc_server_ip` - IP address of Ekahau RTLS Controller (ERC).
* `erc_server_port` - Ekahau RTLS Controller (ERC) UDP listening port.
* `fortipresence` - Enable/disable FortiPresence to monitor the location and activity of WiFi clients even if they don't connect to this WiFi network (default = disable). Valid values: `disable`, `enable`, `enable2`, `foreign`, `both`.

* `fortipresence_ble` - Enable/disable FortiPresence finding and reporting BLE devices. Valid values: `disable`, `enable`.

* `fortipresence_frequency` - FortiPresence report transmit frequency (5 - 65535 sec, default = 30).
* `fortipresence_port` - FortiPresence server UDP listening port (default = 3000).
* `fortipresence_project` - FortiPresence project name (max. 16 characters, default = fortipresence).
* `fortipresence_rogue` - Enable/disable FortiPresence finding and reporting rogue APs. Valid values: `disable`, `enable`.

* `fortipresence_secret` - FortiPresence secret password (max. 16 characters).
* `fortipresence_server` - FortiPresence server IP address.
* `fortipresence_server_addr_type` - FortiPresence server address type (default = ipv4). Valid values: `fqdn`, `ipv4`.

* `fortipresence_server_fqdn` - FQDN of FortiPresence server.
* `fortipresence_unassoc` - Enable/disable FortiPresence finding and reporting unassociated stations. Valid values: `disable`, `enable`.

* `polestar` - Enable/disable PoleStar BLE NAO Track Real Time Location Service (RTLS) support (default = disable). Valid values: `disable`, `enable`.

* `polestar_accumulation_interval` - Time that measurements should be accumulated in seconds (default = 2).
* `polestar_asset_addrgrp_list` - Tags and asset addrgrp list to be reported.
* `polestar_asset_uuid_list1` - Tags and asset UUID list 1 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
* `polestar_asset_uuid_list2` - Tags and asset UUID list 2 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
* `polestar_asset_uuid_list3` - Tags and asset UUID list 3 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
* `polestar_asset_uuid_list4` - Tags and asset UUID list 4 to be reported (string in the format of 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX').
* `polestar_protocol` - Select the protocol to report Measurements, Advertising Data, or Location Data to NAO Cloud. (default = WSS). Valid values: `WSS`.

* `polestar_reporting_interval` - Time between reporting accumulated measurements in seconds (default = 2).
* `polestar_server_fqdn` - FQDN of PoleStar Nao Track Server (default = ws.nao-cloud.com).
* `polestar_server_path` - Path of PoleStar Nao Track Server (default = /v1/token/&lt;access_token&gt;/pst-v2).
* `polestar_server_port` - Port of PoleStar Nao Track Server (default = 443).
* `polestar_server_token` - Access Token of PoleStar Nao Track Server.
* `station_locate` - Enable/disable client station locating services for all clients, whether associated or not (default = disable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWirelessController WtpProfileLbs can be imported using any of these accepted formats:
```
Set import_options = ["wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_wtpprofile_lbs.labelname ObjectWirelessControllerWtpProfileLbs
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
