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

func flattenFmupdateFdsSettingUserAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingFdsCltSslProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingFdsSslProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingFmtrLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingFortiguardAnycast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingFortiguardAnycastSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingLinkdLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingMaxAvIpsVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingMaxWork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverride(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenFmupdateFdsSettingPushOverrideIp(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenFmupdateFdsSettingPushOverridePort(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateFdsSettingPushOverrideStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFdsSettingPushOverrideIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverridePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClient(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "announce_ip"
	if _, ok := i["announce-ip"]; ok {
		result["announce_ip"] = flattenFmupdateFdsSettingPushOverrideToClientAnnounceIp(i["announce-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateFdsSettingPushOverrideToClientStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Port")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingSendReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingSendSetup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverride(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "servlist"
	if _, ok := i["servlist"]; ok {
		result["servlist"] = flattenFmupdateFdsSettingServerOverrideServlist(i["servlist"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateFdsSettingServerOverrideStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFdsSettingServerOverrideServlist(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFmupdateFdsSettingServerOverrideServlistId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistIp6(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := i["service-type"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistServiceType(i["service-type"], d, pre_append)
			tmp["service_type"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-ServiceType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFmupdateFdsSettingServerOverrideServlistId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingSystemSupportFaz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFct(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFdc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFgt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFml(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingSystemSupportFsw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingUmsvcLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUnregDevOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateSchedule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "day"
	if _, ok := i["day"]; ok {
		result["day"] = flattenFmupdateFdsSettingUpdateScheduleDay(i["day"], d, pre_append)
	}

	pre_append = pre + ".0." + "frequency"
	if _, ok := i["frequency"]; ok {
		result["frequency"] = flattenFmupdateFdsSettingUpdateScheduleFrequency(i["frequency"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateFdsSettingUpdateScheduleStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "time"
	if _, ok := i["time"]; ok {
		result["time"] = flattenFmupdateFdsSettingUpdateScheduleTime(i["time"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFdsSettingUpdateScheduleDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateScheduleFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateScheduleStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateScheduleTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateFdsSettingWanipQueryMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFdsSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("user_agent", flattenFmupdateFdsSettingUserAgent(o["User-Agent"], d, "user_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["User-Agent"], "FmupdateFdsSetting-UserAgent"); ok {
			if err = d.Set("user_agent", vv); err != nil {
				return fmt.Errorf("Error reading user_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent: %v", err)
		}
	}

	if err = d.Set("fds_clt_ssl_protocol", flattenFmupdateFdsSettingFdsCltSslProtocol(o["fds-clt-ssl-protocol"], d, "fds_clt_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-clt-ssl-protocol"], "FmupdateFdsSetting-FdsCltSslProtocol"); ok {
			if err = d.Set("fds_clt_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading fds_clt_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_clt_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("fds_ssl_protocol", flattenFmupdateFdsSettingFdsSslProtocol(o["fds-ssl-protocol"], d, "fds_ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-ssl-protocol"], "FmupdateFdsSetting-FdsSslProtocol"); ok {
			if err = d.Set("fds_ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading fds_ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_ssl_protocol: %v", err)
		}
	}

	if err = d.Set("fmtr_log", flattenFmupdateFdsSettingFmtrLog(o["fmtr-log"], d, "fmtr_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmtr-log"], "FmupdateFdsSetting-FmtrLog"); ok {
			if err = d.Set("fmtr_log", vv); err != nil {
				return fmt.Errorf("Error reading fmtr_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmtr_log: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast", flattenFmupdateFdsSettingFortiguardAnycast(o["fortiguard-anycast"], d, "fortiguard_anycast")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-anycast"], "FmupdateFdsSetting-FortiguardAnycast"); ok {
			if err = d.Set("fortiguard_anycast", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_anycast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_anycast: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast_source", flattenFmupdateFdsSettingFortiguardAnycastSource(o["fortiguard-anycast-source"], d, "fortiguard_anycast_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-anycast-source"], "FmupdateFdsSetting-FortiguardAnycastSource"); ok {
			if err = d.Set("fortiguard_anycast_source", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_anycast_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_anycast_source: %v", err)
		}
	}

	if err = d.Set("linkd_log", flattenFmupdateFdsSettingLinkdLog(o["linkd-log"], d, "linkd_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["linkd-log"], "FmupdateFdsSetting-LinkdLog"); ok {
			if err = d.Set("linkd_log", vv); err != nil {
				return fmt.Errorf("Error reading linkd_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading linkd_log: %v", err)
		}
	}

	if err = d.Set("max_av_ips_version", flattenFmupdateFdsSettingMaxAvIpsVersion(o["max-av-ips-version"], d, "max_av_ips_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-av-ips-version"], "FmupdateFdsSetting-MaxAvIpsVersion"); ok {
			if err = d.Set("max_av_ips_version", vv); err != nil {
				return fmt.Errorf("Error reading max_av_ips_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_av_ips_version: %v", err)
		}
	}

	if err = d.Set("max_work", flattenFmupdateFdsSettingMaxWork(o["max-work"], d, "max_work")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-work"], "FmupdateFdsSetting-MaxWork"); ok {
			if err = d.Set("max_work", vv); err != nil {
				return fmt.Errorf("Error reading max_work: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_work: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("push_override", flattenFmupdateFdsSettingPushOverride(o["push-override"], d, "push_override")); err != nil {
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
			if err = d.Set("push_override", flattenFmupdateFdsSettingPushOverride(o["push-override"], d, "push_override")); err != nil {
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
		if err = d.Set("push_override_to_client", flattenFmupdateFdsSettingPushOverrideToClient(o["push-override-to-client"], d, "push_override_to_client")); err != nil {
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
			if err = d.Set("push_override_to_client", flattenFmupdateFdsSettingPushOverrideToClient(o["push-override-to-client"], d, "push_override_to_client")); err != nil {
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

	if err = d.Set("send_report", flattenFmupdateFdsSettingSendReport(o["send_report"], d, "send_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["send_report"], "FmupdateFdsSetting-SendReport"); ok {
			if err = d.Set("send_report", vv); err != nil {
				return fmt.Errorf("Error reading send_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_report: %v", err)
		}
	}

	if err = d.Set("send_setup", flattenFmupdateFdsSettingSendSetup(o["send_setup"], d, "send_setup")); err != nil {
		if vv, ok := fortiAPIPatch(o["send_setup"], "FmupdateFdsSetting-SendSetup"); ok {
			if err = d.Set("send_setup", vv); err != nil {
				return fmt.Errorf("Error reading send_setup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_setup: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_override", flattenFmupdateFdsSettingServerOverride(o["server-override"], d, "server_override")); err != nil {
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
			if err = d.Set("server_override", flattenFmupdateFdsSettingServerOverride(o["server-override"], d, "server_override")); err != nil {
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

	if err = d.Set("system_support_faz", flattenFmupdateFdsSettingSystemSupportFaz(o["system-support-faz"], d, "system_support_faz")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-faz"], "FmupdateFdsSetting-SystemSupportFaz"); ok {
			if err = d.Set("system_support_faz", vv); err != nil {
				return fmt.Errorf("Error reading system_support_faz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_faz: %v", err)
		}
	}

	if err = d.Set("system_support_fct", flattenFmupdateFdsSettingSystemSupportFct(o["system-support-fct"], d, "system_support_fct")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fct"], "FmupdateFdsSetting-SystemSupportFct"); ok {
			if err = d.Set("system_support_fct", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fct: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fct: %v", err)
		}
	}

	if err = d.Set("system_support_fdc", flattenFmupdateFdsSettingSystemSupportFdc(o["system-support-fdc"], d, "system_support_fdc")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fdc"], "FmupdateFdsSetting-SystemSupportFdc"); ok {
			if err = d.Set("system_support_fdc", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fdc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fdc: %v", err)
		}
	}

	if err = d.Set("system_support_fgt", flattenFmupdateFdsSettingSystemSupportFgt(o["system-support-fgt"], d, "system_support_fgt")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fgt"], "FmupdateFdsSetting-SystemSupportFgt"); ok {
			if err = d.Set("system_support_fgt", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fgt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fgt: %v", err)
		}
	}

	if err = d.Set("system_support_fis", flattenFmupdateFdsSettingSystemSupportFis(o["system-support-fis"], d, "system_support_fis")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fis"], "FmupdateFdsSetting-SystemSupportFis"); ok {
			if err = d.Set("system_support_fis", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fis: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fis: %v", err)
		}
	}

	if err = d.Set("system_support_fml", flattenFmupdateFdsSettingSystemSupportFml(o["system-support-fml"], d, "system_support_fml")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fml"], "FmupdateFdsSetting-SystemSupportFml"); ok {
			if err = d.Set("system_support_fml", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fml: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fml: %v", err)
		}
	}

	if err = d.Set("system_support_fsa", flattenFmupdateFdsSettingSystemSupportFsa(o["system-support-fsa"], d, "system_support_fsa")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fsa"], "FmupdateFdsSetting-SystemSupportFsa"); ok {
			if err = d.Set("system_support_fsa", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fsa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fsa: %v", err)
		}
	}

	if err = d.Set("system_support_fts", flattenFmupdateFdsSettingSystemSupportFts(o["system-support-fts"], d, "system_support_fts")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fts"], "FmupdateFdsSetting-SystemSupportFts"); ok {
			if err = d.Set("system_support_fts", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fts: %v", err)
		}
	}

	if err = d.Set("system_support_fsw", flattenFmupdateFdsSettingSystemSupportFsw(o["system-support-fsw"], d, "system_support_fsw")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-support-fsw"], "FmupdateFdsSetting-SystemSupportFsw"); ok {
			if err = d.Set("system_support_fsw", vv); err != nil {
				return fmt.Errorf("Error reading system_support_fsw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_support_fsw: %v", err)
		}
	}

	if err = d.Set("umsvc_log", flattenFmupdateFdsSettingUmsvcLog(o["umsvc-log"], d, "umsvc_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["umsvc-log"], "FmupdateFdsSetting-UmsvcLog"); ok {
			if err = d.Set("umsvc_log", vv); err != nil {
				return fmt.Errorf("Error reading umsvc_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading umsvc_log: %v", err)
		}
	}

	if err = d.Set("unreg_dev_option", flattenFmupdateFdsSettingUnregDevOption(o["unreg-dev-option"], d, "unreg_dev_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["unreg-dev-option"], "FmupdateFdsSetting-UnregDevOption"); ok {
			if err = d.Set("unreg_dev_option", vv); err != nil {
				return fmt.Errorf("Error reading unreg_dev_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unreg_dev_option: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("update_schedule", flattenFmupdateFdsSettingUpdateSchedule(o["update-schedule"], d, "update_schedule")); err != nil {
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
			if err = d.Set("update_schedule", flattenFmupdateFdsSettingUpdateSchedule(o["update-schedule"], d, "update_schedule")); err != nil {
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

	if err = d.Set("wanip_query_mode", flattenFmupdateFdsSettingWanipQueryMode(o["wanip-query-mode"], d, "wanip_query_mode")); err != nil {
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

func expandFmupdateFdsSettingUserAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingFdsCltSslProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingFdsSslProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingFmtrLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingFortiguardAnycast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingFortiguardAnycastSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingLinkdLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingMaxAvIpsVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingMaxWork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandFmupdateFdsSettingPushOverrideIp(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port"], _ = expandFmupdateFdsSettingPushOverridePort(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateFdsSettingPushOverrideStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFmupdateFdsSettingPushOverrideIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverridePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "announce_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandFmupdateFdsSettingPushOverrideToClientAnnounceIp(d, i["announce_ip"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["announce-ip"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateFdsSettingPushOverrideToClientStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpPort(d, i["port"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingSendReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingSendSetup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "servlist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandFmupdateFdsSettingServerOverrideServlist(d, i["servlist"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["servlist"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateFdsSettingServerOverrideStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFmupdateFdsSettingServerOverrideServlist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFmupdateFdsSettingServerOverrideServlistId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFmupdateFdsSettingServerOverrideServlistIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandFmupdateFdsSettingServerOverrideServlistIp6(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFmupdateFdsSettingServerOverrideServlistPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-type"], _ = expandFmupdateFdsSettingServerOverrideServlistServiceType(d, i["service_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFmupdateFdsSettingServerOverrideServlistId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingSystemSupportFaz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFct(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFdc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFgt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFml(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingSystemSupportFsw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingUmsvcLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUnregDevOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "day"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["day"], _ = expandFmupdateFdsSettingUpdateScheduleDay(d, i["day"], pre_append)
	}
	pre_append = pre + ".0." + "frequency"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["frequency"], _ = expandFmupdateFdsSettingUpdateScheduleFrequency(d, i["frequency"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateFdsSettingUpdateScheduleStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["time"], _ = expandFmupdateFdsSettingUpdateScheduleTime(d, i["time"], pre_append)
	}

	return result, nil
}

func expandFmupdateFdsSettingUpdateScheduleDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateScheduleFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateScheduleStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateScheduleTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateFdsSettingWanipQueryMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFdsSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("user_agent"); ok || d.HasChange("user_agent") {
		t, err := expandFmupdateFdsSettingUserAgent(d, v, "user_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["User-Agent"] = t
		}
	}

	if v, ok := d.GetOk("fds_clt_ssl_protocol"); ok || d.HasChange("fds_clt_ssl_protocol") {
		t, err := expandFmupdateFdsSettingFdsCltSslProtocol(d, v, "fds_clt_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-clt-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("fds_ssl_protocol"); ok || d.HasChange("fds_ssl_protocol") {
		t, err := expandFmupdateFdsSettingFdsSslProtocol(d, v, "fds_ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("fmtr_log"); ok || d.HasChange("fmtr_log") {
		t, err := expandFmupdateFdsSettingFmtrLog(d, v, "fmtr_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmtr-log"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_anycast"); ok || d.HasChange("fortiguard_anycast") {
		t, err := expandFmupdateFdsSettingFortiguardAnycast(d, v, "fortiguard_anycast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-anycast"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_anycast_source"); ok || d.HasChange("fortiguard_anycast_source") {
		t, err := expandFmupdateFdsSettingFortiguardAnycastSource(d, v, "fortiguard_anycast_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-anycast-source"] = t
		}
	}

	if v, ok := d.GetOk("linkd_log"); ok || d.HasChange("linkd_log") {
		t, err := expandFmupdateFdsSettingLinkdLog(d, v, "linkd_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["linkd-log"] = t
		}
	}

	if v, ok := d.GetOk("max_av_ips_version"); ok || d.HasChange("max_av_ips_version") {
		t, err := expandFmupdateFdsSettingMaxAvIpsVersion(d, v, "max_av_ips_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-av-ips-version"] = t
		}
	}

	if v, ok := d.GetOk("max_work"); ok || d.HasChange("max_work") {
		t, err := expandFmupdateFdsSettingMaxWork(d, v, "max_work")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-work"] = t
		}
	}

	if v, ok := d.GetOk("push_override"); ok || d.HasChange("push_override") {
		t, err := expandFmupdateFdsSettingPushOverride(d, v, "push_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["push-override"] = t
		}
	}

	if v, ok := d.GetOk("push_override_to_client"); ok || d.HasChange("push_override_to_client") {
		t, err := expandFmupdateFdsSettingPushOverrideToClient(d, v, "push_override_to_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["push-override-to-client"] = t
		}
	}

	if v, ok := d.GetOk("send_report"); ok || d.HasChange("send_report") {
		t, err := expandFmupdateFdsSettingSendReport(d, v, "send_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send_report"] = t
		}
	}

	if v, ok := d.GetOk("send_setup"); ok || d.HasChange("send_setup") {
		t, err := expandFmupdateFdsSettingSendSetup(d, v, "send_setup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send_setup"] = t
		}
	}

	if v, ok := d.GetOk("server_override"); ok || d.HasChange("server_override") {
		t, err := expandFmupdateFdsSettingServerOverride(d, v, "server_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-override"] = t
		}
	}

	if v, ok := d.GetOk("system_support_faz"); ok || d.HasChange("system_support_faz") {
		t, err := expandFmupdateFdsSettingSystemSupportFaz(d, v, "system_support_faz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-faz"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fct"); ok || d.HasChange("system_support_fct") {
		t, err := expandFmupdateFdsSettingSystemSupportFct(d, v, "system_support_fct")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fct"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fdc"); ok || d.HasChange("system_support_fdc") {
		t, err := expandFmupdateFdsSettingSystemSupportFdc(d, v, "system_support_fdc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fdc"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fgt"); ok || d.HasChange("system_support_fgt") {
		t, err := expandFmupdateFdsSettingSystemSupportFgt(d, v, "system_support_fgt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fgt"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fis"); ok || d.HasChange("system_support_fis") {
		t, err := expandFmupdateFdsSettingSystemSupportFis(d, v, "system_support_fis")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fis"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fml"); ok || d.HasChange("system_support_fml") {
		t, err := expandFmupdateFdsSettingSystemSupportFml(d, v, "system_support_fml")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fml"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fsa"); ok || d.HasChange("system_support_fsa") {
		t, err := expandFmupdateFdsSettingSystemSupportFsa(d, v, "system_support_fsa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fsa"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fts"); ok || d.HasChange("system_support_fts") {
		t, err := expandFmupdateFdsSettingSystemSupportFts(d, v, "system_support_fts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fts"] = t
		}
	}

	if v, ok := d.GetOk("system_support_fsw"); ok || d.HasChange("system_support_fsw") {
		t, err := expandFmupdateFdsSettingSystemSupportFsw(d, v, "system_support_fsw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-support-fsw"] = t
		}
	}

	if v, ok := d.GetOk("umsvc_log"); ok || d.HasChange("umsvc_log") {
		t, err := expandFmupdateFdsSettingUmsvcLog(d, v, "umsvc_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["umsvc-log"] = t
		}
	}

	if v, ok := d.GetOk("unreg_dev_option"); ok || d.HasChange("unreg_dev_option") {
		t, err := expandFmupdateFdsSettingUnregDevOption(d, v, "unreg_dev_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unreg-dev-option"] = t
		}
	}

	if v, ok := d.GetOk("update_schedule"); ok || d.HasChange("update_schedule") {
		t, err := expandFmupdateFdsSettingUpdateSchedule(d, v, "update_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-schedule"] = t
		}
	}

	if v, ok := d.GetOk("wanip_query_mode"); ok || d.HasChange("wanip_query_mode") {
		t, err := expandFmupdateFdsSettingWanipQueryMode(d, v, "wanip_query_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanip-query-mode"] = t
		}
	}

	return &obj, nil
}
