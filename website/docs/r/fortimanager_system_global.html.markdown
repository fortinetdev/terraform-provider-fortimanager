---
subcategory: "System Global"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_global"
description: |-
  Global range attributes.
---

# fortimanager_system_global
Global range attributes.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `mc_policy_disabled_adoms`: `fortimanager_system_global_mcpolicydisabledadoms`
>- `ssl_cipher_suites`: `fortimanager_system_global_sslciphersuites`



## Example Usage

```hcl
resource "fortimanager_system_global" "trname" {
  hostname = "FMG-Terr"
  language = "en"
}
```

## Argument Reference


The following arguments are supported:


* `admin_lockout_duration` - Lockout duration(sec) for administration.
* `admin_lockout_method` - Lockout method for administration. ip - Lockout by IP user - Lockout by user Valid values: `ip`, `user`.

* `admin_lockout_threshold` - Lockout threshold for administration.
* `adom_mode` - ADOM mode. normal - Normal ADOM mode. advanced - Advanced ADOM mode. Valid values: `normal`, `advanced`.

* `adom_rev_auto_delete` - Auto delete features for old ADOM revisions. disable - Disable auto delete function for ADOM revision. by-revisions - Auto delete ADOM revisions by maximum number of revisions. by-days - Auto delete ADOM revisions by maximum days. Valid values: `disable`, `by-revisions`, `by-days`.

* `adom_rev_max_backup_revisions` - Maximum number of ADOM revisions to backup.
* `adom_rev_max_days` - Number of days to keep old ADOM revisions.
* `adom_rev_max_revisions` - Maximum number of ADOM revisions to keep.
* `adom_select` - Enable/disable select ADOM after login. disable - Disable select ADOM after login. enable - Enable select ADOM after login. Valid values: `disable`, `enable`.

* `adom_status` - ADOM status. disable - Disable ADOM mode. enable - Enable ADOM mode. Valid values: `disable`, `enable`.

* `apache_mode` - Set apache mode. event - Apache event mode. prefork - Apache prefork mode. Valid values: `event`, `prefork`.

* `api_ip_binding` - Enable/disable source IP check for JSON API request. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `clone_name_option` - set the clone object names option. default - Add a prefix of 'Clone of' to the clone name. keep - Keep the original name for user to edit. Valid values: `default`, `keep`.

* `clt_cert_req` - Require client certificate for GUI login. disable - Disable setting. enable - Require client certificate for GUI login. optional - Optional client certificate for GUI login. Valid values: `disable`, `enable`, `optional`.

* `console_output` - Console output mode. standard - Standard output. more - More page output. Valid values: `standard`, `more`.

* `contentpack_fgt_install` - Enable/disable outbreak alert auto install for FGT ADOMS . disable - Disable the sql report auto outbreak auto install. enable - Enable the sql report auto outbreak auto install. Valid values: `disable`, `enable`.

* `country_flag` - Country flag Status. disable - Disable country flag icon beside ip address. enable - Enable country flag icon beside ip address. Valid values: `disable`, `enable`.

* `create_revision` - Enable/disable create revision by default. disable - Disable create revision by default. enable - Enable create revision by default. Valid values: `disable`, `enable`.

* `daylightsavetime` - Enable/disable daylight saving time. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `detect_unregistered_log_device` - Detect unregistered logging device from log message. disable - Disable attribute function. enable - Enable attribute function. Valid values: `disable`, `enable`.

* `device_view_mode` - Set devices/groups view mode. regular - Regular view mode. tree - Tree view mode. Valid values: `regular`, `tree`.

* `dh_params` - Minimum size of Diffie-Hellman prime for SSH/HTTPS (bits). 1024 - 1024 bits. 1536 - 1536 bits. 2048 - 2048 bits. 3072 - 3072 bits. 4096 - 4096 bits. 6144 - 6144 bits. 8192 - 8192 bits. Valid values: `1024`, `1536`, `2048`, `3072`, `4096`, `6144`, `8192`.

