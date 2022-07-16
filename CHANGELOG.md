## 1.6.0 (Unreleased)

## 1.5.0 (July 15, 2022)

IMPROVEMENTS:

* Support FortiManager version from 7.0.4 to 7.2.0
* Update difference check functionality for arguments
* Fix some argument type convert issues
* Fix error of name can't contain slash

FEATURES:

* **New Resource:** `fortimanager_object_fmg_device_blueprint`
* **New Resource:** `fortimanager_object_fmg_variable`
* **New Resource:** `fortimanager_object_system_npu_ssehascan`
* **New Resource:** `fortimanager_object_vpn_ipsec_fec`
* **New Resource:** `fortimanager_packages_firewall_acl`
* **New Resource:** `fortimanager_packages_firewall_acl6`
* **New Resource:** `fortimanager_packages_firewall_acl6_move`
* **New Resource:** `fortimanager_packages_firewall_acl_move`
* **New Resource:** `fortimanager_securityconsole_assign_package`
* **New Resource:** `fortimanager_system_ha_monitoredips`
* **New Resource:** `fortimanager_system_localinpolicy`
* **New Resource:** `fortimanager_system_localinpolicy6`

## 1.4.0 (Apr 6, 2022)

IMPROVEMENTS:

* Support FortiManager version from 7.0.1 to 7.0.3
* Fix issue of value of argument 'scopetype' and 'dynamic_sort_subtable' is null when import the resource
* Fix argument value type issue for resource fortimanager_object_webfilter_profile

FEATURES:

* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqpvenueurl`
* **New Resource:** `fortimanager_object_system_npu_backgroundssescan`
* **New Resource:** `fortimanager_object_system_npu_hpe`
* **New Resource:** `fortimanager_object_system_npu_ipreassembly`
* **New Resource:** `fortimanager_system_log_ratelimit_ratelimits`
* **New Resource:** `fortimanager_system_global_sslciphersuites`
* **New Resource:** `fortimanager_object_log_npuserver_servergroup`
* **New Resource:** `fortimanager_object_router_aspathlist`
* **New Resource:** `fortimanager_object_system_npu_tcptimeoutprofile`
* **New Resource:** `fortimanager_object_system_npu_isfnpqueues`
* **New Resource:** `fortimanager_object_wirelesscontroller_arrpprofile`
* **New Resource:** `fortimanager_packages_firewall_hyperscalepolicy6`
* **New Resource:** `fortimanager_object_system_npu_npqueues_ipprotocol`
* **New Resource:** `fortimanager_packages_firewall_hyperscalepolicy46`
* **New Resource:** `fortimanager_object_system_npu_npqueues_scheduler`
* **New Resource:** `fortimanager_object_system_npu_swehhash`
* **New Resource:** `fortimanager_object_firewall_ippool_grp`
* **New Resource:** `fortimanager_object_router_accesslist6`
* **New Resource:** `fortimanager_object_firewall_internetservice_entry`
* **New Resource:** `fortimanager_object_wirelesscontroller_nacprofile`
* **New Resource:** `fortimanager_packages_firewall_hyperscalepolicy64`
* **New Resource:** `fortimanager_object_user_connector`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile`
* **New Resource:** `fortimanager_object_wirelesscontroller_address`
* **New Resource:** `fortimanager_object_system_npu_npqueues_ethernettype`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qptermsandconditions`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_icon`
* **New Resource:** `fortimanager_object_system_npu_fpanomaly`
* **New Resource:** `fortimanager_object_system_npu_priorityprotocol`
* **New Resource:** `fortimanager_object_system_npu_npqueues_profile`
* **New Resource:** `fortimanager_object_system_npu_npqueues`
* **New Resource:** `fortimanager_system_log_fospolicystats`
* **New Resource:** `fortimanager_object_router_prefixlist`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovidernai`
* **New Resource:** `fortimanager_system_webproxy`
* **New Resource:** `fortimanager_object_router_prefixlist6`
* **New Resource:** `fortimanager_object_user_deviceaccesslist`
* **New Resource:** `fortimanager_object_system_npu_portcpumap`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge`
* **New Resource:** `fortimanager_object_dlp_fpsensitivity`
* **New Resource:** `fortimanager_object_log_npuserver_serverinfo`
* **New Resource:** `fortimanager_object_global_ips_sensor`
* **New Resource:** `fortimanager_object_endpointcontrol_fctems`
* **New Resource:** `fortimanager_object_system_npu_dswqueuedtsprofile`
* **New Resource:** `fortimanager_packages_firewall_hyperscalepolicy`
* **New Resource:** `fortimanager_object_router_accesslist`
* **New Resource:** `fortimanager_object_user_devicegroup`
* **New Resource:** `fortimanager_object_system_npu_portpathoption`
* **New Resource:** `fortimanager_object_spamfilter_iptrust`
* **New Resource:** `fortimanager_object_cifs_domaincontroller`
* **New Resource:** `fortimanager_object_log_npuserver`
* **New Resource:** `fortimanager_system_hascheduledcheck`
* **New Resource:** `fortimanager_object_system_npu_dosoptions`
* **New Resource:** `fortimanager_object_spamfilter_profile`
* **New Resource:** `fortimanager_object_router_routemap`
* **New Resource:** `fortimanager_object_router_communitylist`
* **New Resource:** `fortimanager_object_wirelesscontroller_addrgrp`
* **New Resource:** `fortimanager_object_system_npu`
* **New Resource:** `fortimanager_object_spamfilter_bwl`
* **New Resource:** `fortimanager_object_system_npu_dswdtsprofile`
* **New Resource:** `fortimanager_system_log_topology`
* **New Resource:** `fortimanager_object_user_devicecategory`
* **New Resource:** `fortimanager_object_spamfilter_bword`
* **New Resource:** `fortimanager_object_spamfilter_dnsbl`
* **New Resource:** `fortimanager_object_firewall_accessproxyvirtualhost`
* **New Resource:** `fortimanager_object_switchcontroller_dsl_policy`
* **New Resource:** `fortimanager_object_system_npu_portnpumap`
* **New Resource:** `fortimanager_object_spamfilter_mheader`
* **New Resource:** `fortimanager_object_system_npu_udptimeoutprofile`
* **New Resource:** `fortimanager_object_system_npu_npqueues_ipservice`

