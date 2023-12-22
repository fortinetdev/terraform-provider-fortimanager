// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateFdsSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFdsSettingUpdate,
		Read:   resourceFmupdateFdsSettingRead,
		Update: resourceFmupdateFdsSettingUpdate,
		Delete: resourceFmupdateFdsSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"user_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_clt_ssl_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_ssl_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmtr_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiguard_anycast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiguard_anycast_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linkd_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_av_ips_version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_work": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"push_override": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
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
			"push_override_to_client": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"announce_ip": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"send_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_setup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_override": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"servlist": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ip6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"service_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"system_support_faz": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"system_support_fct": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"system_support_fdc": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"system_support_fgt": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"system_support_fis": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"system_support_fml": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"system_support_fsa": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"system_support_fts": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"system_support_fsw": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"umsvc_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unreg_dev_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_schedule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"day": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"frequency": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"wanip_query_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateFdsSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateFdsSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFdsSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateFdsSetting")

	return resourceFmupdateFdsSettingRead(d, m)
}

func resourceFmupdateFdsSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteFmupdateFdsSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFdsSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFdsSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateFdsSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFdsSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSetting resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFdsSettingUserAgentFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingFdsCltSslProtocolFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingFdsSslProtocolFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingFmtrLogFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingFortiguardAnycastFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingFortiguardAnycastSourceFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingLinkdLogFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingMaxAvIpsVersionFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingMaxWorkFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideFfa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenFmupdateFdsSettingPushOverrideIpFfa(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenFmupdateFdsSettingPushOverridePortFfa(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateFdsSettingPushOverrideStatusFfa(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFdsSettingPushOverrideIpFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverridePortFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideStatusFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientFfa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "announce_ip"
	if _, ok := i["announce-ip"]; ok {
		result["announce_ip"] = flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpFfa(i["announce-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateFdsSettingPushOverrideToClientStatusFfa(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpFfa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIdFfa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIpFfa(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpPortFfa(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Port")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIdFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIpFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpPortFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientStatusFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingSendReportFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingSendSetupFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideFfa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "servlist"
	if _, ok := i["servlist"]; ok {
		result["servlist"] = flattenFmupdateFdsSettingServerOverrideServlistFfa(i["servlist"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateFdsSettingServerOverrideStatusFfa(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFdsSettingServerOverrideServlistFfa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistIdFfa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistIpFfa(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistIp6Ffa(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistPortFfa(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := i["service-type"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistServiceTypeFfa(i["service-type"], d, pre_append)
			tmp["service_type"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-ServiceType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFmupdateFdsSettingServerOverrideServlistIdFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistIpFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistIp6Ffa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistPortFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistServiceTypeFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideStatusFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingSystemSupportFazFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFctFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFdcFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFgtFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFisFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFmlFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFsaFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFtsFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFswFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingUmsvcLogFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUnregDevOptionFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateScheduleFfa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "day"
	if _, ok := i["day"]; ok {
		result["day"] = flattenFmupdateFdsSettingUpdateScheduleDayFfa(i["day"], d, pre_append)
	}

	pre_append = pre + ".0." + "frequency"
	if _, ok := i["frequency"]; ok {
		result["frequency"] = flattenFmupdateFdsSettingUpdateScheduleFrequencyFfa(i["frequency"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateFdsSettingUpdateScheduleStatusFfa(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "time"
	if _, ok := i["time"]; ok {
		result["time"] = flattenFmupdateFdsSettingUpdateScheduleTimeFfa(i["time"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFdsSettingUpdateScheduleDayFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateScheduleFrequencyFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateScheduleStatusFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateScheduleTimeFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingWanipQueryModeFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFdsSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("user_agent", flattenFmupdateFdsSettingUserAgentFfa(o["User-Agent"], d, "user_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["User-Agent"], "FmupdateFdsSetting-UserAgent"); ok {
			if err = d.Set("user_agent", vv); err != nil {
				return fmt.Errorf("Error reading user_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent: %v", err)
		}
	}

	if err = d.Set("fds_clt_ssl_protocol", flattenFmupdateFdsSettingFdsCltSslProtocolFfa(o["fds-clt-ssl-protocol"], d, "fds_clt_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-clt-ssl-protocol"], "FmupdateFdsSetting-FdsCltSslProtocol"); ok {
			if err = d.Set("fds_clt_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading fds_clt_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_clt_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("fds_ssl_protocol", flattenFmupdateFdsSettingFdsSslProtocolFfa(o["fds-ssl-protocol"], d, "fds_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-ssl-protocol"], "FmupdateFdsSetting-FdsSslProtocol"); ok {
			if err = d.Set("fds_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading fds_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("fmtr_log", flattenFmupdateFdsSettingFmtrLogFfa(o["fmtr-log"], d, "fmtr_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmtr-log"], "FmupdateFdsSetting-FmtrLog"); ok {
			if err = d.Set("fmtr_log", vv); err != nil {
				return fmt.Errorf("Error reading fmtr_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmtr_log: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast", flattenFmupdateFdsSettingFortiguardAnycastFfa(o["fortiguard-anycast"], d, "fortiguard_anycast")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-anycast"], "FmupdateFdsSetting-FortiguardAnycast"); ok {
			if err = d.Set("fortiguard_anycast", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_anycast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_anycast: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast_source", flattenFmupdateFdsSettingFortiguardAnycastSourceFfa(o["fortiguard-anycast-source"], d, "fortiguard_anycast_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-anycast-source"], "FmupdateFdsSetting-FortiguardAnycastSource"); ok {
			if err = d.Set("fortiguard_anycast_source", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_anycast_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_anycast_source: %v", err)
		}
	}

	if err = d.Set("linkd_log", flattenFmupdateFdsSettingLinkdLogFfa(o["linkd-log"], d, "linkd_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["linkd-log"], "FmupdateFdsSetting-LinkdLog"); ok {
			if err = d.Set("linkd_log", vv); err != nil {
				return fmt.Errorf("Error reading linkd_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading linkd_log: %v", err)
		}
	}

	if err = d.Set("max_av_ips_version", flattenFmupdateFdsSettingMaxAvIpsVersionFfa(o["max-av-ips-version"], d, "max_av_ips_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-av-ips-version"], "FmupdateFdsSetting-MaxAvIpsVersion"); ok {
			if err = d.Set("max_av_ips_version", vv); err != nil {
				return fmt.Errorf("Error reading max_av_ips_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_av_ips_version: %v", err)
		}
	}

	if err = d.Set("max_work", flattenFmupdateFdsSettingMaxWorkFfa(o["max-work"], d, "max_work")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-work"], "FmupdateFdsSetting-MaxWork"); ok {
			if err = d.Set("max_work", vv); err != nil {
				return fmt.Errorf("Error reading max_work: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_work: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("push_override", flattenFmupdateFdsSettingPushOverrideFfa(o["push-override"], d, "push_override")); err != nil {
			if vv, ok := fortiAPIPatch(o["push-override"], "FmupdateFdsSetting-PushOverride"); ok {
				if err = d.Set("push_override", vv); err != nil {
					return fmt.Errorf("Error reading push_override: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading push_override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("push_override"); ok {
			if err = d.Set("push_override", flattenFmupdateFdsSettingPushOverrideFfa(o["push-override"], d, "push_override")); err != nil {
				if vv, ok := fortiAPIPatch(o["push-override"], "FmupdateFdsSetting-PushOverride"); ok {
					if err = d.Set("push_override", vv); err != nil {
						return fmt.Errorf("Error reading push_override: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading push_override: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("push_override_to_client", flattenFmupdateFdsSettingPushOverrideToClientFfa(o["push-override-to-client"], d, "push_override_to_client")); err != nil {
			if vv, ok := fortiAPIPatch(o["push-override-to-client"], "FmupdateFdsSetting-PushOverrideToClient"); ok {
				if err = d.Set("push_override_to_client", vv); err != nil {
					return fmt.Errorf("Error reading push_override_to_client: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading push_override_to_client: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("push_override_to_client"); ok {
			if err = d.Set("push_override_to_client", flattenFmupdateFdsSettingPushOverrideToClientFfa(o["push-override-to-client"], d, "push_override_to_client")); err != nil {
				if vv, ok := fortiAPIPatch(o["push-override-to-client"], "FmupdateFdsSetting-PushOverrideToClient"); ok {
					if err = d.Set("push_override_to_client", vv); err != nil {
						return fmt.Errorf("Error reading push_override_to_client: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading push_override_to_client: %v", err)
				}
			}
		}
	}

	if err = d.Set("send_report", flattenFmupdateFdsSettingSendReportFfa(o["send_report"], d, "send_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["send_report"], "FmupdateFdsSetting-SendReport"); ok {
			if err = d.Set("send_report", vv); err != nil {
				return fmt.Errorf("Error reading send_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_report: %v", err)
		}
	}

	if err = d.Set("send_setup", flattenFmupdateFdsSettingSendSetupFfa(o["send_setup"], d, "send_setup")); err != nil {
		if vv, ok := fortiAPIPatch(o["send_setup"], "FmupdateFdsSetting-SendSetup"); ok {
			if err = d.Set("send_setup", vv); err != nil {
				return fmt.Errorf("Error reading send_setup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_setup: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_override", flattenFmupdateFdsSettingServerOverrideFfa(o["server-override"], d, "server_override")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-override"], "FmupdateFdsSetting-ServerOverride"); ok {
				if err = d.Set("server_override", vv); err != nil {
					return fmt.Errorf("Error reading server_override: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_override"); ok {
			if err = d.Set("server_override", flattenFmupdateFdsSettingServerOverrideFfa(o["server-override"], d, "server_override")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-override"], "FmupdateFdsSetting-ServerOverride"); ok {
					if err = d.Set("server_override", vv); err != nil {
						return fmt.Errorf("Error reading server_override: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_override: %v", err)
				}
			}
		}
	}

	if err = d.Set("system_support_faz", flattenFmupdateFdsSettingSystemSupportFazFfa(o["system-support-faz"], d, "system_support_faz")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-faz"], "FmupdateFdsSetting-SystemSupportFaz"); ok {
			if err = d.Set("system_support_faz", vv); err != nil {
				return fmt.Errorf("Error reading system_support_faz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_faz: %v", err)
		}
	}

	if err = d.Set("system_support_fct", flattenFmupdateFdsSettingSystemSupportFctFfa(o["system-support-fct"], d, "system_support_fct")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fct"], "FmupdateFdsSetting-SystemSupportFct"); ok {
			if err = d.Set("system_support_fct", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fct: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fct: %v", err)
		}
	}

	if err = d.Set("system_support_fdc", flattenFmupdateFdsSettingSystemSupportFdcFfa(o["system-support-fdc"], d, "system_support_fdc")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fdc"], "FmupdateFdsSetting-SystemSupportFdc"); ok {
			if err = d.Set("system_support_fdc", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fdc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fdc: %v", err)
		}
	}

	if err = d.Set("system_support_fgt", flattenFmupdateFdsSettingSystemSupportFgtFfa(o["system-support-fgt"], d, "system_support_fgt")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fgt"], "FmupdateFdsSetting-SystemSupportFgt"); ok {
			if err = d.Set("system_support_fgt", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fgt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fgt: %v", err)
		}
	}

	if err = d.Set("system_support_fis", flattenFmupdateFdsSettingSystemSupportFisFfa(o["system-support-fis"], d, "system_support_fis")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fis"], "FmupdateFdsSetting-SystemSupportFis"); ok {
			if err = d.Set("system_support_fis", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fis: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fis: %v", err)
		}
	}

	if err = d.Set("system_support_fml", flattenFmupdateFdsSettingSystemSupportFmlFfa(o["system-support-fml"], d, "system_support_fml")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fml"], "FmupdateFdsSetting-SystemSupportFml"); ok {
			if err = d.Set("system_support_fml", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fml: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fml: %v", err)
		}
	}

	if err = d.Set("system_support_fsa", flattenFmupdateFdsSettingSystemSupportFsaFfa(o["system-support-fsa"], d, "system_support_fsa")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fsa"], "FmupdateFdsSetting-SystemSupportFsa"); ok {
			if err = d.Set("system_support_fsa", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fsa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fsa: %v", err)
		}
	}

	if err = d.Set("system_support_fts", flattenFmupdateFdsSettingSystemSupportFtsFfa(o["system-support-fts"], d, "system_support_fts")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fts"], "FmupdateFdsSetting-SystemSupportFts"); ok {
			if err = d.Set("system_support_fts", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fts: %v", err)
		}
	}

	if err = d.Set("system_support_fsw", flattenFmupdateFdsSettingSystemSupportFswFfa(o["system-support-fsw"], d, "system_support_fsw")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fsw"], "FmupdateFdsSetting-SystemSupportFsw"); ok {
			if err = d.Set("system_support_fsw", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fsw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fsw: %v", err)
		}
	}

	if err = d.Set("umsvc_log", flattenFmupdateFdsSettingUmsvcLogFfa(o["umsvc-log"], d, "umsvc_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["umsvc-log"], "FmupdateFdsSetting-UmsvcLog"); ok {
			if err = d.Set("umsvc_log", vv); err != nil {
				return fmt.Errorf("Error reading umsvc_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading umsvc_log: %v", err)
		}
	}

	if err = d.Set("unreg_dev_option", flattenFmupdateFdsSettingUnregDevOptionFfa(o["unreg-dev-option"], d, "unreg_dev_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["unreg-dev-option"], "FmupdateFdsSetting-UnregDevOption"); ok {
			if err = d.Set("unreg_dev_option", vv); err != nil {
				return fmt.Errorf("Error reading unreg_dev_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unreg_dev_option: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("update_schedule", flattenFmupdateFdsSettingUpdateScheduleFfa(o["update-schedule"], d, "update_schedule")); err != nil {
			if vv, ok := fortiAPIPatch(o["update-schedule"], "FmupdateFdsSetting-UpdateSchedule"); ok {
				if err = d.Set("update_schedule", vv); err != nil {
					return fmt.Errorf("Error reading update_schedule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading update_schedule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("update_schedule"); ok {
			if err = d.Set("update_schedule", flattenFmupdateFdsSettingUpdateScheduleFfa(o["update-schedule"], d, "update_schedule")); err != nil {
				if vv, ok := fortiAPIPatch(o["update-schedule"], "FmupdateFdsSetting-UpdateSchedule"); ok {
					if err = d.Set("update_schedule", vv); err != nil {
						return fmt.Errorf("Error reading update_schedule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading update_schedule: %v", err)
				}
			}
		}
	}

	if err = d.Set("wanip_query_mode", flattenFmupdateFdsSettingWanipQueryModeFfa(o["wanip-query-mode"], d, "wanip_query_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanip-query-mode"], "FmupdateFdsSetting-WanipQueryMode"); ok {
			if err = d.Set("wanip_query_mode", vv); err != nil {
				return fmt.Errorf("Error reading wanip_query_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanip_query_mode: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFdsSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFdsSettingUserAgentFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingFdsCltSslProtocolFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingFdsSslProtocolFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingFmtrLogFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingFortiguardAnycastFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingFortiguardAnycastSourceFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingLinkdLogFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingMaxAvIpsVersionFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingMaxWorkFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandFmupdateFdsSettingPushOverrideIpFfa(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port"], _ = expandFmupdateFdsSettingPushOverridePortFfa(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateFdsSettingPushOverrideStatusFfa(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFmupdateFdsSettingPushOverrideIpFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverridePortFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideStatusFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "announce_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandFmupdateFdsSettingPushOverrideToClientAnnounceIpFfa(d, i["announce_ip"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["announce-ip"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateFdsSettingPushOverrideToClientStatusFfa(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIdFfa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIpFfa(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpPortFfa(d, i["port"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIdFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIpFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpPortFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientStatusFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingSendReportFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingSendSetupFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "servlist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandFmupdateFdsSettingServerOverrideServlistFfa(d, i["servlist"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["servlist"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateFdsSettingServerOverrideStatusFfa(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFmupdateFdsSettingServerOverrideServlistFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFmupdateFdsSettingServerOverrideServlistIdFfa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFmupdateFdsSettingServerOverrideServlistIpFfa(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandFmupdateFdsSettingServerOverrideServlistIp6Ffa(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFmupdateFdsSettingServerOverrideServlistPortFfa(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-type"], _ = expandFmupdateFdsSettingServerOverrideServlistServiceTypeFfa(d, i["service_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFmupdateFdsSettingServerOverrideServlistIdFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistIpFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistIp6Ffa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistPortFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistServiceTypeFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideStatusFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingSystemSupportFazFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFctFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFdcFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFgtFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFisFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFmlFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFsaFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFtsFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFswFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingUmsvcLogFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUnregDevOptionFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateScheduleFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "day"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["day"], _ = expandFmupdateFdsSettingUpdateScheduleDayFfa(d, i["day"], pre_append)
	}
	pre_append = pre + ".0." + "frequency"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["frequency"], _ = expandFmupdateFdsSettingUpdateScheduleFrequencyFfa(d, i["frequency"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateFdsSettingUpdateScheduleStatusFfa(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["time"], _ = expandFmupdateFdsSettingUpdateScheduleTimeFfa(d, i["time"], pre_append)
	}

	return result, nil
}

func expandFmupdateFdsSettingUpdateScheduleDayFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateScheduleFrequencyFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateScheduleStatusFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateScheduleTimeFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingWanipQueryModeFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFdsSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("user_agent"); ok || d.HasChange("user_agent") {
		t, err := expandFmupdateFdsSettingUserAgentFfa(d, v, "user_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["User-Agent"] = t
		}
	}

	if v, ok := d.GetOk("fds_clt_ssl_protocol"); ok || d.HasChange("fds_clt_ssl_protocol") {
		t, err := expandFmupdateFdsSettingFdsCltSslProtocolFfa(d, v, "fds_clt_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-clt-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("fds_ssl_protocol"); ok || d.HasChange("fds_ssl_protocol") {
		t, err := expandFmupdateFdsSettingFdsSslProtocolFfa(d, v, "fds_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("fmtr_log"); ok || d.HasChange("fmtr_log") {
		t, err := expandFmupdateFdsSettingFmtrLogFfa(d, v, "fmtr_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmtr-log"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_anycast"); ok || d.HasChange("fortiguard_anycast") {
		t, err := expandFmupdateFdsSettingFortiguardAnycastFfa(d, v, "fortiguard_anycast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-anycast"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_anycast_source"); ok || d.HasChange("fortiguard_anycast_source") {
		t, err := expandFmupdateFdsSettingFortiguardAnycastSourceFfa(d, v, "fortiguard_anycast_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-anycast-source"] = t
		}
	}

	if v, ok := d.GetOk("linkd_log"); ok || d.HasChange("linkd_log") {
		t, err := expandFmupdateFdsSettingLinkdLogFfa(d, v, "linkd_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["linkd-log"] = t
		}
	}

	if v, ok := d.GetOk("max_av_ips_version"); ok || d.HasChange("max_av_ips_version") {
		t, err := expandFmupdateFdsSettingMaxAvIpsVersionFfa(d, v, "max_av_ips_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-av-ips-version"] = t
		}
	}

	if v, ok := d.GetOk("max_work"); ok || d.HasChange("max_work") {
		t, err := expandFmupdateFdsSettingMaxWorkFfa(d, v, "max_work")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-work"] = t
		}
	}

	if v, ok := d.GetOk("push_override"); ok || d.HasChange("push_override") {
		t, err := expandFmupdateFdsSettingPushOverrideFfa(d, v, "push_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["push-override"] = t
		}
	}

	if v, ok := d.GetOk("push_override_to_client"); ok || d.HasChange("push_override_to_client") {
		t, err := expandFmupdateFdsSettingPushOverrideToClientFfa(d, v, "push_override_to_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["push-override-to-client"] = t
		}
	}

	if v, ok := d.GetOk("send_report"); ok || d.HasChange("send_report") {
		t, err := expandFmupdateFdsSettingSendReportFfa(d, v, "send_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send_report"] = t
		}
	}

	if v, ok := d.GetOk("send_setup"); ok || d.HasChange("send_setup") {
		t, err := expandFmupdateFdsSettingSendSetupFfa(d, v, "send_setup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send_setup"] = t
		}
	}

	if v, ok := d.GetOk("server_override"); ok || d.HasChange("server_override") {
		t, err := expandFmupdateFdsSettingServerOverrideFfa(d, v, "server_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-override"] = t
		}
	}

	if v, ok := d.GetOk("system_support_faz"); ok || d.HasChange("system_support_faz") {
		t, err := expandFmupdateFdsSettingSystemSupportFazFfa(d, v, "system_support_faz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-faz"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fct"); ok || d.HasChange("system_support_fct") {
		t, err := expandFmupdateFdsSettingSystemSupportFctFfa(d, v, "system_support_fct")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fct"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fdc"); ok || d.HasChange("system_support_fdc") {
		t, err := expandFmupdateFdsSettingSystemSupportFdcFfa(d, v, "system_support_fdc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fdc"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fgt"); ok || d.HasChange("system_support_fgt") {
		t, err := expandFmupdateFdsSettingSystemSupportFgtFfa(d, v, "system_support_fgt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fgt"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fis"); ok || d.HasChange("system_support_fis") {
		t, err := expandFmupdateFdsSettingSystemSupportFisFfa(d, v, "system_support_fis")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fis"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fml"); ok || d.HasChange("system_support_fml") {
		t, err := expandFmupdateFdsSettingSystemSupportFmlFfa(d, v, "system_support_fml")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fml"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fsa"); ok || d.HasChange("system_support_fsa") {
		t, err := expandFmupdateFdsSettingSystemSupportFsaFfa(d, v, "system_support_fsa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fsa"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fts"); ok || d.HasChange("system_support_fts") {
		t, err := expandFmupdateFdsSettingSystemSupportFtsFfa(d, v, "system_support_fts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fts"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fsw"); ok || d.HasChange("system_support_fsw") {
		t, err := expandFmupdateFdsSettingSystemSupportFswFfa(d, v, "system_support_fsw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fsw"] = t
		}
	}

	if v, ok := d.GetOk("umsvc_log"); ok || d.HasChange("umsvc_log") {
		t, err := expandFmupdateFdsSettingUmsvcLogFfa(d, v, "umsvc_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["umsvc-log"] = t
		}
	}

	if v, ok := d.GetOk("unreg_dev_option"); ok || d.HasChange("unreg_dev_option") {
		t, err := expandFmupdateFdsSettingUnregDevOptionFfa(d, v, "unreg_dev_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unreg-dev-option"] = t
		}
	}

	if v, ok := d.GetOk("update_schedule"); ok || d.HasChange("update_schedule") {
		t, err := expandFmupdateFdsSettingUpdateScheduleFfa(d, v, "update_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-schedule"] = t
		}
	}

	if v, ok := d.GetOk("wanip_query_mode"); ok || d.HasChange("wanip_query_mode") {
		t, err := expandFmupdateFdsSettingWanipQueryModeFfa(d, v, "wanip_query_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanip-query-mode"] = t
		}
	}

	return &obj, nil
}