* `disable_module` - Disable module list. fortiview-noc - FortiView/NOC-SOC module. fortirecorder - FortiRecorder module. siem - SIEM module. soc - SOC module. ai - AI module. Valid values: `fortiview-noc`, `fortirecorder`, `siem`, `soc`, `ai`.

* `enc_algorithm` - SSL communication encryption algorithms. low - SSL communication using all available encryption algorithms. medium - SSL communication using high and medium encryption algorithms. high - SSL communication using high encryption algorithms. Valid values: `low`, `medium`, `high`.

* `faz_status` - FAZ status. disable - Disable FAZ feature. enable - Enable FAZ feature. Valid values: `disable`, `enable`.

* `fgfm_ca_cert` - set the extra fgfm CA certificates.
* `fgfm_cert_exclusive` - set if the local or CA certificates should be used exclusively. disable - Used certificate best-effort. enable - Used certificate exclusive. Valid values: `disable`, `enable`.

* `fgfm_local_cert` - set the fgfm local certificate.
* `fgfm_ssl_protocol` - set the lowest SSL protocols for fgfmsd. sslv3 - set SSLv3 as the lowest version. tlsv1.0 - set TLSv1.0 as the lowest version. tlsv1.1 - set TLSv1.1 as the lowest version. tlsv1.2 - set TLSv1.2 as the lowest version (default). tlsv1.3 - set TLSv1.3 as the lowest version. Valid values: `sslv3`, `tlsv1.0`, `tlsv1.1`, `tlsv1.2`, `tlsv1.3`.

* `fortiservice_port` - FortiService port (1 - 65535, default = 8013). Used by FortiClient endpoint compliance. Older versions of FortiClient used a different port.
* `gui_curl_timeout` - GUI curl timeout in seconds (5-300 default 30).
* `gui_polling_interval` - GUI polling interval in seconds (1-288000 default 5).
* `ha_member_auto_grouping` - Enable/disable automatically group HA members feature disable - Disable automatically grouping HA members feature. enable - Enable automatically grouping HA members only when group name is unique in your network. Valid values: `disable`, `enable`.

* `hostname` - System hostname.
* `import_ignore_addr_cmt` - Enable/Disable import ignore of address comments. disable - Disable import ignore of address comments. enable - Enable import ignore of address comments. Valid values: `disable`, `enable`.

* `language` - System global language. english - English simch - Simplified Chinese japanese - Japanese korean - Korean spanish - Spanish trach - Traditional Chinese Valid values: `english`, `simch`, `japanese`, `korean`, `spanish`, `trach`.

* `latitude` - fmg location latitude
* `ldap_cache_timeout` - LDAP browser cache timeout (seconds).
* `ldapconntimeout` - LDAP connection timeout (msec).
* `lock_preempt` - Enable/disable ADOM lock override. disable - Disable lock preempt. enable - Enable lock preempt. Valid values: `disable`, `enable`.

* `log_checksum` - Record log file hash value, timestamp, and authentication code at transmission or rolling. none - No record log file checksum. md5 - Record log file's MD5 hash value only. md5-auth - Record log file's MD5 hash value and authentication code. Valid values: `none`, `md5`, `md5-auth`.

* `log_checksum_upload` - Enable/disable upload log checksum with log files. disable - Disable attribute function. enable - Enable attribute function. Valid values: `disable`, `enable`.

* `log_forward_cache_size` - Log forwarding disk cache size (GB).
* `longitude` - fmg location longitude
* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `max_log_forward` - Maximum number of log-forward and aggregation settings.
* `max_running_reports` - Maximum number of reports generating at one time.
* `mc_policy_disabled_adoms` - Mc-Policy-Disabled-Adoms. The structure of `mc_policy_disabled_adoms` block is documented below.
* `multiple_steps_upgrade_in_autolink` - Enable/disable multiple steps upgade in autolink process disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `no_copy_permission_check` - Do not perform permission check to block object changes in different adom during copy and install. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `no_vip_value_check` - Enable/disable skipping policy instead of throwing error when vip has no default or dynamic mapping during policy copy disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `normalized_intf_zone_only` - allow normalized interface to be zone only. disable - Disable SSL low-grade encryption. enable - Enable SSL low-grade encryption. Valid values: `disable`, `enable`.

