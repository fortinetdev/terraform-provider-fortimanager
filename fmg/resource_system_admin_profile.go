// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Admin profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAdminProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminProfileCreate,
		Read:   resourceSystemAdminProfileRead,
		Update: resourceSystemAdminProfileUpdate,
		Delete: resourceSystemAdminProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"adom_lock": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adom_policy_packages": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adom_switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_to_install": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assignment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"change_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_retrieve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_revert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"consistency_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datamask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datamask_custom_fields": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"field_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"field_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"field_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"datamask_custom_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datamask_fields": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"datamask_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"datamask_unmasked_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"deploy_management": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_ap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_forticlient": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_fortiextender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_fortiswitch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_manager": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_op": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_policy_package_lock": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_revision_deletion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_wan_link_load_balance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_management": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extension_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_viewer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgd_center_advanced": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgd_center_fmw_mgmt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgd_center_licensing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgd_center": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_policy_packages": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"import_policy_packages": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intf_mapping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_baseline_cfg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_baseline_ovrd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_objects": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_viewer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_objects": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profileid": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"read_passwd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realtime_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"report_viewer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rpc_permit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"run_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"script_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"set_install_targets": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"super_user_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_setting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"term_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"triage_events": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_incidents": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn_manager": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter": &schema.Schema{
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

func resourceSystemAdminProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAdminProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminProfile resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAdminProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemAdminProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "profileid"))

	return resourceSystemAdminProfileRead(d, m)
}

func resourceSystemAdminProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAdminProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAdminProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "profileid"))

	return resourceSystemAdminProfileRead(d, m)
}

func resourceSystemAdminProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAdminProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAdminProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminProfile resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminProfileAdomLock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileAdomPolicyPackages(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileAdomSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileAllowToInstall(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileAppFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileAssignment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileChangePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileConfigRetrieve(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileConfigRevert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileConsistencyCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDatamask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDatamaskCustomFields(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_category"
		if _, ok := i["field-category"]; ok {
			v := flattenSystemAdminProfileDatamaskCustomFieldsFieldCategory(i["field-category"], d, pre_append)
			tmp["field_category"] = fortiAPISubPartPatch(v, "SystemAdminProfile-DatamaskCustomFields-FieldCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_name"
		if _, ok := i["field-name"]; ok {
			v := flattenSystemAdminProfileDatamaskCustomFieldsFieldName(i["field-name"], d, pre_append)
			tmp["field_name"] = fortiAPISubPartPatch(v, "SystemAdminProfile-DatamaskCustomFields-FieldName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_status"
		if _, ok := i["field-status"]; ok {
			v := flattenSystemAdminProfileDatamaskCustomFieldsFieldStatus(i["field-status"], d, pre_append)
			tmp["field_status"] = fortiAPISubPartPatch(v, "SystemAdminProfile-DatamaskCustomFields-FieldStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_type"
		if _, ok := i["field-type"]; ok {
			v := flattenSystemAdminProfileDatamaskCustomFieldsFieldType(i["field-type"], d, pre_append)
			tmp["field_type"] = fortiAPISubPartPatch(v, "SystemAdminProfile-DatamaskCustomFields-FieldType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAdminProfileDatamaskCustomFieldsFieldCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminProfileDatamaskCustomFieldsFieldName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDatamaskCustomFieldsFieldStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDatamaskCustomFieldsFieldType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDatamaskCustomPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDatamaskFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminProfileDatamaskKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminProfileDatamaskUnmaskedTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeployManagement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeviceAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeviceConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeviceForticlient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeviceFortiextender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeviceFortiswitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeviceManager(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeviceOp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDevicePolicyPackageLock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeviceProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeviceRevisionDeletion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDeviceWanLinkLoadBalance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileEventManagement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileExtensionAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileFabricViewer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileFgdCenterAdvanced(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileFgdCenterFmwMgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileFgdCenterLicensing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileFgdCenter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileGlobalPolicyPackages(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileImportPolicyPackages(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIntfMapping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpsBaselineCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpsBaselineOvrd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpsFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpsObjects(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpv6Trusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpv6Trusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpv6Trusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpv6Trusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpv6Trusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpv6Trusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpv6Trusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpv6Trusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpv6Trusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileIpv6Trusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileLogViewer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfilePolicyObjects(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileProfileid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileReadPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileRealtimeMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileReportViewer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileRpcPermit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileRunReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileScriptAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileSetInstallTargets(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileSuperUserProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileSystemSetting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileTermAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileTriageEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileTrusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenSystemAdminProfileTrusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenSystemAdminProfileTrusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenSystemAdminProfileTrusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenSystemAdminProfileTrusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenSystemAdminProfileTrusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenSystemAdminProfileTrusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenSystemAdminProfileTrusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenSystemAdminProfileTrusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenSystemAdminProfileTrusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenSystemAdminProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileUpdateIncidents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileVpnManager(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileWebFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("adom_lock", flattenSystemAdminProfileAdomLock(o["adom-lock"], d, "adom_lock")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-lock"], "SystemAdminProfile-AdomLock"); ok {
			if err = d.Set("adom_lock", vv); err != nil {
				return fmt.Errorf("Error reading adom_lock: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_lock: %v", err)
		}
	}

	if err = d.Set("adom_policy_packages", flattenSystemAdminProfileAdomPolicyPackages(o["adom-policy-packages"], d, "adom_policy_packages")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-policy-packages"], "SystemAdminProfile-AdomPolicyPackages"); ok {
			if err = d.Set("adom_policy_packages", vv); err != nil {
				return fmt.Errorf("Error reading adom_policy_packages: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_policy_packages: %v", err)
		}
	}

	if err = d.Set("adom_switch", flattenSystemAdminProfileAdomSwitch(o["adom-switch"], d, "adom_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-switch"], "SystemAdminProfile-AdomSwitch"); ok {
			if err = d.Set("adom_switch", vv); err != nil {
				return fmt.Errorf("Error reading adom_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_switch: %v", err)
		}
	}

	if err = d.Set("allow_to_install", flattenSystemAdminProfileAllowToInstall(o["allow-to-install"], d, "allow_to_install")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-to-install"], "SystemAdminProfile-AllowToInstall"); ok {
			if err = d.Set("allow_to_install", vv); err != nil {
				return fmt.Errorf("Error reading allow_to_install: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_to_install: %v", err)
		}
	}

	if err = d.Set("app_filter", flattenSystemAdminProfileAppFilter(o["app-filter"], d, "app_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-filter"], "SystemAdminProfile-AppFilter"); ok {
			if err = d.Set("app_filter", vv); err != nil {
				return fmt.Errorf("Error reading app_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_filter: %v", err)
		}
	}

	if err = d.Set("assignment", flattenSystemAdminProfileAssignment(o["assignment"], d, "assignment")); err != nil {
		if vv, ok := fortiAPIPatch(o["assignment"], "SystemAdminProfile-Assignment"); ok {
			if err = d.Set("assignment", vv); err != nil {
				return fmt.Errorf("Error reading assignment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assignment: %v", err)
		}
	}

	if err = d.Set("change_password", flattenSystemAdminProfileChangePassword(o["change-password"], d, "change_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["change-password"], "SystemAdminProfile-ChangePassword"); ok {
			if err = d.Set("change_password", vv); err != nil {
				return fmt.Errorf("Error reading change_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading change_password: %v", err)
		}
	}

	if err = d.Set("config_retrieve", flattenSystemAdminProfileConfigRetrieve(o["config-retrieve"], d, "config_retrieve")); err != nil {
		if vv, ok := fortiAPIPatch(o["config-retrieve"], "SystemAdminProfile-ConfigRetrieve"); ok {
			if err = d.Set("config_retrieve", vv); err != nil {
				return fmt.Errorf("Error reading config_retrieve: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading config_retrieve: %v", err)
		}
	}

	if err = d.Set("config_revert", flattenSystemAdminProfileConfigRevert(o["config-revert"], d, "config_revert")); err != nil {
		if vv, ok := fortiAPIPatch(o["config-revert"], "SystemAdminProfile-ConfigRevert"); ok {
			if err = d.Set("config_revert", vv); err != nil {
				return fmt.Errorf("Error reading config_revert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading config_revert: %v", err)
		}
	}

	if err = d.Set("consistency_check", flattenSystemAdminProfileConsistencyCheck(o["consistency-check"], d, "consistency_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["consistency-check"], "SystemAdminProfile-ConsistencyCheck"); ok {
			if err = d.Set("consistency_check", vv); err != nil {
				return fmt.Errorf("Error reading consistency_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading consistency_check: %v", err)
		}
	}

	if err = d.Set("datamask", flattenSystemAdminProfileDatamask(o["datamask"], d, "datamask")); err != nil {
		if vv, ok := fortiAPIPatch(o["datamask"], "SystemAdminProfile-Datamask"); ok {
			if err = d.Set("datamask", vv); err != nil {
				return fmt.Errorf("Error reading datamask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading datamask: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("datamask_custom_fields", flattenSystemAdminProfileDatamaskCustomFields(o["datamask-custom-fields"], d, "datamask_custom_fields")); err != nil {
			if vv, ok := fortiAPIPatch(o["datamask-custom-fields"], "SystemAdminProfile-DatamaskCustomFields"); ok {
				if err = d.Set("datamask_custom_fields", vv); err != nil {
					return fmt.Errorf("Error reading datamask_custom_fields: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading datamask_custom_fields: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("datamask_custom_fields"); ok {
			if err = d.Set("datamask_custom_fields", flattenSystemAdminProfileDatamaskCustomFields(o["datamask-custom-fields"], d, "datamask_custom_fields")); err != nil {
				if vv, ok := fortiAPIPatch(o["datamask-custom-fields"], "SystemAdminProfile-DatamaskCustomFields"); ok {
					if err = d.Set("datamask_custom_fields", vv); err != nil {
						return fmt.Errorf("Error reading datamask_custom_fields: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading datamask_custom_fields: %v", err)
				}
			}
		}
	}

	if err = d.Set("datamask_custom_priority", flattenSystemAdminProfileDatamaskCustomPriority(o["datamask-custom-priority"], d, "datamask_custom_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["datamask-custom-priority"], "SystemAdminProfile-DatamaskCustomPriority"); ok {
			if err = d.Set("datamask_custom_priority", vv); err != nil {
				return fmt.Errorf("Error reading datamask_custom_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading datamask_custom_priority: %v", err)
		}
	}

	if err = d.Set("datamask_fields", flattenSystemAdminProfileDatamaskFields(o["datamask-fields"], d, "datamask_fields")); err != nil {
		if vv, ok := fortiAPIPatch(o["datamask-fields"], "SystemAdminProfile-DatamaskFields"); ok {
			if err = d.Set("datamask_fields", vv); err != nil {
				return fmt.Errorf("Error reading datamask_fields: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading datamask_fields: %v", err)
		}
	}

	if err = d.Set("datamask_unmasked_time", flattenSystemAdminProfileDatamaskUnmaskedTime(o["datamask-unmasked-time"], d, "datamask_unmasked_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["datamask-unmasked-time"], "SystemAdminProfile-DatamaskUnmaskedTime"); ok {
			if err = d.Set("datamask_unmasked_time", vv); err != nil {
				return fmt.Errorf("Error reading datamask_unmasked_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading datamask_unmasked_time: %v", err)
		}
	}

	if err = d.Set("deploy_management", flattenSystemAdminProfileDeployManagement(o["deploy-management"], d, "deploy_management")); err != nil {
		if vv, ok := fortiAPIPatch(o["deploy-management"], "SystemAdminProfile-DeployManagement"); ok {
			if err = d.Set("deploy_management", vv); err != nil {
				return fmt.Errorf("Error reading deploy_management: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deploy_management: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemAdminProfileDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemAdminProfile-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("device_ap", flattenSystemAdminProfileDeviceAp(o["device-ap"], d, "device_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-ap"], "SystemAdminProfile-DeviceAp"); ok {
			if err = d.Set("device_ap", vv); err != nil {
				return fmt.Errorf("Error reading device_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_ap: %v", err)
		}
	}

	if err = d.Set("device_config", flattenSystemAdminProfileDeviceConfig(o["device-config"], d, "device_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-config"], "SystemAdminProfile-DeviceConfig"); ok {
			if err = d.Set("device_config", vv); err != nil {
				return fmt.Errorf("Error reading device_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_config: %v", err)
		}
	}

	if err = d.Set("device_forticlient", flattenSystemAdminProfileDeviceForticlient(o["device-forticlient"], d, "device_forticlient")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-forticlient"], "SystemAdminProfile-DeviceForticlient"); ok {
			if err = d.Set("device_forticlient", vv); err != nil {
				return fmt.Errorf("Error reading device_forticlient: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_forticlient: %v", err)
		}
	}

	if err = d.Set("device_fortiextender", flattenSystemAdminProfileDeviceFortiextender(o["device-fortiextender"], d, "device_fortiextender")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-fortiextender"], "SystemAdminProfile-DeviceFortiextender"); ok {
			if err = d.Set("device_fortiextender", vv); err != nil {
				return fmt.Errorf("Error reading device_fortiextender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_fortiextender: %v", err)
		}
	}

	if err = d.Set("device_fortiswitch", flattenSystemAdminProfileDeviceFortiswitch(o["device-fortiswitch"], d, "device_fortiswitch")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-fortiswitch"], "SystemAdminProfile-DeviceFortiswitch"); ok {
			if err = d.Set("device_fortiswitch", vv); err != nil {
				return fmt.Errorf("Error reading device_fortiswitch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_fortiswitch: %v", err)
		}
	}

	if err = d.Set("device_manager", flattenSystemAdminProfileDeviceManager(o["device-manager"], d, "device_manager")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-manager"], "SystemAdminProfile-DeviceManager"); ok {
			if err = d.Set("device_manager", vv); err != nil {
				return fmt.Errorf("Error reading device_manager: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_manager: %v", err)
		}
	}

	if err = d.Set("device_op", flattenSystemAdminProfileDeviceOp(o["device-op"], d, "device_op")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-op"], "SystemAdminProfile-DeviceOp"); ok {
			if err = d.Set("device_op", vv); err != nil {
				return fmt.Errorf("Error reading device_op: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_op: %v", err)
		}
	}

	if err = d.Set("device_policy_package_lock", flattenSystemAdminProfileDevicePolicyPackageLock(o["device-policy-package-lock"], d, "device_policy_package_lock")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-policy-package-lock"], "SystemAdminProfile-DevicePolicyPackageLock"); ok {
			if err = d.Set("device_policy_package_lock", vv); err != nil {
				return fmt.Errorf("Error reading device_policy_package_lock: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_policy_package_lock: %v", err)
		}
	}

	if err = d.Set("device_profile", flattenSystemAdminProfileDeviceProfile(o["device-profile"], d, "device_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-profile"], "SystemAdminProfile-DeviceProfile"); ok {
			if err = d.Set("device_profile", vv); err != nil {
				return fmt.Errorf("Error reading device_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_profile: %v", err)
		}
	}

	if err = d.Set("device_revision_deletion", flattenSystemAdminProfileDeviceRevisionDeletion(o["device-revision-deletion"], d, "device_revision_deletion")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-revision-deletion"], "SystemAdminProfile-DeviceRevisionDeletion"); ok {
			if err = d.Set("device_revision_deletion", vv); err != nil {
				return fmt.Errorf("Error reading device_revision_deletion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_revision_deletion: %v", err)
		}
	}

	if err = d.Set("device_wan_link_load_balance", flattenSystemAdminProfileDeviceWanLinkLoadBalance(o["device-wan-link-load-balance"], d, "device_wan_link_load_balance")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-wan-link-load-balance"], "SystemAdminProfile-DeviceWanLinkLoadBalance"); ok {
			if err = d.Set("device_wan_link_load_balance", vv); err != nil {
				return fmt.Errorf("Error reading device_wan_link_load_balance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_wan_link_load_balance: %v", err)
		}
	}

	if err = d.Set("event_management", flattenSystemAdminProfileEventManagement(o["event-management"], d, "event_management")); err != nil {
		if vv, ok := fortiAPIPatch(o["event-management"], "SystemAdminProfile-EventManagement"); ok {
			if err = d.Set("event_management", vv); err != nil {
				return fmt.Errorf("Error reading event_management: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event_management: %v", err)
		}
	}

	if err = d.Set("extension_access", flattenSystemAdminProfileExtensionAccess(o["extension-access"], d, "extension_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension-access"], "SystemAdminProfile-ExtensionAccess"); ok {
			if err = d.Set("extension_access", vv); err != nil {
				return fmt.Errorf("Error reading extension_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension_access: %v", err)
		}
	}

	if err = d.Set("fabric_viewer", flattenSystemAdminProfileFabricViewer(o["fabric-viewer"], d, "fabric_viewer")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-viewer"], "SystemAdminProfile-FabricViewer"); ok {
			if err = d.Set("fabric_viewer", vv); err != nil {
				return fmt.Errorf("Error reading fabric_viewer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_viewer: %v", err)
		}
	}

	if err = d.Set("fgd_center_advanced", flattenSystemAdminProfileFgdCenterAdvanced(o["fgd-center-advanced"], d, "fgd_center_advanced")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd-center-advanced"], "SystemAdminProfile-FgdCenterAdvanced"); ok {
			if err = d.Set("fgd_center_advanced", vv); err != nil {
				return fmt.Errorf("Error reading fgd_center_advanced: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_center_advanced: %v", err)
		}
	}

	if err = d.Set("fgd_center_fmw_mgmt", flattenSystemAdminProfileFgdCenterFmwMgmt(o["fgd-center-fmw-mgmt"], d, "fgd_center_fmw_mgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd-center-fmw-mgmt"], "SystemAdminProfile-FgdCenterFmwMgmt"); ok {
			if err = d.Set("fgd_center_fmw_mgmt", vv); err != nil {
				return fmt.Errorf("Error reading fgd_center_fmw_mgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_center_fmw_mgmt: %v", err)
		}
	}

	if err = d.Set("fgd_center_licensing", flattenSystemAdminProfileFgdCenterLicensing(o["fgd-center-licensing"], d, "fgd_center_licensing")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd-center-licensing"], "SystemAdminProfile-FgdCenterLicensing"); ok {
			if err = d.Set("fgd_center_licensing", vv); err != nil {
				return fmt.Errorf("Error reading fgd_center_licensing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_center_licensing: %v", err)
		}
	}

	if err = d.Set("fgd_center", flattenSystemAdminProfileFgdCenter(o["fgd_center"], d, "fgd_center")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd_center"], "SystemAdminProfile-FgdCenter"); ok {
			if err = d.Set("fgd_center", vv); err != nil {
				return fmt.Errorf("Error reading fgd_center: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_center: %v", err)
		}
	}

	if err = d.Set("global_policy_packages", flattenSystemAdminProfileGlobalPolicyPackages(o["global-policy-packages"], d, "global_policy_packages")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-policy-packages"], "SystemAdminProfile-GlobalPolicyPackages"); ok {
			if err = d.Set("global_policy_packages", vv); err != nil {
				return fmt.Errorf("Error reading global_policy_packages: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_policy_packages: %v", err)
		}
	}

	if err = d.Set("import_policy_packages", flattenSystemAdminProfileImportPolicyPackages(o["import-policy-packages"], d, "import_policy_packages")); err != nil {
		if vv, ok := fortiAPIPatch(o["import-policy-packages"], "SystemAdminProfile-ImportPolicyPackages"); ok {
			if err = d.Set("import_policy_packages", vv); err != nil {
				return fmt.Errorf("Error reading import_policy_packages: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading import_policy_packages: %v", err)
		}
	}

	if err = d.Set("intf_mapping", flattenSystemAdminProfileIntfMapping(o["intf-mapping"], d, "intf_mapping")); err != nil {
		if vv, ok := fortiAPIPatch(o["intf-mapping"], "SystemAdminProfile-IntfMapping"); ok {
			if err = d.Set("intf_mapping", vv); err != nil {
				return fmt.Errorf("Error reading intf_mapping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intf_mapping: %v", err)
		}
	}

	if err = d.Set("ips_baseline_cfg", flattenSystemAdminProfileIpsBaselineCfg(o["ips-baseline-cfg"], d, "ips_baseline_cfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-baseline-cfg"], "SystemAdminProfile-IpsBaselineCfg"); ok {
			if err = d.Set("ips_baseline_cfg", vv); err != nil {
				return fmt.Errorf("Error reading ips_baseline_cfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_baseline_cfg: %v", err)
		}
	}

	if err = d.Set("ips_baseline_ovrd", flattenSystemAdminProfileIpsBaselineOvrd(o["ips-baseline-ovrd"], d, "ips_baseline_ovrd")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-baseline-ovrd"], "SystemAdminProfile-IpsBaselineOvrd"); ok {
			if err = d.Set("ips_baseline_ovrd", vv); err != nil {
				return fmt.Errorf("Error reading ips_baseline_ovrd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_baseline_ovrd: %v", err)
		}
	}

	if err = d.Set("ips_filter", flattenSystemAdminProfileIpsFilter(o["ips-filter"], d, "ips_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-filter"], "SystemAdminProfile-IpsFilter"); ok {
			if err = d.Set("ips_filter", vv); err != nil {
				return fmt.Errorf("Error reading ips_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_filter: %v", err)
		}
	}

	if err = d.Set("ips_objects", flattenSystemAdminProfileIpsObjects(o["ips-objects"], d, "ips_objects")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-objects"], "SystemAdminProfile-IpsObjects"); ok {
			if err = d.Set("ips_objects", vv); err != nil {
				return fmt.Errorf("Error reading ips_objects: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_objects: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost1", flattenSystemAdminProfileIpv6Trusthost1(o["ipv6_trusthost1"], d, "ipv6_trusthost1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost1"], "SystemAdminProfile-Ipv6Trusthost1"); ok {
			if err = d.Set("ipv6_trusthost1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost1: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost10", flattenSystemAdminProfileIpv6Trusthost10(o["ipv6_trusthost10"], d, "ipv6_trusthost10")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost10"], "SystemAdminProfile-Ipv6Trusthost10"); ok {
			if err = d.Set("ipv6_trusthost10", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost10: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost2", flattenSystemAdminProfileIpv6Trusthost2(o["ipv6_trusthost2"], d, "ipv6_trusthost2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost2"], "SystemAdminProfile-Ipv6Trusthost2"); ok {
			if err = d.Set("ipv6_trusthost2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost2: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost3", flattenSystemAdminProfileIpv6Trusthost3(o["ipv6_trusthost3"], d, "ipv6_trusthost3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost3"], "SystemAdminProfile-Ipv6Trusthost3"); ok {
			if err = d.Set("ipv6_trusthost3", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost3: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost4", flattenSystemAdminProfileIpv6Trusthost4(o["ipv6_trusthost4"], d, "ipv6_trusthost4")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost4"], "SystemAdminProfile-Ipv6Trusthost4"); ok {
			if err = d.Set("ipv6_trusthost4", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost4: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost5", flattenSystemAdminProfileIpv6Trusthost5(o["ipv6_trusthost5"], d, "ipv6_trusthost5")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost5"], "SystemAdminProfile-Ipv6Trusthost5"); ok {
			if err = d.Set("ipv6_trusthost5", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost5: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost6", flattenSystemAdminProfileIpv6Trusthost6(o["ipv6_trusthost6"], d, "ipv6_trusthost6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost6"], "SystemAdminProfile-Ipv6Trusthost6"); ok {
			if err = d.Set("ipv6_trusthost6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost6: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost7", flattenSystemAdminProfileIpv6Trusthost7(o["ipv6_trusthost7"], d, "ipv6_trusthost7")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost7"], "SystemAdminProfile-Ipv6Trusthost7"); ok {
			if err = d.Set("ipv6_trusthost7", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost7: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost8", flattenSystemAdminProfileIpv6Trusthost8(o["ipv6_trusthost8"], d, "ipv6_trusthost8")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost8"], "SystemAdminProfile-Ipv6Trusthost8"); ok {
			if err = d.Set("ipv6_trusthost8", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost8: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost9", flattenSystemAdminProfileIpv6Trusthost9(o["ipv6_trusthost9"], d, "ipv6_trusthost9")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost9"], "SystemAdminProfile-Ipv6Trusthost9"); ok {
			if err = d.Set("ipv6_trusthost9", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost9: %v", err)
		}
	}

	if err = d.Set("log_viewer", flattenSystemAdminProfileLogViewer(o["log-viewer"], d, "log_viewer")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-viewer"], "SystemAdminProfile-LogViewer"); ok {
			if err = d.Set("log_viewer", vv); err != nil {
				return fmt.Errorf("Error reading log_viewer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_viewer: %v", err)
		}
	}

	if err = d.Set("policy_objects", flattenSystemAdminProfilePolicyObjects(o["policy-objects"], d, "policy_objects")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-objects"], "SystemAdminProfile-PolicyObjects"); ok {
			if err = d.Set("policy_objects", vv); err != nil {
				return fmt.Errorf("Error reading policy_objects: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_objects: %v", err)
		}
	}

	if err = d.Set("profileid", flattenSystemAdminProfileProfileid(o["profileid"], d, "profileid")); err != nil {
		if vv, ok := fortiAPIPatch(o["profileid"], "SystemAdminProfile-Profileid"); ok {
			if err = d.Set("profileid", vv); err != nil {
				return fmt.Errorf("Error reading profileid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profileid: %v", err)
		}
	}

	if err = d.Set("read_passwd", flattenSystemAdminProfileReadPasswd(o["read-passwd"], d, "read_passwd")); err != nil {
		if vv, ok := fortiAPIPatch(o["read-passwd"], "SystemAdminProfile-ReadPasswd"); ok {
			if err = d.Set("read_passwd", vv); err != nil {
				return fmt.Errorf("Error reading read_passwd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading read_passwd: %v", err)
		}
	}

	if err = d.Set("realtime_monitor", flattenSystemAdminProfileRealtimeMonitor(o["realtime-monitor"], d, "realtime_monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["realtime-monitor"], "SystemAdminProfile-RealtimeMonitor"); ok {
			if err = d.Set("realtime_monitor", vv); err != nil {
				return fmt.Errorf("Error reading realtime_monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading realtime_monitor: %v", err)
		}
	}

	if err = d.Set("report_viewer", flattenSystemAdminProfileReportViewer(o["report-viewer"], d, "report_viewer")); err != nil {
		if vv, ok := fortiAPIPatch(o["report-viewer"], "SystemAdminProfile-ReportViewer"); ok {
			if err = d.Set("report_viewer", vv); err != nil {
				return fmt.Errorf("Error reading report_viewer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report_viewer: %v", err)
		}
	}

	if err = d.Set("rpc_permit", flattenSystemAdminProfileRpcPermit(o["rpc-permit"], d, "rpc_permit")); err != nil {
		if vv, ok := fortiAPIPatch(o["rpc-permit"], "SystemAdminProfile-RpcPermit"); ok {
			if err = d.Set("rpc_permit", vv); err != nil {
				return fmt.Errorf("Error reading rpc_permit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rpc_permit: %v", err)
		}
	}

	if err = d.Set("run_report", flattenSystemAdminProfileRunReport(o["run-report"], d, "run_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["run-report"], "SystemAdminProfile-RunReport"); ok {
			if err = d.Set("run_report", vv); err != nil {
				return fmt.Errorf("Error reading run_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading run_report: %v", err)
		}
	}

	if err = d.Set("scope", flattenSystemAdminProfileScope(o["scope"], d, "scope")); err != nil {
		if vv, ok := fortiAPIPatch(o["scope"], "SystemAdminProfile-Scope"); ok {
			if err = d.Set("scope", vv); err != nil {
				return fmt.Errorf("Error reading scope: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scope: %v", err)
		}
	}

	if err = d.Set("script_access", flattenSystemAdminProfileScriptAccess(o["script-access"], d, "script_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["script-access"], "SystemAdminProfile-ScriptAccess"); ok {
			if err = d.Set("script_access", vv); err != nil {
				return fmt.Errorf("Error reading script_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading script_access: %v", err)
		}
	}

	if err = d.Set("set_install_targets", flattenSystemAdminProfileSetInstallTargets(o["set-install-targets"], d, "set_install_targets")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-install-targets"], "SystemAdminProfile-SetInstallTargets"); ok {
			if err = d.Set("set_install_targets", vv); err != nil {
				return fmt.Errorf("Error reading set_install_targets: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_install_targets: %v", err)
		}
	}

	if err = d.Set("super_user_profile", flattenSystemAdminProfileSuperUserProfile(o["super-user-profile"], d, "super_user_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["super-user-profile"], "SystemAdminProfile-SuperUserProfile"); ok {
			if err = d.Set("super_user_profile", vv); err != nil {
				return fmt.Errorf("Error reading super_user_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading super_user_profile: %v", err)
		}
	}

	if err = d.Set("system_setting", flattenSystemAdminProfileSystemSetting(o["system-setting"], d, "system_setting")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-setting"], "SystemAdminProfile-SystemSetting"); ok {
			if err = d.Set("system_setting", vv); err != nil {
				return fmt.Errorf("Error reading system_setting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_setting: %v", err)
		}
	}

	if err = d.Set("term_access", flattenSystemAdminProfileTermAccess(o["term-access"], d, "term_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["term-access"], "SystemAdminProfile-TermAccess"); ok {
			if err = d.Set("term_access", vv); err != nil {
				return fmt.Errorf("Error reading term_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading term_access: %v", err)
		}
	}

	if err = d.Set("triage_events", flattenSystemAdminProfileTriageEvents(o["triage-events"], d, "triage_events")); err != nil {
		if vv, ok := fortiAPIPatch(o["triage-events"], "SystemAdminProfile-TriageEvents"); ok {
			if err = d.Set("triage_events", vv); err != nil {
				return fmt.Errorf("Error reading triage_events: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading triage_events: %v", err)
		}
	}

	if err = d.Set("trusthost1", flattenSystemAdminProfileTrusthost1(o["trusthost1"], d, "trusthost1")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost1"], "SystemAdminProfile-Trusthost1"); ok {
			if err = d.Set("trusthost1", vv); err != nil {
				return fmt.Errorf("Error reading trusthost1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost1: %v", err)
		}
	}

	if err = d.Set("trusthost10", flattenSystemAdminProfileTrusthost10(o["trusthost10"], d, "trusthost10")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost10"], "SystemAdminProfile-Trusthost10"); ok {
			if err = d.Set("trusthost10", vv); err != nil {
				return fmt.Errorf("Error reading trusthost10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost10: %v", err)
		}
	}

	if err = d.Set("trusthost2", flattenSystemAdminProfileTrusthost2(o["trusthost2"], d, "trusthost2")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost2"], "SystemAdminProfile-Trusthost2"); ok {
			if err = d.Set("trusthost2", vv); err != nil {
				return fmt.Errorf("Error reading trusthost2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost2: %v", err)
		}
	}

	if err = d.Set("trusthost3", flattenSystemAdminProfileTrusthost3(o["trusthost3"], d, "trusthost3")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost3"], "SystemAdminProfile-Trusthost3"); ok {
			if err = d.Set("trusthost3", vv); err != nil {
				return fmt.Errorf("Error reading trusthost3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost3: %v", err)
		}
	}

	if err = d.Set("trusthost4", flattenSystemAdminProfileTrusthost4(o["trusthost4"], d, "trusthost4")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost4"], "SystemAdminProfile-Trusthost4"); ok {
			if err = d.Set("trusthost4", vv); err != nil {
				return fmt.Errorf("Error reading trusthost4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost4: %v", err)
		}
	}

	if err = d.Set("trusthost5", flattenSystemAdminProfileTrusthost5(o["trusthost5"], d, "trusthost5")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost5"], "SystemAdminProfile-Trusthost5"); ok {
			if err = d.Set("trusthost5", vv); err != nil {
				return fmt.Errorf("Error reading trusthost5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost5: %v", err)
		}
	}

	if err = d.Set("trusthost6", flattenSystemAdminProfileTrusthost6(o["trusthost6"], d, "trusthost6")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost6"], "SystemAdminProfile-Trusthost6"); ok {
			if err = d.Set("trusthost6", vv); err != nil {
				return fmt.Errorf("Error reading trusthost6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost6: %v", err)
		}
	}

	if err = d.Set("trusthost7", flattenSystemAdminProfileTrusthost7(o["trusthost7"], d, "trusthost7")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost7"], "SystemAdminProfile-Trusthost7"); ok {
			if err = d.Set("trusthost7", vv); err != nil {
				return fmt.Errorf("Error reading trusthost7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost7: %v", err)
		}
	}

	if err = d.Set("trusthost8", flattenSystemAdminProfileTrusthost8(o["trusthost8"], d, "trusthost8")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost8"], "SystemAdminProfile-Trusthost8"); ok {
			if err = d.Set("trusthost8", vv); err != nil {
				return fmt.Errorf("Error reading trusthost8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost8: %v", err)
		}
	}

	if err = d.Set("trusthost9", flattenSystemAdminProfileTrusthost9(o["trusthost9"], d, "trusthost9")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost9"], "SystemAdminProfile-Trusthost9"); ok {
			if err = d.Set("trusthost9", vv); err != nil {
				return fmt.Errorf("Error reading trusthost9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost9: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemAdminProfileType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemAdminProfile-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("update_incidents", flattenSystemAdminProfileUpdateIncidents(o["update-incidents"], d, "update_incidents")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-incidents"], "SystemAdminProfile-UpdateIncidents"); ok {
			if err = d.Set("update_incidents", vv); err != nil {
				return fmt.Errorf("Error reading update_incidents: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_incidents: %v", err)
		}
	}

	if err = d.Set("vpn_manager", flattenSystemAdminProfileVpnManager(o["vpn-manager"], d, "vpn_manager")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn-manager"], "SystemAdminProfile-VpnManager"); ok {
			if err = d.Set("vpn_manager", vv); err != nil {
				return fmt.Errorf("Error reading vpn_manager: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_manager: %v", err)
		}
	}

	if err = d.Set("web_filter", flattenSystemAdminProfileWebFilter(o["web-filter"], d, "web_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter"], "SystemAdminProfile-WebFilter"); ok {
			if err = d.Set("web_filter", vv); err != nil {
				return fmt.Errorf("Error reading web_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminProfileAdomLock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileAdomPolicyPackages(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileAdomSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileAllowToInstall(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileAppFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileAssignment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileChangePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileConfigRetrieve(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileConfigRevert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileConsistencyCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDatamask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDatamaskCustomFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["field-category"], _ = expandSystemAdminProfileDatamaskCustomFieldsFieldCategory(d, i["field_category"], pre_append)
		} else {
			tmp["field-category"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["field-name"], _ = expandSystemAdminProfileDatamaskCustomFieldsFieldName(d, i["field_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["field-status"], _ = expandSystemAdminProfileDatamaskCustomFieldsFieldStatus(d, i["field_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["field-type"], _ = expandSystemAdminProfileDatamaskCustomFieldsFieldType(d, i["field_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAdminProfileDatamaskCustomFieldsFieldCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminProfileDatamaskCustomFieldsFieldName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDatamaskCustomFieldsFieldStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDatamaskCustomFieldsFieldType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDatamaskCustomPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDatamaskFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminProfileDatamaskKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminProfileDatamaskUnmaskedTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeployManagement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeviceAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeviceConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeviceForticlient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeviceFortiextender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeviceFortiswitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeviceManager(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeviceOp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDevicePolicyPackageLock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeviceProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeviceRevisionDeletion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDeviceWanLinkLoadBalance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileEventManagement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileExtensionAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileFabricViewer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileFgdCenterAdvanced(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileFgdCenterFmwMgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileFgdCenterLicensing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileFgdCenter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileGlobalPolicyPackages(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileImportPolicyPackages(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIntfMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpsBaselineCfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpsBaselineOvrd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpsFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpsObjects(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpv6Trusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpv6Trusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpv6Trusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpv6Trusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpv6Trusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpv6Trusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpv6Trusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpv6Trusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpv6Trusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileIpv6Trusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileLogViewer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfilePolicyObjects(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileProfileid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileReadPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileRealtimeMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileReportViewer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileRpcPermit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileRunReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileScriptAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileSetInstallTargets(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileSuperUserProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileSystemSetting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTermAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTriageEvents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTrusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTrusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTrusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTrusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTrusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTrusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTrusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTrusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTrusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileTrusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileUpdateIncidents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileVpnManager(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileWebFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("adom_lock"); ok || d.HasChange("adom_lock") {
		t, err := expandSystemAdminProfileAdomLock(d, v, "adom_lock")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-lock"] = t
		}
	}

	if v, ok := d.GetOk("adom_policy_packages"); ok || d.HasChange("adom_policy_packages") {
		t, err := expandSystemAdminProfileAdomPolicyPackages(d, v, "adom_policy_packages")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-policy-packages"] = t
		}
	}

	if v, ok := d.GetOk("adom_switch"); ok || d.HasChange("adom_switch") {
		t, err := expandSystemAdminProfileAdomSwitch(d, v, "adom_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-switch"] = t
		}
	}

	if v, ok := d.GetOk("allow_to_install"); ok || d.HasChange("allow_to_install") {
		t, err := expandSystemAdminProfileAllowToInstall(d, v, "allow_to_install")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-to-install"] = t
		}
	}

	if v, ok := d.GetOk("app_filter"); ok || d.HasChange("app_filter") {
		t, err := expandSystemAdminProfileAppFilter(d, v, "app_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-filter"] = t
		}
	}

	if v, ok := d.GetOk("assignment"); ok || d.HasChange("assignment") {
		t, err := expandSystemAdminProfileAssignment(d, v, "assignment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assignment"] = t
		}
	}

	if v, ok := d.GetOk("change_password"); ok || d.HasChange("change_password") {
		t, err := expandSystemAdminProfileChangePassword(d, v, "change_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-password"] = t
		}
	}

	if v, ok := d.GetOk("config_retrieve"); ok || d.HasChange("config_retrieve") {
		t, err := expandSystemAdminProfileConfigRetrieve(d, v, "config_retrieve")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["config-retrieve"] = t
		}
	}

	if v, ok := d.GetOk("config_revert"); ok || d.HasChange("config_revert") {
		t, err := expandSystemAdminProfileConfigRevert(d, v, "config_revert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["config-revert"] = t
		}
	}

	if v, ok := d.GetOk("consistency_check"); ok || d.HasChange("consistency_check") {
		t, err := expandSystemAdminProfileConsistencyCheck(d, v, "consistency_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["consistency-check"] = t
		}
	}

	if v, ok := d.GetOk("datamask"); ok || d.HasChange("datamask") {
		t, err := expandSystemAdminProfileDatamask(d, v, "datamask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["datamask"] = t
		}
	}

	if v, ok := d.GetOk("datamask_custom_fields"); ok || d.HasChange("datamask_custom_fields") {
		t, err := expandSystemAdminProfileDatamaskCustomFields(d, v, "datamask_custom_fields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["datamask-custom-fields"] = t
		}
	}

	if v, ok := d.GetOk("datamask_custom_priority"); ok || d.HasChange("datamask_custom_priority") {
		t, err := expandSystemAdminProfileDatamaskCustomPriority(d, v, "datamask_custom_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["datamask-custom-priority"] = t
		}
	}

	if v, ok := d.GetOk("datamask_fields"); ok || d.HasChange("datamask_fields") {
		t, err := expandSystemAdminProfileDatamaskFields(d, v, "datamask_fields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["datamask-fields"] = t
		}
	}

	if v, ok := d.GetOk("datamask_key"); ok || d.HasChange("datamask_key") {
		t, err := expandSystemAdminProfileDatamaskKey(d, v, "datamask_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["datamask-key"] = t
		}
	}

	if v, ok := d.GetOk("datamask_unmasked_time"); ok || d.HasChange("datamask_unmasked_time") {
		t, err := expandSystemAdminProfileDatamaskUnmaskedTime(d, v, "datamask_unmasked_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["datamask-unmasked-time"] = t
		}
	}

	if v, ok := d.GetOk("deploy_management"); ok || d.HasChange("deploy_management") {
		t, err := expandSystemAdminProfileDeployManagement(d, v, "deploy_management")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deploy-management"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemAdminProfileDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("device_ap"); ok || d.HasChange("device_ap") {
		t, err := expandSystemAdminProfileDeviceAp(d, v, "device_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-ap"] = t
		}
	}

	if v, ok := d.GetOk("device_config"); ok || d.HasChange("device_config") {
		t, err := expandSystemAdminProfileDeviceConfig(d, v, "device_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-config"] = t
		}
	}

	if v, ok := d.GetOk("device_forticlient"); ok || d.HasChange("device_forticlient") {
		t, err := expandSystemAdminProfileDeviceForticlient(d, v, "device_forticlient")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-forticlient"] = t
		}
	}

	if v, ok := d.GetOk("device_fortiextender"); ok || d.HasChange("device_fortiextender") {
		t, err := expandSystemAdminProfileDeviceFortiextender(d, v, "device_fortiextender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-fortiextender"] = t
		}
	}

	if v, ok := d.GetOk("device_fortiswitch"); ok || d.HasChange("device_fortiswitch") {
		t, err := expandSystemAdminProfileDeviceFortiswitch(d, v, "device_fortiswitch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-fortiswitch"] = t
		}
	}

	if v, ok := d.GetOk("device_manager"); ok || d.HasChange("device_manager") {
		t, err := expandSystemAdminProfileDeviceManager(d, v, "device_manager")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-manager"] = t
		}
	}

	if v, ok := d.GetOk("device_op"); ok || d.HasChange("device_op") {
		t, err := expandSystemAdminProfileDeviceOp(d, v, "device_op")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-op"] = t
		}
	}

	if v, ok := d.GetOk("device_policy_package_lock"); ok || d.HasChange("device_policy_package_lock") {
		t, err := expandSystemAdminProfileDevicePolicyPackageLock(d, v, "device_policy_package_lock")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-policy-package-lock"] = t
		}
	}

	if v, ok := d.GetOk("device_profile"); ok || d.HasChange("device_profile") {
		t, err := expandSystemAdminProfileDeviceProfile(d, v, "device_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-profile"] = t
		}
	}

	if v, ok := d.GetOk("device_revision_deletion"); ok || d.HasChange("device_revision_deletion") {
		t, err := expandSystemAdminProfileDeviceRevisionDeletion(d, v, "device_revision_deletion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-revision-deletion"] = t
		}
	}

	if v, ok := d.GetOk("device_wan_link_load_balance"); ok || d.HasChange("device_wan_link_load_balance") {
		t, err := expandSystemAdminProfileDeviceWanLinkLoadBalance(d, v, "device_wan_link_load_balance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-wan-link-load-balance"] = t
		}
	}

	if v, ok := d.GetOk("event_management"); ok || d.HasChange("event_management") {
		t, err := expandSystemAdminProfileEventManagement(d, v, "event_management")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event-management"] = t
		}
	}

	if v, ok := d.GetOk("extension_access"); ok || d.HasChange("extension_access") {
		t, err := expandSystemAdminProfileExtensionAccess(d, v, "extension_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-access"] = t
		}
	}

	if v, ok := d.GetOk("fabric_viewer"); ok || d.HasChange("fabric_viewer") {
		t, err := expandSystemAdminProfileFabricViewer(d, v, "fabric_viewer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-viewer"] = t
		}
	}

	if v, ok := d.GetOk("fgd_center_advanced"); ok || d.HasChange("fgd_center_advanced") {
		t, err := expandSystemAdminProfileFgdCenterAdvanced(d, v, "fgd_center_advanced")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-center-advanced"] = t
		}
	}

	if v, ok := d.GetOk("fgd_center_fmw_mgmt"); ok || d.HasChange("fgd_center_fmw_mgmt") {
		t, err := expandSystemAdminProfileFgdCenterFmwMgmt(d, v, "fgd_center_fmw_mgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-center-fmw-mgmt"] = t
		}
	}

	if v, ok := d.GetOk("fgd_center_licensing"); ok || d.HasChange("fgd_center_licensing") {
		t, err := expandSystemAdminProfileFgdCenterLicensing(d, v, "fgd_center_licensing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-center-licensing"] = t
		}
	}

	if v, ok := d.GetOk("fgd_center"); ok || d.HasChange("fgd_center") {
		t, err := expandSystemAdminProfileFgdCenter(d, v, "fgd_center")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd_center"] = t
		}
	}

	if v, ok := d.GetOk("global_policy_packages"); ok || d.HasChange("global_policy_packages") {
		t, err := expandSystemAdminProfileGlobalPolicyPackages(d, v, "global_policy_packages")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-policy-packages"] = t
		}
	}

	if v, ok := d.GetOk("import_policy_packages"); ok || d.HasChange("import_policy_packages") {
		t, err := expandSystemAdminProfileImportPolicyPackages(d, v, "import_policy_packages")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["import-policy-packages"] = t
		}
	}

	if v, ok := d.GetOk("intf_mapping"); ok || d.HasChange("intf_mapping") {
		t, err := expandSystemAdminProfileIntfMapping(d, v, "intf_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf-mapping"] = t
		}
	}

	if v, ok := d.GetOk("ips_baseline_cfg"); ok || d.HasChange("ips_baseline_cfg") {
		t, err := expandSystemAdminProfileIpsBaselineCfg(d, v, "ips_baseline_cfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-baseline-cfg"] = t
		}
	}

	if v, ok := d.GetOk("ips_baseline_ovrd"); ok || d.HasChange("ips_baseline_ovrd") {
		t, err := expandSystemAdminProfileIpsBaselineOvrd(d, v, "ips_baseline_ovrd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-baseline-ovrd"] = t
		}
	}

	if v, ok := d.GetOk("ips_filter"); ok || d.HasChange("ips_filter") {
		t, err := expandSystemAdminProfileIpsFilter(d, v, "ips_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-filter"] = t
		}
	}

	if v, ok := d.GetOk("ips_objects"); ok || d.HasChange("ips_objects") {
		t, err := expandSystemAdminProfileIpsObjects(d, v, "ips_objects")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-objects"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost1"); ok || d.HasChange("ipv6_trusthost1") {
		t, err := expandSystemAdminProfileIpv6Trusthost1(d, v, "ipv6_trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost10"); ok || d.HasChange("ipv6_trusthost10") {
		t, err := expandSystemAdminProfileIpv6Trusthost10(d, v, "ipv6_trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost2"); ok || d.HasChange("ipv6_trusthost2") {
		t, err := expandSystemAdminProfileIpv6Trusthost2(d, v, "ipv6_trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost3"); ok || d.HasChange("ipv6_trusthost3") {
		t, err := expandSystemAdminProfileIpv6Trusthost3(d, v, "ipv6_trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost4"); ok || d.HasChange("ipv6_trusthost4") {
		t, err := expandSystemAdminProfileIpv6Trusthost4(d, v, "ipv6_trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost5"); ok || d.HasChange("ipv6_trusthost5") {
		t, err := expandSystemAdminProfileIpv6Trusthost5(d, v, "ipv6_trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost6"); ok || d.HasChange("ipv6_trusthost6") {
		t, err := expandSystemAdminProfileIpv6Trusthost6(d, v, "ipv6_trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost7"); ok || d.HasChange("ipv6_trusthost7") {
		t, err := expandSystemAdminProfileIpv6Trusthost7(d, v, "ipv6_trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost8"); ok || d.HasChange("ipv6_trusthost8") {
		t, err := expandSystemAdminProfileIpv6Trusthost8(d, v, "ipv6_trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost9"); ok || d.HasChange("ipv6_trusthost9") {
		t, err := expandSystemAdminProfileIpv6Trusthost9(d, v, "ipv6_trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("log_viewer"); ok || d.HasChange("log_viewer") {
		t, err := expandSystemAdminProfileLogViewer(d, v, "log_viewer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-viewer"] = t
		}
	}

	if v, ok := d.GetOk("policy_objects"); ok || d.HasChange("policy_objects") {
		t, err := expandSystemAdminProfilePolicyObjects(d, v, "policy_objects")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-objects"] = t
		}
	}

	if v, ok := d.GetOk("profileid"); ok || d.HasChange("profileid") {
		t, err := expandSystemAdminProfileProfileid(d, v, "profileid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profileid"] = t
		}
	}

	if v, ok := d.GetOk("read_passwd"); ok || d.HasChange("read_passwd") {
		t, err := expandSystemAdminProfileReadPasswd(d, v, "read_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["read-passwd"] = t
		}
	}

	if v, ok := d.GetOk("realtime_monitor"); ok || d.HasChange("realtime_monitor") {
		t, err := expandSystemAdminProfileRealtimeMonitor(d, v, "realtime_monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realtime-monitor"] = t
		}
	}

	if v, ok := d.GetOk("report_viewer"); ok || d.HasChange("report_viewer") {
		t, err := expandSystemAdminProfileReportViewer(d, v, "report_viewer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-viewer"] = t
		}
	}

	if v, ok := d.GetOk("rpc_permit"); ok || d.HasChange("rpc_permit") {
		t, err := expandSystemAdminProfileRpcPermit(d, v, "rpc_permit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-permit"] = t
		}
	}

	if v, ok := d.GetOk("run_report"); ok || d.HasChange("run_report") {
		t, err := expandSystemAdminProfileRunReport(d, v, "run_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["run-report"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok || d.HasChange("scope") {
		t, err := expandSystemAdminProfileScope(d, v, "scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	if v, ok := d.GetOk("script_access"); ok || d.HasChange("script_access") {
		t, err := expandSystemAdminProfileScriptAccess(d, v, "script_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script-access"] = t
		}
	}

	if v, ok := d.GetOk("set_install_targets"); ok || d.HasChange("set_install_targets") {
		t, err := expandSystemAdminProfileSetInstallTargets(d, v, "set_install_targets")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-install-targets"] = t
		}
	}

	if v, ok := d.GetOk("super_user_profile"); ok || d.HasChange("super_user_profile") {
		t, err := expandSystemAdminProfileSuperUserProfile(d, v, "super_user_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["super-user-profile"] = t
		}
	}

	if v, ok := d.GetOk("system_setting"); ok || d.HasChange("system_setting") {
		t, err := expandSystemAdminProfileSystemSetting(d, v, "system_setting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-setting"] = t
		}
	}

	if v, ok := d.GetOk("term_access"); ok || d.HasChange("term_access") {
		t, err := expandSystemAdminProfileTermAccess(d, v, "term_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["term-access"] = t
		}
	}

	if v, ok := d.GetOk("triage_events"); ok || d.HasChange("triage_events") {
		t, err := expandSystemAdminProfileTriageEvents(d, v, "triage_events")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["triage-events"] = t
		}
	}

	if v, ok := d.GetOk("trusthost1"); ok || d.HasChange("trusthost1") {
		t, err := expandSystemAdminProfileTrusthost1(d, v, "trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("trusthost10"); ok || d.HasChange("trusthost10") {
		t, err := expandSystemAdminProfileTrusthost10(d, v, "trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("trusthost2"); ok || d.HasChange("trusthost2") {
		t, err := expandSystemAdminProfileTrusthost2(d, v, "trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("trusthost3"); ok || d.HasChange("trusthost3") {
		t, err := expandSystemAdminProfileTrusthost3(d, v, "trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("trusthost4"); ok || d.HasChange("trusthost4") {
		t, err := expandSystemAdminProfileTrusthost4(d, v, "trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("trusthost5"); ok || d.HasChange("trusthost5") {
		t, err := expandSystemAdminProfileTrusthost5(d, v, "trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("trusthost6"); ok || d.HasChange("trusthost6") {
		t, err := expandSystemAdminProfileTrusthost6(d, v, "trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("trusthost7"); ok || d.HasChange("trusthost7") {
		t, err := expandSystemAdminProfileTrusthost7(d, v, "trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("trusthost8"); ok || d.HasChange("trusthost8") {
		t, err := expandSystemAdminProfileTrusthost8(d, v, "trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("trusthost9"); ok || d.HasChange("trusthost9") {
		t, err := expandSystemAdminProfileTrusthost9(d, v, "trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemAdminProfileType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("update_incidents"); ok || d.HasChange("update_incidents") {
		t, err := expandSystemAdminProfileUpdateIncidents(d, v, "update_incidents")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-incidents"] = t
		}
	}

	if v, ok := d.GetOk("vpn_manager"); ok || d.HasChange("vpn_manager") {
		t, err := expandSystemAdminProfileVpnManager(d, v, "vpn_manager")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-manager"] = t
		}
	}

	if v, ok := d.GetOk("web_filter"); ok || d.HasChange("web_filter") {
		t, err := expandSystemAdminProfileWebFilter(d, v, "web_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter"] = t
		}
	}

	return &obj, nil
}
