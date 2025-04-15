// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Admin user.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminUserCreate,
		Read:   resourceSystemAdminUserRead,
		Update: resourceSystemAdminUserUpdate,
		Delete: resourceSystemAdminUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fmgadom": &schema.Schema{
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
			"adom_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adom_exclude": &schema.Schema{
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
			"app_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_filter_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"avatar": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"change_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cors_allow_origin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dashboard": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"column": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"diskio_content_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"diskio_period": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_rate_period": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_rate_topn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_rate_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"moduleid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"num_entries": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"refresh_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"res_cpu_display": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"res_period": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"res_view_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tabid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"time_period": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"widget_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dashboard_tabs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tabid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dev_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ext_auth_accprofile_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ext_auth_adom_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ext_auth_group_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"force_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiai": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hidden": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ips_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ips_filter_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"login_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"meta_data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fieldlength": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fieldname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fieldvalue": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"importance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mobile_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pager_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password_expire": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phone_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_block": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_block_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"policy_package": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_package_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"profileid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_permit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_public_key1": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_public_key2": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_public_key3": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"subject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tacacs_plus_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"th_from_profile": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"th6_from_profile": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"two_factor_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_global_theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"userid": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"web_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"web_filter_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"wildcard": &schema.Schema{
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

func resourceSystemAdminUserCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemAdminUser(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminUser resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemAdminUser(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminUser resource: %v", err)
	}

	d.SetId(getStringKey(d, "userid"))

	return resourceSystemAdminUserRead(d, m)
}

func resourceSystemAdminUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemAdminUser(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminUser resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAdminUser(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminUser resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "userid"))

	return resourceSystemAdminUserRead(d, m)
}

func resourceSystemAdminUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemAdminUser(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemAdminUser(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminUser(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminUser resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminUserAdom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemAdminUserAdomAdomName(i["adom-name"], d, pre_append)
			tmp["adom_name"] = fortiAPISubPartPatch(v, "SystemAdminUser-Adom-AdomName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminUserAdomAdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserAdomAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserAdomExclude(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemAdminUserAdomExcludeAdomName(i["adom-name"], d, pre_append)
			tmp["adom_name"] = fortiAPISubPartPatch(v, "SystemAdminUser-AdomExclude-AdomName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminUserAdomExcludeAdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserAppFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_filter_name"
		if _, ok := i["app-filter-name"]; ok {
			v := flattenSystemAdminUserAppFilterAppFilterName(i["app-filter-name"], d, pre_append)
			tmp["app_filter_name"] = fortiAPISubPartPatch(v, "SystemAdminUser-AppFilter-AppFilterName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminUserAppFilterAppFilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserAvatar(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserChangePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserCorsAllowOrigin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboard(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "column"
		if _, ok := i["column"]; ok {
			v := flattenSystemAdminUserDashboardColumn(i["column"], d, pre_append)
			tmp["column"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-Column")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diskio_content_type"
		if _, ok := i["diskio-content-type"]; ok {
			v := flattenSystemAdminUserDashboardDiskioContentType(i["diskio-content-type"], d, pre_append)
			tmp["diskio_content_type"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-DiskioContentType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diskio_period"
		if _, ok := i["diskio-period"]; ok {
			v := flattenSystemAdminUserDashboardDiskioPeriod(i["diskio-period"], d, pre_append)
			tmp["diskio_period"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-DiskioPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_rate_period"
		if _, ok := i["log-rate-period"]; ok {
			v := flattenSystemAdminUserDashboardLogRatePeriod(i["log-rate-period"], d, pre_append)
			tmp["log_rate_period"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-LogRatePeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_rate_topn"
		if _, ok := i["log-rate-topn"]; ok {
			v := flattenSystemAdminUserDashboardLogRateTopn(i["log-rate-topn"], d, pre_append)
			tmp["log_rate_topn"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-LogRateTopn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_rate_type"
		if _, ok := i["log-rate-type"]; ok {
			v := flattenSystemAdminUserDashboardLogRateType(i["log-rate-type"], d, pre_append)
			tmp["log_rate_type"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-LogRateType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "moduleid"
		if _, ok := i["moduleid"]; ok {
			v := flattenSystemAdminUserDashboardModuleid(i["moduleid"], d, pre_append)
			tmp["moduleid"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-Moduleid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemAdminUserDashboardName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "num_entries"
		if _, ok := i["num-entries"]; ok {
			v := flattenSystemAdminUserDashboardNumEntries(i["num-entries"], d, pre_append)
			tmp["num_entries"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-NumEntries")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "refresh_interval"
		if _, ok := i["refresh-interval"]; ok {
			v := flattenSystemAdminUserDashboardRefreshInterval(i["refresh-interval"], d, pre_append)
			tmp["refresh_interval"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-RefreshInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "res_cpu_display"
		if _, ok := i["res-cpu-display"]; ok {
			v := flattenSystemAdminUserDashboardResCpuDisplay(i["res-cpu-display"], d, pre_append)
			tmp["res_cpu_display"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-ResCpuDisplay")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "res_period"
		if _, ok := i["res-period"]; ok {
			v := flattenSystemAdminUserDashboardResPeriod(i["res-period"], d, pre_append)
			tmp["res_period"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-ResPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "res_view_type"
		if _, ok := i["res-view-type"]; ok {
			v := flattenSystemAdminUserDashboardResViewType(i["res-view-type"], d, pre_append)
			tmp["res_view_type"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-ResViewType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemAdminUserDashboardStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tabid"
		if _, ok := i["tabid"]; ok {
			v := flattenSystemAdminUserDashboardTabid(i["tabid"], d, pre_append)
			tmp["tabid"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-Tabid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time_period"
		if _, ok := i["time-period"]; ok {
			v := flattenSystemAdminUserDashboardTimePeriod(i["time-period"], d, pre_append)
			tmp["time_period"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-TimePeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "widget_type"
		if _, ok := i["widget-type"]; ok {
			v := flattenSystemAdminUserDashboardWidgetType(i["widget-type"], d, pre_append)
			tmp["widget_type"] = fortiAPISubPartPatch(v, "SystemAdminUser-Dashboard-WidgetType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminUserDashboardColumn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardDiskioContentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardDiskioPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardLogRatePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardLogRateTopn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardLogRateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardModuleid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardNumEntries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardRefreshInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardResCpuDisplay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardResPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardResViewType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardTabid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardTimePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardWidgetType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardTabs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemAdminUserDashboardTabsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemAdminUser-DashboardTabs-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tabid"
		if _, ok := i["tabid"]; ok {
			v := flattenSystemAdminUserDashboardTabsTabid(i["tabid"], d, pre_append)
			tmp["tabid"] = fortiAPISubPartPatch(v, "SystemAdminUser-DashboardTabs-Tabid")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminUserDashboardTabsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDashboardTabsTabid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserDevGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserEmailAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserExtAuthAccprofileOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserExtAuthAdomOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserExtAuthGroupMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserFingerprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserFirstName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserForcePasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserHidden(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpsFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ips_filter_name"
		if _, ok := i["ips-filter-name"]; ok {
			v := flattenSystemAdminUserIpsFilterIpsFilterName(i["ips-filter-name"], d, pre_append)
			tmp["ips_filter_name"] = fortiAPISubPartPatch(v, "SystemAdminUser-IpsFilter-IpsFilterName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminUserIpsFilterIpsFilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpv6Trusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpv6Trusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpv6Trusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpv6Trusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpv6Trusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpv6Trusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpv6Trusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpv6Trusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpv6Trusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserIpv6Trusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserLastName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserLoginMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserMetaData(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldlength"
		if _, ok := i["fieldlength"]; ok {
			v := flattenSystemAdminUserMetaDataFieldlength(i["fieldlength"], d, pre_append)
			tmp["fieldlength"] = fortiAPISubPartPatch(v, "SystemAdminUser-MetaData-Fieldlength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldname"
		if _, ok := i["fieldname"]; ok {
			v := flattenSystemAdminUserMetaDataFieldname(i["fieldname"], d, pre_append)
			tmp["fieldname"] = fortiAPISubPartPatch(v, "SystemAdminUser-MetaData-Fieldname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldvalue"
		if _, ok := i["fieldvalue"]; ok {
			v := flattenSystemAdminUserMetaDataFieldvalue(i["fieldvalue"], d, pre_append)
			tmp["fieldvalue"] = fortiAPISubPartPatch(v, "SystemAdminUser-MetaData-Fieldvalue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "importance"
		if _, ok := i["importance"]; ok {
			v := flattenSystemAdminUserMetaDataImportance(i["importance"], d, pre_append)
			tmp["importance"] = fortiAPISubPartPatch(v, "SystemAdminUser-MetaData-Importance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemAdminUserMetaDataStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemAdminUser-MetaData-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminUserMetaDataFieldlength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserMetaDataFieldname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserMetaDataFieldvalue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserMetaDataImportance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserMetaDataStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserMobileNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserPagerNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenSystemAdminUserPasswordExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstrlist2str(v)
}

func flattenSystemAdminUserPhoneNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserPolicyBlock(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_block_name"
		if _, ok := i["policy-block-name"]; ok {
			v := flattenSystemAdminUserPolicyBlockPolicyBlockName(i["policy-block-name"], d, pre_append)
			tmp["policy_block_name"] = fortiAPISubPartPatch(v, "SystemAdminUser-PolicyBlock-PolicyBlockName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminUserPolicyBlockPolicyBlockName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserPolicyPackage(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_package_name"
		if _, ok := i["policy-package-name"]; ok {
			v := flattenSystemAdminUserPolicyPackagePolicyPackageName(i["policy-package-name"], d, pre_append)
			tmp["policy_package_name"] = fortiAPISubPartPatch(v, "SystemAdminUser-PolicyPackage-PolicyPackageName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminUserPolicyPackagePolicyPackageName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserProfileid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserRpcPermit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserSshPublicKey1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserSshPublicKey2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserSshPublicKey3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserSubject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserTacacsPlusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserThFromProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserTh6FromProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserTrusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserTrusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserTrusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserTrusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserTrusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserTrusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserTrusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserTrusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserTrusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserTrusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminUserTwoFactorAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserUseGlobalTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserUserTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserUserType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserUserid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserWebFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "web_filter_name"
		if _, ok := i["web-filter-name"]; ok {
			v := flattenSystemAdminUserWebFilterWebFilterName(i["web-filter-name"], d, pre_append)
			tmp["web_filter_name"] = fortiAPISubPartPatch(v, "SystemAdminUser-WebFilter-WebFilterName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminUserWebFilterWebFilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminUserWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("fmgadom", flattenSystemAdminUserAdom(o["adom"], d, "fmgadom")); err != nil {
			if vv, ok := fortiAPIPatch(o["adom"], "SystemAdminUser-Adom"); ok {
				if err = d.Set("fmgadom", vv); err != nil {
					return fmt.Errorf("Error reading fmgadom: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fmgadom"); ok {
			if err = d.Set("fmgadom", flattenSystemAdminUserAdom(o["adom"], d, "fmgadom")); err != nil {
				if vv, ok := fortiAPIPatch(o["adom"], "SystemAdminUser-Adom"); ok {
					if err = d.Set("fmgadom", vv); err != nil {
						return fmt.Errorf("Error reading fmgadom: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fmgadom: %v", err)
				}
			}
		}
	}

	if err = d.Set("adom_access", flattenSystemAdminUserAdomAccess(o["adom-access"], d, "adom_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-access"], "SystemAdminUser-AdomAccess"); ok {
			if err = d.Set("adom_access", vv); err != nil {
				return fmt.Errorf("Error reading adom_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_access: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("adom_exclude", flattenSystemAdminUserAdomExclude(o["adom-exclude"], d, "adom_exclude")); err != nil {
			if vv, ok := fortiAPIPatch(o["adom-exclude"], "SystemAdminUser-AdomExclude"); ok {
				if err = d.Set("adom_exclude", vv); err != nil {
					return fmt.Errorf("Error reading adom_exclude: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading adom_exclude: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("adom_exclude"); ok {
			if err = d.Set("adom_exclude", flattenSystemAdminUserAdomExclude(o["adom-exclude"], d, "adom_exclude")); err != nil {
				if vv, ok := fortiAPIPatch(o["adom-exclude"], "SystemAdminUser-AdomExclude"); ok {
					if err = d.Set("adom_exclude", vv); err != nil {
						return fmt.Errorf("Error reading adom_exclude: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading adom_exclude: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("app_filter", flattenSystemAdminUserAppFilter(o["app-filter"], d, "app_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["app-filter"], "SystemAdminUser-AppFilter"); ok {
				if err = d.Set("app_filter", vv); err != nil {
					return fmt.Errorf("Error reading app_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading app_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_filter"); ok {
			if err = d.Set("app_filter", flattenSystemAdminUserAppFilter(o["app-filter"], d, "app_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["app-filter"], "SystemAdminUser-AppFilter"); ok {
					if err = d.Set("app_filter", vv); err != nil {
						return fmt.Errorf("Error reading app_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading app_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("avatar", flattenSystemAdminUserAvatar(o["avatar"], d, "avatar")); err != nil {
		if vv, ok := fortiAPIPatch(o["avatar"], "SystemAdminUser-Avatar"); ok {
			if err = d.Set("avatar", vv); err != nil {
				return fmt.Errorf("Error reading avatar: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading avatar: %v", err)
		}
	}

	if err = d.Set("ca", flattenSystemAdminUserCa(o["ca"], d, "ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca"], "SystemAdminUser-Ca"); ok {
			if err = d.Set("ca", vv); err != nil {
				return fmt.Errorf("Error reading ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("change_password", flattenSystemAdminUserChangePassword(o["change-password"], d, "change_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["change-password"], "SystemAdminUser-ChangePassword"); ok {
			if err = d.Set("change_password", vv); err != nil {
				return fmt.Errorf("Error reading change_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading change_password: %v", err)
		}
	}

	if err = d.Set("cors_allow_origin", flattenSystemAdminUserCorsAllowOrigin(o["cors-allow-origin"], d, "cors_allow_origin")); err != nil {
		if vv, ok := fortiAPIPatch(o["cors-allow-origin"], "SystemAdminUser-CorsAllowOrigin"); ok {
			if err = d.Set("cors_allow_origin", vv); err != nil {
				return fmt.Errorf("Error reading cors_allow_origin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cors_allow_origin: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dashboard", flattenSystemAdminUserDashboard(o["dashboard"], d, "dashboard")); err != nil {
			if vv, ok := fortiAPIPatch(o["dashboard"], "SystemAdminUser-Dashboard"); ok {
				if err = d.Set("dashboard", vv); err != nil {
					return fmt.Errorf("Error reading dashboard: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dashboard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dashboard"); ok {
			if err = d.Set("dashboard", flattenSystemAdminUserDashboard(o["dashboard"], d, "dashboard")); err != nil {
				if vv, ok := fortiAPIPatch(o["dashboard"], "SystemAdminUser-Dashboard"); ok {
					if err = d.Set("dashboard", vv); err != nil {
						return fmt.Errorf("Error reading dashboard: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dashboard: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dashboard_tabs", flattenSystemAdminUserDashboardTabs(o["dashboard-tabs"], d, "dashboard_tabs")); err != nil {
			if vv, ok := fortiAPIPatch(o["dashboard-tabs"], "SystemAdminUser-DashboardTabs"); ok {
				if err = d.Set("dashboard_tabs", vv); err != nil {
					return fmt.Errorf("Error reading dashboard_tabs: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dashboard_tabs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dashboard_tabs"); ok {
			if err = d.Set("dashboard_tabs", flattenSystemAdminUserDashboardTabs(o["dashboard-tabs"], d, "dashboard_tabs")); err != nil {
				if vv, ok := fortiAPIPatch(o["dashboard-tabs"], "SystemAdminUser-DashboardTabs"); ok {
					if err = d.Set("dashboard_tabs", vv); err != nil {
						return fmt.Errorf("Error reading dashboard_tabs: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dashboard_tabs: %v", err)
				}
			}
		}
	}

	if err = d.Set("description", flattenSystemAdminUserDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemAdminUser-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("dev_group", flattenSystemAdminUserDevGroup(o["dev-group"], d, "dev_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev-group"], "SystemAdminUser-DevGroup"); ok {
			if err = d.Set("dev_group", vv); err != nil {
				return fmt.Errorf("Error reading dev_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_group: %v", err)
		}
	}

	if err = d.Set("email_address", flattenSystemAdminUserEmailAddress(o["email-address"], d, "email_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-address"], "SystemAdminUser-EmailAddress"); ok {
			if err = d.Set("email_address", vv); err != nil {
				return fmt.Errorf("Error reading email_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_address: %v", err)
		}
	}

	if err = d.Set("ext_auth_accprofile_override", flattenSystemAdminUserExtAuthAccprofileOverride(o["ext-auth-accprofile-override"], d, "ext_auth_accprofile_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ext-auth-accprofile-override"], "SystemAdminUser-ExtAuthAccprofileOverride"); ok {
			if err = d.Set("ext_auth_accprofile_override", vv); err != nil {
				return fmt.Errorf("Error reading ext_auth_accprofile_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ext_auth_accprofile_override: %v", err)
		}
	}

	if err = d.Set("ext_auth_adom_override", flattenSystemAdminUserExtAuthAdomOverride(o["ext-auth-adom-override"], d, "ext_auth_adom_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ext-auth-adom-override"], "SystemAdminUser-ExtAuthAdomOverride"); ok {
			if err = d.Set("ext_auth_adom_override", vv); err != nil {
				return fmt.Errorf("Error reading ext_auth_adom_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ext_auth_adom_override: %v", err)
		}
	}

	if err = d.Set("ext_auth_group_match", flattenSystemAdminUserExtAuthGroupMatch(o["ext-auth-group-match"], d, "ext_auth_group_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["ext-auth-group-match"], "SystemAdminUser-ExtAuthGroupMatch"); ok {
			if err = d.Set("ext_auth_group_match", vv); err != nil {
				return fmt.Errorf("Error reading ext_auth_group_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ext_auth_group_match: %v", err)
		}
	}

	if err = d.Set("fingerprint", flattenSystemAdminUserFingerprint(o["fingerprint"], d, "fingerprint")); err != nil {
		if vv, ok := fortiAPIPatch(o["fingerprint"], "SystemAdminUser-Fingerprint"); ok {
			if err = d.Set("fingerprint", vv); err != nil {
				return fmt.Errorf("Error reading fingerprint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fingerprint: %v", err)
		}
	}

	if err = d.Set("first_name", flattenSystemAdminUserFirstName(o["first-name"], d, "first_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["first-name"], "SystemAdminUser-FirstName"); ok {
			if err = d.Set("first_name", vv); err != nil {
				return fmt.Errorf("Error reading first_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading first_name: %v", err)
		}
	}

	if err = d.Set("force_password_change", flattenSystemAdminUserForcePasswordChange(o["force-password-change"], d, "force_password_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["force-password-change"], "SystemAdminUser-ForcePasswordChange"); ok {
			if err = d.Set("force_password_change", vv); err != nil {
				return fmt.Errorf("Error reading force_password_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading force_password_change: %v", err)
		}
	}

	if err = d.Set("fortiai", flattenSystemAdminUserFortiai(o["fortiai"], d, "fortiai")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiai"], "SystemAdminUser-Fortiai"); ok {
			if err = d.Set("fortiai", vv); err != nil {
				return fmt.Errorf("Error reading fortiai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiai: %v", err)
		}
	}

	if err = d.Set("group", flattenSystemAdminUserGroup(o["group"], d, "group")); err != nil {
		if vv, ok := fortiAPIPatch(o["group"], "SystemAdminUser-Group"); ok {
			if err = d.Set("group", vv); err != nil {
				return fmt.Errorf("Error reading group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}

	if err = d.Set("hidden", flattenSystemAdminUserHidden(o["hidden"], d, "hidden")); err != nil {
		if vv, ok := fortiAPIPatch(o["hidden"], "SystemAdminUser-Hidden"); ok {
			if err = d.Set("hidden", vv); err != nil {
				return fmt.Errorf("Error reading hidden: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hidden: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ips_filter", flattenSystemAdminUserIpsFilter(o["ips-filter"], d, "ips_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["ips-filter"], "SystemAdminUser-IpsFilter"); ok {
				if err = d.Set("ips_filter", vv); err != nil {
					return fmt.Errorf("Error reading ips_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ips_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ips_filter"); ok {
			if err = d.Set("ips_filter", flattenSystemAdminUserIpsFilter(o["ips-filter"], d, "ips_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["ips-filter"], "SystemAdminUser-IpsFilter"); ok {
					if err = d.Set("ips_filter", vv); err != nil {
						return fmt.Errorf("Error reading ips_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ips_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv6_trusthost1", flattenSystemAdminUserIpv6Trusthost1(o["ipv6_trusthost1"], d, "ipv6_trusthost1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost1"], "SystemAdminUser-Ipv6Trusthost1"); ok {
			if err = d.Set("ipv6_trusthost1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost1: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost10", flattenSystemAdminUserIpv6Trusthost10(o["ipv6_trusthost10"], d, "ipv6_trusthost10")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost10"], "SystemAdminUser-Ipv6Trusthost10"); ok {
			if err = d.Set("ipv6_trusthost10", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost10: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost2", flattenSystemAdminUserIpv6Trusthost2(o["ipv6_trusthost2"], d, "ipv6_trusthost2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost2"], "SystemAdminUser-Ipv6Trusthost2"); ok {
			if err = d.Set("ipv6_trusthost2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost2: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost3", flattenSystemAdminUserIpv6Trusthost3(o["ipv6_trusthost3"], d, "ipv6_trusthost3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost3"], "SystemAdminUser-Ipv6Trusthost3"); ok {
			if err = d.Set("ipv6_trusthost3", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost3: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost4", flattenSystemAdminUserIpv6Trusthost4(o["ipv6_trusthost4"], d, "ipv6_trusthost4")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost4"], "SystemAdminUser-Ipv6Trusthost4"); ok {
			if err = d.Set("ipv6_trusthost4", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost4: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost5", flattenSystemAdminUserIpv6Trusthost5(o["ipv6_trusthost5"], d, "ipv6_trusthost5")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost5"], "SystemAdminUser-Ipv6Trusthost5"); ok {
			if err = d.Set("ipv6_trusthost5", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost5: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost6", flattenSystemAdminUserIpv6Trusthost6(o["ipv6_trusthost6"], d, "ipv6_trusthost6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost6"], "SystemAdminUser-Ipv6Trusthost6"); ok {
			if err = d.Set("ipv6_trusthost6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost6: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost7", flattenSystemAdminUserIpv6Trusthost7(o["ipv6_trusthost7"], d, "ipv6_trusthost7")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost7"], "SystemAdminUser-Ipv6Trusthost7"); ok {
			if err = d.Set("ipv6_trusthost7", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost7: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost8", flattenSystemAdminUserIpv6Trusthost8(o["ipv6_trusthost8"], d, "ipv6_trusthost8")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost8"], "SystemAdminUser-Ipv6Trusthost8"); ok {
			if err = d.Set("ipv6_trusthost8", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost8: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost9", flattenSystemAdminUserIpv6Trusthost9(o["ipv6_trusthost9"], d, "ipv6_trusthost9")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6_trusthost9"], "SystemAdminUser-Ipv6Trusthost9"); ok {
			if err = d.Set("ipv6_trusthost9", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost9: %v", err)
		}
	}

	if err = d.Set("last_name", flattenSystemAdminUserLastName(o["last-name"], d, "last_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-name"], "SystemAdminUser-LastName"); ok {
			if err = d.Set("last_name", vv); err != nil {
				return fmt.Errorf("Error reading last_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_name: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenSystemAdminUserLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "SystemAdminUser-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("login_max", flattenSystemAdminUserLoginMax(o["login-max"], d, "login_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-max"], "SystemAdminUser-LoginMax"); ok {
			if err = d.Set("login_max", vv); err != nil {
				return fmt.Errorf("Error reading login_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_max: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("meta_data", flattenSystemAdminUserMetaData(o["meta-data"], d, "meta_data")); err != nil {
			if vv, ok := fortiAPIPatch(o["meta-data"], "SystemAdminUser-MetaData"); ok {
				if err = d.Set("meta_data", vv); err != nil {
					return fmt.Errorf("Error reading meta_data: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading meta_data: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("meta_data"); ok {
			if err = d.Set("meta_data", flattenSystemAdminUserMetaData(o["meta-data"], d, "meta_data")); err != nil {
				if vv, ok := fortiAPIPatch(o["meta-data"], "SystemAdminUser-MetaData"); ok {
					if err = d.Set("meta_data", vv); err != nil {
						return fmt.Errorf("Error reading meta_data: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading meta_data: %v", err)
				}
			}
		}
	}

	if err = d.Set("mobile_number", flattenSystemAdminUserMobileNumber(o["mobile-number"], d, "mobile_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["mobile-number"], "SystemAdminUser-MobileNumber"); ok {
			if err = d.Set("mobile_number", vv); err != nil {
				return fmt.Errorf("Error reading mobile_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mobile_number: %v", err)
		}
	}

	if err = d.Set("pager_number", flattenSystemAdminUserPagerNumber(o["pager-number"], d, "pager_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["pager-number"], "SystemAdminUser-PagerNumber"); ok {
			if err = d.Set("pager_number", vv); err != nil {
				return fmt.Errorf("Error reading pager_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pager_number: %v", err)
		}
	}

	if err = d.Set("password_expire", flattenSystemAdminUserPasswordExpire(o["password-expire"], d, "password_expire")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-expire"], "SystemAdminUser-PasswordExpire"); ok {
			if err = d.Set("password_expire", vv); err != nil {
				return fmt.Errorf("Error reading password_expire: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_expire: %v", err)
		}
	}

	if err = d.Set("phone_number", flattenSystemAdminUserPhoneNumber(o["phone-number"], d, "phone_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["phone-number"], "SystemAdminUser-PhoneNumber"); ok {
			if err = d.Set("phone_number", vv); err != nil {
				return fmt.Errorf("Error reading phone_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phone_number: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("policy_block", flattenSystemAdminUserPolicyBlock(o["policy-block"], d, "policy_block")); err != nil {
			if vv, ok := fortiAPIPatch(o["policy-block"], "SystemAdminUser-PolicyBlock"); ok {
				if err = d.Set("policy_block", vv); err != nil {
					return fmt.Errorf("Error reading policy_block: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading policy_block: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("policy_block"); ok {
			if err = d.Set("policy_block", flattenSystemAdminUserPolicyBlock(o["policy-block"], d, "policy_block")); err != nil {
				if vv, ok := fortiAPIPatch(o["policy-block"], "SystemAdminUser-PolicyBlock"); ok {
					if err = d.Set("policy_block", vv); err != nil {
						return fmt.Errorf("Error reading policy_block: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading policy_block: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("policy_package", flattenSystemAdminUserPolicyPackage(o["policy-package"], d, "policy_package")); err != nil {
			if vv, ok := fortiAPIPatch(o["policy-package"], "SystemAdminUser-PolicyPackage"); ok {
				if err = d.Set("policy_package", vv); err != nil {
					return fmt.Errorf("Error reading policy_package: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading policy_package: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("policy_package"); ok {
			if err = d.Set("policy_package", flattenSystemAdminUserPolicyPackage(o["policy-package"], d, "policy_package")); err != nil {
				if vv, ok := fortiAPIPatch(o["policy-package"], "SystemAdminUser-PolicyPackage"); ok {
					if err = d.Set("policy_package", vv); err != nil {
						return fmt.Errorf("Error reading policy_package: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading policy_package: %v", err)
				}
			}
		}
	}

	if err = d.Set("profileid", flattenSystemAdminUserProfileid(o["profileid"], d, "profileid")); err != nil {
		if vv, ok := fortiAPIPatch(o["profileid"], "SystemAdminUser-Profileid"); ok {
			if err = d.Set("profileid", vv); err != nil {
				return fmt.Errorf("Error reading profileid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profileid: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenSystemAdminUserRadiusServer(o["radius_server"], d, "radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius_server"], "SystemAdminUser-RadiusServer"); ok {
			if err = d.Set("radius_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("rpc_permit", flattenSystemAdminUserRpcPermit(o["rpc-permit"], d, "rpc_permit")); err != nil {
		if vv, ok := fortiAPIPatch(o["rpc-permit"], "SystemAdminUser-RpcPermit"); ok {
			if err = d.Set("rpc_permit", vv); err != nil {
				return fmt.Errorf("Error reading rpc_permit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rpc_permit: %v", err)
		}
	}

	if err = d.Set("ssh_public_key1", flattenSystemAdminUserSshPublicKey1(o["ssh-public-key1"], d, "ssh_public_key1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-public-key1"], "SystemAdminUser-SshPublicKey1"); ok {
			if err = d.Set("ssh_public_key1", vv); err != nil {
				return fmt.Errorf("Error reading ssh_public_key1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_public_key1: %v", err)
		}
	}

	if err = d.Set("ssh_public_key2", flattenSystemAdminUserSshPublicKey2(o["ssh-public-key2"], d, "ssh_public_key2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-public-key2"], "SystemAdminUser-SshPublicKey2"); ok {
			if err = d.Set("ssh_public_key2", vv); err != nil {
				return fmt.Errorf("Error reading ssh_public_key2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_public_key2: %v", err)
		}
	}

	if err = d.Set("ssh_public_key3", flattenSystemAdminUserSshPublicKey3(o["ssh-public-key3"], d, "ssh_public_key3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-public-key3"], "SystemAdminUser-SshPublicKey3"); ok {
			if err = d.Set("ssh_public_key3", vv); err != nil {
				return fmt.Errorf("Error reading ssh_public_key3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_public_key3: %v", err)
		}
	}

	if err = d.Set("subject", flattenSystemAdminUserSubject(o["subject"], d, "subject")); err != nil {
		if vv, ok := fortiAPIPatch(o["subject"], "SystemAdminUser-Subject"); ok {
			if err = d.Set("subject", vv); err != nil {
				return fmt.Errorf("Error reading subject: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subject: %v", err)
		}
	}

	if err = d.Set("tacacs_plus_server", flattenSystemAdminUserTacacsPlusServer(o["tacacs-plus-server"], d, "tacacs_plus_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tacacs-plus-server"], "SystemAdminUser-TacacsPlusServer"); ok {
			if err = d.Set("tacacs_plus_server", vv); err != nil {
				return fmt.Errorf("Error reading tacacs_plus_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tacacs_plus_server: %v", err)
		}
	}

	if err = d.Set("th_from_profile", flattenSystemAdminUserThFromProfile(o["th-from-profile"], d, "th_from_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["th-from-profile"], "SystemAdminUser-ThFromProfile"); ok {
			if err = d.Set("th_from_profile", vv); err != nil {
				return fmt.Errorf("Error reading th_from_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading th_from_profile: %v", err)
		}
	}

	if err = d.Set("th6_from_profile", flattenSystemAdminUserTh6FromProfile(o["th6-from-profile"], d, "th6_from_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["th6-from-profile"], "SystemAdminUser-Th6FromProfile"); ok {
			if err = d.Set("th6_from_profile", vv); err != nil {
				return fmt.Errorf("Error reading th6_from_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading th6_from_profile: %v", err)
		}
	}

	if err = d.Set("trusthost1", flattenSystemAdminUserTrusthost1(o["trusthost1"], d, "trusthost1")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost1"], "SystemAdminUser-Trusthost1"); ok {
			if err = d.Set("trusthost1", vv); err != nil {
				return fmt.Errorf("Error reading trusthost1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost1: %v", err)
		}
	}

	if err = d.Set("trusthost10", flattenSystemAdminUserTrusthost10(o["trusthost10"], d, "trusthost10")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost10"], "SystemAdminUser-Trusthost10"); ok {
			if err = d.Set("trusthost10", vv); err != nil {
				return fmt.Errorf("Error reading trusthost10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost10: %v", err)
		}
	}

	if err = d.Set("trusthost2", flattenSystemAdminUserTrusthost2(o["trusthost2"], d, "trusthost2")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost2"], "SystemAdminUser-Trusthost2"); ok {
			if err = d.Set("trusthost2", vv); err != nil {
				return fmt.Errorf("Error reading trusthost2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost2: %v", err)
		}
	}

	if err = d.Set("trusthost3", flattenSystemAdminUserTrusthost3(o["trusthost3"], d, "trusthost3")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost3"], "SystemAdminUser-Trusthost3"); ok {
			if err = d.Set("trusthost3", vv); err != nil {
				return fmt.Errorf("Error reading trusthost3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost3: %v", err)
		}
	}

	if err = d.Set("trusthost4", flattenSystemAdminUserTrusthost4(o["trusthost4"], d, "trusthost4")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost4"], "SystemAdminUser-Trusthost4"); ok {
			if err = d.Set("trusthost4", vv); err != nil {
				return fmt.Errorf("Error reading trusthost4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost4: %v", err)
		}
	}

	if err = d.Set("trusthost5", flattenSystemAdminUserTrusthost5(o["trusthost5"], d, "trusthost5")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost5"], "SystemAdminUser-Trusthost5"); ok {
			if err = d.Set("trusthost5", vv); err != nil {
				return fmt.Errorf("Error reading trusthost5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost5: %v", err)
		}
	}

	if err = d.Set("trusthost6", flattenSystemAdminUserTrusthost6(o["trusthost6"], d, "trusthost6")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost6"], "SystemAdminUser-Trusthost6"); ok {
			if err = d.Set("trusthost6", vv); err != nil {
				return fmt.Errorf("Error reading trusthost6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost6: %v", err)
		}
	}

	if err = d.Set("trusthost7", flattenSystemAdminUserTrusthost7(o["trusthost7"], d, "trusthost7")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost7"], "SystemAdminUser-Trusthost7"); ok {
			if err = d.Set("trusthost7", vv); err != nil {
				return fmt.Errorf("Error reading trusthost7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost7: %v", err)
		}
	}

	if err = d.Set("trusthost8", flattenSystemAdminUserTrusthost8(o["trusthost8"], d, "trusthost8")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost8"], "SystemAdminUser-Trusthost8"); ok {
			if err = d.Set("trusthost8", vv); err != nil {
				return fmt.Errorf("Error reading trusthost8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost8: %v", err)
		}
	}

	if err = d.Set("trusthost9", flattenSystemAdminUserTrusthost9(o["trusthost9"], d, "trusthost9")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost9"], "SystemAdminUser-Trusthost9"); ok {
			if err = d.Set("trusthost9", vv); err != nil {
				return fmt.Errorf("Error reading trusthost9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost9: %v", err)
		}
	}

	if err = d.Set("two_factor_auth", flattenSystemAdminUserTwoFactorAuth(o["two-factor-auth"], d, "two_factor_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-auth"], "SystemAdminUser-TwoFactorAuth"); ok {
			if err = d.Set("two_factor_auth", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_auth: %v", err)
		}
	}

	if err = d.Set("use_global_theme", flattenSystemAdminUserUseGlobalTheme(o["use-global-theme"], d, "use_global_theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-global-theme"], "SystemAdminUser-UseGlobalTheme"); ok {
			if err = d.Set("use_global_theme", vv); err != nil {
				return fmt.Errorf("Error reading use_global_theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_global_theme: %v", err)
		}
	}

	if err = d.Set("user_theme", flattenSystemAdminUserUserTheme(o["user-theme"], d, "user_theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-theme"], "SystemAdminUser-UserTheme"); ok {
			if err = d.Set("user_theme", vv); err != nil {
				return fmt.Errorf("Error reading user_theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_theme: %v", err)
		}
	}

	if err = d.Set("user_type", flattenSystemAdminUserUserType(o["user_type"], d, "user_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["user_type"], "SystemAdminUser-UserType"); ok {
			if err = d.Set("user_type", vv); err != nil {
				return fmt.Errorf("Error reading user_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_type: %v", err)
		}
	}

	if err = d.Set("userid", flattenSystemAdminUserUserid(o["userid"], d, "userid")); err != nil {
		if vv, ok := fortiAPIPatch(o["userid"], "SystemAdminUser-Userid"); ok {
			if err = d.Set("userid", vv); err != nil {
				return fmt.Errorf("Error reading userid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading userid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("web_filter", flattenSystemAdminUserWebFilter(o["web-filter"], d, "web_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["web-filter"], "SystemAdminUser-WebFilter"); ok {
				if err = d.Set("web_filter", vv); err != nil {
					return fmt.Errorf("Error reading web_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading web_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("web_filter"); ok {
			if err = d.Set("web_filter", flattenSystemAdminUserWebFilter(o["web-filter"], d, "web_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["web-filter"], "SystemAdminUser-WebFilter"); ok {
					if err = d.Set("web_filter", vv); err != nil {
						return fmt.Errorf("Error reading web_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading web_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("wildcard", flattenSystemAdminUserWildcard(o["wildcard"], d, "wildcard")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard"], "SystemAdminUser-Wildcard"); ok {
			if err = d.Set("wildcard", vv); err != nil {
				return fmt.Errorf("Error reading wildcard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminUserFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminUserAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["adom-name"], _ = expandSystemAdminUserAdomAdomName(d, i["adom_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminUserAdomAdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserAdomAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserAdomExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["adom-name"], _ = expandSystemAdminUserAdomExcludeAdomName(d, i["adom_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminUserAdomExcludeAdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserAppFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_filter_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["app-filter-name"], _ = expandSystemAdminUserAppFilterAppFilterName(d, i["app_filter_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminUserAppFilterAppFilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserAvatar(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserChangePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserCorsAllowOrigin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "column"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["column"], _ = expandSystemAdminUserDashboardColumn(d, i["column"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diskio_content_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diskio-content-type"], _ = expandSystemAdminUserDashboardDiskioContentType(d, i["diskio_content_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diskio_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diskio-period"], _ = expandSystemAdminUserDashboardDiskioPeriod(d, i["diskio_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_rate_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-rate-period"], _ = expandSystemAdminUserDashboardLogRatePeriod(d, i["log_rate_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_rate_topn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-rate-topn"], _ = expandSystemAdminUserDashboardLogRateTopn(d, i["log_rate_topn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_rate_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-rate-type"], _ = expandSystemAdminUserDashboardLogRateType(d, i["log_rate_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "moduleid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["moduleid"], _ = expandSystemAdminUserDashboardModuleid(d, i["moduleid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemAdminUserDashboardName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "num_entries"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["num-entries"], _ = expandSystemAdminUserDashboardNumEntries(d, i["num_entries"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "refresh_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["refresh-interval"], _ = expandSystemAdminUserDashboardRefreshInterval(d, i["refresh_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "res_cpu_display"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["res-cpu-display"], _ = expandSystemAdminUserDashboardResCpuDisplay(d, i["res_cpu_display"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "res_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["res-period"], _ = expandSystemAdminUserDashboardResPeriod(d, i["res_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "res_view_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["res-view-type"], _ = expandSystemAdminUserDashboardResViewType(d, i["res_view_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemAdminUserDashboardStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tabid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tabid"], _ = expandSystemAdminUserDashboardTabid(d, i["tabid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["time-period"], _ = expandSystemAdminUserDashboardTimePeriod(d, i["time_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "widget_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["widget-type"], _ = expandSystemAdminUserDashboardWidgetType(d, i["widget_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminUserDashboardColumn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardDiskioContentType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardDiskioPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardLogRatePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardLogRateTopn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardLogRateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardModuleid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardNumEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardRefreshInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardResCpuDisplay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardResPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardResViewType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardTabid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardTimePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardWidgetType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardTabs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemAdminUserDashboardTabsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tabid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tabid"], _ = expandSystemAdminUserDashboardTabsTabid(d, i["tabid"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminUserDashboardTabsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDashboardTabsTabid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserDevGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserEmailAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserExtAuthAccprofileOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserExtAuthAdomOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserExtAuthGroupMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserFingerprint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserFirstName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserForcePasswordChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserHidden(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpsFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ips_filter_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ips-filter-name"], _ = expandSystemAdminUserIpsFilterIpsFilterName(d, i["ips_filter_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminUserIpsFilterIpsFilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpv6Trusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpv6Trusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpv6Trusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpv6Trusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpv6Trusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpv6Trusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpv6Trusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpv6Trusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpv6Trusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserIpv6Trusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserLastName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserLoginMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserMetaData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldlength"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fieldlength"], _ = expandSystemAdminUserMetaDataFieldlength(d, i["fieldlength"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fieldname"], _ = expandSystemAdminUserMetaDataFieldname(d, i["fieldname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldvalue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fieldvalue"], _ = expandSystemAdminUserMetaDataFieldvalue(d, i["fieldvalue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "importance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["importance"], _ = expandSystemAdminUserMetaDataImportance(d, i["importance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemAdminUserMetaDataStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminUserMetaDataFieldlength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserMetaDataFieldname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserMetaDataFieldvalue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserMetaDataImportance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserMetaDataStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserMobileNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserPagerNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminUserPasswordExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserPhoneNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserPolicyBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_block_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["policy-block-name"], _ = expandSystemAdminUserPolicyBlockPolicyBlockName(d, i["policy_block_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminUserPolicyBlockPolicyBlockName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserPolicyPackage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_package_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["policy-package-name"], _ = expandSystemAdminUserPolicyPackagePolicyPackageName(d, i["policy_package_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminUserPolicyPackagePolicyPackageName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserProfileid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserRpcPermit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserSshPublicKey1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminUserSshPublicKey2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminUserSshPublicKey3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminUserSubject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserTacacsPlusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserThFromProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserTh6FromProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserTrusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminUserTrusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminUserTrusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminUserTrusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminUserTrusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminUserTrusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminUserTrusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminUserTrusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminUserTrusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminUserTrusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminUserTwoFactorAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserUseGlobalTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserUserTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserUserType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserUserid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserWebFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "web_filter_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["web-filter-name"], _ = expandSystemAdminUserWebFilterWebFilterName(d, i["web_filter_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminUserWebFilterWebFilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminUserWildcard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminUser(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandSystemAdminUserAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("adom_access"); ok || d.HasChange("adom_access") {
		t, err := expandSystemAdminUserAdomAccess(d, v, "adom_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-access"] = t
		}
	}

	if v, ok := d.GetOk("adom_exclude"); ok || d.HasChange("adom_exclude") {
		t, err := expandSystemAdminUserAdomExclude(d, v, "adom_exclude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-exclude"] = t
		}
	}

	if v, ok := d.GetOk("app_filter"); ok || d.HasChange("app_filter") {
		t, err := expandSystemAdminUserAppFilter(d, v, "app_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-filter"] = t
		}
	}

	if v, ok := d.GetOk("avatar"); ok || d.HasChange("avatar") {
		t, err := expandSystemAdminUserAvatar(d, v, "avatar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["avatar"] = t
		}
	}

	if v, ok := d.GetOk("ca"); ok || d.HasChange("ca") {
		t, err := expandSystemAdminUserCa(d, v, "ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("change_password"); ok || d.HasChange("change_password") {
		t, err := expandSystemAdminUserChangePassword(d, v, "change_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-password"] = t
		}
	}

	if v, ok := d.GetOk("cors_allow_origin"); ok || d.HasChange("cors_allow_origin") {
		t, err := expandSystemAdminUserCorsAllowOrigin(d, v, "cors_allow_origin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cors-allow-origin"] = t
		}
	}

	if v, ok := d.GetOk("dashboard"); ok || d.HasChange("dashboard") {
		t, err := expandSystemAdminUserDashboard(d, v, "dashboard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dashboard"] = t
		}
	}

	if v, ok := d.GetOk("dashboard_tabs"); ok || d.HasChange("dashboard_tabs") {
		t, err := expandSystemAdminUserDashboardTabs(d, v, "dashboard_tabs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dashboard-tabs"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemAdminUserDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dev_group"); ok || d.HasChange("dev_group") {
		t, err := expandSystemAdminUserDevGroup(d, v, "dev_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-group"] = t
		}
	}

	if v, ok := d.GetOk("email_address"); ok || d.HasChange("email_address") {
		t, err := expandSystemAdminUserEmailAddress(d, v, "email_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-address"] = t
		}
	}

	if v, ok := d.GetOk("ext_auth_accprofile_override"); ok || d.HasChange("ext_auth_accprofile_override") {
		t, err := expandSystemAdminUserExtAuthAccprofileOverride(d, v, "ext_auth_accprofile_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-auth-accprofile-override"] = t
		}
	}

	if v, ok := d.GetOk("ext_auth_adom_override"); ok || d.HasChange("ext_auth_adom_override") {
		t, err := expandSystemAdminUserExtAuthAdomOverride(d, v, "ext_auth_adom_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-auth-adom-override"] = t
		}
	}

	if v, ok := d.GetOk("ext_auth_group_match"); ok || d.HasChange("ext_auth_group_match") {
		t, err := expandSystemAdminUserExtAuthGroupMatch(d, v, "ext_auth_group_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-auth-group-match"] = t
		}
	}

	if v, ok := d.GetOk("fingerprint"); ok || d.HasChange("fingerprint") {
		t, err := expandSystemAdminUserFingerprint(d, v, "fingerprint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fingerprint"] = t
		}
	}

	if v, ok := d.GetOk("first_name"); ok || d.HasChange("first_name") {
		t, err := expandSystemAdminUserFirstName(d, v, "first_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["first-name"] = t
		}
	}

	if v, ok := d.GetOk("force_password_change"); ok || d.HasChange("force_password_change") {
		t, err := expandSystemAdminUserForcePasswordChange(d, v, "force_password_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-password-change"] = t
		}
	}

	if v, ok := d.GetOk("fortiai"); ok || d.HasChange("fortiai") {
		t, err := expandSystemAdminUserFortiai(d, v, "fortiai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiai"] = t
		}
	}

	if v, ok := d.GetOk("group"); ok || d.HasChange("group") {
		t, err := expandSystemAdminUserGroup(d, v, "group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	}

	if v, ok := d.GetOk("hidden"); ok || d.HasChange("hidden") {
		t, err := expandSystemAdminUserHidden(d, v, "hidden")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hidden"] = t
		}
	}

	if v, ok := d.GetOk("ips_filter"); ok || d.HasChange("ips_filter") {
		t, err := expandSystemAdminUserIpsFilter(d, v, "ips_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-filter"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost1"); ok || d.HasChange("ipv6_trusthost1") {
		t, err := expandSystemAdminUserIpv6Trusthost1(d, v, "ipv6_trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost10"); ok || d.HasChange("ipv6_trusthost10") {
		t, err := expandSystemAdminUserIpv6Trusthost10(d, v, "ipv6_trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost2"); ok || d.HasChange("ipv6_trusthost2") {
		t, err := expandSystemAdminUserIpv6Trusthost2(d, v, "ipv6_trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost3"); ok || d.HasChange("ipv6_trusthost3") {
		t, err := expandSystemAdminUserIpv6Trusthost3(d, v, "ipv6_trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost4"); ok || d.HasChange("ipv6_trusthost4") {
		t, err := expandSystemAdminUserIpv6Trusthost4(d, v, "ipv6_trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost5"); ok || d.HasChange("ipv6_trusthost5") {
		t, err := expandSystemAdminUserIpv6Trusthost5(d, v, "ipv6_trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost6"); ok || d.HasChange("ipv6_trusthost6") {
		t, err := expandSystemAdminUserIpv6Trusthost6(d, v, "ipv6_trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost7"); ok || d.HasChange("ipv6_trusthost7") {
		t, err := expandSystemAdminUserIpv6Trusthost7(d, v, "ipv6_trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost8"); ok || d.HasChange("ipv6_trusthost8") {
		t, err := expandSystemAdminUserIpv6Trusthost8(d, v, "ipv6_trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost9"); ok || d.HasChange("ipv6_trusthost9") {
		t, err := expandSystemAdminUserIpv6Trusthost9(d, v, "ipv6_trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6_trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("last_name"); ok || d.HasChange("last_name") {
		t, err := expandSystemAdminUserLastName(d, v, "last_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-name"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandSystemAdminUserLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("login_max"); ok || d.HasChange("login_max") {
		t, err := expandSystemAdminUserLoginMax(d, v, "login_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-max"] = t
		}
	}

	if v, ok := d.GetOk("meta_data"); ok || d.HasChange("meta_data") {
		t, err := expandSystemAdminUserMetaData(d, v, "meta_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["meta-data"] = t
		}
	}

	if v, ok := d.GetOk("mobile_number"); ok || d.HasChange("mobile_number") {
		t, err := expandSystemAdminUserMobileNumber(d, v, "mobile_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mobile-number"] = t
		}
	}

	if v, ok := d.GetOk("pager_number"); ok || d.HasChange("pager_number") {
		t, err := expandSystemAdminUserPagerNumber(d, v, "pager_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pager-number"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemAdminUserPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("password_expire"); ok || d.HasChange("password_expire") {
		t, err := expandSystemAdminUserPasswordExpire(d, v, "password_expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expire"] = t
		}
	}

	if v, ok := d.GetOk("phone_number"); ok || d.HasChange("phone_number") {
		t, err := expandSystemAdminUserPhoneNumber(d, v, "phone_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phone-number"] = t
		}
	}

	if v, ok := d.GetOk("policy_block"); ok || d.HasChange("policy_block") {
		t, err := expandSystemAdminUserPolicyBlock(d, v, "policy_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-block"] = t
		}
	}

	if v, ok := d.GetOk("policy_package"); ok || d.HasChange("policy_package") {
		t, err := expandSystemAdminUserPolicyPackage(d, v, "policy_package")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-package"] = t
		}
	}

	if v, ok := d.GetOk("profileid"); ok || d.HasChange("profileid") {
		t, err := expandSystemAdminUserProfileid(d, v, "profileid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profileid"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok || d.HasChange("radius_server") {
		t, err := expandSystemAdminUserRadiusServer(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius_server"] = t
		}
	}

	if v, ok := d.GetOk("rpc_permit"); ok || d.HasChange("rpc_permit") {
		t, err := expandSystemAdminUserRpcPermit(d, v, "rpc_permit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-permit"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key1"); ok || d.HasChange("ssh_public_key1") {
		t, err := expandSystemAdminUserSshPublicKey1(d, v, "ssh_public_key1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key1"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key2"); ok || d.HasChange("ssh_public_key2") {
		t, err := expandSystemAdminUserSshPublicKey2(d, v, "ssh_public_key2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key2"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key3"); ok || d.HasChange("ssh_public_key3") {
		t, err := expandSystemAdminUserSshPublicKey3(d, v, "ssh_public_key3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key3"] = t
		}
	}

	if v, ok := d.GetOk("subject"); ok || d.HasChange("subject") {
		t, err := expandSystemAdminUserSubject(d, v, "subject")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject"] = t
		}
	}

	if v, ok := d.GetOk("tacacs_plus_server"); ok || d.HasChange("tacacs_plus_server") {
		t, err := expandSystemAdminUserTacacsPlusServer(d, v, "tacacs_plus_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tacacs-plus-server"] = t
		}
	}

	if v, ok := d.GetOk("th_from_profile"); ok || d.HasChange("th_from_profile") {
		t, err := expandSystemAdminUserThFromProfile(d, v, "th_from_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["th-from-profile"] = t
		}
	}

	if v, ok := d.GetOk("th6_from_profile"); ok || d.HasChange("th6_from_profile") {
		t, err := expandSystemAdminUserTh6FromProfile(d, v, "th6_from_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["th6-from-profile"] = t
		}
	}

	if v, ok := d.GetOk("trusthost1"); ok || d.HasChange("trusthost1") {
		t, err := expandSystemAdminUserTrusthost1(d, v, "trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("trusthost10"); ok || d.HasChange("trusthost10") {
		t, err := expandSystemAdminUserTrusthost10(d, v, "trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("trusthost2"); ok || d.HasChange("trusthost2") {
		t, err := expandSystemAdminUserTrusthost2(d, v, "trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("trusthost3"); ok || d.HasChange("trusthost3") {
		t, err := expandSystemAdminUserTrusthost3(d, v, "trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("trusthost4"); ok || d.HasChange("trusthost4") {
		t, err := expandSystemAdminUserTrusthost4(d, v, "trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("trusthost5"); ok || d.HasChange("trusthost5") {
		t, err := expandSystemAdminUserTrusthost5(d, v, "trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("trusthost6"); ok || d.HasChange("trusthost6") {
		t, err := expandSystemAdminUserTrusthost6(d, v, "trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("trusthost7"); ok || d.HasChange("trusthost7") {
		t, err := expandSystemAdminUserTrusthost7(d, v, "trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("trusthost8"); ok || d.HasChange("trusthost8") {
		t, err := expandSystemAdminUserTrusthost8(d, v, "trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("trusthost9"); ok || d.HasChange("trusthost9") {
		t, err := expandSystemAdminUserTrusthost9(d, v, "trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_auth"); ok || d.HasChange("two_factor_auth") {
		t, err := expandSystemAdminUserTwoFactorAuth(d, v, "two_factor_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-auth"] = t
		}
	}

	if v, ok := d.GetOk("use_global_theme"); ok || d.HasChange("use_global_theme") {
		t, err := expandSystemAdminUserUseGlobalTheme(d, v, "use_global_theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-global-theme"] = t
		}
	}

	if v, ok := d.GetOk("user_theme"); ok || d.HasChange("user_theme") {
		t, err := expandSystemAdminUserUserTheme(d, v, "user_theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-theme"] = t
		}
	}

	if v, ok := d.GetOk("user_type"); ok || d.HasChange("user_type") {
		t, err := expandSystemAdminUserUserType(d, v, "user_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user_type"] = t
		}
	}

	if v, ok := d.GetOk("userid"); ok || d.HasChange("userid") {
		t, err := expandSystemAdminUserUserid(d, v, "userid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["userid"] = t
		}
	}

	if v, ok := d.GetOk("web_filter"); ok || d.HasChange("web_filter") {
		t, err := expandSystemAdminUserWebFilter(d, v, "web_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok || d.HasChange("wildcard") {
		t, err := expandSystemAdminUserWildcard(d, v, "wildcard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	return &obj, nil
}