* `object_revision_db_max` - Maximum revisions for a single database (10,000-1,000,000 default 100,000).
* `object_revision_mandatory_note` - Enable/disable mandatory note when create revision. disable - Disable object revision. enable - Enable object revision. Valid values: `disable`, `enable`.

* `object_revision_object_max` - Maximum revisions for a single object (10-1000 default 100).
* `object_revision_status` - Enable/disable create revision when modify objects. disable - Disable object revision. enable - Enable object revision. Valid values: `disable`, `enable`.

* `oftp_ssl_protocol` - set the lowest SSL protocols for oftpd. sslv3 - set SSLv3 as the lowest version. tlsv1.0 - set TLSv1.0 as the lowest version. tlsv1.1 - set TLSv1.1 as the lowest version. tlsv1.2 - set TLSv1.2 as the lowest version (default). tlsv1.3 - set TLSv1.3 as the lowest version. Valid values: `sslv3`, `tlsv1.0`, `tlsv1.1`, `tlsv1.2`, `tlsv1.3`.

* `partial_install` - Enable/Disable partial install (install some objects). disable - Disable partial install function. enable - Enable partial install function. Valid values: `disable`, `enable`.

* `partial_install_force` - Enable/Disable partial install when devdb is modified. disable - Disable partial install when devdb is modified. enable - Enable partial install when devdb is modified. Valid values: `disable`, `enable`.

* `partial_install_rev` - Enable/Disable auto creating adom revision for partial install. disable - Disable partial install revision. enable - Enable partial install revision. Valid values: `disable`, `enable`.

* `per_policy_lock` - Enable/Disable per policy lock. disable - Disable per policy lock. enable - Enable per policy lock. Valid values: `disable`, `enable`.

* `perform_improve_by_ha` - Enable/Disable performance improvement by distributing tasks to HA secondary units. disable - Disable performance improvement by HA. enable - Enable performance improvement by HA. Valid values: `disable`, `enable`.

* `policy_object_icon` - show icons of policy objects. disable - Disable icon of policy objects. enable - Enable icon of policy objects. Valid values: `disable`, `enable`.

* `policy_object_in_dual_pane` - show policies and objects in dual pane. disable - Disable polices and objects in dual pane. enable - Enable polices and objects in dual pane. Valid values: `disable`, `enable`.

* `pre_login_banner` - Enable/disable pre-login banner. disable - Disable pre-login banner. enable - Enable pre-login banner. Valid values: `disable`, `enable`.

* `pre_login_banner_message` - Pre-login banner message.
* `private_data_encryption` - Enable/disable private data encryption using an AES 128-bit key. disable - Disable private data encryption using an AES 128-bit key. enable - Enable private data encryption using an AES 128-bit key. Valid values: `disable`, `enable`.

* `remoteauthtimeout` - Remote authentication (RADIUS/LDAP) timeout (sec).
* `save_last_hit_in_adomdb` - Enable/Disable save last-hit value in adomdb. disable - Disable save last-hit value in adomdb. enable - Enable save last-hit value in adomdb. Valid values: `disable`, `enable`.

* `search_all_adoms` - Enable/Disable Search all ADOMs for where-used query. disable - Disable search all ADOMs for where-used queries. enable - Enable search all ADOMs for where-used queries. Valid values: `disable`, `enable`.

* `ssh_enc_algo` - Select one or more SSH ciphers. chacha20-poly1305@openssh.com -  aes128-ctr -  aes192-ctr -  aes256-ctr -  arcfour256 -  arcfour128 -  aes128-cbc -  3des-cbc -  blowfish-cbc -  cast128-cbc -  aes192-cbc -  aes256-cbc -  arcfour -  rijndael-cbc@lysator.liu.se -  aes128-gcm@openssh.com -  aes256-gcm@openssh.com -  Valid values: `chacha20-poly1305@openssh.com`, `aes128-ctr`, `aes192-ctr`, `aes256-ctr`, `arcfour256`, `arcfour128`, `aes128-cbc`, `3des-cbc`, `blowfish-cbc`, `cast128-cbc`, `aes192-cbc`, `aes256-cbc`, `arcfour`, `rijndael-cbc@lysator.liu.se`, `aes128-gcm@openssh.com`, `aes256-gcm@openssh.com`.