## 1.3.6 (Dec 13, 2021)

IMPROVEMENTS:

* **Fix argument type error for following resource:**
    * fortimanager_dvm_cmd_add_device

## 1.3.5 (Nov 5, 2021)

IMPROVEMENTS:

* **Update Document examples;**
* **Fix argument type error for following resources:**
    * fortimanager_packages_firewall_policy
    * fortimanager_object_antivirus_profile
    * fortimanager_object_authentication_scheme
    * fortimanager_object_cli_templategroup
    * fortimanager_object_dnsfilter_profile
    * fortimanager_object_dynamic_address
    * fortimanager_object_firewall_addrgrp
    * fortimanager_object_firewall_addrgrp6
    * fortimanager_object_firewall_internetservicecustomgroup
    * fortimanager_object_firewall_internetservicegroup
    * fortimanager_object_firewall_proxyaddrgrp
    * fortimanager_object_firewall_schedule_group
    * fortimanager_object_firewall_sslsshprofile
    * fortimanager_object_firewall_vip
    * fortimanager_object_firewall_vip46
    * fortimanager_object_firewall_vip6
    * fortimanager_object_firewall_vip64
    * fortimanager_object_firewall_vipgrp
    * fortimanager_object_firewall_vipgrp46
    * fortimanager_object_firewall_vipgrp6
    * fortimanager_object_firewall_vipgrp64
    * fortimanager_object_firewall_wildcardfqdn_group
    * fortimanager_object_switchcontroller_securitypolicy_8021x
    * fortimanager_object_user_domaincontroller
    * fortimanager_object_user_group
    * fortimanager_object_user_krbkeytab
    * fortimanager_object_user_peergrp
    * fortimanager_object_vpn_ssl_web_portal
    * fortimanager_object_vpnmgr_node
    * fortimanager_object_webfilter_ftgdlocalrating
    * fortimanager_object_webfilter_profile
    * fortimanager_object_wirelesscontroller_hotspot20_hsprofile
    * fortimanager_object_wirelesscontroller_vap
    * fortimanager_object_wirelesscontroller_vapgroup
    * fortimanager_object_wirelesscontroller_widsprofile
    * fortimanager_object_wirelesscontroller_wtpprofile
    * fortimanager_packages_authentication_rule
    * fortimanager_packages_firewall_centralsnatmap
    * fortimanager_packages_firewall_interfacepolicy
    * fortimanager_packages_firewall_interfacepolicy6
    * fortimanager_packages_firewall_localinpolicy
    * fortimanager_packages_firewall_localinpolicy6
    * fortimanager_packages_firewall_multicastpolicy
    * fortimanager_packages_firewall_multicastpolicy6
    * fortimanager_packages_firewall_policy46
    * fortimanager_packages_firewall_policy64
    * fortimanager_packages_firewall_proxypolicy
    * fortimanager_packages_firewall_shapingpolicy
    * fortimanager_packages_global_footer_policy
    * fortimanager_packages_global_header_policy
    * fortimanager_packages_firewall_dospolicy
    * fortimanager_packages_firewall_dospolicy6

## 1.3.0 (Jul 17, 2021)

IMPROVEMENTS:

* Support FortiManager version from 6.4 to 7.0

FEATURES:

* **New Resource:** `fortimanager_object_antivirus_mmschecksum`
* **New Resource:** `fortimanager_object_antivirus_notification`
* **New Resource:** `fortimanager_object_emailfilter_blockallowlist`
* **New Resource:** `fortimanager_object_extendercontroller_template`
* **New Resource:** `fortimanager_object_firewall_accessproxy`
* **New Resource:** `fortimanager_object_firewall_accessproxy_move`
* **New Resource:** `fortimanager_object_firewall_carrierendpointbwl`
* **New Resource:** `fortimanager_object_firewall_mmsprofile`
* **New Resource:** `fortimanager_object_switchcontroller_customcommand`
* **New Resource:** `fortimanager_object_user_device`
* **New Resource:** `fortimanager_object_videofilter_profile`
* **New Resource:** `fortimanager_object_videofilter_youtubechannelfilter`
* **New Resource:** `fortimanager_packages_firewall_consolidated_policy`
* **New Resource:** `fortimanager_packages_firewall_consolidated_policy_move`
* **New Resource:** `fortimanager_packages_firewall_policy6`
* **New Resource:** `fortimanager_packages_firewall_policy6_move`
* **New Resource:** `fortimanager_system_log_ratelimit`
* **New Resource:** `fortimanager_system_log_ratelimit_device`
* **New Resource:** `fortimanager_system_socfabri`

## 1.2.0 (Jun 11, 2021)

IMPROVEMENTS:

* Add example for fortimanager_fmupdate_fdssetting_serveroverride
* Add example for fortimanager_fmupdate_fdssetting_updateschedule
* Add example for fortimanager_fmupdate_fwmsetting
* Add example for fortimanager_fmupdate_multilayer
* Add example for fortimanager_fmupdate_publicnetwork
* Add example for fortimanager_fmupdate_serveraccesspriorities
* Add example for fortimanager_fmupdate_serveroverridestatus
* Add example for fortimanager_fmupdate_service
* Add example for fortimanager_fmupdate_webspam_fgdsetting
* Add example for fortimanager_fmupdate_webspam_webproxy
* Add example for fortimanager_object_antivirus_profile
* Add example for fortimanager_object_application_categories
* Add example for fortimanager_object_application_custom
* Add example for fortimanager_object_application_group
* Add example for fortimanager_object_application_list
* Add example for fortimanager_object_authentication_scheme
* Add example for fortimanager_object_cifs_profile
* Add example for fortimanager_object_cli_template
* Add example for fortimanager_object_cli_templategroup
* Add example for fortimanager_object_dlp_filepattern
* Add example for fortimanager_object_dlp_sensitivity
* Add example for fortimanager_object_dlp_sensor
* Add example for fortimanager_object_dnsfilter_profile
* Add example for fortimanager_object_dynamic_address
* Add example for fortimanager_object_dynamic_certificate_local
* Add example for fortimanager_object_dynamic_interface
* Add example for fortimanager_object_dynamic_ippool
* Add example for fortimanager_object_dynamic_multicast_interface
* Add example for fortimanager_object_dynamic_vip
* Add example for fortimanager_object_dynamic_virtualwanlink_members
* Add example for fortimanager_object_dynamic_virtualwanlink_neighbor
* Add example for fortimanager_object_dynamic_virtualwanlink_server
* Add example for fortimanager_object_dynamic_vpntunnel
* Add example for fortimanager_object_emailfilter_bwl
* Add example for fortimanager_object_emailfilter_bword
* Add example for fortimanager_object_emailfilter_dnsbl
* Add example for fortimanager_object_emailfilter_iptrust
* Add example for fortimanager_object_emailfilter_mheader
* Add example for fortimanager_object_emailfilter_profile
* Add example for fortimanager_object_extendercontroller_sim_profile
* Add example for fortimanager_object_filefilter_profile
* Add example for fortimanager_object_firewall_address6
* Add example for fortimanager_object_firewall_address6template
* Add example for fortimanager_object_firewall_addrgrp
* Add example for fortimanager_object_firewall_addrgrp6
* Add example for fortimanager_object_firewall_decryptedtrafficmirror
* Add example for fortimanager_object_firewall_identitybasedroute
* Add example for fortimanager_object_firewall_internetserviceaddition
* Add example for fortimanager_object_firewall_internetservicecustom
* Add example for fortimanager_object_firewall_internetservicename
* Add example for fortimanager_object_firewall_ippool
* Add example for fortimanager_object_firewall_ippool6
* Add example for fortimanager_object_firewall_multicastaddress6
* Add example for fortimanager_object_firewall_profilegroup
* Add example for fortimanager_object_firewall_profileprotocoloptions
* Add example for fortimanager_object_firewall_proxyaddress
* Add example for fortimanager_object_firewall_proxyaddrgrp
* Add example for fortimanager_object_firewall_schedule_group
* Add example for fortimanager_object_firewall_schedule_onetime
* Add example for fortimanager_object_firewall_schedule_recurring
* Add example for fortimanager_object_firewall_service_category
* Add example for fortimanager_object_firewall_service_custom
* Add example for fortimanager_object_firewall_service_group
* Add example for fortimanager_object_firewall_shaper_peripshaper
* Add example for fortimanager_object_firewall_shaper_trafficshaper
* Add example for fortimanager_object_firewall_ssh_localca
* Add example for fortimanager_object_firewall_sslsshprofile
* Add example for fortimanager_object_firewall_vip
* Add example for fortimanager_dvm_cmd_add_device
* Add example for fortimanager_dvm_cmd_del_device
* Add example for fortimanager_dvm_cmd_update_device
* Add example for fortimanager_dvmdb_group
* Add example for fortimanager_dvmdb_revision
* Add example for fortimanager_dvmdb_script
* Add example for fortimanager_fmupdate_analyzer_virusreport
* Add example for fortimanager_fmupdate_avips_advancedlog
* Add example for fortimanager_fmupdate_avips_webproxy
* Add example for fortimanager_fmupdate_customurllist
* Add example for fortimanager_fmupdate_diskquota
* Add example for fortimanager_fmupdate_fctservices
* Add example for fortimanager_fmupdate_fdssetting
* Add example for fortimanager_fmupdate_fdssetting_pushoverride
* Add example for fortimanager_fmupdate_fdssetting_pushoverridetoclient
* Add example for fortimanager_object_firewall_vip46
* Add example for fortimanager_object_firewall_vip6
* Add example for fortimanager_object_firewall_vip64
* Add example for fortimanager_object_firewall_vipgrp
* Add example for fortimanager_object_firewall_vipgrp46
* Add example for fortimanager_object_firewall_vipgrp6
* Add example for fortimanager_object_firewall_vipgrp64
* Add example for fortimanager_object_firewall_wildcardfqdn_custom
* Add example for fortimanager_object_firewall_wildcardfqdn_group
* Add example for fortimanager_object_fsp_vlan
* Add example for fortimanager_object_icap_profile
* Add example for fortimanager_object_icap_server
* Add example for fortimanager_object_ips_custom
* Add example for fortimanager_object_ips_sensor
* Add example for fortimanager_object_log_customfield
* Add example for fortimanager_object_sshfilter_profile
* Add example for fortimanager_object_switchcontroller_lldpprofile
* Add example for fortimanager_object_switchcontroller_qos_dot1pmap
* Add example for fortimanager_object_switchcontroller_qos_qospolicy
* Add example for fortimanager_object_switchcontroller_qos_queuepolicy
* Add example for fortimanager_object_system_customlanguage
* Add example for fortimanager_object_system_dhcp_server
* Add example for fortimanager_object_system_geoipcountry
* Add example for fortimanager_object_system_meta
* Add example for fortimanager_object_system_replacemsggroup
* Add example for fortimanager_object_system_replacemsgimage
* Add example for fortimanager_object_system_sdnconnector
* Add example for fortimanager_object_system_smsserver
* Add example for fortimanager_object_user_adgrp
* Add example for fortimanager_object_user_domaincontroller
* Add example for fortimanager_object_user_fsso
* Add example for fortimanager_object_user_fssopolling
* Add example for fortimanager_object_user_group
* Add example for fortimanager_object_user_krbkeytab
* Add example for fortimanager_object_user_passwordpolicy
* Add example for fortimanager_object_user_peer
* Add example for fortimanager_object_user_pop3
* Add example for fortimanager_object_user_pxgrid
* Add example for fortimanager_object_user_saml
* Add example for fortimanager_object_user_securityexemptlist
* Add example for fortimanager_object_user_tacacs
* Add example for fortimanager_object_user_vcenter
* Add example for fortimanager_object_voip_profile
* Add example for fortimanager_object_vpn_certificate_ca
* Add example for fortimanager_object_vpn_certificate_ocspserver
* Add example for fortimanager_object_vpn_certificate_remote
* Add example for fortimanager_object_vpn_ssl_web_hostchecksoftware
* Add example for fortimanager_object_vpn_ssl_web_portal
* Add example for fortimanager_object_vpn_ssl_web_realm
* Add example for fortimanager_object_vpnmgr_node
* Add example for fortimanager_object_waf_mainclass
* Add example for fortimanager_object_waf_profile
* Add example for fortimanager_object_waf_signature
* Add example for fortimanager_object_waf_subclass
* Add example for fortimanager_object_wanopt_authgroup
* Add example for fortimanager_object_wanopt_profile
* Add example for fortimanager_object_webfilter_categories
* Add example for fortimanager_object_webfilter_content
* Add example for fortimanager_object_webfilter_contentheader
* Add example for fortimanager_object_webfilter_ftgdlocalrating
* Add example for fortimanager_object_webfilter_profile
* Add example for fortimanager_object_webfilter_urlfilter
* Add example for fortimanager_object_webproxy_forwardserver
* Add example for fortimanager_object_webproxy_forwardservergroup
* Add example for fortimanager_object_webproxy_profile
* Add example for fortimanager_object_webproxy_wisp
* Add example for fortimanager_object_wirelesscontroller_bleprofile
* Add example for fortimanager_object_wirelesscontroller_hotspot20_anqp3gppcellular
* Add example for fortimanager_object_wirelesscontroller_hotspot20_anqpipaddresstype
* Add example for fortimanager_object_wirelesscontroller_hotspot20_anqpnetworkauthtype
* Add example for fortimanager_object_wirelesscontroller_hotspot20_h2qpconncapability
* Add example for fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovider
* Add example for fortimanager_object_wirelesscontroller_hotspot20_h2qpwanmetric
* Add example for fortimanager_object_wirelesscontroller_hotspot20_hsprofile
* Add example for fortimanager_object_wirelesscontroller_mpskprofile
* Add example for fortimanager_object_wirelesscontroller_qosprofile
* Add example for fortimanager_object_wirelesscontroller_utmprofile
* Add example for fortimanager_object_wirelesscontroller_vap
* Add example for fortimanager_object_wirelesscontroller_vapgroup
* Add example for fortimanager_object_wirelesscontroller_wagprofile
* Add example for fortimanager_object_wirelesscontroller_widsprofile
* Add example for fortimanager_object_wirelesscontroller_wtpprofile
* Add example for fortimanager_packages_firewall_interfacepolicy6_move
* Add example for fortimanager_packages_firewall_interfacepolicy_move
* Add example for fortimanager_packages_firewall_localinpolicy6_move
* Add example for fortimanager_packages_firewall_localinpolicy_move
* Add example for fortimanager_packages_firewall_multicastpolicy6_move
* Add example for fortimanager_packages_firewall_multicastpolicy_move
* Add example for fortimanager_packages_firewall_policy46_move
* Add example for fortimanager_packages_firewall_policy64_move
* Add example for fortimanager_packages_firewall_policy_move
* Add example for fortimanager_packages_firewall_proxypolicy_move
* Add example for fortimanager_packages_firewall_securitypolicy_move
* Add example for fortimanager_packages_firewall_shapingpolicy_move
* Add example for fortimanager_packages_global_footer_consolidated_policy
* Add example for fortimanager_packages_global_footer_policy6
* Add example for fortimanager_packages_global_footer_shapingpolicy
* Add example for fortimanager_packages_global_header_consolidated_policy
* Add example for fortimanager_packages_global_header_policy6
* Add example for fortimanager_packages_pkg
* Add example for fortimanager_securityconsole_abort
* Add example for fortimanager_securityconsole_install_device
* Add example for fortimanager_securityconsole_install_package
* Add example for fortimanager_securityconsole_install_preview
* Add example for fortimanager_securityconsole_package_cancel_install
* Add example for fortimanager_securityconsole_package_clone
* Add example for fortimanager_securityconsole_package_commit
* Add example for fortimanager_securityconsole_package_move
* Add example for fortimanager_securityconsole_pblock_clone
* Add example for fortimanager_securityconsole_reinstall_package
* Add example for fortimanager_securityconsole_sign_certificate_template
* Add example for fortimanager_system_admin_profile
* Add example for fortimanager_system_admin_setting
* Add example for fortimanager_system_alertconsole
* Add example for fortimanager_system_alertemail
* Add example for fortimanager_system_autodelete
* Add example for fortimanager_system_autodelete_dlpfilesautodeletion
* Add example for fortimanager_system_autodelete_logautodeletion
* Add example for fortimanager_system_autodelete_quarantinefilesautodeletion
* Add example for fortimanager_system_autodelete_reportautodeletion
* Add example for fortimanager_system_backup_allsettings
* Add example for fortimanager_system_certificate_ca
* Add example for fortimanager_system_certificate_crl
* Add example for fortimanager_system_certificate_local
* Add example for fortimanager_system_certificate_oftp
* Add example for fortimanager_system_certificate_remote
* Add example for fortimanager_system_connector
* Add example for fortimanager_system_dm
* Add example for fortimanager_system_dns
* Add example for fortimanager_system_fortiview_autocache
* Add example for fortimanager_system_fortiview_setting
* Add example for fortimanager_system_global
* Add example for fortimanager_system_guiact
* Add example for fortimanager_system_ha
* Add example for fortimanager_system_locallog_disk_filter
* Add example for fortimanager_system_locallog_disk_setting
* Add example for fortimanager_system_locallog_fortianalyzer2_filter
* Add example for fortimanager_system_locallog_fortianalyzer2_setting
* Add example for fortimanager_system_locallog_fortianalyzer3_filter
* Add example for fortimanager_system_locallog_fortianalyzer3_setting
* Add example for fortimanager_system_locallog_fortianalyzer_filter
* Add example for fortimanager_system_locallog_fortianalyzer_setting
* Add example for fortimanager_system_locallog_memory_filter
* Add example for fortimanager_system_locallog_memory_setting
* Add example for fortimanager_system_locallog_setting
* Add example for fortimanager_system_locallog_syslogd2_filter
* Add example for fortimanager_system_locallog_syslogd2_setting
* Add example for fortimanager_system_locallog_syslogd3_filter
* Add example for fortimanager_system_locallog_syslogd3_setting
* Add example for fortimanager_system_locallog_syslogd_filter
* Add example for fortimanager_system_locallog_syslogd_setting
* Add example for fortimanager_system_log_alert
* Add example for fortimanager_system_log_interfacestats
* Add example for fortimanager_system_log_ioc
* Add example for fortimanager_system_log_settings_rollinganalyzer
* Add example for fortimanager_system_log_settings_rollinglocal
* Add example for fortimanager_system_log_settings_rollingregular
* Add example for fortimanager_system_logfetch_clientprofile
* Add example for fortimanager_system_logfetch_serversettings
* Add example for fortimanager_system_mail
* Add example for fortimanager_system_metadata_admins
* Add example for fortimanager_system_ntp
* Add example for fortimanager_system_passwordpolicy
* Add example for fortimanager_system_report_estbrowsetime
* Add example for fortimanager_system_report_setting
* Add example for fortimanager_system_snmp_community
* Add example for fortimanager_system_snmp_sysinfo
* Add example for fortimanager_system_snmp_user
* Add example for fortimanager_system_sql
* Add example for fortimanager_system_sql_tsindexfield
* Add example for fortimanager_system_workflow_approvalmatrix

