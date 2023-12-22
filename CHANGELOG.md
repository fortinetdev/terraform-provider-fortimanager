## 1.11.0 (Unreleased)

## 1.10.0 (Dec 22, 2023)

IMPROVEMENTS:
* Support FortiManager v7.0.10;
* Add variable `clean_session` to support logout session on disabled workspace mode;
* Add resources for central dnat;
* Fix issue when setting block variable to empty block;

FEATURES:
* **New Resource:** `fortimanager_packages_central_dnat`
* **New Resource:** `fortimanager_packages_central_dnat6`
* **New Resource:** `fortimanager_packages_central_dnat6_move`
* **New Resource:** `fortimanager_packages_central_dnat_move`

## 1.9.0 (Oct 13, 2023)

IMPROVEMENTS:

* Support FortiManager Cloud;
* Support FMG v6.4.13, 7.0.9, 7.2.4, 7.4.1; 
* Support System template resources;
* Fix type convert issues;
* Add pkg_folder_path for package resources; 
* Add output attributes for fortimanager_packages_pkg;
* Fix issue of set content to state file even not been created;

FEATURES:

* **New Resource:** `fortimanager_dvm_cmd_add_devlist`
* **New Resource:** `fortimanager_dvm_cmd_del_devlist`
* **New Resource:** `fortimanager_dvm_cmd_update_devlist`
* **New Resource:** `fortimanager_object_antivirus_profile_cifs`
* **New Resource:** `fortimanager_object_antivirus_profile_contentdisarm`
* **New Resource:** `fortimanager_object_antivirus_profile_ftp`
* **New Resource:** `fortimanager_object_antivirus_profile_http`
* **New Resource:** `fortimanager_object_antivirus_profile_imap`
* **New Resource:** `fortimanager_object_antivirus_profile_mapi`
* **New Resource:** `fortimanager_object_antivirus_profile_nacquar`
* **New Resource:** `fortimanager_object_antivirus_profile_nntp`
* **New Resource:** `fortimanager_object_antivirus_profile_outbreakprevention`
* **New Resource:** `fortimanager_object_antivirus_profile_pop3`
* **New Resource:** `fortimanager_object_antivirus_profile_smtp`
* **New Resource:** `fortimanager_object_antivirus_profile_ssh`
* **New Resource:** `fortimanager_object_application_list_defaultnetworkservices`
* **New Resource:** `fortimanager_object_application_list_defaultnetworkservices_move`
* **New Resource:** `fortimanager_object_application_list_entries`
* **New Resource:** `fortimanager_object_application_list_entries_parameters`
* **New Resource:** `fortimanager_object_application_list_entries_parameters_members`
* **New Resource:** `fortimanager_object_application_list_entries_parameters_move`
* **New Resource:** `fortimanager_object_casb_profile`
* **New Resource:** `fortimanager_object_casb_profile_move`
* **New Resource:** `fortimanager_object_casb_profile_saasapplication`
* **New Resource:** `fortimanager_object_casb_profile_saasapplication_accessrule`
* **New Resource:** `fortimanager_object_casb_profile_saasapplication_customcontrol`
* **New Resource:** `fortimanager_object_casb_profile_saasapplication_customcontrol_option`
* **New Resource:** `fortimanager_object_casb_saasapplication`
* **New Resource:** `fortimanager_object_casb_saasapplication_move`
* **New Resource:** `fortimanager_object_casb_useractivity`
* **New Resource:** `fortimanager_object_casb_useractivity_controloptions`
* **New Resource:** `fortimanager_object_casb_useractivity_controloptions_operations`
* **New Resource:** `fortimanager_object_casb_useractivity_match`
* **New Resource:** `fortimanager_object_casb_useractivity_match_rules`
* **New Resource:** `fortimanager_object_casb_useractivity_move`
* **New Resource:** `fortimanager_object_dlp_dictionary_entries`
* **New Resource:** `fortimanager_object_dlp_filepattern_entries`
* **New Resource:** `fortimanager_object_dlp_profile_rule`
* **New Resource:** `fortimanager_object_dlp_profile_rule_move`
* **New Resource:** `fortimanager_object_dlp_sensor_entries`
* **New Resource:** `fortimanager_object_dnsfilter_domainfilter_entries`
* **New Resource:** `fortimanager_object_dnsfilter_domainfilter_entries_move`
* **New Resource:** `fortimanager_object_dnsfilter_profile_dnstranslation`
* **New Resource:** `fortimanager_object_dnsfilter_profile_domainfilter`
* **New Resource:** `fortimanager_object_dnsfilter_profile_ftgddns`
* **New Resource:** `fortimanager_object_dnsfilter_profile_ftgddns_filters`
* **New Resource:** `fortimanager_object_dynamic_address_dynamic_addr_mapping`
* **New Resource:** `fortimanager_object_dynamic_interface_platform_mapping`
* **New Resource:** `fortimanager_object_emailfilter_blockallowlist_entries`
* **New Resource:** `fortimanager_object_emailfilter_blockallowlist_entries_move`
* **New Resource:** `fortimanager_object_emailfilter_bwl_entries`
* **New Resource:** `fortimanager_object_emailfilter_bwl_entries_move`
* **New Resource:** `fortimanager_object_emailfilter_bword_entries`
* **New Resource:** `fortimanager_object_emailfilter_dnsbl_entries`
* **New Resource:** `fortimanager_object_emailfilter_iptrust_entries`
* **New Resource:** `fortimanager_object_emailfilter_mheader_entries`
* **New Resource:** `fortimanager_object_emailfilter_profile_gmail`
* **New Resource:** `fortimanager_object_emailfilter_profile_imap`
* **New Resource:** `fortimanager_object_emailfilter_profile_mapi`
* **New Resource:** `fortimanager_object_emailfilter_profile_msnhotmail`
* **New Resource:** `fortimanager_object_emailfilter_profile_pop3`
* **New Resource:** `fortimanager_object_emailfilter_profile_smtp`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_cellular_controllerreport`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_cellular_modem1`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_cellular_modem1_autoswitch`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_cellular_modem2`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_cellular_modem2_autoswitch`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_alert`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver_move`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_lanextension`
* **New Resource:** `fortimanager_object_extendercontroller_extenderprofile_lanextension_backhaul`
* **New Resource:** `fortimanager_object_extendercontroller_sim_profile_autoswitch_profile`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_cellular`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_cellular_controllerreport`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_cellular_modem1`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_cellular_modem1_autoswitch`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_cellular_modem2`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_cellular_modem2_autoswitch`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_cellular_smsnotification`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_cellular_smsnotification_alert`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_cellular_smsnotification_receiver`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_cellular_smsnotification_receiver_move`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_lanextension`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_lanextension_backhaul`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile_lanextension_backhaul_move`
* **New Resource:** `fortimanager_object_filefilter_profile_rules`
* **New Resource:** `fortimanager_object_filefilter_profile_rules_move`
* **New Resource:** `fortimanager_object_firewall_accessproxy6_apigateway`
* **New Resource:** `fortimanager_object_firewall_accessproxy6_apigateway6`
* **New Resource:** `fortimanager_object_firewall_accessproxy6_apigateway6_quic`
* **New Resource:** `fortimanager_object_firewall_accessproxy6_apigateway6_realservers`
* **New Resource:** `fortimanager_object_firewall_accessproxy6_apigateway6_sslciphersuites`
* **New Resource:** `fortimanager_object_firewall_accessproxy6_apigateway_quic`
* **New Resource:** `fortimanager_object_firewall_accessproxy6_apigateway_realservers`
* **New Resource:** `fortimanager_object_firewall_accessproxy6_apigateway_sslciphersuites`
* **New Resource:** `fortimanager_object_firewall_accessproxy_apigateway`
* **New Resource:** `fortimanager_object_firewall_accessproxy_apigateway6`
* **New Resource:** `fortimanager_object_firewall_accessproxy_apigateway6_quic`
* **New Resource:** `fortimanager_object_firewall_accessproxy_apigateway6_realservers`
* **New Resource:** `fortimanager_object_firewall_accessproxy_apigateway6_sslciphersuites`
* **New Resource:** `fortimanager_object_firewall_accessproxy_apigateway_quic`
* **New Resource:** `fortimanager_object_firewall_accessproxy_apigateway_realservers`
* **New Resource:** `fortimanager_object_firewall_accessproxy_apigateway_sslciphersuites`
* **New Resource:** `fortimanager_object_firewall_accessproxy_realservers`
* **New Resource:** `fortimanager_object_firewall_accessproxy_serverpubkeyauthsettings`
* **New Resource:** `fortimanager_object_firewall_address6template_subnetsegment`
* **New Resource:** `fortimanager_object_firewall_address6template_subnetsegment_values`
* **New Resource:** `fortimanager_object_firewall_address6_list`
* **New Resource:** `fortimanager_object_firewall_address6_subnetsegment`
* **New Resource:** `fortimanager_object_firewall_address6_tagging`
* **New Resource:** `fortimanager_object_firewall_address_list`
* **New Resource:** `fortimanager_object_firewall_address_tagging`
* **New Resource:** `fortimanager_object_firewall_addrgrp6_tagging`
* **New Resource:** `fortimanager_object_firewall_addrgrp_tagging`
* **New Resource:** `fortimanager_object_firewall_casbprofile`
* **New Resource:** `fortimanager_object_firewall_casbprofile_move`
* **New Resource:** `fortimanager_object_firewall_casbprofile_saasapplication`
* **New Resource:** `fortimanager_object_firewall_casbprofile_saasapplication_accessrule`
* **New Resource:** `fortimanager_object_firewall_casbprofile_saasapplication_customcontrol`
* **New Resource:** `fortimanager_object_firewall_casbprofile_saasapplication_customcontrol_option`
* **New Resource:** `fortimanager_object_firewall_gtp`
* **New Resource:** `fortimanager_object_firewall_gtp_apn`
* **New Resource:** `fortimanager_object_firewall_gtp_ieremovepolicy`
* **New Resource:** `fortimanager_object_firewall_gtp_ievalidation`
* **New Resource:** `fortimanager_object_firewall_gtp_imsi`
* **New Resource:** `fortimanager_object_firewall_gtp_ippolicy`
* **New Resource:** `fortimanager_object_firewall_gtp_messageratelimit`
* **New Resource:** `fortimanager_object_firewall_gtp_messageratelimitv0`
* **New Resource:** `fortimanager_object_firewall_gtp_messageratelimitv1`
* **New Resource:** `fortimanager_object_firewall_gtp_messageratelimitv2`
* **New Resource:** `fortimanager_object_firewall_gtp_noippolicy`
* **New Resource:** `fortimanager_object_firewall_gtp_perapnshaper`
* **New Resource:** `fortimanager_object_firewall_gtp_policy`
* **New Resource:** `fortimanager_object_firewall_gtp_policyv2`
* **New Resource:** `fortimanager_object_firewall_identitybasedroute_rule`
* **New Resource:** `fortimanager_object_firewall_internetserviceaddition_entry`
* **New Resource:** `fortimanager_object_firewall_internetserviceaddition_entry_portrange`
* **New Resource:** `fortimanager_object_firewall_internetservicecustom_entry`
* **New Resource:** `fortimanager_object_firewall_internetservicecustom_entry_portrange`
* **New Resource:** `fortimanager_object_firewall_multicastaddress6_tagging`
* **New Resource:** `fortimanager_object_firewall_multicastaddress_tagging`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_cifs`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_cifs_filefilter_entries`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_cifs_filefilter_entries_move`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_cifs_serverkeytab`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_dns`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_ftp`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_http`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_imap`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_mailsignature`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_mapi`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_nntp`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_pop3`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_smtp`
* **New Resource:** `fortimanager_object_firewall_profileprotocoloptions_ssh`
* **New Resource:** `fortimanager_object_firewall_proxyaddress_headergroup`
* **New Resource:** `fortimanager_object_firewall_proxyaddress_tagging`
* **New Resource:** `fortimanager_object_firewall_proxyaddrgrp_tagging`
* **New Resource:** `fortimanager_object_firewall_shapingprofile`
* **New Resource:** `fortimanager_object_firewall_shapingprofile_shapingentries`
* **New Resource:** `fortimanager_object_firewall_shapingprofile_shapingentries_move`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile_dot`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile_ftps`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile_https`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile_imaps`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile_pop3s`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile_smtps`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile_ssh`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile_ssl`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile_sslexempt`
* **New Resource:** `fortimanager_object_firewall_sslsshprofile_sslserver`
* **New Resource:** `fortimanager_object_firewall_vendormac`
* **New Resource:** `fortimanager_object_firewall_vip46_realservers`
* **New Resource:** `fortimanager_object_firewall_vip64_realservers`
* **New Resource:** `fortimanager_object_firewall_vip6_realservers`
* **New Resource:** `fortimanager_object_firewall_vip_quic`
* **New Resource:** `fortimanager_object_fsp_vlan_interface`
* **New Resource:** `fortimanager_object_fsp_vlan_interface_ipv6`
* **New Resource:** `fortimanager_object_fsp_vlan_interface_ipv6_ip6delegatedprefixlist`
* **New Resource:** `fortimanager_object_fsp_vlan_interface_ipv6_ip6extraaddr`
* **New Resource:** `fortimanager_object_fsp_vlan_interface_ipv6_vrrp6`
* **New Resource:** `fortimanager_object_fsp_vlan_interface_secondaryip`
* **New Resource:** `fortimanager_object_fsp_vlan_interface_vrrp`
* **New Resource:** `fortimanager_object_global_ips_sensor_entries`
* **New Resource:** `fortimanager_object_global_ips_sensor_entries_exemptip`
* **New Resource:** `fortimanager_object_global_ips_sensor_override`
* **New Resource:** `fortimanager_object_icap_profile_respmodforwardrules`
* **New Resource:** `fortimanager_object_icap_profile_respmodforwardrules_headergroup`
* **New Resource:** `fortimanager_object_ips_sensor_entries`
* **New Resource:** `fortimanager_object_ips_sensor_entries_exemptip`
* **New Resource:** `fortimanager_object_ips_sensor_entries_move`
* **New Resource:** `fortimanager_object_router_accesslist6_rule`
* **New Resource:** `fortimanager_object_router_accesslist_rule`
* **New Resource:** `fortimanager_object_router_aspathlist_rule`
* **New Resource:** `fortimanager_object_router_communitylist_rule`
* **New Resource:** `fortimanager_object_router_prefixlist6_rule`
* **New Resource:** `fortimanager_object_router_prefixlist_rule`
* **New Resource:** `fortimanager_object_router_routemap_rule`
* **New Resource:** `fortimanager_object_sshfilter_profile_shellcommands`
* **New Resource:** `fortimanager_object_sshfilter_profile_shellcommands_move`
* **New Resource:** `fortimanager_object_switchcontroller_acl_ingress_action`
* **New Resource:** `fortimanager_object_switchcontroller_acl_ingress_classifier`
* **New Resource:** `fortimanager_object_switchcontroller_dynamicportpolicy_policy`
* **New Resource:** `fortimanager_object_switchcontroller_dynamicportpolicy_policy_move`
* **New Resource:** `fortimanager_object_switchcontroller_fortilinksettings_nacports`
* **New Resource:** `fortimanager_object_switchcontroller_lldpprofile_customtlvs`
* **New Resource:** `fortimanager_object_switchcontroller_managedswitch`
* **New Resource:** `fortimanager_object_switchcontroller_managedswitch_customcommand`
* **New Resource:** `fortimanager_object_switchcontroller_managedswitch_dhcpsnoopingstaticclient`
* **New Resource:** `fortimanager_object_switchcontroller_managedswitch_ports`
* **New Resource:** `fortimanager_object_switchcontroller_ptp_profile`
* **New Resource:** `fortimanager_object_switchcontroller_qos_ipdscpmap_map`
* **New Resource:** `fortimanager_object_system_dhcp_server_excluderange`
* **New Resource:** `fortimanager_object_system_dhcp_server_iprange`
* **New Resource:** `fortimanager_object_system_dhcp_server_options`
* **New Resource:** `fortimanager_object_system_dhcp_server_reservedaddress`
* **New Resource:** `fortimanager_object_system_geoipoverride_iprange`
* **New Resource:** `fortimanager_object_system_geoipoverride_ip6range`
* **New Resource:** `fortimanager_object_system_meta_sys_meta_fields`
* **New Resource:** `fortimanager_object_system_replacemsggroup_admin`
* **New Resource:** `fortimanager_object_system_replacemsggroup_alertmail`
* **New Resource:** `fortimanager_object_system_replacemsggroup_auth`
* **New Resource:** `fortimanager_object_system_replacemsggroup_automation`
* **New Resource:** `fortimanager_object_system_replacemsggroup_custommessage`
* **New Resource:** `fortimanager_object_system_replacemsggroup_devicedetectionportal`
* **New Resource:** `fortimanager_object_system_replacemsggroup_fortiguardwf`
* **New Resource:** `fortimanager_object_system_replacemsggroup_ftp`
* **New Resource:** `fortimanager_object_system_replacemsggroup_http`
* **New Resource:** `fortimanager_object_system_replacemsggroup_icap`
* **New Resource:** `fortimanager_object_system_replacemsggroup_mail`
* **New Resource:** `fortimanager_object_system_replacemsggroup_nacquar`
* **New Resource:** `fortimanager_object_system_replacemsggroup_nntp`
* **New Resource:** `fortimanager_object_system_replacemsggroup_spam`
* **New Resource:** `fortimanager_object_system_replacemsggroup_sslvpn`
* **New Resource:** `fortimanager_object_system_replacemsggroup_trafficquota`
* **New Resource:** `fortimanager_object_system_replacemsggroup_utm`
* **New Resource:** `fortimanager_object_system_replacemsggroup_webproxy`
* **New Resource:** `fortimanager_object_system_sdnconnector_compartmentlist`
* **New Resource:** `fortimanager_object_system_sdnconnector_compartmentlist_move`
* **New Resource:** `fortimanager_object_system_sdnconnector_externalaccountlist`
* **New Resource:** `fortimanager_object_system_sdnconnector_externalaccountlist_move`
* **New Resource:** `fortimanager_object_system_sdnconnector_externalip`
* **New Resource:** `fortimanager_object_system_sdnconnector_externalip_move`
* **New Resource:** `fortimanager_object_system_sdnconnector_forwardingrule`
* **New Resource:** `fortimanager_object_system_sdnconnector_forwardingrule_move`
* **New Resource:** `fortimanager_object_system_sdnconnector_gcpprojectlist`
* **New Resource:** `fortimanager_object_system_sdnconnector_gcpprojectlist_move`
* **New Resource:** `fortimanager_object_system_sdnconnector_nic`
* **New Resource:** `fortimanager_object_system_sdnconnector_nic_ip`
* **New Resource:** `fortimanager_object_system_sdnconnector_nic_ip_move`
* **New Resource:** `fortimanager_object_system_sdnconnector_nic_move`
* **New Resource:** `fortimanager_object_system_sdnconnector_ociregionlist`
* **New Resource:** `fortimanager_object_system_sdnconnector_ociregionlist_move`
* **New Resource:** `fortimanager_object_system_sdnconnector_route`
* **New Resource:** `fortimanager_object_system_sdnconnector_routetable`
* **New Resource:** `fortimanager_object_system_sdnconnector_routetable_move`
* **New Resource:** `fortimanager_object_system_sdnconnector_routetable_route`
* **New Resource:** `fortimanager_object_system_sdnconnector_routetable_route_move`
* **New Resource:** `fortimanager_object_system_sdnconnector_route_move`
* **New Resource:** `fortimanager_object_system_sdnproxy`
* **New Resource:** `fortimanager_object_user_domaincontroller_extraserver`
* **New Resource:** `fortimanager_object_user_fssopolling_adgrp`
* **New Resource:** `fortimanager_object_user_group_match`
* **New Resource:** `fortimanager_object_user_nsx_service`
* **New Resource:** `fortimanager_object_user_radius_accountingserver`
* **New Resource:** `fortimanager_object_user_securityexemptlist_rule`
* **New Resource:** `fortimanager_object_user_vcenter_rule`
* **New Resource:** `fortimanager_object_videofilter_profile_fortiguardcategory`
* **New Resource:** `fortimanager_object_videofilter_profile_fortiguardcategory_filters`
* **New Resource:** `fortimanager_object_videofilter_profile_fortiguardcategory_filters_move`
* **New Resource:** `fortimanager_object_virtualpatch_profile`
* **New Resource:** `fortimanager_object_virtualpatch_profile_exemption`
* **New Resource:** `fortimanager_object_voip_profile_msrp`
* **New Resource:** `fortimanager_object_voip_profile_sccp`
* **New Resource:** `fortimanager_object_voip_profile_sip`
* **New Resource:** `fortimanager_object_vpn_ipsec_fec_mappings`
* **New Resource:** `fortimanager_object_vpn_ipsec_fec_mappings_move`
* **New Resource:** `fortimanager_object_vpn_ssl_web_hostchecksoftware_checkitemlist`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal_bookmarkgroup`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal_bookmarkgroup_bookmarks`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal_bookmarkgroup_bookmarks_formdata`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal_bookmarkgroup_bookmarks_move`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal_bookmarkgroup_move`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal_landingpage`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal_landingpage_formdata`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal_macaddrcheckrule`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal_splitdns`
* **New Resource:** `fortimanager_object_vpn_ssl_web_portal_splitdns_move`
* **New Resource:** `fortimanager_object_vpnmgr_node_iprange`
* **New Resource:** `fortimanager_object_vpnmgr_node_ipv4excluderange`
* **New Resource:** `fortimanager_object_vpnmgr_node_protected_subnet`
* **New Resource:** `fortimanager_object_vpnmgr_node_summary_addr`
* **New Resource:** `fortimanager_object_waf_profile_addresslist`
* **New Resource:** `fortimanager_object_waf_profile_constraint`
* **New Resource:** `fortimanager_object_waf_profile_constraint_contentlength`
* **New Resource:** `fortimanager_object_waf_profile_constraint_exception`
* **New Resource:** `fortimanager_object_waf_profile_constraint_headerlength`
* **New Resource:** `fortimanager_object_waf_profile_constraint_hostname`
* **New Resource:** `fortimanager_object_waf_profile_constraint_linelength`
* **New Resource:** `fortimanager_object_waf_profile_constraint_malformed`
* **New Resource:** `fortimanager_object_waf_profile_constraint_maxcookie`
* **New Resource:** `fortimanager_object_waf_profile_constraint_maxheaderline`
* **New Resource:** `fortimanager_object_waf_profile_constraint_maxrangesegment`
* **New Resource:** `fortimanager_object_waf_profile_constraint_maxurlparam`
* **New Resource:** `fortimanager_object_waf_profile_constraint_method`
* **New Resource:** `fortimanager_object_waf_profile_constraint_paramlength`
* **New Resource:** `fortimanager_object_waf_profile_constraint_urlparamlength`
* **New Resource:** `fortimanager_object_waf_profile_constraint_version`
* **New Resource:** `fortimanager_object_waf_profile_method`
* **New Resource:** `fortimanager_object_waf_profile_method_methodpolicy`
* **New Resource:** `fortimanager_object_waf_profile_signature`
* **New Resource:** `fortimanager_object_waf_profile_signature_customsignature`
* **New Resource:** `fortimanager_object_waf_profile_urlaccess`
* **New Resource:** `fortimanager_object_waf_profile_urlaccess_accesspattern`
* **New Resource:** `fortimanager_object_wanopt_profile_cifs`
* **New Resource:** `fortimanager_object_wanopt_profile_ftp`
* **New Resource:** `fortimanager_object_wanopt_profile_http`
* **New Resource:** `fortimanager_object_wanopt_profile_mapi`
* **New Resource:** `fortimanager_object_wanopt_profile_tcp`
* **New Resource:** `fortimanager_object_webproxy_profile_headers`
* **New Resource:** `fortimanager_object_webfilter_contentheader_entries`
* **New Resource:** `fortimanager_object_webfilter_contentheader_entries_move`
* **New Resource:** `fortimanager_object_webfilter_content_entries`
* **New Resource:** `fortimanager_object_webfilter_profile_antiphish`
* **New Resource:** `fortimanager_object_webfilter_profile_antiphish_custompatterns`
* **New Resource:** `fortimanager_object_webfilter_profile_antiphish_inspectionentries`
* **New Resource:** `fortimanager_object_webfilter_profile_filefilter_entries`
* **New Resource:** `fortimanager_object_webfilter_profile_ftgdwf`
* **New Resource:** `fortimanager_object_webfilter_profile_ftgdwf_filters`
* **New Resource:** `fortimanager_object_webfilter_profile_override`
* **New Resource:** `fortimanager_object_webfilter_profile_urlextraction`
* **New Resource:** `fortimanager_object_webfilter_profile_web`
* **New Resource:** `fortimanager_object_webfilter_profile_youtubechannelfilter`
* **New Resource:** `fortimanager_object_webfilter_urlfilter_entries`
* **New Resource:** `fortimanager_object_webfilter_urlfilter_entries_move`
* **New Resource:** `fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv4rules`
* **New Resource:** `fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv4rules_move`
* **New Resource:** `fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules`
* **New Resource:** `fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules_move`
* **New Resource:** `fortimanager_object_wirelesscontroller_apcfgprofile`
* **New Resource:** `fortimanager_object_wirelesscontroller_apcfgprofile_commandlist`
* **New Resource:** `fortimanager_object_wirelesscontroller_apcfgprofile_commandlist_move`
* **New Resource:** `fortimanager_object_wirelesscontroller_bonjourprofile`
* **New Resource:** `fortimanager_object_wirelesscontroller_bonjourprofile_policylist`
* **New Resource:** `fortimanager_object_wirelesscontroller_bonjourprofile_policylist_move`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqp3gppcellular_mccmnclist`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod_authparam`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqproamingconsortium_oilist`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqpvenuename_valuelist`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_anqpvenueurl_valuelist`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist_planinfo`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qpoperatorname_valuelist`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovidernai_nailist`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovider_friendlyname`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovider_servicedescription`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_icon_iconlist`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_qosmap_dscpexcept`
* **New Resource:** `fortimanager_object_wirelesscontroller_hotspot20_qosmap_dscprange`
* **New Resource:** `fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup`
* **New Resource:** `fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_move`
* **New Resource:** `fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey`
* **New Resource:** `fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey_move`
* **New Resource:** `fortimanager_object_wirelesscontroller_vap_macfilterlist`
* **New Resource:** `fortimanager_object_wirelesscontroller_vap_portalmessageoverrides`
* **New Resource:** `fortimanager_object_wirelesscontroller_vap_vlanname`
* **New Resource:** `fortimanager_object_wirelesscontroller_vap_vlanname_move`
* **New Resource:** `fortimanager_object_wirelesscontroller_vap_vlanpool`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile_denymaclist`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile_eslsesdongle`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile_lan`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile_lbs`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile_platform`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile_radio1`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile_radio2`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile_radio3`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile_radio4`
* **New Resource:** `fortimanager_object_wirelesscontroller_wtpprofile_splittunnelingacl`
* **New Resource:** `fortimanager_packages_firewall_dospolicy6_anomaly`
* **New Resource:** `fortimanager_packages_firewall_dospolicy_anomaly`
* **New Resource:** `fortimanager_packages_pblock`
* **New Resource:** `fortimanager_securityconsole_install_objects_v2`
* **New Resource:** `fortimanager_system_admin_group_member`
* **New Resource:** `fortimanager_system_admin_user_adom`
* **New Resource:** `fortimanager_system_admin_user_ipsfilter`
* **New Resource:** `fortimanager_system_csf`
* **New Resource:** `fortimanager_system_csf_fabricconnector`
* **New Resource:** `fortimanager_system_logfetch_clientprofile_devicefilter`
* **New Resource:** `fortimanager_system_snmp_community_hosts`
* **New Resource:** `fortimanager_system_snmp_community_hosts6`
* **New Resource:** `fortimanager_system_template`
* **New Resource:** `fortimanager_system_workflow_approvalmatrix_approver`
* **New Resource:** `fortimanager_systemp_device_profile_fortianalyzer`
* **New Resource:** `fortimanager_systemp_device_profile_fortiguard`
* **New Resource:** `fortimanager_systemp_log_fortianalyzercloud_setting`
* **New Resource:** `fortimanager_systemp_log_fortianalyzer_setting`
* **New Resource:** `fortimanager_systemp_log_syslogd_filter`
* **New Resource:** `fortimanager_systemp_log_syslogd_setting`
* **New Resource:** `fortimanager_systemp_system_centralmanagement`
* **New Resource:** `fortimanager_systemp_system_centralmanagement_serverlist`
* **New Resource:** `fortimanager_systemp_system_emailserver`
* **New Resource:** `fortimanager_systemp_system_global`
* **New Resource:** `fortimanager_systemp_system_ntp`
* **New Resource:** `fortimanager_systemp_system_replacemsg_devicedetectionportal`
* **New Resource:** `fortimanager_systemp_system_replacemsg_nntp`
* **New Resource:** `fortimanager_systemp_system_snmp_sysinfo`
* **New Resource:** `fortimanager_systemp_system_snmp_user`
* **New Resource:** `fortimanager_wantemp_system_sdwan_members_move`
* **New Resource:** `fortimanager_wantemp_system_sdwan_service_move`
* **New Resource:** `fortimanager_wantemp_system_sdwan_zone_move`

## 1.8.0 (Aug 7, 2023)

IMPROVEMENTS:

* Support FMG v6.4.5-6.4.12, 7.0.6, 7.0.7,7.0.8, 7.2.2, 7.2.3, 7.4.0;
* Support Token based authentication
* Update sensitive arguments;
* Fix data type issue;

FEATURES:

* **New Resource:** `fortimanager_object_cloud_orchestaws`
* **New Resource:** `fortimanager_object_cloud_orchestawsconnector`
* **New Resource:** `fortimanager_object_cloud_orchestawstemplate_autoscaleexistingvpc`
* **New Resource:** `fortimanager_object_cloud_orchestawstemplate_autoscalenewvpc`
* **New Resource:** `fortimanager_object_cloud_orchestawstemplate_autoscaletgwnewvpc`
* **New Resource:** `fortimanager_object_cloud_orchestration`
* **New Resource:** `fortimanager_object_firewall_networkservicedynamic`
* **New Resource:** `fortimanager_object_switchcontroller_acl_group`
* **New Resource:** `fortimanager_object_switchcontroller_acl_ingress`
* **New Resource:** `fortimanager_object_system_npu_swtrhash`
* **New Resource:** `fortimanager_object_user_certificate`
* **New Resource:** `fortimanager_securityconsole_template_cli_preview`
* **New Resource:** `fortimanager_system_socfabric_trustedlist`
* **New Resource:** `fortimanager_packages_pblock_firewall_consolidated_policy`
* **New Resource:** `fortimanager_packages_pblock_firewall_policy`
* **New Resource:** `fortimanager_packages_pblock_firewall_policy6`
* **New Resource:** `fortimanager_packages_pblock_firewall_securitypolicy`

## 1.7.0 (Dec 21, 2022)

IMPROVEMENTS:

* Support FortiManager version 7.0.5
* Fix type convert issue
* Update error message in GetDeviceVersion

FEATURES:

* **New Resource:** `fortimanager_fmupdate_fwmsetting_upgradetimeout`
* **New Resource:** `fortimanager_object_dlp_datatype`
* **New Resource:** `fortimanager_object_dlp_dictionary`
* **New Resource:** `fortimanager_object_dlp_profile`
* **New Resource:** `fortimanager_object_extensioncontroller_dataplan`
* **New Resource:** `fortimanager_object_extensioncontroller_extenderprofile`
* **New Resource:** `fortimanager_wan_template`
* **New Resource:** `fortimanager_wantemp_system_sdwan`
* **New Resource:** `fortimanager_wantemp_system_sdwan_duplication`
* **New Resource:** `fortimanager_wantemp_system_sdwan_healthcheck`
* **New Resource:** `fortimanager_wantemp_system_sdwan_healthcheck_sla`
* **New Resource:** `fortimanager_wantemp_system_sdwan_members`
* **New Resource:** `fortimanager_wantemp_system_sdwan_neighbor`
* **New Resource:** `fortimanager_wantemp_system_sdwan_service`
* **New Resource:** `fortimanager_wantemp_system_sdwan_service_sla`
* **New Resource:** `fortimanager_wantemp_system_sdwan_zone`
* **New Resource:** `fortimanager_wantemp_system_virtualwanlink`
* **New Resource:** `fortimanager_wantemp_system_virtualwanlink_healthcheck`
* **New Resource:** `fortimanager_wantemp_system_virtualwanlink_healthcheck_sla`
* **New Resource:** `fortimanager_wantemp_system_virtualwanlink_members`
* **New Resource:** `fortimanager_wantemp_system_virtualwanlink_neighbor`
* **New Resource:** `fortimanager_wantemp_system_virtualwanlink_service`
* **New Resource:** `fortimanager_wantemp_system_virtualwanlink_service_sla`

## 1.6.0 (Sep 26, 2022)

IMPROVEMENTS:

* Support FortiManager version 7.2.1
* Update parameters' Computed feature to allow remove operation of parameters
* Fix some argument type convert issues
* Fix some panic issues
* Upgrade Terraform SDK to v2
* Fix block parameter could not be deleted issue
* Improve performance of move resources
* Fix capital letter issue

FEATURES:

* **New Resource:** `fortimanager_object_firewall_accessproxy6`
* **New Resource:** `fortimanager_object_firewall_accessproxy6_move`
* **New Resource:** `fortimanager_object_fmg_fabric_authorization_template`
* **New Resource:** `fortimanager_object_switchcontroller_dynamicportpolicy`
* **New Resource:** `fortimanager_object_switchcontroller_fortilinksettings`
* **New Resource:** `fortimanager_object_switchcontroller_macpolicy`
* **New Resource:** `fortimanager_object_switchcontroller_switchinterfacetag`
* **New Resource:** `fortimanager_object_switchcontroller_trafficpolicy`
* **New Resource:** `fortimanager_object_switchcontroller_vlanpolicy`
* **New Resource:** `fortimanager_object_user_flexvm`
* **New Resource:** `fortimanager_object_user_json`
* **New Resource:** `fortimanager_object_wirelesscontroller_accesscontrollist`
* **New Resource:** `fortimanager_object_wirelesscontroller_ssidpolicy`
* **New Resource:** `fortimanager_object_wirelesscontroller_syslogprofile`
* **New Resource:** `fortimanager_packages_user_nacpolicy`
* **New Resource:** `fortimanager_packages_user_nacpolicy_move`

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