* `ssh_hostkey_algo` - Select one or more SSH hostkey algorithms. ssh-rsa -  ecdsa-sha2-nistp521 -  rsa-sha2-256 -  rsa-sha2-512 -  ssh-ed25519 -  Valid values: `ssh-rsa`, `ecdsa-sha2-nistp521`, `rsa-sha2-256`, `rsa-sha2-512`, `ssh-ed25519`.

* `ssh_kex_algo` - Select one or more SSH kex algorithms. diffie-hellman-group1-sha1 -  diffie-hellman-group14-sha1 -  diffie-hellman-group14-sha256 -  diffie-hellman-group16-sha512 -  diffie-hellman-group18-sha512 -  diffie-hellman-group-exchange-sha1 -  diffie-hellman-group-exchange-sha256 -  curve25519-sha256@libssh.org -  ecdh-sha2-nistp256 -  ecdh-sha2-nistp384 -  ecdh-sha2-nistp521 -  Valid values: `diffie-hellman-group1-sha1`, `diffie-hellman-group14-sha1`, `diffie-hellman-group14-sha256`, `diffie-hellman-group16-sha512`, `diffie-hellman-group18-sha512`, `diffie-hellman-group-exchange-sha1`, `diffie-hellman-group-exchange-sha256`, `curve25519-sha256@libssh.org`, `ecdh-sha2-nistp256`, `ecdh-sha2-nistp384`, `ecdh-sha2-nistp521`.

* `ssh_mac_algo` - Select one or more SSH MAC algorithms. hmac-md5 -  hmac-md5-etm@openssh.com -  hmac-md5-96 -  hmac-md5-96-etm@openssh.com -  hmac-sha1 -  hmac-sha1-etm@openssh.com -  hmac-sha2-256 -  hmac-sha2-256-etm@openssh.com -  hmac-sha2-512 -  hmac-sha2-512-etm@openssh.com -  hmac-ripemd160 -  hmac-ripemd160@openssh.com -  hmac-ripemd160-etm@openssh.com -  umac-64@openssh.com -  umac-128@openssh.com -  umac-64-etm@openssh.com -  umac-128-etm@openssh.com -  Valid values: `hmac-md5`, `hmac-md5-etm@openssh.com`, `hmac-md5-96`, `hmac-md5-96-etm@openssh.com`, `hmac-sha1`, `hmac-sha1-etm@openssh.com`, `hmac-sha2-256`, `hmac-sha2-256-etm@openssh.com`, `hmac-sha2-512`, `hmac-sha2-512-etm@openssh.com`, `hmac-ripemd160`, `hmac-ripemd160@openssh.com`, `hmac-ripemd160-etm@openssh.com`, `umac-64@openssh.com`, `umac-128@openssh.com`, `umac-64-etm@openssh.com`, `umac-128-etm@openssh.com`.

* `ssh_strong_crypto` - Only allow strong ciphers for SSH when enabled. disable - Disable strong crypto for SSH. enable - Enable strong crypto for SSH. Valid values: `disable`, `enable`.

* `ssl_cipher_suites` - Ssl-Cipher-Suites. The structure of `ssl_cipher_suites` block is documented below.
* `ssl_low_encryption` - SSL low-grade encryption. disable - Disable SSL low-grade encryption. enable - Enable SSL low-grade encryption. Valid values: `disable`, `enable`.

* `ssl_protocol` - SSL protocols. tlsv1.3 - Enable TLSv1.3. tlsv1.2 - Enable TLSv1.2. tlsv1.1 - Enable TLSv1.1. tlsv1.0 - Enable TLSv1.0. sslv3 - Enable SSLv3. Valid values: `tlsv1.3`, `tlsv1.2`, `tlsv1.1`, `tlsv1.0`, `sslv3`.