## 1.1.0 (Jun 6, 2021)

FEATURES:

* **New Resource:** `fortimanager_packages_firewall_policy_move`
* **New Resource:** `fortimanager_packages_firewall_multicastpolicy_move`
* **New Resource:** `fortimanager_packages_firewall_interfacepolicy_move`
* **New Resource:** `fortimanager_packages_firewall_interfacepolicy6_move`
* **New Resource:** `fortimanager_packages_firewall_localinpolicy_move`
* **New Resource:** `fortimanager_packages_firewall_localinpolicy6_move`
* **New Resource:** `fortimanager_packages_firewall_policy64_move`
* **New Resource:** `fortimanager_packages_firewall_dospolicy_move`
* **New Resource:** `fortimanager_packages_firewall_policy46_move`
* **New Resource:** `fortimanager_packages_firewall_dospolicy6_move`
* **New Resource:** `fortimanager_packages_firewall_shapingpolicy_move`
* **New Resource:** `fortimanager_packages_firewall_multicastpolicy6_move`
* **New Resource:** `fortimanager_packages_firewall_centralsnatmap_move`
* **New Resource:** `fortimanager_packages_firewall_proxypolicy_move`
* **New Resource:** `fortimanager_packages_firewall_securitypolicy_move`
* **New Resource:** `fortimanager_packages_pkg`
* **New Resource:** `fortimanager_packages_firewall_policy`
* **New Resource:** `fortimanager_packages_global_header_policy`
* **New Resource:** `fortimanager_packages_global_footer_policy`
* **New Resource:** `fortimanager_packages_global_header_policy6`
* **New Resource:** `fortimanager_packages_global_footer_policy6`
* **New Resource:** `fortimanager_packages_firewall_multicastpolicy`
* **New Resource:** `fortimanager_packages_firewall_interfacepolicy`
* **New Resource:** `fortimanager_packages_firewall_interfacepolicy6`
* **New Resource:** `fortimanager_packages_firewall_localinpolicy`
* **New Resource:** `fortimanager_packages_firewall_localinpolicy6`
* **New Resource:** `fortimanager_packages_firewall_policy64`
* **New Resource:** `fortimanager_packages_firewall_dospolicy`
* **New Resource:** `fortimanager_packages_firewall_policy46`
* **New Resource:** `fortimanager_packages_firewall_dospolicy6`
* **New Resource:** `fortimanager_packages_firewall_shapingpolicy`
* **New Resource:** `fortimanager_packages_global_header_shapingpolicy`
* **New Resource:** `fortimanager_packages_global_footer_shapingpolicy`
* **New Resource:** `fortimanager_packages_firewall_multicastpolicy6`
* **New Resource:** `fortimanager_packages_firewall_centralsnatmap`
* **New Resource:** `fortimanager_packages_firewall_proxypolicy`
* **New Resource:** `fortimanager_packages_authentication_rule`
* **New Resource:** `fortimanager_packages_authentication_setting`
* **New Resource:** `fortimanager_packages_global_header_consolidated_policy`
* **New Resource:** `fortimanager_packages_global_footer_consolidated_policy`
* **New Resource:** `fortimanager_packages_firewall_securitypolicy`
* **New Resource:** `fortimanager_dvmdb_script_execute`
* **New Resource:** `fortimanager_dvm_cmd_add_device`
* **New Resource:** `fortimanager_dvm_cmd_update_device`
* **New Resource:** `fortimanager_dvm_cmd_del_device`
* **New Resource:** `fortimanager_securityconsole_abort`
* **New Resource:** `fortimanager_securityconsole_install_preview`
* **New Resource:** `fortimanager_securityconsole_install_package`
* **New Resource:** `fortimanager_securityconsole_install_device`
* **New Resource:** `fortimanager_securityconsole_package_commit`
* **New Resource:** `fortimanager_securityconsole_package_cancel_install`
* **New Resource:** `fortimanager_securityconsole_package_clone`
* **New Resource:** `fortimanager_securityconsole_package_move`
* **New Resource:** `fortimanager_securityconsole_reinstall_package`
* **New Resource:** `fortimanager_securityconsole_sign_certificate_template`
* **New Resource:** `fortimanager_securityconsole_pblock_clone`
* **New Resource:** `fortimanager_object_webfilter_ftgdlocalcat`
* **New Resource:** `fortimanager_object_webfilter_urlfilter`
* **New Resource:** `fortimanager_object_webfilter_ftgdlocalrating`
* **New Resource:** `fortimanager_object_vpn_certificate_ca`
* **New Resource:** `fortimanager_object_vpn_certificate_remote`
* **New Resource:** `fortimanager_object_system_fortiguard`
* **New Resource:** `fortimanager_object_ips_custom`
* **New Resource:** `fortimanager_object_system_dhcp_server`
* **New Resource:** `fortimanager_object_firewall_address`
* **New Resource:** `fortimanager_object_firewall_address6`
* **New Resource:** `fortimanager_object_firewall_addrgrp`
* **New Resource:** `fortimanager_object_user_adgrp`
* **New Resource:** `fortimanager_object_user_radius`
* **New Resource:** `fortimanager_object_user_ldap`
* **New Resource:** `fortimanager_object_user_local`
* **New Resource:** `fortimanager_object_user_peer`
* **New Resource:** `fortimanager_object_user_peergrp`
* **New Resource:** `fortimanager_object_user_group`
* **New Resource:** `fortimanager_object_firewall_service_custom`
* **New Resource:** `fortimanager_object_firewall_service_group`
* **New Resource:** `fortimanager_object_firewall_schedule_onetime`
* **New Resource:** `fortimanager_object_firewall_schedule_recurring`
* **New Resource:** `fortimanager_object_firewall_ippool`
* **New Resource:** `fortimanager_object_firewall_vip`
* **New Resource:** `fortimanager_object_firewall_vipgrp`
* **New Resource:** `fortimanager_object_firewall_addrgrp6`
* **New Resource:** `fortimanager_object_ips_sensor`
* **New Resource:** `fortimanager_object_log_customfield`
* **New Resource:** `fortimanager_object_user_tacacs`
* **New Resource:** `fortimanager_object_firewall_ldbmonitor`
* **New Resource:** `fortimanager_object_application_list`
* **New Resource:** `fortimanager_object_dlp_sensor`
* **New Resource:** `fortimanager_object_wanopt_peer`
* **New Resource:** `fortimanager_object_wanopt_authgroup`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal`
* **New Resource:** `fortimanager_object_system_replacemsgimage`
* **New Resource:** `fortimanager_object_system_replacemsggroup`
* **New Resource:** `fortimanager_object_webfilter_content`
* **New Resource:** `fortimanager_object_webfilter_contentheader`
* **New Resource:** `fortimanager_object_firewall_schedule_group`
* **New Resource:** `fortimanager_object_firewall_shaper_trafficshaper`
* **New Resource:** `fortimanager_object_firewall_shaper_peripshaper`
* **New Resource:** `fortimanager_object_vpn_ssl_web_hostchecksoftware`
* **New Resource:** `fortimanager_object_wirelesscontroller_vap`
* **New Resource:** `fortimanager_object_wirelesscontroller_vapgroup`
* **New Resource:** `fortimanager_object_webfilter_profile`
* **New Resource:** `fortimanager_object_antivirus_profile`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions`
* **New Resource:** `fortimanager_object_firewall_profilegroup`
* **New Resource:** `fortimanager_object_voip_profile`
* **New Resource:** `fortimanager_object_user_fortitoken`
* **New Resource:** `fortimanager_object_webproxy_forwardserver`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile`
* **New Resource:** `fortimanager_object_dlp_filepattern`
* **New Resource:** `fortimanager_object_icap_server`
* **New Resource:** `fortimanager_object_icap_profile`
* **New Resource:** `fortimanager_object_user_fsso`
* **New Resource:** `fortimanager_object_system_smsserver`
* **New Resource:** `fortimanager_object_vpn_certificate_ocspserver`
* **New Resource:** `fortimanager_object_user_passwordpolicy`
* **New Resource:** `fortimanager_object_user_fssopolling`
* **New Resource:** `fortimanager_object_firewall_ippool6`
* **New Resource:** `fortimanager_object_firewall_vip6`
* **New Resource:** `fortimanager_object_firewall_vipgrp6`
* **New Resource:** `fortimanager_object_firewall_identitybasedroute`
* **New Resource:** `fortimanager_object_wirelesscontroller_widsprofile`
* **New Resource:** `fortimanager_object_wanopt_profile`
* **New Resource:** `fortimanager_object_firewall_multicastaddress`
* **New Resource:** `fortimanager_object_firewall_service_category`
* **New Resource:** `fortimanager_object_system_geoipoverride`
* **New Resource:** `fortimanager_object_application_custom`
* **New Resource:** `fortimanager_object_vpn_ssl_web_realm`
* **New Resource:** `fortimanager_object_firewall_vip64`
* **New Resource:** `fortimanager_object_firewall_vip46`
* **New Resource:** `fortimanager_object_firewall_vipgrp64`
* **New Resource:** `fortimanager_object_firewall_vipgrp46`
* **New Resource:** `fortimanager_object_webproxy_forwardservergroup`
* **New Resource:** `fortimanager_object_user_pop3`
* **New Resource:** `fortimanager_object_webproxy_profile`
* **New Resource:** `fortimanager_object_extendercontroller_sim_profile`
* **New Resource:** `fortimanager_object_user_securityexemptlist`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile`
* **New Resource:** `fortimanager_object_system_customlanguage`
* **New Resource:** `fortimanager_object_system_geoipcountry`
* **New Resource:** `fortimanager_object_webfilter_categories`
* **New Resource:** `fortimanager_object_system_virtualwirepair`
* **New Resource:** `fortimanager_object_firewall_multicastaddress6`
* **New Resource:** `fortimanager_object_waf_mainclass`
* **New Resource:** `fortimanager_object_waf_subclass`
* **New Resource:** `fortimanager_object_waf_signature`
* **New Resource:** `fortimanager_object_waf_profile`
* **New Resource:** `fortimanager_object_dnsfilter_profile`
* **New Resource:** `fortimanager_object_webproxy_wisp`
* **New Resource:** `fortimanager_object_user_krbkeytab`
* **New Resource:** `fortimanager_object_application_categories`
* **New Resource:** `fortimanager_object_application_group`
* **New Resource:** `fortimanager_object_dnsfilter_domainfilter`
* **New Resource:** `fortimanager_object_firewall_internetservice`
* **New Resource:** `fortimanager_object_firewall_internetservicecustom`
* **New Resource:** `fortimanager_object_firewall_proxyaddress`
* **New Resource:** `fortimanager_object_firewall_proxyaddrgrp`
* **New Resource:** `fortimanager_object_switchcontroller_securitypolicy_8021x`
* **New Resource:** `fortimanager_object_switchcontroller_lldpprofile`
* **New Resource:** `fortimanager_object_wirelesscontroller_qosprofile`
* **New Resource:** `fortimanager_object_authentication_scheme`
* **New Resource:** `fortimanager_object_wirelesscontroller_bleprofile`
* **New Resource:** `fortimanager_object_switchcontroller_qos_dot1pmap`
* **New Resource:** `fortimanager_object_switchcontroller_qos_ipdscpmap`
* **New Resource:** `fortimanager_object_switchcontroller_qos_queuepolicy`
* **New Resource:** `fortimanager_object_switchcontroller_qos_qospolicy`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqpvenuename`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqpnetworkauthtype`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqproamingconsortium`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqp3gppcellular`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqpipaddresstype`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qpoperatorname`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qpwanmetric`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qpconncapability`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovider`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_qosmap`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_hsprofile`
* **New Resource:** `fortimanager_object_system_sdnconnector`
* **New Resource:** `fortimanager_object_user_domaincontroller`
* **New Resource:** `fortimanager_object_system_objecttagging`
* **New Resource:** `fortimanager_object_firewall_address6template`
* **New Resource:** `fortimanager_object_firewall_wildcardfqdn_custom`
* **New Resource:** `fortimanager_object_firewall_wildcardfqdn_group`
* **New Resource:** `fortimanager_object_firewall_internetservicegroup`
* **New Resource:** `fortimanager_object_firewall_internetservicecustomgroup`
* **New Resource:** `fortimanager_object_system_externalresource`
* **New Resource:** `fortimanager_object_sshfilter_profile`
* **New Resource:** `fortimanager_object_wirelesscontroller_utmprofile`
* **New Resource:** `fortimanager_object_firewall_ssh_localca`
* **New Resource:** `fortimanager_object_dlp_sensitivity`
* **New Resource:** `fortimanager_object_emailfilter_bword`
* **New Resource:** `fortimanager_object_emailfilter_bwl`
* **New Resource:** `fortimanager_object_emailfilter_mheader`
* **New Resource:** `fortimanager_object_emailfilter_dnsbl`
* **New Resource:** `fortimanager_object_emailfilter_iptrust`
* **New Resource:** `fortimanager_object_emailfilter_profile`
* **New Resource:** `fortimanager_object_emailfilter_fortishield`
* **New Resource:** `fortimanager_object_emailfilter_options`
* **New Resource:** `fortimanager_object_user_exchange`
* **New Resource:** `fortimanager_object_firewall_internetserviceaddition`
* **New Resource:** `fortimanager_object_cifs_profile`
* **New Resource:** `fortimanager_object_user_saml`
* **New Resource:** `fortimanager_object_firewall_trafficclass`
* **New Resource:** `fortimanager_object_wirelesscontroller_wagprofile`
* **New Resource:** `fortimanager_object_credentialstore_domaincontroller`
* **New Resource:** `fortimanager_object_firewall_internetservicename`
* **New Resource:** `fortimanager_object_firewall_decryptedtrafficmirror`
* **New Resource:** `fortimanager_object_filefilter_profile`
* **New Resource:** `fortimanager_object_extendercontroller_dataplan`
* **New Resource:** `fortimanager_object_wirelesscontroller_mpskprofile`
* **New Resource:** `fortimanager_object_dynamic_multicast_interface`
* **New Resource:** `fortimanager_object_dynamic_interface`
* **New Resource:** `fortimanager_object_dynamic_address`
* **New Resource:** `fortimanager_object_vpnmgr_node`
* **New Resource:** `fortimanager_object_system_meta`
* **New Resource:** `fortimanager_object_adom_options`
* **New Resource:** `fortimanager_object_dynamic_vip`
* **New Resource:** `fortimanager_object_dynamic_ippool`
* **New Resource:** `fortimanager_object_dynamic_certificate_local`
* **New Resource:** `fortimanager_object_dynamic_vpntunnel`
* **New Resource:** `fortimanager_object_certificate_template`
* **New Resource:** `fortimanager_object_dynamic_virtualwanlink_members`
* **New Resource:** `fortimanager_object_dynamic_virtualwanlink_server`
* **New Resource:** `fortimanager_object_dynamic_virtualwanlink_neighbor`
* **New Resource:** `fortimanager_object_user_pxgrid`
* **New Resource:** `fortimanager_object_user_clearpass`
* **New Resource:** `fortimanager_object_user_nsx`
* **New Resource:** `fortimanager_object_user_vcenter`
* **New Resource:** `fortimanager_object_cli_template`
* **New Resource:** `fortimanager_object_cli_templategroup`
* **New Resource:** `fortimanager_object_fsp_vlan`
* **New Resource:** `fortimanager_system_global`
* **New Resource:** `fortimanager_system_interface`
* **New Resource:** `fortimanager_system_snmp_sysinfo`
* **New Resource:** `fortimanager_system_snmp_community`
* **New Resource:** `fortimanager_system_snmp_user`
* **New Resource:** `fortimanager_system_route`
* **New Resource:** `fortimanager_system_route6`
* **New Resource:** `fortimanager_system_dns`
* **New Resource:** `fortimanager_system_connector`
* **New Resource:** `fortimanager_system_ha`
* **New Resource:** `fortimanager_system_ha_peer`
* **New Resource:** `fortimanager_system_certificate_ca`
* **New Resource:** `fortimanager_system_certificate_local`
* **New Resource:** `fortimanager_system_certificate_crl`
* **New Resource:** `fortimanager_system_certificate_remote`
* **New Resource:** `fortimanager_system_certificate_oftp`
* **New Resource:** `fortimanager_system_certificate_ssh`
* **New Resource:** `fortimanager_system_saml`
* **New Resource:** `fortimanager_system_saml_fabricidp`
* **New Resource:** `fortimanager_system_saml_serviceproviders`
* **New Resource:** `fortimanager_system_ntp`
* **New Resource:** `fortimanager_system_ntp_ntpserver`
* **New Resource:** `fortimanager_system_backup_allsettings`
* **New Resource:** `fortimanager_system_guiact`
* **New Resource:** `fortimanager_system_metadata_admins`
* **New Resource:** `fortimanager_system_admin_profile`
* **New Resource:** `fortimanager_system_admin_radius`
* **New Resource:** `fortimanager_system_admin_ldap`
* **New Resource:** `fortimanager_system_admin_tacacs`
* **New Resource:** `fortimanager_system_admin_group`
* **New Resource:** `fortimanager_system_admin_user`
* **New Resource:** `fortimanager_system_admin_setting`
* **New Resource:** `fortimanager_system_passwordpolicy`
* **New Resource:** `fortimanager_system_alertemail`
* **New Resource:** `fortimanager_system_syslog`
* **New Resource:** `fortimanager_system_mail`
* **New Resource:** `fortimanager_system_alertevent`
* **New Resource:** `fortimanager_system_alertconsole`
* **New Resource:** `fortimanager_system_locallog_disk_setting`
* **New Resource:** `fortimanager_system_locallog_disk_filter`
* **New Resource:** `fortimanager_system_locallog_memory_setting`
* **New Resource:** `fortimanager_system_locallog_memory_filter`
* **New Resource:** `fortimanager_system_locallog_fortianalyzer_filter`
* **New Resource:** `fortimanager_system_locallog_fortianalyzer_setting`
* **New Resource:** `fortimanager_system_locallog_fortianalyzer2_filter`
* **New Resource:** `fortimanager_system_locallog_fortianalyzer2_setting`
* **New Resource:** `fortimanager_system_locallog_fortianalyzer3_filter`
* **New Resource:** `fortimanager_system_locallog_fortianalyzer3_setting`
* **New Resource:** `fortimanager_system_locallog_syslogd_setting`
* **New Resource:** `fortimanager_system_locallog_syslogd_filter`
* **New Resource:** `fortimanager_system_locallog_syslogd2_setting`
* **New Resource:** `fortimanager_system_locallog_syslogd2_filter`
* **New Resource:** `fortimanager_system_locallog_syslogd3_setting`
* **New Resource:** `fortimanager_system_locallog_syslogd3_filter`
* **New Resource:** `fortimanager_system_locallog_setting`
* **New Resource:** `fortimanager_system_fips`
* **New Resource:** `fortimanager_system_workflow_approvalmatrix`
* **New Resource:** `fortimanager_system_docker`
* **New Resource:** `fortimanager_system_sniffer`
* **New Resource:** `fortimanager_system_dm`
* **New Resource:** `fortimanager_system_log_alert`
* **New Resource:** `fortimanager_system_log_interfacestats`
* **New Resource:** `fortimanager_system_log_ioc`
* **New Resource:** `fortimanager_system_log_maildomain`
* **New Resource:** `fortimanager_system_log_devicedisable`
* **New Resource:** `fortimanager_system_log_settings`
* **New Resource:** `fortimanager_system_log_settings_rollinganalyzer`
* **New Resource:** `fortimanager_system_log_settings_rollinglocal`
* **New Resource:** `fortimanager_system_log_settings_rollingregular`
* **New Resource:** `fortimanager_system_logfetch_serversettings`
* **New Resource:** `fortimanager_system_logfetch_clientprofile`
* **New Resource:** `fortimanager_system_sql`
* **New Resource:** `fortimanager_system_sql_tsindexfield`
* **New Resource:** `fortimanager_system_sql_customskipidx`
* **New Resource:** `fortimanager_system_sql_customindex`
* **New Resource:** `fortimanager_system_report_estbrowsetime`
* **New Resource:** `fortimanager_system_report_autocache`
* **New Resource:** `fortimanager_system_report_setting`
* **New Resource:** `fortimanager_system_fortiview_setting`
* **New Resource:** `fortimanager_system_fortiview_autocache`
* **New Resource:** `fortimanager_system_autodelete`
* **New Resource:** `fortimanager_system_autodelete_quarantinefilesautodeletion`
* **New Resource:** `fortimanager_system_autodelete_dlpfilesautodeletion`
* **New Resource:** `fortimanager_system_autodelete_reportautodeletion`
* **New Resource:** `fortimanager_system_autodelete_logautodeletion`
* **New Resource:** `fortimanager_fmupdate_avips_webproxy`
* **New Resource:** `fortimanager_fmupdate_avips_advancedlog`
* **New Resource:** `fortimanager_fmupdate_webspam_webproxy`
* **New Resource:** `fortimanager_fmupdate_webspam_fgdsetting`
* **New Resource:** `fortimanager_fmupdate_fctservices`
* **New Resource:** `fortimanager_fmupdate_analyzer_virusreport`
* **New Resource:** `fortimanager_fmupdate_service`
* **New Resource:** `fortimanager_fmupdate_publicnetwork`
* **New Resource:** `fortimanager_fmupdate_diskquota`
* **New Resource:** `fortimanager_fmupdate_serveraccesspriorities`
* **New Resource:** `fortimanager_fmupdate_customurllist`
* **New Resource:** `fortimanager_fmupdate_serveroverridestatus`
* **New Resource:** `fortimanager_fmupdate_multilayer`
* **New Resource:** `fortimanager_fmupdate_fdssetting`
* **New Resource:** `fortimanager_fmupdate_fdssetting_serveroverride`
* **New Resource:** `fortimanager_fmupdate_fdssetting_updateschedule`
* **New Resource:** `fortimanager_fmupdate_fdssetting_pushoverridetoclient`
* **New Resource:** `fortimanager_fmupdate_fdssetting_pushoverride`
* **New Resource:** `fortimanager_fmupdate_fwmsetting`
* **New Resource:** `fortimanager_dvmdb_group`
* **New Resource:** `fortimanager_dvmdb_script`
* **New Resource:** `fortimanager_dvmdb_revision`
* **New Resource:** `fortimanager_dvmdb_adom`
