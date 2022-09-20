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
			"admin_lockout_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_lockout_threshold": &schema.Schema{
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
			"faz_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgfm_ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fgfm_local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fgfm_ssl_protocol": &schema.Schema{
				Type:     schema.TypeString,
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
			"import_ignore_addr_cmt": &schema.Schema{
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
			"log_forward_cache_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"longitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"search_all_adoms": &schema.Schema{
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

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemGlobal(obj, adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemGlobal(adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemGlobal(adomv, mkey, nil)
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

func flattenSystemGlobalAdminLockoutDurationSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdminLockoutThresholdSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomModeSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomRevAutoDeleteSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomRevMaxBackupRevisionsSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomRevMaxDaysSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomRevMaxRevisionsSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomSelectSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalAdomStatusSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCloneNameOptionSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCltCertReqSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalConsoleOutputSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCountryFlagSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalCreateRevisionSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDaylightsavetimeSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDetectUnregisteredLogDeviceSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDeviceViewModeSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDhParamsSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalDisableModuleSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalEncAlgorithmSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFazStatusSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFgfmCaCertSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFgfmLocalCertSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalFgfmSslProtocolSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalGuiPollingIntervalSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHaMemberAutoGroupingSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalHostnameSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalImportIgnoreAddrCmtSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLanguageSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLatitudeSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLdapCacheTimeoutSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLdapconntimeoutSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLockPreemptSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLogChecksumSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLogForwardCacheSizeSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalLongitudeSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMaxLogForwardSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMaxRunningReportsSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMcPolicyDisabledAdomsSga(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemGlobalMcPolicyDisabledAdomsAdomNameSga(i["adom-name"], d, pre_append)
			tmp["adom_name"] = fortiAPISubPartPatch(v, "SystemGlobal-McPolicyDisabledAdoms-AdomName")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemGlobalMcPolicyDisabledAdomsAdomNameSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalMultipleStepsUpgradeInAutolinkSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalNormalizedIntfZoneOnlySga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalObjectRevisionDbMaxSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalObjectRevisionMandatoryNoteSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalObjectRevisionObjectMaxSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalObjectRevisionStatusSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalOftpSslProtocolSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPartialInstallSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPartialInstallForceSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPartialInstallRevSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPerPolicyLockSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPerformImproveByHaSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPolicyObjectIconSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPolicyObjectInDualPaneSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPreLoginBannerSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPreLoginBannerMessageSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalPrivateDataEncryptionSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalRemoteauthtimeoutSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSearchAllAdomsSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslCipherSuitesSga(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemGlobalSslCipherSuitesCipherSga(i["cipher"], d, pre_append)
			tmp["cipher"] = fortiAPISubPartPatch(v, "SystemGlobal-SslCipherSuites-Cipher")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenSystemGlobalSslCipherSuitesPrioritySga(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "SystemGlobal-SslCipherSuites-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenSystemGlobalSslCipherSuitesVersionSga(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "SystemGlobal-SslCipherSuites-Version")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemGlobalSslCipherSuitesCipherSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslCipherSuitesPrioritySga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslCipherSuitesVersionSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslLowEncryptionSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalSslProtocolSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalSslStaticKeyCiphersSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTableEntryBlinkSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTaskListSizeSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTftpSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTimezoneSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalTunnelMtuSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalUsgSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalVdomMirrorSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWebserviceProtoSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGlobalWorkflowMaxSessionsSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGlobalWorkspaceModeSga(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("admin_lockout_duration", flattenSystemGlobalAdminLockoutDurationSga(o["admin-lockout-duration"], d, "admin_lockout_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-lockout-duration"], "SystemGlobal-AdminLockoutDuration"); ok {
			if err = d.Set("admin_lockout_duration", vv); err != nil {
				return fmt.Errorf("Error reading admin_lockout_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_lockout_duration: %v", err)
		}
	}

	if err = d.Set("admin_lockout_threshold", flattenSystemGlobalAdminLockoutThresholdSga(o["admin-lockout-threshold"], d, "admin_lockout_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-lockout-threshold"], "SystemGlobal-AdminLockoutThreshold"); ok {
			if err = d.Set("admin_lockout_threshold", vv); err != nil {
				return fmt.Errorf("Error reading admin_lockout_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_lockout_threshold: %v", err)
		}
	}

	if err = d.Set("adom_mode", flattenSystemGlobalAdomModeSga(o["adom-mode"], d, "adom_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-mode"], "SystemGlobal-AdomMode"); ok {
			if err = d.Set("adom_mode", vv); err != nil {
				return fmt.Errorf("Error reading adom_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_mode: %v", err)
		}
	}

	if err = d.Set("adom_rev_auto_delete", flattenSystemGlobalAdomRevAutoDeleteSga(o["adom-rev-auto-delete"], d, "adom_rev_auto_delete")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-rev-auto-delete"], "SystemGlobal-AdomRevAutoDelete"); ok {
			if err = d.Set("adom_rev_auto_delete", vv); err != nil {
				return fmt.Errorf("Error reading adom_rev_auto_delete: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_rev_auto_delete: %v", err)
		}
	}

	if err = d.Set("adom_rev_max_backup_revisions", flattenSystemGlobalAdomRevMaxBackupRevisionsSga(o["adom-rev-max-backup-revisions"], d, "adom_rev_max_backup_revisions")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-rev-max-backup-revisions"], "SystemGlobal-AdomRevMaxBackupRevisions"); ok {
			if err = d.Set("adom_rev_max_backup_revisions", vv); err != nil {
				return fmt.Errorf("Error reading adom_rev_max_backup_revisions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_rev_max_backup_revisions: %v", err)
		}
	}

	if err = d.Set("adom_rev_max_days", flattenSystemGlobalAdomRevMaxDaysSga(o["adom-rev-max-days"], d, "adom_rev_max_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-rev-max-days"], "SystemGlobal-AdomRevMaxDays"); ok {
			if err = d.Set("adom_rev_max_days", vv); err != nil {
				return fmt.Errorf("Error reading adom_rev_max_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_rev_max_days: %v", err)
		}
	}

	if err = d.Set("adom_rev_max_revisions", flattenSystemGlobalAdomRevMaxRevisionsSga(o["adom-rev-max-revisions"], d, "adom_rev_max_revisions")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-rev-max-revisions"], "SystemGlobal-AdomRevMaxRevisions"); ok {
			if err = d.Set("adom_rev_max_revisions", vv); err != nil {
				return fmt.Errorf("Error reading adom_rev_max_revisions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_rev_max_revisions: %v", err)
		}
	}

	if err = d.Set("adom_select", flattenSystemGlobalAdomSelectSga(o["adom-select"], d, "adom_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-select"], "SystemGlobal-AdomSelect"); ok {
			if err = d.Set("adom_select", vv); err != nil {
				return fmt.Errorf("Error reading adom_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_select: %v", err)
		}
	}

	if err = d.Set("adom_status", flattenSystemGlobalAdomStatusSga(o["adom-status"], d, "adom_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-status"], "SystemGlobal-AdomStatus"); ok {
			if err = d.Set("adom_status", vv); err != nil {
				return fmt.Errorf("Error reading adom_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_status: %v", err)
		}
	}

	if err = d.Set("clone_name_option", flattenSystemGlobalCloneNameOptionSga(o["clone-name-option"], d, "clone_name_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["clone-name-option"], "SystemGlobal-CloneNameOption"); ok {
			if err = d.Set("clone_name_option", vv); err != nil {
				return fmt.Errorf("Error reading clone_name_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clone_name_option: %v", err)
		}
	}

	if err = d.Set("clt_cert_req", flattenSystemGlobalCltCertReqSga(o["clt-cert-req"], d, "clt_cert_req")); err != nil {
		if vv, ok := fortiAPIPatch(o["clt-cert-req"], "SystemGlobal-CltCertReq"); ok {
			if err = d.Set("clt_cert_req", vv); err != nil {
				return fmt.Errorf("Error reading clt_cert_req: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clt_cert_req: %v", err)
		}
	}

	if err = d.Set("console_output", flattenSystemGlobalConsoleOutputSga(o["console-output"], d, "console_output")); err != nil {
		if vv, ok := fortiAPIPatch(o["console-output"], "SystemGlobal-ConsoleOutput"); ok {
			if err = d.Set("console_output", vv); err != nil {
				return fmt.Errorf("Error reading console_output: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading console_output: %v", err)
		}
	}

	if err = d.Set("country_flag", flattenSystemGlobalCountryFlagSga(o["country-flag"], d, "country_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["country-flag"], "SystemGlobal-CountryFlag"); ok {
			if err = d.Set("country_flag", vv); err != nil {
				return fmt.Errorf("Error reading country_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country_flag: %v", err)
		}
	}

	if err = d.Set("create_revision", flattenSystemGlobalCreateRevisionSga(o["create-revision"], d, "create_revision")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-revision"], "SystemGlobal-CreateRevision"); ok {
			if err = d.Set("create_revision", vv); err != nil {
				return fmt.Errorf("Error reading create_revision: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_revision: %v", err)
		}
	}

	if err = d.Set("daylightsavetime", flattenSystemGlobalDaylightsavetimeSga(o["daylightsavetime"], d, "daylightsavetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["daylightsavetime"], "SystemGlobal-Daylightsavetime"); ok {
			if err = d.Set("daylightsavetime", vv); err != nil {
				return fmt.Errorf("Error reading daylightsavetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading daylightsavetime: %v", err)
		}
	}

	if err = d.Set("detect_unregistered_log_device", flattenSystemGlobalDetectUnregisteredLogDeviceSga(o["detect-unregistered-log-device"], d, "detect_unregistered_log_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-unregistered-log-device"], "SystemGlobal-DetectUnregisteredLogDevice"); ok {
			if err = d.Set("detect_unregistered_log_device", vv); err != nil {
				return fmt.Errorf("Error reading detect_unregistered_log_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_unregistered_log_device: %v", err)
		}
	}

	if err = d.Set("device_view_mode", flattenSystemGlobalDeviceViewModeSga(o["device-view-mode"], d, "device_view_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-view-mode"], "SystemGlobal-DeviceViewMode"); ok {
			if err = d.Set("device_view_mode", vv); err != nil {
				return fmt.Errorf("Error reading device_view_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_view_mode: %v", err)
		}
	}

	if err = d.Set("dh_params", flattenSystemGlobalDhParamsSga(o["dh-params"], d, "dh_params")); err != nil {
		if vv, ok := fortiAPIPatch(o["dh-params"], "SystemGlobal-DhParams"); ok {
			if err = d.Set("dh_params", vv); err != nil {
				return fmt.Errorf("Error reading dh_params: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dh_params: %v", err)
		}
	}

	if err = d.Set("disable_module", flattenSystemGlobalDisableModuleSga(o["disable-module"], d, "disable_module")); err != nil {
		if vv, ok := fortiAPIPatch(o["disable-module"], "SystemGlobal-DisableModule"); ok {
			if err = d.Set("disable_module", vv); err != nil {
				return fmt.Errorf("Error reading disable_module: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disable_module: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenSystemGlobalEncAlgorithmSga(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["enc-algorithm"], "SystemGlobal-EncAlgorithm"); ok {
			if err = d.Set("enc_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading enc_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("faz_status", flattenSystemGlobalFazStatusSga(o["faz-status"], d, "faz_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-status"], "SystemGlobal-FazStatus"); ok {
			if err = d.Set("faz_status", vv); err != nil {
				return fmt.Errorf("Error reading faz_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_status: %v", err)
		}
	}

	if err = d.Set("fgfm_ca_cert", flattenSystemGlobalFgfmCaCertSga(o["fgfm-ca-cert"], d, "fgfm_ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-ca-cert"], "SystemGlobal-FgfmCaCert"); ok {
			if err = d.Set("fgfm_ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_ca_cert: %v", err)
		}
	}

	if err = d.Set("fgfm_local_cert", flattenSystemGlobalFgfmLocalCertSga(o["fgfm-local-cert"], d, "fgfm_local_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-local-cert"], "SystemGlobal-FgfmLocalCert"); ok {
			if err = d.Set("fgfm_local_cert", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_local_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_local_cert: %v", err)
		}
	}

	if err = d.Set("fgfm_ssl_protocol", flattenSystemGlobalFgfmSslProtocolSga(o["fgfm-ssl-protocol"], d, "fgfm_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm-ssl-protocol"], "SystemGlobal-FgfmSslProtocol"); ok {
			if err = d.Set("fgfm_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading fgfm_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("gui_polling_interval", flattenSystemGlobalGuiPollingIntervalSga(o["gui-polling-interval"], d, "gui_polling_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-polling-interval"], "SystemGlobal-GuiPollingInterval"); ok {
			if err = d.Set("gui_polling_interval", vv); err != nil {
				return fmt.Errorf("Error reading gui_polling_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_polling_interval: %v", err)
		}
	}

	if err = d.Set("ha_member_auto_grouping", flattenSystemGlobalHaMemberAutoGroupingSga(o["ha-member-auto-grouping"], d, "ha_member_auto_grouping")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-member-auto-grouping"], "SystemGlobal-HaMemberAutoGrouping"); ok {
			if err = d.Set("ha_member_auto_grouping", vv); err != nil {
				return fmt.Errorf("Error reading ha_member_auto_grouping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_member_auto_grouping: %v", err)
		}
	}

	if err = d.Set("hostname", flattenSystemGlobalHostnameSga(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "SystemGlobal-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("import_ignore_addr_cmt", flattenSystemGlobalImportIgnoreAddrCmtSga(o["import-ignore-addr-cmt"], d, "import_ignore_addr_cmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["import-ignore-addr-cmt"], "SystemGlobal-ImportIgnoreAddrCmt"); ok {
			if err = d.Set("import_ignore_addr_cmt", vv); err != nil {
				return fmt.Errorf("Error reading import_ignore_addr_cmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading import_ignore_addr_cmt: %v", err)
		}
	}

	if err = d.Set("language", flattenSystemGlobalLanguageSga(o["language"], d, "language")); err != nil {
		if vv, ok := fortiAPIPatch(o["language"], "SystemGlobal-Language"); ok {
			if err = d.Set("language", vv); err != nil {
				return fmt.Errorf("Error reading language: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("latitude", flattenSystemGlobalLatitudeSga(o["latitude"], d, "latitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["latitude"], "SystemGlobal-Latitude"); ok {
			if err = d.Set("latitude", vv); err != nil {
				return fmt.Errorf("Error reading latitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latitude: %v", err)
		}
	}

	if err = d.Set("ldap_cache_timeout", flattenSystemGlobalLdapCacheTimeoutSga(o["ldap-cache-timeout"], d, "ldap_cache_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-cache-timeout"], "SystemGlobal-LdapCacheTimeout"); ok {
			if err = d.Set("ldap_cache_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ldap_cache_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_cache_timeout: %v", err)
		}
	}

	if err = d.Set("ldapconntimeout", flattenSystemGlobalLdapconntimeoutSga(o["ldapconntimeout"], d, "ldapconntimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldapconntimeout"], "SystemGlobal-Ldapconntimeout"); ok {
			if err = d.Set("ldapconntimeout", vv); err != nil {
				return fmt.Errorf("Error reading ldapconntimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldapconntimeout: %v", err)
		}
	}

	if err = d.Set("lock_preempt", flattenSystemGlobalLockPreemptSga(o["lock-preempt"], d, "lock_preempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["lock-preempt"], "SystemGlobal-LockPreempt"); ok {
			if err = d.Set("lock_preempt", vv); err != nil {
				return fmt.Errorf("Error reading lock_preempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lock_preempt: %v", err)
		}
	}

	if err = d.Set("log_checksum", flattenSystemGlobalLogChecksumSga(o["log-checksum"], d, "log_checksum")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-checksum"], "SystemGlobal-LogChecksum"); ok {
			if err = d.Set("log_checksum", vv); err != nil {
				return fmt.Errorf("Error reading log_checksum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_checksum: %v", err)
		}
	}

	if err = d.Set("log_forward_cache_size", flattenSystemGlobalLogForwardCacheSizeSga(o["log-forward-cache-size"], d, "log_forward_cache_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-forward-cache-size"], "SystemGlobal-LogForwardCacheSize"); ok {
			if err = d.Set("log_forward_cache_size", vv); err != nil {
				return fmt.Errorf("Error reading log_forward_cache_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_forward_cache_size: %v", err)
		}
	}

	if err = d.Set("longitude", flattenSystemGlobalLongitudeSga(o["longitude"], d, "longitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["longitude"], "SystemGlobal-Longitude"); ok {
			if err = d.Set("longitude", vv); err != nil {
				return fmt.Errorf("Error reading longitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading longitude: %v", err)
		}
	}

	if err = d.Set("max_log_forward", flattenSystemGlobalMaxLogForwardSga(o["max-log-forward"], d, "max_log_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-forward"], "SystemGlobal-MaxLogForward"); ok {
			if err = d.Set("max_log_forward", vv); err != nil {
				return fmt.Errorf("Error reading max_log_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_forward: %v", err)
		}
	}

	if err = d.Set("max_running_reports", flattenSystemGlobalMaxRunningReportsSga(o["max-running-reports"], d, "max_running_reports")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-running-reports"], "SystemGlobal-MaxRunningReports"); ok {
			if err = d.Set("max_running_reports", vv); err != nil {
				return fmt.Errorf("Error reading max_running_reports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_running_reports: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mc_policy_disabled_adoms", flattenSystemGlobalMcPolicyDisabledAdomsSga(o["mc-policy-disabled-adoms"], d, "mc_policy_disabled_adoms")); err != nil {
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
			if err = d.Set("mc_policy_disabled_adoms", flattenSystemGlobalMcPolicyDisabledAdomsSga(o["mc-policy-disabled-adoms"], d, "mc_policy_disabled_adoms")); err != nil {
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

	if err = d.Set("multiple_steps_upgrade_in_autolink", flattenSystemGlobalMultipleStepsUpgradeInAutolinkSga(o["multiple-steps-upgrade-in-autolink"], d, "multiple_steps_upgrade_in_autolink")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiple-steps-upgrade-in-autolink"], "SystemGlobal-MultipleStepsUpgradeInAutolink"); ok {
			if err = d.Set("multiple_steps_upgrade_in_autolink", vv); err != nil {
				return fmt.Errorf("Error reading multiple_steps_upgrade_in_autolink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiple_steps_upgrade_in_autolink: %v", err)
		}
	}

	if err = d.Set("normalized_intf_zone_only", flattenSystemGlobalNormalizedIntfZoneOnlySga(o["normalized-intf-zone-only"], d, "normalized_intf_zone_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["normalized-intf-zone-only"], "SystemGlobal-NormalizedIntfZoneOnly"); ok {
			if err = d.Set("normalized_intf_zone_only", vv); err != nil {
				return fmt.Errorf("Error reading normalized_intf_zone_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading normalized_intf_zone_only: %v", err)
		}
	}

	if err = d.Set("object_revision_db_max", flattenSystemGlobalObjectRevisionDbMaxSga(o["object-revision-db-max"], d, "object_revision_db_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-revision-db-max"], "SystemGlobal-ObjectRevisionDbMax"); ok {
			if err = d.Set("object_revision_db_max", vv); err != nil {
				return fmt.Errorf("Error reading object_revision_db_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_revision_db_max: %v", err)
		}
	}

	if err = d.Set("object_revision_mandatory_note", flattenSystemGlobalObjectRevisionMandatoryNoteSga(o["object-revision-mandatory-note"], d, "object_revision_mandatory_note")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-revision-mandatory-note"], "SystemGlobal-ObjectRevisionMandatoryNote"); ok {
			if err = d.Set("object_revision_mandatory_note", vv); err != nil {
				return fmt.Errorf("Error reading object_revision_mandatory_note: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_revision_mandatory_note: %v", err)
		}
	}

	if err = d.Set("object_revision_object_max", flattenSystemGlobalObjectRevisionObjectMaxSga(o["object-revision-object-max"], d, "object_revision_object_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-revision-object-max"], "SystemGlobal-ObjectRevisionObjectMax"); ok {
			if err = d.Set("object_revision_object_max", vv); err != nil {
				return fmt.Errorf("Error reading object_revision_object_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_revision_object_max: %v", err)
		}
	}

	if err = d.Set("object_revision_status", flattenSystemGlobalObjectRevisionStatusSga(o["object-revision-status"], d, "object_revision_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-revision-status"], "SystemGlobal-ObjectRevisionStatus"); ok {
			if err = d.Set("object_revision_status", vv); err != nil {
				return fmt.Errorf("Error reading object_revision_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_revision_status: %v", err)
		}
	}

	if err = d.Set("oftp_ssl_protocol", flattenSystemGlobalOftpSslProtocolSga(o["oftp-ssl-protocol"], d, "oftp_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["oftp-ssl-protocol"], "SystemGlobal-OftpSslProtocol"); ok {
			if err = d.Set("oftp_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading oftp_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oftp_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("partial_install", flattenSystemGlobalPartialInstallSga(o["partial-install"], d, "partial_install")); err != nil {
		if vv, ok := fortiAPIPatch(o["partial-install"], "SystemGlobal-PartialInstall"); ok {
			if err = d.Set("partial_install", vv); err != nil {
				return fmt.Errorf("Error reading partial_install: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading partial_install: %v", err)
		}
	}

	if err = d.Set("partial_install_force", flattenSystemGlobalPartialInstallForceSga(o["partial-install-force"], d, "partial_install_force")); err != nil {
		if vv, ok := fortiAPIPatch(o["partial-install-force"], "SystemGlobal-PartialInstallForce"); ok {
			if err = d.Set("partial_install_force", vv); err != nil {
				return fmt.Errorf("Error reading partial_install_force: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading partial_install_force: %v", err)
		}
	}

	if err = d.Set("partial_install_rev", flattenSystemGlobalPartialInstallRevSga(o["partial-install-rev"], d, "partial_install_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["partial-install-rev"], "SystemGlobal-PartialInstallRev"); ok {
			if err = d.Set("partial_install_rev", vv); err != nil {
				return fmt.Errorf("Error reading partial_install_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading partial_install_rev: %v", err)
		}
	}

	if err = d.Set("per_policy_lock", flattenSystemGlobalPerPolicyLockSga(o["per-policy-lock"], d, "per_policy_lock")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-policy-lock"], "SystemGlobal-PerPolicyLock"); ok {
			if err = d.Set("per_policy_lock", vv); err != nil {
				return fmt.Errorf("Error reading per_policy_lock: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_policy_lock: %v", err)
		}
	}

	if err = d.Set("perform_improve_by_ha", flattenSystemGlobalPerformImproveByHaSga(o["perform-improve-by-ha"], d, "perform_improve_by_ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["perform-improve-by-ha"], "SystemGlobal-PerformImproveByHa"); ok {
			if err = d.Set("perform_improve_by_ha", vv); err != nil {
				return fmt.Errorf("Error reading perform_improve_by_ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading perform_improve_by_ha: %v", err)
		}
	}

	if err = d.Set("policy_object_icon", flattenSystemGlobalPolicyObjectIconSga(o["policy-object-icon"], d, "policy_object_icon")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-object-icon"], "SystemGlobal-PolicyObjectIcon"); ok {
			if err = d.Set("policy_object_icon", vv); err != nil {
				return fmt.Errorf("Error reading policy_object_icon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_object_icon: %v", err)
		}
	}

	if err = d.Set("policy_object_in_dual_pane", flattenSystemGlobalPolicyObjectInDualPaneSga(o["policy-object-in-dual-pane"], d, "policy_object_in_dual_pane")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-object-in-dual-pane"], "SystemGlobal-PolicyObjectInDualPane"); ok {
			if err = d.Set("policy_object_in_dual_pane", vv); err != nil {
				return fmt.Errorf("Error reading policy_object_in_dual_pane: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_object_in_dual_pane: %v", err)
		}
	}

	if err = d.Set("pre_login_banner", flattenSystemGlobalPreLoginBannerSga(o["pre-login-banner"], d, "pre_login_banner")); err != nil {
		if vv, ok := fortiAPIPatch(o["pre-login-banner"], "SystemGlobal-PreLoginBanner"); ok {
			if err = d.Set("pre_login_banner", vv); err != nil {
				return fmt.Errorf("Error reading pre_login_banner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pre_login_banner: %v", err)
		}
	}

	if err = d.Set("pre_login_banner_message", flattenSystemGlobalPreLoginBannerMessageSga(o["pre-login-banner-message"], d, "pre_login_banner_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["pre-login-banner-message"], "SystemGlobal-PreLoginBannerMessage"); ok {
			if err = d.Set("pre_login_banner_message", vv); err != nil {
				return fmt.Errorf("Error reading pre_login_banner_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pre_login_banner_message: %v", err)
		}
	}

	if err = d.Set("private_data_encryption", flattenSystemGlobalPrivateDataEncryptionSga(o["private-data-encryption"], d, "private_data_encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-data-encryption"], "SystemGlobal-PrivateDataEncryption"); ok {
			if err = d.Set("private_data_encryption", vv); err != nil {
				return fmt.Errorf("Error reading private_data_encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_data_encryption: %v", err)
		}
	}

	if err = d.Set("remoteauthtimeout", flattenSystemGlobalRemoteauthtimeoutSga(o["remoteauthtimeout"], d, "remoteauthtimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["remoteauthtimeout"], "SystemGlobal-Remoteauthtimeout"); ok {
			if err = d.Set("remoteauthtimeout", vv); err != nil {
				return fmt.Errorf("Error reading remoteauthtimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remoteauthtimeout: %v", err)
		}
	}

	if err = d.Set("search_all_adoms", flattenSystemGlobalSearchAllAdomsSga(o["search-all-adoms"], d, "search_all_adoms")); err != nil {
		if vv, ok := fortiAPIPatch(o["search-all-adoms"], "SystemGlobal-SearchAllAdoms"); ok {
			if err = d.Set("search_all_adoms", vv); err != nil {
				return fmt.Errorf("Error reading search_all_adoms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading search_all_adoms: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenSystemGlobalSslCipherSuitesSga(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
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
			if err = d.Set("ssl_cipher_suites", flattenSystemGlobalSslCipherSuitesSga(o["ssl-cipher-suites"], d, "ssl_cipher_suites")); err != nil {
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

	if err = d.Set("ssl_low_encryption", flattenSystemGlobalSslLowEncryptionSga(o["ssl-low-encryption"], d, "ssl_low_encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-low-encryption"], "SystemGlobal-SslLowEncryption"); ok {
			if err = d.Set("ssl_low_encryption", vv); err != nil {
				return fmt.Errorf("Error reading ssl_low_encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_low_encryption: %v", err)
		}
	}

	if err = d.Set("ssl_protocol", flattenSystemGlobalSslProtocolSga(o["ssl-protocol"], d, "ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-protocol"], "SystemGlobal-SslProtocol"); ok {
			if err = d.Set("ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_protocol: %v", err)
		}
	}

	if err = d.Set("ssl_static_key_ciphers", flattenSystemGlobalSslStaticKeyCiphersSga(o["ssl-static-key-ciphers"], d, "ssl_static_key_ciphers")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-static-key-ciphers"], "SystemGlobal-SslStaticKeyCiphers"); ok {
			if err = d.Set("ssl_static_key_ciphers", vv); err != nil {
				return fmt.Errorf("Error reading ssl_static_key_ciphers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_static_key_ciphers: %v", err)
		}
	}

	if err = d.Set("table_entry_blink", flattenSystemGlobalTableEntryBlinkSga(o["table-entry-blink"], d, "table_entry_blink")); err != nil {
		if vv, ok := fortiAPIPatch(o["table-entry-blink"], "SystemGlobal-TableEntryBlink"); ok {
			if err = d.Set("table_entry_blink", vv); err != nil {
				return fmt.Errorf("Error reading table_entry_blink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading table_entry_blink: %v", err)
		}
	}

	if err = d.Set("task_list_size", flattenSystemGlobalTaskListSizeSga(o["task-list-size"], d, "task_list_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["task-list-size"], "SystemGlobal-TaskListSize"); ok {
			if err = d.Set("task_list_size", vv); err != nil {
				return fmt.Errorf("Error reading task_list_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading task_list_size: %v", err)
		}
	}

	if err = d.Set("tftp", flattenSystemGlobalTftpSga(o["tftp"], d, "tftp")); err != nil {
		if vv, ok := fortiAPIPatch(o["tftp"], "SystemGlobal-Tftp"); ok {
			if err = d.Set("tftp", vv); err != nil {
				return fmt.Errorf("Error reading tftp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tftp: %v", err)
		}
	}

	if err = d.Set("timezone", flattenSystemGlobalTimezoneSga(o["timezone"], d, "timezone")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone"], "SystemGlobal-Timezone"); ok {
			if err = d.Set("timezone", vv); err != nil {
				return fmt.Errorf("Error reading timezone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("tunnel_mtu", flattenSystemGlobalTunnelMtuSga(o["tunnel-mtu"], d, "tunnel_mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-mtu"], "SystemGlobal-TunnelMtu"); ok {
			if err = d.Set("tunnel_mtu", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_mtu: %v", err)
		}
	}

	if err = d.Set("usg", flattenSystemGlobalUsgSga(o["usg"], d, "usg")); err != nil {
		if vv, ok := fortiAPIPatch(o["usg"], "SystemGlobal-Usg"); ok {
			if err = d.Set("usg", vv); err != nil {
				return fmt.Errorf("Error reading usg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usg: %v", err)
		}
	}

	if err = d.Set("vdom_mirror", flattenSystemGlobalVdomMirrorSga(o["vdom-mirror"], d, "vdom_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom-mirror"], "SystemGlobal-VdomMirror"); ok {
			if err = d.Set("vdom_mirror", vv); err != nil {
				return fmt.Errorf("Error reading vdom_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom_mirror: %v", err)
		}
	}

	if err = d.Set("webservice_proto", flattenSystemGlobalWebserviceProtoSga(o["webservice-proto"], d, "webservice_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["webservice-proto"], "SystemGlobal-WebserviceProto"); ok {
			if err = d.Set("webservice_proto", vv); err != nil {
				return fmt.Errorf("Error reading webservice_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webservice_proto: %v", err)
		}
	}

	if err = d.Set("workflow_max_sessions", flattenSystemGlobalWorkflowMaxSessionsSga(o["workflow-max-sessions"], d, "workflow_max_sessions")); err != nil {
		if vv, ok := fortiAPIPatch(o["workflow-max-sessions"], "SystemGlobal-WorkflowMaxSessions"); ok {
			if err = d.Set("workflow_max_sessions", vv); err != nil {
				return fmt.Errorf("Error reading workflow_max_sessions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading workflow_max_sessions: %v", err)
		}
	}

	if err = d.Set("workspace_mode", flattenSystemGlobalWorkspaceModeSga(o["workspace-mode"], d, "workspace_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["workspace-mode"], "SystemGlobal-WorkspaceMode"); ok {
			if err = d.Set("workspace_mode", vv); err != nil {
				return fmt.Errorf("Error reading workspace_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading workspace_mode: %v", err)
		}
	}

	return nil
}

func flattenSystemGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemGlobalAdminLockoutDurationSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLockoutThresholdSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomModeSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomRevAutoDeleteSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomRevMaxBackupRevisionsSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomRevMaxDaysSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomRevMaxRevisionsSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomSelectSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdomStatusSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCloneNameOptionSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCltCertReqSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalConsoleOutputSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCountryFlagSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCreateRevisionSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDaylightsavetimeSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDetectUnregisteredLogDeviceSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDeviceViewModeSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhParamsSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDisableModuleSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalEncAlgorithmSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFazStatusSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgfmCaCertSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgfmLocalCertSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgfmSslProtocolSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiPollingIntervalSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHaMemberAutoGroupingSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHostnameSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalImportIgnoreAddrCmtSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLanguageSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLatitudeSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLdapCacheTimeoutSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLdapconntimeoutSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLockPreemptSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogChecksumSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogForwardCacheSizeSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLongitudeSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMaxLogForwardSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMaxRunningReportsSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMcPolicyDisabledAdomsSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["adom-name"], _ = expandSystemGlobalMcPolicyDisabledAdomsAdomNameSga(d, i["adom_name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemGlobalMcPolicyDisabledAdomsAdomNameSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMultipleStepsUpgradeInAutolinkSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalNormalizedIntfZoneOnlySga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalObjectRevisionDbMaxSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalObjectRevisionMandatoryNoteSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalObjectRevisionObjectMaxSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalObjectRevisionStatusSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalOftpSslProtocolSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPartialInstallSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPartialInstallForceSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPartialInstallRevSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPerPolicyLockSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPerformImproveByHaSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPolicyObjectIconSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPolicyObjectInDualPaneSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPreLoginBannerSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPreLoginBannerMessageSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPrivateDataEncryptionSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRemoteauthtimeoutSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSearchAllAdomsSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslCipherSuitesSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cipher"], _ = expandSystemGlobalSslCipherSuitesCipherSga(d, i["cipher"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandSystemGlobalSslCipherSuitesPrioritySga(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandSystemGlobalSslCipherSuitesVersionSga(d, i["version"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemGlobalSslCipherSuitesCipherSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslCipherSuitesPrioritySga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslCipherSuitesVersionSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslLowEncryptionSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslProtocolSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalSslStaticKeyCiphersSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTableEntryBlinkSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTaskListSizeSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTftpSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTimezoneSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTunnelMtuSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUsgSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVdomMirrorSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWebserviceProtoSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGlobalWorkflowMaxSessionsSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWorkspaceModeSga(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("admin_lockout_duration"); ok || d.HasChange("admin_lockout_duration") {
		t, err := expandSystemGlobalAdminLockoutDurationSga(d, v, "admin_lockout_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-lockout-duration"] = t
		}
	}

	if v, ok := d.GetOk("admin_lockout_threshold"); ok || d.HasChange("admin_lockout_threshold") {
		t, err := expandSystemGlobalAdminLockoutThresholdSga(d, v, "admin_lockout_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-lockout-threshold"] = t
		}
	}

	if v, ok := d.GetOk("adom_mode"); ok || d.HasChange("adom_mode") {
		t, err := expandSystemGlobalAdomModeSga(d, v, "adom_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-mode"] = t
		}
	}

	if v, ok := d.GetOk("adom_rev_auto_delete"); ok || d.HasChange("adom_rev_auto_delete") {
		t, err := expandSystemGlobalAdomRevAutoDeleteSga(d, v, "adom_rev_auto_delete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-rev-auto-delete"] = t
		}
	}

	if v, ok := d.GetOk("adom_rev_max_backup_revisions"); ok || d.HasChange("adom_rev_max_backup_revisions") {
		t, err := expandSystemGlobalAdomRevMaxBackupRevisionsSga(d, v, "adom_rev_max_backup_revisions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-rev-max-backup-revisions"] = t
		}
	}

	if v, ok := d.GetOk("adom_rev_max_days"); ok || d.HasChange("adom_rev_max_days") {
		t, err := expandSystemGlobalAdomRevMaxDaysSga(d, v, "adom_rev_max_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-rev-max-days"] = t
		}
	}

	if v, ok := d.GetOk("adom_rev_max_revisions"); ok || d.HasChange("adom_rev_max_revisions") {
		t, err := expandSystemGlobalAdomRevMaxRevisionsSga(d, v, "adom_rev_max_revisions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-rev-max-revisions"] = t
		}
	}

	if v, ok := d.GetOk("adom_select"); ok || d.HasChange("adom_select") {
		t, err := expandSystemGlobalAdomSelectSga(d, v, "adom_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-select"] = t
		}
	}

	if v, ok := d.GetOk("adom_status"); ok || d.HasChange("adom_status") {
		t, err := expandSystemGlobalAdomStatusSga(d, v, "adom_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-status"] = t
		}
	}

	if v, ok := d.GetOk("clone_name_option"); ok || d.HasChange("clone_name_option") {
		t, err := expandSystemGlobalCloneNameOptionSga(d, v, "clone_name_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clone-name-option"] = t
		}
	}

	if v, ok := d.GetOk("clt_cert_req"); ok || d.HasChange("clt_cert_req") {
		t, err := expandSystemGlobalCltCertReqSga(d, v, "clt_cert_req")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clt-cert-req"] = t
		}
	}

	if v, ok := d.GetOk("console_output"); ok || d.HasChange("console_output") {
		t, err := expandSystemGlobalConsoleOutputSga(d, v, "console_output")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["console-output"] = t
		}
	}

	if v, ok := d.GetOk("country_flag"); ok || d.HasChange("country_flag") {
		t, err := expandSystemGlobalCountryFlagSga(d, v, "country_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country-flag"] = t
		}
	}

	if v, ok := d.GetOk("create_revision"); ok || d.HasChange("create_revision") {
		t, err := expandSystemGlobalCreateRevisionSga(d, v, "create_revision")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-revision"] = t
		}
	}

	if v, ok := d.GetOk("daylightsavetime"); ok || d.HasChange("daylightsavetime") {
		t, err := expandSystemGlobalDaylightsavetimeSga(d, v, "daylightsavetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["daylightsavetime"] = t
		}
	}

	if v, ok := d.GetOk("detect_unregistered_log_device"); ok || d.HasChange("detect_unregistered_log_device") {
		t, err := expandSystemGlobalDetectUnregisteredLogDeviceSga(d, v, "detect_unregistered_log_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-unregistered-log-device"] = t
		}
	}

	if v, ok := d.GetOk("device_view_mode"); ok || d.HasChange("device_view_mode") {
		t, err := expandSystemGlobalDeviceViewModeSga(d, v, "device_view_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-view-mode"] = t
		}
	}

	if v, ok := d.GetOk("dh_params"); ok || d.HasChange("dh_params") {
		t, err := expandSystemGlobalDhParamsSga(d, v, "dh_params")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-params"] = t
		}
	}

	if v, ok := d.GetOk("disable_module"); ok || d.HasChange("disable_module") {
		t, err := expandSystemGlobalDisableModuleSga(d, v, "disable_module")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disable-module"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok || d.HasChange("enc_algorithm") {
		t, err := expandSystemGlobalEncAlgorithmSga(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("faz_status"); ok || d.HasChange("faz_status") {
		t, err := expandSystemGlobalFazStatusSga(d, v, "faz_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-status"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_ca_cert"); ok || d.HasChange("fgfm_ca_cert") {
		t, err := expandSystemGlobalFgfmCaCertSga(d, v, "fgfm_ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_local_cert"); ok || d.HasChange("fgfm_local_cert") {
		t, err := expandSystemGlobalFgfmLocalCertSga(d, v, "fgfm_local_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-local-cert"] = t
		}
	}

	if v, ok := d.GetOk("fgfm_ssl_protocol"); ok || d.HasChange("fgfm_ssl_protocol") {
		t, err := expandSystemGlobalFgfmSslProtocolSga(d, v, "fgfm_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("gui_polling_interval"); ok || d.HasChange("gui_polling_interval") {
		t, err := expandSystemGlobalGuiPollingIntervalSga(d, v, "gui_polling_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-polling-interval"] = t
		}
	}

	if v, ok := d.GetOk("ha_member_auto_grouping"); ok || d.HasChange("ha_member_auto_grouping") {
		t, err := expandSystemGlobalHaMemberAutoGroupingSga(d, v, "ha_member_auto_grouping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-member-auto-grouping"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandSystemGlobalHostnameSga(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("import_ignore_addr_cmt"); ok || d.HasChange("import_ignore_addr_cmt") {
		t, err := expandSystemGlobalImportIgnoreAddrCmtSga(d, v, "import_ignore_addr_cmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["import-ignore-addr-cmt"] = t
		}
	}

	if v, ok := d.GetOk("language"); ok || d.HasChange("language") {
		t, err := expandSystemGlobalLanguageSga(d, v, "language")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["language"] = t
		}
	}

	if v, ok := d.GetOk("latitude"); ok || d.HasChange("latitude") {
		t, err := expandSystemGlobalLatitudeSga(d, v, "latitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latitude"] = t
		}
	}

	if v, ok := d.GetOk("ldap_cache_timeout"); ok || d.HasChange("ldap_cache_timeout") {
		t, err := expandSystemGlobalLdapCacheTimeoutSga(d, v, "ldap_cache_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-cache-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ldapconntimeout"); ok || d.HasChange("ldapconntimeout") {
		t, err := expandSystemGlobalLdapconntimeoutSga(d, v, "ldapconntimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldapconntimeout"] = t
		}
	}

	if v, ok := d.GetOk("lock_preempt"); ok || d.HasChange("lock_preempt") {
		t, err := expandSystemGlobalLockPreemptSga(d, v, "lock_preempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lock-preempt"] = t
		}
	}

	if v, ok := d.GetOk("log_checksum"); ok || d.HasChange("log_checksum") {
		t, err := expandSystemGlobalLogChecksumSga(d, v, "log_checksum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-checksum"] = t
		}
	}

	if v, ok := d.GetOk("log_forward_cache_size"); ok || d.HasChange("log_forward_cache_size") {
		t, err := expandSystemGlobalLogForwardCacheSizeSga(d, v, "log_forward_cache_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-forward-cache-size"] = t
		}
	}

	if v, ok := d.GetOk("longitude"); ok || d.HasChange("longitude") {
		t, err := expandSystemGlobalLongitudeSga(d, v, "longitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["longitude"] = t
		}
	}

	if v, ok := d.GetOk("max_log_forward"); ok || d.HasChange("max_log_forward") {
		t, err := expandSystemGlobalMaxLogForwardSga(d, v, "max_log_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-forward"] = t
		}
	}

	if v, ok := d.GetOk("max_running_reports"); ok || d.HasChange("max_running_reports") {
		t, err := expandSystemGlobalMaxRunningReportsSga(d, v, "max_running_reports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-running-reports"] = t
		}
	}

	if v, ok := d.GetOk("mc_policy_disabled_adoms"); ok || d.HasChange("mc_policy_disabled_adoms") {
		t, err := expandSystemGlobalMcPolicyDisabledAdomsSga(d, v, "mc_policy_disabled_adoms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mc-policy-disabled-adoms"] = t
		}
	}

	if v, ok := d.GetOk("multiple_steps_upgrade_in_autolink"); ok || d.HasChange("multiple_steps_upgrade_in_autolink") {
		t, err := expandSystemGlobalMultipleStepsUpgradeInAutolinkSga(d, v, "multiple_steps_upgrade_in_autolink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-steps-upgrade-in-autolink"] = t
		}
	}

	if v, ok := d.GetOk("normalized_intf_zone_only"); ok || d.HasChange("normalized_intf_zone_only") {
		t, err := expandSystemGlobalNormalizedIntfZoneOnlySga(d, v, "normalized_intf_zone_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["normalized-intf-zone-only"] = t
		}
	}

	if v, ok := d.GetOk("object_revision_db_max"); ok || d.HasChange("object_revision_db_max") {
		t, err := expandSystemGlobalObjectRevisionDbMaxSga(d, v, "object_revision_db_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-revision-db-max"] = t
		}
	}

	if v, ok := d.GetOk("object_revision_mandatory_note"); ok || d.HasChange("object_revision_mandatory_note") {
		t, err := expandSystemGlobalObjectRevisionMandatoryNoteSga(d, v, "object_revision_mandatory_note")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-revision-mandatory-note"] = t
		}
	}

	if v, ok := d.GetOk("object_revision_object_max"); ok || d.HasChange("object_revision_object_max") {
		t, err := expandSystemGlobalObjectRevisionObjectMaxSga(d, v, "object_revision_object_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-revision-object-max"] = t
		}
	}

	if v, ok := d.GetOk("object_revision_status"); ok || d.HasChange("object_revision_status") {
		t, err := expandSystemGlobalObjectRevisionStatusSga(d, v, "object_revision_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-revision-status"] = t
		}
	}

	if v, ok := d.GetOk("oftp_ssl_protocol"); ok || d.HasChange("oftp_ssl_protocol") {
		t, err := expandSystemGlobalOftpSslProtocolSga(d, v, "oftp_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oftp-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("partial_install"); ok || d.HasChange("partial_install") {
		t, err := expandSystemGlobalPartialInstallSga(d, v, "partial_install")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["partial-install"] = t
		}
	}

	if v, ok := d.GetOk("partial_install_force"); ok || d.HasChange("partial_install_force") {
		t, err := expandSystemGlobalPartialInstallForceSga(d, v, "partial_install_force")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["partial-install-force"] = t
		}
	}

	if v, ok := d.GetOk("partial_install_rev"); ok || d.HasChange("partial_install_rev") {
		t, err := expandSystemGlobalPartialInstallRevSga(d, v, "partial_install_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["partial-install-rev"] = t
		}
	}

	if v, ok := d.GetOk("per_policy_lock"); ok || d.HasChange("per_policy_lock") {
		t, err := expandSystemGlobalPerPolicyLockSga(d, v, "per_policy_lock")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-policy-lock"] = t
		}
	}

	if v, ok := d.GetOk("perform_improve_by_ha"); ok || d.HasChange("perform_improve_by_ha") {
		t, err := expandSystemGlobalPerformImproveByHaSga(d, v, "perform_improve_by_ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["perform-improve-by-ha"] = t
		}
	}

	if v, ok := d.GetOk("policy_object_icon"); ok || d.HasChange("policy_object_icon") {
		t, err := expandSystemGlobalPolicyObjectIconSga(d, v, "policy_object_icon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-object-icon"] = t
		}
	}

	if v, ok := d.GetOk("policy_object_in_dual_pane"); ok || d.HasChange("policy_object_in_dual_pane") {
		t, err := expandSystemGlobalPolicyObjectInDualPaneSga(d, v, "policy_object_in_dual_pane")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-object-in-dual-pane"] = t
		}
	}

	if v, ok := d.GetOk("pre_login_banner"); ok || d.HasChange("pre_login_banner") {
		t, err := expandSystemGlobalPreLoginBannerSga(d, v, "pre_login_banner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pre-login-banner"] = t
		}
	}

	if v, ok := d.GetOk("pre_login_banner_message"); ok || d.HasChange("pre_login_banner_message") {
		t, err := expandSystemGlobalPreLoginBannerMessageSga(d, v, "pre_login_banner_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pre-login-banner-message"] = t
		}
	}

	if v, ok := d.GetOk("private_data_encryption"); ok || d.HasChange("private_data_encryption") {
		t, err := expandSystemGlobalPrivateDataEncryptionSga(d, v, "private_data_encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-data-encryption"] = t
		}
	}

	if v, ok := d.GetOk("remoteauthtimeout"); ok || d.HasChange("remoteauthtimeout") {
		t, err := expandSystemGlobalRemoteauthtimeoutSga(d, v, "remoteauthtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remoteauthtimeout"] = t
		}
	}

	if v, ok := d.GetOk("search_all_adoms"); ok || d.HasChange("search_all_adoms") {
		t, err := expandSystemGlobalSearchAllAdomsSga(d, v, "search_all_adoms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-all-adoms"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandSystemGlobalSslCipherSuitesSga(d, v, "ssl_cipher_suites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_low_encryption"); ok || d.HasChange("ssl_low_encryption") {
		t, err := expandSystemGlobalSslLowEncryptionSga(d, v, "ssl_low_encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-low-encryption"] = t
		}
	}

	if v, ok := d.GetOk("ssl_protocol"); ok || d.HasChange("ssl_protocol") {
		t, err := expandSystemGlobalSslProtocolSga(d, v, "ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("ssl_static_key_ciphers"); ok || d.HasChange("ssl_static_key_ciphers") {
		t, err := expandSystemGlobalSslStaticKeyCiphersSga(d, v, "ssl_static_key_ciphers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-static-key-ciphers"] = t
		}
	}

	if v, ok := d.GetOk("table_entry_blink"); ok || d.HasChange("table_entry_blink") {
		t, err := expandSystemGlobalTableEntryBlinkSga(d, v, "table_entry_blink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-entry-blink"] = t
		}
	}

	if v, ok := d.GetOk("task_list_size"); ok || d.HasChange("task_list_size") {
		t, err := expandSystemGlobalTaskListSizeSga(d, v, "task_list_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["task-list-size"] = t
		}
	}

	if v, ok := d.GetOk("tftp"); ok || d.HasChange("tftp") {
		t, err := expandSystemGlobalTftpSga(d, v, "tftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok || d.HasChange("timezone") {
		t, err := expandSystemGlobalTimezoneSga(d, v, "timezone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_mtu"); ok || d.HasChange("tunnel_mtu") {
		t, err := expandSystemGlobalTunnelMtuSga(d, v, "tunnel_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-mtu"] = t
		}
	}

	if v, ok := d.GetOk("usg"); ok || d.HasChange("usg") {
		t, err := expandSystemGlobalUsgSga(d, v, "usg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usg"] = t
		}
	}

	if v, ok := d.GetOk("vdom_mirror"); ok || d.HasChange("vdom_mirror") {
		t, err := expandSystemGlobalVdomMirrorSga(d, v, "vdom_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-mirror"] = t
		}
	}

	if v, ok := d.GetOk("webservice_proto"); ok || d.HasChange("webservice_proto") {
		t, err := expandSystemGlobalWebserviceProtoSga(d, v, "webservice_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webservice-proto"] = t
		}
	}

	if v, ok := d.GetOk("workflow_max_sessions"); ok || d.HasChange("workflow_max_sessions") {
		t, err := expandSystemGlobalWorkflowMaxSessionsSga(d, v, "workflow_max_sessions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workflow-max-sessions"] = t
		}
	}

	if v, ok := d.GetOk("workspace_mode"); ok || d.HasChange("workspace_mode") {
		t, err := expandSystemGlobalWorkspaceModeSga(d, v, "workspace_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workspace-mode"] = t
		}
	}

	return &obj, nil
}