* `ssl_static_key_ciphers` - Enable/disable SSL static key ciphers. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `table_entry_blink` - Enable/disable table entry blink in GUI disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `task_list_size` - Maximum number of completed tasks to keep.
* `tftp` - Enable/disable TFTP in `exec restore image` command (disabled by default in FIPS mode) disable - Disable TFTP enable - Enable TFTP Valid values: `disable`, `enable`.

* `timezone` - Time zone. 00 - (GMT-12:00) Eniwetak, Kwajalein. 01 - (GMT-11:00) Midway Island, Samoa. 02 - (GMT-10:00) Hawaii. 03 - (GMT-9:00) Alaska. 04 - (GMT-8:00) Pacific Time (US & Canada). 05 - (GMT-7:00) Arizona. 06 - (GMT-7:00) Mountain Time (US & Canada). 07 - (GMT-6:00) Central America. 08 - (GMT-6:00) Central Time (US & Canada). 09 - (GMT-6:00) Mexico City. 10 - (GMT-6:00) Saskatchewan. 11 - (GMT-5:00) Bogota, Lima, Quito. 12 - (GMT-5:00) Eastern Time (US & Canada). 13 - (GMT-5:00) Indiana (East). 14 - (GMT-4:00) Atlantic Time (Canada). 15 - (GMT-4:00) La Paz. 16 - (GMT-4:00) Santiago. 17 - (GMT-3:30) Newfoundland. 18 - (GMT-3:00) Brasilia. 19 - (GMT-3:00) Buenos Aires, Georgetown. 20 - (GMT-3:00) Nuuk (Greenland). 21 - (GMT-2:00) Mid-Atlantic (Deprecated). 22 - (GMT-1:00) Azores. 23 - (GMT-1:00) Cape Verde Is. 24 - (GMT) Monrovia. 25 - (GMT) London, Edinburgh. 26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna. 27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague. 28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris. 29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb. 30 - (GMT+1:00) West Central Africa. 31 - (GMT+2:00) Athens, Sofia, Vilnius. 32 - (GMT+2:00) Bucharest. 33 - (GMT+2:00) Cairo. 34 - (GMT+2:00) Harare, Pretoria. 35 - (GMT+2:00) Helsinki, Riga,Tallinn. 36 - (GMT+2:00) Jerusalem. 37 - (GMT+3:00) Baghdad. 38 - (GMT+3:00) Kuwait, Riyadh. 39 - (GMT+3:00) St.Petersburg, Volgograd. 40 - (GMT+3:00) Nairobi. 41 - (GMT+3:30) Tehran. 42 - (GMT+4:00) Abu Dhabi, Muscat. 43 - (GMT+4:00) Baku. 44 - (GMT+4:30) Kabul. 45 - (GMT+5:00) Ekaterinburg. 46 - (GMT+5:00) Islamabad, Karachi, Tashkent. 47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi. 48 - (GMT+5:45) Kathmandu. 49 - (GMT+6:00) Almaty, Novosibirsk. 50 - (GMT+6:00) Astana, Dhaka. 51 - (GMT+5:30) Sri Jayawardenepura. 52 - (GMT+6:30) Rangoon. 53 - (GMT+7:00) Bangkok, Hanoi, Jakarta. 54 - (GMT+7:00) Krasnoyarsk. 55 - (GMT+8:00) Beijing, ChongQing, HongKong, Urumqi. 56 - (GMT+8:00) Irkutsk, Ulaanbaatar. 57 - (GMT+8:00) Kuala Lumpur, Singapore. 58 - (GMT+8:00) Perth. 59 - (GMT+8:00) Taipei. 60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul. 61 - (GMT+9:00) Yakutsk. 62 - (GMT+9:30) Adelaide. 63 - (GMT+9:30) Darwin. 64 - (GMT+10:00) Brisbane. 65 - (GMT+10:00) Canberra, Melbourne, Sydney. 66 - (GMT+10:00) Guam, Port Moresby. 67 - (GMT+10:00) Hobart. 68 - (GMT+10:00) Vladivostok. 69 - (GMT+11:00) Magadan. 70 - (GMT+11:00) Solomon Is., New Caledonia. 71 - (GMT+12:00) Auckland, Wellington. 72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is. 73 - (GMT+13:00) Nuku'alofa. 74 - (GMT-4:30) Caracas. 75 - (GMT+1:00) Namibia. 76 - (GMT-5:00) Brazil-Acre. 77 - (GMT-4:00) Brazil-West. 78 - (GMT-3:00) Brazil-East. 79 - (GMT-2:00) Brazil-DeNoronha. 80 - (GMT+14:00) Kiritimati. 81 - (GMT-7:00) Baja California Sur, Chihuahua. 82 - (GMT+12:45) Chatham Islands. 83 - (GMT+3:00) Minsk. 84 - (GMT+13:00) Samoa. 85 - (GMT+3:00) Istanbul. 86 - (GMT-4:00) Paraguay. 87 - (GMT) Casablanca. 88 - (GMT+3:00) Moscow. 89 - (GMT) Greenwich Mean Time. 90 - (GMT) Dublin. 91 - (GMT) Lisbon. Valid values: `00`, `01`, `02`, `03`, `04`, `05`, `06`, `07`, `08`, `09`, `10`, `11`, `12`, `13`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `22`, `23`, `24`, `25`, `26`, `27`, `28`, `29`, `30`, `31`, `32`, `33`, `34`, `35`, `36`, `37`, `38`, `39`, `40`, `41`, `42`, `43`, `44`, `45`, `46`, `47`, `48`, `49`, `50`, `51`, `52`, `53`, `54`, `55`, `56`, `57`, `58`, `59`, `60`, `61`, `62`, `63`, `64`, `65`, `66`, `67`, `68`, `69`, `70`, `71`, `72`, `73`, `74`, `75`, `76`, `77`, `78`, `79`, `80`, `81`, `82`, `83`, `84`, `85`, `86`, `87`, `88`, `89`, `90`, `91`.

