// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Global range attributes.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGlobalUpdate,
		Read:   resourceSystemGlobalRead,
		Update: resourceSystemGlobalUpdate,
		Delete: resourceSystemGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"admin_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin_lockout_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_lockout_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_lockout_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_grace_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"adom_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adom_rev_auto_delete": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adom_rev_max_backup_revisions": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"adom_rev_max_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"adom_rev_max_revisions": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"adom_select": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adom_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apache_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"api_ip_binding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clone_name_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clt_cert_req": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"console_output": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contentpack_fgt_install": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"country_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"create_revision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"daylightsavetime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detect_unregistered_log_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_view_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_params": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disable_module": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_storage_pool_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fabric_storage_pool_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"faz_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fcp_cfg_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgfm_ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fgfm_cert_exclusive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgfm_deny_unknown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgfm_local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fgfm_peercert_withoutsn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgfm_ssl_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiservice_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"global_ssl_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_curl_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gui_polling_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ha_member_auto_grouping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpd_ssl_protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"import_ignore_addr_cmt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jsonapi_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"latitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_cache_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ldapconntimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lock_preempt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_checksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_checksum_upload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_forward_cache_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"longitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mapclient_ssl_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_log_forward": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_running_reports": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mc_policy_disabled_adoms": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"adom_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"multiple_steps_upgrade_in_autolink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"no_copy_permission_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"no_vip_value_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"normalized_intf_zone_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_revision_db_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"object_revision_mandatory_note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_revision_object_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"object_revision_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oftp_ssl_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"partial_install": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"partial_install_force": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"partial_install_rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_policy_lock": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"perform_improve_by_ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_object_icon": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_object_in_dual_pane": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pre_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pre_login_banner_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_data_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remoteauthtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"save_last_hit_in_adomdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"search_all_adoms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_enc_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_hostkey_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_kex_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_mac_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_strong_crypto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_cipher_suites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_low_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_static_key_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"table_entry_blink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"task_list_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tftp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"usg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webservice_proto": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"workflow_max_sessions": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"workspace_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"workspace_unlock_after_install": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemGlobal(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemGlobal")

	return resourceSystemGlobalRead(d, m)
}

func resourceSystemGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobal resource from API: %v", err)
	}
	return nil
}

func flattenSystemGlobalAdminHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminLockoutDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminLockoutMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminLockoutThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshGraceTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomRevAutoDelete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomRevMaxBackupRevisions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomRevMaxDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomRevMaxRevisions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalApacheMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalApiIpBinding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCloneNameOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCltCertReq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalConsoleOutput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalContentpackFgtInstall(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCountryFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCreateRevision(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDaylightsavetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDetectUnregisteredLogDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDeviceViewMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDhParams(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDisableModule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFabricStoragePoolQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFabricStoragePoolSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFazStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFcpCfgService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFgfmCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFgfmCertExclusive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFgfmDenyUnknown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFgfmLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFgfmPeercertWithoutsn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFgfmSslProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFortiservicePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGlobalSslProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiCurlTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiPollingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHaMemberAutoGrouping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHttpdSslProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalImportIgnoreAddrCmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalJsonapiLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLanguage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLatitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLdapCacheTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLdapconntimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLockPreempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLogChecksum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLogChecksumUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLogForwardCacheSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLongitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalManagementIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalManagementPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMapclientSslProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMaxLogForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMaxRunningReports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMcPolicyDisabledAdoms(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adom_name"
		if _, ok := i["adom-name"]; ok {
			v := flattenSystemGlobalMcPolicyDisabledAdomsAdomName(i["adom-name"], d, pre_append)
			tmp["adom_name"] = fortiAPISubPartPatch(v, "SystemGlobal-McPolicyDisabledAdoms-AdomName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemGlobalMcPolicyDisabledAdomsAdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMultipleStepsUpgradeInAutolink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalNoCopyPermissionCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalNoVipValueCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalNormalizedIntfZoneOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalObjectRevisionDbMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalObjectRevisionMandatoryNote(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalObjectRevisionObjectMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalObjectRevisionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalOftpSslProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPartialInstall(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPartialInstallForce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPartialInstallRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPerPolicyLock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPerformImproveByHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPolicyObjectIcon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPolicyObjectInDualPane(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPreLoginBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPreLoginBannerMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPrivateDataEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalRemoteauthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSaveLastHitInAdomdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSearchAllAdoms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSshEncAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSshHostkeyAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSshKexAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSshMacAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSshStrongCrypto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslCipherSuites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {
			v := flattenSystemGlobalSslCipherSuitesCipher(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "SystemGlobal-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenSystemGlobalSslCipherSuitesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "SystemGlobal-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenSystemGlobalSslCipherSuitesVersion(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "SystemGlobal-SslCipherSuites-Version")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemGlobalSslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslCipherSuitesVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslLowEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSslStaticKeyCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTableEntryBlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTaskListSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTftp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTunnelMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalUsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalVdomMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWebserviceProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalWorkflowMaxSessions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWorkspaceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWorkspaceUnlockAfterInstall(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("admin_host", flattenSystemGlobalAdminHost(o["admin-host"], d, "admin_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-host"], "SystemGlobal-AdminHost"); ok {
			if err = d.Set("admin_host", vv); err != nil {
				return fmt.Errorf("Error reading admin_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_host: %v", err)
		}
	}

	if err = d.Set("admin_lockout_duration", flattenSystemGlobalAdminLockoutDuration(o["admin-lockout-duration"], d, "admin_lockout_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-lockout-duration"], "SystemGlobal-AdminLockoutDuration"); ok {
			if err = d.Set("admin_lockout_duration", vv); err != nil {
				return fmt.Errorf("Error reading admin_lockout_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_lockout_duration: %v", err)
		}
	}

	if err = d.Set("admin_lockout_method", flattenSystemGlobalAdminLockoutMethod(o["admin-lockout-method"], d, "admin_lockout_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-lockout-method"], "SystemGlobal-AdminLockoutMethod"); ok {
			if err = d.Set("admin_lockout_method", vv); err != nil {
				return fmt.Errorf("Error reading admin_lockout_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_lockout_method: %v", err)
		}
	}

	if err = d.Set("admin_lockout_threshold", flattenSystemGlobalAdminLockoutThreshold(o["admin-lockout-threshold"], d, "admin_lockout_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-lockout-threshold"], "SystemGlobal-AdminLockoutThreshold"); ok {
			if err = d.Set("admin_lockout_threshold", vv); err != nil {
				return fmt.Errorf("Error reading admin_lockout_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_lockout_threshold: %v", err)
		}
	}

	if err = d.Set("admin_ssh_grace_time", flattenSystemGlobalAdminSshGraceTime(o["admin-ssh-grace-time"], d, "admin_ssh_grace_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-ssh-grace-time"], "SystemGlobal-AdminSshGraceTime"); ok {
			if err = d.Set("admin_ssh_grace_time", vv); err != nil {
				return fmt.Errorf("Error reading admin_ssh_grace_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_ssh_grace_time: %v", err)
		}
	}

	if err = d.Set("adom_mode", flattenSystemGlobalAdomMode(o["adom-mode"], d, "adom_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-mode"], "SystemGlobal-AdomMode"); ok {
			if err = d.Set("adom_mode", vv); err != nil {
				return fmt.Errorf("Error reading adom_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_mode: %v", err)
		}
	}

	if err = d.Set("adom_rev_auto_delete", flattenSystemGlobalAdomRevAutoDelete(o["adom-rev-auto-delete"], d, "adom_rev_auto_delete")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-rev-auto-delete"], "SystemGlobal-AdomRevAutoDelete"); ok {
			if err = d.Set("adom_rev_auto_delete", vv); err != nil {
				return fmt.Errorf("Error reading adom_rev_auto_delete: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_rev_auto_delete: %v", err)
		}
	}

	if err = d.Set("adom_rev_max_backup_revisions", flattenSystemGlobalAdomRevMaxBackupRevisions(o["adom-rev-max-backup-revisions"], d, "adom_rev_max_backup_revisions")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-rev-max-backup-revisions"], "SystemGlobal-AdomRevMaxBackupRevisions"); ok {
			if err = d.Set("adom_rev_max_backup_revisions", vv); err != nil {
				return fmt.Errorf("Error reading adom_rev_max_backup_revisions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_rev_max_backup_revisions: %v", err)
		}
	}

	if err = d.Set("adom_rev_max_days", flattenSystemGlobalAdomRevMaxDays(o["adom-rev-max-days"], d, "adom_rev_max_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-rev-max-days"], "SystemGlobal-AdomRevMaxDays"); ok {
			if err = d.Set("adom_rev_max_days", vv); err != nil {
				return fmt.Errorf("Error reading adom_rev_max_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_rev_max_days: %v", err)
		}
	}

	if err = d.Set("adom_rev_max_revisions", flattenSystemGlobalAdomRevMaxRevisions(o["adom-rev-max-revisions"], d, "adom_rev_max_revisions")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-rev-max-revisions"], "SystemGlobal-AdomRevMaxRevisions"); ok {
			if err = d.Set("adom_rev_max_revisions", vv); err != nil {
				return fmt.Errorf("Error reading adom_rev_max_revisions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_rev_max_revisions: %v", err)
		}
	}

	if err = d.Set("adom_select", flattenSystemGlobalAdomSelect(o["adom-select"], d, "adom_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-select"], "SystemGlobal-AdomSelect"); ok {
			if err = d.Set("adom_select", vv); err != nil {
				return fmt.Errorf("Error reading adom_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_select: %v", err)
		}
	}

	if err = d.Set("adom_status", flattenSystemGlobalAdomStatus(o["adom-status"], d, "adom_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-status"], "SystemGlobal-AdomStatus"); ok {
			if err = d.Set("adom_status", vv); err != nil {
				return fmt.Errorf("Error reading adom_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_status: %v", err)
		}
	}

	if err = d.Set("apache_mode", flattenSystemGlobalApacheMode(o["apache-mode"], d, "apache_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["apache-mode"], "SystemGlobal-ApacheMode"); ok {
			if err = d.Set("apache_mode", vv); err != nil {
				return fmt.Errorf("Error reading apache_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apache_mode: %v", err)
		}
	}

	if err = d.Set("api_ip_binding", flattenSystemGlobalApiIpBinding(o["api-ip-binding"], d, "api_ip_binding")); err != nil {
		if vv, ok := fortiAPIPatch(o["api-ip-binding"], "SystemGlobal-ApiIpBinding"); ok {
			if err = d.Set("api_ip_binding", vv); err != nil {
				return fmt.Errorf("Error reading api_ip_binding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading api_ip_binding: %v", err)
		}
	}

	if err = d.Set("clone_name_option", flattenSystemGlobalCloneNameOption(o["clone-name-option"], d, "clone_name_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["clone-name-option"], "SystemGlobal-CloneNameOption"); ok {
			if err = d.Set("clone_name_option", vv); err != nil {
				return fmt.Errorf("Error reading clone_name_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clone_name_option: %v", err)
		}
	}

	if err = d.Set("clt_cert_req", flattenSystemGlobalCltCertReq(o["clt-cert-req"], d, "clt_cert_req")); err != nil {
		if vv, ok := fortiAPIPatch(o["clt-cert-req"], "SystemGlobal-CltCertReq"); ok {
			if err = d.Set("clt_cert_req", vv); err != nil {
				return fmt.Errorf("Error reading clt_cert_req: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clt_cert_req: %v", err)
		}
	}

	if err = d.Set("console_output", flattenSystemGlobalConsoleOutput(o["console-output"], d, "console_output")); err != nil {
		if vv, ok := fortiAPIPatch(o["console-output"], "SystemGlobal-ConsoleOutput"); ok {
			if err = d.Set("console_output", vv); err != nil {
				return fmt.Errorf("Error reading console_output: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading console_output: %v", err)
		}
	}

	if err = d.Set("contentpack_fgt_install", flattenSystemGlobalContentpackFgtInstall(o["contentpack-fgt-install"], d, "contentpack_fgt_install")); err != nil {
		if vv, ok := fortiAPIPatch(o["contentpack-fgt-install"], "SystemGlobal-ContentpackFgtInstall"); ok {
			if err = d.Set("contentpack_fgt_install", vv); err != nil {
				return fmt.Errorf("Error reading contentpack_fgt_install: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading contentpack_fgt_install: %v", err)
		}
	}

	if err = d.Set("country_flag", flattenSystemGlobalCountryFlag(o["country-flag"], d, "country_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["country-flag"], "SystemGlobal-CountryFlag"); ok {
			if err = d.Set("country_flag", vv); err != nil {
				return fmt.Errorf("Error reading country_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country_flag: %v", err)
		}
	}

	if err = d.Set("create_revision", flattenSystemGlobalCreateRevision(o["create-revision"], d, "create_revision")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-revision"], "SystemGlobal-CreateRevision"); ok {
			if err = d.Set("create_revision", vv); err != nil {
				return fmt.Errorf("Error reading create_revision: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_revision: %v", err)
		}
	}

	if err = d.Set("daylightsavetime", flattenSystemGlobalDaylightsavetime(o["daylightsavetime"], d, "daylightsavetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["daylightsavetime"], "SystemGlobal-Daylightsavetime"); ok {
			if err = d.Set("daylightsavetime", vv); err != nil {
				return fmt.Errorf("Error reading daylightsavetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading daylightsavetime: %v", err)
		}
	}

	if err = d.Set("detect_unregistered_log_device", flattenSystemGlobalDetectUnregisteredLogDevice(o["detect-unregistered-log-device"], d, "detect_unregistered_log_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-unregistered-log-device"], "SystemGlobal-DetectUnregisteredLogDevice"); ok {
			if err = d.Set("detect_unregistered_log_device", vv); err != nil {
				return fmt.Errorf("Error reading detect_unregistered_log_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_unregistered_log_device: %v", err)
		}
	}

	if err = d.Set("device_view_mode", flattenSystemGlobalDeviceViewMode(o["device-view-mode"], d, "device_view_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-view-mode"], "SystemGlobal-DeviceViewMode"); ok {
			if err = d.Set("device_view_mode", vv); err != nil {
				return fmt.Errorf("Error reading device_view_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_view_mode: %v", err)
		}
	}

	if err = d.Set("dh_params", flattenSystemGlobalDhParams(o["dh-params"], d, "dh_params")); err != nil {
		if vv, ok := fortiAPIPatch(o["dh-params"], "SystemGlobal-DhParams"); ok {
			if err = d.Set("dh_params", vv); err != nil {
				return fmt.Errorf("Error reading dh_params: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dh_params: %v", err)
		}
	}

	if err = d.Set("disable_module", flattenSystemGlobalDisableModule(o["disable-module"], d, "disable_module")); err != nil {
		if vv, ok := fortiAPIPatch(o["disable-module"], "SystemGlobal-DisableModule"); ok {
			if err = d.Set("disable_module", vv); err != nil {
				return fmt.Errorf("Error reading disable_module: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disable_module: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenSystemGlobalEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["enc-algorithm"], "SystemGlobal-EncAlgorithm"); ok {
			if err = d.Set("enc_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading enc_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("fabric_storage_pool_quota", flattenSystemGlobalFabricStoragePoolQuota(o["fabric-storage-pool-quota"], d, "fabric_storage_pool_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-storage-pool-quota"], "SystemGlobal-FabricStoragePoolQuota"); ok {
			if err = d.Set("fabric_storage_pool_quota", vv); err != nil {
				return fmt.Errorf("Error reading fabric_storage_pool_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_storage_pool_quota: %v", err)
		}
	}

	if err = d.Set("fabric_storage_pool_size", flattenSystemGlobalFabricStoragePoolSize(o["fabric-storage-pool-size"], d, "fabric_storage_pool_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-storage-pool-size"], "SystemGlobal-FabricStoragePoolSize"); ok {
			if err = d.Set("fabric_storage_pool_size", vv); err != nil {
				return fmt.Errorf("Error reading fabric_storage_pool_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_storage_pool_size: %v", err)
		}
	}

	if err = d.Set("faz_status", flattenSystemGlobalFazStatus(o["faz-status"], d, "faz_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-status"], "SystemGlobal-FazStatus"); ok {
			if err = d.Set("faz_status", vv); err != nil {
				return fmt.Errorf("Error reading faz_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_status: %v", err)
		}
	}

	if err = d.Set("fcp_cfg_service", flattenSystemGlobalFcpCfgService(o["fcp-cfg-service"], d, "fcp_cfg_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["fcp-cfg-service"], "SystemGlobal-FcpCfgService"); ok {
			if err = d.Set("fcp_cfg_service", vv); err != nil {
				return fmt.Errorf("Error reading fcp_cfg_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fcp_cfg_service: %v", err)
		}
	}

	if err = d.Set("fgfm_ca_cert", flattenSystemGlobalFgfmCaCert(o["fgfm-ca-cert"], d, "fgfm_ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-ca-cert"], "SystemGlobal-FgfmCaCert"); ok {
			if err = d.Set("fgfm_ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_ca_cert: %v", err)
		}
	}

	if err = d.Set("fgfm_cert_exclusive", flattenSystemGlobalFgfmCertExclusive(o["fgfm-cert-exclusive"], d, "fgfm_cert_exclusive")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-cert-exclusive"], "SystemGlobal-FgfmCertExclusive"); ok {
			if err = d.Set("fgfm_cert_exclusive", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_cert_exclusive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_cert_exclusive: %v", err)
		}
	}

	if err = d.Set("fgfm_deny_unknown", flattenSystemGlobalFgfmDenyUnknown(o["fgfm-deny-unknown"], d, "fgfm_deny_unknown")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-deny-unknown"], "SystemGlobal-FgfmDenyUnknown"); ok {
			if err = d.Set("fgfm_deny_unknown", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_deny_unknown: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_deny_unknown: %v", err)
		}
	}

	if err = d.Set("fgfm_local_cert", flattenSystemGlobalFgfmLocalCert(o["fgfm-local-cert"], d, "fgfm_local_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-local-cert"], "SystemGlobal-FgfmLocalCert"); ok {
			if err = d.Set("fgfm_local_cert", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_local_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_local_cert: %v", err)
		}
	}

	if err = d.Set("fgfm_peercert_withoutsn", flattenSystemGlobalFgfmPeercertWithoutsn(o["fgfm-peercert-withoutsn"], d, "fgfm_peercert_withoutsn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-peercert-withoutsn"], "SystemGlobal-FgfmPeercertWithoutsn"); ok {
			if err = d.Set("fgfm_peercert_withoutsn", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_peercert_withoutsn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_peercert_withoutsn: %v", err)
		}
	}

	if err = d.Set("fgfm_ssl_protocol", flattenSystemGlobalFgfmSslProtocol(o["fgfm-ssl-protocol"], d, "fgfm_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-ssl-protocol"], "SystemGlobal-FgfmSslProtocol"); ok {
			if err = d.Set("fgfm_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("fortiservice_port", flattenSystemGlobalFortiservicePort(o["fortiservice-port"], d, "fortiservice_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiservice-port"], "SystemGlobal-FortiservicePort"); ok {
			if err = d.Set("fortiservice_port", vv); err != nil {
				return fmt.Errorf("Error reading fortiservice_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiservice_port: %v", err)
		}
	}

	if err = d.Set("global_ssl_protocol", flattenSystemGlobalGlobalSslProtocol(o["global-ssl-protocol"], d, "global_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-ssl-protocol"], "SystemGlobal-GlobalSslProtocol"); ok {
			if err = d.Set("global_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading global_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("gui_curl_timeout", flattenSystemGlobalGuiCurlTimeout(o["gui-curl-timeout"], d, "gui_curl_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-curl-timeout"], "SystemGlobal-GuiCurlTimeout"); ok {
			if err = d.Set("gui_curl_timeout", vv); err != nil {
				return fmt.Errorf("Error reading gui_curl_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_curl_timeout: %v", err)
		}
	}

	if err = d.Set("gui_polling_interval", flattenSystemGlobalGuiPollingInterval(o["gui-polling-interval"], d, "gui_polling_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-polling-interval"], "SystemGlobal-GuiPollingInterval"); ok {
			if err = d.Set("gui_polling_interval", vv); err != nil {
				return fmt.Errorf("Error reading gui_polling_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_polling_interval: %v", err)
		}
	}

	if err = d.Set("ha_member_auto_grouping", flattenSystemGlobalHaMemberAutoGrouping(o["ha-member-auto-grouping"], d, "ha_member_auto_grouping")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-member-auto-grouping"], "SystemGlobal-HaMemberAutoGrouping"); ok {
			if err = d.Set("ha_member_auto_grouping", vv); err != nil {
				return fmt.Errorf("Error reading ha_member_auto_grouping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_member_auto_grouping: %v", err)
		}
	}

	if err = d.Set("hostname", flattenSystemGlobalHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "SystemGlobal-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("httpd_ssl_protocol", flattenSystemGlobalHttpdSslProtocol(o["httpd-ssl-protocol"], d, "httpd_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["httpd-ssl-protocol"], "SystemGlobal-HttpdSslProtocol"); ok {
			if err = d.Set("httpd_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading httpd_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading httpd_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("import_ignore_addr_cmt", flattenSystemGlobalImportIgnoreAddrCmt(o["import-ignore-addr-cmt"], d, "import_ignore_addr_cmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["import-ignore-addr-cmt"], "SystemGlobal-ImportIgnoreAddrCmt"); ok {
			if err = d.Set("import_ignore_addr_cmt", vv); err != nil {
				return fmt.Errorf("Error reading import_ignore_addr_cmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading import_ignore_addr_cmt: %v", err)
		}
	}

	if err = d.Set("jsonapi_log", flattenSystemGlobalJsonapiLog(o["jsonapi-log"], d, "jsonapi_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["jsonapi-log"], "SystemGlobal-JsonapiLog"); ok {
			if err = d.Set("jsonapi_log", vv); err != nil {
				return fmt.Errorf("Error reading jsonapi_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jsonapi_log: %v", err)
		}
	}

	if err = d.Set("language", flattenSystemGlobalLanguage(o["language"], d, "language")); err != nil {
		if vv, ok := fortiAPIPatch(o["language"], "SystemGlobal-Language"); ok {
			if err = d.Set("language", vv); err != nil {
				return fmt.Errorf("Error reading language: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("latitude", flattenSystemGlobalLatitude(o["latitude"], d, "latitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["latitude"], "SystemGlobal-Latitude"); ok {
			if err = d.Set("latitude", vv); err != nil {
				return fmt.Errorf("Error reading latitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latitude: %v", err)
		}
	}

	if err = d.Set("ldap_cache_timeout", flattenSystemGlobalLdapCacheTimeout(o["ldap-cache-timeout"], d, "ldap_cache_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-cache-timeout"], "SystemGlobal-LdapCacheTimeout"); ok {
			if err = d.Set("ldap_cache_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ldap_cache_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_cache_timeout: %v", err)
		}
	}

	if err = d.Set("ldapconntimeout", flattenSystemGlobalLdapconntimeout(o["ldapconntimeout"], d, "ldapconntimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldapconntimeout"], "SystemGlobal-Ldapconntimeout"); ok {
			if err = d.Set("ldapconntimeout", vv); err != nil {
				return fmt.Errorf("Error reading ldapconntimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldapconntimeout: %v", err)
		}
	}

	if err = d.Set("lock_preempt", flattenSystemGlobalLockPreempt(o["lock-preempt"], d, "lock_preempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["lock-preempt"], "SystemGlobal-LockPreempt"); ok {
			if err = d.Set("lock_preempt", vv); err != nil {
				return fmt.Errorf("Error reading lock_preempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lock_preempt: %v", err)
		}
	}

	if err = d.Set("log_checksum", flattenSystemGlobalLogChecksum(o["log-checksum"], d, "log_checksum")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-checksum"], "SystemGlobal-LogChecksum"); ok {
			if err = d.Set("log_checksum", vv); err != nil {
				return fmt.Errorf("Error reading log_checksum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_checksum: %v", err)
		}
	}

	if err = d.Set("log_checksum_upload", flattenSystemGlobalLogChecksumUpload(o["log-checksum-upload"], d, "log_checksum_upload")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-checksum-upload"], "SystemGlobal-LogChecksumUpload"); ok {
			if err = d.Set("log_checksum_upload", vv); err != nil {
				return fmt.Errorf("Error reading log_checksum_upload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_checksum_upload: %v", err)
		}
	}

	if err = d.Set("log_forward_cache_size", flattenSystemGlobalLogForwardCacheSize(o["log-forward-cache-size"], d, "log_forward_cache_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-forward-cache-size"], "SystemGlobal-LogForwardCacheSize"); ok {
			if err = d.Set("log_forward_cache_size", vv); err != nil {
				return fmt.Errorf("Error reading log_forward_cache_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_forward_cache_size: %v", err)
		}
	}

	if err = d.Set("longitude", flattenSystemGlobalLongitude(o["longitude"], d, "longitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["longitude"], "SystemGlobal-Longitude"); ok {
			if err = d.Set("longitude", vv); err != nil {
				return fmt.Errorf("Error reading longitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading longitude: %v", err)
		}
	}

	if err = d.Set("management_ip", flattenSystemGlobalManagementIp(o["management-ip"], d, "management_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["management-ip"], "SystemGlobal-ManagementIp"); ok {
			if err = d.Set("management_ip", vv); err != nil {
				return fmt.Errorf("Error reading management_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("management_port", flattenSystemGlobalManagementPort(o["management-port"], d, "management_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["management-port"], "SystemGlobal-ManagementPort"); ok {
			if err = d.Set("management_port", vv); err != nil {
				return fmt.Errorf("Error reading management_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading management_port: %v", err)
		}
	}

	if err = d.Set("mapclient_ssl_protocol", flattenSystemGlobalMapclientSslProtocol(o["mapclient-ssl-protocol"], d, "mapclient_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["mapclient-ssl-protocol"], "SystemGlobal-MapclientSslProtocol"); ok {
			if err = d.Set("mapclient_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading mapclient_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mapclient_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("max_log_forward", flattenSystemGlobalMaxLogForward(o["max-log-forward"], d, "max_log_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-forward"], "SystemGlobal-MaxLogForward"); ok {
			if err = d.Set("max_log_forward", vv); err != nil {
				return fmt.Errorf("Error reading max_log_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_forward: %v", err)
		}
	}

	if err = d.Set("max_running_reports", flattenSystemGlobalMaxRunningReports(o["max-running-reports"], d, "max_running_reports")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-running-reports"], "SystemGlobal-MaxRunningReports"); ok {
			if err = d.Set("max_running_reports", vv); err != nil {
				return fmt.Errorf("Error reading max_running_reports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_running_reports: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mc_policy_disabled_adoms", flattenSystemGlobalMcPolicyDisabledAdoms(o["mc-policy-disabled-adoms"], d, "mc_policy_disabled_adoms")); err != nil {
			if vv, ok := fortiAPIPatch(o["mc-policy-disabled-adoms"], "SystemGlobal-McPolicyDisabledAdoms"); ok {
				if err = d.Set("mc_policy_disabled_adoms", vv); err != nil {
					return fmt.Errorf("Error reading mc_policy_disabled_adoms: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mc_policy_disabled_adoms: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mc_policy_disabled_adoms"); ok {
			if err = d.Set("mc_policy_disabled_adoms", flattenSystemGlobalMcPolicyDisabledAdoms(o["mc-policy-disabled-adoms"], d, "mc_policy_disabled_adoms")); err != nil {
				if vv, ok := fortiAPIPatch(o["mc-policy-disabled-adoms"], "SystemGlobal-McPolicyDisabledAdoms"); ok {
					if err = d.Set("mc_policy_disabled_adoms", vv); err != nil {
						return fmt.Errorf("Error reading mc_policy_disabled_adoms: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mc_policy_disabled_adoms: %v", err)
				}
			}
		}
	}

	if err = d.Set("multiple_steps_upgrade_in_autolink", flattenSystemGlobalMultipleStepsUpgradeInAutolink(o["multiple-steps-upgrade-in-autolink"], d, "multiple_steps_upgrade_in_autolink")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiple-steps-upgrade-in-autolink"], "SystemGlobal-MultipleStepsUpgradeInAutolink"); ok {
			if err = d.Set("multiple_steps_upgrade_in_autolink", vv); err != nil {
				return fmt.Errorf("Error reading multiple_steps_upgrade_in_autolink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiple_steps_upgrade_in_autolink: %v", err)
		}
	}

	if err = d.Set("no_copy_permission_check", flattenSystemGlobalNoCopyPermissionCheck(o["no-copy-permission-check"], d, "no_copy_permission_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["no-copy-permission-check"], "SystemGlobal-NoCopyPermissionCheck"); ok {
			if err = d.Set("no_copy_permission_check", vv); err != nil {
				return fmt.Errorf("Error reading no_copy_permission_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading no_copy_permission_check: %v", err)
		}
	}

	if err = d.Set("no_vip_value_check", flattenSystemGlobalNoVipValueCheck(o["no-vip-value-check"], d, "no_vip_value_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["no-vip-value-check"], "SystemGlobal-NoVipValueCheck"); ok {
			if err = d.Set("no_vip_value_check", vv); err != nil {
				return fmt.Errorf("Error reading no_vip_value_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading no_vip_value_check: %v", err)
		}
	}

	if err = d.Set("normalized_intf_zone_only", flattenSystemGlobalNormalizedIntfZoneOnly(o["normalized-intf-zone-only"], d, "normalized_intf_zone_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["normalized-intf-zone-only"], "SystemGlobal-NormalizedIntfZoneOnly"); ok {
			if err = d.Set("normalized_intf_zone_only", vv); err != nil {
				return fmt.Errorf("Error reading normalized_intf_zone_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading normalized_intf_zone_only: %v", err)
		}
	}

	if err = d.Set("object_revision_db_max", flattenSystemGlobalObjectRevisionDbMax(o["object-revision-db-max"], d, "object_revision_db_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-revision-db-max"], "SystemGlobal-ObjectRevisionDbMax"); ok {
			if err = d.Set("object_revision_db_max", vv); err != nil {
				return fmt.Errorf("Error reading object_revision_db_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_revision_db_max: %v", err)
		}
	}

	if err = d.Set("object_revision_mandatory_note", flattenSystemGlobalObjectRevisionMandatoryNote(o["object-revision-mandatory-note"], d, "object_revision_mandatory_note")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-revision-mandatory-note"], "SystemGlobal-ObjectRevisionMandatoryNote"); ok {
			if err = d.Set("object_revision_mandatory_note", vv); err != nil {
				return fmt.Errorf("Error reading object_revision_mandatory_note: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_revision_mandatory_note: %v", err)
		}
	}

	if err = d.Set("object_revision_object_max", flattenSystemGlobalObjectRevisionObjectMax(o["object-revision-object-max"], d, "object_revision_object_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-revision-object-max"], "SystemGlobal-ObjectRevisionObjectMax"); ok {
			if err = d.Set("object_revision_object_max", vv); err != nil {
				return fmt.Errorf("Error reading object_revision_object_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_revision_object_max: %v", err)
		}
	}

	if err = d.Set("object_revision_status", flattenSystemGlobalObjectRevisionStatus(o["object-revision-status"], d, "object_revision_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-revision-status"], "SystemGlobal-ObjectRevisionStatus"); ok {
			if err = d.Set("object_revision_status", vv); err != nil {
				return fmt.Errorf("Error reading object_revision_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_revision_status: %v", err)
		}
	}

	if err = d.Set("oftp_ssl_protocol", flattenSystemGlobalOftpSslProtocol(o["oftp-ssl-protocol"], d, "oftp_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["oftp-ssl-protocol"], "SystemGlobal-OftpSslProtocol"); ok {
			if err = d.Set("oftp_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading oftp_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oftp_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("partial_install", flattenSystemGlobalPartialInstall(o["partial-install"], d, "partial_install")); err != nil {
		if vv, ok := fortiAPIPatch(o["partial-install"], "SystemGlobal-PartialInstall"); ok {
			if err = d.Set("partial_install", vv); err != nil {
				return fmt.Errorf("Error reading partial_install: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading partial_install: %v", err)
		}
	}

	if err = d.Set("partial_install_force", flattenSystemGlobalPartialInstallForce(o["partial-install-force"], d, "partial_install_force")); err != nil {
		if vv, ok := fortiAPIPatch(o["partial-install-force"], "SystemGlobal-PartialInstallForce"); ok {
			if err = d.Set("partial_install_force", vv); err != nil {
				return fmt.Errorf("Error reading partial_install_force: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading partial_install_force: %v", err)
		}
	}

	if err = d.Set("partial_install_rev", flattenSystemGlobalPartialInstallRev(o["partial-install-rev"], d, "partial_install_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["partial-install-rev"], "SystemGlobal-PartialInstallRev"); ok {
			if err = d.Set("partial_install_rev", vv); err != nil {
				return fmt.Errorf("Error reading partial_install_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading partial_install_rev: %v", err)
		}
	}

	if err = d.Set("per_policy_lock", flattenSystemGlobalPerPolicyLock(o["per-policy-lock"], d, "per_policy_lock")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-policy-lock"], "SystemGlobal-PerPolicyLock"); ok {
			if err = d.Set("per_policy_lock", vv); err != nil {
				return fmt.Errorf("Error reading per_policy_lock: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_policy_lock: %v", err)
		}
	}

	if err = d.Set("perform_improve_by_ha", flattenSystemGlobalPerformImproveByHa(o["perform-improve-by-ha"], d, "perform_improve_by_ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["perform-improve-by-ha"], "SystemGlobal-PerformImproveByHa"); ok {
			if err = d.Set("perform_improve_by_ha", vv); err != nil {
				return fmt.Errorf("Error reading perform_improve_by_ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading perform_improve_by_ha: %v", err)
		}
	}

	if err = d.Set("policy_object_icon", flattenSystemGlobalPolicyObjectIcon(o["policy-object-icon"], d, "policy_object_icon")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-object-icon"], "SystemGlobal-PolicyObjectIcon"); ok {
			if err = d.Set("policy_object_icon", vv); err != nil {
				return fmt.Errorf("Error reading policy_object_icon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_object_icon: %v", err)
		}
	}

	if err = d.Set("policy_object_in_dual_pane", flattenSystemGlobalPolicyObjectInDualPane(o["policy-object-in-dual-pane"], d, "policy_object_in_dual_pane")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-object-in-dual-pane"], "SystemGlobal-PolicyObjectInDualPane"); ok {
			if err = d.Set("policy_object_in_dual_pane", vv); err != nil {
				return fmt.Errorf("Error reading policy_object_in_dual_pane: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_object_in_dual_pane: %v", err)
		}
	}

	if err = d.Set("pre_login_banner", flattenSystemGlobalPreLoginBanner(o["pre-login-banner"], d, "pre_login_banner")); err != nil {
		if vv, ok := fortiAPIPatch(o["pre-login-banner"], "SystemGlobal-PreLoginBanner"); ok {
			if err = d.Set("pre_login_banner", vv); err != nil {
				return fmt.Errorf("Error reading pre_login_banner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pre_login_banner: %v", err)
		}
	}

	if err = d.Set("pre_login_banner_message", flattenSystemGlobalPreLoginBannerMessage(o["pre-login-banner-message"], d, "pre_login_banner_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["pre-login-banner-message"], "SystemGlobal-PreLoginBannerMessage"); ok {
			if err = d.Set("pre_login_banner_message", vv); err != nil {
				return fmt.Errorf("Error reading pre_login_banner_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pre_login_banner_message: %v", err)
		}
	}

	if err = d.Set("private_data_encryption", flattenSystemGlobalPrivateDataEncryption(o["private-data-encryption"], d, "private_data_encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-data-encryption"], "SystemGlobal-PrivateDataEncryption"); ok {
			if err = d.Set("private_data_encryption", vv); err != nil {
				return fmt.Errorf("Error reading private_data_encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_data_encryption: %v", err)
		}
	}

	if err = d.Set("remoteauthtimeout", flattenSystemGlobalRemoteauthtimeout(o["remoteauthtimeout"], d, "remoteauthtimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["remoteauthtimeout"], "SystemGlobal-Remoteauthtimeout"); ok {
			if err = d.Set("remoteauthtimeout", vv); err != nil {
				return fmt.Errorf("Error reading remoteauthtimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remoteauthtimeout: %v", err)
		}
	}

	if err = d.Set("save_last_hit_in_adomdb", flattenSystemGlobalSaveLastHitInAdomdb(o["save-last-hit-in-adomdb"], d, "save_last_hit_in_adomdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["save-last-hit-in-adomdb"], "SystemGlobal-SaveLastHitInAdomdb"); ok {
			if err = d.Set("save_last_hit_in_adomdb", vv); err != nil {
				return fmt.Errorf("Error reading save_last_hit_in_adomdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading save_last_hit_in_adomdb: %v", err)
		}
	}

	if err = d.Set("search_all_adoms", flattenSystemGlobalSearchAllAdoms(o["search-all-adoms"], d, "search_all_adoms")); err != nil {
		if vv, ok := fortiAPIPatch(o["search-all-adoms"], "SystemGlobal-SearchAllAdoms"); ok {
			if err = d.Set("search_all_adoms", vv); err != nil {
				return fmt.Errorf("Error reading search_all_adoms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading search_all_adoms: %v", err)
		}
	}

	if err = d.Set("ssh_enc_algo", flattenSystemGlobalSshEncAlgo(o["ssh-enc-algo"], d, "ssh_enc_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-enc-algo"], "SystemGlobal-SshEncAlgo"); ok {
			if err = d.Set("ssh_enc_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_enc_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_enc_algo: %v", err)
		}
	}

	if err = d.Set("ssh_hostkey_algo", flattenSystemGlobalSshHostkeyAlgo(o["ssh-hostkey-algo"], d, "ssh_hostkey_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-hostkey-algo"], "SystemGlobal-SshHostkeyAlgo"); ok {
			if err = d.Set("ssh_hostkey_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_hostkey_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_hostkey_algo: %v", err)
		}
	}

	if err = d.Set("ssh_kex_algo", flattenSystemGlobalSshKexAlgo(o["ssh-kex-algo"], d, "ssh_kex_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-kex-algo"], "SystemGlobal-SshKexAlgo"); ok {
			if err = d.Set("ssh_kex_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_kex_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_kex_algo: %v", err)
		}
	}

	if err = d.Set("ssh_mac_algo", flattenSystemGlobalSshMacAlgo(o["ssh-mac-algo"], d, "ssh_mac_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-mac-algo"], "SystemGlobal-SshMacAlgo"); ok {
			if err = d.Set("ssh_mac_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_mac_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_mac_algo: %v", err)
		}
	}

	if err = d.Set("ssh_strong_crypto", flattenSystemGlobalSshStrongCrypto(o["ssh-strong-crypto"], d, "ssh_strong_crypto")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-strong-crypto"], "SystemGlobal-SshStrongCrypto"); ok {
			if err = d.Set("ssh_strong_crypto", vv); err != nil {
				return fmt.Errorf("Error reading ssh_strong_crypto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_strong_crypto: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenSystemGlobalSslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "SystemGlobal-SslCipherSuites"); ok {
				if err = d.Set("ssl_cipher_suites", vv); err != nil {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenSystemGlobalSslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-cipher-suites"], "SystemGlobal-SslCipherSuites"); ok {
					if err = d.Set("ssl_cipher_suites", vv); err != nil {
						return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_low_encryption", flattenSystemGlobalSslLowEncryption(o["ssl-low-encryption"], d, "ssl_low_encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-low-encryption"], "SystemGlobal-SslLowEncryption"); ok {
			if err = d.Set("ssl_low_encryption", vv); err != nil {
				return fmt.Errorf("Error reading ssl_low_encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_low_encryption: %v", err)
		}
	}

	if err = d.Set("ssl_protocol", flattenSystemGlobalSslProtocol(o["ssl-protocol"], d, "ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-protocol"], "SystemGlobal-SslProtocol"); ok {
			if err = d.Set("ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_protocol: %v", err)
		}
	}

	if err = d.Set("ssl_static_key_ciphers", flattenSystemGlobalSslStaticKeyCiphers(o["ssl-static-key-ciphers"], d, "ssl_static_key_ciphers")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-static-key-ciphers"], "SystemGlobal-SslStaticKeyCiphers"); ok {
			if err = d.Set("ssl_static_key_ciphers", vv); err != nil {
				return fmt.Errorf("Error reading ssl_static_key_ciphers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_static_key_ciphers: %v", err)
		}
	}

	if err = d.Set("table_entry_blink", flattenSystemGlobalTableEntryBlink(o["table-entry-blink"], d, "table_entry_blink")); err != nil {
		if vv, ok := fortiAPIPatch(o["table-entry-blink"], "SystemGlobal-TableEntryBlink"); ok {
			if err = d.Set("table_entry_blink", vv); err != nil {
				return fmt.Errorf("Error reading table_entry_blink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading table_entry_blink: %v", err)
		}
	}

	if err = d.Set("task_list_size", flattenSystemGlobalTaskListSize(o["task-list-size"], d, "task_list_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["task-list-size"], "SystemGlobal-TaskListSize"); ok {
			if err = d.Set("task_list_size", vv); err != nil {
				return fmt.Errorf("Error reading task_list_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading task_list_size: %v", err)
		}
	}

	if err = d.Set("tftp", flattenSystemGlobalTftp(o["tftp"], d, "tftp")); err != nil {
		if vv, ok := fortiAPIPatch(o["tftp"], "SystemGlobal-Tftp"); ok {
			if err = d.Set("tftp", vv); err != nil {
				return fmt.Errorf("Error reading tftp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tftp: %v", err)
		}
	}

	if err = d.Set("timezone", flattenSystemGlobalTimezone(o["timezone"], d, "timezone")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone"], "SystemGlobal-Timezone"); ok {
			if err = d.Set("timezone", vv); err != nil {
				return fmt.Errorf("Error reading timezone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("tunnel_mtu", flattenSystemGlobalTunnelMtu(o["tunnel-mtu"], d, "tunnel_mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-mtu"], "SystemGlobal-TunnelMtu"); ok {
			if err = d.Set("tunnel_mtu", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_mtu: %v", err)
		}
	}

	if err = d.Set("usg", flattenSystemGlobalUsg(o["usg"], d, "usg")); err != nil {
		if vv, ok := fortiAPIPatch(o["usg"], "SystemGlobal-Usg"); ok {
			if err = d.Set("usg", vv); err != nil {
				return fmt.Errorf("Error reading usg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usg: %v", err)
		}
	}

	if err = d.Set("vdom_mirror", flattenSystemGlobalVdomMirror(o["vdom-mirror"], d, "vdom_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom-mirror"], "SystemGlobal-VdomMirror"); ok {
			if err = d.Set("vdom_mirror", vv); err != nil {
				return fmt.Errorf("Error reading vdom_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom_mirror: %v", err)
		}
	}

	if err = d.Set("webservice_proto", flattenSystemGlobalWebserviceProto(o["webservice-proto"], d, "webservice_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["webservice-proto"], "SystemGlobal-WebserviceProto"); ok {
			if err = d.Set("webservice_proto", vv); err != nil {
				return fmt.Errorf("Error reading webservice_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webservice_proto: %v", err)
		}
	}

	if err = d.Set("workflow_max_sessions", flattenSystemGlobalWorkflowMaxSessions(o["workflow-max-sessions"], d, "workflow_max_sessions")); err != nil {
		if vv, ok := fortiAPIPatch(o["workflow-max-sessions"], "SystemGlobal-WorkflowMaxSessions"); ok {
			if err = d.Set("workflow_max_sessions", vv); err != nil {
				return fmt.Errorf("Error reading workflow_max_sessions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading workflow_max_sessions: %v", err)
		}
	}

	if err = d.Set("workspace_mode", flattenSystemGlobalWorkspaceMode(o["workspace-mode"], d, "workspace_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["workspace-mode"], "SystemGlobal-WorkspaceMode"); ok {
			if err = d.Set("workspace_mode", vv); err != nil {
				return fmt.Errorf("Error reading workspace_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading workspace_mode: %v", err)
		}
	}

	if err = d.Set("workspace_unlock_after_install", flattenSystemGlobalWorkspaceUnlockAfterInstall(o["workspace-unlock-after-install"], d, "workspace_unlock_after_install")); err != nil {
		if vv, ok := fortiAPIPatch(o["workspace-unlock-after-install"], "SystemGlobal-WorkspaceUnlockAfterInstall"); ok {
			if err = d.Set("workspace_unlock_after_install", vv); err != nil {
				return fmt.Errorf("Error reading workspace_unlock_after_install: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading workspace_unlock_after_install: %v", err)
		}
	}

	return nil
}

func flattenSystemGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemGlobalAdminHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLockoutDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLockoutMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLockoutThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshGraceTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomRevAutoDelete(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomRevMaxBackupRevisions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomRevMaxDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomRevMaxRevisions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalApacheMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalApiIpBinding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCloneNameOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCltCertReq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalConsoleOutput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalContentpackFgtInstall(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCountryFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCreateRevision(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDaylightsavetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDetectUnregisteredLogDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDeviceViewMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhParams(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDisableModule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFabricStoragePoolQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFabricStoragePoolSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFazStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFcpCfgService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgfmCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgfmCertExclusive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgfmDenyUnknown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgfmLocalCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgfmPeercertWithoutsn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgfmSslProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiservicePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGlobalSslProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiCurlTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiPollingInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHaMemberAutoGrouping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHttpdSslProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalImportIgnoreAddrCmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalJsonapiLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLanguage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLatitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLdapCacheTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLdapconntimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLockPreempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogChecksum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogChecksumUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogForwardCacheSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLongitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalManagementIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalManagementPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMapclientSslProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMaxLogForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMaxRunningReports(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMcPolicyDisabledAdoms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adom_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adom-name"], _ = expandSystemGlobalMcPolicyDisabledAdomsAdomName(d, i["adom_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemGlobalMcPolicyDisabledAdomsAdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMultipleStepsUpgradeInAutolink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalNoCopyPermissionCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalNoVipValueCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalNormalizedIntfZoneOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalObjectRevisionDbMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalObjectRevisionMandatoryNote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalObjectRevisionObjectMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalObjectRevisionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalOftpSslProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPartialInstall(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPartialInstallForce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPartialInstallRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPerPolicyLock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPerformImproveByHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPolicyObjectIcon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPolicyObjectInDualPane(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPreLoginBanner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPreLoginBannerMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPrivateDataEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRemoteauthtimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSaveLastHitInAdomdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSearchAllAdoms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshEncAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSshHostkeyAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSshKexAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSshMacAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSshStrongCrypto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslCipherSuites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cipher"], _ = expandSystemGlobalSslCipherSuitesCipher(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandSystemGlobalSslCipherSuitesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandSystemGlobalSslCipherSuitesVersion(d, i["version"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemGlobalSslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslCipherSuitesVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslLowEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSslStaticKeyCiphers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTableEntryBlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTaskListSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTftp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTimezone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTunnelMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVdomMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWebserviceProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalWorkflowMaxSessions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWorkspaceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWorkspaceUnlockAfterInstall(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("admin_host"); ok || d.HasChange("admin_host") {
		t, err := expandSystemGlobalAdminHost(d, v, "admin_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-host"] = t
		}
	}

	if v, ok := d.GetOk("admin_lockout_duration"); ok || d.HasChange("admin_lockout_duration") {
		t, err := expandSystemGlobalAdminLockoutDuration(d, v, "admin_lockout_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-lockout-duration"] = t
		}
	}

	if v, ok := d.GetOk("admin_lockout_method"); ok || d.HasChange("admin_lockout_method") {
		t, err := expandSystemGlobalAdminLockoutMethod(d, v, "admin_lockout_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-lockout-method"] = t
		}
	}

	if v, ok := d.GetOk("admin_lockout_threshold"); ok || d.HasChange("admin_lockout_threshold") {
		t, err := expandSystemGlobalAdminLockoutThreshold(d, v, "admin_lockout_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-lockout-threshold"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_grace_time"); ok || d.HasChange("admin_ssh_grace_time") {
		t, err := expandSystemGlobalAdminSshGraceTime(d, v, "admin_ssh_grace_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-grace-time"] = t
		}
	}

	if v, ok := d.GetOk("adom_mode"); ok || d.HasChange("adom_mode") {
		t, err := expandSystemGlobalAdomMode(d, v, "adom_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-mode"] = t
		}
	}

	if v, ok := d.GetOk("adom_rev_auto_delete"); ok || d.HasChange("adom_rev_auto_delete") {
		t, err := expandSystemGlobalAdomRevAutoDelete(d, v, "adom_rev_auto_delete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-rev-auto-delete"] = t
		}
	}

	if v, ok := d.GetOk("adom_rev_max_backup_revisions"); ok || d.HasChange("adom_rev_max_backup_revisions") {
		t, err := expandSystemGlobalAdomRevMaxBackupRevisions(d, v, "adom_rev_max_backup_revisions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-rev-max-backup-revisions"] = t
		}
	}

	if v, ok := d.GetOk("adom_rev_max_days"); ok || d.HasChange("adom_rev_max_days") {
		t, err := expandSystemGlobalAdomRevMaxDays(d, v, "adom_rev_max_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-rev-max-days"] = t
		}
	}

	if v, ok := d.GetOk("adom_rev_max_revisions"); ok || d.HasChange("adom_rev_max_revisions") {
		t, err := expandSystemGlobalAdomRevMaxRevisions(d, v, "adom_rev_max_revisions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-rev-max-revisions"] = t
		}
	}

	if v, ok := d.GetOk("adom_select"); ok || d.HasChange("adom_select") {
		t, err := expandSystemGlobalAdomSelect(d, v, "adom_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-select"] = t
		}
	}

	if v, ok := d.GetOk("adom_status"); ok || d.HasChange("adom_status") {
		t, err := expandSystemGlobalAdomStatus(d, v, "adom_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-status"] = t
		}
	}

	if v, ok := d.GetOk("apache_mode"); ok || d.HasChange("apache_mode") {
		t, err := expandSystemGlobalApacheMode(d, v, "apache_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apache-mode"] = t
		}
	}

	if v, ok := d.GetOk("api_ip_binding"); ok || d.HasChange("api_ip_binding") {
		t, err := expandSystemGlobalApiIpBinding(d, v, "api_ip_binding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-ip-binding"] = t
		}
	}

	if v, ok := d.GetOk("clone_name_option"); ok || d.HasChange("clone_name_option") {
		t, err := expandSystemGlobalCloneNameOption(d, v, "clone_name_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clone-name-option"] = t
		}
	}

	if v, ok := d.GetOk("clt_cert_req"); ok || d.HasChange("clt_cert_req") {
		t, err := expandSystemGlobalCltCertReq(d, v, "clt_cert_req")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clt-cert-req"] = t
		}
	}

	if v, ok := d.GetOk("console_output"); ok || d.HasChange("console_output") {
		t, err := expandSystemGlobalConsoleOutput(d, v, "console_output")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["console-output"] = t
		}
	}

	if v, ok := d.GetOk("contentpack_fgt_install"); ok || d.HasChange("contentpack_fgt_install") {
		t, err := expandSystemGlobalContentpackFgtInstall(d, v, "contentpack_fgt_install")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contentpack-fgt-install"] = t
		}
	}

	if v, ok := d.GetOk("country_flag"); ok || d.HasChange("country_flag") {
		t, err := expandSystemGlobalCountryFlag(d, v, "country_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country-flag"] = t
		}
	}

	if v, ok := d.GetOk("create_revision"); ok || d.HasChange("create_revision") {
		t, err := expandSystemGlobalCreateRevision(d, v, "create_revision")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-revision"] = t
		}
	}

	if v, ok := d.GetOk("daylightsavetime"); ok || d.HasChange("daylightsavetime") {
		t, err := expandSystemGlobalDaylightsavetime(d, v, "daylightsavetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["daylightsavetime"] = t
		}
	}

	if v, ok := d.GetOk("detect_unregistered_log_device"); ok || d.HasChange("detect_unregistered_log_device") {
		t, err := expandSystemGlobalDetectUnregisteredLogDevice(d, v, "detect_unregistered_log_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-unregistered-log-device"] = t
		}
	}

	if v, ok := d.GetOk("device_view_mode"); ok || d.HasChange("device_view_mode") {
		t, err := expandSystemGlobalDeviceViewMode(d, v, "device_view_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-view-mode"] = t
		}
	}

	if v, ok := d.GetOk("dh_params"); ok || d.HasChange("dh_params") {
		t, err := expandSystemGlobalDhParams(d, v, "dh_params")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-params"] = t
		}
	}

	if v, ok := d.GetOk("disable_module"); ok || d.HasChange("disable_module") {
		t, err := expandSystemGlobalDisableModule(d, v, "disable_module")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disable-module"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok || d.HasChange("enc_algorithm") {
		t, err := expandSystemGlobalEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("fabric_storage_pool_quota"); ok || d.HasChange("fabric_storage_pool_quota") {
		t, err := expandSystemGlobalFabricStoragePoolQuota(d, v, "fabric_storage_pool_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-storage-pool-quota"] = t
		}
	}

	if v, ok := d.GetOk("fabric_storage_pool_size"); ok || d.HasChange("fabric_storage_pool_size") {
		t, err := expandSystemGlobalFabricStoragePoolSize(d, v, "fabric_storage_pool_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-storage-pool-size"] = t
		}
	}

	if v, ok := d.GetOk("faz_status"); ok || d.HasChange("faz_status") {
		t, err := expandSystemGlobalFazStatus(d, v, "faz_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-status"] = t
		}
	}

	if v, ok := d.GetOk("fcp_cfg_service"); ok || d.HasChange("fcp_cfg_service") {
		t, err := expandSystemGlobalFcpCfgService(d, v, "fcp_cfg_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fcp-cfg-service"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_ca_cert"); ok || d.HasChange("fgfm_ca_cert") {
		t, err := expandSystemGlobalFgfmCaCert(d, v, "fgfm_ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_cert_exclusive"); ok || d.HasChange("fgfm_cert_exclusive") {
		t, err := expandSystemGlobalFgfmCertExclusive(d, v, "fgfm_cert_exclusive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-cert-exclusive"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_deny_unknown"); ok || d.HasChange("fgfm_deny_unknown") {
		t, err := expandSystemGlobalFgfmDenyUnknown(d, v, "fgfm_deny_unknown")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-deny-unknown"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_local_cert"); ok || d.HasChange("fgfm_local_cert") {
		t, err := expandSystemGlobalFgfmLocalCert(d, v, "fgfm_local_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-local-cert"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_peercert_withoutsn"); ok || d.HasChange("fgfm_peercert_withoutsn") {
		t, err := expandSystemGlobalFgfmPeercertWithoutsn(d, v, "fgfm_peercert_withoutsn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-peercert-withoutsn"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_ssl_protocol"); ok || d.HasChange("fgfm_ssl_protocol") {
		t, err := expandSystemGlobalFgfmSslProtocol(d, v, "fgfm_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("fortiservice_port"); ok || d.HasChange("fortiservice_port") {
		t, err := expandSystemGlobalFortiservicePort(d, v, "fortiservice_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiservice-port"] = t
		}
	}

	if v, ok := d.GetOk("global_ssl_protocol"); ok || d.HasChange("global_ssl_protocol") {
		t, err := expandSystemGlobalGlobalSslProtocol(d, v, "global_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("gui_curl_timeout"); ok || d.HasChange("gui_curl_timeout") {
		t, err := expandSystemGlobalGuiCurlTimeout(d, v, "gui_curl_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-curl-timeout"] = t
		}
	}

	if v, ok := d.GetOk("gui_polling_interval"); ok || d.HasChange("gui_polling_interval") {
		t, err := expandSystemGlobalGuiPollingInterval(d, v, "gui_polling_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-polling-interval"] = t
		}
	}

	if v, ok := d.GetOk("ha_member_auto_grouping"); ok || d.HasChange("ha_member_auto_grouping") {
		t, err := expandSystemGlobalHaMemberAutoGrouping(d, v, "ha_member_auto_grouping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-member-auto-grouping"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandSystemGlobalHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("httpd_ssl_protocol"); ok || d.HasChange("httpd_ssl_protocol") {
		t, err := expandSystemGlobalHttpdSslProtocol(d, v, "httpd_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["httpd-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("import_ignore_addr_cmt"); ok || d.HasChange("import_ignore_addr_cmt") {
		t, err := expandSystemGlobalImportIgnoreAddrCmt(d, v, "import_ignore_addr_cmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["import-ignore-addr-cmt"] = t
		}
	}

	if v, ok := d.GetOk("jsonapi_log"); ok || d.HasChange("jsonapi_log") {
		t, err := expandSystemGlobalJsonapiLog(d, v, "jsonapi_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jsonapi-log"] = t
		}
	}

	if v, ok := d.GetOk("language"); ok || d.HasChange("language") {
		t, err := expandSystemGlobalLanguage(d, v, "language")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["language"] = t
		}
	}

	if v, ok := d.GetOk("latitude"); ok || d.HasChange("latitude") {
		t, err := expandSystemGlobalLatitude(d, v, "latitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latitude"] = t
		}
	}

	if v, ok := d.GetOk("ldap_cache_timeout"); ok || d.HasChange("ldap_cache_timeout") {
		t, err := expandSystemGlobalLdapCacheTimeout(d, v, "ldap_cache_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-cache-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ldapconntimeout"); ok || d.HasChange("ldapconntimeout") {
		t, err := expandSystemGlobalLdapconntimeout(d, v, "ldapconntimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldapconntimeout"] = t
		}
	}

	if v, ok := d.GetOk("lock_preempt"); ok || d.HasChange("lock_preempt") {
		t, err := expandSystemGlobalLockPreempt(d, v, "lock_preempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lock-preempt"] = t
		}
	}

	if v, ok := d.GetOk("log_checksum"); ok || d.HasChange("log_checksum") {
		t, err := expandSystemGlobalLogChecksum(d, v, "log_checksum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-checksum"] = t
		}
	}

	if v, ok := d.GetOk("log_checksum_upload"); ok || d.HasChange("log_checksum_upload") {
		t, err := expandSystemGlobalLogChecksumUpload(d, v, "log_checksum_upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-checksum-upload"] = t
		}
	}

	if v, ok := d.GetOk("log_forward_cache_size"); ok || d.HasChange("log_forward_cache_size") {
		t, err := expandSystemGlobalLogForwardCacheSize(d, v, "log_forward_cache_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-forward-cache-size"] = t
		}
	}

	if v, ok := d.GetOk("longitude"); ok || d.HasChange("longitude") {
		t, err := expandSystemGlobalLongitude(d, v, "longitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["longitude"] = t
		}
	}

	if v, ok := d.GetOk("management_ip"); ok || d.HasChange("management_ip") {
		t, err := expandSystemGlobalManagementIp(d, v, "management_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-ip"] = t
		}
	}

	if v, ok := d.GetOk("management_port"); ok || d.HasChange("management_port") {
		t, err := expandSystemGlobalManagementPort(d, v, "management_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-port"] = t
		}
	}

	if v, ok := d.GetOk("mapclient_ssl_protocol"); ok || d.HasChange("mapclient_ssl_protocol") {
		t, err := expandSystemGlobalMapclientSslProtocol(d, v, "mapclient_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapclient-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("max_log_forward"); ok || d.HasChange("max_log_forward") {
		t, err := expandSystemGlobalMaxLogForward(d, v, "max_log_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-forward"] = t
		}
	}

	if v, ok := d.GetOk("max_running_reports"); ok || d.HasChange("max_running_reports") {
		t, err := expandSystemGlobalMaxRunningReports(d, v, "max_running_reports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-running-reports"] = t
		}
	}

	if v, ok := d.GetOk("mc_policy_disabled_adoms"); ok || d.HasChange("mc_policy_disabled_adoms") {
		t, err := expandSystemGlobalMcPolicyDisabledAdoms(d, v, "mc_policy_disabled_adoms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mc-policy-disabled-adoms"] = t
		}
	}

	if v, ok := d.GetOk("multiple_steps_upgrade_in_autolink"); ok || d.HasChange("multiple_steps_upgrade_in_autolink") {
		t, err := expandSystemGlobalMultipleStepsUpgradeInAutolink(d, v, "multiple_steps_upgrade_in_autolink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-steps-upgrade-in-autolink"] = t
		}
	}

	if v, ok := d.GetOk("no_copy_permission_check"); ok || d.HasChange("no_copy_permission_check") {
		t, err := expandSystemGlobalNoCopyPermissionCheck(d, v, "no_copy_permission_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["no-copy-permission-check"] = t
		}
	}

	if v, ok := d.GetOk("no_vip_value_check"); ok || d.HasChange("no_vip_value_check") {
		t, err := expandSystemGlobalNoVipValueCheck(d, v, "no_vip_value_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["no-vip-value-check"] = t
		}
	}

	if v, ok := d.GetOk("normalized_intf_zone_only"); ok || d.HasChange("normalized_intf_zone_only") {
		t, err := expandSystemGlobalNormalizedIntfZoneOnly(d, v, "normalized_intf_zone_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["normalized-intf-zone-only"] = t
		}
	}

	if v, ok := d.GetOk("object_revision_db_max"); ok || d.HasChange("object_revision_db_max") {
		t, err := expandSystemGlobalObjectRevisionDbMax(d, v, "object_revision_db_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-revision-db-max"] = t
		}
	}

	if v, ok := d.GetOk("object_revision_mandatory_note"); ok || d.HasChange("object_revision_mandatory_note") {
		t, err := expandSystemGlobalObjectRevisionMandatoryNote(d, v, "object_revision_mandatory_note")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-revision-mandatory-note"] = t
		}
	}

	if v, ok := d.GetOk("object_revision_object_max"); ok || d.HasChange("object_revision_object_max") {
		t, err := expandSystemGlobalObjectRevisionObjectMax(d, v, "object_revision_object_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-revision-object-max"] = t
		}
	}

	if v, ok := d.GetOk("object_revision_status"); ok || d.HasChange("object_revision_status") {
		t, err := expandSystemGlobalObjectRevisionStatus(d, v, "object_revision_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-revision-status"] = t
		}
	}

	if v, ok := d.GetOk("oftp_ssl_protocol"); ok || d.HasChange("oftp_ssl_protocol") {
		t, err := expandSystemGlobalOftpSslProtocol(d, v, "oftp_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oftp-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("partial_install"); ok || d.HasChange("partial_install") {
		t, err := expandSystemGlobalPartialInstall(d, v, "partial_install")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["partial-install"] = t
		}
	}

	if v, ok := d.GetOk("partial_install_force"); ok || d.HasChange("partial_install_force") {
		t, err := expandSystemGlobalPartialInstallForce(d, v, "partial_install_force")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["partial-install-force"] = t
		}
	}

	if v, ok := d.GetOk("partial_install_rev"); ok || d.HasChange("partial_install_rev") {
		t, err := expandSystemGlobalPartialInstallRev(d, v, "partial_install_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["partial-install-rev"] = t
		}
	}

	if v, ok := d.GetOk("per_policy_lock"); ok || d.HasChange("per_policy_lock") {
		t, err := expandSystemGlobalPerPolicyLock(d, v, "per_policy_lock")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-policy-lock"] = t
		}
	}

	if v, ok := d.GetOk("perform_improve_by_ha"); ok || d.HasChange("perform_improve_by_ha") {
		t, err := expandSystemGlobalPerformImproveByHa(d, v, "perform_improve_by_ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["perform-improve-by-ha"] = t
		}
	}

	if v, ok := d.GetOk("policy_object_icon"); ok || d.HasChange("policy_object_icon") {
		t, err := expandSystemGlobalPolicyObjectIcon(d, v, "policy_object_icon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-object-icon"] = t
		}
	}

	if v, ok := d.GetOk("policy_object_in_dual_pane"); ok || d.HasChange("policy_object_in_dual_pane") {
		t, err := expandSystemGlobalPolicyObjectInDualPane(d, v, "policy_object_in_dual_pane")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-object-in-dual-pane"] = t
		}
	}

	if v, ok := d.GetOk("pre_login_banner"); ok || d.HasChange("pre_login_banner") {
		t, err := expandSystemGlobalPreLoginBanner(d, v, "pre_login_banner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pre-login-banner"] = t
		}
	}

	if v, ok := d.GetOk("pre_login_banner_message"); ok || d.HasChange("pre_login_banner_message") {
		t, err := expandSystemGlobalPreLoginBannerMessage(d, v, "pre_login_banner_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pre-login-banner-message"] = t
		}
	}

	if v, ok := d.GetOk("private_data_encryption"); ok || d.HasChange("private_data_encryption") {
		t, err := expandSystemGlobalPrivateDataEncryption(d, v, "private_data_encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-data-encryption"] = t
		}
	}

	if v, ok := d.GetOk("remoteauthtimeout"); ok || d.HasChange("remoteauthtimeout") {
		t, err := expandSystemGlobalRemoteauthtimeout(d, v, "remoteauthtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remoteauthtimeout"] = t
		}
	}

	if v, ok := d.GetOk("save_last_hit_in_adomdb"); ok || d.HasChange("save_last_hit_in_adomdb") {
		t, err := expandSystemGlobalSaveLastHitInAdomdb(d, v, "save_last_hit_in_adomdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["save-last-hit-in-adomdb"] = t
		}
	}

	if v, ok := d.GetOk("search_all_adoms"); ok || d.HasChange("search_all_adoms") {
		t, err := expandSystemGlobalSearchAllAdoms(d, v, "search_all_adoms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-all-adoms"] = t
		}
	}

	if v, ok := d.GetOk("ssh_enc_algo"); ok || d.HasChange("ssh_enc_algo") {
		t, err := expandSystemGlobalSshEncAlgo(d, v, "ssh_enc_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-enc-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hostkey_algo"); ok || d.HasChange("ssh_hostkey_algo") {
		t, err := expandSystemGlobalSshHostkeyAlgo(d, v, "ssh_hostkey_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hostkey-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_kex_algo"); ok || d.HasChange("ssh_kex_algo") {
		t, err := expandSystemGlobalSshKexAlgo(d, v, "ssh_kex_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-kex-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_mac_algo"); ok || d.HasChange("ssh_mac_algo") {
		t, err := expandSystemGlobalSshMacAlgo(d, v, "ssh_mac_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-mac-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_strong_crypto"); ok || d.HasChange("ssh_strong_crypto") {
		t, err := expandSystemGlobalSshStrongCrypto(d, v, "ssh_strong_crypto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-strong-crypto"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandSystemGlobalSslCipherSuites(d, v, "ssl_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_low_encryption"); ok || d.HasChange("ssl_low_encryption") {
		t, err := expandSystemGlobalSslLowEncryption(d, v, "ssl_low_encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-low-encryption"] = t
		}
	}

	if v, ok := d.GetOk("ssl_protocol"); ok || d.HasChange("ssl_protocol") {
		t, err := expandSystemGlobalSslProtocol(d, v, "ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("ssl_static_key_ciphers"); ok || d.HasChange("ssl_static_key_ciphers") {
		t, err := expandSystemGlobalSslStaticKeyCiphers(d, v, "ssl_static_key_ciphers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-static-key-ciphers"] = t
		}
	}

	if v, ok := d.GetOk("table_entry_blink"); ok || d.HasChange("table_entry_blink") {
		t, err := expandSystemGlobalTableEntryBlink(d, v, "table_entry_blink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-entry-blink"] = t
		}
	}

	if v, ok := d.GetOk("task_list_size"); ok || d.HasChange("task_list_size") {
		t, err := expandSystemGlobalTaskListSize(d, v, "task_list_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["task-list-size"] = t
		}
	}

	if v, ok := d.GetOk("tftp"); ok || d.HasChange("tftp") {
		t, err := expandSystemGlobalTftp(d, v, "tftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok || d.HasChange("timezone") {
		t, err := expandSystemGlobalTimezone(d, v, "timezone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_mtu"); ok || d.HasChange("tunnel_mtu") {
		t, err := expandSystemGlobalTunnelMtu(d, v, "tunnel_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-mtu"] = t
		}
	}

	if v, ok := d.GetOk("usg"); ok || d.HasChange("usg") {
		t, err := expandSystemGlobalUsg(d, v, "usg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usg"] = t
		}
	}

	if v, ok := d.GetOk("vdom_mirror"); ok || d.HasChange("vdom_mirror") {
		t, err := expandSystemGlobalVdomMirror(d, v, "vdom_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-mirror"] = t
		}
	}

	if v, ok := d.GetOk("webservice_proto"); ok || d.HasChange("webservice_proto") {
		t, err := expandSystemGlobalWebserviceProto(d, v, "webservice_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webservice-proto"] = t
		}
	}

	if v, ok := d.GetOk("workflow_max_sessions"); ok || d.HasChange("workflow_max_sessions") {
		t, err := expandSystemGlobalWorkflowMaxSessions(d, v, "workflow_max_sessions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workflow-max-sessions"] = t
		}
	}

	if v, ok := d.GetOk("workspace_mode"); ok || d.HasChange("workspace_mode") {
		t, err := expandSystemGlobalWorkspaceMode(d, v, "workspace_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workspace-mode"] = t
		}
	}

	if v, ok := d.GetOk("workspace_unlock_after_install"); ok || d.HasChange("workspace_unlock_after_install") {
		t, err := expandSystemGlobalWorkspaceUnlockAfterInstall(d, v, "workspace_unlock_after_install")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workspace-unlock-after-install"] = t
		}
	}

	return &obj, nil
}