* `tunnel_mtu` - Maximum transportation unit(68 - 9000).
* `usg` - Enable/disable Fortiguard server restriction. disable - Contact any Fortiguard server enable - Contact Fortiguard server in USA only Valid values: `disable`, `enable`.

* `vdom_mirror` - VDOM mirror. disable - Disable VDOM mirror function. enable - Enable VDOM mirror function. Valid values: `disable`, `enable`.

* `webservice_proto` - Web Service connection support SSL protocols. tlsv1.3 - Web Service connection using TLSv1.3 protocol. tlsv1.2 - Web Service connection using TLSv1.2 protocol. tlsv1.1 - Web Service connection using TLSv1.1 protocol. tlsv1.0 - Web Service connection using TLSv1.0 protocol. sslv3 - Web Service connection using SSLv3 protocol. sslv2 - Web Service connection using SSLv2 protocol. Valid values: `tlsv1.3`, `tlsv1.2`, `tlsv1.1`, `tlsv1.0`, `sslv3`, `sslv2`.

* `workflow_max_sessions` - Maximum number of workflow sessions per ADOM (minimum 100).
* `workspace_mode` - Set workspace mode. disabled - Workspace disabled. normal - Workspace lock mode. workflow - Workspace workflow mode. per-adom - Per-Adom workspace mode. Valid values: `disabled`, `normal`, `workflow`, `per-adom`.

* `workspace_unlock_after_install` - Enable/disable ADOM auto-unlock after device installation. disable - Disable automatically unlock adom after device installation. enable - Enable automatically unlock adom after device installation. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `mc_policy_disabled_adoms` block supports:

* `adom_name` - Adom names.

The `ssl_cipher_suites` block supports:

* `cipher` - Cipher name
* `priority` - SSL/TLS cipher suites priority.
* `version` - SSL/TLS version the cipher suite can be used with. tls1.2-or-below - TLS 1.2 or below. tls1.3 - TLS 1.3 Valid values: `tls1.2-or-below`, `tls1.3`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Global can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_global.labelname SystemGlobal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

